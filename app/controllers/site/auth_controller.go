package site

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var redirectUrl string = "/"

/**
Авторизация с формы
Method POST
Data:
	username string
	password string
*/
func authForm(w http.ResponseWriter, r *http.Request) {
	var auth bool = true
	//var username string = r.Form.Get("username")
	//var password string = r.Form.Get("password")

	//todo проверка авторизации

	if !auth {
		return
	}

	data, err := json.Marshal(map[string]interface{}{
		"id":       1,
		"userhash": "UserHash",
	})

	_, err = r.Cookie("auth")
	if err != nil {
		session := &http.Cookie{
			Name:     "auth",
			Domain:   os.Getenv("APP_DOMAIN"),
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 6),
			Value:    string(data),
			HttpOnly: true,
		}
		http.SetCookie(w, session)
		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
	}
}

//Страница авторизации
func (sc *SiteController) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Referer() != "/auth" {
		redirectUrl = r.Referer()
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
		}
		authForm(w, r)
	}

	sc.View.ParseFiles().ExecuteTemplate(w, "login", map[string]interface{}{
		"Title":      "Авторизация на сайте",
		"Username":   r.Form.Get("username"),
		"FormName":   "form-auth",
		"FormAction": "/auth",
		"FormMethod": "POST",
	})
}
