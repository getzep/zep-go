package client

import (
	"reflect"

	"github.com/getzep/zep-go/v2"
)

// HasBaseEntity is an interface that allows checking if a struct has an embedded BaseEntity
type HasBaseEntity interface {
	GetDescription() string
}

// BaseEntityMetadata contains information extracted from BaseEntity struct tags
type BaseEntityMetadata struct {
	Description string
	Name        string
}

// ExtractBaseEntityMetadata tries to find an embedded BaseEntity by reflection
// and extract metadata from direct struct tags
func ExtractBaseEntityMetadata(entity zep.EntityDefinition) (BaseEntityMetadata, bool) {
	v := reflect.ValueOf(entity)
	
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	if v.Kind() != reflect.Struct {
		return BaseEntityMetadata{}, false
	}
	
	t := v.Type()
	
	// Look for BaseEntity field and its struct tags
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		
		// Check if this is the BaseEntity field
		if field.Type.Name() == "BaseEntity" && field.Anonymous {
			metadata := BaseEntityMetadata{}
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
	
	return BaseEntityMetadata{}, false
} 