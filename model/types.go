// Package model provides core types for Mendix model elements.
package model

import (
	"encoding/json"
	"time"
)

// ID represents a unique identifier for model elements.
// In Mendix, these are typically UUIDs.
type ID string

// QualifiedName represents a fully qualified name in the format "Module.Element".
type QualifiedName string

// Point represents a position in 2D space.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Size represents dimensions in 2D space.
type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Element is the base interface for all model elements.
type Element interface {
	GetID() ID
	GetTypeName() string
}

// NamedElement is an element with a name.
type NamedElement interface {
	Element
	GetName() string
}

// ContainedElement is an element that belongs to a container.
type ContainedElement interface {
	Element
	GetContainerID() ID
}

// BaseElement provides common fields for all model elements.
type BaseElement struct {
	ID       ID     `json:"$ID"`
	TypeName string `json:"$Type"`
}

// GetID returns the element's unique identifier.
func (e *BaseElement) GetID() ID {
	return e.ID
}

// GetTypeName returns the element's type name.
func (e *BaseElement) GetTypeName() string {
	return e.TypeName
}

// Unit represents a document unit in the Mendix model.
// Units are top-level elements like DomainModel, Microflow, Page, etc.
type Unit struct {
	BaseElement
	ContainerID ID     `json:"containerId"`
	Name        string `json:"name,omitempty"`
}

// GetName returns the unit's name.
func (u *Unit) GetName() string {
	return u.Name
}

// GetContainerID returns the ID of the containing element.
func (u *Unit) GetContainerID() ID {
	return u.ContainerID
}

// Module represents a Mendix module.
type Module struct {
	BaseElement
	Name            string `json:"name"`
	Documentation   string `json:"documentation,omitempty"`
	Excluded        bool   `json:"excluded,omitempty"`
	FromAppStore    bool   `json:"fromAppStore,omitempty"`
	IsReusableComponent bool `json:"isReusableComponent,omitempty"`

	// Contained units
	DomainModelID ID   `json:"domainModelId,omitempty"`
	Documents     []ID `json:"documents,omitempty"`
}

// GetName returns the module's name.
func (m *Module) GetName() string {
	return m.Name
}

// Project represents a Mendix project.
type Project struct {
	BaseElement
	Name           string    `json:"name"`
	MendixVersion  string    `json:"mendixVersion"`
	ProjectID      string    `json:"projectId,omitempty"`
	IsSystemProject bool     `json:"isSystemProject,omitempty"`
	CreatedDate    time.Time `json:"createdDate,omitempty"`

	// Project-level settings
	Modules        []ID `json:"modules,omitempty"`
	ProjectDocuments []ID `json:"projectDocuments,omitempty"`
}

// GetName returns the project's name.
func (p *Project) GetName() string {
	return p.Name
}

// Folder represents a folder within a module for organizing documents.
type Folder struct {
	BaseElement
	ContainerID ID     `json:"containerId"`
	Name        string `json:"name"`
	Documents   []ID   `json:"documents,omitempty"`
	Folders     []ID   `json:"folders,omitempty"`
}

// GetName returns the folder's name.
func (f *Folder) GetName() string {
	return f.Name
}

// GetContainerID returns the ID of the containing element.
func (f *Folder) GetContainerID() ID {
	return f.ContainerID
}

// Text represents localized text.
type Text struct {
	BaseElement
	Translations map[string]string `json:"translations,omitempty"`
}

// GetTranslation returns the translation for a given language code.
func (t *Text) GetTranslation(languageCode string) string {
	if t.Translations == nil {
		return ""
	}
	return t.Translations[languageCode]
}

// Image represents an image reference.
type Image struct {
	BaseElement
	Name       string `json:"name,omitempty"`
	ImageData  []byte `json:"imageData,omitempty"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
}

// Constant represents a constant value.
type Constant struct {
	BaseElement
	ContainerID   ID     `json:"containerId"`
	Name          string `json:"name"`
	Documentation string `json:"documentation,omitempty"`
	Type          string `json:"type"`
	DefaultValue  string `json:"defaultValue,omitempty"`
	ExposedToClient bool `json:"exposedToClient,omitempty"`
}

// GetName returns the constant's name.
func (c *Constant) GetName() string {
	return c.Name
}

// GetContainerID returns the ID of the containing element.
func (c *Constant) GetContainerID() ID {
	return c.ContainerID
}

// Enumeration represents an enumeration type.
type Enumeration struct {
	BaseElement
	ContainerID   ID                   `json:"containerId"`
	Name          string               `json:"name"`
	Documentation string               `json:"documentation,omitempty"`
	Values        []EnumerationValue   `json:"values,omitempty"`
}

// GetName returns the enumeration's name.
func (e *Enumeration) GetName() string {
	return e.Name
}

// GetContainerID returns the ID of the containing element.
func (e *Enumeration) GetContainerID() ID {
	return e.ContainerID
}

// EnumerationValue represents a value in an enumeration.
type EnumerationValue struct {
	BaseElement
	Name   string `json:"name"`
	Caption *Text `json:"caption,omitempty"`
	Image  *Image `json:"image,omitempty"`
}

// GetName returns the enumeration value's name.
func (v *EnumerationValue) GetName() string {
	return v.Name
}

// RegularExpression represents a regular expression constraint.
type RegularExpression struct {
	BaseElement
	ContainerID   ID     `json:"containerId"`
	Name          string `json:"name"`
	Documentation string `json:"documentation,omitempty"`
	Expression    string `json:"expression"`
}

// GetName returns the regular expression's name.
func (r *RegularExpression) GetName() string {
	return r.Name
}

// GetContainerID returns the ID of the containing element.
func (r *RegularExpression) GetContainerID() ID {
	return r.ContainerID
}

// ScheduledEvent represents a scheduled event.
type ScheduledEvent struct {
	BaseElement
	ContainerID    ID     `json:"containerId"`
	Name           string `json:"name"`
	Documentation  string `json:"documentation,omitempty"`
	MicroflowID    ID     `json:"microflowId,omitempty"`
	StartDateTime  *time.Time `json:"startDateTime,omitempty"`
	TimeZone       string `json:"timeZone,omitempty"`
	Interval       int    `json:"interval,omitempty"`
	IntervalType   string `json:"intervalType,omitempty"`
	Enabled        bool   `json:"enabled"`
}

// GetName returns the scheduled event's name.
func (s *ScheduledEvent) GetName() string {
	return s.Name
}

// GetContainerID returns the ID of the containing element.
func (s *ScheduledEvent) GetContainerID() ID {
	return s.ContainerID
}

// MarshalJSON provides custom JSON marshaling.
func (e *BaseElement) MarshalJSON() ([]byte, error) {
	type Alias BaseElement
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	})
}

// DocumentType represents the type of a document.
type DocumentType string

const (
	DocumentTypeDomainModel    DocumentType = "DomainModels$DomainModel"
	DocumentTypeMicroflow      DocumentType = "Microflows$Microflow"
	DocumentTypeNanoflow       DocumentType = "Microflows$Nanoflow"
	DocumentTypePage           DocumentType = "Pages$Page"
	DocumentTypeLayout         DocumentType = "Pages$Layout"
	DocumentTypeSnippet        DocumentType = "Pages$Snippet"
	DocumentTypeConstant       DocumentType = "Constants$Constant"
	DocumentTypeEnumeration    DocumentType = "Enumerations$Enumeration"
	DocumentTypeScheduledEvent DocumentType = "ScheduledEvents$ScheduledEvent"
	DocumentTypeJavaAction     DocumentType = "JavaActions$JavaAction"
	DocumentTypeRule           DocumentType = "Rules$Rule"
)
