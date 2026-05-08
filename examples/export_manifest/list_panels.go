package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: list_panels <mpr_path>")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	// Check for mprcontents in same directory
	baseDir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(baseDir, "mprcontents")

	fmt.Printf("📂 Looking for mprcontents at: %s\n", contentsDir)

	// Check if it exists
	if _, err := os.Stat(contentsDir); os.IsNotExist(err) {
		// Fallback to old format
		contentsDir = strings.TrimSuffix(mprPath, ".mpr") + "-mprcontents"
		fmt.Printf("📂 Fallback to: %s\n", contentsDir)
	}

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Error opening MPR: %v", err)
	}
	defer db.Close()

	fmt.Println("🔍 Listing all panels/pages...")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	// Get all units
	query := "SELECT UnitID, ContainmentName FROM Unit"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Query error: %v", err)
	}
	defer rows.Close()

	panelCount := 0
	errorCount := 0
	totalUnits := 0

	for rows.Next() {
		totalUnits++
		var unitIDBytes []byte
		var name string
		if err := rows.Scan(&unitIDBytes, &name); err != nil {
			continue
		}

		// Convert Windows GUID bytes to string
		unitID := guidToString(unitIDBytes)
		if unitID == "" {
			continue
		}

		// Try to load BSON and check type
		content, err := loadUnitContents(contentsDir, unitID)
		if err != nil {
			errorCount++
			if errorCount == 1 {
				fmt.Printf("❌ First load error for %s: %v\n", name, err)
			}
			continue
		}

		bsonType, ok := content["$Type"].(string)
		if !ok {
			continue
		}

		// Check if it's a page or snippet
		if bsonType == "Forms$Page" || bsonType == "Forms$Snippet" {
			panelCount++
			fmt.Printf("%d. %s [%s]\n", panelCount, name, bsonType)

			if strings.Contains(name, "UpdateUoMFactor") {
				fmt.Printf("   ⭐ MATCH! UnitID: %s\n", unitID)
			}
		}
	}

	fmt.Printf("\n✅ Found %d pages/snippets out of %d total units\n", panelCount, totalUnits)
	fmt.Printf("   %d units failed to load BSON\n", errorCount)
}

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

func loadUnitContents(contentsDir, unitID string) (map[string]interface{}, error) {
	if len(unitID) < 5 {
		return nil, fmt.Errorf("invalid unitID: %s", unitID)
	}

	// Remove dashes for directory structure
	idNoDashes := strings.ReplaceAll(unitID, "-", "")
	if len(idNoDashes) < 4 {
		return nil, fmt.Errorf("invalid unitID after removing dashes: %s", unitID)
	}

	// Build path: first 2 chars / next 2 chars / full UUID with dashes.mxunit
	path := filepath.Join(contentsDir, idNoDashes[0:2], idNoDashes[2:4], unitID+".mxunit")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var result map[string]interface{}
	if err := bson.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal BSON: %w", err)
	}

	return result, nil
}
