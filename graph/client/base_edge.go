package client

import (
	"reflect"

	"github.com/getzep/zep-go/v3"
)

// BaseEdgeMetadata contains information extracted from BaseEdge struct tags
type BaseEdgeMetadata struct {
	Description string
	Name        string
}

// ExtractBaseEdgeMetadata tries to find an embedded BaseEdge by reflection
// and extract metadata from direct struct tags
func ExtractBaseEdgeMetadata(entity zep.EdgeDefinition) (BaseEdgeMetadata, bool) {
	v := reflect.ValueOf(entity)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return BaseEdgeMetadata{}, false
	}

	t := v.Type()

	// Look for BaseEdge field and its struct tags
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Check if this is the BaseEdge field
		if field.Type.Name() == "BaseEdge" && field.Anonymous {
			metadata := BaseEdgeMetadata{}
			foundName := false

			// Extract name tag directly
			nameTag := field.Tag.Get("name")
			if nameTag != "" {
				metadata.Name = nameTag
				foundName = true
			}

			// Extract description tag directly
			descTag := field.Tag.Get("description")
			if descTag != "" {
				metadata.Description = descTag
			}

			// Name is required, description is optional
			if foundName {
				return metadata, true
			}
		}
	}

	return BaseEdgeMetadata{}, false
}
