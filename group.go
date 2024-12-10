// This file was auto-generated by Fern from our API Definition.

package zep

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/getzep/zep-go/v2/core"
)

type CreateGroupRequest struct {
	Description *string `json:"description,omitempty" url:"-"`
	// UserIDs     []string `json:"user_ids"`
	FactRatingInstruction *ApidataFactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"-"`
	GroupID               string                        `json:"group_id" url:"-"`
	Name                  *string                       `json:"name,omitempty" url:"-"`
}

type Group struct {
	CreatedAt   *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	ExternalID  *string `json:"external_id,omitempty" url:"external_id,omitempty"`
	// TODO deprecate
	ID          *int    `json:"id,omitempty" url:"id,omitempty"`
	Name        *string `json:"name,omitempty" url:"name,omitempty"`
	ProjectUUID *string `json:"project_uuid,omitempty" url:"project_uuid,omitempty"`
	UUID        *string `json:"uuid,omitempty" url:"uuid,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *Group) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *Group) UnmarshalJSON(data []byte) error {
	type unmarshaler Group
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = Group(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *Group) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}
