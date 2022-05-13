package service

	"github.com/cameronbrill/brill-wtf-go/model"
)

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
	NewLink(string, ...NewLinkOption) (model.Link, error)
	ShortURLToLink(string) (model.Link, error)
}
