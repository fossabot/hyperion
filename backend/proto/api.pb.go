// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	proto/api.proto

It has these top-level messages:
	Ping
*/
package api

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

type Ping struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ping) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "Ping")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	GetPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error) {
	out := new(Ping)
	err := grpc.Invoke(ctx, "/API/GetPing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	GetPing(context.Context, *Ping) (*Ping, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetPing(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPing",
			Handler:    _API_GetPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

func init() { proto.RegisterFile("proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x03, 0xb3, 0x94, 0x14, 0xb8, 0x58, 0x02, 0x32, 0xf3, 0xd2,
	0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x60, 0x5c, 0x23, 0x05, 0x2e, 0x66, 0xc7, 0x00, 0x4f, 0x21, 0x49, 0x2e, 0x76, 0xf7,
	0xd4, 0x12, 0xb0, 0x5a, 0x56, 0x3d, 0x10, 0x25, 0x05, 0xa1, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x46,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x85, 0x6d, 0x87, 0x96, 0x5d, 0x00, 0x00, 0x00,
}
