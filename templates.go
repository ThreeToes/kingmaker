package kingmaker

import (
	"bytes"
	"html/template"
)

func FillTemplate(w *EventContext) (string, error) {
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
