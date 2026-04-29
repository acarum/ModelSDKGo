// Debug tool to check if page BSON contains QualifiedName
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthropics/modelsdk-go"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_bson_qualifiedname <mpr-file>")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		log.Fatalf("Error opening MPR: %v", err)
	}
	defer reader.Close()

	pages, err := reader.ListPages()
	if err != nil {
		log.Fatalf("Error listing pages: %v", err)
	}

	fmt.Printf("Found %d pages\n\n", len(pages))

	// Take first 3 pages as sample
	for i := 0; i < 3 && i < len(pages); i++ {
		page := pages[i]

		fmt.Printf("=== Page %d: %s ===\n", i+1, page.Name)
		fmt.Printf("ID: %s\n", page.ID)
		fmt.Printf("ContainerID: %s\n", page.ContainerID)

		// Read the BSON directly
		contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")
		pageID := string(page.ID)
		cleanID := strings.ReplaceAll(pageID, "-", "")
		dir1 := cleanID[0:2]
		dir2 := cleanID[2:4]
		filePath := filepath.Join(contentsDir, dir1, dir2, pageID+".mxunit")

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading BSON: %v\n\n", err)
			continue
		}

		var content map[string]interface{}
		if err := bson.Unmarshal(data, &content); err != nil {
			fmt.Printf("Error unmarshaling BSON: %v\n\n", err)
			continue
		}

		// List all top-level keys
		fmt.Println("Available BSON fields:")
		for key := range content {
			if strings.HasPrefix(key, "$") || strings.Contains(strings.ToLower(key), "name") || strings.Contains(strings.ToLower(key), "qualified") {
				fmt.Printf("  %s: %v\n", key, content[key])
			}
		}

		// Specifically check for QualifiedName
		if qn, ok := content["QualifiedName"]; ok {
			fmt.Printf("\n✅ QualifiedName FOUND: %v\n", qn)
		} else {
			fmt.Printf("\n❌ QualifiedName field NOT found\n")
		}

		// Check for other name-related fields
		nameFields := []string{"Name", "FullName", "FullyQualifiedName", "ModuleName", "QualifiedName"}
		for _, field := range nameFields {
			if val, ok := content[field]; ok {
				fmt.Printf("  %s: %v\n", field, val)
			}
		}

		fmt.Println()
	}
}
