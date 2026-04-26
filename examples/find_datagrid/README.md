# Find DataGrid Widgets

Search for DataGrid widgets across all pages and snippets in a Mendix project.

## Description

This tool scans a Mendix MPR file and reports all pages and snippets that contain DataGrid widgets (`com.mendix.widget.web.datagrid.Datagrid`). It uses only the modelsdk-go library API without direct database or BSON access.

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
   - List of pages containing DataGrid widgets
   - Module name for each page
   - Page URL (if available)
   - Count of DataGrid widgets per page

3. **Snippets Section**
   - List of snippets containing DataGrid widgets
   - Module name for each snippet
   - Count of DataGrid widgets per snippet

4. **Summary**
   - Total pages with DataGrid
   - Total snippets with DataGrid
   - Total documents found

## Example Output

```
=== DataGrid Widget Finder ===
Project: C:\Projects\MyApp.mpr
MPR Version: 2
Mendix Version: 10.18.0

Searching for widget: com.mendix.widget.web.datagrid.Datagrid

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📄 SCANNING PAGES
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Found 103 pages to scan

✓ Found 2 page(s) with DataGrid:

1. UserList
   Module: Administration
   URL: /admin/users
   DataGrid count: 1

2. ProductCatalog
   Module: Sales
   URL: /products
   DataGrid count: 2

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📝 SCANNING SNIPPETS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Found 15 snippets to scan

✓ Found 1 snippet(s) with DataGrid:

1. OrderGrid
   Module: Sales
   DataGrid count: 1

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📊 SUMMARY
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Pages with DataGrid:    2
Snippets with DataGrid: 1
Total documents:        3
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

## Features

- ✅ Uses only modelsdk-go API (no direct database access)
- ✅ Scans both pages and snippets
- ✅ Progress indicators during scanning
- ✅ Module name resolution
- ✅ Widget count per document
- ✅ Formatted, readable output

## Technical Details

### Widget Detection

The tool searches for widgets with `$Type` field matching:
- Exact match: `com.mendix.widget.web.datagrid.Datagrid`
- Case-insensitive: contains "datagrid"

### API Usage

- `modelsdk.Open()` - Opens MPR file
- `reader.ListModules()` - Gets module list for name mapping
- `reader.ListPages()` - Gets all pages
- `reader.GetPage(id)` - Gets full page details with widgets
- `reader.ListSnippets()` - Gets all snippets with widgets

### Search Algorithm

1. Marshal page/snippet struct to JSON
2. Unmarshal to map[string]interface{} for inspection
3. Recursively search for widgets matching DataGrid type
4. Collect results and display formatted report

## Related Examples

- `find_widgets` - Generic widget finder (accepts widget type as argument)
- `find_custom_widgets` - Find custom widgets by widget ID
- `inspect_page` - Inspect detailed page structure
- `show_custom_widgets` - Show widget details for specific page
