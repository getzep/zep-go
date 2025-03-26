package main

import (
	"context"
	"fmt"
	"os"

	"github.com/getzep/zep-go/v2"
	zepclient "github.com/getzep/zep-go/v2/client"
	"github.com/getzep/zep-go/v2/option"
)

func entityTypes() {
	apiKey := os.Getenv("ZEP_API_KEY")
	if apiKey == "" {
		fmt.Println("ZEP_API_KEY environment variable is not set")
		return
	}

	client := zepclient.NewClient(
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	type Purchase struct {
		zep.BaseEntity  `name:"Purchase" description:"A purchase is an item that was purchased with base entity"`
		ItemName        string  `description:"The name of the item purchased" json:"item_name,omitempty"`
		ItemPrice       float64 `description:"The price of the item" json:"item_price,omitempty"`
		ItemQuantity    int     `description:"The quantity of the item purchased" json:"item_quantity,omitempty"`
		AdditionalNotes string  `description:"Additional notes about the purchase" json:"additional_notes,omitempty"`
	}

	_, err := client.Graph.SetEntityTypes(
		ctx,
		[]zep.EntityDefinition{
			Purchase{},
		},
	)
	if err != nil {
		fmt.Printf("Error setting entity types with base entity: %v\n", err)
		return
	}

	searchFilters := zep.SearchFilters{NodeLabels: []string{"Purchase"}}
	searchResults, err := client.Graph.Search(
		ctx,
		&zep.GraphSearchQuery{
			UserID:        zep.String("<user_id>"),
			Query:         "tickets",
			Scope:         zep.GraphSearchScopeNodes.Ptr(),
			SearchFilters: &searchFilters,
		},
	)
	if err != nil {
		fmt.Printf("Error searching graph: %v\n", err)
		return
	}

	var purchases []Purchase
	for _, node := range searchResults.Nodes {
		var purchase Purchase
		err := zep.UnmarshalNodeAttributes(node.Attributes, &purchase)
		if err != nil {
			fmt.Printf("Error converting node to struct: %v\n", err)
			continue
		}

		purchases = append(purchases, purchase)
	}

	for _, purchase := range purchases {
		fmt.Printf("Item Name: %s\n", purchase.ItemName)
		fmt.Printf("Item Price: %f\n", purchase.ItemPrice)
		fmt.Printf("Item Quantity: %d\n", purchase.ItemQuantity)
		fmt.Printf("Additional Notes: %s\n", purchase.AdditionalNotes)
	}
}
