package zep

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"net/http"
	"strings"
	"time"
)

const (
	APIBasePath           = "api/v1"
	ServerErrorMessage    = "Failed to connect to Zep server. Please check that the server is running, the API URL is correct, and no other process is using the same port"
	MinServerVersion      = "0.17.0"
	MinServerWarningMsg   = "You are using an incompatible Zep server version. Please upgrade to " + MinServerVersion + " or later."
	DefaultRequestTimeout = 30 // In seconds
)

var _ Client = &DefaultClient{}

type Client interface {
	GetFullURL(endpoint string) string
	CheckServer() error
	HandleRequest(requestPromise *http.Request, notFoundMessage string) (*http.Response, error)
}

// NewClient creates a new Client. If client is provided, it will be used to make requests.
// Otherwise, a default client will be created with a 30 second timeout.
func NewClient(serverURL string, apiKey string, client *http.Client) *DefaultClient {
	headers := make(map[string]string)
	if apiKey != "" {
		headers["Authorization"] = "Bearer " + apiKey
	}
	if client == nil {
		client = &http.Client{Timeout: DefaultRequestTimeout * time.Second}
	}

	// Remove trailing slash from server URL
	serverURL = strings.TrimSuffix(serverURL, "/")

	zepClient := &DefaultClient{ServerURL: serverURL, Headers: headers, Client: client}
	err := zepClient.CheckServer()
	if err != nil {
		fmt.Println(err)
	}

	u := NewUserManager(zepClient)
	zepClient.User = u
	m := NewMemoryManager(zepClient)
	zepClient.Memory = m
	d := NewDocumentManager(zepClient)
	zepClient.Document = d

	return zepClient
}

// DefaultClient is the implementation of Client.
type DefaultClient struct {
	ServerURL string
	Headers   map[string]string
	Client    *http.Client
	User      *UserManager
	Memory    *MemoryManager
	Document  *DocumentManager
}

// GetFullURL returns the full URL for the given endpoint.
// It concatenates the server URL, API base path, and endpoint.
func (z *DefaultClient) GetFullURL(endpoint string) string {
	return joinPaths(z.ServerURL, APIBasePath, endpoint)
}

// CheckServer checks if the server is running and returns an error if it is not.
// It also checks if the server version is compatible with this client.
func (z *DefaultClient) CheckServer() error {
	healthCheck := "/healthz"
	healthCheckURL := z.ServerURL + healthCheck

	req, err := http.NewRequest("GET", healthCheckURL, nil)
	if err != nil {
		return err
	}
	for key, value := range z.Headers {
		req.Header.Add(key, value)
	}

	resp, err := z.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &ZepError{Message: ServerErrorMessage}
	}

	zepServerVersion := resp.Header.Get("X-Zep-Version")
	meetsMinVersion, err := isVersionGreaterOrEqual(zepServerVersion)
	if err != nil {
		return err
	}
	if !meetsMinVersion {
		fmt.Println("Warning: " + MinServerWarningMsg)
	}

	return nil
}

// HandleRequest makes the request and returns the response if the request is successful.
// If the request is not successful, it returns an appropriate error:
// - NotFoundError if the status code is 404
// - AuthenticationError if the status code is 401
// - APIError if the status code is anything else
func (z *DefaultClient) HandleRequest(requestPromise *http.Request, notFoundMessage string) (*http.Response, error) {
	response, err := z.Client.Do(requestPromise)
	if err != nil {
		return nil, &ZepError{Message: ServerErrorMessage + ": " + err.Error()}
	}

	switch response.StatusCode {
	case http.StatusOK:
		return response, nil
	case http.StatusNotFound:
		return nil, &NotFoundError{ZepError: ZepError{Message: notFoundMessage}}
	case http.StatusUnauthorized:
		return nil, &AuthenticationError{ZepError: ZepError{Message: "Authentication failed."}}
	default:
		return nil, &APIError{ZepError: ZepError{Message: fmt.Sprintf("Got an unexpected status code: %d", response.StatusCode)}}
	}
}

func isVersionGreaterOrEqual(version string) (bool, error) {
	c, err := semver.NewConstraint(">= " + MinServerVersion)
	if err != nil {
		return false, err
	}
	currentVersion, err := semver.NewVersion(version)
	if err != nil {
		return false, err
	}
	return c.Check(currentVersion), nil
}

func joinPaths(paths ...string) string {
	return strings.Join(paths, "/")
}
