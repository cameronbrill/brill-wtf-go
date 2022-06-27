package rest

import (
	"context"
	"net/http"
	"time"

	"internal/pcontext"
)

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), pcontext.Want, r.URL.Query().Get("want"))
		if r.URL.Query().Get("forever") == "true" {
			ctx = context.WithValue(ctx, pcontext.TTL, time.Duration(0))
		} else {
			ctx = context.WithValue(ctx, pcontext.TTL, 72*time.Hour)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
