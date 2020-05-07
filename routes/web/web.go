package web

import (
	"net/http"

	"github.com/root87x/examples/app/controllers/admin"
	"github.com/root87x/examples/app/controllers/site"
	"github.com/root87x/examples/app/middlewares"
	"github.com/root87x/examples/internal/view"
)

type web struct {
	Handler *http.ServeMux
	Routes  map[string]interface{}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}

// Конструктор
func NewWeb(handler *http.ServeMux) *web {
	view := view.NewView()
	site := &site.SiteController{View: view}
	admin := &admin.AdminController{View: view}

	handler.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return &web{
		Handler: handler,
		Routes: map[string]interface{}{
			"/":            http.HandlerFunc(site.Main),
			"/favicon.ico": http.HandlerFunc(favicon),
			"/contacts":    http.HandlerFunc(site.Contacts),
			"/admin":       middlewares.WEBAuthMiddleware(http.HandlerFunc(admin.Dashboard)),
			"/admin/login": middlewares.WEBGuestMiddleware(http.HandlerFunc(admin.Auth)),
		},
	}
}
