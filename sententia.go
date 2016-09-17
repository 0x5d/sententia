package sententia

import (
	"bytes"
	"text/template"
)

// Make generates a sentence replacing the noun and adjective templates.
func Make(sentence string) (string, error) {
	tmpl, err := template.New("sentence").Parse(sentence)
	if err != nil {
		return "", err
	}
	var buf = &bytes.Buffer{}
	err = tmpl.Execute(buf, &maker{})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
