package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"internal/pcontext"

	"github.com/cameronbrill/brill-wtf-go/model"
	"github.com/cameronbrill/brill-wtf-go/service"
)

type LinkServiceController struct {
	LinkService service.Service
}

func New(svc service.Service) LinkServiceController {
	return LinkServiceController{
		LinkService: svc,
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
	}
}

func (c LinkServiceController) ShortURLToLink(w http.ResponseWriter, r *http.Request) {
	link, err := c.getLink(w, r)
	if err != nil {
		return
	}
	jsonLink, err := json.Marshal(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(jsonLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c LinkServiceController) ShortURLToLinkRedirect(w http.ResponseWriter, r *http.Request) {
	link, err := c.getLink(w, r)
	if err != nil {
		return
	}
	http.Redirect(w, r, link.Original, http.StatusFound)
}

func (c LinkServiceController) getLink(w http.ResponseWriter, r *http.Request) (*model.Link, error) {
	ctx := r.Context()
	shortLink, ok := ctx.Value(pcontext.Link).(string)
	if !ok {
		http.Error(w, "shortLink not found in context", http.StatusBadRequest)
		return nil, fmt.Errorf("shortLink not found in context")
	}
	link, err := c.LinkService.ShortURLToLink(shortLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return &link, nil
}
