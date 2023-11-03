// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/teamjoinrequest.proto

package team_join_request_pb

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

// TeamJoinRequestServiceClient is the client API for TeamJoinRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamJoinRequestServiceClient interface {
	GetAllJoinRequests(ctx context.Context, in *TeamJoinReqeustEmpty, opts ...grpc.CallOption) (*TeamJoinRequestList, error)
	GetJoinRequestById(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error)
	GetJoinRequestByTeamId(ctx context.Context, in *TeamJoinRequestTeamId, opts ...grpc.CallOption) (*TeamJoinRequestList, error)
	CreateJoinRequest(ctx context.Context, in *TeamJoinRequest, opts ...grpc.CallOption) (*TeamJoinRequest, error)
	UpdateJoinRequest(ctx context.Context, in *TeamJoinRequest, opts ...grpc.CallOption) (*TeamJoinRequest, error)
	DeleteJoinRequest(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error)
	ApproveJoinRequest(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error)
	DeclineJoinRequest(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error)
}

type teamJoinRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamJoinRequestServiceClient(cc grpc.ClientConnInterface) TeamJoinRequestServiceClient {
	return &teamJoinRequestServiceClient{cc}
}

func (c *teamJoinRequestServiceClient) GetAllJoinRequests(ctx context.Context, in *TeamJoinReqeustEmpty, opts ...grpc.CallOption) (*TeamJoinRequestList, error) {
	out := new(TeamJoinRequestList)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/GetAllJoinRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamJoinRequestServiceClient) GetJoinRequestById(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error) {
	out := new(TeamJoinRequest)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/GetJoinRequestById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamJoinRequestServiceClient) GetJoinRequestByTeamId(ctx context.Context, in *TeamJoinRequestTeamId, opts ...grpc.CallOption) (*TeamJoinRequestList, error) {
	out := new(TeamJoinRequestList)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/GetJoinRequestByTeamId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamJoinRequestServiceClient) CreateJoinRequest(ctx context.Context, in *TeamJoinRequest, opts ...grpc.CallOption) (*TeamJoinRequest, error) {
	out := new(TeamJoinRequest)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/CreateJoinRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamJoinRequestServiceClient) UpdateJoinRequest(ctx context.Context, in *TeamJoinRequest, opts ...grpc.CallOption) (*TeamJoinRequest, error) {
	out := new(TeamJoinRequest)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/UpdateJoinRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamJoinRequestServiceClient) DeleteJoinRequest(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error) {
	out := new(TeamJoinRequest)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/DeleteJoinRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamJoinRequestServiceClient) ApproveJoinRequest(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error) {
	out := new(TeamJoinRequest)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/ApproveJoinRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamJoinRequestServiceClient) DeclineJoinRequest(ctx context.Context, in *TeamJoinRequestId, opts ...grpc.CallOption) (*TeamJoinRequest, error) {
	out := new(TeamJoinRequest)
	err := c.cc.Invoke(ctx, "/TeamJoinRequestService/DeclineJoinRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamJoinRequestServiceServer is the server API for TeamJoinRequestService service.
// All implementations must embed UnimplementedTeamJoinRequestServiceServer
// for forward compatibility
type TeamJoinRequestServiceServer interface {
	GetAllJoinRequests(context.Context, *TeamJoinReqeustEmpty) (*TeamJoinRequestList, error)
	GetJoinRequestById(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error)
	GetJoinRequestByTeamId(context.Context, *TeamJoinRequestTeamId) (*TeamJoinRequestList, error)
	CreateJoinRequest(context.Context, *TeamJoinRequest) (*TeamJoinRequest, error)
	UpdateJoinRequest(context.Context, *TeamJoinRequest) (*TeamJoinRequest, error)
	DeleteJoinRequest(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error)
	ApproveJoinRequest(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error)
	DeclineJoinRequest(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error)
	mustEmbedUnimplementedTeamJoinRequestServiceServer()
}

// UnimplementedTeamJoinRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamJoinRequestServiceServer struct {
}

func (UnimplementedTeamJoinRequestServiceServer) GetAllJoinRequests(context.Context, *TeamJoinReqeustEmpty) (*TeamJoinRequestList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllJoinRequests not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) GetJoinRequestById(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJoinRequestById not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) GetJoinRequestByTeamId(context.Context, *TeamJoinRequestTeamId) (*TeamJoinRequestList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJoinRequestByTeamId not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) CreateJoinRequest(context.Context, *TeamJoinRequest) (*TeamJoinRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJoinRequest not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) UpdateJoinRequest(context.Context, *TeamJoinRequest) (*TeamJoinRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJoinRequest not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) DeleteJoinRequest(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJoinRequest not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) ApproveJoinRequest(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveJoinRequest not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) DeclineJoinRequest(context.Context, *TeamJoinRequestId) (*TeamJoinRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclineJoinRequest not implemented")
}
func (UnimplementedTeamJoinRequestServiceServer) mustEmbedUnimplementedTeamJoinRequestServiceServer() {
}

// UnsafeTeamJoinRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamJoinRequestServiceServer will
// result in compilation errors.
type UnsafeTeamJoinRequestServiceServer interface {
	mustEmbedUnimplementedTeamJoinRequestServiceServer()
}

func RegisterTeamJoinRequestServiceServer(s grpc.ServiceRegistrar, srv TeamJoinRequestServiceServer) {
	s.RegisterService(&TeamJoinRequestService_ServiceDesc, srv)
}

func _TeamJoinRequestService_GetAllJoinRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinReqeustEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).GetAllJoinRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/GetAllJoinRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).GetAllJoinRequests(ctx, req.(*TeamJoinReqeustEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamJoinRequestService_GetJoinRequestById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).GetJoinRequestById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/GetJoinRequestById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).GetJoinRequestById(ctx, req.(*TeamJoinRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamJoinRequestService_GetJoinRequestByTeamId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinRequestTeamId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).GetJoinRequestByTeamId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/GetJoinRequestByTeamId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).GetJoinRequestByTeamId(ctx, req.(*TeamJoinRequestTeamId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamJoinRequestService_CreateJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).CreateJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/CreateJoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).CreateJoinRequest(ctx, req.(*TeamJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamJoinRequestService_UpdateJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).UpdateJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/UpdateJoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).UpdateJoinRequest(ctx, req.(*TeamJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamJoinRequestService_DeleteJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).DeleteJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/DeleteJoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).DeleteJoinRequest(ctx, req.(*TeamJoinRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamJoinRequestService_ApproveJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).ApproveJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/ApproveJoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).ApproveJoinRequest(ctx, req.(*TeamJoinRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamJoinRequestService_DeclineJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamJoinRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamJoinRequestServiceServer).DeclineJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TeamJoinRequestService/DeclineJoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamJoinRequestServiceServer).DeclineJoinRequest(ctx, req.(*TeamJoinRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamJoinRequestService_ServiceDesc is the grpc.ServiceDesc for TeamJoinRequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamJoinRequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TeamJoinRequestService",
	HandlerType: (*TeamJoinRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllJoinRequests",
			Handler:    _TeamJoinRequestService_GetAllJoinRequests_Handler,
		},
		{
			MethodName: "GetJoinRequestById",
			Handler:    _TeamJoinRequestService_GetJoinRequestById_Handler,
		},
		{
			MethodName: "GetJoinRequestByTeamId",
			Handler:    _TeamJoinRequestService_GetJoinRequestByTeamId_Handler,
		},
		{
			MethodName: "CreateJoinRequest",
			Handler:    _TeamJoinRequestService_CreateJoinRequest_Handler,
		},
		{
			MethodName: "UpdateJoinRequest",
			Handler:    _TeamJoinRequestService_UpdateJoinRequest_Handler,
		},
		{
			MethodName: "DeleteJoinRequest",
			Handler:    _TeamJoinRequestService_DeleteJoinRequest_Handler,
		},
		{
			MethodName: "ApproveJoinRequest",
			Handler:    _TeamJoinRequestService_ApproveJoinRequest_Handler,
		},
		{
			MethodName: "DeclineJoinRequest",
			Handler:    _TeamJoinRequestService_DeclineJoinRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/teamjoinrequest.proto",
}
