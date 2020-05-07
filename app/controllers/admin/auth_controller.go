package admin

import "net/http"

func (ac *AdminController) Auth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, login"))
}
