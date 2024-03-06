
<div class="crd-schema-version">
  <h2 class="headline-with-link">
    <a class="header-link" href="#">
      <i class="fa fa-link"></i>
    </a>Chart Configuration Reference
  </h2>

{{- range . }}
  <h3 class="headline-with-link">
    <a class="header-link" href="#{{.Slug}}">
      <i class="fa fa-link"></i>
    </a>{{ .Title | html }}
  </h3>
  {{- if .Description }}
  <h4 class="headline-with-link">
    <a class="header-link" href="#">
      <i class="fa fa-link"></i>
    </a>{{ .Description | html }}
  </h4>
  {{- end }}
  {{- range .Rows }}
  <div class="property depth-0">
    <div class="property-header">
      <h3 class="property-path headline-with-link">
        <a class="header-link" href="#{{.Slug}}">
          <i class="fa fa-link"></i>
        </a>.{{- .FullPath -}}
      </h3>
    </div>
    <div class="property-body">
      <div class="property-meta">
        {{- if .Title -}}
        <span class="property-title">{{- .Title | html  -}}</span><br />
        {{- end -}}
        {{- if ne (len .Types) 0 -}}
          {{- range $index, $element := .Types -}}
            {{ if $index }}, {{ end }}<span class="property-type">{{- $element }}</span>
          {{- end -}}&nbsp;
        {{- end }}
      </div>
      <div class="property-description">
        {{- if .Description -}}
          {{- .Description | html -}}
        {{- end -}}
      </div>
    </div>
  </div>
  {{- end }}
{{- end -}}
</div>
