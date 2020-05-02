package web

import (
	"net/http"
)

func adminMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, admin"))
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, login"))
}
