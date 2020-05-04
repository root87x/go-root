package web

import (
	"net/http"

	"github.com/root87x/examples/middlewares"

	"github.com/root87x/examples/internal/handler"

	"github.com/root87x/examples/internal/template"
)

type web struct {
	Handler *http.ServeMux
	Routes  map[string]interface{}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}

func main(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handler.WEBError(w, r, 404)
		return
	}
	tmpl := template.ParseWithBlocks([]string{"./views/layout.html", "./views/pages/main.html"})
	tmpl.ExecuteTemplate(w, "main", map[string]string{
		"title": "Main page",
	})
}

func contacts(w http.ResponseWriter, r *http.Request) {
	tmpl := template.ParseWithBlocks([]string{"./views/layout.html", "./views/pages/contacts.html"})
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
