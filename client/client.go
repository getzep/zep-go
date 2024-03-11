// This file was auto-generated by Fern from our API Definition.

package client

import (
	collection "github.com/getzep/zep-go/collection"
	core "github.com/getzep/zep-go/core"
	document "github.com/getzep/zep-go/document"
	memory "github.com/getzep/zep-go/memory"
	messages "github.com/getzep/zep-go/messages"
	option "github.com/getzep/zep-go/option"
	search "github.com/getzep/zep-go/search"
	session "github.com/getzep/zep-go/session"
	user "github.com/getzep/zep-go/user"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Document   *document.Client
	Collection *collection.Client
	Session    *session.Client
	Memory     *memory.Client
	Messages   *messages.Client
	Search     *search.Client
	User       *user.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:     options.ToHeader(),
		Document:   document.NewClient(opts...),
		Collection: collection.NewClient(opts...),
		Session:    session.NewClient(opts...),
		Memory:     memory.NewClient(opts...),
		Messages:   messages.NewClient(opts...),
		Search:     search.NewClient(opts...),
		User:       user.NewClient(opts...),
	}
}