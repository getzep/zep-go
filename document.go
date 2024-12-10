// This file was auto-generated by Fern from our API Definition.

package zep

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/getzep/zep-go/v2/core"
)

type CreateDocumentCollectionRequest struct {
	Description *string                `json:"description,omitempty" url:"-"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" url:"-"`
}

type GetDocumentListRequest struct {
	DocumentIDs []string `json:"document_ids,omitempty" url:"-"`
	UUIDs       []string `json:"uuids,omitempty" url:"-"`
}

type DocumentSearchPayload struct {
	// Limit the number of returned documents
	Limit *int `json:"-" url:"limit,omitempty"`
	// Document metadata to filter on.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
	MinScore *float64               `json:"min_score,omitempty" url:"-"`
	// The lambda parameter for the MMR Reranking Algorithm.
	MmrLambda *float64 `json:"mmr_lambda,omitempty" url:"-"`
	// The type of search to perform. Defaults to "similarity". Must be one of "similarity" or "mmr".
	SearchType *SearchType `json:"search_type,omitempty" url:"-"`
	// The search text.
	Text *string `json:"text,omitempty" url:"-"`
}

type ApidataDocument struct {
	Content    *string                `json:"content,omitempty" url:"content,omitempty"`
	CreatedAt  *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Embedding  []float64              `json:"embedding,omitempty" url:"embedding,omitempty"`
	IsEmbedded *bool                  `json:"is_embedded,omitempty" url:"is_embedded,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	UpdatedAt  *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UUID       *string                `json:"uuid,omitempty" url:"uuid,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *ApidataDocument) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApidataDocument) UnmarshalJSON(data []byte) error {
	type unmarshaler ApidataDocument
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApidataDocument(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApidataDocument) String() string {
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

type ApidataDocumentCollection struct {
	CreatedAt             *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	Description           *string                `json:"description,omitempty" url:"description,omitempty"`
	DocumentCount         *int                   `json:"document_count,omitempty" url:"document_count,omitempty"`
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

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *ApidataDocumentCollection) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApidataDocumentCollection) UnmarshalJSON(data []byte) error {
	type unmarshaler ApidataDocumentCollection
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApidataDocumentCollection(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApidataDocumentCollection) String() string {
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

type ApidataDocumentSearchResponse struct {
	CurrentPage *int                        `json:"current_page,omitempty" url:"current_page,omitempty"`
	QueryVector []float64                   `json:"query_vector,omitempty" url:"query_vector,omitempty"`
	ResultCount *int                        `json:"result_count,omitempty" url:"result_count,omitempty"`
	Results     []*ApidataDocumentWithScore `json:"results,omitempty" url:"results,omitempty"`
	TotalPages  *int                        `json:"total_pages,omitempty" url:"total_pages,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *ApidataDocumentSearchResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApidataDocumentSearchResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ApidataDocumentSearchResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApidataDocumentSearchResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApidataDocumentSearchResponse) String() string {
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

type ApidataDocumentWithScore struct {
	Content    *string                `json:"content,omitempty" url:"content,omitempty"`
	CreatedAt  *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Embedding  []float64              `json:"embedding,omitempty" url:"embedding,omitempty"`
	IsEmbedded *bool                  `json:"is_embedded,omitempty" url:"is_embedded,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	Score      *float64               `json:"score,omitempty" url:"score,omitempty"`
	UpdatedAt  *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UUID       *string                `json:"uuid,omitempty" url:"uuid,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *ApidataDocumentWithScore) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApidataDocumentWithScore) UnmarshalJSON(data []byte) error {
	type unmarshaler ApidataDocumentWithScore
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApidataDocumentWithScore(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApidataDocumentWithScore) String() string {
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

type CreateDocumentRequest struct {
	Content    string                 `json:"content" url:"content"`
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateDocumentRequest) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateDocumentRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateDocumentRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateDocumentRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

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

type UpdateDocumentListRequest struct {
	DocumentID *string                `json:"document_id,omitempty" url:"document_id,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	UUID       string                 `json:"uuid" url:"uuid"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdateDocumentListRequest) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateDocumentListRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateDocumentListRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateDocumentListRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

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

type UpdateDocumentCollectionRequest struct {
	Description *string                `json:"description,omitempty" url:"-"`
	Metadata    map[string]interface{} `json:"metadata,omitempty" url:"-"`
}

type UpdateDocumentRequest struct {
	DocumentID *string                `json:"document_id,omitempty" url:"-"`
	Metadata   map[string]interface{} `json:"metadata,omitempty" url:"-"`
}
