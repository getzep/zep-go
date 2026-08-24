package zep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatEdgeDateRange(t *testing.T) {
	t.Run("both dates present", func(t *testing.T) {
		edge := &Edge{
			ValidAt:   String("2023-01-01T10:00:00Z"),
			InvalidAt: String("2023-01-02T15:00:00Z"),
		}
		result := formatEdgeDateRange(edge)
		assert.Equal(t, "2023-01-01 10:00:00 - 2023-01-02 15:00:00", result)
	})

	t.Run("only valid_at present", func(t *testing.T) {
		edge := &Edge{
			ValidAt: String("2023-01-01T10:00:00Z"),
		}
		result := formatEdgeDateRange(edge)
		assert.Equal(t, "2023-01-01 10:00:00 - present", result)
	})

	t.Run("no dates present", func(t *testing.T) {
		edge := &Edge{}
		result := formatEdgeDateRange(edge)
		assert.Equal(t, "date unknown - present", result)
	})
}

func TestComposeContextString(t *testing.T) {
	t.Run("with facts, entities, and episodes", func(t *testing.T) {
		edges := []*Edge{
			{
				Fact:    String("John likes coffee"),
				ValidAt: String("2023-01-01T10:00:00Z"),
			},
		}

		nodes := []*Node{
			{
				Name:   String("John"),
				Labels: []string{"Person", "Entity"},
				Attributes: map[string]interface{}{
					"age":    30,
					"labels": []string{"Person"},
				},
				Summary: String("A coffee enthusiast"),
			},
		}

		// v4 splits v3's overloaded pair (spec-3 8.4): role_name is the sender
		// and role is the enum v3 spelled role_type. The rendering is unchanged.
		episodes := []*Episode{
			{
				RoleName:  String("user"),
				Role:      String("user"),
				Content:   String("I love coffee"),
				CreatedAt: String("2023-01-01T12:00:00Z"),
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
		edges := []*Edge{
			{
				Fact:    String("John likes coffee"),
				ValidAt: String("2023-01-01T10:00:00Z"),
			},
		}

		nodes := []*Node{
			{
				Name:    String("John"),
				Summary: String("A person"),
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
		nodes := []*Node{
			{
				Name:    String("Test"),
				Labels:  []string{"Entity"},
				Summary: String("Test entity"),
			},
		}

		result := ComposeContextString(nil, nodes, nil)

		assert.Contains(t, result, "Name: Test")
		assert.NotContains(t, result, "Label: Entity")
		assert.Contains(t, result, "Summary: Test entity")
	})

	t.Run("episode with only a role and no sender name", func(t *testing.T) {
		episodes := []*Episode{
			{
				Role:      String("assistant"),
				Content:   String("Hello there"),
				CreatedAt: String("2023-01-01T12:00:00Z"),
			},
		}

		result := ComposeContextString(nil, nil, episodes)

		assert.Contains(t, result, "(assistant): Hello there")
	})

	t.Run("episode with only a sender name and no role", func(t *testing.T) {
		episodes := []*Episode{
			{
				RoleName:  String("Alice"),
				Content:   String("Hello there"),
				CreatedAt: String("2023-01-01T12:00:00Z"),
			},
		}

		result := ComposeContextString(nil, nil, episodes)

		assert.Contains(t, result, "Alice: Hello there")
	})

	t.Run("empty inputs", func(t *testing.T) {
		result := ComposeContextString(nil, nil, nil)

		assert.Contains(t, result, "FACTS and ENTITIES represent")
		assert.NotContains(t, result, "EPISODES")
		assert.Contains(t, result, "<FACTS>\n\n</FACTS>")
		assert.Contains(t, result, "<ENTITIES>\n\n</ENTITIES>")
	})
}
