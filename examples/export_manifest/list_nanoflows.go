package main

import (
	"database/sql"
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
		fmt.Println("Usage: go run list_nanoflows.go <path_to_mpr> [search_pattern]")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	searchPattern := ""
	if len(os.Args) > 2 {
		searchPattern = strings.ToLower(os.Args[2])
	}

	contentsDir := extractContentsDir(mprPath)
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		data, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		typeStr, _ := data["$Type"].(string)
		if typeStr == "Microflows$Nanoflow" {
			name, _ := data["Name"].(string)
			qname, _ := data["QualifiedName"].(string)
			
			if searchPattern == "" || strings.Contains(strings.ToLower(name), searchPattern) || strings.Contains(strings.ToLower(qname), searchPattern) {
				fmt.Printf("Name: %s\n", name)
				if qname != "" {
					fmt.Printf("QualifiedName: %s\n", qname)
				}
				fmt.Println("---")
				count++
			}
		}
	}

	fmt.Printf("\nFound %d nanoflow(s)\n", count)
}

func extractContentsDir(mprPath string) string {
	dir := filepath.Dir(mprPath)
	base := filepath.Base(mprPath)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(dir, base, "model")
}

func loadUnitContents(contentsDir, unitID string) (map[string]interface{}, error) {
	filePath := filepath.Join(contentsDir, unitID)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := bson.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		blob[0:4], blob[4:6], blob[6:8], blob[8:10], blob[10:16])
}
