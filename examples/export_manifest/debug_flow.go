package main

import (
	"database/sql"
	"encoding/json"
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
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run debug_flow.go <path_to_mpr> <flow_qualified_name>")
		fmt.Println("Example: go run debug_flow.go \"C:\\path\\to\\project.mpr\" \"MyModule.MyNanoflow\"")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	flowQualifiedName := os.Args[2]

	fmt.Printf("🔍 Debugging flow: %s\n", flowQualifiedName)
	fmt.Printf("📦 MPR: %s\n\n", mprPath)

	// Extract contents folder
	contentsDir := extractContentsDir(mprPath)
	fmt.Printf("📁 Contents directory: %s\n\n", contentsDir)

	// Open database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	fmt.Println("📄 Loading flow...")
	flowData, err := loadFlowByQualifiedName(db, contentsDir, flowQualifiedName)
	if err != nil {
		log.Fatal("Failed to load flow:", err)
	}

	if flowData == nil {
		log.Fatal("Flow not found")
	}

	fmt.Println("✅ Flow loaded successfully\n")

	// Print the entire flow structure
	prettyJSON, _ := json.MarshalIndent(flowData, "", "  ")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📋 Flow Structure:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println(string(prettyJSON))

	// Search for commands
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🔎 Searching for Command calls...")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	commands := findAllCommands(flowData)
	if len(commands) == 0 {
		fmt.Println("❌ No commands found")
	} else {
		fmt.Printf("✅ Found %d command(s):\n", len(commands))
		for i, cmd := range commands {
			fmt.Printf("  [%d] %s\n", i+1, cmd)
		}
	}
}

func extractContentsDir(mprPath string) string {
	dir := filepath.Dir(mprPath)
	base := filepath.Base(mprPath)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(dir, base, "model")
}

func loadFlowByQualifiedName(db *sql.DB, contentsDir, qualifiedName string) (map[string]interface{}, error) {
	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		data, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		if qn, ok := data["QualifiedName"].(string); ok {
			if qn == qualifiedName {
				return data, nil
			}
		}
	}

	return nil, fmt.Errorf("flow not found: %s", qualifiedName)
}

func loadUnitContents(contentsDir, unitID string) (map[string]interface{}, error) {
	filePath := filepath.Join(contentsDir, unitID)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := bson.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		blob[0:4], blob[4:6], blob[6:8], blob[8:10], blob[10:16])
}

func findAllCommands(data interface{}) []string {
	var commands []string
	seen := make(map[string]bool)

	var search func(interface{})
	search = func(d interface{}) {
		switch v := d.(type) {
		case map[string]interface{}:
			if typeStr, ok := v["$Type"].(string); ok {
				// MicroflowCall
				if typeStr == "Microflows$MicroflowCall" {
					if microflow, ok := v["Microflow"].(string); ok {
						if strings.Contains(microflow, "Command") || strings.Contains(microflow, "CallCommand") {
							// Extract CommandName parameter
							if params, ok := v["ParameterMappings"].(primitive.A); ok {
								for i, param := range params {
									if i == 0 {
										continue
									}
									if paramMap, ok := param.(map[string]interface{}); ok {
										if argValue, ok := paramMap["ArgumentValue"].(map[string]interface{}); ok {
											if constant, ok := argValue["Constant"].(string); ok {
												if !seen[constant] {
													commands = append(commands, fmt.Sprintf("MicroflowCall: %s (Command: %s)", microflow, constant))
													seen[constant] = true
												}
											}
										}
									}
								}
							}
							if !seen[microflow] {
								commands = append(commands, fmt.Sprintf("MicroflowCall: %s", microflow))
								seen[microflow] = true
							}
						}
					}
				}

				// ExternalAction
				if typeStr == "Microflows$ExternalAction" {
					if appName, ok := v["AppName"].(string); ok {
						if commandName, ok := v["CommandName"].(string); ok {
							fullCommand := appName + "." + commandName
							if !seen[fullCommand] {
								commands = append(commands, fmt.Sprintf("ExternalAction: %s", fullCommand))
								seen[fullCommand] = true
							}
						}
					}
				}

				// JavaAction
				if typeStr == "Microflows$JavaAction" {
					if actionName, ok := v["ActionName"].(string); ok {
						if strings.Contains(actionName, "Command") {
							if !seen[actionName] {
								commands = append(commands, fmt.Sprintf("JavaAction: %s", actionName))
								seen[actionName] = true
							}
						}
					}
				}
			}

			// Recurse
			for _, val := range v {
				search(val)
			}

		case primitive.A:
			for i, item := range v {
				if i == 0 {
					continue
				}
				search(item)
			}

		case []interface{}:
			for _, item := range v {
				search(item)
			}
		}
	}

	search(data)
	return commands
}
