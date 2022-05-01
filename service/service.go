package service

import (
	"fmt"

	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/generator"
)

type s struct {
	// a database dependency would go here but instead we're going to have a static map
	m map[string]string
}

// New instantiates a new service.
func New( /* a database connection would be injected here */ ) *s {
	return &s{
		m: map[string]string{
			"abc": "https://github.com/cameronbrill/create-go-app",
		},
	}
}

func (s *s) NewShortURL(orig string) (string, error) {
	l, err := generator.NewXKPassword(&config.GeneratorConfig{
		NumWords:           3,
		WordLenMin:         3,
		WordLenMax:         6,
		SeparatorCharacter: "-",
		CaseTransform:      "LOWER",
		PaddingType:        "FIXED",
	}).Generate()
	if err != nil {
		return "", fmt.Errorf("generating link: %w", err)
	}
	return l, nil
}

func (s *s) ShortToLong(short string) (string, error) {
	return "", nil
}
