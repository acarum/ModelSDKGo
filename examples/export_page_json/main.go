package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthropics/modelsdk-go"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: export_page_json <mpr_file_path> <page_name> <output_json_file>")
		fmt.Println("Example: export_page_json MyApp.mpr StateMachine_Details output.json")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	pageName := os.Args[2]
	outputFile := os.Args[3]

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("Searching for page: %s\n", pageName)

	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
		os.Exit(1)
	}

	var pageID string
	foundPage := false

	for _, page := range pages {
		if page.Name == pageName {
			fmt.Printf("Page found: %s (ID: %s)\n", page.Name, page.ID)
			pageID = string(page.ID)
			foundPage = true
			break
		}
	}

	if !foundPage {
		fmt.Printf("Page '%s' not found!\n", pageName)
		os.Exit(1)
	}

	bsonData, err := loadPageBSON(mprPath, pageID)
	if err != nil {
		fmt.Printf("Error loading page BSON: %v\n", err)
		os.Exit(1)
	}

	var pageData map[string]interface{}
	if err := bson.Unmarshal(bsonData, &pageData); err != nil {
		fmt.Printf("Error parsing BSON: %v\n", err)
		os.Exit(1)
	}

	// Export to JSON
	jsonData, err := json.MarshalIndent(pageData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n✓ Exported to: %s\n", outputFile)
	fmt.Printf("File size: %d bytes (%.2f MB)\n", len(jsonData), float64(len(jsonData))/1024/1024)

	// Search for CustomWidget patterns
	fmt.Println("\nSearching for CustomWidget patterns...")
	customWidgetCount := strings.Count(string(jsonData), "CustomWidget")
	mxToSignalCount := strings.Count(string(jsonData), "MxToSignal")
	siemensCount := strings.Count(string(jsonData), "siemens")

	fmt.Printf("  'CustomWidget' appears: %d times\n", customWidgetCount)
	fmt.Printf("  'MxToSignal' appears: %d times\n", mxToSignalCount)
	fmt.Printf("  'siemens' appears: %d times\n", siemensCount)
}

func loadPageBSON(mprPath, pageID string) ([]byte, error) {
	dir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(dir, "mprcontents")

	if stat, err := os.Stat(contentsDir); err != nil || !stat.IsDir() {
		return nil, fmt.Errorf("mprcontents folder not found")
	}

	cleanID := strings.ReplaceAll(pageID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid page ID: %s", pageID)
	}

	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, pageID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return data, nil
}
