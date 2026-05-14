// debug_access_rights dumps MicroflowAccesses from Security$ModuleSecurity
// and AccessRules from DomainModels$Entity for a given module, showing which
// roles are referenced in each entry.
// Usage: go run examples/debug_access_rights/main.go <mpr_path> <module_name>
package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: debug_access_rights <mpr_path> <module_name>")
		os.Exit(1)
	}
	mprPath := os.Args[1]
	filterModule := os.Args[2]

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var uidBytes []byte
		rows.Scan(&uidBytes)
		uid := bytesToUUID(uidBytes)
		data, err := readUnit(contentsDir, uid)
		if err != nil {
			continue
		}
		t, _ := data["$Type"].(string)

		switch t {
		case "Security$ModuleSecurity":
			// Check if it belongs to filterModule by looking at module name in ModuleRoles
			moduleName := inferModuleName(data)
			if moduleName != filterModule {
				continue
			}
			fmt.Printf("\n=== Security$ModuleSecurity (unit: %s) module: %s ===\n", uid, moduleName)

			// ModuleRoles
			if arr, ok := data["ModuleRoles"].(primitive.A); ok {
				fmt.Printf("  ModuleRoles (%d):\n", len(arr)-1)
				for i, item := range arr {
					if i == 0 {
						continue
					}
					rm := toMap(item)
					if rm == nil {
						continue
					}
					fmt.Printf("    [%d] Name=%q $ID=%v\n", i, rm["Name"], rm["$ID"])
				}
			}

			// MicroflowAccesses
			dumpAccessArray(data, "MicroflowAccesses")
			// NanoflowAccesses
			dumpAccessArray(data, "NanoflowAccesses")

		case "DomainModels$DomainModel":
			// Find entities and their access rules
			entities, _ := data["Entities"].(primitive.A)
			for i, item := range entities {
				if i == 0 {
					continue
				}
				e := toMap(item)
				if e == nil {
					continue
				}
				eName, _ := e["Name"].(string)
				accessRules, _ := e["AccessRules"].(primitive.A)
				if len(accessRules) <= 1 {
					continue
				}
				// Check if any rule references filterModule
				hasModule := false
				for j, ar := range accessRules {
					if j == 0 {
						continue
					}
					arm := toMap(ar)
					if arm == nil {
						continue
					}
					if mrArr, ok := arm["ModuleRoles"].(primitive.A); ok {
						for k, mr := range mrArr {
							if k == 0 {
								continue
							}
							if s, ok := mr.(string); ok && strings.HasPrefix(s, filterModule+".") {
								hasModule = true
							}
						}
					}
				}
				if !hasModule {
					continue
				}
				fmt.Printf("\n=== Entity %s (in DomainModel unit: %s) ===\n", eName, uid)
				for j, ar := range accessRules {
					if j == 0 {
						continue
					}
					arm := toMap(ar)
					if arm == nil {
						continue
					}
					mrArr, _ := arm["ModuleRoles"].(primitive.A)
					roles := []string{}
					for k, mr := range mrArr {
						if k == 0 {
							continue
						}
						if s, ok := mr.(string); ok {
							roles = append(roles, s)
						}
					}
					fmt.Printf("  AccessRule[%d] ModuleRoles=%v AllowedCreate=%v AllowedDelete=%v\n",
						j, roles, arm["AllowedCreate"], arm["AllowedDelete"])
					// MemberAccesses
					if ma, ok := arm["MemberAccesses"].(primitive.A); ok {
						for k, maItem := range ma {
							if k == 0 {
								continue
							}
							mam := toMap(maItem)
							if mam == nil {
								continue
							}
							fmt.Printf("    MemberAccess: Attr=%v AccessRights=%v\n", mam["Attribute"], mam["AccessRights"])
						}
					}
				}
			}
		}
	}
}

func inferModuleName(data map[string]interface{}) string {
	if arr, ok := data["ModuleRoles"].(primitive.A); ok {
		for i, item := range arr {
			if i == 0 {
				continue
			}
			rm := toMap(item)
			if rm == nil {
				continue
			}
			if name, _ := rm["Name"].(string); name != "" {
				// The qualified role name is stored in MicroflowAccesses as "Module.Role"
				// We need to find it another way — look at MicroflowAccesses
				break
			}
		}
	}
	// Try MicroflowAccesses to find module prefix
	if arr, ok := data["MicroflowAccesses"].(primitive.A); ok {
		for i, item := range arr {
			if i == 0 {
				continue
			}
			rm := toMap(item)
			if rm == nil {
				continue
			}
			if mf, _ := rm["Microflow"].(string); strings.Contains(mf, ".") {
				return strings.SplitN(mf, ".", 2)[0]
			}
		}
	}
	// Fallback: look at ModuleRoles[1].Name and find its module from qualified name
	return "<unknown>"
}

func dumpAccessArray(data map[string]interface{}, field string) {
	arr, ok := data[field].(primitive.A)
	if !ok || len(arr) <= 1 {
		fmt.Printf("  %s: (none)\n", field)
		return
	}
	fmt.Printf("  %s (%d):\n", field, len(arr)-1)
	for i, item := range arr {
		if i == 0 {
			continue
		}
		rm := toMap(item)
		if rm == nil {
			fmt.Printf("    [%d] <non-map: %T>\n", i, item)
			continue
		}
		mfName := ""
		if v, ok := rm["Microflow"].(string); ok {
			mfName = v
		}
		if v, ok := rm["Nanoflow"].(string); ok {
			mfName = v
		}
		role := ""
		if v, ok := rm["ModuleRole"].(string); ok {
			role = v
		}
		allowed := rm["Allowed"]
		fmt.Printf("    [%d] Flow=%-50s Role=%-40s Allowed=%v\n", i, mfName, role, allowed)
	}
}

func toMap(v interface{}) map[string]interface{} {
	switch t := v.(type) {
	case map[string]interface{}:
		return t
	case primitive.D:
		m := make(map[string]interface{}, len(t))
		for _, e := range t {
			m[e.Key] = e.Value
		}
		return m
	}
	return nil
}

func readUnit(contentsDir, uid string) (map[string]interface{}, error) {
	clean := strings.ReplaceAll(uid, "-", "")
	if len(clean) < 4 {
		return nil, fmt.Errorf("short uid")
	}
	fpath := filepath.Join(contentsDir, clean[:2], clean[2:4], uid+".mxunit")
	raw, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := bson.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func bytesToUUID(b []byte) string {
	if len(b) != 16 {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString([]byte{b[3], b[2], b[1], b[0]}),
		hex.EncodeToString([]byte{b[5], b[4]}),
		hex.EncodeToString([]byte{b[7], b[6]}),
		hex.EncodeToString(b[8:10]),
		hex.EncodeToString(b[10:16]))
}
