package generatedclient

import (
	"context"
	"net/http"

	"github.com/getzep/zep-go/v4/core"
	"github.com/getzep/zep-go/v4/internal"
)

type Client struct {
	caller *internal.Caller
}

func New(httpClient core.HTTPClient) *Client {
	return &Client{
		caller: internal.NewCaller(&internal.CallerParams{
			Client: httpClient,
		}),
	}
}

func (c *Client) Mutate(ctx context.Context, opts ...core.IdempotentRequestOption) error {
	options := core.NewIdempotentRequestOptions(opts...)
	return c.call(ctx, http.MethodPost, options.ToHeader())
}

func (c *Client) GetRead(ctx context.Context, opts ...core.RequestOption) error {
	options := core.NewRequestOptions(opts...)
	return c.call(ctx, http.MethodGet, options.ToHeader())
}

func (c *Client) PostRead(ctx context.Context, opts ...core.RequestOption) error {
	options := core.NewRequestOptions(opts...)
	return c.call(ctx, http.MethodPost, options.ToHeader())
}

func (c *Client) call(ctx context.Context, method string, headers http.Header) error {
	_, err := c.caller.Call(ctx, &internal.CallParams{
		URL:                "http://localhost/fixture",
		Method:             method,
		Headers:            headers,
		MaxAttempts:        2,
		ResponseIsOptional: true,
	})
	return err
}
