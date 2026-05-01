package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run list_pages.go <mpr_path>")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	fmt.Printf("📦 MPR: %s\n\n", mprPath)

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	fmt.Println("📄 Pages found:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	count := 0
	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		if typeName, ok := content["$Type"].(string); ok && typeName == "Pages$Page" {
			name := extractNameFromContents(content)
			qualifiedName := ""
			if qn, ok := content["QualifiedName"].(string); ok {
				qualifiedName = qn
			}

			count++
			fmt.Printf("%3d. %-40s (QualifiedName: %s)\n", count, name, qualifiedName)

			// Show if it starts with "Counter"
			if strings.HasPrefix(name, "Counter") {
				fmt.Printf("     👉 MATCHES Counter*\n")
			}
		}
	}

	fmt.Printf("\n✅ Total: %d page(s)\n", count)
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		blob[3], blob[2], blob[1], blob[0],
		blob[5], blob[4],
		blob[7], blob[6],
		blob[8], blob[9],
		blob[10], blob[11], blob[12], blob[13], blob[14], blob[15])
}

func loadUnitContents(contentsDir string, unitID string) (map[string]interface{}, error) {
	cleanID := strings.ReplaceAll(unitID, "-", "")
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func extractNameFromContents(content map[string]interface{}) string {
	if name, ok := content["Name"].(string); ok {
		return name
	}
	if qualifiedName, ok := content["QualifiedName"].(string); ok {
		parts := strings.Split(qualifiedName, ".")
		if len(parts) > 0 {
			return parts[len(parts)-1]
		}
	}
	return ""
}
