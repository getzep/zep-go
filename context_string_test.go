package zep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatEdgeDateRange(t *testing.T) {
	t.Run("both dates present", func(t *testing.T) {
		validAt := "2023-01-01T10:00:00Z"
		invalidAt := "2023-01-02T15:00:00Z"
		edge := &EntityEdge{
			ValidAt:   &validAt,
			InvalidAt: &invalidAt,
		}
		result := formatEdgeDateRange(edge)
		assert.Equal(t, "2023-01-01 10:00:00 - 2023-01-02 15:00:00", result)
	})

	t.Run("only valid_at present", func(t *testing.T) {
		validAt := "2023-01-01T10:00:00Z"
		edge := &EntityEdge{
			ValidAt: &validAt,
		}
		result := formatEdgeDateRange(edge)
		assert.Equal(t, "2023-01-01 10:00:00 - present", result)
	})

	t.Run("no dates present", func(t *testing.T) {
		edge := &EntityEdge{}
		result := formatEdgeDateRange(edge)
		assert.Equal(t, "date unknown - present", result)
	})
}

func TestComposeContextString(t *testing.T) {
	t.Run("with facts, entities, and episodes", func(t *testing.T) {
		validAt := "2023-01-01T10:00:00Z"
		edges := []*EntityEdge{
			{
				Fact:    "John likes coffee",
				ValidAt: &validAt,
			},
		}

		labels := []string{"Person", "Entity"}
		attributes := map[string]interface{}{
			"age":    30,
			"labels": []string{"Person"},
		}
		nodes := []*EntityNode{
			{
				Name:       "John",
				Labels:     labels,
				Attributes: attributes,
				Summary:    "A coffee enthusiast",
			},
		}

		role := "user"
		roleType := RoleTypeUserRole
		episodes := []*Episode{
			{
				Role:      &role,
				RoleType:  &roleType,
				Content:   "I love coffee",
				CreatedAt: "2023-01-01T12:00:00Z",
			},
		}

		result := ComposeContextString(edges, nodes, episodes)

		assert.Contains(t, result, "FACTS and ENTITIES, and EPISODES represent")
		assert.Contains(t, result, "John likes coffee (2023-01-01 10:00:00 - present)")
		assert.Contains(t, result, "Name: John")
		assert.Contains(t, result, "Label: Person")
		assert.Contains(t, result, "Attributes:")
		assert.Contains(t, result, "  age: 30")
		assert.Contains(t, result, "Summary: A coffee enthusiast")
		assert.Contains(t, result, "user (user): I love coffee (2023-01-01 12:00:00)")
		assert.Contains(t, result, "<EPISODES>")
	})

	t.Run("without episodes", func(t *testing.T) {
		validAt := "2023-01-01T10:00:00Z"
		edges := []*EntityEdge{
			{
				Fact:    "John likes coffee",
				ValidAt: &validAt,
			},
		}

		nodes := []*EntityNode{
			{
				Name:    "John",
				Summary: "A person",
			},
		}

		result := ComposeContextString(edges, nodes, nil)

		assert.Contains(t, result, "FACTS and ENTITIES represent")
		assert.NotContains(t, result, ", and EPISODES")
		assert.NotContains(t, result, "<EPISODES>")
		assert.Contains(t, result, "John likes coffee")
		assert.Contains(t, result, "Name: John")
		assert.Contains(t, result, "Summary: A person")
	})

	t.Run("entity with only Entity label filtered out", func(t *testing.T) {
		labels := []string{"Entity"}
		nodes := []*EntityNode{
			{
				Name:    "Test",
				Labels:  labels,
				Summary: "Test entity",
			},
		}

		result := ComposeContextString(nil, nodes, nil)

		assert.Contains(t, result, "Name: Test")
		assert.NotContains(t, result, "Label: Entity")
		assert.Contains(t, result, "Summary: Test entity")
	})

	t.Run("episode with only role_type", func(t *testing.T) {
		roleType := RoleTypeAssistantRole
		episodes := []*Episode{
			{
				RoleType:  &roleType,
				Content:   "Hello there",
				CreatedAt: "2023-01-01T12:00:00Z",
			},
		}

		result := ComposeContextString(nil, nil, episodes)

		assert.Contains(t, result, "(assistant): Hello there")
	})

	t.Run("empty inputs", func(t *testing.T) {
		result := ComposeContextString(nil, nil, nil)

		assert.Contains(t, result, "FACTS and ENTITIES represent")
		assert.NotContains(t, result, "EPISODES")
		assert.Contains(t, result, "<FACTS>\n\n</FACTS>")
		assert.Contains(t, result, "<ENTITIES>\n\n</ENTITIES>")
	})
}
