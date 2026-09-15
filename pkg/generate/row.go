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

// applyAnnotations overwrites the annotations that schema sets itself, and leaves the others at the
// value the row already carries.
func applyAnnotations(row *Row, schema *jsonschema.Schema) {
	if schema.Title != "" {
		row.Title = schema.Title
	}
	if schema.Description != "" {
		row.Description = schema.Description
	}
	if schema.Default != nil {
		row.DefaultValue = defaultValueFromSchema(schema.Default)
	}
	if schema.Examples != nil {
		row.Examples = examplesFromSchema(schema.Examples)
	}
}

func defaultValueFromSchema(defaultValue *any) string {
	defaultJson, err := json.Marshal(defaultValue)
	if err != nil {
		return stringFromAny(defaultValue)
	}
	return string(defaultJson)
}

func examplesFromSchema(schemaExamples []any) []string {
	var examples []string
	for _, example := range schemaExamples {
		exampleJson, err := json.Marshal(example)
		if err != nil {
			examples = append(examples, stringFromAny(example))
		} else {
			examples = append(examples, string(exampleJson))
		}
	}
	return examples
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

		// JSON Schema 2020-12 allows annotations next to `$ref`, where they describe the referencing
		// property rather than the referenced schema. Applying them lets one definition be reused in
		// several places that need their own title, description, default or examples. Only the row of
		// the referencing property itself is affected, not the rows of the referenced schema's own
		// properties.
		fullPath := key.MergedPropertyPath(path, name)
		for i := range rows {
			if rows[i].FullPath == fullPath {
				applyAnnotations(&rows[i], schema)
			}
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
		row.Examples = examplesFromSchema(schema.Examples)
	}

	if schema.Default != nil {
		row.DefaultValue = defaultValueFromSchema(schema.Default)
	}

	return row
}
