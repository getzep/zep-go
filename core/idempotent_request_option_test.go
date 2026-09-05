package core_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/getzep/zep-go/v4/internal/testfixtures/generatedclient"
	"github.com/getzep/zep-go/v4/option"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type recordingHTTPClient struct {
	mu         sync.Mutex
	statuses   []int
	requestLog []*http.Request
}

func (c *recordingHTTPClient) Do(request *http.Request) (*http.Response, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	requestCopy := request.Clone(request.Context())
	requestCopy.Header = request.Header.Clone()
	c.requestLog = append(c.requestLog, requestCopy)

	status := http.StatusOK
	if len(c.requestLog) <= len(c.statuses) {
		status = c.statuses[len(c.requestLog)-1]
	}
	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("")),
	}, nil
}

func (c *recordingHTTPClient) requests() []*http.Request {
	c.mu.Lock()
	defer c.mu.Unlock()

	requests := make([]*http.Request, len(c.requestLog))
	copy(requests, c.requestLog)
	return requests
}

func TestGeneratedMutationGeneratesUUIDV4(t *testing.T) {
	httpClient := &recordingHTTPClient{}
	client := generatedclient.New(httpClient)

	require.NoError(t, client.Mutate(context.Background()))
	requests := httpClient.requests()
	require.Len(t, requests, 1)
	parsed, err := uuid.Parse(requests[0].Header.Get("Idempotency-Key"))
	require.NoError(t, err)
	require.Equal(t, uuid.Version(4), parsed.Version())
}

func TestGeneratedMutationReusesKeyAcrossRetries(t *testing.T) {
	httpClient := &recordingHTTPClient{statuses: []int{http.StatusInternalServerError, http.StatusOK}}
	client := generatedclient.New(httpClient)

	require.NoError(t, client.Mutate(context.Background()))
	requests := httpClient.requests()
	require.Len(t, requests, 2)
	require.NotEmpty(t, requests[0].Header.Get("Idempotency-Key"))
	require.Equal(t, requests[0].Header.Get("Idempotency-Key"), requests[1].Header.Get("Idempotency-Key"))
}

func TestGeneratedMutationPreservesCallerKey(t *testing.T) {
	httpClient := &recordingHTTPClient{}
	client := generatedclient.New(httpClient)
	idempotencyKey := "caller-key"

	require.NoError(t, client.Mutate(context.Background(), option.WithIdempotencyKey(&idempotencyKey)))
	requests := httpClient.requests()
	require.Len(t, requests, 1)
	require.Equal(t, idempotencyKey, requests[0].Header.Get("Idempotency-Key"))
}

func TestGeneratedGetReadOmitsIdempotencyKey(t *testing.T) {
	httpClient := &recordingHTTPClient{}
	client := generatedclient.New(httpClient)

	require.NoError(t, client.GetRead(context.Background()))
	requests := httpClient.requests()
	require.Len(t, requests, 1)
	require.Empty(t, requests[0].Header.Get("Idempotency-Key"))
}

func TestGeneratedPostReadOmitsIdempotencyKey(t *testing.T) {
	httpClient := &recordingHTTPClient{}
	client := generatedclient.New(httpClient)

	require.NoError(t, client.PostRead(context.Background()))
	requests := httpClient.requests()
	require.Len(t, requests, 1)
	require.Empty(t, requests[0].Header.Get("Idempotency-Key"))
}
