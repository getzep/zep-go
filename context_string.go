package zep

import (
	"fmt"
	"strings"
	"time"
)

// dateFormat defines the format for date strings
const dateFormat = "2006-01-02 15:04:05"

// templateString defines the template for context information
const templateString = `
FACTS and ENTITIES represent relevant context to the current conversation.

# These are the most relevant facts and their valid date ranges
# format: FACT (Date range: from - to)
<FACTS>
%s
</FACTS>

# These are the most relevant entities
# ENTITY_NAME: entity summary
<ENTITIES>
%s
</ENTITIES>
`

// formatEdgeDateRange formats the date range of an entity edge.
func formatEdgeDateRange(edge *EntityEdge) string {
	validAt := "date unknown"
	invalidAt := "present"

	if edge.ValidAt != nil && *edge.ValidAt != "" {
		if t, err := time.Parse(time.RFC3339, *edge.ValidAt); err == nil {
			validAt = t.Format(dateFormat)
		}
	}

	if edge.InvalidAt != nil && *edge.InvalidAt != "" {
		if t, err := time.Parse(time.RFC3339, *edge.InvalidAt); err == nil {
			invalidAt = t.Format(dateFormat)
		}
	}

	return fmt.Sprintf("%s - %s", validAt, invalidAt)
}

// ComposeContextString composes a search context from entity edges and nodes.
func ComposeContextString(edges []*EntityEdge, nodes []*EntityNode) string {
	var facts []string
	for _, edge := range edges {
		fact := fmt.Sprintf("  - %s (%s)", edge.Fact, formatEdgeDateRange(edge))
		facts = append(facts, fact)
	}

	var entities []string
	for _, node := range nodes {
		entity := fmt.Sprintf("  - %s: %s", node.Name, node.Summary)
		entities = append(entities, entity)
	}

	factsStr := strings.Join(facts, "\n")
	entitiesStr := strings.Join(entities, "\n")

	return fmt.Sprintf(templateString, factsStr, entitiesStr)
}
