// Debug tool to check if pages have QualifiedName in the database
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_page_qualifiedname <mpr-file>")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	
	// Open SQLite database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Check what columns exist in Unit table
	fmt.Println("=== Unit Table Schema ===")
	rows, err := db.Query("PRAGMA table_info(Unit)")
	if err != nil {
		log.Fatalf("Error querying table info: %v", err)
	}
	defer rows.Close()

	var cols []string
	for rows.Next() {
		var cid int
		var name, typeName string
		var notnull, pk int
		var dfltValue interface{}
		if err := rows.Scan(&cid, &name, &typeName, &notnull, &dfltValue, &pk); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Printf("%d. %s (%s)\n", cid, name, typeName)
		cols = append(cols, name)
	}
	
	fmt.Println("\n=== Sample Page Data ===")
	
	// Query a sample page to see all available columns
	query := fmt.Sprintf("SELECT %s FROM Unit WHERE UnitID IN (SELECT UnitID FROM Unit LIMIT 1)", "*")
	row := db.QueryRow(query)
	
	// Create a slice to hold the column values
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i := range columns {
		columnPointers[i] = &columns[i]
	}
	
	if err := row.Scan(columnPointers...); err != nil {
		log.Fatalf("Error scanning row: %v", err)
	}
	
	for i, col := range cols {
		fmt.Printf("%s: %v\n", col, columns[i])
	}
	
	fmt.Println("\n=== Checking for QualifiedName in Unit table ===")
	// Try to query QualifiedName column
	testQuery := "SELECT UnitID, QualifiedName FROM Unit LIMIT 1"
	testRow := db.QueryRow(testQuery)
	var unitID, qualName string
	err = testRow.Scan(&unitID, &qualName)
	if err != nil {
		fmt.Printf("❌ QualifiedName column does NOT exist: %v\n", err)
	} else {
		fmt.Printf("✅ QualifiedName column EXISTS!\n")
		fmt.Printf("Sample: UnitID=%s, QualifiedName=%s\n", unitID, qualName)
		
		// Query all pages with QualifiedName
		fmt.Println("\n=== Pages with QualifiedName ===")
		pageRows, err := db.Query(`
			SELECT UnitID, QualifiedName, ContainmentName 
			FROM Unit 
			WHERE QualifiedName LIKE '%.%' 
			LIMIT 10
		`)
		if err != nil {
			log.Fatalf("Error querying pages: %v", err)
		}
		defer pageRows.Close()
		
		for pageRows.Next() {
			var id, qn, cn string
			if err := pageRows.Scan(&id, &qn, &cn); err != nil {
				continue
			}
			fmt.Printf("QualifiedName: %s (ContainmentName: %s)\n", qn, cn)
		}
	}
}
