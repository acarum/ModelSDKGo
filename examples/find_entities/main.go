package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

type EntityInfo struct {
	Name       string
	Module     string
	Location   string
	EntityType string
}

type ModuleInfo struct {
	Name string
	ID   string
}

type DomainModelInfo struct {
	Name    string
	UnitID  string
	Content map[string]interface{}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: find_entities <mpr_file_path>")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	fmt.Printf("Opening MPR: %s\n", mprPath)

	// Open MPR database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get contents directory (MPR v2)
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// Load all domain models
	domainModels, err := listDomainModels(db, contentsDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSearching for external entities in non-marketplace modules...\n\n")

	totalEntities := 0
	externalEntities := []EntityInfo{}

	// For each domain model
	for _, dm := range domainModels {
		moduleName := extractModuleName(dm.Name)

		// Skip marketplace modules
		if isMarketplaceModule(moduleName) {
			continue
		}

		// Get entities from domain model
		entities := findExternalEntitiesInDomainModel(dm, moduleName)

		if len(entities) > 0 {
			fmt.Printf("📦 Module: %s\n", moduleName)
			fmt.Printf("  External Entities: %d\n", len(entities))
			for _, entity := range entities {
				fmt.Printf("    ✓ %s\n", entity.Name)
				if entity.Location != "" {
					fmt.Printf("      Location: %s\n", entity.Location)
				}
				externalEntities = append(externalEntities, entity)
			}
			fmt.Println()
		}

		totalEntities += len(entities)
	}

	// Summary
	fmt.Println("=== Summary ===")
	fmt.Printf("Total external entities found: %d\n", totalEntities)
	fmt.Printf("Modules with external entities: %d\n", countModules(externalEntities))
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
	}

	for _, indicator := range marketplaceIndicators {
		if strings.HasPrefix(moduleName, indicator) {
			return true
		}
	}

	return false
}

func findExternalEntitiesInDomainModel(dm DomainModelInfo, moduleName string) []EntityInfo {
	var entities []EntityInfo

	// Look for Entities array in the domain model
	// BSON unmarshals arrays as bson.A (primitive.A), not []interface{}
	// But bson.A is just an alias for []interface{}, so we can convert it
	entitiesRaw, hasEntities := dm.Content["Entities"]
	if !hasEntities {
		return entities
	}

	// Convert to []interface{} regardless of the underlying type
	var entitiesArray []interface{}
	switch v := entitiesRaw.(type) {
	case []interface{}:
		entitiesArray = v
	case bson.A:
		entitiesArray = []interface{}(v)
	default:
		// Try type assertion via reflection
		entitiesArray = []interface{}(v.([]interface{}))
	}

	if entitiesArray == nil {
		return entities
	}

	for _, entityItem := range entitiesArray {
		// Skip the array length marker (first element is usually 3 or 1)
		if intVal, ok := entityItem.(int32); ok && (intVal == 1 || intVal == 2 || intVal == 3) {
			continue
		}

		// Also try int64 (long)
		if intVal, ok := entityItem.(int64); ok && (intVal == 1 || intVal == 2 || intVal == 3) {
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
					location := extractLocation(entityMap)

					if entityName != "" {
						entities = append(entities, EntityInfo{
							Name:       entityName,
							Module:     moduleName,
							Location:   location,
							EntityType: entityType,
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

	// Check for RemoteSource field (could be a reference or object)
	if remoteSource, ok := entityMap["RemoteSource"]; ok && remoteSource != nil {
		// Check if it's not an empty string or null value
		if strVal, ok := remoteSource.(string); ok && strVal != "" {
			return true
		}
		if mapVal, ok := remoteSource.(map[string]interface{}); ok && len(mapVal) > 0 {
			return true
		}
	}

	return false
}

func extractLocation(entityMap map[string]interface{}) string {
	// First check Source field (OData entities)
	if source, ok := entityMap["Source"].(map[string]interface{}); ok {
		if sourceType, ok := source["$Type"].(string); ok {
			if entityTypeName, ok := source["EntityTypeName"].(string); ok {
				return fmt.Sprintf("OData: %s (Type: %s)", entityTypeName, sourceType)
			}
			return sourceType
		}
	}

	// Fallback to other location indicators
	if location, ok := entityMap["Location"].(string); ok && !strings.Contains(location, ";") {
		return location
	}
	if remoteSource, ok := entityMap["RemoteSourceDocument"].(string); ok {
		return remoteSource
	}
	if remoteSource, ok := entityMap["RemoteSource"].(string); ok {
		return remoteSource
	}
	return "External"
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}

	// Windows GUID format: mixed-endian
	// Bytes 0-3: little-endian (reverse)
	// Bytes 4-5: little-endian (reverse)
	// Bytes 6-7: little-endian (reverse)
	// Bytes 8-15: big-endian (as-is)
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

func countModules(entities []EntityInfo) int {
	modules := make(map[string]bool)
	for _, entity := range entities {
		modules[entity.Module] = true
	}
	return len(modules)
}
