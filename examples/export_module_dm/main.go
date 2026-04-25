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
		fmt.Println("Usage: export_module_dm <mpr_file_path> <module_name>")
		fmt.Println("Example: export_module_dm project.mpr OpcenterEXFN_ReferenceData_Connector")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	targetModule := os.Args[2]

	fmt.Printf("Opening MPR: %s\n", mprPath)
	fmt.Printf("Looking for module: %s\n\n", targetModule)

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Find domain models
	rows, err := db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, ContentsHash
		FROM Unit
		WHERE ContainmentName = 'DomainModel'
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		containerIDStr := blobToUUID(containerID)

		// Load domain model contents
		dmContents, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		if typeStr, ok := dmContents["$Type"].(string); ok {
			if typeStr == "DomainModels$DomainModel" {
				// Get module name from container
				moduleContents, err := loadUnitContents(contentsDir, containerIDStr)
				if err != nil {
					continue
				}

				moduleName := ""
				if name, ok := moduleContents["Name"].(string); ok {
					moduleName = name
				}

				if moduleName == targetModule {
					fmt.Printf("Found domain model for module: %s\n\n", moduleName)

					// Export to JSON
					jsonData, err := json.MarshalIndent(dmContents, "", "  ")
					if err != nil {
						log.Fatal(err)
					}

					outputFile := fmt.Sprintf("%s_DomainModel.json", moduleName)
					if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
						log.Fatal(err)
					}

					fmt.Printf("Exported to: %s\n", outputFile)
					fmt.Printf("Size: %d bytes\n\n", len(jsonData))

					// Analyze entities
					if entitiesArray, ok := dmContents["Entities"].([]interface{}); ok {
						fmt.Printf("Entities array has %d items\n", len(entitiesArray))
						entityCount := 0
						for _, item := range entitiesArray {
							if intVal, ok := item.(int32); ok && (intVal == 1 || intVal == 2 || intVal == 3) {
								continue
							}
							if entityMap, ok := item.(map[string]interface{}); ok {
								entityCount++
								if name, ok := entityMap["Name"].(string); ok {
									entityType := entityMap["$Type"]
									fmt.Printf("  [%d] %s (Type: %v)\n", entityCount, name, entityType)

									// Check for external markers
									if loc, ok := entityMap["Location"].(string); ok && !strings.Contains(loc, ";") {
										fmt.Printf("      Location: %s\n", loc)
									}
									if remote, ok := entityMap["RemoteSource"]; ok && remote != nil {
										fmt.Printf("      RemoteSource: %v\n", remote)
									}
								}
							}
						}
					}

					return
				}
			}
		}
	}

	fmt.Printf("Module '%s' not found\n", targetModule)
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
