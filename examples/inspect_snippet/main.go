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
		fmt.Println("Usage: inspect_snippet <mpr_file_path> <snippet_name>")
		fmt.Println("Example: inspect_snippet MyApp.mpr MySnippet")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	snippetName := os.Args[2]

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		log.Fatalf("Error opening MPR: %v", err)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("Searching for snippet: %s\n\n", snippetName)

	// Get all snippets
	snippets, err := reader.ListSnippets()
	if err != nil {
		log.Fatalf("Error listing snippets: %v", err)
	}

	// Find the specific snippet
	var snippetID string
	foundSnippet := false

	for _, snippet := range snippets {
		if snippet.Name == snippetName {
			fmt.Printf("=== Snippet Found ===\n")
			fmt.Printf("Name: %s\n", snippet.Name)
			fmt.Printf("ID: %s\n", snippet.ID)
			fmt.Printf("Type: %s\n", snippet.TypeName)
			if snippet.EntityID != "" {
				fmt.Printf("Entity ID: %s\n", snippet.EntityID)
			}
			if snippet.Documentation != "" {
				fmt.Printf("Documentation: %s\n", snippet.Documentation)
			}
			fmt.Println()
			snippetID = string(snippet.ID)
			foundSnippet = true
			break
		}
	}

	if !foundSnippet {
		fmt.Printf("Snippet '%s' not found!\n", snippetName)
		fmt.Printf("\nAvailable snippets:\n")
		for i, snippet := range snippets {
			fmt.Printf("%3d. %s (ID: %s)\n", i+1, snippet.Name, snippet.ID)
			if i >= 19 {
				fmt.Printf("... and %d more\n", len(snippets)-20)
				break
			}
		}
		os.Exit(1)
	}

	// Read the raw BSON content directly from the .mxunit file
	bsonData, err := loadSnippetBSON(mprPath, snippetID)
	if err != nil {
		log.Fatalf("Error loading snippet BSON: %v", err)
	}

	// Parse BSON to generic map
	var snippetData map[string]interface{}
	if err := bson.Unmarshal(bsonData, &snippetData); err != nil {
		log.Fatalf("Error parsing BSON: %v", err)
	}

	// Show snippet structure
	fmt.Println("=== Snippet Structure (JSON) ===")
	jsonData, err := json.MarshalIndent(snippetData, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling snippet to JSON: %v", err)
	}

	// Truncate if too long
	if len(jsonData) > 10000 {
		fmt.Println(string(jsonData[:10000]))
		fmt.Printf("\n... [truncated, total size: %d bytes] ...\n\n", len(jsonData))
	} else {
		fmt.Println(string(jsonData))
		fmt.Println()
	}

	// Search for widgets in the snippet
	fmt.Println("=== Widgets Found ===")
	widgets := findAllWidgets(snippetData)

	if len(widgets) == 0 {
		fmt.Println("No widgets found in this snippet.")
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

	// List custom widgets specifically
	fmt.Println("\n=== Custom Widgets (Detailed) ===")
	customWidgets := findCustomWidgets(snippetData)
	if len(customWidgets) == 0 {
		fmt.Println("No custom widgets found in this snippet.")
	} else {
		fmt.Printf("Found %d custom widget instance(s):\n\n", len(customWidgets))
		for i, widget := range customWidgets {
			fmt.Printf("Custom Widget #%d:\n", i+1)

			// Print all top-level properties
			for key, value := range widget {
				if key == "$ID" {
					// Print ID in a readable format
					if idMap, ok := value.(map[string]interface{}); ok {
						if data, ok := idMap["Data"].(string); ok {
							fmt.Printf("  %s: %s\n", key, data)
						}
					}
					continue
				}

				if key == "$Type" {
					fmt.Printf("  %s: %s\n", key, value)
					continue
				}

				// Special handling for Type field to show widgetId
				if key == "Type" {
					if typeMap, ok := value.(map[string]interface{}); ok {
						fmt.Printf("  %s: {CustomWidgets$CustomWidgetType}\n", key)
						if widgetId, ok := typeMap["WidgetId"].(string); ok && widgetId != "" {
							fmt.Printf("    widgetId: %s\n", widgetId)
						}
					}
					continue
				}

				// Format value based on type
				formattedValue := formatPropertyValue(value, "    ")
				if formattedValue != "" {
					fmt.Printf("  %s: %s\n", key, formattedValue)
				}
			}

			// Print Object.Properties details if present
			if obj, ok := widget["Object"].(map[string]interface{}); ok {
				if props, ok := obj["Properties"].([]interface{}); ok {
					fmt.Printf("\n  Widget Properties (%d):\n", len(props))
					for j, prop := range props {
						if propMap, ok := prop.(map[string]interface{}); ok {
							fmt.Printf("    Property #%d:\n", j+1)
							printWidgetProperty(propMap, "      ")
						}
					}
				}
			}

			fmt.Println()
			fmt.Println(strings.Repeat("-", 80))
			fmt.Println()
		}
	}
}

// loadSnippetBSON loads the raw BSON data for a snippet from the mprcontents folder
func loadSnippetBSON(mprPath, snippetID string) ([]byte, error) {
	// Check for MPR v2 (mprcontents folder)
	dir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(dir, "mprcontents")

	if stat, err := os.Stat(contentsDir); err != nil || !stat.IsDir() {
		return nil, fmt.Errorf("mprcontents folder not found (MPR v1 not supported by this tool)")
	}

	// Remove dashes from UUID for directory structure
	cleanID := strings.ReplaceAll(snippetID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid snippet ID: %s", snippetID)
	}

	// Build path: mprcontents/{first2}/{next2}/{uuid}.mxunit
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, snippetID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read snippet file %s: %w", filePath, err)
	}

	return data, nil
}

func findAllWidgets(snippet interface{}) []string {
	widgets := []string{}

	// Convert to JSON for searching
	jsonData, err := json.Marshal(snippet)
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

// findCustomWidgets finds all custom widgets (those with a dot in the type name)
func findCustomWidgets(snippet interface{}) []map[string]interface{} {
	var customWidgets []map[string]interface{}

	// Convert to JSON for searching
	jsonData, err := json.Marshal(snippet)
	if err != nil {
		return customWidgets
	}

	// Parse JSON
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return customWidgets
	}

	// Recursively search for custom widgets
	searchForCustomWidgets(data, &customWidgets)

	return customWidgets
}

func searchForCustomWidgets(data interface{}, customWidgets *[]map[string]interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this object is a custom widget (has $Type with CustomWidgets$CustomWidget)
		if typeVal, ok := v["$Type"].(string); ok {
			if typeVal == "CustomWidgets$CustomWidget" {
				*customWidgets = append(*customWidgets, v)
			}
		}
		// Recursively search in all fields
		for _, value := range v {
			searchForCustomWidgets(value, customWidgets)
		}
	case []interface{}:
		// Recursively search in arrays
		for _, item := range v {
			searchForCustomWidgets(item, customWidgets)
		}
	}
}

// formatPropertyValue formats a property value for display
func formatPropertyValue(value interface{}, indent string) string {
	switch v := value.(type) {
	case string:
		if v == "" {
			return ""
		}
		if len(v) > 100 {
			return v[:100] + "..."
		}
		return v
	case bool:
		return fmt.Sprintf("%v", v)
	case float64:
		return fmt.Sprintf("%.0f", v)
	case int:
		return fmt.Sprintf("%d", v)
	case nil:
		return "null"
	case map[string]interface{}:
		// Check if it has a $Type field
		if typeVal, ok := v["$Type"].(string); ok {
			return fmt.Sprintf("{%s}", typeVal)
		}
		return "{object}"
	case []interface{}:
		if len(v) == 0 {
			return "[]"
		}
		return fmt.Sprintf("[%d items]", len(v))
	default:
		return fmt.Sprintf("%v", v)
	}
}

// printWidgetProperty prints detailed information about a widget property
func printWidgetProperty(prop map[string]interface{}, indent string) {
	// Print property type pointer if available
	if typePtr, ok := prop["TypePointer"].(map[string]interface{}); ok {
		if data, ok := typePtr["Data"].(string); ok {
			fmt.Printf("%sTypePointer: %s\n", indent, data)
		}
	}

	// Print value details
	if value, ok := prop["Value"].(map[string]interface{}); ok {
		fmt.Printf("%sValue:\n", indent)

		// Check for common value fields
		if expr, ok := value["Expression"].(string); ok && expr != "" {
			fmt.Printf("%s  Expression: %s\n", indent, strings.TrimSpace(expr))
		}
		if prim, ok := value["PrimitiveValue"].(string); ok && prim != "" {
			fmt.Printf("%s  PrimitiveValue: %s\n", indent, prim)
		}
		if mf, ok := value["Microflow"].(string); ok && mf != "" {
			fmt.Printf("%s  Microflow: %s\n", indent, mf)
		}
		if nf, ok := value["Nanoflow"].(string); ok && nf != "" {
			fmt.Printf("%s  Nanoflow: %s\n", indent, nf)
		}
		if form, ok := value["Form"].(string); ok && form != "" {
			fmt.Printf("%s  Form: %s\n", indent, form)
		}

		// Check for DataSource with MicroflowSettings
		if ds, ok := value["DataSource"].(map[string]interface{}); ok {
			if dsType, ok := ds["$Type"].(string); ok {
				fmt.Printf("%s  DataSource Type: %s\n", indent, dsType)
			}
			if mfSettings, ok := ds["MicroflowSettings"].(map[string]interface{}); ok {
				if mf, ok := mfSettings["Microflow"].(string); ok && mf != "" {
					fmt.Printf("%s  DataSource Microflow: %s\n", indent, mf)
				}
			}
		}

		// Check for nested Objects
		if objects, ok := value["Objects"].([]interface{}); ok && len(objects) > 0 {
			fmt.Printf("%s  Objects: [%d items]\n", indent, len(objects))
			for i, obj := range objects {
				if objMap, ok := obj.(map[string]interface{}); ok {
					fmt.Printf("%s    Object #%d:\n", indent, i+1)
					if objType, ok := objMap["$Type"].(string); ok {
						fmt.Printf("%s      Type: %s\n", indent, objType)
					}
					// Print object properties
					if objProps, ok := objMap["Properties"].([]interface{}); ok {
						for j, objProp := range objProps {
							if objPropMap, ok := objProp.(map[string]interface{}); ok {
								fmt.Printf("%s      Property #%d:\n", indent, j+1)
								if objValue, ok := objPropMap["Value"].(map[string]interface{}); ok {
									if primVal, ok := objValue["PrimitiveValue"].(string); ok && primVal != "" {
										fmt.Printf("%s        PrimitiveValue: %s\n", indent, primVal)
									}
									if exprVal, ok := objValue["Expression"].(string); ok && exprVal != "" {
										fmt.Printf("%s        Expression: %s\n", indent, strings.TrimSpace(exprVal))
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
