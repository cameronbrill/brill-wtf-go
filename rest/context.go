package rest

import (
	"context"
	"net/http"
	"time"

	"internal/pcontext"

	"github.com/go-chi/chi/v5"
)

func shortLinkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := chi.URLParam(r, "link")

		ctx := context.WithValue(r.Context(), pcontext.Link, link)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := r.URL.Query().Get("link")

		if link == "" {
			http.Redirect(w, r, "/", http.StatusFound)
		}

		ctx := context.WithValue(r.Context(), pcontext.Link, link)
		ctx = context.WithValue(ctx, pcontext.Want, r.URL.Query().Get("want"))
		if r.URL.Query().Get("forever") == "true" {
			ctx = context.WithValue(ctx, pcontext.TTL, time.Duration(0))
		} else {
			ctx = context.WithValue(ctx, pcontext.TTL, 72*time.Hour)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
