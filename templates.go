package kingmaker

import (
	"bytes"
	"html/template"
)

func fillTemplate(w *World) (string, error) {
	tmpl, err := template.New("event").Parse(w.Event.Template)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, w)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
