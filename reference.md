# Reference
## Agent
<details><summary><code>client.Agent.Create(request) -> *zep.Agent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.CreateAgentRequest{
    AgentID: "agent_id",
    Name: "name",
    SecurityDomain: "security_domain",
}
client.Agent.Create(
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

**request:** `*zep.CreateAgentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.List(request) -> *zep.AgentPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.AgentListRequest{}
client.Agent.List(
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*zep.AgentListRequestOrderBy` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*zep.AgentListRequestOrder` — asc or desc
    
</dd>
</dl>

<dl>
<dd>

**agentID:** `*string` — Filters to the exact developer-assigned Agent identifier.
    
</dd>
</dl>

<dl>
<dd>

**deploymentID:** `*string` — Filters to the exact deployment identifier.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Filters to Agents whose metadata contains these key-value pairs.
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 

Filters to Agents whose name, Agent identifier, deployment identifier,
or deployment version contains this text. Matching is case-insensitive.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*zep.AgentStatus` — Filters to lifecycle status. Omit to include every status.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — Filters to the exact deployment version.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Get(AgentUUID) -> *zep.Agent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Get(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Delete(AgentUUID, request) -> *zep.AgentDeleteResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.DeleteAgentRequest{
    ExpectedRevision: 1,
}
client.Agent.Delete(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` — The current Agent revision used for optimistic concurrency.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Update(AgentUUID, request) -> *zep.Agent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.PatchAgentRequest{
    ExpectedRevision: 1,
}
client.Agent.Update(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**deploymentID:** `*string` — The customer deployment this Agent represents. Set to null to clear it.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A human-readable description of the Agent. Set to null to clear it.
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` — The current Agent revision used for optimistic concurrency.
    
</dd>
</dl>

<dl>
<dd>

**memorySettings:** `*zep.AgentMemorySettings` — Replacement Agent Skills settings. Set to null to restore defaults.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Replacement developer metadata. Set to null to clear it.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The Agent's display name.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — The customer deployment version this Agent represents. Set to null to clear it.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.DeclareBreakingChange(AgentUUID, request) -> *zep.AgentBreakingChange</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.DeclareAgentBreakingChangeRequest{
    ExpectedRevision: 1,
    Version: "version",
}
client.Agent.DeclareBreakingChange(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.GetContext(AgentUUID, request) -> *zep.AgentContext</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.GetAgentContextRequest{}
client.Agent.GetContext(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**environments:** `[]*zep.AgentContextResource` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**maxCharacters:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**models:** `[]*zep.AgentContextResource` 
    
</dd>
</dl>

<dl>
<dd>

**objective:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tools:** `[]*zep.AgentContextResource` 
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

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
request := &zep.BatchListRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**status:** `*zep.BatchListRequestStatus` — Batch status filter
    
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
request := &zep.BatchListItemsRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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
            Type: zep.BatchItemInputTypeGraphEpisode,
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
request := &zep.ContextTemplateListRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

## DebugLog
<details><summary><code>client.DebugLog.Enable() -> *zep.DebugLoggingEnablement</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enables debug logging for the project for one hour, or for 24 hours on the Enterprise plan. Episodes ingested while debug logging is enabled have a debug log (see `graph.episode.get_debug_logs`). On the Flex, Flex Plus, and Enterprise plans, the operation also enables ingestion tracing (see `graph.episode.list_ingestion_traces`); `ingestion_trace_enabled` reports the result. A repeated call starts the duration again.
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
client.DebugLog.Enable(
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

**contentPolicy:** `*zep.GraphContentPolicyRequest` 

Content policy additions for the graph. The graph binds the current
project content policy plus these additions, and the binding does not
change after creation.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — A description of the graph.
    
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
request := &zep.GraphListRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*zep.GraphListRequestOrderBy` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*zep.GraphListRequestOrder` — asc or desc
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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
request := map[string]any{
    "key": "value",
}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**request:** `zep.CloneGraphRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.GetContentPolicy(GraphUUID) -> *zep.GraphContentPolicy</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the content policy the graph bound at creation. The policy of a graph does not change after creation. A graph without a content policy returns revision 0 with no categories and no rules.
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
client.Graph.GetContentPolicy(
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.ListContentPolicyEvents(GraphUUID, request) -> *zep.ContentPolicyEventPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the content policy decisions recorded for a graph, newest first. Each event carries identifiers only. A graph without a content policy returns an empty list.
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
request := &zep.ContentPolicyEventListRequest{}
client.Graph.ListContentPolicyEvents(
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.ContentPolicyEventListFilters` — Exact-match filters that limit the events in the list.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**recencyBias:** `*zep.GraphContextRequestRecencyBias` — Adjusts result selection to favor more recent graph data.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Edge search does not support node_labels, exclude_node_labels, or content_policy_violated. content_policy_violated is supported on episode search only. An edge has two nodes, so a node label filter is not defined for edge search. To filter edges by node label, use graph.edge.list, graph.node.list_neighbors, or graph.get_subgraph. To search the edges of nodes that have a label, do these two steps. First, call graph.search_nodes with node_labels. Then, call graph.search_edges with connected_node_uuids set to the UUIDs of the returned nodes.
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
request := &zep.SearchEdgesRequest{
    Query: "query",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages. Search ranks a limited set of results, so iteration ends after that set. To get the top N results, set `page_size` to N and read the first page.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query.
    
</dd>
</dl>

<dl>
<dd>

**bfsOriginNodeUUIDs:** `[]string` — Nodes used as BFS origins for graph-distance-aware retrieval.
    
</dd>
</dl>

<dl>
<dd>

**centerNodeUUID:** `*string` 

The node to rank results by distance from. Required when reranker is
node_distance.
    
</dd>
</dl>

<dl>
<dd>

**filterLowQualityResults:** `*bool` 

The default is false. When true, Zep removes low-quality results at the knee of the
cross-encoder score curve. The filter never removes a result with a score
of 0.9 or more. It keeps at least 5 results, or all results when fewer
than 5 are available. The filter only removes results, so a page still
returns up to page_size results. Use this field only with reranker
cross_encoder. It is not supported for episode search. Invalid
combinations return a 400 error. The cutoff is approximate, not guaranteed.
When the cross-encoder does not score the results, for example after a
timeout, Zep returns the results without the filter.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.EdgeSearchFilters` — Filters constraining which items are returned.
    
</dd>
</dl>

<dl>
<dd>

**mmrLambda:** `*float64` 

The diversity weighting used for maximal marginal relevance reranking.
Required when reranker is mmr.
    
</dd>
</dl>

<dl>
<dd>

**reranker:** `*zep.SearchEdgesRequestReranker` — The reranking strategy applied to retrieved results. Defaults to rrf.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Episode search supports only the metadata_filters and content_policy_violated filters. Every other filter returns unsupported_filter.
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
request := &zep.SearchEpisodesRequest{
    Query: "query",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages. Search ranks a limited set of results, so iteration ends after that set. To get the top N results, set `page_size` to N and read the first page.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query.
    
</dd>
</dl>

<dl>
<dd>

**bfsOriginNodeUUIDs:** `[]string` — Nodes used as BFS origins for graph-distance-aware retrieval.
    
</dd>
</dl>

<dl>
<dd>

**centerNodeUUID:** `*string` 

The node to rank results by distance from. Required when reranker is
node_distance.
    
</dd>
</dl>

<dl>
<dd>

**filterLowQualityResults:** `*bool` 

The default is false. When true, Zep removes low-quality results at the knee of the
cross-encoder score curve. The filter never removes a result with a score
of 0.9 or more. It keeps at least 5 results, or all results when fewer
than 5 are available. The filter only removes results, so a page still
returns up to page_size results. Use this field only with reranker
cross_encoder. It is not supported for episode search. Invalid
combinations return a 400 error. The cutoff is approximate, not guaranteed.
When the cross-encoder does not score the results, for example after a
timeout, Zep returns the results without the filter.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.EpisodeSearchFilters` — Filters constraining which items are returned.
    
</dd>
</dl>

<dl>
<dd>

**mmrLambda:** `*float64` 

The diversity weighting used for maximal marginal relevance reranking.
Required when reranker is mmr.
    
</dd>
</dl>

<dl>
<dd>

**reranker:** `*zep.SearchEpisodesRequestReranker` — The reranking strategy applied to retrieved results. Defaults to rrf.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Node search does not support edge_types, exclude_edge_types, edge_uuids, connected_node_uuids, source_node_uuids, target_node_uuids, date_filters, or content_policy_violated. content_policy_violated is supported on episode search only. The edge filters select edges. To get the nodes that an edge connects, use graph.search_edges or graph.node.list_neighbors. To filter nodes by created_at, use graph.node.list.
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
request := &zep.SearchNodesRequest{
    Query: "query",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages. Search ranks a limited set of results, so iteration ends after that set. To get the top N results, set `page_size` to N and read the first page.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query.
    
</dd>
</dl>

<dl>
<dd>

**bfsOriginNodeUUIDs:** `[]string` — Nodes used as BFS origins for graph-distance-aware retrieval.
    
</dd>
</dl>

<dl>
<dd>

**centerNodeUUID:** `*string` 

The node to rank results by distance from. Required when reranker is
node_distance.
    
</dd>
</dl>

<dl>
<dd>

**filterLowQualityResults:** `*bool` 

The default is false. When true, Zep removes low-quality results at the knee of the
cross-encoder score curve. The filter never removes a result with a score
of 0.9 or more. It keeps at least 5 results, or all results when fewer
than 5 are available. The filter only removes results, so a page still
returns up to page_size results. Use this field only with reranker
cross_encoder. It is not supported for episode search. Invalid
combinations return a 400 error. The cutoff is approximate, not guaranteed.
When the cross-encoder does not score the results, for example after a
timeout, Zep returns the results without the filter.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.NodeSearchFilters` — Filters constraining which items are returned.
    
</dd>
</dl>

<dl>
<dd>

**mmrLambda:** `*float64` 

The diversity weighting used for maximal marginal relevance reranking.
Required when reranker is mmr.
    
</dd>
</dl>

<dl>
<dd>

**reranker:** `*zep.SearchNodesRequestReranker` — The reranking strategy applied to retrieved results. Defaults to rrf.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Observation search supports no filters. Every filter returns unsupported_filter.
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
request := &zep.SearchObservationsRequest{
    Query: "query",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages. Search ranks a limited set of results, so iteration ends after that set. To get the top N results, set `page_size` to N and read the first page.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query.
    
</dd>
</dl>

<dl>
<dd>

**bfsOriginNodeUUIDs:** `[]string` — Nodes used as BFS origins for graph-distance-aware retrieval.
    
</dd>
</dl>

<dl>
<dd>

**centerNodeUUID:** `*string` 

The node to rank results by distance from. Required when reranker is
node_distance.
    
</dd>
</dl>

<dl>
<dd>

**filterLowQualityResults:** `*bool` 

The default is false. When true, Zep removes low-quality results at the knee of the
cross-encoder score curve. The filter never removes a result with a score
of 0.9 or more. It keeps at least 5 results, or all results when fewer
than 5 are available. The filter only removes results, so a page still
returns up to page_size results. Use this field only with reranker
cross_encoder. It is not supported for episode search. Invalid
combinations return a 400 error. The cutoff is approximate, not guaranteed.
When the cross-encoder does not score the results, for example after a
timeout, Zep returns the results without the filter.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.ObservationSearchFilters` — Filters constraining which items are returned.
    
</dd>
</dl>

<dl>
<dd>

**mmrLambda:** `*float64` 

The diversity weighting used for maximal marginal relevance reranking.
Required when reranker is mmr.
    
</dd>
</dl>

<dl>
<dd>

**reranker:** `*zep.SearchObservationsRequestReranker` — The reranking strategy applied to retrieved results. Defaults to rrf.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Thread-summary search supports only the metadata_filters filter, which matches the metadata of the thread episodes. Every other filter returns unsupported_filter.
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
request := &zep.SearchThreadSummariesRequest{
    Query: "query",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages. Search ranks a limited set of results, so iteration ends after that set. To get the top N results, set `page_size` to N and read the first page.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` — The search query.
    
</dd>
</dl>

<dl>
<dd>

**bfsOriginNodeUUIDs:** `[]string` — Nodes used as BFS origins for graph-distance-aware retrieval.
    
</dd>
</dl>

<dl>
<dd>

**centerNodeUUID:** `*string` 

The node to rank results by distance from. Required when reranker is
node_distance.
    
</dd>
</dl>

<dl>
<dd>

**filterLowQualityResults:** `*bool` 

The default is false. When true, Zep removes low-quality results at the knee of the
cross-encoder score curve. The filter never removes a result with a score
of 0.9 or more. It keeps at least 5 results, or all results when fewer
than 5 are available. The filter only removes results, so a page still
returns up to page_size results. Use this field only with reranker
cross_encoder. It is not supported for episode search. Invalid
combinations return a 400 error. The cutoff is approximate, not guaranteed.
When the cross-encoder does not score the results, for example after a
timeout, Zep returns the results without the filter.
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.ThreadSummarySearchFilters` — Filters constraining which items are returned.
    
</dd>
</dl>

<dl>
<dd>

**mmrLambda:** `*float64` 

The diversity weighting used for maximal marginal relevance reranking.
Required when reranker is mmr.
    
</dd>
</dl>

<dl>
<dd>

**reranker:** `*zep.SearchThreadSummariesRequestReranker` — The reranking strategy applied to retrieved results. Defaults to rrf.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**depth:** `*int` — The maximum traversal depth from the seed nodes. Defaults to 1.
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*zep.SubgraphRequestDirection` 

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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set `hyperedges_enabled` to `true` to enable extraction. A member needs the `project.edit` permission. Project API keys receive `403 permission_denied`.
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

<dl>
<dd>

**hyperedgesEnabled:** `*bool` 

Set to true to enable hyperedge extraction. This setting cannot be disabled.
A member needs project.edit. Project API keys receive 403 permission_denied.
    
</dd>
</dl>

<dl>
<dd>

**includePolicyViolatingEpisodes:** `*bool` 

When true, episode reads on graphs with a content policy include the
episodes that violated the policy.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.GetContentPolicy() -> *zep.ContentPolicy</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the current content policy revision of the project. A new graph binds this revision at creation.
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
client.Project.GetContentPolicy(
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

<details><summary><code>client.Project.SetContentPolicy(request) -> *zep.ContentPolicy</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replaces the project content policy and creates a new immutable revision. Graphs that already exist keep the revision they bound. An empty policy (no categories and no rules) removes the content policy for new graphs.
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
request := &zep.ContentPolicyRequest{}
client.Project.SetContentPolicy(
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

**categories:** `[]*zep.ContentPolicyCategoryRequest` 

The categories of the policy. Maximum 16. An empty list with no rules
means no content policy.
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*zep.ContentPolicyRuleRequest` — The rules of the policy. Maximum 32.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Project.ListContentPolicyRevisions() -> *zep.ContentPolicyRevisionPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists every revision of the project content policy, newest first, including revision 0.
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
request := &zep.ProjectListContentPolicyRevisionsRequest{}
client.Project.ListContentPolicyRevisions(
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

<details><summary><code>client.Project.GetContentPolicyRevision(RevisionUUID) -> *zep.ContentPolicy</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Project.GetContentPolicyRevision(
    context.TODO(),
    "revision_uuid",
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

**revisionUUID:** `string` — Revision UUID
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replaces the entity types and the edge types that the project uses.
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
request := &zep.TaskListRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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
request := &zep.ThreadListRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*zep.ThreadListRequestOrderBy` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*zep.ThreadListRequestOrder` — asc or desc
    
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
request := &zep.ThreadGetContextRequest{}
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
request := &zep.ThreadListEpisodesRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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
request := &zep.ThreadListMessagesRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

**order:** `*zep.ThreadListMessagesRequestOrder` — Sort direction: asc or desc
    
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

## TraceConnection
<details><summary><code>client.TraceConnection.List() -> *zep.TraceConnectionPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List trace connections in the current Zep project.
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
request := &zep.TraceConnectionListRequest{}
client.TraceConnection.List(
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

<details><summary><code>client.TraceConnection.Create(request) -> *zep.TraceConnection</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Verify a provider credential before Zep stores it. Both providers accept 1 to 60 requests per minute and an optional retention window. Example requests: `{"name":"Support traces","provider":"braintrust","credential":"secret","requests_per_minute":60,"retention_days":30}` and `{"name":"Support traces","provider":"langsmith","credential":"secret","requests_per_minute":60,"retention_days":30}`.
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
request := &zep.CreateTraceConnectionRequest{}
client.TraceConnection.Create(
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

**apiUrl:** `*string` — APIURL is an optional custom provider endpoint.
    
</dd>
</dl>

<dl>
<dd>

**credential:** `*string` — Credential is the write-only provider credential.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name is the connection name.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*zep.CreateTraceConnectionRequestProvider` — Provider is the Braintrust or LangSmith provider.
    
</dd>
</dl>

<dl>
<dd>

**requestsPerMinute:** `*int` — RequestsPerMinute is the request rate limit from 1 to 60 per minute.
    
</dd>
</dl>

<dl>
<dd>

**retentionDays:** `*int` — RetentionDays is an optional retention window in days.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TraceConnection.Get(ConnectionUUID) -> *zep.TraceConnection</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read one trace connection. The response includes a credential hint and never includes the credential.
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
client.TraceConnection.Get(
    context.TODO(),
    "connection_uuid",
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

**connectionUUID:** `string` — Trace connection UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TraceConnection.Delete(ConnectionUUID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a trace connection that no active trace import uses.
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
client.TraceConnection.Delete(
    context.TODO(),
    "connection_uuid",
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

**connectionUUID:** `string` — Trace connection UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TraceConnection.Update(ConnectionUUID, request) -> *zep.TraceConnection</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Verify changed settings before Zep stores them. Both providers accept 1 to 60 requests per minute and an optional retention window. Example request: `{"requests_per_minute":60,"retention_days":30}`.
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
request := &zep.UpdateTraceConnectionRequest{}
client.TraceConnection.Update(
    context.TODO(),
    "connection_uuid",
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

**connectionUUID:** `string` — Trace connection UUID
    
</dd>
</dl>

<dl>
<dd>

**apiUrl:** `*string` — APIURL is the new custom provider endpoint, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**credential:** `*string` — Credential is the new write-only provider credential, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name is the new connection name, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**requestsPerMinute:** `*int` — RequestsPerMinute is the new request rate limit, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**retentionDays:** `*int` — RetentionDays is the new retention window in days, when supplied.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TraceConnection.Verify(ConnectionUUID) -> *zep.TraceConnection</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Check the provider credential and update the connection status.
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
client.TraceConnection.Verify(
    context.TODO(),
    "connection_uuid",
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

**connectionUUID:** `string` — Trace connection UUID
    
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

Requires attribute-based access control to be enabled for your account.
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

Requires attribute-based access control to be enabled for your account.
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

Requires attribute-based access control to be enabled for your account.
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

Requires attribute-based access control to be enabled for your account.
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

Requires attribute-based access control to be enabled for your account.
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

Requires attribute-based access control to be enabled for your account.
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

Requires attribute-based access control to be enabled for your account.
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

Requires attribute-based access control to be enabled for your account.
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

Requires attribute-based access control to be enabled for your account.
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

Requires attribute-based access control to be enabled for your account.
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

Requires attribute-based access control to be enabled for your account.
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
request := &zep.UserGroupListForUserRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

**contentPolicy:** `*zep.GraphContentPolicyRequest` 

Content policy additions for the user's graph. The graph binds the
current project content policy plus these additions, and the binding
does not change after creation.
    
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

**userID:** `*string` 

An optional external identity identifier, such as an identity-provider
subject claim. Zep uses it to match an identity-provider sign-in to this
user. It is not an address; use uuid to address the user in v4. It is
unique within the project, can contain at most 500 characters, and
cannot be changed after creation.
    
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
request := &zep.UserListRequest{}
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

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*zep.UserListRequestOrderBy` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*zep.UserListRequestOrder` — asc or desc
    
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

## Agent Learning
<details><summary><code>client.Agent.Learning.Get(AgentUUID) -> *zep.AgentLearningState</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.LearningGetRequest{
    TaskFamily: "task_family",
}
client.Agent.Learning.Get(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `string` — Task family
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Learning.ListRuns(AgentUUID) -> *zep.Pagev4AgentSkillCompilationOutcome</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.LearningListRunsRequest{}
client.Agent.Learning.ListRuns(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` — Task family
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

## Agent LiteralPolicy
<details><summary><code>client.Agent.LiteralPolicy.Get(AgentUUID) -> *zep.AgentLiteralPolicy</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.LiteralPolicy.Get(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.LiteralPolicy.Update(AgentUUID, request) -> *zep.AgentLiteralPolicy</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.UpdateAgentLiteralPolicyRequest{
    AllowlistedClasses: &zep.AgentLiteralPolicyClasses{
        Environments: []string{
            "environments",
        },
        Tools: []string{
            "tools",
        },
    },
    AllowlistedValues: &zep.AgentLiteralPolicyValues{
        Environments: []string{
            "environments",
        },
        Tools: []string{
            "tools",
        },
    },
    Default: agent.UpdateAgentLiteralPolicyRequestDefaultParameterize,
}
client.Agent.LiteralPolicy.Update(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**allowlistedClasses:** `*zep.AgentLiteralPolicyClasses` — Typed reusable literal classes.
    
</dd>
</dl>

<dl>
<dd>

**allowlistedValues:** `*zep.AgentLiteralPolicyValues` — Exact reusable values scoped by tools and environments.
    
</dd>
</dl>

<dl>
<dd>

**default_:** `agent.UpdateAgentLiteralPolicyRequestDefault` — The remediation for literals that are not explicitly allowed.
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*int` — The current literal-policy revision used for optimistic concurrency.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Skill
<details><summary><code>client.Agent.Skill.Create(AgentUUID, request) -> *zep.AgentSkill</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.CreateAgentSkillRequest{
    Definition: &zep.SkillDefinition{},
}
client.Agent.Skill.Create(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**definition:** `*zep.SkillDefinition` 
    
</dd>
</dl>

<dl>
<dd>

**skillID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.List(AgentUUID, request) -> *zep.Pagev4AgentSkillSearchHit</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.AgentSkillListRequest{}
client.Agent.Skill.List(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**environments:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tools:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Search(AgentUUID, request) -> *zep.AgentSkillSearchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.AgentSkillSearchRequest{
    Query: "query",
}
client.Agent.Skill.Search(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**environments:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**includeMarkdown:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**markdownFormat:** `*agent.AgentSkillSearchRequestMarkdownFormat` 

MarkdownFormat selects the form of inline `markdown`. `agent` (the
default) returns the Skill text that an Agent needs: no evidence
markers, the description only in the frontmatter, and no empty
sections. `full` returns the stored SKILL.md with its evidence markers.
    
</dd>
</dl>

<dl>
<dd>

**query:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tools:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Get(AgentUUID, SkillUUID) -> *zep.AgentSkill</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Skill.Get(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Approve(AgentUUID, SkillUUID, request) -> *zep.AgentSkillAdmissionDecision</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.ApproveAgentSkillRequest{
    ExpectedVersion: 1,
}
client.Agent.Skill.Approve(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedVersion:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Retire(AgentUUID, SkillUUID, request) -> *zep.AgentSkill</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.RetireAgentSkillRequest{
    ExpectedVersion: 1,
}
client.Agent.Skill.Retire(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedVersion:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.CreateVersion(AgentUUID, SkillUUID, request) -> *zep.AgentSkillVersion</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.CreateAgentSkillVersionRequest{
    Definition: &zep.SkillDefinition{},
    ExpectedVersion: 1,
}
client.Agent.Skill.CreateVersion(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**definition:** `*zep.SkillDefinition` 
    
</dd>
</dl>

<dl>
<dd>

**expectedVersion:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Split
<details><summary><code>client.Agent.Split.Plan(AgentUUID, request) -> *zep.AgentSplitPlan</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.CreateAgentSplitPlanRequest{
    Destinations: []*zep.AgentSplitPlanDestinationRequest{
        &zep.AgentSplitPlanDestinationRequest{
            Agent: &zep.CreateAgentRequest{
                AgentID: "agent_id",
                Name: "name",
                SecurityDomain: "security_domain",
            },
            Skills: []*zep.AgentSplitPlanSkillSelectionRequest{
                &zep.AgentSplitPlanSkillSelectionRequest{
                    SkillUUID: "skill_uuid",
                    SkillVersionUUID: "skill_version_uuid",
                    Version: 1,
                },
            },
        },
    },
    ExpectedRevision: 1,
    Rationale: "rationale",
    ReviewAcknowledged: true,
}
client.Agent.Split.Plan(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Source Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**destinations:** `[]*zep.AgentSplitPlanDestinationRequest` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**rationale:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**reviewAcknowledged:** `bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent TaskFamily
<details><summary><code>client.Agent.TaskFamily.List(AgentUUID) -> *zep.AgentTaskFamilyDescriptorPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.TaskFamilyListRequest{}
client.Agent.TaskFamily.List(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

<details><summary><code>client.Agent.TaskFamily.Rename(AgentUUID, request) -> *zep.RenameAgentTaskFamilyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.RenameAgentTaskFamilyRequest{
    FromSlug: "from_slug",
    ToSlug: "to_slug",
}
client.Agent.TaskFamily.Rename(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**fromSlug:** `string` — The current slug of the live descriptor.
    
</dd>
</dl>

<dl>
<dd>

**toSlug:** `string` — The new slug, or the slug of the live descriptor to merge into.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent TraceImport
<details><summary><code>client.Agent.TraceImport.List(AgentUUID) -> *zep.TraceImportPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List trace imports that belong to this Agent.
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
request := &agent.TraceImportListRequest{}
client.Agent.TraceImport.List(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

<details><summary><code>client.Agent.TraceImport.Create(AgentUUID, request) -> *zep.TraceImportCreateResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a scheduled import or queue a one-time import. A scheduled import returns HTTP 201 with `import`. A one-time import returns HTTP 202 with `import`, `run`, and `task`. Example request: `{"connection_uuid":"8c78a85e-eac2-4f57-b5f5-59a68a1e77a1","provider_project_id":"project-123","name":"Support traces","selection":{"trace_ids":["trace-123"]},"mapping":{"task_family":{"source":"fixed","value":"support"}}}`.
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
request := &agent.CreateTraceImportRequest{}
client.Agent.TraceImport.Create(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**connectionUUID:** `*string` — ConnectionUUID identifies the project trace connection.
    
</dd>
</dl>

<dl>
<dd>

**learnFrom:** `*bool` — LearnFrom controls whether imported evidence can support Skills.
    
</dd>
</dl>

<dl>
<dd>

**mapping:** `*zep.TraceMapping` — Mapping defines how source traces become Traces. Example: {"task_family":{"source":"fixed","value":"support"}}.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name is the import display name.
    
</dd>
</dl>

<dl>
<dd>

**onSourceChange:** `*zep.TraceImportOnSourceChange` — OnSourceChange defines how changed source traces are handled.
    
</dd>
</dl>

<dl>
<dd>

**providerProjectID:** `*string` — ProviderProjectID is the source provider project identifier.
    
</dd>
</dl>

<dl>
<dd>

**requireReview:** `*bool` — RequireReview requires review for candidates backed by this import.
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*zep.TraceImportSchedule` — Schedule enables recurring imports when supplied. Example: {"interval_hours":4,"start_from":"now","settle_minutes":5,"max_open_hours":24}.
    
</dd>
</dl>

<dl>
<dd>

**selection:** `*zep.TraceImportSelection` — Selection defines the traces to import. Example: {"trace_ids":["trace_123"]}.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.TraceImport.Get(AgentUUID, ImportUUID) -> *zep.TraceImport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read one trace import. The response does not include its source cursor.
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
client.Agent.TraceImport.Get(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.TraceImport.Delete(AgentUUID, ImportUUID) -> *zep.Task</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set `traces=delete` to queue asynchronous Trace deletion. The default keeps Traces. Example query: `?traces=delete`.
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
request := &agent.TraceImportDeleteRequest{}
client.Agent.TraceImport.Delete(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>

<dl>
<dd>

**traces:** `*agent.TraceImportDeleteRequestTraces` — Whether to delete imported Traces
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.TraceImport.Update(AgentUUID, ImportUUID, request) -> *zep.TraceImport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use `expected_revision` to reject a stale update. Example request: `{"expected_revision":1,"name":"Updated support traces"}`.
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
request := &agent.UpdateTraceImportRequest{}
client.Agent.TraceImport.Update(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*int` — ExpectedRevision is the current import revision.
    
</dd>
</dl>

<dl>
<dd>

**learnFrom:** `*bool` — LearnFrom is the new learning setting, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**mapping:** `*zep.TraceMapping` — Mapping is the new mapping, when supplied. Example: {"task_family":{"source":"fixed","value":"support"}}.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name is the new import name, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**onSourceChange:** `*agent.UpdateTraceImportRequestOnSourceChange` — OnSourceChange is the new change policy, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**requireReview:** `*bool` — RequireReview is the new review setting, when supplied.
    
</dd>
</dl>

<dl>
<dd>

**schedule:** `*zep.TraceImportSchedule` — Schedule is the new schedule, when supplied. Example: {"interval_hours":4,"start_from":"now","settle_minutes":5,"max_open_hours":24}.
    
</dd>
</dl>

<dl>
<dd>

**selection:** `*zep.TraceImportSelection` — Selection is the new selection, when supplied. Example: {"trace_ids":["trace_123"]}.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.TraceImport.Pause(AgentUUID, ImportUUID) -> *zep.TraceImport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pause a scheduled trace import. One-time imports cannot be paused.
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
client.Agent.TraceImport.Pause(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.TraceImport.Resume(AgentUUID, ImportUUID) -> *zep.TraceImport</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resume a scheduled trace import. Zep verifies credentials first when they caused the pause.
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
client.Agent.TraceImport.Resume(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Trace
<details><summary><code>client.Agent.Trace.Create(AgentUUID, request) -> *zep.AgentTrace</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.CreateAgentTraceRequest{
    Objective: "objective",
}
client.Agent.Trace.Create(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**learnFrom:** `*bool` — Whether eligible evidence may participate in learning. Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Developer metadata used to organize and filter Traces.
    
</dd>
</dl>

<dl>
<dd>

**objective:** `string` — The task-instance objective captured for this attempt.
    
</dd>
</dl>

<dl>
<dd>

**parentTraceUUID:** `*string` — The parent Trace UUID when this attempt retries an earlier attempt.
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` 

The task-family slug. Optional when inferred task families are enabled
for the account: the service assigns one after the Trace closes.
    
</dd>
</dl>

<dl>
<dd>

**traceID:** `*string` — The optional customer-assigned identifier, unique within the Agent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.List(AgentUUID, request) -> *zep.AgentTracePage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.AgentTraceListRequest{}
client.Agent.Trace.List(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*agent.TraceListRequestOrderBy` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*agent.TraceListRequestOrder` — asc or desc
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `*string` — Filter by owning Trace import UUID
    
</dd>
</dl>

<dl>
<dd>

**sourceKind:** `*agent.TraceListRequestSourceKind` — Source kind
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Filters to Traces whose metadata contains these key-value pairs.
    
</dd>
</dl>

<dl>
<dd>

**outcome:** `*zep.AgentTraceOutcome` — Filters to one reported outcome.
    
</dd>
</dl>

<dl>
<dd>

**parentTraceUUID:** `*string` — Filters to direct retries of this parent Trace.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*zep.AgentTraceLifecycle` — Filters to one lifecycle state.
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` — Filters to an exact task-family slug.
    
</dd>
</dl>

<dl>
<dd>

**verification:** `*zep.AgentTraceVerification` — Filters to one verification strength.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.Get(AgentUUID, TraceUUID) -> *zep.AgentTrace</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Trace.Get(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.Delete(AgentUUID, TraceUUID, request) -> *zep.AgentTraceSourceDeletionResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.DeleteAgentTraceRequest{
    ExpectedRevision: 1,
}
client.Agent.Trace.Delete(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.DeleteAgentTraceRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.Update(AgentUUID, TraceUUID, request) -> *zep.AgentTrace</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.PatchAgentTraceRequest{
    ExpectedRevision: 1,
}
client.Agent.Trace.Update(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` — The current Trace revision used for optimistic concurrency.
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` — Replacement task-family slug. Set to null to clear it.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.Abandon(AgentUUID, TraceUUID, request) -> *zep.AgentTraceFinalizationResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.AbandonAgentTraceRequest{
    HighestAcceptedSequence: 1,
}
client.Agent.Trace.Abandon(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**assertionUUID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**highestAcceptedSequence:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**missingSequenceRanges:** `[]*zep.AgentTraceSequenceRange` 
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**verification:** `*zep.AgentTraceVerification` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.Close(AgentUUID, TraceUUID, request) -> *zep.AgentTraceFinalizationResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.CloseAgentTraceRequest{}
client.Agent.Trace.Close(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**closingSequence:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**missingSequenceRanges:** `[]*zep.AgentTraceSequenceRange` 
    
</dd>
</dl>

<dl>
<dd>

**outcome:** `*zep.AgentTraceOutcome` 
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**verification:** `*zep.AgentTraceVerification` 
    
</dd>
</dl>

<dl>
<dd>

**verifier:** `*zep.AgentTraceCloseVerifier` 
    
</dd>
</dl>

<dl>
<dd>

**verifierID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.CorrectTaskFamily(AgentUUID, TraceUUID, request) -> *zep.AgentTraceFinalizationResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.CorrectAgentTraceTaskFamilyRequest{
    ExpectedRevision: 1,
    Reason: "reason",
}
client.Agent.Trace.CorrectTaskFamily(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` — The current Trace revision used for optimistic concurrency.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `string` — Why the terminal Trace classification is being corrected.
    
</dd>
</dl>

<dl>
<dd>

**taskFamily:** `*string` — Replacement task-family slug. Set to null to clear it.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.ListEvents(AgentUUID, TraceUUID) -> *zep.AgentTraceEventPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.TraceListEventsRequest{}
client.Agent.Trace.ListEvents(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

<details><summary><code>client.Agent.Trace.AppendEvent(AgentUUID, TraceUUID, request) -> *zep.AgentTraceEvent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.AppendAgentTraceEventRequest{
    EventType: zep.AgentTraceEventTypeInputReference,
}
client.Agent.Trace.AppendEvent(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**content:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**context:** `*zep.AgentTraceEventContext` 
    
</dd>
</dl>

<dl>
<dd>

**eventID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `zep.AgentTraceEventType` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**occurredAt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sequence:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `[]*zep.AgentTraceEventSource` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.DeleteEvent(AgentUUID, TraceUUID, EventUUID, request) -> *zep.AgentTraceSourceDeletionResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.DeleteAgentTraceRequest{
    ExpectedRevision: 1,
}
client.Agent.Trace.DeleteEvent(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
    "event_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**eventUUID:** `string` — Event UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.DeleteAgentTraceRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.Reopen(AgentUUID, TraceUUID, request) -> *zep.AgentTrace</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.ReopenAgentTraceRequest{
    ExpectedRevision: 1,
}
client.Agent.Trace.Reopen(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` — The current Trace revision used for optimistic concurrency.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.GetSummary(AgentUUID, TraceUUID) -> *zep.AgentTraceSummaryVersion</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Trace.GetSummary(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Trace.ListSummaryVersions(AgentUUID, TraceUUID) -> *zep.AgentTraceSummaryPage</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.TraceListSummaryVersionsRequest{}
client.Agent.Trace.ListSummaryVersions(
    context.TODO(),
    "agent_uuid",
    "trace_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` — Trace UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

## Agent Verifier
<details><summary><code>client.Agent.Verifier.List(AgentUUID, request) -> *zep.Pagev4AgentVerifier</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.AgentVerifierListRequest{}
client.Agent.Verifier.List(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Page cursor
    
</dd>
</dl>

<dl>
<dd>

**class:** `*string` — Filters to an exact immutable verifier class.
    
</dd>
</dl>

<dl>
<dd>

**principalID:** `*string` — Filters to registrations bound to this exact principal identifier.
    
</dd>
</dl>

<dl>
<dd>

**principalType:** `*string` — Filters to registrations bound to this principal type.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*zep.AgentVerifierStatus` — Filters to active or revoked registrations. Omit to include both.
    
</dd>
</dl>

<dl>
<dd>

**verifierID:** `*string` — Filters to an exact developer-assigned verifier identifier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Verifier.Get(AgentUUID, VerifierUUID) -> *zep.AgentVerifier</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Verifier.Get(
    context.TODO(),
    "agent_uuid",
    "verifier_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**verifierUUID:** `string` — Verifier UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Verifier.Update(AgentUUID, VerifierUUID, request) -> *zep.AgentVerifier</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.PatchAgentVerifierRequest{
    ExpectedRevision: 1,
}
client.Agent.Verifier.Update(
    context.TODO(),
    "agent_uuid",
    "verifier_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**verifierUUID:** `string` — Verifier UUID
    
</dd>
</dl>

<dl>
<dd>

**capabilities:** `map[string]any` — Replacement capabilities. Set to null to clear them.
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` — The current verifier revision used for optimistic concurrency.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Replacement metadata. Set to null to clear it.
    
</dd>
</dl>

<dl>
<dd>

**permittedStrengths:** `[]zep.AgentVerificationStrength` — Replacement permitted strengths. The list cannot be empty or null.
    
</dd>
</dl>

<dl>
<dd>

**principalBindings:** `[]*zep.AgentVerifierPrincipalBindingInput` — Replacement exact principal bindings. The list cannot be empty or null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Verifier.InvalidateEvidence(AgentUUID, VerifierUUID, request) -> *zep.AgentVerifierEvidenceInvalidationResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.InvalidateAgentVerifierEvidenceRequest{
    Reason: "reason",
    VerifierRevisions: []int{
        1,
    },
}
client.Agent.Verifier.InvalidateEvidence(
    context.TODO(),
    "agent_uuid",
    "verifier_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**verifierUUID:** `string` — Verifier UUID
    
</dd>
</dl>

<dl>
<dd>

**reason:** `string` — Why evidence from these verifier revisions is invalid.
    
</dd>
</dl>

<dl>
<dd>

**verifierRevisions:** `[]int` — Verifier revisions whose accepted evidence is no longer valid.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Verifier.Revoke(AgentUUID, VerifierUUID, request) -> *zep.AgentVerifier</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &agent.RevokeAgentVerifierRequest{
    ExpectedRevision: 1,
}
client.Agent.Verifier.Revoke(
    context.TODO(),
    "agent_uuid",
    "verifier_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**verifierUUID:** `string` — Verifier UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` — The current verifier revision used for optimistic concurrency.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Skill Candidate
<details><summary><code>client.Agent.Skill.Candidate.List(AgentUUID) -> *zep.Pagev4AgentSkillCandidateReviewSummary</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.CandidateListRequest{}
client.Agent.Skill.Candidate.List(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

<details><summary><code>client.Agent.Skill.Candidate.Get(AgentUUID, ReviewUUID) -> *zep.AgentSkillCandidateReview</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Skill.Candidate.Get(
    context.TODO(),
    "agent_uuid",
    "review_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**reviewUUID:** `string` — Candidate review UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Skill Evaluation
<details><summary><code>client.Agent.Skill.Evaluation.CreateForCandidate(AgentUUID, ReviewUUID, CandidateUUID, request) -> *zep.AgentSkillCandidateEvaluation</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.CreateAgentSkillCandidateEvaluationRequest{
    EvaluatedBy: skill.CreateAgentSkillCandidateEvaluationRequestEvaluatedByCustomer,
    EvaluatorVersion: "evaluator_version",
    ExpectedRevision: 1,
    Verdict: skill.CreateAgentSkillCandidateEvaluationRequestVerdictSucceeded,
}
client.Agent.Skill.Evaluation.CreateForCandidate(
    context.TODO(),
    "agent_uuid",
    "review_uuid",
    "candidate_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**reviewUUID:** `string` — Candidate review UUID
    
</dd>
</dl>

<dl>
<dd>

**candidateUUID:** `string` — Candidate UUID
    
</dd>
</dl>

<dl>
<dd>

**controlledTest:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**environment:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**evaluatedBy:** `skill.CreateAgentSkillCandidateEvaluationRequestEvaluatedBy` 
    
</dd>
</dl>

<dl>
<dd>

**evaluatorVersion:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**evidenceReference:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**metrics:** `map[string]float64` 
    
</dd>
</dl>

<dl>
<dd>

**verdict:** `skill.CreateAgentSkillCandidateEvaluationRequestVerdict` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Evaluation.Create(AgentUUID, SkillUUID, request) -> *zep.AgentSkillEvaluation</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.CreateAgentSkillEvaluationRequest{
    EvaluatedBy: skill.CreateAgentSkillEvaluationRequestEvaluatedByCustomer,
    EvaluatorVersion: "evaluator_version",
    SkillVersion: 1,
    Verdict: skill.CreateAgentSkillEvaluationRequestVerdictSucceeded,
}
client.Agent.Skill.Evaluation.Create(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**controlledTest:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**environment:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**evaluatedBy:** `skill.CreateAgentSkillEvaluationRequestEvaluatedBy` 
    
</dd>
</dl>

<dl>
<dd>

**evaluatorVersion:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**evidenceReference:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metrics:** `map[string]float64` 
    
</dd>
</dl>

<dl>
<dd>

**skillVersion:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**useUUID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**verdict:** `skill.CreateAgentSkillEvaluationRequestVerdict` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Skill Publication
<details><summary><code>client.Agent.Skill.Publication.Lookup(AgentUUID) -> *zep.AgentSkillPublicationLineage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Find a destination Skill by its source Agent, Skill, and version. The endpoint returns 404 until an automatic publication is ready or after it is invalidated.
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
request := &skill.PublicationLookupRequest{
    SourceAgentUUID: "source_agent_uuid",
    SourceSkillUUID: "source_skill_uuid",
    SourceVersion: 1,
}
client.Agent.Skill.Publication.Lookup(
    context.TODO(),
    "agent_uuid",
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

**agentUUID:** `string` — Destination Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**sourceAgentUUID:** `string` — Source Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**sourceSkillUUID:** `string` — Source Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**sourceVersion:** `int` — Source version number
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Publication.Get(AgentUUID, SkillUUID) -> *zep.AgentSkillPublicationLineage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the publication that created a destination Skill. Use its source version and policy identity to verify a copied Skill. The endpoint returns 404 for unpublished and invalidated Skills.
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
client.Agent.Skill.Publication.Get(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Destination Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Destination Skill UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Skill Evidence
<details><summary><code>client.Agent.Skill.Evidence.List(AgentUUID, SkillUUID) -> *zep.Pagev4AgentSkillEvidence</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.EvidenceListRequest{}
client.Agent.Skill.Evidence.List(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

## Agent Skill Relation
<details><summary><code>client.Agent.Skill.Relation.List(AgentUUID, SkillUUID) -> *zep.Pagev4AgentSkillRelationship</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.RelationListRequest{}
client.Agent.Skill.Relation.List(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*skill.RelationListRequestKindItem` — Relationship kinds
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

## Agent Skill Version
<details><summary><code>client.Agent.Skill.Version.RestoreVersion(AgentUUID, SkillUUID, request) -> *zep.AgentSkillVersion</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.RestoreAgentSkillVersionRequest{
    ExpectedVersion: 1,
    Version: 1,
}
client.Agent.Skill.Version.RestoreVersion(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**expectedVersion:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Version.List(AgentUUID, SkillUUID) -> *zep.Pagev4AgentSkillVersion</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.VersionListRequest{}
client.Agent.Skill.Version.List(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**markdownFormat:** `*skill.VersionListRequestMarkdownFormat` — Markdown form: agent (default) or full
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Version.Compare(AgentUUID, SkillUUID, request) -> *zep.AgentSkillVersionComparison</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.CompareAgentSkillVersionsRequest{
    FromVersion: 1,
    ToVersion: 1,
}
client.Agent.Skill.Version.Compare(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**fromVersion:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**toVersion:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Version.Get(AgentUUID, SkillUUID, Version) -> *zep.AgentSkillVersion</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.VersionGetRequest{}
client.Agent.Skill.Version.Get(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
    1,
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**version:** `int` — Skill version
    
</dd>
</dl>

<dl>
<dd>

**markdownFormat:** `*skill.VersionGetRequestMarkdownFormat` — Markdown form: agent (default) or full
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Skill Use
<details><summary><code>client.Agent.Skill.Use.Create(AgentUUID, SkillUUID, request) -> *zep.AgentSkillUse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &skill.CreateAgentSkillUseRequest{
    SearchID: "search_id",
    TraceUUID: "trace_uuid",
    Usage: skill.CreateAgentSkillUseRequestUsageSelected,
}
client.Agent.Skill.Use.Create(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**outcome:** `*zep.AddAgentSkillUseOutcomeRequest` 
    
</dd>
</dl>

<dl>
<dd>

**reasonCode:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**searchID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**traceUUID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**usage:** `skill.CreateAgentSkillUseRequestUsage` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Use.AddOutcome(AgentUUID, SkillUUID, UseUUID, request) -> *zep.AgentSkillUseOutcome</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &zep.AddAgentSkillUseOutcomeRequest{
    Outcome: zep.AddAgentSkillUseOutcomeRequestOutcomeSucceeded,
}
client.Agent.Skill.Use.AddOutcome(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
    "use_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**useUUID:** `string` — Skill use UUID
    
</dd>
</dl>

<dl>
<dd>

**request:** `*zep.AddAgentSkillUseOutcomeRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent Skill Export
<details><summary><code>client.Agent.Skill.Export.Create(AgentUUID, SkillUUID, Version) -> *zep.Task</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Skill.Export.Create(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
    1,
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**version:** `int` — Skill version
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.Skill.Export.Get(AgentUUID, SkillUUID, Version, TaskUUID) -> string</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Agent.Skill.Export.Get(
    context.TODO(),
    "agent_uuid",
    "skill_uuid",
    1,
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**skillUUID:** `string` — Skill UUID
    
</dd>
</dl>

<dl>
<dd>

**version:** `int` — Skill version
    
</dd>
</dl>

<dl>
<dd>

**taskUUID:** `string` — Export Task UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent TraceImport Run
<details><summary><code>client.Agent.TraceImport.Run.List(AgentUUID, ImportUUID) -> *zep.TraceImportRunPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List runs that belong to this trace import.
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
request := &traceimport.RunListRequest{}
client.Agent.TraceImport.Run.List(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

<details><summary><code>client.Agent.TraceImport.Run.Create(AgentUUID, ImportUUID) -> *zep.Task</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a pending run and its Task. This operation does not start the run.
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
client.Agent.TraceImport.Run.Create(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Agent.TraceImport.Run.Get(AgentUUID, ImportUUID, RunUUID) -> *zep.TraceImportRun</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read one run that belongs to this trace import.
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
client.Agent.TraceImport.Run.Get(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
    "run_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>

<dl>
<dd>

**runUUID:** `string` — Run UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Agent TraceImport Run Item
<details><summary><code>client.Agent.TraceImport.Run.Item.List(AgentUUID, ImportUUID, RunUUID) -> *zep.TraceImportRunItemPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List trace results for this run. Run items do not include source payload.
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
request := &run.ItemListRequest{}
client.Agent.TraceImport.Run.Item.List(
    context.TODO(),
    "agent_uuid",
    "import_uuid",
    "run_uuid",
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

**agentUUID:** `string` — Agent UUID
    
</dd>
</dl>

<dl>
<dd>

**importUUID:** `string` — Trace import UUID
    
</dd>
</dl>

<dl>
<dd>

**runUUID:** `string` — Run UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

## Graph DocumentSummary
<details><summary><code>client.Graph.DocumentSummary.List(GraphUUID, request) -> *zep.DocumentSummaryPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the document summaries of a graph and returns one page of results. The default `page_size` is 50, and the maximum `page_size` is 100.
You can sort by `uuid` or `created_at`. The default sort is `uuid` in descending order.
You cannot use these `filters` fields: `node_labels`, `exclude_node_labels`, `edge_types`, `exclude_edge_types`, `edge_uuids`, `episode_uuids`, `connected_node_uuids`, `source_node_uuids`, `target_node_uuids`, and `content_policy_violated`.
`date_filters` supports only the `created_at` field.
If the request has an unsupported or unknown filter field, the API returns an `unsupported_filter` error.
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
request := &graph.DocumentSummaryListRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.DocumentSummaryListRequestOrderBy` — Sort key: uuid (default) or created_at
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.DocumentSummaryListRequestOrder` — Sort direction: asc or desc (default desc)
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.DocumentSummaryListFilters` 

Filters that limit the items in the list. The operation description
names the filter fields that the operation does not support.
    
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
request := &graph.EpisodeListForDocumentRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**documentID:** `string` — Document ID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**metadata:** `map[string]any` — Metadata to store on the episode. Max 10 keys. Values must be strings, numbers, booleans, or arrays of scalars.
    
</dd>
</dl>

<dl>
<dd>

**referenceTime:** `*string` 

The episode's reference time, used for temporal reasoning rather than
ingestion time. Defaults to the ingestion time.
    
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

**type_:** `*graph.AddEpisodeRequestType` — The data format of the episode: text, json, or message. Defaults to text.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the episodes of a graph. `filters.mentioned_node_uuids` restricts
the results to episodes that mention any of the listed node UUIDs. The
list can also contain episode UUIDs: an episode UUID matches that episode,
so one request can return a known set of episodes. At most 256 entries.
`filters.metadata_filters` restricts the results to episodes whose stored
metadata matches the predicate.
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
request := &graph.EpisodeListRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.EpisodeListRequestOrderBy` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.EpisodeListRequestOrder` — Sort direction: asc or desc
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.EpisodeListFilters` — Filters that limit the episodes in the list.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**episodeUUID:** `string` — Episode UUID
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata to merge onto the episode; a key set to null is removed. Max 10 keys after the merge. Values must be strings, numbers, booleans, or arrays of scalars.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Episode.GetDebugLogs(GraphUUID, EpisodeUUID) -> *zep.EpisodeDebugLog</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the ingestion workflow log of an episode. The log exists only when debug logging was enabled for the project when the episode was ingested (see `debug_log.enable`). The log holds episode content, so an API key with an ABAC policy needs an explicit grant of this action; the `readonly` macro does not grant it.
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
client.Graph.Episode.GetDebugLogs(
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

<details><summary><code>client.Graph.Episode.ListIngestionTraces(GraphUUID, EpisodeUUID) -> *zep.IngestionTracePage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the ingestion traces of an episode, oldest first. Each trace records the input and the output of one ingestion step, with an explanation on each output entry that has one. Traces exist only when ingestion tracing was enabled for the project when the episode was ingested (see `debug_log.enable`). An episode with no traces returns a page with an empty `items` array. Traces hold episode content, prompt input, and model output, so an API key with an ABAC policy needs an explicit grant of this action; the `readonly` macro does not grant it.
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
request := &graph.EpisodeListIngestionTracesRequest{}
client.Graph.Episode.ListIngestionTraces(
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**episodeUUID:** `string` — Episode UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
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

## Graph Edge
<details><summary><code>client.Graph.Edge.Add(GraphUUID, request) -> *zep.AddEdgesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds 1 to 100 edges. A name creates a node when deduplicate is false. When deduplicate is true, Zep matches a node by name first.
Example: {"edges":[{"fact":"Ada works at Acme Corp","fact_name":"WORKS_AT","source_node":{"uuid":"f47ac10b-58cc-4372-a567-0e02b2c3d479"},"target_node":{"uuid":"f47ac10b-58cc-4372-a567-0e02b2c3d480"}},{"fact":"Ada leads a team","fact_name":"LEADS","source_node":{"name":"Ada Lovelace","labels":["Person"]},"target_node":{"name":"Engineering","labels":["Department"]}}],"deduplicate":false}
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
request := &graph.AddEdgesRequest{
    Edges: []*zep.EdgeInput{
        &zep.EdgeInput{
            Fact: "Ada works at Acme Corp",
            FactName: "WORKS_AT",
            SourceNode: &zep.EdgeNodeRef{},
            TargetNode: &zep.EdgeNodeRef{},
        },
    },
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**deduplicate:** `*bool` 

When true, Zep compares each edge with graph edges and can merge a
duplicate or invalidate a contradicted edge. This adds an LLM call per
edge. The default is false.
    
</dd>
</dl>

<dl>
<dd>

**edges:** `[]*zep.EdgeInput` — The edges to add to the graph. The request accepts 1 to 100 edges.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the edges of a graph and returns one page of results. The default `page_size` is 50, and the maximum `page_size` is 100.
You can sort by `uuid`, `created_at`, or `valid_at`. The default sort is `uuid` in descending order. An edge with no `valid_at` sorts as if its `valid_at` is 1970-01-01T00:00:00Z. Thus it comes before an edge with a later `valid_at` in ascending order, and after an edge with an earlier `valid_at`.
You can use all `filters` fields except `content_policy_violated`.
If the request has an unsupported or unknown filter field, the API returns an `unsupported_filter` error.
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
request := &graph.EdgeListRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.EdgeListRequestOrderBy` — Sort key: uuid (default), created_at, or valid_at
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.EdgeListRequestOrder` — Sort direction: asc or desc (default desc)
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.EdgeListFilters` 

Filters that limit the items in the list. The operation description
names the filter fields that the operation does not support.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates one edge. When the edge belongs to a hyperedge, changing fact
rewrites it on every member of that hyperedge in one all-or-nothing
write, because the members share it. Attribute-only edits touch this
edge alone.
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

## Graph Hyperedge
<details><summary><code>client.Graph.Hyperedge.Add(GraphUUID, request) -> *zep.AddHyperedgeResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates one hyperedge: a fact that relates more than two nodes, written
onto one member edge per node pair. The member edges must span at least
three distinct nodes, since two nodes are a pair of edges rather than a
hyperedge; a single pair uses graph.edge.add. Zep assigns the hyperedge
identifier and every member edge identifier at accept time, and the
members become readable when the task completes.
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
request := &graph.AddHyperedgeRequest{
    Edges: []*zep.HyperedgeInput{
        &zep.HyperedgeInput{
            Name: "name",
            SourceNode: &zep.EdgeNodeRef{},
            TargetNode: &zep.EdgeNodeRef{},
        },
    },
    Fact: "fact",
}
client.Graph.Hyperedge.Add(
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**attributes:** `map[string]any` — Additional attributes to store on every member edge.
    
</dd>
</dl>

<dl>
<dd>

**edges:** `[]*zep.HyperedgeInput` 

The member edges to create. Each pair becomes one edge carrying the
shared fact.
    
</dd>
</dl>

<dl>
<dd>

**expiredAt:** `*string` — The time at which the fact was superseded or invalidated.
    
</dd>
</dl>

<dl>
<dd>

**fact:** `string` — The natural-language fact, written onto every member edge.
    
</dd>
</dl>

<dl>
<dd>

**invalidAt:** `*string` — The time at which the fact stopped being true.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` — Metadata attached to the episode created for this hyperedge.
    
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

<details><summary><code>client.Graph.Hyperedge.List(GraphUUID, request) -> *zep.HyperedgePage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the graph's hyperedges. A hyperedge is listed while it has at
least one member edge. Supported filters are node_uuids, edge_uuids and
episode_uuids: a hyperedge matches a list when any of its members does,
and must match every list supplied.
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
request := &graph.HyperedgeListRequest{}
client.Graph.Hyperedge.List(
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.HyperedgeListRequestOrderBy` — Sort key: uuid or created_at
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.HyperedgeListRequestOrder` — Sort direction: asc or desc
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.HyperedgeListFilters` — Filters that limit the hyperedges in the list.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Hyperedge.Get(GraphUUID, HyperedgeUUID) -> *zep.Hyperedge</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns one hyperedge assembled from its member edges. The fact and the
validity timestamps are shared by every member and are reported on the
hyperedge; each member reports only its own name and endpoints.
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
client.Graph.Hyperedge.Get(
    context.TODO(),
    "graph_uuid",
    "hyperedge_uuid",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**hyperedgeUUID:** `string` — Hyperedge UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Hyperedge.Delete(GraphUUID, HyperedgeUUID) -> *zep.AsyncResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes every member edge of the hyperedge. After the task completes the
hyperedge and each of its members are gone.
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
client.Graph.Hyperedge.Delete(
    context.TODO(),
    "graph_uuid",
    "hyperedge_uuid",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**hyperedgeUUID:** `string` — Hyperedge UUID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Hyperedge.Update(GraphUUID, HyperedgeUUID, request) -> *zep.Hyperedge</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the shared fact and writes it onto every member edge in one
all-or-nothing write: either every member carries the new fact or none
does. Only fact is accepted, because it is the only field the members
share. name belongs to each member edge, and membership changes use
create_edge and delete_edge. Updating fact on one member through
graph.edge.update cascades the same way.
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
request := &graph.PatchHyperedgeRequest{}
client.Graph.Hyperedge.Update(
    context.TODO(),
    "graph_uuid",
    "hyperedge_uuid",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**hyperedgeUUID:** `string` — Hyperedge UUID
    
</dd>
</dl>

<dl>
<dd>

**fact:** `*string` — The natural-language fact, rewritten onto every member edge.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Hyperedge.CreateEdge(GraphUUID, HyperedgeUUID, request) -> *zep.AddHyperedgeEdgeResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Adds one member edge to an existing hyperedge. The new member inherits the
hyperedge's fact and timestamps and joins its episodes, so only its own
name and node pair are supplied.
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
request := &graph.AddHyperedgeEdgeRequest{
    Name: "name",
    SourceNode: &zep.EdgeNodeRef{},
    TargetNode: &zep.EdgeNodeRef{},
}
client.Graph.Hyperedge.CreateEdge(
    context.TODO(),
    "graph_uuid",
    "hyperedge_uuid",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**hyperedgeUUID:** `string` — Hyperedge UUID
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the new member edge, in upper snake case.
    
</dd>
</dl>

<dl>
<dd>

**sourceNode:** `*zep.EdgeNodeRef` — The source node of the new member.
    
</dd>
</dl>

<dl>
<dd>

**targetNode:** `*zep.EdgeNodeRef` — The target node of the new member.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Graph.Hyperedge.DeleteEdge(GraphUUID, HyperedgeUUID, EdgeUUID) -> *zep.AsyncResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes one member edge from the hyperedge. The remaining members stay in
the hyperedge, and deleting the last member removes the hyperedge itself.
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
client.Graph.Hyperedge.DeleteEdge(
    context.TODO(),
    "graph_uuid",
    "hyperedge_uuid",
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**hyperedgeUUID:** `string` — Hyperedge UUID
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the nodes of a graph and returns one page of results. The default `page_size` is 50, and the maximum `page_size` is 100.
You can sort by `uuid`, `created_at`, or `degree`. The default sort is `uuid` in descending order.
You cannot use these `filters` fields: `connected_node_uuids`, `source_node_uuids`, `target_node_uuids`, and `content_policy_violated`.
If the request has an unsupported or unknown filter field, the API returns an `unsupported_filter` error.
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
request := &graph.NodeListRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.NodeListRequestOrderBy` — Sort key: uuid (default), created_at, or degree
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.NodeListRequestOrder` — Sort direction: asc or desc (default desc)
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.NodeListFilters` 

Filters that limit the items in the list. The operation description
names the filter fields that the operation does not support.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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
request := &graph.NeighborsRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**nodeUUID:** `string` — Node UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.NodeListNeighborsRequestOrderBy` — Sort field
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.NodeListNeighborsRequestOrder` — Sort direction: asc or desc
    
</dd>
</dl>

<dl>
<dd>

**direction:** `*graph.NeighborsRequestDirection` — The edge orientation to follow from the node: in, out, or both.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the observations of a graph and returns one page of results. The default `page_size` is 50, and the maximum `page_size` is 100.
You can sort by `uuid` or `created_at`. The default sort is `uuid` in descending order.
You cannot use these `filters` fields: `node_labels`, `exclude_node_labels`, `edge_types`, `exclude_edge_types`, `edge_uuids`, `episode_uuids`, `connected_node_uuids`, `source_node_uuids`, `target_node_uuids`, `metadata_filters`, and `content_policy_violated`.
`date_filters` supports only the `created_at` field.
If the request has an unsupported or unknown filter field, the API returns an `unsupported_filter` error.
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
request := &graph.ObservationListRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.ObservationListRequestOrderBy` — Sort key: uuid (default) or created_at
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.ObservationListRequestOrder` — Sort direction: asc or desc (default desc)
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.ObservationListFilters` 

Filters that limit the items in the list. The operation description
names the filter fields that the operation does not support.
    
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
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

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists the thread summaries of a graph and returns one page of results. The default `page_size` is 50, and the maximum `page_size` is 100.
You can sort by `uuid` or `created_at`. The default sort is `uuid` in descending order.
You cannot use these `filters` fields: `node_labels`, `exclude_node_labels`, `edge_types`, `exclude_edge_types`, `edge_uuids`, `episode_uuids`, `connected_node_uuids`, `source_node_uuids`, `target_node_uuids`, and `content_policy_violated`.
`date_filters` supports only the `created_at` field.
If the request has an unsupported or unknown filter field, the API returns an `unsupported_filter` error.
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
request := &graph.ThreadSummaryListRequest{}
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

**graphUUID:** `string` — The UUID of the graph. For a user graph, you can also use the UUID of the user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*graph.ThreadSummaryListRequestOrderBy` — Sort key: uuid (default) or created_at
    
</dd>
</dl>

<dl>
<dd>

**order:** `*graph.ThreadSummaryListRequestOrder` — Sort direction: asc or desc (default desc)
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*zep.ThreadSummaryListFilters` 

Filters that limit the items in the list. The operation description
names the filter fields that the operation does not support.
    
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

**metadata:** `map[string]any` — Metadata to merge onto the message; a key set to null is removed. Max 10 keys after the merge. Values must be strings, numbers, booleans, or arrays of scalars.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TraceConnection Project
<details><summary><code>client.TraceConnection.Project.List(ConnectionUUID) -> *zep.TraceProviderProjectPage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List provider projects with the `page_size` and `cursor` query parameters.
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
request := &traceconnection.ProjectListRequest{}
client.TraceConnection.Project.List(
    context.TODO(),
    "connection_uuid",
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

**connectionUUID:** `string` — Trace connection UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Provider page cursor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TraceConnection Trace
<details><summary><code>client.TraceConnection.Trace.Get(ConnectionUUID, request) -> *zep.SourceTraceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read one trace and optionally preview a mapping. Example body: `{"provider_project_id":"project_123","trace_id":"trace-123","mapping":{"task_family":{"source":"fixed","value":"support"}}}`.
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
request := &traceconnection.SourceTraceGetRequest{
    ProviderProjectID: "project_123",
    TraceID: "trace_123",
}
client.TraceConnection.Trace.Get(
    context.TODO(),
    "connection_uuid",
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

**connectionUUID:** `string` — Trace connection UUID
    
</dd>
</dl>

<dl>
<dd>

**mapping:** `*zep.TraceMapping` — Mapping is an optional mapping preview configuration. Example: {"task_family":{"source":"fixed","value":"support"}}.
    
</dd>
</dl>

<dl>
<dd>

**providerProjectID:** `string` — ProviderProjectID is the provider project identifier.
    
</dd>
</dl>

<dl>
<dd>

**traceID:** `string` — TraceID is the provider trace identifier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TraceConnection.Trace.List(ConnectionUUID, request) -> *zep.SourceTracePage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Filter provider traces. Example body: `{"provider_project_id":"project_123","filter":{"started_after":"2026-01-01T00:00:00Z"}}`.
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
request := &traceconnection.SourceTraceListRequest{
    ProviderProjectID: "project_123",
}
client.TraceConnection.Trace.List(
    context.TODO(),
    "connection_uuid",
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

**connectionUUID:** `string` — Trace connection UUID
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of items in one page. The SDK method returns a pager. In Python, `.items` holds only the first page. In TypeScript, `.data` holds only the first page. To read every page, iterate the pager: use `for` or `list(...)` in Python, or `for await` in TypeScript. In Go, the method returns one page, and `Iterator()` reads all pages.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque page cursor
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*zep.TraceFilter` — Filter contains provider trace filters. Example: {"started_after":"2026-01-01T00:00:00Z"}.
    
</dd>
</dl>

<dl>
<dd>

**providerProjectID:** `string` — ProviderProjectID is the provider project identifier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

