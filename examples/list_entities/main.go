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
		fmt.Println("Usage: list_entities <mpr_file_path>")
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

	// Find all Entity units
	rows, err := db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, ContentsHash
		FROM Unit
	`)
	if err != nil {
		log.Fatalf("Failed to query units: %v", err)
	}
	defer rows.Close()

	entityCount := 0
	externalCount := 0

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

		// Check if it's an Entity
		if typeStr, ok := contents["$Type"].(string); ok {
			if typeStr == "DomainModels$Entity" || typeStr == "DomainModels$ExternalEntity" {
				entityCount++
				name := extractName(contents)

				// Check if external
				isExternal := false
				location := ""

				if typeStr == "DomainModels$ExternalEntity" {
					isExternal = true
					externalCount++
				}

				if loc, ok := contents["Location"].(string); ok && loc != "" {
					location = loc
					if !isExternal {
						isExternal = true
						externalCount++
					}
				}

				if remoteSource, ok := contents["RemoteSourceDocument"].(string); ok && remoteSource != "" {
					location = remoteSource
					if !isExternal {
						isExternal = true
						externalCount++
					}
				}

				if entityCount <= 50 { // Print first 50 entities
					if isExternal {
						fmt.Printf("✓ EXTERNAL Entity: %s\n", name)
						if location != "" {
							fmt.Printf("  Location: %s\n", location)
						}
						fmt.Printf("  Type: %s\n\n", typeStr)
					} else {
						fmt.Printf("  Entity: %s (Type: %s)\n", name, typeStr)
					}
				}
			}
		}
	}

	fmt.Println("=== Summary ===")
	fmt.Printf("Total entities found: %d\n", entityCount)
	fmt.Printf("External entities: %d\n", externalCount)
	fmt.Printf("Regular entities: %d\n", entityCount-externalCount)
}

func extractName(contents map[string]interface{}) string {
	if name, ok := contents["Name"].(string); ok {
		return name
	}
	return "Unknown"
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
