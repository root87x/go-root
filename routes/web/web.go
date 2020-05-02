package web

import (
	"net/http"

	"github.com/root87x/examples/middlewares"

	"github.com/mattn/go-zglob"

	"github.com/root87x/examples/internal/template"
)

type web struct {
	Handler *http.ServeMux
	Routes  map[string]interface{}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	tmpl := template.Parse([]string{"./views/layout.html", "./views/notfound.html"})
	tmpl.ExecuteTemplate(w, "notfound", map[string]string{
		"title": "Page Not found",
	})
}

func main(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r)
		return
	}

	path := []string{"./views/layout.html", "./views/pages/main.html"}
	blocks, _ := zglob.Glob("./views/blocks/*.html")
	path = append(path, blocks...)

	tmpl := template.Parse(path)

	tmpl.ExecuteTemplate(w, "main", map[string]string{
		"title": "Main page",
	})
}

func contacts(w http.ResponseWriter, r *http.Request) {
	path := []string{"./views/layout.html", "./views/pages/contacts.html"}
	blocks, _ := zglob.Glob("./views/blocks/*.html")
	path = append(path, blocks...)

	tmpl := template.Parse(path)
	tmpl.ExecuteTemplate(w, "contacts", map[string]string{
		"title": "Contacts",
	})
}

func NewWeb(handler *http.ServeMux) *web {
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
