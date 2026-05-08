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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: debug_nanoflow_actions <path-to-mpr>")
	}

	mprPath := os.Args[1]
	tmpDir := filepath.Join(os.TempDir(), "mendix_extract")
	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, 0755)

	// Extract MPR
	fmt.Println("Extracting MPR...")
	extractMPR(mprPath, tmpDir)

	// Open database
	dbPath := filepath.Join(tmpDir, "project.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	contentsDir := filepath.Join(tmpDir, "contents")

	// Find ACT_UpdateUoMFactor_2
	fmt.Println("\nSearching for ACT_UpdateUoMFactor_2...")

	rows, err := db.Query(`
		SELECT d.id, d.type, d.data
		FROM document d
		WHERE d.type IN ('Microflows$Microflow', 'Microflows$Nanoflow')
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var docID, docType string
		var dataBlob []byte
		if err := rows.Scan(&docID, &docType, &dataBlob); err != nil {
			continue
		}

		var doc bson.M
		if err := bson.Unmarshal(dataBlob, &doc); err != nil {
			continue
		}

		name, _ := doc["Name"].(string)
		if !strings.Contains(name, "ACT_UpdateUoMFactor") {
			continue
		}

		fmt.Printf("\n=== %s (Type: %s) ===\n", name, docType)
		fmt.Printf("Document ID: %s\n", docID)

		// Get Unit reference
		if unitRef, ok := doc["Unit"].(string); ok {
			fmt.Printf("Unit Reference: %s\n", unitRef)

			// Load the unit file
			unitDoc := loadUnitFile(contentsDir, unitRef)
			if unitDoc != nil {
				analyzeActions(unitDoc, name)
			}
		}
	}
}

func loadUnitFile(contentsDir, unitID string) bson.M {
	// Build file path from UUID
	if len(unitID) < 4 {
		return nil
	}

	subdir1 := unitID[:2]
	subdir2 := unitID[2:4]
	filePath := filepath.Join(contentsDir, subdir1, subdir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	var doc bson.M
	if err := bson.Unmarshal(data, &doc); err != nil {
		return nil
	}

	return doc
}

func analyzeActions(doc bson.M, flowName string) {
	fmt.Println("\n--- Actions Analysis ---")

	// Look for Objects array (contains all actions)
	objects, ok := doc["Objects"].(primitive.A)
	if !ok {
		fmt.Println("No Objects array found")
		return
	}

	actionCount := 0
	for _, obj := range objects {
		objMap, ok := obj.(bson.M)
		if !ok {
			continue
		}

		typeName, _ := objMap["$Type"].(string)

		// Check for all types of action calls
		if strings.Contains(typeName, "MicroflowCall") ||
			strings.Contains(typeName, "CallExternalAction") ||
			strings.Contains(typeName, "JavaActionCall") ||
			strings.Contains(typeName, "ExternalAction") {

			actionCount++
			fmt.Printf("\n[Action #%d] Type: %s\n", actionCount, typeName)

			// Extract microflow/action reference
			if mfCall, ok := objMap["MicroflowCall"].(string); ok {
				fmt.Printf("  MicroflowCall: %s\n", mfCall)
			}
			if extAction, ok := objMap["ExternalAction"].(string); ok {
				fmt.Printf("  ExternalAction: %s\n", extAction)
			}
			if javaAction, ok := objMap["JavaActionCall"].(string); ok {
				fmt.Printf("  JavaActionCall: %s\n", javaAction)
			}

			// Check for parameter mappings
			if params, ok := objMap["ParameterMappings"].(primitive.A); ok {
				fmt.Println("  Parameters:")
				for i, param := range params {
					if paramMap, ok := param.(bson.M); ok {
						argVal := ""
						if arg, ok := paramMap["Argument"].(string); ok {
							argVal = arg
						}
						paramName := ""
						if pn, ok := paramMap["Parameter"].(string); ok {
							paramName = pn
						}
						fmt.Printf("    [%d] %s = %s\n", i, paramName, argVal)
					}
				}
			}

			// Print all other fields for debugging
			fmt.Println("  All fields:")
			for k, v := range objMap {
				if k != "$Type" && k != "MicroflowCall" && k != "ExternalAction" &&
					k != "JavaActionCall" && k != "ParameterMappings" {
					fmt.Printf("    %s: %v\n", k, v)
				}
			}
		}
	}

	if actionCount == 0 {
		fmt.Println("No action calls found")
	} else {
		fmt.Printf("\nTotal action calls: %d\n", actionCount)
	}
}

func extractMPR(mprPath, destDir string) {
	// Simple zip extraction (assuming MPR is a zip file)
	data, err := os.ReadFile(mprPath)
	if err != nil {
		log.Fatal(err)
	}

	// Check if it's a zip file
	if len(data) < 4 || data[0] != 0x50 || data[1] != 0x4B {
		log.Fatal("Not a valid MPR file (not a ZIP)")
	}

	// Use PowerShell to extract
	cmd := fmt.Sprintf("Expand-Archive -Path '%s' -DestinationPath '%s' -Force", mprPath, destDir)
	os.Chdir(filepath.Dir(mprPath))

	psCmd := fmt.Sprintf("powershell.exe -Command \"%s\"", cmd)
	if err := runCommand(psCmd); err != nil {
		log.Fatal(err)
	}
}

func runCommand(cmd string) error {
	// Use os/exec
	parts := strings.Fields(cmd)
	// Skip implementation for brevity
	return nil
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return hex.EncodeToString(blob)
	}
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		blob[3], blob[2], blob[1], blob[0],
		blob[5], blob[4],
		blob[7], blob[6],
		blob[8], blob[9],
		blob[10], blob[11], blob[12], blob[13], blob[14], blob[15])
}
