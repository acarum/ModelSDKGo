// Debug tool to see what ListSnippets() actually returns
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/anthropics/modelsdk-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_snippet <mpr-file>")
		os.Exit(1)
	}

	reader, err := modelsdk.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	snippets, err := reader.ListSnippets()
	if err != nil {
		fmt.Printf("Error listing snippets: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d snippets\n\n", len(snippets))

	for _, snippet := range snippets {
		if snippet.Name == "MySnippet" {
			fmt.Printf("=== MySnippet from ListSnippets() ===\n\n")
			
			// Convert to JSON to see structure
			data, err := json.MarshalIndent(snippet, "", "  ")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			
			fmt.Printf("%s\n\n", string(data))
			
			// Check if Widget field is populated
			fmt.Printf("Widget field type: %T\n", snippet.Widget)
			fmt.Printf("Widget is nil: %v\n", snippet.Widget == nil)
			
			if snippet.Widget != nil {
				fmt.Printf("Widget details:\n")
				widgetData, _ := json.MarshalIndent(snippet.Widget, "  ", "  ")
				fmt.Printf("  %s\n", string(widgetData))
			}
			
			break
		}
	}
}
