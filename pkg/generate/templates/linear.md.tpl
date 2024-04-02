
## Chart configuration reference

{{ range . -}}
### {{ .Title }} {#{{ .Slug }}}

{{- if .Description -}}
  {{- .Description -}}
{{- end }}

{{ range .Rows -}}
---

`.{{ .FullPath }}`

{{ if eq (len .Types) 1 -}}
  **Type:** `{{ index .Types 0 }}`
{{- else if gt (len .Types) 1 -}}
  **Types:**
  {{- range $index, $element := .Types -}}
    {{ if $index }}, {{ end }}`{{- $element }}`
  {{- end }}
{{- end }}

{{ if .Title -}}
**{{ .Title }}**
{{- end }}

{{ if .Description -}}
  {{- .Description }}
{{- end }}

{{ end }}
{{ end }}
