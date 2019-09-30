// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test_protos

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

type Outer struct {
	OuterValue           *Inner   `protobuf:"bytes,1,opt,name=outer_value,json=outerValue,proto3" json:"outer_value,omitempty"`
	OuterNum             int64    `protobuf:"varint,2,opt,name=outer_num,json=outerNum,proto3" json:"outer_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Outer) Reset()         { *m = Outer{} }
func (m *Outer) String() string { return proto.CompactTextString(m) }
func (*Outer) ProtoMessage()    {}
func (*Outer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Outer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outer.Unmarshal(m, b)
}
func (m *Outer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outer.Marshal(b, m, deterministic)
}
func (m *Outer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outer.Merge(m, src)
}
func (m *Outer) XXX_Size() int {
	return xxx_messageInfo_Outer.Size(m)
}
func (m *Outer) XXX_DiscardUnknown() {
	xxx_messageInfo_Outer.DiscardUnknown(m)
}

var xxx_messageInfo_Outer proto.InternalMessageInfo

func (m *Outer) GetOuterValue() *Inner {
	if m != nil {
		return m.OuterValue
	}
	return nil
}

func (m *Outer) GetOuterNum() int64 {
	if m != nil {
		return m.OuterNum
	}
	return 0
}

type Inner struct {
	InnerValue           string   `protobuf:"bytes,3,opt,name=inner_value,json=innerValue,proto3" json:"inner_value,omitempty"`
	InnerNum             int64    `protobuf:"varint,4,opt,name=inner_num,json=innerNum,proto3" json:"inner_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inner) Reset()         { *m = Inner{} }
func (m *Inner) String() string { return proto.CompactTextString(m) }
func (*Inner) ProtoMessage()    {}
func (*Inner) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Inner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inner.Unmarshal(m, b)
}
func (m *Inner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inner.Marshal(b, m, deterministic)
}
func (m *Inner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inner.Merge(m, src)
}
func (m *Inner) XXX_Size() int {
	return xxx_messageInfo_Inner.Size(m)
}
func (m *Inner) XXX_DiscardUnknown() {
	xxx_messageInfo_Inner.DiscardUnknown(m)
}

var xxx_messageInfo_Inner proto.InternalMessageInfo

func (m *Inner) GetInnerValue() string {
	if m != nil {
		return m.InnerValue
	}
	return ""
}

func (m *Inner) GetInnerNum() int64 {
	if m != nil {
		return m.InnerNum
	}
	return 0
}

type OuterWithExtra struct {
	OuterValue           *Inner   `protobuf:"bytes,1,opt,name=outer_value,json=outerValue,proto3" json:"outer_value,omitempty"`
	OuterNum             int64    `protobuf:"varint,2,opt,name=outer_num,json=outerNum,proto3" json:"outer_num,omitempty"`
	ExtraField           string   `protobuf:"bytes,3,opt,name=extra_field,json=extraField,proto3" json:"extra_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OuterWithExtra) Reset()         { *m = OuterWithExtra{} }
func (m *OuterWithExtra) String() string { return proto.CompactTextString(m) }
func (*OuterWithExtra) ProtoMessage()    {}
func (*OuterWithExtra) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *OuterWithExtra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OuterWithExtra.Unmarshal(m, b)
}
func (m *OuterWithExtra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OuterWithExtra.Marshal(b, m, deterministic)
}
func (m *OuterWithExtra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OuterWithExtra.Merge(m, src)
}
func (m *OuterWithExtra) XXX_Size() int {
	return xxx_messageInfo_OuterWithExtra.Size(m)
}
func (m *OuterWithExtra) XXX_DiscardUnknown() {
	xxx_messageInfo_OuterWithExtra.DiscardUnknown(m)
}

var xxx_messageInfo_OuterWithExtra proto.InternalMessageInfo

func (m *OuterWithExtra) GetOuterValue() *Inner {
	if m != nil {
		return m.OuterValue
	}
	return nil
}

func (m *OuterWithExtra) GetOuterNum() int64 {
	if m != nil {
		return m.OuterNum
	}
	return 0
}

func (m *OuterWithExtra) GetExtraField() string {
	if m != nil {
		return m.ExtraField
	}
	return ""
}

func init() {
	proto.RegisterType((*Outer)(nil), "test_protos.Outer")
	proto.RegisterType((*Inner)(nil), "test_protos.Inner")
	proto.RegisterType((*OuterWithExtra)(nil), "test_protos.OuterWithExtra")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x8e, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0xc6, 0x5d, 0x63, 0xc4, 0xcc, 0x82, 0xc5, 0x16, 0x12, 0x4c, 0xe1, 0x71, 0xd5, 0x55, 0x87,
	0x24, 0x6f, 0xa0, 0x44, 0x48, 0xa1, 0xc2, 0xf9, 0x0f, 0xab, 0x63, 0xd5, 0x31, 0x2e, 0xdc, 0xed,
	0xe9, 0xee, 0x6c, 0xd0, 0xda, 0xd7, 0xf0, 0x61, 0x65, 0xe6, 0x54, 0x14, 0x2c, 0x2c, 0xec, 0xbe,
	0xf9, 0xf1, 0xcd, 0x8f, 0x0f, 0x80, 0x30, 0x52, 0xf9, 0x18, 0x3a, 0xea, 0x8c, 0xe6, 0x5c, 0x4b,
	0x8e, 0xf9, 0x35, 0x0c, 0x4f, 0x13, 0x61, 0x30, 0x33, 0xd0, 0x1d, 0x87, 0x7a, 0x65, 0x9b, 0x84,
	0x63, 0x95, 0xa9, 0x42, 0x4f, 0x4d, 0xf9, 0xad, 0x5b, 0x2e, 0xbc, 0xc7, 0x50, 0x81, 0xd4, 0x2e,
	0xb9, 0x65, 0x26, 0x30, 0xea, 0x9f, 0x7c, 0x6a, 0xc7, 0xeb, 0x99, 0x2a, 0x06, 0xd5, 0x96, 0x80,
	0x93, 0xd4, 0xe6, 0x73, 0x18, 0xca, 0x87, 0xd9, 0x03, 0xed, 0x38, 0x7c, 0xa8, 0x07, 0x99, 0x2a,
	0x46, 0x15, 0x08, 0xfa, 0xd2, 0xf4, 0x05, 0xd6, 0x6c, 0xf4, 0x1a, 0x01, 0xac, 0x79, 0x55, 0xb0,
	0x2d, 0x13, 0xaf, 0x1c, 0x3d, 0xcc, 0x9f, 0x29, 0xd8, 0xff, 0xdf, 0xca, 0x13, 0x91, 0xd5, 0xf5,
	0xbd, 0xc3, 0xe6, 0xee, 0x73, 0xa2, 0xa0, 0x23, 0x26, 0xd3, 0x37, 0x05, 0xfa, 0x1c, 0x23, 0x9d,
	0x61, 0x58, 0xb9, 0x5b, 0x34, 0x07, 0xb0, 0xc3, 0xe7, 0x85, 0xb7, 0xe1, 0xe5, 0xb0, 0x71, 0xe8,
	0xa9, 0xc2, 0xa7, 0x84, 0x91, 0xcc, 0xcf, 0x1d, 0xb2, 0x7c, 0xf7, 0x17, 0x96, 0xaf, 0x99, 0x05,
	0x4c, 0x44, 0x49, 0x01, 0x6d, 0xeb, 0xfc, 0x92, 0xdd, 0x18, 0x8e, 0x31, 0x46, 0xbb, 0xc4, 0xf8,
	0x77, 0xd1, 0xbe, 0xba, 0xd9, 0x14, 0x32, 0x7b, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x49, 0x25, 0x23,
	0x8c, 0xe8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	TestUnaryClientRequest(ctx context.Context, in *Outer, opts ...grpc.CallOption) (*Outer, error)
	TestStreamingServerMessages(ctx context.Context, in *Outer, opts ...grpc.CallOption) (TestService_TestStreamingServerMessagesClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestUnaryClientRequest(ctx context.Context, in *Outer, opts ...grpc.CallOption) (*Outer, error) {
	out := new(Outer)
	err := c.cc.Invoke(ctx, "/test_protos.TestService/TestUnaryClientRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestStreamingServerMessages(ctx context.Context, in *Outer, opts ...grpc.CallOption) (TestService_TestStreamingServerMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/test_protos.TestService/TestStreamingServerMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestStreamingServerMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_TestStreamingServerMessagesClient interface {
	Recv() (*Outer, error)
	grpc.ClientStream
}

type testServiceTestStreamingServerMessagesClient struct {
	grpc.ClientStream
}

func (x *testServiceTestStreamingServerMessagesClient) Recv() (*Outer, error) {
	m := new(Outer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	TestUnaryClientRequest(context.Context, *Outer) (*Outer, error)
	TestStreamingServerMessages(*Outer, TestService_TestStreamingServerMessagesServer) error
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) TestUnaryClientRequest(ctx context.Context, req *Outer) (*Outer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestUnaryClientRequest not implemented")
}
func (*UnimplementedTestServiceServer) TestStreamingServerMessages(req *Outer, srv TestService_TestStreamingServerMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method TestStreamingServerMessages not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_TestUnaryClientRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Outer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestUnaryClientRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test_protos.TestService/TestUnaryClientRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestUnaryClientRequest(ctx, req.(*Outer))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestStreamingServerMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Outer)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestStreamingServerMessages(m, &testServiceTestStreamingServerMessagesServer{stream})
}

type TestService_TestStreamingServerMessagesServer interface {
	Send(*Outer) error
	grpc.ServerStream
}

type testServiceTestStreamingServerMessagesServer struct {
	grpc.ServerStream
}

func (x *testServiceTestStreamingServerMessagesServer) Send(m *Outer) error {
	return x.ServerStream.SendMsg(m)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test_protos.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestUnaryClientRequest",
			Handler:    _TestService_TestUnaryClientRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestStreamingServerMessages",
			Handler:       _TestService_TestStreamingServerMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}
