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

type ManifestReport struct {
	ProjectName    string
	MendixVersion  string
	GeneratedAt    string
	Entities       map[string][]EntityInfo // by module
	MicroflowCalls []MicroflowCallInfo     // all calls
	Widgets        []WidgetInfo            // signal manager widgets
}

type ReportOptions struct {
	IncludeEntities   bool
	IncludeAttributes bool
	IncludeMicroflows bool
	IncludeWidgets    bool
}

func main() {
	// Define CLI flags
	includeEntities := flag.Bool("include-entities", false, "Include external entities in the report")
	includeAttributes := flag.Bool("include-attributes", true, "Include entity attributes in the report")
	includeMicroflows := flag.Bool("include-microflows", true, "Include microflow/action calls in the report")
	includeWidgets := flag.Bool("include-widgets", true, "Include Signal Manager widgets in the report")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: export_manifest [options] <mpr_file_path> <output_md_path>\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  <mpr_file_path>   Path to the Mendix MPR file\n")
		fmt.Fprintf(os.Stderr, "  <output_md_path>  Path to the output Markdown report\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  export_manifest MyApp.mpr manifest_report.md\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -include-entities=true MyApp.mpr manifest_report.md\n")
		fmt.Fprintf(os.Stderr, "  export_manifest -include-attributes=false MyApp.mpr manifest_report.md\n")
	}
	flag.Parse()

	// Check positional arguments
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	mprPath := flag.Arg(0)
	outputPath := flag.Arg(1)

	// Store global MPR path for module traversal
	globalMPRPath = mprPath

	// Create report options
	options := ReportOptions{
		IncludeEntities:   *includeEntities,
		IncludeAttributes: *includeAttributes,
		IncludeMicroflows: *includeMicroflows,
		IncludeWidgets:    *includeWidgets,
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
	fmt.Println("🔎 Scanning for external entities...")
	err = collectExternalEntities(db, contentsDir, &report)
	if err != nil {
		fmt.Printf("Error collecting entities: %v\n", err)
		os.Exit(1)
	}

	// ==== SECTION 2: Microflow/Action Calls ====
	fmt.Println("\n🔎 Scanning for microflow/action calls...")
	err = collectMicroflowCalls(db, contentsDir, &report)
	if err != nil {
		fmt.Printf("Error collecting microflow calls: %v\n", err)
		os.Exit(1)
	}

	// ==== SECTION 3: Signal Manager Widgets ====
	fmt.Println("\n🔎 Scanning for Signal Manager widgets...")
	err = collectSignalManagerWidgets(reader, mprPath, &report)
	if err != nil {
		fmt.Printf("Error collecting widgets: %v\n", err)
		os.Exit(1)
	}

	// Generate Markdown report
	fmt.Printf("\n📝 Generating report: %s\n", outputPath)
	err = generateMarkdownReport(&report, outputPath, &options)
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
				if isExternalEntity(entityMap) {
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

func extractPublishedFrom(entityMap map[string]interface{}, entityName string) string {
	// Extract from Source field (OData entities)
	if source, ok := entityMap["Source"].(map[string]interface{}); ok {
		if entityTypeName, ok := source["EntityTypeName"].(string); ok {
			// EntityTypeName usually contains the service/app name
			return entityTypeName
		}
	}

	// Fallback to other location indicators
	if remoteSource, ok := entityMap["RemoteSourceDocument"].(string); ok {
		return remoteSource
	}
	if remoteSource, ok := entityMap["RemoteSource"].(string); ok {
		return remoteSource
	}

	// Use entity name as fallback (often the entity name matches the service name)
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
	} else {
		fmt.Fprintf(file, "- **External Entities:** _Excluded from report_\n")
	}
	if options.IncludeMicroflows {
		fmt.Fprintf(file, "- **Microflow/Action Calls:** %d\n", len(report.MicroflowCalls))
	} else {
		fmt.Fprintf(file, "- **Microflow/Action Calls:** _Excluded from report_\n")
	}
	if options.IncludeWidgets {
		fmt.Fprintf(file, "- **Signal Manager Widgets:** %d subscription(s)\n\n", len(report.Widgets))
	} else {
		fmt.Fprintf(file, "- **Signal Manager Widgets:** _Excluded from report_\n\n")
	}
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
				fmt.Fprintf(file, "Found %d external entity/entities:\n\n", len(entities))

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
					} else {
						fmt.Fprintf(file, "_Attributes excluded from report_\n\n")
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

	// Footer
	fmt.Fprintf(file, "_Report generated by export_manifest tool_\n")

	return nil
}
