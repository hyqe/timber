package timber

import (
	"bytes"
	"io"
	"text/template"
)

var (
	DATA      = template.Must(template.New("DATA").Parse("{{.Data}}\n"))
	STATUS    = template.Must(template.New("STATUS").Parse("{{.Level}}: {{.Data}}\n"))
	TIMESTAMP = template.Must(template.New("TIMESTAMP").Parse("{{.CreatedAt}} {{.Level}}: {{.Data}}\n"))
)

// Formatter is responsible for generating the final text output.
type Formatter = func(l Log) io.Reader

func defaultFormatter(tmpl *template.Template) Formatter {
	return func(l Log) io.Reader {
		var buff bytes.Buffer
		tmpl.Execute(&buff, l)
		return &buff
	}
}
