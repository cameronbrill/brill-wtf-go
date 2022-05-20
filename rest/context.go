package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"internal/pcontext"
)

type requestBody struct {
	Original string `json:"original"`
}

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		link := r.URL.Query().Get("short")
		if link == "" {
			var bod requestBody
			err := json.NewDecoder(r.Body).Decode(&bod)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			link = bod.Original
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
