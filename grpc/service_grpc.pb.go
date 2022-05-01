// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LinkServiceClient is the client API for LinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkServiceClient interface {
	NewLink(ctx context.Context, in *NewLinkRequest, opts ...grpc.CallOption) (*NewLinkResponse, error)
	ShortURLToLink(ctx context.Context, in *ShortURLToLinkRequest, opts ...grpc.CallOption) (*ShortURLToLinkResponse, error)
}

type linkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkServiceClient(cc grpc.ClientConnInterface) LinkServiceClient {
	return &linkServiceClient{cc}
}

func (c *linkServiceClient) NewLink(ctx context.Context, in *NewLinkRequest, opts ...grpc.CallOption) (*NewLinkResponse, error) {
	out := new(NewLinkResponse)
	err := c.cc.Invoke(ctx, "/grpc.LinkService/NewLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkServiceClient) ShortURLToLink(ctx context.Context, in *ShortURLToLinkRequest, opts ...grpc.CallOption) (*ShortURLToLinkResponse, error) {
	out := new(ShortURLToLinkResponse)
	err := c.cc.Invoke(ctx, "/grpc.LinkService/ShortURLToLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServiceServer is the server API for LinkService service.
// All implementations must embed UnimplementedLinkServiceServer
// for forward compatibility
type LinkServiceServer interface {
	NewLink(context.Context, *NewLinkRequest) (*NewLinkResponse, error)
	ShortURLToLink(context.Context, *ShortURLToLinkRequest) (*ShortURLToLinkResponse, error)
	mustEmbedUnimplementedLinkServiceServer()
}

// UnimplementedLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLinkServiceServer struct {
}

func (UnimplementedLinkServiceServer) NewLink(context.Context, *NewLinkRequest) (*NewLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewLink not implemented")
}
func (UnimplementedLinkServiceServer) ShortURLToLink(context.Context, *ShortURLToLinkRequest) (*ShortURLToLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortURLToLink not implemented")
}
func (UnimplementedLinkServiceServer) mustEmbedUnimplementedLinkServiceServer() {}

// UnsafeLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkServiceServer will
// result in compilation errors.
type UnsafeLinkServiceServer interface {
	mustEmbedUnimplementedLinkServiceServer()
}

func RegisterLinkServiceServer(s grpc.ServiceRegistrar, srv LinkServiceServer) {
	s.RegisterService(&LinkService_ServiceDesc, srv)
}

func _LinkService_NewLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).NewLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LinkService/NewLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).NewLink(ctx, req.(*NewLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkService_ShortURLToLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortURLToLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).ShortURLToLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.LinkService/ShortURLToLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).ShortURLToLink(ctx, req.(*ShortURLToLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkService_ServiceDesc is the grpc.ServiceDesc for LinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.LinkService",
	HandlerType: (*LinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewLink",
			Handler:    _LinkService_NewLink_Handler,
		},
		{
			MethodName: "ShortURLToLink",
			Handler:    _LinkService_ShortURLToLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
