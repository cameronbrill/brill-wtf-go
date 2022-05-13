package client

import (
	"context"

	tgrpc "github.com/cameronbrill/brill-wtf-go/grpc"
	"github.com/cameronbrill/brill-wtf-go/model"
	"github.com/cameronbrill/brill-wtf-go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type svc struct {
	client tgrpc.LinkServiceClient
}

func New(conn string) (service.Service, error) {
	c, err := grpc.Dial(conn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &svc{
		client: tgrpc.NewLinkServiceClient(c),
	}, nil
}

func (s *svc) NewLink(orig string, options ...service.NewLinkOption) (model.Link, error) {
	link, err := s.client.NewLink(context.Background(), &tgrpc.NewLinkRequest{Original: orig})
	var Link model.Link
	if err != nil {
		return Link, err
	}
	Link = unmarshalLink(link.Link)
	return Link, nil
}

func (s *svc) ShortURLToLink(shortURL string) (model.Link, error) {
	link, err := s.client.ShortURLToLink(context.Background(), &tgrpc.ShortURLToLinkRequest{Short: shortURL})
	var Link model.Link
	if err != nil {
		return Link, err
	}
	Link = unmarshalLink(link.Link)
	return Link, nil
}

func unmarshalLink(u *tgrpc.Link) model.Link {
	return model.Link{
		ID:       u.Id,
		Original: u.Original,
		Short:    u.Short,
	}
}
