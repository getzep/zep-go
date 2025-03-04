package client

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// PropertyType defines the supported property types
type PropertyType string

const (
	Text    PropertyType = "Text"
	Number  PropertyType = "Number"
	Float   PropertyType = "Float"
	Boolean PropertyType = "Boolean"
)

// IsValid checks if the property type is valid
func (pt PropertyType) IsValid() bool {
	switch pt {
	case Text, Number, Float, Boolean:
		return true
	default:
		return false
	}
}

// Property represents a property definition for an entity type
type Property struct {
	Type        PropertyType `json:"type"`
	Description string       `json:"description"`
}

// EntitySchema represents a schema for an entity type
type EntitySchema struct {
	Name       string              `json:"name"`
	Properties map[string]Property `json:"properties"`
}

// Validate checks if the schema is valid
func (s EntitySchema) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("entity schema name is required")
	}

	if len(s.Properties) == 0 {
		return fmt.Errorf("entity schema must have at least one property")
	}

	if len(s.Properties) > 10 {
		return fmt.Errorf("entity schema cannot have more than 10 properties")
	}

	for name, prop := range s.Properties {
		if name == "" {
			return fmt.Errorf("property name cannot be empty")
		}

		if !prop.Type.IsValid() {
			return fmt.Errorf("invalid property type for %s: %s", name, prop.Type)
		}

		if prop.Description == "" {
			return fmt.Errorf("property description is required for %s", name)
		}
	}

	return nil
}

// EntitySchemaToJSON converts an entity schema to JSON
func EntitySchemaToJSON(schema EntitySchema) (string, error) {
	if err := schema.Validate(); err != nil {
		return "", err
	}

	// Convert the schema to the format expected by the server
	type property struct {
		Name        string       `json:"name"`
		Type        PropertyType `json:"type"`
		Description string       `json:"description"`
	}

	type entityType struct {
		Name       string     `json:"name"`
		Properties []property `json:"properties"`
	}

	et := entityType{
		Name:       schema.Name,
		Properties: make([]property, 0, len(schema.Properties)),
	}

	for name, prop := range schema.Properties {
		et.Properties = append(et.Properties, property{
			Name:        name,
			Type:        prop.Type,
			Description: prop.Description,
		})
	}

	data, err := json.Marshal(et)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// JSONToEntitySchema converts JSON to an entity schema
func JSONToEntitySchema(jsonStr string) (EntitySchema, error) {
	// Parse the JSON
	type property struct {
		Name        string       `json:"name"`
		Type        PropertyType `json:"type"`
		Description string       `json:"description"`
	}

	type entityType struct {
		Name       string     `json:"name"`
		Properties []property `json:"properties"`
	}

	var et entityType
	if err := json.Unmarshal([]byte(jsonStr), &et); err != nil {
		return EntitySchema{}, err
	}

	// Convert to EntitySchema
	schema := EntitySchema{
		Name:       et.Name,
		Properties: make(map[string]Property),
	}

	for _, prop := range et.Properties {
		schema.Properties[prop.Name] = Property{
			Type:        prop.Type,
			Description: prop.Description,
		}
	}

	// Validate the schema
	if err := schema.Validate(); err != nil {
		return EntitySchema{}, err
	}

	return schema, nil
}

// ExtractEntitySchema extracts an entity schema from a struct using reflection
func ExtractEntitySchema(entity interface{}, name string) (EntitySchema, error) {
	schema := EntitySchema{
		Name:       name,
		Properties: make(map[string]Property),
	}

	t := reflect.TypeOf(entity)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return schema, fmt.Errorf("entity must be a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the entity tag
		tag := field.Tag.Get("entity")
		if tag == "" || tag == "-" {
			continue
		}

		// Get the json tag for the field name
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		if jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" {
				fieldName = parts[0]
			}
		}

		// Parse the tag
		var propType PropertyType
		var description string = tag

		switch field.Type.Kind() {
		case reflect.String:
			propType = Text
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			propType = Number
		case reflect.Float32, reflect.Float64:
			propType = Float
		case reflect.Bool:
			propType = Boolean
		default:
			// Skip unsupported types
			continue
		}

		// Add the property to the schema
		schema.Properties[fieldName] = Property{
			Type:        propType,
			Description: description,
		}
	}

	// Validate the schema
	if err := schema.Validate(); err != nil {
		return schema, err
	}

	return schema, nil
}

// CreateEntityFromSchema creates a new entity instance from a schema and data
func CreateEntityFromSchema(schema EntitySchema, data map[string]interface{}, entityPtr interface{}) error {
	v := reflect.ValueOf(entityPtr)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("entityPtr must be a non-nil pointer")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("entityPtr must point to a struct")
	}

	t := v.Type()

	// Map of JSON field names to struct fields
	fieldMap := make(map[string]reflect.StructField)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the json tag
		jsonTag := field.Tag.Get("json")
		fieldName := field.Name
		if jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" {
				fieldName = parts[0]
			}
		}

		fieldMap[fieldName] = field
	}

	// Set values from data
	for key, value := range data {
		field, ok := fieldMap[key]
		if !ok {
			continue
		}

		fieldValue := v.FieldByName(field.Name)
		if !fieldValue.CanSet() {
			continue
		}

		// Convert value to the appropriate type
		switch field.Type.Kind() {
		case reflect.String:
			if str, ok := value.(string); ok {
				fieldValue.SetString(str)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var intVal int64
			switch v := value.(type) {
			case int:
				intVal = int64(v)
			case int64:
				intVal = v
			case float64:
				intVal = int64(v)
			}
			fieldValue.SetInt(intVal)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var uintVal uint64
			switch v := value.(type) {
			case uint:
				uintVal = uint64(v)
			case uint64:
				uintVal = v
			case int:
				uintVal = uint64(v)
			case float64:
				uintVal = uint64(v)
			}
			fieldValue.SetUint(uintVal)
		case reflect.Float32, reflect.Float64:
			var floatVal float64
			switch v := value.(type) {
			case float64:
				floatVal = v
			case int:
				floatVal = float64(v)
			}
			fieldValue.SetFloat(floatVal)
		case reflect.Bool:
			if b, ok := value.(bool); ok {
				fieldValue.SetBool(b)
			}
		}
	}

	return nil
}
