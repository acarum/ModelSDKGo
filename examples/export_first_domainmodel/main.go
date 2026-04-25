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
		fmt.Println("Usage: export_first_domainmodel <mpr_file_path>")
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

	rows, err := db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, ContentsHash
		FROM Unit
	`)
	if err != nil {
		log.Fatalf("Failed to query units: %v", err)
	}
	defer rows.Close()

	dmCount := 0
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

		if typeStr, ok := contents["$Type"].(string); ok {
			if typeStr == "DomainModels$DomainModel" {
				dmCount++

				// Export first 3 domain models
				if dmCount <= 3 {
					jsonData, err := json.MarshalIndent(contents, "", "  ")
					if err != nil {
						continue
					}

					outputFile := fmt.Sprintf("DomainModel_%d.json", dmCount)
					if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
						continue
					}

					fmt.Printf("Exported Domain Model #%d to: %s\n", dmCount, outputFile)
					fmt.Printf("  Container ID: %s\n", blobToUUID(containerID))
					fmt.Printf("  Containment Name: %s\n", containmentName)
					fmt.Printf("  Size: %d bytes\n", len(jsonData))

					// Try to find name in various places
					if name, ok := contents["Name"].(string); ok {
						fmt.Printf("  Name field: %s\n", name)
					} else {
						fmt.Printf("  Name field: NOT FOUND\n")
					}

					if qualName, ok := contents["QualifiedName"].(string); ok {
						fmt.Printf("  QualifiedName field: %s\n", qualName)
					}

					// Show top-level keys
					fmt.Printf("  Top-level keys: ")
					keys := []string{}
					for k := range contents {
						keys = append(keys, k)
					}
					fmt.Printf("%v\n\n", keys)
				}

				if dmCount >= 3 {
					break
				}
			}
		}
	}

	fmt.Printf("\nTotal Domain Models found: %d\n", dmCount)
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
