package controller

import (
	"encoding/json"
	"net/http"

	"internal/pcontext"

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
	link, err := c.LinkService.NewLink(originalLink)
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
	link := c.getLink(w, r)
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
	link := c.getLink(w, r)
	http.Redirect(w, r, link.Original, http.StatusFound)
}

func (c LinkServiceController) getLink(w http.ResponseWriter, r *http.Request) *service.Link {
	ctx := r.Context()
	shortLink, ok := ctx.Value("link").(string)
	if !ok {
		http.Error(w, "shortLink not found in context", http.StatusBadRequest)
		return nil
	}
	link, err := c.LinkService.ShortURLToLink(shortLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return &link
}
