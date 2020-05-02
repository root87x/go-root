package template

import (
	"html/template"
	"log"

	"github.com/mattn/go-zglob"
)

func ParseGlob(path string) *template.Template {
	matches, err := zglob.Glob(path)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err := template.ParseFiles(matches...)
	if err != nil {
		log.Fatalln(err)
	}

	return tmpl
}

func Parse(path []string) *template.Template {
	tmpl, err := template.ParseFiles(path...)
	if err != nil {
		log.Fatalln(err)
	}

	return tmpl
}
