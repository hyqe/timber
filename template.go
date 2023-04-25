package timber

import "html/template"

const (
	DATA      = "{{.Data}}"
	STATUS    = "{{.Level}} {{.Data}}"
	TIMESTAMP = "{{.CreatedAt}} {{.Level}} {{.Data}}"
)

var (
	tmplData      = mustParseTemplate(DATA)
	tmplSTATUS    = mustParseTemplate(STATUS)
	tmplTIMESTAMP = mustParseTemplate(TIMESTAMP)
)

func mustParseTemplate(tmpl string) *template.Template {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	return t
}
