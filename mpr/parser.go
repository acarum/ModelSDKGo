package mpr

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/anthropics/modelsdk-go/domainmodel"
	"github.com/anthropics/modelsdk-go/microflows"
	"github.com/anthropics/modelsdk-go/model"
	"github.com/anthropics/modelsdk-go/pages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// extractBsonID extracts an ID string from various BSON ID representations.
// Mendix stores IDs as Binary with Subtype/Data or as primitive.Binary.
func extractBsonID(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return blobToUUID(val)
	case primitive.Binary:
		return blobToUUID(val.Data)
	case map[string]interface{}:
		// Binary UUID stored as {Subtype: 0, Data: "base64..."}
		if data, ok := val["Data"].(string); ok {
			decoded, err := base64.StdEncoding.DecodeString(data)
			if err == nil {
				return blobToUUID(decoded)
			}
		}
		// Also try $ID field
		if id, ok := val["$ID"]; ok {
			return extractBsonID(id)
		}
	}

	return ""
}

// extractInt extracts an integer from various BSON number types.
func extractInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int32:
		return int(val)
	case int64:
		return int(val)
	case int:
		return val
	case float64:
		return int(val)
	}
	return 0
}

// extractString extracts a string from various BSON representations.
func extractString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// extractBool extracts a boolean from BSON, with default value.
func extractBool(v interface{}, defaultVal bool) bool {
	if v == nil {
		return defaultVal
	}
	if b, ok := v.(bool); ok {
		return b
	}
	return defaultVal
}

// toStringMap converts a value to map[string]interface{}, handling both
// map[string]interface{} and primitive.D (mongo-driver BSON embedded doc).
func toStringMap(v interface{}) map[string]interface{} {
	if m, ok := v.(map[string]interface{}); ok {
		return m
	}
	if d, ok := v.(primitive.D); ok {
		m := make(map[string]interface{}, len(d))
		for _, e := range d {
			m[e.Key] = e.Value
		}
		return m
	}
	return nil
}

// extractBsonArray extracts items from a Mendix BSON array.
func extractBsonArray(v interface{}) []interface{} {
	if v == nil {
		return nil
	}

	arr, ok := v.(primitive.A)
	if !ok {
		// Try regular slice
		if slice, ok := v.([]interface{}); ok {
			// Check if first element is the array type indicator
			if len(slice) > 0 {
				if typeIndicator, ok := slice[0].(int32); ok && typeIndicator == 3 {
					// Skip the type indicator
					return slice[1:]
				}
			}
			return slice
		}
		return nil
	}

	// primitive.A is []interface{} underneath
	slice := []interface{}(arr)

	// Check if first element is the array type indicator (3)
	if len(slice) > 0 {
		if typeIndicator, ok := slice[0].(int32); ok && typeIndicator == 3 {
			// Skip the type indicator
			return slice[1:]
		}
	}

	return slice
}

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

	// Parse entities - use extractBsonArray to handle Mendix array format
	entities := extractBsonArray(raw["Entities"])
	for _, e := range entities {
		if entityMap, ok := e.(map[string]interface{}); ok {
			entity := parseEntity(entityMap)
			dm.Entities = append(dm.Entities, entity)
		}
	}

	// Parse associations
	associations := extractBsonArray(raw["Associations"])
	for _, a := range associations {
		if assocMap, ok := a.(map[string]interface{}); ok {
			assoc := parseAssociation(assocMap)
			dm.Associations = append(dm.Associations, assoc)
		}
	}

	// Parse annotations
	annotations := extractBsonArray(raw["Annotations"])
	for _, a := range annotations {
		if annotMap, ok := a.(map[string]interface{}); ok {
			annot := parseAnnotation(annotMap)
			dm.Annotations = append(dm.Annotations, annot)
		}
	}

	return dm, nil
}

func parseEntity(raw map[string]interface{}) *domainmodel.Entity {
	entity := &domainmodel.Entity{}

	// Use extractBsonID to handle various ID formats (string, binary, base64)
	entity.ID = model.ID(extractBsonID(raw["$ID"]))
	if typeName, ok := raw["$Type"].(string); ok {
		entity.TypeName = typeName
	}
	if name, ok := raw["Name"].(string); ok {
		entity.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		entity.Documentation = doc
	}

	// Parse location - handle both int32 and float64
	if loc, ok := raw["Location"].(map[string]interface{}); ok {
		entity.Location.X = extractInt(loc["x"])
		entity.Location.Y = extractInt(loc["y"])
	}

	// Parse persistable - default to true if not specified
	entity.Persistable = true
	if persistable, ok := raw["Persistable"].(bool); ok {
		entity.Persistable = persistable
	}

	// Parse generalization (parent entity)
	if gen := raw["Generalization"]; gen != nil {
		if genMap, ok := gen.(map[string]interface{}); ok {
			if genID := extractBsonID(genMap["$ID"]); genID != "" {
				entity.GeneralizationID = model.ID(genID)
			}
			// Handle direct entity reference
			if entityRef := extractBsonID(genMap["Generalization"]); entityRef != "" {
				entity.GeneralizationID = model.ID(entityRef)
			}
		}
	}

	// Parse attributes using extractBsonArray
	attrs := extractBsonArray(raw["Attributes"])
	for _, a := range attrs {
		if attrMap, ok := a.(map[string]interface{}); ok {
			attr := parseAttribute(attrMap)
			entity.Attributes = append(entity.Attributes, attr)
		}
	}

	// Parse indexes
	indexes := extractBsonArray(raw["Indexes"])
	for _, i := range indexes {
		if indexMap, ok := i.(map[string]interface{}); ok {
			index := parseIndex(indexMap)
			entity.Indexes = append(entity.Indexes, index)
		}
	}

	// Parse access rules
	rules := extractBsonArray(raw["AccessRules"])
	for _, r := range rules {
		if ruleMap, ok := r.(map[string]interface{}); ok {
			rule := parseAccessRule(ruleMap)
			entity.AccessRules = append(entity.AccessRules, rule)
		}
	}

	// Parse validation rules
	validations := extractBsonArray(raw["ValidationRules"])
	for _, v := range validations {
		if validMap, ok := v.(map[string]interface{}); ok {
			validation := parseValidationRule(validMap)
			entity.ValidationRules = append(entity.ValidationRules, validation)
		}
	}

	// Parse event handlers
	handlers := extractBsonArray(raw["EventHandlers"])
	for _, h := range handlers {
		if handlerMap, ok := h.(map[string]interface{}); ok {
			handler := parseEventHandler(handlerMap)
			entity.EventHandlers = append(entity.EventHandlers, handler)
		}
	}

	return entity
}

func parseAttribute(raw map[string]interface{}) *domainmodel.Attribute {
	attr := &domainmodel.Attribute{}

	attr.ID = model.ID(extractBsonID(raw["$ID"]))
	attr.TypeName = extractString(raw["$Type"])
	attr.Name = extractString(raw["Name"])
	attr.Documentation = extractString(raw["Documentation"])

	// Parse attribute type - Mendix uses "NewType" field
	if attrType, ok := raw["NewType"].(map[string]interface{}); ok {
		attr.Type = parseAttributeType(attrType)
	} else if attrType, ok := raw["Type"].(map[string]interface{}); ok {
		// Fallback to "Type" for older format
		attr.Type = parseAttributeType(attrType)
	}

	// Parse default value
	if val, ok := raw["Value"].(map[string]interface{}); ok {
		attr.Value = parseAttributeValue(val)
	}

	return attr
}

func parseAttributeValue(raw map[string]interface{}) *domainmodel.AttributeValue {
	typeName := extractString(raw["$Type"])
	defaultValue := extractString(raw["DefaultValue"])

	switch typeName {
	case "DomainModels$StoredValue":
		return &domainmodel.AttributeValue{
			Type:         "StoredValue",
			DefaultValue: defaultValue,
		}
	case "DomainModels$CalculatedValue":
		return &domainmodel.AttributeValue{
			Type:        "CalculatedValue",
			MicroflowID: model.ID(extractBsonID(raw["Microflow"])),
		}
	default:
		return &domainmodel.AttributeValue{
			DefaultValue: defaultValue,
		}
	}
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

	assoc.ID = model.ID(extractBsonID(raw["$ID"]))
	assoc.TypeName = extractString(raw["$Type"])
	assoc.Name = extractString(raw["Name"])
	assoc.Documentation = extractString(raw["Documentation"])
	assoc.ParentID = model.ID(extractBsonID(raw["Parent"]))
	assoc.ChildID = model.ID(extractBsonID(raw["Child"]))
	assoc.Type = domainmodel.AssociationType(extractString(raw["Type"]))
	assoc.Owner = domainmodel.AssociationOwner(extractString(raw["Owner"]))

	return assoc
}

func parseAnnotation(raw map[string]interface{}) *domainmodel.Annotation {
	annot := &domainmodel.Annotation{}

	annot.ID = model.ID(extractBsonID(raw["$ID"]))
	annot.TypeName = extractString(raw["$Type"])
	annot.Caption = extractString(raw["Caption"])

	if loc, ok := raw["Location"].(map[string]interface{}); ok {
		annot.Location.X = extractInt(loc["x"])
		annot.Location.Y = extractInt(loc["y"])
	}

	return annot
}

func parseIndex(raw map[string]interface{}) *domainmodel.Index {
	index := &domainmodel.Index{}

	index.ID = model.ID(extractBsonID(raw["$ID"]))
	index.Name = extractString(raw["Name"])

	// Parse index attributes
	attrs := extractBsonArray(raw["Attributes"])
	for _, a := range attrs {
		if attrMap, ok := a.(map[string]interface{}); ok {
			attrID := extractBsonID(attrMap["Attribute"])
			if attrID != "" {
				index.AttributeIDs = append(index.AttributeIDs, model.ID(attrID))
			}
		}
	}

	return index
}

func parseAccessRule(raw map[string]interface{}) *domainmodel.AccessRule {
	rule := &domainmodel.AccessRule{}

	rule.ID = model.ID(extractBsonID(raw["$ID"]))
	rule.AllowCreate = extractBool(raw["AllowCreate"], false)
	rule.AllowRead = extractBool(raw["AllowRead"], false)
	rule.AllowWrite = extractBool(raw["AllowWrite"], false)
	rule.AllowDelete = extractBool(raw["AllowDelete"], false)
	rule.XPathConstraint = extractString(raw["XPathConstraint"])

	// Parse module roles
	roles := extractBsonArray(raw["ModuleRoles"])
	for _, r := range roles {
		roleID := extractBsonID(r)
		if roleID != "" {
			rule.ModuleRoles = append(rule.ModuleRoles, model.ID(roleID))
		}
	}

	return rule
}

func parseValidationRule(raw map[string]interface{}) *domainmodel.ValidationRule {
	rule := &domainmodel.ValidationRule{}

	rule.ID = model.ID(extractBsonID(raw["$ID"]))
	rule.AttributeID = model.ID(extractBsonID(raw["Attribute"]))
	rule.Type = extractString(raw["$Type"])

	// Parse error message
	if errMsg, ok := raw["ErrorMessage"].(map[string]interface{}); ok {
		rule.ErrorMessage = parseText(errMsg)
	}

	return rule
}

func parseEventHandler(raw map[string]interface{}) *domainmodel.EventHandler {
	handler := &domainmodel.EventHandler{}

	handler.ID = model.ID(extractBsonID(raw["$ID"]))
	handler.Event = domainmodel.EventType(extractString(raw["Event"]))
	handler.MicroflowID = model.ID(extractBsonID(raw["Microflow"]))
	handler.RaiseErrorOnFalse = extractBool(raw["RaiseErrorOnFalse"], false)

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
	page.TypeName = "Forms$Page"
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
	if titleMap := toStringMap(raw["Title"]); titleMap != nil {
		page.Title = parseText(titleMap)
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

	// Try flat map format: {"en_US": "...", "en_GB": "..."}
	if translations, ok := raw["Translations"].(map[string]interface{}); ok {
		for lang, val := range translations {
			if str, ok := val.(string); ok {
				text.Translations[lang] = str
			}
		}
	}

	// MPR v2 format: Title is {$Type: "Texts$Text", Items: [{$Type: "Texts$Translation", LanguageCode: "en_US", Text: "..."}]}
	// So translations are under raw["Items"]
	for _, key := range []string{"Items", "Translations"} {
		if items := extractBsonArray(raw[key]); items != nil {
			for _, item := range items {
				m := toStringMap(item)
				if m == nil {
					continue
				}
				lang, _ := m["LanguageCode"].(string)
				txt, _ := m["Text"].(string)
				if lang != "" && txt != "" {
					text.Translations[lang] = txt
				}
			}
			break
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
	layout.TypeName = "Forms$Layout"
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

	// Load from mprcontents folder using the proper directory structure
	return r.loadUnitContents(unitID)
}

// parseSnippet parses snippet contents from BSON.
func (r *Reader) parseSnippet(unitID, containerID string, contents []byte) (*pages.Snippet, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	snippet := &pages.Snippet{}
	snippet.ID = model.ID(unitID)
	snippet.TypeName = "Pages$Snippet"
	snippet.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		snippet.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		snippet.Documentation = doc
	}
	if entityID := extractID(raw["Entity"]); entityID != "" {
		snippet.EntityID = model.ID(entityID)
	}

	return snippet, nil
}

// parseJavaAction parses Java action contents from BSON.
func (r *Reader) parseJavaAction(unitID, containerID string, contents []byte) (*JavaAction, error) {
	contents, err := r.resolveContents(unitID, contents)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := bson.Unmarshal(contents, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON: %w", err)
	}

	ja := &JavaAction{}
	ja.ID = model.ID(unitID)
	ja.TypeName = "JavaActions$JavaAction"
	ja.ContainerID = model.ID(containerID)

	if name, ok := raw["Name"].(string); ok {
		ja.Name = name
	}
	if doc, ok := raw["Documentation"].(string); ok {
		ja.Documentation = doc
	}

	return ja, nil
}

// extractID extracts an ID from various BSON representations.
// IDs in Mendix BSON can be strings, binary UUIDs, or nested structures.
func extractID(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return blobToUUID(val)
	case map[string]interface{}:
		// Could be a reference structure with $ID
		if id, ok := val["$ID"].(string); ok {
			return id
		}
		if id, ok := val["$ID"].([]byte); ok {
			return blobToUUID(id)
		}
	}

	return ""
}

// WriteJSON serializes the given element to JSON.
func WriteJSON(element interface{}) ([]byte, error) {
	return json.MarshalIndent(element, "", "  ")
}
