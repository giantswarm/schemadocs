### {{.Title}}
{{- if .Name}}
Properties within the `.{{.Name}}` top-level object
{{- end }}
{{- if .Description }}
{{ .Description }}
{{- end }}

| **Property** | **Description** | **More Details** |
| :----------- | :-------------- | :--------------- |
{{ range .Rows }}
{{- template "row.md.tpl" . }}
{{- end }}
