package service

import (
	"errors"
	"fmt"
	"net/url"

	myErr "github.com/cameronbrill/brill-wtf-go/errors"
	"github.com/cameronbrill/brill-wtf-go/model"
	"github.com/cameronbrill/brill-wtf-go/transport"

	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/generator"
	"github.com/go-redis/redis/v8"
)

type svc struct {
	src transport.Storage
}

// New instantiates a new service.
func New(options ...ServiceOption) *svc {
	var s svc
	for _, option := range options {
		if err := option(&s); err != nil {
			panic(err)
		}
	}
	if err := s.src.Connect(); err != nil {
		panic(err)
	}
	return &s
}

func (s *svc) NewLink(orig string, options ...NewLinkOption) (model.Link, error) {
	link := model.Link{
		Original: orig,
	}
	for _, option := range options {
		err := option(&link)
		if err != nil {
			return link, err
		}
	}

	if err := isURL(orig); err != nil {
		return link, fmt.Errorf("invalid original URL: %w", err)
	}

	if link.Want != "" {
		if _, err := s.src.Get(link.Want); !(errors.Is(err, myErr.ErrNotFound) || errors.Is(err, redis.Nil)) {
			return link, fmt.Errorf("short URL already exists: %s", link.Want)
		}
		link.Short = link.Want
	} else {
		l, err := generator.NewXKPassword(&config.GeneratorConfig{
			NumWords:           3,
			WordLenMin:         3,
			WordLenMax:         6,
			SeparatorCharacter: "-",
			CaseTransform:      "LOWER",
			PaddingType:        "FIXED",
		}).Generate()
		if err != nil {
			return link, fmt.Errorf("generating link: %w", err)
		}
		link.Short = l
	}

	err := s.src.Set(link.Short, link)
	if err != nil {
		return link, fmt.Errorf("saving link: %w", err)
	}
	return link, nil
}

func (s *svc) ShortURLToLink(short string) (model.Link, error) {
	link, err := s.src.Get(short)
	if err != nil {
		return model.Link{}, fmt.Errorf("invalid short URL (%s): %w", short, err)
	}
	return link, nil
}

func isURL(str string) error {
	u, err := url.Parse(str)
	if err != nil {
		return err
	}
	if u.Scheme == "" {
		return fmt.Errorf("missing scheme, replace %s with https://%s", str, str)
	}
	if u.Host == "" {
		return fmt.Errorf("invalid URL: %s", str)
	}
	return nil
}
