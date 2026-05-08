package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: analyze_panel <mpr_path> <panel_name>")
		os.Exit(1)
	}

	mprPath := os.Args[1]
	panelName := os.Args[2]

	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		log.Fatalf("Error opening MPR: %v", err)
	}
	defer db.Close()

	// Check for mprcontents in same directory
	baseDir := strings.TrimSuffix(mprPath, "OC EX System.mpr")
	contentsDir := baseDir + "mprcontents"
	
	// Check if it exists
	if _, err := os.Stat(contentsDir); os.IsNotExist(err) {
		// Fallback to old format
		contentsDir = strings.TrimSuffix(mprPath, ".mpr") + "-mprcontents"
	}

	fmt.Printf("🔍 Analyzing panel: %s\n", panelName)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	// Find the panel (search by LIKE to handle partial matches)
	var unitID string
	var foundName string
	query := "SELECT UnitID, ContainmentName FROM Unit WHERE ContainmentName LIKE ?"
	rows, err := db.Query(query, "%"+panelName+"%")
	if err != nil {
		log.Fatalf("Query error: %v", err)
	}
	defer rows.Close()

	var candidates []struct {
		ID   string
		Name string
	}

	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		candidates = append(candidates, struct {
			ID   string
			Name string
		}{id, name})
	}

	if len(candidates) == 0 {
		log.Fatalf("Panel matching '%s' not found", panelName)
	}

	if len(candidates) > 1 {
		fmt.Println("⚠️  Multiple matches found:")
		for i, c := range candidates {
			fmt.Printf("  %d. %s\n", i+1, c.Name)
		}
		fmt.Println("\nUsing first match...")
	}

	unitID = candidates[0].ID
	foundName = candidates[0].Name

	fmt.Printf("✅ Found panel: %s\n", foundName)
	fmt.Printf("   UnitID: %s\n\n", unitID)

	// Load panel content
	content, err := loadUnitContents(contentsDir, unitID)
	if err != nil {
		log.Fatalf("Error loading content: %v", err)
	}

	// Extract all microflow/nanoflow references
	fmt.Println("📋 Level 0: Panel Direct References")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	refs := extractAllFlowReferences(content, 0)
	
	if len(refs) == 0 {
		fmt.Println("  ❌ No microflow/nanoflow references found")
	} else {
		for i, ref := range refs {
			fmt.Printf("  %d. %s → %s\n", i+1, ref.FieldType, ref.FlowName)
			
			// Try to analyze this flow
			analyzeFlow(db, contentsDir, ref.FlowName, 1, 5)
		}
	}
}

type FlowReference struct {
	FieldType string // "Microflow", "Nanoflow", "MicroflowCall"
	FlowName  string
}

func extractAllFlowReferences(data map[string]interface{}, depth int) []FlowReference {
	var refs []FlowReference
	seen := make(map[string]bool)

	var traverse func(v interface{}, path string)
	traverse = func(v interface{}, path string) {
		switch val := v.(type) {
		case map[string]interface{}:
			// Check for flow references
			if mf, ok := val["Microflow"].(string); ok && mf != "" && mf != "<parameter>" {
				key := "Microflow:" + mf
				if !seen[key] {
					refs = append(refs, FlowReference{"Microflow", mf})
					seen[key] = true
				}
			}
			if nf, ok := val["Nanoflow"].(string); ok && nf != "" && nf != "<parameter>" {
				key := "Nanoflow:" + nf
				if !seen[key] {
					refs = append(refs, FlowReference{"Nanoflow", nf})
					seen[key] = true
				}
			}
			if mc, ok := val["MicroflowCall"].(string); ok && mc != "" && mc != "<parameter>" {
				key := "MicroflowCall:" + mc
				if !seen[key] {
					refs = append(refs, FlowReference{"MicroflowCall", mc})
					seen[key] = true
				}
			}

			// Recurse
			for k, nested := range val {
				traverse(nested, path+"."+k)
			}
		case primitive.A:
			for i, item := range val {
				traverse(item, fmt.Sprintf("%s[%d]", path, i))
			}
		case []interface{}:
			for i, item := range val {
				traverse(item, fmt.Sprintf("%s[%d]", path, i))
			}
		}
	}

	traverse(data, "root")
	return refs
}

func analyzeFlow(db *sql.DB, contentsDir, flowName string, level, maxLevel int) {
	if level > maxLevel {
		return
	}

	indent := strings.Repeat("  ", level)
	
	// Parse module.name
	parts := strings.Split(flowName, ".")
	if len(parts) != 2 {
		fmt.Printf("%s  └─ ⚠️  Invalid format: %s\n", indent, flowName)
		return
	}

	moduleName := parts[0]
	mfName := parts[1]

	// Find the flow in database
	var unitID string
	query := "SELECT UnitID FROM Unit WHERE ContainmentName = ?"
	rows, err := db.Query(query, mfName)
	if err != nil {
		fmt.Printf("%s  └─ ❌ Query error: %v\n", indent, err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var candidateID string
		if err := rows.Scan(&candidateID); err != nil {
			continue
		}

		// Check module
		candidateModule := findModuleByTraversal(candidateID, db)
		if candidateModule == moduleName {
			unitID = candidateID
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("%s  └─ ⚠️  Not found in Unit table (might be inline nanoflow)\n", indent)
		return
	}

	fmt.Printf("%s  └─ ✅ Found as Unit: %s\n", indent, unitID)

	// Load flow content
	content, err := loadUnitContents(contentsDir, unitID)
	if err != nil {
		fmt.Printf("%s      └─ ❌ Load error: %v\n", indent, err)
		return
	}

	// Extract what this flow calls
	fmt.Printf("\n%s📋 Level %d: %s calls:\n", indent, level, mfName)
	fmt.Printf("%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n", indent)
	
	refs := extractAllFlowReferences(content, level)
	
	if len(refs) == 0 {
		fmt.Printf("%s  (no further calls)\n", indent)
	} else {
		for i, ref := range refs {
			fmt.Printf("%s  %d. %s → %s\n", indent, i+1, ref.FieldType, ref.FlowName)
			
			// Check if this is a command call
			if ref.FieldType == "MicroflowCall" || ref.FieldType == "Microflow" {
				// Recurse
				analyzeFlow(db, contentsDir, ref.FlowName, level+1, maxLevel)
			}
		}
	}
	fmt.Println()
}

func findModuleByTraversal(unitID string, db *sql.DB) string {
	currentID := unitID
	for {
		var containerID sql.NullString
		var containmentName string
		query := "SELECT ContainerID, ContainmentName FROM Unit WHERE UnitID = ?"
		err := db.QueryRow(query, currentID).Scan(&containerID, &containmentName)
		if err != nil {
			return ""
		}

		if !containerID.Valid {
			return containmentName
		}

		currentID = containerID.String
	}
}

func loadUnitContents(contentsDir, unitID string) (map[string]interface{}, error) {
	if len(unitID) < 5 {
		return nil, fmt.Errorf("invalid unitID: %s", unitID)
	}

	path := fmt.Sprintf("%s/%s/%s/%s.mxunit", contentsDir, unitID[0:2], unitID[2:4], unitID)
	
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
