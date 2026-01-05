// Package pages provides types for Mendix pages, layouts, and widgets.
package pages

import (
	"github.com/anthropics/modelsdk-go/model"
)

// Page represents a page in the Mendix model.
type Page struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`
	Title         *model.Text `json:"title,omitempty"`
	URL           string   `json:"url,omitempty"`
	LayoutID      model.ID `json:"layoutId,omitempty"`
	LayoutCall    *LayoutCall `json:"layoutCall,omitempty"`
	AllowedRoles  []model.ID `json:"allowedRoles,omitempty"`
	Parameters    []*PageParameter `json:"parameters,omitempty"`
	PopupWidth    int      `json:"popupWidth,omitempty"`
	PopupHeight   int      `json:"popupHeight,omitempty"`
	PopupResizable bool    `json:"popupResizable,omitempty"`
	MarkAsUsed    bool     `json:"markAsUsed"`
	Excluded      bool     `json:"excluded"`
}

// GetName returns the page's name.
func (p *Page) GetName() string {
	return p.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (p *Page) GetContainerID() model.ID {
	return p.ContainerID
}

// Layout represents a layout in the Mendix model.
type Layout struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`
	LayoutType    LayoutType `json:"layoutType"`
	MainPlaceholderID model.ID `json:"mainPlaceholderId,omitempty"`
	Widget        Widget   `json:"widget,omitempty"`
}

// GetName returns the layout's name.
func (l *Layout) GetName() string {
	return l.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (l *Layout) GetContainerID() model.ID {
	return l.ContainerID
}

// LayoutType represents the type of layout.
type LayoutType string

const (
	LayoutTypeResponsive     LayoutType = "Responsive"
	LayoutTypeTablet         LayoutType = "Tablet"
	LayoutTypePhone          LayoutType = "Phone"
	LayoutTypeModalPopup     LayoutType = "ModalPopup"
	LayoutTypePopup          LayoutType = "Popup"
	LayoutTypeLegacy         LayoutType = "Legacy"
)

// Snippet represents a reusable page snippet.
type Snippet struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`
	EntityID      model.ID `json:"entityId,omitempty"`
	Parameters    []*SnippetParameter `json:"parameters,omitempty"`
	Widget        Widget   `json:"widget,omitempty"`
}

// GetName returns the snippet's name.
func (s *Snippet) GetName() string {
	return s.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (s *Snippet) GetContainerID() model.ID {
	return s.ContainerID
}

// BuildingBlock represents a building block.
type BuildingBlock struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`
	Widget        Widget   `json:"widget,omitempty"`
	TemplateID    string   `json:"templateId,omitempty"`
}

// GetName returns the building block's name.
func (bb *BuildingBlock) GetName() string {
	return bb.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (bb *BuildingBlock) GetContainerID() model.ID {
	return bb.ContainerID
}

// PageTemplate represents a page template.
type PageTemplate struct {
	model.BaseElement
	ContainerID    model.ID `json:"containerId"`
	Name           string   `json:"name"`
	Documentation  string   `json:"documentation,omitempty"`
	DisplayName    *model.Text `json:"displayName,omitempty"`
	LayoutID       model.ID `json:"layoutId,omitempty"`
	PageTemplateType PageTemplateType `json:"pageTemplateType"`
	Widget         Widget   `json:"widget,omitempty"`
}

// GetName returns the page template's name.
func (pt *PageTemplate) GetName() string {
	return pt.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (pt *PageTemplate) GetContainerID() model.ID {
	return pt.ContainerID
}

// PageTemplateType represents the type of page template.
type PageTemplateType string

const (
	PageTemplateTypeStandard PageTemplateType = "Standard"
	PageTemplateTypeEdit     PageTemplateType = "Edit"
	PageTemplateTypeSelect   PageTemplateType = "Select"
)

// LayoutCall represents a call to a layout with argument bindings.
type LayoutCall struct {
	model.BaseElement
	LayoutID  model.ID `json:"layoutId"`
	Arguments []*LayoutCallArgument `json:"arguments,omitempty"`
}

// LayoutCallArgument represents an argument binding in a layout call.
type LayoutCallArgument struct {
	model.BaseElement
	ParameterID model.ID `json:"parameterId"`
	Widget      Widget   `json:"widget,omitempty"`
}

// PageParameter represents a parameter of a page.
type PageParameter struct {
	model.BaseElement
	ContainerID model.ID `json:"containerId"`
	Name        string   `json:"name"`
	EntityID    model.ID `json:"entityId,omitempty"`
}

// GetName returns the parameter's name.
func (p *PageParameter) GetName() string {
	return p.Name
}

// GetContainerID returns the ID of the containing page.
func (p *PageParameter) GetContainerID() model.ID {
	return p.ContainerID
}

// SnippetParameter represents a parameter of a snippet.
type SnippetParameter struct {
	model.BaseElement
	ContainerID model.ID `json:"containerId"`
	Name        string   `json:"name"`
	EntityID    model.ID `json:"entityId,omitempty"`
	Type        string   `json:"type,omitempty"`
}

// GetName returns the parameter's name.
func (p *SnippetParameter) GetName() string {
	return p.Name
}

// GetContainerID returns the ID of the containing snippet.
func (p *SnippetParameter) GetContainerID() model.ID {
	return p.ContainerID
}

// Widget is the base interface for all page widgets.
type Widget interface {
	GetID() model.ID
	GetTypeName() string
	GetName() string
}

// BaseWidget provides common fields for all widgets.
type BaseWidget struct {
	model.BaseElement
	Name     string `json:"name"`
	Class    string `json:"class,omitempty"`
	Style    string `json:"style,omitempty"`
	TabIndex int    `json:"tabIndex,omitempty"`
}

// GetName returns the widget's name.
func (w *BaseWidget) GetName() string {
	return w.Name
}

// Placeholder Widgets

// LayoutPlaceholder represents a placeholder in a layout.
type LayoutPlaceholder struct {
	BaseWidget
}

// ConditionalVisibilitySettings represents visibility conditions.
type ConditionalVisibilitySettings struct {
	model.BaseElement
	Expression     string     `json:"expression,omitempty"`
	ModuleRoles    []model.ID `json:"moduleRoles,omitempty"`
	SourceVariable *PageVariable `json:"sourceVariable,omitempty"`
	Attribute      model.ID   `json:"attribute,omitempty"`
}

// PageVariable represents a page variable reference.
type PageVariable struct {
	model.BaseElement
	UseAllPages bool     `json:"useAllPages"`
	PageID      model.ID `json:"pageId,omitempty"`
	Widget      string   `json:"widget,omitempty"`
}

// Container Widgets

// LayoutGrid represents a layout grid container.
type LayoutGrid struct {
	BaseWidget
	Rows []*LayoutGridRow `json:"rows,omitempty"`
}

// LayoutGridRow represents a row in a layout grid.
type LayoutGridRow struct {
	model.BaseElement
	Columns []*LayoutGridColumn `json:"columns,omitempty"`
}

// LayoutGridColumn represents a column in a layout grid.
type LayoutGridColumn struct {
	model.BaseElement
	Weight   int    `json:"weight"`
	Widgets  []Widget `json:"widgets,omitempty"`
}

// Container represents a generic container widget.
type Container struct {
	BaseWidget
	Widgets []Widget `json:"widgets,omitempty"`
	RenderMode ContainerRenderMode `json:"renderMode,omitempty"`
}

// ContainerRenderMode represents how a container is rendered.
type ContainerRenderMode string

const (
	ContainerRenderModeDiv  ContainerRenderMode = "Div"
	ContainerRenderModeForm ContainerRenderMode = "Form"
)

// GroupBox represents a group box container.
type GroupBox struct {
	BaseWidget
	Caption    *model.Text `json:"caption,omitempty"`
	Collapsible bool       `json:"collapsible"`
	Collapsed   bool       `json:"collapsed"`
	Widgets    []Widget    `json:"widgets,omitempty"`
}

// TabContainer represents a tab container.
type TabContainer struct {
	BaseWidget
	TabPages      []*TabPage `json:"tabPages,omitempty"`
	DefaultPageID model.ID   `json:"defaultPageId,omitempty"`
}

// TabPage represents a page within a tab container.
type TabPage struct {
	model.BaseElement
	Name      string      `json:"name"`
	Caption   *model.Text `json:"caption,omitempty"`
	Widgets   []Widget    `json:"widgets,omitempty"`
	RefreshOnShow bool    `json:"refreshOnShow,omitempty"`
}

// GetName returns the tab page's name.
func (tp *TabPage) GetName() string {
	return tp.Name
}

// ScrollContainer represents a scrollable container.
type ScrollContainer struct {
	BaseWidget
	ScrollBehavior ScrollBehavior `json:"scrollBehavior"`
	Widgets        []Widget       `json:"widgets,omitempty"`
}

// ScrollBehavior represents how scrolling behaves.
type ScrollBehavior string

const (
	ScrollBehaviorVertical   ScrollBehavior = "Vertical"
	ScrollBehaviorHorizontal ScrollBehavior = "Horizontal"
	ScrollBehaviorBoth       ScrollBehavior = "Both"
)

// DataView represents a data view widget.
type DataView struct {
	BaseWidget
	DataSource   DataSource `json:"dataSource,omitempty"`
	Editable     bool       `json:"editable"`
	ReadOnly     bool       `json:"readOnly,omitempty"`
	ShowFooter   bool       `json:"showFooter"`
	Widgets      []Widget   `json:"widgets,omitempty"`
	FooterWidgets []Widget  `json:"footerWidgets,omitempty"`
	NoEntityMessage *model.Text `json:"noEntityMessage,omitempty"`
}

// ListView represents a list view widget.
type ListView struct {
	BaseWidget
	DataSource      DataSource `json:"dataSource,omitempty"`
	Editable        bool       `json:"editable"`
	ClickAction     ClientAction `json:"clickAction,omitempty"`
	PageSize        int        `json:"pageSize,omitempty"`
	Widgets         []Widget   `json:"widgets,omitempty"`
	Templates       []*ListViewTemplate `json:"templates,omitempty"`
}

// ListViewTemplate represents a template in a list view.
type ListViewTemplate struct {
	model.BaseElement
	Widgets []Widget `json:"widgets,omitempty"`
}

// TemplateGrid represents a template grid widget.
type TemplateGrid struct {
	BaseWidget
	DataSource        DataSource `json:"dataSource,omitempty"`
	NumberOfColumns   int        `json:"numberOfColumns"`
	NumberOfRows      int        `json:"numberOfRows"`
	SelectionMode     SelectionMode `json:"selectionMode"`
	SelectFirst       bool       `json:"selectFirst"`
	Widgets           []Widget   `json:"widgets,omitempty"`
	ControlBarWidgets []Widget   `json:"controlBarWidgets,omitempty"`
}

// DataGrid represents a data grid widget.
type DataGrid struct {
	BaseWidget
	DataSource         DataSource `json:"dataSource,omitempty"`
	Columns            []*DataGridColumn `json:"columns,omitempty"`
	SelectionMode      SelectionMode `json:"selectionMode"`
	SelectFirst        bool       `json:"selectFirst"`
	ShowPagingButtons  bool       `json:"showPagingButtons"`
	ShowEmptyRows      bool       `json:"showEmptyRows,omitempty"`
	WidthUnit          WidthUnit  `json:"widthUnit,omitempty"`
	ControlBarWidgets  []Widget   `json:"controlBarWidgets,omitempty"`
}

// DataGridColumn represents a column in a data grid.
type DataGridColumn struct {
	model.BaseElement
	Name          string      `json:"name,omitempty"`
	Caption       *model.Text `json:"caption,omitempty"`
	AttributePath string      `json:"attributePath,omitempty"`
	Editable      bool        `json:"editable"`
	Aggregate     AggregateFunction `json:"aggregate,omitempty"`
	AggregateCaption *model.Text `json:"aggregateCaption,omitempty"`
	ShowTooltip   bool        `json:"showTooltip,omitempty"`
}

// AggregateFunction represents an aggregate function for columns.
type AggregateFunction string

const (
	AggregateFunctionNone    AggregateFunction = "None"
	AggregateFunctionAverage AggregateFunction = "Average"
	AggregateFunctionCount   AggregateFunction = "Count"
	AggregateFunctionMaximum AggregateFunction = "Maximum"
	AggregateFunctionMinimum AggregateFunction = "Minimum"
	AggregateFunctionSum     AggregateFunction = "Sum"
)

// SelectionMode represents how selection works.
type SelectionMode string

const (
	SelectionModeNone   SelectionMode = "None"
	SelectionModeSingle SelectionMode = "Single"
	SelectionModeMulti  SelectionMode = "Multi"
)

// WidthUnit represents the unit for widths.
type WidthUnit string

const (
	WidthUnitPercentage WidthUnit = "Percentage"
	WidthUnitPixels     WidthUnit = "Pixels"
)

// Data Sources

// DataSource represents a data source for widgets.
type DataSource interface {
	isDataSource()
}

// EntityPathSource retrieves data via an entity path.
type EntityPathSource struct {
	model.BaseElement
	EntityPath string `json:"entityPath"`
}

func (EntityPathSource) isDataSource() {}

// DatabaseSource retrieves data from the database.
type DatabaseSource struct {
	model.BaseElement
	EntityID        model.ID    `json:"entityId"`
	XPathConstraint string      `json:"xPathConstraint,omitempty"`
	Sorting         []*GridSort `json:"sorting,omitempty"`
}

func (DatabaseSource) isDataSource() {}

// GridSort represents sorting configuration.
type GridSort struct {
	model.BaseElement
	AttributePath string `json:"attributePath"`
	Direction     SortDirection `json:"direction"`
}

// SortDirection represents the sort direction.
type SortDirection string

const (
	SortDirectionAscending  SortDirection = "Ascending"
	SortDirectionDescending SortDirection = "Descending"
)

// MicroflowSource retrieves data from a microflow.
type MicroflowSource struct {
	model.BaseElement
	MicroflowID model.ID `json:"microflowId"`
}

func (MicroflowSource) isDataSource() {}

// NanoflowSource retrieves data from a nanoflow.
type NanoflowSource struct {
	model.BaseElement
	NanoflowID model.ID `json:"nanoflowId"`
}

func (NanoflowSource) isDataSource() {}

// ListenToWidgetSource listens to another widget.
type ListenToWidgetSource struct {
	model.BaseElement
	WidgetID model.ID `json:"widgetId"`
}

func (ListenToWidgetSource) isDataSource() {}

// AssociationSource retrieves data via association.
type AssociationSource struct {
	model.BaseElement
	EntityPath string `json:"entityPath"`
}

func (AssociationSource) isDataSource() {}

// Input Widgets

// TextBox represents a text input widget.
type TextBox struct {
	BaseWidget
	AttributePath      string        `json:"attributePath,omitempty"`
	FormattingInfo     *FormattingInfo `json:"formattingInfo,omitempty"`
	Placeholder        *model.Text   `json:"placeholder,omitempty"`
	MaxLength          int           `json:"maxLength,omitempty"`
	IsPassword         bool          `json:"isPassword,omitempty"`
	ReadOnly           bool          `json:"readOnly,omitempty"`
	OnChangeAction     ClientAction  `json:"onChangeAction,omitempty"`
	OnEnterAction      ClientAction  `json:"onEnterAction,omitempty"`
}

// TextArea represents a multi-line text input widget.
type TextArea struct {
	BaseWidget
	AttributePath  string       `json:"attributePath,omitempty"`
	Placeholder    *model.Text  `json:"placeholder,omitempty"`
	MaxLength      int          `json:"maxLength,omitempty"`
	CounterMessage *model.Text  `json:"counterMessage,omitempty"`
	Rows           int          `json:"rows,omitempty"`
	ReadOnly       bool         `json:"readOnly,omitempty"`
	OnChangeAction ClientAction `json:"onChangeAction,omitempty"`
}

// FormattingInfo represents formatting configuration.
type FormattingInfo struct {
	model.BaseElement
	DecimalPrecision int    `json:"decimalPrecision,omitempty"`
	GroupDigits      bool   `json:"groupDigits,omitempty"`
	EnumFormat       string `json:"enumFormat,omitempty"`
}

// DatePicker represents a date picker widget.
type DatePicker struct {
	BaseWidget
	AttributePath  string       `json:"attributePath,omitempty"`
	Placeholder    *model.Text  `json:"placeholder,omitempty"`
	DateFormat     string       `json:"dateFormat,omitempty"`
	ReadOnly       bool         `json:"readOnly,omitempty"`
	OnChangeAction ClientAction `json:"onChangeAction,omitempty"`
}

// DropDown represents a drop-down selection widget.
type DropDown struct {
	BaseWidget
	AttributePath  string       `json:"attributePath,omitempty"`
	EmptyOption    *model.Text  `json:"emptyOption,omitempty"`
	ReadOnly       bool         `json:"readOnly,omitempty"`
	OnChangeAction ClientAction `json:"onChangeAction,omitempty"`
}

// ReferenceSelector represents a reference selector widget.
type ReferenceSelector struct {
	BaseWidget
	AttributePath      string       `json:"attributePath,omitempty"`
	EmptyOption        *model.Text  `json:"emptyOption,omitempty"`
	SelectorSource     SelectorSource `json:"selectorSource,omitempty"`
	ReadOnly           bool         `json:"readOnly,omitempty"`
	OnChangeAction     ClientAction `json:"onChangeAction,omitempty"`
}

// SelectorSource represents the source for a reference selector.
type SelectorSource interface {
	isSelectorSource()
}

// DatabaseSelectorSource uses a database query.
type DatabaseSelectorSource struct {
	model.BaseElement
	XPathConstraint string `json:"xPathConstraint,omitempty"`
}

func (DatabaseSelectorSource) isSelectorSource() {}

// MicroflowSelectorSource uses a microflow.
type MicroflowSelectorSource struct {
	model.BaseElement
	MicroflowID model.ID `json:"microflowId"`
}

func (MicroflowSelectorSource) isSelectorSource() {}

// ReferenceSetSelector represents a reference set selector widget.
type ReferenceSetSelector struct {
	BaseWidget
	AttributePath      string         `json:"attributePath,omitempty"`
	SelectorSource     SelectorSource `json:"selectorSource,omitempty"`
	ReadOnly           bool           `json:"readOnly,omitempty"`
	OnChangeAction     ClientAction   `json:"onChangeAction,omitempty"`
	ShowSelectPage     bool           `json:"showSelectPage,omitempty"`
	SelectPageID       model.ID       `json:"selectPageId,omitempty"`
}

// CheckBox represents a checkbox widget.
type CheckBox struct {
	BaseWidget
	AttributePath  string       `json:"attributePath,omitempty"`
	ReadOnly       bool         `json:"readOnly,omitempty"`
	OnChangeAction ClientAction `json:"onChangeAction,omitempty"`
}

// RadioButtons represents a radio button group widget.
type RadioButtons struct {
	BaseWidget
	AttributePath   string       `json:"attributePath,omitempty"`
	RenderDirection RenderDirection `json:"renderDirection,omitempty"`
	ReadOnly        bool         `json:"readOnly,omitempty"`
	OnChangeAction  ClientAction `json:"onChangeAction,omitempty"`
}

// RenderDirection represents the direction for rendering.
type RenderDirection string

const (
	RenderDirectionHorizontal RenderDirection = "Horizontal"
	RenderDirectionVertical   RenderDirection = "Vertical"
)

// FileManager represents a file manager widget.
type FileManager struct {
	BaseWidget
	Type           FileManagerType `json:"type"`
	AllowedExtensions string       `json:"allowedExtensions,omitempty"`
	MaxFileSize    int             `json:"maxFileSize,omitempty"`
	ShowButton     bool            `json:"showButton,omitempty"`
}

// FileManagerType represents the type of file manager.
type FileManagerType string

const (
	FileManagerTypeUpload   FileManagerType = "Upload"
	FileManagerTypeDownload FileManagerType = "Download"
	FileManagerTypeBoth     FileManagerType = "Both"
)

// ImageUploader represents an image uploader widget.
type ImageUploader struct {
	BaseWidget
	AllowedExtensions string `json:"allowedExtensions,omitempty"`
	MaxFileSize       int    `json:"maxFileSize,omitempty"`
	ThumbnailSize     int    `json:"thumbnailSize,omitempty"`
}

// Button Widgets

// ActionButton represents a button that triggers an action.
type ActionButton struct {
	BaseWidget
	Caption        *model.Text   `json:"caption,omitempty"`
	Tooltip        *model.Text   `json:"tooltip,omitempty"`
	Icon           *Icon         `json:"icon,omitempty"`
	ButtonStyle    ButtonStyle   `json:"buttonStyle,omitempty"`
	RenderMode     ButtonRenderMode `json:"renderMode,omitempty"`
	Action         ClientAction  `json:"action,omitempty"`
}

// ButtonStyle represents the style of a button.
type ButtonStyle string

const (
	ButtonStyleDefault   ButtonStyle = "Default"
	ButtonStylePrimary   ButtonStyle = "Primary"
	ButtonStyleSecondary ButtonStyle = "Secondary"
	ButtonStyleSuccess   ButtonStyle = "Success"
	ButtonStyleWarning   ButtonStyle = "Warning"
	ButtonStyleDanger    ButtonStyle = "Danger"
	ButtonStyleInverse   ButtonStyle = "Inverse"
	ButtonStyleLink      ButtonStyle = "Link"
)

// ButtonRenderMode represents how a button is rendered.
type ButtonRenderMode string

const (
	ButtonRenderModeButton ButtonRenderMode = "Button"
	ButtonRenderModeLink   ButtonRenderMode = "Link"
)

// Icon represents an icon.
type Icon struct {
	model.BaseElement
	Type    IconType `json:"type"`
	Name    string   `json:"name,omitempty"`
	ImageID model.ID `json:"imageId,omitempty"`
}

// IconType represents the type of icon.
type IconType string

const (
	IconTypeGlyph IconType = "Glyph"
	IconTypeImage IconType = "Image"
)

// DropDownButton represents a dropdown button.
type DropDownButton struct {
	BaseWidget
	Caption     *model.Text `json:"caption,omitempty"`
	Tooltip     *model.Text `json:"tooltip,omitempty"`
	Icon        *Icon       `json:"icon,omitempty"`
	ButtonStyle ButtonStyle `json:"buttonStyle,omitempty"`
	Items       []*DropDownButtonItem `json:"items,omitempty"`
}

// DropDownButtonItem represents an item in a dropdown button.
type DropDownButtonItem struct {
	model.BaseElement
	Caption *model.Text  `json:"caption,omitempty"`
	Action  ClientAction `json:"action,omitempty"`
}

// LinkButton represents a link-style button.
type LinkButton struct {
	BaseWidget
	Caption *model.Text  `json:"caption,omitempty"`
	Tooltip *model.Text  `json:"tooltip,omitempty"`
	Action  ClientAction `json:"action,omitempty"`
}

// Text and Display Widgets

// Text represents a static text widget.
type Text struct {
	BaseWidget
	Caption    *model.Text `json:"caption,omitempty"`
	RenderMode TextRenderMode `json:"renderMode,omitempty"`
}

// TextRenderMode represents how text is rendered.
type TextRenderMode string

const (
	TextRenderModeText     TextRenderMode = "Text"
	TextRenderModeH1       TextRenderMode = "H1"
	TextRenderModeH2       TextRenderMode = "H2"
	TextRenderModeH3       TextRenderMode = "H3"
	TextRenderModeH4       TextRenderMode = "H4"
	TextRenderModeH5       TextRenderMode = "H5"
	TextRenderModeH6       TextRenderMode = "H6"
	TextRenderModeParagraph TextRenderMode = "Paragraph"
)

// DynamicText represents dynamic text based on an attribute.
type DynamicText struct {
	BaseWidget
	AttributePath string         `json:"attributePath,omitempty"`
	RenderMode    TextRenderMode `json:"renderMode,omitempty"`
}

// Label represents a label widget.
type Label struct {
	BaseWidget
	Caption *model.Text `json:"caption,omitempty"`
	ForID   model.ID    `json:"forId,omitempty"`
}

// Title represents a page title widget.
type Title struct {
	BaseWidget
}

// DynamicImage represents a dynamic image widget.
type DynamicImage struct {
	BaseWidget
	DefaultImage   model.ID `json:"defaultImage,omitempty"`
	Width          int      `json:"width,omitempty"`
	WidthUnit      WidthUnit `json:"widthUnit,omitempty"`
	Height         int      `json:"height,omitempty"`
	Responsive     bool     `json:"responsive"`
	OnClickAction  ClientAction `json:"onClickAction,omitempty"`
}

// StaticImage represents a static image widget.
type StaticImage struct {
	BaseWidget
	ImageID       model.ID `json:"imageId,omitempty"`
	Width         int      `json:"width,omitempty"`
	WidthUnit     WidthUnit `json:"widthUnit,omitempty"`
	Height        int      `json:"height,omitempty"`
	Responsive    bool     `json:"responsive"`
	OnClickAction ClientAction `json:"onClickAction,omitempty"`
}

// ClientActions

// ClientAction represents an action triggered by client interaction.
type ClientAction interface {
	isClientAction()
}

// NoClientAction represents no action.
type NoClientAction struct {
	model.BaseElement
}

func (NoClientAction) isClientAction() {}

// PageClientAction opens a page.
type PageClientAction struct {
	model.BaseElement
	PageID         model.ID       `json:"pageId"`
	PageSettings   *PageSettings  `json:"pageSettings,omitempty"`
}

func (PageClientAction) isClientAction() {}

// PageSettings represents page display settings.
type PageSettings struct {
	model.BaseElement
	FormLocation FormLocation `json:"formLocation"`
}

// FormLocation represents where a form is displayed.
type FormLocation string

const (
	FormLocationContent FormLocation = "Content"
	FormLocationPopup   FormLocation = "Popup"
	FormLocationModal   FormLocation = "Modal"
)

// MicroflowClientAction calls a microflow.
type MicroflowClientAction struct {
	model.BaseElement
	MicroflowID model.ID `json:"microflowId"`
}

func (MicroflowClientAction) isClientAction() {}

// NanoflowClientAction calls a nanoflow.
type NanoflowClientAction struct {
	model.BaseElement
	NanoflowID model.ID `json:"nanoflowId"`
}

func (NanoflowClientAction) isClientAction() {}

// ClosePageClientAction closes the current page.
type ClosePageClientAction struct {
	model.BaseElement
}

func (ClosePageClientAction) isClientAction() {}

// SaveChangesClientAction saves changes.
type SaveChangesClientAction struct {
	model.BaseElement
	ClosePage bool `json:"closePage"`
}

func (SaveChangesClientAction) isClientAction() {}

// CancelChangesClientAction cancels changes.
type CancelChangesClientAction struct {
	model.BaseElement
	ClosePage bool `json:"closePage"`
}

func (CancelChangesClientAction) isClientAction() {}

// CreateObjectClientAction creates an object.
type CreateObjectClientAction struct {
	model.BaseElement
	EntityID model.ID `json:"entityId"`
	PageID   model.ID `json:"pageId,omitempty"`
}

func (CreateObjectClientAction) isClientAction() {}

// DeleteClientAction deletes an object.
type DeleteClientAction struct {
	model.BaseElement
	ClosePage bool `json:"closePage"`
}

func (DeleteClientAction) isClientAction() {}

// SignOutClientAction signs out the user.
type SignOutClientAction struct {
	model.BaseElement
}

func (SignOutClientAction) isClientAction() {}

// ShowHomePageClientAction shows the home page.
type ShowHomePageClientAction struct {
	model.BaseElement
}

func (ShowHomePageClientAction) isClientAction() {}

// LinkClientAction opens a link.
type LinkClientAction struct {
	model.BaseElement
	LinkType LinkType `json:"linkType"`
	Address  string   `json:"address,omitempty"`
}

func (LinkClientAction) isClientAction() {}

// LinkType represents the type of link.
type LinkType string

const (
	LinkTypeWeb   LinkType = "Web"
	LinkTypeEmail LinkType = "Email"
	LinkTypePhone LinkType = "Phone"
)

// Navigation Widgets

// NavigationTree represents a navigation tree widget.
type NavigationTree struct {
	BaseWidget
	Items []*NavigationItem `json:"items,omitempty"`
}

// NavigationItem represents an item in navigation.
type NavigationItem struct {
	model.BaseElement
	Caption  *model.Text  `json:"caption,omitempty"`
	Icon     *Icon        `json:"icon,omitempty"`
	Action   ClientAction `json:"action,omitempty"`
	SubItems []*NavigationItem `json:"subItems,omitempty"`
}

// MenuBar represents a menu bar widget.
type MenuBar struct {
	BaseWidget
	MenuSource MenuSource `json:"menuSource,omitempty"`
}

// MenuSource represents the source for a menu.
type MenuSource interface {
	isMenuSource()
}

// NavigationMenuSource uses a navigation profile.
type NavigationMenuSource struct {
	model.BaseElement
}

func (NavigationMenuSource) isMenuSource() {}

// CustomMenuSource uses custom items.
type CustomMenuSource struct {
	model.BaseElement
	Items []*NavigationItem `json:"items,omitempty"`
}

func (CustomMenuSource) isMenuSource() {}

// SimpleMenuBar represents a simple menu bar.
type SimpleMenuBar struct {
	BaseWidget
	Orientation MenuOrientation `json:"orientation"`
	MenuSource  MenuSource      `json:"menuSource,omitempty"`
}

// MenuOrientation represents menu orientation.
type MenuOrientation string

const (
	MenuOrientationHorizontal MenuOrientation = "Horizontal"
	MenuOrientationVertical   MenuOrientation = "Vertical"
)

// Table Widget

// Table represents a table widget.
type Table struct {
	BaseWidget
	Rows []*TableRow `json:"rows,omitempty"`
}

// TableRow represents a row in a table.
type TableRow struct {
	model.BaseElement
	Cells []*TableCell `json:"cells,omitempty"`
}

// TableCell represents a cell in a table.
type TableCell struct {
	model.BaseElement
	ColumnSpan int      `json:"columnSpan"`
	RowSpan    int      `json:"rowSpan"`
	Widgets    []Widget `json:"widgets,omitempty"`
}

// Pluggable Widgets (Custom Widgets)

// PluggableWidget represents a pluggable/custom widget.
type PluggableWidget struct {
	BaseWidget
	WidgetID   string                     `json:"widgetId"`
	Properties map[string]interface{}     `json:"properties,omitempty"`
}

// HTMLSnippet represents an HTML snippet widget.
type HTMLSnippet struct {
	BaseWidget
	Type    HTMLSnippetType `json:"type"`
	Content string          `json:"content,omitempty"`
	ExternalURL string      `json:"externalUrl,omitempty"`
}

// HTMLSnippetType represents the type of HTML snippet.
type HTMLSnippetType string

const (
	HTMLSnippetTypeHTML     HTMLSnippetType = "HTML"
	HTMLSnippetTypeScript   HTMLSnippetType = "Script"
	HTMLSnippetTypeStyle    HTMLSnippetType = "Style"
	HTMLSnippetTypeExternal HTMLSnippetType = "External"
)

// SnippetCallWidget represents a snippet call widget.
type SnippetCallWidget struct {
	BaseWidget
	SnippetID model.ID `json:"snippetId"`
}

// Report Widgets

// ReportPane represents a report pane widget.
type ReportPane struct {
	BaseWidget
	ReportWidgets []Widget `json:"reportWidgets,omitempty"`
}

// ReportChart represents a report chart widget.
type ReportChart struct {
	BaseWidget
	ChartType ReportChartType `json:"chartType"`
}

// ReportChartType represents the type of report chart.
type ReportChartType string

const (
	ReportChartTypeLine ReportChartType = "Line"
	ReportChartTypeBar  ReportChartType = "Bar"
	ReportChartTypePie  ReportChartType = "Pie"
)

// ReportParameter represents a report parameter widget.
type ReportParameter struct {
	BaseWidget
	ParameterID model.ID `json:"parameterId"`
}

// ReportDateRangeSelector represents a date range selector for reports.
type ReportDateRangeSelector struct {
	BaseWidget
	FromParameter model.ID `json:"fromParameter"`
	ToParameter   model.ID `json:"toParameter"`
}

// Visibility and Conditional Rendering

// ConditionalContainer represents a conditionally visible container.
type ConditionalContainer struct {
	BaseWidget
	Condition         VisibilityCondition `json:"condition,omitempty"`
	Widgets           []Widget            `json:"widgets,omitempty"`
	AlternativeWidgets []Widget           `json:"alternativeWidgets,omitempty"`
}

// VisibilityCondition represents a visibility condition.
type VisibilityCondition interface {
	isVisibilityCondition()
}

// AttributeCondition bases visibility on an attribute.
type AttributeCondition struct {
	model.BaseElement
	AttributePath string `json:"attributePath"`
	Operator      ConditionOperator `json:"operator"`
	Value         string `json:"value,omitempty"`
}

func (AttributeCondition) isVisibilityCondition() {}

// ConditionOperator represents a condition operator.
type ConditionOperator string

const (
	ConditionOperatorEquals    ConditionOperator = "Equals"
	ConditionOperatorNotEquals ConditionOperator = "NotEquals"
	ConditionOperatorEmpty     ConditionOperator = "Empty"
	ConditionOperatorNotEmpty  ConditionOperator = "NotEmpty"
)

// ModuleRoleCondition bases visibility on module roles.
type ModuleRoleCondition struct {
	model.BaseElement
	Roles []model.ID `json:"roles,omitempty"`
}

func (ModuleRoleCondition) isVisibilityCondition() {}

// ExpressionCondition uses an expression.
type ExpressionCondition struct {
	model.BaseElement
	Expression string `json:"expression"`
}

func (ExpressionCondition) isVisibilityCondition() {}
