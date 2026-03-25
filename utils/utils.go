package utils

import (
	"bytes"
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, path string) error {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, Artistdata); err != nil {
		return err
	}

	buf.WriteTo(w)
	return nil
}
