# Reference
## Batch
<details><summary><code>client.Batch.List() -> *zep.BatchPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.BatchListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Status: zep.String(
        "status",
    ),
}
client.Batch.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Batch status filter
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.Create(request) -> *zep.Batch</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateBatchRequest{}
client.Batch.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ignoreRoles:** `[]string` 

Message roles to skip during graph extraction for thread message items in
this batch.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to store on the batch.
    
</dd>
</dl>

<dl>
<dd>

**strictOntology:** `*bool` 

When true, prevents extraction of generic entity nodes that do not match
the configured ontology for episodes in this batch.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.Get(BatchUUID) -> *zep.Batch</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Batch.Get(
    context.TODO(),
    "batch_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchUUID:** `string` — Batch UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.Delete(BatchUUID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Batch.Delete(
    context.TODO(),
    "batch_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchUUID:** `string` — Batch UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.ListItems(BatchUUID) -> *zep.BatchItemPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.BatchListItemsRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.Batch.ListItems(
    context.TODO(),
    "batch_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchUUID:** `string` — Batch UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.AddItems(BatchUUID, request) -> *zep.BatchItemsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.AddBatchItemsRequest{
    Items: []*zep.BatchItemInput{
        &zep.BatchItemInput{
            Type: zep.V4BatchItemInputTypeGraphEpisode,
        },
    },
}
client.Batch.AddItems(
    context.TODO(),
    "batch_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchUUID:** `string` — Batch UUID
    
</dd>
</dl>

<dl>
<dd>

**items:** `[]*zep.BatchItemInput` — The batch items to append, each identified by its type field.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.Process(BatchUUID) -> *zep.ProcessBatchResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Batch.Process(
    context.TODO(),
    "batch_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchUUID:** `string` — Batch UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Context
<details><summary><code>client.Context.CreateTemplate(request) -> *zep.ContextTemplate</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateContextTemplateRequest{}
client.Context.CreateTemplate(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.CreateContextTemplateRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Context.ListTemplates(request) -> *zep.ContextTemplatePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.ContextTemplateListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.Context.ListTemplates(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Filters results to the context template with this exact name.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Context.GetTemplate(TemplateUUID) -> *zep.ContextTemplate</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Context.GetTemplate(
    context.TODO(),
    "template_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateUUID:** `string` — Template UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Context.UpdateTemplate(TemplateUUID, request) -> *zep.ContextTemplate</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateContextTemplateRequest{}
client.Context.UpdateTemplate(
    context.TODO(),
    "template_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateUUID:** `string` — Template UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.CreateContextTemplateRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Context.DeleteTemplate(TemplateUUID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Context.DeleteTemplate(
    context.TODO(),
    "template_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateUUID:** `string` — Template UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Graph
<details><summary><code>client.Graph.Create(request) -> *zep.Graph</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateGraphRequest{}
client.Graph.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**description:** `*string` — A description of the graph.
    
</dd>
</dl>

<dl>
<dd>

**graphID:** `*string` — An optional developer-assigned identifier for the graph.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — A display name for the graph.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — The graph's IANA time zone.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.List(request) -> *zep.GraphPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GraphListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    OrderBy: zep.String(
        "order_by",
    ),
    Order: zep.String(
        "order",
    ),
}
client.Graph.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — asc or desc
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 

Filters results to graphs whose name, description, or graph ID contains
this text.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Lookup(request) -> *zep.Graph</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.LookupRequest{}
client.Graph.Lookup(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.LookupRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Get(GraphUUID) -> *zep.Graph</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Get(
    context.TODO(),
    "graph_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Delete(GraphUUID) -> *zep.GraphDeleteResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Delete(
    context.TODO(),
    "graph_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Update(GraphUUID, request) -> *zep.Graph</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.PatchGraphRequest{}
client.Graph.Update(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A description of the graph.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The graph's display name.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — The graph's IANA time zone.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Clone(GraphUUID, request) -> *zep.CloneGraphResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CloneGraphRequest{}
client.Graph.Clone(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**targetGraphID:** `*string` — An optional name for the cloned graph.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.GetContext(GraphUUID, request) -> *zep.GraphContextResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GraphContextRequest{
    Query: "query",
}
client.Graph.GetContext(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.SearchFilters` 

Filters constraining which graph data can be selected for the context
block.
    
</dd>
</dl>

<dl>
<dd>

**includeResults:** `*bool` — When true, includes the raw graph results selected for the context block.
    
</dd>
</dl>

<dl>
<dd>

**maxCharacters:** `*int` — The maximum number of characters in the assembled context block.
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query used to assemble the context block.
    
</dd>
</dl>

<dl>
<dd>

**recencyBias:** `*zep.V4GraphContextRequestRecencyBias` — Adjusts result selection to favor more recent graph data.
    
</dd>
</dl>

<dl>
<dd>

**templateUUID:** `*string` — The UUID of a context template used to render the context block.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.GetInstructions(GraphUUID) -> *zep.Instructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.GetInstructions(
    context.TODO(),
    "graph_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SetInstructions(GraphUUID, request) -> *zep.Instructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.Instructions{}
client.Graph.SetInstructions(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.Instructions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.GetObservationSteering(GraphUUID) -> *zep.ObservationSteering</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.GetObservationSteering(
    context.TODO(),
    "graph_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SetObservationSteering(GraphUUID, request) -> *zep.ObservationSteering</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.ObservationSteering{}
client.Graph.SetObservationSteering(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.ObservationSteering` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.GetOntology(GraphUUID) -> *zep.Ontology</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.GetOntology(
    context.TODO(),
    "graph_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SetOntology(GraphUUID, request) -> *zep.Ontology</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.Ontology{}
client.Graph.SetOntology(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.Ontology` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SearchEdges(GraphUUID, request) -> *zep.EdgePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GraphSearchEdgesRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchRequest{
        Query: "query",
    },
}
client.Graph.SearchEdges(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SearchEpisodes(GraphUUID, request) -> *zep.EpisodePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GraphSearchEpisodesRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchRequest{
        Query: "query",
    },
}
client.Graph.SearchEpisodes(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SearchNodes(GraphUUID, request) -> *zep.NodePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GraphSearchNodesRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchRequest{
        Query: "query",
    },
}
client.Graph.SearchNodes(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SearchObservations(GraphUUID, request) -> *zep.ObservationPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GraphSearchObservationsRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchRequest{
        Query: "query",
    },
}
client.Graph.SearchObservations(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.SearchThreadSummaries(GraphUUID, request) -> *zep.ThreadSummaryPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GraphSearchThreadSummariesRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchRequest{
        Query: "query",
    },
}
client.Graph.SearchThreadSummaries(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.GetSubgraph(GraphUUID, request) -> *zep.SubgraphResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.SubgraphRequest{
    SeedNodeUUIDs: []string{
        "seed_node_uuids",
    },
}
client.Graph.GetSubgraph(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**depth:** `*int` — The maximum traversal depth from the seed nodes. Defaults to 1.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*zep.V4SubgraphRequestDirection` 

The edge orientation to follow during expansion: in, out, or both.
Defaults to both.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.SearchFilters` — Filters constraining the traversed edges and included nodes.
    
</dd>
</dl>

<dl>
<dd>

**maxEdges:** `*int` — The maximum number of edges in the response. Defaults to 200.
    
</dd>
</dl>

<dl>
<dd>

**maxNodes:** `*int` — The maximum number of nodes in the response. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**seedNodeUUIDs:** `[]string` — The node UUIDs to expand from, in traversal-priority order.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Warm(GraphUUID) -> *zep.AsyncResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Warm(
    context.TODO(),
    "graph_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Lookup
<details><summary><code>client.Lookup.Batch(request) -> *zep.LookupBatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.BatchLookupRequest{}
client.Lookup.Batch(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphs:** `[]string` — Developer-assigned graph IDs to resolve to UUIDs.
    
</dd>
</dl>

<dl>
<dd>

**threads:** `[]string` — Developer-assigned thread IDs to resolve to UUIDs.
    
</dd>
</dl>

<dl>
<dd>

**users:** `[]string` — Developer-assigned user IDs to resolve to UUIDs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Project
<details><summary><code>client.Project.Get() -> *zep.Project</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Project.Get(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.Update(request) -> *zep.Project</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.PatchProjectRequest{}
client.Project.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**defaultTimeZone:** `*string` 

The project's IANA fallback time zone. Set to null to clear the existing
value.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.GetInstructions() -> *zep.Instructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Project.GetInstructions(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.SetInstructions(request) -> *zep.Instructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.Instructions{}
client.Project.SetInstructions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.Instructions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.GetObservationSteering() -> *zep.ObservationSteering</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Project.GetObservationSteering(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.SetObservationSteering(request) -> *zep.ObservationSteering</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.ObservationSteering{}
client.Project.SetObservationSteering(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.ObservationSteering` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.GetOntology() -> *zep.Ontology</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Project.GetOntology(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.SetOntology(request) -> *zep.Ontology</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.Ontology{}
client.Project.SetOntology(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.Ontology` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.GetUserSummaryInstructions() -> *zep.UserSummaryInstructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Project.GetUserSummaryInstructions(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.SetUserSummaryInstructions(request) -> *zep.UserSummaryInstructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.UserSummaryInstructions{}
client.Project.SetUserSummaryInstructions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.UserSummaryInstructions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Task
<details><summary><code>client.Task.List() -> *zep.TaskPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.TaskListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.Task.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Task.Get(TaskUUID) -> *zep.Task</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Task.Get(
    context.TODO(),
    "task_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**taskUUID:** `string` — Task UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Thread
<details><summary><code>client.Thread.List() -> *zep.ThreadPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.ThreadListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    OrderBy: zep.String(
        "order_by",
    ),
    Order: zep.String(
        "order",
    ),
    UserUUID: zep.String(
        "user_uuid",
    ),
}
client.Thread.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — asc or desc
    
</dd>
</dl>

<dl>
<dd>

**userUUID:** `*string` — Filter by user UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.Create(request) -> *zep.Thread</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateThreadRequest{
    UserUUID: "user_uuid",
}
client.Thread.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadID:** `*string` — An optional developer-assigned identifier for the thread.
    
</dd>
</dl>

<dl>
<dd>

**userUUID:** `string` — The UUID of the user this thread belongs to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.Lookup(request) -> *zep.Thread</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.LookupRequest{}
client.Thread.Lookup(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.LookupRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.Get(ThreadUUID) -> *zep.Thread</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Thread.Get(
    context.TODO(),
    "thread_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.Delete(ThreadUUID) -> *zep.ThreadDeleteResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Thread.Delete(
    context.TODO(),
    "thread_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.GetContext(ThreadUUID) -> *zep.ThreadContextResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.ThreadGetContextRequest{
    TemplateUUID: zep.String(
        "template_uuid",
    ),
}
client.Thread.GetContext(
    context.TODO(),
    "thread_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>

<dl>
<dd>

**templateUUID:** `*string` — Context template UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.ListEpisodes(ThreadUUID) -> *zep.EpisodePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.ThreadListEpisodesRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.Thread.ListEpisodes(
    context.TODO(),
    "thread_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.ListMessages(ThreadUUID) -> *zep.MessagePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.ThreadListMessagesRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.Thread.ListMessages(
    context.TODO(),
    "thread_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.AddMessages(ThreadUUID, request) -> *zep.AddMessagesResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.AddMessagesRequest{
    Messages: []*zep.AddMessage{
        &zep.AddMessage{},
    },
}
client.Thread.AddMessages(
    context.TODO(),
    "thread_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>

<dl>
<dd>

**ignoreRoles:** `[]string` 

Message roles to skip during graph extraction; the messages are still
stored.
    
</dd>
</dl>

<dl>
<dd>

**messages:** `[]*zep.AddMessage` — The messages to add to the thread.
    
</dd>
</dl>

<dl>
<dd>

**returnContext:** `*bool` 

When true, returns the context block for the thread's most recent
messages.
    
</dd>
</dl>

<dl>
<dd>

**strictOntology:** `*bool` 

When true, prevents extraction of generic entity nodes that do not match
the configured ontology.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.GetSummary(ThreadUUID) -> *zep.ThreadSummary</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Thread.GetSummary(
    context.TODO(),
    "thread_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## UserGroup
<details><summary><code>client.UserGroup.Create(request) -> *zep.UserGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateUserGroupRequest{
    Name: "name",
}
client.UserGroup.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**description:** `*string` — A description of the user group.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the user group.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.List(request) -> *zep.UserGroupPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.UserGroupListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchListRequest{},
}
client.UserGroup.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.Get(GroupUUID) -> *zep.UserGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserGroup.Get(
    context.TODO(),
    "group_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.Delete(GroupUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserGroup.Delete(
    context.TODO(),
    "group_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.Update(GroupUUID, request) -> *zep.UserGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.PatchUserGroupRequest{}
client.UserGroup.Update(
    context.TODO(),
    "group_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A description of the user group.
    
</dd>
</dl>

<dl>
<dd>

**expectedVersion:** `*int` — The user group's current version, used to detect concurrent updates.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the user group.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.ListMemberCandidates(GroupUUID, request) -> *zep.UserPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.UserGroupListMemberCandidatesRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchListRequest{},
}
client.UserGroup.ListMemberCandidates(
    context.TODO(),
    "group_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.AddMembers(GroupUUID, request) -> *zep.MembershipMutationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.MutateMembersRequest{
    UserUUIDs: []string{
        "user_uuids",
    },
}
client.UserGroup.AddMembers(
    context.TODO(),
    "group_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.MutateMembersRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.ListMembers(GroupUUID, request) -> *zep.UserPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.UserGroupListMembersRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.SearchListRequest{},
}
client.UserGroup.ListMembers(
    context.TODO(),
    "group_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.SearchListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.RemoveMembers(GroupUUID, request) -> *zep.MembershipMutationResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.MutateMembersRequest{
    UserUUIDs: []string{
        "user_uuids",
    },
}
client.UserGroup.RemoveMembers(
    context.TODO(),
    "group_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.MutateMembersRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.RemoveMember(GroupUUID, UserUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserGroup.RemoveMember(
    context.TODO(),
    "group_uuid",
    "user_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**groupUUID:** `string` — User group UUID
    
</dd>
</dl>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGroup.ListForUser(UserUUID) -> *zep.UserGroupPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires a project API key, or an account-admin bearer token with the X-Zep-Project header. The account must be entitled to attribute-based access control.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.UserGroupListForUserRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.UserGroup.ListForUser(
    context.TODO(),
    "user_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## User
<details><summary><code>client.User.Create(request) -> *zep.User</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateUserRequest{}
client.User.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disableDefaultOntology:** `*bool` — When true, disables the default ontology for the user's graph.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — The email address of the user.
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `*string` — The user's first name.
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `*string` — The user's last name.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to store on the user.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — The user's IANA time zone.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — An optional developer-assigned identifier for the user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.List(request) -> *zep.UserPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.UserListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    OrderBy: zep.String(
        "order_by",
    ),
    Order: zep.String(
        "order",
    ),
}
client.User.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*string` — asc or desc
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Filters results to users whose user ID, email, or name contains this text.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Lookup(request) -> *zep.User</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.LookupRequest{}
client.User.Lookup(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*zep.LookupRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Get(UserUUID) -> *zep.User</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.Get(
    context.TODO(),
    "user_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Delete(UserUUID) -> *zep.UserDeleteResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.Delete(
    context.TODO(),
    "user_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.Update(UserUUID, request) -> *zep.User</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.PatchUserRequest{}
client.User.Update(
    context.TODO(),
    "user_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>

<dl>
<dd>

**disableDefaultOntology:** `*bool` — When true, disables the default ontology for the user's graph.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — The email address of the user.
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `*string` — The user's first name.
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `*string` — The user's last name.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to merge onto the user; a key set to null is removed.
    
</dd>
</dl>

<dl>
<dd>

**timeZone:** `*string` — The user's IANA time zone.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetNode(UserUUID) -> *zep.Node</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetNode(
    context.TODO(),
    "user_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.GetSummaryInstructions(UserUUID) -> *zep.UserSummaryInstructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.User.GetSummaryInstructions(
    context.TODO(),
    "user_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.User.SetSummaryInstructions(UserUUID, request) -> *zep.UserSummaryInstructions</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.UserSummaryInstructions{}
client.User.SetSummaryInstructions(
    context.TODO(),
    "user_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userUUID:** `string` — User UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.UserSummaryInstructions` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Graph DocumentSummary
<details><summary><code>client.Graph.DocumentSummary.List(GraphUUID, request) -> *zep.DocumentSummaryPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.DocumentSummaryListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.ArtifactListRequest{},
}
client.Graph.DocumentSummary.List(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.ArtifactListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Graph Episode
<details><summary><code>client.Graph.Episode.ListForDocument(GraphUUID, DocumentID) -> *zep.EpisodePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.EpisodeListForDocumentRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.Graph.Episode.ListForDocument(
    context.TODO(),
    "graph_uuid",
    "document_id",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**documentID:** `string` — Document ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Episode.Add(GraphUUID, request) -> *zep.AddEpisodeResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.AddEpisodeRequest{
    Data: "data",
}
client.Graph.Episode.Add(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**createdAt:** `*string` 

The episode's reference time, used for temporal reasoning rather than
ingestion time.
    
</dd>
</dl>

<dl>
<dd>

**data:** `string` — The episode content to add to the graph.
    
</dd>
</dl>

<dl>
<dd>

**documentID:** `*string` — Groups this episode as a chunk of a document on the graph.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to store on the episode.
    
</dd>
</dl>

<dl>
<dd>

**sourceDescription:** `*string` — A description of the source of this episode.
    
</dd>
</dl>

<dl>
<dd>

**strictOntology:** `*bool` 

When true, prevents extraction of generic entity nodes that do not match
the configured ontology.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*graph.V4AddEpisodeRequestType` — The data format of the episode: text, json, or message. Defaults to text.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Episode.List(GraphUUID, request) -> *zep.EpisodePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.EpisodeListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.ArtifactListRequest{},
}
client.Graph.Episode.List(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.ArtifactListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Episode.Get(GraphUUID, EpisodeUUID) -> *zep.Episode</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Episode.Get(
    context.TODO(),
    "graph_uuid",
    "episode_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**episodeUUID:** `string` — Episode UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Episode.Delete(GraphUUID, EpisodeUUID) -> *zep.AsyncResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Episode.Delete(
    context.TODO(),
    "graph_uuid",
    "episode_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**episodeUUID:** `string` — Episode UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Episode.Update(GraphUUID, EpisodeUUID, request) -> *zep.Episode</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.PatchEpisodeRequest{}
client.Graph.Episode.Update(
    context.TODO(),
    "graph_uuid",
    "episode_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**episodeUUID:** `string` — Episode UUID
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to merge onto the episode; a key set to null is removed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Graph Edge
<details><summary><code>client.Graph.Edge.Add(GraphUUID, request) -> *zep.AddEdgeResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.AddEdgeRequest{
    Fact: "fact",
    FactName: "fact_name",
    SourceNode: &zep.EdgeNodeRef{},
    TargetNode: &zep.EdgeNodeRef{},
}
client.Graph.Edge.Add(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**attributes:** `map[string]any` — Additional attributes to store on the edge.
    
</dd>
</dl>

<dl>
<dd>

**expiredAt:** `*string` — The time at which the fact was superseded or invalidated.
    
</dd>
</dl>

<dl>
<dd>

**fact:** `string` 

The fact text describing the relationship between the source and target
nodes.
    
</dd>
</dl>

<dl>
<dd>

**factName:** `string` — The name of the edge, in upper snake case, for example RELATES_TO.
    
</dd>
</dl>

<dl>
<dd>

**invalidAt:** `*string` — The time at which the fact stopped being true.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata attached to the episode created for this edge.
    
</dd>
</dl>

<dl>
<dd>

**sourceNode:** `*zep.EdgeNodeRef` 

The source node of the edge, referenced by uuid or created or matched by
name.
    
</dd>
</dl>

<dl>
<dd>

**targetNode:** `*zep.EdgeNodeRef` 

The target node of the edge, referenced by uuid or created or matched by
name.
    
</dd>
</dl>

<dl>
<dd>

**validAt:** `*string` — The time at which the fact became true.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Edge.List(GraphUUID, request) -> *zep.EdgePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.EdgeListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.ArtifactListRequest{},
}
client.Graph.Edge.List(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.ArtifactListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Edge.Get(GraphUUID, EdgeUUID) -> *zep.Edge</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Edge.Get(
    context.TODO(),
    "graph_uuid",
    "edge_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**edgeUUID:** `string` — Edge UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Edge.Delete(GraphUUID, EdgeUUID) -> *zep.AsyncResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Edge.Delete(
    context.TODO(),
    "graph_uuid",
    "edge_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**edgeUUID:** `string` — Edge UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Edge.Update(GraphUUID, EdgeUUID, request) -> *zep.Edge</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.PatchEdgeRequest{}
client.Graph.Edge.Update(
    context.TODO(),
    "graph_uuid",
    "edge_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**edgeUUID:** `string` — Edge UUID
    
</dd>
</dl>

<dl>
<dd>

**attributes:** `map[string]any` 

Additional attributes to merge onto the edge; a key set to null is
removed.
    
</dd>
</dl>

<dl>
<dd>

**fact:** `*string` 

The fact text describing the relationship between the source and target
nodes.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Graph Node
<details><summary><code>client.Graph.Node.Add(GraphUUID, request) -> *zep.AddNodesResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.AddNodesRequest{
    Nodes: []*zep.NodeInput{
        &zep.NodeInput{
            Name: "name",
        },
    },
}
client.Graph.Node.Add(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**nodes:** `[]*zep.NodeInput` — The nodes to add to the graph.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Node.List(GraphUUID, request) -> *zep.NodePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.NodeListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.ArtifactListRequest{},
}
client.Graph.Node.List(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.ArtifactListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Node.Get(GraphUUID, NodeUUID) -> *zep.Node</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Node.Get(
    context.TODO(),
    "graph_uuid",
    "node_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**nodeUUID:** `string` — Node UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Node.Delete(GraphUUID, NodeUUID) -> *zep.AsyncResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Node.Delete(
    context.TODO(),
    "graph_uuid",
    "node_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**nodeUUID:** `string` — Node UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Node.Update(GraphUUID, NodeUUID, request) -> *zep.Node</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.PatchNodeRequest{}
client.Graph.Node.Update(
    context.TODO(),
    "graph_uuid",
    "node_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**nodeUUID:** `string` — Node UUID
    
</dd>
</dl>

<dl>
<dd>

**attributes:** `map[string]any` 

Additional attributes to merge onto the node; a key set to null is
removed.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The node's name.
    
</dd>
</dl>

<dl>
<dd>

**summary:** `*string` — A summary of the node.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Node.ListNeighbors(GraphUUID, NodeUUID, request) -> *zep.NeighborPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.NeighborsRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
}
client.Graph.Node.ListNeighbors(
    context.TODO(),
    "graph_uuid",
    "node_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**nodeUUID:** `string` — Node UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*graph.V4NeighborsRequestDirection` — The edge orientation to follow from the node: in, out, or both.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.SearchFilters` — Filters constraining the connecting edges and the neighbor nodes.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Graph Observation
<details><summary><code>client.Graph.Observation.List(GraphUUID, request) -> *zep.ObservationPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.ObservationListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.ArtifactListRequest{},
}
client.Graph.Observation.List(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.ArtifactListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Observation.Get(GraphUUID, ObservationUUID) -> *zep.Observation</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Graph.Observation.Get(
    context.TODO(),
    "graph_uuid",
    "observation_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**observationUUID:** `string` — Observation UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Graph ThreadSummary
<details><summary><code>client.Graph.ThreadSummary.List(GraphUUID, request) -> *zep.ThreadSummaryPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &graph.ThreadSummaryListRequest{
    Limit: zep.Int(
        1,
    ),
    Cursor: zep.String(
        "cursor",
    ),
    Body: &zep.ArtifactListRequest{},
}
client.Graph.ThreadSummary.List(
    context.TODO(),
    "graph_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**graphUUID:** `string` — Graph UUID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Page size
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.ArtifactListRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Thread Message
<details><summary><code>client.Thread.Message.Get(ThreadUUID, MessageUUID) -> *zep.Message</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Thread.Message.Get(
    context.TODO(),
    "thread_uuid",
    "message_uuid",
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>

<dl>
<dd>

**messageUUID:** `string` — Message UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Thread.Message.Update(ThreadUUID, MessageUUID, request) -> *zep.Message</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &thread.PatchMessageRequest{}
client.Thread.Message.Update(
    context.TODO(),
    "thread_uuid",
    "message_uuid",
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**threadUUID:** `string` — Thread UUID
    
</dd>
</dl>

<dl>
<dd>

**messageUUID:** `string` — Message UUID
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to merge onto the message; a key set to null is removed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

