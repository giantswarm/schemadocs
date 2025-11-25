| `
{{- .FullPath -}}
` |
{{- if or .Title .Description -}}
  {{- if .Title }} **{{ .Title }}**{{ end -}}
  {{- if and .Title .Description }} - {{ end -}}
  {{- .Description -}}
{{- else -}}
  **None**
{{- end -}}
|
{{- if ne (len .Types) 0 -}}
    **Type {{- if gt (len .Types) 1 -}} s {{- end -}} :** `
    {{- range $index, $element := .Types -}}
        {{ if $index }}, {{ end }}{{ $element }}
    {{- end -}}`<br/>
{{- end }}
{{- if ne (len .Examples) 0 -}}
    **Example {{- if gt (len .Examples) 1 -}} s {{- end -}} :** `
    {{- range $index, $element := .Examples -}}
        {{ if $index }}, {{ end }}{{ $element }}
    {{- end -}}`<br/>
{{- end }}
{{- if ne (len .KeyPatterns) 0 -}}
    **Key pattern {{- if gt (len .KeyPatterns) 1 -}} s {{- end -}} :**
    {{- range .KeyPatterns -}}
        <br/>`{{ index $.KeyPatternMappings . }}`=`{{ . }}`
    {{- end -}}<br/>
{{- end -}}
{{- if ne .ConstValue nil -}}
    **Must have value:** `{{ .ConstValue }}`<br/>
{{- end -}}
{{- if eq (len .EnumValues) 1 -}}
    **Allowed value:** `{{ index .EnumValues 0 }}`<br/>
{{- end -}}
{{- if gt (len .EnumValues) 1 -}}
    **Allowed values:** {{ range $index, $element := .EnumValues -}}{{ if $index }}, {{ end }}`{{ $element }}`{{- end -}}<br/>
{{- end -}}
{{- if ne .ValuePattern "" -}}
    **Value pattern:** `{{ .ValuePattern }}`<br/>
{{- end -}}
{{- if ne .DefaultValue "" -}}
    **Default:** `{{ .DefaultValue }}`
{{- end -}}
|
