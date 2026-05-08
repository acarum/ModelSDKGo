package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
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
	
	// Search for ACT_UpdateUoMFactor nanoflows
	fmt.Println("🔍 Searching for ACT_UpdateUoMFactor* nanoflows...")
	
	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		fmt.Printf("Error querying units: %v\n", err)
		return
	}
	defer rows.Close()
	
	foundCount := 0
	
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
		
		// Check $Type
		bsonType, ok := content["$Type"].(string)
		if !ok {
			continue
		}
		
		// Only check Microflows and Nanoflows
		if bsonType != "Microflows$Microflow" && bsonType != "Microflows$Nanoflow" {
			continue
		}
		
		// Check Name
		name, ok := content["Name"].(string)
		if !ok || !strings.Contains(name, "ACT_UpdateUoMFactor") {
			continue
		}
		
		foundCount++
		fmt.Printf("\n✅ Found: %s\n", name)
		fmt.Printf("   Type: %s\n", bsonType)
		fmt.Printf("   UnitID: %s\n", unitID)
		
		// Find module
		moduleName := findModuleByTraversal(unitID, db, contentsDir)
		fmt.Printf("   Module: %s\n", moduleName)
		fmt.Printf("   Qualified: %s.%s\n", moduleName, name)
		
		// Check if it calls other microflows
		calls := extractMicroflowCalls(content, moduleName)
		if len(calls) > 0 {
			fmt.Printf("   Calls:\n")
			for _, call := range calls {
				fmt.Printf("     - %s\n", call)
			}
		} else {
			fmt.Printf("   Calls: (none found)\n")
		}
	}
	
	if foundCount == 0 {
		fmt.Println("\n❌ No ACT_UpdateUoMFactor* nanoflows found in database!")
	} else {
		fmt.Printf("\n✅ Found %d ACT_UpdateUoMFactor* nanoflow(s)\n", foundCount)
	}
}

func extractMicroflowCalls(content map[string]interface{}, defaultModule string) []string {
	var calls []string
	seen := make(map[string]bool)
	
	var extract func(obj interface{})
	extract = func(obj interface{}) {
		switch v := obj.(type) {
		case map[string]interface{}:
			// Look for MicroflowCall action
			if action, ok := v["$Type"].(string); ok && (action == "Microflows$MicroflowCall" || action == "Microflows$NanoflowCall") {
				if mfName, ok := v["MicroflowCall"].(string); ok && mfName != "" {
					qualifiedName := ensureQualifiedName(mfName, defaultModule)
					if !seen[qualifiedName] {
						calls = append(calls, qualifiedName)
						seen[qualifiedName] = true
					}
				}
			}
			
			// Recurse
			for _, val := range v {
				extract(val)
			}
			
		case []interface{}:
			for _, item := range v {
				extract(item)
			}
		}
	}
	
	extract(content)
	return calls
}

func ensureQualifiedName(name string, defaultModule string) string {
	if strings.Contains(name, ".") {
		return name
	}
	if defaultModule != "" {
		return defaultModule + "." + name
	}
	return name
}

func findModuleByTraversal(unitID string, db *sql.DB, contentsDir string) string {
	// Build module map
	moduleMap := make(map[string]string)
	
	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		return ""
	}
	defer rows.Close()
	
	for rows.Next() {
		var unitIDBytes []byte
		if err := rows.Scan(&unitIDBytes); err != nil {
			continue
		}
		
		unitGUID := guidToString(unitIDBytes)
		if unitGUID == "" {
			continue
		}
		
		content, err := loadUnitContents(contentsDir, unitGUID)
		if err != nil {
			continue
		}
		
		if typeVal, ok := content["$Type"].(string); ok && strings.Contains(typeVal, "Projects$Module") {
			if moduleName, ok := content["Name"].(string); ok && moduleName != "" {
				moduleMap[unitGUID] = moduleName
			}
		}
	}
	
	// Traverse up
	guidBytes := stringToWindowsGUID(unitID)
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

func stringToWindowsGUID(uuidStr string) []byte {
	cleaned := strings.ReplaceAll(uuidStr, "-", "")
	if len(cleaned) != 32 {
		return nil
	}
	
	guid := make([]byte, 16)
	fmt.Sscanf(cleaned[0:2], "%02x", &guid[3])
	fmt.Sscanf(cleaned[2:4], "%02x", &guid[2])
	fmt.Sscanf(cleaned[4:6], "%02x", &guid[1])
	fmt.Sscanf(cleaned[6:8], "%02x", &guid[0])
	fmt.Sscanf(cleaned[8:10], "%02x", &guid[5])
	fmt.Sscanf(cleaned[10:12], "%02x", &guid[4])
	fmt.Sscanf(cleaned[12:14], "%02x", &guid[7])
	fmt.Sscanf(cleaned[14:16], "%02x", &guid[6])
	for i := 8; i < 16; i++ {
		fmt.Sscanf(cleaned[i*2:i*2+2], "%02x", &guid[i])
	}
	return guid
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
