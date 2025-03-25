package client

import (
	"context"

	"github.com/getzep/zep-go/v2"
)

func (c *Client) SetEntityTypes(ctx context.Context, entityTypes map[string]zep.EntityTypeDefinition) (*zep.SuccessResponse, error) {
	var entitySchemas []*zep.EntityType
	for name, entityType := range entityTypes {
		entitySchema, err := ExtractEntitySchema(entityType.Interface, name)
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
		entityType := &zep.EntityType{
			Name:        name,
			Description: entityType.Description,
			Properties:  entityProperties,
		}
		entitySchemas = append(entitySchemas, entityType)
	}
	request := &zep.EntityTypeRequest{
		EntityTypes: entitySchemas,
	}
	return c.SetEntityTypesInternal(ctx, request)
}
