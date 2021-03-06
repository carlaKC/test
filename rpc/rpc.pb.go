// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ExampleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleRequest) Reset()         { *m = ExampleRequest{} }
func (m *ExampleRequest) String() string { return proto.CompactTextString(m) }
func (*ExampleRequest) ProtoMessage()    {}
func (*ExampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *ExampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleRequest.Unmarshal(m, b)
}
func (m *ExampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleRequest.Marshal(b, m, deterministic)
}
func (m *ExampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleRequest.Merge(m, src)
}
func (m *ExampleRequest) XXX_Size() int {
	return xxx_messageInfo_ExampleRequest.Size(m)
}
func (m *ExampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleRequest proto.InternalMessageInfo

type ExampleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleResponse) Reset()         { *m = ExampleResponse{} }
func (m *ExampleResponse) String() string { return proto.CompactTextString(m) }
func (*ExampleResponse) ProtoMessage()    {}
func (*ExampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *ExampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleResponse.Unmarshal(m, b)
}
func (m *ExampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleResponse.Marshal(b, m, deterministic)
}
func (m *ExampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleResponse.Merge(m, src)
}
func (m *ExampleResponse) XXX_Size() int {
	return xxx_messageInfo_ExampleResponse.Size(m)
}
func (m *ExampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExampleRequest)(nil), "rpc.ExampleRequest")
	proto.RegisterType((*ExampleResponse)(nil), "rpc.ExampleResponse")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x28, 0x51, 0x12, 0xe0, 0xe2, 0x73, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x0d,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x12, 0xe4, 0xe2, 0x87, 0x8b, 0x14, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x1a, 0xa5, 0x72, 0x71, 0xc1, 0x84, 0x02, 0x9c, 0x85, 0xc2, 0x11, 0xbc, 0x82, 0x64,
	0x21, 0x61, 0x3d, 0x90, 0x7d, 0xa8, 0x66, 0x48, 0x89, 0xa0, 0x0a, 0x42, 0x8c, 0x51, 0x92, 0x6b,
	0xba, 0xfc, 0x64, 0x32, 0x93, 0x84, 0x90, 0x98, 0x7e, 0x99, 0xa1, 0x7e, 0x2a, 0x44, 0x12, 0x46,
	0x17, 0x15, 0x24, 0x27, 0xb1, 0x81, 0x9d, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x30, 0xac,
	0xbc, 0x77, 0xc2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleRPCClient is the client API for ExampleRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleRPCClient interface {
	ExampleRpc(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error)
}

type exampleRPCClient struct {
	cc *grpc.ClientConn
}

func NewExampleRPCClient(cc *grpc.ClientConn) ExampleRPCClient {
	return &exampleRPCClient{cc}
}

func (c *exampleRPCClient) ExampleRpc(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResponse, error) {
	out := new(ExampleResponse)
	err := c.cc.Invoke(ctx, "/rpc.ExampleRPC/ExampleRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleRPCServer is the server API for ExampleRPC service.
type ExampleRPCServer interface {
	ExampleRpc(context.Context, *ExampleRequest) (*ExampleResponse, error)
}

// UnimplementedExampleRPCServer can be embedded to have forward compatible implementations.
type UnimplementedExampleRPCServer struct {
}

func (*UnimplementedExampleRPCServer) ExampleRpc(ctx context.Context, req *ExampleRequest) (*ExampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleRpc not implemented")
}

func RegisterExampleRPCServer(s *grpc.Server, srv ExampleRPCServer) {
	s.RegisterService(&_ExampleRPC_serviceDesc, srv)
}

func _ExampleRPC_ExampleRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleRPCServer).ExampleRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ExampleRPC/ExampleRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleRPCServer).ExampleRpc(ctx, req.(*ExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ExampleRPC",
	HandlerType: (*ExampleRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExampleRpc",
			Handler:    _ExampleRPC_ExampleRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
