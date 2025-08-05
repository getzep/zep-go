package zep

import (
	"encoding/json"
	"fmt"
)

// GraphOntologyOptions contains options for graph ontology operations
type GraphOntologyOptions struct {
	UserID  *string
	GraphID *string
}

// GraphOntologyOption is a functional option for configuring graph ontology operations
type GraphOntologyOption func(*GraphOntologyOptions)

// ForUser allows to specify a user ID when setting entity/edge types.
func ForUser(userID string) GraphOntologyOption {
	return func(opts *GraphOntologyOptions) {
		opts.UserID = &userID
	}
}

// ForGraph allows to specify a graph ID when setting entity/edge types.
func ForGraph(graphID string) GraphOntologyOption {
	return func(opts *GraphOntologyOptions) {
		opts.GraphID = &graphID
	}
}

type BaseEntity struct{}

type EntityDefinition interface {
	// This is a marker interface - any struct embedding BaseEntity will satisfy it
	isEntityDefinition()
}

// isEntityDefinition is a marker method of EntityDefinition interface
func (BaseEntity) isEntityDefinition() {}

// UnmarshalNodeAttributes unmarshals a map[string]interface{} into a struct
// that embeds BaseEntity
func UnmarshalNodeAttributes(attributes map[string]interface{}, dest EntityDefinition) error {
	jsonData, err := json.Marshal(attributes)
	if err != nil {
		return fmt.Errorf("error marshaling node attributes: %w", err)
	}

	err = json.Unmarshal(jsonData, dest)
	if err != nil {
		return fmt.Errorf("error unmarshaling to struct: %w", err)
	}

	return nil
}

type BaseEdge struct{}

type EdgeDefinitionWithSourceTargets struct {
	EdgeModel     EdgeDefinition           `json:"edge_model"`
	SourceTargets []EntityEdgeSourceTarget `json:"source_targets,omitempty"`
}

type EdgeDefinition interface {
	// This is a marker interface - any struct embedding BaseEdge will satisfy it
	isEdgeDefinition()
}

// isEdgeDefinition is a marker method of EntityDefinition interface
func (BaseEdge) isEdgeDefinition() {}

// UnmarshalEdgeAttributes unmarshals a map[string]interface{} into a struct
// that embeds BaseEdge
func UnmarshalEdgeAttributes(attributes map[string]interface{}, dest EdgeDefinition) error {
	jsonData, err := json.Marshal(attributes)
	if err != nil {
		return fmt.Errorf("error marshaling node attributes: %w", err)
	}

	err = json.Unmarshal(jsonData, dest)
	if err != nil {
		return fmt.Errorf("error unmarshaling to struct: %w", err)
	}

	return nil
}
