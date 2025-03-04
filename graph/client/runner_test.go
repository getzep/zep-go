package client

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Person defines a person entity using struct tags
type Person struct {
	FirstName string  `entity:"The person's first name" json:"first_name,omitempty"`
	LastName  string  `entity:"The person's last name" json:"last_name,omitempty"`
	Age       int     `entity:"The person's age" json:"age,omitempty"`
	Height    float64 `entity:"The person's height in meters" json:"height,omitempty"`
	IsActive  bool    `entity:"Whether the person is active" json:"is_active,omitempty"`
}

func TestEntitySchema(t *testing.T) {
	t.Run("ExtractEntitySchema", func(t *testing.T) {
		// Define a struct with entity tags
		person := Person{}

		// Extract schema from the struct
		schema, err := ExtractEntitySchema(person, "Person")
		assert.NoError(t, err)
		assert.Equal(t, "Person", schema.Name)
		assert.Len(t, schema.Properties, 5)

		// Check property types
		assert.Equal(t, Text, schema.Properties["first_name"].Type)
		assert.Equal(t, Number, schema.Properties["age"].Type)
		assert.Equal(t, Float, schema.Properties["height"].Type)
		assert.Equal(t, Boolean, schema.Properties["is_active"].Type)

		// Check property descriptions
		assert.Equal(t, "The person's first name", schema.Properties["first_name"].Description)
		assert.Equal(t, "The person's last name", schema.Properties["last_name"].Description)
		assert.Equal(t, "The person's age", schema.Properties["age"].Description)
		assert.Equal(t, "The person's height in meters", schema.Properties["height"].Description)
		assert.Equal(t, "Whether the person is active", schema.Properties["is_active"].Description)
	})

	t.Run("EntitySchemaToJSON", func(t *testing.T) {
		// Define a schema
		schema := EntitySchema{
			Name: "Person",
			Properties: map[string]Property{
				"first_name": {Type: Text, Description: "The person's first name"},
				"last_name":  {Type: Text, Description: "The person's last name"},
				"age":        {Type: Number, Description: "The person's age"},
				"height":     {Type: Float, Description: "The person's height in meters"},
				"is_active":  {Type: Boolean, Description: "Whether the person is active"},
			},
		}

		// Convert schema to JSON
		jsonStr, err := EntitySchemaToJSON(schema)
		assert.NoError(t, err)

		// Verify JSON contains expected properties
		var jsonMap map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &jsonMap)
		assert.NoError(t, err)

		assert.Equal(t, "Person", jsonMap["name"])
		properties := jsonMap["properties"].([]interface{})
		assert.Len(t, properties, 5)

		// Check that all properties are in the JSON
		propertyNames := make(map[string]bool)
		for _, prop := range properties {
			propMap := prop.(map[string]interface{})
			propertyNames[propMap["name"].(string)] = true

			// Verify each property has type and description
			assert.Contains(t, propMap, "type")
			assert.Contains(t, propMap, "description")
		}

		assert.Contains(t, propertyNames, "first_name")
		assert.Contains(t, propertyNames, "last_name")
		assert.Contains(t, propertyNames, "age")
		assert.Contains(t, propertyNames, "height")
		assert.Contains(t, propertyNames, "is_active")
	})

	t.Run("JSONToEntitySchema", func(t *testing.T) {
		// Define a JSON schema
		jsonStr := `{
			"name": "Person",
			"properties": [
				{"name": "first_name", "type": "Text", "description": "The person's first name"},
				{"name": "last_name", "type": "Text", "description": "The person's last name"},
				{"name": "age", "type": "Number", "description": "The person's age"},
				{"name": "height", "type": "Float", "description": "The person's height in meters"},
				{"name": "is_active", "type": "Boolean", "description": "Whether the person is active"}
			]
		}`

		// Convert JSON to schema
		schema, err := JSONToEntitySchema(jsonStr)
		assert.NoError(t, err)

		// Verify schema properties
		assert.Equal(t, "Person", schema.Name)
		assert.Len(t, schema.Properties, 5)

		assert.Equal(t, Text, schema.Properties["first_name"].Type)
		assert.Equal(t, Text, schema.Properties["last_name"].Type)
		assert.Equal(t, Number, schema.Properties["age"].Type)
		assert.Equal(t, Float, schema.Properties["height"].Type)
		assert.Equal(t, Boolean, schema.Properties["is_active"].Type)

		assert.Equal(t, "The person's first name", schema.Properties["first_name"].Description)
		assert.Equal(t, "The person's last name", schema.Properties["last_name"].Description)
		assert.Equal(t, "The person's age", schema.Properties["age"].Description)
		assert.Equal(t, "The person's height in meters", schema.Properties["height"].Description)
		assert.Equal(t, "Whether the person is active", schema.Properties["is_active"].Description)
	})

	t.Run("RoundTrip", func(t *testing.T) {
		// Define a struct with entity tags
		person := Person{}

		// Extract schema from the struct
		schema, err := ExtractEntitySchema(person, "Person")
		assert.NoError(t, err)

		// Convert schema to JSON
		jsonStr, err := EntitySchemaToJSON(schema)
		assert.NoError(t, err)

		// Convert JSON back to schema
		recreatedSchema, err := JSONToEntitySchema(jsonStr)
		assert.NoError(t, err)

		// Verify recreated schema matches original
		assert.Equal(t, schema.Name, recreatedSchema.Name)
		assert.Equal(t, len(schema.Properties), len(recreatedSchema.Properties))

		for name, prop := range schema.Properties {
			recreatedProp, ok := recreatedSchema.Properties[name]
			assert.True(t, ok)
			assert.Equal(t, prop.Type, recreatedProp.Type)
			assert.Equal(t, prop.Description, recreatedProp.Description)
		}
	})

	t.Run("CreateEntityFromSchema", func(t *testing.T) {
		// Define a schema
		schema := EntitySchema{
			Name: "Person",
			Properties: map[string]Property{
				"first_name": {Type: Text, Description: "The person's first name"},
				"last_name":  {Type: Text, Description: "The person's last name"},
				"age":        {Type: Number, Description: "The person's age"},
				"height":     {Type: Float, Description: "The person's height in meters"},
				"is_active":  {Type: Boolean, Description: "Whether the person is active"},
			},
		}

		// Define data
		data := map[string]interface{}{
			"first_name": "John",
			"last_name":  "Doe",
			"age":        30,
			"height":     1.75,
			"is_active":  true,
		}

		// Create entity from schema and data
		var person Person
		err := CreateEntityFromSchema(schema, data, &person)
		assert.NoError(t, err)

		// Verify entity values
		assert.Equal(t, "John", person.FirstName)
		assert.Equal(t, "Doe", person.LastName)
		assert.Equal(t, 30, person.Age)
		assert.Equal(t, 1.75, person.Height)
		assert.Equal(t, true, person.IsActive)
	})

	t.Run("ValidationFailure", func(t *testing.T) {
		// Test invalid schema (no name)
		schema := EntitySchema{
			Properties: map[string]Property{
				"first_name": {Type: Text, Description: "The person's first name"},
			},
		}

		_, err := EntitySchemaToJSON(schema)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")

		// Test invalid schema (no properties)
		schema = EntitySchema{
			Name:       "Person",
			Properties: map[string]Property{},
		}

		_, err = EntitySchemaToJSON(schema)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must have at least one property")

		// Test invalid schema (too many properties)
		schema = EntitySchema{
			Name: "Person",
			Properties: map[string]Property{
				"prop1":  {Type: Text, Description: "Property 1"},
				"prop2":  {Type: Text, Description: "Property 2"},
				"prop3":  {Type: Text, Description: "Property 3"},
				"prop4":  {Type: Text, Description: "Property 4"},
				"prop5":  {Type: Text, Description: "Property 5"},
				"prop6":  {Type: Text, Description: "Property 6"},
				"prop7":  {Type: Text, Description: "Property 7"},
				"prop8":  {Type: Text, Description: "Property 8"},
				"prop9":  {Type: Text, Description: "Property 9"},
				"prop10": {Type: Text, Description: "Property 10"},
				"prop11": {Type: Text, Description: "Property 11"},
			},
		}

		_, err = EntitySchemaToJSON(schema)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot have more than 10 properties")

		// Test invalid schema (invalid property type)
		schema = EntitySchema{
			Name: "Person",
			Properties: map[string]Property{
				"first_name": {Type: "InvalidType", Description: "The person's first name"},
			},
		}

		_, err = EntitySchemaToJSON(schema)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid property type")

		// Test invalid schema (missing description)
		schema = EntitySchema{
			Name: "Person",
			Properties: map[string]Property{
				"first_name": {Type: Text, Description: ""},
			},
		}

		_, err = EntitySchemaToJSON(schema)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "description is required")
	})
}
