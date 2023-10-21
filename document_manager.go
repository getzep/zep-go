package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DocumentManager struct {
	Client ZepClient
}

// NewDocumentManager creates a new DocumentManager
func NewDocumentManager(client ZepClient) *DocumentManager {
	return &DocumentManager{Client: client}
}

func (d *DocumentManager) AddCollection(params AddCollectionParams) (*DocumentCollectionModel, error) {
	if params.EmbeddingDimensions <= 0 {
		return nil, errors.New("embeddingDimensions must be a positive integer")
	}

	collection := DocumentCollectionModel{
		Name:                params.Name,
		Description:         params.Description,
		Metadata:            params.Metadata,
		EmbeddingDimensions: &params.EmbeddingDimensions,
		IsAutoEmbedded:      params.IsAutoEmbedded,
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

func (d *DocumentManager) GetCollection(name string) (*DocumentCollectionModel, error) {
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

	return &responseData, nil
}

func (d *DocumentManager) UpdateCollection(params UpdateCollectionParams) (*DocumentCollectionModel, error) {
	if params.Description == nil && params.Metadata == nil {
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

func (d *DocumentManager) ListCollections() ([]DocumentCollectionModel, error) {
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

	return responseData, nil
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
	Description         *string
	Metadata            *map[string]any
	IsAutoEmbedded      *bool
}

type UpdateCollectionParams struct {
	Name        string
	Description *string
	Metadata    *map[string]any
}
