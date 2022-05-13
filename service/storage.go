package service

import (
	myErr "github.com/cameronbrill/brill-wtf-go/errors"
	"github.com/cameronbrill/brill-wtf-go/model"
)

type Storage interface {
	Connect() error
	Get(key string) (model.Link, error)
	Set(key string, value model.Link) error
}

type BasicStorage struct {
	m map[string]model.Link
}

func (s *BasicStorage) Connect() error {
	s.m = make(map[string]model.Link)
	return nil
}

func (b *BasicStorage) Get(key string) (model.Link, error) {
	link, ok := b.m[key]
	if !ok {
		return model.Link{}, myErr.ErrNotFound
	}
	return link, nil
}

func (b *BasicStorage) Set(key string, value model.Link) error {
	b.m[key] = value
	return nil
}
