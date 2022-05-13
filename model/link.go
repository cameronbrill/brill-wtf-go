package model

import "time"

// Link is a Link business object.
type Link struct {
	ID       int64
	Original string
	Short    string
	Want     string
	TTL      time.Duration
}
