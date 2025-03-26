package client

import (
	"context"
	"fmt"

	"github.com/getzep/zep-go/v2"
)

// SetEntityTypes sets entity types for the project using structs that embed BaseEntity with struct tags.
// It takes a slice of EntityDefinition, which is satisfied by any struct that embeds BaseEntity.
func (c *Client) SetEntityTypes(ctx context.Context, entities []zep.EntityDefinition) (*zep.SuccessResponse, error) {
	var entitySchemas []*zep.EntityType

	for i, entityStruct := range entities {
		// Try to extract metadata from embedded BaseEntity struct tags
		metadata, found := ExtractBaseEntityMetadata(entityStruct)
		if !found {
			return nil, fmt.Errorf("entity at index %d does not have a BaseEntity with required name tag", i)
		}
		
		// Name is always from the struct tag
		entityName := metadata.Name
		
		// Extract entity schema as usual
		entitySchema, err := ExtractEntitySchema(entityStruct, entityName)
		if err != nil {
			return nil, err
		}
		
		entityProperties := make([]*zep.EntityProperty, 0, len(entitySchema.Properties))
		for name, property := range entitySchema.Properties {
			entityProperties = append(entityProperties, &zep.EntityProperty{
				Name:        name,
				Type:        zep.EntityPropertyType(property.Type),
				Description: property.Description,
			})
		}
		
		// If description is not provided in struct tag, use a default or empty string
		description := metadata.Description
		if description == "" {
			description = fmt.Sprintf("Entity type for %s", entityName)
		}
		
		entityType := &zep.EntityType{
			Name:        entityName,
			Description: description,
			Properties:  entityProperties,
		}
		
		entitySchemas = append(entitySchemas, entityType)
	}

	request := &zep.EntityTypeRequest{
		EntityTypes: entitySchemas,
	}
	return c.SetEntityTypesInternal(ctx, request)
}
