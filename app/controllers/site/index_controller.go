package site

import "net/http"

func (sc *SiteController) Main(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		sc.PageNotFound(w, r)
		return
	}

	sc.View.ParseFiles().ExecuteTemplate(w, "main", map[string]string{
		"title": "Main page",
	})
}
