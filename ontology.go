package zep

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Declaring an ontology with structs
//
// Graph.SetOntology and Project.SetOntology accept an *Ontology holding
// []*EntityType and []*EdgeType. Building those by hand means repeating every
// property's name, type and description as data. The declarations below let an
// ontology be written once, as structs, and derive the payload from them:
//
//	type Traveler struct {
//	    zep.EntityBase `description:"Someone who takes trips."`
//
//	    HomeCity string `description:"The city they live in" identity:"true"`
//	    Trips    int    `description:"How many trips they have taken"`
//	}
//
//	type TraveledTo struct {
//	    zep.EdgeBase `description:"A traveler visiting a destination."`
//
//	    Purpose string `description:"Why they went"`
//	}
//
//	ontology, err := zep.BuildOntology(
//	    zep.Entities{"Traveler": Traveler{}},
//	    zep.Edges{
//	        "TRAVELED_TO": {
//	            Model: TraveledTo{},
//	            SourceTargets: []*zep.EdgeSourceTarget{
//	                {SourceEntityType: zep.String("Traveler"), TargetEntityType: zep.String("Destination")},
//	            },
//	        },
//	    },
//	)
//	if err != nil {
//	    return err
//	}
//	client.Graph.SetOntology(ctx, graphUUID, ontology)
//
// The same *Ontology goes to Project.SetOntology for the project default. v3
// addressed many graphs in one call; v4 has one ontology endpoint per scope, so
// a caller targeting several graphs sends the same payload once per graph.
//
// These are plain functions rather than client methods on purpose: the
// generated clients already define SetOntology, so extending them collides.

// EntityBase is embedded in a struct to declare it an entity type. Its
// description tag describes the type, which is what the extraction model reads
// to decide whether a node belongs to it, so the tag is required.
type EntityBase struct{}

// EdgeBase is embedded in a struct to declare it an edge type. Its description
// tag is required, for the same reason as EntityBase's.
type EdgeBase struct{}

// Entities maps an entity type name to a struct that declares it. The API
// expects PascalCase names.
type Entities map[string]interface{}

// EdgeDeclaration pairs an edge model with the entity type pairs the edge is
// allowed to connect. Leave SourceTargets nil to allow any pair.
type EdgeDeclaration struct {
	Model         interface{}
	SourceTargets []*EdgeSourceTarget
}

// Edges maps an edge type name to its declaration. The API expects
// SCREAMING_SNAKE_CASE names.
type Edges map[string]EdgeDeclaration

// BuildOntology derives an *Ontology from the given entity and edge structs.
//
// Every exported field of a declaration becomes a property, and needs a
// description tag. A field tagged `zep:"-"` is left out. A field tagged
// `identity:"true"` is also listed in the type's identity properties, which is
// how the extraction model tells two nodes of the same type apart.
//
// Types are emitted in name order, so the same declarations always produce the
// same payload.
func BuildOntology(entities Entities, edges Edges) (*Ontology, error) {
	ontology := &Ontology{}

	for _, name := range sortedNames(len(entities), func(add func(string)) {
		for name := range entities {
			add(name)
		}
	}) {
		declaration, err := readDeclaration(entities[name], entityBaseType, name)
		if err != nil {
			return nil, err
		}
		ontology.EntityTypes = append(ontology.EntityTypes, &EntityType{
			Name:               String(name),
			Description:        String(declaration.description),
			Properties:         declaration.properties,
			IdentityProperties: declaration.identityProperties,
		})
	}

	for _, name := range sortedNames(len(edges), func(add func(string)) {
		for name := range edges {
			add(name)
		}
	}) {
		declaration, err := readDeclaration(edges[name].Model, edgeBaseType, name)
		if err != nil {
			return nil, err
		}
		ontology.EdgeTypes = append(ontology.EdgeTypes, &EdgeType{
			Name:          String(name),
			Description:   String(declaration.description),
			Properties:    declaration.properties,
			SourceTargets: edges[name].SourceTargets,
		})
	}

	return ontology, nil
}

// declaration is what one struct says about the type it declares.
type declaration struct {
	description        string
	properties         []*EntityProperty
	identityProperties []string
}

// The two marker types, compared by reflect.Type so that a struct of the same
// name declared elsewhere is not mistaken for one of them.
var (
	entityBaseType = reflect.TypeOf(EntityBase{})
	edgeBaseType   = reflect.TypeOf(EdgeBase{})
)

// readDeclaration reads a type declaration off a struct by reflection. marker is
// the base struct the declaration must embed, which is what separates an entity
// declaration from an edge one.
func readDeclaration(model interface{}, marker reflect.Type, name string) (declaration, error) {
	var read declaration

	if strings.TrimSpace(name) == "" {
		return read, fmt.Errorf("a type name cannot be empty")
	}

	structType := reflect.TypeOf(model)
	if structType == nil {
		return read, fmt.Errorf("%s: declaration is nil", name)
	}
	for structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	if structType.Kind() != reflect.Struct {
		return read, fmt.Errorf("%s: declaration must be a struct, got %s", name, structType.Kind())
	}

	// The marker is read before the fields so that a declaration embedding the
	// wrong one is reported as such, rather than as a bad property.
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Anonymous && field.Type == marker {
			read.description = strings.TrimSpace(field.Tag.Get("description"))
			break
		}
	}
	if read.description == "" {
		if !embedsMarker(structType, marker) {
			return declaration{}, fmt.Errorf("%s: declaration must embed zep.%s", name, marker.Name())
		}
		return declaration{}, fmt.Errorf(
			`%s: the embedded zep.%s needs a description tag; the extraction model reads it to decide what belongs to this type`,
			name,
			marker.Name(),
		)
	}

	properties, identityProperties, err := readFields(structType, name)
	if err != nil {
		return declaration{}, err
	}
	read.properties = properties
	read.identityProperties = identityProperties

	return read, nil
}

// embedsMarker reports whether the struct embeds the given marker at all, which
// separates "embedded the wrong marker" from "left the description off".
func embedsMarker(structType, marker reflect.Type) bool {
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Anonymous && field.Type == marker {
			return true
		}
	}
	return false
}

// readFields turns a declaration's exported fields into properties. An embedded
// struct that is not a marker has its own fields promoted, so shared properties
// can be declared once and embedded.
func readFields(structType reflect.Type, name string) ([]*EntityProperty, []string, error) {
	var (
		properties         []*EntityProperty
		identityProperties []string
	)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		if field.Anonymous {
			if field.Type == entityBaseType || field.Type == edgeBaseType {
				continue
			}
			embedded := field.Type
			for embedded.Kind() == reflect.Ptr {
				embedded = embedded.Elem()
			}
			if embedded.Kind() == reflect.Struct {
				embeddedProperties, embeddedIdentity, err := readFields(embedded, name)
				if err != nil {
					return nil, nil, err
				}
				properties = append(properties, embeddedProperties...)
				identityProperties = append(identityProperties, embeddedIdentity...)
				continue
			}
		}
		if !field.IsExported() || field.Tag.Get("zep") == "-" {
			continue
		}

		propertyName := field.Name
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			if first := strings.Split(jsonTag, ",")[0]; first != "" && first != "-" {
				propertyName = first
			}
		}

		description := strings.TrimSpace(field.Tag.Get("description"))
		if description == "" {
			return nil, nil, fmt.Errorf(
				`%s.%s: a property needs a description tag; add description:"..." to the field, or zep:"-" to leave it out of the ontology`,
				name,
				field.Name,
			)
		}

		fieldType := field.Type
		for fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}
		propertyType, ok := propertyTypeForKind(fieldType.Kind())
		if !ok {
			return nil, nil, fmt.Errorf(
				"%s.%s: %s cannot be an ontology property; use a string, an integer, a float, or a bool",
				name,
				field.Name,
				field.Type,
			)
		}

		properties = append(properties, &EntityProperty{
			Name:        String(propertyName),
			Description: String(description),
			Type:        propertyType.Ptr(),
		})

		identity, err := identityTag(field, name)
		if err != nil {
			return nil, nil, err
		}
		if identity {
			identityProperties = append(identityProperties, propertyName)
		}
	}

	return properties, identityProperties, nil
}

// identityTag reports whether the field is tagged as an identity property. A tag
// that is present but not a bool is an error rather than a silent false, which
// is the mistake worth catching.
func identityTag(field reflect.StructField, name string) (bool, error) {
	raw, ok := field.Tag.Lookup("identity")
	if !ok {
		return false, nil
	}
	identity, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf(
			"%s.%s: the identity tag must be true or false, got %q",
			name,
			field.Name,
			raw,
		)
	}
	return identity, nil
}

// propertyTypeForKind maps a Go kind onto the property type the API accepts.
func propertyTypeForKind(kind reflect.Kind) (EntityPropertyType, bool) {
	switch kind {
	case reflect.String:
		return EntityPropertyTypeText, true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return EntityPropertyTypeInt, true
	case reflect.Float32, reflect.Float64:
		return EntityPropertyTypeFloat, true
	case reflect.Bool:
		return EntityPropertyTypeBoolean, true
	default:
		return "", false
	}
}

// sortedNames collects map keys through add and returns them in order, so that
// the same declarations always produce the same payload.
func sortedNames(size int, collect func(add func(string))) []string {
	names := make([]string, 0, size)
	collect(func(name string) {
		names = append(names, name)
	})
	sort.Strings(names)
	return names
}
