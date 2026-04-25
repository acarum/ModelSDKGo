package main

import (
	"fmt"
	"os"

	"github.com/anthropics/modelsdk-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: list_document_types <path-to-mpr-file>")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n\n", reader.Version())

	// Get all modules
	modules, err := reader.ListModules()
	if err != nil {
		fmt.Printf("Error listing modules: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d modules:\n\n", len(modules))

	for _, m := range modules {
		fmt.Printf("Module: %s (ID: %s)\n", m.Name, m.ID)

		// Try to get pages for this module
		fmt.Println("  Checking for pages...")
		pages, err := reader.ListPages()
		if err == nil {
			modulePages := 0
			for _, p := range pages {
				if string(p.ContainerID) == string(m.ID) {
					modulePages++
					if modulePages <= 3 {
						fmt.Printf("    - %s\n", p.Name)
					}
				}
			}
			if modulePages > 0 {
				fmt.Printf("  Total pages: %d\n", modulePages)
			} else {
				fmt.Println("  No pages found")
			}
		}

		fmt.Println()
	}
}
