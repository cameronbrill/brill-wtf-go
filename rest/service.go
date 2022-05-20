package rest

import (
	"github.com/cameronbrill/brill-wtf-go/rest/controller"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/cameronbrill/brill-wtf-go/web"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
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

	r.Use(cors.Default().Handler)

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
