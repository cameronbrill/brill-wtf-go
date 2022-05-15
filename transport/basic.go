package transport

import (
	myErr "github.com/cameronbrill/brill-wtf-go/errors"
	"github.com/cameronbrill/brill-wtf-go/model"
)

type Basic struct {
	m map[string]model.Link
}

func (s *Basic) Connect() error {
	s.m = make(map[string]model.Link)
	return nil
}

func (b *Basic) Get(key string) (model.Link, error) {
	link, ok := b.m[key]
	if !ok {
		return model.Link{}, myErr.ErrNotFound
	}
	return link, nil
}

func (b *Basic) Set(key string, value model.Link) error {
	b.m[key] = value
	return nil
}
