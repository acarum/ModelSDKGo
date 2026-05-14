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
	if len(os.Args) < 3 {
		log.Fatal("Usage: inspect_microflow_bson.exe <path-to-mpr> <microflow-qualified-name>")
	}

	mprPath := os.Args[1]
	microflowName := os.Args[2]

	fmt.Printf("🔍 Opening MPR: %s\n", mprPath)

	// Open MPR
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

	// Parse qualified name
	parts := strings.Split(microflowName, ".")
	if len(parts) != 2 {
		log.Fatalf("Microflow name must be in format: Module.Name")
	}
	moduleName := parts[0]
	mfName := parts[1]

	fmt.Printf("\n🔍 Searching for %s in module %s...\n", mfName, moduleName)

	// Build cache to find the microflow
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	cache := buildMicroflowNameToUnitIDCacheLazy(db, contentsDir, []string{microflowName})

	unitID, found := cache[microflowName]
	if !found {
		log.Fatalf("❌ Microflow not found: %s", microflowName)
	}

	fmt.Printf("\n✅ Found microflow: UnitID=%s\n", unitID)

	// Load BSON content
	content, err := loadUnitContents(contentsDir, unitID)
	if err != nil {
		log.Fatalf("Failed to load unit: %v", err)
	}

	fmt.Printf("\n📄 BSON Content (top-level keys):\n")
	for key := range content {
		fmt.Printf("  - %s\n", key)
	}

	// Pretty print full BSON structure
	fmt.Printf("\n📋 Full BSON structure:\n")
	jsonData, _ := json.MarshalIndent(content, "", "  ")
	fmt.Println(string(jsonData))
}

func buildMicroflowNameToUnitIDCacheLazy(db *sql.DB, contentsDir string, microflowNames []string) map[string]string {
	cache := make(map[string]string)
	needed := make(map[string]bool)
	for _, name := range microflowNames {
		needed[name] = true
	}

	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		return cache
	}
	defer rows.Close()

	foundCount := 0
	for rows.Next() {
		if foundCount >= len(needed) {
			break
		}

		var unitIDBytes []byte
		if err := rows.Scan(&unitIDBytes); err != nil {
			continue
		}

		unitID := guidToString(unitIDBytes)
		if unitID == "" {
			continue
		}

		content, err := loadUnitContents(contentsDir, unitID)
		if err != nil {
			continue
		}

		bsonType, ok := content["$Type"].(string)
		if !ok || (bsonType != "Microflows$Microflow" && bsonType != "Microflows$Nanoflow") {
			continue
		}

		name, ok := content["Name"].(string)
		if !ok || name == "" {
			continue
		}

		moduleName := findModuleByTraversal(unitID, db)
		if moduleName == "" {
			continue
		}

		qualifiedName := moduleName + "." + name
		if needed[qualifiedName] {
			cache[qualifiedName] = unitID
			foundCount++
		}
	}

	return cache
}

func loadUnitContents(contentsDir string, unitID string) (map[string]interface{}, error) {
	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid unit ID: %s", unitID)
	}

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

func findModuleByTraversal(unitID string, db *sql.DB) string {
	currentID := unitID
	for i := 0; i < 20; i++ {
		var containerIDBytes []byte
		var bsonTypeBytes []byte

		query := "SELECT ContainerID, Type FROM Unit WHERE UnitID = ?"
		err := db.QueryRow(query, stringToGUID(currentID)).Scan(&containerIDBytes, &bsonTypeBytes)
		if err != nil {
			return ""
		}

		typeName := string(bsonTypeBytes)
		if typeName == "Projects$Module" {
			var nameBytes []byte
			if err := db.QueryRow("SELECT Name FROM Unit WHERE UnitID = ?", stringToGUID(currentID)).Scan(&nameBytes); err == nil {
				return string(nameBytes)
			}
			return ""
		}

		if len(containerIDBytes) == 0 {
			break
		}
		currentID = guidToString(containerIDBytes)
	}
	return ""
}

func guidToString(guidBytes []byte) string {
	if len(guidBytes) != 16 {
		return ""
	}
	b := make([]byte, 16)
	copy(b, guidBytes)
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15])
}

func stringToGUID(guidStr string) []byte {
	guidStr = strings.ReplaceAll(guidStr, "-", "")
	if len(guidStr) != 32 {
		return nil
	}

	b := make([]byte, 16)
	for i := 0; i < 16; i++ {
		fmt.Sscanf(guidStr[i*2:i*2+2], "%02x", &b[i])
	}

	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	return b
}
