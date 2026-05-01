package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run debug_page.go <mpr_path> [page_name]")
		fmt.Println("Example: go run debug_page.go \"path/to/file.mpr\" \"PageName\"")
		fmt.Println("If page_name is omitted, lists all pages")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	// Open database
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	// If no page name provided, list all pages
	if len(os.Args) < 3 {
		listAllPages(db, contentsDir)
		return
	}

	pageName := os.Args[2]
	fmt.Printf("🔍 Debugging page: %s\n", pageName)
	fmt.Printf("📦 MPR: %s\n\n", mprPath)

	// Load page
	fmt.Println("📄 Loading page...")
	pageData, err := loadPageByQualifiedNameDebug(db, contentsDir, pageName)
	if err != nil {
		fmt.Printf("❌ Error loading page: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ Page loaded successfully\n")

	// Inspect FormCall structure
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📋 FormCall Arguments:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if formCall, ok := pageData["FormCall"].(map[string]interface{}); ok {
		if args, ok := formCall["Arguments"].(primitive.A); ok {
			fmt.Printf("Found %d argument(s)\n\n", len(args)-1) // -1 for count element

			for i, arg := range args {
				if i == 0 {
					continue // Skip count
				}

				if argMap, ok := arg.(map[string]interface{}); ok {
					if param, ok := argMap["Parameter"].(string); ok {
						fmt.Printf("[Argument %d]\n", i)
						fmt.Printf("  Parameter: %s\n", param)

						// Check if this is Right placeholder
						if strings.HasSuffix(param, ".Right") {
							fmt.Printf("  ✨ This is the RIGHT placeholder!\n\n")

							// Inspect widgets
							if widgets, ok := argMap["Widgets"]; ok {
								fmt.Println("  📦 Widgets structure:")
								inspectWidgets(widgets, "    ")
							}
						} else {
							fmt.Printf("  (Not Right placeholder)\n\n")
						}
					}
				}
			}
		}
	}
}

func listAllPages(db *sql.DB, contentsDir string) {
	fmt.Println("📄 Listing all pages...")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		if typeName, ok := content["$Type"].(string); ok && typeName == "Forms$Page" {
			name := extractNameFromContents(content)
			count++
			fmt.Printf("%3d. %s\n", count, name)
		}
	}

	fmt.Printf("\n✅ Total: %d page(s)\n", count)
}

func inspectWidgets(data interface{}, indent string) {
	switch v := data.(type) {
	case primitive.A:
		fmt.Printf("%sArray with %d elements\n", indent, len(v))
		for i, item := range v {
			if i == 0 {
				fmt.Printf("%s[0] = %v (count element)\n", indent, item)
				continue
			}
			fmt.Printf("%s[%d]:\n", indent, i)
			inspectWidget(item, indent+"  ")
		}

	case []interface{}:
		fmt.Printf("%sSlice with %d elements\n", indent, len(v))
		for i, item := range v {
			fmt.Printf("%s[%d]:\n", indent, i)
			inspectWidget(item, indent+"  ")
		}

	default:
		fmt.Printf("%s%T: %v\n", indent, data, data)
	}
}

func inspectWidget(data interface{}, indent string) {
	switch v := data.(type) {
	case map[string]interface{}:
		if typeStr, ok := v["$Type"].(string); ok {
			fmt.Printf("%s$Type: %s\n", indent, typeStr)

			if name, ok := v["Name"].(string); ok {
				fmt.Printf("%sName: %s\n", indent, name)
			}

			// Check for DivContainer with class
			if typeStr == "Forms$DivContainer" {
				if appearance, ok := v["Appearance"].(map[string]interface{}); ok {
					if class, ok := appearance["Class"].(string); ok {
						if strings.Contains(class, "vertical-command-bar") {
							fmt.Printf("%s🎯 FOUND VERTICAL-COMMAND-BAR!\n", indent)
						}
					}
				}

				// Check for OnClickAction on container
				if onClickAction, ok := v["OnClickAction"].(map[string]interface{}); ok {
					actionType, _ := onClickAction["$Type"].(string)
					fmt.Printf("%s🔗 Container OnClickAction: %s\n", indent, actionType)

					// For any action type, show full structure
					if strings.Contains(actionType, "CreateObjectClientAction") || strings.Contains(actionType, "FormAction") {
						fmt.Printf("%s  Full structure:\n", indent)
						for key, val := range onClickAction {
							fmt.Printf("%s    %s: %v\n", indent, key, val)
						}
					}
				}
			}

			// Check for ActionButton
			if typeStr == "Forms$ActionButton" {
				fmt.Printf("%s🔘 ACTION BUTTON: %v\n", indent, v["Name"])

				if captionTemplate, ok := v["CaptionTemplate"].(map[string]interface{}); ok {
					if template, ok := captionTemplate["Template"].(map[string]interface{}); ok {
						if items, ok := template["Items"].(primitive.A); ok {
							for i, item := range items {
								if i == 0 {
									continue
								}
								if itemMap, ok := item.(map[string]interface{}); ok {
									langCode, _ := itemMap["LanguageCode"].(string)
									text, _ := itemMap["Text"].(string)
									fmt.Printf("%s    Translation[%s]: '%s'\n", indent, langCode, text)
								}
							}
						}
					}
				}
			}

			// Check for DynamicText - might contain actual button labels!
			if typeStr == "Forms$DynamicText" {
				name, _ := v["Name"].(string)
				if name == "text66" || name == "text69" || name == "text68" || name == "text70" || name == "text73" {
					fmt.Printf("%s📝 DYNAMIC TEXT: %v\n", indent, name)

					// Check Content field
					if content, ok := v["Content"].(map[string]interface{}); ok {
						// Check Template.Items
						if template, ok := content["Template"].(map[string]interface{}); ok {
							if items, ok := template["Items"].(primitive.A); ok {
								fmt.Printf("%s    Template.Items:\n", indent)
								for i, item := range items {
									if i == 0 {
										continue
									}
									if itemMap, ok := item.(map[string]interface{}); ok {
										langCode, _ := itemMap["LanguageCode"].(string)
										text, _ := itemMap["Text"].(string)
										fmt.Printf("%s      [%s]: '%s'\n", indent, langCode, text)
									}
								}
							}
						}

						// Check Fallback.Items
						if fallback, ok := content["Fallback"].(map[string]interface{}); ok {
							if items, ok := fallback["Items"].(primitive.A); ok {
								fmt.Printf("%s    Fallback.Items:\n", indent)
								for i, item := range items {
									if i == 0 {
										continue
									}
									if itemMap, ok := item.(map[string]interface{}); ok {
										langCode, _ := itemMap["LanguageCode"].(string)
										text, _ := itemMap["Text"].(string)
										fmt.Printf("%s      [%s]: '%s'\n", indent, langCode, text)
									}
								}
							}
						}
					}
				}
			}

			// Recursively inspect nested structures
			for key, val := range v {
				if key == "$Type" || key == "Name" || key == "Appearance" || key == "Children" {
					continue // Already printed
				}

				// Look for nested widgets
				if key == "Widgets" || key == "Items" {
					fmt.Printf("%s%s:\n", indent, key)
					inspectWidgets(val, indent+"  ")
				}
			}
		}

	default:
		fmt.Printf("%s%T\n", indent, data)
	}
}

func loadPageByQualifiedNameDebug(db *sql.DB, contentsDir string, qualifiedName string) (map[string]interface{}, error) {
	// Extract page name (support both "Module.Page" and "Page" formats)
	var pageName string
	parts := strings.Split(qualifiedName, ".")
	if len(parts) == 2 {
		pageName = parts[1]
	} else if len(parts) == 1 {
		pageName = qualifiedName
	} else {
		return nil, fmt.Errorf("invalid qualified name format: %s", qualifiedName)
	}

	fmt.Printf("Searching for page with name: %s\n", pageName)

	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		if typeName, ok := content["$Type"].(string); ok && typeName == "Forms$Page" {
			name := extractNameFromContents(content)
			if name == pageName {
				return content, nil
			}
		}
	}

	return nil, fmt.Errorf("page not found: %s", qualifiedName)
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		blob[3], blob[2], blob[1], blob[0],
		blob[5], blob[4],
		blob[7], blob[6],
		blob[8], blob[9],
		blob[10], blob[11], blob[12], blob[13], blob[14], blob[15])
}

func loadUnitContents(contentsDir string, unitID string) (map[string]interface{}, error) {
	cleanID := strings.ReplaceAll(unitID, "-", "")
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func extractNameFromContents(content map[string]interface{}) string {
	if name, ok := content["Name"].(string); ok {
		return name
	}
	if qualifiedName, ok := content["QualifiedName"].(string); ok {
		parts := strings.Split(qualifiedName, ".")
		if len(parts) > 0 {
			return parts[len(parts)-1]
		}
	}
	return ""
}
