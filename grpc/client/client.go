package client

import (
	"context"

	tgrpc "github.com/cameronbrill/go-project-template/grpc"
	"github.com/cameronbrill/go-project-template/service"
	"google.golang.org/grpc"
)

type svc struct {
	client tgrpc.UserServiceClient
}

func New(conn string) (service.Service, error) {
	c, err := grpc.Dial(conn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &svc{
		client: tgrpc.NewUserServiceClient(c),
	}, nil
}

func (s *svc) GetUsers(ids []int64) (map[int64]service.User, error) {
	users, err := s.client.GetUsers(context.Background(), &tgrpc.GetUsersRequest{Ids: ids})
	if err != nil {
		return nil, err
	}
	var ret map[int64]service.User
	for _, u := range users.Users {
		ret[u.Id] = unmarshalUser(u)
	}
	return ret, nil
}

func (s *svc) GetUser(id int64) (service.User, error) {
	u, err := s.client.GetUsers(context.Background(), &tgrpc.GetUsersRequest{Ids: []int64{id}})
	var user service.User
	if err != nil {
		return user, err
	}
	user = unmarshalUser(u.Users[0])
	return user, nil
}

func unmarshalUser(u *tgrpc.User) service.User {
	return service.User{
		ID:   u.Id,
		Name: u.Name,
	}
}
