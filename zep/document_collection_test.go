package zep

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocumentCollection_Status(t *testing.T) {
	zepClient := NewClient("", "", nil)

	t.Run("Returns ready when document count and embedded count match", func(t *testing.T) {
		documentCollectionModel := DocumentCollectionModel{
			Name:                  "testCollection",
			DocumentCount:         1,
			DocumentEmbeddedCount: 1,
		}
		documentCollection := NewDocumentCollection(zepClient, documentCollectionModel)

		assert.Equal(t, "ready", documentCollection.Status())
	})

	t.Run("Returns pending when document count and embedded count do not match", func(t *testing.T) {
		documentCollectionModel := DocumentCollectionModel{
			Name:                  "testCollection",
			DocumentCount:         1,
			DocumentEmbeddedCount: 0,
		}
		documentCollection := NewDocumentCollection(zepClient, documentCollectionModel)

		assert.Equal(t, "pending", documentCollection.Status())
	})
}

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
	// I can easily refactor these tests if needed.
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

	t.Run("Handles error on upload", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		documents := []Document{
			{
				Content: "test content",
			},
		}
		uuids, err := documentCollection.AddDocuments(documents)
		assert.NotNil(t, err)
		assert.Nil(t, uuids)
	})
}

func TestDocumentCollection_UpdateDocument(t *testing.T) {
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

	t.Run("Updates one document", func(t *testing.T) {
		testUUID := "testUUID"
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}

		updateDocumentParams := UpdateDocumentParams{
			UUID: testUUID,
		}

		err := documentCollection.UpdateDocument(updateDocumentParams)
		assert.Nil(t, err)
	})

	t.Run("Handles error on update", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		updateDocumentParams := UpdateDocumentParams{
			UUID: "testUUID",
		}

		err := documentCollection.UpdateDocument(updateDocumentParams)
		assert.NotNil(t, err)
		assert.IsType(t, &APIError{}, err)
	})
}

func TestDocumentCollection_DeleteDocument(t *testing.T) {
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

	t.Run("Deletes one document", func(t *testing.T) {
		testUUID := "testUUID"
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}

		err := documentCollection.DeleteDocument(testUUID)
		assert.Nil(t, err)
	})

	t.Run("Handles error on delete", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		err := documentCollection.DeleteDocument("testUUID")
		assert.NotNil(t, err)
		assert.IsType(t, &APIError{}, err)
	})
}

func TestDocumentCollection_GetDocument(t *testing.T) {
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

	t.Run("Gets one document", func(t *testing.T) {
		testUUID := "testUUID"
		testDocument := Document{
			UUID:    testUUID,
			Content: "test content",
		}
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(testDocument)
		}

		document, err := documentCollection.GetDocument(testUUID)
		assert.Nil(t, err)
		assert.Equal(t, testDocument, *document)
	})

	t.Run("Handles error on get", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		document, err := documentCollection.GetDocument("testUUID")
		assert.NotNil(t, err)
		assert.IsType(t, &APIError{}, err)
		assert.Nil(t, document)
	})
}

func TestDocumentCollection_GetDocuments(t *testing.T) {
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

	t.Run("Gets one document", func(t *testing.T) {
		testUUIDs := []string{"testUUID", "testUUID2"}
		testDocuments := []Document{
			{
				UUID:    testUUIDs[0],
				Content: "test content",
			},
			{
				UUID:    testUUIDs[1],
				Content: "test content",
			},
		}
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(testDocuments)
		}

		documents, err := documentCollection.GetDocuments(testUUIDs)
		assert.Nil(t, err)
		assert.Equal(t, testDocuments, documents)
	})

	t.Run("Handles error on get", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		documents, err := documentCollection.GetDocuments([]string{"testUUID"})
		assert.NotNil(t, err)
		assert.IsType(t, &APIError{}, err)
		assert.Nil(t, documents)
	})
}

func TestDocumentCollection_SearchReturnQueryVector(t *testing.T) {
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

	t.Run("Gets one document", func(t *testing.T) {
		testUUIDs := []string{"testUUID"}
		testDocuments := []Document{
			{
				UUID:    testUUIDs[0],
				Content: "test content",
			},
		}

		queryVector := []float32{0.1, -0.2, 0.3}
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			response := struct {
				Results     []Document `json:"results"`
				QueryVector []float32  `json:"query_vector"`
			}{
				Results:     testDocuments,
				QueryVector: queryVector,
			}
			json.NewEncoder(rw).Encode(response)
		}

		query := SearchQuery{
			Embedding: &[]float32{0.1, 0.3},
		}

		limit := 1
		documents, queryVector, err := documentCollection.SearchReturnQueryVector(query, &limit)

		assert.Nil(t, err)
		assert.Equal(t, testDocuments, documents)
		assert.Equal(t, queryVector, queryVector)
	})

	t.Run("Handles error on search", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		query := SearchQuery{
			Embedding: &[]float32{0.1, 0.3},
		}

		limit := 1
		documents, queryVector, err := documentCollection.SearchReturnQueryVector(query, &limit)

		assert.NotNil(t, err)
		assert.IsType(t, &APIError{}, err)
		assert.Nil(t, documents)
		assert.Nil(t, queryVector)
	})
}

func TestDocumentCollection_Search(t *testing.T) {
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

	t.Run("Gets one document", func(t *testing.T) {
		testUUIDs := []string{"testUUID"}
		testDocuments := []Document{
			{
				UUID:    testUUIDs[0],
				Content: "test content",
			},
		}

		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			response := struct {
				Results     []Document `json:"results"`
				QueryVector []float32  `json:"query_vector"`
			}{
				Results: testDocuments,
			}
			json.NewEncoder(rw).Encode(response)
		}

		query := SearchQuery{
			Embedding: &[]float32{0.1, 0.3},
		}

		limit := 1
		documents, err := documentCollection.Search(query, &limit)

		assert.Nil(t, err)
		assert.Equal(t, testDocuments, documents)
	})

	t.Run("Handles error on search", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		query := SearchQuery{
			Embedding: &[]float32{0.1, 0.3},
		}

		limit := 1
		documents, err := documentCollection.Search(query, &limit)

		assert.NotNil(t, err)
		assert.IsType(t, &APIError{}, err)
		assert.Nil(t, documents)
	})
}

func TestDocumentCollection_CreateIndex(t *testing.T) {
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

	t.Run("Creates index", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}

		err := documentCollection.CreateIndex(true)
		assert.Nil(t, err)
	})

	t.Run("Handles error on create index", func(t *testing.T) {
		*responseFunc = func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		err := documentCollection.CreateIndex(false)
		assert.NotNil(t, err)
		assert.IsType(t, &APIError{}, err)
	})
}
