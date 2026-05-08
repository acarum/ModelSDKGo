package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"

	modelsdk "github.com/anthropics/modelsdk-go"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: debug_microflow_calls.exe <path-to-mpr>")
	}

	mprPath := os.Args[1]

	fmt.Printf("🔍 Opening MPR: %s\n", mprPath)

	// Open MPR with modelsdk
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		log.Fatalf("Failed to open MPR: %v", err)
	}
	defer reader.Close()

	// Open database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Show schema of Unit table
	fmt.Printf("\n📋 Schema of Unit table:\n")
	schemaRows, err := db.Query("PRAGMA table_info(Unit)")
	if err != nil {
		log.Fatalf("Schema query failed: %v", err)
	}
	
	fmt.Printf("  Columns:\n")
	for schemaRows.Next() {
		var cid int
		var name, typ string
		var notnull, dflt_value, pk interface{}
		if err := schemaRows.Scan(&cid, &name, &typ, &notnull, &dflt_value, &pk); err == nil {
			fmt.Printf("    %d. %s (%s)\n", cid, name, typ)
		}
	}
	schemaRows.Close()

	// List ALL units using correct column names
	fmt.Printf("\n📋 Listing first 30 units:\n")
	listQuery := "SELECT UnitID, ContainmentName FROM Unit LIMIT 30"
	listRows, err := db.Query(listQuery)
	if err != nil {
		log.Fatalf("List query failed: %v", err)
	}
	
	count := 0
	for listRows.Next() {
		var unitIDBytes []byte
		var containmentName string
		if err := listRows.Scan(&unitIDBytes, &containmentName); err == nil {
			unitIDStr := guidToString(unitIDBytes)
			fmt.Printf("  %d. %s (ID: %s)\n", count+1, containmentName, unitIDStr[:8]+"...")
			count++
		}
	}
	listRows.Close()

	os.Exit(0) // Exit here to see the schema

	// Find a microflow from Section 2 that should exist
	microflowName := "SUB_HideAndUnhideCommand"
	moduleName := "OpcenterEXFN_ReferenceData_Connector"

	fmt.Printf("\n🔍 Searching for %s.%s (using Name field)...\n", moduleName, microflowName)

	query := "SELECT UnitID, ContainerID FROM Unit WHERE Name = ?"
	rows, err := db.Query(query, microflowName)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	var unitID string
	found := false
	var foundModule string
	for rows.Next() {
		var candidateID, containerID string
		if err := rows.Scan(&candidateID, &containerID); err != nil {
			continue
		}

		// Check module
		candidateModule := findModuleByTraversal(candidateID, db)
		fmt.Printf("  Candidate: UnitID=%s, Module=%s\n", candidateID, candidateModule)
		
		// Accept first match regardless of module
		if !found {
			unitID = candidateID
			foundModule = candidateModule
			found = true
		}
		
		// But prefer exact module match
		if candidateModule == moduleName {
			unitID = candidateID
			foundModule = candidateModule
			fmt.Printf("  ✅ Exact module match! Using this one.\n")
			break
		}
	}

	if !found {
		log.Fatalf("❌ Microflow not found!")
	}

	fmt.Printf("\n✅ Found microflow: UnitID=%s, Module=%s\n", unitID, foundModule)

	// Get contents directory path
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")
	fmt.Printf("📁 Contents dir: %s\n", contentsDir)

	// Load BSON content
	data, err := loadUnitContents(contentsDir, unitID)
	if err != nil {
		log.Fatalf("Failed to load unit: %v", err)
	}

	fmt.Printf("\n📄 BSON Content (top-level keys):\n")
	for key := range data {
		fmt.Printf("  - %s\n", key)
	}

	// Search for MicroflowCall patterns
	fmt.Printf("\n🔎 Searching for microflow calls...\n")
	calls := findMicroflowCalls(data, 0, 10)
	
	if len(calls) == 0 {
		fmt.Printf("\n❌ No microflow calls found!\n")
		fmt.Printf("\n📋 Full BSON structure (first 150 lines):\n")
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		lines := strings.Split(string(jsonData), "\n")
		for i, line := range lines {
			if i >= 150 {
				fmt.Printf("... (truncated at line %d)\n", i)
				break
			}
			fmt.Println(line)
		}
	} else {
		fmt.Printf("\n✅ Found %d microflow calls:\n", len(calls))
		for _, call := range calls {
			fmt.Printf("  - %s\n", call)
		}
	}
}

func loadUnitContents(contentsDir string, unitID string) (map[string]interface{}, error) {
	// Remove dashes from UUID for directory structure
	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid unit ID: %s", unitID)
	}

	// Build path: mprcontents/{first2}/{next2}/{uuid}.mxunit
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	rawData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read unit file: %w", err)
	}

	var content map[string]interface{}
	if err := bson.Unmarshal(rawData, &content); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	return content, nil
}

func findMicroflowCalls(obj interface{}, depth int, maxDepth int) []string {
	if depth > maxDepth {
		return nil
	}

	var calls []string
	seen := make(map[string]bool)

	var extract func(o interface{}, d int)
	extract = func(o interface{}, d int) {
		if d > maxDepth {
			return
		}

		switch v := o.(type) {
		case map[string]interface{}:
			// Check for microflow call
			if typeStr, ok := v["$Type"].(string); ok {
				if strings.Contains(typeStr, "MicroflowCall") || strings.Contains(typeStr, "NanoflowCall") {
					fmt.Printf("    [Depth %d] Found $Type: %s\n", d, typeStr)
					// Print all fields
					for key, val := range v {
						fmt.Printf("      %s: %v (type: %T)\n", key, val, val)
					}
					
					// Try to extract microflow name from various fields
					for key, val := range v {
						if strings.Contains(strings.ToLower(key), "microflow") || 
						   strings.Contains(strings.ToLower(key), "nanoflow") {
							if strVal, ok := val.(string); ok && strVal != "" && strVal != "<parameter>" {
								if !seen[strVal] {
									calls = append(calls, strVal)
									seen[strVal] = true
								}
							}
						}
					}
				}
			}

			// Recurse
			for _, val := range v {
				extract(val, d+1)
			}
		case []interface{}:
			for _, item := range v {
				extract(item, d+1)
			}
		}
	}

	extract(obj, depth)
	return calls
}

func findModuleByTraversal(unitID string, db *sql.DB) string {
	currentID := unitID
	for i := 0; i < 20; i++ {
		var containerID sql.NullString
		var typeName string
		
		query := "SELECT ContainerID, Type FROM Unit WHERE UnitID = ?"
		err := db.QueryRow(query, currentID).Scan(&containerID, &typeName)
		if err != nil {
			return ""
		}

		if typeName == "Projects$Module" {
			var name string
			if err := db.QueryRow("SELECT Name FROM Unit WHERE UnitID = ?", currentID).Scan(&name); err == nil {
				return name
			}
			return ""
		}

		if !containerID.Valid {
			break
		}
		currentID = containerID.String
	}
	return ""
}

func guidToString(guidBytes []byte) string {
	if len(guidBytes) != 16 {
		return ""
	}
	// Reverse the Windows GUID encoding
	b := make([]byte, 16)
	copy(b, guidBytes)
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15])
}
