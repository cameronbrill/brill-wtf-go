package link

import (
	"strings"
	"testing"
)

func TestNewLink(t *testing.T) {
        type linkTest struct {
                name string
                options []Option
                isWord bool
                wantTinyURL string
        }

        tests := []linkTest{
                {name: "default", wantTinyURL: "dfsadf"},
                {name: "specify word pattern", options: []Option{WithWordPattern()}, isWord: true, wantTinyURL: "this-is-word"},
                {name: "specify character pattern", options: []Option{WithCharacterPattern()}, isWord: false, wantTinyURL: "dfsadf"},
                {name: "specify tinyUrl", options: []Option{TinyURL("potatoe-potato")}, isWord: false, wantTinyURL: "potatoe-potato"},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        l, err := New(tt.options...)
                        if err != nil {
                                t.Fatalf("unexpected error: %v", err)
                        }
                        switch tt.isWord {
                        case true:
                                if len(strings.Split((&l).String(), "-")) != len(strings.Split(tt.wantTinyURL, "-")) {
                                        t.Fatalf("actual != expected [%s != %s]", &l, tt.wantTinyURL)
                                }
                        case false:
                                if len(tt.wantTinyURL) != len((&l).String()) {
                                        t.Fatalf("actual != expected [%s != %s]", &l, tt.wantTinyURL)
                                }
                        }
                })
        }
}