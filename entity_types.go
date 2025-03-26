package zep

import (
	"encoding/json"
	"fmt"
)

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
