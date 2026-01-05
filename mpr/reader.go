// Package mpr provides functionality for reading and writing Mendix project files (.mpr).
package mpr

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthropics/modelsdk-go/domainmodel"
	"github.com/anthropics/modelsdk-go/microflows"
	"github.com/anthropics/modelsdk-go/model"
	"github.com/anthropics/modelsdk-go/pages"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
)

// MPRVersion represents the MPR file format version.
type MPRVersion int

const (
	// MPRVersionV1 is the original single-file format.
	MPRVersionV1 MPRVersion = 1
	// MPRVersionV2 uses mprcontents folder (Mendix 10.18+).
	MPRVersionV2 MPRVersion = 2
)

// Reader provides methods to read Mendix project files.
type Reader struct {
	path        string
	db          *sql.DB
	version     MPRVersion
	contentsDir string
	readOnly    bool
}

// OpenOptions configures how the MPR file is opened.
type OpenOptions struct {
	// ReadOnly opens the database in read-only mode.
	ReadOnly bool
}

// Open opens an MPR file for reading.
func Open(path string) (*Reader, error) {
	return OpenWithOptions(path, OpenOptions{ReadOnly: true})
}

// OpenWithOptions opens an MPR file with the specified options.
func OpenWithOptions(path string, opts OpenOptions) (*Reader, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("mpr file not found: %s", path)
	}

	r := &Reader{
		path:     path,
		readOnly: opts.ReadOnly,
	}

	// Check for MPR v2 (mprcontents folder)
	dir := filepath.Dir(path)
	contentsDir := filepath.Join(dir, "mprcontents")
	if stat, err := os.Stat(contentsDir); err == nil && stat.IsDir() {
		r.version = MPRVersionV2
		r.contentsDir = contentsDir
	} else {
		r.version = MPRVersionV1
	}

	// Open SQLite database
	dsn := path
	if opts.ReadOnly {
		dsn = fmt.Sprintf("file:%s?mode=ro", path)
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	r.db = db

	// Verify it's a valid MPR file
	if err := r.verify(); err != nil {
		r.Close()
		return nil, err
	}

	return r, nil
}

// Close closes the reader and releases resources.
func (r *Reader) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

// Path returns the path to the MPR file.
func (r *Reader) Path() string {
	return r.path
}

// Version returns the MPR file format version.
func (r *Reader) Version() MPRVersion {
	return r.version
}

// verify checks that the file is a valid MPR database.
func (r *Reader) verify() error {
	// Check for Unit table which is required
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name = 'Unit'").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to query tables: %w", err)
	}
	if count == 0 {
		return errors.New("not a valid MPR file: Unit table not found")
	}
	return nil
}

// GetMendixVersion returns the Mendix version used to create the project.
func (r *Reader) GetMendixVersion() (string, error) {
	var version string
	// Try new schema first
	err := r.db.QueryRow("SELECT _ProductVersion FROM _MetaData LIMIT 1").Scan(&version)
	if err != nil {
		// Try old schema
		err = r.db.QueryRow("SELECT MendixVersion FROM _MetaData LIMIT 1").Scan(&version)
		if err != nil {
			return "", fmt.Errorf("failed to get Mendix version: %w", err)
		}
	}
	return version, nil
}

// blobToUUID converts a 16-byte blob to a UUID string.
func blobToUUID(blob []byte) string {
	if len(blob) != 16 {
		return hex.EncodeToString(blob)
	}
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		blob[0:4], blob[4:6], blob[6:8], blob[8:10], blob[10:16])
}

// getTypeFromContents extracts the $Type field from BSON contents.
func getTypeFromContents(contents []byte) string {
	if len(contents) == 0 {
		return ""
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return ""
	}

	if typeName, ok := raw["$Type"].(string); ok {
		return typeName
	}
	return ""
}

// listUnitsByType returns all units matching the given type prefix.
func (r *Reader) listUnitsByType(typePrefix string) ([]rawUnit, error) {
	rows, err := r.db.Query(`
		SELECT UnitID, ContainerID, ContainmentName, Contents
		FROM Unit
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	var units []rawUnit
	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan unit row: %w", err)
		}

		typeName := getTypeFromContents(contents)
		if typePrefix == "" || strings.HasPrefix(typeName, typePrefix) {
			units = append(units, rawUnit{
				ID:              blobToUUID(unitID),
				ContainerID:     blobToUUID(containerID),
				ContainmentName: containmentName,
				Type:            typeName,
				Contents:        contents,
			})
		}
	}

	return units, nil
}

// rawUnit holds raw unit data from the database.
type rawUnit struct {
	ID              string
	ContainerID     string
	ContainmentName string
	Type            string
	Contents        []byte
}

// ListModules returns all modules in the project.
func (r *Reader) ListModules() ([]*model.Module, error) {
	units, err := r.listUnitsByType("Projects$Module")
	if err != nil {
		return nil, err
	}

	var modules []*model.Module
	for _, u := range units {
		module, err := r.parseModule(u.ID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse module %s: %w", u.ID, err)
		}
		modules = append(modules, module)
	}

	return modules, nil
}

// GetModule retrieves a module by ID.
func (r *Reader) GetModule(id model.ID) (*model.Module, error) {
	modules, err := r.ListModules()
	if err != nil {
		return nil, err
	}

	for _, m := range modules {
		if m.ID == id {
			return m, nil
		}
	}

	return nil, fmt.Errorf("module not found: %s", id)
}

// GetModuleByName retrieves a module by name.
func (r *Reader) GetModuleByName(name string) (*model.Module, error) {
	modules, err := r.ListModules()
	if err != nil {
		return nil, err
	}

	for _, m := range modules {
		if m.Name == name {
			return m, nil
		}
	}

	return nil, fmt.Errorf("module not found: %s", name)
}

// ListDomainModels returns all domain models in the project.
func (r *Reader) ListDomainModels() ([]*domainmodel.DomainModel, error) {
	units, err := r.listUnitsByType("DomainModels$DomainModel")
	if err != nil {
		return nil, err
	}

	var domainModels []*domainmodel.DomainModel
	for _, u := range units {
		dm, err := r.parseDomainModel(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse domain model %s: %w", u.ID, err)
		}
		domainModels = append(domainModels, dm)
	}

	return domainModels, nil
}

// GetDomainModel retrieves a domain model by module ID.
func (r *Reader) GetDomainModel(moduleID model.ID) (*domainmodel.DomainModel, error) {
	domainModels, err := r.ListDomainModels()
	if err != nil {
		return nil, err
	}

	for _, dm := range domainModels {
		if dm.ContainerID == moduleID {
			return dm, nil
		}
	}

	return nil, fmt.Errorf("domain model not found for module: %s", moduleID)
}

// ListMicroflows returns all microflows in the project.
func (r *Reader) ListMicroflows() ([]*microflows.Microflow, error) {
	units, err := r.listUnitsByType("Microflows$Microflow")
	if err != nil {
		return nil, err
	}

	var result []*microflows.Microflow
	for _, u := range units {
		mf, err := r.parseMicroflow(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse microflow %s: %w", u.ID, err)
		}
		result = append(result, mf)
	}

	return result, nil
}

// GetMicroflow retrieves a microflow by ID.
func (r *Reader) GetMicroflow(id model.ID) (*microflows.Microflow, error) {
	microflowsList, err := r.ListMicroflows()
	if err != nil {
		return nil, err
	}

	for _, mf := range microflowsList {
		if mf.ID == id {
			return mf, nil
		}
	}

	return nil, fmt.Errorf("microflow not found: %s", id)
}

// ListNanoflows returns all nanoflows in the project.
func (r *Reader) ListNanoflows() ([]*microflows.Nanoflow, error) {
	units, err := r.listUnitsByType("Microflows$Nanoflow")
	if err != nil {
		return nil, err
	}

	var result []*microflows.Nanoflow
	for _, u := range units {
		nf, err := r.parseNanoflow(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse nanoflow %s: %w", u.ID, err)
		}
		result = append(result, nf)
	}

	return result, nil
}

// GetNanoflow retrieves a nanoflow by ID.
func (r *Reader) GetNanoflow(id model.ID) (*microflows.Nanoflow, error) {
	nanoflows, err := r.ListNanoflows()
	if err != nil {
		return nil, err
	}

	for _, nf := range nanoflows {
		if nf.ID == id {
			return nf, nil
		}
	}

	return nil, fmt.Errorf("nanoflow not found: %s", id)
}

// ListPages returns all pages in the project.
func (r *Reader) ListPages() ([]*pages.Page, error) {
	units, err := r.listUnitsByType("Pages$Page")
	if err != nil {
		return nil, err
	}

	var result []*pages.Page
	for _, u := range units {
		page, err := r.parsePage(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse page %s: %w", u.ID, err)
		}
		result = append(result, page)
	}

	return result, nil
}

// GetPage retrieves a page by ID.
func (r *Reader) GetPage(id model.ID) (*pages.Page, error) {
	pagesList, err := r.ListPages()
	if err != nil {
		return nil, err
	}

	for _, p := range pagesList {
		if p.ID == id {
			return p, nil
		}
	}

	return nil, fmt.Errorf("page not found: %s", id)
}

// ListLayouts returns all layouts in the project.
func (r *Reader) ListLayouts() ([]*pages.Layout, error) {
	units, err := r.listUnitsByType("Pages$Layout")
	if err != nil {
		return nil, err
	}

	var result []*pages.Layout
	for _, u := range units {
		layout, err := r.parseLayout(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse layout %s: %w", u.ID, err)
		}
		result = append(result, layout)
	}

	return result, nil
}

// GetLayout retrieves a layout by ID.
func (r *Reader) GetLayout(id model.ID) (*pages.Layout, error) {
	layouts, err := r.ListLayouts()
	if err != nil {
		return nil, err
	}

	for _, l := range layouts {
		if l.ID == id {
			return l, nil
		}
	}

	return nil, fmt.Errorf("layout not found: %s", id)
}

// ListEnumerations returns all enumerations in the project.
func (r *Reader) ListEnumerations() ([]*model.Enumeration, error) {
	units, err := r.listUnitsByType("Enumerations$Enumeration")
	if err != nil {
		return nil, err
	}

	var result []*model.Enumeration
	for _, u := range units {
		enum, err := r.parseEnumeration(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse enumeration %s: %w", u.ID, err)
		}
		result = append(result, enum)
	}

	return result, nil
}

// GetEnumeration retrieves an enumeration by ID.
func (r *Reader) GetEnumeration(id model.ID) (*model.Enumeration, error) {
	enums, err := r.ListEnumerations()
	if err != nil {
		return nil, err
	}

	for _, e := range enums {
		if e.ID == id {
			return e, nil
		}
	}

	return nil, fmt.Errorf("enumeration not found: %s", id)
}

// ListConstants returns all constants in the project.
func (r *Reader) ListConstants() ([]*model.Constant, error) {
	units, err := r.listUnitsByType("Constants$Constant")
	if err != nil {
		return nil, err
	}

	var result []*model.Constant
	for _, u := range units {
		constant, err := r.parseConstant(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse constant %s: %w", u.ID, err)
		}
		result = append(result, constant)
	}

	return result, nil
}

// GetConstant retrieves a constant by ID.
func (r *Reader) GetConstant(id model.ID) (*model.Constant, error) {
	constants, err := r.ListConstants()
	if err != nil {
		return nil, err
	}

	for _, c := range constants {
		if c.ID == id {
			return c, nil
		}
	}

	return nil, fmt.Errorf("constant not found: %s", id)
}

// ListScheduledEvents returns all scheduled events in the project.
func (r *Reader) ListScheduledEvents() ([]*model.ScheduledEvent, error) {
	units, err := r.listUnitsByType("ScheduledEvents$ScheduledEvent")
	if err != nil {
		return nil, err
	}

	var result []*model.ScheduledEvent
	for _, u := range units {
		event, err := r.parseScheduledEvent(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse scheduled event %s: %w", u.ID, err)
		}
		result = append(result, event)
	}

	return result, nil
}

// GetScheduledEvent retrieves a scheduled event by ID.
func (r *Reader) GetScheduledEvent(id model.ID) (*model.ScheduledEvent, error) {
	events, err := r.ListScheduledEvents()
	if err != nil {
		return nil, err
	}

	for _, e := range events {
		if e.ID == id {
			return e, nil
		}
	}

	return nil, fmt.Errorf("scheduled event not found: %s", id)
}

// ListSnippets returns all snippets in the project.
func (r *Reader) ListSnippets() ([]*pages.Snippet, error) {
	units, err := r.listUnitsByType("Pages$Snippet")
	if err != nil {
		return nil, err
	}

	var result []*pages.Snippet
	for _, u := range units {
		snippet, err := r.parseSnippet(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse snippet %s: %w", u.ID, err)
		}
		result = append(result, snippet)
	}

	return result, nil
}

// ListJavaActions returns all Java actions in the project.
func (r *Reader) ListJavaActions() ([]*JavaAction, error) {
	units, err := r.listUnitsByType("JavaActions$JavaAction")
	if err != nil {
		return nil, err
	}

	var result []*JavaAction
	for _, u := range units {
		ja, err := r.parseJavaAction(u.ID, u.ContainerID, u.Contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse java action %s: %w", u.ID, err)
		}
		result = append(result, ja)
	}

	return result, nil
}

// JavaAction represents a Java action.
type JavaAction struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`
}

// GetName returns the Java action's name.
func (ja *JavaAction) GetName() string {
	return ja.Name
}

// GetContainerID returns the container ID.
func (ja *JavaAction) GetContainerID() model.ID {
	return ja.ContainerID
}

// ListUnits returns all units with their IDs and types.
func (r *Reader) ListUnits() ([]*UnitInfo, error) {
	units, err := r.listUnitsByType("")
	if err != nil {
		return nil, err
	}

	var result []*UnitInfo
	for _, u := range units {
		result = append(result, &UnitInfo{
			ID:              model.ID(u.ID),
			ContainerID:     model.ID(u.ContainerID),
			ContainmentName: u.ContainmentName,
			Type:            u.Type,
		})
	}

	return result, nil
}

// UnitInfo contains basic information about a unit.
type UnitInfo struct {
	ID              model.ID
	ContainerID     model.ID
	ContainmentName string
	Type            string
}

// ExportJSON exports the entire model as JSON.
func (r *Reader) ExportJSON() ([]byte, error) {
	modules, err := r.ListModules()
	if err != nil {
		modules = nil // Continue even if modules fail
	}

	domainModels, err := r.ListDomainModels()
	if err != nil {
		domainModels = nil
	}

	microflowsList, err := r.ListMicroflows()
	if err != nil {
		microflowsList = nil
	}

	nanoflows, err := r.ListNanoflows()
	if err != nil {
		nanoflows = nil
	}

	pagesList, err := r.ListPages()
	if err != nil {
		pagesList = nil
	}

	layouts, err := r.ListLayouts()
	if err != nil {
		layouts = nil
	}

	enumerations, err := r.ListEnumerations()
	if err != nil {
		enumerations = nil
	}

	constants, err := r.ListConstants()
	if err != nil {
		constants = nil
	}

	export := map[string]interface{}{
		"modules":      modules,
		"domainModels": domainModels,
		"microflows":   microflowsList,
		"nanoflows":    nanoflows,
		"pages":        pagesList,
		"layouts":      layouts,
		"enumerations": enumerations,
		"constants":    constants,
	}

	return json.MarshalIndent(export, "", "  ")
}

// GetUnitTypes returns a count of units by type.
func (r *Reader) GetUnitTypes() (map[string]int, error) {
	units, err := r.listUnitsByType("")
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	for _, u := range units {
		counts[u.Type]++
	}

	return counts, nil
}
