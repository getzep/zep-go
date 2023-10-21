package main

import "time"

type Session struct {
	UUID      *string                `json:"uuid,omitempty"`
	CreatedAt *time.Time             `json:"created_at,omitempty"`
	UpdatedAt *time.Time             `json:"updated_at,omitempty"`
	DeletedAt *time.Time             `json:"deleted_at,omitempty"`
	SessionID string                 `json:"session_id"`
	Metadata  map[string]interface{} `json:"metadata"`
	UserID    *string                `json:"user_id,omitempty"`
}

type Message struct {
	UUID       *string                 `json:"uuid,omitempty"`
	CreatedAt  *time.Time              `json:"created_at,omitempty"`
	Role       string                  `json:"role"`
	Content    string                  `json:"content"`
	TokenCount *int                    `json:"token_count,omitempty"`
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
}

type Summary struct {
	UUID              string    `json:"uuid"`
	CreatedAt         time.Time `json:"created_at"`
	Content           string    `json:"content"`
	RecentMessageUUID string    `json:"recent_message_uuid"`
	TokenCount        int       `json:"token_count"`
}

type Memory struct {
	Messages   *[]Message              `json:"messages,omitempty"`
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Summary    *Summary                `json:"summary,omitempty"`
	UUID       *string                 `json:"uuid,omitempty"`
	CreatedAt  *time.Time              `json:"created_at,omitempty"`
	TokenCount *int                    `json:"token_count,omitempty"`
}

type MemorySearchPayload struct {
	Text       *string                 `json:"text,omitempty"`
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	SearchType *string                 `json:"search_type,omitempty"`
	MMRLambda  *float64                `json:"mmr_lambda,omitempty"`
}

type MemorySearchResult struct {
	Message  *Message                `json:"message,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	Summary  *string                 `json:"summary,omitempty"`
	Dist     *float64                `json:"dist,omitempty"`
}
