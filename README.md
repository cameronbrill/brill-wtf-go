[![Go Report Card](https://goreportcard.com/badge/github.com/cameronbrill/brill-wtf-go)](https://goreportcard.com/report/github.com/cameronbrill/brill-wtf-go)
[![GoDoc](https://godoc.org/github.com/cameronbrill/brill-wtf-go?status.svg)](https://godoc.org/github.com/cameronbrill/brill-wtf-go)

# go project template

Contains common things that I add to all of my go projects.

`make all`

### adding new executable

add any `cmd/*` path to `Makefile` like so:

```
SUB_DIRS=example grpc/server rest
```
