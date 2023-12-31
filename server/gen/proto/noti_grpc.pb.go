// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/noti.proto

package proto

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

// NotiServiceClient is the client API for NotiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotiServiceClient interface {
	GetAllNotiByUserId(ctx context.Context, in *GetAllNotiByUserIdRequest, opts ...grpc.CallOption) (*GetAllNotiByUserIdResponse, error)
	ReadNoti(ctx context.Context, in *ReadNotiRequest, opts ...grpc.CallOption) (*ReadNotiResponse, error)
}

type notiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotiServiceClient(cc grpc.ClientConnInterface) NotiServiceClient {
	return &notiServiceClient{cc}
}

func (c *notiServiceClient) GetAllNotiByUserId(ctx context.Context, in *GetAllNotiByUserIdRequest, opts ...grpc.CallOption) (*GetAllNotiByUserIdResponse, error) {
	out := new(GetAllNotiByUserIdResponse)
	err := c.cc.Invoke(ctx, "/NotiService/GetAllNotiByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notiServiceClient) ReadNoti(ctx context.Context, in *ReadNotiRequest, opts ...grpc.CallOption) (*ReadNotiResponse, error) {
	out := new(ReadNotiResponse)
	err := c.cc.Invoke(ctx, "/NotiService/ReadNoti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotiServiceServer is the server API for NotiService service.
// All implementations must embed UnimplementedNotiServiceServer
// for forward compatibility
type NotiServiceServer interface {
	GetAllNotiByUserId(context.Context, *GetAllNotiByUserIdRequest) (*GetAllNotiByUserIdResponse, error)
	ReadNoti(context.Context, *ReadNotiRequest) (*ReadNotiResponse, error)
	mustEmbedUnimplementedNotiServiceServer()
}

// UnimplementedNotiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotiServiceServer struct {
}

func (UnimplementedNotiServiceServer) GetAllNotiByUserId(context.Context, *GetAllNotiByUserIdRequest) (*GetAllNotiByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotiByUserId not implemented")
}
func (UnimplementedNotiServiceServer) ReadNoti(context.Context, *ReadNotiRequest) (*ReadNotiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNoti not implemented")
}
func (UnimplementedNotiServiceServer) mustEmbedUnimplementedNotiServiceServer() {}

// UnsafeNotiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotiServiceServer will
// result in compilation errors.
type UnsafeNotiServiceServer interface {
	mustEmbedUnimplementedNotiServiceServer()
}

func RegisterNotiServiceServer(s grpc.ServiceRegistrar, srv NotiServiceServer) {
	s.RegisterService(&NotiService_ServiceDesc, srv)
}

func _NotiService_GetAllNotiByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNotiByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotiServiceServer).GetAllNotiByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotiService/GetAllNotiByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotiServiceServer).GetAllNotiByUserId(ctx, req.(*GetAllNotiByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotiService_ReadNoti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadNotiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotiServiceServer).ReadNoti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotiService/ReadNoti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotiServiceServer).ReadNoti(ctx, req.(*ReadNotiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotiService_ServiceDesc is the grpc.ServiceDesc for NotiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotiService",
	HandlerType: (*NotiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllNotiByUserId",
			Handler:    _NotiService_GetAllNotiByUserId_Handler,
		},
		{
			MethodName: "ReadNoti",
			Handler:    _NotiService_ReadNoti_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/noti.proto",
}
