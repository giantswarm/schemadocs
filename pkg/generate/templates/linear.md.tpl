
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
    </a>{{.Title}}
  </h3>
  {{- range .Rows }}
  <div class="property depth-0">
    <div class="property-header">
      <h4 class="property-path headline-with-link">
        <a class="header-link" href="#{{.Slug}}">
          <i class="fa fa-link"></i>
        </a>.{{- .FullPath -}}
      </h4>
    </div>
    <div class="property-body">
      <div class="property-meta">
        {{- if ne (len .Types) 0 -}}
          {{- range $index, $element := .Types -}}
            {{ if $index }}, {{ end }}<span class="property-type"><em>{{- $element }}</em></span>
          {{- end -}}&nbsp;<br />
        {{- end }}
        {{- if .Title -}}
        <span class="property-title">{{- .Title -}}.&nbsp;</span>
        {{- end -}}
        {{- if .Description -}}
        <span class="property-description">{{- .Description -}}</span><br />
        {{- end -}}
      </span>
      </div>
    </div>
  </div>
  {{- end }}
{{- end -}}
</div>
