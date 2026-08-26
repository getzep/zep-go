package zep

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type traveler struct {
	EntityBase `description:"Someone who takes trips."`

	HomeCity string  `description:"The city they live in" identity:"true"`
	Trips    int     `description:"How many trips they have taken"`
	Rating   float64 `description:"Their average trip rating"`
	Frequent bool    `description:"Whether they travel often"`
}

type traveledTo struct {
	EdgeBase `description:"A traveler visiting a destination."`

	Purpose string `description:"Why they went"`
}

func TestBuildOntologyProducesTheDocumentedPayload(t *testing.T) {
	ontology, err := BuildOntology(
		Entities{"Traveler": traveler{}},
		Edges{
			"TRAVELED_TO": {
				Model: traveledTo{},
				SourceTargets: []*EdgeSourceTarget{
					{SourceEntityType: String("Traveler"), TargetEntityType: String("Destination")},
				},
			},
		},
	)
	require.NoError(t, err)

	body, err := json.Marshal(ontology)
	require.NoError(t, err)

	assert.JSONEq(t, `{
		"entity_types": [{
			"name": "Traveler",
			"description": "Someone who takes trips.",
			"identity_properties": ["HomeCity"],
			"properties": [
				{"name": "HomeCity", "description": "The city they live in", "type": "text"},
				{"name": "Trips", "description": "How many trips they have taken", "type": "int"},
				{"name": "Rating", "description": "Their average trip rating", "type": "float"},
				{"name": "Frequent", "description": "Whether they travel often", "type": "boolean"}
			]
		}],
		"edge_types": [{
			"name": "TRAVELED_TO",
			"description": "A traveler visiting a destination.",
			"properties": [
				{"name": "Purpose", "description": "Why they went", "type": "text"}
			],
			"source_targets": [
				{"source_entity_type": "Traveler", "target_entity_type": "Destination"}
			]
		}]
	}`, string(body))
}

func TestBuildOntologyEmitsTypesInNameOrder(t *testing.T) {
	// Map iteration order is random, so the same declarations have to be sorted
	// into a stable payload rather than whatever order the range produced.
	type destination struct {
		EntityBase `description:"A place a traveler goes."`

		Country string `description:"The country it is in"`
	}

	for i := 0; i < 16; i++ {
		ontology, err := BuildOntology(
			Entities{"Traveler": traveler{}, "Destination": destination{}},
			nil,
		)
		require.NoError(t, err)
		require.Len(t, ontology.EntityTypes, 2)
		assert.Equal(t, "Destination", *ontology.EntityTypes[0].GetName())
		assert.Equal(t, "Traveler", *ontology.EntityTypes[1].GetName())
	}
}

func TestBuildOntologyOmitsIdentityPropertiesWhenNoneAreTagged(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country string `description:"The country it is in"`
	}

	ontology, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.NoError(t, err)

	body, err := json.Marshal(ontology.EntityTypes[0])
	require.NoError(t, err)
	assert.NotContains(t, string(body), "identity_properties")
}

func TestBuildOntologyReadsThePropertyNameOffTheJSONTag(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country string `json:"country" description:"The country it is in" identity:"true"`
	}

	ontology, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.NoError(t, err)
	assert.Equal(t, "country", *ontology.EntityTypes[0].GetProperties()[0].GetName())
	// The identity list has to name the property as sent, not the Go field.
	assert.Equal(t, []string{"country"}, ontology.EntityTypes[0].GetIdentityProperties())
}

func TestBuildOntologySkipsFieldsTaggedOut(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country  string `description:"The country it is in"`
		Internal string `zep:"-"`
	}

	ontology, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.NoError(t, err)
	require.Len(t, ontology.EntityTypes[0].GetProperties(), 1)
	assert.Equal(t, "Country", *ontology.EntityTypes[0].GetProperties()[0].GetName())
}

func TestBuildOntologyRejectsAMisspelledZepTag(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country  string `description:"The country it is in"`
		Internal string `zap:"-"`
	}

	// A misspelled zep tag has to be an error rather than a silent no-op: left
	// alone, Internal would stay in the ontology despite the intent to
	// exclude it, with nothing to say so.
	_, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Place.Internal")
	assert.Contains(t, err.Error(), `"zap"`)
	assert.Contains(t, err.Error(), `"zep"`)
}

func TestBuildOntologyRejectsAMisspelledIdentityTag(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country string `description:"The country it is in" idenity:"true"`
	}

	// A misspelled identity tag has to be an error rather than a silent
	// no-op: left alone, the identity flag would simply be dropped.
	_, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Place.Country")
	assert.Contains(t, err.Error(), `"idenity"`)
	assert.Contains(t, err.Error(), `"identity"`)
}

func TestBuildOntologyRejectsAMisspelledDescriptionTag(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country string `descripton:"The country it is in"`
	}

	// A misspelled description tag has to be reported as the likely typo it
	// is, not as a plain missing description.
	_, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Place.Country")
	assert.Contains(t, err.Error(), `"descripton"`)
	assert.Contains(t, err.Error(), `"description"`)
}

func TestBuildOntologyReadsThroughAPointerField(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Population *int `description:"How many people live there"`
	}

	ontology, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.NoError(t, err)
	assert.Equal(t, EntityPropertyTypeInt, *ontology.EntityTypes[0].GetProperties()[0].GetType())
}

func TestBuildOntologyAcceptsAPointerToADeclaration(t *testing.T) {
	ontology, err := BuildOntology(Entities{"Traveler": &traveler{}}, nil)
	require.NoError(t, err)
	assert.Equal(t, "Traveler", *ontology.EntityTypes[0].GetName())
}

func TestBuildOntologyRejectsAPropertyWithNoDescription(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country string
	}

	_, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.Error(t, err)
	// The message has to name the field, which is the only way to find it.
	assert.Contains(t, err.Error(), "Place.Country")
}

func TestBuildOntologyRejectsATypeWithNoDescription(t *testing.T) {
	type place struct {
		EntityBase

		Country string `description:"The country it is in"`
	}

	_, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Place")
}

func TestBuildOntologyRejectsAnUnsupportedPropertyType(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Neighbors []string `description:"Places next to it"`
	}

	_, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Place.Neighbors")
}

func TestBuildOntologyRejectsAnEntityDeclaredAsAnEdge(t *testing.T) {
	// EntityBase and EdgeBase are what keep the two apart; passing one where the
	// other belongs is a mistake worth a message rather than an empty type.
	_, err := BuildOntology(nil, Edges{"TRAVELED_TO": {Model: traveler{}}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must embed zep.EdgeBase")
}

func TestBuildOntologyRejectsANonBoolIdentityTag(t *testing.T) {
	type place struct {
		EntityBase `description:"A place."`

		Country string `description:"The country it is in" identity:"yes please"`
	}

	_, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "identity tag must be true or false")
}

func TestBuildOntologyRejectsANonStruct(t *testing.T) {
	_, err := BuildOntology(Entities{"Place": "not a struct"}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must be a struct")
}

func TestBuildOntologyRejectsANilDeclaration(t *testing.T) {
	_, err := BuildOntology(Entities{"Place": nil}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "is nil")
}

func TestBuildOntologyReturnsAnEmptyOntologyForNoTypes(t *testing.T) {
	// Sending an empty ontology is how a graph is handed back to the project
	// default, so it has to build rather than error.
	ontology, err := BuildOntology(nil, nil)
	require.NoError(t, err)
	assert.Empty(t, ontology.GetEntityTypes())
	assert.Empty(t, ontology.GetEdgeTypes())
}

func TestBuildOntologyPromotesPropertiesFromAnEmbeddedStruct(t *testing.T) {
	// Shared properties are declared once and embedded, which is what a reader
	// coming from the Python DSL's class inheritance expects to work.
	type audited struct {
		RecordedBy string `description:"Who recorded it"`
	}
	type place struct {
		EntityBase `description:"A place."`
		audited

		Country string `description:"The country it is in" identity:"true"`
	}

	ontology, err := BuildOntology(Entities{"Place": place{}}, nil)
	require.NoError(t, err)

	properties := ontology.EntityTypes[0].GetProperties()
	require.Len(t, properties, 2)
	assert.Equal(t, "RecordedBy", *properties[0].GetName())
	assert.Equal(t, "Country", *properties[1].GetName())
	assert.Equal(t, []string{"Country"}, ontology.EntityTypes[0].GetIdentityProperties())
}

func TestBuildOntologyRejectsAnEdgeDeclaredAsAnEntity(t *testing.T) {
	_, err := BuildOntology(Entities{"Traveler": traveledTo{}}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must embed zep.EntityBase")
}
