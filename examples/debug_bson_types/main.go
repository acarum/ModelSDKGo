package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: debug_bson_types <mpr_file_path> <module_name>")
		fmt.Println("Example: debug_bson_types project.mpr OpcenterEXFN_ReferenceData_Connector")
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

		// Load module
		moduleContents, err := loadUnitContents(contentsDir, containerIDStr)
		if err != nil {
			continue
		}

		moduleName := ""
		if name, ok := moduleContents["Name"].(string); ok {
			moduleName = name
		}

		if moduleName == targetModule {
			// Load domain model
			dmContents, err := loadUnitContents(contentsDir, unitIDStr)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Found domain model for module: %s\n\n", moduleName)
			fmt.Println("Top-level keys and types:")

			for key, value := range dmContents {
				valueType := reflect.TypeOf(value)
				fmt.Printf("  %s: %v\n", key, valueType)

				if key == "Entities" {
					fmt.Println("\n    === Entities field details ===")
					if arr, ok := value.([]interface{}); ok {
						fmt.Printf("    Converted to []interface{}, length: %d\n", len(arr))
						for i := 0; i < 5 && i < len(arr); i++ {
							itemType := reflect.TypeOf(arr[i])
							fmt.Printf("      [%d] %v", i, itemType)
							if intVal, ok := arr[i].(int32); ok {
								fmt.Printf(" = %d", intVal)
							} else if intVal, ok := arr[i].(int64); ok {
								fmt.Printf(" = %d", intVal)
							} else if mapVal, ok := arr[i].(map[string]interface{}); ok {
								if name, ok := mapVal["Name"].(string); ok {
									fmt.Printf(" (Name: %s)", name)
								}
							}
							fmt.Println()
						}
					} else if arr, ok := value.(bson.A); ok {
						fmt.Printf("    BSON Array (bson.A), length: %d\n", len(arr))
					} else {
						fmt.Printf("    Type: %v, Value: %v\n", reflect.TypeOf(value), value)
					}
				}
			}

			return
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
