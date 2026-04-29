---
name: export-page-manifest
description: Generate detailed page layout reports from Mendix MPR files. Extract Main/Right placeholders, tabs, widgets, buttons, DataGrid 2 columns and action buttons, button-to-page navigation flows. Export JSON page catalogs with module filtering. Use for page structure documentation, UI catalog generation, button navigation mapping, DataGrid toolbar analysis, modal panel content extraction, automation pipelines, page type filtering (engineering vs runtime).
userInvocable: true
---

# Mendix Page Manifest Generator

Generate comprehensive page layout reports from Mendix MPR files, including placeholder content, tab structures, widget captions, DataGrid 2 configurations, and button navigation flows.

## When to Use This Tool

Use `export_page_manifest.exe` when you need to:

- **Export page catalog as JSON**: Get a machine-readable list of all pages with module and layout info
- **Document page layouts**: Extract Main/Right placeholder structures with tabs and widgets
- **Analyze button flows**: Map button → nanoflow → page navigation chains
- **Catalog DataGrid 2 widgets**: Extract columns and toolbar action buttons
- **Extract modal panel content**: Document EXFN_ModalPanel layout with input fields and buttons
- **Filter pages by type**: Generate reports for engineering pages vs runtime/user-facing pages
- **Audit UI consistency**: Review widget captions and button labels across pages
- **Generate UI specifications**: Create technical documentation for page designs
- **Automate documentation pipelines**: Integrate page metadata into CI/CD or documentation generators

## Tool Location

The compiled executable is located at:
```
c:\Workspaces\Mendix\SDK\ModelSDKGo-acarum\examples\export_page_manifest\export_page_manifest.exe
```

Navigate to this directory before running commands, or use the full path.

## Usage Modes

### Mode 1: All Pages Report

Generate a report for all pages in the MPR (excluding marketplace modules):

```powershell
.\export_page_manifest.exe <mpr_file_path>
```

**Output:** `buttons_manifest.md` (contains all pages with Main/Right placeholders)

### Mode 2: Single Page Report

Generate a report for one specific page:

```powershell
.\export_page_manifest.exe <mpr_file_path> <page_name>
```

**Output:** `<PageName>_manifest.md`

### Mode 3: Filtered by Page Type

Generate reports only for pages matching a type filter:

```powershell
.\export_page_manifest.exe --type <eng|runtime> <mpr_file_path>
```

**Output:** `buttons_manifest.md` (filtered pages only)

### Mode 4: JSON Page List Export

Export a JSON list of all pages with name, module, and layout (excludes UI and marketplace modules):

```powershell
.\export_page_manifest.exe --pages <mpr_file_path>
```

**Output:** `page_list.json` (JSON array with page metadata)

**JSON Structure Example:**

```json
[
  {
    "name": "AuditTrailRecord_View",
    "module": "EXFN_AuditTrailViewer",
    "layout": "Atlas_Default"
  },
  {
    "name": "PANEL_UpdateUoMDimension",
    "module": "OpcenterEXFN_ReferenceData",
    "layout": "OpcenterEXFN_DISW_DesignSystem.EXFN_ModalPanel"
  },
  {
    "name": "PANEL_CreateBaseUoM_UoMDimension",
    "module": "OpcenterEXFN_ReferenceData",
    "layout": "OpcenterEXFN_DISW_DesignSystem.EXFN_ModalPanel"
  }
]
```

**Filtering Applied:**
- ✅ **Includes**: Business/application modules
- ❌ **Excludes**: UI modules (containing DISW, DesignSystem, _UI, UI_)
- ❌ **Excludes**: Marketplace modules (Atlas, Administration, System, etc.)

## Command-Line Options

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `--type` | `eng`, `runtime` | _(none)_ | Filter pages by type: engineering pages or runtime/user-facing pages |
| `--pages` | _(flag)_ | `false` | Export JSON page list (name, module, layout) instead of Markdown report |
| `--filter` | Wildcard pattern | _(none)_ | Filter page names with wildcard (e.g., `PANEL_*`) |
| `<mpr_file_path>` | File path | _(required)_ | Path to Mendix MPR file (absolute or relative, must have .mpr extension) |
| `[page_name]` | Page name | _(none)_ | Optional: analyze only this page (exact name match) |

## Example Commands

### Basic Page Analysis

```powershell
# All pages in project
.\export_page_manifest.exe "C:\Projects\MyApp.mpr"

# Single page report
.\export_page_manifest.exe "C:\Projects\MyApp.mpr" "PANEL_ChangePackage"

# Relative path (MPR in current directory)
.\export_page_manifest.exe "MyApp.mpr"
```

### Page Type Filtering

```powershell
# Only engineering pages
.\export_page_manifest.exe --type eng "MyApp.mpr"

# Only runtime/user-facing pages
.\export_page_manifest.exe --type runtime "MyApp.mpr"

# Filter by type and single page
.\export_page_manifest.exe --type eng "MyApp.mpr" "StateMachine_Details"
```

### JSON Export Examples

```powershell
# Export page list to JSON
.\export_page_manifest.exe --pages "C:\Projects\MyApp.mpr"

# Filter pages with wildcard pattern
.\export_page_manifest.exe --pages --filter "PANEL_*" "MyApp.mpr"

# Combine filters (type + pages)
.\export_page_manifest.exe --pages --type eng "MyApp.mpr"

# Process JSON output
$pages = Get-Content "page_list.json" | ConvertFrom-Json
Write-Host "Total pages: $($pages.Count)"
$pages | Group-Object module | Format-Table Name, Count
```

### Automation Examples

```powershell
# Generate reports for multiple projects
$projects = Get-ChildItem "C:\Projects\*.mpr"
foreach ($mpr in $projects) {
    .\export_page_manifest.exe $mpr.FullName
}

# Generate report and open in editor
.\export_page_manifest.exe "MyApp.mpr" "PANEL_ChangePackage"
code "PANEL_ChangePackage_manifest.md"

# Export JSON page list and analyze
.\export_page_manifest.exe --pages "MyApp.mpr"
$json = Get-Content "page_list.json" -Raw | ConvertFrom-Json
$json | Where-Object { $_.layout -like "*ModalPanel*" } | Select-Object name, module
```

## Output Structure

### Report Sections

Each page report contains:

1. **Page Header**: Page name and title (with en_US translation)
2. **Main Placeholder**:
   - Layout template path
   - Contents table (for modal panels): all input widgets with captions
   - Tabs (if present): nested widgets and DataGrid 2 details per tab
   - Buttons: toolbar buttons with captions
   - Button → Page Navigation: mapping of buttons to opened pages
   - Panel Details: widgets on each opened page
3. **Right Placeholder**: Same structure as Main (for command bars)
4. **Summary**: Total button count

### Example Output: Modal Panel Page

```markdown
# PANEL_ChangePackage

---

## 📄 PANEL_ChangePackage

**Title:** Change Package

### Main

**Layout:** `OpcenterEXFN_DISW_DesignSystem.EXFN_ModalPanel.Main`

#### Contents

| # | Type | Name | Caption |
|---|---|---|---|
| 1 | TextArea | operationTextArea |  |

**Buttons:**

| # | Name | Caption |
|---|---|---|
| 1 | CPSaveActionButton | Save |
| 2 | notesCancelActionButton | Cancel |

---

## Summary

- **Total Buttons:** 2
```

### Example Output: Tabbed Page

```markdown
## 📄 StateMachine_Details

**Title:** State Machine Details

### Main

**Layout:** `Atlas_Default.Main`

#### Tab: General

**Caption:** General Information

| # | Type | Name | Caption |
|---|---|---|---|
| 1 | TextBox | nameTextBox | Name |
| 2 | TextArea | descriptionTextArea | Description |

**DataGrid 2: statesDataGrid**

**Columns:**

| # | Column ID | Caption |
|---|---|---|
| 1 | stateName | State Name |
| 2 | stateType | Type |
| 3 | isInitial | Initial State |

**Action Buttons:**

| # | Name | Caption | Nanoflow | Opens Page |
|---|---|---|---|---|
| 1 | addStateButton | Add State | ACT_AddState | PANEL_AddState |
| 2 | editStateButton | Edit | ACT_EditState | PANEL_EditState |
| 3 | deleteStateButton | Delete | ACT_DeleteState | - |
```

### Example Output: Vertical Command Bar

```markdown
### Right

**Layout:** `Atlas_Default.Right`

**Vertical CommandBar**

| | Name |
|---|---|
| Container | `rightContainer` |
| VerticalCommandBarClass | `verticalCommandBarContainer` |

**Vertical CommandBar Buttons:**

| # | Name | Caption |
|---|---|---|
| 1 | helpButton | Help |
| 2 | settingsButton | Settings |

**Button → Page Navigation:**

| Button Caption | Nanoflow | Page Opened |
|---|---|---|
| Settings | ACT_OpenSettings | PANEL_Settings |

**Panel: `PANEL_Settings`**

| Widget Type | Name | Caption |
|---|---|---|
| CheckBox | enableNotifications | Enable Notifications |
| DropDown | themeSelector | Theme |
```

## Content Extraction Details

### Supported Widget Types

The tool extracts captions from:

- **Input widgets**: TextBox, TextArea, DatePicker, DropDown, CheckBox, RadioButton
- **Display widgets**: Label, StaticLabel, Image, DynamicImage
- **Action widgets**: ActionButton
- **Data widgets**: DataView, ListView, ReferenceSelector, InputReferenceSelector, FileManager
- **DataGrid 2**: Custom widget with columns and action buttons

### DataGrid 2 Analysis

For DataGrid 2 widgets (`com.mendix.widget.web.datagrid.Datagrid`), the tool extracts:

- **Columns**: ID and caption (from `caption` or `header` properties)
- **Action Buttons**: Toolbar buttons only (excludes inline row buttons with RenderType "Link")
- **Button Configuration**: Name, caption, parent container, nanoflow, opened page

### Caption Extraction Logic

Captions are extracted using these patterns (in order of priority):

1. **Direct Caption** (string or Texts$Text with translations)
2. **CaptionTemplate.Template** (ActionButton, en_US only)
3. **LabelTemplate.Template** (input widgets)
4. **Tooltip** (fallback for icon-only buttons)

**Language Priority**: English (en_US) preferred, falls back to first available language.

### Button Navigation Resolution

For each action button with a nanoflow:

1. Tool opens the nanoflow document from the MPR
2. Searches for `Microflows$ShowFormAction` inside the nanoflow
3. Extracts the `FormSettings.Form` property (page reference)
4. Displays full navigation chain: **Button → Nanoflow → Page**
5. Loads and documents widgets on the opened page

### Layout Placeholder Detection

The tool identifies:

- **Main placeholder**: `Atlas_Default.Main`, `EXFN_ModalPanel.Main`, etc.
- **Right placeholder**: `Atlas_Default.Right`, custom layouts
- **Root container**: First DivContainer in placeholder
- **Vertical command bar**: DivContainer with CSS class containing "vertical-command-bar"

### Modal Panel Special Handling

Pages with `EXFN_ModalPanel` layout show:

- **Contents section**: Table of all input widgets (excluding ActionButton and DataView)
- **Buttons section**: Modal action buttons (Save, Cancel, etc.)
- **Simplified layout**: No Vertical CommandBar header (not applicable for modals)

## Typical Output Examples

### Progress Logging

```
Opened: C:\Projects\MyApp.mpr
MPR Version: 2

Scanning 45 pages...

[45/45] Scanning: PANEL_OperatorLandingStart

Pages with Main/Right placeholders: 12

Exported buttons manifest to: buttons_manifest.md
```

### Single Page Mode

```
Opened: C:\Projects\MyApp.mpr
MPR Version: 2

Scanning 45 pages...

[23/45] Scanning: PANEL_ChangePackage
[DEBUG] PANEL_ChangePackage → Main=true Right=false

Pages with Main/Right placeholders: 1

Exported page manifest to: PANEL_ChangePackage_manifest.md
```

### JSON Export Mode

```
Opened: C:\Projects\MyApp.mpr
MPR Version: 2

Scanning 51 pages...

Filtering modules:
  ✗ Excluded UI module: OpcenterEXFN_DISW_DesignSystem (2 pages)
  ✗ Excluded marketplace module: Atlas_Core (3 pages)
  ✓ Included business module: OpcenterEXFN_ReferenceData (49 pages)
  ✓ Included business module: EXFN_AuditTrailViewer (2 pages)

Exported page list to: page_list.json
Total pages exported: 51
```

**Output file (page_list.json):**
```json
[
  {
    "name": "AuditTrailRecord_View",
    "module": "EXFN_AuditTrailViewer",
    "layout": "Atlas_Default"
  },
  {
    "name": "PANEL_UpdateUoMDimension",
    "module": "OpcenterEXFN_ReferenceData",
    "layout": "OpcenterEXFN_DISW_DesignSystem.EXFN_ModalPanel"
  }
]
```

## Marketplace Module Filtering

The tool automatically excludes pages from marketplace/system modules:

- Marketplace, Community, AppStore
- Atlas, Administration, System
- CommunityCommons, NanoflowCommons
- DataWidgets, WebActions
- DISW_DesignSystem

**Rationale**: Focus on custom application pages, not framework/library pages.

## Troubleshooting

### Error: "ERROR: MPR file path is required"

**Cause**: No MPR file path provided as command-line argument  
**Solution**: Provide MPR file path: `.\export_page_manifest.exe "MyApp.mpr"`

### Error: "ERROR: File must have .mpr extension"

**Cause**: Provided file does not have .mpr extension  
**Solution**: Ensure you're pointing to a valid Mendix project file (*.mpr)

### Error: "ERROR: MPR file does not exist"

**Cause**: Specified MPR file path does not exist on disk  
**Solution**: Verify the file path is correct and the file exists

### Error: "mprcontents folder not found"

**Cause**: Invalid MPR path or corrupt MPR file  
**Solution**: Verify MPR file exists and is a valid Mendix Version 2 format (Mendix 7.0+)

### Error: "BSON unmarshal failed"

**Cause**: Corrupted page document or unsupported Mendix version  
**Solution**: Try opening the MPR in Mendix Studio Pro to verify project integrity

### Warning: Page not found in report

**Possible causes:**
1. Page has no Main/Right placeholders (e.g., popup or custom layout)
2. Page is in a marketplace module (automatically excluded)
3. Page type doesn't match `--type` filter

**Solution**: Run without filters to see all qualifying pages, or check page layout in Mendix Studio Pro

### Empty Contents Section

**Cause**: Page layout is not `EXFN_ModalPanel` (Contents section only shown for modal panels)  
**Expected behavior**: Standard pages show Tabs/Buttons sections instead

### Missing Button Captions

**Possible causes:**
1. Icon-only button (caption is empty, tooltip used as fallback)
2. Caption in non-English language (tool prefers en_US, falls back to first language)

**Solution**: Check button configuration in Mendix Studio Pro

### DataGrid 2 Not Detected

**Cause**: Widget ID mismatch (tool searches for `com.mendix.widget.web.datagrid.Datagrid`)  
**Solution**: Verify widget type in page JSON or Mendix Studio Pro

### Button Navigation Shows Empty Page

**Cause**: Nanoflow doesn't contain `ShowFormAction`, or page reference is not resolvable  
**Solution**: Inspect nanoflow in Mendix Studio Pro to verify ShowPage action exists

## Advanced Use Cases

### Page Type Filtering Implementation

Currently, the `--type` flag is implemented with placeholder logic:

- `--type eng`: Returns all pages (TODO: add engineering page detection)
- `--type runtime`: Returns all pages (TODO: add runtime page detection)

**Customization**: To implement custom filtering logic, modify the `matchesPageType` function in the source code with your page naming conventions or metadata checks.

### Extracting Button Navigation for Flow Diagrams

```powershell
# Generate report and extract navigation data
.\export_page_manifest.exe "MyApp.mpr"

# Parse Markdown to extract Button → Page Navigation sections
$content = Get-Content "buttons_manifest.md"
$content | Select-String "Button → Page Navigation:" -Context 0,5
```

### Extracting DataGrid Configurations

To extract all DataGrid 2 configurations from a project:

```powershell
# Generate full report
.\export_page_manifest.exe "MyApp.mpr"

# Search for DataGrid 2 sections in output
Select-String "DataGrid 2:" "buttons_manifest.md" -Context 5,15
```

## Integration Examples

### PowerShell Automation: Generate Reports for All Projects

```powershell
$projects = @(
    "C:\Projects\App1.mpr",
    "C:\Projects\App2.mpr",
    "C:\Projects\App3.mpr"
)

foreach ($mpr in $projects) {
    $appName = [System.IO.Path]::GetFileNameWithoutExtension($mpr)
    Write-Host "Processing: $appName"
    
    .\export_page_manifest.exe $mpr
    
    # Rename output with project name
    if (Test-Path "buttons_manifest.md") {
        Move-Item "buttons_manifest.md" "${appName}_buttons_manifest.md" -Force
    }
}

Write-Host "Reports generated: $(Get-ChildItem '*_buttons_manifest.md').Count"
```

### Extract Button Navigation to CSV

```powershell
# Generate report
.\export_page_manifest.exe "MyApp.mpr"

# Parse navigation sections and convert to CSV
$content = Get-Content "buttons_manifest.md" -Raw
$navSections = [regex]::Matches($content, '(?s)\*\*Button → Page Navigation:\*\*\s+\|.*?\n\|.*?\n\|(.*?)(?=\n\n)')

$results = @()
foreach ($match in $navSections) {
    $rows = $match.Groups[1].Value -split '\n'
    foreach ($row in $rows) {
        if ($row -match '\|\s*(.+?)\s*\|\s*(.+?)\s*\|\s*(.+?)\s*\|') {
            $results += [PSCustomObject]@{
                ButtonCaption = $matches[1].Trim()
                Nanoflow = $matches[2].Trim()
                PageOpened = $matches[3].Trim()
            }
        }
    }
}

$results | Export-Csv -Path "button_navigation.csv" -NoTypeInformation
Write-Host "Exported $($results.Count) navigation entries to button_navigation.csv"
```

## FAQ

**Q: Can I filter pages by module?**  
A: Not directly via command-line flag. The tool automatically excludes marketplace modules. For custom module filtering, use the full report and filter the output Markdown by page name prefix.

**Q: Does the tool modify the MPR file?**  
A: No, it's read-only. The tool only reads from the SQLite database and BSON files.

**Q: What Mendix versions are supported?**  
A: MPR Version 2 format (Mendix 7.0+). Tested with Mendix 10.x and 11.x.

**Q: Can I export to JSON instead of Markdown?**  
A: Yes! Use the `--pages` flag to export a JSON list of all pages with metadata (name, module, layout). Example: `.\export_page_manifest.exe --pages "MyApp.mpr"` → `page_list.json`. For full BSON structure of a single page, use `export_page_json.exe`.

**Q: How do I find the exact page name?**  
A: Use `find_pages.exe` to list all pages in the MPR, then copy the exact name to use as filter.

**Q: Why doesn't my DataGrid 2 show columns?**  
A: The tool uses type-based property extraction. If columns aren't detected, the widget may use custom property names. Check the widget configuration in Mendix Studio Pro.

**Q: Can I customize the Markdown format?**  
A: Not via command-line. Customization requires modifying the source code (exportMarkdown function).

**Q: How do I get button → page navigation for nested dialogs?**  
A: The tool resolves one level of navigation (button → nanoflow → page). For deeper chains, generate reports for each page in the chain separately.

## Related Tools

- **export_page_json**: Export full page BSON structure to JSON (for debugging)
- **find_pages**: List all pages in an MPR with module and ID
- **inspect_page**: Detailed page structure inspection with widget tree
- **find_custom_widgets**: Find pages using specific custom widget types
- **find_datagrid**: Find all pages with DataGrid 2 widgets

## Best Practices

1. **Start with JSON export** (`--pages`) to get a full page catalog, then generate detailed reports for specific pages
2. **Use JSON mode for automation** to integrate page lists into build pipelines or documentation tools
3. **Combine filters** (`--pages --filter "PANEL_*"`) to focus on specific page naming patterns
4. **Use `--type` filter** (once implemented) to separate engineering UI from runtime UI
5. **Store reports in version control** to track page layout changes over time
6. **Generate before/after reports** when refactoring page layouts
7. **Use single-page mode** for detailed analysis during development
8. **Parse navigation sections** to generate flow diagrams or test plans

## Output File Management

- **Default output location**: Current directory (where the command is executed)
- **Single page report**: `<PageName>_manifest.md`
- **All pages report**: `buttons_manifest.md`
- **JSON page list**: `page_list.json`
- **Overwrite behavior**: Files are overwritten without confirmation

**Output Format by Mode:**

| Mode | Flag | Output File | Format |
|------|------|-------------|--------|
| All pages | _(none)_ | `buttons_manifest.md` | Markdown |
| Single page | `[page_name]` | `<PageName>_manifest.md` | Markdown |
| Page list | `--pages` | `page_list.json` | JSON |
| Filtered | `--type eng` | `buttons_manifest.md` | Markdown |

**Tip**: Navigate to the output directory or move generated files to a central documentation folder for long-term storage.
