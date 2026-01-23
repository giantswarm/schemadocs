package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/santhosh-tekuri/jsonschema/v6"

	pkgerror "github.com/giantswarm/schemadocs/pkg/error"
	"github.com/giantswarm/schemadocs/pkg/generate/templates"
	"github.com/giantswarm/schemadocs/pkg/key"
)

// UniversalLoader implements jsonschema.URLLoader for both HTTP/HTTPS and file URLs
type UniversalLoader struct{}

func (u UniversalLoader) Load(urlStr string) (any, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %s", urlStr)
	}

	var data []byte

	switch parsedURL.Scheme {
	case "http", "https":
		// Handle HTTP/HTTPS URLs
		// This is intentionally loading from a variable URL as part of JSON schema resolution
		resp, err := http.Get(urlStr) // #nosec G107
		if err != nil {
			return nil, err
		}
		defer func() {
			if closeErr := resp.Body.Close(); closeErr != nil {
				// Log the error but don't override the main error
				// In a real application, you might want to log this properly
				_ = closeErr // Explicitly ignore the error to satisfy linters
			}
		}()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

	case "file":
		// Handle file:// URLs
		filePath := parsedURL.Path
		// This is intentionally reading from a variable path as part of JSON schema resolution
		data, err = os.ReadFile(filePath) // #nosec G304
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unsupported protocol scheme: %s", parsedURL.Scheme)
	}

	// Parse JSON and return the parsed data
	var jsonData any
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("invalid JSON in %s: %w", urlStr, err)
	}

	return jsonData, nil
}

func Generate(schemaPath string, layout string) (string, error) {
	compiler := jsonschema.NewCompiler()

	// Add universal loader for both HTTP/HTTPS and file URLs
	compiler.UseLoader(UniversalLoader{})

	schema, err := compiler.Compile(schemaPath)

	if err != nil {
		return "", fmt.Errorf("%s: %w", err.Error(), pkgerror.ErrInvalidSchemaFile)
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
		if sections[i].Title != sections[j].Title {
			return sections[i].Title < sections[j].Title
		}
		if sections[i].Path != sections[j].Path {
			return sections[i].Path < sections[j].Path
		}
		return sections[i].Name < sections[j].Name
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

var multipleLinebreaksRegex = regexp.MustCompile(`\n{2,}`)

func toMarkdown(sections []Section, layout string) (string, error) {
	var t *template.Template
	var err error
	if layout == "linear" {
		t, err = templates.GetLinearTemplate()
	} else {
		t, err = templates.GetDefaultTemplate()
	}
	if err != nil {
		return "", fmt.Errorf("%s: %w", err.Error(), pkgerror.ErrDocsGeneration)
	}

	var tplBuffer bytes.Buffer
	err = t.Execute(&tplBuffer, sections)
	if err != nil {
		return "", fmt.Errorf("%s: %w", err.Error(), pkgerror.ErrDocsGeneration)
	}
	ret := tplBuffer.String()

	// Replace any occurrences of more than two linebreaks with exactly two linebreaks
	// (high effort to do this consistently in a Go template)
	ret = multipleLinebreaksRegex.ReplaceAllString(ret, "\n\n")

	// Only keep exactly one newline at the end since editors are normally configured to do that
	ret = strings.TrimRight(ret, "\n") + "\n"

	return ret, nil
}
