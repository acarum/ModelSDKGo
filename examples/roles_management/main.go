package main

import (
	"crypto/rand"
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
		fmt.Println("Usage:")
		fmt.Println("  Export pages×roles: roles_management <mpr_file_path> [output.csv]")
		fmt.Println("  Export user roles:  roles_management <mpr_file_path> --user-roles [output.csv]")
		fmt.Println("  Import:             roles_management <mpr_file_path> --import <input.csv> [--clone-from <UserRole>]")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	// Detect import mode
	for i, arg := range os.Args {
		if arg == "--import" && i+1 < len(os.Args) {
			cloneFrom := "User"
			for j, a := range os.Args {
				if a == "--clone-from" && j+1 < len(os.Args) {
					cloneFrom = os.Args[j+1]
				}
			}
			runImport(mprPath, os.Args[i+1], cloneFrom)
			return
		}
	}

	// Detect --user-roles mode
	for i, arg := range os.Args {
		if arg == "--user-roles" {
			outputPath := filepath.Join("examples", "roles_management", "user_roles_matrix.csv")
			if i+1 < len(os.Args) && !strings.HasPrefix(os.Args[i+1], "--") {
				outputPath = os.Args[i+1]
			}
			runUserRolesExport(mprPath, outputPath)
			return
		}
	}

	// ── Export mode ──────────────────────────────────────────────────────────
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

	// Step 1: collect all navigation items (pages, nanoflows, microflows)
	fmt.Println("Scanning navigation...")
	navItems := collectNavItems(db, contentsDir)
	navPages := 0
	navFlows := 0
	for _, ni := range navItems {
		if ni.ActionType == "Page" {
			navPages++
		} else {
			navFlows++
		}
	}
	fmt.Printf("  Navigation items: %d (%d pages, %d nanoflow/microflow)\n", len(navItems), navPages, navFlows)

	// Step 2: collect ALL module roles
	fmt.Println("Collecting module roles...")
	allRoles := collectAllModuleRoles(db, contentsDir)
	fmt.Printf("  Module roles: %d\n", len(allRoles))

	// Step 3: build page → allowed roles map
	fmt.Println("Collecting page access rules...")
	pageRoles := collectPageRoles(db, contentsDir)
	fmt.Printf("  Pages scanned: %d\n", len(pageRoles))

	// Step 4: resolve nav items → pageEntry structs
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

	// Sort nav items: pages first (alphabetically), then flows (alphabetically)
	sort.Slice(navItems, func(i, j int) bool {
		ti, tj := navItems[i].ActionType, navItems[j].ActionType
		if ti != tj {
			// "Page" before "Nanoflow"/"Microflow"
			if ti == "Page" {
				return true
			}
			if tj == "Page" {
				return false
			}
		}
		return navItems[i].Target < navItems[j].Target
	})

	entries := make([]pageEntry, 0, len(navItems))
	for _, ni := range navItems {
		parts := strings.SplitN(ni.Target, ".", 2)
		modName, itemName := "", ni.Target
		if len(parts) == 2 {
			modName, itemName = parts[0], parts[1]
		}

		if ni.ActionType == "Page" {
			info, found := qnToInfo[ni.Target]
			allowedRoles := map[string]bool{}
			if found {
				allowedRoles = info.AllowedRoles
			}
			entries = append(entries, pageEntry{
				ModuleName:   modName,
				PageName:     itemName,
				NavType:      "Page",
				Caption:      ni.Caption,
				AllowedRoles: allowedRoles,
			})
		} else {
			entries = append(entries, pageEntry{
				ModuleName:   modName,
				PageName:     itemName,
				NavType:      ni.ActionType,
				Caption:      ni.Caption,
				AllowedRoles: map[string]bool{},
			})
		}
	}

	// Step 5: write CSV
	writeCSV(outputPath, entries, allRoles)
	fmt.Printf("\n✓ CSV written: %s\n", outputPath)
	fmt.Printf("  %d nav items (rows) x %d roles (columns)\n", len(entries), len(allRoles))
}

// ─────────────────────────────────────────────────────────────────────────────
// Navigation collection
// ─────────────────────────────────────────────────────────────────────────────

// navItem represents a single navigation menu item with its action.
type navItem struct {
	Caption    string
	ActionType string // "Page", "Nanoflow", "Microflow", "Unknown"
	Target     string // qualified name: "Module.PageName" / "Module.NanoflowName" / etc.
}

// collectNavItems scans Navigation$NavigationDocument units and returns all
// navigation items (pages, nanoflows, microflows) reachable from any menu profile.
// Duplicates by (ActionType+Target) are eliminated.
func collectNavItems(db *sql.DB, contentsDir string) []navItem {
	seen := make(map[string]bool)
	var result []navItem
	found := false
	scanUnits(db, contentsDir, func(data map[string]interface{}) {
		if data["$Type"] != "Navigation$NavigationDocument" {
			return
		}
		found = true
		extractNavItems(data, seen, &result)
	})
	_ = found
	return result
}

// extractNavItems recursively walks a nav document node and appends found items.
func extractNavItems(data interface{}, seen map[string]bool, items *[]navItem) {
	switch v := data.(type) {
	case primitive.D:
		extractNavItems(dToMap(v), seen, items)
	case map[string]interface{}:
		// If this node has an "Action" field it is a menu item.
		if rawAction, hasAction := v["Action"]; hasAction {
			item := navItem{Caption: extractNavCaption(v)}
			if actionMap, ok := toMap(rawAction); ok {
				actionType, _ := actionMap["$Type"].(string)
				switch actionType {
				case "Pages$ShowPageClientAction", "Forms$FormAction":
					item.ActionType = "Page"
					// Try FormSettings.Form first, then direct Form field.
					if fs, ok := toMap(actionMap["FormSettings"]); ok {
						if form, ok := fs["Form"].(string); ok && strings.Contains(form, ".") {
							item.Target = form
						}
					}
					if item.Target == "" {
						if form, ok := actionMap["Form"].(string); ok && strings.Contains(form, ".") {
							item.Target = form
						}
					}
				case "Pages$CallNanoflowClientAction", "Forms$CallNanoflowClientAction":
					item.ActionType = "Nanoflow"
					if nf, ok := actionMap["Nanoflow"].(string); ok {
						item.Target = nf
					}
				case "Pages$CallMicroflowClientAction", "Forms$CallMicroflowClientAction":
					item.ActionType = "Microflow"
					if mf, ok := actionMap["Microflow"].(string); ok {
						item.Target = mf
					}
				default:
					item.ActionType = "Unknown"
				}
			}
			key := item.ActionType + "|" + item.Target
			if item.Target != "" && !seen[key] {
				seen[key] = true
				*items = append(*items, item)
			}
		}
		for _, val := range v {
			extractNavItems(val, seen, items)
		}
	case primitive.A:
		for i, item := range v {
			if i == 0 {
				continue // skip BSON count prefix
			}
			extractNavItems(item, seen, items)
		}
	case []interface{}:
		for _, item := range v {
			extractNavItems(item, seen, items)
		}
	}
}

// extractNavCaption extracts the display text from a navigation item node.
func extractNavCaption(itemMap map[string]interface{}) string {
	captionData, ok := itemMap["Caption"]
	if !ok {
		return ""
	}
	cm, ok := toMap(captionData)
	if !ok {
		return ""
	}
	if arr, ok := cm["Items"].(primitive.A); ok {
		for i, capItem := range arr {
			if i == 0 {
				continue
			}
			if capMap, ok := toMap(capItem); ok {
				if text, ok := capMap["Text"].(string); ok && text != "" {
					return text
				}
			}
		}
	}
	return ""
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
				if roleMap, ok := toMap(item); ok {
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
	case primitive.D:
		m := dToMap(v)
		extractRolesFromProjectSecurity(m, roleSet)
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
	NavType      string // "Page", "Nanoflow", "Microflow", "Unknown"
	Caption      string
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

	// Header: Caption, Module, Page, NavType, <role1>, <role2>, ...
	header := []string{"Caption", "Module", "Page", "NavType"}
	header = append(header, allRoles...)
	if err := w.Write(header); err != nil {
		log.Fatalf("csv write header: %v", err)
	}

	for _, e := range entries {
		row := []string{e.Caption, e.ModuleName, e.PageName, e.NavType}
		for _, role := range allRoles {
			if e.NavType != "Page" {
				// Nanoflow / Microflow items don't have page-level access rules
				row = append(row, "N/A")
				continue
			}
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
// IMPORT MODE
// ─────────────────────────────────────────────────────────────────────────────

// importCSVRow holds a parsed row from the import CSV.
type importCSVRow struct {
	Module     string
	Page       string
	RoleAccess map[string]bool // "Module.Role" → true (Yes) / false (No)
}

// runUserRolesExport exports a CSV with UserRoles as rows and ModuleRoles as columns.
// Each cell is "Yes" if the UserRole contains that ModuleRole, "No" otherwise.
func runUserRolesExport(mprPath, outputPath string) {
	fmt.Printf("User-roles export\n")
	fmt.Printf("  MPR: %s\n", mprPath)
	fmt.Printf("  Out: %s\n\n", outputPath)

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("sqlite3 open: %v", err)
	}
	defer db.Close()
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// --- collect all UserRoles and their assigned ModuleRoles from ProjectSecurity ---
	type userRoleEntry struct {
		Name        string
		ModuleRoles map[string]bool // "Module.Role" → true
	}
	var userRoles []userRoleEntry

	scanUnits(db, contentsDir, func(data map[string]interface{}) {
		if data["$Type"] != "Security$ProjectSecurity" {
			return
		}
		arr, _ := data["UserRoles"].(primitive.A)
		for i, item := range arr {
			if i == 0 {
				continue
			}
			rm, ok := toMap(item)
			if !ok {
				continue
			}
			name, _ := rm["Name"].(string)
			if name == "" {
				continue
			}
			entry := userRoleEntry{Name: name, ModuleRoles: make(map[string]bool)}
			mrArr, _ := rm["ModuleRoles"].(primitive.A)
			for j, mr := range mrArr {
				if j == 0 {
					continue
				}
				if s, ok := mr.(string); ok && strings.Contains(s, ".") {
					entry.ModuleRoles[s] = true
				}
			}
			userRoles = append(userRoles, entry)
		}
	})

	if len(userRoles) == 0 {
		log.Fatal("No UserRoles found in Security$ProjectSecurity")
	}
	fmt.Printf("  User roles: %d\n", len(userRoles))

	// --- collect all ModuleRoles sorted ---
	allModuleRoles := collectAllModuleRoles(db, contentsDir)
	fmt.Printf("  Module roles: %d\n", len(allModuleRoles))

	// --- write CSV ---
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		log.Fatalf("mkdir: %v", err)
	}
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("create csv: %v", err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	// header: UserRole, <ModuleRole1>, <ModuleRole2>, ...
	header := append([]string{"UserRole"}, allModuleRoles...)
	w.Write(header)

	for _, ur := range userRoles {
		row := make([]string, 1+len(allModuleRoles))
		row[0] = ur.Name
		for j, mr := range allModuleRoles {
			if ur.ModuleRoles[mr] {
				row[1+j] = "Yes"
			} else {
				row[1+j] = "No"
			}
		}
		w.Write(row)
	}

	fmt.Printf("\n✓ CSV written: %s\n", outputPath)
	fmt.Printf("  %d user roles (rows) x %d module roles (columns)\n", len(userRoles), len(allModuleRoles))
}

// runImport reads a CSV file (same format as export) and:
//  1. Identifies new Module.Role columns not yet in the MPR
//  2. Creates missing modules (Projects$Module + Security$ModuleSecurity)
//  3. Adds missing ModuleRole to existing Security$ModuleSecurity
//  4. Appends new Module.Role to the "User" UserRole in Security$ProjectSecurity
//  5. Updates AllowedModuleRoles on each Forms$Page based on Yes cells for new roles
func runImport(mprPath, csvPath, cloneFrom string) {
	fmt.Printf("Import mode\n")
	fmt.Printf("  MPR: %s\n", mprPath)
	fmt.Printf("  CSV: %s\n", csvPath)
	fmt.Printf("  Clone UserRole from: %s\n\n", cloneFrom)

	// Backup
	backupPath := mprPath + ".bak"
	if err := copyFile(mprPath, backupPath); err != nil {
		log.Fatalf("backup failed: %v", err)
	}
	fmt.Printf("✓ Backup: %s\n\n", backupPath)

	// Parse CSV
	roleColumns, rows := parseImportCSV(csvPath)
	if len(rows) == 0 {
		fmt.Println("No rows in CSV, nothing to do.")
		return
	}
	fmt.Printf("CSV: %d rows, %d role columns\n", len(rows), len(roleColumns))

	// Open MPR (read-write raw SQL only — SDK writer has bugs on MPR v2)
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("sqlite3 open: %v", err)
	}
	defer db.Close()
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Find project root UUID (the unit where ContainerID = UnitID)
	projectRootUID := findProjectRootUID(db)
	if projectRootUID == "" {
		log.Fatal("Cannot find project root unit")
	}

	// Build moduleUnitID: moduleName → UUID string (from BSON scan)
	moduleUnitID := make(map[string]string)
	scanUnitsWithUnitID(db, contentsDir, func(data map[string]interface{}, unitID string, _ string) {
		if data["$Type"] == "Projects$ModuleImpl" {
			if name, _ := data["Name"].(string); name != "" {
				moduleUnitID[name] = unitID
			}
		}
	})

	// Scan all Security$ModuleSecurity units
	existingModuleRoles := make(map[string]bool) // "Module.Role" → true
	moduleSecurityUID := make(map[string]string) // moduleName → Security unit UUID

	scanUnitsWithUnitID(db, contentsDir, func(data map[string]interface{}, unitID string, containerIDRaw string) {
		if data["$Type"] != "Security$ModuleSecurity" {
			return
		}
		modName := findModuleByTraversalFromBytes(containerIDRaw, db, func(uid string) string {
			for name, id := range moduleUnitID {
				if id == uid {
					return name
				}
			}
			return ""
		})
		if modName != "" {
			moduleSecurityUID[modName] = unitID
		}
		if arr, ok := data["ModuleRoles"].(primitive.A); ok {
			for i, item := range arr {
				if i == 0 {
					continue
				}
				if roleMap, ok := item.(map[string]interface{}); ok {
					if roleName, _ := roleMap["Name"].(string); roleName != "" && modName != "" {
						existingModuleRoles[modName+"."+roleName] = true
					}
				}
			}
		}
	})
	fmt.Printf("Existing module roles: %d\n", len(existingModuleRoles))

	// Determine new roles
	var newRoles []string
	for _, qr := range roleColumns {
		if !existingModuleRoles[qr] {
			newRoles = append(newRoles, qr)
		}
	}
	sort.Strings(newRoles)
	fmt.Printf("New roles to create: %d\n", len(newRoles))
	for _, r := range newRoles {
		fmt.Printf("  + %s\n", r)
	}

	// Group new roles by module
	newRolesByModule := make(map[string][]string)
	for _, qr := range newRoles {
		parts := strings.SplitN(qr, ".", 2)
		if len(parts) == 2 {
			newRolesByModule[parts[0]] = append(newRolesByModule[parts[0]], parts[1])
		}
	}

	// Append module roles to Security$ModuleSecurity (only for modules that already exist)
	skippedRoles := make(map[string]bool)
	for modName, roleNames := range newRolesByModule {
		// Skip if module does not exist in the MPR
		if _, exists := moduleUnitID[modName]; !exists {
			fmt.Printf("\n⚠ Skipping module '%s': not found in MPR (create the module in Mendix Studio Pro first)\n", modName)
			for _, rn := range roleNames {
				skippedRoles[modName+"."+rn] = true
			}
			continue
		}

		// Ensure Security$ModuleSecurity exists for this module
		secUID, hasSec := moduleSecurityUID[modName]
		if !hasSec {
			secUID = generateUUIDString()
			secBytes := buildEmptyModuleSecurityBSON(secUID)
			if err := writeUnitToMPR(db, contentsDir, secUID, moduleUnitID[modName], "ModuleSecurity", secBytes); err != nil {
				log.Fatalf("writeUnit ModuleSecurity %s: %v", modName, err)
			}
			moduleSecurityUID[modName] = secUID
		}

		// Load + append roles
		secData, err := readUnit(contentsDir, secUID)
		if err != nil {
			log.Fatalf("readUnit %s: %v", secUID, err)
		}
		for _, roleName := range roleNames {
			secData = appendModuleRole(secData, roleName)
			existingModuleRoles[modName+"."+roleName] = true
			fmt.Printf("  ✓ Role created: %s.%s\n", modName, roleName)
		}
		secBytes, _ := bson.Marshal(secData)
		if err := updateUnitFile(contentsDir, secUID, secBytes); err != nil {
			log.Fatalf("updateUnit %s: %v", secUID, err)
		}
	}

	// Sync ALL roles from CSV to ProjectSecurity UserRoles (idempotent: skips duplicates, skips unknown modules)
	var effectiveNewRoles []string
	for _, r := range newRoles {
		if !skippedRoles[r] {
			effectiveNewRoles = append(effectiveNewRoles, r)
		}
	}
	// Also collect all CSV roles that are NOT skipped (for idempotent UserRole sync)
	var allEffectiveRoles []string
	for _, r := range roleColumns {
		if !skippedRoles[r] {
			allEffectiveRoles = append(allEffectiveRoles, r)
		}
	}
	fmt.Print("\nUpdating ProjectSecurity UserRoles... ")
	if err := updateProjectSecurityUserRole(db, contentsDir, allEffectiveRoles); err != nil {
		fmt.Printf("warning: %v\n", err)
	} else {
		fmt.Printf("✓\n")
	}

	// Create new UserRoles for each effective ModuleRole from CSV that has no dedicated UserRole yet
	fmt.Print("Creating new UserRoles in ProjectSecurity... ")
	created, err := createUserRolesForNewModuleRoles(db, contentsDir, allEffectiveRoles, cloneFrom)
	if err != nil {
		fmt.Printf("\n\nERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ +%d new UserRoles\n", created)

	// Update page AllowedModuleRoles
	fmt.Println("\nUpdating page access rules...")
	pageUnitIDMap := make(map[string]string)
	scanUnitsWithUnitID(db, contentsDir, func(data map[string]interface{}, unitID string, _ string) {
		if data["$Type"] == "Forms$Page" {
			if name, _ := data["Name"].(string); name != "" {
				pageUnitIDMap[name] = unitID
			}
		}
	})

	newRoleSet := make(map[string]bool)
	for _, r := range effectiveNewRoles {
		newRoleSet[r] = true
	}

	pagesUpdated := 0
	for _, row := range rows {
		uid, ok := pageUnitIDMap[row.Page]
		if !ok {
			continue
		}
		pageData, err := readUnit(contentsDir, uid)
		if err != nil {
			continue
		}
		changed := false
		for qr, yes := range row.RoleAccess {
			if yes && newRoleSet[qr] {
				pageData = appendAllowedModuleRole(pageData, qr)
				changed = true
			}
		}
		if changed {
			pageBytes, _ := bson.Marshal(pageData)
			if err := updateUnitFile(contentsDir, uid, pageBytes); err != nil {
				log.Printf("  update page %s: %v", row.Page, err)
				continue
			}
			pagesUpdated++
			fmt.Printf("  ✓ %s.%s\n", row.Module, row.Page)
		}
	}
	// Clone microflow/nanoflow and entity access rights for new roles
	if len(effectiveNewRoles) > 0 {
		fmt.Print("\nCloning flow accesses (microflows + nanoflows)... ")
		flowsUpdated, flowErr := cloneFlowAccesses(db, contentsDir, effectiveNewRoles, cloneFrom)
		if flowErr != nil {
			fmt.Printf("warning: %v\n", flowErr)
		} else {
			fmt.Printf("✓ %d flows updated\n", flowsUpdated)
		}

		fmt.Print("Cloning entity access rules... ")
		dmUpdated, dmErr := cloneEntityAccessRules(db, contentsDir, effectiveNewRoles, cloneFrom)
		if dmErr != nil {
			fmt.Printf("warning: %v\n", dmErr)
		} else {
			fmt.Printf("✓ %d domain models updated\n", dmUpdated)
		}
	}

	fmt.Printf("\n✓ Import complete. Pages updated: %d\n", pagesUpdated)
}

// parseImportCSV reads the CSV and returns role columns and rows.
func parseImportCSV(csvPath string) (roleColumns []string, rows []importCSVRow) {
	f, err := os.Open(csvPath)
	if err != nil {
		log.Fatalf("open csv: %v", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("read csv: %v", err)
	}
	if len(records) < 2 {
		return
	}
	header := records[0]

	// Locate Module and Page columns by name (order-independent)
	moduleIdx, pageIdx := -1, -1
	for i, col := range header {
		switch strings.TrimSpace(col) {
		case "Module":
			moduleIdx = i
		case "Page":
			pageIdx = i
		}
	}
	if moduleIdx == -1 || pageIdx == -1 {
		log.Fatalf("parseImportCSV: CSV must have 'Module' and 'Page' columns (got header: %v)", header)
	}

	// Collect role columns (any column containing ".")
	roleColIdx := make([]int, 0)
	for i, col := range header {
		col = strings.TrimSpace(col)
		if strings.Contains(col, ".") {
			roleColumns = append(roleColumns, col)
			roleColIdx = append(roleColIdx, i)
		}
	}

	for _, rec := range records[1:] {
		if len(rec) <= max(moduleIdx, pageIdx) {
			continue
		}
		row := importCSVRow{
			Module:     strings.TrimSpace(rec[moduleIdx]),
			Page:       strings.TrimSpace(rec[pageIdx]),
			RoleAccess: make(map[string]bool),
		}
		for j, col := range roleColumns {
			idx := roleColIdx[j]
			if idx < len(rec) {
				row.RoleAccess[col] = strings.EqualFold(strings.TrimSpace(rec[idx]), "yes")
			}
		}
		rows = append(rows, row)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// appendModuleRole adds a Security$ModuleRole entry to a ModuleSecurity BSON map.
func appendModuleRole(data map[string]interface{}, roleName string) map[string]interface{} {
	arr, _ := data["ModuleRoles"].(primitive.A)
	if arr == nil {
		arr = primitive.A{int32(3)}
	}
	newRole := bson.M{
		"$ID":         newBinaryID(),
		"$Type":       "Security$ModuleRole",
		"Description": "",
		"Name":        roleName,
	}
	arr = append(arr, newRole)
	data["ModuleRoles"] = arr
	return data
}

// appendAllowedModuleRole appends a "Module.Role" string to a Page's AllowedModuleRoles.
func appendAllowedModuleRole(data map[string]interface{}, qualifiedRole string) map[string]interface{} {
	arr, _ := data["AllowedModuleRoles"].(primitive.A)
	if arr == nil {
		arr = primitive.A{int32(1)}
	}
	arr = append(arr, qualifiedRole)
	data["AllowedModuleRoles"] = arr
	return data
}

// dToMap converts a primitive.D (BSON ordered doc) to map[string]interface{}.
// bson.Unmarshal into map[string]interface{} decodes sub-documents as primitive.D,
// so all direct map[string]interface{} assertions on sub-docs must go through this.
func dToMap(d primitive.D) map[string]interface{} {
	m := make(map[string]interface{}, len(d))
	for _, e := range d {
		m[e.Key] = e.Value
	}
	return m
}

// toMap converts either map[string]interface{} or primitive.D to map[string]interface{}.
func toMap(v interface{}) (map[string]interface{}, bool) {
	switch t := v.(type) {
	case map[string]interface{}:
		return t, true
	case primitive.D:
		return dToMap(t), true
	}
	return nil, false
}

// updateProjectSecurityUserRole appends newRoles to the ModuleRoles of every UserRole
// that already contains at least one role from the same module (i.e., same prefix).
func updateProjectSecurityUserRole(db *sql.DB, contentsDir string, newRoles []string) error {
	var secUnitID string
	var secData map[string]interface{}
	scanUnitsWithUnitID(db, contentsDir, func(data map[string]interface{}, unitID string, _ string) {
		if data["$Type"] == "Security$ProjectSecurity" {
			secUnitID = unitID
			secData = data
		}
	})
	if secData == nil {
		return fmt.Errorf("Security$ProjectSecurity not found")
	}

	// Build set of module prefixes for new roles (e.g. "OpcenterEXFN_ReferenceData")
	newRoleModules := make(map[string]bool)
	for _, nr := range newRoles {
		if p := strings.SplitN(nr, ".", 2); len(p) == 2 {
			newRoleModules[p[0]] = true
		}
	}

	userRoles, _ := secData["UserRoles"].(primitive.A)
	for i, item := range userRoles {
		if i == 0 {
			continue
		}
		roleMap, ok := toMap(item)
		if !ok {
			continue
		}
		mrArr, _ := roleMap["ModuleRoles"].(primitive.A)
		if mrArr == nil {
			continue
		}
		// Check if this UserRole already has any role from one of the new role modules
		hasModule := false
		for j, mr := range mrArr {
			if j == 0 {
				continue
			}
			if s, ok := mr.(string); ok {
				if p := strings.SplitN(s, ".", 2); len(p) == 2 && newRoleModules[p[0]] {
					hasModule = true
					break
				}
			}
		}
		if !hasModule {
			continue
		}
		// Build set of existing roles in this UserRole to avoid duplicates
		existing := make(map[string]bool)
		for j, mr := range mrArr {
			if j == 0 {
				continue
			}
			if s, ok := mr.(string); ok {
				existing[s] = true
			}
		}
		for _, nr := range newRoles {
			if !existing[nr] {
				mrArr = append(mrArr, nr)
			}
		}
		roleMap["ModuleRoles"] = mrArr
		userRoles[i] = primitive.D(mapToD(roleMap))
	}
	secData["UserRoles"] = userRoles
	secBytes, err := bson.Marshal(secData)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	return updateUnitFile(contentsDir, secUnitID, secBytes)
}

// mapToD converts map[string]interface{} back to primitive.D preserving key order where possible.
func mapToD(m map[string]interface{}) primitive.D {
	d := make(primitive.D, 0, len(m))
	for k, v := range m {
		d = append(d, primitive.E{Key: k, Value: v})
	}
	return d
}

// createUserRolesForNewModuleRoles creates a new UserRole in Security$ProjectSecurity
// for each new ModuleRole, cloning the existing UserRole that already contains a role
// from the same module. The new UserRole is named after the role part (e.g. "Pippo").
// Returns the number of UserRoles actually created.
func createUserRolesForNewModuleRoles(db *sql.DB, contentsDir string, newRoles []string, cloneFrom string) (int, error) {
	var secUnitID string
	var secData map[string]interface{}
	scanUnitsWithUnitID(db, contentsDir, func(data map[string]interface{}, unitID string, _ string) {
		if data["$Type"] == "Security$ProjectSecurity" {
			secUnitID = unitID
			secData = data
		}
	})
	if secData == nil {
		return 0, fmt.Errorf("Security$ProjectSecurity not found")
	}

	userRoles, _ := secData["UserRoles"].(primitive.A)

	// Validate that the cloneFrom UserRole exists
	cloneFromIdx := -1
	existingUserRoleNames := make(map[string]bool)
	for i, item := range userRoles {
		if i == 0 {
			continue
		}
		rm, ok := toMap(item)
		if !ok {
			continue
		}
		if name, _ := rm["Name"].(string); name != "" {
			existingUserRoleNames[name] = true
			if name == cloneFrom {
				cloneFromIdx = i
			}
		}
	}
	if cloneFromIdx == -1 {
		// List available UserRole names for a helpful error message
		names := make([]string, 0, len(existingUserRoleNames))
		for n := range existingUserRoleNames {
			names = append(names, fmt.Sprintf("%q", n))
		}
		sort.Strings(names)
		return 0, fmt.Errorf("UserRole %q not found in Security$ProjectSecurity.\nAvailable UserRoles: %s\nUse --clone-from <UserRole> to specify which UserRole to clone", cloneFrom, strings.Join(names, ", "))
	}

	created := 0
	for _, qr := range newRoles {
		parts := strings.SplitN(qr, ".", 2)
		if len(parts) != 2 {
			continue
		}
		modName, roleName := parts[0], parts[1]

		// Skip if UserRole with this name already exists
		if existingUserRoleNames[roleName] {
			continue
		}

		// Clone the specified UserRole, replacing same-module roles with the new one
		sourceRM, _ := toMap(userRoles[cloneFromIdx])
		sourceMR, _ := sourceRM["ModuleRoles"].(primitive.A)

		// Build new ModuleRoles: copy all existing, replace same-module role with new one
		newMR := primitive.A{sourceMR[0]} // keep count prefix
		for j, mr := range sourceMR {
			if j == 0 {
				continue
			}
			if s, ok := mr.(string); ok && strings.HasPrefix(s, modName+".") {
				// Replace module-specific role with new one
				newMR = append(newMR, qr)
			} else {
				newMR = append(newMR, mr)
			}
		}

		newUserRole := primitive.D{
			{Key: "$ID", Value: newBinaryID()},
			{Key: "$Type", Value: "Security$UserRole"},
			{Key: "Name", Value: roleName},
			{Key: "Description", Value: ""},
			{Key: "ModuleRoles", Value: newMR},
		}
		userRoles = append(userRoles, newUserRole)
		existingUserRoleNames[roleName] = true
		fmt.Printf("\n  ✓ UserRole '%s' created (cloned from '%s')", roleName, sourceRM["Name"])
		created++
	}

	if created == 0 {
		return 0, nil
	}

	secData["UserRoles"] = userRoles
	secBytes, err := bson.Marshal(secData)
	if err != nil {
		return 0, fmt.Errorf("marshal: %w", err)
	}
	return created, updateUnitFile(contentsDir, secUnitID, secBytes)
}

// ─────────────────────────────────────────────────────────────────────────────
// Access-rights cloning
// ─────────────────────────────────────────────────────────────────────────────

// buildSourceRoleMap builds a map from newRole → sourceRole (cloneFrom in the same module).
// E.g., for newRole "OpcenterEXFN_ReferenceData.PippoTest" and cloneFrom "User",
// the sourceRole is "OpcenterEXFN_ReferenceData.User".
func buildSourceRoleMap(newRoles []string, cloneFrom string) map[string]string {
	m := make(map[string]string)
	for _, nr := range newRoles {
		parts := strings.SplitN(nr, ".", 2)
		if len(parts) == 2 {
			m[nr] = parts[0] + "." + cloneFrom
		}
	}
	return m
}

// cloneFlowAccesses copies AllowedModuleRoles access from the cloneFrom role to each new role
// for all Microflows$Microflow and Microflows$Nanoflow units.
// Returns the number of flow units updated.
func cloneFlowAccesses(db *sql.DB, contentsDir string, newRoles []string, cloneFrom string) (int, error) {
	if len(newRoles) == 0 {
		return 0, nil
	}
	sourceMap := buildSourceRoleMap(newRoles, cloneFrom)
	updated := 0
	var lastErr error

	scanUnitsWithUnitID(db, contentsDir, func(data map[string]interface{}, unitID string, _ string) {
		t, _ := data["$Type"].(string)
		if t != "Microflows$Microflow" && t != "Microflows$Nanoflow" {
			return
		}
		arr, _ := data["AllowedModuleRoles"].(primitive.A)
		if len(arr) <= 1 {
			return
		}
		existing := make(map[string]bool)
		for i, item := range arr {
			if i == 0 {
				continue
			}
			if s, ok := item.(string); ok {
				existing[s] = true
			}
		}
		changed := false
		for newRole, sourceRole := range sourceMap {
			if existing[sourceRole] && !existing[newRole] {
				arr = append(arr, newRole)
				existing[newRole] = true
				changed = true
			}
		}
		if !changed {
			return
		}
		data["AllowedModuleRoles"] = arr
		b, _ := bson.Marshal(data)
		if err := updateUnitFile(contentsDir, unitID, b); err != nil {
			lastErr = err
		} else {
			updated++
		}
	})

	return updated, lastErr
}

// cloneEntityAccessRules clones entity access rules from the cloneFrom role to each new role
// for all DomainModels$DomainModel units (entities are embedded in the domain model BSON).
// Returns the number of domain model units updated.
func cloneEntityAccessRules(db *sql.DB, contentsDir string, newRoles []string, cloneFrom string) (int, error) {
	if len(newRoles) == 0 {
		return 0, nil
	}
	sourceMap := buildSourceRoleMap(newRoles, cloneFrom)
	updated := 0
	var lastErr error

	scanUnitsWithUnitID(db, contentsDir, func(data map[string]interface{}, unitID string, _ string) {
		if data["$Type"] != "DomainModels$DomainModel" {
			return
		}
		entities, _ := data["Entities"].(primitive.A)
		if len(entities) <= 1 {
			return
		}
		dmChanged := false
		for i, entityItem := range entities {
			if i == 0 {
				continue
			}
			entityMap, ok := toMap(entityItem)
			if !ok {
				continue
			}
			accessRules, _ := entityMap["AccessRules"].(primitive.A)
			if len(accessRules) <= 1 {
				continue
			}
			entityChanged := false
			origLen := len(accessRules)
			for ri := 1; ri < origLen; ri++ {
				ruleMap, ok := toMap(accessRules[ri])
				if !ok {
					continue
				}
				amr, _ := ruleMap["AllowedModuleRoles"].(primitive.A)
				for newRole, sourceRole := range sourceMap {
					sourceFound, newFound := false, false
					for j, mr := range amr {
						if j == 0 {
							continue
						}
						if s, ok := mr.(string); ok {
							if s == sourceRole {
								sourceFound = true
							}
							if s == newRole {
								newFound = true
							}
						}
					}
					if sourceFound && !newFound {
						accessRules = append(accessRules, cloneAccessRule(ruleMap, newRole))
						entityChanged = true
					}
				}
			}
			if entityChanged {
				entityMap["AccessRules"] = accessRules
				entities[i] = entityMap
				dmChanged = true
			}
		}
		if !dmChanged {
			return
		}
		data["Entities"] = entities
		b, _ := bson.Marshal(data)
		if err := updateUnitFile(contentsDir, unitID, b); err != nil {
			lastErr = err
		} else {
			updated++
		}
	})

	return updated, lastErr
}

// cloneAccessRule creates a copy of a DomainModels$AccessRule map with a new $ID,
// the given newRole in AllowedModuleRoles, and cloned MemberAccesses with new $IDs.
func cloneAccessRule(ruleMap map[string]interface{}, newRole string) map[string]interface{} {
	clone := make(map[string]interface{}, len(ruleMap))
	for k, v := range ruleMap {
		clone[k] = v
	}
	clone["$ID"] = newBinaryID()
	clone["AllowedModuleRoles"] = primitive.A{int32(1), newRole}
	if maArr, ok := ruleMap["MemberAccesses"].(primitive.A); ok && len(maArr) > 1 {
		newMA := make(primitive.A, 0, len(maArr))
		newMA = append(newMA, maArr[0]) // keep count prefix
		for j, maItem := range maArr {
			if j == 0 {
				continue
			}
			maMap, ok := toMap(maItem)
			if !ok {
				newMA = append(newMA, maItem)
				continue
			}
			newMAMap := make(map[string]interface{}, len(maMap))
			for k, v := range maMap {
				newMAMap[k] = v
			}
			newMAMap["$ID"] = newBinaryID()
			newMA = append(newMA, newMAMap)
		}
		clone["MemberAccesses"] = newMA
	}
	return clone
}

// buildEmptyModuleSecurityBSON returns BSON bytes for an empty Security$ModuleSecurity unit.
func buildEmptyModuleSecurityBSON(unitID string) []byte {
	doc := bson.M{
		"$ID":         newBinaryIDFromUUID(unitID),
		"$Type":       "Security$ModuleSecurity",
		"ModuleRoles": primitive.A{int32(3)},
	}
	b, _ := bson.Marshal(doc)
	return b
}

// newBinaryID returns a new random primitive.Binary for Mendix $ID fields.
func newBinaryID() primitive.Binary {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("rand: %v", err)
	}
	return primitive.Binary{Subtype: 0x00, Data: b}
}

// newBinaryIDFromUUID converts a UUID string to primitive.Binary.
func newBinaryIDFromUUID(uuidStr string) primitive.Binary {
	cleaned := strings.ReplaceAll(uuidStr, "-", "")
	b, err := hex.DecodeString(cleaned)
	if err != nil || len(b) != 16 {
		return newBinaryID()
	}
	return primitive.Binary{Subtype: 0x00, Data: b}
}

// generateUUIDString returns a random UUID v4 string.
func generateUUIDString() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("rand: %v", err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// writeUnitToMPR inserts a new unit (mxunit file + SQLite row).
// containmentName is the ContainmentName value (e.g. "Modules", "ModuleSecurity", "Documents").
func writeUnitToMPR(db *sql.DB, contentsDir, unitID, containerID, containmentName string, data []byte) error {
	clean := strings.ReplaceAll(unitID, "-", "")
	dir := filepath.Join(contentsDir, clean[:2], clean[2:4])
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, unitID+".mxunit"), data, 0644); err != nil {
		return err
	}
	uidBytes := stringToWindowsGUID(unitID)
	cidBytes := stringToWindowsGUID(containerID)
	_, err := db.Exec(`INSERT INTO Unit (UnitID, ContainerID, ContainmentName, ContentsHash) VALUES (?, ?, ?, '')`,
		uidBytes, cidBytes, containmentName)
	return err
}

// findProjectRootUID returns the UUID of the Projects$Project unit (self-referencing containerID).
func findProjectRootUID(db *sql.DB) string {
	rows, err := db.Query("SELECT UnitID, ContainerID FROM Unit")
	if err != nil {
		return ""
	}
	defer rows.Close()
	for rows.Next() {
		var uid, cid []byte
		rows.Scan(&uid, &cid)
		if string(uid) == string(cid) {
			return bytesToUUID(uid)
		}
	}
	return ""
}

// buildModuleBSON returns BSON bytes for a new Projects$ModuleImpl unit.
func buildModuleBSON(unitID, moduleName string) []byte {
	doc := bson.M{
		"$ID":          newBinaryIDFromUUID(unitID),
		"$Type":        "Projects$ModuleImpl",
		"Name":         moduleName,
		"Excluded":     false,
		"FromAppStore": false,
	}
	b, _ := bson.Marshal(doc)
	return b
}

// updateUnitFile overwrites an existing mxunit file.
func updateUnitFile(contentsDir, unitID string, data []byte) error {
	clean := strings.ReplaceAll(unitID, "-", "")
	if len(clean) < 4 {
		return fmt.Errorf("invalid uid: %s", unitID)
	}
	fpath := filepath.Join(contentsDir, clean[:2], clean[2:4], unitID+".mxunit")
	return os.WriteFile(fpath, data, 0644)
}

// scanUnitsWithUnitID iterates all units providing unitID and raw containerID bytes as string.
func scanUnitsWithUnitID(db *sql.DB, contentsDir string, fn func(data map[string]interface{}, unitID string, containerIDRaw string)) {
	rows, err := db.Query("SELECT UnitID, ContainerID FROM Unit")
	if err != nil {
		log.Printf("scanUnitsWithUnitID query: %v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var idBytes, cidBytes []byte
		if err := rows.Scan(&idBytes, &cidBytes); err != nil {
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
		fn(data, uid, string(cidBytes))
	}
}

// findModuleByTraversalFromBytes traces up from a raw-bytes containerID string.
func findModuleByTraversalFromBytes(containerIDRaw string, db *sql.DB, lookupName func(uid string) string) string {
	current := []byte(containerIDRaw)
	for i := 0; i < 20; i++ {
		guid := guidToString(current)
		if name := lookupName(guid); name != "" {
			return name
		}
		var parent []byte
		err := db.QueryRow("SELECT ContainerID FROM Unit WHERE UnitID = ?", current).Scan(&parent)
		if err != nil || len(parent) == 0 {
			break
		}
		current = parent
	}
	return ""
}

// copyFile copies src to dst.
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
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
