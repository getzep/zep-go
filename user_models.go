package main

import "time"

type User struct {
	UUID      *string                 `json:"uuid,omitempty"`
	ID        *int                    `json:"id,omitempty"`
	CreatedAt *time.Time              `json:"created_at,omitempty"`
	UpdatedAt *time.Time              `json:"updated_at,omitempty"`
	DeletedAt *time.Time              `json:"deleted_at,omitempty"`
	UserID    string                  `json:"user_id"`
	Email     *string                 `json:"email,omitempty"`
	FirstName *string                 `json:"first_name,omitempty"`
	LastName  *string                 `json:"last_name,omitempty"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"`
}

type CreateUserRequest struct {
	UserID    string                  `json:"user_id"`
	Email     *string                 `json:"email,omitempty"`
	FirstName *string                 `json:"first_name,omitempty"`
	LastName  *string                 `json:"last_name,omitempty"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateUserRequest struct {
	UUID      *string                 `json:"uuid,omitempty"`
	UserID    string                  `json:"user_id"`
	Email     *string                 `json:"email,omitempty"`
	FirstName *string                 `json:"first_name,omitempty"`
	LastName  *string                 `json:"last_name,omitempty"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"`
}
