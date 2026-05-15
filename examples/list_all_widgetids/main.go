package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/anthropics/modelsdk-go"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: list_all_widgetids <mpr_file_path>")
		fmt.Println("Example: list_all_widgetids MyApp.mpr")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n\n", reader.Path())

	widgetIDs := make(map[string]int)

	// Search in snippets
	fmt.Println("=== Scanning Snippets ===")
	snippets, err := reader.ListSnippets()
	if err != nil {
		fmt.Printf("Error listing snippets: %v\n", err)
	} else {
		fmt.Printf("Processing %d snippets...\n", len(snippets))
		for i, snippet := range snippets {
			if i%10 == 0 {
				fmt.Printf("\r[%d/%d]", i, len(snippets))
			}

			bsonData, err := loadUnitBSON(mprPath, string(snippet.ID))
			if err != nil {
				continue
			}

			var snippetData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &snippetData); err != nil {
				continue
			}

			collectWidgetIDs(snippetData, widgetIDs)
		}
		fmt.Printf("\r[%d/%d] ✓\n", len(snippets), len(snippets))
	}

	// Search in pages
	fmt.Println("\n=== Scanning Pages ===")
	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
	} else {
		fmt.Printf("Processing %d pages...\n", len(pages))
		for i, page := range pages {
			if i%10 == 0 {
				fmt.Printf("\r[%d/%d]", i, len(pages))
			}

			bsonData, err := loadUnitBSON(mprPath, string(page.ID))
			if err != nil {
				continue
			}

			var pageData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &pageData); err != nil {
				continue
			}

			collectWidgetIDs(pageData, widgetIDs)
		}
		fmt.Printf("\r[%d/%d] ✓\n", len(pages), len(pages))
	}

	// Search in layouts
	fmt.Println("\n=== Scanning Layouts ===")
	layouts, err := reader.ListLayouts()
	if err != nil {
		fmt.Printf("Error listing layouts: %v\n", err)
	} else {
		fmt.Printf("Processing %d layouts...\n", len(layouts))
		for i, layout := range layouts {
			if i%10 == 0 {
				fmt.Printf("\r[%d/%d]", i, len(layouts))
			}

			bsonData, err := loadUnitBSON(mprPath, string(layout.ID))
			if err != nil {
				continue
			}

			var layoutData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &layoutData); err != nil {
				continue
			}

			collectWidgetIDs(layoutData, widgetIDs)
		}
		fmt.Printf("\r[%d/%d] ✓\n", len(layouts), len(layouts))
	}

	// Sort and display
	fmt.Println("\n=== CUSTOM WIDGET IDs ===")
	if len(widgetIDs) == 0 {
		fmt.Println("No custom widgets found.")
	} else {
		sortedIDs := sortByCount(widgetIDs)
		for _, item := range sortedIDs {
			fmt.Printf("  %-60s : %d occurrences\n", item.ID, item.Count)
		}
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total unique widget IDs: %d\n", len(widgetIDs))
}

type WidgetIDCount struct {
	ID    string
	Count int
}

func sortByCount(m map[string]int) []WidgetIDCount {
	var items []WidgetIDCount
	for id, count := range m {
		items = append(items, WidgetIDCount{id, count})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Count == items[j].Count {
			return items[i].ID < items[j].ID
		}
		return items[i].Count > items[j].Count
	})
	return items
}

func loadUnitBSON(mprPath, unitID string) ([]byte, error) {
	dir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(dir, "mprcontents")

	if stat, err := os.Stat(contentsDir); err != nil || !stat.IsDir() {
		return nil, fmt.Errorf("mprcontents folder not found")
	}

	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid unit ID: %s", unitID)
	}

	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read unit file %s: %w", filePath, err)
	}

	return data, nil
}

func collectWidgetIDs(data interface{}, widgetIDs map[string]int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	var parsed interface{}
	if err := json.Unmarshal(jsonData, &parsed); err != nil {
		return
	}

	searchForWidgetIDs(parsed, widgetIDs)
}

func searchForWidgetIDs(data interface{}, widgetIDs map[string]int) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a CustomWidget
		if typeVal, ok := v["$Type"].(string); ok && typeVal == "CustomWidgets$CustomWidget" {
			// Extract the widgetId from the Type field
			if typeField, ok := v["Type"].(map[string]interface{}); ok {
				if widgetID, ok := typeField["WidgetId"].(string); ok && widgetID != "" {
					widgetIDs[widgetID]++
				}
			}
		}

		// Recursively search in all fields
		for _, value := range v {
			searchForWidgetIDs(value, widgetIDs)
		}

	case []interface{}:
		// Recursively search in arrays
		for _, item := range v {
			searchForWidgetIDs(item, widgetIDs)
		}
	}
}
