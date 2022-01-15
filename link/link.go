package link

import (
	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/generator"
	"github.com/pkg/errors"
)

type Link struct {
	url     string
	tinyUrl string
	pattern Pattern
}

func New(setters ...Option) (Link, error) {
	args := DefaultOptions()

	for _, setter := range setters {
		setter(args)
	}

	link := &Link{
		url:     args.URL,
		tinyUrl: args.TinyURL,
		pattern: args.Pattern,
	}

	if link.tinyUrl != "" {
		// TODO make sure this doesn't exist already
		return *link, nil
	}

	if link.tinyUrl == "" {
		// generate link
		err := link.new()
		if err != nil {
			return *link, err
		}
	}

	return *link, nil
}

func (link *Link) new() error {
	switch link.pattern {
	case Patterns.Character:
		link.tinyUrl = RandStringBytesMaskImprSrcSB(6)
		// validate no clashing
		return nil
	case Patterns.Words:
		l, err := generator.NewXKPassword(&config.GeneratorConfig{
			NumWords:           3,
			WordLenMin:         3,
			WordLenMax:         6,
			SeparatorCharacter: "-",
			CaseTransform:      "LOWER",
			PaddingType:        "FIXED",
		}).Generate()
		if err != nil {
			return errors.Wrap(err, "generating link")
		}
		link.tinyUrl = l
	default:
		return errors.New("invalid link generation pattern")
	}
	return nil
}

func (link *Link) String() string {
	return link.tinyUrl
}
