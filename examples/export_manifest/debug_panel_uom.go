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
	
	// Find PANEL_UpdateUoMFactor
	fmt.Println("🔍 Searching for PANEL_UpdateUoMFactor...")
	
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
		
		// Check if this is PANEL_UpdateUoMFactor
		name, ok := content["Name"].(string)
		if !ok || name != "PANEL_UpdateUoMFactor" {
			continue
		}
		
		fmt.Printf("\n✅ Found PANEL_UpdateUoMFactor (UnitID: %s)\n\n", unitID)
		
		// Search for inline nanoflows
		fmt.Println("🔍 Searching for inline nanoflows in page BSON...")
		findInlineNanoflows(content, 0)
		
		// Search for ACT_UpdateUoMFactor references
		fmt.Println("\n🔍 Searching for ACT_UpdateUoMFactor references...")
		findStringReferences(content, "ACT_UpdateUoMFactor", 0)
		
		break
	}
}

func findInlineNanoflows(obj interface{}, depth int) {
	if depth > 20 {
		return
	}
	
	indent := strings.Repeat("  ", depth)
	
	switch v := obj.(type) {
	case map[string]interface{}:
		// Check $Type
		if typeVal, ok := v["$Type"].(string); ok {
			if strings.Contains(typeVal, "Nanoflow") || strings.Contains(typeVal, "anoflow") {
				fmt.Printf("%s📦 Found type: %s\n", indent, typeVal)
				
				// Check for Name
				if nameVal, ok := v["Name"]; ok {
					fmt.Printf("%s   Name: %v\n", indent, nameVal)
				}
				
				// Print some context
				for key := range v {
					if key != "$Type" && key != "Name" {
						fmt.Printf("%s   Has field: %s\n", indent, key)
					}
				}
				fmt.Println()
			}
		}
		
		// Recurse
		for key, val := range v {
			if key == "$Type" || key == "Name" {
				continue
			}
			findInlineNanoflows(val, depth+1)
		}
		
	case primitive.A:
		for _, item := range v {
			findInlineNanoflows(item, depth+1)
		}
		
	case []interface{}:
		for _, item := range v {
			findInlineNanoflows(item, depth+1)
		}
	}
}

func findStringReferences(obj interface{}, searchStr string, depth int) {
	if depth > 20 {
		return
	}
	
	indent := strings.Repeat("  ", depth)
	
	switch v := obj.(type) {
	case string:
		if strings.Contains(v, searchStr) {
			fmt.Printf("%s📝 Found string: %s\n", indent, v)
		}
		
	case map[string]interface{}:
		for key, val := range v {
			if strVal, ok := val.(string); ok && strings.Contains(strVal, searchStr) {
				fmt.Printf("%s📋 Field '%s': %s\n", indent, key, strVal)
			}
			findStringReferences(val, searchStr, depth+1)
		}
		
	case primitive.A:
		for _, item := range v {
			findStringReferences(item, searchStr, depth+1)
		}
		
	case []interface{}:
		for _, item := range v {
			findStringReferences(item, searchStr, depth+1)
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
