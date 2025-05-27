package client

import (
	"context"
	"fmt"

	"github.com/getzep/zep-go/v2"
)

// SetEntityTypes sets entity types for the project using structs that embed BaseEntity with struct tags.
// It takes a slice of EntityDefinition, which is satisfied by any struct that embeds BaseEntity.
func (c *Client) SetEntityTypes(ctx context.Context, entities []zep.EntityDefinition, edges []zep.EdgeDefinitionWithSourceTargets) (*zep.SuccessResponse, error) {
	var entitySchemas []*zep.EntityType
	var edgeSchemas []*zep.EdgeType

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

	for i, edgeWithSourceTargets := range edges {
		edgeStruct := edgeWithSourceTargets.EdgeModel
		// Try to extract metadata from embedded BaseEntity struct tags
		metadata, found := ExtractBaseEdgeMetadata(edgeStruct)
		if !found {
			return nil, fmt.Errorf("entity at index %d does not have a BaseEdge with required name tag", i)
		}

		// Name is always from the struct tag
		edgeName := metadata.Name

		// Extract entity schema as usual
		edgeSchema, err := ExtractEdgeSchema(edgeStruct, edgeName)
		if err != nil {
			return nil, err
		}

		entityProperties := make([]*zep.EntityProperty, 0, len(edgeSchema.Properties))
		for name, property := range edgeSchema.Properties {
			entityProperties = append(entityProperties, &zep.EntityProperty{
				Name:        name,
				Type:        zep.EntityPropertyType(property.Type),
				Description: property.Description,
			})
		}

		// If description is not provided in struct tag, use a default or empty string
		description := metadata.Description
		if description == "" {
			description = fmt.Sprintf("Entity type for %s", edgeName)
		}
		var sourceTargets []*zep.EntityEdgeSourceTarget
		if edgeWithSourceTargets.SourceTargets != nil {
			for _, sourceTarget := range edgeWithSourceTargets.SourceTargets {
				sourceTargets = append(sourceTargets, &zep.EntityEdgeSourceTarget{
					Source: sourceTarget.Source,
					Target: sourceTarget.Target,
				})
			}
		}
		edgeType := &zep.EdgeType{
			Name:          edgeName,
			Description:   description,
			Properties:    entityProperties,
			SourceTargets: sourceTargets,
		}

		edgeSchemas = append(edgeSchemas, edgeType)
	}

	request := &zep.EntityTypeRequest{
		EntityTypes: entitySchemas,
		EdgeTypes:   edgeSchemas,
	}
	return c.SetEntityTypesInternal(ctx, request)
}
