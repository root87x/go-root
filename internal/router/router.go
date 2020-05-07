package router

import (
	"net/http"
)

type router struct {
	handler *http.ServeMux
	routes  map[string]interface{}
}

func (r *router) Map() {
	for path, fnc := range r.routes {
		p := string(path)
		f := fnc.(http.Handler)
		r.handler.Handle(p, f)
	}
}

// Установить карту маршрутов
func (r *router) SetRoutes(routes map[string]interface{}) {
	r.routes = routes
}

// Добавить карту маршрутов
func (r *router) AddRoutes(routes map[string]interface{}) {
	for path, fnc := range routes {
		r.routes[path] = fnc.(http.Handler)
	}
}

func NewRouter(handler *http.ServeMux) *router {
	return &router{
		handler: handler,
		routes:  map[string]interface{}{},
	}
}
