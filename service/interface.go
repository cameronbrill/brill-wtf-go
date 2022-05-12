package service

import "errors"

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")

// Link is a Link business object.
type Link struct {
	ID       int64
	Original string
	Short    string
	want     string
}

type NewLinkOption func(*Link) error

// Service defines the interface exposed by this package.
type Service interface {
	NewLink(string, ...NewLinkOption) (Link, error)
	ShortURLToLink(string) (Link, error)
}

func WithShort(want string) NewLinkOption {
	return func(l *Link) error {
		l.want = want
		return nil
	}
}
