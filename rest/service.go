package rest

import (
	"os"

	"github.com/cameronbrill/brill-wtf-go/rest/controller"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/cameronbrill/brill-wtf-go/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type linkRouter struct {
	c        *controller.LinkServiceController
	renderer web.Renderer
}

func RegisterLinkServiceRouter(svc service.Service, r *chi.Mux, opts ...Option) {
	var router linkRouter

	for _, opt := range opts {
		err := opt(&router)
		if err != nil {
			panic(err)
		}
	}

	ctrl := controller.New(svc, router.renderer)
	router.c = &ctrl

	ogns := []string{"https://brill.wtf", "http://brill.wtf", "https://www.brill.wtf", "http://www.brill.wtf"}
	if os.Getenv("ENV") == "dev" {
		ogns = []string{"*"}
	}
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   ogns,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/link", func(subRouter chi.Router) {
		subRouter.Use(linkCtx)
		subRouter.Get("/", router.c.ShortURLToLink)
		subRouter.Post("/", router.c.NewLink)
	})
}

type Option func(*linkRouter) error

func WithRenderer(renderer web.Renderer) Option {
	return func(r *linkRouter) error {
		r.renderer = renderer
		return nil
	}
}
