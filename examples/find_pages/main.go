package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anthropics/modelsdk-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: find_pages <mpr_file_path>")
		fmt.Println("Example: find_pages MyApp.mpr")
		os.Exit(1)
	}

	mprPath := os.Args[1]

	// Open the MPR file
	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		log.Fatalf("Error opening MPR: %v", err)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n", reader.Version())

	// Get Mendix version
	version, err := reader.GetMendixVersion()
	if err != nil {
		fmt.Printf("Warning: Could not get Mendix version: %v\n", err)
	} else {
		fmt.Printf("Mendix Version: %s\n", version)
	}
	fmt.Println()

	// Get all pages
	pages, err := reader.ListPages()
	if err != nil {
		log.Fatalf("Error listing pages: %v", err)
	}

	fmt.Println("=== All Pages ===")
	fmt.Printf("Total: %d pages\n\n", len(pages))

	if len(pages) == 0 {
		fmt.Println("No pages found in this project.")
		return
	}

	// Print each page with details
	for i, page := range pages {
		fmt.Printf("%d. %s\n", i+1, page.Name)
		fmt.Printf("   ID: %s\n", page.ID)
		fmt.Printf("   Type: %s\n", page.TypeName)

		// Print URL if available
		if page.URL != "" {
			fmt.Printf("   URL: %s\n", page.URL)
		}

		// Print module info if available
		if page.ContainerID != "" {
			fmt.Printf("   Module ID: %s\n", page.ContainerID)
		}

		fmt.Println()
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total pages found: %d\n", len(pages))
}
