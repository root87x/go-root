package admin

import (
	"log"
	"net/http"
)

func (ac *AdminController) Auth(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
	}

	ac.View.ParseFiles().ExecuteTemplate(w, "login_admin", map[string]interface{}{
		"Title":      "Авторизация",
		"Username":   r.Form.Get("username"),
		"FormName":   "form-auth",
		"FormAction": "/admin/auth",
		"FormMethod": "POST",
	})
}
