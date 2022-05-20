package model

import "time"

// Link is a Link business object.
type Link struct {
	ID       int64         `json:"-"`
	Original string        `json:"original"`
	Short    string        `json:"short"`
	Want     string        `json:"-"`
	TTL      time.Duration `json:"-"`
}
