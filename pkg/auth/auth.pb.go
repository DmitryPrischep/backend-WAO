// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/auth/auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Token struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf4896e73d518a2, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UserData struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Agent                string   `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf4896e73d518a2, []int{1}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *UserData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserData) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

type Nothing struct {
	Null                 bool     `protobuf:"varint,1,opt,name=null,proto3" json:"null,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf4896e73d518a2, []int{2}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func (m *Nothing) GetNull() bool {
	if m != nil {
		return m.Null
	}
	return false
}

func init() {
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*UserData)(nil), "auth.UserData")
	proto.RegisterType((*Nothing)(nil), "auth.Nothing")
}

func init() { proto.RegisterFile("pkg/auth/auth.proto", fileDescriptor_adf4896e73d518a2) }

var fileDescriptor_adf4896e73d518a2 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xb7, 0xba, 0xad, 0x75, 0x16, 0x3d, 0x8c, 0x1e, 0x4a, 0x61, 0x41, 0x72, 0x50, 0x4f,
	0x2b, 0xe8, 0x13, 0xc8, 0xee, 0xd9, 0x43, 0xd1, 0x07, 0x88, 0x3a, 0x24, 0x4b, 0x43, 0x52, 0xd2,
	0x89, 0xbe, 0x81, 0xcf, 0x2d, 0x9d, 0x58, 0xd1, 0xbd, 0x84, 0xf9, 0x92, 0xff, 0x0f, 0x1f, 0x03,
	0x17, 0x43, 0x6f, 0xee, 0x74, 0x62, 0x2b, 0xc7, 0x66, 0x88, 0x81, 0x03, 0x2e, 0xa7, 0x59, 0xad,
	0xa1, 0x7c, 0x0e, 0x3d, 0x79, 0xbc, 0x84, 0xf2, 0x43, 0xbb, 0x44, 0x4d, 0x71, 0x55, 0xdc, 0x9e,
	0x76, 0x19, 0x54, 0x07, 0xf5, 0xcb, 0x48, 0x71, 0xa7, 0x59, 0x4f, 0x09, 0x17, 0xcc, 0xde, 0xcf,
	0x09, 0x01, 0x6c, 0xa1, 0x1e, 0xf4, 0x38, 0x7e, 0x86, 0xf8, 0xde, 0x1c, 0xc9, 0xc3, 0x2f, 0x4f,
	0x0d, 0x6d, 0xc8, 0x73, 0x73, 0x9c, 0x1b, 0x02, 0x6a, 0x0d, 0x27, 0x4f, 0x81, 0xed, 0xde, 0x1b,
	0x44, 0x58, 0xfa, 0xe4, 0x9c, 0xfc, 0x58, 0x77, 0x32, 0xdf, 0x7f, 0x15, 0xb0, 0x7a, 0x4c, 0x6c,
	0xb7, 0x96, 0xde, 0x7a, 0x8a, 0x78, 0x03, 0xd5, 0x36, 0x92, 0x66, 0xc2, 0xf3, 0x8d, 0xe8, 0xcf,
	0x42, 0xed, 0x2a, 0xb3, 0xf8, 0xab, 0x05, 0x5e, 0x43, 0x29, 0x1d, 0xfc, 0x7b, 0xdf, 0x1e, 0x94,
	0x24, 0x57, 0xed, 0xc8, 0x11, 0xd3, 0xff, 0xe0, 0x59, 0x86, 0x1f, 0x35, 0xb5, 0x78, 0xad, 0x64,
	0x4f, 0x0f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xcc, 0x81, 0x77, 0x3e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthCheckerClient interface {
	Create(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*Token, error)
	Check(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserData, error)
	Delete(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Nothing, error)
}

type authCheckerClient struct {
	cc *grpc.ClientConn
}

func NewAuthCheckerClient(cc *grpc.ClientConn) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) Create(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth.AuthChecker/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) Check(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/auth.AuthChecker/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) Delete(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/auth.AuthChecker/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
type AuthCheckerServer interface {
	Create(context.Context, *UserData) (*Token, error)
	Check(context.Context, *Token) (*UserData, error)
	Delete(context.Context, *Token) (*Nothing, error)
}

// UnimplementedAuthCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedAuthCheckerServer struct {
}

func (*UnimplementedAuthCheckerServer) Create(ctx context.Context, req *UserData) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAuthCheckerServer) Check(ctx context.Context, req *Token) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedAuthCheckerServer) Delete(ctx context.Context, req *Token) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAuthCheckerServer(s *grpc.Server, srv AuthCheckerServer) {
	s.RegisterService(&_AuthChecker_serviceDesc, srv)
}

func _AuthChecker_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthChecker/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Create(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthChecker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Check(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthChecker/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Delete(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthChecker_Create_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _AuthChecker_Check_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthChecker_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/auth/auth.proto",
}
