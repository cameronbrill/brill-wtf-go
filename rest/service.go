package rest

import (
	"net/http"

	"github.com/cameronbrill/brill-wtf-go/rest/controller"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/cameronbrill/brill-wtf-go/web"

	"github.com/go-chi/chi/v5"
)

type linkRouter struct {
	c        *controller.LinkServiceController
	renderer web.Renderer
}

func RegisterLinkServiceRouter(svc service.Service, r *chi.Mux, opts ...Option) {
	var router linkRouter

	for _, opt := range opts {
		opt(&router)
	}

	ctrl := controller.New(svc, router.renderer)
	router.c = &ctrl

	r.Mount("/", router.routes(true))
	r.Mount("/links", router.routes())
	r.Mount("/link", router.routes())
}

func (r linkRouter) routes(args ...bool) chi.Router {
	mountedRouter := chi.NewRouter()

	mountedRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome to the link service\nroute to /new?link={link} to create a new short link\notherwise, route to /{link} to get the original link"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mountedRouter.Route("/new", func(subRouter chi.Router) {
		subRouter.Use(linkCtx)
		subRouter.HandleFunc("/", r.c.NewLink)
	})

	mountedRouter.Route("/{link}", func(subRouter chi.Router) {
		subRouter.Use(shortLinkCtx)
		handlerFunc := r.c.ShortURLToLink
		if len(args) == 1 && args[0] {
			handlerFunc = r.c.ShortURLToLinkRedirect
		}
		subRouter.Get("/", handlerFunc)
	})
	return mountedRouter
}

type Option func(*linkRouter) error

func WithRenderer(renderer web.Renderer) Option {
	return func(r *linkRouter) error {
		r.renderer = renderer
		return nil
	}
}
