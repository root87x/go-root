package web

import (
	"net/http"

	"github.com/mattn/go-zglob"
	"github.com/root87x/examples/middlewares"

	"github.com/root87x/examples/internal/handler"

	"github.com/root87x/examples/internal/template"
)

type web struct {
	Handler *http.ServeMux
	Routes  map[string]interface{}
}

var templates []string

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}

// Страница не найдена
func PageNotFound(w http.ResponseWriter, r *http.Request) {
	templates = append(templates, []string{"./views/site/layout.html", "./views/site/pages/notfound.html"}...)
	err := handler.NewHandler(w, r).Status(404).Template(templates)
	err.ExecuteTemplate(w, "notfound", map[string]interface{}{
		"Title": "Страница не найдена",
	})
}

// Главная страница
func main(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		PageNotFound(w, r)
		return
	}
	templates = append(templates, []string{"./views/site/layout.html", "./views/site/pages/main.html"}...)
	tmpl, _ := template.Parse(templates)
	tmpl.ExecuteTemplate(w, "main", map[string]string{
		"title": "Main page",
	})
}

// Контакты
func contacts(w http.ResponseWriter, r *http.Request) {
	templates = append(templates, []string{"./views/site/layout.html", "./views/site/pages/contacts.html"}...)
	tmpl, _ := template.Parse(templates)
	tmpl.ExecuteTemplate(w, "contacts", map[string]string{
		"title": "Contacts",
	})
}

// Конструктор
func NewWeb(handler *http.ServeMux) *web {
	matches, _ := zglob.Glob("./views/site/blocks/*.html")
	templates = matches

	handler.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return &web{
		Handler: handler,
		Routes: map[string]interface{}{
			"/":            http.HandlerFunc(main),
			"/favicon.ico": http.HandlerFunc(favicon),
			"/contacts":    http.HandlerFunc(contacts),
			"/admin":       middlewares.WEBAuthMiddleware(http.HandlerFunc(adminMain)),
			"/admin/login": middlewares.WEBGuestMiddleware(http.HandlerFunc(adminLogin)),
		},
	}
}
