// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	EmailValidation(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*None, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ValidateJWT(ctx context.Context, in *ValidateJWTRequest, opts ...grpc.CallOption) (*ValidateJWTResponse, error)
	GetUserFromID(ctx context.Context, in *GetUserFromIDRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserForForeignUID(ctx context.Context, in *GetForeignUserRequest, opts ...grpc.CallOption) (*GetForeignUserResponse, error)
	GetUserIDsForUsername(ctx context.Context, in *GetUserIDsForUsernameRequest, opts ...grpc.CallOption) (*GetUserIDsForUsernameResponse, error)
	BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*BanUserResponse, error)
	SetUserRank(ctx context.Context, in *SetRankRequest, opts ...grpc.CallOption) (*None, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*None, error)
	AddAuditEvent(ctx context.Context, in *NewAuditEventRequest, opts ...grpc.CallOption) (*None, error)
	GetAuditEvents(ctx context.Context, in *AuditEventsListRequest, opts ...grpc.CallOption) (*AuditListResponse, error)
	GetFollowers(ctx context.Context, in *FollowerReq, opts ...grpc.CallOption) (*FollowerResp, error)
	AddFolllower(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*None, error)
	SetProfile(ctx context.Context, in *ProfileReq, opts ...grpc.CallOption) (*None, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) EmailValidation(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/proto.UserService/EmailValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateJWT(ctx context.Context, in *ValidateJWTRequest, opts ...grpc.CallOption) (*ValidateJWTResponse, error) {
	out := new(ValidateJWTResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/ValidateJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFromID(ctx context.Context, in *GetUserFromIDRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUserFromID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserForForeignUID(ctx context.Context, in *GetForeignUserRequest, opts ...grpc.CallOption) (*GetForeignUserResponse, error) {
	out := new(GetForeignUserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUserForForeignUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserIDsForUsername(ctx context.Context, in *GetUserIDsForUsernameRequest, opts ...grpc.CallOption) (*GetUserIDsForUsernameResponse, error) {
	out := new(GetUserIDsForUsernameResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUserIDsForUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*BanUserResponse, error) {
	out := new(BanUserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/BanUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetUserRank(ctx context.Context, in *SetRankRequest, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/proto.UserService/SetUserRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/proto.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddAuditEvent(ctx context.Context, in *NewAuditEventRequest, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/proto.UserService/AddAuditEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAuditEvents(ctx context.Context, in *AuditEventsListRequest, opts ...grpc.CallOption) (*AuditListResponse, error) {
	out := new(AuditListResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetAuditEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFollowers(ctx context.Context, in *FollowerReq, opts ...grpc.CallOption) (*FollowerResp, error) {
	out := new(FollowerResp)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddFolllower(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/proto.UserService/AddFolllower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetProfile(ctx context.Context, in *ProfileReq, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/proto.UserService/SetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	EmailValidation(context.Context, *ValidationRequest) (*None, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ValidateJWT(context.Context, *ValidateJWTRequest) (*ValidateJWTResponse, error)
	GetUserFromID(context.Context, *GetUserFromIDRequest) (*UserResponse, error)
	GetUserForForeignUID(context.Context, *GetForeignUserRequest) (*GetForeignUserResponse, error)
	GetUserIDsForUsername(context.Context, *GetUserIDsForUsernameRequest) (*GetUserIDsForUsernameResponse, error)
	BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error)
	SetUserRank(context.Context, *SetRankRequest) (*None, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*None, error)
	AddAuditEvent(context.Context, *NewAuditEventRequest) (*None, error)
	GetAuditEvents(context.Context, *AuditEventsListRequest) (*AuditListResponse, error)
	GetFollowers(context.Context, *FollowerReq) (*FollowerResp, error)
	AddFolllower(context.Context, *AddFollowReq) (*None, error)
	SetProfile(context.Context, *ProfileReq) (*None, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) EmailValidation(context.Context, *ValidationRequest) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailValidation not implemented")
}
func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) ValidateJWT(context.Context, *ValidateJWTRequest) (*ValidateJWTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateJWT not implemented")
}
func (UnimplementedUserServiceServer) GetUserFromID(context.Context, *GetUserFromIDRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFromID not implemented")
}
func (UnimplementedUserServiceServer) GetUserForForeignUID(context.Context, *GetForeignUserRequest) (*GetForeignUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserForForeignUID not implemented")
}
func (UnimplementedUserServiceServer) GetUserIDsForUsername(context.Context, *GetUserIDsForUsernameRequest) (*GetUserIDsForUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIDsForUsername not implemented")
}
func (UnimplementedUserServiceServer) BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanUser not implemented")
}
func (UnimplementedUserServiceServer) SetUserRank(context.Context, *SetRankRequest) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserRank not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) AddAuditEvent(context.Context, *NewAuditEventRequest) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuditEvent not implemented")
}
func (UnimplementedUserServiceServer) GetAuditEvents(context.Context, *AuditEventsListRequest) (*AuditListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditEvents not implemented")
}
func (UnimplementedUserServiceServer) GetFollowers(context.Context, *FollowerReq) (*FollowerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedUserServiceServer) AddFolllower(context.Context, *AddFollowReq) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFolllower not implemented")
}
func (UnimplementedUserServiceServer) SetProfile(context.Context, *ProfileReq) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProfile not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_EmailValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EmailValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/EmailValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EmailValidation(ctx, req.(*ValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ValidateJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateJWTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ValidateJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ValidateJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ValidateJWT(ctx, req.(*ValidateJWTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFromID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFromIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFromID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUserFromID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFromID(ctx, req.(*GetUserFromIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserForForeignUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForeignUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserForForeignUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUserForForeignUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserForForeignUID(ctx, req.(*GetForeignUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserIDsForUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIDsForUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserIDsForUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUserIDsForUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserIDsForUsername(ctx, req.(*GetUserIDsForUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BanUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BanUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/BanUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BanUser(ctx, req.(*BanUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetUserRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetUserRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SetUserRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetUserRank(ctx, req.(*SetRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddAuditEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAuditEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAuditEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/AddAuditEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAuditEvent(ctx, req.(*NewAuditEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAuditEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditEventsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAuditEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetAuditEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAuditEvents(ctx, req.(*AuditEventsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetFollowers(ctx, req.(*FollowerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddFolllower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddFolllower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/AddFolllower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddFolllower(ctx, req.(*AddFollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetProfile(ctx, req.(*ProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmailValidation",
			Handler:    _UserService_EmailValidation_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "ValidateJWT",
			Handler:    _UserService_ValidateJWT_Handler,
		},
		{
			MethodName: "GetUserFromID",
			Handler:    _UserService_GetUserFromID_Handler,
		},
		{
			MethodName: "GetUserForForeignUID",
			Handler:    _UserService_GetUserForForeignUID_Handler,
		},
		{
			MethodName: "GetUserIDsForUsername",
			Handler:    _UserService_GetUserIDsForUsername_Handler,
		},
		{
			MethodName: "BanUser",
			Handler:    _UserService_BanUser_Handler,
		},
		{
			MethodName: "SetUserRank",
			Handler:    _UserService_SetUserRank_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "AddAuditEvent",
			Handler:    _UserService_AddAuditEvent_Handler,
		},
		{
			MethodName: "GetAuditEvents",
			Handler:    _UserService_GetAuditEvents_Handler,
		},
		{
			MethodName: "GetFollowers",
			Handler:    _UserService_GetFollowers_Handler,
		},
		{
			MethodName: "AddFolllower",
			Handler:    _UserService_AddFolllower_Handler,
		},
		{
			MethodName: "SetProfile",
			Handler:    _UserService_SetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice.proto",
}
