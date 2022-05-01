package controller

import (
	"context"

	"github.com/cameronbrill/brill-wtf-go/grpc"
	"github.com/cameronbrill/brill-wtf-go/service"
)

type LinkServiceController struct {
	LinkService service.Service
	grpc.UnimplementedLinkServiceServer
}

func New(svc service.Service) grpc.LinkServiceServer {
	return LinkServiceController{
		LinkService: svc,
	}
}

func (c LinkServiceController) NewLink(ctx context.Context, req *grpc.NewLinkRequest) (*grpc.NewLinkResponse, error) {
	link, err := c.LinkService.NewLink(req.Original)
	if err != nil {
		return nil, err
	}

	var resp grpc.NewLinkResponse

	resp.Link = marshalLink(link)

	return &resp, nil
}

func (c LinkServiceController) ShortURLToLink(ctx context.Context, req *grpc.ShortURLToLinkRequest) (*grpc.ShortURLToLinkResponse, error) {
	link, err := c.LinkService.ShortURLToLink(req.Short)
	if err != nil {
		return nil, err
	}

	var resp grpc.ShortURLToLinkResponse

	resp.Link = marshalLink(link)

	return &resp, nil
}

func marshalLink(l service.Link) *grpc.Link {
	return &grpc.Link{
		Id:       l.ID,
		Original: l.Original,
		Short:    l.Short,
	}
}
