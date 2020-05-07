package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/root87x/examples/internal/router"
	apiv1 "github.com/root87x/examples/routes/api/v1"
	web "github.com/root87x/examples/routes/web"
	"github.com/root87x/examples/server"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalln("File .env not found: " + err.Error())
	}
}

func main() {
	handler := http.NewServeMux()

	srv := server.NewServer(handler)
	srv.SetAddrHost(os.Getenv("APP_ADDR_HOST"))
	srv.SetReadTimeout(time.Second * 5000)
	srv.SetWriteTimeout(time.Second * 5000)

	api := apiv1.NewAPIV1(handler)
	web := web.NewWeb(handler)

	router := router.NewRouter(handler)
	router.SetRoutes(api.Routes)
	router.AddRoutes(web.Routes)
	router.Map()

	err := srv.Listen()
	if err != nil {
		log.Fatalln(err)
	}
}
