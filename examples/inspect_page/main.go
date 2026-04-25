package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthropics/modelsdk-go"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: inspect_page <mpr_file_path> <page_name>")
		fmt.Println("Example: inspect_page MyApp.mpr StateMachine_Details")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	pageName := os.Args[2]

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		log.Fatalf("Error opening MPR: %v", err)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("Searching for page: %s\n\n", pageName)

	// Get all pages
	pages, err := reader.ListPages()
	if err != nil {
		log.Fatalf("Error listing pages: %v", err)
	}

	// Find the specific page
	var pageID string
	foundPage := false

	for _, page := range pages {
		if page.Name == pageName {
			fmt.Printf("=== Page Found ===\n")
			fmt.Printf("Name: %s\n", page.Name)
			fmt.Printf("ID: %s\n", page.ID)
			fmt.Printf("Type: %s\n\n", page.TypeName)
			pageID = string(page.ID)
			foundPage = true
			break
		}
	}

	if !foundPage {
		fmt.Printf("Page '%s' not found!\n", pageName)
		os.Exit(1)
	}

	// Read the raw BSON content directly from the .mxunit file
	bsonData, err := loadPageBSON(mprPath, pageID)
	if err != nil {
		log.Fatalf("Error loading page BSON: %v", err)
	}

	// Parse BSON to generic map
	var pageData map[string]interface{}
	if err := bson.Unmarshal(bsonData, &pageData); err != nil {
		log.Fatalf("Error parsing BSON: %v", err)
	}

	// Show page structure
	fmt.Println("=== Page Structure (JSON) ===")
	jsonData, err := json.MarshalIndent(pageData, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling page to JSON: %v", err)
	}

	// Truncate if too long
	if len(jsonData) > 10000 {
		fmt.Println(string(jsonData[:10000]))
		fmt.Printf("\n... [truncated, total size: %d bytes] ...\n\n", len(jsonData))
	} else {
		fmt.Println(string(jsonData))
		fmt.Println()
	}

	// Search for widgets in the page
	fmt.Println("=== Widgets Found ===")
	widgets := findAllWidgets(pageData)

	if len(widgets) == 0 {
		fmt.Println("No widgets found in this page.")
	} else {
		fmt.Printf("Total widgets/elements: %d\n\n", len(widgets))

		// Count by type
		typeCounts := make(map[string]int)
		for _, widget := range widgets {
			typeCounts[widget]++
		}

		// Sort by count (descending)
		type TypeCount struct {
			Type  string
			Count int
		}
		var sorted []TypeCount
		for t, c := range typeCounts {
			sorted = append(sorted, TypeCount{t, c})
		}

		// Sort function
		for i := 0; i < len(sorted); i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[j].Count > sorted[i].Count {
					sorted[i], sorted[j] = sorted[j], sorted[i]
				}
			}
		}

		fmt.Println("Widget types (sorted by frequency):")
		for i, item := range sorted {
			fmt.Printf("%3d. %-50s : %d occurrences\n", i+1, item.Type, item.Count)
		}
	}
}

// loadPageBSON loads the raw BSON data for a page from the mprcontents folder
func loadPageBSON(mprPath, pageID string) ([]byte, error) {
	// Check for MPR v2 (mprcontents folder)
	dir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(dir, "mprcontents")

	if stat, err := os.Stat(contentsDir); err != nil || !stat.IsDir() {
		return nil, fmt.Errorf("mprcontents folder not found (MPR v1 not supported by this tool)")
	}

	// Remove dashes from UUID for directory structure
	cleanID := strings.ReplaceAll(pageID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid page ID: %s", pageID)
	}

	// Build path: mprcontents/{first2}/{next2}/{uuid}.mxunit
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, pageID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read page file %s: %w", filePath, err)
	}

	return data, nil
}

func findAllWidgets(page interface{}) []string {
	widgets := []string{}

	// Convert to JSON for searching
	jsonData, err := json.Marshal(page)
	if err != nil {
		return widgets
	}

	// Parse JSON
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return widgets
	}

	// Recursively search for widgets
	searchForWidgets(data, &widgets)

	return widgets
}

func searchForWidgets(data interface{}, widgets *[]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this object has a $Type field
		if typeVal, ok := v["$Type"].(string); ok {
			*widgets = append(*widgets, typeVal)
		}
		// Recursively search in all fields
		for _, value := range v {
			searchForWidgets(value, widgets)
		}
	case []interface{}:
		// Recursively search in arrays
		for _, item := range v {
			searchForWidgets(item, widgets)
		}
	}
}
