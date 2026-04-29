// Debug tool to trace container hierarchy for pages
package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	
	_ "github.com/mattn/go-sqlite3"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_container_hierarchy <mpr-file>")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	
	// Open SQLite database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Get a sample page
	var pageUnitID, pageContainerID []byte
	var pageName string
	err = db.QueryRow(`
		SELECT UnitID, ContainerID, ContainmentName 
		FROM Unit 
		LIMIT 1
	`).Scan(&pageUnitID, &pageContainerID, &pageName)
	
	if err != nil {
		log.Fatalf("Error getting sample page: %v", err)
	}
	
	pageID := blobToUUID(pageUnitID)
	containerID := blobToUUID(pageContainerID)
	
	fmt.Printf("=== Sample Unit ===\n")
	fmt.Printf("UnitID: %s\n", pageID)
	fmt.Printf("ContainerID: %s\n", containerID)
	fmt.Printf("ContainmentName: %s\n\n", pageName)
	
	// Trace the container hierarchy
	fmt.Println("=== Container Hierarchy ===")
	currentID := pageContainerID
	level := 0
	
	for len(currentID) > 0 && level < 10 {
		var unitID, containerID []byte
		var containmentName string
		
		err := db.QueryRow(`
			SELECT UnitID, ContainerID, ContainmentName 
			FROM Unit 
			WHERE UnitID = ?
		`, currentID).Scan(&unitID, &containerID, &containmentName)
		
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("%sLevel %d: %s (ROOT - no parent found)\n", 
					strings.Repeat("  ", level), level, blobToUUID(currentID))
				break
			}
			log.Fatalf("Error tracing hierarchy: %v", err)
		}
		
		fmt.Printf("%sLevel %d: %s (ContainmentName: %s)\n", 
			strings.Repeat("  ", level), level, blobToUUID(unitID), containmentName)
		
		// Check if this is a module by querying the Contents
		// Modules have $Type starting with "Projects$Module"
		var contents []byte
		contentErr := db.QueryRow(`
			SELECT ContentsHash FROM Unit WHERE UnitID = ?
		`, unitID).Scan(&contents)
		
		if contentErr == nil {
			// For MPR v2, we can't easily read contents, but we can check ContainmentName
			if containmentName == "Modules" {
				fmt.Printf("%s  ^ This is likely a MODULE\n", strings.Repeat("  ", level))
			}
		}
		
		if len(containerID) == 0 || hex.EncodeToString(containerID) == "00000000000000000000000000000000" {
			fmt.Printf("%s  ^ This is the ROOT\n", strings.Repeat("  ", level))
			break
		}
		
		currentID = containerID
		level++
	}
	
	fmt.Println("\n=== Testing Module Detection ===")
	// Query modules directly
	rows, err := db.Query(`
		SELECT UnitID, ContainmentName 
		FROM Unit 
		WHERE ContainmentName = 'Modules'
		LIMIT 5
	`)
	if err == nil {
		defer rows.Close()
		fmt.Println("Units with ContainmentName = 'Modules':")
		for rows.Next() {
			var unitID []byte
			var cname string
			rows.Scan(&unitID, &cname)
			fmt.Printf("  - %s (%s)\n", blobToUUID(unitID), cname)
		}
	}
}
