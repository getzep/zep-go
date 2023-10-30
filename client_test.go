package zep

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZepClient(t *testing.T) {
	serverURL := "http://localhost"
	apiKey := "testKey"
	client := &http.Client{}

	zepClient := NewClient(serverURL, apiKey, client)

	assert.Equal(t, serverURL, zepClient.ServerURL)
	assert.Equal(t, apiKey, zepClient.Headers["Authorization"][7:])
	assert.Equal(t, client, zepClient.Client)
}

func TestCheckServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("X-Zep-Version", string(MinServerVersion))
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	zepClient := NewClient(server.URL, "", nil)

	err := zepClient.CheckServer()

	assert.Nil(t, err)
}

func TestCheckServer_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	zepClient := NewClient(server.URL, "", nil)

	err := zepClient.CheckServer()

	assert.NotNil(t, err)
	assert.Equal(t, ServerErrorMessage, err.Error())
}
