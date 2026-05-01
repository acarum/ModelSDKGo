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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run debug_page_flows.go <path_to_mpr> <page_name>")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	pageName := os.Args[2]

	fmt.Printf("🔍 Finding flows in page: %s\n", pageName)
	fmt.Printf("📦 MPR: %s\n\n", mprPath)

	contentsDir := extractContentsDir(mprPath)
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pageData, err := loadPageByName(db, contentsDir, pageName)
	if err != nil || pageData == nil {
		log.Fatal("Failed to load page")
	}

	fmt.Println("✅ Page loaded\n")

	flows := findAllFlowsWithContext(pageData)
	if len(flows) == 0 {
		fmt.Println("❌ No flows found")
	} else {
		fmt.Printf("✅ Found %d flow(s):\n\n", len(flows))
		for i, flow := range flows {
			fmt.Printf("[%d] Button: %s\n", i+1, flow.ButtonCaption)
			fmt.Printf("    Caption: %s\n", flow.Caption)
			fmt.Printf("    Flow Type: %s\n", flow.FlowType)
			fmt.Printf("    Flow Name: %s\n\n", flow.FlowName)
		}
	}
}

type FlowInfo struct {
	ButtonCaption string
	Caption       string
	FlowType      string
	FlowName      string
}

func findAllFlowsWithContext(pageData map[string]interface{}) []FlowInfo {
	var flows []FlowInfo

	var search func(interface{}, string, string)
	search = func(data interface{}, buttonName, caption string) {
		switch v := data.(type) {
		case map[string]interface{}:
			// Check for ActionButton
			if typeStr, ok := v["$Type"].(string); ok {
				if strings.Contains(typeStr, "ActionButton") {
					if name, ok := v["Name"].(string); ok {
						buttonName = name
					}

					// Try to get caption from DynamicText in same container
					if parent, ok := v["Parent"].(map[string]interface{}); ok {
						caption = extractCaptionFromParent(parent)
					}
				}
			}

			// Check for OnClickAction
			if onClickAction, ok := v["OnClickAction"].(map[string]interface{}); ok {
				if actionType, ok := onClickAction["$Type"].(string); ok {
					if strings.Contains(actionType, "CallNanoflowClientAction") {
						if nanoflow, ok := onClickAction["Nanoflow"].(string); ok {
							flows = append(flows, FlowInfo{
								ButtonCaption: buttonName,
								Caption:       caption,
								FlowType:      "Nanoflow",
								FlowName:      nanoflow,
							})
						}
					} else if strings.Contains(actionType, "CallMicroflowClientAction") {
						if microflow, ok := onClickAction["Microflow"].(string); ok {
							flows = append(flows, FlowInfo{
								ButtonCaption: buttonName,
								Caption:       caption,
								FlowType:      "Microflow",
								FlowName:      microflow,
							})
						}
					}
				}
			}

			// Recurse
			for _, val := range v {
				search(val, buttonName, caption)
			}

		case primitive.A:
			for i, item := range v {
				if i == 0 {
					continue
				}
				search(item, buttonName, caption)
			}

		case []interface{}:
			for _, item := range v {
				search(item, buttonName, caption)
			}
		}
	}

	search(pageData, "", "")
	return flows
}

func extractCaptionFromParent(parent map[string]interface{}) string {
	if widgets, ok := parent["Widgets"].(primitive.A); ok {
		for i, widget := range widgets {
			if i == 0 {
				continue
			}
			if widgetMap, ok := widget.(map[string]interface{}); ok {
				if typeStr, ok := widgetMap["$Type"].(string); ok {
					if strings.Contains(typeStr, "DynamicText") {
						if content, ok := widgetMap["Content"].(map[string]interface{}); ok {
							if template, ok := content["Template"].(map[string]interface{}); ok {
								if items, ok := template["Items"].(primitive.A); ok {
									for j, item := range items {
										if j == 0 {
											continue
										}
										if itemMap, ok := item.(map[string]interface{}); ok {
											if langCode, ok := itemMap["LanguageCode"].(string); ok {
												if langCode == "en_US" {
													if text, ok := itemMap["Text"].(string); ok {
														return text
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return ""
}

func extractContentsDir(mprPath string) string {
	dir := filepath.Dir(mprPath)
	base := filepath.Base(mprPath)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(dir, base, "model")
}

func loadPageByName(db *sql.DB, contentsDir, pageName string) (map[string]interface{}, error) {
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

		if typeStr, ok := data["$Type"].(string); ok {
			if typeStr == "Forms$Page" {
				// Try both Name and QualifiedName
				if name, ok := data["Name"].(string); ok {
					if name == pageName {
						return data, nil
					}
				}
				if qname, ok := data["QualifiedName"].(string); ok {
					if qname == pageName {
						return data, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("page not found")
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
