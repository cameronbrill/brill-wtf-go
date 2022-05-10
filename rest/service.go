package rest

import (
	"net/http"

	"github.com/cameronbrill/brill-wtf-go/rest/controller"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/go-chi/chi/v5"
)

type linkRouter struct {
	c *controller.LinkServiceController
}

func RegisterLinkServiceRouter(svc service.Service, r *chi.Mux) {
	var router linkRouter
	ctrl := controller.New(svc)
	router.c = &ctrl

	r.Mount("/", router.routes(true))
	r.Mount("/links", router.routes())
	r.Mount("/link", router.routes())
}

func (r linkRouter) routes(args ...bool) chi.Router {
	mountedRouter := chi.NewRouter()

	mountedRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to the link service\nroute to /new/{link} to create a new short link\notherwise, route to /{link} to get the original link"))
	})

	mountedRouter.Route("/new/{link}", func(subRouter chi.Router) {
		subRouter.Use(linkCtx)
		subRouter.HandleFunc("/", r.c.NewLink)
	})

	mountedRouter.Route("/{link}", func(subRouter chi.Router) {
		subRouter.Use(linkCtx)
		handlerFunc := r.c.ShortURLToLink
		if len(args) == 1 && args[0] {
			handlerFunc = r.c.ShortURLToLinkRedirect
		}
		subRouter.Get("/", handlerFunc)
	})
	return mountedRouter
}
