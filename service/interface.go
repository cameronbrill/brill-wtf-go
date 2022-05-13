package service

import (
	"github.com/cameronbrill/brill-wtf-go/model"
)

// Service defines the interface exposed by this package.
type Service interface {
	NewLink(string, ...NewLinkOption) (model.Link, error)
	ShortURLToLink(string) (model.Link, error)
}
