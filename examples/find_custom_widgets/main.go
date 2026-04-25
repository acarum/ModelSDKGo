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
	if len(os.Args) < 3 {
		fmt.Println("Usage: find_custom_widgets <mpr_file_path> <widget_id>")
		fmt.Println("Example: find_custom_widgets MyApp.mpr siemens.mxtosignal.MxToSignal")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	widgetID := os.Args[2]

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("\n=== Searching for custom widgets with widgetId: %s ===\n\n", widgetID)

	totalFound := 0

	// Search in snippets
	fmt.Println("=== Searching Snippets ===")
	snippets, err := reader.ListSnippets()
	if err != nil {
		fmt.Printf("Error listing snippets: %v\n", err)
	} else {
		fmt.Printf("Scanning %d snippets...\n\n", len(snippets))
		for i, snippet := range snippets {
			fmt.Printf("\r[%d/%d] Scanning snippet: %s", i+1, len(snippets), snippet.Name)

			bsonData, err := loadUnitBSON(mprPath, string(snippet.ID))
			if err != nil {
				continue
			}

			var snippetData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &snippetData); err != nil {
				continue
			}

			widgets := findWidgetsByWidgetID(snippetData, widgetID)
			if len(widgets) > 0 {
				fmt.Printf("\n\n✓ Snippet: %s\n", snippet.Name)
				fmt.Printf("  ID: %s\n", snippet.ID)
				fmt.Printf("  Found %d widget(s) with widgetId '%s':\n", len(widgets), widgetID)
				for j, widget := range widgets {
					widgetName := "[unnamed]"
					if name, ok := widget["Name"].(string); ok && name != "" {
						widgetName = name
					}
					fmt.Printf("    %d. Widget Name: %s\n", j+1, widgetName)

					// Extract AppName and Signal Name from Properties
					printSignalSubscriptions(widget)
				}
				fmt.Println()
				totalFound += len(widgets)
			}
		}
		fmt.Println()
	}

	// Search in pages
	fmt.Println("\n=== Searching Pages ===")
	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
	} else {
		fmt.Printf("Scanning %d pages...\n\n", len(pages))
		for i, page := range pages {
			fmt.Printf("\r[%d/%d] Scanning page: %s", i+1, len(pages), page.Name)

			bsonData, err := loadUnitBSON(mprPath, string(page.ID))
			if err != nil {
				continue
			}

			var pageData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &pageData); err != nil {
				continue
			}

			widgets := findWidgetsByWidgetID(pageData, widgetID)
			if len(widgets) > 0 {
				fmt.Printf("\n\n✓ Page: %s\n", page.Name)
				fmt.Printf("  ID: %s\n", page.ID)
				fmt.Printf("  Found %d widget(s) with widgetId '%s':\n", len(widgets), widgetID)
				for j, widget := range widgets {
					widgetName := "[unnamed]"
					if name, ok := widget["Name"].(string); ok && name != "" {
						widgetName = name
					}
					fmt.Printf("    %d. Widget Name: %s\n", j+1, widgetName)

					// Extract AppName and Signal Name from Properties
					printSignalSubscriptions(widget)
				}
				fmt.Println()
				totalFound += len(widgets)
			}
		}
		fmt.Println()
	}

	// Search in layouts
	fmt.Println("\n=== Searching Layouts ===")
	layouts, err := reader.ListLayouts()
	if err != nil {
		fmt.Printf("Error listing layouts: %v\n", err)
	} else {
		fmt.Printf("Scanning %d layouts...\n\n", len(layouts))
		for i, layout := range layouts {
			fmt.Printf("\r[%d/%d] Scanning layout: %s", i+1, len(layouts), layout.Name)

			bsonData, err := loadUnitBSON(mprPath, string(layout.ID))
			if err != nil {
				continue
			}

			var layoutData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &layoutData); err != nil {
				continue
			}

			widgets := findWidgetsByWidgetID(layoutData, widgetID)
			if len(widgets) > 0 {
				fmt.Printf("\n\n✓ Layout: %s\n", layout.Name)
				fmt.Printf("  ID: %s\n", layout.ID)
				fmt.Printf("  Found %d widget(s) with widgetId '%s':\n", len(widgets), widgetID)
				for j, widget := range widgets {
					widgetName := "[unnamed]"
					if name, ok := widget["Name"].(string); ok && name != "" {
						widgetName = name
					}
					fmt.Printf("    %d. Widget Name: %s\n", j+1, widgetName)

					// Extract AppName and Signal Name from Properties
					printSignalSubscriptions(widget)
				}
				fmt.Println()
				totalFound += len(widgets)
			}
		}
		fmt.Println()
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total widgets with widgetId '%s': %d\n", widgetID, totalFound)
}

// loadUnitBSON loads the raw BSON data for a unit from the mprcontents folder
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

// findWidgetsByWidgetID finds all custom widgets with the specified widgetId
func findWidgetsByWidgetID(data interface{}, widgetID string) []map[string]interface{} {
	var matchingWidgets []map[string]interface{}

	// Convert to JSON for searching
	jsonData, err := json.Marshal(data)
	if err != nil {
		return matchingWidgets
	}

	var parsed interface{}
	if err := json.Unmarshal(jsonData, &parsed); err != nil {
		return matchingWidgets
	}

	searchForMatchingWidgets(parsed, widgetID, &matchingWidgets)
	return matchingWidgets
}

func searchForMatchingWidgets(data interface{}, widgetID string, matchingWidgets *[]map[string]interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a CustomWidget
		if typeVal, ok := v["$Type"].(string); ok && typeVal == "CustomWidgets$CustomWidget" {
			// Check if it has the matching widgetId in the Type field
			if typeField, ok := v["Type"].(map[string]interface{}); ok {
				if wid, ok := typeField["WidgetId"].(string); ok && wid == widgetID {
					*matchingWidgets = append(*matchingWidgets, v)
				}
			}
		}

		// Recursively search in all fields
		for _, value := range v {
			searchForMatchingWidgets(value, widgetID, matchingWidgets)
		}

	case []interface{}:
		// Recursively search in arrays
		for _, item := range v {
			searchForMatchingWidgets(item, widgetID, matchingWidgets)
		}
	}
}

// printSignalSubscriptions extracts and prints AppName and Signal Name from widget properties
func printSignalSubscriptions(widget map[string]interface{}) {
	// Collect all primitive values recursively from the widget structure
	values := make([]string, 0)
	collectPrimitiveValues(widget, &values)

	var appName, signalName string

	// Search for AppName and Signal Name patterns in collected values
	for _, value := range values {
		valueLower := strings.ToLower(value)

		// Look for AppName pattern
		if strings.Contains(valueLower, "appname") || strings.HasSuffix(valueLower, "_appname") {
			appName = value
		}

		// Look for Signal Name pattern (SN suffix, or contains "signal" without "appname")
		if strings.HasSuffix(valueLower, "_sn") || (strings.Contains(valueLower, "signal") && !strings.Contains(valueLower, "appname")) {
			signalName = value
		}
	}

	if appName != "" || signalName != "" {
		if appName != "" {
			fmt.Printf("       AppName: %s\n", appName)
		}
		if signalName != "" {
			fmt.Printf("       Signal Name: %s\n", signalName)
		}
	} else {
		fmt.Printf("       [No AppName or Signal Name found in properties]\n")
	}
}

// collectPrimitiveValues recursively collects all string values from PrimitiveValue fields
func collectPrimitiveValues(data interface{}, values *[]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a PrimitiveValue
		if primVal, ok := v["PrimitiveValue"]; ok {
			if strVal, ok := primVal.(string); ok && strVal != "" {
				*values = append(*values, strVal)
			}
		}

		// Recursively search in all fields
		for _, value := range v {
			collectPrimitiveValues(value, values)
		}

	case []interface{}:
		// Recursively search in arrays
		for _, item := range v {
			collectPrimitiveValues(item, values)
		}
	}
}

// extractPropertyValue extracts the value from a property (PrimitiveValue, Expression, etc.)
func extractPropertyValue(propMap map[string]interface{}) string {
	// Try PrimitiveValue first
	if value, ok := propMap["PrimitiveValue"]; ok {
		if valueMap, ok := value.(map[string]interface{}); ok {
			if val, ok := valueMap["Value"]; ok {
				return fmt.Sprintf("%v", val)
			}
		}
	}

	// Try Expression
	if value, ok := propMap["Expression"]; ok {
		if valueMap, ok := value.(map[string]interface{}); ok {
			if val, ok := valueMap["Value"]; ok {
				return fmt.Sprintf("%v", val)
			}
		}
	}

	// Try Microflow
	if value, ok := propMap["Microflow"]; ok {
		if valueMap, ok := value.(map[string]interface{}); ok {
			if val, ok := valueMap["QualifiedName"]; ok {
				return fmt.Sprintf("[Microflow: %v]", val)
			}
		}
	}

	return "[value not found]"
}
