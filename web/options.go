package web

import "time"

type RenderOption func(*PageRenderer) error

func WithShortURL(val string) RenderOption {
	return func(r *PageRenderer) error {
		r.data.Short = val
		return nil
	}
}

func WithOriginalURL(val string) RenderOption {
	return func(r *PageRenderer) error {
		r.data.Orig = val
		return nil
	}
}

func WithTTL(val time.Duration) RenderOption {
	return func(r *PageRenderer) error {
		r.data.TTL = val
		return nil
	}
}
