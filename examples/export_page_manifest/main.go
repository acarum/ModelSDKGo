package main

import (
	"database/sql"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthropics/modelsdk-go"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GalleryInfo holds sort and search widget names found inside a Gallery widget
type GalleryInfo struct {
	SortBy        []string // names of DropdownSort widgets
	Search        []string // names of DatagridTextFilter widgets
	TileContainer string   // name of DivContainer with Class "exfn-tile-container"
}

// WidgetCaption holds a single widget with an extractable caption
type WidgetCaption struct {
	WidgetType string
	WidgetName string
	Caption    string
	// For DataGrid2: columns and action buttons are nested under the widget entry
	Columns       []DataGridColumn
	ActionButtons []DataGridActionButton
	// For Gallery: sort and search info
	Gallery *GalleryInfo
}

// DataGridActionButton represents an action button inside a DataGrid 2 widget
type DataGridActionButton struct {
	Name          string
	Caption       string
	ContainerName string // name of the direct parent DivContainer
	NanoflowName  string // nanoflow called on click (from parent DivContainer.OnClickAction)
	ShowPageName  string // page opened by ShowFormAction inside that nanoflow
}

// DataGridColumn represents a column in a DataGrid 2 widget
type DataGridColumn struct {
	Name    string
	Caption string
}

// TabInfo represents a tab found inside a TabContainer
type TabInfo struct {
	Name    string
	Caption string
	Widgets []WidgetCaption
}

// PlaceholderContent holds widgets and tabs found in a layout placeholder
type PlaceholderContent struct {
	Name                    string
	Parameter               string          // full parameter path e.g. "Atlas_Default.Main"
	RootContainerName       string          // name of the first DivContainer in the placeholder
	CommandBarContainerName string          // name of DivContainer with CSS class containing "vertical-command-bar"
	Widgets                 []WidgetCaption // all widgets with captions (for Contents section)
	Tabs                    []TabInfo
	Buttons                 []DataGridActionButton // direct ActionButtons outside of tabs (e.g. Right command bar)
}

// PageReport represents the full report for one page
type PageReport struct {
	PageName  string
	PageTitle string
	PageID    string
	MprPath   string
	Main      *PlaceholderContent
	Right     *PlaceholderContent
}

func main() {
	typeFlag := flag.String("type", "", "Filter pages by type: 'eng' or 'runtime' (default: all)")
	filterFlag := flag.String("filter", "", "Wildcard pattern to match page names (e.g. PANEL_*)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: export_page_manifest [--type eng|runtime] [--filter PATTERN] <mpr_file_path> [page_name]")
		fmt.Println("Example: export_page_manifest MyApp.mpr")
		fmt.Println("         export_page_manifest MyApp.mpr StateMachine_Details")
		fmt.Println("         export_page_manifest --filter PANEL_* MyApp.mpr")
		fmt.Println("         export_page_manifest --type eng MyApp.mpr")
		os.Exit(1)
	}

	if *typeFlag != "" && *typeFlag != "eng" && *typeFlag != "runtime" {
		log.Fatalf("Invalid --type value %q: must be 'eng' or 'runtime'", *typeFlag)
	}

	mprPath := args[0]
	pageFilter := ""
	if len(args) > 1 {
		pageFilter = args[1]
	}
	wildcardFilter := *filterFlag
	outputDir := filepath.Join("examples", "export_page_manifest")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		log.Fatalf("Error opening MPR: %v", err)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n\n", reader.Version())

	// Build module map for marketplace filtering
	modules, err := reader.ListModules()
	if err != nil {
		log.Fatalf("Error listing modules: %v", err)
	}
	moduleMap := make(map[string]string) // containerID -> moduleName
	for _, m := range modules {
		moduleMap[string(m.ID)] = m.Name
	}

	pages, err := reader.ListPages()
	if err != nil {
		log.Fatalf("Error listing pages: %v", err)
	}

	fmt.Printf("Scanning %d pages...\n\n", len(pages))

	var reports []PageReport

	for i, page := range pages {
		if pageFilter != "" && page.Name != pageFilter {
			continue
		}
		if wildcardFilter != "" {
			matched, matchErr := filepath.Match(wildcardFilter, page.Name)
			if matchErr != nil || !matched {
				continue
			}
		}

		// Skip marketplace modules
		if moduleName, ok := moduleMap[string(page.ContainerID)]; ok {
			if isMarketplaceModule(moduleName) {
				continue
			}
		}

		// Filter by page type if requested
		if *typeFlag != "" && !matchesPageType(page.Name, *typeFlag) {
			continue
		}

		fmt.Printf("\r[%d/%d] Scanning: %s", i+1, len(pages), page.Name)

		bsonData, err := loadPageBSON(mprPath, string(page.ID))
		if err != nil {
			if pageFilter != "" {
				fmt.Printf("\n[DEBUG] loadPageBSON failed for %s (ID=%s): %v\n", page.Name, page.ID, err)
			}
			continue
		}

		var pageData map[string]interface{}
		if err := bson.Unmarshal(bsonData, &pageData); err != nil {
			if pageFilter != "" {
				fmt.Printf("\n[DEBUG] bson.Unmarshal failed for %s: %v\n", page.Name, err)
			}
			continue
		}

		report := extractPlaceholders(pageData, page.Name, string(page.ID), mprPath)
		report.MprPath = mprPath
		// Extract page title using en_US
		if page.Title != nil {
			report.PageTitle = page.Title.GetTranslation("en_US")
		}
		if pageFilter != "" {
			fmt.Printf("\n[DEBUG] %s → Main=%v Right=%v\n", page.Name, report.Main != nil, report.Right != nil)
		}
		if report.Main != nil || report.Right != nil {
			reports = append(reports, report)
		}
	}

	fmt.Printf("\n\nPages with Main/Right placeholders: %d\n\n", len(reports))

	if wildcardFilter != "" {
		// One file per page
		for _, r := range reports {
			outputFile := filepath.Join(outputDir, r.PageName+"_manifest.md")
			exportMarkdown([]PageReport{r}, outputFile, r.PageName)
		}
	} else {
		outputFileName := "buttons_manifest.md"
		if pageFilter != "" {
			outputFileName = pageFilter + "_manifest.md"
		}
		outputFile := filepath.Join(outputDir, outputFileName)
		exportMarkdown(reports, outputFile, pageFilter)
	}
}

// matchesPageType returns true if the page matches the requested type filter.
// type "eng"     → engineering pages (TODO: add your criteria here)
// type "runtime" → runtime/user-facing pages (TODO: add your criteria here)
func matchesPageType(pageName, pageType string) bool {
	switch pageType {
	case "eng":
		// TODO: replace with actual eng-page detection logic
		return true
	case "runtime":
		// TODO: replace with actual runtime-page detection logic
		return true
	}
	return true
}

// extractPlaceholders parses FormCall.Arguments to find Main and Right placeholders
func extractPlaceholders(pageData map[string]interface{}, pageName, pageID, mprPath string) PageReport {
	report := PageReport{PageName: pageName, PageID: pageID}

	formCall, ok := getMap(pageData, "FormCall")
	if !ok {
		return report
	}

	args := getArray(formCall, "Arguments")
	for _, arg := range args {
		argMap, ok := arg.(map[string]interface{})
		if !ok {
			continue
		}

		// Check if it's a FormCallArgument
		argType, _ := argMap["$Type"].(string)
		if argType != "Forms$FormCallArgument" {
			continue
		}

		param, _ := argMap["Parameter"].(string)
		placeholderName := placeholderSuffix(param) // e.g. "Main", "Right"

		widgets := getArray(argMap, "Widgets")
		tabs := findTabs(widgets)

		// Resolve ShowPage for DataGrid action buttons inside tabs
		for ti := range tabs {
			for wi := range tabs[ti].Widgets {
				for bi := range tabs[ti].Widgets[wi].ActionButtons {
					btn := &tabs[ti].Widgets[wi].ActionButtons[bi]
					if btn.NanoflowName != "" {
						btn.ShowPageName = loadNanoflowShowPage(mprPath, btn.NanoflowName)
					}
				}
			}
		}

		// Collect direct ActionButtons at placeholder level (e.g. Right vertical-command-bar)
		// Only when there are no tabs, to avoid duplicating DataGrid toolbar buttons
		var directButtons []DataGridActionButton
		if len(tabs) == 0 {
			directButtons = findAllActionButtons(widgets)
			// Resolve ShowPage for each button's nanoflow
			for i := range directButtons {
				if directButtons[i].NanoflowName != "" {
					directButtons[i].ShowPageName = loadNanoflowShowPage(mprPath, directButtons[i].NanoflowName)
				}
			}
		}

		content := &PlaceholderContent{
			Name:                    placeholderName,
			Parameter:               param,
			RootContainerName:       findFirstContainerName(widgets),
			CommandBarContainerName: findContainerByClass(widgets, "vertical-command-bar"),
			Widgets:                 findAllWidgetsSummary(widgets),
			Tabs:                    tabs,
			Buttons:                 directButtons,
		}

		switch strings.ToLower(placeholderName) {
		case "main":
			report.Main = content
		case "right":
			report.Right = content
		}
	}

	return report
}

// placeholderSuffix extracts the last part of a placeholder path (e.g. "Atlas_Default.Main" → "Main")
func placeholderSuffix(param string) string {
	parts := strings.Split(param, ".")
	if len(parts) == 0 {
		return param
	}
	return parts[len(parts)-1]
}

// findTabs recursively searches for Forms$TabPage inside Forms$TabContainer
func findTabs(data interface{}) []TabInfo {
	var tabs []TabInfo

	switch v := data.(type) {
	case map[string]interface{}:
		t, _ := v["$Type"].(string)
		if t == "Forms$TabPage" {
			tab := TabInfo{
				Name:    getStr(v, "Name"),
				Caption: extractCaption(v),
				Widgets: findWidgetCaptions(v),
			}
			tabs = append(tabs, tab)
			// Still recurse to find nested tabs
			for _, val := range v {
				tabs = append(tabs, findTabs(val)...)
			}
			return tabs
		}
		// Recurse into all fields
		for _, val := range v {
			tabs = append(tabs, findTabs(val)...)
		}
	case primitive.A:
		for _, item := range v {
			tabs = append(tabs, findTabs(item)...)
		}
	case []interface{}:
		for _, item := range v {
			tabs = append(tabs, findTabs(item)...)
		}
	}

	return tabs
}

// extractBinaryKey returns a consistent string key from a BSON binary field (primitive.Binary or {Data: base64}).
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

// buildTypePointerMap traverses a widget's type tree to map TypePointer.Data → PropertyKey
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
	case primitive.A:
		for _, item := range v {
			buildTypePointerMap(item, m)
		}
	case []interface{}:
		for _, item := range v {
			buildTypePointerMap(item, m)
		}
	}
}

// getTypePointerData extracts a consistent key from a WidgetProperty's TypePointer field
func getTypePointerData(prop map[string]interface{}) string {
	return extractBinaryKey(prop["TypePointer"])
}

// findContainerByClass returns the Name of the first Forms$DivContainer whose
// Appearance.Class contains the given CSS class substring.
func findContainerByClass(data interface{}, cssClass string) string {
	switch v := data.(type) {
	case map[string]interface{}:
		if v["$Type"] == "Forms$DivContainer" {
			if app, ok := v["Appearance"].(map[string]interface{}); ok {
				if cls, ok := app["Class"].(string); ok && strings.Contains(cls, cssClass) {
					return getStr(v, "Name")
				}
			}
		}
		for _, val := range v {
			if r := findContainerByClass(val, cssClass); r != "" {
				return r
			}
		}
	case primitive.A:
		for _, item := range v {
			if r := findContainerByClass(item, cssClass); r != "" {
				return r
			}
		}
	case []interface{}:
		for _, item := range v {
			if r := findContainerByClass(item, cssClass); r != "" {
				return r
			}
		}
	}
	return ""
}

// findFirstContainerName returns the Name of the first Forms$DivContainer found in data.
func findFirstContainerName(data interface{}) string {
	switch v := data.(type) {
	case map[string]interface{}:
		if v["$Type"] == "Forms$DivContainer" {
			return getStr(v, "Name")
		}
		for _, val := range v {
			if r := findFirstContainerName(val); r != "" {
				return r
			}
		}
	case primitive.A:
		for _, item := range v {
			if r := findFirstContainerName(item); r != "" {
				return r
			}
		}
	case []interface{}:
		for _, item := range v {
			if r := findFirstContainerName(item); r != "" {
				return r
			}
		}
	}
	return ""
}

// findAllActionButtons recursively collects all Forms$ActionButton nodes,
// propagating the nanoflow name from any enclosing DivContainer.OnClickAction.
func findAllActionButtons(data interface{}) []DataGridActionButton {
	return findAllActionButtonsWithNF(data, "", "")
}

func findAllActionButtonsWithNF(data interface{}, inheritedNanoflow, inheritedContainer string) []DataGridActionButton {
	var result []DataGridActionButton
	switch v := data.(type) {
	case map[string]interface{}:
		// If this DivContainer has a nanoflow OnClickAction, propagate nanoflow and container name to children
		nf := inheritedNanoflow
		containerName := inheritedContainer
		if v["$Type"] == "Forms$DivContainer" {
			if oca, ok := v["OnClickAction"].(map[string]interface{}); ok {
				if oca["$Type"] == "Forms$CallNanoflowClientAction" {
					if n, ok := oca["Nanoflow"].(string); ok && n != "" {
						nf = n
						containerName = getStr(v, "Name")
					}
				}
			}
		}
		if v["$Type"] == "Forms$ActionButton" {
			result = append(result, DataGridActionButton{
				Name:          getStr(v, "Name"),
				Caption:       extractCaption(v),
				ContainerName: containerName,
				NanoflowName:  nf,
			})
			return result
		}
		for _, val := range v {
			result = append(result, findAllActionButtonsWithNF(val, nf, containerName)...)
		}
	case primitive.A:
		for _, item := range v {
			result = append(result, findAllActionButtonsWithNF(item, inheritedNanoflow, inheritedContainer)...)
		}
	case []interface{}:
		for _, item := range v {
			result = append(result, findAllActionButtonsWithNF(item, inheritedNanoflow, inheritedContainer)...)
		}
	}
	return result
}

// loadNanoflowShowPage opens the MPR and searches for a nanoflow by full qualified name
// (e.g. "Module.NanoflowName") and returns the page opened by its ShowFormAction, if any.
func loadNanoflowShowPage(mprPath, nanoflowFullName string) string {
	if nanoflowFullName == "" {
		return ""
	}
	parts := strings.SplitN(nanoflowFullName, ".", 2)
	if len(parts) != 2 {
		return ""
	}
	nfShortName := parts[1]

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		return ""
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var idBytes []byte
		if err := rows.Scan(&idBytes); err != nil {
			continue
		}
		uid := bytesToUUID(idBytes)
		if uid == "" {
			continue
		}
		p1, p2 := uid[:2], uid[2:4]
		data, err := os.ReadFile(filepath.Join(contentsDir, p1, p2, uid+".mxunit"))
		if err != nil {
			continue
		}
		var m map[string]interface{}
		if bson.Unmarshal(data, &m) != nil {
			continue
		}
		if m["$Type"] != "Microflows$Nanoflow" {
			continue
		}
		name, _ := m["Name"].(string)
		if name != nfShortName {
			continue
		}
		// Found it — scan for ShowFormAction
		return findShowFormPage(m)
	}
	return ""
}

// findShowFormPage recursively searches for Microflows$ShowFormAction and returns FormSettings.Form
func findShowFormPage(data interface{}) string {
	switch v := data.(type) {
	case map[string]interface{}:
		if v["$Type"] == "Microflows$ShowFormAction" {
			if fs, ok := v["FormSettings"].(map[string]interface{}); ok {
				if f, ok := fs["Form"].(string); ok && f != "" {
					return f
				}
			}
		}
		for _, val := range v {
			if r := findShowFormPage(val); r != "" {
				return r
			}
		}
	case primitive.A:
		for _, item := range v {
			if r := findShowFormPage(item); r != "" {
				return r
			}
		}
	case []interface{}:
		for _, item := range v {
			if r := findShowFormPage(item); r != "" {
				return r
			}
		}
	}
	return ""
}

// loadPanelWidgets opens the MPR and searches for a page by full qualified name
// (e.g. "Module.PageName") and returns the widgets found on it.
func loadPanelWidgets(mprPath, fullPageName string) []WidgetCaption {
	if fullPageName == "" {
		return nil
	}
	parts := strings.SplitN(fullPageName, ".", 2)
	if len(parts) != 2 {
		return nil
	}
	pageShortName := parts[1]

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		return nil
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var idBytes []byte
		if err := rows.Scan(&idBytes); err != nil {
			continue
		}
		uid := bytesToUUID(idBytes)
		if uid == "" {
			continue
		}
		p1, p2 := uid[:2], uid[2:4]
		data, err := os.ReadFile(filepath.Join(contentsDir, p1, p2, uid+".mxunit"))
		if err != nil {
			continue
		}
		var m map[string]interface{}
		if bson.Unmarshal(data, &m) != nil {
			continue
		}
		if m["$Type"] != "Forms$Page" {
			continue
		}
		name, _ := m["Name"].(string)
		if name != pageShortName {
			continue
		}
		return findWidgetCaptions(m)
	}
	return nil
}

// bytesToUUID converts a 16-byte SQL BLOB to UUID string (Windows GUID byte order)
func bytesToUUID(b []byte) string {
	if len(b) != 16 {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString([]byte{b[3], b[2], b[1], b[0]}),
		hex.EncodeToString([]byte{b[5], b[4]}),
		hex.EncodeToString([]byte{b[7], b[6]}),
		hex.EncodeToString(b[8:10]),
		hex.EncodeToString(b[10:16]))
}

// extractDataGridActionButtons recursively collects all Forms$ActionButton nodes inside a DataGrid 2 widget,
// propagating the nanoflow from any enclosing DivContainer.OnClickAction and excluding row-inline buttons.
func extractDataGridActionButtons(data interface{}) []DataGridActionButton {
	return extractDataGridActionButtonsWithNF(data, "")
}

func extractDataGridActionButtonsWithNF(data interface{}, inheritedNF string) []DataGridActionButton {
	var result []DataGridActionButton
	switch v := data.(type) {
	case map[string]interface{}:
		nf := inheritedNF
		if v["$Type"] == "Forms$DivContainer" {
			if oca, ok := v["OnClickAction"].(map[string]interface{}); ok {
				if oca["$Type"] == "Forms$CallNanoflowClientAction" {
					if n, ok := oca["Nanoflow"].(string); ok && n != "" {
						nf = n
					}
				}
			}
		}
		if v["$Type"] == "Forms$ActionButton" {
			// Skip inline row buttons (RenderType "Link"); only include toolbar buttons (RenderType "Button")
			if getStr(v, "RenderType") == "Link" {
				return result
			}
			result = append(result, DataGridActionButton{
				Name:         getStr(v, "Name"),
				Caption:      extractCaption(v),
				NanoflowName: nf,
			})
			return result
		}
		for _, val := range v {
			result = append(result, extractDataGridActionButtonsWithNF(val, nf)...)
		}
	case primitive.A:
		for _, item := range v {
			result = append(result, extractDataGridActionButtonsWithNF(item, inheritedNF)...)
		}
	case []interface{}:
		for _, item := range v {
			result = append(result, extractDataGridActionButtonsWithNF(item, inheritedNF)...)
		}
	}
	return result
}

// extractDataGridColumns extracts column names and captions from a DataGrid 2 custom widget
func extractDataGridColumns(widget map[string]interface{}) []DataGridColumn {
	// Build TypePointer.Data → PropertyKey map from this widget's type definition
	typeMap := make(map[string]string)
	buildTypePointerMap(widget, typeMap)

	obj, ok := widget["Object"].(map[string]interface{})
	if !ok {
		return nil
	}

	// Find the "columns" property in the widget's top-level properties
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

		var columns []DataGridColumn
		colObjects := toSlice(val["Objects"])
		for colIdx, colObj := range colObjects {
			colMap, ok := colObj.(map[string]interface{})
			if !ok || colMap["$Type"] != "CustomWidgets$WidgetObject" {
				continue
			}
			colProps := getArray(colMap, "Properties")
			caption := ""
			colName := fmt.Sprintf("col%d", colIdx+1)
			for _, cp := range colProps {
				cpMap, ok := cp.(map[string]interface{})
				if !ok {
					continue
				}
				propKey := typeMap[getTypePointerData(cpMap)]
				valMap, ok := cpMap["Value"].(map[string]interface{})
				if !ok {
					continue
				}
				switch propKey {
				case "caption", "header":
					// Caption stored as TextTemplate (Forms$ClientTemplate)
					if tt, ok := valMap["TextTemplate"].(map[string]interface{}); ok {
						if tmpl, ok := tt["Template"].(map[string]interface{}); ok {
							caption = extractTextsText(tmpl)
						}
					}
				case "id":
					if s, ok := valMap["PrimitiveValue"].(string); ok && s != "" {
						colName = s
					}
				}
			}
			// Fallback: if no typed match, scan all properties for any TextTemplate
			if caption == "" {
				for _, cp := range colProps {
					cpMap, ok := cp.(map[string]interface{})
					if !ok {
						continue
					}
					valMap, ok := cpMap["Value"].(map[string]interface{})
					if !ok {
						continue
					}
					if tt, ok := valMap["TextTemplate"].(map[string]interface{}); ok {
						if tmpl, ok := tt["Template"].(map[string]interface{}); ok {
							if t := extractTextsText(tmpl); t != "" {
								caption = t
								break
							}
						}
					}
				}
			}
			columns = append(columns, DataGridColumn{Name: colName, Caption: caption})
		}
		return columns
	}
	return nil
}

// findWidgetCaptions recursively collects all widgets with non-empty captions
func findWidgetCaptions(data interface{}) []WidgetCaption {
	var result []WidgetCaption

	switch v := data.(type) {
	case map[string]interface{}:
		t, _ := v["$Type"].(string)
		if t != "" && t != "Forms$TabPage" {
			short := shortType(t)
			// Skip widget property descriptors (custom widget metadata)
			if !strings.HasPrefix(short, "Widget") {
				// Special handling for CustomWidget: check WidgetId
				if t == "CustomWidgets$CustomWidget" {
					widgetID := ""
					if typeNode, ok := v["Type"].(map[string]interface{}); ok {
						widgetID, _ = typeNode["WidgetId"].(string)
					}
					if widgetID == "com.mendix.widget.web.datagrid.Datagrid" {
						name := getStr(v, "Name")
						cols := extractDataGridColumns(v)
						btns := extractDataGridActionButtons(v)
						result = append(result, WidgetCaption{
							WidgetType:    "DataGrid2",
							WidgetName:    name,
							Caption:       "(Data Grid 2)",
							Columns:       cols,
							ActionButtons: btns,
						})
						return result // don't recurse into the DataGrid internals
					}
				} else {
					cap := extractCaption(v)
					if cap != "" {
						result = append(result, WidgetCaption{
							WidgetType: short,
							WidgetName: getStr(v, "Name"),
							Caption:    cap,
						})
					}
				}
			}
		}
		for _, val := range v {
			result = append(result, findWidgetCaptions(val)...)
		}
	case primitive.A:
		for _, item := range v {
			result = append(result, findWidgetCaptions(item)...)
		}
	case []interface{}:
		for _, item := range v {
			result = append(result, findWidgetCaptions(item)...)
		}
	}

	return result
}

// extractGalleryInfo scans a Gallery widget node for DropdownSort, DatagridTextFilter children
// and the first DivContainer with Class "exfn-tile-container".
func extractGalleryInfo(data interface{}) *GalleryInfo {
	var sortBy, search []string
	tileContainer := ""
	var walk func(v interface{})
	walk = func(v interface{}) {
		switch m := v.(type) {
		case map[string]interface{}:
			t, _ := m["$Type"].(string)
			if t == "CustomWidgets$CustomWidget" {
				if typeNode, ok := m["Type"].(map[string]interface{}); ok {
					wid, _ := typeNode["WidgetId"].(string)
					name := getStr(m, "Name")
					switch wid {
					case "com.mendix.widget.web.dropdownsort.DropdownSort":
						sortBy = append(sortBy, name)
						return // don't recurse into this widget's type definition
					case "com.mendix.widget.web.datagridtextfilter.DatagridTextFilter":
						search = append(search, name)
						return
					}
				}
			}
			if t == "Forms$DivContainer" && tileContainer == "" {
				if app, ok := m["Appearance"].(map[string]interface{}); ok {
					if cls, _ := app["Class"].(string); cls == "exfn-tile-container" {
						tileContainer = getStr(m, "Name")
					}
				}
			}
			for _, val := range m {
				walk(val)
			}
		case primitive.A:
			for _, item := range m {
				walk(item)
			}
		case []interface{}:
			for _, item := range m {
				walk(item)
			}
		}
	}
	walk(data)
	if len(sortBy) == 0 && len(search) == 0 && tileContainer == "" {
		return nil
	}
	return &GalleryInfo{SortBy: sortBy, Search: search, TileContainer: tileContainer}
}

// findAllWidgetsSummary recursively collects all widget-like nodes with their type, name and caption.
// Unlike findWidgetCaptions, it includes widgets even when caption is empty.
func findAllWidgetsSummary(data interface{}) []WidgetCaption {
	var result []WidgetCaption
	// widget types to include (skip containers, layout helpers, etc.)
	include := map[string]bool{
		"Forms$TextBox": true, "Forms$TextArea": true, "Forms$DatePicker": true,
		"Forms$DropDown": true, "Forms$CheckBox": true, "Forms$RadioButton": true,
		"Forms$Label": true, "Forms$StaticLabel": true, "Forms$ActionButton": true,
		"Forms$DataView": true, "Forms$ListView": true, "Forms$ReferenceSelector": true,
		"Forms$InputReferenceSelector": true, "Forms$FileManager": true,
		"Forms$Image": true, "Forms$DynamicImage": true,
	}
	switch v := data.(type) {
	case map[string]interface{}:
		t, _ := v["$Type"].(string)
		if t == "CustomWidgets$CustomWidget" {
			if typeNode, ok := v["Type"].(map[string]interface{}); ok {
				wid, _ := typeNode["WidgetId"].(string)
				switch wid {
				case "com.mendix.widget.web.gallery.Gallery":
					gi := extractGalleryInfo(v)
					result = append(result, WidgetCaption{
						WidgetType: "Gallery",
						WidgetName: getStr(v, "Name"),
						Caption:    "",
						Gallery:    gi,
					})
					return result // don't recurse into Gallery internals
				case "com.mendix.widget.web.datagrid.Datagrid":
					cols := extractDataGridColumns(v)
					result = append(result, WidgetCaption{
						WidgetType: "DataGrid2",
						WidgetName: getStr(v, "Name"),
						Caption:    "",
						Columns:    cols,
					})
					return result // don't recurse into DataGrid2 internals
				case "com.mendix.widget.custom.switch.Switch":
					result = append(result, WidgetCaption{
						WidgetType: "Switch",
						WidgetName: getStr(v, "Name"),
						Caption:    "",
					})
					return result
				}
			}
		}
		if include[t] {
			cap := extractCaption(v)
			result = append(result, WidgetCaption{
				WidgetType: shortType(t),
				WidgetName: getStr(v, "Name"),
				Caption:    cap,
			})
		}
		for _, val := range v {
			result = append(result, findAllWidgetsSummary(val)...)
		}
	case primitive.A:
		for _, item := range v {
			result = append(result, findAllWidgetsSummary(item)...)
		}
	case []interface{}:
		for _, item := range v {
			result = append(result, findAllWidgetsSummary(item)...)
		}
	}
	return result
}

// shortType strips the "Forms$" prefix for readability
func shortType(t string) string {
	parts := strings.SplitN(t, "$", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return t
}

// extractTextsText extracts the first non-empty text from a Texts$Text node (Items → Texts$Translation → Text)
// Prefers en_US, falls back to first available language.
func extractTextsText(node map[string]interface{}) string {
	items := getArray(node, "Items")
	var first string
	for _, item := range items {
		t, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		if t["$Type"] != "Texts$Translation" {
			continue
		}
		text, _ := t["Text"].(string)
		if text == "" {
			continue
		}
		lang, _ := t["LanguageCode"].(string)
		if lang == "en_US" {
			return text
		}
		if first == "" {
			first = text
		}
	}
	return first
}

// extractTextsTextEnUS extracts ONLY the en_US text from a Texts$Text node.
// Returns empty string if en_US is absent or empty (does not fall back to other languages).
func extractTextsTextEnUS(node map[string]interface{}) string {
	for _, item := range getArray(node, "Items") {
		t, ok := item.(map[string]interface{})
		if !ok || t["$Type"] != "Texts$Translation" {
			continue
		}
		if lang, _ := t["LanguageCode"].(string); lang == "en_US" {
			text, _ := t["Text"].(string)
			return text // may be empty — caller decides
		}
	}
	return ""
}

// extractCaption gets caption from a widget — handles multiple Mendix caption patterns:
// 1. Caption (plain string)
// 2. Caption (Forms$Text / Texts$Text with Translations/Items)
// 3. CaptionTemplate.Template (Texts$Text) — used by ActionButton
// 4. LabelTemplate.Template (Texts$Text) — used by input widgets
func extractCaption(data map[string]interface{}) string {
	// Pattern 1 & 2: direct Caption field
	if cap, ok := data["Caption"]; ok {
		switch v := cap.(type) {
		case string:
			if v != "" {
				return v
			}
		case map[string]interface{}:
			capType, _ := v["$Type"].(string)
			// Texts$Text (used by TabPage and others)
			if capType == "Texts$Text" {
				if t := extractTextsText(v); t != "" {
					return t
				}
			}
			// Legacy Forms$Text with Translations array
			translations := getArray(v, "Translations")
			for _, tr := range translations {
				tm, ok := tr.(map[string]interface{})
				if !ok {
					continue
				}
				if text, ok := tm["Text"].(string); ok && text != "" {
					return text
				}
			}
			if text, ok := v["Text"].(string); ok && text != "" {
				return text
			}
		}
	}

	// Pattern 3: CaptionTemplate.Template (Forms$ActionButton etc.)
	// Only use en_US text; if empty, icon-only button falls through to Tooltip (Pattern 5).
	if ct, ok := data["CaptionTemplate"].(map[string]interface{}); ok {
		if tmpl, ok := ct["Template"].(map[string]interface{}); ok {
			if t := extractTextsTextEnUS(tmpl); t != "" {
				return t
			}
		}
	}

	// Pattern 4: LabelTemplate.Template (input widgets)
	if lt, ok := data["LabelTemplate"].(map[string]interface{}); ok {
		if tmpl, ok := lt["Template"].(map[string]interface{}); ok {
			if t := extractTextsText(tmpl); t != "" {
				return t
			}
		}
	}

	// Pattern 5: Tooltip (fallback for icon-only buttons with no caption text)
	if tt, ok := data["Tooltip"].(map[string]interface{}); ok {
		if t := extractTextsText(tt); t != "" {
			return t
		}
	}

	return ""
}

// --- helpers ---

func getMap(data map[string]interface{}, key string) (map[string]interface{}, bool) {
	v, ok := data[key]
	if !ok {
		return nil, false
	}
	m, ok := v.(map[string]interface{})
	return m, ok
}

// toSlice converts primitive.A or []interface{} to []interface{}, skipping non-object elements
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
	// Mendix BSON arrays begin with an integer count — keep only maps and arrays
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

func getStr(data map[string]interface{}, key string) string {
	v, ok := data[key]
	if !ok {
		return ""
	}
	s, _ := v.(string)
	return s
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

	return os.ReadFile(filePath)
}

// isMarketplaceModule checks if a module name belongs to a marketplace/system module
func isMarketplaceModule(moduleName string) bool {
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
		"DISW_DesignSystem",
	}
	for _, indicator := range marketplaceIndicators {
		if strings.HasPrefix(moduleName, indicator) {
			return true
		}
	}
	return false
}

// writePlaceholderSection writes one Main/Right section with tabs and widget captions
func writePlaceholderSection(file *os.File, sectionName string, ph *PlaceholderContent, mprPath string) {
	isModalPanel := ph != nil && strings.Contains(ph.Parameter, "EXFN_ModalPanel")
	fmt.Fprintf(file, "### %s\n\n", sectionName)
	if ph == nil {
		fmt.Fprintf(file, "_No %s placeholder in this page._\n\n", sectionName)
		return
	}
	fmt.Fprintf(file, "**Layout:** `%s`\n\n", ph.Parameter)

	// Contents section: all widgets with name and caption (only for EXFN_ModalPanel)
	if isModalPanel && len(ph.Widgets) > 0 {
		fmt.Fprintf(file, "#### Contents\n\n")
		fmt.Fprintf(file, "| # | Type | Name | Caption |\n")
		fmt.Fprintf(file, "|---|---|---|---|\n")
		row := 0
		for _, w := range ph.Widgets {
			if w.WidgetType == "ActionButton" || w.WidgetType == "DataView" {
				continue
			}
			row++
			fmt.Fprintf(file, "| %d | %s | %s | %s |\n", row,
				strings.ReplaceAll(w.WidgetType, "|", "\\|"),
				strings.ReplaceAll(w.WidgetName, "|", "\\|"),
				strings.ReplaceAll(w.Caption, "|", "\\|"))
		}
		fmt.Fprintf(file, "\n")
		// Gallery subsections
		for _, w := range ph.Widgets {
			if w.WidgetType != "Gallery" || w.Gallery == nil {
				continue
			}
			fmt.Fprintf(file, "##### Gallery: %s\n\n", w.WidgetName)
			if w.Gallery.TileContainer != "" {
				fmt.Fprintf(file, "- **Tile container:** %s\n", w.Gallery.TileContainer)
			}
			if len(w.Gallery.SortBy) > 0 {
				fmt.Fprintf(file, "- **Sort by:** %s\n", strings.Join(w.Gallery.SortBy, ", "))
			}
			if len(w.Gallery.Search) > 0 {
				fmt.Fprintf(file, "- **Search:** %s\n", strings.Join(w.Gallery.Search, ", "))
			}
			fmt.Fprintf(file, "\n")
		}
		// DataGrid2 subsections
		for _, w := range ph.Widgets {
			if w.WidgetType != "DataGrid2" || len(w.Columns) == 0 {
				continue
			}
			fmt.Fprintf(file, "##### DataGrid2: %s\n\n", w.WidgetName)
			fmt.Fprintf(file, "| # | Caption |\n")
			fmt.Fprintf(file, "|---|---|\n")
			for i, col := range w.Columns {
				fmt.Fprintf(file, "| %d | %s |\n", i+1,
					strings.ReplaceAll(col.Caption, "|", "\\|"))
			}
			fmt.Fprintf(file, "\n")
		}
	}

	if len(ph.Tabs) == 0 {
		if len(ph.Buttons) == 0 {
			if !isModalPanel {
				fmt.Fprintf(file, "_No tabs found in %s placeholder._\n\n", sectionName)
			}
			return
		}
		if isModalPanel {
			// Skip Vertical CommandBar header for modal panel layout
			fmt.Fprintf(file, "**Buttons:**\n\n")
		} else {
			// No tabs, but has direct command-bar buttons (e.g. Right vertical-command-bar)
			fmt.Fprintf(file, "**Vertical CommandBar**\n\n")
			fmt.Fprintf(file, "| | Name |\n")
			fmt.Fprintf(file, "|---|---|\n")
			rootName := ph.RootContainerName
			if rootName == "" {
				rootName = "-"
			}
			cbName := ph.CommandBarContainerName
			if cbName == "" {
				cbName = "-"
			}
			fmt.Fprintf(file, "| Container | `%s` |\n", rootName)
			fmt.Fprintf(file, "| VerticalCommandBarClass | `%s` |\n\n", cbName)
			fmt.Fprintf(file, "**Vertical CommandBar Buttons:**\n\n")
		}
		fmt.Fprintf(file, "| # | Name | Caption |\n")
		fmt.Fprintf(file, "|---|---|---|\n")
		for i, btn := range ph.Buttons {
			btnCap := btn.Caption
			if btnCap == "" {
				btnCap = "(no caption)"
			}
			fmt.Fprintf(file, "| %d | %s | %s |\n", i+1,
				strings.ReplaceAll(btn.Name, "|", "\\|"),
				strings.ReplaceAll(btnCap, "|", "\\|"))
		}
		fmt.Fprintf(file, "\n")

		// Button → Page navigation section (only buttons that open a page)
		var navButtons []DataGridActionButton
		for _, btn := range ph.Buttons {
			if btn.ShowPageName != "" {
				navButtons = append(navButtons, btn)
			}
		}
		if len(navButtons) > 0 {
			fmt.Fprintf(file, "**Button → Page Navigation:**\n\n")
			fmt.Fprintf(file, "| Button Caption | Nanoflow | Page Opened |\n")
			fmt.Fprintf(file, "|---|---|---|\n")
			for _, btn := range navButtons {
				nfShort := btn.NanoflowName
				if idx := strings.LastIndex(nfShort, "."); idx >= 0 {
					nfShort = nfShort[idx+1:]
				}
				pageShort := btn.ShowPageName
				if idx := strings.LastIndex(pageShort, "."); idx >= 0 {
					pageShort = pageShort[idx+1:]
				}
				fmt.Fprintf(file, "| %s | %s | %s |\n",
					strings.ReplaceAll(btn.Caption, "|", "\\|"),
					strings.ReplaceAll(nfShort, "|", "\\|"),
					strings.ReplaceAll(pageShort, "|", "\\|"))
			}
			fmt.Fprintf(file, "\n")

			// Per-panel widget details: one subsection per unique page
			seen := make(map[string]bool)
			for _, btn := range navButtons {
				if seen[btn.ShowPageName] {
					continue
				}
				seen[btn.ShowPageName] = true
				pageShort := btn.ShowPageName
				if idx := strings.LastIndex(pageShort, "."); idx >= 0 {
					pageShort = pageShort[idx+1:]
				}
				fmt.Fprintf(file, "**Panel: `%s`**\n\n", pageShort)
				widgets := loadPanelWidgets(mprPath, btn.ShowPageName)
				if len(widgets) == 0 {
					fmt.Fprintf(file, "_No widgets with captions found._\n\n")
				} else {
					fmt.Fprintf(file, "| Widget Type | Name | Caption |\n")
					fmt.Fprintf(file, "|---|---|---|\n")
					for _, w := range widgets {
						wCap := w.Caption
						if wCap == "" {
							wCap = "(no caption)"
						}
						fmt.Fprintf(file, "| %s | %s | %s |\n",
							strings.ReplaceAll(w.WidgetType, "|", "\\|"),
							strings.ReplaceAll(w.WidgetName, "|", "\\|"),
							strings.ReplaceAll(wCap, "|", "\\|"))
					}
					fmt.Fprintf(file, "\n")
				}
			}
		}

		return
	}
	// Tabs overview list
	fmt.Fprintf(file, "**Tabs:**\n\n")
	fmt.Fprintf(file, "| # | Name | Caption |\n")
	fmt.Fprintf(file, "|---|---|---|\n")
	for i, tab := range ph.Tabs {
		tc := tab.Caption
		if tc == "" {
			tc = "(no caption)"
		}
		fmt.Fprintf(file, "| %d | `%s` | %s |\n", i+1,
			strings.ReplaceAll(tab.Name, "|", "\\|"),
			strings.ReplaceAll(tc, "|", "\\|"))
	}
	fmt.Fprintf(file, "\n")

	for i, tab := range ph.Tabs {
		tabCaption := tab.Caption
		if tabCaption == "" {
			tabCaption = "(no caption)"
		}
		fmt.Fprintf(file, "#### Tab %d — `%s`", i+1, tab.Name)
		if tab.Caption != "" {
			fmt.Fprintf(file, " _%s_", strings.ReplaceAll(tab.Caption, "|", "\\|"))
		}
		fmt.Fprintf(file, "\n\n")
		if len(tab.Widgets) == 0 {
			fmt.Fprintf(file, "_No widgets with captions found._\n\n")
		} else {
			fmt.Fprintf(file, "| Widget Type | Name | Caption |\n")
			fmt.Fprintf(file, "|---|---|---|\n")
			for _, w := range tab.Widgets {
				fmt.Fprintf(file, "| %s | %s | %s |\n",
					strings.ReplaceAll(w.WidgetType, "|", "\\|"),
					strings.ReplaceAll(w.WidgetName, "|", "\\|"),
					strings.ReplaceAll(w.Caption, "|", "\\|"))
				// If DataGrid2, add column sub-table
				if w.WidgetType == "DataGrid2" && len(w.Columns) > 0 {
					fmt.Fprintf(file, "\n  **Columns of `%s`:**\n\n", w.WidgetName)
					fmt.Fprintf(file, "  | # | Column Name | Caption |\n")
					fmt.Fprintf(file, "  |---|---|---|\n")
					for j, col := range w.Columns {
						colName := col.Name
						if colName == "" {
							colName = fmt.Sprintf("col%d", j+1)
						}
						colCap := col.Caption
						if colCap == "" {
							colCap = "(no caption)"
						}
						fmt.Fprintf(file, "  | %d | %s | %s |\n", j+1,
							strings.ReplaceAll(colName, "|", "\\|"),
							strings.ReplaceAll(colCap, "|", "\\|"))
					}
					fmt.Fprintf(file, "\n")
				}
				// If DataGrid2, add action buttons sub-table
				if w.WidgetType == "DataGrid2" && len(w.ActionButtons) > 0 {
					fmt.Fprintf(file, "  **Contextual CommandBar Buttons of `%s`:**\n\n", w.WidgetName)
					fmt.Fprintf(file, "  | # | Name | Caption |\n")
					fmt.Fprintf(file, "  |---|---|---|\n")
					for j, btn := range w.ActionButtons {
						btnCap := btn.Caption
						if btnCap == "" {
							btnCap = "(no caption)"
						}
						fmt.Fprintf(file, "  | %d | %s | %s |\n", j+1,
							strings.ReplaceAll(btn.Name, "|", "\\|"),
							strings.ReplaceAll(btnCap, "|", "\\|"))
					}
					fmt.Fprintf(file, "\n")

					// Button → Page navigation for this DataGrid
					var navBtns []DataGridActionButton
					for _, btn := range w.ActionButtons {
						if btn.ShowPageName != "" {
							navBtns = append(navBtns, btn)
						}
					}
					if len(navBtns) > 0 {
						fmt.Fprintf(file, "  **Button \u2192 Page Navigation:**\n\n")
						fmt.Fprintf(file, "  | Button Caption | Nanoflow | Page Opened |\n")
						fmt.Fprintf(file, "  |---|---|---|\n")
						for _, btn := range navBtns {
							nfShort := btn.NanoflowName
							if idx := strings.LastIndex(nfShort, "."); idx >= 0 {
								nfShort = nfShort[idx+1:]
							}
							pageShort := btn.ShowPageName
							if idx := strings.LastIndex(pageShort, "."); idx >= 0 {
								pageShort = pageShort[idx+1:]
							}
							fmt.Fprintf(file, "  | %s | %s | %s |\n",
								strings.ReplaceAll(btn.Caption, "|", "\\|"),
								strings.ReplaceAll(nfShort, "|", "\\|"),
								strings.ReplaceAll(pageShort, "|", "\\|"))
						}
						fmt.Fprintf(file, "\n")

						// Panel widget details
						seen := make(map[string]bool)
						for _, btn := range navBtns {
							if seen[btn.ShowPageName] {
								continue
							}
							seen[btn.ShowPageName] = true
							pageShort := btn.ShowPageName
							if idx := strings.LastIndex(pageShort, "."); idx >= 0 {
								pageShort = pageShort[idx+1:]
							}
							fmt.Fprintf(file, "  **Panel: `%s`**\n\n", pageShort)
							panelWidgets := loadPanelWidgets(mprPath, btn.ShowPageName)
							if len(panelWidgets) == 0 {
								fmt.Fprintf(file, "  _No widgets with captions found._\n\n")
							} else {
								fmt.Fprintf(file, "  | Widget Type | Name | Caption |\n")
								fmt.Fprintf(file, "  |---|---|---|\n")
								for _, pw := range panelWidgets {
									pwCap := pw.Caption
									if pwCap == "" {
										pwCap = "(no caption)"
									}
									fmt.Fprintf(file, "  | %s | %s | %s |\n",
										strings.ReplaceAll(pw.WidgetType, "|", "\\|"),
										strings.ReplaceAll(pw.WidgetName, "|", "\\|"),
										strings.ReplaceAll(pwCap, "|", "\\|"))
								}
								fmt.Fprintf(file, "\n")
							}
						}
					}
				}
			}
			fmt.Fprintf(file, "\n")
		}
	}
}

// exportMarkdown writes the full report to a .md file
func exportMarkdown(reports []PageReport, outputFile string, pageFilter string) {
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating Markdown file: %v", err)
	}
	defer file.Close()

	totalMain := 0
	totalRight := 0
	totalTabs := 0
	totalButtons := 0
	isModalPanelReport := false

	for _, r := range reports {
		if r.Main != nil {
			totalMain++
			totalTabs += len(r.Main.Tabs)
			totalButtons += len(r.Main.Buttons)
			if strings.Contains(r.Main.Parameter, "EXFN_ModalPanel") {
				isModalPanelReport = true
			}
		}
		if r.Right != nil {
			totalRight++
			totalTabs += len(r.Right.Tabs)
		}
	}

	title := "Page Manifest — Layout Placeholders & Tabs"
	if pageFilter != "" {
		title = pageFilter
	}
	fmt.Fprintf(file, "# %s\n\n", title)
	fmt.Fprintf(file, "---\n\n")

	for _, r := range reports {
		fmt.Fprintf(file, "## 📄 %s\n\n", r.PageName)
		if r.PageTitle != "" {
			fmt.Fprintf(file, "**Title:** %s\n\n", r.PageTitle)
		}

		writePlaceholderSection(file, "Main", r.Main, r.MprPath)

		// Skip Right section if the layout is EXFN_ModalPanel (modal panel has no Right placeholder)
		mainParam := ""
		if r.Main != nil {
			mainParam = r.Main.Parameter
		}
		if !strings.Contains(mainParam, "EXFN_ModalPanel") {
			writePlaceholderSection(file, "Right", r.Right, r.MprPath)
		}

		fmt.Fprintf(file, "---\n\n")
	}

	fmt.Fprintf(file, "## Summary\n\n")
	if !isModalPanelReport {
		fmt.Fprintf(file, "- **Pages with placeholders:** %d\n", len(reports))
	}
	if isModalPanelReport {
		fmt.Fprintf(file, "- **Total Buttons:** %d\n", totalButtons)
	} else {
		fmt.Fprintf(file, "- **Total tabs:** %d\n", totalTabs)
	}

	fmt.Printf("✓ Markdown exported to: %s\n", outputFile)
	fmt.Printf("  Pages with Main/Right: %d  |  Total tabs: %d\n", len(reports), totalTabs)
}
