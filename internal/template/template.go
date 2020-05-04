package template

import (
	"html/template"
	"log"

	"github.com/mattn/go-zglob"
)

func Parse(path []string) *template.Template {
	tmpl, err := template.ParseFiles(path...)
	if err != nil {
		log.Fatalln(err)
	}

	return tmpl
}

func ParseWithBlocks(path []string) *template.Template {
	matches, _ := zglob.Glob("./views/blocks/*.html")
	matches = append(matches, path...)
	tmpl := Parse(matches)

	return tmpl
}
