package server

import (
	"fmt"
	"log"
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

func (s server) SetAddrHost(value string) {
	if len(value) < 1 {
		return
	}
	s.srv.Addr = value
}

func (s server) ListenWithTls(certFile string, keyFile string) error {
	fmt.Println("Start server with tls ON " + s.srv.Addr)
	return s.srv.ListenAndServeTLS(certFile, keyFile)
}

func (s server) Listen() error {
	log.Println("Start server ON " + s.srv.Addr)
	return s.srv.ListenAndServe()
}

func NewServer(handler *http.ServeMux) *server {
	return &server{
		srv: &http.Server{
			Handler: handler,
		},
	}
}
