package generate

import (
	"sort"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
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
	return Section{
		Name:        name,
		Title:       property.Title,
		Description: property.Description,
		Path:        path,
		Slug:        strings.ReplaceAll(property.Title, " ", "-"),
		Rows:        sortedRows(RowsFromSchema(property, path, name, []string{})),
	}
}

func NewSection(title string, rows []Row) Section {
	return Section{
		Title: title,
		Slug:  strings.ReplaceAll(title, " ", "-"),
		Rows:  sortedRows(rows),
	}
}

func sortedRows(rows []Row) []Row {
	sort.SliceStable(rows, func(i, j int) bool { return rows[i].FullPath < rows[j].FullPath })
	return rows
}
