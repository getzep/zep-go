package zep

import "time"

type Document struct {
	UUID       string                 `json:"uuid,omitempty"`
	CreatedAt  time.Time              `json:"created_at,omitempty"`
	UpdatedAt  time.Time              `json:"updated_at,omitempty"`
	DocumentID string                 `json:"document_id,omitempty"`
	Content    string                 `json:"content"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	IsEmbedded bool                   `json:"is_embedded,omitempty"`
	Embedding  []float32              `json:"embedding,omitempty"`
	Score      float64                `json:"score,omitempty"`
}

type DocumentCollectionModel struct {
	UUID                  string                 `json:"uuid,omitempty"`
	CreatedAt             time.Time              `json:"created_at,omitempty"`
	UpdatedAt             time.Time              `json:"updated_at,omitempty"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty"`
	EmbeddingDimensions   int                    `json:"embedding_dimensions,omitempty"`
	IsAutoEmbedded        bool                   `json:"is_auto_embedded,omitempty"`
	IsIndexed             bool                   `json:"is_indexed,omitempty"`
	DocumentCount         int                    `json:"document_count,omitempty"`
	DocumentEmbeddedCount int                    `json:"document_embedded_count,omitempty"`
	IsNormalized          bool                   `json:"is_normalized,omitempty"`
}
