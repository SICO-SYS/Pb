// Code generated by protoc-gen-go.
// source: He.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	He.proto

It has these top-level messages:
	ResponseMsg
	OpenRequest
	APIKeypair
	APIToken
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResponseMsg struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *ResponseMsg) Reset()                    { *m = ResponseMsg{} }
func (m *ResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*ResponseMsg) ProtoMessage()               {}
func (*ResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResponseMsg) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type OpenRequest struct {
	Request string `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *OpenRequest) Reset()                    { *m = OpenRequest{} }
func (m *OpenRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()               {}
func (*OpenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OpenRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

type APIKeypair struct {
	Key    string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
}

func (m *APIKeypair) Reset()                    { *m = APIKeypair{} }
func (m *APIKeypair) String() string            { return proto.CompactTextString(m) }
func (*APIKeypair) ProtoMessage()               {}
func (*APIKeypair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *APIKeypair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *APIKeypair) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type APIToken struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *APIToken) Reset()                    { *m = APIToken{} }
func (m *APIToken) String() string            { return proto.CompactTextString(m) }
func (*APIToken) ProtoMessage()               {}
func (*APIToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *APIToken) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *APIToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ResponseMsg)(nil), "pb.ResponseMsg")
	proto.RegisterType((*OpenRequest)(nil), "pb.OpenRequest")
	proto.RegisterType((*APIKeypair)(nil), "pb.APIKeypair")
	proto.RegisterType((*APIToken)(nil), "pb.APIToken")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Open service

type OpenClient interface {
	RegUser(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*APIKeypair, error)
}

type openClient struct {
	cc *grpc.ClientConn
}

func NewOpenClient(cc *grpc.ClientConn) OpenClient {
	return &openClient{cc}
}

func (c *openClient) RegUser(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*APIKeypair, error) {
	out := new(APIKeypair)
	err := grpc.Invoke(ctx, "/pb.Open/RegUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Open service

type OpenServer interface {
	RegUser(context.Context, *OpenRequest) (*APIKeypair, error)
}

func RegisterOpenServer(s *grpc.Server, srv OpenServer) {
	s.RegisterService(&_Open_serviceDesc, srv)
}

func _Open_RegUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenServer).RegUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Open/RegUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenServer).RegUser(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Open_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Open",
	HandlerType: (*OpenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegUser",
			Handler:    _Open_RegUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "He.proto",
}

// Client API for Secret service

type SecretClient interface {
	APIAuth(ctx context.Context, in *APIKeypair, opts ...grpc.CallOption) (*ResponseMsg, error)
}

type secretClient struct {
	cc *grpc.ClientConn
}

func NewSecretClient(cc *grpc.ClientConn) SecretClient {
	return &secretClient{cc}
}

func (c *secretClient) APIAuth(ctx context.Context, in *APIKeypair, opts ...grpc.CallOption) (*ResponseMsg, error) {
	out := new(ResponseMsg)
	err := grpc.Invoke(ctx, "/pb.Secret/APIAuth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Secret service

type SecretServer interface {
	APIAuth(context.Context, *APIKeypair) (*ResponseMsg, error)
}

func RegisterSecretServer(s *grpc.Server, srv SecretServer) {
	s.RegisterService(&_Secret_serviceDesc, srv)
}

func _Secret_APIAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeypair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).APIAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Secret/APIAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).APIAuth(ctx, req.(*APIKeypair))
	}
	return interceptor(ctx, in, info, handler)
}

var _Secret_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Secret",
	HandlerType: (*SecretServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "APIAuth",
			Handler:    _Secret_APIAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "He.proto",
}

func init() { proto.RegisterFile("He.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd7, 0xb9, 0x75, 0xdb, 0x1b, 0xa8, 0x3c, 0x44, 0xca, 0x4e, 0x23, 0x17, 0x77, 0x90,
	0x1e, 0x3a, 0xe9, 0x3d, 0x37, 0x8b, 0x88, 0x25, 0xea, 0x1f, 0x60, 0xeb, 0xa3, 0x4a, 0xb1, 0x89,
	0x49, 0x7a, 0xe8, 0x7f, 0x2f, 0x49, 0x23, 0x16, 0xd9, 0xed, 0xfb, 0x92, 0xf7, 0xfb, 0xde, 0xc7,
	0x83, 0xf5, 0x3d, 0xa5, 0x4a, 0x4b, 0x2b, 0x71, 0xae, 0x2a, 0x76, 0x84, 0xad, 0x20, 0xa3, 0x64,
	0x67, 0xe8, 0xd1, 0x34, 0x88, 0xb0, 0xa8, 0xe5, 0x3b, 0x25, 0xd1, 0x3e, 0x3a, 0x2c, 0x85, 0xd7,
	0x78, 0x09, 0x67, 0x5f, 0xa6, 0x49, 0xe6, 0xfb, 0xe8, 0xb0, 0x11, 0x4e, 0xb2, 0x1b, 0xd8, 0x3e,
	0x29, 0xea, 0x04, 0x7d, 0xf7, 0x64, 0x2c, 0x26, 0xb0, 0xd2, 0xa3, 0xf4, 0xdc, 0x46, 0xfc, 0x5a,
	0x96, 0x03, 0xf0, 0xb2, 0x78, 0xa0, 0x41, 0xbd, 0x7d, 0x6a, 0x17, 0xd4, 0xd2, 0x10, 0x66, 0x9c,
	0xc4, 0x6b, 0x88, 0x0d, 0xd5, 0x9a, 0x6c, 0x48, 0x0f, 0x8e, 0x65, 0xb0, 0xe6, 0x65, 0xf1, 0x22,
	0x5b, 0xea, 0x4e, 0x50, 0x57, 0xb0, 0xb4, 0xee, 0x2b, 0x40, 0xa3, 0xc9, 0xee, 0x60, 0xe1, 0x4a,
	0xe1, 0x2d, 0xac, 0x04, 0x35, 0xaf, 0x86, 0x34, 0x5e, 0xa4, 0xaa, 0x4a, 0x27, 0x4d, 0x77, 0xe7,
	0xee, 0xe1, 0xaf, 0x11, 0x9b, 0x65, 0x39, 0xc4, 0xcf, 0x7e, 0xa7, 0xe3, 0x78, 0x59, 0xf0, 0xde,
	0x7e, 0xe0, 0xbf, 0xb1, 0x9d, 0xcf, 0x99, 0x9c, 0x89, 0xcd, 0xaa, 0xd8, 0x9f, 0xf0, 0xf8, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0x4c, 0xec, 0xa5, 0x4e, 0x01, 0x00, 0x00,
}
