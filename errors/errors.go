package errors

import "errors"

// ErrNotFound signifies that a single requested object was not found.
var ErrNotFound = errors.New("not found")
var ErrAlreadyExits = errors.New("already exists")
