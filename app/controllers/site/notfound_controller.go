package site

import (
	"net/http"

	"github.com/root87x/examples/internal/handler"
)

func (sc *SiteController) PageNotFound(w http.ResponseWriter, r *http.Request) {
	tmpl := handler.NewHandler(w, r).Status(404).Template()
	tmpl.ExecuteTemplate(w, "notfound", map[string]interface{}{
		"Title": "Страница не найдена",
	})
}
