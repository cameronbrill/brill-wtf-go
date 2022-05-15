package service

import (
	"time"

	"github.com/cameronbrill/brill-wtf-go/model"
	"github.com/cameronbrill/brill-wtf-go/redis"
)

type NewLinkOption func(*model.Link) error

func WithShortURL(want string) NewLinkOption {
	return func(l *model.Link) error {
		l.Want = want
		return nil
	}
}

func WithTTL(ttl time.Duration) NewLinkOption {
	return func(l *model.Link) error {
		l.TTL = ttl
		return nil
	}
}

type ServiceOption func(*svc) error

func WithBasicStorage() ServiceOption {
	return func(svc *svc) error {
		return nil
	}
}

func WithRedisStorage() ServiceOption {
	return func(svc *svc) error {
		return nil
	}
}
