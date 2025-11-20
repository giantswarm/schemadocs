package key

import (
	"fmt"
	"regexp"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	GlobalPropertyName = "global"
	ItemNameSuffix     = "[*]"
	OtherSectionTitle  = "Other"
)

func NameFromPattern(pattern *regexp.Regexp, patterns []string, defaultName string) string {
	if pattern == nil {
		return defaultName
	}
	return NameFromPatternString(pattern.String(), patterns)
}

func NameFromPatternString(patternStr string, patterns []string) string {
	for index, patternItem := range patterns {
		if patternStr == patternItem {
			if index == 0 {
				return "PATTERN"
			}
			return fmt.Sprintf("PATTERN_%d", index+1)
		}
	}
	return patternStr
}

func ListItemName(name string) string {
	return name + ItemNameSuffix
}

func MergedPropertyPath(left string, right string) string {
	if left != "" && right != "" {
		return fmt.Sprintf("%s.%s", left, right)
	} else if left != "" {
		return left
	}
	return right
}

func SchemaIsPrimitive(schema *jsonschema.Schema) bool {
	return len(schema.Properties) == 0 &&
		len(schema.PatternProperties) == 0 &&
		(schema.AdditionalProperties == nil || schema.AdditionalProperties == false) &&
		(schema.Items == nil || schema.Items == false) &&
		schema.Items2020 == nil &&
		schema.Ref == nil &&
		schema.DynamicRef == nil &&
		schema.RecursiveRef == nil &&
		schema.OneOf == nil &&
		schema.AllOf == nil
}

func SchemaIsPresentable(schema *jsonschema.Schema) bool {
	return SchemaIsPrimitive(schema) ||
		schema.Title != "" ||
		schema.Items2020 != nil ||
		(schema.Items != nil && schema.Items != false) ||
		(schema.AdditionalProperties != nil && schema.AdditionalProperties != false)
}
