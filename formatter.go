package main

import (
	"strings"
	"text/template"
)

const sourceDocTemplate = `# {{.Type}}

{{.Description}}

## Supported Methods
{{ if ne .Get "" }}
* **Get:** {{.Get}}
{{- else }}
* **Get**
{{- end }}
{{- if ne .List "" }}
* **List:** {{.List}}
{{- else }}
* **List**
{{- end }}
{{- if ne .Search "" }}
* **Search:** {{.Search}}
{{- end }}
`

func (s *SourceDoc) FormatMarkdown() (string, error) {
	t := template.Must(template.New("sourceDoc").Parse(sourceDocTemplate))
	out := strings.Builder{}

	err := t.Execute(&out, s)

	if err != nil {
		return "", err
	}

	return out.String(), nil
}
