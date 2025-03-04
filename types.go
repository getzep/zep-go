// This file was auto-generated by Fern from our API Definition.

package zep

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/getzep/zep-go/v2/internal"
)

type APIError struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *APIError) GetMessage() *string {
	if a == nil {
		return nil
	}
	return a.Message
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
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *APIError) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
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
	rawJSON         json.RawMessage
}

func (a *ApidataFactRatingExamples) GetHigh() *string {
	if a == nil {
		return nil
	}
	return a.High
}

func (a *ApidataFactRatingExamples) GetLow() *string {
	if a == nil {
		return nil
	}
	return a.Low
}

func (a *ApidataFactRatingExamples) GetMedium() *string {
	if a == nil {
		return nil
	}
	return a.Medium
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
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApidataFactRatingExamples) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
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
	rawJSON         json.RawMessage
}

func (e *EntityEdge) GetCreatedAt() string {
	if e == nil {
		return ""
	}
	return e.CreatedAt
}

func (e *EntityEdge) GetEpisodes() []string {
	if e == nil {
		return nil
	}
	return e.Episodes
}

func (e *EntityEdge) GetExpiredAt() *string {
	if e == nil {
		return nil
	}
	return e.ExpiredAt
}

func (e *EntityEdge) GetFact() string {
	if e == nil {
		return ""
	}
	return e.Fact
}

func (e *EntityEdge) GetInvalidAt() *string {
	if e == nil {
		return nil
	}
	return e.InvalidAt
}

func (e *EntityEdge) GetName() string {
	if e == nil {
		return ""
	}
	return e.Name
}

func (e *EntityEdge) GetSourceNodeUUID() string {
	if e == nil {
		return ""
	}
	return e.SourceNodeUUID
}

func (e *EntityEdge) GetTargetNodeUUID() string {
	if e == nil {
		return ""
	}
	return e.TargetNodeUUID
}

func (e *EntityEdge) GetUUID() string {
	if e == nil {
		return ""
	}
	return e.UUID
}

func (e *EntityEdge) GetValidAt() *string {
	if e == nil {
		return nil
	}
	return e.ValidAt
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
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityEdge) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EntityNode struct {
	// Additional attributes of the node. Dependent on node labels
	Attributes map[string]interface{} `json:"attributes,omitempty" url:"attributes,omitempty"`
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
	rawJSON         json.RawMessage
}

func (e *EntityNode) GetAttributes() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.Attributes
}

func (e *EntityNode) GetCreatedAt() string {
	if e == nil {
		return ""
	}
	return e.CreatedAt
}

func (e *EntityNode) GetLabels() []string {
	if e == nil {
		return nil
	}
	return e.Labels
}

func (e *EntityNode) GetName() string {
	if e == nil {
		return ""
	}
	return e.Name
}

func (e *EntityNode) GetSummary() string {
	if e == nil {
		return ""
	}
	return e.Summary
}

func (e *EntityNode) GetUUID() string {
	if e == nil {
		return ""
	}
	return e.UUID
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
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EntityNode) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
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
	rawJSON         json.RawMessage
}

func (e *Episode) GetContent() string {
	if e == nil {
		return ""
	}
	return e.Content
}

func (e *Episode) GetCreatedAt() string {
	if e == nil {
		return ""
	}
	return e.CreatedAt
}

func (e *Episode) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *Episode) GetSource() *GraphDataType {
	if e == nil {
		return nil
	}
	return e.Source
}

func (e *Episode) GetSourceDescription() *string {
	if e == nil {
		return nil
	}
	return e.SourceDescription
}

func (e *Episode) GetUUID() string {
	if e == nil {
		return ""
	}
	return e.UUID
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
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *Episode) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EpisodeResponse struct {
	Episodes []*Episode `json:"episodes,omitempty" url:"episodes,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EpisodeResponse) GetEpisodes() []*Episode {
	if e == nil {
		return nil
	}
	return e.Episodes
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
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EpisodeResponse) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EpisodeType = interface{}

type Fact struct {
	Content   string  `json:"content" url:"content"`
	CreatedAt string  `json:"created_at" url:"created_at"`
	ExpiredAt *string `json:"expired_at,omitempty" url:"expired_at,omitempty"`
	// Deprecated
	Fact           string   `json:"fact" url:"fact"`
	InvalidAt      *string  `json:"invalid_at,omitempty" url:"invalid_at,omitempty"`
	Name           *string  `json:"name,omitempty" url:"name,omitempty"`
	Rating         *float64 `json:"rating,omitempty" url:"rating,omitempty"`
	SourceNodeName *string  `json:"source_node_name,omitempty" url:"source_node_name,omitempty"`
	TargetNodeName *string  `json:"target_node_name,omitempty" url:"target_node_name,omitempty"`
	UUID           string   `json:"uuid" url:"uuid"`
	ValidAt        *string  `json:"valid_at,omitempty" url:"valid_at,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *Fact) GetContent() string {
	if f == nil {
		return ""
	}
	return f.Content
}

func (f *Fact) GetCreatedAt() string {
	if f == nil {
		return ""
	}
	return f.CreatedAt
}

func (f *Fact) GetExpiredAt() *string {
	if f == nil {
		return nil
	}
	return f.ExpiredAt
}

func (f *Fact) GetFact() string {
	if f == nil {
		return ""
	}
	return f.Fact
}

func (f *Fact) GetInvalidAt() *string {
	if f == nil {
		return nil
	}
	return f.InvalidAt
}

func (f *Fact) GetName() *string {
	if f == nil {
		return nil
	}
	return f.Name
}

func (f *Fact) GetRating() *float64 {
	if f == nil {
		return nil
	}
	return f.Rating
}

func (f *Fact) GetSourceNodeName() *string {
	if f == nil {
		return nil
	}
	return f.SourceNodeName
}

func (f *Fact) GetTargetNodeName() *string {
	if f == nil {
		return nil
	}
	return f.TargetNodeName
}

func (f *Fact) GetUUID() string {
	if f == nil {
		return ""
	}
	return f.UUID
}

func (f *Fact) GetValidAt() *string {
	if f == nil {
		return nil
	}
	return f.ValidAt
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
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *Fact) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FactRatingExamples struct {
	High   *string `json:"high,omitempty" url:"high,omitempty"`
	Low    *string `json:"low,omitempty" url:"low,omitempty"`
	Medium *string `json:"medium,omitempty" url:"medium,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *FactRatingExamples) GetHigh() *string {
	if f == nil {
		return nil
	}
	return f.High
}

func (f *FactRatingExamples) GetLow() *string {
	if f == nil {
		return nil
	}
	return f.Low
}

func (f *FactRatingExamples) GetMedium() *string {
	if f == nil {
		return nil
	}
	return f.Medium
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
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *FactRatingExamples) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
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
	rawJSON         json.RawMessage
}

func (f *FactRatingInstruction) GetExamples() *FactRatingExamples {
	if f == nil {
		return nil
	}
	return f.Examples
}

func (f *FactRatingInstruction) GetInstruction() *string {
	if f == nil {
		return nil
	}
	return f.Instruction
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
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *FactRatingInstruction) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FactsResponse struct {
	Facts []*Fact `json:"facts,omitempty" url:"facts,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *FactsResponse) GetFacts() []*Fact {
	if f == nil {
		return nil
	}
	return f.Facts
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
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *FactsResponse) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
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
	Classifications map[string]string `json:"classifications,omitempty" url:"classifications,omitempty"`
	CreatedAt       *string           `json:"created_at,omitempty" url:"created_at,omitempty"`
	DeletedAt       *string           `json:"deleted_at,omitempty" url:"deleted_at,omitempty"`
	EndedAt         *string           `json:"ended_at,omitempty" url:"ended_at,omitempty"`
	// Deprecated
	FactRatingInstruction *FactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"fact_rating_instruction,omitempty"`
	// Deprecated
	Facts []string `json:"facts,omitempty" url:"facts,omitempty"`
	ID    *int     `json:"id,omitempty" url:"id,omitempty"`
	// Deprecated
	Metadata    map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	ProjectUUID *string                `json:"project_uuid,omitempty" url:"project_uuid,omitempty"`
	SessionID   *string                `json:"session_id,omitempty" url:"session_id,omitempty"`
	// Deprecated
	UpdatedAt *string `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UserID    *string `json:"user_id,omitempty" url:"user_id,omitempty"`
	UUID      *string `json:"uuid,omitempty" url:"uuid,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *Session) GetClassifications() map[string]string {
	if s == nil {
		return nil
	}
	return s.Classifications
}

func (s *Session) GetCreatedAt() *string {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *Session) GetDeletedAt() *string {
	if s == nil {
		return nil
	}
	return s.DeletedAt
}

func (s *Session) GetEndedAt() *string {
	if s == nil {
		return nil
	}
	return s.EndedAt
}

func (s *Session) GetFactRatingInstruction() *FactRatingInstruction {
	if s == nil {
		return nil
	}
	return s.FactRatingInstruction
}

func (s *Session) GetFacts() []string {
	if s == nil {
		return nil
	}
	return s.Facts
}

func (s *Session) GetID() *int {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *Session) GetMetadata() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.Metadata
}

func (s *Session) GetProjectUUID() *string {
	if s == nil {
		return nil
	}
	return s.ProjectUUID
}

func (s *Session) GetSessionID() *string {
	if s == nil {
		return nil
	}
	return s.SessionID
}

func (s *Session) GetUpdatedAt() *string {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *Session) GetUserID() *string {
	if s == nil {
		return nil
	}
	return s.UserID
}

func (s *Session) GetUUID() *string {
	if s == nil {
		return nil
	}
	return s.UUID
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
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *Session) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type SessionFactRatingExamples = interface{}

type SessionFactRatingInstruction = interface{}

type SuccessResponse struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SuccessResponse) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
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
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SuccessResponse) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
