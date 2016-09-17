package sententia

import (
	"bytes"
	"text/template"
)

// Make generates a sentence replacing the noun and adjective templates.
func Make(sentence string) (string, error) {
	tmpl, err := template.New("sentence").Funcs(funcs).Parse(sentence)
	if err != nil {
		return "", err
	}
	var buf = &bytes.Buffer{}
	err = tmpl.Execute(buf, nil)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
