package handler

import (
	"html/template"
	"log"
	"net/http"

	tpl "github.com/root87x/examples/internal/template"
)

type Handler struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *Handler) Template(pathTpl []string) *template.Template {
	templ, err := tpl.Parse(pathTpl)
	if err != nil {
		log.Fatalln(err)
	}

	return templ
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
