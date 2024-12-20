// This file was auto-generated by Fern from our API Definition.

package zep

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/getzep/zep-go/v2/core"
)

type APIError struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *APIError) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *APIError) UnmarshalJSON(data []byte) error {
	type unmarshaler APIError
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = APIError(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *APIError) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AddedFact = interface{}

type ApidataFactRatingExamples struct {
	High   *string `json:"high,omitempty" url:"high,omitempty"`
	Low    *string `json:"low,omitempty" url:"low,omitempty"`
	Medium *string `json:"medium,omitempty" url:"medium,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *ApidataFactRatingExamples) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApidataFactRatingExamples) UnmarshalJSON(data []byte) error {
	type unmarshaler ApidataFactRatingExamples
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApidataFactRatingExamples(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApidataFactRatingExamples) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ApidataFactRatingInstruction struct {
	// Examples is a list of examples that demonstrate how facts might be rated based on your instruction. You should provide
	// an example of a highly rated example, a low rated example, and a medium (or in between example). For example, if you are rating
	// based on relevance to a trip planning application, your examples might be:
	// High: "Joe's dream vacation is Bali"
	// Medium: "Joe has a fear of flying",
	// Low: "Joe's favorite food is Japanese",
	Examples *ApidataFactRatingExamples `json:"examples,omitempty" url:"examples,omitempty"`
	// A string describing how to rate facts as they apply to your application. A trip planning application may
	// use something like "relevancy to planning a trip, the user's preferences when traveling,
	// or the user's travel history."
	Instruction *string `json:"instruction,omitempty" url:"instruction,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *ApidataFactRatingInstruction) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApidataFactRatingInstruction) UnmarshalJSON(data []byte) error {
	type unmarshaler ApidataFactRatingInstruction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApidataFactRatingInstruction(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApidataFactRatingInstruction) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ClassifySessionResponse = interface{}

type CommunityNode = interface{}

type DocumentCollectionResponse = interface{}

type DocumentResponse = interface{}

type DocumentSearchResult = interface{}

type DocumentSearchResultPage = interface{}

type EntityEdge struct {
	// Creation time of the edge
	CreatedAt string `json:"created_at" url:"created_at"`
	// List of episode ids that reference these entity edges
	Episodes []string `json:"episodes,omitempty" url:"episodes,omitempty"`
	// Datetime of when the node was invalidated
	ExpiredAt *string `json:"expired_at,omitempty" url:"expired_at,omitempty"`
	// Fact representing the edge and nodes that it connects
	Fact string `json:"fact" url:"fact"`
	// Datetime of when the fact stopped being true
	InvalidAt *string `json:"invalid_at,omitempty" url:"invalid_at,omitempty"`
	// Name of the edge, relation name
	Name string `json:"name" url:"name"`
	// UUID of the source node
	SourceNodeUUID string `json:"source_node_uuid" url:"source_node_uuid"`
	// UUID of the target node
	TargetNodeUUID string `json:"target_node_uuid" url:"target_node_uuid"`
	// UUID of the edge
	UUID string `json:"uuid" url:"uuid"`
	// Datetime of when the fact became true
	ValidAt *string `json:"valid_at,omitempty" url:"valid_at,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EntityEdge) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EntityEdge) UnmarshalJSON(data []byte) error {
	type unmarshaler EntityEdge
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntityEdge(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityEdge) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EntityNode struct {
	// Creation time of the node
	CreatedAt string `json:"created_at" url:"created_at"`
	// Labels associated with the node
	Labels []string `json:"labels,omitempty" url:"labels,omitempty"`
	// Name of the node
	Name string `json:"name" url:"name"`
	// Regional summary of surrounding edges
	Summary string `json:"summary" url:"summary"`
	// UUID of the node
	UUID string `json:"uuid" url:"uuid"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EntityNode) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EntityNode) UnmarshalJSON(data []byte) error {
	type unmarshaler EntityNode
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EntityNode(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityNode) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type Episode struct {
	Content           string         `json:"content" url:"content"`
	CreatedAt         string         `json:"created_at" url:"created_at"`
	Name              *string        `json:"name,omitempty" url:"name,omitempty"`
	Source            *GraphDataType `json:"source,omitempty" url:"source,omitempty"`
	SourceDescription *string        `json:"source_description,omitempty" url:"source_description,omitempty"`
	UUID              string         `json:"uuid" url:"uuid"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *Episode) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *Episode) UnmarshalJSON(data []byte) error {
	type unmarshaler Episode
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = Episode(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *Episode) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EpisodeResponse struct {
	Episodes []*Episode `json:"episodes,omitempty" url:"episodes,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (e *EpisodeResponse) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EpisodeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EpisodeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EpisodeResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties

	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EpisodeResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EpisodeType = interface{}

type Fact struct {
	Content   string  `json:"content" url:"content"`
	CreatedAt string  `json:"created_at" url:"created_at"`
	ExpiredAt *string `json:"expired_at,omitempty" url:"expired_at,omitempty"`
	// Deprecated. This field will be removed in the future, please use `content` instead.
	Fact           string   `json:"fact" url:"fact"`
	InvalidAt      *string  `json:"invalid_at,omitempty" url:"invalid_at,omitempty"`
	Name           *string  `json:"name,omitempty" url:"name,omitempty"`
	Rating         *float64 `json:"rating,omitempty" url:"rating,omitempty"`
	SourceNodeName *string  `json:"source_node_name,omitempty" url:"source_node_name,omitempty"`
	TargetNodeName *string  `json:"target_node_name,omitempty" url:"target_node_name,omitempty"`
	UUID           string   `json:"uuid" url:"uuid"`
	ValidAt        *string  `json:"valid_at,omitempty" url:"valid_at,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (f *Fact) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *Fact) UnmarshalJSON(data []byte) error {
	type unmarshaler Fact
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = Fact(value)

	extraProperties, err := core.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties

	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *Fact) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FactRatingExamples struct {
	High   *string `json:"high,omitempty" url:"high,omitempty"`
	Low    *string `json:"low,omitempty" url:"low,omitempty"`
	Medium *string `json:"medium,omitempty" url:"medium,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (f *FactRatingExamples) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FactRatingExamples) UnmarshalJSON(data []byte) error {
	type unmarshaler FactRatingExamples
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FactRatingExamples(value)

	extraProperties, err := core.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties

	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FactRatingExamples) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FactRatingInstruction struct {
	// Examples is a list of examples that demonstrate how facts might be rated based on your instruction. You should provide
	// an example of a highly rated example, a low rated example, and a medium (or in between example). For example, if you are rating
	// based on relevance to a trip planning application, your examples might be:
	// High: "Joe's dream vacation is Bali"
	// Medium: "Joe has a fear of flying",
	// Low: "Joe's favorite food is Japanese",
	Examples *FactRatingExamples `json:"examples,omitempty" url:"examples,omitempty"`
	// A string describing how to rate facts as they apply to your application. A trip planning application may
	// use something like "relevancy to planning a trip, the user's preferences when traveling,
	// or the user's travel history."
	Instruction *string `json:"instruction,omitempty" url:"instruction,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (f *FactRatingInstruction) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FactRatingInstruction) UnmarshalJSON(data []byte) error {
	type unmarshaler FactRatingInstruction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FactRatingInstruction(value)

	extraProperties, err := core.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties

	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FactRatingInstruction) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FactsResponse struct {
	Facts []*Fact `json:"facts,omitempty" url:"facts,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (f *FactsResponse) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FactsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler FactsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FactsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties

	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FactsResponse) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type GraphDataType string

const (
	GraphDataTypeText    GraphDataType = "text"
	GraphDataTypeJSON    GraphDataType = "json"
	GraphDataTypeMessage GraphDataType = "message"
)

func NewGraphDataTypeFromString(s string) (GraphDataType, error) {
	switch s {
	case "text":
		return GraphDataTypeText, nil
	case "json":
		return GraphDataTypeJSON, nil
	case "message":
		return GraphDataTypeMessage, nil
	}
	var t GraphDataType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GraphDataType) Ptr() *GraphDataType {
	return &g
}

type MemoryType string

const (
	MemoryTypePerpetual        MemoryType = "perpetual"
	MemoryTypeSummaryRetriever MemoryType = "summary_retriever"
	MemoryTypeMessageWindow    MemoryType = "message_window"
)

func NewMemoryTypeFromString(s string) (MemoryType, error) {
	switch s {
	case "perpetual":
		return MemoryTypePerpetual, nil
	case "summary_retriever":
		return MemoryTypeSummaryRetriever, nil
	case "message_window":
		return MemoryTypeMessageWindow, nil
	}
	var t MemoryType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (m MemoryType) Ptr() *MemoryType {
	return &m
}

type SearchType string

const (
	SearchTypeSimilarity SearchType = "similarity"
	SearchTypeMmr        SearchType = "mmr"
)

func NewSearchTypeFromString(s string) (SearchType, error) {
	switch s {
	case "similarity":
		return SearchTypeSimilarity, nil
	case "mmr":
		return SearchTypeMmr, nil
	}
	var t SearchType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SearchType) Ptr() *SearchType {
	return &s
}

type Session struct {
	Classifications       map[string]string             `json:"classifications,omitempty" url:"classifications,omitempty"`
	CreatedAt             *string                       `json:"created_at,omitempty" url:"created_at,omitempty"`
	DeletedAt             *string                       `json:"deleted_at,omitempty" url:"deleted_at,omitempty"`
	EndedAt               *string                       `json:"ended_at,omitempty" url:"ended_at,omitempty"`
	FactRatingInstruction *ApidataFactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"fact_rating_instruction,omitempty"`
	Facts                 []string                      `json:"facts,omitempty" url:"facts,omitempty"`
	// TODO deprecate
	ID          *int                   `json:"id,omitempty" url:"id,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	ProjectUUID *string                `json:"project_uuid,omitempty" url:"project_uuid,omitempty"`
	SessionID   *string                `json:"session_id,omitempty" url:"session_id,omitempty"`
	UpdatedAt   *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// Must be a pointer to allow for null values
	UserID *string `json:"user_id,omitempty" url:"user_id,omitempty"`
	UUID   *string `json:"uuid,omitempty" url:"uuid,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *Session) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Session) UnmarshalJSON(data []byte) error {
	type unmarshaler Session
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Session(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Session) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type SessionFactRatingExamples = interface{}

type SessionFactRatingInstruction = interface{}

type SuccessResponse struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SuccessResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SuccessResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SuccessResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SuccessResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SuccessResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type UpdateGroupRequest = interface{}
