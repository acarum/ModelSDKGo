package main

import (
	"fmt"
	"os"

	"github.com/anthropics/modelsdk-go"
)

func main() {
	mprPath := "C:\\Workspaces\\Mendix\\MDUI\\System_Mendix_CLI\\OC EX System.mpr"

	reader, err := modelsdk.Open(mprPath)
	if err != nil {
		fmt.Printf("Error opening: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	fmt.Printf("Opened: %s\n", reader.Path())
	fmt.Printf("MPR Version: %d\n", reader.Version())

	// Try to load a known unit
	testUnitID := "00603f05-c35e-4956-b137-57867b5bb9bc"
	fmt.Printf("\nTrying to load unit: %s\n", testUnitID)

	// This is a hack to test loadUnitContents directly
	// We'll see if we can load modules first
	modules, err := reader.ListModules()
	if err != nil {
		fmt.Printf("Error listing modules: %v\n", err)
	} else {
		fmt.Printf("Found %d modules\n", len(modules))
		for _, m := range modules {
			fmt.Printf("  - %s (ID: %s)\n", m.Name, m.ID)
		}
	}
}
