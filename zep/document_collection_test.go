package zep

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocumentCollection_AddDocuments(t *testing.T) {
	var responseFunc *http.HandlerFunc // using a pointer, so each test can change the behaviour of the server
	responseFunc = new(http.HandlerFunc)
	*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK) // default response to pass health check.
	}
	server := httptest.NewServer(responseFunc)
	defer server.Close()

	client := &http.Client{}
	zepClient := NewClient(server.URL, "", client)
	documentCollectionModel := DocumentCollectionModel{
		Name: "testCollection",
	}
	documentCollection := NewDocumentCollection(zepClient, documentCollectionModel)

	// TODO: Review test syntax with Daniel. This is a deviation from the other tests, so I want to make sure it's still acceptable.
	t.Run("Adds one document", func(t *testing.T) {
		testUUIDS := []string{"testUUID"}
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(testUUIDS)
		}
		documents := []Document{
			{
				Content: "test content",
			},
		}
		uuids, err := documentCollection.AddDocuments(documents)
		assert.Nil(t, err)

		assert.Equal(t, testUUIDS, uuids)
	})

	t.Run("Adds multiple documents", func(t *testing.T) {
		testUUIDS := []string{"testUUID1", "testUUID2"}
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(testUUIDS)
		}
		documents := []Document{
			{
				Content: "test content",
			},
			{
				Content: "test content",
			},
		}
		uuids, err := documentCollection.AddDocuments(documents)
		assert.Nil(t, err)
		assert.Equal(t, testUUIDS, uuids)
	})
}
