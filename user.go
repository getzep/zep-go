// This file was auto-generated by Fern from our API Definition.

package zep

type CreateUserRequest struct {
	// The email address of the user.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The first name of the user.
	FirstName *string `json:"first_name,omitempty" url:"first_name,omitempty"`
	// The last name of the user.
	LastName *string `json:"last_name,omitempty" url:"last_name,omitempty"`
	// The metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// The unique identifier of the user.
	UserID *string `json:"user_id,omitempty" url:"user_id,omitempty"`
}

type UserListOrderedRequest struct {
	// Page number for pagination, starting from 1
	PageNumber *int `json:"-" url:"pageNumber,omitempty"`
	// Number of users to retrieve per page
	PageSize *int `json:"-" url:"pageSize,omitempty"`
}

type UpdateUserRequest struct {
	// The email address of the user.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The first name of the user.
	FirstName *string `json:"first_name,omitempty" url:"first_name,omitempty"`
	// The last name of the user.
	LastName *string `json:"last_name,omitempty" url:"last_name,omitempty"`
	// The metadata to update
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
}
