package handler

import (
	"net/http"
)

type Handler struct {
	w http.ResponseWriter
	r *http.Request
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
