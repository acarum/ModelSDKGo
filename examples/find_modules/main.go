package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mprPath := "C:\\Workspaces\\Mendix\\MDUI\\System_Mendix_CLI\\OC EX System.mpr"
	contentsDir := "C:\\Workspaces\\Mendix\\MDUI\\System_Mendix_CLI\\mprcontents"

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", mprPath))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Query all units
	rows, err := db.Query(`SELECT UnitID, ContainerID, ContainmentName, ContentsHash FROM Unit`)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	totalUnits := 0
	unitsWithFiles := 0

	for rows.Next() {
		var unitID, containerID []byte
		var containmentName, contentsHash string

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			continue
		}

		totalUnits++

		// Convert binary ID to UUID string
		uuid := fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
			unitID[0], unitID[1], unitID[2], unitID[3],
			unitID[4], unitID[5],
			unitID[6], unitID[7],
			unitID[8], unitID[9],
			unitID[10], unitID[11], unitID[12], unitID[13], unitID[14], unitID[15])

		// Check if file exists
		cleanID := strings.ReplaceAll(uuid, "-", "")
		filePath := filepath.Join(contentsDir, cleanID[0:2], cleanID[2:4], uuid+".mxunit")

		if _, err := os.Stat(filePath); err == nil {
			unitsWithFiles++

			// Read and check type
			data, err := os.ReadFile(filePath)
			if err == nil && len(data) > 0 {
				// Simple check for "Projects$Module" in the data
				if strings.Contains(string(data), "Projects$Module") {
					fmt.Printf("Found module: %s (ContainmentName: %s)\n", uuid, containmentName)
				}
			}
		}
	}

	fmt.Printf("\nTotal units: %d\n", totalUnits)
	fmt.Printf("Units with files: %d\n", unitsWithFiles)
}
