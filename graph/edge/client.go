// This file was auto-generated by Fern from our API Definition.

package edge

import (
	context "context"
	v2 "github.com/getzep/zep-go/v2"
	core "github.com/getzep/zep-go/v2/core"
	internal "github.com/getzep/zep-go/v2/internal"
	option "github.com/getzep/zep-go/v2/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		options.APIKey = os.Getenv("ZEP_API_KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Get all edges for a group
func (c *Client) GetByGroupID(
	ctx context.Context,
	// Group ID
	groupID string,
	opts ...option.RequestOption,
) ([]*v2.EntityEdge, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.getzep.com/api/v2",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/graph/edge/group/%v",
		groupID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response []*v2.EntityEdge
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get all edges for a user
func (c *Client) GetByUserID(
	ctx context.Context,
	// User ID
	userID string,
	opts ...option.RequestOption,
) ([]*v2.EntityEdge, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.getzep.com/api/v2",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/graph/edge/user/%v",
		userID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response []*v2.EntityEdge
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get a specific edge by its UUID
func (c *Client) Get(
	ctx context.Context,
	// Edge UUID
	_uuid string,
	opts ...option.RequestOption,
) (*v2.EntityEdge, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.getzep.com/api/v2",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/graph/edge/%v",
		_uuid,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &v2.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *v2.EntityEdge
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete an edge by UUID
func (c *Client) Delete(
	ctx context.Context,
	// Edge UUID
	__uuid string,
	opts ...option.RequestOption,
) (*v2.SuccessResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.getzep.com/api/v2",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/graph/edge/%v",
		__uuid,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &v2.BadRequestError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &v2.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *v2.SuccessResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
