package controller

import (
	"encoding/json"
	"errors"
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

type requestBody struct {
	Original string `json:"original"`
}

func (c LinkServiceController) NewLink(w http.ResponseWriter, r *http.Request) {
	var bod requestBody
	err := json.NewDecoder(r.Body).Decode(&bod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	originalLink := bod.Original
	ctx := r.Context()
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
	shortLink := r.URL.Query().Get("want")
	if shortLink == "" {
		http.Error(w, "shortLink not found in context", http.StatusNotFound)
		return
	}
	var link model.Link
	link, err := c.LinkService.ShortURLToLink(shortLink)
	if err != nil {
		if errors.Is(err, cerrors.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shouldRedirect := r.URL.Query().Get("should-redirect") != ""
	if shouldRedirect {
		http.Redirect(w, r, link.Original, http.StatusFound)
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
