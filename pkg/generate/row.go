package generate

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"

	"github.com/giantswarm/schemadocs/pkg/key"
)

type Row struct {
	Path               string
	Name               string
	FullPath           string
	Title              string
	Slug               string
	Description        string
	Types              []string
	KeyPatterns        []string
	KeyPatternMappings map[string]string
	ValuePattern       string
	DefaultValue       string
	ConstValue         *string
	EnumValues         []string
	Examples           []string
	Required           bool
	Primitive          bool
	Presentable        bool
}

func RowsFromSchema(schema *jsonschema.Schema, path string, name string, keyPatterns []string) []Row {
	// Sorting happens outside of `RowsFromSchema`, so the order in this slice doesn't matter
	var rows []Row

	if schema.Ref != nil || schema.RecursiveRef != nil {
		if schema.Ref != nil {
			rows = append(rows, RowsFromSchema(schema.Ref, path, name, keyPatterns)...)
		}
		if schema.RecursiveRef != nil {
			rows = append(rows, RowsFromSchema(schema.RecursiveRef, path, name, keyPatterns)...)
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
			var patternRegex *regexp.Regexp
			if additionalProperties.Pattern != nil {
				keyPatterns = append(keyPatterns, additionalProperties.Pattern.String())
				patternRegex = regexp.MustCompile(additionalProperties.Pattern.String())
			}
			additionalPropertyKey := key.NameFromPattern(patternRegex, keyPatterns, "*")
			rows = append(rows, RowsFromSchema(additionalProperties, row.FullPath, additionalPropertyKey, keyPatterns)...)
		}
	}

	if schema.PatternProperties != nil {
		// Sort pattern keys for deterministic output
		var patternStrings []string
		for pattern := range schema.PatternProperties {
			patternStrings = append(patternStrings, pattern.String())
		}
		sort.Strings(patternStrings)

		for _, patternStr := range patternStrings {
			// Find the corresponding pattern and property
			for pattern, patternProperty := range schema.PatternProperties {
				if pattern.String() == patternStr {
					keyPatterns = append(keyPatterns, pattern.String())
					patternRegex := regexp.MustCompile(pattern.String())
					patternName := key.NameFromPattern(patternRegex, keyPatterns, "*")
					rows = append(rows, RowsFromSchema(patternProperty, row.FullPath, patternName, keyPatterns)...)
					break
				}
			}
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

	for oneOfIndex, oneOfSchema := range schema.OneOf {
		rows = append(rows, RowsFromSchema(oneOfSchema, path, fmt.Sprintf("%s[option#%d]", name, oneOfIndex+1), keyPatterns)...)
	}

	for _, allOfSchema := range schema.AllOf {
		// Exclude first generated row to avoid repetition (row would be equal to the parent of `allOf`)
		rows = append(rows, RowsFromSchema(allOfSchema, path, name, keyPatterns)[1:]...)
	}

	return rows
}

func stringFromAny(a any) string {
	if stringer, ok := a.(fmt.Stringer); ok {
		return stringer.String()
	} else if s, ok := a.(string); ok {
		return s
	} else if v, ok := a.(float64); ok {
		return strconv.FormatFloat(v, 'f', -1, 64)
	} else {
		panic(fmt.Sprintf("Unsupported any type: %T", a))
	}
}

func NewRow(schema *jsonschema.Schema, path string, name string, keyPatterns []string) Row {
	keyPatternMappings := make(map[string]string)
	for _, keyPattern := range keyPatterns {
		keyPatternMappings[keyPattern] = key.NameFromPatternString(keyPattern, keyPatterns)
	}

	var types []string
	if schema.Types != nil {
		types = schema.Types.ToStrings()
	}

	row := Row{
		Path:               path,
		Name:               name,
		FullPath:           key.MergedPropertyPath(path, name),
		Title:              schema.Title,
		Slug:               strings.ToLower(strings.ReplaceAll(key.MergedPropertyPath(path, name), ".", "-")),
		Description:        schema.Description,
		Types:              types,
		Primitive:          key.SchemaIsPrimitive(schema),
		KeyPatterns:        keyPatterns,
		KeyPatternMappings: keyPatternMappings,
	}

	row.Presentable = (row.Primitive || row.Path != "" && row.Path != key.GlobalPropertyName) &&
		row.Name != "" &&
		schema.AllOf == nil

	if schema.Const != nil {
		s := stringFromAny(*schema.Const)
		row.ConstValue = &s
	}

	if schema.Enum != nil {
		for _, enumValue := range schema.Enum.Values {
			row.EnumValues = append(row.EnumValues, stringFromAny(enumValue))
		}
	}

	if schema.Pattern != nil {
		row.ValuePattern = schema.Pattern.String()
	}

	if schema.Examples != nil {
		var examples []string
		for _, example := range schema.Examples {
			exampleJson, err := json.Marshal(example)
			if err != nil {
				examples = append(examples, stringFromAny(example))
			} else {
				examples = append(examples, string(exampleJson))
			}
		}
		row.Examples = examples
	}

	if schema.Default != nil {
		defaultJson, err := json.Marshal(schema.Default)
		if err != nil {
			row.DefaultValue = stringFromAny(schema.Default)
		} else {
			row.DefaultValue = string(defaultJson)
		}
	}

	return row
}
