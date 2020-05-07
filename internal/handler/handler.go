package handler

import (
	"html/template"
	"net/http"

	"github.com/root87x/examples/internal/view"
)

type Handler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *Handler) Template() *template.Template {
	tmpl := view.NewView().ParseFiles()
	return tmpl
}

func (h *Handler) Message(msg string) {
	h.w.Write([]byte(msg))
}

func (h *Handler) Status(code int) *Handler {
	h.w.WriteHeader(code)

	return h
}

func NewHandler(w http.ResponseWriter, r *http.Request) *Handler {
	return &Handler{
		w: w,
		r: r,
	}
}
