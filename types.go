// This file was auto-generated by Fern from our API Definition.

package zep

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/getzep/zep-go/core"
)

type APIError struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`

	_rawJSON json.RawMessage
}

func (a *APIError) UnmarshalJSON(data []byte) error {
	type unmarshaler APIError
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = APIError(value)
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

type ClassifySessionResponse struct {
	Class *string `json:"class,omitempty" url:"class,omitempty"`
	Name  *string `json:"name,omitempty" url:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ClassifySessionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ClassifySessionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClassifySessionResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClassifySessionResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreateDocumentRequest struct {
	Content    *string                `json:"content,omitempty" url:"content,omitempty"`
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Embedding  []float64              `json:"embedding,omitempty" url:"embedding,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CreateDocumentRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateDocumentRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateDocumentRequest(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateDocumentRequest) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type DocumentCollectionResponse struct {
	CreatedAt   *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// Number of documents in the collection
	DocumentCount *int `json:"document_count,omitempty" url:"document_count,omitempty"`
	// Number of documents with embeddings
	DocumentEmbeddedCount *int                   `json:"document_embedded_count,omitempty" url:"document_embedded_count,omitempty"`
	EmbeddingDimensions   *int                   `json:"embedding_dimensions,omitempty" url:"embedding_dimensions,omitempty"`
	EmbeddingModelName    *string                `json:"embedding_model_name,omitempty" url:"embedding_model_name,omitempty"`
	IsAutoEmbedded        *bool                  `json:"is_auto_embedded,omitempty" url:"is_auto_embedded,omitempty"`
	IsIndexed             *bool                  `json:"is_indexed,omitempty" url:"is_indexed,omitempty"`
	IsNormalized          *bool                  `json:"is_normalized,omitempty" url:"is_normalized,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	Name                  *string                `json:"name,omitempty" url:"name,omitempty"`
	UpdatedAt             *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UUID                  *string                `json:"uuid,omitempty" url:"uuid,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DocumentCollectionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DocumentCollectionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DocumentCollectionResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DocumentCollectionResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DocumentResponse struct {
	Content    *string                `json:"content,omitempty" url:"content,omitempty"`
	CreatedAt  *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Embedding  []float64              `json:"embedding,omitempty" url:"embedding,omitempty"`
	IsEmbedded *bool                  `json:"is_embedded,omitempty" url:"is_embedded,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	UpdatedAt  *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UUID       *string                `json:"uuid,omitempty" url:"uuid,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DocumentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DocumentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DocumentResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DocumentResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DocumentSearchResult struct {
	Content    *string                `json:"content,omitempty" url:"content,omitempty"`
	CreatedAt  *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Embedding  []float64              `json:"embedding,omitempty" url:"embedding,omitempty"`
	IsEmbedded *bool                  `json:"is_embedded,omitempty" url:"is_embedded,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	Score      *float64               `json:"score,omitempty" url:"score,omitempty"`
	UpdatedAt  *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UUID       *string                `json:"uuid,omitempty" url:"uuid,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DocumentSearchResult) UnmarshalJSON(data []byte) error {
	type unmarshaler DocumentSearchResult
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DocumentSearchResult(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DocumentSearchResult) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DocumentSearchResultPage struct {
	CurrentPage *int                    `json:"current_page,omitempty" url:"current_page,omitempty"`
	QueryVector []float64               `json:"query_vector,omitempty" url:"query_vector,omitempty"`
	ResultCount *int                    `json:"result_count,omitempty" url:"result_count,omitempty"`
	Results     []*DocumentSearchResult `json:"results,omitempty" url:"results,omitempty"`
	TotalPages  *int                    `json:"total_pages,omitempty" url:"total_pages,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DocumentSearchResultPage) UnmarshalJSON(data []byte) error {
	type unmarshaler DocumentSearchResultPage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DocumentSearchResultPage(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DocumentSearchResultPage) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type Memory struct {
	// Most recent list of facts derived from the session. Included only with perpetual memory type.
	Facts []string `json:"facts,omitempty" url:"facts,omitempty"`
	// A list of message objects, where each message contains a role and content.
	Messages []*Message `json:"messages,omitempty" url:"messages,omitempty"`
	// A dictionary containing metadata associated with the memory.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// A Summary object.
	Summary *Summary `json:"summary,omitempty" url:"summary,omitempty"`

	_rawJSON json.RawMessage
}

func (m *Memory) UnmarshalJSON(data []byte) error {
	type unmarshaler Memory
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = Memory(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *Memory) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type MemorySearchResult struct {
	Embedding []float64              `json:"embedding,omitempty" url:"embedding,omitempty"`
	Message   *Message               `json:"message,omitempty" url:"message,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	Score     *float64               `json:"score,omitempty" url:"score,omitempty"`
	Summary   *Summary               `json:"summary,omitempty" url:"summary,omitempty"`

	_rawJSON json.RawMessage
}

func (m *MemorySearchResult) UnmarshalJSON(data []byte) error {
	type unmarshaler MemorySearchResult
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MemorySearchResult(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MemorySearchResult) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type Message struct {
	// The content of the message.
	Content *string `json:"content,omitempty" url:"content,omitempty"`
	// The timestamp of when the message was created.
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The metadata associated with the message.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// The role of the sender of the message (e.g., "user", "assistant").
	Role *string `json:"role,omitempty" url:"role,omitempty"`
	// The type of the role (e.g., "user", "system").
	RoleType *RoleType `json:"role_type,omitempty" url:"role_type,omitempty"`
	// The number of tokens in the message.
	TokenCount *int `json:"token_count,omitempty" url:"token_count,omitempty"`
	// The timestamp of when the message was last updated.
	UpdatedAt *string `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The unique identifier of the message.
	UUID *string `json:"uuid,omitempty" url:"uuid,omitempty"`

	_rawJSON json.RawMessage
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type unmarshaler Message
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = Message(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *Message) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type MessageListResponse struct {
	// A list of message objects.
	Messages []*Message `json:"messages,omitempty" url:"messages,omitempty"`
	// The number of messages returned.
	RowCount *int `json:"row_count,omitempty" url:"row_count,omitempty"`
	// The total number of messages.
	TotalCount *int `json:"total_count,omitempty" url:"total_count,omitempty"`

	_rawJSON json.RawMessage
}

func (m *MessageListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler MessageListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MessageListResponse(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *MessageListResponse) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type ModelsZepDataClass struct {
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	Name        *string `json:"name,omitempty" url:"name,omitempty"`
	Type        *string `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelsZepDataClass) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelsZepDataClass
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelsZepDataClass(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelsZepDataClass) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type Question struct {
	Question *string `json:"question,omitempty" url:"question,omitempty"`

	_rawJSON json.RawMessage
}

func (q *Question) UnmarshalJSON(data []byte) error {
	type unmarshaler Question
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*q = Question(value)
	q._rawJSON = json.RawMessage(data)
	return nil
}

func (q *Question) String() string {
	if len(q._rawJSON) > 0 {
		if value, err := core.StringifyJSON(q._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(q); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", q)
}

type RoleType string

const (
	RoleTypeNoRole        RoleType = "norole"
	RoleTypeSystemRole    RoleType = "system"
	RoleTypeAssistantRole RoleType = "assistant"
	RoleTypeUserRole      RoleType = "user"
	RoleTypeFunctionRole  RoleType = "function"
	RoleTypeToolRole      RoleType = "tool"
)

func NewRoleTypeFromString(s string) (RoleType, error) {
	switch s {
	case "norole":
		return RoleTypeNoRole, nil
	case "system":
		return RoleTypeSystemRole, nil
	case "assistant":
		return RoleTypeAssistantRole, nil
	case "user":
		return RoleTypeUserRole, nil
	case "function":
		return RoleTypeFunctionRole, nil
	case "tool":
		return RoleTypeToolRole, nil
	}
	var t RoleType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RoleType) Ptr() *RoleType {
	return &r
}

type SearchScope string

const (
	SearchScopeMessages SearchScope = "messages"
	SearchScopeSummary  SearchScope = "summary"
)

func NewSearchScopeFromString(s string) (SearchScope, error) {
	switch s {
	case "messages":
		return SearchScopeMessages, nil
	case "summary":
		return SearchScopeSummary, nil
	}
	var t SearchScope
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SearchScope) Ptr() *SearchScope {
	return &s
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
	Classifications map[string]string      `json:"classifications,omitempty" url:"classifications,omitempty"`
	CreatedAt       *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	DeletedAt       *string                `json:"deleted_at,omitempty" url:"deleted_at,omitempty"`
	Facts           []string               `json:"facts,omitempty" url:"facts,omitempty"`
	ID              *int                   `json:"id,omitempty" url:"id,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	ProjectUUID     *string                `json:"project_uuid,omitempty" url:"project_uuid,omitempty"`
	SessionID       *string                `json:"session_id,omitempty" url:"session_id,omitempty"`
	UpdatedAt       *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// Must be a pointer to allow for null values
	UserID *string `json:"user_id,omitempty" url:"user_id,omitempty"`
	UUID   *string `json:"uuid,omitempty" url:"uuid,omitempty"`

	_rawJSON json.RawMessage
}

func (s *Session) UnmarshalJSON(data []byte) error {
	type unmarshaler Session
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Session(value)
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

type SessionListResponse struct {
	ResponseCount *int       `json:"response_count,omitempty" url:"response_count,omitempty"`
	Sessions      []*Session `json:"sessions,omitempty" url:"sessions,omitempty"`
	TotalCount    *int       `json:"total_count,omitempty" url:"total_count,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SessionListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SessionListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SessionListResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SessionListResponse) String() string {
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

type SuccessResponse struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SuccessResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SuccessResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SuccessResponse(value)
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

type Summary struct {
	// The content of the summary.
	Content *string `json:"content,omitempty" url:"content,omitempty"`
	// The timestamp of when the summary was created.
	CreatedAt           *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	RelatedMessageUUIDs []string               `json:"related_message_uuids,omitempty" url:"related_message_uuids,omitempty"`
	// The number of tokens in the summary.
	TokenCount *int `json:"token_count,omitempty" url:"token_count,omitempty"`
	// The unique identifier of the summary.
	UUID *string `json:"uuid,omitempty" url:"uuid,omitempty"`

	_rawJSON json.RawMessage
}

func (s *Summary) UnmarshalJSON(data []byte) error {
	type unmarshaler Summary
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Summary(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Summary) String() string {
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

type SummaryListResponse struct {
	RowCount   *int       `json:"row_count,omitempty" url:"row_count,omitempty"`
	Summaries  []*Summary `json:"summaries,omitempty" url:"summaries,omitempty"`
	TotalCount *int       `json:"total_count,omitempty" url:"total_count,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SummaryListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SummaryListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SummaryListResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SummaryListResponse) String() string {
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

type UpdateDocumentListRequest struct {
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	UUID       string                 `json:"uuid" url:"uuid"`

	_rawJSON json.RawMessage
}

func (u *UpdateDocumentListRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateDocumentListRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateDocumentListRequest(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateDocumentListRequest) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type User struct {
	CreatedAt    *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	DeletedAt    *string                `json:"deleted_at,omitempty" url:"deleted_at,omitempty"`
	Email        *string                `json:"email,omitempty" url:"email,omitempty"`
	FirstName    *string                `json:"first_name,omitempty" url:"first_name,omitempty"`
	ID           *int                   `json:"id,omitempty" url:"id,omitempty"`
	LastName     *string                `json:"last_name,omitempty" url:"last_name,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	ProjectUUID  *string                `json:"project_uuid,omitempty" url:"project_uuid,omitempty"`
	SessionCount *int                   `json:"session_count,omitempty" url:"session_count,omitempty"`
	UpdatedAt    *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UserID       *string                `json:"user_id,omitempty" url:"user_id,omitempty"`
	UUID         *string                `json:"uuid,omitempty" url:"uuid,omitempty"`

	_rawJSON json.RawMessage
}

func (u *User) UnmarshalJSON(data []byte) error {
	type unmarshaler User
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = User(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *User) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserListResponse struct {
	RowCount   *int    `json:"row_count,omitempty" url:"row_count,omitempty"`
	TotalCount *int    `json:"total_count,omitempty" url:"total_count,omitempty"`
	Users      []*User `json:"users,omitempty" url:"users,omitempty"`

	_rawJSON json.RawMessage
}

func (u *UserListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserListResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserListResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
