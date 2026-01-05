// Package mpr provides functionality for reading and writing Mendix project files (.mpr).
package mpr

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anthropics/modelsdk-go/domainmodel"
	"github.com/anthropics/modelsdk-go/microflows"
	"github.com/anthropics/modelsdk-go/model"
	"github.com/anthropics/modelsdk-go/pages"

	_ "github.com/mattn/go-sqlite3"
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
	path       string
	db         *sql.DB
	version    MPRVersion
	contentsDir string
	readOnly   bool
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
	// Check for expected tables
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name IN ('Unit', 'Document', 'Module')").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to query tables: %w", err)
	}
	if count == 0 {
		return errors.New("not a valid MPR file: required tables not found")
	}
	return nil
}

// GetMendixVersion returns the Mendix version used to create the project.
func (r *Reader) GetMendixVersion() (string, error) {
	var version string
	err := r.db.QueryRow("SELECT MendixVersion FROM _MetaData LIMIT 1").Scan(&version)
	if err != nil {
		return "", fmt.Errorf("failed to get Mendix version: %w", err)
	}
	return version, nil
}

// ListModules returns all modules in the project.
func (r *Reader) ListModules() ([]*model.Module, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.ContentsHash, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'Projects$Module'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query modules: %w", err)
	}
	defer rows.Close()

	var modules []*model.Module
	for rows.Next() {
		var unitID, containerID string
		var contentsHash sql.NullString
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contentsHash, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan module row: %w", err)
		}

		module, err := r.parseModule(unitID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse module %s: %w", unitID, err)
		}

		modules = append(modules, module)
	}

	return modules, nil
}

// GetModule retrieves a module by ID.
func (r *Reader) GetModule(id model.ID) (*model.Module, error) {
	var contents []byte
	err := r.db.QueryRow(`
		SELECT Contents FROM Unit WHERE UnitID = ? AND Type = 'Projects$Module'
	`, string(id)).Scan(&contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("module not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query module: %w", err)
	}

	return r.parseModule(string(id), contents)
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
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'DomainModels$DomainModel'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query domain models: %w", err)
	}
	defer rows.Close()

	var domainModels []*domainmodel.DomainModel
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan domain model row: %w", err)
		}

		dm, err := r.parseDomainModel(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse domain model %s: %w", unitID, err)
		}

		domainModels = append(domainModels, dm)
	}

	return domainModels, nil
}

// GetDomainModel retrieves a domain model by module ID.
func (r *Reader) GetDomainModel(moduleID model.ID) (*domainmodel.DomainModel, error) {
	var unitID, containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT UnitID, ContainerID, Contents
		FROM Unit
		WHERE ContainerID = ? AND Type = 'DomainModels$DomainModel'
	`, string(moduleID)).Scan(&unitID, &containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("domain model not found for module: %s", moduleID)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query domain model: %w", err)
	}

	return r.parseDomainModel(unitID, containerID, contents)
}

// ListMicroflows returns all microflows in the project.
func (r *Reader) ListMicroflows() ([]*microflows.Microflow, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'Microflows$Microflow'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query microflows: %w", err)
	}
	defer rows.Close()

	var result []*microflows.Microflow
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan microflow row: %w", err)
		}

		mf, err := r.parseMicroflow(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse microflow %s: %w", unitID, err)
		}

		result = append(result, mf)
	}

	return result, nil
}

// GetMicroflow retrieves a microflow by ID.
func (r *Reader) GetMicroflow(id model.ID) (*microflows.Microflow, error) {
	var containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT ContainerID, Contents FROM Unit WHERE UnitID = ? AND Type = 'Microflows$Microflow'
	`, string(id)).Scan(&containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("microflow not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query microflow: %w", err)
	}

	return r.parseMicroflow(string(id), containerID, contents)
}

// ListNanoflows returns all nanoflows in the project.
func (r *Reader) ListNanoflows() ([]*microflows.Nanoflow, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'Microflows$Nanoflow'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query nanoflows: %w", err)
	}
	defer rows.Close()

	var result []*microflows.Nanoflow
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan nanoflow row: %w", err)
		}

		nf, err := r.parseNanoflow(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse nanoflow %s: %w", unitID, err)
		}

		result = append(result, nf)
	}

	return result, nil
}

// GetNanoflow retrieves a nanoflow by ID.
func (r *Reader) GetNanoflow(id model.ID) (*microflows.Nanoflow, error) {
	var containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT ContainerID, Contents FROM Unit WHERE UnitID = ? AND Type = 'Microflows$Nanoflow'
	`, string(id)).Scan(&containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("nanoflow not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query nanoflow: %w", err)
	}

	return r.parseNanoflow(string(id), containerID, contents)
}

// ListPages returns all pages in the project.
func (r *Reader) ListPages() ([]*pages.Page, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'Pages$Page'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query pages: %w", err)
	}
	defer rows.Close()

	var result []*pages.Page
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan page row: %w", err)
		}

		page, err := r.parsePage(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse page %s: %w", unitID, err)
		}

		result = append(result, page)
	}

	return result, nil
}

// GetPage retrieves a page by ID.
func (r *Reader) GetPage(id model.ID) (*pages.Page, error) {
	var containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT ContainerID, Contents FROM Unit WHERE UnitID = ? AND Type = 'Pages$Page'
	`, string(id)).Scan(&containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("page not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query page: %w", err)
	}

	return r.parsePage(string(id), containerID, contents)
}

// ListLayouts returns all layouts in the project.
func (r *Reader) ListLayouts() ([]*pages.Layout, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'Pages$Layout'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query layouts: %w", err)
	}
	defer rows.Close()

	var result []*pages.Layout
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan layout row: %w", err)
		}

		layout, err := r.parseLayout(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse layout %s: %w", unitID, err)
		}

		result = append(result, layout)
	}

	return result, nil
}

// GetLayout retrieves a layout by ID.
func (r *Reader) GetLayout(id model.ID) (*pages.Layout, error) {
	var containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT ContainerID, Contents FROM Unit WHERE UnitID = ? AND Type = 'Pages$Layout'
	`, string(id)).Scan(&containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("layout not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query layout: %w", err)
	}

	return r.parseLayout(string(id), containerID, contents)
}

// ListEnumerations returns all enumerations in the project.
func (r *Reader) ListEnumerations() ([]*model.Enumeration, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'Enumerations$Enumeration'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query enumerations: %w", err)
	}
	defer rows.Close()

	var result []*model.Enumeration
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan enumeration row: %w", err)
		}

		enum, err := r.parseEnumeration(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse enumeration %s: %w", unitID, err)
		}

		result = append(result, enum)
	}

	return result, nil
}

// GetEnumeration retrieves an enumeration by ID.
func (r *Reader) GetEnumeration(id model.ID) (*model.Enumeration, error) {
	var containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT ContainerID, Contents FROM Unit WHERE UnitID = ? AND Type = 'Enumerations$Enumeration'
	`, string(id)).Scan(&containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("enumeration not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query enumeration: %w", err)
	}

	return r.parseEnumeration(string(id), containerID, contents)
}

// ListConstants returns all constants in the project.
func (r *Reader) ListConstants() ([]*model.Constant, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'Constants$Constant'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query constants: %w", err)
	}
	defer rows.Close()

	var result []*model.Constant
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan constant row: %w", err)
		}

		constant, err := r.parseConstant(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse constant %s: %w", unitID, err)
		}

		result = append(result, constant)
	}

	return result, nil
}

// GetConstant retrieves a constant by ID.
func (r *Reader) GetConstant(id model.ID) (*model.Constant, error) {
	var containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT ContainerID, Contents FROM Unit WHERE UnitID = ? AND Type = 'Constants$Constant'
	`, string(id)).Scan(&containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("constant not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query constant: %w", err)
	}

	return r.parseConstant(string(id), containerID, contents)
}

// ListScheduledEvents returns all scheduled events in the project.
func (r *Reader) ListScheduledEvents() ([]*model.ScheduledEvent, error) {
	rows, err := r.db.Query(`
		SELECT Unit.UnitID, Unit.ContainerID, Unit.Contents
		FROM Unit
		WHERE Unit.Type = 'ScheduledEvents$ScheduledEvent'
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query scheduled events: %w", err)
	}
	defer rows.Close()

	var result []*model.ScheduledEvent
	for rows.Next() {
		var unitID, containerID string
		var contents []byte

		if err := rows.Scan(&unitID, &containerID, &contents); err != nil {
			return nil, fmt.Errorf("failed to scan scheduled event row: %w", err)
		}

		event, err := r.parseScheduledEvent(unitID, containerID, contents)
		if err != nil {
			return nil, fmt.Errorf("failed to parse scheduled event %s: %w", unitID, err)
		}

		result = append(result, event)
	}

	return result, nil
}

// GetScheduledEvent retrieves a scheduled event by ID.
func (r *Reader) GetScheduledEvent(id model.ID) (*model.ScheduledEvent, error) {
	var containerID string
	var contents []byte
	err := r.db.QueryRow(`
		SELECT ContainerID, Contents FROM Unit WHERE UnitID = ? AND Type = 'ScheduledEvents$ScheduledEvent'
	`, string(id)).Scan(&containerID, &contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("scheduled event not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query scheduled event: %w", err)
	}

	return r.parseScheduledEvent(string(id), containerID, contents)
}

// GetUnit retrieves a raw unit by ID and type.
func (r *Reader) GetUnit(id model.ID, unitType string) ([]byte, error) {
	var contents []byte
	err := r.db.QueryRow(`
		SELECT Contents FROM Unit WHERE UnitID = ? AND Type = ?
	`, string(id), unitType).Scan(&contents)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("unit not found: %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query unit: %w", err)
	}

	return contents, nil
}

// ListUnits returns all units with their IDs and types.
func (r *Reader) ListUnits() ([]*UnitInfo, error) {
	rows, err := r.db.Query(`
		SELECT UnitID, ContainerID, Type FROM Unit
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	var units []*UnitInfo
	for rows.Next() {
		var unitID, containerID, unitType string
		if err := rows.Scan(&unitID, &containerID, &unitType); err != nil {
			return nil, fmt.Errorf("failed to scan unit row: %w", err)
		}

		units = append(units, &UnitInfo{
			ID:          model.ID(unitID),
			ContainerID: model.ID(containerID),
			Type:        unitType,
		})
	}

	return units, nil
}

// UnitInfo contains basic information about a unit.
type UnitInfo struct {
	ID          model.ID
	ContainerID model.ID
	Type        string
}

// ExportJSON exports the entire model as JSON.
func (r *Reader) ExportJSON() ([]byte, error) {
	modules, err := r.ListModules()
	if err != nil {
		return nil, err
	}

	domainModels, err := r.ListDomainModels()
	if err != nil {
		return nil, err
	}

	microflowsList, err := r.ListMicroflows()
	if err != nil {
		return nil, err
	}

	nanoflows, err := r.ListNanoflows()
	if err != nil {
		return nil, err
	}

	pagesList, err := r.ListPages()
	if err != nil {
		return nil, err
	}

	layouts, err := r.ListLayouts()
	if err != nil {
		return nil, err
	}

	export := map[string]interface{}{
		"modules":      modules,
		"domainModels": domainModels,
		"microflows":   microflowsList,
		"nanoflows":    nanoflows,
		"pages":        pagesList,
		"layouts":      layouts,
	}

	return json.MarshalIndent(export, "", "  ")
}
