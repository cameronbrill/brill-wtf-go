package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")
		ctx := context.WithValue(r.Context(), "link", link)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
