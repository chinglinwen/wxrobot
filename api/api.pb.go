// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
        api.proto

It has these top-level messages:
        TextRequest
        TextReply
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

// The request message containing the user's name.
type TextRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *TextRequest) Reset()                    { *m = TextRequest{} }
func (m *TextRequest) String() string            { return proto.CompactTextString(m) }
func (*TextRequest) ProtoMessage()               {}
func (*TextRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TextRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TextRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// The response message containing the greetings
type TextReply struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *TextReply) Reset()                    { *m = TextReply{} }
func (m *TextReply) String() string            { return proto.CompactTextString(m) }
func (*TextReply) ProtoMessage()               {}
func (*TextReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TextReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TextReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*TextRequest)(nil), "TextRequest")
	proto.RegisterType((*TextReply)(nil), "TextReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	GetText(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetText(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := grpc.Invoke(ctx, "/Api/GetText", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	GetText(context.Context, *TextRequest) (*TextReply, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_GetText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Api/GetText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetText(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetText",
			Handler:    _Api_GetText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe5, 0xe2, 0x0e, 0x49, 0xad, 0x28, 0x09, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x4c, 0x10, 0x31, 0x10,
	0x5b, 0xc9, 0x94, 0x8b, 0x13, 0xa2, 0xad, 0x20, 0xa7, 0x52, 0x48, 0x84, 0x8b, 0x35, 0xb5, 0xa8,
	0x28, 0xbf, 0x08, 0xaa, 0x0b, 0xc2, 0x01, 0x69, 0x4b, 0x49, 0x2c, 0x49, 0x84, 0x69, 0x03, 0xb1,
	0x8d, 0x74, 0xb8, 0x98, 0x1d, 0x0b, 0x32, 0x85, 0x54, 0xb9, 0xd8, 0xdd, 0x53, 0x4b, 0x40, 0x06,
	0x08, 0xf1, 0xe8, 0x21, 0x59, 0x2f, 0xc5, 0xa5, 0x07, 0x37, 0x55, 0x89, 0x21, 0x89, 0x0d, 0xec,
	0x44, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xf3, 0x29, 0xf7, 0xaf, 0x00, 0x00, 0x00,
}
