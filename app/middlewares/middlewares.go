package middlewares

import (
	"net/http"
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
func WEBAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("auth")
		// Если пусто на душе, то отправляем на исповедь
		if err != nil {
			http.Redirect(w, r, "/admin/login", 302)
		} else {
			// Здесь делаем доп.проверку по авторизации
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Permission denied"))
		}
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
