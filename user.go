// This file was auto-generated by Fern from our API Definition.

package zep

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/getzep/zep-go/v2/core"
)

type CreateUserRequest struct {
	// The email address of the user.
	Email *string `json:"email,omitempty" url:"-"`
	// Optional instruction to use for fact rating.
	FactRatingInstruction *FactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"-"`
	// The first name of the user.
	FirstName *string `json:"first_name,omitempty" url:"-"`
	// The last name of the user.
	LastName *string `json:"last_name,omitempty" url:"-"`
	// The metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
	// The unique identifier of the user.
	UserID *string `json:"user_id,omitempty" url:"-"`
}

type UserListOrderedRequest struct {
	// Page number for pagination, starting from 1
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
	// Number of users to retrieve per page
	PageSize *int `json:"-" url:"pageSize,omitempty"`
}

type User struct {
	CreatedAt             *string                `json:"created_at,omitempty" url:"created_at,omitempty"`
	DeletedAt             *string                `json:"deleted_at,omitempty" url:"deleted_at,omitempty"`
	Email                 *string                `json:"email,omitempty" url:"email,omitempty"`
	FactRatingInstruction *FactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"fact_rating_instruction,omitempty"`
	FirstName             *string                `json:"first_name,omitempty" url:"first_name,omitempty"`
	ID                    *int                   `json:"id,omitempty" url:"id,omitempty"`
	LastName              *string                `json:"last_name,omitempty" url:"last_name,omitempty"`
	Metadata              map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	ProjectUUID           *string                `json:"project_uuid,omitempty" url:"project_uuid,omitempty"`
	SessionCount          *int                   `json:"session_count,omitempty" url:"session_count,omitempty"`
	UpdatedAt             *string                `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UserID                *string                `json:"user_id,omitempty" url:"user_id,omitempty"`
	UUID                  *string                `json:"uuid,omitempty" url:"uuid,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *User) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *User) UnmarshalJSON(data []byte) error {
	type unmarshaler User
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = User(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

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

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UserListResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserListResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

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

type UpdateUserRequest struct {
	// The email address of the user.
	Email *string `json:"email,omitempty" url:"-"`
	// Optional instruction to use for fact rating.
	FactRatingInstruction *FactRatingInstruction `json:"fact_rating_instruction,omitempty" url:"-"`
	// The first name of the user.
	FirstName *string `json:"first_name,omitempty" url:"-"`
	// The last name of the user.
	LastName *string `json:"last_name,omitempty" url:"-"`
	// The metadata to update
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"-"`
}
