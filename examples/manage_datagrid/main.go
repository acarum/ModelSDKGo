package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/anthropics/modelsdk-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const defaultDataGridWidgetID = "com.mendix.widget.web.datagrid.Datagrid"
const targetCustomDateTimeFormat = ",MMM dd, yyyy  hh:mm:ss a"
const dateTimeModificationDescription = "set DateTime format from custom to default and remove custom format value when it matches the target pattern"

var formatDateTimeExpressionRegex = regexp.MustCompile(`(?i)^\s*formatdatetime\s*\(\s*(.+?)\s*,\s*'([^']*)'\s*\)\s*$`)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  Find mode:            manage_datagrid <mpr_file_path> [widget_id] [--dump-json]")
		fmt.Println("  DefaultDateTme mode:  manage_datagrid <mpr_file_path> [widget_id] --defaultDateTme [--dump-target-json] [--only-page ModuleName.PageName] [--force-page-diff]")
		fmt.Println("\nExamples:")
		fmt.Println("  manage_datagrid MyApp.mpr")
		fmt.Println("  manage_datagrid MyApp.mpr com.mendix.widget.web.datagrid.Datagrid --dump-json")
		fmt.Println("  manage_datagrid MyApp.mpr --defaultDateTme")
		fmt.Println("  manage_datagrid MyApp.mpr com.mendix.widget.web.datagrid.Datagrid --defaultDateTme")
		fmt.Println("  manage_datagrid MyApp.mpr --defaultDateTme --dump-target-json")
		fmt.Println("  manage_datagrid MyApp.mpr --defaultDateTme --only-page MyModule.MyPage")
		fmt.Println("  manage_datagrid MyApp.mpr --defaultDateTme --dump-target-json --only-page MyModule.MyPage")
		fmt.Println("  manage_datagrid MyApp.mpr --defaultDateTme --only-page MyModule.MyPage --force-page-diff")
		fmt.Println("\nNote:")
		fmt.Printf("  Default widget_id: %s\n", defaultDataGridWidgetID)
		os.Exit(1)
	}

	mprPath := os.Args[1]
	sourceWidgetID := defaultDataGridWidgetID

	// Check modes and flags
	defaultDateTmeMode := false
	dumpJSON := false
	dumpTargetJSON := false
	forcePageDiff := false
	onlyPage := ""
	widgetIDSet := false

	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--dump-json" {
			dumpJSON = true
			continue
		}
		if arg == "--defaultDateTme" {
			defaultDateTmeMode = true
			continue
		}
		if arg == "--dump-target-json" {
			dumpTargetJSON = true
			continue
		}
		if arg == "--force-page-diff" {
			forcePageDiff = true
			continue
		}
		if arg == "--only-page" && i+1 < len(os.Args) {
			onlyPage = os.Args[i+1]
			i++ // skip next arg
			continue
		}

		if strings.HasPrefix(arg, "--") {
			fmt.Printf("Unknown flag: %s\n", arg)
			os.Exit(1)
		}

		if !widgetIDSet {
			sourceWidgetID = arg
			widgetIDSet = true
			continue
		}

		fmt.Printf("Unexpected extra argument: %s\n", arg)
		os.Exit(1)
	}

	if !widgetIDSet {
		fmt.Printf("No widget_id provided, using default: %s\n", sourceWidgetID)
	}

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())

	if defaultDateTmeMode {
		fmt.Printf("\n=== DEFAULT DATETIME MODE ===\n")
		fmt.Printf("Searching for widgets with widgetId: %s\n", sourceWidgetID)
		fmt.Printf("Will reset custom DateTime formatting to default only when custom format is: %s\n\n", targetCustomDateTimeFormat)
		if dumpTargetJSON {
			fmt.Println("Target page JSON dump mode: ENABLED")
		}
		if onlyPage != "" {
			fmt.Printf("Filter: Only page %s will be modified\n", onlyPage)
		}
		if forcePageDiff {
			fmt.Println("Force page diff mode: ENABLED")
		}
		fmt.Println()
		performDefaultDateTme(reader, mprPath, sourceWidgetID, dumpTargetJSON, onlyPage, forcePageDiff)
	} else {
		fmt.Printf("\n=== Searching for custom widgets with widgetId: %s ===\n", sourceWidgetID)
		if dumpJSON {
			fmt.Printf("JSON dump mode: ENABLED\n")
		}
		fmt.Println()
		performSearch(reader, mprPath, sourceWidgetID, dumpJSON)
	}
}

type UnitDefaultDateTmeInfo struct {
	UnitType         string
	UnitID           string
	UnitName         string
	QualifiedName    string
	GridCount        int
	ColumnsUpdated   int
	ColumnNames      []string
	DateTimeColumns  []string
	ModificationType string
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

func performDefaultDateTme(reader *modelsdk.Reader, mprPath, widgetID string, dumpTargetJSON bool, onlyPage string, forcePageDiff bool) {
	var unitsToUpdate []UnitDefaultDateTmeInfo
	dumpedTargetPages := 0

	// Build module map (moduleID -> moduleName)
	moduleMap := make(map[string]string)
	modules, err := reader.ListModules()
	if err == nil {
		for _, module := range modules {
			moduleMap[string(module.ID)] = module.Name
		}
	}

	// Build container hierarchy map (unitID -> containerID)
	containerMap := make(map[string]string)
	units, err := reader.ListUnits()
	if err == nil {
		for _, unit := range units {
			containerMap[string(unit.ID)] = string(unit.ContainerID)
		}
	}

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

			widgets := findWidgetsByWidgetIDDirect(snippetData, widgetID)
			if len(widgets) == 0 {
				continue
			}

			columnsUpdated := 0
			columnNames := make([]string, 0)
			dateTimeColumns := make([]string, 0)
			for _, widget := range widgets {
				updated, updatedColumns, dtColumns := collectDateTimeModificationsInWidget(widget)
				columnsUpdated += updated
				columnNames = appendUniqueStrings(columnNames, updatedColumns...)
				dateTimeColumns = appendUniqueStrings(dateTimeColumns, dtColumns...)
			}

			if columnsUpdated > 0 || len(dateTimeColumns) > 0 {
				unitsToUpdate = append(unitsToUpdate, UnitDefaultDateTmeInfo{
					UnitType:         "Snippet",
					UnitID:           string(snippet.ID),
					UnitName:         snippet.Name,
					GridCount:        len(widgets),
					ColumnsUpdated:   columnsUpdated,
					ColumnNames:      columnNames,
					DateTimeColumns:  dateTimeColumns,
					ModificationType: dateTimeModificationDescription,
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
			moduleName := resolveModuleNameFromContainerID(string(page.ContainerID), moduleMap, containerMap)
			if moduleName == "" {
				moduleName = "[unknown]"
			}
			pageIdentifier := moduleName + "." + page.Name

			if onlyPage != "" && !matchesOnlyPageFilter(pageIdentifier, page.Name, onlyPage) {
				continue
			}

			bsonData, err := loadUnitBSON(mprPath, string(page.ID))
			if err != nil {
				continue
			}

			var pageData map[string]interface{}
			if err := bson.Unmarshal(bsonData, &pageData); err != nil {
				continue
			}

			widgets := findWidgetsByWidgetIDDirect(pageData, widgetID)
			if len(widgets) == 0 {
				continue
			}

			columnsUpdated := 0
			columnNames := make([]string, 0)
			dateTimeColumns := make([]string, 0)
			for _, widget := range widgets {
				updated, updatedColumns, dtColumns := collectDateTimeModificationsInWidget(widget)
				columnsUpdated += updated
				columnNames = appendUniqueStrings(columnNames, updatedColumns...)
				dateTimeColumns = appendUniqueStrings(dateTimeColumns, dtColumns...)
			}

			if columnsUpdated > 0 || len(dateTimeColumns) > 0 {
				unitsToUpdate = append(unitsToUpdate, UnitDefaultDateTmeInfo{
					UnitType:         "Page",
					UnitID:           string(page.ID),
					UnitName:         page.Name,
					QualifiedName:    pageIdentifier,
					GridCount:        len(widgets),
					ColumnsUpdated:   columnsUpdated,
					ColumnNames:      columnNames,
					DateTimeColumns:  dateTimeColumns,
					ModificationType: dateTimeModificationDescription,
				})

				if dumpTargetJSON && columnsUpdated > 0 {
					if err := dumpUnitJSON(pageData, "TARGET_PAGE", page.Name); err != nil {
						fmt.Printf("  ✗ Error dumping target page JSON for %s: %v\n", page.Name, err)
					} else {
						dumpedTargetPages++
						fmt.Printf("  ✓ Target page JSON saved: TARGET_PAGE_%s.json\n", sanitizeFilename(page.Name))
					}
				}
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

			widgets := findWidgetsByWidgetIDDirect(layoutData, widgetID)
			if len(widgets) == 0 {
				continue
			}

			columnsUpdated := 0
			columnNames := make([]string, 0)
			dateTimeColumns := make([]string, 0)
			for _, widget := range widgets {
				updated, updatedColumns, dtColumns := collectDateTimeModificationsInWidget(widget)
				columnsUpdated += updated
				columnNames = appendUniqueStrings(columnNames, updatedColumns...)
				dateTimeColumns = appendUniqueStrings(dateTimeColumns, dtColumns...)
			}

			if columnsUpdated > 0 || len(dateTimeColumns) > 0 {
				unitsToUpdate = append(unitsToUpdate, UnitDefaultDateTmeInfo{
					UnitType:         "Layout",
					UnitID:           string(layout.ID),
					UnitName:         layout.Name,
					GridCount:        len(widgets),
					ColumnsUpdated:   columnsUpdated,
					ColumnNames:      columnNames,
					DateTimeColumns:  dateTimeColumns,
					ModificationType: dateTimeModificationDescription,
				})
			}
		}
	}

	if len(unitsToUpdate) == 0 {
		fmt.Printf("\nNo custom DateTime format found under Content/ContentParams for widgetId '%s'.\n", widgetID)
		if onlyPage != "" {
			fmt.Printf("\n=== Pending Updates (Before Modification) ===\n")
			fmt.Printf("  Page: %s\n", onlyPage)
			fmt.Printf("  Columns to modify: none\n")
			fmt.Printf("  DateTime columns: none\n")
			fmt.Printf("  Modification: %s\n", dateTimeModificationDescription)
			fmt.Printf("Note: Filter was set to only-page %s\n", onlyPage)
			if forcePageDiff {
				if err := forceDumpPageBeforeAfterDiff(reader, mprPath, onlyPage); err != nil {
					fmt.Printf("✗ Force page diff failed: %v\n", err)
				} else {
					fmt.Println("✓ Force page diff generated (BEFORE/AFTER/DIFF)")
				}
			}
		}
		return
	}

	if dumpTargetJSON {
		fmt.Printf("\nTarget page JSON files saved: %d\n", dumpedTargetPages)
	}

	totalWidgets := 0
	totalColumns := 0
	fmt.Printf("\n=== Pending Updates (Before Modification) ===\n")
	for _, unit := range unitsToUpdate {
		displayName := unit.UnitName
		if unit.UnitType == "Page" && unit.QualifiedName != "" {
			displayName = unit.QualifiedName
		}
		fmt.Printf("  %s: %s (%d grid(s), %d column update(s))\n", unit.UnitType, displayName, unit.GridCount, unit.ColumnsUpdated)
		if len(unit.ColumnNames) > 0 {
			fmt.Printf("    Columns to modify: %s\n", strings.Join(unit.ColumnNames, ", "))
		} else {
			fmt.Printf("    Columns to modify: none\n")
		}
		if len(unit.DateTimeColumns) > 0 {
			fmt.Printf("    DateTime columns: %s\n", strings.Join(unit.DateTimeColumns, ", "))
		} else {
			fmt.Printf("    DateTime columns: none\n")
		}
		if unit.ModificationType != "" {
			fmt.Printf("    Modification: %s\n", unit.ModificationType)
		}
		totalWidgets += unit.GridCount
		totalColumns += unit.ColumnsUpdated
	}
	fmt.Printf("\nTotal grids: %d\n", totalWidgets)
	fmt.Printf("Total column updates: %d\n", totalColumns)

	fmt.Printf("\nApply updates and save to mxunit files? (yes/no): ")
	inputReader := bufio.NewReader(os.Stdin)
	response, _ := inputReader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response != "yes" && response != "y" {
		fmt.Println("Operation cancelled.")
		return
	}

	fmt.Println("\n=== Saving Updates ===")
	successCount := 0
	errorCount := 0
	pageDumpPairs := 0
	diffReports := 0
	modifiedPageSummaries := make([]string, 0)

	for _, unit := range unitsToUpdate {
		fmt.Printf("Processing %s: %s... ", unit.UnitType, unit.UnitName)

		bsonData, err := loadUnitBSON(mprPath, unit.UnitID)
		if err != nil {
			fmt.Printf("✗ Error loading: %v\n", err)
			errorCount++
			continue
		}

		var unitData map[string]interface{}
		if err := bson.Unmarshal(bsonData, &unitData); err != nil {
			fmt.Printf("✗ Error unmarshaling: %v\n", err)
			errorCount++
			continue
		}

		shouldDumpPageArtifacts := dumpTargetJSON || forcePageDiff
		shouldGeneratePageDiff := forcePageDiff

		var beforeData interface{}
		if unit.UnitType == "Page" && (shouldDumpPageArtifacts || shouldGeneratePageDiff) {
			beforeData, err = deepCopyJSONValue(unitData)
			if err != nil {
				fmt.Printf("✗ Error preparing before snapshot: %v\n", err)
				errorCount++
				continue
			}
		}

		widgets := findWidgetsByWidgetIDDirect(unitData, widgetID)
		columnsUpdated := 0
		for _, widget := range widgets {
			columnsUpdated += resetDateTimeCustomFormattingInWidget(widget)
		}

		if columnsUpdated == 0 {
			fmt.Printf("⚠ No changes to save\n")
			continue
		}

		if unit.UnitType == "Page" && shouldDumpPageArtifacts {
			beforeFile := fmt.Sprintf("TARGET_PAGE_BEFORE_%s.json", sanitizeFilename(unit.UnitName))
			if err := dumpJSONToFile(beforeData, beforeFile); err != nil {
				fmt.Printf("✗ Error dumping BEFORE JSON: %v\n", err)
				errorCount++
				continue
			}
		}

		updatedBSON, err := bson.Marshal(unitData)
		if err != nil {
			fmt.Printf("✗ Error marshaling: %v\n", err)
			errorCount++
			continue
		}

		if err := saveUnitBSON(mprPath, unit.UnitID, updatedBSON); err != nil {
			fmt.Printf("✗ Error saving: %v\n", err)
			errorCount++
			continue
		}

		if unit.UnitType == "Page" {
			diffs := make([]JSONDiff, 0)

			if shouldDumpPageArtifacts {
				afterFile := fmt.Sprintf("TARGET_PAGE_AFTER_%s.json", sanitizeFilename(unit.UnitName))
				if err := dumpJSONToFile(unitData, afterFile); err != nil {
					fmt.Printf("✗ Error dumping AFTER JSON: %v\n", err)
					errorCount++
					continue
				}
				pageDumpPairs++
			}

			if shouldGeneratePageDiff {
				diffs = collectJSONDiffs(beforeData, unitData)
				diffFile := fmt.Sprintf("TARGET_PAGE_DIFF_%s.txt", sanitizeFilename(unit.UnitName))
				if err := writeJSONDiffReport(diffFile, diffs); err != nil {
					fmt.Printf("✗ Error writing DIFF report: %v\n", err)
					errorCount++
					continue
				}
				diffReports++
			}

			pageDisplayName := unit.QualifiedName
			if pageDisplayName == "" {
				pageDisplayName = unit.UnitName
			}
			modifiedColumns := "none"
			if len(unit.ColumnNames) > 0 {
				modifiedColumns = strings.Join(unit.ColumnNames, ", ")
			}
			modifiedPageSummaries = append(modifiedPageSummaries, fmt.Sprintf("%s: %s", pageDisplayName, modifiedColumns))
			if shouldDumpPageArtifacts && shouldGeneratePageDiff {
				fmt.Printf("✓ Updated %d column(s) [BEFORE/AFTER dumped, DIFF: %d change(s)]\n", columnsUpdated, len(diffs))
			} else if shouldDumpPageArtifacts {
				fmt.Printf("✓ Updated %d column(s) [BEFORE/AFTER dumped]\n", columnsUpdated)
			} else if shouldGeneratePageDiff {
				fmt.Printf("✓ Updated %d column(s) [DIFF: %d change(s)]\n", columnsUpdated, len(diffs))
			} else {
				fmt.Printf("✓ Updated %d column(s)\n", columnsUpdated)
			}
			successCount++
			continue
		}

		fmt.Printf("✓ Updated %d column(s)\n", columnsUpdated)
		successCount++
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Successfully processed: %d unit(s)\n", successCount)
	if errorCount > 0 {
		fmt.Printf("Failed: %d unit(s)\n", errorCount)
	}
	if pageDumpPairs > 0 {
		fmt.Printf("Page dump BEFORE/AFTER pairs: %d\n", pageDumpPairs)
	}
	if diffReports > 0 {
		fmt.Printf("Diff reports generated: %d\n", diffReports)
	}
	if len(modifiedPageSummaries) > 0 {
		fmt.Println("\nModified pages (ModuleName.PageName: col1, col2, ...):")
		for _, summaryLine := range modifiedPageSummaries {
			fmt.Printf("  %s\n", summaryLine)
		}
	}
	fmt.Println("\n✓ Default DateTime update completed!")
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
	return dumpJSONToFile(data, filename)
}

func dumpJSONToFile(data interface{}, filename string) error {

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

type JSONDiff struct {
	Path   string
	Before string
	After  string
}

func deepCopyJSONValue(v interface{}) (interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var copy interface{}
	if err := json.Unmarshal(data, &copy); err != nil {
		return nil, err
	}

	return copy, nil
}

func collectJSONDiffs(before, after interface{}) []JSONDiff {
	diffs := make([]JSONDiff, 0)
	collectJSONDiffsRecursive("$", before, after, &diffs)
	return diffs
}

func collectJSONDiffsRecursive(path string, before, after interface{}, diffs *[]JSONDiff) {
	beforeMap, beforeIsMap := before.(map[string]interface{})
	afterMap, afterIsMap := after.(map[string]interface{})
	if beforeIsMap && afterIsMap {
		keysMap := make(map[string]struct{})
		for k := range beforeMap {
			keysMap[k] = struct{}{}
		}
		for k := range afterMap {
			keysMap[k] = struct{}{}
		}

		keys := make([]string, 0, len(keysMap))
		for k := range keysMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			b, bOK := beforeMap[key]
			a, aOK := afterMap[key]
			nextPath := path + "." + key
			if !bOK {
				*diffs = append(*diffs, JSONDiff{Path: nextPath, Before: "<missing>", After: toJSONValueString(a)})
				continue
			}
			if !aOK {
				*diffs = append(*diffs, JSONDiff{Path: nextPath, Before: toJSONValueString(b), After: "<missing>"})
				continue
			}
			collectJSONDiffsRecursive(nextPath, b, a, diffs)
		}
		return
	}

	beforeArr, beforeIsArr := before.([]interface{})
	afterArr, afterIsArr := after.([]interface{})
	if beforeIsArr && afterIsArr {
		maxLen := len(beforeArr)
		if len(afterArr) > maxLen {
			maxLen = len(afterArr)
		}

		for i := 0; i < maxLen; i++ {
			nextPath := fmt.Sprintf("%s[%d]", path, i)
			if i >= len(beforeArr) {
				*diffs = append(*diffs, JSONDiff{Path: nextPath, Before: "<missing>", After: toJSONValueString(afterArr[i])})
				continue
			}
			if i >= len(afterArr) {
				*diffs = append(*diffs, JSONDiff{Path: nextPath, Before: toJSONValueString(beforeArr[i]), After: "<missing>"})
				continue
			}
			collectJSONDiffsRecursive(nextPath, beforeArr[i], afterArr[i], diffs)
		}
		return
	}

	if toJSONValueString(before) != toJSONValueString(after) {
		*diffs = append(*diffs, JSONDiff{Path: path, Before: toJSONValueString(before), After: toJSONValueString(after)})
	}
}

func toJSONValueString(v interface{}) string {
	if v == nil {
		return "null"
	}
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}

func writeJSONDiffReport(filename string, diffs []JSONDiff) error {
	var b strings.Builder
	b.WriteString("JSON DIFF REPORT\n")
	b.WriteString(fmt.Sprintf("Total differences: %d\n\n", len(diffs)))

	for i, d := range diffs {
		b.WriteString(fmt.Sprintf("%d) %s\n", i+1, d.Path))
		b.WriteString(fmt.Sprintf("   - before: %s\n", d.Before))
		b.WriteString(fmt.Sprintf("   - after : %s\n\n", d.After))
	}

	return os.WriteFile(filename, []byte(b.String()), 0644)
}

func forceDumpPageBeforeAfterDiff(reader *modelsdk.Reader, mprPath, qualifiedPage string) error {
	modules, err := reader.ListModules()
	if err != nil {
		return fmt.Errorf("list modules: %w", err)
	}

	moduleMap := make(map[string]string)
	for _, module := range modules {
		moduleMap[string(module.ID)] = module.Name
	}

	containerMap := make(map[string]string)
	units, err := reader.ListUnits()
	if err != nil {
		return fmt.Errorf("list units: %w", err)
	}
	for _, unit := range units {
		containerMap[string(unit.ID)] = string(unit.ContainerID)
	}

	pages, err := reader.ListPages()
	if err != nil {
		return fmt.Errorf("list pages: %w", err)
	}

	var targetPageID string
	var targetPageName string
	var targetQualified string
	for _, page := range pages {
		moduleName := resolveModuleNameFromContainerID(string(page.ContainerID), moduleMap, containerMap)
		if moduleName == "" {
			moduleName = "[unknown]"
		}
		candidateQualified := moduleName + "." + page.Name
		if candidateQualified != qualifiedPage {
			continue
		}
		targetPageID = string(page.ID)
		targetPageName = page.Name
		targetQualified = candidateQualified
		break
	}

	if targetPageID == "" {
		requestedPageName := qualifiedPage
		if idx := strings.LastIndex(qualifiedPage, "."); idx >= 0 && idx+1 < len(qualifiedPage) {
			requestedPageName = qualifiedPage[idx+1:]
		}

		matches := 0
		for _, page := range pages {
			if page.Name != requestedPageName {
				continue
			}
			moduleName := resolveModuleNameFromContainerID(string(page.ContainerID), moduleMap, containerMap)
			if moduleName == "" {
				moduleName = "[unknown]"
			}
			targetPageID = string(page.ID)
			targetPageName = page.Name
			targetQualified = moduleName + "." + page.Name
			matches++
		}

		if matches > 1 {
			return fmt.Errorf("multiple pages matched by name '%s', use exact ModuleName.PageName", requestedPageName)
		}
	}

	if targetPageID == "" {
		return fmt.Errorf("page not found: %s", qualifiedPage)
	}

	bsonData, err := loadUnitBSON(mprPath, targetPageID)
	if err != nil {
		return fmt.Errorf("load page bson: %w", err)
	}

	var pageData map[string]interface{}
	if err := bson.Unmarshal(bsonData, &pageData); err != nil {
		return fmt.Errorf("unmarshal page bson: %w", err)
	}

	beforeData, err := deepCopyJSONValue(pageData)
	if err != nil {
		return fmt.Errorf("copy before data: %w", err)
	}

	beforeFile := fmt.Sprintf("TARGET_PAGE_BEFORE_%s.json", sanitizeFilename(targetPageName))
	afterFile := fmt.Sprintf("TARGET_PAGE_AFTER_%s.json", sanitizeFilename(targetPageName))
	diffFile := fmt.Sprintf("TARGET_PAGE_DIFF_%s.txt", sanitizeFilename(targetPageName))

	if err := dumpJSONToFile(beforeData, beforeFile); err != nil {
		return fmt.Errorf("write before file: %w", err)
	}
	if err := dumpJSONToFile(pageData, afterFile); err != nil {
		return fmt.Errorf("write after file: %w", err)
	}

	diffs := collectJSONDiffs(beforeData, pageData)
	if err := writeJSONDiffReport(diffFile, diffs); err != nil {
		return fmt.Errorf("write diff report: %w", err)
	}

	fmt.Printf("  ✓ Force diff page resolved: %s\n", targetQualified)

	return nil
}

func matchesOnlyPageFilter(pageIdentifier, pageName, onlyPage string) bool {
	if onlyPage == "" {
		return true
	}

	if pageIdentifier == onlyPage {
		return true
	}

	requestedPageName := onlyPage
	if idx := strings.LastIndex(onlyPage, "."); idx >= 0 && idx+1 < len(onlyPage) {
		requestedPageName = onlyPage[idx+1:]
	}

	return pageName == requestedPageName
}

func resolveModuleNameFromContainerID(containerID string, moduleMap map[string]string, containerMap map[string]string) string {
	if name, ok := moduleMap[containerID]; ok {
		return name
	}

	visited := make(map[string]struct{})
	currentID := containerID
	for currentID != "" {
		if _, seen := visited[currentID]; seen {
			break
		}
		visited[currentID] = struct{}{}

		if name, ok := moduleMap[currentID]; ok {
			return name
		}

		nextID, ok := containerMap[currentID]
		if !ok || nextID == currentID {
			break
		}

		currentID = nextID
	}

	return ""
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

func findWidgetsByWidgetIDDirect(data interface{}, widgetID string) []map[string]interface{} {
	var matchingWidgets []map[string]interface{}
	searchForMatchingWidgetsDirect(data, widgetID, &matchingWidgets)
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

func searchForMatchingWidgetsDirect(data interface{}, widgetID string, matchingWidgets *[]map[string]interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		if typeVal, ok := v["$Type"].(string); ok && typeVal == "CustomWidgets$CustomWidget" {
			if typeField, ok := v["Type"].(map[string]interface{}); ok {
				if wid, ok := typeField["WidgetId"].(string); ok {
					widLower := strings.ToLower(wid)
					searchLower := strings.ToLower(widgetID)
					if widLower == searchLower || strings.Contains(widLower, searchLower) {
						*matchingWidgets = append(*matchingWidgets, v)
					}
				}
			}
		}

		for _, value := range v {
			searchForMatchingWidgetsDirect(value, widgetID, matchingWidgets)
		}

	case []interface{}:
		for _, item := range v {
			searchForMatchingWidgetsDirect(item, widgetID, matchingWidgets)
		}

	case primitive.A:
		for _, item := range v {
			searchForMatchingWidgetsDirect(item, widgetID, matchingWidgets)
		}
	}
}

func resetDateTimeCustomFormattingInWidget(widget map[string]interface{}) int {
	updatedColumns, _, _ := collectDateTimeModificationsInWidget(widget)
	return updatedColumns
}

func collectDateTimeModificationsInWidget(widget map[string]interface{}) (int, []string, []string) {
	typeMap := make(map[string]string)
	buildTypePointerMap(widget, typeMap)

	obj, ok := widget["Object"].(map[string]interface{})
	if !ok {
		return 0, nil, nil
	}

	props := getArray(obj, "Properties")
	for _, p := range props {
		pm, ok := p.(map[string]interface{})
		if !ok {
			continue
		}
		if typeMap[getTypePointerData(pm)] != "columns" {
			continue
		}

		val, ok := pm["Value"].(map[string]interface{})
		if !ok {
			continue
		}

		updatedColumns := 0
		columnNames := make([]string, 0)
		dateTimeColumns := make([]string, 0)
		colObjects := toSlice(val["Objects"])
		for colIdx, colObj := range colObjects {
			colMap, ok := colObj.(map[string]interface{})
			if !ok || colMap["$Type"] != "CustomWidgets$WidgetObject" {
				continue
			}

			colProps := getArray(colMap, "Properties")
			isLastUpdatedOn := columnTargetsLastUpdatedOn(colProps)

			// Check if this column has legacy dynamicText marker property
			hasLegacyMarker := hasLegacyDynamicTextProperty(colProps)

			columnUpdated := false
			columnUsesDateTime := false

			// Align LastUpdatedOn to CreatedOn behavior: dynamicText -> attribute
			if isLastUpdatedOn && hasLegacyMarker && normalizeDynamicTextToAttribute(colProps) {
				columnUpdated = true
			}

			for _, cp := range colProps {
				cpMap, ok := cp.(map[string]interface{})
				if !ok {
					continue
				}

				if valMap, ok := cpMap["Value"].(map[string]interface{}); ok {
					if hasDateTimeFormatInContentParams(valMap) {
						// Only track DateTime usage on LastUpdatedOn column
						if isLastUpdatedOn {
							columnUsesDateTime = true
						}
					}
					// Apply DateTime expression reset only to LastUpdatedOn
					if isLastUpdatedOn && resetDateTimeInContentParams(valMap) > 0 {
						columnUpdated = true
					}
				}
			}

			columnName := resolveDataGridColumnName(colProps, typeMap, colIdx)
			if columnUsesDateTime {
				dateTimeColumns = appendUniqueStrings(dateTimeColumns, columnName)
			}

			if columnUpdated {
				updatedColumns++
				columnNames = appendUniqueStrings(columnNames, columnName)
			}
		}

		return updatedColumns, columnNames, dateTimeColumns
	}

	return 0, nil, nil
}

func hasLegacyDynamicTextProperty(colProps []interface{}) bool {
	for _, cp := range colProps {
		cpMap, ok := cp.(map[string]interface{})
		if !ok {
			continue
		}

		if valMap, ok := cpMap["Value"].(map[string]interface{}); ok {
			if primVal, ok := valMap["PrimitiveValue"].(string); ok && primVal == "dynamicText" {
				return true
			}
		}
	}
	return false
}

func normalizeDynamicTextToAttribute(colProps []interface{}) bool {
	updated := false
	for _, cp := range colProps {
		cpMap, ok := cp.(map[string]interface{})
		if !ok {
			continue
		}

		valMap, ok := cpMap["Value"].(map[string]interface{})
		if !ok {
			continue
		}

		if primVal, ok := valMap["PrimitiveValue"].(string); ok && strings.EqualFold(strings.TrimSpace(primVal), "dynamicText") {
			valMap["PrimitiveValue"] = "attribute"
			updated = true
		}
	}

	return updated
}

func columnTargetsLastUpdatedOn(colProps []interface{}) bool {
	for _, cp := range colProps {
		cpMap, ok := cp.(map[string]interface{})
		if !ok {
			continue
		}

		valMap, ok := cpMap["Value"].(map[string]interface{})
		if !ok {
			continue
		}

		if containsLastUpdatedOnReference(valMap) {
			return true
		}
	}

	return false
}

func containsLastUpdatedOnReference(data interface{}) bool {
	switch v := data.(type) {
	case map[string]interface{}:
		if attrRef, ok := v["AttributeRef"].(map[string]interface{}); ok {
			if attr, ok := attrRef["Attribute"].(string); ok {
				attrLower := strings.ToLower(attr)
				if attrLower == "lastupdatedon" || strings.HasSuffix(attrLower, ".lastupdatedon") {
					return true
				}
			}
		}

		if expr, ok := v["Expression"].(string); ok {
			exprLower := strings.ToLower(expr)
			if strings.Contains(exprLower, "lastupdatedon") {
				return true
			}
		}

		for _, value := range v {
			if containsLastUpdatedOnReference(value) {
				return true
			}
		}

	case []interface{}:
		for _, item := range v {
			if containsLastUpdatedOnReference(item) {
				return true
			}
		}

	case primitive.A:
		for _, item := range v {
			if containsLastUpdatedOnReference(item) {
				return true
			}
		}
	}

	return false
}

func hasDateTimeFormatInContentParams(data interface{}) bool {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			normalizedKey := normalizeKey(key)
			if isDateTimeFormatKey(normalizedKey) || isCustomDateTimeFormatKey(normalizedKey) {
				return true
			}
			if normalizedKey == "expression" && isDateTimeExpressionValue(value) {
				return true
			}
			if hasDateTimeFormatInContentParams(value) {
				return true
			}
		}
	case []interface{}:
		for _, item := range v {
			if hasDateTimeFormatInContentParams(item) {
				return true
			}
		}
	case primitive.A:
		for _, item := range v {
			if hasDateTimeFormatInContentParams(item) {
				return true
			}
		}
	}

	return false
}

func resolveDataGridColumnName(colProps []interface{}, typeMap map[string]string, colIdx int) string {
	fallback := fmt.Sprintf("col%d", colIdx+1)
	for _, cp := range colProps {
		cpMap, ok := cp.(map[string]interface{})
		if !ok {
			continue
		}

		propKey := strings.ToLower(typeMap[getTypePointerData(cpMap)])
		if propKey != "id" && propKey != "attribute" {
			continue
		}

		valMap, ok := cpMap["Value"].(map[string]interface{})
		if !ok {
			continue
		}

		if columnName, ok := extractComparableString(valMap); ok && columnName != "" {
			return columnName
		}
	}

	return fallback
}

func appendUniqueStrings(base []string, values ...string) []string {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			continue
		}
		exists := false
		for _, existing := range base {
			if existing == value {
				exists = true
				break
			}
		}
		if !exists {
			base = append(base, value)
		}
	}
	return base
}

func buildTypePointerMap(data interface{}, m map[string]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		if v["$Type"] == "CustomWidgets$WidgetPropertyType" {
			key, _ := v["PropertyKey"].(string)
			if d := extractBinaryKey(v["$ID"]); d != "" && key != "" {
				m[d] = key
			}
		}
		for _, val := range v {
			buildTypePointerMap(val, m)
		}
	case []interface{}:
		for _, item := range v {
			buildTypePointerMap(item, m)
		}
	case primitive.A:
		for _, item := range v {
			buildTypePointerMap(item, m)
		}
	}
}

func extractBinaryKey(v interface{}) string {
	switch id := v.(type) {
	case primitive.Binary:
		return string(id.Data)
	case map[string]interface{}:
		if data, ok := id["Data"].(string); ok {
			return data
		}
	case string:
		return id
	}
	return ""
}

func getTypePointerData(prop map[string]interface{}) string {
	return extractBinaryKey(prop["TypePointer"])
}

func toSlice(v interface{}) []interface{} {
	var items []interface{}
	switch arr := v.(type) {
	case primitive.A:
		items = []interface{}(arr)
	case []interface{}:
		items = arr
	default:
		return nil
	}

	var result []interface{}
	for _, item := range items {
		switch item.(type) {
		case map[string]interface{}, primitive.A, []interface{}:
			result = append(result, item)
		}
	}
	return result
}

func getArray(data map[string]interface{}, key string) []interface{} {
	v, ok := data[key]
	if !ok {
		return nil
	}
	return toSlice(v)
}

func resetDateTimeInContentParams(data interface{}) int {
	switch v := data.(type) {
	case map[string]interface{}:
		changes := resetDateTimeInContentParamsMap(v)
		for _, value := range v {
			changes += resetDateTimeInContentParams(value)
		}
		return changes

	case []interface{}:
		changes := 0
		for _, item := range v {
			changes += resetDateTimeInContentParams(item)
		}
		return changes

	case primitive.A:
		changes := 0
		for _, item := range v {
			changes += resetDateTimeInContentParams(item)
		}
		return changes
	}

	return 0
}

func resetDateTimeInContentParamsMap(v map[string]interface{}) int {
	changes := 0

	formatKey, hasFormat := findMapKey(v, isDateTimeFormatKey)
	customFormatKey, hasCustomFormat := findMapKey(v, isCustomDateTimeFormatKey)

	if hasFormat && hasCustomFormat {
		if isCustomSelectorValue(v[formatKey]) {
			customFormatValue, ok := extractComparableString(v[customFormatKey])
			if ok && isTargetDateTimeFormat(customFormatValue) {
				if updatedValue, changed := replaceCustomWithDefault(v[formatKey]); changed {
					v[formatKey] = updatedValue
					changes++
				}

				delete(v, customFormatKey)
				changes++
			}
		}
	}

	for key, value := range v {
		if normalizeKey(key) != "expression" {
			continue
		}

		expression, ok := value.(string)
		if !ok {
			continue
		}

		updatedExpression, changed := replaceTargetFormatDateTimeExpression(expression)
		if changed {
			v[key] = updatedExpression
			changes++
		}
	}

	return changes
}

func isDateTimeExpressionValue(value interface{}) bool {
	expression, ok := value.(string)
	if !ok {
		return false
	}

	_, _, matched := parseFormatDateTimeExpression(expression)
	return matched
}

func replaceTargetFormatDateTimeExpression(expression string) (string, bool) {
	expressionValue, formatValue, ok := parseFormatDateTimeExpression(expression)
	if !ok || !isTargetDateTimeFormat(formatValue) {
		return expression, false
	}

	return expressionValue, true
}

func parseFormatDateTimeExpression(expression string) (string, string, bool) {
	matches := formatDateTimeExpressionRegex.FindStringSubmatch(expression)
	if len(matches) != 3 {
		return "", "", false
	}

	return strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2]), true
}

func isTargetDateTimeFormat(value string) bool {
	target := strings.TrimSpace(targetCustomDateTimeFormat)
	if strings.TrimSpace(value) == target {
		return true
	}

	trimmedTarget := strings.TrimPrefix(target, ",")
	return strings.TrimSpace(value) == trimmedTarget
}

func replaceCustomWithDefault(value interface{}) (interface{}, bool) {
	switch v := value.(type) {
	case string:
		if strings.EqualFold(strings.TrimSpace(v), "custom") {
			return "default", true
		}
		return value, false

	case map[string]interface{}:
		changed := false

		if pv, ok := v["PrimitiveValue"].(string); ok && strings.EqualFold(strings.TrimSpace(pv), "custom") {
			v["PrimitiveValue"] = "default"
			changed = true
		}
		if rawValue, ok := v["Value"].(string); ok && strings.EqualFold(strings.TrimSpace(rawValue), "custom") {
			v["Value"] = "default"
			changed = true
		}
		if key, ok := v["_Key"].(string); ok && strings.EqualFold(strings.TrimSpace(key), "custom") {
			v["_Key"] = "default"
			changed = true
		}

		return v, changed
	}

	return value, false
}

func isCustomSelectorValue(value interface{}) bool {
	_, changed := replaceCustomWithDefault(value)
	return changed
}

func extractComparableString(value interface{}) (string, bool) {
	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v), true
	case map[string]interface{}:
		if s, ok := v["PrimitiveValue"].(string); ok {
			return strings.TrimSpace(s), true
		}
		if s, ok := v["Value"].(string); ok {
			return strings.TrimSpace(s), true
		}
		if s, ok := v["_Key"].(string); ok {
			return strings.TrimSpace(s), true
		}
	}

	return "", false
}

func findMapKey(v map[string]interface{}, predicate func(string) bool) (string, bool) {
	for key := range v {
		if predicate(normalizeKey(key)) {
			return key, true
		}
	}
	return "", false
}

func isDateTimeFormatKey(normalizedKey string) bool {
	return normalizedKey == "formatdatetime" || normalizedKey == "formatdatatime"
}

func isCustomDateTimeFormatKey(normalizedKey string) bool {
	return strings.Contains(normalizedKey, "custom") &&
		(strings.Contains(normalizedKey, "formatdatetime") || strings.Contains(normalizedKey, "formatdatatime"))
}

func normalizeKey(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(strings.TrimSpace(s)) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
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
