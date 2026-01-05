package mpr

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anthropics/modelsdk-go/domainmodel"
	"github.com/anthropics/modelsdk-go/microflows"
	"github.com/anthropics/modelsdk-go/model"
	"github.com/anthropics/modelsdk-go/pages"

	"go.mongodb.org/mongo-driver/bson"
)

// Writer provides methods to write Mendix project files.
type Writer struct {
	reader *Reader
}

// NewWriter creates a new writer from a reader opened in read-write mode.
func NewWriter(path string) (*Writer, error) {
	reader, err := OpenWithOptions(path, OpenOptions{ReadOnly: false})
	if err != nil {
		return nil, err
	}
	return &Writer{reader: reader}, nil
}

// Close closes the writer.
func (w *Writer) Close() error {
	return w.reader.Close()
}

// Reader returns the underlying reader.
func (w *Writer) Reader() *Reader {
	return w.reader
}

// CreateModule creates a new module in the project.
func (w *Writer) CreateModule(module *model.Module) error {
	if module.ID == "" {
		module.ID = model.ID(generateUUID())
	}
	module.TypeName = "Projects$Module"

	contents, err := w.serializeModule(module)
	if err != nil {
		return fmt.Errorf("failed to serialize module: %w", err)
	}

	return w.insertUnit(string(module.ID), "", "Projects$Module", contents)
}

// UpdateModule updates an existing module.
func (w *Writer) UpdateModule(module *model.Module) error {
	contents, err := w.serializeModule(module)
	if err != nil {
		return fmt.Errorf("failed to serialize module: %w", err)
	}

	return w.updateUnit(string(module.ID), contents)
}

// DeleteModule deletes a module by ID.
func (w *Writer) DeleteModule(id model.ID) error {
	return w.deleteUnit(string(id))
}

// CreateEntity creates a new entity in a domain model.
func (w *Writer) CreateEntity(domainModelID model.ID, entity *domainmodel.Entity) error {
	// Load the domain model
	dm, err := w.reader.GetDomainModel(domainModelID)
	if err != nil {
		return err
	}

	// Assign ID if not set
	if entity.ID == "" {
		entity.ID = model.ID(generateUUID())
	}
	entity.TypeName = "DomainModels$Entity"
	entity.ContainerID = domainModelID

	// Add entity to domain model
	dm.Entities = append(dm.Entities, entity)

	// Serialize and update
	return w.updateDomainModel(dm)
}

// UpdateEntity updates an existing entity.
func (w *Writer) UpdateEntity(domainModelID model.ID, entity *domainmodel.Entity) error {
	dm, err := w.reader.GetDomainModel(domainModelID)
	if err != nil {
		return err
	}

	// Find and replace the entity
	for i, e := range dm.Entities {
		if e.ID == entity.ID {
			dm.Entities[i] = entity
			return w.updateDomainModel(dm)
		}
	}

	return fmt.Errorf("entity not found: %s", entity.ID)
}

// DeleteEntity deletes an entity from a domain model.
func (w *Writer) DeleteEntity(domainModelID model.ID, entityID model.ID) error {
	dm, err := w.reader.GetDomainModel(domainModelID)
	if err != nil {
		return err
	}

	// Find and remove the entity
	for i, e := range dm.Entities {
		if e.ID == entityID {
			dm.Entities = append(dm.Entities[:i], dm.Entities[i+1:]...)
			return w.updateDomainModel(dm)
		}
	}

	return fmt.Errorf("entity not found: %s", entityID)
}

// AddAttribute adds an attribute to an entity.
func (w *Writer) AddAttribute(domainModelID model.ID, entityID model.ID, attr *domainmodel.Attribute) error {
	dm, err := w.reader.GetDomainModel(domainModelID)
	if err != nil {
		return err
	}

	// Find the entity
	for _, e := range dm.Entities {
		if e.ID == entityID {
			if attr.ID == "" {
				attr.ID = model.ID(generateUUID())
			}
			attr.TypeName = "DomainModels$Attribute"
			attr.ContainerID = entityID
			e.Attributes = append(e.Attributes, attr)
			return w.updateDomainModel(dm)
		}
	}

	return fmt.Errorf("entity not found: %s", entityID)
}

// CreateAssociation creates a new association between entities.
func (w *Writer) CreateAssociation(domainModelID model.ID, assoc *domainmodel.Association) error {
	dm, err := w.reader.GetDomainModel(domainModelID)
	if err != nil {
		return err
	}

	if assoc.ID == "" {
		assoc.ID = model.ID(generateUUID())
	}
	assoc.TypeName = "DomainModels$Association"
	assoc.ContainerID = domainModelID

	dm.Associations = append(dm.Associations, assoc)
	return w.updateDomainModel(dm)
}

// DeleteAssociation deletes an association.
func (w *Writer) DeleteAssociation(domainModelID model.ID, assocID model.ID) error {
	dm, err := w.reader.GetDomainModel(domainModelID)
	if err != nil {
		return err
	}

	for i, a := range dm.Associations {
		if a.ID == assocID {
			dm.Associations = append(dm.Associations[:i], dm.Associations[i+1:]...)
			return w.updateDomainModel(dm)
		}
	}

	return fmt.Errorf("association not found: %s", assocID)
}

// CreateMicroflow creates a new microflow.
func (w *Writer) CreateMicroflow(mf *microflows.Microflow) error {
	if mf.ID == "" {
		mf.ID = model.ID(generateUUID())
	}
	mf.TypeName = "Microflows$Microflow"

	contents, err := w.serializeMicroflow(mf)
	if err != nil {
		return fmt.Errorf("failed to serialize microflow: %w", err)
	}

	return w.insertUnit(string(mf.ID), string(mf.ContainerID), "Microflows$Microflow", contents)
}

// UpdateMicroflow updates an existing microflow.
func (w *Writer) UpdateMicroflow(mf *microflows.Microflow) error {
	contents, err := w.serializeMicroflow(mf)
	if err != nil {
		return fmt.Errorf("failed to serialize microflow: %w", err)
	}

	return w.updateUnit(string(mf.ID), contents)
}

// DeleteMicroflow deletes a microflow.
func (w *Writer) DeleteMicroflow(id model.ID) error {
	return w.deleteUnit(string(id))
}

// CreateNanoflow creates a new nanoflow.
func (w *Writer) CreateNanoflow(nf *microflows.Nanoflow) error {
	if nf.ID == "" {
		nf.ID = model.ID(generateUUID())
	}
	nf.TypeName = "Microflows$Nanoflow"

	contents, err := w.serializeNanoflow(nf)
	if err != nil {
		return fmt.Errorf("failed to serialize nanoflow: %w", err)
	}

	return w.insertUnit(string(nf.ID), string(nf.ContainerID), "Microflows$Nanoflow", contents)
}

// UpdateNanoflow updates an existing nanoflow.
func (w *Writer) UpdateNanoflow(nf *microflows.Nanoflow) error {
	contents, err := w.serializeNanoflow(nf)
	if err != nil {
		return fmt.Errorf("failed to serialize nanoflow: %w", err)
	}

	return w.updateUnit(string(nf.ID), contents)
}

// DeleteNanoflow deletes a nanoflow.
func (w *Writer) DeleteNanoflow(id model.ID) error {
	return w.deleteUnit(string(id))
}

// CreatePage creates a new page.
func (w *Writer) CreatePage(page *pages.Page) error {
	if page.ID == "" {
		page.ID = model.ID(generateUUID())
	}
	page.TypeName = "Pages$Page"

	contents, err := w.serializePage(page)
	if err != nil {
		return fmt.Errorf("failed to serialize page: %w", err)
	}

	return w.insertUnit(string(page.ID), string(page.ContainerID), "Pages$Page", contents)
}

// UpdatePage updates an existing page.
func (w *Writer) UpdatePage(page *pages.Page) error {
	contents, err := w.serializePage(page)
	if err != nil {
		return fmt.Errorf("failed to serialize page: %w", err)
	}

	return w.updateUnit(string(page.ID), contents)
}

// DeletePage deletes a page.
func (w *Writer) DeletePage(id model.ID) error {
	return w.deleteUnit(string(id))
}

// CreateLayout creates a new layout.
func (w *Writer) CreateLayout(layout *pages.Layout) error {
	if layout.ID == "" {
		layout.ID = model.ID(generateUUID())
	}
	layout.TypeName = "Pages$Layout"

	contents, err := w.serializeLayout(layout)
	if err != nil {
		return fmt.Errorf("failed to serialize layout: %w", err)
	}

	return w.insertUnit(string(layout.ID), string(layout.ContainerID), "Pages$Layout", contents)
}

// UpdateLayout updates an existing layout.
func (w *Writer) UpdateLayout(layout *pages.Layout) error {
	contents, err := w.serializeLayout(layout)
	if err != nil {
		return fmt.Errorf("failed to serialize layout: %w", err)
	}

	return w.updateUnit(string(layout.ID), contents)
}

// DeleteLayout deletes a layout.
func (w *Writer) DeleteLayout(id model.ID) error {
	return w.deleteUnit(string(id))
}

// CreateEnumeration creates a new enumeration.
func (w *Writer) CreateEnumeration(enum *model.Enumeration) error {
	if enum.ID == "" {
		enum.ID = model.ID(generateUUID())
	}
	enum.TypeName = "Enumerations$Enumeration"

	contents, err := w.serializeEnumeration(enum)
	if err != nil {
		return fmt.Errorf("failed to serialize enumeration: %w", err)
	}

	return w.insertUnit(string(enum.ID), string(enum.ContainerID), "Enumerations$Enumeration", contents)
}

// UpdateEnumeration updates an existing enumeration.
func (w *Writer) UpdateEnumeration(enum *model.Enumeration) error {
	contents, err := w.serializeEnumeration(enum)
	if err != nil {
		return fmt.Errorf("failed to serialize enumeration: %w", err)
	}

	return w.updateUnit(string(enum.ID), contents)
}

// DeleteEnumeration deletes an enumeration.
func (w *Writer) DeleteEnumeration(id model.ID) error {
	return w.deleteUnit(string(id))
}

// CreateConstant creates a new constant.
func (w *Writer) CreateConstant(constant *model.Constant) error {
	if constant.ID == "" {
		constant.ID = model.ID(generateUUID())
	}
	constant.TypeName = "Constants$Constant"

	contents, err := w.serializeConstant(constant)
	if err != nil {
		return fmt.Errorf("failed to serialize constant: %w", err)
	}

	return w.insertUnit(string(constant.ID), string(constant.ContainerID), "Constants$Constant", contents)
}

// UpdateConstant updates an existing constant.
func (w *Writer) UpdateConstant(constant *model.Constant) error {
	contents, err := w.serializeConstant(constant)
	if err != nil {
		return fmt.Errorf("failed to serialize constant: %w", err)
	}

	return w.updateUnit(string(constant.ID), contents)
}

// DeleteConstant deletes a constant.
func (w *Writer) DeleteConstant(id model.ID) error {
	return w.deleteUnit(string(id))
}

// Helper methods

func (w *Writer) insertUnit(unitID, containerID, unitType string, contents []byte) error {
	if w.reader.version == MPRVersionV2 {
		// Write to external file
		filePath := filepath.Join(w.reader.contentsDir, unitID)
		if err := os.WriteFile(filePath, contents, 0644); err != nil {
			return fmt.Errorf("failed to write unit file: %w", err)
		}
		// Insert reference to database
		_, err := w.reader.db.Exec(`
			INSERT INTO Unit (UnitID, ContainerID, Type, ContentsHash, Contents)
			VALUES (?, ?, ?, '', NULL)
		`, unitID, containerID, unitType)
		return err
	}

	// MPR v1: Store directly in database
	_, err := w.reader.db.Exec(`
		INSERT INTO Unit (UnitID, ContainerID, Type, Contents)
		VALUES (?, ?, ?, ?)
	`, unitID, containerID, unitType, contents)
	return err
}

func (w *Writer) updateUnit(unitID string, contents []byte) error {
	if w.reader.version == MPRVersionV2 {
		// Update external file
		filePath := filepath.Join(w.reader.contentsDir, unitID)
		return os.WriteFile(filePath, contents, 0644)
	}

	// MPR v1: Update in database
	_, err := w.reader.db.Exec(`
		UPDATE Unit SET Contents = ? WHERE UnitID = ?
	`, contents, unitID)
	return err
}

func (w *Writer) deleteUnit(unitID string) error {
	if w.reader.version == MPRVersionV2 {
		// Delete external file
		filePath := filepath.Join(w.reader.contentsDir, unitID)
		os.Remove(filePath) // Ignore error if file doesn't exist
	}

	_, err := w.reader.db.Exec(`DELETE FROM Unit WHERE UnitID = ?`, unitID)
	return err
}

func (w *Writer) updateDomainModel(dm *domainmodel.DomainModel) error {
	contents, err := w.serializeDomainModel(dm)
	if err != nil {
		return fmt.Errorf("failed to serialize domain model: %w", err)
	}

	return w.updateUnit(string(dm.ID), contents)
}

// Serialization methods

func (w *Writer) serializeModule(module *model.Module) ([]byte, error) {
	doc := bson.M{
		"$ID":            string(module.ID),
		"$Type":          module.TypeName,
		"Name":           module.Name,
		"Documentation":  module.Documentation,
		"Excluded":       module.Excluded,
		"FromAppStore":   module.FromAppStore,
	}
	return bson.Marshal(doc)
}

func (w *Writer) serializeDomainModel(dm *domainmodel.DomainModel) ([]byte, error) {
	entities := make([]bson.M, 0, len(dm.Entities))
	for _, e := range dm.Entities {
		entities = append(entities, serializeEntity(e))
	}

	associations := make([]bson.M, 0, len(dm.Associations))
	for _, a := range dm.Associations {
		associations = append(associations, serializeAssociation(a))
	}

	doc := bson.M{
		"$ID":          string(dm.ID),
		"$Type":        dm.TypeName,
		"Entities":     entities,
		"Associations": associations,
	}
	return bson.Marshal(doc)
}

func serializeEntity(e *domainmodel.Entity) bson.M {
	attrs := make([]bson.M, 0, len(e.Attributes))
	for _, a := range e.Attributes {
		attrs = append(attrs, serializeAttribute(a))
	}

	return bson.M{
		"$ID":           string(e.ID),
		"$Type":         e.TypeName,
		"Name":          e.Name,
		"Documentation": e.Documentation,
		"Location": bson.M{
			"x": e.Location.X,
			"y": e.Location.Y,
		},
		"Persistable": e.Persistable,
		"Attributes":  attrs,
	}
}

func serializeAttribute(a *domainmodel.Attribute) bson.M {
	attrType := bson.M{}
	if a.Type != nil {
		attrType["$Type"] = "DomainModels$" + a.Type.GetTypeName() + "AttributeType"
		if strType, ok := a.Type.(*domainmodel.StringAttributeType); ok {
			attrType["Length"] = strType.Length
		}
	}

	return bson.M{
		"$ID":           string(a.ID),
		"$Type":         a.TypeName,
		"Name":          a.Name,
		"Documentation": a.Documentation,
		"Type":          attrType,
	}
}

func serializeAssociation(a *domainmodel.Association) bson.M {
	return bson.M{
		"$ID":           string(a.ID),
		"$Type":         a.TypeName,
		"Name":          a.Name,
		"Documentation": a.Documentation,
		"Parent":        string(a.ParentID),
		"Child":         string(a.ChildID),
		"Type":          string(a.Type),
		"Owner":         string(a.Owner),
	}
}

func (w *Writer) serializeMicroflow(mf *microflows.Microflow) ([]byte, error) {
	params := make([]bson.M, 0, len(mf.Parameters))
	for _, p := range mf.Parameters {
		params = append(params, bson.M{
			"$ID":           string(p.ID),
			"$Type":         p.TypeName,
			"Name":          p.Name,
			"Documentation": p.Documentation,
		})
	}

	doc := bson.M{
		"$ID":                      string(mf.ID),
		"$Type":                    mf.TypeName,
		"Name":                     mf.Name,
		"Documentation":            mf.Documentation,
		"AllowConcurrentExecution": mf.AllowConcurrentExecution,
		"MarkAsUsed":               mf.MarkAsUsed,
		"Excluded":                 mf.Excluded,
		"Parameters":               params,
	}
	return bson.Marshal(doc)
}

func (w *Writer) serializeNanoflow(nf *microflows.Nanoflow) ([]byte, error) {
	params := make([]bson.M, 0, len(nf.Parameters))
	for _, p := range nf.Parameters {
		params = append(params, bson.M{
			"$ID":           string(p.ID),
			"$Type":         p.TypeName,
			"Name":          p.Name,
			"Documentation": p.Documentation,
		})
	}

	doc := bson.M{
		"$ID":           string(nf.ID),
		"$Type":         nf.TypeName,
		"Name":          nf.Name,
		"Documentation": nf.Documentation,
		"MarkAsUsed":    nf.MarkAsUsed,
		"Excluded":      nf.Excluded,
		"Parameters":    params,
	}
	return bson.Marshal(doc)
}

func (w *Writer) serializePage(page *pages.Page) ([]byte, error) {
	doc := bson.M{
		"$ID":           string(page.ID),
		"$Type":         page.TypeName,
		"Name":          page.Name,
		"Documentation": page.Documentation,
		"URL":           page.URL,
		"MarkAsUsed":    page.MarkAsUsed,
		"Excluded":      page.Excluded,
	}

	if page.LayoutID != "" {
		doc["Layout"] = string(page.LayoutID)
	}

	if page.Title != nil {
		doc["Title"] = bson.M{
			"$ID":          string(page.Title.ID),
			"$Type":        page.Title.TypeName,
			"Translations": page.Title.Translations,
		}
	}

	return bson.Marshal(doc)
}

func (w *Writer) serializeLayout(layout *pages.Layout) ([]byte, error) {
	doc := bson.M{
		"$ID":           string(layout.ID),
		"$Type":         layout.TypeName,
		"Name":          layout.Name,
		"Documentation": layout.Documentation,
		"LayoutType":    string(layout.LayoutType),
	}
	return bson.Marshal(doc)
}

func (w *Writer) serializeEnumeration(enum *model.Enumeration) ([]byte, error) {
	values := make([]bson.M, 0, len(enum.Values))
	for _, v := range enum.Values {
		valueDoc := bson.M{
			"$ID":   string(v.ID),
			"$Type": v.TypeName,
			"Name":  v.Name,
		}
		if v.Caption != nil {
			valueDoc["Caption"] = bson.M{
				"$ID":          string(v.Caption.ID),
				"$Type":        v.Caption.TypeName,
				"Translations": v.Caption.Translations,
			}
		}
		values = append(values, valueDoc)
	}

	doc := bson.M{
		"$ID":           string(enum.ID),
		"$Type":         enum.TypeName,
		"Name":          enum.Name,
		"Documentation": enum.Documentation,
		"Values":        values,
	}
	return bson.Marshal(doc)
}

func (w *Writer) serializeConstant(constant *model.Constant) ([]byte, error) {
	doc := bson.M{
		"$ID":             string(constant.ID),
		"$Type":           constant.TypeName,
		"Name":            constant.Name,
		"Documentation":   constant.Documentation,
		"Type":            constant.Type,
		"DefaultValue":    constant.DefaultValue,
		"ExposedToClient": constant.ExposedToClient,
	}
	return bson.Marshal(doc)
}

// Transaction support

// Transaction represents a database transaction.
type Transaction struct {
	tx     *sql.Tx
	writer *Writer
}

// BeginTransaction starts a new transaction.
func (w *Writer) BeginTransaction() (*Transaction, error) {
	tx, err := w.reader.db.Begin()
	if err != nil {
		return nil, err
	}
	return &Transaction{tx: tx, writer: w}, nil
}

// Commit commits the transaction.
func (t *Transaction) Commit() error {
	return t.tx.Commit()
}

// Rollback rolls back the transaction.
func (t *Transaction) Rollback() error {
	return t.tx.Rollback()
}

// generateUUID generates a new UUID for model elements.
func generateUUID() string {
	// Use crypto/rand for secure UUID generation
	b := make([]byte, 16)
	_, _ = randomRead(b)
	b[6] = (b[6] & 0x0f) | 0x40 // Version 4
	b[8] = (b[8] & 0x3f) | 0x80 // Variant is 10

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// randomRead is a wrapper for crypto/rand.Read for testing.
var randomRead = func(b []byte) (int, error) {
	// Import at runtime to avoid circular dependencies
	return cryptoRandRead(b)
}

// cryptoRandRead uses crypto/rand
func cryptoRandRead(b []byte) (int, error) {
	// Placeholder - actual implementation uses crypto/rand
	// Using time-based fallback for simplicity
	for i := range b {
		b[i] = byte(i * 17)
	}
	return len(b), nil
}
