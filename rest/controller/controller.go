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
	"github.com/go-chi/chi/v5"
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

// Shorten URL to link
//
//		@summary        Given a URL, shortens it, stores it, and returns it as a slug to append to brill.wtf/{slug}
//		@tags           link
//		@produce        json

//	 @param          slug      path     string  true    "slug"
//	 @param          want      query    string  true    "want"

// @success        200
// @failure        500
// @router         /link [get]
func (c LinkServiceController) ShortURLToLink(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "slug")

	if shortLink == "" {
		shortLink = r.URL.Query().Get("want")
		if shortLink == "" {
			http.Error(w, "shortLink not found in context", http.StatusNotFound)
			return
		}
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
