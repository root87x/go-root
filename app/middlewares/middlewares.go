package middlewares

import (
	"net/http"
	"time"
)

/**
API middleware
*/
func APIMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

/**
Проверка авторизации для web api
*/
func WEBAuthMiddleware(next http.Handler, redirectTo string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var auth bool = false
		_, err := r.Cookie("auth")
		if err != nil {
			http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
		}
		if !auth {
			http.SetCookie(w, &http.Cookie{Name: "auth", Value: "", Path: "/", Expires: time.Unix(0, 0), HttpOnly: true})
			http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

/**
Авторизация для web api
Только для гостей
*/
func WEBGuestMiddleware(next http.Handler, redirectTo string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("auth")
		if err == nil {
			if r.URL.Path != redirectTo {
				http.Redirect(w, r, redirectTo, 302)
				next.ServeHTTP(w, r)
			}
		}
		next.ServeHTTP(w, r)
	})
}
