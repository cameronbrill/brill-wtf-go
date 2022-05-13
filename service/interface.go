package service

import (
	"errors"

	"github.com/cameronbrill/brill-wtf-go/model"
)

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")
var ErrAlreadyExits = errors.New("already exists")

// Service defines the interface exposed by this package.
type Service interface {
	NewLink(string, ...NewLinkOption) (model.Link, error)
	ShortURLToLink(string) (model.Link, error)
}
