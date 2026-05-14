package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	mprPath := "C:\\Workspaces\\Mendix\\MDUI\\System_Mendix_CLI\\OC EX System.mpr"

	// Open database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Find ACT_UpdateUoMFactor_1
	fmt.Println("🔍 Searching for ACT_UpdateUoMFactor_1...")

	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		fmt.Printf("Error querying units: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var unitIDBytes []byte
		if err := rows.Scan(&unitIDBytes); err != nil {
			continue
		}

		unitID := guidToString(unitIDBytes)
		if unitID == "" {
			continue
		}

		// Load BSON content
		content, err := loadUnitContents(contentsDir, unitID)
		if err != nil {
			continue
		}

		// Check Name
		name, ok := content["Name"].(string)
		if !ok || name != "ACT_UpdateUoMFactor_1" {
			continue
		}

		fmt.Printf("\n✅ Found ACT_UpdateUoMFactor_1 (UnitID: %s)\n\n", unitID)

		// Find all $Type fields
		fmt.Println("🔍 All action types in this nanoflow:")
		findAllTypes(content, make(map[string]int))

		// Find all fields that might contain references
		fmt.Println("\n🔍 All fields that might be microflow/nanoflow references:")
		findReferenceFields(content, 0)

		break
	}
}

func findAllTypes(obj interface{}, typeCounts map[string]int) map[string]int {
	switch v := obj.(type) {
	case map[string]interface{}:
		if typeVal, ok := v["$Type"].(string); ok {
			typeCounts[typeVal]++
		}

		for _, val := range v {
			findAllTypes(val, typeCounts)
		}

	case primitive.A:
		for _, item := range v {
			findAllTypes(item, typeCounts)
		}

	case []interface{}:
		for _, item := range v {
			findAllTypes(item, typeCounts)
		}
	}

	// Print results
	if len(typeCounts) > 0 {
		for t, count := range typeCounts {
			fmt.Printf("  - %s: %d occurrence(s)\n", t, count)
		}
	}

	return typeCounts
}

func findReferenceFields(obj interface{}, depth int) {
	if depth > 15 {
		return
	}

	indent := strings.Repeat("  ", depth)

	switch v := obj.(type) {
	case map[string]interface{}:
		for key, val := range v {
			// Look for fields that might contain microflow references
			if key == "Microflow" || key == "Nanoflow" || key == "MicroflowCall" ||
				key == "ActionCall" || key == "Action" || strings.Contains(key, "Flow") {
				if strVal, ok := val.(string); ok && strVal != "" {
					fmt.Printf("%s📋 %s: %s\n", indent, key, strVal)
				}
			}
		}

		// Recurse
		for _, val := range v {
			findReferenceFields(val, depth+1)
		}

	case primitive.A:
		for _, item := range v {
			findReferenceFields(item, depth+1)
		}

	case []interface{}:
		for _, item := range v {
			findReferenceFields(item, depth+1)
		}
	}
}

// Helper functions
func guidToString(guidBytes []byte) string {
	if len(guidBytes) != 16 {
		return ""
	}

	// Windows GUID byte order
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		guidBytes[3], guidBytes[2], guidBytes[1], guidBytes[0],
		guidBytes[5], guidBytes[4],
		guidBytes[7], guidBytes[6],
		guidBytes[8], guidBytes[9],
		guidBytes[10], guidBytes[11], guidBytes[12], guidBytes[13], guidBytes[14], guidBytes[15])
}

func loadUnitContents(contentsDir, unitID string) (map[string]interface{}, error) {
	cleanID := strings.ReplaceAll(unitID, "-", "")
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var content map[string]interface{}
	if err := bson.Unmarshal(data, &content); err != nil {
		return nil, err
	}

	return content, nil
}
