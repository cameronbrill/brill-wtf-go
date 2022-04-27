package controller

import (
	"context"

	"github.com/cameronbrill/go-project-template/grpc"
	"github.com/cameronbrill/go-project-template/service"
)

type userServiceController struct {
	userService service.Service
	grpc.UnimplementedUserServiceServer
}

func New(svc service.Service) grpc.UserServiceServer {
	return userServiceController{
		userService: svc,
	}
}

func (c userServiceController) GetUsers(ctx context.Context, req *grpc.GetUsersRequest) (*grpc.GetUsersResponse, error) {
	resMap, err := c.userService.GetUsers(req.GetIds())
	if err != nil {
		return nil, err
	}

	var resp grpc.GetUsersResponse
	for _, user := range resMap {
		resp.Users = append(resp.Users, &grpc.User{
			Id:   user.ID,
			Name: user.Name,
		})
	}

	return &resp, nil
}
