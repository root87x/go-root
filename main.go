package main

import (
	"log"
	"net/http"
	"time"

	"github.com/root87x/examples/router"
	apiv1 "github.com/root87x/examples/routes/api/v1"
	web "github.com/root87x/examples/routes/web"
	"github.com/root87x/examples/server"
)

func main() {
	handler := http.NewServeMux()

	srv := server.NewServer(":8000", handler)
	srv.SetReadTimeout(time.Second * 5000)
	srv.SetWriteTimeout(time.Second * 5000)

	api := apiv1.NewAPIV1(handler)
	web := web.NewWeb(handler)

	router := router.NewRouter(api.Handler)
	router.SetRoutes(api.Routes)
	router.AddRoutes(web.Routes)
	router.Map()

	err := srv.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
