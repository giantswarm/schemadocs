
## Chart configuration reference

{{ range . -}}
### {{ .Title }} {#{{ .Slug }}}

{{- if .Description -}}
  {{- .Description -}}
{{- end }}

{{ range .Rows -}}
{{- $row := . -}}
---

`.{{ .FullPath }}`

{{ if eq (len .Types) 1 -}}
  **Type:** `{{ index .Types 0 }}`
{{- else if gt (len .Types) 1 -}}
  **Types:**
  {{- range $index, $element := .Types -}}
    {{ if $index }}, {{ else }} {{ end }}`{{- $element }}`
  {{- end }}
{{- end }}

{{ if .Title -}}
**{{ .Title }}**
{{- end }}

{{ if .Description -}}
  {{- .Description }}
{{- end }}

{{ if ne (len .Examples) 0 -}}
    **Example {{- if gt (len .Examples) 1 -}} s {{- end -}} :**
    {{- range $index, $element := .Examples -}}
        {{ if $index }}, {{ else }} {{ end }}`{{ $element }}`
    {{- end -}}
{{- end }}

{{ if ne (len .KeyPatterns) 0 -}}
    **Key pattern {{- if gt (len .KeyPatterns) 1 -}} s {{- end -}} :**
    {{- range $index, $element := .KeyPatterns -}}
        {{- if $index }}, {{ else }} {{ end }}`{{ index $row.KeyPatternMappings $element }}`=`{{ $element }}`
    {{- end -}}
{{- end }}

{{ if ne .ConstValue nil -}}
    **Must have value:** `{{ .ConstValue }}`
{{- end }}

{{ if eq (len .EnumValues) 1 -}}
    **Allowed value:** `{{ index .EnumValues 0 }}`
{{- end }}
{{ if gt (len .EnumValues) 1 -}}
    **Allowed values:** {{ range $index, $element := .EnumValues -}}{{ if $index }}, {{ end }}`{{ $element }}`{{- end -}}
{{- end }}

{{ if ne .ValuePattern "" -}}
    **Value pattern:** `{{ .ValuePattern }}`
{{- end }}

{{ if ne .DefaultValue "" -}}
    **Default:** `{{ .DefaultValue }}`
{{- end }}

{{ end }}
{{ end }}
