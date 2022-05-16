package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"internal/pcontext"

	cerrors "github.com/cameronbrill/brill-wtf-go/errors"
	"github.com/cameronbrill/brill-wtf-go/model"
	"github.com/cameronbrill/brill-wtf-go/service"
	"github.com/cameronbrill/brill-wtf-go/web"
)

type LinkServiceController struct {
	LinkService service.Service
	renderer    web.Renderer
}

func New(svc service.Service, renderer web.Renderer) LinkServiceController {
	return LinkServiceController{
		LinkService: svc,
		renderer:    renderer,
	}
}

func (c LinkServiceController) NewLink(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	originalLink, ok := ctx.Value(pcontext.Link).(string)
	if !ok {
		http.Error(w, "originalLink not found in context", http.StatusBadRequest)
		return
	}
	want, ok := ctx.Value(pcontext.Want).(string)
	if !ok {
		http.Error(w, "want not string in context", http.StatusInternalServerError)
		return
	}
	ttl, ok := ctx.Value(pcontext.TTL).(time.Duration)
	if !ok {
		ttl = 72 * time.Hour
	}
	link, err := c.LinkService.NewLink(originalLink, service.WithShortURL(want), service.WithTTL(ttl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonLink, err := json.Marshal(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c LinkServiceController) ShortURLToLink(w http.ResponseWriter, r *http.Request) {
	link, err := c.getLink(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonLink, err := json.Marshal(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c LinkServiceController) ShortURLToLinkRedirect(w http.ResponseWriter, r *http.Request) {
	link, err := c.getLink(r)
	if err != nil {
		if errors.Is(err, cerrors.ErrNotFound) && c.renderer != nil {
			c.renderer.Render(w, r, "404", web.WithShortURL(link.Short))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, link.Original, http.StatusFound)
}

func (c LinkServiceController) getLink(r *http.Request) (*model.Link, error) {
	ctx := r.Context()
	shortLink, ok := ctx.Value(pcontext.Link).(string)
	if !ok {
		return nil, fmt.Errorf("shortLink not found in context")
	}
	link, err := c.LinkService.ShortURLToLink(shortLink)
	if err != nil {
		return &model.Link{Short: shortLink}, err
	}
	return &link, nil
}
