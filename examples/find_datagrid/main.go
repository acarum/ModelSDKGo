// Example: Find DataGrid widget usage across all pages and snippets
//
// This example demonstrates how to search for DataGrid widgets
// across all pages and snippets in a Mendix project using only the modelsdk-go API.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/modelsdk-go"
)

const (
	DataGridWidgetType = "com.mendix.widget.web.datagrid.Datagrid"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: find_datagrid <path-to-mpr-file>")
		fmt.Println("Example: find_datagrid MyApp.mpr")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR file: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("=== DataGrid Widget Finder ===\n")
	fmt.Printf("Project: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n", reader.Version())

	version, err := reader.GetMendixVersion()
	if err == nil {
		fmt.Printf("Mendix Version: %s\n", version)
	}

	fmt.Printf("\nSearching for widget: %s\n\n", DataGridWidgetType)

	// Get all modules to map module IDs to names
	modules, err := reader.ListModules()
	if err != nil {
		fmt.Printf("Error listing modules: %v\n", err)
		os.Exit(1)
	}

	moduleMap := make(map[string]string)
	for _, m := range modules {
		moduleMap[string(m.ID)] = m.Name
	}

	// Search in pages
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📄 SCANNING PAGES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	pagesWithDataGrid := searchPagesForDataGrid(reader, moduleMap)

	// Search in snippets
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📝 SCANNING SNIPPETS")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	snippetsWithDataGrid := searchSnippetsForDataGrid(reader, moduleMap)

	// Print summary
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📊 SUMMARY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("Pages with DataGrid:    %d\n", len(pagesWithDataGrid))
	fmt.Printf("Snippets with DataGrid: %d\n", len(snippetsWithDataGrid))
	fmt.Printf("Total documents:        %d\n", len(pagesWithDataGrid)+len(snippetsWithDataGrid))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// DocumentInfo holds information about a document (page or snippet) containing DataGrid
type DocumentInfo struct {
	Name         string
	Module       string
	WidgetCount  int
	URL          string
	DocumentType string
}

func searchPagesForDataGrid(reader *modelsdk.Reader, moduleMap map[string]string) []DocumentInfo {
	var results []DocumentInfo

	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
		return results
	}

	fmt.Printf("Found %d pages to scan\n\n", len(pages))

	for i, page := range pages {
		// Show progress
		fmt.Printf("\r[%d/%d] Scanning pages...", i+1, len(pages))

		// Get the full page details
		fullPage, err := reader.GetPage(page.ID)
		if err != nil {
			continue
		}

		// Get module name
		moduleName := "Unknown"
		if name, ok := moduleMap[string(fullPage.ContainerID)]; ok {
			moduleName = name
		}

		// Search for DataGrid widgets in the page
		widgets := findDataGridWidgets(fullPage)

		if len(widgets) > 0 {
			results = append(results, DocumentInfo{
				Name:         fullPage.Name,
				Module:       moduleName,
				WidgetCount:  len(widgets),
				URL:          fullPage.URL,
				DocumentType: "Page",
			})
		}
	}

	fmt.Printf("\r") // Clear progress line

	// Print results
	if len(results) == 0 {
		fmt.Println("✗ No pages with DataGrid found")
	} else {
		fmt.Printf("✓ Found %d page(s) with DataGrid:\n\n", len(results))
		for i, doc := range results {
			fmt.Printf("%d. %s\n", i+1, doc.Name)
			fmt.Printf("   Module: %s\n", doc.Module)
			if doc.URL != "" {
				fmt.Printf("   URL: %s\n", doc.URL)
			}
			fmt.Printf("   DataGrid count: %d\n\n", doc.WidgetCount)
		}
	}

	return results
}

func searchSnippetsForDataGrid(reader *modelsdk.Reader, moduleMap map[string]string) []DocumentInfo {
	var results []DocumentInfo

	snippets, err := reader.ListSnippets()
	if err != nil {
		fmt.Printf("Error listing snippets: %v\n", err)
		return results
	}

	fmt.Printf("Found %d snippets to scan\n\n", len(snippets))

	for i, snippet := range snippets {
		// Show progress
		fmt.Printf("\r[%d/%d] Scanning snippets...", i+1, len(snippets))

		// Get module name
		moduleName := "Unknown"
		if name, ok := moduleMap[string(snippet.ContainerID)]; ok {
			moduleName = name
		}

		// Search for DataGrid widgets in the snippet
		// Note: ListSnippets returns snippets with Widget field populated
		widgets := findDataGridWidgets(snippet)

		if len(widgets) > 0 {
			results = append(results, DocumentInfo{
				Name:         snippet.Name,
				Module:       moduleName,
				WidgetCount:  len(widgets),
				DocumentType: "Snippet",
			})
		}
	}

	fmt.Printf("\r") // Clear progress line

	// Print results
	if len(results) == 0 {
		fmt.Println("✗ No snippets with DataGrid found")
	} else {
		fmt.Printf("✓ Found %d snippet(s) with DataGrid:\n\n", len(results))
		for i, doc := range results {
			fmt.Printf("%d. %s\n", i+1, doc.Name)
			fmt.Printf("   Module: %s\n", doc.Module)
			fmt.Printf("   DataGrid count: %d\n\n", doc.WidgetCount)
		}
	}

	return results
}

// findDataGridWidgets searches for DataGrid widgets in any object (page or snippet)
func findDataGridWidgets(obj interface{}) []map[string]interface{} {
	var results []map[string]interface{}

	// Convert object to JSON for inspection
	data, err := json.Marshal(obj)
	if err != nil {
		return results
	}

	var objMap map[string]interface{}
	if err := json.Unmarshal(data, &objMap); err != nil {
		return results
	}

	// Search recursively
	searchDataGridRecursive(objMap, &results)
	return results
}

// searchDataGridRecursive recursively searches for DataGrid widgets
func searchDataGridRecursive(obj interface{}, results *[]map[string]interface{}) {
	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this object is a DataGrid widget
		if typeField, ok := v["$Type"].(string); ok {
			// Check for exact match with DataGrid widget type
			if strings.Contains(typeField, "datagrid") ||
				strings.Contains(strings.ToLower(typeField), "datagrid") ||
				typeField == DataGridWidgetType {
				*results = append(*results, v)
			}
		}

		// Recursively search all fields
		for _, value := range v {
			searchDataGridRecursive(value, results)
		}

	case []interface{}:
		// Recursively search array elements
		for _, item := range v {
			searchDataGridRecursive(item, results)
		}
	}
}
