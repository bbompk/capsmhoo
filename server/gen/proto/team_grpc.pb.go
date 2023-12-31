// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: team.proto

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

const (
	TeamService_GetAllTeams_FullMethodName           = "/TeamService/GetAllTeams"
	TeamService_GetTeamById_FullMethodName           = "/TeamService/GetTeamById"
	TeamService_GetTeamByUserId_FullMethodName       = "/TeamService/GetTeamByUserId"
	TeamService_CreateTeam_FullMethodName            = "/TeamService/CreateTeam"
	TeamService_UpdateTeam_FullMethodName            = "/TeamService/UpdateTeam"
	TeamService_DeleteTeam_FullMethodName            = "/TeamService/DeleteTeam"
	TeamService_AddStudentToTeam_FullMethodName      = "/TeamService/AddStudentToTeam"
	TeamService_RemoveStudentFromTeam_FullMethodName = "/TeamService/RemoveStudentFromTeam"
)

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	GetAllTeams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TeamList, error)
	GetTeamById(ctx context.Context, in *TeamId, opts ...grpc.CallOption) (*Team, error)
	GetTeamByUserId(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Team, error)
	CreateTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*Team, error)
	UpdateTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*Team, error)
	DeleteTeam(ctx context.Context, in *TeamId, opts ...grpc.CallOption) (*Team, error)
	AddStudentToTeam(ctx context.Context, in *TeamAndStudentID, opts ...grpc.CallOption) (*Student, error)
	RemoveStudentFromTeam(ctx context.Context, in *TeamAndStudentID, opts ...grpc.CallOption) (*Empty, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) GetAllTeams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TeamList, error) {
	out := new(TeamList)
	err := c.cc.Invoke(ctx, TeamService_GetAllTeams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamById(ctx context.Context, in *TeamId, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetTeamById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamByUserId(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetTeamByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) CreateTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_CreateTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) UpdateTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_UpdateTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) DeleteTeam(ctx context.Context, in *TeamId, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_DeleteTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) AddStudentToTeam(ctx context.Context, in *TeamAndStudentID, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, TeamService_AddStudentToTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) RemoveStudentFromTeam(ctx context.Context, in *TeamAndStudentID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TeamService_RemoveStudentFromTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility
type TeamServiceServer interface {
	GetAllTeams(context.Context, *Empty) (*TeamList, error)
	GetTeamById(context.Context, *TeamId) (*Team, error)
	GetTeamByUserId(context.Context, *UserId) (*Team, error)
	CreateTeam(context.Context, *Team) (*Team, error)
	UpdateTeam(context.Context, *Team) (*Team, error)
	DeleteTeam(context.Context, *TeamId) (*Team, error)
	AddStudentToTeam(context.Context, *TeamAndStudentID) (*Student, error)
	RemoveStudentFromTeam(context.Context, *TeamAndStudentID) (*Empty, error)
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (UnimplementedTeamServiceServer) GetAllTeams(context.Context, *Empty) (*TeamList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTeams not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamById(context.Context, *TeamId) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamById not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamByUserId(context.Context, *UserId) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamByUserId not implemented")
}
func (UnimplementedTeamServiceServer) CreateTeam(context.Context, *Team) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (UnimplementedTeamServiceServer) UpdateTeam(context.Context, *Team) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedTeamServiceServer) DeleteTeam(context.Context, *TeamId) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedTeamServiceServer) AddStudentToTeam(context.Context, *TeamAndStudentID) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudentToTeam not implemented")
}
func (UnimplementedTeamServiceServer) RemoveStudentFromTeam(context.Context, *TeamAndStudentID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStudentFromTeam not implemented")
}
func (UnimplementedTeamServiceServer) mustEmbedUnimplementedTeamServiceServer() {}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	s.RegisterService(&TeamService_ServiceDesc, srv)
}

func _TeamService_GetAllTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetAllTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetAllTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetAllTeams(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamById(ctx, req.(*TeamId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamByUserId(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Team)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_CreateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).CreateTeam(ctx, req.(*Team))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Team)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_UpdateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).UpdateTeam(ctx, req.(*Team))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_DeleteTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).DeleteTeam(ctx, req.(*TeamId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_AddStudentToTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamAndStudentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).AddStudentToTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_AddStudentToTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).AddStudentToTeam(ctx, req.(*TeamAndStudentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_RemoveStudentFromTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamAndStudentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).RemoveStudentFromTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_RemoveStudentFromTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).RemoveStudentFromTeam(ctx, req.(*TeamAndStudentID))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamService_ServiceDesc is the grpc.ServiceDesc for TeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTeams",
			Handler:    _TeamService_GetAllTeams_Handler,
		},
		{
			MethodName: "GetTeamById",
			Handler:    _TeamService_GetTeamById_Handler,
		},
		{
			MethodName: "GetTeamByUserId",
			Handler:    _TeamService_GetTeamByUserId_Handler,
		},
		{
			MethodName: "CreateTeam",
			Handler:    _TeamService_CreateTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _TeamService_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _TeamService_DeleteTeam_Handler,
		},
		{
			MethodName: "AddStudentToTeam",
			Handler:    _TeamService_AddStudentToTeam_Handler,
		},
		{
			MethodName: "RemoveStudentFromTeam",
			Handler:    _TeamService_RemoveStudentFromTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team.proto",
}
