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
	if len(os.Args) < 2 {
		fmt.Println("Usage: check_module <mpr_file_path> [module_name_prefix]")
		fmt.Println("Example: check_module project.mpr OpcenterEXFN")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	modulePrefix := ""
	if len(os.Args) >= 3 {
		modulePrefix = os.Args[2]
	}

	fmt.Printf("Opening MPR: %s\n", mprPath)
	if modulePrefix != "" {
		fmt.Printf("Looking for modules with prefix: %s\n\n", modulePrefix)
	} else {
		fmt.Println("Listing all modules\n")
	}

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Find all units
	rows, err := db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, ContentsHash
		FROM Unit
	`)
	if err != nil {
		log.Fatalf("Failed to query units: %v", err)
	}
	defer rows.Close()

	moduleCount := 0
	domainModelCount := 0
	entityCount := 0
	externalEntityCount := 0

	modules := make(map[string]bool)
	domainModels := []string{}
	entities := []string{}

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

		typeStr, ok := contents["$Type"].(string)
		if !ok {
			continue
		}

		name := extractName(contents)

		// Check for Module
		if typeStr == "Projects$Module" {
			if modulePrefix == "" || strings.Contains(name, modulePrefix) {
				moduleCount++
				modules[name] = true
				fmt.Printf("📦 Module: %s\n", name)
			}
		}

		// Check for DomainModel
		if typeStr == "DomainModels$DomainModel" {
			moduleName := extractModuleName(name)
			if modulePrefix == "" || strings.Contains(moduleName, modulePrefix) {
				domainModelCount++
				domainModels = append(domainModels, name)
				fmt.Printf("  📊 Domain Model: %s\n", name)

				// Analyze entities in this domain model
				if entitiesArray, ok := contents["Entities"].([]interface{}); ok {
					fmt.Printf("     Entities array length: %d\n", len(entitiesArray))

					// Try to find entity references
					entityRefs := 0
					for _, item := range entitiesArray {
						if intVal, ok := item.(int32); ok && intVal == 1 {
							continue // Skip array marker
						}
						if strVal, ok := item.(string); ok {
							entityRefs++
							if entityRefs <= 5 {
								fmt.Printf("     - Entity reference: %s\n", strVal)
							}
						}
						if mapVal, ok := item.(map[string]interface{}); ok {
							entityName := extractName(mapVal)
							if entityType, ok := mapVal["$Type"].(string); ok {
								fmt.Printf("     - Embedded entity: %s (Type: %s)\n", entityName, entityType)
							}
						}
					}
					if entityRefs > 5 {
						fmt.Printf("     ... and %d more entity references\n", entityRefs-5)
					}
				}
			}
		}

		// Check for Entity
		if typeStr == "DomainModels$Entity" || typeStr == "DomainModels$ExternalEntity" {
			moduleName := extractModuleName(name)
			if modulePrefix == "" || strings.Contains(moduleName, modulePrefix) {
				entityCount++
				isExternal := typeStr == "DomainModels$ExternalEntity"

				// Check other external indicators
				if !isExternal {
					if loc, ok := contents["Location"].(string); ok && loc != "" {
						isExternal = true
					}
					if remote, ok := contents["RemoteSourceDocument"].(string); ok && remote != "" {
						isExternal = true
					}
					if remote, ok := contents["RemoteSource"]; ok && remote != nil {
						isExternal = true
					}
				}

				if isExternal {
					externalEntityCount++
					entities = append(entities, name)
					fmt.Printf("  ✓ EXTERNAL Entity: %s\n", name)

					// Export first external entity to JSON
					if externalEntityCount == 1 {
						jsonData, _ := json.MarshalIndent(contents, "", "  ")
						outputFile := "first_external_entity.json"
						os.WriteFile(outputFile, jsonData, 0644)
						fmt.Printf("    (Exported to %s for analysis)\n", outputFile)
					}
				}
			}
		}
	}

	fmt.Println("\n=== Summary ===")
	fmt.Printf("Modules found: %d\n", moduleCount)
	fmt.Printf("Domain Models found: %d\n", domainModelCount)
	fmt.Printf("Total Entities found: %d\n", entityCount)
	fmt.Printf("External Entities found: %d\n", externalEntityCount)

	if len(domainModels) > 0 {
		fmt.Println("\nDomain Models:")
		for _, dm := range domainModels {
			fmt.Printf("  - %s\n", dm)
		}
	}
}

func extractName(contents map[string]interface{}) string {
	if name, ok := contents["Name"].(string); ok {
		return name
	}
	return "Unknown"
}

func extractModuleName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) > 0 {
		return parts[0]
	}
	return fullName
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
