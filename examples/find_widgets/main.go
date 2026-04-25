// Example: Find widget usage across all pages
//
// This example demonstrates how to search for a specific widget type
// across all pages in a Mendix project and display its properties.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/modelsdk-go"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: find_widgets <path-to-mpr-file> <widget-type>")
		fmt.Println("Example: find_widgets MyApp.mpr DataGrid")
		fmt.Println("         find_widgets MyApp.mpr Button")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	widgetType := os.Args[2]

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR file: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n", reader.Version())

	version, err := reader.GetMendixVersion()
	if err == nil {
		fmt.Printf("Mendix Version: %s\n", version)
	}

	fmt.Printf("\n=== Searching for widget type: %s ===\n", widgetType)
	fmt.Println("(Excluding marketplace modules)\n")

	// Get all modules to map module IDs to names
	modules, err := reader.ListModules()
	if err != nil {
		fmt.Printf("Error listing modules: %v\n", err)
		os.Exit(1)
	}

	moduleMap := make(map[string]string)
	marketplaceModules := make(map[string]bool)
	for _, m := range modules {
		moduleMap[string(m.ID)] = m.Name
		if isMarketplaceModule(m.Name) {
			marketplaceModules[m.Name] = true
		}
	}

	// Get all pages
	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Scanning %d pages...\n\n", len(pages))

	foundCount := 0
	skippedPages := 0

	for i, page := range pages {
		// Get the full page details
		fullPage, err := reader.GetPage(page.ID)
		if err != nil {
			continue
		}

		// Get module name and check if it's marketplace
		moduleName := "Unknown"
		if name, ok := moduleMap[string(fullPage.ContainerID)]; ok {
			moduleName = name
		}

		// Show progress
		fmt.Printf("\r[%d/%d] Scanning page: %s (%s)", i+1, len(pages), fullPage.Name, moduleName)

		// Skip marketplace modules
		if isMarketplaceModule(moduleName) {
			skippedPages++
			continue
		}

		// Search for widgets in the page
		widgets := findWidgetsByType(fullPage, widgetType)

		if len(widgets) > 0 {
			fmt.Printf("\n\n") // Clear progress line
			fmt.Printf("Page: %s\n", fullPage.Name)
			fmt.Printf("Module: %s\n", moduleName)
			if fullPage.URL != "" {
				fmt.Printf("URL: %s\n", fullPage.URL)
			}
			fmt.Printf("Found %d widget(s):\n", len(widgets))

			for i, widget := range widgets {
				fmt.Printf("\n  Widget #%d:\n", i+1)
				fmt.Printf("    Type: %s\n", widget["$Type"])

				// Print all properties
				fmt.Printf("    Properties:\n")
				for key, value := range widget {
					if key == "$Type" {
						continue
					}

					// Format the value for display
					valueStr := formatValue(value)
					if valueStr != "" {
						fmt.Printf("      %s: %s\n", key, valueStr)
					}
				}
			}

			fmt.Println()
			fmt.Println(strings.Repeat("-", 80))
			fmt.Println()

			foundCount++
		}
	}

	// Also search in layouts
	layouts, err := reader.ListLayouts()
	skippedLayouts := 0
	if err == nil {
		fmt.Printf("\n\nScanning %d layouts...\n\n", len(layouts))

		for i, layout := range layouts {
			fullLayout, err := reader.GetLayout(layout.ID)
			if err != nil {
				continue
			}

			// Get module name and check if it's marketplace
			moduleName := "Unknown"
			if name, ok := moduleMap[string(fullLayout.ContainerID)]; ok {
				moduleName = name
			}

			// Show progress
			fmt.Printf("\r[%d/%d] Scanning layout: %s (%s)", i+1, len(layouts), fullLayout.Name, moduleName)

			// Skip marketplace modules
			if isMarketplaceModule(moduleName) {
				skippedLayouts++
				continue
			}

			widgets := findWidgetsByType(fullLayout, widgetType)

			if len(widgets) > 0 {
				fmt.Printf("\n\n") // Clear progress line
				fmt.Printf("Layout: %s\n", fullLayout.Name)
				fmt.Printf("Module: %s\n", moduleName)
				fmt.Printf("Layout Type: %s\n", fullLayout.LayoutType)
				fmt.Printf("Found %d widget(s):\n", len(widgets))

				for i, widget := range widgets {
					fmt.Printf("\n  Widget #%d:\n", i+1)
					fmt.Printf("    Type: %s\n", widget["$Type"])

					fmt.Printf("    Properties:\n")
					for key, value := range widget {
						if key == "$Type" {
							continue
						}

						valueStr := formatValue(value)
						if valueStr != "" {
							fmt.Printf("      %s: %s\n", key, valueStr)
						}
					}
				}

				fmt.Println()
				fmt.Println(strings.Repeat("-", 80))
				fmt.Println()

				foundCount++
			}
		}
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total modules: %d\n", len(modules))
	fmt.Printf("Marketplace modules skipped: %d\n", len(marketplaceModules))
	fmt.Printf("Pages scanned: %d (skipped: %d)\n", len(pages)-skippedPages, skippedPages)
	fmt.Printf("Layouts scanned: %d (skipped: %d)\n", len(layouts)-skippedLayouts, skippedLayouts)
	fmt.Printf("Total pages/layouts with '%s' widget: %d\n", widgetType, foundCount)
}

// isMarketplaceModule checks if a module is from the marketplace
func isMarketplaceModule(moduleName string) bool {
	// Common marketplace module prefixes/patterns
	marketplaceIndicators := []string{
		"Marketplace",
		"Community",
		"AppStore",
		"Atlas",
		"Administration",
		"System",
		"CommunityCommons",
		"NanoflowCommons",
		"DataWidgets",
		"WebActions",
	}

	for _, indicator := range marketplaceIndicators {
		if strings.HasPrefix(moduleName, indicator) {
			return true
		}
	}

	return false
}

// findWidgetsByType recursively searches for widgets of a specific type
func findWidgetsByType(obj interface{}, widgetType string) []map[string]interface{} {
	var results []map[string]interface{}

	// Convert object to map for inspection
	data, err := json.Marshal(obj)
	if err != nil {
		return results
	}

	var objMap map[string]interface{}
	if err := json.Unmarshal(data, &objMap); err != nil {
		return results
	}

	results = searchWidgets(objMap, widgetType, results)
	return results
}

// searchWidgets recursively searches through a map structure
func searchWidgets(obj interface{}, widgetType string, results []map[string]interface{}) []map[string]interface{} {
	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this object has a $Type field matching our search
		if typeField, ok := v["$Type"].(string); ok {
			if strings.Contains(strings.ToLower(typeField), strings.ToLower(widgetType)) {
				results = append(results, v)
			}
		}

		// Recursively search all fields
		for _, value := range v {
			results = searchWidgets(value, widgetType, results)
		}

	case []interface{}:
		// Recursively search array elements
		for _, item := range v {
			results = searchWidgets(item, widgetType, results)
		}
	}

	return results
}

// formatValue formats a value for display
func formatValue(value interface{}) string {
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
	case map[string]interface{}:
		// For nested objects, show type if available
		if typeField, ok := v["$Type"].(string); ok {
			return fmt.Sprintf("{%s}", typeField)
		}
		return "{object}"
	case []interface{}:
		if len(v) == 0 {
			return ""
		}
		return fmt.Sprintf("[%d items]", len(v))
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}
