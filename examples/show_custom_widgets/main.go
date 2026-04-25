package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthropics/modelsdk-go"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: show_custom_widgets <mpr_file_path> <page_name>")
		fmt.Println("Example: show_custom_widgets MyApp.mpr StateMachine_Details")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	pageName := os.Args[2]

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("Searching for page: %s\n\n", pageName)

	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
		os.Exit(1)
	}

	var pageID string
	foundPage := false

	for _, page := range pages {
		if page.Name == pageName {
			fmt.Printf("=== Page Found ===\n")
			fmt.Printf("Name: %s\n", page.Name)
			fmt.Printf("ID: %s\n\n", page.ID)
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

	fmt.Println("=== Custom Widgets in Page ===")
	customWidgets := findCustomWidgets(pageData)

	if len(customWidgets) == 0 {
		fmt.Println("No custom widgets found.")
	} else {
		fmt.Printf("Found %d custom widget(s):\n\n", len(customWidgets))
		for i, widget := range customWidgets {
			fmt.Printf("%d. Type: %s\n", i+1, widget.Type)
			if widget.Object != "" {
				fmt.Printf("   Object: %s\n", widget.Object)
			}
			if widget.ID != "" {
				fmt.Printf("   ID: %s\n", widget.ID)
			}
			fmt.Println()
		}
	}
}

type CustomWidgetInfo struct {
	Type   string
	Object string
	ID     string
}

func findCustomWidgets(data interface{}) []CustomWidgetInfo {
	var widgets []CustomWidgetInfo
	searchCustomWidgets(data, &widgets)
	return widgets
}

func searchCustomWidgets(data interface{}, widgets *[]CustomWidgetInfo) {
	switch v := data.(type) {
	case map[string]interface{}:
		if typeVal, ok := v["$Type"].(string); ok {
			if typeVal == "CustomWidgets$CustomWidget" {
				widget := CustomWidgetInfo{Type: typeVal}

				// Try to extract the object/type name
				if objVal, ok := v["object"].(string); ok {
					widget.Object = objVal
				}
				if objVal, ok := v["Object"].(string); ok {
					widget.Object = objVal
				}
				if typeObj, ok := v["type"]; ok {
					if typeMap, ok := typeObj.(map[string]interface{}); ok {
						if objName, ok := typeMap["qualifiedName"].(string); ok {
							widget.Object = objName
						}
					}
				}

				// Try to extract ID
				if idVal, ok := v["$ID"]; ok {
					widget.ID = fmt.Sprintf("%v", idVal)
				}

				*widgets = append(*widgets, widget)
			}
		}

		// Recursively search
		for _, value := range v {
			searchCustomWidgets(value, widgets)
		}
	case []interface{}:
		for _, item := range v {
			searchCustomWidgets(item, widgets)
		}
	}
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
