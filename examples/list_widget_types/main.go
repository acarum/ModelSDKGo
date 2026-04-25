package main

import (
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
		fmt.Println("Usage: list_all_widget_types <mpr_file_path>")
		fmt.Println("Example: list_all_widget_types MyApp.mpr")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n", reader.Version())

	version, err := reader.GetMendixVersion()
	if err == nil {
		fmt.Printf("Mendix Version: %s\n", version)
	}
	fmt.Println()

	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
		os.Exit(1)
	}

	layouts, err := reader.ListLayouts()
	if err != nil {
		fmt.Printf("Error listing layouts: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Scanning %d pages and %d layouts...\n\n", len(pages), len(layouts))

	widgetTypes := make(map[string]int)

	// Scan all pages
	for i, page := range pages {
		if i%20 == 0 {
			fmt.Printf("Processing pages: %d/%d\r", i, len(pages))
		}

		bsonData, err := loadPageBSON(mprPath, string(page.ID))
		if err != nil {
			continue
		}

		var pageData map[string]interface{}
		if err := bson.Unmarshal(bsonData, &pageData); err != nil {
			continue
		}

		collectTypes(pageData, widgetTypes)
	}
	fmt.Printf("Processing pages: %d/%d ✓\n", len(pages), len(pages))

	// Scan all layouts
	for i, layout := range layouts {
		if i%10 == 0 {
			fmt.Printf("Processing layouts: %d/%d\r", i, len(layouts))
		}

		bsonData, err := loadPageBSON(mprPath, string(layout.ID))
		if err != nil {
			continue
		}

		var layoutData map[string]interface{}
		if err := bson.Unmarshal(bsonData, &layoutData); err != nil {
			continue
		}

		collectTypes(layoutData, widgetTypes)
	}
	fmt.Printf("Processing layouts: %d/%d ✓\n\n", len(layouts), len(layouts))

	// Categorize widget types
	customWidgets := make(map[string]int)
	mendixWidgets := make(map[string]int)

	for widgetType, count := range widgetTypes {
		if strings.HasPrefix(widgetType, "Forms$") ||
			strings.HasPrefix(widgetType, "Texts$") ||
			strings.HasPrefix(widgetType, "DataTypes$") ||
			strings.HasPrefix(widgetType, "DomainModels$") ||
			strings.HasPrefix(widgetType, "CustomWidgets$") ||
			strings.HasPrefix(widgetType, "Pages$") {
			mendixWidgets[widgetType] = count
		} else {
			customWidgets[widgetType] = count
		}
	}

	// Sort and display custom widgets
	fmt.Println("=== CUSTOM WIDGETS ===")
	if len(customWidgets) == 0 {
		fmt.Println("No custom widgets found.")
	} else {
		sortedCustom := sortByCount(customWidgets)
		for _, item := range sortedCustom {
			fmt.Printf("  %s: %d occurrences\n", item.Type, item.Count)
		}
	}

	fmt.Println()
	fmt.Printf("=== SUMMARY ===\n")
	fmt.Printf("Total unique types: %d\n", len(widgetTypes))
	fmt.Printf("Custom widgets: %d\n", len(customWidgets))
	fmt.Printf("Mendix built-in: %d\n", len(mendixWidgets))
}

type TypeCount struct {
	Type  string
	Count int
}

func sortByCount(m map[string]int) []TypeCount {
	var items []TypeCount
	for t, c := range m {
		items = append(items, TypeCount{t, c})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})
	return items
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

func collectTypes(data interface{}, types map[string]int) {
	switch v := data.(type) {
	case map[string]interface{}:
		if typeVal, ok := v["$Type"].(string); ok {
			types[typeVal]++
		}
		for _, value := range v {
			collectTypes(value, types)
		}
	case []interface{}:
		for _, item := range v {
			collectTypes(item, types)
		}
	}
}
