package site

import (
	"net/http"
)

func (sc *SiteController) PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	sc.View.ParseFiles().ExecuteTemplate(w, "notfound", map[string]interface{}{
		"Title": "Страница не найдена",
	})
}
