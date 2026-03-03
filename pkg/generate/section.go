package generate

import (
	"sort"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

type Section struct {
	Name        string
	Title       string
	Slug        string
	Description string
	Path        string
	Rows        []Row
}

func SectionFromSchema(property *jsonschema.Schema, path, name string) Section {
	title := property.Title
	if title == "" {
		title = name
	}
	return Section{
		Name:        name,
		Title:       title,
		Description: property.Description,
		Path:        path,
		Slug:        strings.ToLower(strings.ReplaceAll(title, " ", "-")),
		Rows:        sortedRows(RowsFromSchema(property, path, name, []string{})),
	}
}

func NewSection(title string, rows []Row) Section {
	return Section{
		Title: title,
		Slug:  strings.ToLower(strings.ReplaceAll(title, " ", "-")),
		Rows:  sortedRows(rows),
	}
}

func sortedRows(rows []Row) []Row {
	sort.SliceStable(rows, func(i, j int) bool { return rows[i].FullPath < rows[j].FullPath })
	return rows
}
