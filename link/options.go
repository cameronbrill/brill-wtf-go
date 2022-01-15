package link

type Options struct {
	URL     string
	TinyURL string // the string for https://brill.wtf/${TinyURL}
	Pattern Pattern
}

type Option func(*Options)

func DefaultOptions() *Options {
	return &Options{
		Pattern: Patterns.Character,
	}
}

func URL(url string) Option {
	return func(o *Options) {
		o.URL = url
	}
}

func TinyURL(postfix string) Option {
	return func(o *Options) {
		o.TinyURL = postfix
	}
}

// WithWordPattern() generates short urls with random words
// example: https://brill.wtf/potato-bean-soup
func WithWordPattern() Option {
	return func(o *Options) {
		o.Pattern = Patterns.Words
	}
}

// WithCharacterPattern() generates short urls with random characters
// example: https://brill.wtf/a84nbfjEE6b3
func WithCharacterPattern() Option {
	return func(o *Options) {
		o.Pattern = Patterns.Character
	}
}
