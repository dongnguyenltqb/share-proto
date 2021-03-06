// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/user.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	message "share-proto/proto-gen/message"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("rpc/user.proto", fileDescriptor_27d7a9c2ccec3127) }

var fileDescriptor_27d7a9c2ccec3127 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2a, 0x48, 0xd6,
	0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96,
	0x12, 0xca, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x45, 0x92, 0x30, 0x3a, 0xcf, 0xc4, 0xc5, 0x12,
	0x5a, 0x9c, 0x5a, 0x24, 0x14, 0xcd, 0x25, 0xe1, 0x9e, 0x5a, 0x02, 0x62, 0xba, 0x24, 0x96, 0x24,
	0x3a, 0x55, 0x3a, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x87, 0xe4, 0x67, 0xa7, 0xe6, 0x09, 0x89, 0xe8,
	0x41, 0x75, 0xea, 0x21, 0x89, 0x4a, 0xa9, 0xc3, 0x45, 0x71, 0x69, 0x0c, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0x51, 0x62, 0x10, 0x2a, 0xe4, 0x92, 0x0e, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0xc5, 0x6e, 0xbe,
	0x0a, 0xdc, 0x24, 0x3c, 0xaa, 0xa4, 0xb4, 0x88, 0x51, 0x05, 0xb7, 0xd2, 0x92, 0x8b, 0x13, 0xa4,
	0xc0, 0x27, 0x3f, 0x3d, 0x13, 0xd9, 0x03, 0xce, 0x45, 0xa9, 0x29, 0xa9, 0x79, 0x25, 0x99, 0x89,
	0x39, 0xc5, 0x52, 0x08, 0x51, 0xb0, 0x2a, 0xb8, 0x56, 0x1b, 0x2e, 0x1e, 0x90, 0xd6, 0xa0, 0xd4,
	0xf4, 0xcc, 0xe2, 0x92, 0xd4, 0x22, 0x21, 0x41, 0x84, 0xc5, 0x50, 0x2b, 0xa5, 0xc4, 0xe1, 0x42,
	0x30, 0x55, 0x30, 0xdd, 0x4e, 0xd2, 0x51, 0x92, 0xc5, 0x19, 0x89, 0x45, 0xa9, 0xba, 0xe0, 0x00,
	0xd6, 0x07, 0x93, 0xba, 0xe9, 0xa9, 0x79, 0xfa, 0x45, 0x05, 0xc9, 0x49, 0x6c, 0x60, 0xae, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x3b, 0x9a, 0xf8, 0xa0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserDataByAccessToken(ctx context.Context, in *message.AccessToken, opts ...grpc.CallOption) (*message.GetUserDataByAccessTokenResult, error)
	UpdateUserDataByAccessToken(ctx context.Context, in *message.UpdateUserDataByAccessToken, opts ...grpc.CallOption) (*message.UpdateUserDataByAccessTokenResult, error)
	UserLogin(ctx context.Context, in *message.Credentials, opts ...grpc.CallOption) (*message.LoginResult, error)
	UserRegister(ctx context.Context, in *message.UserData, opts ...grpc.CallOption) (*message.RegisterResult, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserDataByAccessToken(ctx context.Context, in *message.AccessToken, opts ...grpc.CallOption) (*message.GetUserDataByAccessTokenResult, error) {
	out := new(message.GetUserDataByAccessTokenResult)
	err := c.cc.Invoke(ctx, "/rpc.User/GetUserDataByAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserDataByAccessToken(ctx context.Context, in *message.UpdateUserDataByAccessToken, opts ...grpc.CallOption) (*message.UpdateUserDataByAccessTokenResult, error) {
	out := new(message.UpdateUserDataByAccessTokenResult)
	err := c.cc.Invoke(ctx, "/rpc.User/UpdateUserDataByAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogin(ctx context.Context, in *message.Credentials, opts ...grpc.CallOption) (*message.LoginResult, error) {
	out := new(message.LoginResult)
	err := c.cc.Invoke(ctx, "/rpc.User/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserRegister(ctx context.Context, in *message.UserData, opts ...grpc.CallOption) (*message.RegisterResult, error) {
	out := new(message.RegisterResult)
	err := c.cc.Invoke(ctx, "/rpc.User/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserDataByAccessToken(context.Context, *message.AccessToken) (*message.GetUserDataByAccessTokenResult, error)
	UpdateUserDataByAccessToken(context.Context, *message.UpdateUserDataByAccessToken) (*message.UpdateUserDataByAccessTokenResult, error)
	UserLogin(context.Context, *message.Credentials) (*message.LoginResult, error)
	UserRegister(context.Context, *message.UserData) (*message.RegisterResult, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserDataByAccessToken(ctx context.Context, req *message.AccessToken) (*message.GetUserDataByAccessTokenResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDataByAccessToken not implemented")
}
func (*UnimplementedUserServer) UpdateUserDataByAccessToken(ctx context.Context, req *message.UpdateUserDataByAccessToken) (*message.UpdateUserDataByAccessTokenResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDataByAccessToken not implemented")
}
func (*UnimplementedUserServer) UserLogin(ctx context.Context, req *message.Credentials) (*message.LoginResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserServer) UserRegister(ctx context.Context, req *message.UserData) (*message.RegisterResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserDataByAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AccessToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserDataByAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/GetUserDataByAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserDataByAccessToken(ctx, req.(*message.AccessToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserDataByAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.UpdateUserDataByAccessToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserDataByAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UpdateUserDataByAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserDataByAccessToken(ctx, req.(*message.UpdateUserDataByAccessToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*message.Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRegister(ctx, req.(*message.UserData))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserDataByAccessToken",
			Handler:    _User_GetUserDataByAccessToken_Handler,
		},
		{
			MethodName: "UpdateUserDataByAccessToken",
			Handler:    _User_UpdateUserDataByAccessToken_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _User_UserRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/user.proto",
}
