package view

import (
	"html/template"
	"log"

	"github.com/mattn/go-zglob"
)

type View struct {
	viewCollection []string
	viewPath       string
}

// Парсинг шаблонов
func (v *View) ParseFiles() *template.Template {
	tmpl, err := template.ParseFiles(v.viewCollection...)
	if err != nil {
		log.Fatalln(err)
	}
	return tmpl
}

// Конструктор
func NewView(viewPath string) *View {
	paths, _ := zglob.Glob(viewPath + "/**/*.html")
	return &View{
		viewCollection: paths,
		viewPath:       viewPath,
	}
}
