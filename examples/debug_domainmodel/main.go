package main

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: debug_domainmodel <mpr_file_path> <module_name>")
		fmt.Println("Example: debug_domainmodel project.mpr Documents")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	targetModule := os.Args[2]

	fmt.Printf("Opening MPR: %s\n", mprPath)
	fmt.Printf("Looking for domain model in module: %s\n\n", targetModule)

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Find the domain model for the target module
	rows, err := db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, ContentsHash
		FROM Unit
	`)
	if err != nil {
		log.Fatalf("Failed to query units: %v", err)
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		contents, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Check if it's a DomainModel
		if typeStr, ok := contents["$Type"].(string); ok {
			if typeStr == "DomainModels$DomainModel" {
				name := extractName(contents)
				if strings.HasPrefix(name, targetModule+".") || name == targetModule {
					found = true
					fmt.Printf("Found domain model: %s\n\n", name)

					// Export to JSON
					jsonData, err := json.MarshalIndent(contents, "", "  ")
					if err != nil {
						log.Fatalf("Failed to marshal JSON: %v", err)
					}

					outputFile := fmt.Sprintf("%s_DomainModel.json", targetModule)
					if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
						log.Fatalf("Failed to write file: %v", err)
					}

					fmt.Printf("Domain model exported to: %s\n", outputFile)
					fmt.Printf("File size: %d bytes\n", len(jsonData))

					// Print summary
					if entities, ok := contents["Entities"].([]interface{}); ok {
						fmt.Printf("Total items in Entities array: %d\n", len(entities))
					}

					break
				}
			}
		}
	}

	if !found {
		fmt.Printf("No domain model found for module: %s\n", targetModule)
	}
}

func extractName(contents map[string]interface{}) string {
	if name, ok := contents["Name"].(string); ok {
		return name
	}
	return ""
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}

	var formatted [16]byte
	formatted[0] = blob[3]
	formatted[1] = blob[2]
	formatted[2] = blob[1]
	formatted[3] = blob[0]
	formatted[4] = blob[5]
	formatted[5] = blob[4]
	formatted[6] = blob[7]
	formatted[7] = blob[6]
	copy(formatted[8:], blob[8:16])

	hexStr := hex.EncodeToString(formatted[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hexStr[0:8],
		hexStr[8:12],
		hexStr[12:16],
		hexStr[16:20],
		hexStr[20:32])
}

func loadUnitContents(contentsDir string, unitID string) (map[string]interface{}, error) {
	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid unit ID: %s", unitID)
	}

	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read unit file: %w", err)
	}

	var content map[string]interface{}
	if err := bson.Unmarshal(data, &content); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	return content, nil
}
