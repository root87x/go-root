// Package apiV1
package v1

import (
	"encoding/json"
	"net/http"

	"github.com/root87x/examples/middlewares"
)

type apiV1 struct {
	Handler *http.ServeMux
	Routes  map[string]interface{}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	dataMap := map[string]interface{}{
		"msg": "Hello, World",
	}

	data, _ := json.Marshal(dataMap)

	w.Write(data)
}

// Конструктор
func NewAPIV1(handler *http.ServeMux) *apiV1 {
	return &apiV1{
		Handler: handler,
		Routes: map[string]interface{}{
			"/api/v1": middlewares.APIMiddleware(http.HandlerFunc(mainPage)),
		},
	}
}
