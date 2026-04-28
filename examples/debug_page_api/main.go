// Debug tool to see what GetPage() actually returns
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/anthropics/modelsdk-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_page_api <mpr-file>")
		os.Exit(1)
	}

	reader, err := modelsdk.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	pages, err := reader.ListPages()
	if err != nil {
		fmt.Printf("Error listing pages: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d pages\n\n", len(pages))

	// Find StateMachine_Details page
	for _, page := range pages {
		if page.Name == "StateMachine_Details" {
			fmt.Printf("=== StateMachine_Details from GetPage() ===\n\n")

			// Get full page details
			fullPage, err := reader.GetPage(page.ID)
			if err != nil {
				fmt.Printf("Error getting page: %v\n", err)
				continue
			}

			// Show just the top-level structure
			fmt.Printf("ID: %s\n", fullPage.ID)
			fmt.Printf("Name: %s\n", fullPage.Name)
			fmt.Printf("URL: %s\n", fullPage.URL)
			fmt.Printf("ContainerID: %s\n", fullPage.ContainerID)
			fmt.Printf("TypeName: %s\n", fullPage.TypeName)

			// Check LayoutCall
			if fullPage.LayoutCall != nil {
				fmt.Printf("\nLayoutCall present: Yes\n")
				fmt.Printf("LayoutCall.LayoutID: %s\n", fullPage.LayoutCall.LayoutID)
				fmt.Printf("LayoutCall.Arguments count: %d\n", len(fullPage.LayoutCall.Arguments))

				// Check if arguments have widgets
				for i, arg := range fullPage.LayoutCall.Arguments {
					fmt.Printf("\nArgument #%d:\n", i+1)
					fmt.Printf("  ParameterID: %s\n", arg.ParameterID)
					fmt.Printf("  Widget type: %T\n", arg.Widget)
					fmt.Printf("  Widget is nil: %v\n", arg.Widget == nil)

					if arg.Widget != nil {
						// Try to see widget details
						widgetData, _ := json.MarshalIndent(arg.Widget, "    ", "  ")
						fmt.Printf("  Widget: %s\n", string(widgetData))
					}
				}
			} else {
				fmt.Printf("\nLayoutCall: nil\n")
			}

			break
		}
	}
}
