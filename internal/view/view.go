package view

import (
	"html/template"
	"log"

	"github.com/mattn/go-zglob"
)

type View struct {
	pathes []string
}

// Парсинг шаблонов
func (v *View) ParseFiles() *template.Template {
	tmpl, err := template.ParseFiles(v.pathes...)
	if err != nil {
		log.Fatalln(err)
	}
	return tmpl
}

// Конструктор
func NewView() *View {
	paths, _ := zglob.Glob("./views/**/*.html")
	return &View{
		pathes: paths,
	}
}
