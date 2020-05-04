package handler

import (
	"net/http"

	"github.com/root87x/examples/internal/template"
)

func WEBError(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	tmpl := template.ParseWithBlocks([]string{"./views/layout.html", "./views/pages/notfound.html"})
	tmpl.ExecuteTemplate(w, "notfound", map[string]string{
		"title": "Page Not found",
	})
}
