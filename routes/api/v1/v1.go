// Package apiV1
package v1

import (
	"net/http"

	v1 "github.com/root87x/examples/app/controllers/api/v1"
	"github.com/root87x/examples/app/middlewares"
)

type apiV1 struct {
	Handler *http.ServeMux
	Routes  map[string]interface{}
}

// Конструктор
func NewAPIV1(handler *http.ServeMux) *apiV1 {
	return &apiV1{
		Handler: handler,
		Routes: map[string]interface{}{
			"/api/v1": middlewares.APIMiddleware(http.HandlerFunc(v1.Test)),
		},
	}
}
