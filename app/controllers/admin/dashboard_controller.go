package admin

import "net/http"

func (ac *AdminController) Dashboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, admin"))
}
