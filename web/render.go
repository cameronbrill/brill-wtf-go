package web

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Renderer interface {
	Render(http.ResponseWriter, *http.Request, Page, ...RenderOption)
}

type Page string

const (
	NotFound Page = "404"
)

func New() Renderer {
	return &PageRenderer{
		data: &pageData{},
	}
}

type PageRenderer struct {
	data *pageData
}

type pageData struct {
	Short string
	Orig  string
	TTL   time.Duration
}

func (p *PageRenderer) Render(w http.ResponseWriter, r *http.Request, page Page, opts ...RenderOption) {
	for _, opt := range opts {
		opt(p)
	}

	t, err := template.ParseFiles(fmt.Sprintf("web/%s.gohtml", page))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, p.data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
