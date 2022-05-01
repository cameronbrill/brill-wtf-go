package service

import "errors"

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// Link is a Link business object.
type Link struct {
	ID       int64
	Original string
	Short    string
}

// Service defines the interface exposed by this package.
type Service interface {
	NewLink(string) (Link, error)
	ShortURLToLink(string) (Link, error)
}
