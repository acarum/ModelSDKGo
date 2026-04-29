// Debug tool to trace page -> folder -> module hierarchy
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}
	return fmt.Sprintf("%08x-%04x-%04x-%02x%02x-%012x",
		uint32(blob[0])<<24|uint32(blob[1])<<16|uint32(blob[2])<<8|uint32(blob[3]),
		uint16(blob[4])<<8|uint16(blob[5]),
		uint16(blob[6])<<8|uint16(blob[7]),
		blob[8], blob[9],
		uint64(blob[10])<<40|uint64(blob[11])<<32|uint64(blob[12])<<24|uint64(blob[13])<<16|uint64(blob[14])<<8|uint64(blob[15]))
}

func getTypeFromHash(mprPath string, contentsHash string) string {
	// For MPR v2, read from mprcontents folder
	mprContentsDir := strings.TrimSuffix(mprPath, ".mpr") + "-mprcontents"
	filePath := fmt.Sprintf("%s/%s.mxunit", mprContentsDir, contentsHash)
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	
	var raw map[string]interface{}
	if err := bson.Unmarshal(data, &raw); err != nil {
		return ""
	}
	
	if typeName, ok := raw["$Type"].(string); ok {
		return typeName
	}
	return ""
}

func getNameFromHash(mprPath string, contentsHash string) string {
	// For MPR v2, read from mprcontents folder
	mprContentsDir := strings.TrimSuffix(mprPath, ".mpr") + "-mprcontents"
	filePath := fmt.Sprintf("%s/%s.mxunit", mprContentsDir, contentsHash)
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	
	var raw map[string]interface{}
	if err := bson.Unmarshal(data, &raw); err != nil {
		return ""
	}
	
	if name, ok := raw["Name"].(string); ok {
		return name
	}
	return ""
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_page_to_module <mpr-file>")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	
	// Open SQLite database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Find a real page in the database
	var pageUnitID, pageContainerID []byte
	var pageContentsHash, pageContainmentName string
	err = db.QueryRow(`
		SELECT UnitID, ContainerID, ContentsHash, ContainmentName 
		FROM Unit 
		WHERE ContentsHash IS NOT NULL
		LIMIT 1
	`).Scan(&pageUnitID, &pageContainerID, &pageContentsHash, &pageContainmentName)
	
	if err != nil {
		log.Fatalf("Error getting sample page: %v", err)
	}
	
	// Check if this is actually a page
	pageType := getTypeFromHash(mprPath, pageContentsHash)
	pageName := getNameFromHash(mprPath, pageContentsHash)
	
	fmt.Printf("=== Sample Unit ===\n")
	fmt.Printf("UnitID: %s\n", blobToUUID(pageUnitID))
	fmt.Printf("Type: %s\n", pageType)
	fmt.Printf("Name: %s\n", pageName)
	fmt.Printf("ContainmentName: %s\n", pageContainmentName)
	fmt.Printf("ContainerID: %s\n\n", blobToUUID(pageContainerID))
	
	// Trace the container hierarchy
	fmt.Println("=== Container Hierarchy (bottom-up) ===")
	currentID := pageContainerID
	level := 1
	
	for len(currentID) > 0 && level < 10 {
		var unitID, containerID []byte
		var contentsHash, containmentName string
		
		err := db.QueryRow(`
			SELECT UnitID, ContainerID, ContentsHash, ContainmentName 
			FROM Unit 
			WHERE UnitID = ?
		`, currentID).Scan(&unitID, &containerID, &contentsHash, &containmentName)
		
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("%sLevel %d: %s (NOT FOUND in Unit table)\n", 
					strings.Repeat("  ", level), level, blobToUUID(currentID))
				break
			}
			log.Fatalf("Error tracing hierarchy: %v", err)
		}
		
		typeName := getTypeFromHash(mprPath, contentsHash)
		name := getNameFromHash(mprPath, contentsHash)
		
		fmt.Printf("%sLevel %d: %s\n", strings.Repeat("  ", level), level, blobToUUID(unitID))
		fmt.Printf("%s  Type: %s\n", strings.Repeat("  ", level), typeName)
		fmt.Printf("%s  Name: %s\n", strings.Repeat("  ", level), name)
		fmt.Printf("%s  ContainmentName: %s\n", strings.Repeat("  ", level), containmentName)
		
		// Check if this is a module
		if strings.HasPrefix(typeName, "Projects$Module") || containmentName == "Modules" {
			fmt.Printf("%s  ✅ THIS IS A MODULE!\n", strings.Repeat("  ", level))
			fmt.Printf("\n=== RESULT ===\n")
			fmt.Printf("Page '%s' belongs to Module '%s'\n", pageName, name)
			fmt.Printf("Qualified Name would be: %s.%s\n", name, pageName)
			break
		}
		
		if len(containerID) == 0 {
			fmt.Printf("%s  ^ This is the ROOT\n", strings.Repeat("  ", level))
			break
		}
		
		currentID = containerID
		level++
	}
}
