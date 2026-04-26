---
name: mendix-report-generator
description: 'Generate comprehensive manifest reports from Mendix MPR files. Use for: analyzing Mendix projects, exporting project documentation, finding Signal Manager widgets, batch processing multiple MPR files, customizing report sections with flags, troubleshooting MPR analysis issues.'
argument-hint: 'Specify MPR file path or directory for batch mode'
---

# Mendix Manifest Report Generator

A Go-based tool that analyzes Mendix MPR files (v2 format with SQLite + BSON) and generates detailed markdown reports containing project structure, microflows, widgets, navigation, and security information.

## When to Use

**Generate reports when you need to:**
- Document Mendix project architecture and dependencies
- Find and catalog Signal Manager widget subscriptions
- Analyze microflow/action calls across the project
- Extract navigation structure and page accessibility rules
- Compare multiple Mendix projects side-by-side
- Audit external entities and attributes
- Create batch reports for entire project directories

**Troubleshoot with this skill when:**
- Reports take too long to generate
- Missing expected content in reports
- Batch mode not finding MPR files
- Need to customize report sections

## Tool Location

```
c:\Workspaces\Mendix\SDK\ModelSDKGo-acarum\examples\export_manifest\
```

Executable: `export_manifest.exe`
Source: `main.go`

## Operation Modes

### Single File Mode

Analyze one MPR file and generate a report.

**Basic usage:**
```powershell
.\export_manifest.exe "path\to\project.mpr"
```

**Auto-detect MPR in current directory:**
```powershell
.\export_manifest.exe .
```

**Custom output directory:**
```powershell
.\export_manifest.exe -output-dir="reports" "project.mpr"
```

**Output:** `<project-name>-manifest.md` in the same directory as the MPR or in `-output-dir`

### Batch Mode

Recursively scan a directory tree for all MPR files and generate reports for each.

**Basic batch:**
```powershell
.\export_manifest.exe -source-dir="C:\Projects\Mendix\"
```

**With custom output:**
```powershell
.\export_manifest.exe -source-dir="." -output-dir="reports"
```

**Output:** 
- Individual reports: `<output-dir>/<project-name>-manifest.md`
- Index file: `<output-dir>/index.md` with wikilinks `[[filename]]` to all reports

**Excluded directories:** `.mendix-cache`, `deployment`, `bin`, `.svn`, `.git`, `node_modules`, `.idea`, `.vscode`

## Report Customization Flags

All flags work in both single file and batch modes.

| Flag | Default | Description |
|------|---------|-------------|
| `-include-entities` | `false` | Include external entities in report |
| `-include-attributes` | `true` | Include entity attributes (requires entities enabled) |
| `-include-microflows` | `true` | Include microflow/action calls analysis |
| `-include-widgets` | `true` | Include Signal Manager widget subscriptions |
| `-include-navigation` | `true` | Include navigation hierarchy |
| `-include-roles` | `false` | Include system roles and page accessibility |
| `-output-dir` | `""` | Directory for output files |
| `-source-dir` | `""` | Source directory for batch mode |

### Flag Examples

**Widgets-only report:**
```powershell
.\export_manifest.exe -include-entities=false `
  -include-microflows=false `
  -include-navigation=false `
  "project.mpr"
```

**Full security audit:**
```powershell
.\export_manifest.exe -include-roles=true `
  -include-entities=true `
  "project.mpr"
```

**Minimal navigation report:**
```powershell
.\export_manifest.exe -include-microflows=false `
  -include-widgets=false `
  "project.mpr"
```

## Report Sections

Generated markdown reports contain:

### 1. Project Overview
- Project name
- Mendix version
- Generation timestamp

### 2. Microflow/Action Calls (optional, default: ON)
- Table of microflows with their called actions/sub-microflows
- Module and document names
- Call hierarchy

### 3. Signal Manager Subscriptions (optional, default: ON)
- Siemens MxToSignal widget instances
- Signal names and app names
- Subscription filter presence (Yes/No)
- Document type (Page/Snippet) and location

**Performance:** Pre-filtering optimization scans only documents containing widgets (95%+ reduction in processing time)

### 4. System Roles & Page Access (optional, default: OFF)
- User roles defined in the project
- Page/snippet accessibility rules per role
- Security matrix

### 5. Navigation Items (optional, default: ON)
- Navigation hierarchy structure
- Menu items and their targets
- Page references

### 6. External Entities (optional, default: OFF)
- Entities from external modules
- Associations and attributes

## Performance Characteristics

### Typical Execution Times

| Project Size | Documents | Microflows | Time (Optimized) |
|--------------|-----------|------------|------------------|
| Small (OC EX System) | 118 | 124 | ~8-10s |
| Medium (RPT) | 214 | 396 | ~30-35s |
| Large (CMX) | 220 | 412 | ~40-45s |
| Extra Large (MFT) | 332 | 635 | ~60-70s |

### Optimization Features

**Widget Pre-filtering:**
- Fast byte-search (`bytes.Contains`) before BSON unmarshaling
- Reduces document processing by 94-98%
- Only loads documents that actually contain target widgets

**Progress Logging:**
- `[Phase X/Y]` indicators show current analysis stage
- Document count and filter statistics
- Real-time progress updates

## Troubleshooting

### Issue: Batch mode finds fewer files than expected

**Check excluded directories:**
- By default, `bin`, `deployment`, `.mendix-cache` are skipped
- Verify MPR files aren't in excluded paths
- Use single file mode for specific MPRs in excluded directories

**Verification command:**
```powershell
Get-ChildItem -Path "C:\Projects\" -Filter "*.mpr" -Recurse | 
  Where-Object { $_.FullName -notmatch "bin|deployment|\.mendix-cache" }
```

### Issue: Report generation is slow

**Disable unused sections:**
```powershell
.\export_manifest.exe -include-entities=false `
  -include-roles=false `
  "large-project.mpr"
```

**Check for large microflow counts:**
- Microflow analysis is currently the main bottleneck
- Projects with 500+ microflows may take 60s+

**Monitor with progress logging:**
- Watch for phase completion times
- Identify which phase is slowest

### Issue: Missing Signal Manager widgets in report

**Verify widget ID:**
- Tool searches for `siemens.mxtosignal.MxToSignal`
- Custom widget IDs require code modification

**Check pre-filtering output:**
- Look for "Found X documents to scan (filtered from Y total)"
- If X is 0, no documents contain the widget ID string

**Disable widget section to test:**
```powershell
.\export_manifest.exe -include-widgets=false "project.mpr"
```

### Issue: Cannot find MPR file with "." shortcut

**Ensure only one MPR in directory:**
- "." auto-detect works only with single MPR files
- Multiple MPRs require explicit path

**Use full path instead:**
```powershell
.\export_manifest.exe "$(Get-Location)\ProjectName.mpr"
```

### Issue: Report contains "No" but should show subscription filters

**Working as designed:**
- Subscription Filter column shows Yes/No (not actual filter values)
- "Yes" = filter is configured
- "No" = filter is empty/null
- This is a recent UI improvement for readability

## Workflow: Generate Batch Reports for Project Portfolio

1. **Navigate to tool directory:**
   ```powershell
   cd c:\Workspaces\Mendix\SDK\ModelSDKGo-acarum\examples\export_manifest\
   ```

2. **Verify tool is compiled:**
   ```powershell
   if (-not (Test-Path .\export_manifest.exe)) { go build -o export_manifest.exe }
   ```

3. **Run batch analysis:**
   ```powershell
   .\export_manifest.exe -source-dir="C:\Projects\Mendix\" `
     -output-dir="C:\Projects\Mendix\reports"
   ```

4. **Review index file:**
   - Open `C:\Projects\Mendix\reports\index.md`
   - Contains wikilinks to all generated reports: `[[filename]]`
   - Shows generation timestamp and project count

5. **Optional: Regenerate with different flags:**
   ```powershell
   # Security-focused reports
   .\export_manifest.exe -source-dir="C:\Projects\Mendix\" `
     -output-dir="C:\Projects\Mendix\security-reports" `
     -include-roles=true
   ```

## Workflow: Debug Missing Widgets

1. **Run with widgets-only mode:**
   ```powershell
   .\export_manifest.exe -include-entities=false `
     -include-microflows=false `
     -include-navigation=false `
     -include-roles=false `
     "project.mpr"
   ```

2. **Check pre-filtering statistics:**
   - Look for: "📊 Found X documents to scan (filtered from Y total, Z% reduction)"
   - If X = 0: Widget ID not found in any documents
   - If X > 0 but subscriptions = 0: Widget exists but has no subscriptions

3. **Verify widget configuration in Mendix:**
   - Open project in Mendix Studio Pro
   - Search for widget type: "MxToSignal"
   - Check Signal Name and Subscription Filter properties

4. **Compare with known-good project:**
   - Run tool on project confirmed to have widgets
   - Compare document counts and widget findings

## Workflow: Customize Report for Specific Audience

**For architects (high-level overview):**
```powershell
.\export_manifest.exe -include-attributes=false `
  -include-microflows=false `
  -include-widgets=false `
  "project.mpr"
```
Output: Navigation structure and security roles only

**For developers (technical details):**
```powershell
.\export_manifest.exe -include-entities=true `
  -include-attributes=true `
  -include-roles=true `
  "project.mpr"
```
Output: Full entity model, microflows, and access control

**For integration teams (Signal Manager focus):**
```powershell
.\export_manifest.exe -include-entities=false `
  -include-microflows=false `
  -include-navigation=false `
  "project.mpr"
```
Output: Only Signal Manager subscriptions

## Technical Details

### MPR File Format (Mendix v2)

- **Database:** SQLite3 file (`.mpr`)
- **Documents:** BSON files in `mprcontents/{xx}/{yy}/{uuid}.mxunit`
- **UUID encoding:** Windows GUID format (mixed-endian)

### Key Document Types

- `Navigation$NavigationDocument` - Navigation structure
- `Security$ProjectSecurity` - System-wide security
- `Security$ModuleSecurity` - Module-level security
- `Forms$Page` - UI pages
- `Forms$Snippet` - Reusable UI components
- `DomainModels$EntityImpl` - Data entities
- `Microflows$Microflow` - Business logic flows
- `CustomWidgets$CustomWidget` - Widget configurations

### Dependencies

- `github.com/anthropics/modelsdk-go` - MPR SDK
- `github.com/mattn/go-sqlite3` - SQLite3 driver
- `go.mongodb.org/mongo-driver/bson` - BSON parsing

## Code Modifications

### Change Widget ID

Edit `main.go` line ~1595:
```go
widgetID := "siemens.mxtosignal.MxToSignal"  // Change to your widget ID
```

### Add Custom Report Section

1. Create collection function (follow pattern of `collectSignalManagerWidgets`)
2. Add data structure to `ManifestReport` struct (~line 90)
3. Add flag to CLI flags section (~line 246)
4. Call collection function in `processSingleFile` (~line 500)
5. Add markdown generation in `generateMarkdownReport` (~line 3200)

### Modify Exclusion List

Edit `findMPRFilesRecursive` function (~line 130):
```go
excludeDirs := []string{
    ".mendix-cache",
    "deployment",
    "bin",
    // Add your directories here
}
```

## Best Practices

1. **Always use absolute paths** for MPR files on Windows
2. **Run batch mode from tool directory** to avoid path issues
3. **Test flags on small project first** before large batch operations
4. **Keep reports in version control** for change tracking
5. **Use output-dir to separate report generations** by date/purpose
6. **Review index.md first** in batch mode to see all projects found

## Related Tools

- **Mendix SDK:** Full programmatic access to project model
- **Mendix Studio Pro:** Official IDE for manual inspection
- **Mendix mx-assist:** AI-powered analysis within Studio Pro

## Example Outputs

**Index.md structure:**
```markdown
# Manifest Reports Index

**Generated:** 2026-04-26 14:30:00
**Total Reports:** 3

## Reports

- [[Opcenter EX DS Complex Manufacturing-manifest]]
- [[Opcenter EX DS Manufacturing-manifest]]
- [[Opcenter EX DS Repetitive Manufacturing-manifest]]
```

**Signal Manager section:**
```markdown
## Signal Manager Subscriptions

| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |
|--------|---------------|----------|-------------|----------|---------------------|
| MyModule | Page | DetailPage | SignalName | AppName | Yes |
```
