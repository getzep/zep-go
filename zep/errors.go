package zep

type ZepError struct {
	Message      string
	Code         *int
	ResponseData interface{}
}

func (e *ZepError) Error() string {
	return e.Message
}

// APIError is a general error returned by the Zep API
type APIError struct {
	ZepError
}

// NotFoundError is returned when a resource is not found
type NotFoundError struct {
	ZepError
}

// AuthenticationError is returned when authentication fails
type AuthenticationError struct {
	ZepError
}
