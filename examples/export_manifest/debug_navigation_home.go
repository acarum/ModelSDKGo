package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run debug_navigation_home.go <path-to-mpr>")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	// Extract MPR
	tempDir, db, contentsDir, err := openMPR(mprPath)
	if err != nil {
		log.Fatalf("Failed to open MPR: %v", err)
	}
	defer os.RemoveAll(tempDir)
	defer db.Close()

	fmt.Printf("Analyzing Navigation in: %s\n\n", mprPath)

	// Find NavigationDocument
	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to query units: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var unitID []byte
		err := rows.Scan(&unitID)
		if err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		if typeName, ok := content["$Type"].(string); ok && typeName == "Navigation$NavigationDocument" {
			fmt.Println("Found NavigationDocument")
			
			// Navigate to first item (Home)
			if profiles, ok := content["Profiles"].(primitive.A); ok {
				for i, profile := range profiles {
					if i == 0 {
						continue
					}
					
					if profileMap, ok := profile.(map[string]interface{}); ok {
						if menu, ok := profileMap["Menu"].(map[string]interface{}); ok {
							if items, ok := menu["Items"].(primitive.A); ok {
								for j, item := range items {
									if j == 0 {
										continue
									}
									
									if itemMap, ok := item.(map[string]interface{}); ok {
										// Extract caption
										caption := extractCaption(itemMap)
										
										if caption == "Home" {
											fmt.Printf("\n=== HOME Navigation Item ===\n")
											
											// Print full structure
											jsonData, _ := json.MarshalIndent(itemMap, "", "  ")
											fmt.Println(string(jsonData))
											
											// Analyze Action
											if action, ok := itemMap["Action"].(map[string]interface{}); ok {
												fmt.Printf("\n=== Action Details ===\n")
												actionJSON, _ := json.MarshalIndent(action, "", "  ")
												fmt.Println(string(actionJSON))
												
												// Try to resolve target page
												if actionType, ok := action["$Type"].(string); ok {
													fmt.Printf("\nAction Type: %s\n", actionType)
													
													// Extract nanoflow reference
													if actionType == "Pages$CallNanoflowClientAction" {
														if nanoflow, ok := action["Nanoflow"].(string); ok {
															fmt.Printf("Nanoflow Reference: %s\n", nanoflow)
															
															// Load nanoflow and find ShowPage
															pages := findPagesInNanoflow(db, contentsDir, nanoflow)
															fmt.Printf("Pages found in nanoflow: %v\n", pages)
														}
													}
												}
											}
											
											return
										}
									}
								}
							}
						}
					}
					
					break
				}
			}
		}
	}
	
	fmt.Println("Home navigation item not found")
}

func extractCaption(itemMap map[string]interface{}) string {
	if captionData, ok := itemMap["Caption"].(map[string]interface{}); ok {
		if items, ok := captionData["Items"].(primitive.A); ok {
			for i, capItem := range items {
				if i == 0 {
					continue
				}
				if capMap, ok := capItem.(map[string]interface{}); ok {
					if text, ok := capMap["Text"].(string); ok {
						return text
					}
				}
			}
		}
	}
	return ""
}

func findPagesInNanoflow(db *sql.DB, contentsDir string, nanoflowRef string) []string {
	var pages []string
	
	// Query all units
	query := `SELECT UnitID FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return pages
	}
	defer rows.Close()
	
	for rows.Next() {
		var unitID []byte
		if err := rows.Scan(&unitID); err != nil {
			continue
		}
		
		unitIDStr := blobToUUID(unitID)
		
		// Try loading as UUID first
		if unitIDStr == nanoflowRef {
			content, err := loadUnitContents(contentsDir, unitIDStr)
			if err == nil {
				if typeName, ok := content["$Type"].(string); ok && typeName == "Microflows$Nanoflow" {
					fmt.Printf("\nFound nanoflow by UUID\n")
					pages = findAllShowPagesInFlow(content)
					return pages
				}
			}
		}
		
		// Try matching by name
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}
		
		if typeName, ok := content["$Type"].(string); ok && typeName == "Microflows$Nanoflow" {
			name := extractNameFromContents(content)
			if name == nanoflowRef || extractNameFromContents(content) == nanoflowRef {
				fmt.Printf("\nFound nanoflow by name: %s\n", name)
				pages = findAllShowPagesInFlow(content)
				return pages
			}
		}
	}
	
	return pages
}
