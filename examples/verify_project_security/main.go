// verify_project_security reads Security$ProjectSecurity from an MPR and prints
// each UserRole with its assigned ModuleRoles.
// Usage: go run examples/verify_project_security/main.go <mpr_path> [role_to_check]
// Exit code 0 = role found in at least one UserRole (or no role given), 1 = not found / error.
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
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: verify_project_security <mpr_path> [role_to_check]")
		os.Exit(1)
	}
	mprPath := os.Args[1]
	checkRole := ""
	if len(os.Args) >= 3 {
		checkRole = os.Args[2]
	}

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Find Security$ProjectSecurity unit
	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	foundRole := false
	for rows.Next() {
		var uidBytes []byte
		rows.Scan(&uidBytes)
		uid := bytesToUUID(uidBytes)
		data, err := readUnit(contentsDir, uid)
		if err != nil || data["$Type"] != "Security$ProjectSecurity" {
			continue
		}
		fmt.Printf("Security$ProjectSecurity (unit: %s)\n", uid)
		userRoles, _ := data["UserRoles"].(primitive.A)
		for i, item := range userRoles {
			if i == 0 {
				continue
			}
			rm := toMap(item)
			if rm == nil {
				fmt.Printf("  [%d] <non-map type: %T>\n", i, item)
				continue
			}
			name, _ := rm["Name"].(string)
			mrArr, _ := rm["ModuleRoles"].(primitive.A)
			fmt.Printf("  UserRole[%d]: %q\n", i, name)
			for j, mr := range mrArr {
				if j == 0 {
					continue
				}
				if s, ok := mr.(string); ok {
					fmt.Printf("    - %s\n", s)
					if checkRole != "" && s == checkRole {
						foundRole = true
					}
				}
			}
		}
		break
	}

	if checkRole != "" {
		if foundRole {
			fmt.Printf("\n✓ Role '%s' IS present in ProjectSecurity\n", checkRole)
			os.Exit(0)
		} else {
			fmt.Printf("\n✗ Role '%s' NOT found in ProjectSecurity\n", checkRole)
			os.Exit(1)
		}
	}
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
