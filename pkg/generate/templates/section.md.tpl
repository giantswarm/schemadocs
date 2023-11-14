### {{.Title}}
{{- if .Name}}
{{- if eq .Path "global" }}
Properties within the `.global.{{.Name}}` object
{{- else }}
Properties within the `.{{.Name}}` top-level object
{{- end }}
{{- end }}
{{- if .Description }}
{{ .Description }}
{{- end }}

| **Property** | **Description** | **More Details** |
| :----------- | :-------------- | :--------------- |
{{ range .Rows }}
{{- template "row.md.tpl" . }}
{{- end }}
