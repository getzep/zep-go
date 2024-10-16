// This file was auto-generated by Fern from our API Definition.

package zep

type AddDataRequest struct {
	Data    *string        `json:"data,omitempty" url:"-"`
	GroupID *string        `json:"group_id,omitempty" url:"-"`
	Type    *GraphDataType `json:"type,omitempty" url:"-"`
	UserID  *string        `json:"user_id,omitempty" url:"-"`
}

type GraphSearchQuery struct {
	// Node to rerank around for node distance reranking
	CenterNodeUUID *string `json:"center_node_uuid,omitempty" url:"-"`
	// one of user_id or group_id must be provided
	GroupID *string `json:"group_id,omitempty" url:"-"`
	// The maximum number of facts to retrieve
	Limit *int `json:"limit,omitempty" url:"-"`
	// minimum similarity score for a result to be returned
	MinScore *float64 `json:"min_score,omitempty" url:"-"`
	// weighting for maximal marginal relevance
	MmrLambda *float64 `json:"mmr_lambda,omitempty" url:"-"`
	// The string to search for (required)
	Query string `json:"query" url:"-"`
	// Defaults to RRF
	Reranker *Reranker `json:"reranker,omitempty" url:"-"`
	// Defaults to Edges. Nodes and Communities will be added in the future.
	Scope *GraphSearchScope `json:"scope,omitempty" url:"-"`
	// one of user_id or group_id must be provided
	UserID *string `json:"user_id,omitempty" url:"-"`
}
