# ModelSDK Go

A Go library for reading and modifying Mendix projects on disk. This is a Go alternative to the [Mendix Model SDK](https://docs.mendix.com/apidocs-mxsdk/mxsdk/) and [Mendix Platform SDK](https://www.npmjs.com/package/mendixmodelsdk).

## Features

- **Read MPR Files**: Open and parse Mendix project files (.mpr) in both v1 and v2 formats
- **Type-Safe Access**: Strongly-typed Go structs for all Mendix model elements
- **Modify Models**: Create, update, and delete entities, attributes, associations, microflows, pages, and more
- **Export to JSON**: Convert the entire model to JSON format for analysis
- **No Cloud Required**: Works directly with local project files - no Mendix Platform API needed

## Installation

```bash
go get github.com/anthropics/modelsdk-go
```

**Note**: This library requires CGO for SQLite support. Make sure you have a C compiler installed.

## Quick Start

### Reading a Project

```go
package main

import (
    "fmt"
    "github.com/anthropics/modelsdk-go"
)

func main() {
    // Open a Mendix project
    reader, err := modelsdk.Open("/path/to/MyApp.mpr")
    if err != nil {
        panic(err)
    }
    defer reader.Close()

    // List all modules
    modules, _ := reader.ListModules()
    for _, m := range modules {
        fmt.Printf("Module: %s\n", m.Name)
    }

    // Get domain model for a module
    dm, _ := reader.GetDomainModel(modules[0].ID)
    for _, entity := range dm.Entities {
        fmt.Printf("  Entity: %s\n", entity.Name)
        for _, attr := range entity.Attributes {
            fmt.Printf("    - %s: %s\n", attr.Name, attr.Type.GetTypeName())
        }
    }

    // List microflows
    microflows, _ := reader.ListMicroflows()
    fmt.Printf("Total microflows: %d\n", len(microflows))

    // List pages
    pages, _ := reader.ListPages()
    fmt.Printf("Total pages: %d\n", len(pages))
}
```

### Modifying a Project

```go
package main

import (
    "github.com/anthropics/modelsdk-go"
)

func main() {
    // Open for writing
    writer, err := modelsdk.OpenForWriting("/path/to/MyApp.mpr")
    if err != nil {
        panic(err)
    }
    defer writer.Close()

    reader := writer.Reader()
    modules, _ := reader.ListModules()
    dm, _ := reader.GetDomainModel(modules[0].ID)

    // Create a new entity
    customer := modelsdk.NewEntity("Customer")
    writer.CreateEntity(dm.ID, customer)

    // Add attributes
    writer.AddAttribute(dm.ID, customer.ID, modelsdk.NewStringAttribute("Name", 200))
    writer.AddAttribute(dm.ID, customer.ID, modelsdk.NewStringAttribute("Email", 254))
    writer.AddAttribute(dm.ID, customer.ID, modelsdk.NewBooleanAttribute("IsActive"))
    writer.AddAttribute(dm.ID, customer.ID, modelsdk.NewDateTimeAttribute("CreatedDate", true))

    // Create another entity
    order := modelsdk.NewEntity("Order")
    writer.CreateEntity(dm.ID, order)

    // Create an association
    assoc := modelsdk.NewAssociation("Customer_Order", customer.ID, order.ID)
    writer.CreateAssociation(dm.ID, assoc)
}
```

## MPR File Format

Mendix projects are stored in `.mpr` files which are SQLite databases containing BSON-encoded model elements.

### MPR v1 (Mendix < 10.18)
- Single `.mpr` file containing all model data
- Documents stored as BSON blobs in SQLite

### MPR v2 (Mendix >= 10.18)
- `.mpr` file contains references and metadata
- `mprcontents/` folder contains individual document files
- Better for Git versioning and large projects

The library automatically detects and handles both formats.

## Model Structure

```
Project
├── Modules
│   ├── Domain Model
│   │   ├── Entities
│   │   │   ├── Attributes
│   │   │   ├── Indexes
│   │   │   ├── Access Rules
│   │   │   ├── Validation Rules
│   │   │   └── Event Handlers
│   │   ├── Associations
│   │   └── Annotations
│   ├── Microflows
│   │   ├── Parameters
│   │   └── Activities & Flows
│   ├── Nanoflows
│   ├── Pages
│   │   ├── Widgets
│   │   └── Data Sources
│   ├── Layouts
│   ├── Snippets
│   ├── Enumerations
│   ├── Constants
│   ├── Scheduled Events
│   └── Java Actions
└── Project Documents
```

## API Reference

### Core Types

| Type | Description |
|------|-------------|
| `modelsdk.ID` | Unique identifier for model elements (UUID) |
| `modelsdk.Module` | Represents a Mendix module |
| `modelsdk.Project` | Represents a Mendix project |
| `modelsdk.DomainModel` | Contains entities and associations |
| `modelsdk.Entity` | An entity in the domain model |
| `modelsdk.Attribute` | An attribute of an entity |
| `modelsdk.Association` | A relationship between entities |
| `modelsdk.Microflow` | A microflow (server-side logic) |
| `modelsdk.Nanoflow` | A nanoflow (client-side logic) |
| `modelsdk.Page` | A page in the UI |
| `modelsdk.Layout` | A page layout template |

### Reader Methods

```go
// Open a project
reader, _ := modelsdk.Open("path/to/project.mpr")
defer reader.Close()

// Metadata
reader.Path()                    // Get file path
reader.Version()                 // Get MPR version (1 or 2)
reader.GetMendixVersion()        // Get Mendix Studio Pro version

// Modules
reader.ListModules()             // List all modules
reader.GetModule(id)             // Get module by ID
reader.GetModuleByName(name)     // Get module by name

// Domain Models
reader.ListDomainModels()        // List all domain models
reader.GetDomainModel(moduleID)  // Get domain model for module

// Microflows & Nanoflows
reader.ListMicroflows()          // List all microflows
reader.GetMicroflow(id)          // Get microflow by ID
reader.ListNanoflows()           // List all nanoflows
reader.GetNanoflow(id)           // Get nanoflow by ID

// Pages & Layouts
reader.ListPages()               // List all pages
reader.GetPage(id)               // Get page by ID
reader.ListLayouts()             // List all layouts
reader.GetLayout(id)             // Get layout by ID

// Other
reader.ListEnumerations()        // List all enumerations
reader.ListConstants()           // List all constants
reader.ListScheduledEvents()     // List all scheduled events
reader.ExportJSON()              // Export entire model as JSON
```

### Writer Methods

```go
// Open for writing
writer, _ := modelsdk.OpenForWriting("path/to/project.mpr")
defer writer.Close()

// Access the reader
reader := writer.Reader()

// Modules
writer.CreateModule(module)
writer.UpdateModule(module)
writer.DeleteModule(id)

// Entities
writer.CreateEntity(domainModelID, entity)
writer.UpdateEntity(domainModelID, entity)
writer.DeleteEntity(domainModelID, entityID)

// Attributes
writer.AddAttribute(domainModelID, entityID, attribute)

// Associations
writer.CreateAssociation(domainModelID, association)
writer.DeleteAssociation(domainModelID, associationID)

// Microflows & Nanoflows
writer.CreateMicroflow(microflow)
writer.UpdateMicroflow(microflow)
writer.DeleteMicroflow(id)
writer.CreateNanoflow(nanoflow)
writer.UpdateNanoflow(nanoflow)
writer.DeleteNanoflow(id)

// Pages & Layouts
writer.CreatePage(page)
writer.UpdatePage(page)
writer.DeletePage(id)
writer.CreateLayout(layout)
writer.UpdateLayout(layout)
writer.DeleteLayout(id)

// Other
writer.CreateEnumeration(enumeration)
writer.CreateConstant(constant)
```

### Helper Functions

```go
// Create attributes
modelsdk.NewStringAttribute(name, length)
modelsdk.NewIntegerAttribute(name)
modelsdk.NewDecimalAttribute(name)
modelsdk.NewBooleanAttribute(name)
modelsdk.NewDateTimeAttribute(name, localize)
modelsdk.NewEnumerationAttribute(name, enumID)

// Create entities
modelsdk.NewEntity(name)                 // Persistable entity
modelsdk.NewNonPersistableEntity(name)   // Non-persistable entity

// Create associations
modelsdk.NewAssociation(name, parentID, childID)      // Reference (1:N)
modelsdk.NewReferenceSetAssociation(name, p, c)       // Reference set (M:N)

// Create flows
modelsdk.NewMicroflow(name)
modelsdk.NewNanoflow(name)

// Create pages
modelsdk.NewPage(name)

// Generate IDs
modelsdk.GenerateID()
```

## Package Structure

```
github.com/anthropics/modelsdk-go/
├── modelsdk.go          # Main package with convenience functions
├── model/               # Core model types (ID, Module, Project, etc.)
├── domainmodel/         # Domain model types (Entity, Attribute, Association)
├── microflows/          # Microflow and Nanoflow types
├── pages/               # Page, Layout, and Widget types
├── mpr/                 # MPR file reader and writer
└── examples/            # Example applications
    ├── read_project/    # Reading a Mendix project
    └── modify_project/  # Modifying a Mendix project
```

## Examples

### Read Project Information

```bash
cd examples/read_project
go run main.go /path/to/MyApp.mpr
```

### Modify Project

```bash
cd examples/modify_project
go run main.go /path/to/MyApp.mpr
```

### Inspect DataGrid Column Properties

Use the DataGrid utility in `examples/manage_datagrid` to inspect DataGrid custom widget columns and print resolved column-level properties.

Run from repository root:

```bash
go run ./examples/manage_datagrid /path/to/MyApp.mpr --showColumnProperties
```

Filter to a single page:

```bash
go run ./examples/manage_datagrid /path/to/MyApp.mpr --showColumnProperties --only-page ModuleName.PageName
```

When `--showColumnProperties` is enabled, each detected DataGrid column includes:

- `attribute`
- `dynamicText`
- `tooltip`
- `caption`
- `dateFormat`
- `customDateFormat`

Example (Windows):

```powershell
go run ./examples/manage_datagrid "C:\Workspaces\Mendix\MDUI\MxCLI\System_CLI\OC EX System.mpr" --showColumnProperties --only-page CustomModule.UoM_Master
```

**Warning**: Always backup your `.mpr` file before modifying it!

## Compatibility

- Supports Mendix Studio Pro versions 8.x, 9.x, 10.x
- Supports both MPR v1 and MPR v2 formats
- Tested on Linux, macOS, and Windows

## Comparison with Official SDK

| Feature | Mendix Model SDK (TypeScript) | modelsdk-go |
|---------|-------------------------------|-------------|
| Language | TypeScript/JavaScript | Go |
| Runtime | Node.js | Native binary |
| Cloud Required | Yes (Platform API) | No |
| Local Files | No | Yes |
| Real-time Collaboration | Yes | No |
| Read Operations | Yes | Yes |
| Write Operations | Yes | Yes |
| Type Safety | Yes (TypeScript) | Yes (Go) |

## Resources

- [Mendix Model SDK Documentation](https://docs.mendix.com/apidocs-mxsdk/mxsdk/)
- [Mendix Metamodel Documentation](https://docs.mendix.com/apidocs-mxsdk/mxsdk/mendix-metamodel/)
- [MPR File Format Discussion](https://community.mendix.com/link/space/studio-pro/questions/86892)

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
