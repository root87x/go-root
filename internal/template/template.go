package template

import (
	"html/template"
	"log"
)

var Template *template.Template

func Parse(path []string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(path...)
	if err != nil {
		log.Fatalln(err)
	}

	return tmpl, err
}
