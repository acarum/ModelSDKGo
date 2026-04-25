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
	if len(os.Args) < 2 {
		fmt.Println("Usage: list_actions <mpr_file_path>")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	fmt.Printf("Opening MPR: %s\n", mprPath)

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Failed to open MPR file: %v", err)
	}
	defer db.Close()

	mprDir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(mprDir, "mprcontents")

	microflows, err := listMicroflows(db, contentsDir)
	if err != nil {
		log.Fatalf("Failed to list microflows: %v", err)
	}

	fmt.Printf("Found %d microflows. Analyzing action types...\n\n", len(microflows))

	actionTypes := make(map[string]int)
	actionExamples := make(map[string]string)

	for _, mf := range microflows {
		// Convert to JSON for easier processing
		jsonData, err := json.Marshal(mf.Content)
		if err != nil {
			continue
		}
		var contentMap map[string]interface{}
		if err := json.Unmarshal(jsonData, &contentMap); err != nil {
			continue
		}

		findActionTypes(contentMap, actionTypes, actionExamples, mf.Name)
	}

	fmt.Println("=== Action Types Found ===")
	for actionType, count := range actionTypes {
		fmt.Printf("%s: %d occurrences\n", actionType, count)
		if example, ok := actionExamples[actionType]; ok {
			fmt.Printf("  Example microflow: %s\n", example)
		}
	}
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

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		typeName := getTypeFromContents(content)
		if typeName != "Microflows$Microflow" {
			continue
		}

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

func findActionTypes(obj interface{}, types map[string]int, examples map[string]string, microflowName string) {
	switch v := obj.(type) {
	case map[string]interface{}:
		if typeStr, ok := v["$Type"].(string); ok {
			if typeStr == "Microflows$ActionActivity" {
				if action, ok := v["Action"].(map[string]interface{}); ok {
					if actionType, ok := action["$Type"].(string); ok {
						types[actionType]++
						if _, exists := examples[actionType]; !exists {
							examples[actionType] = microflowName
						}
					}
				}
			}
		}

		for _, value := range v {
			findActionTypes(value, types, examples, microflowName)
		}

	case []interface{}:
		for _, item := range v {
			findActionTypes(item, types, examples, microflowName)
		}
	}
}
