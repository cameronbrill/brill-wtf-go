package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cameronbrill/brill-wtf-go/rest"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/cameronbrill/brill-wtf-go/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	var opt service.ServiceOption
	if os.Getenv("ENV") == "dev" {
		println("starting dev server...")
		opt = service.WithBasicStorage()
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"http://localhost:3000"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}))
	} else {
		println("starting server...")
		opt = service.WithRedisStorage()
	}
	svc := service.New(opt)
	rest.RegisterLinkServiceRouter(svc, r, rest.WithRenderer(web.New()))
	port := 3333
	fmt.Printf("listening on port :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}
