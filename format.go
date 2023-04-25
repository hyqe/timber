package timber

import (
	"bytes"
	"text/template"
)

// Formatter is responsible for generating the final text output.
type Formatter func(l Log) string

var (
	STATUS    = template.Must(template.New("STATUS").Parse("{{.Level}}: {{.Message}}"))
	TIMESTAMP = template.Must(template.New("TIMESTAMP").Parse("{{.CreatedAt}} {{.Level}}: {{.Message}}"))
)

// TEMPLATE formatter for logs
func TEMPLATE(t *template.Template) Formatter {
	return func(l Log) string {
		var buff bytes.Buffer
		t.Execute(&buff, l.Flat())
		return buff.String()
	}
}

// JSON Formatter for logs
func JSON(l Log) string {
	return string(l.JSON())
}
