package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mprPath := `C:\Workspaces\Mendix\MDUI\ProductionCoordinator\Opcenter EX DS Production Coordinator.mpr`
	if len(os.Args) > 1 {
		mprPath = os.Args[1]
	}
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", mprPath))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Show distinct ContainmentName values and counts
	rows, err := db.Query(`SELECT ContainmentName, COUNT(*) as cnt FROM Unit GROUP BY ContainmentName ORDER BY cnt DESC LIMIT 40`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var cnt int
		rows.Scan(&name, &cnt)
		fmt.Printf("ContainmentName=%-50s count=%d\n", "'"+name+"'", cnt)
	}
}
