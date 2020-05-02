package middlewares

import (
	"fmt"
	"net/http"
)

// API middleware
func APIMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Проверка авторизации для web api
func WEBAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//authToken := "a1239876SHA-123"
		cookie, err := r.Cookie("auth")
		if err != nil {
			http.Redirect(w, r, "/admin/login", 302)
			next.ServeHTTP(w, r)
		}

		value := cookie.Value
		fmt.Println(value)

		next.ServeHTTP(w, r)
	})
}

/**
Авторизация для web api
Только для гостей
*/
func WEBGuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cook, _ := r.Cookie("auth")
		if cook != nil {
			if r.URL.Path != "/admin/login" {
				http.Redirect(w, r, r.Referer(), 302)
				next.ServeHTTP(w, r)
			}
		}
		next.ServeHTTP(w, r)
	})
}
