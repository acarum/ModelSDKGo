// Example: Find Signal Manager widget subscriptions across all pages and snippets
//
// This example demonstrates how to search for Signal Manager widgets
// and extract their signal subscriptions using only the modelsdk-go API.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/modelsdk-go"
)

const (
	SignalManagerWidgetID = "siemens.mxtosignal.MxToSignal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: find_signal_manager <path-to-mpr-file>")
		fmt.Println("Example: find_signal_manager MyApp.mpr")
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

	fmt.Printf("=== Signal Manager Widget Finder ===\n")
	fmt.Printf("Project: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n", reader.Version())

	version, err := reader.GetMendixVersion()
	if err == nil {
		fmt.Printf("Mendix Version: %s\n", version)
	}

	fmt.Printf("\nSearching for widget: %s\n\n", SignalManagerWidgetID)

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

	pageSubscriptions := searchPagesForSignalManager(reader, moduleMap)

	// Search in snippets
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📝 SCANNING SNIPPETS")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	snippetSubscriptions := searchSnippetsForSignalManager(reader, moduleMap)

	// Print summary
	allSubscriptions := append(pageSubscriptions, snippetSubscriptions...)

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📊 SIGNAL SUBSCRIPTIONS SUMMARY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if len(allSubscriptions) == 0 {
		fmt.Println("No signal subscriptions found")
	} else {
		fmt.Printf("\nTotal subscriptions: %d\n\n", len(allSubscriptions))

		// Group by module
		moduleGroups := make(map[string][]SignalSubscription)
		for _, sub := range allSubscriptions {
			moduleGroups[sub.Module] = append(moduleGroups[sub.Module], sub)
		}

		for module, subs := range moduleGroups {
			fmt.Printf("Module: %s (%d subscription(s))\n", module, len(subs))
			for i, sub := range subs {
				fmt.Printf("  %d. %s (%s)\n", i+1, sub.DocumentName, sub.DocumentType)
				fmt.Printf("     Signal Name: %s\n", sub.SignalName)
				if sub.AppName != "" {
					fmt.Printf("     App Name: %s\n", sub.AppName)
				}
				if sub.SubscriptionFilter != "" {
					fmt.Printf("     Filter: %s\n", sub.SubscriptionFilter)
				}
				fmt.Println()
			}
		}
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("Pages with Signal Manager:    %d\n", countDocuments(pageSubscriptions))
	fmt.Printf("Snippets with Signal Manager: %d\n", countDocuments(snippetSubscriptions))
	fmt.Printf("Total subscriptions:          %d\n", len(allSubscriptions))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// SignalSubscription holds information about a signal subscription
type SignalSubscription struct {
	Module             string
	DocumentName       string
	DocumentType       string
	SignalName         string
	AppName            string
	SubscriptionFilter string
}

func countDocuments(subscriptions []SignalSubscription) int {
	docs := make(map[string]bool)
	for _, sub := range subscriptions {
		key := sub.DocumentType + ":" + sub.DocumentName
		docs[key] = true
	}
	return len(docs)
}

func searchPagesForSignalManager(reader *modelsdk.Reader, moduleMap map[string]string) []SignalSubscription {
	var subscriptions []SignalSubscription

	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
		return subscriptions
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

		// Search for Signal Manager widgets and extract subscriptions
		pageSubs := extractSignalSubscriptions(fullPage, fullPage.Name, "Page", moduleName)
		subscriptions = append(subscriptions, pageSubs...)
	}

	fmt.Printf("\r") // Clear progress line

	// Print results
	uniqueDocs := countDocuments(subscriptions)
	if uniqueDocs == 0 {
		fmt.Println("✗ No pages with Signal Manager found")
	} else {
		fmt.Printf("✓ Found Signal Manager in %d page(s) with %d subscription(s)\n", uniqueDocs, len(subscriptions))
	}

	return subscriptions
}

func searchSnippetsForSignalManager(reader *modelsdk.Reader, moduleMap map[string]string) []SignalSubscription {
	var subscriptions []SignalSubscription

	snippets, err := reader.ListSnippets()
	if err != nil {
		fmt.Printf("Error listing snippets: %v\n", err)
		return subscriptions
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

		// Search for Signal Manager widgets and extract subscriptions
		snippetSubs := extractSignalSubscriptions(snippet, snippet.Name, "Snippet", moduleName)
		subscriptions = append(subscriptions, snippetSubs...)
	}

	fmt.Printf("\r") // Clear progress line

	// Print results
	uniqueDocs := countDocuments(subscriptions)
	if uniqueDocs == 0 {
		fmt.Println("✗ No snippets with Signal Manager found")
	} else {
		fmt.Printf("✓ Found Signal Manager in %d snippet(s) with %d subscription(s)\n", uniqueDocs, len(subscriptions))
	}

	return subscriptions
}

// extractSignalSubscriptions finds Signal Manager widgets and extracts their subscriptions
func extractSignalSubscriptions(obj interface{}, docName, docType, module string) []SignalSubscription {
	var subscriptions []SignalSubscription

	// Convert object to JSON for inspection
	data, err := json.Marshal(obj)
	if err != nil {
		return subscriptions
	}

	var objMap map[string]interface{}
	if err := json.Unmarshal(data, &objMap); err != nil {
		return subscriptions
	}

	// Find all Signal Manager widgets
	widgets := findSignalManagerWidgets(objMap)

	// Extract subscriptions from each widget
	for _, widget := range widgets {
		// Extract all properties from the widget
		props := extractWidgetProperties(widget)

		// Try to find signal subscription data
		// Look for patterns: SignalName, AppName, SubscriptionFilter
		sub := SignalSubscription{
			Module:       module,
			DocumentName: docName,
			DocumentType: docType,
		}

		// Search for signal-related properties
		for key, value := range props {
			keyLower := strings.ToLower(key)

			if strings.Contains(keyLower, "signalname") || strings.HasSuffix(keyLower, "_sn") {
				sub.SignalName = value
			}
			if strings.Contains(keyLower, "appname") {
				sub.AppName = value
			}
			if strings.Contains(keyLower, "subscriptionfilter") || strings.Contains(keyLower, "filter") {
				sub.SubscriptionFilter = value
			}
		}

		// Only add if we found at least a signal name
		if sub.SignalName != "" {
			subscriptions = append(subscriptions, sub)
		}
	}

	return subscriptions
}

// findSignalManagerWidgets searches for Signal Manager widgets
func findSignalManagerWidgets(data interface{}) []map[string]interface{} {
	var widgets []map[string]interface{}
	searchForSignalManager(data, &widgets)
	return widgets
}

func searchForSignalManager(data interface{}, widgets *[]map[string]interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a CustomWidget with Signal Manager widgetId
		if typeField, ok := v["$Type"].(string); ok {
			if typeField == "CustomWidgets$CustomWidget" {
				// Check the Type.WidgetId field
				if typeObj, ok := v["Type"].(map[string]interface{}); ok {
					if widgetId, ok := typeObj["WidgetId"].(string); ok {
						if widgetId == SignalManagerWidgetID {
							*widgets = append(*widgets, v)
							return // Found it, no need to recurse deeper
						}
					}
				}
			}
		}

		// Recursively search all fields
		for _, value := range v {
			searchForSignalManager(value, widgets)
		}

	case []interface{}:
		// Recursively search array elements
		for _, item := range v {
			searchForSignalManager(item, widgets)
		}
	}
}

// extractWidgetProperties recursively extracts all property values from a widget
func extractWidgetProperties(widget map[string]interface{}) map[string]string {
	props := make(map[string]string)
	extractPropertiesRecursive(widget, "", props)
	return props
}

func extractPropertiesRecursive(data interface{}, path string, props map[string]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check for PrimitiveValue which contains actual values
		if primVal, ok := v["PrimitiveValue"]; ok {
			if strVal, ok := primVal.(string); ok && strVal != "" {
				if path != "" {
					props[path] = strVal
				}
			}
		}

		// Recursively process all fields
		for key, value := range v {
			newPath := key
			if path != "" {
				newPath = path + "." + key
			}
			extractPropertiesRecursive(value, newPath, props)
		}

	case []interface{}:
		// Process array elements
		for i, item := range v {
			newPath := fmt.Sprintf("%s[%d]", path, i)
			extractPropertiesRecursive(item, newPath, props)
		}

	case string:
		// Store string values
		if path != "" && v != "" {
			props[path] = v
		}
	}
}
