package main

import (
	"database/sql"
	"encoding/hex"
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
		fmt.Println("Usage: find_external_entities <mpr_file_path>")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	fmt.Printf("Opening MPR: %s\n\n", mprPath)

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// First, build a map of ContainerID -> Module Name
	fmt.Println("Building module map...")
	moduleMap := make(map[string]string)

	rows, err := db.Query(`
		SELECT UnitID, ContainmentName
		FROM Unit
		WHERE ContainmentName = 'DomainModel'
	`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var unitID []byte
		var containmentName string

		if err := rows.Scan(&unitID, &containmentName); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)

		// Find the parent module by querying the container
		var containerID []byte
		err := db.QueryRow(`
			SELECT ContainerID FROM Unit WHERE UnitID = ?
		`, unitID).Scan(&containerID)

		if err == nil {
			containerIDStr := blobToUUID(containerID)

			// Load the container to get module name
			contents, err := loadUnitContents(contentsDir, containerIDStr)
			if err == nil {
				if name, ok := contents["Name"].(string); ok {
					moduleMap[unitIDStr] = name
					fmt.Printf("  Domain Model %s -> Module: %s\n", unitIDStr[:8], name)
				}
			}
		}
	}
	rows.Close()

	fmt.Printf("\nFound %d modules with domain models\n\n", len(moduleMap))

	// Now find all entities
	fmt.Println("Searching for external entities...")

	rows2, err := db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, ContentsHash
		FROM Unit
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()

	totalEntities := 0
	externalEntities := 0
	entitiesByModule := make(map[string][]string)

	for rows2.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows2.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
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

		// Check for entities
		if typeStr == "DomainModels$Entity" || typeStr == "DomainModels$ExternalEntity" {
			totalEntities++

			name := "Unknown"
			if n, ok := contents["Name"].(string); ok {
				name = n
			}

			// Get module name from container
			containerIDStr := blobToUUID(containerID)
			moduleName := moduleMap[containerIDStr]
			if moduleName == "" {
				moduleName = "Unknown Module"
			}

			// Check if external
			isExternal := typeStr == "DomainModels$ExternalEntity"

			if !isExternal {
				// Check other indicators
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
				externalEntities++
				entitiesByModule[moduleName] = append(entitiesByModule[moduleName], name)

				if externalEntities <= 20 {
					fmt.Printf("✓ EXTERNAL: %s.%s (Type: %s)\n", moduleName, name, typeStr)
				}
			}
		}
	}

	fmt.Println("\n=== Summary ===")
	fmt.Printf("Total entities: %d\n", totalEntities)
	fmt.Printf("External entities: %d\n", externalEntities)
	fmt.Printf("Modules with external entities: %d\n\n", len(entitiesByModule))

	if len(entitiesByModule) > 0 {
		fmt.Println("External entities by module:")
		for module, entities := range entitiesByModule {
			fmt.Printf("  %s: %d entities\n", module, len(entities))
			for i, entity := range entities {
				if i < 5 {
					fmt.Printf("    - %s\n", entity)
				}
			}
			if len(entities) > 5 {
				fmt.Printf("    ... and %d more\n", len(entities)-5)
			}
		}
	}
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
