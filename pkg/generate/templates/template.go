package templates

import (
	_ "embed"
	"text/template"
)

//go:embed docs.md.tpl
var docs string

//go:embed linear.md.tpl
var linear string

//go:embed row.md.tpl
var row string

//go:embed section.md.tpl
var section string

func GetDefaultTemplate() (*template.Template, error) {
	docsTpl, err := template.New("docs.md.tpl").Parse(docs)
	if err != nil {
		return nil, err
	}

	sectionTpl, err := docsTpl.New("section.md.tpl").Parse(section)
	if err != nil {
		return nil, err
	}

	_, err = sectionTpl.New("row.md.tpl").Parse(row)
	if err != nil {
		return nil, err
	}

	return docsTpl, nil
}

func GetLinearTemplate() (*template.Template, error) {
	docsTpl, err := template.New("linear.md.tpl").Parse(linear)
	if err != nil {
		return nil, err
	}

	return docsTpl, nil
}
