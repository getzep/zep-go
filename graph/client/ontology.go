package client

import (
	"context"
	"fmt"

	"github.com/getzep/zep-go/v4"
	"github.com/getzep/zep-go/v4/option"
)

// SetEntityTypes sets the entity and edge types for one graph, replacing any
// existing ontology on it. It takes a slice of EntityDefinition, satisfied by
// any struct embedding BaseEntity, and a slice of EdgeDefinitionWithSourceTargets,
// whose EdgeModel is satisfied by any struct embedding BaseEdge.
//
// The graph is now an explicit argument. v3 took ForUsers and ForGraphs and
// fanned out server-side; v4 has one ontology endpoint per scope (spec-3 14.4),
// so a caller targeting several graphs calls this once per graph, and a caller
// targeting the project default uses Project.SetOntology with the same
// zep.BuildOntology output.
// BuildOntology turns the struct-tag ontology definitions into the Ontology the
// v4 API accepts. It is the whole value the hand-written layer adds: the wire
// shape is a list of entity and edge types, and this derives it from Go structs
// so a caller declares an ontology once, in the type system.
//
// v3 addressed many graphs in one call through ForUsers and ForGraphs. v4 has
// no fan-out endpoint (spec-3 14.4), so the caller passes one scope: either
// Graph.SetEntityTypes with a graph UUID, or Project.SetOntology for the
// project default.
func BuildOntology(
	entities []zep.EntityDefinition,
	edges []zep.EdgeDefinitionWithSourceTargets,
) (*zep.Ontology, error) {
	entitySchemas := make([]*zep.EntityType, 0, len(entities))
	for i, entityStruct := range entities {
		metadata, found := ExtractBaseEntityMetadata(entityStruct)
		if !found {
			return nil, fmt.Errorf("entity at index %d does not have a BaseEntity with required name tag", i)
		}
		entitySchema, err := ExtractEntitySchema(entityStruct, metadata.Name)
		if err != nil {
			return nil, err
		}
		properties := make([]*zep.EntityProperty, 0, len(entitySchema.Properties))
		for name, property := range entitySchema.Properties {
			properties = append(properties, &zep.EntityProperty{
				Name:        name,
				Type:        zep.EntityPropertyType(property.Type),
				Description: property.Description,
			})
		}
		description := metadata.Description
		if description == "" {
			description = fmt.Sprintf("Entity type for %s", metadata.Name)
		}
		entitySchemas = append(entitySchemas, &zep.EntityType{
			Name:        metadata.Name,
			Description: description,
			Properties:  properties,
		})
	}

	edgeSchemas := make([]*zep.EdgeType, 0, len(edges))
	for i, edgeWithSourceTargets := range edges {
		edgeStruct := edgeWithSourceTargets.EdgeModel
		metadata, found := ExtractBaseEdgeMetadata(edgeStruct)
		if !found {
			return nil, fmt.Errorf("edge at index %d does not have a BaseEdge with required name tag", i)
		}
		edgeSchema, err := ExtractEdgeSchema(edgeStruct, metadata.Name)
		if err != nil {
			return nil, err
		}
		properties := make([]*zep.EntityProperty, 0, len(edgeSchema.Properties))
		for name, property := range edgeSchema.Properties {
			properties = append(properties, &zep.EntityProperty{
				Name:        name,
				Type:        zep.EntityPropertyType(property.Type),
				Description: property.Description,
			})
		}
		description := metadata.Description
		if description == "" {
			description = fmt.Sprintf("Entity type for %s", metadata.Name)
		}
		var sourceTargets []*zep.EdgeSourceTarget
		for _, sourceTarget := range edgeWithSourceTargets.SourceTargets {
			sourceTargets = append(sourceTargets, &zep.EdgeSourceTarget{
				SourceEntityType: sourceTarget.SourceEntityType,
				TargetEntityType: sourceTarget.TargetEntityType,
			})
		}
		edgeSchemas = append(edgeSchemas, &zep.EdgeType{
			Name:          metadata.Name,
			Description:   description,
			Properties:    properties,
			SourceTargets: sourceTargets,
		})
	}

	return &zep.Ontology{EntityTypes: entitySchemas, EdgeTypes: edgeSchemas}, nil
}

func (c *Client) SetEntityTypes(
	ctx context.Context,
	graphUUID string,
	entities []zep.EntityDefinition,
	edges []zep.EdgeDefinitionWithSourceTargets,
	opts ...option.IdempotentRequestOption,
) (*zep.Ontology, error) {
	ontology, err := BuildOntology(entities, edges)
	if err != nil {
		return nil, err
	}
	return c.SetOntology(ctx, graphUUID, ontology, opts...)
}
