package main

import (
	"fmt"
	"net/http"

	"github.com/cameronbrill/brill-wtf-go/rest"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	svc := service.New()
	rest.RegisterLinkServiceRouter(svc, r)
	port := 3333
	fmt.Printf("listening on port :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}
