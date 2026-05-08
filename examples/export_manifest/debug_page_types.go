package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	mprPath := `C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr`
	tempDir, db, err := openMPR(mprPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)
	defer db.Close()

	contentsDir := filepath.Join(tempDir, "mprcontents")

	query := `SELECT UnitID FROM Unit LIMIT 100`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer rows.Close()

	typesFound := make(map[string]int)
	pageTypes := make(map[string][]string) // type -> list of names

	for rows.Next() {
		var unitIDBlob []byte
		if err := rows.Scan(&unitIDBlob); err != nil {
			continue
		}

		unitID := hex.EncodeToString(unitIDBlob)
		filePath := filepath.Join(contentsDir, unitID[:2], unitID[2:4], unitID+".mxunit")

		data, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		var doc map[string]interface{}
		err = bson.Unmarshal(data, &doc)
		if err != nil {
			continue
		}

		typeField, ok := doc["$Type"].(string)
		if !ok {
			continue
		}

		typesFound[typeField]++

		// If it looks like a page type, store the name
		if len(typeField) > 4 && typeField[:5] == "Pages" {
			name, _ := doc["Name"].(string)
			pageTypes[typeField] = append(pageTypes[typeField], name)
		}
	}

	fmt.Println("\nAll types found:")
	for t, count := range typesFound {
		fmt.Printf("  %s: %d\n", t, count)
	}

	fmt.Println("\nPage types with names:")
	for t, names := range pageTypes {
		fmt.Printf("  %s:\n", t)
		for i, name := range names {
			if i >= 5 {
				fmt.Printf("    ... and %d more\n", len(names)-5)
				break
			}
			fmt.Printf("    - %s\n", name)
		}
	}
}
