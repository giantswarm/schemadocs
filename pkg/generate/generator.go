package generate

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"

	"github.com/giantswarm/microerror"
	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"

	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
	"github.com/giantswarm/schemadocs/pkg/generate/templates"
	"github.com/giantswarm/schemadocs/pkg/key"
)

func Generate(schemaPath string, layout string) (string, error) {
	compiler := jsonschema.NewCompiler()
	compiler.ExtractAnnotations = true
	schema, err := compiler.Compile(schemaPath)

	if err != nil {
		return "", microerror.Maskf(pkgerror.InvalidSchemaFile, "%s", err.Error())
	}

	sections := sectionsFromSchema(schema, "")
	return toMarkdown(sections, layout)
}

func sectionsFromSchema(schema *jsonschema.Schema, path string) []Section {
	var sections []Section
	var otherSectionRows []Row

	if key.SchemaIsPrimitive(schema) {
		otherSectionRows = append(otherSectionRows, RowsFromSchema(schema, path, schema.Title, []string{})...)
	} else if key.SchemaIsPresentable(schema) && path != key.GlobalPropertyName {
		sections = append(sections, SectionFromSchema(schema, path, schema.Title))
	} else {
		for name, property := range schema.Properties {
			if key.SchemaIsPrimitive(property) {
				otherSectionRows = append(otherSectionRows, RowsFromSchema(property, path, name, []string{})...)
			} else if name == key.GlobalPropertyName {
				globalSections := sectionsFromSchema(property, key.GlobalPropertyName)
				sections = append(sections, globalSections...)
			} else {
				sections = append(sections, SectionFromSchema(property, path, name))
			}
		}
	}

	sort.SliceStable(sections, func(i, j int) bool {
		if sections[i].Title == sections[j].Title {
			return sections[i].Path < sections[j].Path
		}
		return sections[i].Title < sections[j].Title
	})

	if len(otherSectionRows) > 0 {
		otherSectionRows = sortedRows(otherSectionRows)

		var otherSectionTitle string
		if path != "" {
			otherSectionTitle = fmt.Sprintf("%s %s", key.OtherSectionTitle, path)
		} else {
			otherSectionTitle = key.OtherSectionTitle
		}

		section := NewSection(otherSectionTitle, otherSectionRows)
		sections = append(sections, section)
	}

	return sections
}

func toMarkdown(sections []Section, layout string) (string, error) {
	var t *template.Template
	var err error
	if layout == "linear" {
		t, err = templates.GetLinearTemplate()
	} else {
		t, err = templates.GetDefaultTemplate()
	}
	if err != nil {
		return "", microerror.Maskf(pkgerror.DocsGenerationError, "%s", err.Error())
	}

	var tplBuffer bytes.Buffer
	err = t.Execute(&tplBuffer, sections)
	if err != nil {
		return "", microerror.Maskf(pkgerror.DocsGenerationError, "%s", err.Error())
	}
	return tplBuffer.String(), nil
}
