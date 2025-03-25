package main

import (
	"context"
	"encoding/json"
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
		ItemName        string  `description:"The name of the item purchased" json:"item_name,omitempty"`
		ItemPrice       float64 `description:"The price of the item" json:"item_price,omitempty"`
		ItemQuantity    int     `description:"The quantity of the item purchased" json:"item_quantity,omitempty"`
		AdditionalNotes string  `description:"Additional notes about the purchase" json:"additional_notes,omitempty"`
	}
	_, err := client.Graph.SetEntityTypes(
		ctx,
		map[string]zep.EntityTypeDefinition{
			"Purchase": {
				Description: "A purchase is an item that was purchased",
				Interface:   Purchase{},
			},
		},
	)
	if err != nil {
		fmt.Printf("Error setting entity types: %v\n", err)
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

	nodes := searchResults.Nodes

	var purchases []Purchase
	for _, node := range nodes {
		var purchase Purchase
		jsonData, err := json.Marshal(node.Attributes)
		if err != nil {
			fmt.Printf("Error marshaling node attributes: %v\n", err)
			continue
		}
		err = json.Unmarshal(jsonData, &purchase)
		if err != nil {
			fmt.Printf("Error converting node to struct: %v\n", err)
			continue
		}

		purchases = append(purchases, purchase)
	}

	fmt.Printf("Purchases: %v\n", purchases)
}
