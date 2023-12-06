package generate

import (
	"encoding/json"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemadocs/pkg/key"
)

type Row struct {
	Path               string
	Name               string
	FullPath           string
	Title              string
	Description        string
	Types              []string
	KeyPatterns        []string
	KeyPatternMappings map[string]string
	ValuePattern       string
	DefaultValue       string
	Examples           []string
	Required           bool
	Primitive          bool
	Presentable        bool
}

func RowsFromSchema(schema *jsonschema.Schema, path string, name string, keyPatterns []string) []Row {
	// Sorting happens outside of `RowsFromSchema`, so the order in this slice doesn't matter
	var rows []Row

	if schema.Ref != nil || schema.RecursiveRef != nil || schema.DynamicRef != nil {
		if schema.Ref != nil {
			rows = append(rows, RowsFromSchema(schema.Ref, path, name, keyPatterns)...)
		}
		if schema.RecursiveRef != nil {
			rows = append(rows, RowsFromSchema(schema.RecursiveRef, path, name, keyPatterns)...)
		}
		if schema.DynamicRef != nil {
			rows = append(rows, RowsFromSchema(schema.DynamicRef, path, name, keyPatterns)...)
		}

		return rows
	}

	row := NewRow(schema, path, name, keyPatterns)
	if row.Presentable {
		rows = append(rows, row)
	}

	for propertyName, property := range schema.Properties {
		rows = append(rows, RowsFromSchema(property, row.FullPath, propertyName, keyPatterns)...)
	}

	if schema.AdditionalProperties != nil {
		additionalProperties, ok := schema.AdditionalProperties.(*jsonschema.Schema)
		if ok {
			if additionalProperties.Pattern != nil {
				keyPatterns = append(keyPatterns, additionalProperties.Pattern.String())
			}
			additionalPropertyKey := key.NameFromPattern(additionalProperties.Pattern, keyPatterns, "*")
			rows = append(rows, RowsFromSchema(additionalProperties, row.FullPath, additionalPropertyKey, keyPatterns)...)
		}
	}

	if schema.PatternProperties != nil {
		for pattern, patternProperty := range schema.PatternProperties {
			keyPatterns = append(keyPatterns, pattern.String())
			patternName := key.NameFromPattern(pattern, keyPatterns, "*")
			rows = append(rows, RowsFromSchema(patternProperty, row.FullPath, patternName, keyPatterns)...)
		}
	}

	if schema.Items != nil {
		switch schema.Items.(type) {
		case *jsonschema.Schema:
			rows = append(rows, RowsFromSchema(schema.Items.(*jsonschema.Schema), path, key.ListItemName(name), keyPatterns)...)
		case []*jsonschema.Schema:
			for _, item := range schema.Items.([]*jsonschema.Schema) {
				rows = append(rows, RowsFromSchema(item, path, key.ListItemName(name), keyPatterns)...)
			}
		}
	}

	if schema.Items2020 != nil {
		rows = append(rows, RowsFromSchema(schema.Items2020, path, key.ListItemName(name), keyPatterns)...)
	}

	return rows
}

func NewRow(schema *jsonschema.Schema, path string, name string, keyPatterns []string) Row {
	keyPatternMappings := make(map[string]string)
	for _, keyPattern := range keyPatterns {
		keyPatternMappings[keyPattern] = key.NameFromPatternString(keyPattern, keyPatterns)
	}

	row := Row{
		Path:               path,
		Name:               name,
		FullPath:           key.MergedPropertyPath(path, name),
		Title:              schema.Title,
		Description:        schema.Description,
		Types:              schema.Types,
		Primitive:          key.SchemaIsPrimitive(schema),
		KeyPatterns:        keyPatterns,
		KeyPatternMappings: keyPatternMappings,
	}

	row.Presentable = (row.Primitive || row.Path != "" && row.Path != key.GlobalPropertyName) && row.Name != ""

	if schema.Pattern != nil {
		row.ValuePattern = schema.Pattern.String()
	}

	if schema.Examples != nil {
		var examples []string
		for _, example := range schema.Examples {
			exampleJson, err := json.Marshal(example)
			if err != nil {
				examples = append(examples, fmt.Sprintf("%v", example))
			} else {
				examples = append(examples, string(exampleJson))
			}
		}
		row.Examples = examples
	}

	if schema.Default != nil {
		defaultJson, err := json.Marshal(schema.Default)
		if err != nil {
			row.DefaultValue = fmt.Sprintf("%v", schema.Default)
		} else {
			row.DefaultValue = string(defaultJson)
		}
	}

	return row
}
