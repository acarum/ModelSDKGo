// debug_security: dumps Security$* BSON units from an MPR file.
// Usage: go run examples/debug_security/main.go <mpr_path>
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_security <mpr_path>")
		os.Exit(1)
	}
	mprPath := os.Args[1]
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", mprPath))
	if err != nil {
		fmt.Printf("open db: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	rows, err := db.Query("SELECT UnitID FROM Unit")
	if err != nil {
		fmt.Printf("query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var idBytes []byte
		rows.Scan(&idBytes)
		uid := bytesToUUID(idBytes)
		if uid == "" {
			continue
		}
		clean := strings.ReplaceAll(uid, "-", "")
		fpath := filepath.Join(contentsDir, clean[:2], clean[2:4], uid+".mxunit")
		raw, err := os.ReadFile(fpath)
		if err != nil {
			continue
		}
		var m map[string]interface{}
		if err := bson.Unmarshal(raw, &m); err != nil {
			continue
		}
		t, _ := m["$Type"].(string)
		if !strings.Contains(t, "Security") {
			continue
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		s := string(b)
		if len(s) > 3000 {
			s = s[:3000] + "\n... (truncated)"
		}
		fmt.Printf("\n=== [%d] %s (unit: %s) ===\n%s\n", count, t, uid, s)
		count++
	}
	fmt.Printf("\nTotal security units: %d\n", count)
}

func bytesToUUID(b []byte) string {
	if len(b) != 16 {
		return ""
	}
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		b[3], b[2], b[1], b[0],
		b[5], b[4],
		b[7], b[6],
		b[8], b[9],
		b[10], b[11], b[12], b[13], b[14], b[15])
}
