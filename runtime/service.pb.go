// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime/service.proto

package runtime // import "github.com/socketfunc/proto/runtime"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import packet "github.com/socketfunc/proto/packet"

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

type StreamSend_Type int32

const (
	StreamSend_SEND      StreamSend_Type = 0
	StreamSend_REPLY     StreamSend_Type = 1
	StreamSend_BROADCAST StreamSend_Type = 2
)

var StreamSend_Type_name = map[int32]string{
	0: "SEND",
	1: "REPLY",
	2: "BROADCAST",
}
var StreamSend_Type_value = map[string]int32{
	"SEND":      0,
	"REPLY":     1,
	"BROADCAST": 2,
}

func (x StreamSend_Type) String() string {
	return proto.EnumName(StreamSend_Type_name, int32(x))
}
func (StreamSend_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_40afd5f82bcda706, []int{1, 0}
}

type StreamRequest struct {
	Packet               *packet.Packet `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40afd5f82bcda706, []int{0}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetPacket() *packet.Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

type StreamSend struct {
	Type                 StreamSend_Type `protobuf:"varint,1,opt,name=type,proto3,enum=runtime.StreamSend_Type" json:"type,omitempty"`
	Packet               *packet.Packet  `protobuf:"bytes,2,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StreamSend) Reset()         { *m = StreamSend{} }
func (m *StreamSend) String() string { return proto.CompactTextString(m) }
func (*StreamSend) ProtoMessage()    {}
func (*StreamSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40afd5f82bcda706, []int{1}
}
func (m *StreamSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamSend.Unmarshal(m, b)
}
func (m *StreamSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamSend.Marshal(b, m, deterministic)
}
func (dst *StreamSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamSend.Merge(dst, src)
}
func (m *StreamSend) XXX_Size() int {
	return xxx_messageInfo_StreamSend.Size(m)
}
func (m *StreamSend) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamSend.DiscardUnknown(m)
}

var xxx_messageInfo_StreamSend proto.InternalMessageInfo

func (m *StreamSend) GetType() StreamSend_Type {
	if m != nil {
		return m.Type
	}
	return StreamSend_SEND
}

func (m *StreamSend) GetPacket() *packet.Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

type HealthzRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthzRequest) Reset()         { *m = HealthzRequest{} }
func (m *HealthzRequest) String() string { return proto.CompactTextString(m) }
func (*HealthzRequest) ProtoMessage()    {}
func (*HealthzRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40afd5f82bcda706, []int{2}
}
func (m *HealthzRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthzRequest.Unmarshal(m, b)
}
func (m *HealthzRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthzRequest.Marshal(b, m, deterministic)
}
func (dst *HealthzRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthzRequest.Merge(dst, src)
}
func (m *HealthzRequest) XXX_Size() int {
	return xxx_messageInfo_HealthzRequest.Size(m)
}
func (m *HealthzRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthzRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthzRequest proto.InternalMessageInfo

type HealthzResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthzResponse) Reset()         { *m = HealthzResponse{} }
func (m *HealthzResponse) String() string { return proto.CompactTextString(m) }
func (*HealthzResponse) ProtoMessage()    {}
func (*HealthzResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_40afd5f82bcda706, []int{3}
}
func (m *HealthzResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthzResponse.Unmarshal(m, b)
}
func (m *HealthzResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthzResponse.Marshal(b, m, deterministic)
}
func (dst *HealthzResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthzResponse.Merge(dst, src)
}
func (m *HealthzResponse) XXX_Size() int {
	return xxx_messageInfo_HealthzResponse.Size(m)
}
func (m *HealthzResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthzResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthzResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StreamRequest)(nil), "runtime.StreamRequest")
	proto.RegisterType((*StreamSend)(nil), "runtime.StreamSend")
	proto.RegisterType((*HealthzRequest)(nil), "runtime.HealthzRequest")
	proto.RegisterType((*HealthzResponse)(nil), "runtime.HealthzResponse")
	proto.RegisterEnum("runtime.StreamSend_Type", StreamSend_Type_name, StreamSend_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeServiceClient is the client API for RuntimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeServiceClient interface {
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (RuntimeService_StreamClient, error)
	Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error)
}

type runtimeServiceClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeServiceClient(cc *grpc.ClientConn) RuntimeServiceClient {
	return &runtimeServiceClient{cc}
}

func (c *runtimeServiceClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (RuntimeService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeService_serviceDesc.Streams[0], "/runtime.RuntimeService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RuntimeService_StreamClient interface {
	Recv() (*StreamSend, error)
	grpc.ClientStream
}

type runtimeServiceStreamClient struct {
	grpc.ClientStream
}

func (x *runtimeServiceStreamClient) Recv() (*StreamSend, error) {
	m := new(StreamSend)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeServiceClient) Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/runtime.RuntimeService/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeServiceServer is the server API for RuntimeService service.
type RuntimeServiceServer interface {
	Stream(*StreamRequest, RuntimeService_StreamServer) error
	Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error)
}

func RegisterRuntimeServiceServer(s *grpc.Server, srv RuntimeServiceServer) {
	s.RegisterService(&_RuntimeService_serviceDesc, srv)
}

func _RuntimeService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RuntimeServiceServer).Stream(m, &runtimeServiceStreamServer{stream})
}

type RuntimeService_StreamServer interface {
	Send(*StreamSend) error
	grpc.ServerStream
}

type runtimeServiceStreamServer struct {
	grpc.ServerStream
}

func (x *runtimeServiceStreamServer) Send(m *StreamSend) error {
	return x.ServerStream.SendMsg(m)
}

func _RuntimeService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.RuntimeService/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).Healthz(ctx, req.(*HealthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.RuntimeService",
	HandlerType: (*RuntimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _RuntimeService_Healthz_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _RuntimeService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "runtime/service.proto",
}

func init() { proto.RegisterFile("runtime/service.proto", fileDescriptor_service_40afd5f82bcda706) }

var fileDescriptor_service_40afd5f82bcda706 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0xcc, 0x96, 0xfc, 0xd2, 0x5f, 0x9f, 0x34, 0xc6, 0x2d, 0x6a, 0xc8, 0x49, 0x56, 0x14, 0x11,
	0xd9, 0x48, 0x3c, 0x88, 0x17, 0xa1, 0xb5, 0x05, 0x0f, 0xa2, 0x65, 0xd3, 0x8b, 0xde, 0xd2, 0xf8,
	0xb4, 0x41, 0xf3, 0xc7, 0x64, 0x23, 0xd4, 0xaf, 0x20, 0xf8, 0x99, 0xa5, 0xd9, 0xa5, 0xfe, 0xc5,
	0xd3, 0x63, 0xdf, 0xcc, 0xec, 0xcc, 0xec, 0xc2, 0x7a, 0x59, 0x67, 0x32, 0x49, 0xd1, 0xaf, 0xb0,
	0x7c, 0x4e, 0x62, 0xe4, 0x45, 0x99, 0xcb, 0x9c, 0xb6, 0xf5, 0xda, 0xeb, 0x15, 0x51, 0xfc, 0x80,
	0xd2, 0x57, 0x43, 0xa1, 0xec, 0x18, 0xba, 0xa1, 0x2c, 0x31, 0x4a, 0x05, 0x3e, 0xd5, 0x58, 0x49,
	0xba, 0x0b, 0x96, 0x22, 0xb8, 0x64, 0x8b, 0xec, 0xad, 0x04, 0x36, 0xd7, 0xfc, 0x71, 0x33, 0x84,
	0x46, 0xd9, 0x1b, 0x01, 0x50, 0xca, 0x10, 0xb3, 0x5b, 0x7a, 0x00, 0xa6, 0x9c, 0x17, 0xd8, 0x88,
	0xec, 0xc0, 0xe5, 0xda, 0x94, 0x7f, 0x50, 0xf8, 0x64, 0x5e, 0xa0, 0x68, 0x58, 0x9f, 0x4c, 0x5a,
	0x7f, 0x9a, 0xec, 0x83, 0xb9, 0x50, 0xd1, 0xff, 0x60, 0x86, 0xa3, 0xcb, 0xa1, 0x63, 0xd0, 0x0e,
	0xfc, 0x13, 0xa3, 0xf1, 0xc5, 0xb5, 0x43, 0x68, 0x17, 0x3a, 0x03, 0x71, 0xd5, 0x1f, 0x9e, 0xf5,
	0xc3, 0x89, 0xd3, 0x62, 0x0e, 0xd8, 0xe7, 0x18, 0x3d, 0xca, 0xd9, 0x8b, 0xae, 0xc2, 0xd6, 0x60,
	0x75, 0xb9, 0xa9, 0x8a, 0x3c, 0xab, 0x30, 0x78, 0x25, 0x60, 0x0b, 0x15, 0x2d, 0x54, 0xaf, 0x44,
	0x4f, 0xc0, 0x52, 0x21, 0xe9, 0xc6, 0xb7, 0xd4, 0xfa, 0x1e, 0xaf, 0xf7, 0x4b, 0x1b, 0x66, 0x1c,
	0x12, 0x7a, 0x0a, 0x6d, 0x6d, 0x40, 0x37, 0x97, 0x9c, 0xaf, 0x21, 0x3c, 0xf7, 0x27, 0xa0, 0xb2,
	0x30, 0x63, 0xb0, 0x73, 0xb3, 0x7d, 0x9f, 0xc8, 0x59, 0x3d, 0xe5, 0x71, 0x9e, 0xfa, 0x55, 0xbe,
	0xe8, 0x7c, 0x57, 0x67, 0xb1, 0xdf, 0xfc, 0x8d, 0xaf, 0x85, 0x53, 0xab, 0x39, 0x1e, 0xbd, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xab, 0x6a, 0xcb, 0xa9, 0xe1, 0x01, 0x00, 0x00,
}
