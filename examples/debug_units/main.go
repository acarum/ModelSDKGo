package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mprPath := "C:\\Workspaces\\Mendix\\MDUI\\System_Mendix_CLI\\OC EX System.mpr"

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", mprPath))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Query the first few units to see their structure
	rows, err := db.Query(`SELECT UnitID, ContainerID, ContainmentName, ContentsHash FROM Unit LIMIT 10`)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	fmt.Println("First 10 units:")
	for rows.Next() {
		var unitID, containerID []byte
		var containmentName, contentsHash string

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			fmt.Printf("Scan error: %v\n", err)
			continue
		}

		// Convert binary ID to UUID string
		uuid := fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
			unitID[0], unitID[1], unitID[2], unitID[3],
			unitID[4], unitID[5],
			unitID[6], unitID[7],
			unitID[8], unitID[9],
			unitID[10], unitID[11], unitID[12], unitID[13], unitID[14], unitID[15])

		fmt.Printf("  UUID: %s\n", uuid)
		fmt.Printf("    ContainmentName: %s\n", containmentName)
		fmt.Printf("    ContentsHash: %s\n", contentsHash)
		fmt.Println()
	}
}
