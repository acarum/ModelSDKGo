package main

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/anthropics/modelsdk-go"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Data structures
type EntityInfo struct {
	Name           string   `json:"Name"`
	Module         string   `json:"Module"`
	Attributes     []string `json:"Attributes"`
	PublishedFrom  string   `json:"PublishedFrom"`
	EntityTypeName string   `json:"EntityTypeName"`
}

type DomainModelInfo struct {
	Name    string
	UnitID  string
	Content map[string]interface{}
}

type MicroflowCallInfo struct {
	MicroflowName string `json:"MicroflowName"`
	Module        string `json:"Module"`
	ActivityName  string `json:"ActivityName"`
	Caption       string `json:"Caption"`
	CallType      string `json:"CallType"`   // "MicroflowCall", "JavaAction", or "ExternalAction"
	TargetName    string `json:"TargetName"` // Microflow/JavaAction/ExternalAction name
	AppName       string `json:"AppName"`
	CommandName   string `json:"CommandName"`
}

type MicroflowInfo struct {
	ID         string
	Name       string
	ModuleName string
	Content    map[string]interface{}
}

type WidgetInfo struct {
	DocumentName       string `json:"DocumentName"`
	DocumentType       string `json:"DocumentType"` // "Page" or "Snippet"
	Module             string `json:"Module"`
	SignalName         string `json:"SignalName"`
	AppName            string `json:"AppName"`
	SubscriptionFilter string `json:"SubscriptionFilter"`
}

type NavigationItem struct {
	ItemName     string   `json:"ItemName"`
	Caption      string   `json:"Caption"`
	Target       string   `json:"-"`          // Page or microflow name - excluded from JSON export
	TargetPage   string   `json:"TargetPage"` // Resolved page qualified name(s) - qualified name (Module.PageName) for Page items, or page(s) opened by Microflow/Nanoflow
	Module       string   `json:"Module"`
	MenuDocument string   `json:"-"`            // Name of the menu document - excluded from JSON export
	ItemType     string   `json:"-"`            // "Page", "Microflow", "Nanoflow" - excluded from JSON
	ParentItem   string   `json:"ParentItem"`   // For hierarchical structure
	Level        int      `json:"Level"`        // Indentation level
	AllowedRoles []string `json:"AllowedRoles"` // User roles that can access this item
}

type SystemRole struct {
	Name   string `json:"Name"`
	Module string `json:"Module"`
}

type PageAccessInfo struct {
	PageName     string   `json:"PageName"`
	Module       string   `json:"Module"`
	DocumentType string   `json:"DocumentType"` // "Page" or "Snippet"
	AllowedRoles []string `json:"AllowedRoles"`
	IsPublic     bool     `json:"IsPublic"` // No role restrictions
}

type PageCommandButton struct {
	ButtonName    string `json:"ButtonName"`
	Caption       string `json:"Caption"`
	ActionType    string `json:"-"`                 // Type of action (e.g., "CallNanoflowClientAction", "CallMicroflowClientAction")
	ActionName    string `json:"-"`                 // Nanoflow/Microflow name
	TargetPage    string `json:"TargetPage"`        // Resolved page from ShowPage action (if any)
	TargetCommand string `json:"TargetCommandName"` // Command extracted from Save button in TargetPage
	TargetAppName string `json:"TargetAppName"`     // App name extracted from TargetCommand
}

type PageCommandInfo struct {
	PageName string              `json:"PageName"` // Qualified page name (Module.PageName)
	Commands []PageCommandButton `json:"Commands"`
}

// PageCommandSummary stores simplified page command information
type PageCommandSummary struct {
	PageName       string   `json:"PageName"`
	Module         string   `json:"Module"`
	TargetAppNames []string `json:"TargetAppNames"`
	TargetCommands []string `json:"TargetCommands"`
}

type ManifestReport struct {
	ProjectName        string                  `json:"ProjectName"`
	MendixVersion      string                  `json:"MendixVersion"`
	MPRPath            string                  `json:"MPRPath"`
	GeneratedAt        string                  `json:"GeneratedAt"`
	Entities           map[string][]EntityInfo `json:"Entities"`                   // by module
	MicroflowCalls     []MicroflowCallInfo     `json:"MicroflowCalls"`             // all calls
	Widgets            []WidgetInfo            `json:"SignalManagerSubscriptions"` // signal manager widgets
	NavigationItems    []NavigationItem        `json:"NavigationItems"`            // navigation menu items
	SystemRoles        []SystemRole            `json:"SystemRoles"`                // system roles
	PageAccess         []PageAccessInfo        `json:"-"`                          // page accessibility - excluded from export
	NavigationCommands []PageCommandInfo       `json:"NavigationPageCommands"`     // command bar actions from navigation pages
	PagesAnalysis      []PageAnalysisInfo      `json:"PageCommandsHierarchy"`      // detailed pages/panels analysis with recursive hierarchy
	PageCommands       []PageCommandSummary    `json:"PageCommands"`               // simplified view of all page/panel commands
}

type ReportOptions struct {
	IncludeEntities               bool
	IncludeAttributes             bool
	IncludeMicroflows             bool
	IncludeWidgets                bool
	IncludeNavigation             bool
	IncludeRoles                  bool
	IncludePageCommands           bool
	IncludePagesCommandsHierarchy bool
}

// MicroflowCallHierarchy represents a recursive call tree
type MicroflowCallHierarchy struct {
	Name  string
	Level int `json:"-"`
	Calls []MicroflowCallHierarchy
}

// PageAnalysisInfo stores detailed analysis of pages/panels
type PageAnalysisInfo struct {
	Name           string
	Module         string
	MicroflowCalls []string
	CallHierarchy  []MicroflowCallHierarchy
	TargetCommands []string
}

// ReportEntry represents a report generation result
type ReportEntry struct {
	ProjectName string `json:"ProjectName"`
	MPRPath     string `json:"MPRPath"`
	ReportPath  string `json:"ReportPath"`
	Success     bool   `json:"Success"`
	Error       string `json:"Error"`
}

// findMPRFilesRecursive searches for all .mpr files in a directory tree
func findMPRFilesRecursive(rootDir string) ([]string, error) {
	var mprFiles []string

	// Directories to exclude from scanning
	excludeDirs := []string{
		".mendix-cache",
		"deployment",
		"bin",
		".svn",
		".git",
		"node_modules",
		".idea",
		".vscode",
	}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip excluded directories
		if info.IsDir() {
			dirName := info.Name()
			for _, exclude := range excludeDirs {
				if dirName == exclude {
					return filepath.SkipDir
				}
			}
		}

		// Add .mpr files to the list
		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".mpr" {
			mprFiles = append(mprFiles, path)
		}
		return nil
	})

	return mprFiles, err
}

// generateIndexFile creates an index.md file listing all generated reports
func generateIndexFile(indexPath string, entries []ReportEntry) error {
	file, err := os.Create(indexPath)
	if err != nil {
		return fmt.Errorf("failed to create index file: %w", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "# Manifest Reports Index\n\n")
	fmt.Fprintf(file, "**Generated:** %s  \n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(file, "**Total Reports:** %d  \n\n", len(entries))
	fmt.Fprintf(file, "---\n\n")

	// Successful reports
	successCount := 0
	for _, entry := range entries {
		if entry.Success {
			successCount++
		}
	}

	if successCount > 0 {
		fmt.Fprintf(file, "## Generated Reports (%d)\n\n", successCount)
		fmt.Fprintf(file, "| App Name | Report |\n")
		fmt.Fprintf(file, "|----------|--------|\n")

		for _, entry := range entries {
			if entry.Success {
				reportFilename := filepath.Base(entry.ReportPath)
				appName := entry.ProjectName
				if appName == "" {
					appName = strings.TrimSuffix(filepath.Base(entry.MPRPath), ".mpr")
				}
				// Use wikilink syntax [[filename]]
				fmt.Fprintf(file, "| %s | [[%s]] |\n", appName, reportFilename)
			}
		}
		fmt.Fprintf(file, "\n")
	}

	// Failed reports
	failedCount := len(entries) - successCount
	if failedCount > 0 {
		fmt.Fprintf(file, "## Failed Reports (%d)\n\n", failedCount)
		fmt.Fprintf(file, "| MPR File | Error |\n")
		fmt.Fprintf(file, "|----------|-------|\n")

		for _, entry := range entries {
			if !entry.Success {
				mprFilename := filepath.Base(entry.MPRPath)
				fmt.Fprintf(file, "| %s | %s |\n", mprFilename, entry.Error)
			}
		}
		fmt.Fprintf(file, "\n")
	}

	fmt.Fprintf(file, "---\n\n")
	fmt.Fprintf(file, "_Index generated by export_manifest tool_\n")

	return nil
}

// countSuccessful counts successful report generations
func countSuccessful(entries []ReportEntry) int {
	count := 0
	for _, entry := range entries {
		if entry.Success {
			count++
		}
	}
	return count
}

// countFailed counts failed report generations
func countFailed(entries []ReportEntry) int {
	count := 0
	for _, entry := range entries {
		if !entry.Success {
			count++
		}
	}
	return count
}

func main() {
	// Define CLI flags
	includeEntities := flag.Bool("include-entities", true, "Include external entities in the report")
	includeAttributes := flag.Bool("include-attributes", true, "Include entity attributes in the report")
	includeMicroflows := flag.Bool("include-microflows", true, "Include microflow/action calls in the report")
	includeWidgets := flag.Bool("include-widgets", true, "Include Signal Manager widgets in the report")
	includeNavigation := flag.Bool("include-navigation", true, "Include navigation items in the report")
	includeRoles := flag.Bool("include-roles", false, "Include system roles and page accessibility in the report")
	includePageCommands := flag.Bool("include-page-commands", true, "Include command bar actions from navigation pages in the report")
	includePagesCommandsHierarchy := flag.Bool("include-pages-commands-hierarchy", true, "Include detailed pages/panels analysis with recursive microflow hierarchy (up to 5 levels)")
	outputDir := flag.String("output-dir", "", "Output directory for the report file (optional)")
	sourceDir := flag.String("source-dir", "", "Source directory to scan for MPR files recursively (batch mode)")
	outputFormat := flag.String("output-format", "md", "Output format: 'md' (Markdown), 'json' (JSON), or 'both' (Markdown + JSON)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: export_manifest [options] <mpr_file_path>\n")
		fmt.Fprintf(os.Stderr, "   OR: export_manifest [options] -source-dir=<directory>\n\n")
		fmt.Fprintf(os.Stderr, "Single File Mode:\n")
		fmt.Fprintf(os.Stderr, "  <mpr_file_path>   Path to the Mendix MPR file (use '.' to auto-detect in current directory)\n")
		fmt.Fprintf(os.Stderr, "  Output: <mpr_filename>-manifest.md (or .json)\n\n")
		fmt.Fprintf(os.Stderr, "Batch Mode (-source-dir):\n")
		fmt.Fprintf(os.Stderr, "  Recursively scans directory for all .mpr files and generates reports\n")
		fmt.Fprintf(os.Stderr, "  Output: <source-dir>/reports/ or <output-dir>/ with index.md/index.json\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  export_manifest MyApp.mpr\n")
		fmt.Fprintf(os.Stderr, "  export_manifest .\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -output-format=json MyApp.mpr\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -output-format=both MyApp.mpr\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -source-dir=\"C:\\\\Projects\\\\Mendix\"\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -source-dir=. -output-dir=\"reports\" -output-format=json\n")
	}
	flag.Parse()

	// Validate output format
	validFormats := map[string]bool{"md": true, "json": true, "both": true}
	if !validFormats[*outputFormat] {
		fmt.Fprintf(os.Stderr, "❌ Invalid output format '%s'. Must be 'md', 'json', or 'both'\n", *outputFormat)
		flag.Usage()
		os.Exit(1)
	}

	// Create report options
	options := ReportOptions{
		IncludeEntities:               *includeEntities,
		IncludeAttributes:             *includeAttributes,
		IncludeMicroflows:             *includeMicroflows,
		IncludeWidgets:                *includeWidgets,
		IncludeNavigation:             *includeNavigation,
		IncludeRoles:                  *includeRoles,
		IncludePageCommands:           *includePageCommands,
		IncludePagesCommandsHierarchy: *includePagesCommandsHierarchy,
	}

	// Check if batch mode (source-dir) or single file mode
	if *sourceDir != "" {
		// Batch mode: process all MPR files in directory
		processBatchMode(*sourceDir, *outputDir, *outputFormat, &options)
	} else {
		// Single file mode
		if flag.NArg() < 1 {
			flag.Usage()
			os.Exit(1)
		}
		processSingleFile(flag.Arg(0), *outputDir, *outputFormat, &options)
	}
}

// processBatchMode scans a directory recursively for MPR files and generates reports
func processBatchMode(sourceDir string, outputDir string, outputFormat string, options *ReportOptions) {
	fmt.Printf("📁 Batch mode: Scanning directory %s\n", sourceDir)

	// Search for all .mpr files recursively
	mprFiles, err := findMPRFilesRecursive(sourceDir)
	if err != nil {
		fmt.Printf("❌ Error searching for MPR files: %v\n", err)
		os.Exit(1)
	}

	if len(mprFiles) == 0 {
		fmt.Printf("❌ No .mpr files found in directory tree\n")
		os.Exit(1)
	}

	fmt.Printf("📊 Found %d MPR file(s)\n", len(mprFiles))
	for i, file := range mprFiles {
		fmt.Printf("  [%d] %s\n", i+1, file)
	}
	fmt.Println()

	// Determine output directory
	finalOutputDir := outputDir
	if finalOutputDir == "" {
		// Create default output directory in source directory
		finalOutputDir = filepath.Join(sourceDir, "reports")
	}

	// Create output directory
	err = os.MkdirAll(finalOutputDir, 0755)
	if err != nil {
		fmt.Printf("❌ Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("💾 Output directory: %s\n\n", finalOutputDir)

	// Process each MPR file
	reportEntries := []ReportEntry{}

	for i, mprPath := range mprFiles {
		fmt.Printf("[%d/%d] Processing: %s\n", i+1, len(mprFiles), filepath.Base(mprPath))

		// Generate report filenames based on format
		mprFilename := filepath.Base(mprPath)
		mprNameWithoutExt := strings.TrimSuffix(mprFilename, filepath.Ext(mprFilename))

		var reportPathMD, reportPathJSON string
		if outputFormat == "md" || outputFormat == "both" {
			outputFilenameMD := mprNameWithoutExt + "-manifest.md"
			reportPathMD = filepath.Join(finalOutputDir, outputFilenameMD)
		}
		if outputFormat == "json" || outputFormat == "both" {
			outputFilenameJSON := mprNameWithoutExt + ".json"
			reportPathJSON = filepath.Join(finalOutputDir, outputFilenameJSON)
		}

		entry := ReportEntry{
			MPRPath:    mprPath,
			ReportPath: reportPathMD, // Use MD path for backward compatibility
		}

		// Store global MPR path for module traversal
		globalMPRPath = mprPath

		// Open and process MPR
		reader, err := modelsdk.Open(mprPath)
		if err != nil {
			fmt.Printf("  ❌ Error opening MPR: %v\n\n", err)
			entry.Success = false
			entry.Error = err.Error()
			reportEntries = append(reportEntries, entry)
			continue
		}

		// Extract project name from MPR filename
		projectName := strings.TrimSuffix(filepath.Base(mprPath), ".mpr")
		entry.ProjectName = projectName

		// Get Mendix version
		mendixVersion, err := reader.GetMendixVersion()
		if err != nil {
			mendixVersion = "Unknown"
		}

		// Generate report
		report := ManifestReport{
			ProjectName:   projectName,
			MendixVersion: mendixVersion,
			MPRPath:       mprPath,
			GeneratedAt:   time.Now().Format("2006-01-02 15:04:05"),
			Entities:      make(map[string][]EntityInfo),
		}

		// Open SQL database
		db, err := sql.Open("sqlite3", mprPath)
		if err != nil {
			fmt.Printf("  ❌ Error opening database: %v\n\n", err)
			entry.Success = false
			entry.Error = err.Error()
			reportEntries = append(reportEntries, entry)
			continue
		}

		contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

		// Collect data based on options (silent mode - no progress output)
		if options.IncludeEntities {
			collectExternalEntities(db, contentsDir, &report)
		}
		if options.IncludeMicroflows {
			collectMicroflowCalls(db, contentsDir, &report)
		}
		if options.IncludeWidgets {
			collectSignalManagerWidgets(reader, mprPath, &report)
		}
		if options.IncludeRoles || options.IncludeNavigation {
			pageAccessMap := make(map[string][]string)
			if options.IncludeRoles {
				collectSystemRolesAndPageAccess(db, contentsDir, &report)
				for _, pageAccess := range report.PageAccess {
					qualifiedName := pageAccess.Module + "." + pageAccess.PageName
					pageAccessMap[qualifiedName] = pageAccess.AllowedRoles
				}
			}
			if options.IncludeNavigation {
				collectNavigationItems(db, contentsDir, &report, pageAccessMap)
			}
		}

		db.Close()

		// Generate reports based on format
		if outputFormat == "md" || outputFormat == "both" {
			err = generateMarkdownReport(&report, reportPathMD, options)
			if err != nil {
				fmt.Printf("  ❌ Error generating Markdown report: %v\n\n", err)
				entry.Success = false
				entry.Error = err.Error()
				reportEntries = append(reportEntries, entry)
				continue
			}
		}
		if outputFormat == "json" || outputFormat == "both" {
			err = generateJSONReport(&report, reportPathJSON, options)
			if err != nil {
				fmt.Printf("  ❌ Error generating JSON report: %v\n\n", err)
				entry.Success = false
				entry.Error = err.Error()
				reportEntries = append(reportEntries, entry)
				continue
			}
			// Update entry.ReportPath to JSON if JSON-only
			if outputFormat == "json" {
				entry.ReportPath = reportPathJSON
			}
		}

		entry.Success = true
		reportEntries = append(reportEntries, entry)

		if outputFormat == "md" {
			fmt.Printf("  ✅ Report generated: %s\n\n", filepath.Base(reportPathMD))
		} else if outputFormat == "json" {
			fmt.Printf("  ✅ Report generated: %s\n\n", filepath.Base(reportPathJSON))
		} else {
			fmt.Printf("  ✅ Reports generated: %s, %s\n\n", filepath.Base(reportPathMD), filepath.Base(reportPathJSON))
		}
	}

	// Generate index files based on format
	if outputFormat == "md" || outputFormat == "both" {
		indexPath := filepath.Join(finalOutputDir, "index.md")
		err = generateIndexFile(indexPath, reportEntries)
		if err != nil {
			fmt.Printf("❌ Error generating index.md: %v\n", err)
			os.Exit(1)
		}
	}
	if outputFormat == "json" || outputFormat == "both" {
		indexJSONPath := filepath.Join(finalOutputDir, "index.json")
		err = generateIndexJSON(indexJSONPath, reportEntries)
		if err != nil {
			fmt.Printf("❌ Error generating index.json: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("\n📊 Summary: %d successful, %d failed\n", countSuccessful(reportEntries), countFailed(reportEntries))
	if outputFormat == "md" {
		fmt.Printf("📝 Index file: %s\n", filepath.Join(finalOutputDir, "index.md"))
	} else if outputFormat == "json" {
		fmt.Printf("📝 Index file: %s\n", filepath.Join(finalOutputDir, "index.json"))
	} else {
		fmt.Printf("📝 Index files: %s, %s\n", filepath.Join(finalOutputDir, "index.md"), filepath.Join(finalOutputDir, "index.json"))
	}
}

// processSingleFile processes a single MPR file
func processSingleFile(mprPathArg string, outputDir string, outputFormat string, options *ReportOptions) {
	mprPath := mprPathArg

	// If "." is passed, search for .mpr file in current directory
	if mprPath == "." {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}

		// Search for .mpr files in current directory
		mprFiles, err := filepath.Glob(filepath.Join(currentDir, "*.mpr"))
		if err != nil {
			fmt.Printf("Error searching for .mpr files: %v\n", err)
			os.Exit(1)
		}

		if len(mprFiles) == 0 {
			fmt.Printf("❌ No .mpr files found in current directory\n")
			os.Exit(1)
		} else if len(mprFiles) > 1 {
			fmt.Printf("❌ Multiple .mpr files found in current directory:\n")
			for _, file := range mprFiles {
				fmt.Printf("  - %s\n", filepath.Base(file))
			}
			fmt.Printf("\nPlease specify which file to use.\n")
			os.Exit(1)
		} else {
			mprPath = mprFiles[0]
			fmt.Printf("🔍 Found MPR file: %s\n", filepath.Base(mprPath))
		}
	}

	// Generate output filename base from MPR filename
	mprFilename := filepath.Base(mprPath)
	mprNameWithoutExt := strings.TrimSuffix(mprFilename, filepath.Ext(mprFilename))

	// Determine output paths based on format
	var finalOutputPathMD, finalOutputPathJSON string
	if outputFormat == "md" || outputFormat == "both" {
		outputFilenameMD := mprNameWithoutExt + "-manifest.md"
		if outputDir != "" {
			finalOutputPathMD = filepath.Join(outputDir, outputFilenameMD)
		} else {
			finalOutputPathMD = outputFilenameMD
		}
	}
	if outputFormat == "json" || outputFormat == "both" {
		outputFilenameJSON := mprNameWithoutExt + ".json"
		if outputDir != "" {
			finalOutputPathJSON = filepath.Join(outputDir, outputFilenameJSON)
		} else {
			finalOutputPathJSON = outputFilenameJSON
		}
	}

	// Create output directory if needed
	if outputDir != "" {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Store global MPR path for module traversal
	globalMPRPath = mprPath

	fmt.Printf("🔍 Opening MPR: %s\n", mprPath)

	// Open MPR with modelsdk
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening MPR: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	// Get project info
	projectName := filepath.Base(mprPath)
	projectName = strings.TrimSuffix(projectName, ".mpr")

	mendixVersion := "Unknown"
	if version, err := reader.GetMendixVersion(); err == nil {
		mendixVersion = version
	}

	fmt.Printf("📦 Project: %s\n", projectName)
	fmt.Printf("📌 Mendix Version: %s\n", mendixVersion)
	fmt.Println()

	// Initialize report
	report := ManifestReport{
		ProjectName:   projectName,
		MendixVersion: mendixVersion,
		MPRPath:       mprPath,
		GeneratedAt:   time.Now().Format("2006-01-02 15:04:05"),
		Entities:      make(map[string][]EntityInfo),
	}

	// Open SQL database for entity queries
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Count enabled sections
	totalPhases := 0
	if options.IncludeEntities {
		totalPhases++
	}
	if options.IncludeMicroflows {
		totalPhases++
	}
	if options.IncludeWidgets {
		totalPhases++
	}
	if options.IncludeRoles || options.IncludeNavigation {
		totalPhases++
	}
	if options.IncludeNavigation {
		totalPhases++
	}
	if options.IncludePageCommands {
		totalPhases++
	}
	if options.IncludePagesCommandsHierarchy {
		totalPhases++ // Phase for pages analysis
	}
	totalPhases++ // Final report generation

	fmt.Printf("\n📋 Analysis plan: %d phase(s) to complete\n", totalPhases)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	currentPhase := 0

	// ==== SECTION 1: External Entities ====
	if options.IncludeEntities {
		currentPhase++
		fmt.Printf("\n[Phase %d/%d] 🔎 Scanning for external entities...\n", currentPhase, totalPhases)
		err = collectExternalEntities(db, contentsDir, &report)
		if err != nil {
			fmt.Printf("Error collecting entities: %v\n", err)
			os.Exit(1)
		}
	}

	// ==== SECTION 2: Microflow/Action Calls ====
	if options.IncludeMicroflows {
		currentPhase++
		fmt.Printf("\n[Phase %d/%d] 🔎 Scanning for microflow/action calls...\n", currentPhase, totalPhases)
		err = collectMicroflowCalls(db, contentsDir, &report)
		if err != nil {
			fmt.Printf("Error collecting microflow calls: %v\n", err)
			os.Exit(1)
		}
	}

	// ==== SECTION 3: Signal Manager Widgets ====
	if options.IncludeWidgets {
		currentPhase++
		fmt.Printf("\n[Phase %d/%d] 🔎 Scanning for Signal Manager widgets...\n", currentPhase, totalPhases)
		err = collectSignalManagerWidgets(reader, mprPath, &report)
		if err != nil {
			fmt.Printf("Error collecting widgets: %v\n", err)
			os.Exit(1)
		}
	}

	// ==== SECTION 4: System Roles & Page Accessibility ====
	// Collect this first to use for navigation items
	pageAccessMap := make(map[string][]string) // Map: pageQualifiedName -> []roleNames
	if options.IncludeRoles || options.IncludeNavigation {
		currentPhase++
		fmt.Printf("\n[Phase %d/%d] 🔎 Scanning for system roles and page accessibility...\n", currentPhase, totalPhases)
		err = collectSystemRolesAndPageAccess(db, contentsDir, &report)
		if err != nil {
			fmt.Printf("Error collecting roles and page access: %v\n", err)
			os.Exit(1)
		}

		// Build map for quick lookup
		for _, pageAccess := range report.PageAccess {
			pageAccessMap[pageAccess.PageName] = pageAccess.AllowedRoles
		}
	}

	// ==== SECTION 5: Navigation Items ====
	if options.IncludeNavigation {
		currentPhase++
		fmt.Printf("\n[Phase %d/%d] 🔎 Scanning for navigation items...\n", currentPhase, totalPhases)
		err = collectNavigationItems(db, contentsDir, &report, pageAccessMap)
		if err != nil {
			fmt.Printf("Error collecting navigation items: %v\n", err)
			os.Exit(1)
		}
	}

	// ==== SECTION 6: Page Commands ====
	if options.IncludePageCommands {
		currentPhase++
		fmt.Printf("\n[Phase %d/%d] 🖋️ Scanning for page commands...\n", currentPhase, totalPhases)
		pageCommands, err := collectPageCommands(db, contentsDir, mprPath, report.NavigationItems, report.MicroflowCalls)
		if err != nil {
			fmt.Printf("Error collecting page commands: %v\n", err)
			os.Exit(1)
		}
		report.NavigationCommands = pageCommands
		fmt.Printf("  ✅ Completed: Page commands collected\n")
		fmt.Printf("  📊 Result: %d page(s) with commands\n", len(report.NavigationCommands))
	}

	// ==== SECTION 7: Pages Analysis (Detailed with Recursive Hierarchy) ====
	if options.IncludePagesCommandsHierarchy {
		currentPhase++
		fmt.Printf("\n[Phase %d/%d] 🔍 Analyzing pages/panels with recursive microflow hierarchy...\n", currentPhase, totalPhases)
		pagesAnalysis, err := collectPagesWithMicroflows(db, contentsDir)
		if err != nil {
			fmt.Printf("Error analyzing pages: %v\n", err)
			os.Exit(1)
		}
		report.PagesAnalysis = pagesAnalysis
		fmt.Printf("  ✅ Completed: Pages analysis with recursive hierarchy\n")
		fmt.Printf("  📊 Result: %d page(s)/panel(s) analyzed\n", len(report.PagesAnalysis))

		// Populate PageCommands (simplified view) from PagesAnalysis
		for _, page := range report.PagesAnalysis {
			appNames := make([]string, 0)
			cmdNames := make([]string, 0)
			for _, cmd := range page.TargetCommands {
				if idx := strings.LastIndex(cmd, "."); idx >= 0 {
					appNames = append(appNames, cmd[:idx])
					cmdNames = append(cmdNames, cmd[idx+1:])
				} else {
					appNames = append(appNames, "-")
					cmdNames = append(cmdNames, cmd)
				}
			}
			report.PageCommands = append(report.PageCommands, PageCommandSummary{
				PageName:       page.Name,
				Module:         page.Module,
				TargetAppNames: appNames,
				TargetCommands: cmdNames,
			})
		}
	}

	// Generate reports based on format
	currentPhase++
	if outputFormat == "md" || outputFormat == "both" {
		fmt.Printf("\n[Phase %d/%d] 📝 Generating Markdown report...\n", currentPhase, totalPhases)
		fmt.Printf("  📄 Output file: %s\n", finalOutputPathMD)
		err = generateMarkdownReport(&report, finalOutputPathMD, options)
		if err != nil {
			fmt.Printf("Error generating Markdown report: %v\n", err)
			os.Exit(1)
		}
	}
	if outputFormat == "json" || outputFormat == "both" {
		if outputFormat == "both" {
			fmt.Printf("\n[Phase %d/%d] 📝 Generating JSON report...\n", currentPhase, totalPhases)
		} else {
			fmt.Printf("\n[Phase %d/%d] 📝 Generating JSON report...\n", currentPhase, totalPhases)
		}
		fmt.Printf("  📄 Output file: %s\n", finalOutputPathJSON)
		err = generateJSONReport(&report, finalOutputPathJSON, options)
		if err != nil {
			fmt.Printf("Error generating JSON report: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("\n✅ Report generated successfully!")
	fmt.Printf("\n📊 Summary:\n")
	if options.IncludeEntities {
		totalEntities := 0
		for _, entities := range report.Entities {
			totalEntities += len(entities)
		}
		fmt.Printf("  • External Entities: %d (in %d modules)\n", totalEntities, len(report.Entities))
	}
	if options.IncludeMicroflows {
		fmt.Printf("  • Microflow Calls: %d\n", len(report.MicroflowCalls))
	}
	if options.IncludeWidgets {
		fmt.Printf("  • Signal Manager Subscriptions: %d\n", len(report.Widgets))
	}
	if options.IncludeNavigation {
		fmt.Printf("  • Navigation Items: %d\n", len(report.NavigationItems))
	}
	if options.IncludeRoles {
		fmt.Printf("  • System Roles: %d\n", len(report.SystemRoles))
		fmt.Printf("  • Page Access Rules: %d\n", len(report.PageAccess))
	}
	if outputFormat == "md" {
		fmt.Printf("\n📁 Report saved to: %s\n", finalOutputPathMD)
	} else if outputFormat == "json" {
		fmt.Printf("\n📁 Report saved to: %s\n", finalOutputPathJSON)
	} else {
		fmt.Printf("\n📁 Reports saved to:\n")
		fmt.Printf("  • Markdown: %s\n", finalOutputPathMD)
		fmt.Printf("  • JSON: %s\n", finalOutputPathJSON)
	}
}

// ==== SECTION 1: External Entities Collection ====

func collectExternalEntities(db *sql.DB, contentsDir string, report *ManifestReport) error {
	// Load all domain models
	fmt.Printf("  🔍 Loading domain models...\n")
	domainModels, err := listDomainModels(db, contentsDir)
	if err != nil {
		return fmt.Errorf("failed to list domain models: %w", err)
	}
	fmt.Printf("  📦 Found %d domain model(s) to analyze\n", len(domainModels))

	totalEntities := 0
	processedModules := 0

	// For each domain model
	for _, dm := range domainModels {
		processedModules++
		moduleName := extractModuleName(dm.Name)

		// Skip marketplace modules
		if isMarketplaceModule(moduleName) {
			continue
		}

		// Get external entities from domain model
		entities := findExternalEntitiesInDomainModel(dm, moduleName)

		if len(entities) > 0 {
			report.Entities[moduleName] = entities
			totalEntities += len(entities)
			fmt.Printf("  ✓ Module '%s': found %d external entities\n", moduleName, len(entities))
		}
	}

	fmt.Printf("\n  ✅ Completed: Analyzed %d modules\n", processedModules)
	fmt.Printf("  📊 Result: %d external entities in %d modules\n", totalEntities, len(report.Entities))
	return nil
}

func listDomainModels(db *sql.DB, contentsDir string) ([]DomainModelInfo, error) {
	rows, err := db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, ContentsHash
		FROM Unit
		WHERE ContainmentName = 'DomainModel'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	var domainModels []DomainModelInfo

	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		containerIDStr := blobToUUID(containerID)

		// Load contents
		contents, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Check if it's a DomainModel
		if typeStr, ok := contents["$Type"].(string); ok {
			if typeStr == "DomainModels$DomainModel" {
				// Get module name from container
				containerContents, err := loadUnitContents(contentsDir, containerIDStr)
				if err == nil {
					name := extractNameFromContents(containerContents)
					domainModels = append(domainModels, DomainModelInfo{
						Name:    name,
						UnitID:  unitIDStr,
						Content: contents,
					})
				}
			}
		}
	}

	return domainModels, nil
}

func findExternalEntitiesInDomainModel(dm DomainModelInfo, moduleName string) []EntityInfo {
	var entities []EntityInfo

	// Look for Entities array in the domain model
	entitiesRaw, hasEntities := dm.Content["Entities"]
	if !hasEntities {
		return entities
	}

	// Convert to []interface{} regardless of the underlying type (handles bson.A)
	var entitiesArray []interface{}
	switch v := entitiesRaw.(type) {
	case []interface{}:
		entitiesArray = v
	case bson.A:
		entitiesArray = []interface{}(v)
	default:
		return entities
	}

	if entitiesArray == nil {
		return entities
	}

	for _, entityItem := range entitiesArray {
		// Skip the array length marker (first element is usually 1, 2, or 3)
		if intVal, ok := entityItem.(int32); ok && (intVal >= 1 && intVal <= 3) {
			continue
		}
		if intVal, ok := entityItem.(int64); ok && (intVal >= 1 && intVal <= 3) {
			continue
		}

		// Try to get entity as embedded map
		if entityMap, ok := entityItem.(map[string]interface{}); ok {
			// Check type - should be DomainModels$EntityImpl or DomainModels$ExternalEntity
			entityType := ""
			if typeStr, ok := entityMap["$Type"].(string); ok {
				entityType = typeStr
			}

			// Only process if it's an entity
			if entityType == "DomainModels$EntityImpl" || entityType == "DomainModels$ExternalEntity" {
				if isExternalEntity(entityMap) && isPersistableEntity(entityMap) {
					entityName := extractNameFromContents(entityMap)
					publishedFrom := extractPublishedFrom(entityMap, entityName)
					attributes := extractAttributes(entityMap)

					if entityName != "" {
						entities = append(entities, EntityInfo{
							Name:          entityName,
							Module:        moduleName,
							Attributes:    attributes,
							PublishedFrom: publishedFrom,
						})
					}
				}
			}
		}
	}

	return entities
}

func isExternalEntity(entityMap map[string]interface{}) bool {
	// Check for $Type
	if typeStr, ok := entityMap["$Type"].(string); ok {
		// ExternalEntity is explicitly external
		if typeStr == "DomainModels$ExternalEntity" {
			return true
		}
	}

	// Check for Source field with OData type (MAIN INDICATOR!)
	if source, ok := entityMap["Source"].(map[string]interface{}); ok {
		if sourceType, ok := source["$Type"].(string); ok {
			// OData entities have Rest$ODataEntityTypeSource
			if sourceType == "Rest$ODataEntityTypeSource" {
				return true
			}
			// Also check for other external source types
			if strings.Contains(sourceType, "OData") || strings.Contains(sourceType, "Rest") || strings.Contains(sourceType, "External") {
				return true
			}
		}
	}

	// Check for IsExternal flag
	if isExternal, ok := entityMap["IsExternal"].(bool); ok && isExternal {
		return true
	}

	// Check for RemoteSourceDocument (older format)
	if remoteSource, ok := entityMap["RemoteSourceDocument"].(string); ok && remoteSource != "" {
		return true
	}

	// Check for RemoteSource field
	if remoteSource, ok := entityMap["RemoteSource"]; ok && remoteSource != nil {
		if strVal, ok := remoteSource.(string); ok && strVal != "" {
			return true
		}
		if mapVal, ok := remoteSource.(map[string]interface{}); ok && len(mapVal) > 0 {
			return true
		}
	}

	return false
}

func isPersistableEntity(entityMap map[string]interface{}) bool {
	// Check MaybeGeneralization.Persistable field
	// This field indicates if the entity is persistable (stored in database)
	// Non-persistable entities are typically transient/temporary objects used for operations
	if maybeGen, ok := entityMap["MaybeGeneralization"].(map[string]interface{}); ok {
		if persistable, ok := maybeGen["Persistable"].(bool); ok {
			return persistable
		}
	}

	// If Persistable field is not found, assume it's persistable (default)
	// This ensures we don't accidentally filter out entities from older Mendix versions
	return true
}

func extractPublishedFrom(entityMap map[string]interface{}, entityName string) string {
	// Extract from Source field (OData entities)
	if source, ok := entityMap["Source"].(map[string]interface{}); ok {
		// First try EntityTypeName (direct service name)
		if entityTypeName, ok := source["EntityTypeName"].(string); ok && entityTypeName != "" {
			return entityTypeName
		}

		// Try SourceDocument (format: "ModuleName.ServiceName")
		if sourceDoc, ok := source["SourceDocument"].(string); ok && sourceDoc != "" {
			// Extract service name from "ModuleName.ServiceName" format
			parts := strings.Split(sourceDoc, ".")
			if len(parts) == 2 {
				return parts[1] // Return "ServiceName"
			}
			return sourceDoc // Return as-is if format is different
		}

		// Try RemoteName as fallback
		if remoteName, ok := source["RemoteName"].(string); ok && remoteName != "" {
			return remoteName
		}
	}

	// Fallback to other location indicators
	if remoteSource, ok := entityMap["RemoteSourceDocument"].(string); ok && remoteSource != "" {
		return remoteSource
	}
	if remoteSource, ok := entityMap["RemoteSource"].(string); ok && remoteSource != "" {
		return remoteSource
	}

	// Use entity name as fallback
	if entityName != "" {
		return entityName
	}

	return "Unknown"
}

func extractAttributes(entityMap map[string]interface{}) []string {
	var attributes []string

	// Look for Attributes array
	attributesRaw, hasAttributes := entityMap["Attributes"]
	if !hasAttributes {
		return attributes
	}

	// Convert to []interface{} (handles bson.A)
	var attributesArray []interface{}
	switch v := attributesRaw.(type) {
	case []interface{}:
		attributesArray = v
	case bson.A:
		attributesArray = []interface{}(v)
	default:
		return attributes
	}

	for _, attrItem := range attributesArray {
		// Skip array length markers
		if intVal, ok := attrItem.(int32); ok && (intVal >= 1 && intVal <= 3) {
			continue
		}
		if intVal, ok := attrItem.(int64); ok && (intVal >= 1 && intVal <= 3) {
			continue
		}

		// Try to get attribute as map
		if attrMap, ok := attrItem.(map[string]interface{}); ok {
			if name, ok := attrMap["Name"].(string); ok && name != "" {
				// Get type info if available
				typeInfo := extractAttributeType(attrMap)

				if typeInfo != "" {
					attributes = append(attributes, fmt.Sprintf("%s (%s)", name, typeInfo))
				} else {
					attributes = append(attributes, name)
				}
			}
		}
	}

	return attributes
}

func extractAttributeType(attrMap map[string]interface{}) string {
	// Try standard Type field with $Type
	if attrType, ok := attrMap["Type"].(map[string]interface{}); ok {
		if typeName, ok := attrType["$Type"].(string); ok {
			// Extract just the type name (e.g., "StringAttributeType" -> "String")
			return extractTypeName(typeName)
		}
	}

	// Try RemoteAttributeTypeName (for OData external entities)
	if remoteType, ok := attrMap["RemoteAttributeTypeName"].(string); ok && remoteType != "" {
		// OData types like "Edm.String", "Edm.Int32", etc.
		return formatODataType(remoteType)
	}

	// Try ValueType (alternative field)
	if valueType, ok := attrMap["ValueType"].(string); ok && valueType != "" {
		return valueType
	}

	// Check if it's a reference attribute (association)
	if attrMap["$Type"] != nil {
		if typeStr, ok := attrMap["$Type"].(string); ok {
			if strings.Contains(typeStr, "AssociationAttribute") {
				return "Reference"
			}
			if strings.Contains(typeStr, "Calculated") {
				return "Calculated"
			}
		}
	}

	// Fallback: try to infer from attribute name patterns
	if name, ok := attrMap["Name"].(string); ok {
		return inferTypeFromName(name)
	}

	return ""
}

func formatODataType(odataType string) string {
	// Convert OData types to readable format
	// "Edm.String" -> "String"
	// "Edm.Int32" -> "Integer"
	// "Edm.DateTime" -> "DateTime"
	if strings.HasPrefix(odataType, "Edm.") {
		baseType := strings.TrimPrefix(odataType, "Edm.")

		// Map common OData types to Mendix types
		switch baseType {
		case "Int32", "Int64", "Int16":
			return "Integer"
		case "Decimal", "Double":
			return "Decimal"
		case "Boolean":
			return "Boolean"
		case "DateTime", "DateTimeOffset":
			return "DateTime"
		case "Guid":
			return "String (GUID)"
		default:
			return baseType
		}
	}
	return odataType
}

func inferTypeFromName(name string) string {
	// Infer type based on common naming patterns
	nameLower := strings.ToLower(name)

	if strings.HasSuffix(nameLower, "_id") || strings.HasSuffix(nameLower, "id") {
		return "String"
	}
	if strings.HasPrefix(nameLower, "is") || strings.HasPrefix(nameLower, "has") || strings.HasPrefix(nameLower, "do") {
		return "Boolean"
	}
	if nameLower == "succeeded" || strings.HasSuffix(nameLower, "enabled") || strings.HasSuffix(nameLower, "frozen") || strings.HasSuffix(nameLower, "locked") || strings.HasSuffix(nameLower, "hidden") {
		return "Boolean"
	}
	if strings.Contains(nameLower, "date") || strings.Contains(nameLower, "time") || strings.HasSuffix(nameLower, "edon") || strings.HasSuffix(nameLower, "atedon") {
		return "DateTime"
	}
	if strings.Contains(nameLower, "count") || strings.Contains(nameLower, "seed") || strings.Contains(nameLower, "increment") || strings.Contains(nameLower, "sequence") {
		return "Integer"
	}
	if strings.Contains(nameLower, "multiplier") || strings.Contains(nameLower, "addend") || strings.Contains(nameLower, "exponent") || strings.Contains(nameLower, "factor") {
		return "Decimal"
	}
	if strings.Contains(nameLower, "amount") || strings.Contains(nameLower, "value") || strings.Contains(nameLower, "max") || strings.Contains(nameLower, "min") {
		return "Integer/Decimal"
	}

	// Default fallback
	return "String"
}

func extractTypeName(fullType string) string {
	// Convert "DomainModels$StringAttributeType" -> "String"
	// Convert "DomainModels$IntegerAttributeType" -> "Integer"
	parts := strings.Split(fullType, "$")
	if len(parts) > 1 {
		typeName := parts[1]
		typeName = strings.TrimSuffix(typeName, "AttributeType")
		return typeName
	}
	return fullType
}

// ==== SECTION 2: Microflow/Action Calls Collection ====

func collectMicroflowCalls(db *sql.DB, contentsDir string, report *ManifestReport) error {
	// List all microflows
	fmt.Printf("  🔍 Loading microflows...\n")
	microflows, err := listMicroflows(db, contentsDir)
	if err != nil {
		return fmt.Errorf("failed to list microflows: %w", err)
	}
	fmt.Printf("  📦 Found %d microflow(s) to analyze\n", len(microflows))

	totalCalls := 0
	processedCount := 0
	targetMicroflow := "EXFN_ServiceLayer.CallCommand_MF"
	targetJavaAction := "EXFN_ServiceLayer.CallCommandAction"

	// Search for calls in each microflow
	for _, mf := range microflows {
		processedCount++
		// Skip marketplace modules
		if isMarketplaceModule(mf.ModuleName) {
			continue
		}

		calls := findMicroflowCalls(mf, targetMicroflow, targetJavaAction, db, contentsDir)

		for _, call := range calls {
			report.MicroflowCalls = append(report.MicroflowCalls, MicroflowCallInfo{
				MicroflowName: mf.Name,
				Module:        mf.ModuleName,
				ActivityName:  call.ActivityName,
				Caption:       call.Caption,
				CallType:      call.CallType,
				TargetName:    call.TargetName,
				AppName:       call.AppName,
				CommandName:   call.CommandName,
			})
			totalCalls++
		}
	}

	fmt.Printf("\n  ✅ Completed: Analyzed %d microflows\n", processedCount)
	fmt.Printf("  📊 Result: %d microflow/action calls found\n", totalCalls)
	return nil
}

func listMicroflows(db *sql.DB, contentsDir string) ([]MicroflowInfo, error) {
	query := `SELECT UnitID, ContainerID, ContainmentName, ContentsHash FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var microflows []MicroflowInfo
	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			continue
		}

		// Convert UUID blob to string
		unitIDStr := blobToUUID(unitID)

		// Load content from mprcontents
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Extract type from BSON content
		typeName := ""
		if t, ok := content["$Type"].(string); ok {
			typeName = t
		}

		// Only process Microflows and Nanoflows
		if typeName != "Microflows$Microflow" && typeName != "Microflows$Nanoflow" {
			continue
		}

		// Skip microflows excluded from project
		if excluded, ok := content["ExcludedFromProject"].(bool); ok && excluded {
			continue
		}

		// Extract name from BSON content
		name := extractNameFromContents(content)

		// Try to get module name from ObjectCollection > Objects > MicroflowParameter > VariableType > Entity
		moduleName := extractModuleFromMicroflowParams(content)

		// If that fails, use ContainmentName as fallback (often a folder, but better than nothing)
		if moduleName == "" {
			moduleName = containmentName
		}

		microflows = append(microflows, MicroflowInfo{
			ID:         unitIDStr,
			Name:       name,
			ModuleName: moduleName,
			Content:    content,
		})
	}

	return microflows, nil
}

// extractModuleFromMicroflowParams tries to extract module name from microflow parameters or allowed roles
func extractModuleFromMicroflowParams(content map[string]interface{}) string {
	// Strategy 1: Navigate ObjectCollection > Objects > find MicroflowParameter > VariableType > Entity
	if objColl, ok := content["ObjectCollection"].(map[string]interface{}); ok {
		if objects, ok := objColl["Objects"].(bson.A); ok {
			for _, obj := range objects {
				if objMap, ok := obj.(map[string]interface{}); ok {
					// Check if it's a MicroflowParameter
					if objType, ok := objMap["$Type"].(string); ok && objType == "Microflows$MicroflowParameter" {
						// Get VariableType > Entity
						if varType, ok := objMap["VariableType"].(map[string]interface{}); ok {
							if entity, ok := varType["Entity"].(string); ok && entity != "" {
								// Entity is in format "Module.EntityName", extract module
								return extractModuleName(entity)
							}
						}
					}
				}
			}
		}
	}

	// Strategy 2: Extract from AllowedModuleRoles (format: "Module.Role")
	if roles, ok := content["AllowedModuleRoles"].(bson.A); ok {
		for _, role := range roles {
			if roleStr, ok := role.(string); ok && roleStr != "" {
				// Extract module from "Module.Role"
				moduleName := extractModuleName(roleStr)
				if moduleName != roleStr && moduleName != "" {
					return moduleName
				}
			}
		}
	}

	return ""
}

type TempMicroflowCall struct {
	ActivityName string
	Caption      string
	TargetName   string
	CallType     string
	AppName      string
	CommandName  string
}

func findMicroflowCalls(mf MicroflowInfo, targetMicroflow string, targetJavaAction string, db *sql.DB, contentsDir string) []TempMicroflowCall {
	var calls []TempMicroflowCall

	// Convert content to searchable format
	jsonData, err := json.Marshal(mf.Content)
	if err != nil {
		return calls
	}

	var contentMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &contentMap); err != nil {
		return calls
	}

	// Search for ActionActivity with MicroflowCall, JavaAction, or ExternalAction
	// Pass the original BSON content for parameter detection
	searchActivities(contentMap, targetMicroflow, targetJavaAction, &calls, mf.Content, mf.Name, db, contentsDir)

	return calls
}

// resolveVariableValue searches the microflow for variable assignments or parameters
func resolveVariableValue(originalContent map[string]interface{}, variableName string) string {
	// Remove $ prefix if present
	cleanVarName := strings.TrimPrefix(variableName, "$")

	// First check if it's a microflow parameter
	if objColl, ok := originalContent["ObjectCollection"].(map[string]interface{}); ok {
		if objectsRaw, ok := objColl["Objects"]; ok {
			// Handle both bson.A and []interface{}
			var objects []interface{}
			switch v := objectsRaw.(type) {
			case bson.A:
				objects = []interface{}(v)
			case []interface{}:
				objects = v
			default:
				return ""
			}

			for _, obj := range objects {
				if objMap, ok := obj.(map[string]interface{}); ok {
					if objType, ok := objMap["$Type"].(string); ok && objType == "Microflows$MicroflowParameter" {
						if paramName, ok := objMap["Name"].(string); ok && paramName == cleanVarName {
							return "<parameter>"
						}
					}
				}
			}
		}
	}

	// Then search for variable assignments
	var result string
	findVariableAssignment(originalContent, cleanVarName, &result)
	return result
}

func findVariableAssignment(obj interface{}, varName string, result *string) {
	if *result != "" {
		return // Already found
	}

	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this is a ChangeVariableActivity
		if typeStr, ok := v["$Type"].(string); ok {
			if typeStr == "Microflows$ChangeVariableActivity" {
				action, ok := v["Action"].(map[string]interface{})
				if !ok {
					break
				}

				// Check variable name
				if chgVar, ok := action["ChangeVariableName"].(string); ok && chgVar == varName {
					// Extract the value expression
					if valueExpr, ok := action["Value"].(string); ok {
						// Clean up the expression (remove quotes, etc.)
						cleanValue := strings.Trim(valueExpr, "'\"")
						*result = cleanValue
						return
					}
				}
			}
		}

		// Recursively search all fields
		for _, value := range v {
			findVariableAssignment(value, varName, result)
			if *result != "" {
				return
			}
		}

	case []interface{}:
		for _, item := range v {
			findVariableAssignment(item, varName, result)
			if *result != "" {
				return
			}
		}
	}
}

func searchActivities(obj interface{}, targetMicroflow string, targetJavaAction string, calls *[]TempMicroflowCall, contentMap map[string]interface{}, microflowName string, db *sql.DB, contentsDir string) {
	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this is an ActionActivity
		if typeStr, ok := v["$Type"].(string); ok {
			if typeStr == "Microflows$ActionActivity" {
				// Check for all three call types
				checkMicroflowCall(v, targetMicroflow, calls, contentMap)
				checkJavaAction(v, targetJavaAction, calls, contentMap, microflowName, db, contentsDir)
				checkExternalAction(v, calls, contentMap)
			}
		}

		// Recursively search all fields
		for _, value := range v {
			searchActivities(value, targetMicroflow, targetJavaAction, calls, contentMap, microflowName, db, contentsDir)
		}

	case []interface{}:
		for _, item := range v {
			searchActivities(item, targetMicroflow, targetJavaAction, calls, contentMap, microflowName, db, contentsDir)
		}
	}
}

func checkMicroflowCall(activity map[string]interface{}, targetMicroflow string, calls *[]TempMicroflowCall, contentMap map[string]interface{}) {
	action, ok := activity["Action"].(map[string]interface{})
	if !ok {
		return
	}

	actionType, ok := action["$Type"].(string)
	if !ok || actionType != "Microflows$MicroflowCallAction" {
		return
	}

	microflowCallObj, ok := action["MicroflowCall"].(map[string]interface{})
	if !ok {
		return
	}

	microflowCall, ok := microflowCallObj["Microflow"].(string)
	if !ok {
		return
	}

	// Filter: only calls to targetMicroflow (EXFN_ServiceLayer.CallCommand_MF)
	if !strings.Contains(microflowCall, targetMicroflow) {
		return
	}

	// Extract details
	caption := ""
	if cap, ok := activity["Caption"].(string); ok {
		caption = cap
	}

	activityName := ""
	if name, ok := activity["Name"].(string); ok {
		activityName = name
	}

	// Extract AppName and CommandName from ParameterMappings
	appName := ""
	commandName := ""
	if paramMappings, ok := microflowCallObj["ParameterMappings"].([]interface{}); ok {
		for _, mapping := range paramMappings {
			if mappingMap, ok := mapping.(map[string]interface{}); ok {
				if param, ok := mappingMap["Parameter"].(string); ok {
					argument := ""
					if arg, ok := mappingMap["Argument"].(string); ok {
						argument = arg
					}

					if strings.Contains(strings.ToLower(param), "appname") {
						appName = argument
						// Resolve variable if it starts with $
						if strings.HasPrefix(appName, "$") {
							if resolved := resolveVariableValue(contentMap, appName); resolved != "" {
								appName = resolved
							}
						}
					} else if strings.Contains(strings.ToLower(param), "commandname") {
						commandName = argument
						// Resolve variable if it starts with $
						if strings.HasPrefix(commandName, "$") {
							if resolved := resolveVariableValue(contentMap, commandName); resolved != "" {
								commandName = resolved
							}
						}
					}
				}
			}
		}
	}

	*calls = append(*calls, TempMicroflowCall{
		ActivityName: activityName,
		Caption:      caption,
		TargetName:   microflowCall,
		CallType:     "MicroflowCall",
		AppName:      appName,
		CommandName:  commandName,
	})
}

func checkJavaAction(activity map[string]interface{}, targetJavaAction string, calls *[]TempMicroflowCall, contentMap map[string]interface{}, microflowName string, db *sql.DB, contentsDir string) {
	action, ok := activity["Action"].(map[string]interface{})
	if !ok {
		return
	}

	actionType, ok := action["$Type"].(string)
	if !ok || actionType != "Microflows$JavaActionCallAction" {
		return
	}

	// Get JavaAction name
	javaActionName := ""
	if ja, ok := action["JavaAction"].(string); ok {
		javaActionName = ja
	}
	if javaActionName == "" {
		if ja, ok := action["JavaActionQualifiedName"].(string); ok {
			javaActionName = ja
		}
	}

	if javaActionName == "" {
		return
	}

	// Filter: only calls to targetJavaAction (EXFN_ServiceLayer.CallCommandAction)
	if !strings.Contains(javaActionName, targetJavaAction) {
		return
	}

	// Extract details
	caption := ""
	if cap, ok := activity["Caption"].(string); ok {
		caption = cap
	}

	activityName := ""
	if name, ok := activity["Name"].(string); ok {
		activityName = name
	}

	// Extract AppName and CommandName from ParameterMappings
	appName := ""
	commandName := ""

	if paramMappings, ok := action["ParameterMappings"].([]interface{}); ok {
		for _, mapping := range paramMappings {
			if mappingMap, ok := mapping.(map[string]interface{}); ok {
				paramName := ""
				if param, ok := mappingMap["Parameter"].(string); ok {
					paramName = param
				}

				argument := ""
				if valueObj, ok := mappingMap["Value"].(map[string]interface{}); ok {
					if arg, ok := valueObj["Argument"].(string); ok {
						argument = strings.TrimSpace(arg)
					}
				}

				if strings.Contains(strings.ToLower(paramName), "appname") {
					appName = argument
					// Resolve variable if it starts with $
					if strings.HasPrefix(appName, "$") {
						if resolved := resolveVariableValue(contentMap, appName); resolved != "" {
							appName = resolved
						}
					}
				} else if strings.Contains(strings.ToLower(paramName), "commandname") {
					commandName = argument
					// Resolve variable if it starts with $
					if strings.HasPrefix(commandName, "$") {
						if resolved := resolveVariableValue(contentMap, commandName); resolved != "" {
							commandName = resolved
						}
					}
				}
			}
		}
	}

	// If AppName or CommandName are empty, "empty", or placeholders, use specific placeholders for JavaAction
	if appName == "" || strings.ToLower(appName) == "empty" || appName == "'AppName'" || appName == "<parameter>" {
		appName = "<JavaAppName>"
	}
	if commandName == "" || strings.ToLower(commandName) == "empty" || commandName == "'CommandName'" || commandName == "<parameter>" {
		commandName = "<JavaCommandName>"
	}

	*calls = append(*calls, TempMicroflowCall{
		ActivityName: activityName,
		Caption:      caption,
		TargetName:   javaActionName,
		CallType:     "JavaAction",
		AppName:      appName,
		CommandName:  commandName,
	})
}

func checkExternalAction(activity map[string]interface{}, calls *[]TempMicroflowCall, contentMap map[string]interface{}) {
	action, ok := activity["Action"].(map[string]interface{})
	if !ok {
		return
	}

	actionType, ok := action["$Type"].(string)
	if !ok || actionType != "Microflows$CallExternalAction" {
		return
	}

	// Get Name field (CommandName)
	externalActionName := ""
	if name, ok := action["Name"].(string); ok {
		externalActionName = name
	}

	if externalActionName == "" {
		return
	}

	// Extract details
	caption := ""
	if cap, ok := activity["Caption"].(string); ok {
		caption = cap
	}

	activityName := ""
	if actName, ok := activity["Name"].(string); ok {
		activityName = actName
	}

	// Extract AppName from ConsumedODataService (part after ".")
	appName := ""
	if consumedService, ok := action["ConsumedODataService"].(string); ok {
		parts := strings.Split(consumedService, ".")
		if len(parts) > 0 {
			appName = parts[len(parts)-1]
		}
	}

	*calls = append(*calls, TempMicroflowCall{
		ActivityName: activityName,
		Caption:      caption,
		TargetName:   externalActionName,
		CallType:     "ExternalAction",
		AppName:      appName,
		CommandName:  externalActionName,
	})
}

// ==== SECTION 3: Signal Manager Widgets Collection ====

func collectSignalManagerWidgets(reader *modelsdk.Reader, mprPath string, report *ManifestReport) error {
	widgetID := "siemens.mxtosignal.MxToSignal"
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	totalWidgets := 0
	pagesCount := 0
	snippetsCount := 0

	// Open database connection for module extraction
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	// ===== OPTIMIZATION: Pre-filter documents containing widget =====
	fmt.Printf("  🔍 Pre-filtering documents with Signal Manager widgets...\n")

	// Get all pages
	pages, err := reader.ListPages()
	if err != nil {
		return fmt.Errorf("failed to list pages: %w", err)
	}

	// Pre-filter pages that contain the widget ID
	var filteredPages []string
	for _, page := range pages {
		if containsWidgetID(contentsDir, string(page.ID), widgetID) {
			filteredPages = append(filteredPages, string(page.ID))
		}
	}

	// Get all snippets
	snippets, err := reader.ListSnippets()
	if err != nil {
		return fmt.Errorf("failed to list snippets: %w", err)
	}

	// Pre-filter snippets that contain the widget ID
	var filteredSnippets []string
	for _, snippet := range snippets {
		if containsWidgetID(contentsDir, string(snippet.ID), widgetID) {
			filteredSnippets = append(filteredSnippets, string(snippet.ID))
		}
	}

	totalDocs := len(pages) + len(snippets)
	filteredDocs := len(filteredPages) + len(filteredSnippets)
	fmt.Printf("  📊 Found %d documents to scan (filtered from %d total, %.1f%% reduction)\n",
		filteredDocs, totalDocs, float64(totalDocs-filteredDocs)/float64(totalDocs)*100)

	// Process filtered pages
	if len(filteredPages) > 0 {
		fmt.Printf("  📄 Scanning %d filtered page(s)...\n", len(filteredPages))
	}

	for _, pageID := range filteredPages {
		// Load page BSON (we already know it contains the widget)
		pageContent, err := loadUnitContents(contentsDir, pageID)
		if err != nil {
			continue
		}

		// Extract module name and document name from BSON
		moduleName, docName := extractModuleAndNameFromBSON(pageContent, pageID, db)

		// Find widgets with the specified widgetId
		widgets := findWidgetsByWidgetID(pageContent, widgetID)

		for _, widget := range widgets {
			// Extract all signal subscriptions from this widget
			subscriptions := extractSignalSubscriptions(widget, docName, "Page", moduleName)

			for _, sub := range subscriptions {
				report.Widgets = append(report.Widgets, sub)
				totalWidgets++
				pagesCount++
			}
		}
	}

	// Process filtered snippets
	if len(filteredSnippets) > 0 {
		fmt.Printf("  📄 Scanning %d filtered snippet(s)...\n", len(filteredSnippets))
	}

	for _, snippetID := range filteredSnippets {
		// Load snippet BSON (we already know it contains the widget)
		snippetContent, err := loadUnitContents(contentsDir, snippetID)
		if err != nil {
			continue
		}

		// Extract module name and document name from BSON
		moduleName, docName := extractModuleAndNameFromBSON(snippetContent, snippetID, db)

		// Find widgets with the specified widgetId
		widgets := findWidgetsByWidgetID(snippetContent, widgetID)

		for _, widget := range widgets {
			// Extract all signal subscriptions from this widget
			subscriptions := extractSignalSubscriptions(widget, docName, "Snippet", moduleName)

			for _, sub := range subscriptions {
				report.Widgets = append(report.Widgets, sub)
				totalWidgets++
				snippetsCount++
			}
		}
	}

	fmt.Printf("\n  ✅ Completed: Analyzed %d filtered pages and %d filtered snippets\n", len(filteredPages), len(filteredSnippets))
	fmt.Printf("  📊 Result: %d signal subscription(s) found (%d in pages, %d in snippets)\n",
		totalWidgets, pagesCount, snippetsCount)
	return nil
}

// findWidgetsByWidgetID finds all custom widgets with the specified widgetId
func findWidgetsByWidgetID(data map[string]interface{}, widgetID string) []map[string]interface{} {
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
			// Check if it has the matching widgetId in the Type field
			if typeField, ok := v["Type"].(map[string]interface{}); ok {
				if wid, ok := typeField["WidgetId"].(string); ok && wid == widgetID {
					*matchingWidgets = append(*matchingWidgets, v)
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

// KeyValue stores a key-value pair for primitive values
type KeyValue struct {
	Key   string
	Value string
}

// extractSignalSubscriptions extracts all signal subscriptions from a widget
func extractSignalSubscriptions(widget map[string]interface{}, documentName, documentType, moduleName string) []WidgetInfo {
	var subscriptions []WidgetInfo

	// Collect all primitive values with their keys to understand structure
	values := make([]KeyValue, 0)
	collectPrimitiveValuesWithKeys(widget, "", &values)

	// Group values by subscription object index
	// Pattern: Object_Properties_4_Value_Objects_X_Properties_Y_Value (for PrimitiveValue)
	// Pattern: Object_Properties_4_Value_Objects_X_Properties_Y_Value_Expression (for Expression)
	subscriptionMap := make(map[int]map[string]string) // [objectIndex][fieldName]value

	for _, kv := range values {
		// Parse pattern: ...Objects_X_Properties_Y...
		var objectIdx, propIdx int
		if strings.Contains(kv.Key, "Objects_") && strings.Contains(kv.Key, "_Properties_") {
			// Try to extract indices
			parts := strings.Split(kv.Key, "_")
			for i, part := range parts {
				if part == "Objects" && i+1 < len(parts) {
					fmt.Sscanf(parts[i+1], "%d", &objectIdx)
				}
				if part == "Properties" && i+1 < len(parts) {
					fmt.Sscanf(parts[i+1], "%d", &propIdx)
				}
			}

			if _, ok := subscriptionMap[objectIdx]; !ok {
				subscriptionMap[objectIdx] = make(map[string]string)
			}

			// Determine field type based on property index and key pattern
			// Properties_1 = SignalName (PrimitiveValue)
			// Properties_2 = AppName (PrimitiveValue)
			// Properties_3 = SubscriptionFilter (Expression)
			// Properties_6 = ThrottleDuration (PrimitiveValue)
			// Properties_7 = Blocking (PrimitiveValue)
			if propIdx == 1 {
				subscriptionMap[objectIdx]["signalname"] = kv.Value
			} else if propIdx == 2 {
				subscriptionMap[objectIdx]["appname"] = kv.Value
			} else if propIdx == 3 && strings.Contains(kv.Key, "_Expression") {
				// SubscriptionFilter is Properties_3 with Expression
				subscriptionMap[objectIdx]["filter"] = kv.Value
			}
		}
	}

	// Convert map to slice of WidgetInfo
	for _, props := range subscriptionMap {
		sub := WidgetInfo{
			DocumentName:       documentName,
			DocumentType:       documentType,
			Module:             moduleName,
			SignalName:         props["signalname"],
			AppName:            props["appname"],
			SubscriptionFilter: props["filter"],
		}
		// Only add if at least one field is populated
		if sub.AppName != "" || sub.SignalName != "" || sub.SubscriptionFilter != "" {
			subscriptions = append(subscriptions, sub)
		}
	}

	return subscriptions
}

// collectPrimitiveValuesWithKeys recursively collects string values with their key paths
func collectPrimitiveValuesWithKeys(data interface{}, keyPath string, values *[]KeyValue) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a PrimitiveValue
		if primVal, ok := v["PrimitiveValue"]; ok {
			if strVal, ok := primVal.(string); ok && strVal != "" {
				*values = append(*values, KeyValue{Key: keyPath, Value: strVal})
			}
		}

		// Check for Expression field (used for SubscriptionFilter and other expression-based properties)
		if exprVal, ok := v["Expression"]; ok {
			if strVal, ok := exprVal.(string); ok && strVal != "" {
				// Clean up expression value (remove quotes and newlines)
				cleanVal := strings.Trim(strings.TrimSpace(strVal), "'\"")
				if cleanVal != "" {
					*values = append(*values, KeyValue{Key: keyPath + "_Expression", Value: cleanVal})
				}
			}
		}

		// Recursively search in all fields
		for key, value := range v {
			newPath := keyPath
			if newPath != "" {
				newPath += "_" + key
			} else {
				newPath = key
			}
			collectPrimitiveValuesWithKeys(value, newPath, values)
		}

	case []interface{}:
		// Recursively search in arrays
		for i, item := range v {
			newPath := fmt.Sprintf("%s_%d", keyPath, i)
			collectPrimitiveValuesWithKeys(item, newPath, values)
		}
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

// ==== Utility Functions ====

func extractModuleName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) > 0 {
		return parts[0]
	}
	return fullName
}

func extractNameFromContents(contents map[string]interface{}) string {
	if name, ok := contents["Name"].(string); ok {
		return name
	}
	return ""
}

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

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}

	// Windows GUID format: mixed-endian
	var formatted [16]byte

	// Part 1: 4 bytes (little-endian)
	formatted[0] = blob[3]
	formatted[1] = blob[2]
	formatted[2] = blob[1]
	formatted[3] = blob[0]

	// Part 2: 2 bytes (little-endian)
	formatted[4] = blob[5]
	formatted[5] = blob[4]

	// Part 3: 2 bytes (little-endian)
	formatted[6] = blob[7]
	formatted[7] = blob[6]

	// Part 4-5: 8 bytes (big-endian, copy as-is)
	copy(formatted[8:], blob[8:16])

	hexStr := hex.EncodeToString(formatted[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hexStr[0:8],
		hexStr[8:12],
		hexStr[12:16],
		hexStr[16:20],
		hexStr[20:32])
}

// uuidToBlob converts UUID string to Windows GUID blob format (mixed-endian)
func uuidToBlob(uuid string) []byte {
	// Remove dashes
	cleanUUID := strings.ReplaceAll(uuid, "-", "")

	// Decode hex string
	bytes, err := hex.DecodeString(cleanUUID)
	if err != nil || len(bytes) != 16 {
		return make([]byte, 16)
	}

	// Convert from UUID format to Windows GUID format (reverse byte order for first 3 parts)
	blob := make([]byte, 16)

	// Part 1: 4 bytes (reverse)
	blob[0] = bytes[3]
	blob[1] = bytes[2]
	blob[2] = bytes[1]
	blob[3] = bytes[0]

	// Part 2: 2 bytes (reverse)
	blob[4] = bytes[5]
	blob[5] = bytes[4]

	// Part 3: 2 bytes (reverse)
	blob[6] = bytes[7]
	blob[7] = bytes[6]

	// Part 4-5: 8 bytes (copy as-is)
	copy(blob[8:], bytes[8:16])

	return blob
}

func loadUnitContents(contentsDir string, unitID string) (map[string]interface{}, error) {
	// Remove dashes from UUID for directory structure
	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid unit ID: %s", unitID)
	}

	// Build path: mprcontents/{first2}/{next2}/{uuid}.mxunit
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read unit file: %w", err)
	}

	var content map[string]interface{}
	if err := bson.Unmarshal(data, &content); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	return content, nil
}

// containsWidgetID performs fast pre-filtering by checking if the raw BSON file
// contains the widget ID string without full parsing. This avoids expensive
// BSON unmarshaling for documents that don't contain the widget.
func containsWidgetID(contentsDir, unitID, widgetID string) bool {
	// Remove dashes from UUID for directory structure
	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return false
	}

	// Build path: mprcontents/{first2}/{next2}/{uuid}.mxunit
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	// Read file as raw bytes
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}

	// Fast byte search - check if widget ID appears anywhere in the file
	// This is much faster than unmarshaling BSON
	return bytes.Contains(data, []byte(widgetID))
}

// stringToWindowsGUID converts a UUID string to Windows GUID binary format
func stringToWindowsGUID(uuidStr string) []byte {
	cleaned := strings.ReplaceAll(uuidStr, "-", "")
	if len(cleaned) != 32 {
		return nil
	}

	bytes, err := hex.DecodeString(cleaned)
	if err != nil || len(bytes) != 16 {
		return nil
	}

	// Windows GUID encoding: swap bytes for first 3 groups
	// Bytes 0-3: little-endian (reverse)
	bytes[0], bytes[1], bytes[2], bytes[3] = bytes[3], bytes[2], bytes[1], bytes[0]
	// Bytes 4-5: little-endian (reverse)
	bytes[4], bytes[5] = bytes[5], bytes[4]
	// Bytes 6-7: little-endian (reverse)
	bytes[6], bytes[7] = bytes[7], bytes[6]
	// Bytes 8-15: big-endian (unchanged)

	return bytes
}

// extractModuleAndNameFromBSON extracts the module name and document name from page/snippet BSON content
func extractModuleAndNameFromBSON(content map[string]interface{}, unitID string, db *sql.DB) (module, name string) {
	// Try to get Name field
	if docName, ok := content["Name"].(string); ok && docName != "" {
		name = docName
	}

	// Try to get module from ContainmentName field first
	if containmentName, ok := content["ContainmentName"].(string); ok && containmentName != "" {
		// ContainmentName format: "ModuleName.PageName" or "ModuleName.Folder.PageName"
		parts := strings.Split(containmentName, ".")
		if len(parts) > 0 {
			module = parts[0]
		}
		// If name wasn't in Name field, try to get it from last part of ContainmentName
		if name == "" && len(parts) > 0 {
			name = parts[len(parts)-1]
		}
	}

	// If module not found, try AllowedModuleRoles array (works for pages)
	if module == "" {
		if allowedRoles, ok := content["AllowedModuleRoles"].(bson.A); ok && len(allowedRoles) > 1 {
			// AllowedModuleRoles format: [1, "ModuleName.RoleName", ...]
			if roleStr, ok := allowedRoles[1].(string); ok && roleStr != "" {
				parts := strings.Split(roleStr, ".")
				if len(parts) > 0 {
					module = parts[0]
				}
			}
		}
	}

	// Last resort for snippets: trace up the hierarchy to find the module
	if module == "" && db != nil {
		module = findModuleByTraversal(unitID, db)
	}

	return module, name
}

// findModuleByTraversal traces up the unit hierarchy to find the parent module
func findModuleByTraversal(unitID string, db *sql.DB) string {
	if db == nil {
		return ""
	}

	guidBytes := stringToWindowsGUID(unitID)
	if guidBytes == nil {
		return ""
	}

	// Build a map of all modules (GUID -> Module Name)
	moduleMap := make(map[string]string)

	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		return ""
	}
	defer rows.Close()

	// Get MPR path from database connection
	mprPath := extractMPRPath(db)
	if mprPath == "" {
		return ""
	}
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	for rows.Next() {
		var unitIDBytes []byte
		if err := rows.Scan(&unitIDBytes); err != nil {
			continue
		}

		unitGUID := guidToString(unitIDBytes)
		if unitGUID == "" {
			continue
		}

		// Read BSON to check if it's a module
		cleanID := strings.ReplaceAll(unitGUID, "-", "")
		dir1 := cleanID[0:2]
		dir2 := cleanID[2:4]
		filePath := filepath.Join(contentsDir, dir1, dir2, unitGUID+".mxunit")

		data, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		var content map[string]interface{}
		if err := bson.Unmarshal(data, &content); err != nil {
			continue
		}

		// Check if this is a module
		if typeVal, ok := content["$Type"].(string); ok && strings.Contains(typeVal, "Projects$Module") {
			if moduleName, ok := content["Name"].(string); ok && moduleName != "" {
				moduleMap[unitGUID] = moduleName
			}
		}
	}

	// Now trace up the hierarchy
	currentID := guidBytes
	for i := 0; i < 20; i++ {
		var parentID []byte
		err := db.QueryRow("SELECT ContainerID FROM Unit WHERE UnitID = ?", currentID).Scan(&parentID)
		if err != nil || len(parentID) == 0 {
			break
		}

		// Check if this parent is a module
		parentGUID := guidToString(parentID)
		if moduleName, found := moduleMap[parentGUID]; found {
			return moduleName
		}

		currentID = parentID
	}

	return ""
}

// extractMPRPath gets the MPR file path from the database connection
// This is a workaround since sql.DB doesn't expose the connection string
var globalMPRPath string

func extractMPRPath(db *sql.DB) string {
	return globalMPRPath
}

// guidToString converts Windows GUID bytes to UUID string format
func guidToString(guidBytes []byte) string {
	if len(guidBytes) != 16 {
		return ""
	}
	// Reverse the Windows GUID encoding
	b := make([]byte, 16)
	copy(b, guidBytes)
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15])
}

// parseAttribute extracts name and type from attribute string
// Format: "Name (Type)" or just "Name"
func parseAttribute(attr string) (name, attrType string) {
	// Check if attribute has type in parentheses
	if strings.Contains(attr, "(") && strings.Contains(attr, ")") {
		parts := strings.Split(attr, "(")
		if len(parts) == 2 {
			name = strings.TrimSpace(parts[0])
			attrType = strings.TrimSuffix(strings.TrimSpace(parts[1]), ")")
			return name, attrType
		}
	}
	// No type specified, return just the name
	return attr, "-"
}

// ==== SECTION 4: Navigation Items Collection ====

// findAllShowPagesInFlow recursively searches BSON structure for all ShowPage/ShowFormAction activities
// Returns a slice of qualified page names (Module.PageName format)
func findAllShowPagesInFlow(data interface{}) []string {
	var pages []string

	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a ShowPage or ShowFormAction
		if typeStr, ok := v["$Type"].(string); ok {
			if strings.Contains(typeStr, "ShowPageAction") || strings.Contains(typeStr, "ShowFormAction") {
				// Try to extract page from FormSettings.Form
				if formSettings, ok := v["FormSettings"].(map[string]interface{}); ok {
					if formName, ok := formSettings["Form"].(string); ok && formName != "" {
						// Keep qualified name (Module.PageName)
						pages = append(pages, formName)
					}
				}
			}
		}

		// Recursively search all fields
		for _, val := range v {
			childPages := findAllShowPagesInFlow(val)
			pages = append(pages, childPages...)
		}

	case primitive.A: // BSON array
		for i, item := range v {
			if i == 0 {
				continue // Skip count element
			}
			childPages := findAllShowPagesInFlow(item)
			pages = append(pages, childPages...)
		}

	case []interface{}: // Regular slice
		for _, item := range v {
			childPages := findAllShowPagesInFlow(item)
			pages = append(pages, childPages...)
		}
	}

	return pages
}

// findAllShowPagesInFlowRecursive searches for ShowPage actions recursively through called flows
func findAllShowPagesInFlowRecursive(db *sql.DB, contentsDir string, flowContent map[string]interface{}, visited map[string]bool) []string {
	var pages []string

	// Find direct ShowPage actions
	directPages := findAllShowPagesInFlow(flowContent)
	pages = append(pages, directPages...)

	// Find called microflows/nanoflows
	calledFlows := findCalledFlows(flowContent)

	// Recursively search in called flows
	for _, flowRef := range calledFlows {
		// Skip if already visited (avoid infinite loops)
		if visited[flowRef] {
			continue
		}
		visited[flowRef] = true

		// Load the called flow
		flowData := loadFlowByReference(db, contentsDir, flowRef)
		if flowData != nil {
			// Recursively search
			childPages := findAllShowPagesInFlowRecursive(db, contentsDir, flowData, visited)
			pages = append(pages, childPages...)
		}
	}

	return pages
}

// findCalledFlows extracts all microflow/nanoflow references from a flow
func findCalledFlows(data interface{}) []string {
	var flows []string

	switch v := data.(type) {
	case map[string]interface{}:
		if typeStr, ok := v["$Type"].(string); ok {
			// Check for MicroflowCall action
			if strings.Contains(typeStr, "MicroflowCallAction") {
				// MicroflowCall is a map, extract the microflow reference from it
				if microflowCallMap, ok := v["MicroflowCall"].(map[string]interface{}); ok {
					// Try to get "Microflow" field which should be the UUID/reference
					if microflowRef, ok := microflowCallMap["Microflow"].(string); ok && microflowRef != "" {
						flows = append(flows, microflowRef)
					}
				}
			}
			// Check for NanoflowCall action
			if strings.Contains(typeStr, "NanoflowCallAction") {
				// NanoflowCall is a map, extract the nanoflow reference from it
				if nanoflowCallMap, ok := v["NanoflowCall"].(map[string]interface{}); ok {
					// Try to get "Nanoflow" field which should be the UUID/reference
					if nanoflowRef, ok := nanoflowCallMap["Nanoflow"].(string); ok && nanoflowRef != "" {
						flows = append(flows, nanoflowRef)
					}
				}
			}
		}

		// Recursively search all fields
		for _, val := range v {
			childFlows := findCalledFlows(val)
			flows = append(flows, childFlows...)
		}

	case primitive.A: // BSON array
		for i, item := range v {
			if i == 0 {
				continue // Skip count element
			}
			childFlows := findCalledFlows(item)
			flows = append(flows, childFlows...)
		}

	case []interface{}: // Regular slice
		for _, item := range v {
			childFlows := findCalledFlows(item)
			flows = append(flows, childFlows...)
		}
	}

	return flows
}

// loadFlowByReference loads a microflow or nanoflow by its reference (UUID or qualified name)
func loadFlowByReference(db *sql.DB, contentsDir string, flowRef string) map[string]interface{} {
	// Try loading by UUID first
	if content, err := loadUnitContents(contentsDir, flowRef); err == nil {
		if typeName, ok := content["$Type"].(string); ok {
			if typeName == "Microflows$Microflow" || typeName == "Microflows$Nanoflow" {
				return content
			}
		}
	}

	// Try loading by qualified name
	microflows, err := listMicroflows(db, contentsDir)
	if err != nil {
		return nil
	}

	for _, mf := range microflows {
		fullName := mf.ModuleName + "." + mf.Name
		if fullName == flowRef || mf.Name == flowRef {
			return mf.Content
		}
	}

	return nil
}

// resolveTargetPage determines the actual page(s) opened by a navigation item
// For Page items: returns qualified page name (Module.PageName)
// For Microflow/Nanoflow items: loads the flow and searches for ShowPage activities
// Returns comma-separated list of qualified page names if multiple pages found
func resolveTargetPage(db *sql.DB, contentsDir string, itemType string, target string, module string) string {
	// For Page items, return the qualified name as-is
	if strings.Contains(itemType, "Page") || strings.Contains(itemType, "FormAction") {
		return target // Return qualified name (Module.PageName)
	}

	// For Microflow/Nanoflow items, load the flow and search for ShowPage activities
	if strings.Contains(itemType, "Microflow") || strings.Contains(itemType, "Nanoflow") {
		// Construct qualified name if we have module
		qualifiedName := target
		if module != "" && !strings.Contains(target, ".") {
			qualifiedName = module + "." + target
		}

		// Load all microflows/nanoflows
		microflows, err := listMicroflows(db, contentsDir)
		if err != nil {
			return "" // Failed to load microflows
		}

		// Find the matching microflow/nanoflow
		for _, mf := range microflows {
			// Try to match by qualified name or just name
			fullName := mf.ModuleName + "." + mf.Name
			if fullName == qualifiedName || mf.Name == target {
				// Search for all ShowPage activities recursively (follows called flows)
				visited := make(map[string]bool)
				visited[fullName] = true
				pages := findAllShowPagesInFlowRecursive(db, contentsDir, mf.Content, visited)
				if len(pages) > 0 {
					// Return comma-separated list of page names
					return strings.Join(pages, ", ")
				}
				break
			}
		}

		// Also check nanoflows by UUID (they use same structure as microflows)
		query := `SELECT UnitID FROM Unit`
		rows, err := db.Query(query)
		if err != nil {
			return ""
		}
		defer rows.Close()

		for rows.Next() {
			var unitID []byte
			if err := rows.Scan(&unitID); err != nil {
				continue
			}

			unitIDStr := blobToUUID(unitID)

			// Check if UUID matches target
			if unitIDStr == target {
				content, err := loadUnitContents(contentsDir, unitIDStr)
				if err != nil {
					continue
				}

				// Check if this is a Nanoflow
				if typeName, ok := content["$Type"].(string); ok && typeName == "Microflows$Nanoflow" {
					// Search for ShowPage activities recursively
					visited := make(map[string]bool)
					visited[unitIDStr] = true
					pages := findAllShowPagesInFlowRecursive(db, contentsDir, content, visited)
					if len(pages) > 0 {
						return strings.Join(pages, ", ")
					}
					break
				}
			}

			// Also try matching by name
			content, err := loadUnitContents(contentsDir, unitIDStr)
			if err != nil {
				continue
			}

			// Check if this is a Nanoflow
			if typeName, ok := content["$Type"].(string); ok && typeName == "Microflows$Nanoflow" {
				name := extractNameFromContents(content)
				if name == target || (module+"."+name) == qualifiedName {
					// Search for ShowPage activities recursively
					visited := make(map[string]bool)
					fullName := module + "." + name
					visited[fullName] = true
					pages := findAllShowPagesInFlowRecursive(db, contentsDir, content, visited)
					if len(pages) > 0 {
						return strings.Join(pages, ", ")
					}
					break
				}
			}
		}
	}

	return "" // No page found or unsupported item type
}

// loadPageByQualifiedName loads a page's BSON data by its qualified name (Module.PageName)
func loadPageByQualifiedName(db *sql.DB, contentsDir string, qualifiedName string) (map[string]interface{}, error) {
	// Parse module and page name
	parts := strings.Split(qualifiedName, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid qualified name format: %s", qualifiedName)
	}
	pageName := parts[1]

	// Query all units to find the page
	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Check if this is a Page with matching name
		if typeName, ok := content["$Type"].(string); ok && typeName == "Forms$Page" {
			name := extractNameFromContents(content)
			if name == pageName {
				return content, nil
			}
		}
	}

	return nil, fmt.Errorf("page not found: %s", qualifiedName)
}

// findRightPlaceholder finds the Right placeholder's widgets array in a page
func findRightPlaceholder(pageData map[string]interface{}) interface{} {
	// Navigate: FormCall -> Arguments -> find Parameter ending with ".Right" -> return Widgets
	if formCall, ok := pageData["FormCall"].(map[string]interface{}); ok {
		if args, ok := formCall["Arguments"].(primitive.A); ok {
			for i, arg := range args {
				if i == 0 {
					continue // Skip count
				}
				if argMap, ok := arg.(map[string]interface{}); ok {
					if argType, ok := argMap["$Type"].(string); ok && argType == "Forms$FormCallArgument" {
						if param, ok := argMap["Parameter"].(string); ok {
							if strings.HasSuffix(param, ".Right") {
								// Return the Widgets array
								return argMap["Widgets"]
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// loadSnippetContent loads snippet BSON from mprcontents folder
func loadSnippetContent(contentsDir, snippetID string) (map[string]interface{}, error) {
	// Remove dashes from UUID for directory structure
	cleanID := strings.ReplaceAll(snippetID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid snippet ID: %s", snippetID)
	}

	// Build path: mprcontents/{first2}/{next2}/{uuid}.mxunit
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, snippetID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read snippet file %s: %w", filePath, err)
	}

	// Parse BSON
	var snippetData map[string]interface{}
	if err := bson.Unmarshal(data, &snippetData); err != nil {
		return nil, fmt.Errorf("failed to parse snippet BSON: %w", err)
	}

	return snippetData, nil
}

// findSnippetInWidgets searches for snippet widget in widgets array and returns snippet ID
// Searches recursively through nested widgets
func findSnippetInWidgets(widgets interface{}) string {
	switch v := widgets.(type) {
	case map[string]interface{}:
		// Check if THIS is a snippet widget
		if typeStr, ok := v["$Type"].(string); ok {
			if strings.Contains(typeStr, "Snippet") {
				// Extract snippet reference - different locations depending on type
				if snippetRef, ok := v["Snippet"].(string); ok && snippetRef != "" {
					return snippetRef
				}
				// Forms$SnippetCall has Form field directly
				if form, ok := v["Form"].(string); ok && form != "" {
					return form
				}
				// Forms$SnippetCallWidget has SnippetCall sub-object with Form
				if snippetCall, ok := v["SnippetCall"].(map[string]interface{}); ok {
					if form, ok := snippetCall["Form"].(string); ok && form != "" {
						return form
					}
				}
			}
		}
		// Recursively search nested fields
		for key, val := range v {
			if key == "$ID" || key == "$Type" {
				continue
			}
			if result := findSnippetInWidgets(val); result != "" {
				return result
			}
		}
	case primitive.A:
		for i, item := range v {
			if i == 0 {
				continue // Skip count
			}
			if result := findSnippetInWidgets(item); result != "" {
				return result
			}
		}
	case []interface{}:
		for _, item := range v {
			if result := findSnippetInWidgets(item); result != "" {
				return result
			}
		}
	}
	return ""
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// resolveRightPlaceholderWidgets checks if Right placeholder contains a snippet and loads it
// Returns (widgets, snippetLoaded) where snippetLoaded indicates if a snippet was successfully resolved
func resolveRightPlaceholderWidgets(rightWidgets interface{}, contentsDir string, db *sql.DB, mprPath string) (interface{}, bool) {
	if rightWidgets == nil {
		return nil, false
	}

	// Check if there's a snippet in the widgets
	snippetQName := findSnippetInWidgets(rightWidgets)
	if snippetQName == "" {
		// No snippet found, return original widgets
		return rightWidgets, false
	}

	// Convert qualified name to UUID using reader
	snippetUUID, err := getSnippetUUIDByQualifiedName(mprPath, snippetQName)
	if err != nil {
		return rightWidgets, false
	}

	// Load snippet content
	snippetData, err := loadSnippetContent(contentsDir, snippetUUID)
	if err != nil {
		return rightWidgets, false
	}

	// Extract widgets from snippet - they should be in the Widget or Widgets field
	if widgets, ok := snippetData["Widget"]; ok {
		return widgets, true
	}
	if widgets, ok := snippetData["Widgets"]; ok {
		return widgets, true
	}

	// Snippet loaded but no widgets found, return original
	return rightWidgets, false
}

// getSnippetUUIDByQualifiedName retrieves the UUID of a snippet by its qualified name using reader
func getSnippetUUIDByQualifiedName(mprPath, qualifiedName string) (string, error) {
	// Open reader to get snippets
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		return "", fmt.Errorf("failed to open MPR: %w", err)
	}
	defer reader.Close()

	// List all snippets
	snippets, err := reader.ListSnippets()
	if err != nil {
		return "", fmt.Errorf("failed to list snippets: %w", err)
	}

	// Extract just the snippet name from qualified name (Module.SnippetName -> SnippetName)
	parts := strings.Split(qualifiedName, ".")
	snippetName := qualifiedName
	if len(parts) == 2 {
		snippetName = parts[1]
	}

	// Find snippet by name
	for _, snippet := range snippets {
		if snippet.Name == snippetName {
			return string(snippet.ID), nil
		}
	}

	return "", fmt.Errorf("snippet not found: %s (tried name: %s)", qualifiedName, snippetName)
}

// findFirstCommandBarContainer recursively searches for the first DivContainer with command bar class
// Supports various class name patterns: vertical-command-bar, verticalCommandBar, command-bar, commandBar, etc.
func findFirstCommandBarContainer(data interface{}) map[string]interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if this is a DivContainer with command bar class
		if typeStr, ok := v["$Type"].(string); ok && typeStr == "Forms$DivContainer" {
			if appearance, ok := v["Appearance"].(map[string]interface{}); ok {
				if class, ok := appearance["Class"].(string); ok {
					classLower := strings.ToLower(class)
					// Match various command bar class patterns
					if strings.Contains(classLower, "vertical-command-bar") ||
						strings.Contains(classLower, "verticalcommandbar") ||
						(strings.Contains(classLower, "vertical") && strings.Contains(classLower, "command")) {
						return v // Found it!
					}
				}
			}
		}

		// Recursively search all fields
		for _, val := range v {
			if result := findFirstCommandBarContainer(val); result != nil {
				return result
			}
		}

	case primitive.A: // BSON array
		for i, item := range v {
			if i == 0 {
				continue // Skip count
			}
			if result := findFirstCommandBarContainer(item); result != nil {
				return result
			}
		}

	case []interface{}: // Regular slice
		for _, item := range v {
			if result := findFirstCommandBarContainer(item); result != nil {
				return result
			}
		}
	}

	return nil
}

// extractButtonCaption extracts caption from an ActionButton widget
func extractButtonCaption(button map[string]interface{}) string {
	// Try CaptionTemplate.Template first (ActionButton standard)
	if captionTemplate, ok := button["CaptionTemplate"].(map[string]interface{}); ok {
		if template, ok := captionTemplate["Template"].(map[string]interface{}); ok {
			// Extract text from Texts$Text structure - look for Items array
			if items, ok := template["Items"].(primitive.A); ok {
				for i, item := range items {
					if i == 0 {
						continue // Skip count
					}
					if itemMap, ok := item.(map[string]interface{}); ok {
						// Check if this is a Translation with Text
						if text, ok := itemMap["Text"].(string); ok && text != "" {
							return text
						}
					}
				}
			}
		}
	}

	// Try direct Caption field
	if caption, ok := button["Caption"]; ok {
		if captionStr, ok := caption.(string); ok {
			return captionStr
		}
		// Caption might be a Texts$Text structure
		if captionMap, ok := caption.(map[string]interface{}); ok {
			if items, ok := captionMap["Items"].(primitive.A); ok {
				for i, item := range items {
					if i == 0 {
						continue
					}
					if itemMap, ok := item.(map[string]interface{}); ok {
						if text, ok := itemMap["Text"].(string); ok && text != "" {
							return text
						}
					}
				}
			}
		}
	}

	// Try Tooltip as fallback (for icon-only buttons)
	if tooltip, ok := button["Tooltip"].(map[string]interface{}); ok {
		if items, ok := tooltip["Items"].(primitive.A); ok {
			for i, item := range items {
				if i == 0 {
					continue
				}
				if itemMap, ok := item.(map[string]interface{}); ok {
					if text, ok := itemMap["Text"].(string); ok && text != "" {
						return text
					}
				}
			}
		}
	}

	return ""
}

// extractActionButtons finds all ActionButton widgets in a container and extracts their details
// Note: In Mendix, OnClickAction is on the parent DivContainer, and Caption is in sibling DynamicText
func extractActionButtons(container map[string]interface{}, db *sql.DB, contentsDir string) []PageCommandButton {
	var buttons []PageCommandButton

	// Recursive search that only follows Widgets field, not all fields
	var searchWidgets func(data interface{}, parentAction map[string]interface{})
	searchWidgets = func(data interface{}, parentAction map[string]interface{}) {
		switch v := data.(type) {
		case map[string]interface{}:
			typeStr, _ := v["$Type"].(string)

			// Capture OnClickAction from DivContainer to pass to children
			currentAction := parentAction
			if typeStr == "Forms$DivContainer" {
				if onClickAction, ok := v["OnClickAction"].(map[string]interface{}); ok {
					currentAction = onClickAction
				}

				// When we find a DivContainer, check its Widgets for ActionButton + DynamicText pairs
				if widgets, ok := v["Widgets"].(primitive.A); ok {
					var actionButton map[string]interface{}
					var dynamicText map[string]interface{}

					for i, widget := range widgets {
						if i == 0 {
							continue // Skip count
						}
						if widgetMap, ok := widget.(map[string]interface{}); ok {
							widgetType, _ := widgetMap["$Type"].(string)
							if widgetType == "Forms$ActionButton" {
								actionButton = widgetMap
							} else if widgetType == "Forms$DynamicText" {
								dynamicText = widgetMap
							}
						}
					}

					// If we found both, extract button with caption from DynamicText
					if actionButton != nil && currentAction != nil {
						button := PageCommandButton{
							ButtonName: extractNameFromContents(actionButton),
							Caption:    extractCaptionFromDynamicText(dynamicText),
						}

						// Extract action information
						if actionType, ok := currentAction["$Type"].(string); ok {
							button.ActionType = actionType

							if strings.Contains(actionType, "CallNanoflowClientAction") {
								if nanoflow, ok := currentAction["Nanoflow"].(string); ok {
									button.ActionName = nanoflow
									nanoflowPages := loadNanoflowShowPages(db, contentsDir, nanoflow)
									if len(nanoflowPages) > 0 {
										button.TargetPage = strings.Join(nanoflowPages, ", ")
									}
								}
							} else if strings.Contains(actionType, "CallMicroflowClientAction") {
								if microflow, ok := currentAction["Microflow"].(string); ok {
									button.ActionName = microflow
									microflowPages := loadMicroflowShowPages(db, contentsDir, microflow)
									if len(microflowPages) > 0 {
										button.TargetPage = strings.Join(microflowPages, ", ")
									}
								}
							} else if strings.Contains(actionType, "ShowPage") {
								if pageSettings, ok := currentAction["PageSettings"].(map[string]interface{}); ok {
									if page, ok := pageSettings["Page"].(string); ok {
										button.TargetPage = page
									}
								}
							} else if strings.Contains(actionType, "CreateObjectClientAction") {
								button.ActionName = "Create Object"
								if pageSettings, ok := currentAction["PageSettings"].(map[string]interface{}); ok {
									if form, ok := pageSettings["Form"].(string); ok {
										button.TargetPage = form
									}
								}
							} else if strings.Contains(actionType, "FormAction") {
								if formSettings, ok := currentAction["FormSettings"].(map[string]interface{}); ok {
									if form, ok := formSettings["Form"].(string); ok {
										button.TargetPage = form
									}
								}
							}
						}

						buttons = append(buttons, button)
					}

					// Continue recursively searching in widgets
					for i, widget := range widgets {
						if i == 0 {
							continue
						}
						searchWidgets(widget, currentAction)
					}
				}
			} else if typeStr == "Forms$DataView" {
				// DataView can contain nested DivContainers, search in its Widgets
				if widgets, ok := v["Widgets"].(primitive.A); ok {
					for i, widget := range widgets {
						if i == 0 {
							continue
						}
						searchWidgets(widget, parentAction)
					}
				}
			}

		case primitive.A:
			for i, item := range v {
				if i == 0 {
					continue
				}
				searchWidgets(item, parentAction)
			}
		}
	}

	searchWidgets(container, nil)
	return buttons
}

// extractCaptionFromDynamicText extracts the caption from a DynamicText widget (en_US language)
func extractCaptionFromDynamicText(dynamicText map[string]interface{}) string {
	if dynamicText == nil {
		return ""
	}

	// Navigate: Content -> Template -> Items -> find en_US translation
	if content, ok := dynamicText["Content"].(map[string]interface{}); ok {
		if template, ok := content["Template"].(map[string]interface{}); ok {
			if items, ok := template["Items"].(primitive.A); ok {
				for i, item := range items {
					if i == 0 {
						continue // Skip count
					}
					if itemMap, ok := item.(map[string]interface{}); ok {
						if langCode, ok := itemMap["LanguageCode"].(string); ok && langCode == "en_US" {
							if text, ok := itemMap["Text"].(string); ok && text != "" {
								return text
							}
						}
					}
				}

				// Fallback to first non-empty text if en_US not found
				for i, item := range items {
					if i == 0 {
						continue
					}
					if itemMap, ok := item.(map[string]interface{}); ok {
						if text, ok := itemMap["Text"].(string); ok && text != "" {
							return text
						}
					}
				}
			}
		}
	}

	return ""
}

// findCommandByButtonHeuristic extracts command name from button caption + page entity name
// Example: Hide button on StateMachine_Master -> looks for HideStateMachine command
func findCommandByButtonHeuristic(pageQualifiedName string, buttonCaption string, allMicroflowCalls []MicroflowCallInfo) string {
	// Extract page name without module prefix
	pageName := pageQualifiedName
	if strings.Contains(pageName, ".") {
		parts := strings.Split(pageName, ".")
		pageName = parts[len(parts)-1]
	}

	// Extract entity name from page name (e.g., "StateMachine_Master" -> "StateMachine")
	entityName := pageName
	suffixes := []string{"_Master", "_Details", "_Edit", "_New", "_SingleSelection"}

	// Remove all known suffixes (some pages have multiple, e.g., "_Master_SingleSelection")
	changed := true
	for changed {
		changed = false
		for _, suffix := range suffixes {
			if strings.HasSuffix(entityName, suffix) {
				entityName = strings.TrimSuffix(entityName, suffix)
				changed = true
				break
			}
		}
	}

	if entityName == "" || buttonCaption == "" {
		return ""
	}

	// Build expected command name (e.g., "Hide" + "StateMachine" = "HideStateMachine")
	expectedCommandName := buttonCaption + entityName

	// Strategy 1: Try exact match
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			if call.CommandName == expectedCommandName {
				return call.AppName + "." + call.CommandName
			}
		}
	}

	// Strategy 2: Try case-insensitive match
	expectedLower := strings.ToLower(expectedCommandName)
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			if strings.ToLower(call.CommandName) == expectedLower {
				return call.AppName + "." + call.CommandName
			}
		}
	}

	// Strategy 3: Try with "Set" prefix (e.g., "Set" + "Initial" + "Status" = "SetStatusAsInitial" or "SetInitialStatus")
	// Try both orders: SetCaptionEntity and SetEntityCaption
	patterns := []string{
		"Set" + buttonCaption + entityName,        // SetInitialStatus
		"Set" + entityName + buttonCaption,        // SetStatusInitial
		"Set" + entityName + "As" + buttonCaption, // SetStatusAsInitial
	}

	for _, pattern := range patterns {
		patternLower := strings.ToLower(pattern)
		for _, call := range allMicroflowCalls {
			if call.CallType == "ExternalAction" {
				if strings.ToLower(call.CommandName) == patternLower {
					return call.AppName + "." + call.CommandName
				}
			}
		}
	}

	// Strategy 4: Try commands that contain both caption and entity name
	captionLower := strings.ToLower(buttonCaption)
	entityLower := strings.ToLower(entityName)
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			commandLower := strings.ToLower(call.CommandName)
			if strings.Contains(commandLower, captionLower) && strings.Contains(commandLower, entityLower) {
				return call.AppName + "." + call.CommandName
			}
		}
	}

	return ""
}

// findSaveButtonCommand finds the "Save" button in a page and extracts the command it calls
// Uses multiple strategies: direct button search, flow analysis, and heuristic name matching
func findSaveButtonCommand(db *sql.DB, contentsDir string, pageQualifiedName string, allMicroflowCalls []MicroflowCallInfo) string {
	// Load the target page
	pageData, err := loadPageByQualifiedName(db, contentsDir, pageQualifiedName)
	if err != nil || pageData == nil {
		// If page loading fails, try heuristic approach
		return findCommandByPageNameHeuristic(pageQualifiedName, allMicroflowCalls)
	}

	// First try to find "Save" button specifically
	saveButtonFlow := findSaveButtonFlow(pageData)
	if saveButtonFlow != "" {
		command := extractCommandFromFlow(db, contentsDir, saveButtonFlow)
		if command != "" {
			return command
		}
	}

	// If no Save button found, try to find any flow that might contain a command
	// This handles PANEL pages that might have different button structures
	flows := findAllFlowsInPage(pageData)
	for _, flow := range flows {
		command := extractCommandFromFlow(db, contentsDir, flow)
		if command != "" {
			return command
		}
	}

	// If no command found through flow analysis, try heuristic approach
	return findCommandByPageNameHeuristic(pageQualifiedName, allMicroflowCalls)
}

// findCommandByPageNameHeuristic extracts the command name from page name patterns
// Example: PANEL_CreateStateMachine -> looks for CreateStateMachine in microflow calls
func findCommandByPageNameHeuristic(pageQualifiedName string, allMicroflowCalls []MicroflowCallInfo) string {
	// Extract page name without module prefix
	pageName := pageQualifiedName
	if strings.Contains(pageName, ".") {
		parts := strings.Split(pageName, ".")
		pageName = parts[len(parts)-1]
	}

	// Common patterns: PANEL_CreateXXX, PANEL_UpdateXXX, etc.
	patterns := []string{
		"PANEL_Create",
		"PANEL_Update",
		"Panel_Create",
		"Panel_Update",
		"_Create",
		"_Update",
	}

	var actionName string
	var operation string // "Create" or "Update"

	for _, pattern := range patterns {
		if strings.Contains(pageName, pattern) {
			actionName = strings.Replace(pageName, pattern, "", 1)
			if strings.HasSuffix(pattern, "Create") {
				operation = "Create"
			} else if strings.HasSuffix(pattern, "Update") {
				operation = "Update"
			}
			break
		}
	}

	if actionName == "" {
		return ""
	}

	// Strip "Base" prefix if present (e.g., "BaseUoMDimension" -> "UoMDimension")
	alternativeActionName := actionName
	if strings.HasPrefix(actionName, "Base") {
		alternativeActionName = strings.TrimPrefix(actionName, "Base")
	}

	// Strategy 1: Try exact match with operation prefix (e.g., "CreateStateMachine")
	exactMatch := operation + actionName
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			if call.CommandName == exactMatch {
				return call.AppName + "." + call.CommandName
			}
		}
	}

	// Strategy 1b: Try exact match with operation prefix and alternative name (e.g., "CreateUoMDimension")
	if alternativeActionName != actionName {
		exactMatchAlt := operation + alternativeActionName
		for _, call := range allMicroflowCalls {
			if call.CallType == "ExternalAction" {
				if call.CommandName == exactMatchAlt {
					return call.AppName + "." + call.CommandName
				}
			}
		}
	}

	// Strategy 2: Try exact match with just the action name (e.g., "StateMachine" -> "StateMachine" command)
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			if call.CommandName == actionName {
				return call.AppName + "." + call.CommandName
			}
		}
	}

	// Strategy 3: Try match where command name ends with action name (e.g., "CreateStateMachine", "UpdateStateMachine")
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			if strings.HasSuffix(call.CommandName, actionName) {
				// Prefer commands that start with the same operation
				if strings.HasPrefix(call.CommandName, operation) {
					return call.AppName + "." + call.CommandName
				}
			}
		}
	}

	// Strategy 4: Try match where command name contains action name
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			if strings.Contains(call.CommandName, actionName) {
				// Prefer commands that start with the same operation
				if strings.HasPrefix(call.CommandName, operation) {
					return call.AppName + "." + call.CommandName
				}
			}
		}
	}

	// Strategy 5: Last resort - any command containing action name
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" {
			if strings.Contains(call.CommandName, actionName) {
				return call.AppName + "." + call.CommandName
			}
		}
	}

	return ""
}

// findSaveButtonFlow recursively searches for a button with caption "Save" and returns its flow name
func findSaveButtonFlow(pageData map[string]interface{}) string {
	var search func(data interface{}, parentAction map[string]interface{}) string
	search = func(data interface{}, parentAction map[string]interface{}) string {
		switch v := data.(type) {
		case map[string]interface{}:
			typeStr, _ := v["$Type"].(string)

			// Capture OnClickAction from DivContainer
			currentAction := parentAction
			if typeStr == "Forms$DivContainer" {
				if onClickAction, ok := v["OnClickAction"].(map[string]interface{}); ok {
					currentAction = onClickAction
				}

				// Check Widgets for ActionButton + DynamicText pairs
				if widgets, ok := v["Widgets"].(primitive.A); ok {
					var actionButton map[string]interface{}
					var dynamicText map[string]interface{}

					for i, widget := range widgets {
						if i == 0 {
							continue
						}
						if widgetMap, ok := widget.(map[string]interface{}); ok {
							widgetType, _ := widgetMap["$Type"].(string)
							if widgetType == "Forms$ActionButton" {
								actionButton = widgetMap
							} else if widgetType == "Forms$DynamicText" {
								dynamicText = widgetMap
							}
						}
					}

					// Check if this is a Save button
					if actionButton != nil && dynamicText != nil {
						caption := extractCaptionFromDynamicText(dynamicText)
						if caption == "Save" {
							// Get the flow name from inherited or own action
							actionToUse := currentAction
							if onClickAction, ok := actionButton["OnClickAction"].(map[string]interface{}); ok {
								actionToUse = onClickAction
							}

							if actionToUse != nil {
								if actionType, ok := actionToUse["$Type"].(string); ok {
									if strings.Contains(actionType, "CallNanoflowClientAction") {
										if nanoflow, ok := actionToUse["Nanoflow"].(string); ok {
											return nanoflow
										}
									} else if strings.Contains(actionType, "CallMicroflowClientAction") {
										if microflow, ok := actionToUse["Microflow"].(string); ok {
											return microflow
										}
									}
								}
							}
						}
					}
				}
			}

			// Recursively search
			for _, val := range v {
				if result := search(val, currentAction); result != "" {
					return result
				}
			}

		case primitive.A:
			for i, item := range v {
				if i == 0 {
					continue
				}
				if result := search(item, parentAction); result != "" {
					return result
				}
			}

		case []interface{}:
			for _, item := range v {
				if result := search(item, parentAction); result != "" {
					return result
				}
			}
		}
		return ""
	}

	return search(pageData, nil)
}

// findAllFlowsInPage finds all nanoflow/microflow references in a page
func findAllFlowsInPage(pageData map[string]interface{}) []string {
	var flows []string
	seen := make(map[string]bool)

	var search func(data interface{})
	search = func(data interface{}) {
		switch v := data.(type) {
		case map[string]interface{}:
			// Check for OnClickAction with nanoflow/microflow
			if onClickAction, ok := v["OnClickAction"].(map[string]interface{}); ok {
				if actionType, ok := onClickAction["$Type"].(string); ok {
					if strings.Contains(actionType, "CallNanoflowClientAction") {
						if nanoflow, ok := onClickAction["Nanoflow"].(string); ok && !seen[nanoflow] {
							flows = append(flows, nanoflow)
							seen[nanoflow] = true
						}
					} else if strings.Contains(actionType, "CallMicroflowClientAction") {
						if microflow, ok := onClickAction["Microflow"].(string); ok && !seen[microflow] {
							flows = append(flows, microflow)
							seen[microflow] = true
						}
					}
				}
			}

			// Recursively search
			for _, val := range v {
				search(val)
			}

		case primitive.A:
			for i, item := range v {
				if i == 0 {
					continue
				}
				search(item)
			}

		case []interface{}:
			for _, item := range v {
				search(item)
			}
		}
	}

	search(pageData)
	return flows
}

// extractCommandFromFlow loads a flow and extracts CommandName from MicroflowCall, ExternalAction, or JavaAction
func extractCommandFromFlow(db *sql.DB, contentsDir string, flowName string) string {
	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		data, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Check if this is a Nanoflow or Microflow
		typeStr, _ := data["$Type"].(string)
		if typeStr != "Microflows$Nanoflow" && typeStr != "Microflows$Microflow" {
			continue
		}

		// Check if this is the flow we're looking for (match by QualifiedName or Name)
		matched := false
		if qualifiedName, ok := data["QualifiedName"].(string); ok {
			if qualifiedName == flowName {
				matched = true
			}
		}
		if !matched {
			if name, ok := data["Name"].(string); ok {
				// Match by simple name or if flowName ends with .Name
				if name == flowName || strings.HasSuffix(flowName, "."+name) {
					matched = true
				}
			}
		}

		if matched {
			// Search for command calls in the flow
			return searchFlowForCommand(data)
		}
	}

	return ""
}

// searchFlowForCommand recursively searches for MicroflowCall, ExternalAction, or JavaAction and extracts CommandName
func searchFlowForCommand(data interface{}) string {
	switch v := data.(type) {
	case map[string]interface{}:
		if typeStr, ok := v["$Type"].(string); ok {
			// Check for MicroflowCall
			if typeStr == "Microflows$MicroflowCall" {
				if microflow, ok := v["Microflow"].(string); ok {
					// Check if it contains "Command" in the name
					if strings.Contains(microflow, "Command") || strings.Contains(microflow, "CallCommand") {
						// Try to extract CommandName parameter
						if params, ok := v["ParameterMappings"].(primitive.A); ok {
							for i, param := range params {
								if i == 0 {
									continue
								}
								if paramMap, ok := param.(map[string]interface{}); ok {
									if argValue, ok := paramMap["ArgumentValue"].(map[string]interface{}); ok {
										if constant, ok := argValue["Constant"].(string); ok {
											return constant
										}
									}
								}
							}
						}
						return microflow // Fallback to microflow name
					}
				}
			}

			// Check for ExternalAction (OData calls)
			if typeStr == "Microflows$ExternalAction" {
				if appName, ok := v["AppName"].(string); ok {
					if commandName, ok := v["CommandName"].(string); ok {
						return appName + "." + commandName
					}
				}
			}

			// Check for JavaAction
			if typeStr == "Microflows$JavaAction" {
				if javaAction, ok := v["JavaAction"].(string); ok {
					if strings.Contains(javaAction, "Command") {
						// Try to extract CommandName from parameters
						if params, ok := v["ParameterMappings"].(primitive.A); ok {
							for i, param := range params {
								if i == 0 {
									continue
								}
								if paramMap, ok := param.(map[string]interface{}); ok {
									if argValue, ok := paramMap["ArgumentValue"].(map[string]interface{}); ok {
										if constant, ok := argValue["Constant"].(string); ok {
											return constant
										}
									}
								}
							}
						}
						return javaAction
					}
				}
			}
		}

		// Recursively search all fields
		for _, val := range v {
			if result := searchFlowForCommand(val); result != "" {
				return result
			}
		}

	case primitive.A:
		for i, item := range v {
			if i == 0 {
				continue
			}
			if result := searchFlowForCommand(item); result != "" {
				return result
			}
		}

	case []interface{}:
		for _, item := range v {
			if result := searchFlowForCommand(item); result != "" {
				return result
			}
		}
	}

	return ""
}

// loadNanoflowShowPages loads a nanoflow and finds all ShowPage actions
func loadNanoflowShowPages(db *sql.DB, contentsDir string, nanoflowName string) []string {
	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Check if this is a Nanoflow with matching name
		if typeName, ok := content["$Type"].(string); ok && typeName == "Microflows$Nanoflow" {
			name := extractNameFromContents(content)
			// Match by simple name or qualified name
			if name == nanoflowName || strings.HasSuffix(nanoflowName, "."+name) {
				visited := make(map[string]bool)
				visited[nanoflowName] = true
				return findAllShowPagesInFlowRecursive(db, contentsDir, content, visited)
			}
		}
	}

	return nil
}

// loadMicroflowShowPages loads a microflow and finds all ShowPage actions
func loadMicroflowShowPages(db *sql.DB, contentsDir string, microflowName string) []string {
	microflows, err := listMicroflows(db, contentsDir)
	if err != nil {
		return nil
	}

	for _, mf := range microflows {
		// Match by simple name or qualified name
		if mf.Name == microflowName || strings.HasSuffix(microflowName, "."+mf.Name) {
			visited := make(map[string]bool)
			fullName := mf.ModuleName + "." + mf.Name
			visited[fullName] = true
			return findAllShowPagesInFlowRecursive(db, contentsDir, mf.Content, visited)
		}
	}

	return nil
}

// collectPageCommands extracts command bar actions from pages listed in navigation items
func collectPageCommands(db *sql.DB, contentsDir string, mprPath string, navigationItems []NavigationItem, allMicroflowCalls []MicroflowCallInfo) ([]PageCommandInfo, error) {
	// Get unique list of target pages from navigation items (filter out "-" and empty strings)
	uniquePages := make(map[string]bool)
	for _, item := range navigationItems {
		if item.TargetPage != "" && item.TargetPage != "-" {
			uniquePages[item.TargetPage] = true
		}
	}

	fmt.Printf("  📄 Found %d unique page(s) to analyze\n", len(uniquePages))

	var pageCommands []PageCommandInfo
	analyzed := 0
	skipped := 0

	for pageName := range uniquePages {
		// Load page BSON
		pageData, err := loadPageByQualifiedName(db, contentsDir, pageName)
		if err != nil {
			// Skip pages that can't be loaded
			skipped++
			continue
		}

		// Find Right placeholder
		rightWidgets := findRightPlaceholder(pageData)
		if rightWidgets == nil {
			skipped++
			continue
		}

		// Resolve snippet if present in Right placeholder
		resolvedWidgets, _ := resolveRightPlaceholderWidgets(rightWidgets, contentsDir, db, mprPath)

		// Find first command bar container
		commandBarContainer := findFirstCommandBarContainer(resolvedWidgets)
		if commandBarContainer == nil {
			skipped++
			continue
		}

		// Extract action buttons
		buttons := extractActionButtons(commandBarContainer, db, contentsDir)
		if len(buttons) > 0 {
			// For each button, find the command it calls
			for i := range buttons {
				// Strategy 1: If button opens a target page (PANEL_*), find command in that page
				if buttons[i].TargetPage != "" && buttons[i].TargetPage != "-" {
					buttons[i].TargetCommand = findSaveButtonCommand(db, contentsDir, buttons[i].TargetPage, allMicroflowCalls)
					// Extract AppName and CommandName from TargetCommand (format: "AppName.CommandName")
					if buttons[i].TargetCommand != "" {
						if idx := strings.LastIndex(buttons[i].TargetCommand, "."); idx >= 0 {
							buttons[i].TargetAppName = buttons[i].TargetCommand[:idx]
							buttons[i].TargetCommand = buttons[i].TargetCommand[idx+1:] // Keep only CommandName
						}
					}
				} else if buttons[i].ActionName != "" {
					// Strategy 2: If button calls a nanoflow/microflow directly, try to find command in that flow
					if strings.Contains(buttons[i].ActionType, "CallNanoflowClientAction") ||
						strings.Contains(buttons[i].ActionType, "CallMicroflowClientAction") {
						buttons[i].TargetCommand = extractCommandFromFlow(db, contentsDir, buttons[i].ActionName)
						// Extract AppName and CommandName from TargetCommand
						if buttons[i].TargetCommand != "" {
							if idx := strings.LastIndex(buttons[i].TargetCommand, "."); idx >= 0 {
								buttons[i].TargetAppName = buttons[i].TargetCommand[:idx]
								buttons[i].TargetCommand = buttons[i].TargetCommand[idx+1:] // Keep only CommandName
							}
						}

						// Strategy 3: If no command found in flow, use heuristic based on button caption + page entity
						if buttons[i].TargetCommand == "" {
							buttons[i].TargetCommand = findCommandByButtonHeuristic(pageName, buttons[i].Caption, allMicroflowCalls)
							// Extract AppName and CommandName from TargetCommand
							if buttons[i].TargetCommand != "" {
								if idx := strings.LastIndex(buttons[i].TargetCommand, "."); idx >= 0 {
									buttons[i].TargetAppName = buttons[i].TargetCommand[:idx]
									buttons[i].TargetCommand = buttons[i].TargetCommand[idx+1:] // Keep only CommandName
								}
							}
						}
					}
				}
			}

			pageCommands = append(pageCommands, PageCommandInfo{
				PageName: pageName,
				Commands: buttons,
			})
			analyzed++
		} else {
			skipped++
		}
	}

	fmt.Printf("  ✓ Analyzed %d page(s) with commands, skipped %d page(s)\n", analyzed, skipped)

	// Validate that all Target Commands exist in allMicroflowCalls
	validateTargetCommands(pageCommands, allMicroflowCalls)

	return pageCommands, nil
}

// validateTargetCommands checks that all Target Commands are present in the Microflow/Action Calls list
func validateTargetCommands(pageCommands []PageCommandInfo, allMicroflowCalls []MicroflowCallInfo) {
	// Build a set of all available commands (format: "AppName.CommandName")
	availableCommands := make(map[string]bool)
	for _, call := range allMicroflowCalls {
		if call.CallType == "ExternalAction" && call.AppName != "" && call.CommandName != "" {
			commandKey := call.AppName + "." + call.CommandName
			availableCommands[commandKey] = true
		}
	}

	// Check each Target Command
	missingCommands := make(map[string]bool)
	for _, pageInfo := range pageCommands {
		for _, button := range pageInfo.Commands {
			if button.TargetCommand != "" && button.TargetCommand != "-" {
				if !availableCommands[button.TargetCommand] {
					missingCommands[button.TargetCommand] = true
				}
			}
		}
	}

	// Report missing commands
	if len(missingCommands) > 0 {
		fmt.Printf("  ⚠️  Warning: %d command(s) not found in Microflow/Action Calls:\n", len(missingCommands))
		for cmd := range missingCommands {
			fmt.Printf("     - %s\n", cmd)
		}
	} else {
		fmt.Printf("  ✓ All Target Commands validated successfully\n")
	}
}

func collectNavigationItems(db *sql.DB, contentsDir string, report *ManifestReport, pageAccessMap map[string][]string) error {
	// Query NavigationDocument only
	fmt.Printf("  🔍 Looking for NavigationDocument...\n")
	query := `SELECT UnitID FROM Unit`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unitID []byte
		err := rows.Scan(&unitID)
		if err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		contentMap, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		typeName := ""
		if t, ok := contentMap["$Type"].(string); ok {
			typeName = t
		}

		if typeName == "Navigation$NavigationDocument" {
			// Extract from NavigationDocument with hierarchy
			items := extractNavigationHierarchy(contentMap, db, contentsDir, pageAccessMap)
			report.NavigationItems = append(report.NavigationItems, items...)
			fmt.Printf("  📋 Extracting navigation hierarchy...\n")
			break // Only process first NavigationDocument
		}
	}

	fmt.Printf("\n  ✅ Completed: Navigation structure extracted\n")
	fmt.Printf("  📊 Result: %d navigation item(s) found\n", len(report.NavigationItems))
	return nil
}

// extractNavigationHierarchy extracts the complete navigation hierarchy from NavigationDocument
func extractNavigationHierarchy(content map[string]interface{}, db *sql.DB, contentsDir string, pageAccessMap map[string][]string) []NavigationItem {
	var items []NavigationItem

	// Get Profiles array
	if profiles, ok := content["Profiles"]; ok {
		if profilesArr, ok := profiles.(primitive.A); ok {
			for i, profile := range profilesArr {
				if i == 0 {
					continue // Skip count
				}

				if profileMap, ok := profile.(map[string]interface{}); ok {
					// Get Menu
					if menu, ok := profileMap["Menu"]; ok {
						if menuMap, ok := menu.(map[string]interface{}); ok {
							// Get Items
							if menuItems, ok := menuMap["Items"]; ok {
								if itemsArr, ok := menuItems.(primitive.A); ok {
									// Process each top-level item
									for j, item := range itemsArr {
										if j == 0 {
											continue // Skip count
										}

										if itemMap, ok := item.(map[string]interface{}); ok {
											// Extract recursively
											extracted := extractNavigationItemRecursive(itemMap, db, contentsDir, "", 1, pageAccessMap)
											items = append(items, extracted...)
										}
									}
								}
							}
						}
					}
				}

				// Process only first profile (usually Responsive)
				break
			}
		}
	}

	return items
}

// extractNavigationItemRecursive extracts a navigation item and its children recursively
func extractNavigationItemRecursive(itemMap map[string]interface{}, db *sql.DB, contentsDir string, parentCaption string, level int, pageAccessMap map[string][]string) []NavigationItem {
	var items []NavigationItem

	// Extract Caption
	caption := ""
	if captionData, ok := itemMap["Caption"]; ok {
		if captionMap, ok := captionData.(map[string]interface{}); ok {
			if captionItems, ok := captionMap["Items"]; ok {
				if capArr, ok := captionItems.(primitive.A); ok {
					for i, capItem := range capArr {
						if i == 0 {
							continue
						}
						if capMap, ok := capItem.(map[string]interface{}); ok {
							if text, ok := capMap["Text"].(string); ok {
								caption = text
								break
							}
						}
					}
				}
			}
		}
	}

	// Check if item has sub-items
	if subItems, ok := itemMap["Items"]; ok {
		if subItemsArr, ok := subItems.(primitive.A); ok {
			hasChildren := len(subItemsArr) > 1 // More than just count

			if hasChildren {
				// This is a menu group - add it without action
				groupItem := NavigationItem{
					ItemName:     caption,
					Caption:      caption,
					Target:       "",
					Module:       "",
					MenuDocument: "",
					ItemType:     "MenuGroup",
					ParentItem:   parentCaption,
					Level:        level,
					AllowedRoles: []string{}, // Menu groups don't have specific roles
				}
				items = append(items, groupItem)

				// Process children
				for i, subItem := range subItemsArr {
					if i == 0 {
						continue
					}
					if subItemMap, ok := subItem.(map[string]interface{}); ok {
						childItems := extractNavigationItemRecursive(subItemMap, db, contentsDir, caption, level+1, pageAccessMap)
						items = append(items, childItems...)
					}
				}

				return items
			}
		}
	}

	// This is a leaf item - extract action
	target := ""
	targetQualifiedName := "" // For role lookup
	itemType := "Unknown"
	moduleName := ""

	if action, ok := itemMap["Action"]; ok {
		if actionMap, ok := action.(map[string]interface{}); ok {
			// Check $Type to determine action type
			if actionType, ok := actionMap["$Type"].(string); ok {
				itemType = actionType

				// Extract target based on action type
				if strings.Contains(actionType, "ShowPage") || strings.Contains(actionType, "FormAction") {
					// Extract page reference from FormSettings
					if formSettings, ok := actionMap["FormSettings"]; ok {
						if formMap, ok := formSettings.(map[string]interface{}); ok {
							if formID, ok := formMap["Form"].(string); ok {
								// formID is the qualified name (e.g., "Module.PageName")
								targetQualifiedName = formID
								target = formID

								// Try to load as UUID first, but formID might be qualified name
								formContent, err := loadUnitContents(contentsDir, formID)
								if err == nil {
									if formType, ok := formContent["$Type"].(string); ok {
										if formType == "Menus$MenuDocument" {
											// This item opens a menu - expand its items
											menuItems := extractMenuItemsFromMenuDocument(formContent, "", "", level+1)
											for i := range menuItems {
												menuItems[i].ParentItem = caption
											}
											items = append(items, menuItems...)
											return items
										}
									}
								}

								// Extract module name from qualified name
								parts := strings.Split(formID, ".")
								if len(parts) == 2 {
									moduleName = parts[0]
								}
							}
						}
					}
				} else if strings.Contains(actionType, "Nanoflow") || strings.Contains(actionType, "Microflow") {
					itemType = "Nanoflow"

					// Extract nanoflow/microflow reference
					if strings.Contains(actionType, "Nanoflow") {
						if nanoflowRef, ok := actionMap["Nanoflow"].(string); ok {
							target = nanoflowRef
						}
					} else if strings.Contains(actionType, "Microflow") {
						if microflowRef, ok := actionMap["Microflow"].(string); ok {
							target = microflowRef
						}
					}
				}
			}
		}
	}

	// Add leaf item
	// Get allowed roles from page access map
	allowedRoles := []string{}
	if targetQualifiedName != "" {
		if roles, ok := pageAccessMap[targetQualifiedName]; ok {
			allowedRoles = roles
		}
	}

	// Resolve target page(s) - extract page name for Page items, or find pages opened by Microflow/Nanoflow
	targetPage := resolveTargetPage(db, contentsDir, itemType, target, moduleName)

	leafItem := NavigationItem{
		ItemName:     caption,
		Caption:      caption,
		Target:       target,
		TargetPage:   targetPage,
		Module:       moduleName,
		MenuDocument: "",
		ItemType:     itemType,
		ParentItem:   parentCaption,
		Level:        level,
		AllowedRoles: allowedRoles,
	}
	items = append(items, leafItem)

	return items
}

// extractNavigationItemsFromBSON - DEPRECATED, kept for compatibility
func extractNavigationItemsFromBSON(content map[string]interface{}, moduleName string, db *sql.DB, contentsDir string) []NavigationItem {
	// Old implementation - no longer used
	return []NavigationItem{}
}

// extractMenuItemsFromMenuDocument extracts menu items directly from a MenuDocument
func extractMenuItemsFromMenuDocument(content map[string]interface{}, documentName string, moduleName string, level int) []NavigationItem {
	var items []NavigationItem

	// Check if ItemCollection exists
	itemCollection, hasCollection := content["ItemCollection"]
	if !hasCollection {
		return items
	}

	// Look for ItemCollection → Items in menu document
	collectionMap, ok := itemCollection.(map[string]interface{})
	if !ok {
		return items
	}

	itemsArr, ok := collectionMap["Items"]
	if !ok {
		return items
	}

	// Handle primitive.A (BSON array type)
	var itemsList []interface{}
	if primitiveArr, ok := itemsArr.(primitive.A); ok {
		itemsList = primitiveArr
	} else if arrInterface, ok := itemsArr.([]interface{}); ok {
		itemsList = arrInterface
	} else {
		return items
	}

	// Skip the first element (count) and process menu items
	for i, item := range itemsList {
		if i == 0 {
			// First element is the count, skip it
			continue
		}

		if itemMap, ok := item.(map[string]interface{}); ok {
			navItem := parseMenuItemSimple(itemMap, moduleName, documentName, level)
			if navItem.ItemName != "" || navItem.Caption != "" {
				items = append(items, navItem)
			}
		}
	}

	return items
}

// parseMenuItemSimple parses a menu item without database lookups (simpler version)
func parseMenuItemSimple(itemMap map[string]interface{}, moduleName string, documentName string, level int) NavigationItem {
	item := NavigationItem{
		Module:       moduleName,
		MenuDocument: documentName,
		Level:        level,
	}

	// Extract caption from Text structure
	if caption, ok := itemMap["Caption"]; ok {
		if captionMap, ok := caption.(map[string]interface{}); ok {
			// Handle Items array (translations)
			if itemsArr, ok := captionMap["Items"]; ok {
				// Handle primitive.A for caption items
				var itemsList []interface{}
				if primitiveArr, ok := itemsArr.(primitive.A); ok {
					itemsList = primitiveArr
				} else if arrInterface, ok := itemsArr.([]interface{}); ok {
					itemsList = arrInterface
				}

				if itemsList != nil {
					// Skip first element (count), get first translation
					for i, trans := range itemsList {
						if i == 0 {
							continue // Skip count
						}
						if transMap, ok := trans.(map[string]interface{}); ok {
							if text, ok := transMap["Text"]; ok {
								if textStr, ok := text.(string); ok {
									item.Caption = textStr
									break // Use first translation
								}
							}
						}
					}
				}
			}
		}
	}

	// Use caption as name if not set
	if item.ItemName == "" {
		item.ItemName = item.Caption
	}

	// Extract action type and target
	if action, ok := itemMap["Action"]; ok {
		if actionMap, ok := action.(map[string]interface{}); ok {
			if actionType, ok := actionMap["$Type"]; ok {
				actionTypeStr := fmt.Sprintf("%v", actionType)

				switch {
				case strings.Contains(actionTypeStr, "FormAction"):
					item.ItemType = "Page"
					// Extract page from FormSettings
					if formSettings, ok := actionMap["FormSettings"]; ok {
						if settingsMap, ok := formSettings.(map[string]interface{}); ok {
							if form, ok := settingsMap["Form"]; ok {
								item.Target = fmt.Sprintf("%v", form)
							}
						}
					}

				case strings.Contains(actionTypeStr, "CallMicroflowClientAction"):
					item.ItemType = "Microflow"
					if microflow, ok := actionMap["Microflow"]; ok {
						item.Target = fmt.Sprintf("%v", microflow)
					}

				case strings.Contains(actionTypeStr, "CallNanoflowClientAction"):
					item.ItemType = "Nanoflow"
					if nanoflow, ok := actionMap["Nanoflow"]; ok {
						item.Target = fmt.Sprintf("%v", nanoflow)
					}

				default:
					item.ItemType = "Action"
					item.Target = actionTypeStr
				}
			}
		}
	}

	return item
}

// extractMenuItems recursively extracts menu items from MenuDocument
func extractMenuItems(menuRef interface{}, db *sql.DB, contentsDir string, moduleName string, level int) []NavigationItem {
	var items []NavigationItem

	// Handle reference to menu document
	if refStr, ok := menuRef.(string); ok {
		// Load menu document by reference
		menuContent := loadDocumentByReference(db, contentsDir, refStr)
		if menuContent != nil {
			return extractMenuItems(menuContent, db, contentsDir, moduleName, level)
		}
		return items
	}

	// Handle menu document map
	menuMap, ok := menuRef.(map[string]interface{})
	if !ok {
		return items
	}

	// Look for Items array in menu document
	if itemsArr, ok := menuMap["Items"]; ok {
		if itemsList, ok := itemsArr.([]interface{}); ok {
			for _, item := range itemsList {
				if itemMap, ok := item.(map[string]interface{}); ok {
					navItem := parseMenuItem(itemMap, db, contentsDir, moduleName, level)
					if navItem.ItemName != "" || navItem.Caption != "" {
						items = append(items, navItem)
					}

					// Check for sub-items
					if subMenu, ok := itemMap["SubMenu"]; ok {
						subItems := extractMenuItems(subMenu, db, contentsDir, moduleName, level+1)
						items = append(items, subItems...)
					}
				}
			}
		}
	}

	return items
}

// parseMenuItem parses a single menu item from BSON
func parseMenuItem(itemMap map[string]interface{}, db *sql.DB, contentsDir string, moduleName string, level int) NavigationItem {
	item := NavigationItem{
		Module: moduleName,
		Level:  level,
	}

	// Extract caption
	if caption, ok := itemMap["Caption"]; ok {
		if captionMap, ok := caption.(map[string]interface{}); ok {
			// Handle translation structure
			if textVal, ok := captionMap["Text"]; ok {
				if text, ok := textVal.(string); ok {
					item.Caption = text
				}
			}
		} else if captionStr, ok := caption.(string); ok {
			item.Caption = captionStr
		}
	}

	// Extract name/title
	if name, ok := itemMap["Name"]; ok {
		if nameStr, ok := name.(string); ok {
			item.ItemName = nameStr
		}
	}

	// If no name, use caption
	if item.ItemName == "" {
		item.ItemName = item.Caption
	}

	// Determine target and type
	// Check for Page action
	if action, ok := itemMap["Action"]; ok {
		if actionMap, ok := action.(map[string]interface{}); ok {
			// Check action type
			if actionType, ok := actionMap["$Type"]; ok {
				actionTypeStr := fmt.Sprintf("%v", actionType)

				switch {
				case strings.Contains(actionTypeStr, "PageClientAction"):
					item.ItemType = "Page"
					// Extract page reference
					if pageRef, ok := actionMap["Page"]; ok {
						item.Target = resolvePageReference(pageRef, db, contentsDir)
					}

				case strings.Contains(actionTypeStr, "MicroflowClientAction"):
					item.ItemType = "Microflow"
					// Extract microflow reference
					if mfRef, ok := actionMap["Microflow"]; ok {
						item.Target = resolveMicroflowReference(mfRef, db, contentsDir)
					}

				case strings.Contains(actionTypeStr, "NanoflowClientAction"):
					item.ItemType = "Nanoflow"
					// Extract nanoflow reference
					if nfRef, ok := actionMap["Nanoflow"]; ok {
						item.Target = resolveNanoflowReference(nfRef, db, contentsDir)
					}
				}
			}
		}
	}

	return item
}

// extractPageReference creates a navigation item from a page reference
func extractPageReference(pageRef interface{}, db *sql.DB, contentsDir string, moduleName string, itemName string, level int) NavigationItem {
	return NavigationItem{
		ItemName: itemName,
		Caption:  itemName,
		Target:   resolvePageReference(pageRef, db, contentsDir),
		Module:   moduleName,
		ItemType: "Page",
		Level:    level,
	}
}

// resolvePageReference resolves a page reference to page name
func resolvePageReference(pageRef interface{}, db *sql.DB, contentsDir string) string {
	refID := extractReferenceID(pageRef)
	if refID == "" {
		return ""
	}

	// Convert UUID string to Windows GUID binary
	refIDBytes := stringToWindowsGUID(refID)
	if refIDBytes == nil {
		return refID
	}

	// Query for page name
	var name string
	query := `SELECT ContainmentName FROM Unit WHERE UnitID = ?`
	err := db.QueryRow(query, refIDBytes).Scan(&name)
	if err == nil && name != "" {
		return name
	}

	return refID
}

// resolveMicroflowReference resolves a microflow reference to microflow name
func resolveMicroflowReference(mfRef interface{}, db *sql.DB, contentsDir string) string {
	refID := extractReferenceID(mfRef)
	if refID == "" {
		return ""
	}

	// Convert UUID string to Windows GUID binary
	refIDBytes := stringToWindowsGUID(refID)
	if refIDBytes == nil {
		return refID
	}

	// Query for microflow name
	var name string
	query := `SELECT ContainmentName FROM Unit WHERE UnitID = ?`
	err := db.QueryRow(query, refIDBytes).Scan(&name)
	if err == nil && name != "" {
		return name
	}

	return refID
}

// resolveNanoflowReference resolves a nanoflow reference to nanoflow name
func resolveNanoflowReference(nfRef interface{}, db *sql.DB, contentsDir string) string {
	refID := extractReferenceID(nfRef)
	if refID == "" {
		return ""
	}

	// Convert UUID string to Windows GUID binary
	refIDBytes := stringToWindowsGUID(refID)
	if refIDBytes == nil {
		return refID
	}

	// Query for nanoflow name
	var name string
	query := `SELECT ContainmentName FROM Unit WHERE UnitID = ?`
	err := db.QueryRow(query, refIDBytes).Scan(&name)
	if err == nil && name != "" {
		return name
	}

	return refID
}

// extractReferenceID extracts unit ID from a reference object
func extractReferenceID(ref interface{}) string {
	if refStr, ok := ref.(string); ok {
		return refStr
	}

	if refMap, ok := ref.(map[string]interface{}); ok {
		// Try $ID field
		if id, ok := refMap["$ID"]; ok {
			return fmt.Sprintf("%v", id)
		}
		// Try Unit field
		if unit, ok := refMap["Unit"]; ok {
			return fmt.Sprintf("%v", unit)
		}
	}

	return ""
}

// loadDocumentByReference loads a document's BSON content by reference
func loadDocumentByReference(db *sql.DB, contentsDir string, refID string) map[string]interface{} {
	// Convert UUID string to Windows GUID binary
	refIDBytes := stringToWindowsGUID(refID)
	if refIDBytes == nil {
		// If conversion fails, try loading directly from mprcontents
		contentMap, err := loadUnitContents(contentsDir, refID)
		if err != nil {
			return nil
		}
		return contentMap
	}

	// Query for document contents
	query := `SELECT ContentsHash FROM Unit WHERE UnitID = ?`
	var contentsHash interface{}
	err := db.QueryRow(query, refIDBytes).Scan(&contentsHash)

	if err != nil {
		// Try loading from mprcontents
		contentMap, err := loadUnitContents(contentsDir, refID)
		if err != nil {
			return nil
		}
		return contentMap
	}

	// Load from mprcontents using the UUID string
	contentMap, err := loadUnitContents(contentsDir, refID)
	if err != nil {
		return nil
	}

	return contentMap
}

// getModuleNameFromContainerID gets module name from a container ID
func getModuleNameFromContainerID(db *sql.DB, contentsDir string, containerID string) string {
	// Convert containerID string to binary format for query
	containerIDBytes := stringToWindowsGUID(containerID)
	if containerIDBytes == nil {
		return "Unknown"
	}

	// Traverse up the hierarchy to find the module (Projects$ModuleImpl)
	currentID := containerIDBytes
	maxDepth := 10 // Prevent infinite loops

	for depth := 0; depth < maxDepth; depth++ {
		query := `SELECT ContainmentName, ContainerID, UnitID FROM Unit WHERE UnitID = ?`
		var name string
		var parentID, unitID []byte

		err := db.QueryRow(query, currentID).Scan(&name, &parentID, &unitID)
		if err != nil {
			return "Unknown"
		}

		// Load the unit content to check its type
		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err == nil {
			// Check if this is a ModuleImpl
			if typeName, ok := content["$Type"].(string); ok {
				if typeName == "Projects$ModuleImpl" || typeName == "Projects$Module" {
					// Extract module name from BSON content (not ContainmentName)
					if moduleName, ok := content["Name"].(string); ok && moduleName != "" {
						return moduleName
					}
					// Fallback to ContainmentName if Name field is not available
					if name != "" {
						return name
					}
				}
			}
		}

		// Move to parent
		if len(parentID) == 0 {
			break
		}
		currentID = parentID
	}

	return "Unknown"
}

// ==== System Roles & Page Accessibility Collection ====

func collectSystemRolesAndPageAccess(db *sql.DB, contentsDir string, report *ManifestReport) error {
	fmt.Printf("  🔍 Loading security configuration...\n")

	// Step 1: Collect all System Roles from project security documents
	systemRoles := make(map[string]SystemRole) // key: RoleID, value: SystemRole
	roleIDToName := make(map[string]string)    // for quick lookups

	// Step 2: Collect module roles and map them to system roles
	moduleRoleToSystemRoles := make(map[string][]string) // key: ModuleRoleID, value: []SystemRoleNames

	// Step 3: Collect page access information
	var pageAccessList []PageAccessInfo

	// Query all units
	query := `SELECT UnitID, ContainerID, ContainmentName, ContentsHash FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		containerIDStr := blobToUUID(containerID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Extract document type
		typeName := ""
		if t, ok := content["$Type"].(string); ok {
			typeName = t
		}

		// Process ProjectSecurity documents to extract System Roles
		if typeName == "Security$ProjectSecurity" {
			if userRoles, ok := content["UserRoles"]; ok {
				extractSystemRoles(userRoles, systemRoles, roleIDToName, moduleRoleToSystemRoles)
			}
		}

		// Process Page and Snippet documents to extract AllowedModuleRoles
		if typeName == "Forms$Page" || typeName == "Forms$Snippet" {
			// Get module name from containerID
			moduleName := getModuleNameFromContainerID(db, contentsDir, containerIDStr)
			if moduleName == "" {
				moduleName = "Unknown"
			}

			// Get page name from BSON content "Name" field
			pageName := containmentName // Fallback
			if name, ok := content["Name"].(string); ok && name != "" {
				pageName = name
			}

			// Get qualified name from BSON content or construct it
			qualifiedName := ""
			if qName, ok := content["QualifiedName"].(string); ok && qName != "" {
				qualifiedName = qName
			} else {
				// Construct from module + pageName
				qualifiedName = moduleName + "." + pageName
			}

			pageAccess := extractPageAccess(content, qualifiedName, moduleName, typeName, moduleRoleToSystemRoles)
			pageAccessList = append(pageAccessList, pageAccess)
		}
	}

	// Convert systemRoles map to slice
	for _, role := range systemRoles {
		report.SystemRoles = append(report.SystemRoles, role)
	}

	report.PageAccess = pageAccessList

	fmt.Printf("\n  ✅ Completed: Security configuration analyzed\n")
	fmt.Printf("  📊 Result: %d system role(s) and %d page(s)/snippet(s) with access rules\n", len(report.SystemRoles), len(report.PageAccess))
	return nil
}

// extractSystemRoles extracts system roles from ProjectSecurity UserRoles
func extractSystemRoles(userRoles interface{}, systemRoles map[string]SystemRole, roleIDToName map[string]string, moduleRoleToSystemRoles map[string][]string) {
	// UserRoles is directly a primitive.A array
	var itemsList []interface{}
	if primitiveArr, ok := userRoles.(primitive.A); ok {
		itemsList = primitiveArr
	} else if arrInterface, ok := userRoles.([]interface{}); ok {
		itemsList = arrInterface
	}

	for i, item := range itemsList {
		if i == 0 {
			continue // Skip count element
		}
		if roleMap, ok := item.(map[string]interface{}); ok {
			roleID := ""
			roleName := ""

			// Get role ID (it's a primitive.Binary, need to convert to string)
			if binID, ok := roleMap["$ID"].(primitive.Binary); ok {
				roleID = blobToUUID(binID.Data)
			} else if strID, ok := roleMap["$ID"].(string); ok {
				roleID = strID
			}

			// Get role name
			if name, ok := roleMap["Name"].(string); ok {
				roleName = name
			}

			if roleID != "" && roleName != "" {
				systemRoles[roleID] = SystemRole{
					Name:   roleName,
					Module: "System", // System roles are at project level
				}
				roleIDToName[roleID] = roleName

				// Extract ModuleRoles from SystemRole to build reverse mapping
				if moduleRolesData, ok := roleMap["ModuleRoles"]; ok {
					var moduleRolesList []interface{}
					if primitiveArr, ok := moduleRolesData.(primitive.A); ok {
						moduleRolesList = primitiveArr
					} else if arrInterface, ok := moduleRolesData.([]interface{}); ok {
						moduleRolesList = arrInterface
					}

					for j, mrItem := range moduleRolesList {
						if j == 0 {
							continue // Skip count
						}
						// ModuleRoles contains qualified names as strings
						if moduleRoleQName, ok := mrItem.(string); ok {
							// Add this SystemRole to the list for this ModuleRole
							if _, exists := moduleRoleToSystemRoles[moduleRoleQName]; !exists {
								moduleRoleToSystemRoles[moduleRoleQName] = []string{}
							}
							moduleRoleToSystemRoles[moduleRoleQName] = append(moduleRoleToSystemRoles[moduleRoleQName], roleName)
						}
					}
				}
			}
		}
	}
}

// extractModuleRoleMapping maps module roles to system roles
// extractPageAccess extracts page access information
func extractPageAccess(content map[string]interface{}, pageName string, moduleName string, docType string, moduleRoleToSystemRoles map[string][]string) PageAccessInfo {
	// Clean up document type to show only "Page" or "Snippet"
	cleanType := strings.TrimPrefix(docType, "Forms$")

	pageAccess := PageAccessInfo{
		PageName:     pageName,
		Module:       moduleName,
		DocumentType: cleanType,
		AllowedRoles: []string{},
		IsPublic:     true, // Assume public unless roles are found
	}

	// Extract AllowedModuleRoles
	if allowedRoles, ok := content["AllowedModuleRoles"]; ok {
		// AllowedModuleRoles is directly a primitive.A array
		var itemsList []interface{}
		if primitiveArr, ok := allowedRoles.(primitive.A); ok {
			itemsList = primitiveArr
		} else if arrInterface, ok := allowedRoles.([]interface{}); ok {
			itemsList = arrInterface
		}

		for i, item := range itemsList {
			if i == 0 {
				continue // Skip count element
			}
			refID := extractReferenceID(item)
			if refID != "" {
				// Map module role to system roles
				if systemRoles, exists := moduleRoleToSystemRoles[refID]; exists {
					pageAccess.AllowedRoles = append(pageAccess.AllowedRoles, systemRoles...)
				}
			}
		}
	}

	// If roles were found, it's not public
	if len(pageAccess.AllowedRoles) > 0 {
		pageAccess.IsPublic = false
	}

	return pageAccess
}

// ==== Pages Analysis with Recursive Hierarchy (Section 6) ====

// shouldExcludeModule checks if a module should be excluded from pages analysis
// Excludes marketplace modules and UI framework modules
func shouldExcludeModule(moduleName string) bool {
	// Standard Mendix modules
	if moduleName == "Administration" || moduleName == "System" {
		return true
	}

	// Marketplace modules (typically start with EXFN_)
	if strings.HasPrefix(moduleName, "EXFN_") {
		return true
	}

	// UI framework modules (contain _DISW_ or DesignSystem)
	if strings.Contains(moduleName, "_DISW_") || strings.Contains(moduleName, "DesignSystem") {
		return true
	}

	// Atlas UI modules
	if strings.HasPrefix(moduleName, "Atlas_") {
		return true
	}

	return false
}

// collectPagesWithMicroflows collects all pages/panels and analyzes their microflow/nanoflow calls with recursive hierarchy
func collectPagesWithMicroflows(db *sql.DB, contentsDir string) ([]PageAnalysisInfo, error) {
	// Query all Units
	query := `SELECT UnitID, ContainerID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	var pagesAnalysis []PageAnalysisInfo
	microflowCache := make(map[string]map[string]interface{})   // Cache: microflowName -> content
	hierarchyCache := make(map[string][]MicroflowCallHierarchy) // Cache: microflowName -> hierarchy

	processedCount := 0
	cacheHits := 0
	scannedCount := 0
	pagesFoundCount := 0
	fileReadErrors := 0
	bsonParseErrors := 0

	for rows.Next() {
		scannedCount++
		var unitIDBlob, containerIDBlob []byte
		if err := rows.Scan(&unitIDBlob, &containerIDBlob); err != nil {
			continue
		}

		unitID := blobToUUID(unitIDBlob)
		if unitID == "" {
			continue
		}

		// Remove dashes for file path
		cleanID := strings.ReplaceAll(unitID, "-", "")
		filePath := filepath.Join(contentsDir, cleanID[:2], cleanID[2:4], unitID+".mxunit")

		// Read BSON
		data, err := os.ReadFile(filePath)
		if err != nil {
			fileReadErrors++
			continue
		}

		var doc map[string]interface{}
		err = bson.Unmarshal(data, &doc)
		if err != nil {
			bsonParseErrors++
			continue
		}

		// Only process Forms$Page (pages and panels)
		typeField, ok := doc["$Type"].(string)
		if !ok || typeField != "Forms$Page" {
			continue
		}

		pagesFoundCount++

		// Get page name and module
		pageName, ok := doc["Name"].(string)
		if !ok || pageName == "" {
			continue
		}

		// Get module name from ContainerID
		containerID := blobToUUID(containerIDBlob)
		moduleName := getModuleNameFromContainerID(db, contentsDir, containerID)

		// Filter out marketplace and UI modules
		if shouldExcludeModule(moduleName) {
			continue
		}

		// Extract microflow/nanoflow calls from page
		microflowCalls := extractMicroflowCallsFromPage(doc)

		// Build call hierarchy for each microflow (up to 5 levels deep)
		var callHierarchy []MicroflowCallHierarchy
		for _, mfName := range microflowCalls {
			// Check cache first
			if cachedHierarchy, found := hierarchyCache[mfName]; found {
				callHierarchy = append(callHierarchy, MicroflowCallHierarchy{
					Name:  mfName,
					Level: 0,
					Calls: cachedHierarchy,
				})
				cacheHits++
			} else {
				// Build hierarchy and cache it
				hierarchy := buildMicroflowCallHierarchy(mfName, 0, 5, db, contentsDir, microflowCache, make(map[string]bool))
				hierarchyCache[mfName] = hierarchy.Calls
				callHierarchy = append(callHierarchy, hierarchy)
			}
		}

		// Filter hierarchy to keep only nodes with commands
		filteredHierarchy := filterHierarchyByCommands(callHierarchy, db, contentsDir, microflowCache)

		// Extract target commands from filtered hierarchy
		targetCommands := extractCommandsFromHierarchy(filteredHierarchy, db, contentsDir, microflowCache)

		pagesAnalysis = append(pagesAnalysis, PageAnalysisInfo{
			Name:           pageName,
			Module:         moduleName,
			MicroflowCalls: microflowCalls,
			CallHierarchy:  filteredHierarchy,
			TargetCommands: targetCommands,
		})

		processedCount++
	}

	fmt.Printf("  📦 Scanned %d units, %d file read errors, %d BSON parse errors\n", scannedCount, fileReadErrors, bsonParseErrors)
	fmt.Printf("  📦 Found %d Forms$Page documents, processed %d pages\n", pagesFoundCount, processedCount)
	fmt.Printf("  📦 Built call hierarchies: %d microflows cached, %d hierarchies reused\n", len(microflowCache), cacheHits)
	return pagesAnalysis, nil
}

// buildMicroflowCallHierarchy builds a recursive call tree for a microflow/nanoflow
func buildMicroflowCallHierarchy(name string, currentLevel int, maxDepth int, db *sql.DB, contentsDir string, cache map[string]map[string]interface{}, visited map[string]bool) MicroflowCallHierarchy {
	hierarchy := MicroflowCallHierarchy{
		Name:  name,
		Level: currentLevel,
		Calls: []MicroflowCallHierarchy{},
	}

	// Stop recursion if max depth reached or already visited (cycle detection)
	if currentLevel >= maxDepth || visited[name] {
		return hierarchy
	}

	// Mark as visited
	visited[name] = true

	// Load microflow content (check cache first)
	var content map[string]interface{}
	var ok bool
	if content, ok = cache[name]; !ok {
		// Not in cache - load from database on-the-fly
		content = loadMicroflowByName(db, contentsDir, name)
		if content != nil {
			cache[name] = content
		} else {
			// Microflow not found
			delete(visited, name)
			return hierarchy
		}
	}

	// Extract called microflows/nanoflows
	calledFlows := extractMicroflowCallsFromMicroflow(content)

	// Recursively build hierarchy for each called flow
	for _, calledFlow := range calledFlows {
		childHierarchy := buildMicroflowCallHierarchy(calledFlow, currentLevel+1, maxDepth, db, contentsDir, cache, visited)
		hierarchy.Calls = append(hierarchy.Calls, childHierarchy)
	}

	// Unmark visited for this path (allow other branches)
	delete(visited, name)

	return hierarchy
}

// extractMicroflowCallsFromPage extracts microflow/nanoflow names from a page document
func extractMicroflowCallsFromPage(pageDoc map[string]interface{}) []string {
	var microflows []string
	seen := make(map[string]bool)

	var traverse func(interface{})
	traverse = func(v interface{}) {
		switch val := v.(type) {
		case map[string]interface{}:
			// Check for Nanoflow field (direct reference)
			if nanoflowField, ok := val["Nanoflow"].(string); ok && nanoflowField != "" {
				if !seen[nanoflowField] {
					microflows = append(microflows, nanoflowField)
					seen[nanoflowField] = true
				}
			}
			// Check for Microflow field (direct reference)
			if microflowField, ok := val["Microflow"].(string); ok && microflowField != "" {
				if !seen[microflowField] {
					microflows = append(microflows, microflowField)
					seen[microflowField] = true
				}
			}
			// Check for MicroflowCall field
			if mfCall, ok := val["MicroflowCall"].(string); ok && mfCall != "" {
				if !seen[mfCall] {
					microflows = append(microflows, mfCall)
					seen[mfCall] = true
				}
			}
			// Traverse nested objects
			for _, v2 := range val {
				traverse(v2)
			}
		case []interface{}:
			for _, item := range val {
				traverse(item)
			}
		case primitive.A:
			for _, item := range val {
				traverse(item)
			}
		}
	}

	traverse(pageDoc)
	return microflows
}

// extractMicroflowCallsFromMicroflow extracts microflow/nanoflow calls from a microflow document
func extractMicroflowCallsFromMicroflow(microflowDoc map[string]interface{}) []string {
	var microflows []string
	seen := make(map[string]bool)

	var traverse func(interface{})
	traverse = func(v interface{}) {
		switch val := v.(type) {
		case map[string]interface{}:
			// Check for $Type = "Microflows$MicroflowCall" or "Microflows$NanoflowCall"
			if typeField, ok := val["$Type"].(string); ok {
				if typeField == "Microflows$MicroflowCall" || typeField == "Microflows$NanoflowCall" {
					// Extract MicroflowCall field (can be "MicroflowCall" or "Microflow")
					mfCallName := ""
					if mfCall, ok := val["MicroflowCall"].(string); ok && mfCall != "" {
						mfCallName = mfCall
					} else if microflow, ok := val["Microflow"].(string); ok && microflow != "" {
						mfCallName = microflow
					}

					if mfCallName != "" && !seen[mfCallName] {
						microflows = append(microflows, mfCallName)
						seen[mfCallName] = true
					}
				}
			}
			// Check for direct Nanoflow field reference
			if nanoflowField, ok := val["Nanoflow"].(string); ok && nanoflowField != "" {
				if !seen[nanoflowField] {
					microflows = append(microflows, nanoflowField)
					seen[nanoflowField] = true
				}
			}
			// Check for direct Microflow field reference
			if microflowField, ok := val["Microflow"].(string); ok && microflowField != "" {
				if !seen[microflowField] {
					microflows = append(microflows, microflowField)
					seen[microflowField] = true
				}
			}
			// Traverse nested objects
			for _, v2 := range val {
				traverse(v2)
			}
		case []interface{}:
			for _, item := range val {
				traverse(item)
			}
		case primitive.A:
			for _, item := range val {
				traverse(item)
			}
		}
	}

	traverse(microflowDoc)
	return microflows
}

// loadMicroflowByName loads a microflow/nanoflow by name from the database
func loadMicroflowByName(db *sql.DB, contentsDir string, name string) map[string]interface{} {
	// Extract simple name from qualified name (Module.Name -> Name)
	simpleName := name
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		simpleName = name[idx+1:]
	}

	// Query all Units (brute force scan)
	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var unitIDBlob []byte
		if err := rows.Scan(&unitIDBlob); err != nil {
			continue
		}

		// Use blobToUUID instead of hex.EncodeToString
		unitID := blobToUUID(unitIDBlob)
		if unitID == "" {
			continue
		}

		// Remove dashes for file path
		cleanID := strings.ReplaceAll(unitID, "-", "")
		filePath := filepath.Join(contentsDir, cleanID[:2], cleanID[2:4], unitID+".mxunit")

		data, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		var doc map[string]interface{}
		err = bson.Unmarshal(data, &doc)
		if err != nil {
			continue
		}

		// Check if this is a Microflow or Nanoflow with matching name
		typeField, ok := doc["$Type"].(string)
		if !ok {
			continue
		}
		if typeField != "Microflows$Microflow" && typeField != "Microflows$Nanoflow" {
			continue
		}

		nameField, ok := doc["Name"].(string)
		// Match both simple name and qualified name
		if ok && (nameField == name || nameField == simpleName) {
			return doc
		}
	}

	return nil
}

// extractCommandsFromHierarchy extracts all external action commands from a call hierarchy
func extractCommandsFromHierarchy(hierarchy []MicroflowCallHierarchy, db *sql.DB, contentsDir string, cache map[string]map[string]interface{}) []string {
	var commands []string
	seen := make(map[string]bool)

	var traverse func([]MicroflowCallHierarchy)
	traverse = func(nodes []MicroflowCallHierarchy) {
		for _, node := range nodes {
			// Load microflow content
			var content map[string]interface{}
			var ok bool
			if content, ok = cache[node.Name]; !ok {
				// Not in cache - load on-the-fly
				content = loadMicroflowByName(db, contentsDir, node.Name)
				if content != nil {
					cache[node.Name] = content
				}
			}

			if content != nil {
				// Extract commands from this microflow
				nodeCommands := extractExternalActionsFromContent(content)
				for _, cmd := range nodeCommands {
					if !seen[cmd] {
						commands = append(commands, cmd)
						seen[cmd] = true
					}
				}
			}

			// Traverse children
			if len(node.Calls) > 0 {
				traverse(node.Calls)
			}
		}
	}

	traverse(hierarchy)
	return commands
}

// extractExternalActionsFromContent extracts ExternalAction and CallExternalAction commands from microflow content
func extractExternalActionsFromContent(content map[string]interface{}) []string {
	var commands []string
	seen := make(map[string]bool)

	var traverse func(interface{})
	traverse = func(v interface{}) {
		switch val := v.(type) {
		case map[string]interface{}:
			// Check for $Type
			if typeField, ok := val["$Type"].(string); ok {
				// Pattern 1: Microflows$ExternalAction (legacy)
				if typeField == "Microflows$ExternalAction" {
					appName, _ := val["AppName"].(string)
					commandName, _ := val["CommandName"].(string)

					// Clean values (remove quotes and whitespace)
					appName = strings.Trim(strings.TrimSpace(appName), "'\"")
					commandName = strings.Trim(strings.TrimSpace(commandName), "'\"")

					if appName != "" && commandName != "" {
						cmd := appName + "." + commandName
						if !seen[cmd] {
							commands = append(commands, cmd)
							seen[cmd] = true
						}
					}
				}
				// Pattern 2: Microflows$CallExternalAction (OData Services)
				if typeField == "Microflows$CallExternalAction" {
					consumedService, _ := val["ConsumedODataService"].(string)
					actionName, _ := val["Name"].(string)
					if consumedService != "" && actionName != "" {
						// Parse service name (format: "Module.ServiceName")
						parts := strings.Split(consumedService, ".")
						serviceName := consumedService
						if len(parts) > 1 {
							serviceName = parts[len(parts)-1]
						}
						cmd := serviceName + "." + actionName
						if !seen[cmd] {
							commands = append(commands, cmd)
							seen[cmd] = true
						}
					}
				}
				// Pattern 3: Microflows$MicroflowCall calling CallCommand_MF pattern
				if typeField == "Microflows$MicroflowCall" {
					// Check if calling a CallCommand_MF or similar pattern
					var microflowName string
					if mfCall, ok := val["MicroflowCall"].(string); ok && mfCall != "" {
						microflowName = mfCall
					} else if microflow, ok := val["Microflow"].(string); ok && microflow != "" {
						microflowName = microflow
					}

					// If calling CallCommand_MF or CallCommandAction, extract AppName and CommandName
					if strings.Contains(microflowName, "CallCommand") {
						// Try to extract AppName and CommandName from ParameterMappings
						appName := ""
						commandName := ""

						if paramMappings, ok := val["ParameterMappings"].(primitive.A); ok {
							for _, pm := range paramMappings {
								if pmMap, ok := pm.(map[string]interface{}); ok {
									paramName, _ := pmMap["Parameter"].(string)

									// Extract argument - can be in "Argument" or "Value.Argument"
									argument := ""
									if arg, ok := pmMap["Argument"].(string); ok {
										argument = arg
									} else if valueObj, ok := pmMap["Value"].(map[string]interface{}); ok {
										if arg, ok := valueObj["Argument"].(string); ok {
											argument = arg
										}
									}

									// Extract parameter name (last part after dot)
									if strings.Contains(paramName, ".") {
										parts := strings.Split(paramName, ".")
										paramName = parts[len(parts)-1]
									}

									// Clean argument thoroughly
									argument = strings.TrimSpace(argument)
									argument = strings.TrimPrefix(argument, "$")
									argument = strings.Trim(argument, "'\"")
									argument = strings.ReplaceAll(argument, "\n", "")
									argument = strings.ReplaceAll(argument, "\r", "")
									argument = strings.ReplaceAll(argument, "\\n", "")
									argument = strings.TrimSpace(argument)

									// Skip invalid values
									if argument != "" && strings.ToLower(argument) != "empty" {
										if strings.Contains(strings.ToLower(paramName), "appname") {
											appName = argument
										} else if strings.Contains(strings.ToLower(paramName), "commandname") {
											commandName = argument
										}
									}
								}
							}
						}

						// Build command identifier
						var cmd string
						if appName != "" && commandName != "" {
							cmd = appName + "." + commandName
						} else if commandName != "" {
							cmd = commandName
						} else if appName != "" {
							cmd = appName
						} else {
							cmd = "DynamicCommand"
						}

						if !seen[cmd] {
							commands = append(commands, cmd)
							seen[cmd] = true
						}
					}
				}
				// Pattern 4: Microflows$JavaActionCallAction
				if typeField == "Microflows$JavaActionCallAction" {
					// Try to extract AppName and CommandName from ParameterMappings
					appName := ""
					commandName := ""

					if paramMappings, ok := val["ParameterMappings"].(primitive.A); ok {
						for _, pm := range paramMappings {
							if pmMap, ok := pm.(map[string]interface{}); ok {
								paramName, _ := pmMap["Parameter"].(string)

								// Extract argument - can be in "Argument" or "Value.Argument"
								argument := ""
								if arg, ok := pmMap["Argument"].(string); ok {
									argument = arg
								} else if valueObj, ok := pmMap["Value"].(map[string]interface{}); ok {
									if arg, ok := valueObj["Argument"].(string); ok {
										argument = arg
									}
								}

								// Extract parameter name (last part after dot)
								if strings.Contains(paramName, ".") {
									parts := strings.Split(paramName, ".")
									paramName = parts[len(parts)-1]
								}

								// Clean argument thoroughly
								argument = strings.TrimSpace(argument)
								argument = strings.TrimPrefix(argument, "$")
								argument = strings.Trim(argument, "'\"")
								argument = strings.ReplaceAll(argument, "\n", "")
								argument = strings.ReplaceAll(argument, "\r", "")
								argument = strings.ReplaceAll(argument, "\\n", "")
								argument = strings.TrimSpace(argument)

								// Skip invalid values
								if argument != "" && strings.ToLower(argument) != "empty" {
									if strings.Contains(strings.ToLower(paramName), "appname") {
										appName = argument
									} else if strings.Contains(strings.ToLower(paramName), "commandname") {
										commandName = argument
									}
								}
							}
						}
					}

					// Build command identifier
					var cmd string
					if appName != "" && commandName != "" {
						cmd = appName + "." + commandName
					} else if commandName != "" {
						cmd = commandName
					} else if appName != "" {
						cmd = appName
					} else {
						// If no parameters, use JavaAction name
						if javaActionName, ok := val["JavaAction"].(string); ok && javaActionName != "" {
							cmd = javaActionName
						} else {
							cmd = "JavaActionCommand"
						}
					}

					if !seen[cmd] {
						commands = append(commands, cmd)
						seen[cmd] = true
					}
				}
			}
			// Traverse nested objects
			for _, v2 := range val {
				traverse(v2)
			}
		case []interface{}:
			for _, item := range val {
				traverse(item)
			}
		case primitive.A:
			for _, item := range val {
				traverse(item)
			}
		}
	}

	traverse(content)
	return commands
}

// nodeHasCommands checks if a microflow node contains external action commands
func nodeHasCommands(nodeName string, db *sql.DB, contentsDir string, cache map[string]map[string]interface{}) bool {
	var content map[string]interface{}
	var ok bool

	// Try to get from cache first
	if content, ok = cache[nodeName]; !ok {
		// Not in cache - load on-the-fly
		content = loadMicroflowByName(db, contentsDir, nodeName)
		if content != nil {
			cache[nodeName] = content
		}
	}

	if content == nil {
		return false
	}

	// Check if this microflow has any external actions
	commands := extractExternalActionsFromContent(content)
	return len(commands) > 0
}

// filterHierarchyByCommands filters the hierarchy to keep only nodes with commands (or descendants with commands)
func filterHierarchyByCommands(hierarchy []MicroflowCallHierarchy, db *sql.DB, contentsDir string, cache map[string]map[string]interface{}) []MicroflowCallHierarchy {
	var filtered []MicroflowCallHierarchy

	for _, node := range hierarchy {
		// Recursively filter children first
		filteredCalls := []MicroflowCallHierarchy{}
		if len(node.Calls) > 0 {
			filteredCalls = filterHierarchyByCommands(node.Calls, db, contentsDir, cache)
		}

		// Keep this node if:
		// 1. It has commands directly, OR
		// 2. It has children that have commands (after filtering)
		hasCommands := nodeHasCommands(node.Name, db, contentsDir, cache)
		hasValidChildren := len(filteredCalls) > 0

		if hasCommands || hasValidChildren {
			// Keep this node
			newNode := MicroflowCallHierarchy{
				Name:  node.Name,
				Calls: filteredCalls,
			}
			filtered = append(filtered, newNode)
		}
		// Otherwise, skip this node (it has no commands and no valid children)
	}

	return filtered
}

// ==== Markdown Report Generation ====

func generateMarkdownReport(report *ManifestReport, outputPath string, options *ReportOptions) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Header
	fmt.Fprintf(file, "# Manifest Report: %s\n\n", report.ProjectName)
	fmt.Fprintf(file, "**Mendix Version:** %s  \n", report.MendixVersion)
	fmt.Fprintf(file, "**MPR File:** %s  \n", report.MPRPath)
	fmt.Fprintf(file, "**Generated:** %s  \n\n", report.GeneratedAt)
	fmt.Fprintf(file, "---\n\n")

	// Summary
	totalEntities := 0
	for _, entities := range report.Entities {
		totalEntities += len(entities)
	}

	fmt.Fprintf(file, "## Summary\n\n")
	if options.IncludeEntities {
		fmt.Fprintf(file, "- **External Entities:** %d (across %d modules)\n", totalEntities, len(report.Entities))
	}
	if options.IncludeMicroflows {
		fmt.Fprintf(file, "- **Microflow/Action Calls:** %d\n", len(report.MicroflowCalls))
	}
	if options.IncludeWidgets {
		fmt.Fprintf(file, "- **Signal Manager Subscriptions:** %d subscription(s)\n", len(report.Widgets))
	}
	if options.IncludeNavigation {
		fmt.Fprintf(file, "- **Navigation Items:** %d\n", len(report.NavigationItems))
	}
	if options.IncludePageCommands {
		totalCommands := 0
		validatedCommands := 0
		for _, pageInfo := range report.NavigationCommands {
			for _, button := range pageInfo.Commands {
				totalCommands++
				if button.TargetCommand != "" && button.TargetCommand != "-" {
					validatedCommands++
				}
			}
		}
		fmt.Fprintf(file, "- **Pages with Commands:** %d page(s) analyzed\n", len(report.NavigationCommands))
		if totalCommands > 0 {
			fmt.Fprintf(file, "- **Command Buttons:** %d total, %d with extracted commands\n", totalCommands, validatedCommands)
		}
	}
	if options.IncludeRoles {
		fmt.Fprintf(file, "- **System Roles:** %d\n", len(report.SystemRoles))
		fmt.Fprintf(file, "- **Pages/Snippets:** %d\n", len(report.PageAccess))
	}
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "---\n\n")

	// Section 1: External Entities
	if options.IncludeEntities {
		fmt.Fprintf(file, "## 1. External Entities\n\n")
		fmt.Fprintf(file, "External OData entities used in the project, grouped by module.\n\n")

		if len(report.Entities) == 0 {
			fmt.Fprintf(file, "_No external entities found._\n\n")
		} else {
			// Sort modules alphabetically for consistent output
			moduleNames := make([]string, 0, len(report.Entities))
			for moduleName := range report.Entities {
				moduleNames = append(moduleNames, moduleName)
			}
			// Simple bubble sort for consistency
			for i := 0; i < len(moduleNames); i++ {
				for j := i + 1; j < len(moduleNames); j++ {
					if moduleNames[i] > moduleNames[j] {
						moduleNames[i], moduleNames[j] = moduleNames[j], moduleNames[i]
					}
				}
			}

			for _, moduleName := range moduleNames {
				entities := report.Entities[moduleName]
				fmt.Fprintf(file, "### Module: %s\n\n", moduleName)

				for _, entity := range entities {
					// Header for each entity
					fmt.Fprintf(file, "#### Entity: %s\n\n", entity.Name)
					fmt.Fprintf(file, "**Published From:** %s\n\n", entity.PublishedFrom)

					// Attributes table (one attribute per row)
					if options.IncludeAttributes {
						if len(entity.Attributes) > 0 {
							fmt.Fprintf(file, "| Attribute | Type |\n")
							fmt.Fprintf(file, "|-----------|------|\n")
							for _, attr := range entity.Attributes {
								// Parse attribute to extract name and type
								// Format is "Name (Type)" or just "Name"
								attrName, attrType := parseAttribute(attr)
								fmt.Fprintf(file, "| %s | %s |\n", attrName, attrType)
							}
							fmt.Fprintf(file, "\n")
						} else {
							fmt.Fprintf(file, "_No attributes_\n\n")
						}
					}
				}
			}
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Section 2: Microflow/Action Calls
	if options.IncludeMicroflows {
		fmt.Fprintf(file, "## 2. Microflow/Action Calls\n\n")
		fmt.Fprintf(file, "Microflow and action calls found in the project (MicroflowCall, JavaAction, ExternalAction).\n\n")

		if len(report.MicroflowCalls) == 0 {
			fmt.Fprintf(file, "_No microflow/action calls found._\n\n")
		} else {
			fmt.Fprintf(file, "Found %d call(s):\n\n", len(report.MicroflowCalls))
			fmt.Fprintf(file, "| Microflow | Module | Call Type | AppName | CommandName |\n")
			fmt.Fprintf(file, "|-----------|--------|-----------|---------|-------------|\n")

			for _, call := range report.MicroflowCalls {
				appName := call.AppName
				if appName == "" {
					appName = "-"
				}
				commandName := call.CommandName
				if commandName == "" {
					commandName = "-"
				}

				fmt.Fprintf(file, "| %s | %s | %s | %s | %s |\n",
					call.MicroflowName,
					call.Module,
					call.CallType,
					appName,
					commandName,
				)
			}
			fmt.Fprintf(file, "\n")
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Section 3: Signal Manager Subscriptions
	if options.IncludeWidgets {
		fmt.Fprintf(file, "## 3. Signal Manager Subscriptions\n\n")
		fmt.Fprintf(file, "Signal subscriptions from Signal Manager widgets (siemens.mxtosignal.MxToSignal).\n\n")

		if len(report.Widgets) == 0 {
			fmt.Fprintf(file, "_No signal subscriptions found._\n\n")
		} else {
			fmt.Fprintf(file, "Found %d subscription(s):\n\n", len(report.Widgets))
			fmt.Fprintf(file, "| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |\n")
			fmt.Fprintf(file, "|--------|---------------|----------|-------------|----------|---------------------|\n")

			for _, widget := range report.Widgets {
				module := widget.Module
				if module == "" {
					module = "-"
				}
				docType := widget.DocumentType
				if docType == "" {
					docType = "-"
				}
				docName := widget.DocumentName
				if docName == "" {
					docName = "-"
				}
				signalName := widget.SignalName
				if signalName == "" {
					signalName = "-"
				}
				appName := widget.AppName
				if appName == "" {
					appName = "-"
				}

				// Show Yes/No instead of actual filter value
				filterPresent := "No"
				if widget.SubscriptionFilter != "" {
					filterPresent = "Yes"
				}

				fmt.Fprintf(file, "| %s | %s | %s | %s | %s | %s |\n",
					module,
					docType,
					docName,
					signalName,
					appName,
					filterPresent,
				)
			}
			fmt.Fprintf(file, "\n")
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Section 4: Navigation Items
	if options.IncludeNavigation {
		fmt.Fprintf(file, "## 4. Navigation Items\n\n")

		if len(report.NavigationItems) == 0 {
			fmt.Fprintf(file, "_No navigation items found._\n\n")
		} else {
			fmt.Fprintf(file, "| Parent Node | Node | Target Page | User Roles |\n")
			fmt.Fprintf(file, "|-------------|------|-------------|------------|\n")

			for _, item := range report.NavigationItems {
				parent := item.ParentItem
				if parent == "" {
					parent = "-"
				}

				caption := item.Caption
				if caption == "" {
					caption = item.ItemName
				}
				if caption == "" {
					caption = "-"
				}

				targetPage := item.TargetPage
				if targetPage == "" {
					targetPage = "-"
				}

				roles := "-"
				if len(item.AllowedRoles) > 0 {
					roles = strings.Join(item.AllowedRoles, ", ")
				}

				fmt.Fprintf(file, "| %s | %s | %s | %s |\n", parent, caption, targetPage, roles)
			}
			fmt.Fprintf(file, "\n")
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Section 5: Navigation Page Commands (from navigation pages)
	if options.IncludePageCommands {
		sectionNum := 5
		if options.IncludeRoles {
			sectionNum = 6
		}
		fmt.Fprintf(file, "## %d. Navigation Page Commands\n\n", sectionNum)
		fmt.Fprintf(file, "Command bar actions extracted from navigation pages. Shows buttons in the vertical command bar of the Right placeholder.\n\n")

		if len(report.NavigationCommands) == 0 {
			fmt.Fprintf(file, "_No page commands found._\n\n")
		} else {
			fmt.Fprintf(file, "Found commands in %d page(s):\n\n", len(report.NavigationCommands))

			for _, pageInfo := range report.NavigationCommands {
				fmt.Fprintf(file, "### %s\n\n", pageInfo.PageName)
				fmt.Fprintf(file, "| Caption | Target Page | Target AppName | Target CommandName |\n")
				fmt.Fprintf(file, "|---------|-------------|----------------|--------------------|\n")

				for _, cmd := range pageInfo.Commands {
					caption := cmd.Caption
					if caption == "" {
						caption = "-"
					}

					targetPage := cmd.TargetPage
					if targetPage == "" {
						targetPage = "-"
					}

					targetAppName := "-"
					targetCommandName := "-"

					if cmd.TargetCommand != "" {
						if idx := strings.LastIndex(cmd.TargetCommand, "."); idx >= 0 {
							// Extract AppName before the dot and CommandName after the dot
							targetAppName = cmd.TargetCommand[:idx]
							targetCommandName = cmd.TargetCommand[idx+1:]
						} else {
							// No dot found, use the whole string as command name
							targetCommandName = cmd.TargetCommand
						}
					}

					fmt.Fprintf(file, "| %s | %s | %s | %s |\n", caption, targetPage, targetAppName, targetCommandName)
				}
				fmt.Fprintf(file, "\n")
			}
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Section 5/6: System Roles & Page Accessibility
	if options.IncludeRoles {
		sectionNum := 5
		if options.IncludePageCommands {
			sectionNum = 7
		}
		fmt.Fprintf(file, "## %d. System Roles\n\n", sectionNum)

		// Part 1: System Roles List
		fmt.Fprintf(file, "### %d.1 System Roles\n\n", sectionNum)
		if len(report.SystemRoles) == 0 {
			fmt.Fprintf(file, "_No system roles found._\n\n")
		} else {
			fmt.Fprintf(file, "Found %d system role(s):\n\n", len(report.SystemRoles))
			fmt.Fprintf(file, "| Role Name | Module |\n")
			fmt.Fprintf(file, "|-----------|--------|\n")

			for _, role := range report.SystemRoles {
				fmt.Fprintf(file, "| %s | %s |\n", role.Name, role.Module)
			}
			fmt.Fprintf(file, "\n")
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Section 7: Pages/Panels Commands Hierarchy (if enabled)
	if options.IncludePagesCommandsHierarchy && len(report.PagesAnalysis) > 0 {
		// Calculate section number based on enabled sections
		sectionNum := 5
		if options.IncludePageCommands {
			sectionNum++
		}
		if options.IncludeRoles {
			sectionNum++
		}
		fmt.Fprintf(file, "## %d. Pages/Panels Commands Hierarchy\n\n", sectionNum)
		fmt.Fprintf(file, "Microflows and nanoflows called by each page/panel, showing recursive call hierarchy up to 5 levels (in YAML structure). Microflows called transitively are loaded on-the-fly from the database when needed.\n\n")
		fmt.Fprintf(file, "```yaml\n")
		fmt.Fprintf(file, "pages:\n")

		for _, page := range report.PagesAnalysis {
			fmt.Fprintf(file, "  - name: %s\n", page.Name)
			fmt.Fprintf(file, "    module: %s\n", page.Module)
			fmt.Fprintf(file, "    flows:\n")

			// Write hierarchy using recursive helper
			if len(page.CallHierarchy) > 0 {
				writeHierarchy(file, page.CallHierarchy, "      ", false)
			} else {
				fmt.Fprintf(file, "      []\n")
			}

			fmt.Fprintf(file, "    target_commands:\n")
			if len(page.TargetCommands) > 0 {
				for _, cmd := range page.TargetCommands {
					fmt.Fprintf(file, "      - %s\n", cmd)
				}
			} else {
				fmt.Fprintf(file, "      []\n")
			}
			fmt.Fprintf(file, "\n")
		}

		fmt.Fprintf(file, "```\n\n")
		fmt.Fprintf(file, "---\n\n")
	}

	// Section 8: PageCommands (simplified view - if enabled)
	if options.IncludePagesCommandsHierarchy && len(report.PagesAnalysis) > 0 {
		// Calculate section number based on enabled sections
		sectionNum := 5
		if options.IncludePageCommands {
			sectionNum++
		}
		if options.IncludeRoles {
			sectionNum++
		}
		sectionNum++ // Increment for the previous Pages/Panels Commands Hierarchy section

		fmt.Fprintf(file, "## %d. PageCommands\n\n", sectionNum)
		fmt.Fprintf(file, "Simplified view showing only the target commands for each page/panel.\n\n")

		// Create a table with columns: Page/Panel, Module, Target AppName, Target CommandName
		fmt.Fprintf(file, "| Page/Panel | Module | Target AppName | Target CommandName |\n")
		fmt.Fprintf(file, "|------------|--------|----------------|--------------------|\n")

		for _, page := range report.PagesAnalysis {
			appNamesStr := ""
			commandsStr := ""

			if len(page.TargetCommands) > 0 {
				appNames := make([]string, 0, len(page.TargetCommands))
				commandNames := make([]string, 0, len(page.TargetCommands))

				for _, cmd := range page.TargetCommands {
					if idx := strings.LastIndex(cmd, "."); idx >= 0 {
						// Extract AppName before the dot and CommandName after the dot
						appNames = append(appNames, cmd[:idx])
						commandNames = append(commandNames, cmd[idx+1:])
					} else {
						// No dot found, use "-" for app name and whole string as command name
						appNames = append(appNames, "-")
						commandNames = append(commandNames, cmd)
					}
				}

				appNamesStr = strings.Join(appNames, "<br>")
				commandsStr = strings.Join(commandNames, "<br>")
			} else {
				appNamesStr = "-"
				commandsStr = "-"
			}

			fmt.Fprintf(file, "| %s | %s | %s | %s |\n", page.Name, page.Module, appNamesStr, commandsStr)
		}

		fmt.Fprintf(file, "\n---\n\n")
	}

	// Footer
	fmt.Fprintf(file, "_Report generated by export_manifest tool_\n")

	return nil
}

// writeHierarchy recursively writes microflow call hierarchy in YAML format
// insideCalls indicates if we're already inside a "calls:" section (to avoid repeating it)
func writeHierarchy(file *os.File, hierarchy []MicroflowCallHierarchy, indent string, insideCalls bool) {
	for _, node := range hierarchy {
		// Always write "name:" for all nodes
		fmt.Fprintf(file, "%s- name: %s\n", indent, node.Name)
		if len(node.Calls) > 0 {
			if !insideCalls {
				// First level: write "calls:"
				fmt.Fprintf(file, "%s  calls:\n", indent)
				writeHierarchy(file, node.Calls, indent+"    ", true)
			} else {
				// Nested levels: just indent, no "calls:" keyword
				writeHierarchy(file, node.Calls, indent+"  ", true)
			}
		}
	}
}

// generateJSONReport generates a JSON report file respecting the provided options
func generateJSONReport(report *ManifestReport, outputPath string, options *ReportOptions) error {
	// Create a filtered copy of the report based on options
	filteredReport := ManifestReport{
		ProjectName:   report.ProjectName,
		MendixVersion: report.MendixVersion,
		MPRPath:       report.MPRPath,
		GeneratedAt:   report.GeneratedAt,
	}

	// Include entities if enabled
	if options.IncludeEntities {
		if options.IncludeAttributes {
			// Include entities with attributes
			filteredReport.Entities = report.Entities
		} else {
			// Include entities but remove attributes
			filteredReport.Entities = make(map[string][]EntityInfo)
			for module, entities := range report.Entities {
				filteredEntities := make([]EntityInfo, len(entities))
				for i, entity := range entities {
					filteredEntities[i] = EntityInfo{
						Name:           entity.Name,
						Module:         entity.Module,
						Attributes:     []string{}, // Empty attributes
						PublishedFrom:  entity.PublishedFrom,
						EntityTypeName: entity.EntityTypeName,
					}
				}
				filteredReport.Entities[module] = filteredEntities
			}
		}
	} else {
		filteredReport.Entities = make(map[string][]EntityInfo)
	}

	// Include microflows if enabled
	if options.IncludeMicroflows {
		filteredReport.MicroflowCalls = report.MicroflowCalls
	} else {
		filteredReport.MicroflowCalls = []MicroflowCallInfo{}
	}

	// Include widgets if enabled
	if options.IncludeWidgets {
		filteredReport.Widgets = report.Widgets
	} else {
		filteredReport.Widgets = []WidgetInfo{}
	}

	// Include navigation if enabled
	if options.IncludeNavigation {
		filteredReport.NavigationItems = report.NavigationItems
	} else {
		filteredReport.NavigationItems = []NavigationItem{}
	}

	// Include roles if enabled
	if options.IncludeRoles {
		filteredReport.SystemRoles = report.SystemRoles
		filteredReport.PageAccess = report.PageAccess
	} else {
		filteredReport.SystemRoles = []SystemRole{}
		filteredReport.PageAccess = []PageAccessInfo{}
	}

	// Include page commands if enabled
	if options.IncludePageCommands {
		filteredReport.NavigationCommands = report.NavigationCommands
	} else {
		filteredReport.NavigationCommands = []PageCommandInfo{}
	}

	// Include pages analysis if enabled
	if options.IncludePagesCommandsHierarchy {
		filteredReport.PagesAnalysis = report.PagesAnalysis
		filteredReport.PageCommands = report.PageCommands // Simplified view
	} else {
		filteredReport.PagesAnalysis = []PageAnalysisInfo{}
		filteredReport.PageCommands = []PageCommandSummary{}
	}

	// Marshal with pretty-print (2 spaces indentation)
	jsonData, err := json.MarshalIndent(&filteredReport, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to file
	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

// IndexReport represents the index file structure for batch mode
type IndexReport struct {
	GeneratedAt  string        `json:\"GeneratedAt\"`
	TotalReports int           `json:\"TotalReports\"`
	SuccessCount int           `json:\"SuccessCount\"`
	FailedCount  int           `json:\"FailedCount\"`
	Reports      []ReportEntry `json:\"Reports\"`
}

// generateIndexJSON creates an index.json file listing all generated reports
func generateIndexJSON(indexPath string, entries []ReportEntry) error {
	index := IndexReport{
		GeneratedAt:  time.Now().Format("2006-01-02 15:04:05"),
		TotalReports: len(entries),
		SuccessCount: countSuccessful(entries),
		FailedCount:  countFailed(entries),
		Reports:      entries,
	}

	// Marshal with pretty-print (2 spaces indentation)
	jsonData, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index JSON: %w", err)
	}

	// Write to file
	err = os.WriteFile(indexPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write index JSON file: %w", err)
	}

	return nil
}
