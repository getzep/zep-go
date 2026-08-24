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
FACTS and ENTITIES%s represent relevant context to the current conversation.

# These are the most relevant facts and their valid date ranges
# format: FACT (Date range: from - to)
<FACTS>
%s
</FACTS>

# These are the most relevant entities
# Name: ENTITY_NAME
# Label: entity_label (if present)
# Attributes: (if present)
#   attr_name: attr_value
# Summary: entity summary
<ENTITIES>
%s
</ENTITIES>
%s
`

// deref returns the pointed-to string, or "" when the pointer is nil. Every v4
// response field is optional in the generated Go client, so the composer reads
// them through this rather than dereferencing directly.
func deref(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// formatEdgeDateRange formats the date range of an entity edge.
func formatEdgeDateRange(edge *Edge) string {
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

// ComposeContextString composes a search context from entity edges, nodes, and episodes.
func ComposeContextString(edges []*Edge, nodes []*Node, episodes []*Episode) string {
	var facts []string
	for _, edge := range edges {
		fact := fmt.Sprintf("  - %s (%s)", deref(edge.Fact), formatEdgeDateRange(edge))
		facts = append(facts, fact)
	}

	var entities []string
	for _, node := range nodes {
		var entityParts []string
		entityParts = append(entityParts, fmt.Sprintf("Name: %s", deref(node.Name)))

		// Add label if present (excluding 'Entity' from labels)
		if node.Labels != nil && len(node.Labels) > 0 {
			labels := make([]string, 0, len(node.Labels))
			for _, label := range node.Labels {
				if label != "Entity" {
					labels = append(labels, label)
				}
			}
			if len(labels) > 0 {
				entityParts = append(entityParts, fmt.Sprintf("Label: %s", labels[0]))
			}
		}

		// Add attributes if present (excluding 'labels' attribute)
		if node.Attributes != nil && len(node.Attributes) > 0 {
			hasNonLabelAttributes := false
			for key := range node.Attributes {
				if key != "labels" {
					hasNonLabelAttributes = true
					break
				}
			}
			if hasNonLabelAttributes {
				entityParts = append(entityParts, "Attributes:")
				for key, value := range node.Attributes {
					if key != "labels" {
						entityParts = append(entityParts, fmt.Sprintf("  %s: %v", key, value))
					}
				}
			}
		}

		// Add summary if present
		if summary := deref(node.Summary); summary != "" {
			entityParts = append(entityParts, fmt.Sprintf("Summary: %s", summary))
		}

		entity := strings.Join(entityParts, "\n")
		entities = append(entities, entity)
	}

	// Format episodes
	var episodesList []string
	if episodes != nil {
		for _, episode := range episodes {
			// 8.4 resolves v3's overloaded pair: v4's role carries the enum
			// v3 spelled role_type, and role_name carries the sender name v3
			// spelled role. The rendered prefix is unchanged.
			roleName := deref(episode.RoleName)
			roleType := deref(episode.Role)

			var rolePrefix string
			switch {
			case roleName != "" && roleType != "":
				rolePrefix = fmt.Sprintf("%s (%s): ", roleName, roleType)
			case roleName != "":
				rolePrefix = fmt.Sprintf("%s: ", roleName)
			case roleType != "":
				rolePrefix = fmt.Sprintf("(%s): ", roleType)
			}

			timestamp := "date unknown"
			if createdAt := deref(episode.CreatedAt); createdAt != "" {
				if t, err := time.Parse(time.RFC3339, createdAt); err == nil {
					timestamp = t.Format(dateFormat)
				}
			}

			episodeStr := fmt.Sprintf("  - %s%s (%s)", rolePrefix, deref(episode.Content), timestamp)
			episodesList = append(episodesList, episodeStr)
		}
	}

	factsStr := strings.Join(facts, "\n")
	entitiesStr := strings.Join(entities, "\n")
	episodesStr := strings.Join(episodesList, "\n")

	// Determine if episodes section should be included
	episodesHeader := ""
	episodesSection := ""
	if len(episodesList) > 0 {
		episodesHeader = ", and EPISODES"
		episodesSection = fmt.Sprintf("\n# These are the most relevant episodes\n<EPISODES>\n%s\n</EPISODES>", episodesStr)
	}

	return fmt.Sprintf(templateString, episodesHeader, factsStr, entitiesStr, episodesSection)
}
