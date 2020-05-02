package server

import (
	"fmt"
	"net/http"
	"time"
)

type server struct {
	srv *http.Server
}

func (s server) SetReadTimeout(value time.Duration) {
	s.srv.WriteTimeout = value
}

func (s server) SetWriteTimeout(value time.Duration) {
	s.srv.ReadTimeout = value
}

func (s server) StartWithTls(certFile string, keyFile string) error {
	fmt.Println("Start server with tls ON " + s.srv.Addr)
	return s.srv.ListenAndServeTLS(certFile, keyFile)
}

func (s server) Start() error {
	fmt.Println("Start server ON " + s.srv.Addr)
	return s.srv.ListenAndServe()
}

func NewServer(ip string, handler *http.ServeMux) *server {
	return &server{
		srv: &http.Server{
			Addr:    ip,
			Handler: handler,
		},
	}
}
