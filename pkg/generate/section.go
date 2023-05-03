package generate

import (
	"sort"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Section struct {
	Name        string
	Title       string
	Description string
	Rows        []Row
}

func SectionFromSchema(property *jsonschema.Schema, name string) Section {
	return Section{
		Name:        name,
		Title:       property.Title,
		Description: property.Description,
		Rows:        sortedRows(RowsFromSchema(property, "", name, []string{})),
	}
}

func NewSection(title string, rows []Row) Section {
	return Section{
		Title: title,
		Rows:  sortedRows(rows),
	}
}

func sortedRows(rows []Row) []Row {
	sort.SliceStable(rows, func(i, j int) bool { return rows[i].FullPath < rows[j].FullPath })
	return rows
}
