package site

import "net/http"

func (sc *SiteController) Contacts(w http.ResponseWriter, r *http.Request) {
	sc.View.ParseFiles().ExecuteTemplate(w, "contacts", map[string]string{
		"title": "Contacts",
	})
}
