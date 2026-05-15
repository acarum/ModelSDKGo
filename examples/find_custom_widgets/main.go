package main

import (
	"bufio"
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
		fmt.Println("Usage:")
		fmt.Println("  Find mode:    find_custom_widgets <mpr_file_path> <widget_id> [--dump-json]")
		fmt.Println("  Replace mode: find_custom_widgets <mpr_file_path> <source_widget_id> --replace <dest_widget_id>")
		fmt.Println("\nExamples:")
		fmt.Println("  find_custom_widgets MyApp.mpr siemens.mxtosignal.MxToSignal")
		fmt.Println("  find_custom_widgets MyApp.mpr dependencygraph --dump-json")
		fmt.Println("  find_custom_widgets MyApp.mpr siemens.DependencyGraph.DependencyGraph --replace siemens.dependencyGraph.DependencyGraph")
		fmt.Println("\nNote:")
		fmt.Println("  Replace mode uses binary byte replacement to preserve BSON structure.")
		fmt.Println("  Source and destination widget IDs must have the same length.")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	sourceWidgetID := os.Args[2]

	// Check modes and flags
	replaceMode := false
	dumpJSON := false
	var destWidgetID string

	for i := 3; i < len(os.Args); i++ {
		if os.Args[i] == "--replace" && i+1 < len(os.Args) {
			replaceMode = true
			destWidgetID = os.Args[i+1]
			i++ // skip next arg
		} else if os.Args[i] == "--dump-json" {
			dumpJSON = true
		}
	}

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())

	if replaceMode {
		fmt.Printf("\n=== REPLACE MODE ===\n")
		fmt.Printf("Searching for widgets with widgetId: %s\n", sourceWidgetID)
		fmt.Printf("Will replace with: %s\n\n", destWidgetID)
		performReplace(reader, mprPath, sourceWidgetID, destWidgetID)
	} else {
		fmt.Printf("\n=== Searching for custom widgets with widgetId: %s ===\n", sourceWidgetID)
		if dumpJSON {
			fmt.Printf("JSON dump mode: ENABLED\n")
		}
		fmt.Println()
		performSearch(reader, mprPath, sourceWidgetID, dumpJSON)
	}
}

type UnitReplaceInfo struct {
	UnitType string
	UnitID   string
	UnitName string
	Widgets  []map[string]interface{}
}

func performSearch(reader *modelsdk.Reader, mprPath, widgetID string, dumpJSON bool) {
	totalFound := 0
	dumpedFiles := 0

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
					actualWidgetID := extractWidgetID(widget)
					fmt.Printf("    %d. Widget Name: %s\n", j+1, widgetName)
					fmt.Printf("       WidgetId: %s\n", actualWidgetID)

					// Extract AppName and Signal Name from Properties
					printSignalSubscriptions(widget)
				}

				// Dump JSON if requested
				if dumpJSON {
					if err := dumpUnitJSON(snippetData, "Snippet", snippet.Name); err != nil {
						fmt.Printf("  ✗ Error dumping JSON: %v\n", err)
					} else {
						fmt.Printf("  ✓ JSON saved to: SNIPPET_%s.json\n", sanitizeFilename(snippet.Name))
						dumpedFiles++
					}
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
					actualWidgetID := extractWidgetID(widget)
					fmt.Printf("    %d. Widget Name: %s\n", j+1, widgetName)
					fmt.Printf("       WidgetId: %s\n", actualWidgetID)

					// Extract AppName and Signal Name from Properties
					printSignalSubscriptions(widget)
				}

				// Dump JSON if requested
				if dumpJSON {
					if err := dumpUnitJSON(pageData, "Page", page.Name); err != nil {
						fmt.Printf("  ✗ Error dumping JSON: %v\n", err)
					} else {
						fmt.Printf("  ✓ JSON saved to: PAGE_%s.json\n", sanitizeFilename(page.Name))
						dumpedFiles++
					}
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
					actualWidgetID := extractWidgetID(widget)
					fmt.Printf("    %d. Widget Name: %s\n", j+1, widgetName)
					fmt.Printf("       WidgetId: %s\n", actualWidgetID)

					// Extract AppName and Signal Name from Properties
					printSignalSubscriptions(widget)
				}

				// Dump JSON if requested
				if dumpJSON {
					if err := dumpUnitJSON(layoutData, "Layout", layout.Name); err != nil {
						fmt.Printf("  ✗ Error dumping JSON: %v\n", err)
					} else {
						fmt.Printf("  ✓ JSON saved to: LAYOUT_%s.json\n", sanitizeFilename(layout.Name))
						dumpedFiles++
					}
				}

				fmt.Println()
				totalFound += len(widgets)
			}
		}
		fmt.Println()
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total widgets with widgetId '%s': %d\n", widgetID, totalFound)
	if dumpJSON && dumpedFiles > 0 {
		fmt.Printf("JSON files saved: %d\n", dumpedFiles)
	}
}

func performReplace(reader *modelsdk.Reader, mprPath, sourceWidgetID, destWidgetID string) {
	var unitsToReplace []UnitReplaceInfo

	// Search in snippets
	fmt.Println("=== Scanning Snippets ===")
	snippets, err := reader.ListSnippets()
	if err != nil {
		fmt.Printf("Error listing snippets: %v\n", err)
	} else {
		fmt.Printf("Scanning %d snippets...\n", len(snippets))
		for _, snippet := range snippets {
			bsonData, err := loadUnitBSON(mprPath, string(snippet.ID))
			if err != nil {
				continue
			}

			var snippetData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &snippetData); err != nil {
				continue
			}

			widgets := findWidgetsByWidgetID(snippetData, sourceWidgetID)
			if len(widgets) > 0 {
				unitsToReplace = append(unitsToReplace, UnitReplaceInfo{
					UnitType: "Snippet",
					UnitID:   string(snippet.ID),
					UnitName: snippet.Name,
					Widgets:  widgets,
				})
			}
		}
	}

	// Search in pages
	fmt.Println("\n=== Scanning Pages ===")
	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
	} else {
		fmt.Printf("Scanning %d pages...\n", len(pages))
		for _, page := range pages {
			bsonData, err := loadUnitBSON(mprPath, string(page.ID))
			if err != nil {
				continue
			}

			var pageData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &pageData); err != nil {
				continue
			}

			widgets := findWidgetsByWidgetID(pageData, sourceWidgetID)
			if len(widgets) > 0 {
				unitsToReplace = append(unitsToReplace, UnitReplaceInfo{
					UnitType: "Page",
					UnitID:   string(page.ID),
					UnitName: page.Name,
					Widgets:  widgets,
				})
			}
		}
	}

	// Search in layouts
	fmt.Println("\n=== Scanning Layouts ===")
	layouts, err := reader.ListLayouts()
	if err != nil {
		fmt.Printf("Error listing layouts: %v\n", err)
	} else {
		fmt.Printf("Scanning %d layouts...\n", len(layouts))
		for _, layout := range layouts {
			bsonData, err := loadUnitBSON(mprPath, string(layout.ID))
			if err != nil {
				continue
			}

			var layoutData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &layoutData); err != nil {
				continue
			}

			widgets := findWidgetsByWidgetID(layoutData, sourceWidgetID)
			if len(widgets) > 0 {
				unitsToReplace = append(unitsToReplace, UnitReplaceInfo{
					UnitType: "Layout",
					UnitID:   string(layout.ID),
					UnitName: layout.Name,
					Widgets:  widgets,
				})
			}
		}
	}

	if len(unitsToReplace) == 0 {
		fmt.Printf("\n✗ No widgets found with widgetId '%s'\n", sourceWidgetID)
		return
	}

	// Display summary
	totalWidgets := 0
	fmt.Printf("\n=== Found Widgets to Replace ===\n")
	for _, unit := range unitsToReplace {
		fmt.Printf("  %s: %s (%d widget(s))\n", unit.UnitType, unit.UnitName, len(unit.Widgets))
		totalWidgets += len(unit.Widgets)
	}
	fmt.Printf("\nTotal: %d widget(s) in %d unit(s)\n", totalWidgets, len(unitsToReplace))

	// Ask for confirmation
	fmt.Printf("\nReplace '%s' with '%s' in all found widgets? (yes/no): ", sourceWidgetID, destWidgetID)
	inputReader := bufio.NewReader(os.Stdin)
	response, _ := inputReader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response != "yes" && response != "y" {
		fmt.Println("Operation cancelled.")
		return
	}

	// Perform replacement using binary byte replacement
	fmt.Println("\n=== Performing Replacement ===")
	successCount := 0
	errorCount := 0

	// Convert patterns to bytes
	sourceBytes := []byte(sourceWidgetID)
	destBytes := []byte(destWidgetID)

	// Validate lengths match
	if len(sourceBytes) != len(destBytes) {
		fmt.Printf("✗ Error: Pattern lengths must match!\n")
		fmt.Printf("  Source: %d bytes\n", len(sourceBytes))
		fmt.Printf("  Target: %d bytes\n", len(destBytes))
		return
	}

	for _, unit := range unitsToReplace {
		fmt.Printf("Processing %s: %s... ", unit.UnitType, unit.UnitName)

		// Load the BSON data as raw bytes
		bsonData, err := loadUnitBSON(mprPath, unit.UnitID)
		if err != nil {
			fmt.Printf("✗ Error loading: %v\n", err)
			errorCount++
			continue
		}

		// Perform binary byte replacement
		replacedCount := replaceBytesInData(bsonData, sourceBytes, destBytes)

		if replacedCount == 0 {
			fmt.Printf("⚠ No replacements made\n")
			errorCount++
			continue
		}

		// Save the modified bytes
		if err := saveUnitBSON(mprPath, unit.UnitID, bsonData); err != nil {
			fmt.Printf("✗ Error saving: %v\n", err)
			errorCount++
			continue
		}

		fmt.Printf("✓ Replaced %d occurrence(s)\n", replacedCount)
		successCount++
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Successfully processed: %d unit(s)\n", successCount)
	if errorCount > 0 {
		fmt.Printf("Failed: %d unit(s)\n", errorCount)
	}
	fmt.Println("\n✓ Replacement completed!")
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

// saveUnitBSON saves the modified BSON data back to the mxunit file
func saveUnitBSON(mprPath, unitID string, data []byte) error {
	dir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(dir, "mprcontents")

	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return fmt.Errorf("invalid unit ID: %s", unitID)
	}

	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write unit file %s: %w", filePath, err)
	}

	return nil
}

// dumpUnitJSON saves the unit data as a formatted JSON file
func dumpUnitJSON(data interface{}, unitType, unitName string) error {
	filename := fmt.Sprintf("%s_%s.json", strings.ToUpper(unitType), sanitizeFilename(unitName))

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// sanitizeFilename removes invalid characters from filenames
func sanitizeFilename(name string) string {
	// Replace invalid characters with underscore
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := name
	for _, char := range invalidChars {
		result = strings.ReplaceAll(result, char, "_")
	}
	return result
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
			// Check if it has the matching widgetId in the Type field (case insensitive substring match)
			if typeField, ok := v["Type"].(map[string]interface{}); ok {
				if wid, ok := typeField["WidgetId"].(string); ok {
					// Case insensitive match: exact or substring
					widLower := strings.ToLower(wid)
					searchLower := strings.ToLower(widgetID)
					if widLower == searchLower || strings.Contains(widLower, searchLower) {
						*matchingWidgets = append(*matchingWidgets, v)
					}
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

// replaceBytesInData performs binary byte-level replacement in a byte array
// This preserves the exact BSON structure without unmarshaling/marshaling
func replaceBytesInData(data []byte, oldPattern, newPattern []byte) int {
	if len(oldPattern) != len(newPattern) {
		return 0
	}

	replacements := 0
	dataLen := len(data)
	patternLen := len(oldPattern)

	for i := 0; i <= dataLen-patternLen; i++ {
		// Check if pattern matches at position i
		match := true
		for j := 0; j < patternLen; j++ {
			if data[i+j] != oldPattern[j] {
				match = false
				break
			}
		}

		if match {
			// Replace bytes in place
			for j := 0; j < patternLen; j++ {
				data[i+j] = newPattern[j]
			}
			replacements++
			i += patternLen - 1 // Skip past the replaced pattern
		}
	}

	return replacements
}

// replaceWidgetID replaces the widgetId in all matching custom widgets
func replaceWidgetID(data interface{}, sourceWidgetID, destWidgetID string) int {
	count := 0
	replaceInData(data, sourceWidgetID, destWidgetID, &count)
	return count
}

// replaceWidgetIDInBSON replaces widgetId directly in BSON structure without JSON conversion
func replaceWidgetIDInBSON(data interface{}, sourceWidgetID, destWidgetID string) int {
	count := 0
	replaceInBSONData(data, sourceWidgetID, destWidgetID, &count)
	return count
}

func replaceInBSONData(data interface{}, sourceWidgetID, destWidgetID string, count *int) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a CustomWidget with matching widgetId
		if typeVal, ok := v["$Type"].(string); ok && typeVal == "CustomWidgets$CustomWidget" {
			if typeField, ok := v["Type"].(map[string]interface{}); ok {
				if wid, ok := typeField["WidgetId"].(string); ok {
					widLower := strings.ToLower(wid)
					searchLower := strings.ToLower(sourceWidgetID)
					if widLower == searchLower || strings.Contains(widLower, searchLower) {
						// Replace the widgetId directly
						typeField["WidgetId"] = destWidgetID
						*count++
					}
				}
			}
		}

		// Recursively process all fields
		for _, value := range v {
			replaceInBSONData(value, sourceWidgetID, destWidgetID, count)
		}

	case []interface{}:
		// Recursively process arrays
		for _, item := range v {
			replaceInBSONData(item, sourceWidgetID, destWidgetID, count)
		}
	}
}

func replaceInData(data interface{}, sourceWidgetID, destWidgetID string, count *int) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a CustomWidget with matching widgetId
		if typeVal, ok := v["$Type"].(string); ok && typeVal == "CustomWidgets$CustomWidget" {
			if typeField, ok := v["Type"].(map[string]interface{}); ok {
				if wid, ok := typeField["WidgetId"].(string); ok {
					widLower := strings.ToLower(wid)
					searchLower := strings.ToLower(sourceWidgetID)
					if widLower == searchLower || strings.Contains(widLower, searchLower) {
						// Replace the widgetId
						typeField["WidgetId"] = destWidgetID
						*count++
					}
				}
			}
		}

		// Recursively process all fields
		for _, value := range v {
			replaceInData(value, sourceWidgetID, destWidgetID, count)
		}

	case []interface{}:
		// Recursively process arrays
		for _, item := range v {
			replaceInData(item, sourceWidgetID, destWidgetID, count)
		}
	}
}

// extractWidgetID extracts the full widgetId from a widget
func extractWidgetID(widget map[string]interface{}) string {
	if typeField, ok := widget["Type"].(map[string]interface{}); ok {
		if widgetID, ok := typeField["WidgetId"].(string); ok {
			return widgetID
		}
	}
	return "[unknown]"
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
