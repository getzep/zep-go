package zep

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

const DefaultBatchSize = 500
const MaxConcurrentBatches = 5

type DocumentCollection struct {
	DocumentCollectionModel
	Client Client `json:"-"`
}

func NewDocumentCollection(client Client, params DocumentCollectionModel) *DocumentCollection {
	return &DocumentCollection{DocumentCollectionModel: params, Client: client}
}

func (d *DocumentCollection) Status() string {
	if d.DocumentCount != 0 && d.DocumentEmbeddedCount != 0 && d.DocumentCount == d.DocumentEmbeddedCount {
		return "ready"
	}
	return "pending"
}

func (d *DocumentCollection) AddDocuments(documents []Document) ([]string, error) {
	if d.Name == "" {
		return nil, errors.New("collection name must be provided")
	}
	if len(documents) == 0 {
		return nil, errors.New("no documents provided")
	}

	var wg sync.WaitGroup
	uuids := make(chan string)
	errSem := make(chan bool, 1)
	errs := make(chan error, 1)
	// Limit concurrent batches.
	sem := make(chan bool, MaxConcurrentBatches)

	for i := 0; i < len(documents); i += DefaultBatchSize {
		// Use a semaphore to limit concurrent batches to MaxConcurrentBatches.
		sem <- true
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			batch := documents[i:min(i+DefaultBatchSize, len(documents))]
			batchUUIDs, err := d.uploadBatch(batch)
			if err != nil {
				errSem <- true
				// Only return the first error.
				if len(errs) == 0 {
					errs <- err
				}
				<-errSem
				return
			}
			for _, uuid := range batchUUIDs {
				uuids <- uuid
			}
			<-sem
		}(i)
	}

	// Wait for all batches to finish.
	go func() {
		wg.Wait()
		close(uuids)
		close(errs)
	}()

	result := make([]string, 0, len(documents))
	for uuid := range uuids {
		result = append(result, uuid)
	}

	if len(errs) > 0 {
		return nil, <-errs
	}

	return result, nil
}

func (d *DocumentCollection) uploadBatch(batch []Document) ([]string, error) {
	batchJSON, err := json.Marshal(batch)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", d.Client.GetFullURL("/collection/"+d.Name+"/document"), bytes.NewBuffer(batchJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := d.Client.HandleRequest(request, fmt.Sprintf("Failed to add documents to collection %s", d.Name))
	if err != nil {
		return nil, err
	}

	var responseData []string
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (d *DocumentCollection) UpdateDocument(params UpdateDocumentParams) error {
	if d.Name == "" {
		return errors.New("collection name must be provided")
	}
	if params.UUID == "" {
		return errors.New("document must have a uuid")
	}

	documentJSON, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("PATCH", d.Client.GetFullURL("/collection/"+d.Name+"/document/"+params.UUID), bytes.NewBuffer(documentJSON))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	_, err = d.Client.HandleRequest(request, fmt.Sprintf("Failed to update document %s", params.UUID))
	if err != nil {
		return err
	}

	return nil
}

func (d *DocumentCollection) DeleteDocument(uuid string) error {
	if d.Name == "" {
		return errors.New("collection name must be provided")
	}
	if uuid == "" {
		return errors.New("document must have a uuid")
	}

	request, err := http.NewRequest("DELETE", d.Client.GetFullURL("/collection/"+d.Name+"/document/uuid/"+uuid), nil)
	if err != nil {
		return err
	}

	_, err = d.Client.HandleRequest(request, fmt.Sprintf("Failed to delete document %s", uuid))
	if err != nil {
		return err
	}

	return nil
}

func (d *DocumentCollection) GetDocument(uuid string) (*Document, error) {
	if d.Name == "" {
		return nil, errors.New("collection name must be provided")
	}
	if uuid == "" {
		return nil, errors.New("document must have a uuid")
	}

	request, err := http.NewRequest("GET", d.Client.GetFullURL("/collection/"+d.Name+"/document/"+uuid), nil)
	if err != nil {
		return nil, err
	}

	response, err := d.Client.HandleRequest(request, fmt.Sprintf("Failed to get document %s", uuid))
	if err != nil {
		return nil, err
	}

	var document Document
	err = json.NewDecoder(response.Body).Decode(&document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (d *DocumentCollection) GetDocuments(uuids []string) ([]Document, error) {
	if len(uuids) == 0 {
		return nil, errors.New("no uuids provided")
	}
	if d.Name == "" {
		return nil, errors.New("collection name must be provided")
	}

	uuidsJSON, err := json.Marshal(uuids)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", d.Client.GetFullURL("/collection/"+d.Name+"/document/list/get"), bytes.NewBuffer(uuidsJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := d.Client.HandleRequest(request, "failed to get documents")
	if err != nil {
		return nil, err
	}

	var documents []Document
	err = json.NewDecoder(response.Body).Decode(&documents)
	if err != nil {
		return nil, err
	}

	return documents, nil
}
func (d *DocumentCollection) SearchReturnQueryVector(query SearchQuery, limit *int) ([]Document, []float32, error) {
	if d.Name == "" {
		return nil, nil, errors.New("collection name must be provided")
	}
	if query.Text == nil && query.Embedding == nil && query.Metadata == nil {
		return nil, nil, errors.New("search query must have at least one of text, embedding, or metadata")
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, nil, err
	}

	url := d.Client.GetFullURL("/collection/" + d.Name + "/search")
	if limit != nil {
		url += "?limit=" + strconv.Itoa(*limit)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(queryJSON))
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := d.Client.HandleRequest(request, "Failed to search collection")
	if err != nil {
		return nil, nil, err
	}

	var results struct {
		Results     []Document `json:"results"`
		QueryVector []float32  `json:"query_vector"`
	}
	err = json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		return nil, nil, err
	}

	return results.Results, results.QueryVector, nil
}

func (d *DocumentCollection) Search(query SearchQuery, limit *int) ([]Document, error) {
	documents, _, err := d.SearchReturnQueryVector(query, limit)
	return documents, err
}

func (d *DocumentCollection) CreateIndex(force bool) error {
	if d.Name == "" {
		return errors.New("collection name must be provided")
	}

	url := d.Client.GetFullURL("/collection/" + d.Name + "/index/create")
	if force {
		url += "?force=true"
	}

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	_, err = d.Client.HandleRequest(request, "Failed to create index")
	return err
}

type SearchQuery struct {
	Text       *string
	Metadata   *map[string]any
	Embedding  *[]float32
	SearchType *string
	MmrLambda  *float64
}

type UpdateDocumentParams struct {
	UUID       string
	DocumentID string
	Metadata   *map[string]any
}
