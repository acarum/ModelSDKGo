# Find Signal Manager Subscriptions

Search for Signal Manager widgets and extract their signal subscriptions across all pages and snippets in a Mendix project.

## Description

This tool scans a Mendix MPR file and reports all Signal Manager widgets (`siemens.mxtosignal.MxToSignal`) along with their signal subscriptions including:
- Signal Name
- App Name
- Subscription Filter

It uses only the modelsdk-go library API without direct database or BSON access.

## Usage

```bash
go run main.go <path-to-mpr-file>
```

## Example

```bash
go run main.go "C:\Projects\MyApp.mpr"
```

## Output

The tool produces a formatted report with:

1. **Project Information**
   - Project path
   - MPR version
   - Mendix version

2. **Pages Section**
   - List of pages containing Signal Manager widgets
   - Number of subscriptions found

3. **Snippets Section**
   - List of snippets containing Signal Manager widgets
   - Number of subscriptions found

4. **Signal Subscriptions Summary**
   - Grouped by module
   - Shows Signal Name, App Name, and Filter for each subscription

5. **Summary**
   - Total pages with Signal Manager
   - Total snippets with Signal Manager
   - Total subscriptions found

## Example Output

```
=== Signal Manager Widget Finder ===
Project: C:\Projects\MyApp.mpr
MPR Version: 2
Mendix Version: 11.9.0

Searching for widget: siemens.mxtosignal.MxToSignal

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📄 SCANNING PAGES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Found 103 pages to scan

✓ Found Signal Manager in 1 page(s) with 2 subscription(s)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📝 SCANNING SNIPPETS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Found 15 snippets to scan

✓ Found Signal Manager in 1 snippet(s) with 1 subscription(s)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📊 SIGNAL SUBSCRIPTIONS SUMMARY
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Total subscriptions: 3

Module: OpcenterEXFN_ReferenceData (3 subscription(s))
  1. StateMachine_Details (Page)
     Signal Name: SN
     App Name: APPName
     Filter: filter0

  2. StateMachine_Details (Page)
     Signal Name: SN2
     App Name: AN2
     Filter: filter

  3. MySnippet (Snippet)
     Signal Name: SNIPPET_SN
     App Name: SNIPPET_APPName
     Filter: SNIPPET_filter0

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Pages with Signal Manager:    1
Snippets with Signal Manager: 1
Total subscriptions:          3
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

## Features

- ✅ Uses only modelsdk-go API (no direct database access)
- ✅ Scans both pages and snippets
- ✅ Extracts signal subscription details
- ✅ Progress indicators during scanning
- ✅ Module name resolution
- ✅ Groups results by module
- ✅ Formatted, readable output

## Technical Details

### Widget Detection

The tool searches for widgets with:
- `$Type` = `"CustomWidgets$CustomWidget"`
- `Type.WidgetId` = `"siemens.mxtosignal.MxToSignal"`

### Subscription Extraction

The tool extracts the following properties from each widget:
- **Signal Name**: Found in properties containing "signalname" or ending with "_sn"
- **App Name**: Found in properties containing "appname"
- **Subscription Filter**: Found in properties containing "subscriptionfilter" or "filter"

### API Usage

- `modelsdk.Open()` - Opens MPR file
- `reader.ListModules()` - Gets module list for name mapping
- `reader.ListPages()` - Gets all pages
- `reader.GetPage(id)` - Gets full page details with widgets
- `reader.ListSnippets()` - Gets all snippets with widgets

### Search Algorithm

1. Marshal page/snippet struct to JSON
2. Unmarshal to map[string]interface{} for inspection
3. Recursively search for CustomWidget with Signal Manager widgetId
4. Extract properties from each widget
5. Match properties to subscription fields (SignalName, AppName, Filter)
6. Collect and display results grouped by module

## Comparison with find_custom_widgets

This tool is similar to `find_custom_widgets` but:
- Uses only modelsdk-go API (no direct BSON access)
- Specifically targets Signal Manager widgets
- Automatically extracts subscription information
- Provides grouped summary output

## Related Examples

- `find_datagrid` - Find DataGrid widgets using API-only approach
- `find_custom_widgets` - Find custom widgets by ID using BSON access
- `export_manifest` - Full project manifest with Signal Manager subscriptions
- `inspect_page` - Inspect detailed page structure

## Building

```bash
go build -o find_signal_manager.exe
```

## Running

```bash
.\find_signal_manager.exe "path\to\project.mpr"
```
