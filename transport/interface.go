package transport

import "github.com/cameronbrill/brill-wtf-go/model"

type Storage interface {
	Connect() error
	Get(key string) (model.Link, error)
	Set(key string, value model.Link) error
}
