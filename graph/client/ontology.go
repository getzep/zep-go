package client

import (
	"context"

	"github.com/getzep/zep-go/v2"
)

func (c *Client) SetEntityTypes(ctx context.Context, entityTypes map[string]interface{}) (*zep.SuccessResponse, error) {
	var entitySchemas []*zep.EntityType
	for name, entityType := range entityTypes {
		entitySchema, err := ExtractEntitySchema(entityType, name)
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
			Name:       name,
			Properties: entityProperties,
		}
		entitySchemas = append(entitySchemas, entityType)
	}
	request := &zep.ApidataEntityTypeRequest{
		EntityTypes: entitySchemas,
	}
	return c.SetEntityTypesInternal(ctx, request)
}
