package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type key string

const (
	linkKey key = "link"
)

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")
		if l := r.URL.Query().Get("link"); l != "" {
			link = l
		}
		ctx := context.WithValue(r.Context(), linkKey, link)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
