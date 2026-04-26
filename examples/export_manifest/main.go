package main

import (
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
	Name           string
	Module         string
	Attributes     []string
	PublishedFrom  string
	EntityTypeName string
}

type DomainModelInfo struct {
	Name    string
	UnitID  string
	Content map[string]interface{}
}

type MicroflowCallInfo struct {
	MicroflowName string
	Module        string
	ActivityName  string
	Caption       string
	CallType      string // "MicroflowCall", "JavaAction", or "ExternalAction"
	TargetName    string // Microflow/JavaAction/ExternalAction name
	AppName       string
	CommandName   string
}

type MicroflowInfo struct {
	ID         string
	Name       string
	ModuleName string
	Content    map[string]interface{}
}

type WidgetInfo struct {
	DocumentName       string
	DocumentType       string // "Page" or "Snippet"
	Module             string
	SignalName         string
	AppName            string
	SubscriptionFilter string
}

type NavigationItem struct {
	ItemName     string
	Caption      string
	Target       string // Page or microflow name
	Module       string
	MenuDocument string   // Name of the menu document (e.g., "System", "System Counters")
	ItemType     string   // "Page", "Microflow", "Nanoflow"
	ParentItem   string   // For hierarchical structure
	Level        int      // Indentation level
	AllowedRoles []string // User roles that can access this item
}

type SystemRole struct {
	Name   string
	Module string
}

type PageAccessInfo struct {
	PageName     string
	Module       string
	DocumentType string // "Page" or "Snippet"
	AllowedRoles []string
	IsPublic     bool // No role restrictions
}

type ManifestReport struct {
	ProjectName     string
	MendixVersion   string
	MPRPath         string
	GeneratedAt     string
	Entities        map[string][]EntityInfo // by module
	MicroflowCalls  []MicroflowCallInfo     // all calls
	Widgets         []WidgetInfo            // signal manager widgets
	NavigationItems []NavigationItem        // navigation menu items
	SystemRoles     []SystemRole            // system roles
	PageAccess      []PageAccessInfo        // page accessibility
}

type ReportOptions struct {
	IncludeEntities   bool
	IncludeAttributes bool
	IncludeMicroflows bool
	IncludeWidgets    bool
	IncludeNavigation bool
	IncludeRoles      bool
}

func main() {
	// Define CLI flags
	includeEntities := flag.Bool("include-entities", false, "Include external entities in the report")
	includeAttributes := flag.Bool("include-attributes", true, "Include entity attributes in the report")
	includeMicroflows := flag.Bool("include-microflows", true, "Include microflow/action calls in the report")
	includeWidgets := flag.Bool("include-widgets", true, "Include Signal Manager widgets in the report")
	includeNavigation := flag.Bool("include-navigation", true, "Include navigation items in the report")
	includeRoles := flag.Bool("include-roles", false, "Include system roles and page accessibility in the report")
	outputDir := flag.String("output-dir", "", "Output directory for the report file (optional)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: export_manifest [options] <mpr_file_path>\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  <mpr_file_path>   Path to the Mendix MPR file (use '.' to auto-detect in current directory)\n\n")
		fmt.Fprintf(os.Stderr, "The output report will be named: <mpr_filename>-manifest.md\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  export_manifest MyApp.mpr\n")
		fmt.Fprintf(os.Stderr, "  export_manifest .\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -include-entities=true MyApp.mpr\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -output-dir=\"reports\" .\n")
	}
	flag.Parse()

	// Check positional arguments
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	mprPath := flag.Arg(0)

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

	// Generate output filename from MPR filename
	mprFilename := filepath.Base(mprPath)
	mprNameWithoutExt := strings.TrimSuffix(mprFilename, filepath.Ext(mprFilename))
	outputFilename := mprNameWithoutExt + "-manifest.md"

	// Construct final output path
	finalOutputPath := outputFilename
	if *outputDir != "" {
		// Join with outputDir
		finalOutputPath = filepath.Join(*outputDir, outputFilename)
		
		// Create output directory if it doesn't exist
		err := os.MkdirAll(*outputDir, 0755)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Store global MPR path for module traversal
	globalMPRPath = mprPath

	// Create report options
	options := ReportOptions{
		IncludeEntities:   *includeEntities,
		IncludeAttributes: *includeAttributes,
		IncludeMicroflows: *includeMicroflows,
		IncludeWidgets:    *includeWidgets,
		IncludeNavigation: *includeNavigation,
		IncludeRoles:      *includeRoles,
	}

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

	// ==== SECTION 1: External Entities ====
	if options.IncludeEntities {
		fmt.Println("🔎 Scanning for external entities...")
		err = collectExternalEntities(db, contentsDir, &report)
		if err != nil {
			fmt.Printf("Error collecting entities: %v\n", err)
			os.Exit(1)
		}
	}

	// ==== SECTION 2: Microflow/Action Calls ====
	if options.IncludeMicroflows {
		fmt.Println("\n🔎 Scanning for microflow/action calls...")
		err = collectMicroflowCalls(db, contentsDir, &report)
		if err != nil {
			fmt.Printf("Error collecting microflow calls: %v\n", err)
			os.Exit(1)
		}
	}

	// ==== SECTION 3: Signal Manager Widgets ====
	if options.IncludeWidgets {
		fmt.Println("\n🔎 Scanning for Signal Manager widgets...")
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
		fmt.Println("\n🔎 Scanning for system roles and page accessibility...")
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
		fmt.Println("\n🔎 Scanning for navigation items...")
		err = collectNavigationItems(db, contentsDir, &report, pageAccessMap)
		if err != nil {
			fmt.Printf("Error collecting navigation items: %v\n", err)
			os.Exit(1)
		}
	}

	// Generate Markdown report
	fmt.Printf("\n📝 Generating report: %s\n", finalOutputPath)
	err = generateMarkdownReport(&report, finalOutputPath, &options)
	if err != nil {
		fmt.Printf("Error generating report: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Report generated successfully!")
}

// ==== SECTION 1: External Entities Collection ====

func collectExternalEntities(db *sql.DB, contentsDir string, report *ManifestReport) error {
	// Load all domain models
	domainModels, err := listDomainModels(db, contentsDir)
	if err != nil {
		return fmt.Errorf("failed to list domain models: %w", err)
	}

	totalEntities := 0

	// For each domain model
	for _, dm := range domainModels {
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

	fmt.Printf("  📊 Total: %d external entities in %d modules\n", totalEntities, len(report.Entities))
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
	microflows, err := listMicroflows(db, contentsDir)
	if err != nil {
		return fmt.Errorf("failed to list microflows: %w", err)
	}

	totalCalls := 0
	targetMicroflow := "EXFN_ServiceLayer.CallCommand_MF"
	targetJavaAction := "EXFN_ServiceLayer.CallCommandAction"

	// Search for calls in each microflow
	for _, mf := range microflows {
		// Skip marketplace modules
		if isMarketplaceModule(mf.ModuleName) {
			continue
		}

		calls := findMicroflowCalls(mf, targetMicroflow, targetJavaAction)

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

	fmt.Printf("  📊 Total: %d microflow/action calls found\n", totalCalls)
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

		// Only process Microflows
		if typeName != "Microflows$Microflow" {
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

func findMicroflowCalls(mf MicroflowInfo, targetMicroflow string, targetJavaAction string) []TempMicroflowCall {
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
	searchActivities(contentMap, targetMicroflow, targetJavaAction, &calls, mf.Content)

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

func searchActivities(obj interface{}, targetMicroflow string, targetJavaAction string, calls *[]TempMicroflowCall, contentMap map[string]interface{}) {
	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this is an ActionActivity
		if typeStr, ok := v["$Type"].(string); ok {
			if typeStr == "Microflows$ActionActivity" {
				// Check for all three call types
				checkMicroflowCall(v, targetMicroflow, calls, contentMap)
				checkJavaAction(v, targetJavaAction, calls, contentMap)
				checkExternalAction(v, calls, contentMap)
			}
		}

		// Recursively search all fields
		for _, value := range v {
			searchActivities(value, targetMicroflow, targetJavaAction, calls, contentMap)
		}

	case []interface{}:
		for _, item := range v {
			searchActivities(item, targetMicroflow, targetJavaAction, calls, contentMap)
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

func checkJavaAction(activity map[string]interface{}, targetJavaAction string, calls *[]TempMicroflowCall, contentMap map[string]interface{}) {
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

	// Search in pages
	pages, err := reader.ListPages()
	if err != nil {
		return fmt.Errorf("failed to list pages: %w", err)
	}

	for _, page := range pages {
		// Load page BSON
		pageContent, err := loadUnitContents(contentsDir, string(page.ID))
		if err != nil {
			continue
		}

		// Extract module name and document name from BSON (with DB fallback)
		moduleName, docName := extractModuleAndNameFromBSON(pageContent, string(page.ID), db)
		if docName == "" {
			docName = page.Name // fallback to SDK name
		}

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

	// Search in snippets
	snippets, err := reader.ListSnippets()
	if err != nil {
		return fmt.Errorf("failed to list snippets: %w", err)
	}

	for _, snippet := range snippets {
		// Load snippet BSON
		snippetContent, err := loadUnitContents(contentsDir, string(snippet.ID))
		if err != nil {
			continue
		}

		// Extract module name and document name from BSON (with DB fallback)
		moduleName, docName := extractModuleAndNameFromBSON(snippetContent, string(snippet.ID), db)
		if docName == "" {
			docName = snippet.Name // fallback to SDK name
		}

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

	fmt.Printf("  📊 Total: %d signal subscription(s) found (%d in pages, %d in snippets)\n",
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

func collectNavigationItems(db *sql.DB, contentsDir string, report *ManifestReport, pageAccessMap map[string][]string) error {
	// Query NavigationDocument only
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
			break // Only process first NavigationDocument
		}
	}

	fmt.Printf("✅ Found %d navigation item(s)\n", len(report.NavigationItems))
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

	leafItem := NavigationItem{
		ItemName:     caption,
		Caption:      caption,
		Target:       target,
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

	fmt.Printf("✅ Found %d system role(s) and %d page(s)/snippet(s)\n", len(report.SystemRoles), len(report.PageAccess))
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
		fmt.Fprintf(file, "- **Signal Manager Widgets:** %d subscription(s)\n", len(report.Widgets))
	}
	if options.IncludeNavigation {
		fmt.Fprintf(file, "- **Navigation Items:** %d\n", len(report.NavigationItems))
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

	// Section 3: Signal Manager Widgets
	if options.IncludeWidgets {
		fmt.Fprintf(file, "## 3. Signal Manager Widgets\n\n")
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
				filter := widget.SubscriptionFilter
				if filter == "" {
					filter = "-"
				}

				fmt.Fprintf(file, "| %s | %s | %s | %s | %s | %s |\n",
					module,
					docType,
					docName,
					signalName,
					appName,
					filter,
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
			fmt.Fprintf(file, "| Parent Node | Node | User Roles |\n")
			fmt.Fprintf(file, "|-------------|------|------------|\n")

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

				roles := "-"
				if len(item.AllowedRoles) > 0 {
					roles = strings.Join(item.AllowedRoles, ", ")
				}

				fmt.Fprintf(file, "| %s | %s | %s |\n", parent, caption, roles)
			}
			fmt.Fprintf(file, "\n")
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Section 5: System Roles & Page Accessibility
	if options.IncludeRoles {
		fmt.Fprintf(file, "## 5. System Roles & Page Accessibility\n\n")

		// Part 1: System Roles List
		fmt.Fprintf(file, "### 5.1 System Roles\n\n")
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

		// Part 2: Page Accessibility - Group by Role
		fmt.Fprintf(file, "### 5.2 Page Accessibility by Role\n\n")

		// Create a map: SystemRole -> []Pages
		roleToPages := make(map[string][]string)
		for _, pageAccess := range report.PageAccess {
			if pageAccess.IsPublic {
				// Public pages accessible by all roles
				roleToPages["Public (No Restrictions)"] = append(roleToPages["Public (No Restrictions)"], fmt.Sprintf("%s (%s)", pageAccess.PageName, pageAccess.Module))
			} else {
				for _, roleName := range pageAccess.AllowedRoles {
					roleToPages[roleName] = append(roleToPages[roleName], fmt.Sprintf("%s (%s)", pageAccess.PageName, pageAccess.Module))
				}
			}
		}

		if len(roleToPages) == 0 {
			fmt.Fprintf(file, "_No page access information found._\n\n")
		} else {
			fmt.Fprintf(file, "| System Role | Accessible Pages Count | Pages |\n")
			fmt.Fprintf(file, "|-------------|------------------------|-------|\n")

			for roleName, pages := range roleToPages {
				pagesStr := strings.Join(pages, ", ")
				if len(pagesStr) > 100 {
					pagesStr = pagesStr[:100] + "..."
				}
				fmt.Fprintf(file, "| %s | %d | %s |\n", roleName, len(pages), pagesStr)
			}
			fmt.Fprintf(file, "\n")
		}

		// Part 3: Page Accessibility - View by Page
		fmt.Fprintf(file, "### 5.3 Page Accessibility by Page\n\n")

		if len(report.PageAccess) == 0 {
			fmt.Fprintf(file, "_No pages found._\n\n")
		} else {
			fmt.Fprintf(file, "Found %d page(s)/snippet(s):\n\n", len(report.PageAccess))
			fmt.Fprintf(file, "| Page Name | Module | Type | Allowed Roles | Access |\n")
			fmt.Fprintf(file, "|-----------|--------|------|---------------|--------|\n")

			for _, pageAccess := range report.PageAccess {
				rolesStr := strings.Join(pageAccess.AllowedRoles, ", ")
				if rolesStr == "" {
					rolesStr = "-"
				}
				accessType := "Restricted"
				if pageAccess.IsPublic {
					accessType = "**Public**"
				}

				fmt.Fprintf(file, "| %s | %s | %s | %s | %s |\n",
					pageAccess.PageName,
					pageAccess.Module,
					pageAccess.DocumentType,
					rolesStr,
					accessType,
				)
			}
			fmt.Fprintf(file, "\n")
		}

		fmt.Fprintf(file, "---\n\n")
	}

	// Footer
	fmt.Fprintf(file, "_Report generated by export_manifest tool_\n")

	return nil
}
