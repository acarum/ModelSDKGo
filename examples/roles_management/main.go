package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/anthropics/modelsdk-go"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: roles_management <mpr_file_path> [output.csv]")
		fmt.Println("Example: roles_management MyApp.mpr")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	outputPath := filepath.Join("examples", "roles_management", "roles_matrix.csv")
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		log.Fatalf("mkdir: %v", err)
	}

	// Build module map via modelsdk (moduleUnitID → moduleName)
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		log.Fatalf("modelsdk.Open: %v", err)
	}
	defer reader.Close()
	fmt.Printf("Opened: %s\n", reader.Path())

	modules, err := reader.ListModules()
	if err != nil {
		log.Fatalf("ListModules: %v", err)
	}
	moduleMap := make(map[string]string) // moduleUnitID → moduleName
	for _, m := range modules {
		moduleMap[string(m.ID)] = m.Name
	}

	// Open raw SQLite for BSON unit scanning
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("sqlite3 open: %v", err)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Step 1: collect all page qualified names reachable from navigation
	fmt.Println("Scanning navigation...")
	navPages := collectNavPages(db, contentsDir)
	fmt.Printf("  Navigation pages: %d\n", len(navPages))

	// Step 2: collect ALL module roles
	fmt.Println("Collecting module roles...")
	allRoles := collectAllModuleRoles(db, contentsDir)
	fmt.Printf("  Module roles: %d\n", len(allRoles))

	// Step 3: build page → allowed roles map
	fmt.Println("Collecting page access rules...")
	pageRoles := collectPageRoles(db, contentsDir)
	fmt.Printf("  Pages scanned: %d\n", len(pageRoles))

	// Step 4: resolve nav pages → PageEntry structs
	// Build lookup: "Module.PageName" → pageRoleInfo using moduleMap
	pageNameToModule := make(map[string]string) // pageName → moduleName
	pagesFromSDK, err := reader.ListPages()
	if err != nil {
		log.Fatalf("ListPages: %v", err)
	}
	for _, p := range pagesFromSDK {
		modName := findModuleByTraversal(string(p.ID), db, moduleMap)
		if modName != "" {
			pageNameToModule[p.Name] = modName
		}
	}
	fmt.Printf("  Pages mapped to modules: %d / %d\n", len(pageNameToModule), len(pagesFromSDK))

	qnToInfo := make(map[string]pageRoleInfo)
	for pageName, info := range pageRoles {
		modName := pageNameToModule[pageName]
		if modName == "" {
			continue
		}
		qn := modName + "." + pageName
		qnToInfo[qn] = info
	}

	navList := make([]string, 0, len(navPages))
	for qn := range navPages {
		navList = append(navList, qn)
	}
	sort.Strings(navList)

	entries := make([]pageEntry, 0, len(navList))
	for _, qn := range navList {
		parts := strings.SplitN(qn, ".", 2)
		modName, pgName := "", qn
		if len(parts) == 2 {
			modName, pgName = parts[0], parts[1]
		}
		info, found := qnToInfo[qn]
		if !found {
			entries = append(entries, pageEntry{
				ModuleName:   modName,
				PageName:     pgName,
				AllowedRoles: map[string]bool{},
			})
			continue
		}
		entries = append(entries, pageEntry{
			ModuleName:   modName,
			PageName:     pgName,
			AllowedRoles: info.AllowedRoles,
		})
	}

	// Step 5: write CSV
	writeCSV(outputPath, entries, allRoles)
	fmt.Printf("\n✓ CSV written: %s\n", outputPath)
	fmt.Printf("  %d pages (rows) x %d roles (columns)\n", len(entries), len(allRoles))
}

// ─────────────────────────────────────────────────────────────────────────────
// Navigation collection
// ─────────────────────────────────────────────────────────────────────────────

// collectNavPages scans Navigation$NavigationDocument units and returns
// a set of all page qualified names ("Module.PageName") reachable from any menu.
func collectNavPages(db *sql.DB, contentsDir string) map[string]bool {
	result := make(map[string]bool)
	scanUnits(db, contentsDir, func(data map[string]interface{}) {
		if data["$Type"] != "Navigation$NavigationDocument" {
			return
		}
		findFormRefs(data, result)
	})
	return result
}

// findFormRefs recursively extracts all "Module.PageName" strings from Form /
// FormSettings.Form fields anywhere in the subtree.
func findFormRefs(data interface{}, refs map[string]bool) {
	switch v := data.(type) {
	case map[string]interface{}:
		if form, ok := v["Form"].(string); ok && strings.Contains(form, ".") {
			refs[form] = true
		}
		if fs, ok := v["FormSettings"].(map[string]interface{}); ok {
			if form, ok := fs["Form"].(string); ok && strings.Contains(form, ".") {
				refs[form] = true
			}
		}
		for _, val := range v {
			findFormRefs(val, refs)
		}
	case primitive.A:
		for i, item := range v {
			if i == 0 {
				continue // skip count prefix
			}
			findFormRefs(item, refs)
		}
	case []interface{}:
		for _, item := range v {
			findFormRefs(item, refs)
		}
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Module role collection
// ─────────────────────────────────────────────────────────────────────────────

// collectAllModuleRoles returns a sorted list of all "Module.Role" strings from:
//  1. Security$ProjectSecurity.UserRoles[].ModuleRoles (primary)
//  2. All Forms$Page.AllowedModuleRoles (fallback when security not found)
func collectAllModuleRoles(db *sql.DB, contentsDir string) []string {
	roleSet := make(map[string]bool)

	scanUnits(db, contentsDir, func(data map[string]interface{}) {
		if data["$Type"] == "Security$ProjectSecurity" {
			extractRolesFromProjectSecurity(data, roleSet)
		}
	})

	if len(roleSet) == 0 {
		scanUnits(db, contentsDir, func(data map[string]interface{}) {
			if data["$Type"] != "Forms$Page" {
				return
			}
			if arr, ok := data["AllowedModuleRoles"].(primitive.A); ok {
				for i, item := range arr {
					if i == 0 {
						continue
					}
					if s, ok := item.(string); ok && strings.Contains(s, ".") {
						roleSet[s] = true
					}
				}
			}
		})
	}

	roles := make([]string, 0, len(roleSet))
	for r := range roleSet {
		roles = append(roles, r)
	}
	sort.Strings(roles)
	return roles
}

func extractRolesFromProjectSecurity(data interface{}, roleSet map[string]bool) {
	switch v := data.(type) {
	case map[string]interface{}:
		if arr, ok := v["UserRoles"].(primitive.A); ok {
			for i, item := range arr {
				if i == 0 {
					continue
				}
				if roleMap, ok := item.(map[string]interface{}); ok {
					if mrArr, ok := roleMap["ModuleRoles"].(primitive.A); ok {
						for j, mr := range mrArr {
							if j == 0 {
								continue
							}
							if s, ok := mr.(string); ok && strings.Contains(s, ".") {
								roleSet[s] = true
							}
						}
					}
				}
			}
		}
		for _, val := range v {
			extractRolesFromProjectSecurity(val, roleSet)
		}
	case primitive.A:
		for i, item := range v {
			if i == 0 {
				continue
			}
			extractRolesFromProjectSecurity(item, roleSet)
		}
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Page access collection
// ─────────────────────────────────────────────────────────────────────────────

type pageRoleInfo struct {
	PageName     string
	AllowedRoles map[string]bool // "Module.Role" → true; empty = public
}

// collectPageRoles scans all Forms$Page units and returns map[pageName] → pageRoleInfo
func collectPageRoles(db *sql.DB, contentsDir string) map[string]pageRoleInfo {
	pageRoles := make(map[string]pageRoleInfo)

	scanUnits(db, contentsDir, func(data map[string]interface{}) {
		if data["$Type"] != "Forms$Page" {
			return
		}
		pageName, _ := data["Name"].(string)
		if pageName == "" {
			return
		}

		allowedRoles := make(map[string]bool)
		if arr, ok := data["AllowedModuleRoles"].(primitive.A); ok {
			for i, item := range arr {
				if i == 0 {
					continue
				}
				if s, ok := item.(string); ok && strings.Contains(s, ".") {
					allowedRoles[s] = true
				}
			}
		}

		pageRoles[pageName] = pageRoleInfo{PageName: pageName, AllowedRoles: allowedRoles}
	})

	return pageRoles
}

// ─────────────────────────────────────────────────────────────────────────────
// CSV output
// ─────────────────────────────────────────────────────────────────────────────

type pageEntry struct {
	ModuleName   string
	PageName     string
	AllowedRoles map[string]bool
}

func writeCSV(outputPath string, entries []pageEntry, allRoles []string) {
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("create csv: %v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// Header: Module, Page, <role1>, <role2>, ...
	header := []string{"Module", "Page"}
	header = append(header, allRoles...)
	if err := w.Write(header); err != nil {
		log.Fatalf("csv write header: %v", err)
	}

	for _, e := range entries {
		row := []string{e.ModuleName, e.PageName}
		for _, role := range allRoles {
			switch {
			case len(e.AllowedRoles) == 0:
				row = append(row, "PUBLIC")
			case e.AllowedRoles[role]:
				row = append(row, "Yes")
			default:
				row = append(row, "No")
			}
		}
		if err := w.Write(row); err != nil {
			log.Fatalf("csv write row: %v", err)
		}
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Low-level MPR unit scanning helpers
// ─────────────────────────────────────────────────────────────────────────────

func scanUnits(db *sql.DB, contentsDir string, fn func(data map[string]interface{})) {
	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		log.Printf("query units: %v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var idBytes []byte
		if err := rows.Scan(&idBytes); err != nil {
			continue
		}
		uid := bytesToUUID(idBytes)
		if uid == "" {
			continue
		}
		data, err := readUnit(contentsDir, uid)
		if err != nil {
			continue
		}
		fn(data)
	}
}

func scanUnitsWithContainer(db *sql.DB, contentsDir string, fn func(data map[string]interface{}, containerID string)) {
	rows, err := db.Query("SELECT UnitID, ContainerID FROM Unit")
	if err != nil {
		log.Printf("query units with container: %v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var idBytes, containerBytes []byte
		if err := rows.Scan(&idBytes, &containerBytes); err != nil {
			continue
		}
		uid := bytesToUUID(idBytes)
		if uid == "" {
			continue
		}
		// Use raw bytes string as containerID key — matches moduleMap keys built via string(m.ID)
		containerID := string(containerBytes)
		data, err := readUnit(contentsDir, uid)
		if err != nil {
			continue
		}
		fn(data, containerID)
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

// bytesToUUID converts a 16-byte SQL BLOB to UUID string (Windows GUID byte order).
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

// findModuleByTraversal traces up the unit hierarchy to find the parent module name.
// pageID is a UUID string. moduleMap maps UUID string → module name.
func findModuleByTraversal(pageID string, db *sql.DB, moduleMap map[string]string) string {
	guidBytes := stringToWindowsGUID(pageID)
	if guidBytes == nil {
		return ""
	}
	currentID := guidBytes
	for i := 0; i < 20; i++ {
		var parentID []byte
		err := db.QueryRow("SELECT ContainerID FROM Unit WHERE UnitID = ?", currentID).Scan(&parentID)
		if err != nil || len(parentID) == 0 {
			break
		}
		parentGUID := guidToString(parentID)
		if moduleName, found := moduleMap[parentGUID]; found {
			return moduleName
		}
		currentID = parentID
	}
	return ""
}

// guidToString converts Windows GUID bytes (16 bytes, little-endian first 3 groups) to UUID string.
func guidToString(b []byte) string {
	if len(b) != 16 {
		return ""
	}
	cp := make([]byte, 16)
	copy(cp, b)
	cp[0], cp[1], cp[2], cp[3] = cp[3], cp[2], cp[1], cp[0]
	cp[4], cp[5] = cp[5], cp[4]
	cp[6], cp[7] = cp[7], cp[6]
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		cp[0], cp[1], cp[2], cp[3], cp[4], cp[5], cp[6], cp[7],
		cp[8], cp[9], cp[10], cp[11], cp[12], cp[13], cp[14], cp[15])
}

// stringToWindowsGUID converts UUID string to Windows GUID bytes.
func stringToWindowsGUID(uuidStr string) []byte {
	cleaned := strings.ReplaceAll(uuidStr, "-", "")
	if len(cleaned) != 32 {
		return nil
	}
	b, err := hex.DecodeString(cleaned)
	if err != nil || len(b) != 16 {
		return nil
	}
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	return b
}
