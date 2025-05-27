package client

import (
	"fmt"
	"reflect"
	"strings"
)

// PropertyType defines the supported property types
type PropertyType string

const (
	Text    PropertyType = "Text"
	Int     PropertyType = "Int"
	Float   PropertyType = "Float"
	Boolean PropertyType = "Boolean"
)

// IsValid checks if the property type is valid
func (pt PropertyType) IsValid() bool {
	switch pt {
	case Text, Int, Float, Boolean:
		return true
	default:
		return false
	}
}

// Property represents a property definition for a schema
type Property struct {
	Type        PropertyType `json:"type"`
	Description string       `json:"description"`
}

// Schema represents a schema for an entity or edge type
type Schema struct {
	Name       string              `json:"name"`
	Properties map[string]Property `json:"properties"`
}

// Validate checks if the schema is valid
func (s Schema) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("schema name is required")
	}

	if len(s.Properties) == 0 {
		return fmt.Errorf("schema must have at least one property")
	}

	if len(s.Properties) > 10 {
		return fmt.Errorf("schema cannot have more than 10 properties")
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

// ExtractSchema extracts a schema from a struct using reflection
func ExtractSchema(obj interface{}, name string) (Schema, error) {
	schema := Schema{
		Name:       name,
		Properties: make(map[string]Property),
	}

	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return schema, fmt.Errorf("object must be a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the description tag
		tag := field.Tag.Get("description")
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
			propType = Int
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

// ExtractEntitySchema extracts an entity schema from a struct using reflection
// This is a wrapper around ExtractSchema for backward compatibility
func ExtractEntitySchema(entity interface{}, name string) (Schema, error) {
	return ExtractSchema(entity, name)
}

// ExtractEdgeSchema extracts an edge schema from a struct using reflection
// This is a wrapper around ExtractSchema for backward compatibility
func ExtractEdgeSchema(edge interface{}, name string) (Schema, error) {
	return ExtractSchema(edge, name)
}

// EntitySchema is an alias for Schema for backward compatibility
type EntitySchema = Schema

// EdgeSchema is an alias for Schema for backward compatibility
type EdgeSchema = Schema
