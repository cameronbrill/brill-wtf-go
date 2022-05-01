package main

import (
	"net"
	"os"

	tgrpc "github.com/cameronbrill/brill-wtf-go/grpc"
	"github.com/cameronbrill/brill-wtf-go/grpc/controller"
	"github.com/cameronbrill/brill-wtf-go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	svc := service.New()
	ctrl := controller.New(svc)

	server := grpc.NewServer()
	tgrpc.RegisterUserServiceServer(server, ctrl)
	reflection.Register(server)

	conn, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}

	print("starting grpc server...\n")
	err = server.Serve(conn)
	if err != nil {
		panic(err)
	}
}
