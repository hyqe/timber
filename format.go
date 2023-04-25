package timber

import (
	"bytes"
	"io"
	"text/template"
)

// Formatter is responsible for generating the final text output.
type Formatter func(l Log) io.Reader

var (
	STATUS    = template.Must(template.New("STATUS").Parse("{{.Level}}: {{.Data}}\n"))
	TIMESTAMP = template.Must(template.New("TIMESTAMP").Parse("{{.CreatedAt}} {{.Level}}: {{.Data}}\n"))
)

// TEMPLATE formatter for logs
func TEMPLATE(tmpl *template.Template) Formatter {
	return func(l Log) io.Reader {
		var buff bytes.Buffer
		tmpl.Execute(&buff, l.Flat())
		return &buff
	}
}

// JSON Formatter for logs
func JSON(l Log) io.Reader {
	return bytes.NewReader(l.JSON())
}
