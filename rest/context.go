package rest

import (
	"context"
	"net/http"

	"internal/pcontext"

	"github.com/go-chi/chi/v5"
)

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")

		if l := r.URL.Query().Get("link"); l != "" {
			link = l
		}

		ctx := context.WithValue(r.Context(), pcontext.Link, link)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
