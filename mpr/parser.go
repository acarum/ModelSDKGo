package mpr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anthropics/modelsdk-go/domainmodel"
	"github.com/anthropics/modelsdk-go/microflows"
	"github.com/anthropics/modelsdk-go/model"
	"github.com/anthropics/modelsdk-go/pages"

	"go.mongodb.org/mongo-driver/bson"
)

// parseModule parses module contents from BSON.
func (r *Reader) parseModule(unitID string, contents []byte) (*model.Module, error) {
	// For MPR v2, contents might be a reference to an external file
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	// Parse BSON contents
	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	module := &model.Module{}
	module.ID = model.ID(unitID)
	module.TypeName = "Projects$Module"

	if name, ok := raw["Name"].(string); ok {
		module.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		module.Documentation = doc
	}
	if excluded, ok := raw["Excluded"].(bool); ok {
		module.Excluded = excluded
	}
	if fromAppStore, ok := raw["FromAppStore"].(bool); ok {
		module.FromAppStore = fromAppStore
	}

	return module, nil
}

// parseDomainModel parses domain model contents from BSON.
func (r *Reader) parseDomainModel(unitID, containerID string, contents []byte) (*domainmodel.DomainModel, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	dm := &domainmodel.DomainModel{}
	dm.ID = model.ID(unitID)
	dm.TypeName = "DomainModels$DomainModel"
	dm.ContainerID = model.ID(containerID)

	// Parse entities
	if entities, ok := raw["Entities"].([]interface{}); ok {
		for _, e := range entities {
			if entityMap, ok := e.(map[string]interface{}); ok {
				entity := parseEntity(entityMap)
				dm.Entities = append(dm.Entities, entity)
			}
		}
	}

	// Parse associations
	if associations, ok := raw["Associations"].([]interface{}); ok {
		for _, a := range associations {
			if assocMap, ok := a.(map[string]interface{}); ok {
				assoc := parseAssociation(assocMap)
				dm.Associations = append(dm.Associations, assoc)
			}
		}
	}

	// Parse annotations
	if annotations, ok := raw["Annotations"].([]interface{}); ok {
		for _, a := range annotations {
			if annotMap, ok := a.(map[string]interface{}); ok {
				annot := parseAnnotation(annotMap)
				dm.Annotations = append(dm.Annotations, annot)
			}
		}
	}

	return dm, nil
}

func parseEntity(raw map[string]interface{}) *domainmodel.Entity {
	entity := &domainmodel.Entity{}

	if id, ok := raw["$ID"].(string); ok {
		entity.ID = model.ID(id)
	}
	if typeName, ok := raw["$Type"].(string); ok {
		entity.TypeName = typeName
	}
	if name, ok := raw["Name"].(string); ok {
		entity.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		entity.Documentation = doc
	}

	// Parse location
	if loc, ok := raw["Location"].(map[string]interface{}); ok {
		if x, ok := loc["x"].(int32); ok {
			entity.Location.X = int(x)
		}
		if y, ok := loc["y"].(int32); ok {
			entity.Location.Y = int(y)
		}
	}

	// Parse persistable
	if persistable, ok := raw["Persistable"].(bool); ok {
		entity.Persistable = persistable
	}

	// Parse attributes
	if attrs, ok := raw["Attributes"].([]interface{}); ok {
		for _, a := range attrs {
			if attrMap, ok := a.(map[string]interface{}); ok {
				attr := parseAttribute(attrMap)
				entity.Attributes = append(entity.Attributes, attr)
			}
		}
	}

	// Parse indexes
	if indexes, ok := raw["Indexes"].([]interface{}); ok {
		for _, i := range indexes {
			if indexMap, ok := i.(map[string]interface{}); ok {
				index := parseIndex(indexMap)
				entity.Indexes = append(entity.Indexes, index)
			}
		}
	}

	// Parse access rules
	if rules, ok := raw["AccessRules"].([]interface{}); ok {
		for _, r := range rules {
			if ruleMap, ok := r.(map[string]interface{}); ok {
				rule := parseAccessRule(ruleMap)
				entity.AccessRules = append(entity.AccessRules, rule)
			}
		}
	}

	// Parse validation rules
	if validations, ok := raw["ValidationRules"].([]interface{}); ok {
		for _, v := range validations {
			if validMap, ok := v.(map[string]interface{}); ok {
				validation := parseValidationRule(validMap)
				entity.ValidationRules = append(entity.ValidationRules, validation)
			}
		}
	}

	// Parse event handlers
	if handlers, ok := raw["EventHandlers"].([]interface{}); ok {
		for _, h := range handlers {
			if handlerMap, ok := h.(map[string]interface{}); ok {
				handler := parseEventHandler(handlerMap)
				entity.EventHandlers = append(entity.EventHandlers, handler)
			}
		}
	}

	return entity
}

func parseAttribute(raw map[string]interface{}) *domainmodel.Attribute {
	attr := &domainmodel.Attribute{}

	if id, ok := raw["$ID"].(string); ok {
		attr.ID = model.ID(id)
	}
	if typeName, ok := raw["$Type"].(string); ok {
		attr.TypeName = typeName
	}
	if name, ok := raw["Name"].(string); ok {
		attr.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		attr.Documentation = doc
	}

	// Parse attribute type
	if attrType, ok := raw["Type"].(map[string]interface{}); ok {
		attr.Type = parseAttributeType(attrType)
	}

	return attr
}

func parseAttributeType(raw map[string]interface{}) domainmodel.AttributeType {
	typeName, _ := raw["$Type"].(string)

	switch typeName {
	case "DomainModels$StringAttributeType":
		t := &domainmodel.StringAttributeType{}
		if length, ok := raw["Length"].(int32); ok {
			t.Length = int(length)
		}
		return t
	case "DomainModels$IntegerAttributeType":
		return &domainmodel.IntegerAttributeType{}
	case "DomainModels$LongAttributeType":
		return &domainmodel.LongAttributeType{}
	case "DomainModels$DecimalAttributeType":
		return &domainmodel.DecimalAttributeType{}
	case "DomainModels$BooleanAttributeType":
		return &domainmodel.BooleanAttributeType{}
	case "DomainModels$DateTimeAttributeType":
		t := &domainmodel.DateTimeAttributeType{}
		if localize, ok := raw["LocalizeDate"].(bool); ok {
			t.LocalizeDate = localize
		}
		return t
	case "DomainModels$EnumerationAttributeType":
		t := &domainmodel.EnumerationAttributeType{}
		if enumID, ok := raw["Enumeration"].(string); ok {
			t.EnumerationID = model.ID(enumID)
		}
		return t
	case "DomainModels$AutoNumberAttributeType":
		return &domainmodel.AutoNumberAttributeType{}
	case "DomainModels$BinaryAttributeType":
		return &domainmodel.BinaryAttributeType{}
	case "DomainModels$HashedStringAttributeType":
		return &domainmodel.HashedStringAttributeType{}
	default:
		return &domainmodel.StringAttributeType{} // Default fallback
	}
}

func parseAssociation(raw map[string]interface{}) *domainmodel.Association {
	assoc := &domainmodel.Association{}

	if id, ok := raw["$ID"].(string); ok {
		assoc.ID = model.ID(id)
	}
	if typeName, ok := raw["$Type"].(string); ok {
		assoc.TypeName = typeName
	}
	if name, ok := raw["Name"].(string); ok {
		assoc.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		assoc.Documentation = doc
	}
	if parentID, ok := raw["Parent"].(string); ok {
		assoc.ParentID = model.ID(parentID)
	}
	if childID, ok := raw["Child"].(string); ok {
		assoc.ChildID = model.ID(childID)
	}
	if assocType, ok := raw["Type"].(string); ok {
		assoc.Type = domainmodel.AssociationType(assocType)
	}
	if owner, ok := raw["Owner"].(string); ok {
		assoc.Owner = domainmodel.AssociationOwner(owner)
	}

	return assoc
}

func parseAnnotation(raw map[string]interface{}) *domainmodel.Annotation {
	annot := &domainmodel.Annotation{}

	if id, ok := raw["$ID"].(string); ok {
		annot.ID = model.ID(id)
	}
	if typeName, ok := raw["$Type"].(string); ok {
		annot.TypeName = typeName
	}
	if caption, ok := raw["Caption"].(string); ok {
		annot.Caption = caption
	}

	if loc, ok := raw["Location"].(map[string]interface{}); ok {
		if x, ok := loc["x"].(int32); ok {
			annot.Location.X = int(x)
		}
		if y, ok := loc["y"].(int32); ok {
			annot.Location.Y = int(y)
		}
	}

	return annot
}

func parseIndex(raw map[string]interface{}) *domainmodel.Index {
	index := &domainmodel.Index{}

	if id, ok := raw["$ID"].(string); ok {
		index.ID = model.ID(id)
	}
	if name, ok := raw["Name"].(string); ok {
		index.Name = name
	}

	return index
}

func parseAccessRule(raw map[string]interface{}) *domainmodel.AccessRule {
	rule := &domainmodel.AccessRule{}

	if id, ok := raw["$ID"].(string); ok {
		rule.ID = model.ID(id)
	}
	if allowCreate, ok := raw["AllowCreate"].(bool); ok {
		rule.AllowCreate = allowCreate
	}
	if allowRead, ok := raw["AllowRead"].(bool); ok {
		rule.AllowRead = allowRead
	}
	if allowWrite, ok := raw["AllowWrite"].(bool); ok {
		rule.AllowWrite = allowWrite
	}
	if allowDelete, ok := raw["AllowDelete"].(bool); ok {
		rule.AllowDelete = allowDelete
	}
	if xpath, ok := raw["XPathConstraint"].(string); ok {
		rule.XPathConstraint = xpath
	}

	return rule
}

func parseValidationRule(raw map[string]interface{}) *domainmodel.ValidationRule {
	rule := &domainmodel.ValidationRule{}

	if id, ok := raw["$ID"].(string); ok {
		rule.ID = model.ID(id)
	}
	if attrID, ok := raw["Attribute"].(string); ok {
		rule.AttributeID = model.ID(attrID)
	}

	return rule
}

func parseEventHandler(raw map[string]interface{}) *domainmodel.EventHandler {
	handler := &domainmodel.EventHandler{}

	if id, ok := raw["$ID"].(string); ok {
		handler.ID = model.ID(id)
	}
	if event, ok := raw["Event"].(string); ok {
		handler.Event = domainmodel.EventType(event)
	}
	if mfID, ok := raw["Microflow"].(string); ok {
		handler.MicroflowID = model.ID(mfID)
	}
	if raiseError, ok := raw["RaiseErrorOnFalse"].(bool); ok {
		handler.RaiseErrorOnFalse = raiseError
	}

	return handler
}

// parseMicroflow parses microflow contents from BSON.
func (r *Reader) parseMicroflow(unitID, containerID string, contents []byte) (*microflows.Microflow, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	mf := &microflows.Microflow{}
	mf.ID = model.ID(unitID)
	mf.TypeName = "Microflows$Microflow"
	mf.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		mf.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		mf.Documentation = doc
	}
	if concurrent, ok := raw["AllowConcurrentExecution"].(bool); ok {
		mf.AllowConcurrentExecution = concurrent
	}
	if markAsUsed, ok := raw["MarkAsUsed"].(bool); ok {
		mf.MarkAsUsed = markAsUsed
	}
	if excluded, ok := raw["Excluded"].(bool); ok {
		mf.Excluded = excluded
	}

	// Parse parameters
	if params, ok := raw["Parameters"].([]interface{}); ok {
		for _, p := range params {
			if paramMap, ok := p.(map[string]interface{}); ok {
				param := parseMicroflowParameter(paramMap)
				mf.Parameters = append(mf.Parameters, param)
			}
		}
	}

	// Parse object collection (flow elements)
	if objCollection, ok := raw["ObjectCollection"].(map[string]interface{}); ok {
		mf.ObjectCollection = parseMicroflowObjectCollection(objCollection)
	}

	return mf, nil
}

func parseMicroflowParameter(raw map[string]interface{}) *microflows.MicroflowParameter {
	param := &microflows.MicroflowParameter{}

	if id, ok := raw["$ID"].(string); ok {
		param.ID = model.ID(id)
	}
	if name, ok := raw["Name"].(string); ok {
		param.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		param.Documentation = doc
	}

	return param
}

func parseMicroflowObjectCollection(raw map[string]interface{}) *microflows.MicroflowObjectCollection {
	collection := &microflows.MicroflowObjectCollection{}

	if id, ok := raw["$ID"].(string); ok {
		collection.ID = model.ID(id)
	}

	// Parse objects would go here - complex nested parsing

	return collection
}

// parseNanoflow parses nanoflow contents from BSON.
func (r *Reader) parseNanoflow(unitID, containerID string, contents []byte) (*microflows.Nanoflow, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	nf := &microflows.Nanoflow{}
	nf.ID = model.ID(unitID)
	nf.TypeName = "Microflows$Nanoflow"
	nf.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		nf.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		nf.Documentation = doc
	}
	if markAsUsed, ok := raw["MarkAsUsed"].(bool); ok {
		nf.MarkAsUsed = markAsUsed
	}
	if excluded, ok := raw["Excluded"].(bool); ok {
		nf.Excluded = excluded
	}

	// Parse parameters
	if params, ok := raw["Parameters"].([]interface{}); ok {
		for _, p := range params {
			if paramMap, ok := p.(map[string]interface{}); ok {
				param := parseMicroflowParameter(paramMap)
				nf.Parameters = append(nf.Parameters, param)
			}
		}
	}

	return nf, nil
}

// parsePage parses page contents from BSON.
func (r *Reader) parsePage(unitID, containerID string, contents []byte) (*pages.Page, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	page := &pages.Page{}
	page.ID = model.ID(unitID)
	page.TypeName = "Pages$Page"
	page.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		page.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		page.Documentation = doc
	}
	if url, ok := raw["URL"].(string); ok {
		page.URL = url
	}
	if layoutID, ok := raw["Layout"].(string); ok {
		page.LayoutID = model.ID(layoutID)
	}
	if markAsUsed, ok := raw["MarkAsUsed"].(bool); ok {
		page.MarkAsUsed = markAsUsed
	}
	if excluded, ok := raw["Excluded"].(bool); ok {
		page.Excluded = excluded
	}

	// Parse title
	if title, ok := raw["Title"].(map[string]interface{}); ok {
		page.Title = parseText(title)
	}

	// Parse parameters
	if params, ok := raw["Parameters"].([]interface{}); ok {
		for _, p := range params {
			if paramMap, ok := p.(map[string]interface{}); ok {
				param := parsePageParameter(paramMap)
				page.Parameters = append(page.Parameters, param)
			}
		}
	}

	return page, nil
}

func parseText(raw map[string]interface{}) *model.Text {
	text := &model.Text{}

	if id, ok := raw["$ID"].(string); ok {
		text.ID = model.ID(id)
	}

	text.Translations = make(map[string]string)
	if translations, ok := raw["Translations"].(map[string]interface{}); ok {
		for lang, val := range translations {
			if str, ok := val.(string); ok {
				text.Translations[lang] = str
			}
		}
	}

	return text
}

func parsePageParameter(raw map[string]interface{}) *pages.PageParameter {
	param := &pages.PageParameter{}

	if id, ok := raw["$ID"].(string); ok {
		param.ID = model.ID(id)
	}
	if name, ok := raw["Name"].(string); ok {
		param.Name = name
	}
	if entityID, ok := raw["Entity"].(string); ok {
		param.EntityID = model.ID(entityID)
	}

	return param
}

// parseLayout parses layout contents from BSON.
func (r *Reader) parseLayout(unitID, containerID string, contents []byte) (*pages.Layout, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	layout := &pages.Layout{}
	layout.ID = model.ID(unitID)
	layout.TypeName = "Pages$Layout"
	layout.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		layout.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		layout.Documentation = doc
	}
	if layoutType, ok := raw["LayoutType"].(string); ok {
		layout.LayoutType = pages.LayoutType(layoutType)
	}

	return layout, nil
}

// parseEnumeration parses enumeration contents from BSON.
func (r *Reader) parseEnumeration(unitID, containerID string, contents []byte) (*model.Enumeration, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	enum := &model.Enumeration{}
	enum.ID = model.ID(unitID)
	enum.TypeName = "Enumerations$Enumeration"
	enum.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		enum.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		enum.Documentation = doc
	}

	// Parse values
	if values, ok := raw["Values"].([]interface{}); ok {
		for _, v := range values {
			if valMap, ok := v.(map[string]interface{}); ok {
				value := parseEnumerationValue(valMap)
				enum.Values = append(enum.Values, value)
			}
		}
	}

	return enum, nil
}

func parseEnumerationValue(raw map[string]interface{}) model.EnumerationValue {
	value := model.EnumerationValue{}

	if id, ok := raw["$ID"].(string); ok {
		value.ID = model.ID(id)
	}
	if name, ok := raw["Name"].(string); ok {
		value.Name = name
	}
	if caption, ok := raw["Caption"].(map[string]interface{}); ok {
		value.Caption = parseText(caption)
	}

	return value
}

// parseConstant parses constant contents from BSON.
func (r *Reader) parseConstant(unitID, containerID string, contents []byte) (*model.Constant, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	constant := &model.Constant{}
	constant.ID = model.ID(unitID)
	constant.TypeName = "Constants$Constant"
	constant.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		constant.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		constant.Documentation = doc
	}
	if constType, ok := raw["Type"].(string); ok {
		constant.Type = constType
	}
	if defaultValue, ok := raw["DefaultValue"].(string); ok {
		constant.DefaultValue = defaultValue
	}
	if exposed, ok := raw["ExposedToClient"].(bool); ok {
		constant.ExposedToClient = exposed
	}

	return constant, nil
}

// parseScheduledEvent parses scheduled event contents from BSON.
func (r *Reader) parseScheduledEvent(unitID, containerID string, contents []byte) (*model.ScheduledEvent, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	event := &model.ScheduledEvent{}
	event.ID = model.ID(unitID)
	event.TypeName = "ScheduledEvents$ScheduledEvent"
	event.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		event.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		event.Documentation = doc
	}
	if mfID, ok := raw["Microflow"].(string); ok {
		event.MicroflowID = model.ID(mfID)
	}
	if enabled, ok := raw["Enabled"].(bool); ok {
		event.Enabled = enabled
	}
	if interval, ok := raw["Interval"].(int32); ok {
		event.Interval = int(interval)
	}
	if intervalType, ok := raw["IntervalType"].(string); ok {
		event.IntervalType = intervalType
	}

	return event, nil
}

// resolveContents handles MPR v2 external file references.
func (r *Reader) resolveContents(unitID string, contents []byte) ([]byte, error) {
	// For MPR v1, contents are stored directly in the database
	if r.version == MPRVersionV1 {
		return contents, nil
	}

	// For MPR v2, check if contents is a reference to an external file
	// Contents might be empty or contain just a hash
	if len(contents) > 0 {
		// Check if it's actual BSON content (starts with length prefix)
		if len(contents) >= 4 {
			return contents, nil
		}
	}

	// Look for the external file in mprcontents
	externalPath := filepath.Join(r.contentsDir, unitID)
	if _, err := os.Stat(externalPath); err == nil {
		return os.ReadFile(externalPath)
	}

	// Try with common extensions
	for _, ext := range []string{".mxunit", ".json", ""} {
		path := filepath.Join(r.contentsDir, unitID+ext)
		if data, err := os.ReadFile(path); err == nil {
			return data, nil
		}
	}

	return contents, nil
}

// WriteJSON serializes the given element to JSON.
func WriteJSON(element interface{}) ([]byte, error) {
	return json.MarshalIndent(element, "", "  ")
}
