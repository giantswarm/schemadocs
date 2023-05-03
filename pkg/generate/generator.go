package generate

import (
	"bytes"
	"sort"

	"github.com/giantswarm/microerror"
	"github.com/santhosh-tekuri/jsonschema/v5"

	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
	"github.com/giantswarm/schemadocs/pkg/generate/templates"
	"github.com/giantswarm/schemadocs/pkg/key"
)

func Generate(schemaPath string) (string, error) {
	compiler := jsonschema.NewCompiler()
	compiler.ExtractAnnotations = true
	schema, err := compiler.Compile(schemaPath)

	if err != nil {
		return "", microerror.Maskf(pkgerror.InvalidSchemaFile, err.Error())
	}

	sections := sectionsFromSchema(schema)
	return toMarkdown(sections)
}

func sectionsFromSchema(schema *jsonschema.Schema) []Section {
	var sections []Section
	var otherSectionRows []Row

	if key.SchemaIsPrimitive(schema) {
		otherSectionRows = append(otherSectionRows, RowsFromSchema(schema, "", schema.Title, []string{})...)
	} else if key.SchemaIsPresentable(schema) {
		sections = append(sections, SectionFromSchema(schema, schema.Title))
	} else {
		for name, property := range schema.Properties {
			if key.SchemaIsPrimitive(property) {
				otherSectionRows = append(otherSectionRows, RowsFromSchema(property, "", name, []string{})...)
			} else {
				sections = append(sections, SectionFromSchema(property, name))
			}
		}
	}

	sort.SliceStable(sections, func(i, j int) bool { return sections[i].Title < sections[j].Title })

	if len(otherSectionRows) > 0 {
		section := NewSection(key.OtherSectionTitle, otherSectionRows)
		sections = append(sections, section)
	}

	return sections
}

func toMarkdown(sections []Section) (string, error) {
	/*t, err := template.ParseFiles(
		path.Join(pkgPath, "templates/docs.md.tpl"),
		path.Join(pkgPath, "templates/row.md.tpl"),
		path.Join(pkgPath, "templates/section.md.tpl"),
	)*/

	t, err := templates.GetTemplate()

	if err != nil {
		return "", microerror.Maskf(pkgerror.DocsGenerationError, err.Error())
	}

	var tplBuffer bytes.Buffer
	err = t.Execute(&tplBuffer, sections)
	if err != nil {
		return "", microerror.Maskf(pkgerror.DocsGenerationError, err.Error())
	}
	return tplBuffer.String(), nil
}
