package zep

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DocumentManager struct {
	Client Client
}

// NewDocumentManager creates a new DocumentManager
func NewDocumentManager(client Client) *DocumentManager {
	return &DocumentManager{Client: client}
}

func (d *DocumentManager) AddCollection(params AddCollectionParams) (*DocumentCollection, error) {
	if params.EmbeddingDimensions <= 0 {
		return nil, errors.New("embeddingDimensions must be a positive integer")
	}

	collection := DocumentCollection{
		Client: d.Client,
		DocumentCollectionModel: DocumentCollectionModel{
			Name:                params.Name,
			Description:         params.Description,
			Metadata:            params.Metadata,
			EmbeddingDimensions: params.EmbeddingDimensions,
			IsAutoEmbedded:      params.IsAutoEmbedded,
		},
	}

	collectionJSON, err := json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", d.Client.GetFullURL("/collection/"+params.Name), bytes.NewBuffer(collectionJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	_, err = d.Client.HandleRequest(request, fmt.Sprintf("Failed to add collection %s", params.Name))
	if err != nil {
		return nil, err
	}

	return d.GetCollection(collection.Name)
}

func (d *DocumentManager) GetCollection(name string) (*DocumentCollection, error) {
	if name == "" {
		return nil, errors.New("collection name must be provided")
	}

	request, err := http.NewRequest("GET", d.Client.GetFullURL("/collection/"+name), nil)
	if err != nil {
		return nil, err
	}

	response, err := d.Client.HandleRequest(request, fmt.Sprintf("No collection found for name %s", name))
	if err != nil {
		return nil, err
	}

	var responseData DocumentCollectionModel
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return &DocumentCollection{
		Client:                  d.Client,
		DocumentCollectionModel: responseData,
	}, nil
}

func (d *DocumentManager) UpdateCollection(params UpdateCollectionParams) (*DocumentCollection, error) {
	if params.Description == "" && len(params.Metadata) == 0 {
		return nil, errors.New("either description or metadata must be provided")
	}

	collection := DocumentCollectionModel{
		Name:        params.Name,
		Description: params.Description,
		Metadata:    params.Metadata,
	}

	collectionJSON, err := json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", d.Client.GetFullURL("/collection/"+collection.Name), bytes.NewBuffer(collectionJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	_, err = d.Client.HandleRequest(request, fmt.Sprintf("Failed to update collection %s", collection.Name))
	if err != nil {
		return nil, err
	}

	return d.GetCollection(collection.Name)
}

func (d *DocumentManager) ListCollections() ([]DocumentCollection, error) {
	request, err := http.NewRequest("GET", d.Client.GetFullURL("/collection"), nil)
	if err != nil {
		return nil, err
	}

	response, err := d.Client.HandleRequest(request, "Failed to list collections")
	if err != nil {
		return nil, err
	}

	var responseData []DocumentCollectionModel
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	collections := make([]DocumentCollection, len(responseData))
	for i, collection := range responseData {
		collections[i] = DocumentCollection{
			Client:                  d.Client,
			DocumentCollectionModel: collection,
		}
	}

	return collections, nil
}

func (d *DocumentManager) DeleteCollection(collectionName string) error {
	if collectionName == "" {
		return errors.New("collection name must be provided")
	}

	request, err := http.NewRequest("DELETE", d.Client.GetFullURL("/collection/"+collectionName), nil)
	if err != nil {
		return err
	}

	_, err = d.Client.HandleRequest(request, fmt.Sprintf("Failed to delete collection %s", collectionName))
	if err != nil {
		return err
	}

	return nil
}

type AddCollectionParams struct {
	Name                string
	EmbeddingDimensions int
	Description         string
	Metadata            map[string]any
	IsAutoEmbedded      bool
}

type UpdateCollectionParams struct {
	Name        string
	Description string
	Metadata    map[string]any
}
