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
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: inspect_microflow <mpr_file_path> <microflow_name>")
		fmt.Println("Example: inspect_microflow project.mpr MyMicroflow")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	targetName := os.Args[2]

	fmt.Printf("Opening MPR: %s\n", mprPath)
	fmt.Printf("Looking for microflow: %s\n\n", targetName)

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Failed to open MPR file: %v", err)
	}
	defer db.Close()

	// Get MPR base directory for reading .mxunit files
	mprDir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(mprDir, "mprcontents")

	microflows, err := listMicroflows(db, contentsDir)
	if err != nil {
		log.Fatalf("Failed to list microflows: %v", err)
	}

	fmt.Printf("Found %d microflows total\n\n", len(microflows))

	// Find matching microflow
	for _, mf := range microflows {
		if strings.Contains(mf.Name, targetName) {
			fmt.Printf("=== Microflow: %s ===\n", mf.Name)
			fmt.Printf("Module: %s\n", mf.ModuleName)
			fmt.Printf("ID: %s\n\n", mf.ID)

			// Export to JSON
			jsonData, err := json.MarshalIndent(mf.Content, "", "  ")
			if err != nil {
				log.Fatalf("Failed to marshal JSON: %v", err)
			}

			// Save to file
			filename := fmt.Sprintf("%s.json", strings.ReplaceAll(mf.Name, ".", "_"))
			if err := os.WriteFile(filename, jsonData, 0644); err != nil {
				log.Fatalf("Failed to write file: %v", err)
			}

			fmt.Printf("Exported to: %s\n", filename)
			fmt.Printf("File size: %.2f KB\n", float64(len(jsonData))/1024)

			// Search for ActionActivity
			countActivities(mf.Content)

			return
		}
	}

	fmt.Printf("Microflow '%s' not found\n", targetName)
}

type MicroflowInfo struct {
	ID         string
	Name       string
	ModuleName string
	Content    map[string]interface{}
}

func listMicroflows(db *sql.DB, contentsDir string) ([]MicroflowInfo, error) {
	query := `SELECT UnitID, ContainerID, ContainmentName, ContentsHash FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var microflows []MicroflowInfo
	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		// Convert UUID blob to string
		unitIDStr := blobToUUID(unitID)

		// Load content from mprcontents
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Extract type from BSON content
		typeName := getTypeFromContents(content)

		// Only process Microflows
		if typeName != "Microflows$Microflow" {
			continue
		}

		// Extract name from BSON content
		name := extractNameFromContents(content)

		microflows = append(microflows, MicroflowInfo{
			ID:         unitIDStr,
			Name:       name,
			ModuleName: containmentName,
			Content:    content,
		})
	}

	return microflows, nil
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		blob[3], blob[2], blob[1], blob[0],
		blob[5], blob[4],
		blob[7], blob[6],
		blob[8], blob[9], blob[10], blob[11], blob[12], blob[13], blob[14], blob[15])
}

func getTypeFromContents(contents map[string]interface{}) string {
	if typeName, ok := contents["$Type"].(string); ok {
		return typeName
	}
	return ""
}

func extractNameFromContents(contents map[string]interface{}) string {
	if name, ok := contents["Name"].(string); ok {
		return name
	}
	return ""
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

func countActivities(obj interface{}) {
	actionCount := 0
	microflowCallCount := 0

	var search func(interface{})
	search = func(o interface{}) {
		switch v := o.(type) {
		case map[string]interface{}:
			if typeStr, ok := v["$Type"].(string); ok {
				if typeStr == "Microflows$ActionActivity" {
					actionCount++
					// Check action type
					if action, ok := v["Action"].(map[string]interface{}); ok {
						if actionType, ok := action["$Type"].(string); ok {
							if actionType == "Microflows$MicroflowCall" {
								microflowCallCount++
								fmt.Printf("\n=== MicroflowCall Activity Found ===\n")
								if caption, ok := v["Caption"].(string); ok {
									fmt.Printf("Caption: %s\n", caption)
								}
								if name, ok := v["Name"].(string); ok {
									fmt.Printf("Name: %s\n", name)
								}
								// Print entire action
								actionJSON, _ := json.MarshalIndent(action, "", "  ")
								fmt.Printf("Action:\n%s\n", string(actionJSON))
							}
						}
					}
				}
			}
			for _, value := range v {
				search(value)
			}
		case []interface{}:
			for _, item := range v {
				search(item)
			}
		}
	}

	search(obj)

	fmt.Printf("\n=== Activity Summary ===\n")
	fmt.Printf("Total ActionActivity: %d\n", actionCount)
	fmt.Printf("MicroflowCall actions: %d\n", microflowCallCount)
}
