// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: comics.proto

package protobufs

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

// ComicsClient is the client API for Comics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComicsClient interface {
	GetComics(ctx context.Context, in *GetComicsRequest, opts ...grpc.CallOption) (*GetComicsResponse, error)
	GetComicFeeds(ctx context.Context, in *DescribeComicsRequest, opts ...grpc.CallOption) (*DescribeComicsResponse, error)
	UpdateComic(ctx context.Context, in *UpdateComicRequest, opts ...grpc.CallOption) (*GetComicsResponse, error)
	MarkAsRead(ctx context.Context, in *MarkReadRequest, opts ...grpc.CallOption) (*MarkReadResponse, error)
	MarkItemRead(ctx context.Context, in *MarkItemReadRequest, opts ...grpc.CallOption) (*MarkReadResponse, error)
	RefreshRssFeed(ctx context.Context, in *GetComicsRequest, opts ...grpc.CallOption) (*RefreshRssFeedResponse, error)
}

type comicsClient struct {
	cc grpc.ClientConnInterface
}

func NewComicsClient(cc grpc.ClientConnInterface) ComicsClient {
	return &comicsClient{cc}
}

func (c *comicsClient) GetComics(ctx context.Context, in *GetComicsRequest, opts ...grpc.CallOption) (*GetComicsResponse, error) {
	out := new(GetComicsResponse)
	err := c.cc.Invoke(ctx, "/Comics/GetComics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicsClient) GetComicFeeds(ctx context.Context, in *DescribeComicsRequest, opts ...grpc.CallOption) (*DescribeComicsResponse, error) {
	out := new(DescribeComicsResponse)
	err := c.cc.Invoke(ctx, "/Comics/GetComicFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicsClient) UpdateComic(ctx context.Context, in *UpdateComicRequest, opts ...grpc.CallOption) (*GetComicsResponse, error) {
	out := new(GetComicsResponse)
	err := c.cc.Invoke(ctx, "/Comics/UpdateComic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicsClient) MarkAsRead(ctx context.Context, in *MarkReadRequest, opts ...grpc.CallOption) (*MarkReadResponse, error) {
	out := new(MarkReadResponse)
	err := c.cc.Invoke(ctx, "/Comics/MarkAsRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicsClient) MarkItemRead(ctx context.Context, in *MarkItemReadRequest, opts ...grpc.CallOption) (*MarkReadResponse, error) {
	out := new(MarkReadResponse)
	err := c.cc.Invoke(ctx, "/Comics/MarkItemRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comicsClient) RefreshRssFeed(ctx context.Context, in *GetComicsRequest, opts ...grpc.CallOption) (*RefreshRssFeedResponse, error) {
	out := new(RefreshRssFeedResponse)
	err := c.cc.Invoke(ctx, "/Comics/RefreshRssFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComicsServer is the server API for Comics service.
// All implementations must embed UnimplementedComicsServer
// for forward compatibility
type ComicsServer interface {
	GetComics(context.Context, *GetComicsRequest) (*GetComicsResponse, error)
	GetComicFeeds(context.Context, *DescribeComicsRequest) (*DescribeComicsResponse, error)
	UpdateComic(context.Context, *UpdateComicRequest) (*GetComicsResponse, error)
	MarkAsRead(context.Context, *MarkReadRequest) (*MarkReadResponse, error)
	MarkItemRead(context.Context, *MarkItemReadRequest) (*MarkReadResponse, error)
	RefreshRssFeed(context.Context, *GetComicsRequest) (*RefreshRssFeedResponse, error)
	mustEmbedUnimplementedComicsServer()
}

// UnimplementedComicsServer must be embedded to have forward compatible implementations.
type UnimplementedComicsServer struct {
}

func (UnimplementedComicsServer) GetComics(context.Context, *GetComicsRequest) (*GetComicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComics not implemented")
}
func (UnimplementedComicsServer) GetComicFeeds(context.Context, *DescribeComicsRequest) (*DescribeComicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComicFeeds not implemented")
}
func (UnimplementedComicsServer) UpdateComic(context.Context, *UpdateComicRequest) (*GetComicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComic not implemented")
}
func (UnimplementedComicsServer) MarkAsRead(context.Context, *MarkReadRequest) (*MarkReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsRead not implemented")
}
func (UnimplementedComicsServer) MarkItemRead(context.Context, *MarkItemReadRequest) (*MarkReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkItemRead not implemented")
}
func (UnimplementedComicsServer) RefreshRssFeed(context.Context, *GetComicsRequest) (*RefreshRssFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshRssFeed not implemented")
}
func (UnimplementedComicsServer) mustEmbedUnimplementedComicsServer() {}

// UnsafeComicsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComicsServer will
// result in compilation errors.
type UnsafeComicsServer interface {
	mustEmbedUnimplementedComicsServer()
}

func RegisterComicsServer(s grpc.ServiceRegistrar, srv ComicsServer) {
	s.RegisterService(&Comics_ServiceDesc, srv)
}

func _Comics_GetComics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComicsServer).GetComics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comics/GetComics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComicsServer).GetComics(ctx, req.(*GetComicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comics_GetComicFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeComicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComicsServer).GetComicFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comics/GetComicFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComicsServer).GetComicFeeds(ctx, req.(*DescribeComicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comics_UpdateComic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateComicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComicsServer).UpdateComic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comics/UpdateComic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComicsServer).UpdateComic(ctx, req.(*UpdateComicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comics_MarkAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComicsServer).MarkAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comics/MarkAsRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComicsServer).MarkAsRead(ctx, req.(*MarkReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comics_MarkItemRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkItemReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComicsServer).MarkItemRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comics/MarkItemRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComicsServer).MarkItemRead(ctx, req.(*MarkItemReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comics_RefreshRssFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComicsServer).RefreshRssFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Comics/RefreshRssFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComicsServer).RefreshRssFeed(ctx, req.(*GetComicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Comics_ServiceDesc is the grpc.ServiceDesc for Comics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Comics",
	HandlerType: (*ComicsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComics",
			Handler:    _Comics_GetComics_Handler,
		},
		{
			MethodName: "GetComicFeeds",
			Handler:    _Comics_GetComicFeeds_Handler,
		},
		{
			MethodName: "UpdateComic",
			Handler:    _Comics_UpdateComic_Handler,
		},
		{
			MethodName: "MarkAsRead",
			Handler:    _Comics_MarkAsRead_Handler,
		},
		{
			MethodName: "MarkItemRead",
			Handler:    _Comics_MarkItemRead_Handler,
		},
		{
			MethodName: "RefreshRssFeed",
			Handler:    _Comics_RefreshRssFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comics.proto",
}
