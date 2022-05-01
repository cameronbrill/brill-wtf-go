package client

import (
	"context"

	tgrpc "github.com/cameronbrill/brill-wtf-go/grpc"
	"github.com/cameronbrill/brill-wtf-go/service"
	"google.golang.org/grpc"
)

type svc struct {
	client tgrpc.LinkServiceClient
}

func New(conn string) (service.Service, error) {
	c, err := grpc.Dial(conn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &svc{
		client: tgrpc.NewLinkServiceClient(c),
	}, nil
}

func (s *svc) NewLink(orig string) (service.Link, error) {
	link, err := s.client.NewLink(context.Background(), &tgrpc.NewLinkRequest{Original: orig})
	var Link service.Link
	if err != nil {
		return Link, err
	}
	Link = unmarshalLink(link.Link)
	return Link, nil
}

func (s *svc) ShortURLToLink(shortURL string) (service.Link, error) {
	link, err := s.client.ShortURLToLink(context.Background(), &tgrpc.ShortURLToLinkRequest{Short: shortURL})
	var Link service.Link
	if err != nil {
		return Link, err
	}
	Link = unmarshalLink(link.Link)
	return Link, nil
}

func unmarshalLink(u *tgrpc.Link) service.Link {
	return service.Link{
		ID:       u.Id,
		Original: u.Original,
		Short:    u.Short,
	}
}
