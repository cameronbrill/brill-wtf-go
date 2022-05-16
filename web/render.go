package web

import (
	"net/http"
)

type Renderer interface {
	Render(http.ResponseWriter, *http.Request, Page, ...RenderOption)
}

type Page string
