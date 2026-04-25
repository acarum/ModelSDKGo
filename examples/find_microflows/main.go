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
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: finds_microflows <mpr_file_path> [target_microflow_name] [java_action_name] [external_action_name]")
		fmt.Println("Example: finds_microflows project.mpr EXFN_ServiceLayer.CallCommand_MF CallCommandAction UpdateStatus")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	targetMicroflow := "EXFN_ServiceLayer.CallCommand_MF"
	if len(os.Args) >= 3 {
		targetMicroflow = os.Args[2]
	}

	targetJavaAction := "CallCommandAction"
	if len(os.Args) >= 4 {
		targetJavaAction = os.Args[3]
	}

	targetExternalAction := "" // Search all by default
	if len(os.Args) >= 5 {
		targetExternalAction = os.Args[4]
	}

	fmt.Printf("Opening MPR: %s\n", mprPath)
	fmt.Printf("Searching for microflows calling: %s\n", targetMicroflow)
	fmt.Printf("Searching for Java Action: %s\n", targetJavaAction)
	if targetExternalAction != "" {
		fmt.Printf("Searching for External Action: %s\n\n", targetExternalAction)
	} else {
		fmt.Printf("Searching for all External Actions\n\n")
	}

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Failed to open MPR file: %v", err)
	}
	defer db.Close()

	// Get MPR base directory for reading .mxunit files
	mprDir := filepath.Dir(mprPath)
	contentsDir := filepath.Join(mprDir, "mprcontents")

	microflows, err := listMicroflows(db, contentsDir)
	if err != nil {
		log.Fatalf("Failed to list microflows: %v", err)
	}

	if targetExternalAction != "" {
		fmt.Printf("Found %d microflows. Searching for calls to '%s', Java Action '%s', and External Action '%s'...\n\n", len(microflows), targetMicroflow, targetJavaAction, targetExternalAction)
	} else {
		fmt.Printf("Found %d microflows. Searching for calls to '%s', Java Action '%s', and all External Actions...\n\n", len(microflows), targetMicroflow, targetJavaAction)
	}

	foundCount := 0
	for _, mf := range microflows {
		calls, err := findMicroflowCalls(mf, targetMicroflow, targetJavaAction, targetExternalAction)
		if err != nil {
			fmt.Printf("Warning: Failed to parse microflow %s: %v\n", mf.Name, err)
			continue
		}

		if len(calls) > 0 {
			foundCount++
			fmt.Printf("✓ Microflow: %s\n", mf.Name)
			fmt.Printf("  Module: %s\n", mf.ModuleName)
			fmt.Printf("  Calls found: %d\n", len(calls))
			for i, call := range calls {
				fmt.Printf("  [%d] Type: %s\n", i+1, call.CallType)
				if call.CallType == "MicroflowCall" {
					fmt.Printf("      Microflow: %s\n", call.CalledMicroflow)
				} else if call.CallType == "JavaAction" {
					fmt.Printf("      JavaAction: %s\n", call.JavaAction)
				} else if call.CallType == "ExternalAction" {
					fmt.Printf("      ExternalAction: %s\n", call.ExternalAction)
				}
				if call.ActivityName != "" {
					fmt.Printf("      Activity: %s\n", call.ActivityName)
				}
				if call.Caption != "" {
					fmt.Printf("      Caption: %s\n", call.Caption)
				}
				// Always show AppName and CommandName
				fmt.Printf("      AppName: %s\n", call.AppName)
				fmt.Printf("      CommandName: %s\n", call.CommandName)
			}
			fmt.Println()
		}
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total microflows scanned: %d\n", len(microflows))
	if targetExternalAction != "" {
		fmt.Printf("Microflows calling '%s', Java Action '%s', or External Action '%s': %d\n", targetMicroflow, targetJavaAction, targetExternalAction, foundCount)
	} else {
		fmt.Printf("Microflows calling '%s', Java Action '%s', or any External Action: %d\n", targetMicroflow, targetJavaAction, foundCount)
	}
}

type MicroflowInfo struct {
	ID         string
	Name       string
	ModuleName string
	Content    map[string]interface{}
}

type MicroflowCall struct {
	ActivityName    string
	Caption         string
	CalledMicroflow string
	JavaAction      string
	ExternalAction  string
	CallType        string // "MicroflowCall", "JavaAction", or "ExternalAction"
	AppName         string
	CommandName     string
}

func listMicroflows(db *sql.DB, contentsDir string) ([]MicroflowInfo, error) {
	query := `SELECT UnitID, ContainerID, ContainmentName, ContentsHash FROM Unit`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var microflows []MicroflowInfo
	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		// Convert UUID blob to string
		unitIDStr := blobToUUID(unitID)

		// Load content from mprcontents
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			// Skip units that can't be loaded
			continue
		}

		// Extract type from BSON content
		typeName := getTypeFromContents(content)

		// Only process Microflows
		if typeName != "Microflows$Microflow" {
			continue
		}

		// Extract name from BSON content
		name := extractNameFromContents(content)

		microflows = append(microflows, MicroflowInfo{
			ID:         unitIDStr,
			Name:       name,
			ModuleName: containmentName,
			Content:    content,
		})
	}

	return microflows, nil
}

func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return ""
	}
	// Windows GUID format: Data1 (4 bytes LE), Data2 (2 bytes LE), Data3 (2 bytes LE), Data4 (8 bytes BE)
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		blob[3], blob[2], blob[1], blob[0], // Data1 (little-endian)
		blob[5], blob[4], // Data2 (little-endian)
		blob[7], blob[6], // Data3 (little-endian)
		blob[8], blob[9], blob[10], blob[11], blob[12], blob[13], blob[14], blob[15]) // Data4 (big-endian)
}

func getTypeFromContents(contents map[string]interface{}) string {
	if typeName, ok := contents["$Type"].(string); ok {
		return typeName
	}
	return ""
}

func extractNameFromContents(contents map[string]interface{}) string {
	if name, ok := contents["Name"].(string); ok {
		return name
	}
	return ""
}

func loadUnitContents(contentsDir string, unitID string) (map[string]interface{}, error) {
	// Remove dashes from UUID for directory structure
	cleanID := strings.ReplaceAll(unitID, "-", "")
	if len(cleanID) < 4 {
		return nil, fmt.Errorf("invalid unit ID: %s", unitID)
	}

	// Build path: mprcontents/{first2}/{next2}/{uuid}.mxunit
	dir1 := cleanID[0:2]
	dir2 := cleanID[2:4]
	filePath := filepath.Join(contentsDir, dir1, dir2, unitID+".mxunit")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read unit file: %w", err)
	}

	var content map[string]interface{}
	if err := bson.Unmarshal(data, &content); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	return content, nil
}

func findMicroflowCalls(mf MicroflowInfo, targetMicroflow string, targetJavaAction string, targetExternalAction string) ([]MicroflowCall, error) {
	var calls []MicroflowCall

	// Convert content to JSON for easier searching
	jsonData, err := json.Marshal(mf.Content)
	if err != nil {
		return nil, err
	}

	var contentMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &contentMap); err != nil {
		return nil, err
	}

	// Search for ActionActivity with MicroflowCall, JavaAction, or ExternalAction
	searchActivities(contentMap, targetMicroflow, targetJavaAction, targetExternalAction, &calls)

	return calls, nil
}

func searchActivities(obj interface{}, targetMicroflow string, targetJavaAction string, targetExternalAction string, calls *[]MicroflowCall) {
	switch v := obj.(type) {
	case map[string]interface{}:
		// Check if this is an ActionActivity
		if typeStr, ok := v["$Type"].(string); ok {
			if typeStr == "Microflows$ActionActivity" {
				// Check for MicroflowCall, JavaAction, and ExternalAction
				checkMicroflowCall(v, targetMicroflow, calls)
				checkJavaAction(v, targetJavaAction, calls)
				checkExternalAction(v, targetExternalAction, calls)
			}
		}

		// Recursively search all fields
		for _, value := range v {
			searchActivities(value, targetMicroflow, targetJavaAction, targetExternalAction, calls)
		}

	case []interface{}:
		for _, item := range v {
			searchActivities(item, targetMicroflow, targetJavaAction, targetExternalAction, calls)
		}
	}
}

func checkMicroflowCall(activity map[string]interface{}, targetMicroflow string, calls *[]MicroflowCall) {
	// Get Action field
	action, ok := activity["Action"].(map[string]interface{})
	if !ok {
		return
	}

	// Check if Action is MicroflowCallAction (not MicroflowCall!)
	actionType, ok := action["$Type"].(string)
	if !ok || actionType != "Microflows$MicroflowCallAction" {
		return
	}

	// Get the MicroflowCall object from the action
	microflowCallObj, ok := action["MicroflowCall"].(map[string]interface{})
	if !ok {
		return
	}

	// Get the microflow being called
	microflowCall, ok := microflowCallObj["Microflow"].(string)
	if !ok {
		return
	}

	// Check if it matches the target (exact match or contains)
	if microflowCall == targetMicroflow || strings.Contains(microflowCall, targetMicroflow) {
		// Extract activity details
		caption := ""
		if cap, ok := activity["Caption"].(string); ok {
			caption = cap
		}

		activityName := ""
		if name, ok := activity["Name"].(string); ok {
			activityName = name
		}

		// Extract parameters (AppName and CommandName)
		appName := ""
		commandName := ""
		if paramMappings, ok := microflowCallObj["ParameterMappings"].([]interface{}); ok {
			for _, mapping := range paramMappings {
				if mappingMap, ok := mapping.(map[string]interface{}); ok {
					// Get Parameter name and Argument value
					if param, ok := mappingMap["Parameter"].(string); ok {
						argument := ""
						if arg, ok := mappingMap["Argument"].(string); ok {
							argument = arg
						}

						// Check if parameter name contains AppName or CommandName
						if strings.Contains(strings.ToLower(param), "appname") {
							appName = argument
						} else if strings.Contains(strings.ToLower(param), "commandname") {
							commandName = argument
						}
					}
				}
			}
		}

		*calls = append(*calls, MicroflowCall{
			ActivityName:    activityName,
			Caption:         caption,
			CalledMicroflow: microflowCall,
			CallType:        "MicroflowCall",
			AppName:         appName,
			CommandName:     commandName,
		})
	}
}

func checkJavaAction(activity map[string]interface{}, targetJavaAction string, calls *[]MicroflowCall) {
	// Get Action field
	action, ok := activity["Action"].(map[string]interface{})
	if !ok {
		return
	}

	// Check if Action is JavaActionCallAction
	actionType, ok := action["$Type"].(string)
	if !ok || actionType != "Microflows$JavaActionCallAction" {
		return
	}

	// Get the JavaAction being called (could be in different fields)
	javaActionName := ""

	// Try "JavaAction" field
	if ja, ok := action["JavaAction"].(string); ok {
		javaActionName = ja
	}

	// Try "JavaActionQualifiedName" field
	if javaActionName == "" {
		if ja, ok := action["JavaActionQualifiedName"].(string); ok {
			javaActionName = ja
		}
	}

	if javaActionName == "" {
		return
	}

	// Check if it matches the target (exact match or contains)
	if javaActionName == targetJavaAction || strings.Contains(javaActionName, targetJavaAction) {
		// Extract activity details
		caption := ""
		if cap, ok := activity["Caption"].(string); ok {
			caption = cap
		}

		activityName := ""
		if name, ok := activity["Name"].(string); ok {
			activityName = name
		}

		// Extract parameters (AppName and CommandName)
		appName := ""
		commandName := ""
		if paramMappings, ok := action["ParameterMappings"].([]interface{}); ok {
			for _, mapping := range paramMappings {
				if mappingMap, ok := mapping.(map[string]interface{}); ok {
					// Get Parameter name
					paramName := ""
					if param, ok := mappingMap["Parameter"].(string); ok {
						paramName = param
					}

					// Get Argument from Value object
					argument := ""
					if valueObj, ok := mappingMap["Value"].(map[string]interface{}); ok {
						if arg, ok := valueObj["Argument"].(string); ok {
							argument = strings.TrimSpace(arg) // Remove newlines and spaces
						}
					}

					// Check if parameter name contains AppName or CommandName
					if strings.Contains(strings.ToLower(paramName), "appname") {
						appName = argument
					} else if strings.Contains(strings.ToLower(paramName), "commandname") {
						commandName = argument
					}
				}
			}
		}

		*calls = append(*calls, MicroflowCall{
			ActivityName: activityName,
			Caption:      caption,
			JavaAction:   javaActionName,
			CallType:     "JavaAction",
			AppName:      appName,
			CommandName:  commandName,
		})
	}
}

func checkExternalAction(activity map[string]interface{}, targetExternalAction string, calls *[]MicroflowCall) {
	// Get Action field
	action, ok := activity["Action"].(map[string]interface{})
	if !ok {
		return
	}

	// Check if Action is CallExternalAction
	actionType, ok := action["$Type"].(string)
	if !ok || actionType != "Microflows$CallExternalAction" {
		return
	}

	// Get the Name field (used as CommandName)
	externalActionName := ""
	if name, ok := action["Name"].(string); ok {
		externalActionName = name
	}

	if externalActionName == "" {
		return
	}

	// Check if it matches the target (if target is specified)
	if targetExternalAction != "" {
		if externalActionName != targetExternalAction && !strings.Contains(externalActionName, targetExternalAction) {
			return
		}
	}

	// Extract activity details
	caption := ""
	if cap, ok := activity["Caption"].(string); ok {
		caption = cap
	}

	activityName := ""
	if actName, ok := activity["Name"].(string); ok {
		activityName = actName
	}

	// Extract AppName from ConsumedODataService (part after the ".")
	appName := ""
	if consumedService, ok := action["ConsumedODataService"].(string); ok {
		// Split by "." and take the last part
		parts := strings.Split(consumedService, ".")
		if len(parts) > 0 {
			appName = parts[len(parts)-1]
		}
	}

	*calls = append(*calls, MicroflowCall{
		ActivityName:   activityName,
		Caption:        caption,
		ExternalAction: externalActionName,
		CallType:       "ExternalAction",
		AppName:        appName,            // From ConsumedODataService (after ".")
		CommandName:    externalActionName, // Name of the action
	})
}
