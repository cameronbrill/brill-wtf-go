package service

import (
	"fmt"
	"net/url"

	"github.com/cameronbrill/brill-wtf-go/model"
	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/generator"
)

type s struct {
	src Storage
}

// New instantiates a new service.
func New(options ...ServiceOption) *s {
	var svc s
	for _, option := range options {
		if err := option(&svc); err != nil {
			panic(err)
		}
	}
	if err := svc.src.Connect(); err != nil {
		panic(err)
	}
	return &svc
}

func (s *s) NewLink(orig string, options ...NewLinkOption) (model.Link, error) {
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
		if _, err := s.src.Get(link.Want); err != ErrNotFound {
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

	s.src.Set(link.Short, link)
	return link, nil
}

func (s *s) ShortURLToLink(short string) (model.Link, error) {
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
