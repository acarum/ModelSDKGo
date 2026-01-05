// Package microflows provides types for Mendix microflows and nanoflows.
package microflows

import (
	"github.com/anthropics/modelsdk-go/model"
)

// Microflow represents a microflow in the Mendix model.
type Microflow struct {
	model.BaseElement
	ContainerID          model.ID   `json:"containerId"`
	Name                 string     `json:"name"`
	Documentation        string     `json:"documentation,omitempty"`
	AllowConcurrentExecution bool   `json:"allowConcurrentExecution"`
	MarkAsUsed           bool       `json:"markAsUsed"`
	Excluded             bool       `json:"excluded"`

	// Return type
	ReturnType DataType `json:"returnType,omitempty"`

	// Parameters
	Parameters []*MicroflowParameter `json:"parameters,omitempty"`

	// Flow elements
	ObjectCollection *MicroflowObjectCollection `json:"objectCollection,omitempty"`

	// Allowed module roles for execution
	AllowedModuleRoles []model.ID `json:"allowedModuleRoles,omitempty"`

	// Concurrent execution settings
	ConcurrentExecutionSettings *ConcurrentExecutionSettings `json:"concurrentExecutionSettings,omitempty"`
}

// GetName returns the microflow's name.
func (m *Microflow) GetName() string {
	return m.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (m *Microflow) GetContainerID() model.ID {
	return m.ContainerID
}

// Nanoflow represents a nanoflow in the Mendix model.
// Nanoflows run on the client side and have restrictions on which activities can be used.
type Nanoflow struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`
	MarkAsUsed    bool     `json:"markAsUsed"`
	Excluded      bool     `json:"excluded"`

	// Return type
	ReturnType DataType `json:"returnType,omitempty"`

	// Parameters
	Parameters []*MicroflowParameter `json:"parameters,omitempty"`

	// Flow elements
	ObjectCollection *MicroflowObjectCollection `json:"objectCollection,omitempty"`
}

// GetName returns the nanoflow's name.
func (n *Nanoflow) GetName() string {
	return n.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (n *Nanoflow) GetContainerID() model.ID {
	return n.ContainerID
}

// Rule represents a rule in the Mendix model.
type Rule struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`

	// Return type (always boolean)
	ReturnType DataType `json:"returnType,omitempty"`

	// Parameters
	Parameters []*MicroflowParameter `json:"parameters,omitempty"`

	// Flow elements
	ObjectCollection *MicroflowObjectCollection `json:"objectCollection,omitempty"`
}

// GetName returns the rule's name.
func (r *Rule) GetName() string {
	return r.Name
}

// GetContainerID returns the ID of the containing folder/module.
func (r *Rule) GetContainerID() model.ID {
	return r.ContainerID
}

// MicroflowParameter represents a parameter of a microflow.
type MicroflowParameter struct {
	model.BaseElement
	ContainerID   model.ID `json:"containerId"`
	Name          string   `json:"name"`
	Documentation string   `json:"documentation,omitempty"`
	Type          DataType `json:"type"`
}

// GetName returns the parameter's name.
func (p *MicroflowParameter) GetName() string {
	return p.Name
}

// GetContainerID returns the ID of the containing microflow.
func (p *MicroflowParameter) GetContainerID() model.ID {
	return p.ContainerID
}

// MicroflowObjectCollection contains all objects and flows in a microflow.
type MicroflowObjectCollection struct {
	model.BaseElement
	Objects []MicroflowObject `json:"objects,omitempty"`
}

// MicroflowObject is the base interface for all microflow objects.
type MicroflowObject interface {
	GetID() model.ID
	GetPosition() model.Point
}

// BaseMicroflowObject provides common fields for microflow objects.
type BaseMicroflowObject struct {
	model.BaseElement
	Position       model.Point `json:"position"`
	Size           model.Size  `json:"size,omitempty"`
	RelativeMiddle model.Point `json:"relativeMiddle,omitempty"`
}

// GetPosition returns the object's position.
func (o *BaseMicroflowObject) GetPosition() model.Point {
	return o.Position
}

// SequenceFlow represents a flow connection between objects.
type SequenceFlow struct {
	model.BaseElement
	OriginID        model.ID `json:"originId"`
	DestinationID   model.ID `json:"destinationId"`
	OriginConnectionIndex int `json:"originConnectionIndex"`
	DestinationConnectionIndex int `json:"destinationConnectionIndex"`
	CaseValue       *CaseValue `json:"caseValue,omitempty"`
	IsErrorHandler  bool       `json:"isErrorHandler,omitempty"`
}

// CaseValue represents a case value for a decision flow.
type CaseValue interface {
	isCaseValue()
}

// NoCase represents no case (default flow).
type NoCase struct {
	model.BaseElement
}

func (NoCase) isCaseValue() {}

// EnumerationCase represents an enumeration case value.
type EnumerationCase struct {
	model.BaseElement
	Value string `json:"value"`
}

func (EnumerationCase) isCaseValue() {}

// InheritanceCase represents an inheritance/type case value.
type InheritanceCase struct {
	model.BaseElement
	EntityID model.ID `json:"entityId"`
}

func (InheritanceCase) isCaseValue() {}

// BooleanCase represents a boolean case value.
type BooleanCase struct {
	model.BaseElement
	Value bool `json:"value"`
}

func (BooleanCase) isCaseValue() {}

// Annotation represents an annotation in a microflow.
type Annotation struct {
	BaseMicroflowObject
	Caption string `json:"caption"`
}

// AnnotationFlow connects an annotation to an object.
type AnnotationFlow struct {
	model.BaseElement
	OriginID      model.ID `json:"originId"`
	DestinationID model.ID `json:"destinationId"`
}

// Events

// StartEvent represents the start of a microflow.
type StartEvent struct {
	BaseMicroflowObject
}

// EndEvent represents the end of a microflow.
type EndEvent struct {
	BaseMicroflowObject
	ReturnValue string `json:"returnValue,omitempty"`
}

// ContinueEvent represents a continue in a loop.
type ContinueEvent struct {
	BaseMicroflowObject
}

// BreakEvent represents a break in a loop.
type BreakEvent struct {
	BaseMicroflowObject
}

// ErrorEvent represents an error throw in a microflow.
type ErrorEvent struct {
	BaseMicroflowObject
}

// Decisions and Control Flow

// ExclusiveSplit represents an exclusive decision (if/else).
type ExclusiveSplit struct {
	BaseMicroflowObject
	Caption           string   `json:"caption,omitempty"`
	Documentation     string   `json:"documentation,omitempty"`
	SplitCondition    SplitCondition `json:"splitCondition,omitempty"`
	ErrorHandlingType ErrorHandlingType `json:"errorHandlingType,omitempty"`
}

// ExclusiveMerge represents a merge point for exclusive splits.
type ExclusiveMerge struct {
	BaseMicroflowObject
}

// InheritanceSplit represents a type-based decision.
type InheritanceSplit struct {
	BaseMicroflowObject
	Caption       string   `json:"caption,omitempty"`
	Documentation string   `json:"documentation,omitempty"`
	VariableName  string   `json:"variableName"`
	ErrorHandlingType ErrorHandlingType `json:"errorHandlingType,omitempty"`
}

// SplitCondition represents the condition for a split.
type SplitCondition interface {
	isSplitCondition()
}

// ExpressionSplitCondition represents an expression-based split condition.
type ExpressionSplitCondition struct {
	model.BaseElement
	Expression string `json:"expression"`
}

func (ExpressionSplitCondition) isSplitCondition() {}

// RuleSplitCondition represents a rule-based split condition.
type RuleSplitCondition struct {
	model.BaseElement
	RuleID model.ID `json:"ruleId"`
	ParameterMappings []*RuleCallParameterMapping `json:"parameterMappings,omitempty"`
}

func (RuleSplitCondition) isSplitCondition() {}

// RuleCallParameterMapping maps a parameter to a value.
type RuleCallParameterMapping struct {
	model.BaseElement
	ParameterID model.ID `json:"parameterId"`
	Argument    string   `json:"argument"`
}

// LoopedActivity represents a loop construct.
type LoopedActivity struct {
	BaseMicroflowObject
	Caption       string `json:"caption,omitempty"`
	Documentation string `json:"documentation,omitempty"`
	IteratedListVariableName string `json:"iteratedListVariableName"`
	LoopVariableName string `json:"loopVariableName"`
	ObjectCollection *MicroflowObjectCollection `json:"objectCollection,omitempty"`
	ErrorHandlingType ErrorHandlingType `json:"errorHandlingType,omitempty"`
}

// ErrorHandlingType represents how errors are handled.
type ErrorHandlingType string

const (
	ErrorHandlingTypeAbort    ErrorHandlingType = "Abort"
	ErrorHandlingTypeContinue ErrorHandlingType = "Continue"
	ErrorHandlingTypeCustom   ErrorHandlingType = "Custom"
	ErrorHandlingTypeRollback ErrorHandlingType = "Rollback"
)

// Activities

// Activity is the base interface for all activities.
type Activity interface {
	MicroflowObject
	IsActivity()
}

// BaseActivity provides common fields for activities.
type BaseActivity struct {
	BaseMicroflowObject
	Caption           string            `json:"caption,omitempty"`
	Documentation     string            `json:"documentation,omitempty"`
	ErrorHandlingType ErrorHandlingType `json:"errorHandlingType,omitempty"`
	AutoGenerateCaption bool            `json:"autoGenerateCaption"`
}

// IsActivity marks this as an activity.
func (a *BaseActivity) IsActivity() {}

// ActionActivity wraps an action.
type ActionActivity struct {
	BaseActivity
	Action MicroflowAction `json:"action,omitempty"`
}

// MicroflowAction is the base interface for all microflow actions.
type MicroflowAction interface {
	isMicroflowAction()
}

// Object Actions

// CreateObjectAction creates a new object.
type CreateObjectAction struct {
	model.BaseElement
	EntityID       model.ID `json:"entityId"`
	OutputVariable string   `json:"outputVariable,omitempty"`
	Commit         CommitType `json:"commit"`
	InitialMembers []*MemberChange `json:"initialMembers,omitempty"`
}

func (CreateObjectAction) isMicroflowAction() {}

// ChangeObjectAction changes an existing object.
type ChangeObjectAction struct {
	model.BaseElement
	ChangeVariable string         `json:"changeVariable"`
	Commit         CommitType     `json:"commit"`
	RefreshInClient bool          `json:"refreshInClient"`
	Changes        []*MemberChange `json:"changes,omitempty"`
}

func (ChangeObjectAction) isMicroflowAction() {}

// DeleteObjectAction deletes an object.
type DeleteObjectAction struct {
	model.BaseElement
	DeleteVariable  string `json:"deleteVariable"`
	RefreshInClient bool   `json:"refreshInClient"`
}

func (DeleteObjectAction) isMicroflowAction() {}

// CommitObjectsAction commits one or more objects.
type CommitObjectsAction struct {
	model.BaseElement
	CommitVariable  string `json:"commitVariable"`
	WithEvents      bool   `json:"withEvents"`
	RefreshInClient bool   `json:"refreshInClient"`
}

func (CommitObjectsAction) isMicroflowAction() {}

// RollbackObjectAction rolls back an object.
type RollbackObjectAction struct {
	model.BaseElement
	RollbackVariable string `json:"rollbackVariable"`
	RefreshInClient  bool   `json:"refreshInClient"`
}

func (RollbackObjectAction) isMicroflowAction() {}

// MemberChange represents a change to a member (attribute or association).
type MemberChange struct {
	model.BaseElement
	AttributeID   model.ID       `json:"attributeId,omitempty"`
	AssociationID model.ID       `json:"associationId,omitempty"`
	Type          MemberChangeType `json:"type"`
	Value         string         `json:"value,omitempty"`
}

// MemberChangeType represents how a member is changed.
type MemberChangeType string

const (
	MemberChangeTypeSet MemberChangeType = "Set"
	MemberChangeTypeAdd MemberChangeType = "Add"
	MemberChangeTypeRemove MemberChangeType = "Remove"
)

// CommitType represents how objects are committed.
type CommitType string

const (
	CommitTypeYes CommitType = "Yes"
	CommitTypeNo CommitType = "No"
	CommitTypeYesWithEvents CommitType = "YesWithEvents"
	CommitTypeNoEvent CommitType = "NoEvent"
)

// Retrieve Actions

// RetrieveAction retrieves objects from the database.
type RetrieveAction struct {
	model.BaseElement
	OutputVariable string        `json:"outputVariable"`
	Source         RetrieveSource `json:"source,omitempty"`
}

func (RetrieveAction) isMicroflowAction() {}

// RetrieveSource represents the source for a retrieve action.
type RetrieveSource interface {
	isRetrieveSource()
}

// DatabaseRetrieveSource retrieves from the database.
type DatabaseRetrieveSource struct {
	model.BaseElement
	EntityID      model.ID `json:"entityId"`
	XPathConstraint string `json:"xPathConstraint,omitempty"`
	Range         *Range   `json:"range,omitempty"`
	Sorting       []*SortItem `json:"sorting,omitempty"`
}

func (DatabaseRetrieveSource) isRetrieveSource() {}

// AssociationRetrieveSource retrieves via association.
type AssociationRetrieveSource struct {
	model.BaseElement
	StartVariable string   `json:"startVariable"`
	AssociationID model.ID `json:"associationId"`
}

func (AssociationRetrieveSource) isRetrieveSource() {}

// Range specifies a range for retrieval.
type Range struct {
	model.BaseElement
	RangeType   RangeType `json:"rangeType"`
	Limit       string    `json:"limit,omitempty"`
	Offset      string    `json:"offset,omitempty"`
}

// RangeType represents the type of range.
type RangeType string

const (
	RangeTypeAll    RangeType = "All"
	RangeTypeFirst  RangeType = "First"
	RangeTypeCustom RangeType = "Custom"
)

// SortItem represents a sort specification.
type SortItem struct {
	model.BaseElement
	AttributeID model.ID   `json:"attributeId"`
	Direction   SortDirection `json:"direction"`
}

// SortDirection represents sort order.
type SortDirection string

const (
	SortDirectionAscending  SortDirection = "Ascending"
	SortDirectionDescending SortDirection = "Descending"
)

// AggregateListAction aggregates a list.
type AggregateListAction struct {
	model.BaseElement
	InputVariable  string          `json:"inputVariable"`
	OutputVariable string          `json:"outputVariable"`
	Function       AggregateFunction `json:"function"`
	AttributeID    model.ID        `json:"attributeId,omitempty"`
}

func (AggregateListAction) isMicroflowAction() {}

// AggregateFunction represents an aggregate function.
type AggregateFunction string

const (
	AggregateFunctionCount   AggregateFunction = "Count"
	AggregateFunctionSum     AggregateFunction = "Sum"
	AggregateFunctionAverage AggregateFunction = "Average"
	AggregateFunctionMin     AggregateFunction = "Minimum"
	AggregateFunctionMax     AggregateFunction = "Maximum"
)

// ListOperationAction performs list operations.
type ListOperationAction struct {
	model.BaseElement
	Operation      ListOperation `json:"operation,omitempty"`
	OutputVariable string        `json:"outputVariable,omitempty"`
}

func (ListOperationAction) isMicroflowAction() {}

// ListOperation represents a list operation.
type ListOperation interface {
	isListOperation()
}

// HeadOperation gets the first element.
type HeadOperation struct {
	model.BaseElement
	ListVariable string `json:"listVariable"`
}

func (HeadOperation) isListOperation() {}

// TailOperation gets all but the first element.
type TailOperation struct {
	model.BaseElement
	ListVariable string `json:"listVariable"`
}

func (TailOperation) isListOperation() {}

// FindOperation finds an element.
type FindOperation struct {
	model.BaseElement
	ListVariable string `json:"listVariable"`
	Expression   string `json:"expression"`
}

func (FindOperation) isListOperation() {}

// FilterOperation filters a list.
type FilterOperation struct {
	model.BaseElement
	ListVariable string `json:"listVariable"`
	Expression   string `json:"expression"`
}

func (FilterOperation) isListOperation() {}

// SortOperation sorts a list.
type SortOperation struct {
	model.BaseElement
	ListVariable string      `json:"listVariable"`
	Sorting      []*SortItem `json:"sorting,omitempty"`
}

func (SortOperation) isListOperation() {}

// UnionOperation unions two lists.
type UnionOperation struct {
	model.BaseElement
	ListVariable1 string `json:"listVariable1"`
	ListVariable2 string `json:"listVariable2"`
}

func (UnionOperation) isListOperation() {}

// IntersectOperation intersects two lists.
type IntersectOperation struct {
	model.BaseElement
	ListVariable1 string `json:"listVariable1"`
	ListVariable2 string `json:"listVariable2"`
}

func (IntersectOperation) isListOperation() {}

// SubtractOperation subtracts one list from another.
type SubtractOperation struct {
	model.BaseElement
	ListVariable1 string `json:"listVariable1"`
	ListVariable2 string `json:"listVariable2"`
}

func (SubtractOperation) isListOperation() {}

// ContainsOperation checks if a list contains an object.
type ContainsOperation struct {
	model.BaseElement
	ListVariable   string `json:"listVariable"`
	ObjectVariable string `json:"objectVariable"`
}

func (ContainsOperation) isListOperation() {}

// EqualsOperation checks if two lists are equal.
type EqualsOperation struct {
	model.BaseElement
	ListVariable1 string `json:"listVariable1"`
	ListVariable2 string `json:"listVariable2"`
}

func (EqualsOperation) isListOperation() {}

// Variable Actions

// CreateVariableAction creates a variable.
type CreateVariableAction struct {
	model.BaseElement
	VariableName string   `json:"variableName"`
	DataType     DataType `json:"dataType,omitempty"`
	InitialValue string   `json:"initialValue,omitempty"`
}

func (CreateVariableAction) isMicroflowAction() {}

// ChangeVariableAction changes a variable.
type ChangeVariableAction struct {
	model.BaseElement
	VariableName string `json:"variableName"`
	Value        string `json:"value"`
}

func (ChangeVariableAction) isMicroflowAction() {}

// CastAction casts an object to a more specific type.
type CastAction struct {
	model.BaseElement
	ObjectVariable string `json:"objectVariable"`
	OutputVariable string `json:"outputVariable"`
}

func (CastAction) isMicroflowAction() {}

// Client Actions

// ShowPageAction shows a page.
type ShowPageAction struct {
	model.BaseElement
	PageID             model.ID   `json:"pageId,omitempty"`
	PageSettings       *PageSettings `json:"pageSettings,omitempty"`
	PassedObject       string     `json:"passedObject,omitempty"`
	OverridePageTitle  *model.Text `json:"overridePageTitle,omitempty"`
}

func (ShowPageAction) isMicroflowAction() {}

// PageSettings represents page display settings.
type PageSettings struct {
	model.BaseElement
	Location  PageLocation `json:"location"`
	ModalForm bool         `json:"modalForm"`
}

// PageLocation represents where a page is shown.
type PageLocation string

const (
	PageLocationContent PageLocation = "Content"
	PageLocationPopup   PageLocation = "Popup"
	PageLocationModal   PageLocation = "Modal"
)

// ShowHomePageAction shows the home page.
type ShowHomePageAction struct {
	model.BaseElement
}

func (ShowHomePageAction) isMicroflowAction() {}

// ClosePageAction closes the current page.
type ClosePageAction struct {
	model.BaseElement
	NumberOfPages int `json:"numberOfPages"`
}

func (ClosePageAction) isMicroflowAction() {}

// ShowMessageAction shows a message to the user.
type ShowMessageAction struct {
	model.BaseElement
	Template *model.Text `json:"template,omitempty"`
	Type     MessageType `json:"type"`
	Blocking bool        `json:"blocking"`
}

func (ShowMessageAction) isMicroflowAction() {}

// MessageType represents the type of message.
type MessageType string

const (
	MessageTypeInformation MessageType = "Information"
	MessageTypeWarning     MessageType = "Warning"
	MessageTypeError       MessageType = "Error"
)

// ValidationFeedbackAction shows validation feedback.
type ValidationFeedbackAction struct {
	model.BaseElement
	ObjectVariable string      `json:"objectVariable"`
	AttributeID    model.ID    `json:"attributeId,omitempty"`
	AssociationID  model.ID    `json:"associationId,omitempty"`
	Template       *model.Text `json:"template,omitempty"`
}

func (ValidationFeedbackAction) isMicroflowAction() {}

// DownloadFileAction downloads a file.
type DownloadFileAction struct {
	model.BaseElement
	FileDocument string `json:"fileDocument"`
	ShowInBrowser bool  `json:"showInBrowser"`
}

func (DownloadFileAction) isMicroflowAction() {}

// Integration Actions

// MicroflowCallAction calls another microflow.
type MicroflowCallAction struct {
	model.BaseElement
	MicroflowID        model.ID                      `json:"microflowId"`
	OutputVariable     string                        `json:"outputVariable,omitempty"`
	UseReturnVariable  bool                          `json:"useReturnVariable"`
	ParameterMappings  []*MicroflowCallParameterMapping `json:"parameterMappings,omitempty"`
}

func (MicroflowCallAction) isMicroflowAction() {}

// MicroflowCallParameterMapping maps a parameter to an argument.
type MicroflowCallParameterMapping struct {
	model.BaseElement
	ParameterID model.ID `json:"parameterId"`
	Argument    string   `json:"argument"`
}

// JavaActionCallAction calls a Java action.
type JavaActionCallAction struct {
	model.BaseElement
	JavaActionID       model.ID                     `json:"javaActionId"`
	OutputVariable     string                       `json:"outputVariable,omitempty"`
	UseReturnVariable  bool                         `json:"useReturnVariable"`
	ParameterMappings  []*JavaActionParameterMapping `json:"parameterMappings,omitempty"`
}

func (JavaActionCallAction) isMicroflowAction() {}

// JavaActionParameterMapping maps a Java action parameter.
type JavaActionParameterMapping struct {
	model.BaseElement
	ParameterID model.ID `json:"parameterId"`
	Argument    string   `json:"argument"`
}

// WebServiceCallAction calls a web service.
type WebServiceCallAction struct {
	model.BaseElement
	ServiceID              model.ID `json:"serviceId,omitempty"`
	OperationName          string   `json:"operationName,omitempty"`
	SendMappingID          model.ID `json:"sendMappingId,omitempty"`
	ReceiveMappingID       model.ID `json:"receiveMappingId,omitempty"`
	OutputVariable         string   `json:"outputVariable,omitempty"`
	UseReturnVariable      bool     `json:"useReturnVariable"`
	TimeoutExpression      string   `json:"timeoutExpression,omitempty"`
}

func (WebServiceCallAction) isMicroflowAction() {}

// RestCallAction calls a REST service.
type RestCallAction struct {
	model.BaseElement
	HttpConfiguration  *HttpConfiguration `json:"httpConfiguration,omitempty"`
	RequestHandling    RequestHandling    `json:"requestHandling,omitempty"`
	ResultHandling     ResultHandling     `json:"resultHandling,omitempty"`
	ErrorHandling      *ErrorHandling     `json:"errorHandling,omitempty"`
	OutputVariable     string             `json:"outputVariable,omitempty"`
	UseReturnVariable  bool               `json:"useReturnVariable"`
	TimeoutExpression  string             `json:"timeoutExpression,omitempty"`
}

func (RestCallAction) isMicroflowAction() {}

// HttpConfiguration represents HTTP configuration for a REST call.
type HttpConfiguration struct {
	model.BaseElement
	HttpMethod     HttpMethod   `json:"httpMethod"`
	LocationTemplate string     `json:"locationTemplate,omitempty"`
	CustomLocation string       `json:"customLocation,omitempty"`
	CustomHeaders  []*HttpHeader `json:"customHeaders,omitempty"`
}

// HttpMethod represents an HTTP method.
type HttpMethod string

const (
	HttpMethodGet    HttpMethod = "GET"
	HttpMethodPost   HttpMethod = "POST"
	HttpMethodPut    HttpMethod = "PUT"
	HttpMethodPatch  HttpMethod = "PATCH"
	HttpMethodDelete HttpMethod = "DELETE"
)

// HttpHeader represents an HTTP header.
type HttpHeader struct {
	model.BaseElement
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RequestHandling represents how a request is handled.
type RequestHandling interface {
	isRequestHandling()
}

// SimpleRequestHandling represents simple request handling.
type SimpleRequestHandling struct {
	model.BaseElement
	ParameterEntityID model.ID `json:"parameterEntityId,omitempty"`
}

func (SimpleRequestHandling) isRequestHandling() {}

// MappingRequestHandling uses an export mapping.
type MappingRequestHandling struct {
	model.BaseElement
	MappingID        model.ID `json:"mappingId"`
	ContentType      string   `json:"contentType,omitempty"`
	ParameterVariable string  `json:"parameterVariable,omitempty"`
}

func (MappingRequestHandling) isRequestHandling() {}

// CustomRequestHandling uses a custom body.
type CustomRequestHandling struct {
	model.BaseElement
	Template string `json:"template,omitempty"`
}

func (CustomRequestHandling) isRequestHandling() {}

// ResultHandling represents how a result is handled.
type ResultHandling interface {
	isResultHandling()
}

// ResultHandlingNone represents no result handling.
type ResultHandlingNone struct {
	model.BaseElement
}

func (ResultHandlingNone) isResultHandling() {}

// ResultHandlingString returns the result as a string.
type ResultHandlingString struct {
	model.BaseElement
}

func (ResultHandlingString) isResultHandling() {}

// ResultHandlingMapping uses an import mapping.
type ResultHandlingMapping struct {
	model.BaseElement
	MappingID        model.ID `json:"mappingId"`
	ResultEntityID   model.ID `json:"resultEntityId,omitempty"`
	ResultVariable   string   `json:"resultVariable,omitempty"`
}

func (ResultHandlingMapping) isResultHandling() {}

// ErrorHandling represents error handling configuration.
type ErrorHandling struct {
	model.BaseElement
	ErrorHandlingType ErrorHandlingType `json:"errorHandlingType"`
}

// LogMessageAction logs a message.
type LogMessageAction struct {
	model.BaseElement
	LogLevel     LogLevel    `json:"logLevel"`
	LogNodeName  string      `json:"logNodeName,omitempty"`
	MessageTemplate *model.Text `json:"messageTemplate,omitempty"`
	IncludeLastStackTrace bool `json:"includeLastStackTrace"`
}

func (LogMessageAction) isMicroflowAction() {}

// LogLevel represents a log level.
type LogLevel string

const (
	LogLevelTrace    LogLevel = "Trace"
	LogLevelDebug    LogLevel = "Debug"
	LogLevelInfo     LogLevel = "Info"
	LogLevelWarning  LogLevel = "Warning"
	LogLevelError    LogLevel = "Error"
	LogLevelCritical LogLevel = "Critical"
)

// ImportMappingCallAction calls an import mapping.
type ImportMappingCallAction struct {
	model.BaseElement
	MappingID        model.ID `json:"mappingId"`
	SourceVariable   string   `json:"sourceVariable,omitempty"`
	OutputVariable   string   `json:"outputVariable,omitempty"`
	ContentType      string   `json:"contentType,omitempty"`
}

func (ImportMappingCallAction) isMicroflowAction() {}

// ExportMappingCallAction calls an export mapping.
type ExportMappingCallAction struct {
	model.BaseElement
	MappingID        model.ID `json:"mappingId"`
	SourceVariable   string   `json:"sourceVariable,omitempty"`
	OutputVariable   string   `json:"outputVariable,omitempty"`
	ContentType      string   `json:"contentType,omitempty"`
}

func (ExportMappingCallAction) isMicroflowAction() {}

// ConcurrentExecutionSettings represents settings for concurrent execution.
type ConcurrentExecutionSettings struct {
	model.BaseElement
	Enabled        bool   `json:"enabled"`
	NumberOfThreads int   `json:"numberOfThreads,omitempty"`
}

// Data Types

// DataType represents a data type in Mendix.
type DataType interface {
	isDataType()
	GetTypeName() string
}

// BooleanType represents a boolean type.
type BooleanType struct {
	model.BaseElement
}

func (BooleanType) isDataType()             {}
func (BooleanType) GetTypeName() string { return "Boolean" }

// IntegerType represents an integer type.
type IntegerType struct {
	model.BaseElement
}

func (IntegerType) isDataType()             {}
func (IntegerType) GetTypeName() string { return "Integer" }

// LongType represents a long type.
type LongType struct {
	model.BaseElement
}

func (LongType) isDataType()             {}
func (LongType) GetTypeName() string { return "Long" }

// DecimalType represents a decimal type.
type DecimalType struct {
	model.BaseElement
}

func (DecimalType) isDataType()             {}
func (DecimalType) GetTypeName() string { return "Decimal" }

// StringType represents a string type.
type StringType struct {
	model.BaseElement
}

func (StringType) isDataType()             {}
func (StringType) GetTypeName() string { return "String" }

// DateTimeType represents a date/time type.
type DateTimeType struct {
	model.BaseElement
}

func (DateTimeType) isDataType()             {}
func (DateTimeType) GetTypeName() string { return "DateTime" }

// ObjectType represents an object type.
type ObjectType struct {
	model.BaseElement
	EntityID model.ID `json:"entityId"`
}

func (ObjectType) isDataType()             {}
func (ObjectType) GetTypeName() string { return "Object" }

// ListType represents a list type.
type ListType struct {
	model.BaseElement
	EntityID model.ID `json:"entityId"`
}

func (ListType) isDataType()             {}
func (ListType) GetTypeName() string { return "List" }

// EnumerationType represents an enumeration type.
type EnumerationType struct {
	model.BaseElement
	EnumerationID model.ID `json:"enumerationId"`
}

func (EnumerationType) isDataType()             {}
func (EnumerationType) GetTypeName() string { return "Enumeration" }

// VoidType represents no return type.
type VoidType struct {
	model.BaseElement
}

func (VoidType) isDataType()             {}
func (VoidType) GetTypeName() string { return "Void" }

// BinaryType represents a binary type.
type BinaryType struct {
	model.BaseElement
}

func (BinaryType) isDataType()             {}
func (BinaryType) GetTypeName() string { return "Binary" }
