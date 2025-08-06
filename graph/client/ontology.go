package client

import (
	"context"
	"fmt"

	"github.com/getzep/zep-go/v3"
)

// SetOntology sets entity end edge types for the target, replacing any existing entity/edge types set for the target. If no user/graph target is set, it will default to the project target.
// It takes a slice of EntityDefinition, which is satisfied by any struct that embeds BaseEntity, and a slice of EdgeDefinition, which is satisfied by any struct that embeds BaseEdge
func (c *Client) SetOntology(
	ctx context.Context,
	entities []zep.EntityDefinition,
	edges []zep.EdgeDefinitionWithSourceTargets,
	options ...zep.GraphOntologyOption,
) (*zep.SuccessResponse, error) {
	return c.SetEntityTypes(ctx, entities, edges, options...)
}

// SetEntityTypes sets entity end edge types for the target, replacing any existing entity/edge types set for the target. If no user/graph target is set, it will default to the project target.
// It takes a slice of EntityDefinition, which is satisfied by any struct that embeds BaseEntity, and a slice of EdgeDefinition, which is satisfied by any struct that embeds BaseEdge
func (c *Client) SetEntityTypes(
	ctx context.Context,
	entities []zep.EntityDefinition,
	edges []zep.EdgeDefinitionWithSourceTargets,
	options ...zep.GraphOntologyOption,
) (*zep.SuccessResponse, error) {
	opts := &zep.GraphOntologyOptions{}
	for _, option := range options {
		option(opts)
	}

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
		GraphIDs:    opts.GraphIDs,
		UserIDs:     opts.UserIDs,
	}
	return c.SetEntityTypesInternal(ctx, request)
}
