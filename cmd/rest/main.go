package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cameronbrill/brill-wtf-go/rest"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/cameronbrill/brill-wtf-go/web"
	"gopkg.in/yaml.v2"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/docgen/raml"
)

//  @title          brill.wtf link shortener
//  @version        1.0.0
//  @description    This is a simple RESTful API for shortening and receiving links.

//  @contact.name   Cameron Brill
//  @contact.url    https://cameronbrill.me

//  @license.name   MIT License
//  @license.url    https://github.com/cameronbrill/brill-wtf-go/blob/master/LICENSE

// @host       localhost:3000
// @basePath   /
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
	} else if os.Getenv("ENV") == "gen" {
		println("starting gen server...")
		opt = service.WithBasicStorage()
	} else {
		println("starting server...")
		opt = service.WithRedisStorage()
	}
	svc := service.New(opt)
	rest.RegisterLinkServiceRouter(svc, r, rest.WithRenderer(web.New()))
	if os.Getenv("ENV") == "gen" {
		if err := os.Remove("routes.raml"); err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}

		f, err := os.Create("routes.raml")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		ramlDocs := &raml.RAML{
			Title:     "RAML Representation of RESTful API",
			BaseUri:   "http://api.go-chi-docgen-example.com/v1",
			Version:   "v1.0",
			MediaType: "application/json",
		}

		if err := chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			handlerInfo := docgen.GetFuncInfo(handler)
			resource := &raml.Resource{
				DisplayName: strings.ToUpper(method) + " " + route,
				Description: "Handler Function: " + handlerInfo.Func + "\nComment: " + handlerInfo.Comment,
			}

			return ramlDocs.Add(method, route, resource)
		}); err != nil {
			log.Fatalf("error: %v", err)
		}

		raml, err := yaml.Marshal(ramlDocs)

		if err != nil {
			log.Fatal(err)
		}

		if _, err = f.Write(append([]byte("#%RAML 1.0\n---\n"), raml...)); err != nil { // For the RAML document to be valid, the first line of the file must begin with the text "#%RAML 1.0" followed by an newline character.
			log.Fatal(err)
		}

		return
	}
	port := 3333
	fmt.Printf("listening on port :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}
