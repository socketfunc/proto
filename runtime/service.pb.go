// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime/service.proto

package runtime // import "github.com/socketfunc/proto/runtime"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import store "github.com/socketfunc/proto/store"

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

type Cmd int32

const (
	Cmd_STREAM Cmd = 0
	Cmd_STORE  Cmd = 1
)

var Cmd_name = map[int32]string{
	0: "STREAM",
	1: "STORE",
}
var Cmd_value = map[string]int32{
	"STREAM": 0,
	"STORE":  1,
}

func (x Cmd) String() string {
	return proto.EnumName(Cmd_name, int32(x))
}
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{0}
}

type Store_Cmd int32

const (
	Store_Cmd_GET    Store_Cmd = 0
	Store_Cmd_PUT    Store_Cmd = 1
	Store_Cmd_MODIFY Store_Cmd = 2
	Store_Cmd_DEL    Store_Cmd = 3
)

var Store_Cmd_name = map[int32]string{
	0: "GET",
	1: "PUT",
	2: "MODIFY",
	3: "DEL",
}
var Store_Cmd_value = map[string]int32{
	"GET":    0,
	"PUT":    1,
	"MODIFY": 2,
	"DEL":    3,
}

func (x Store_Cmd) String() string {
	return proto.EnumName(Store_Cmd_name, int32(x))
}
func (Store_Cmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{1}
}

type Receive struct {
	Cmd                  Cmd            `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Cmd" json:"cmd,omitempty"`
	StreamRequest        *StreamRequest `protobuf:"bytes,2,opt,name=stream_request,json=streamRequest,proto3" json:"stream_request,omitempty"`
	StoreResponse        *StoreResponse `protobuf:"bytes,3,opt,name=store_response,json=storeResponse,proto3" json:"store_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Receive) Reset()         { *m = Receive{} }
func (m *Receive) String() string { return proto.CompactTextString(m) }
func (*Receive) ProtoMessage()    {}
func (*Receive) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{0}
}
func (m *Receive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receive.Unmarshal(m, b)
}
func (m *Receive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receive.Marshal(b, m, deterministic)
}
func (dst *Receive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receive.Merge(dst, src)
}
func (m *Receive) XXX_Size() int {
	return xxx_messageInfo_Receive.Size(m)
}
func (m *Receive) XXX_DiscardUnknown() {
	xxx_messageInfo_Receive.DiscardUnknown(m)
}

var xxx_messageInfo_Receive proto.InternalMessageInfo

func (m *Receive) GetCmd() Cmd {
	if m != nil {
		return m.Cmd
	}
	return Cmd_STREAM
}

func (m *Receive) GetStreamRequest() *StreamRequest {
	if m != nil {
		return m.StreamRequest
	}
	return nil
}

func (m *Receive) GetStoreResponse() *StoreResponse {
	if m != nil {
		return m.StoreResponse
	}
	return nil
}

type Send struct {
	Cmd                  Cmd           `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Cmd" json:"cmd,omitempty"`
	StreamSend           *StreamSend   `protobuf:"bytes,2,opt,name=stream_send,json=streamSend,proto3" json:"stream_send,omitempty"`
	StoreRequest         *StoreRequest `protobuf:"bytes,3,opt,name=store_request,json=storeRequest,proto3" json:"store_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Send) Reset()         { *m = Send{} }
func (m *Send) String() string { return proto.CompactTextString(m) }
func (*Send) ProtoMessage()    {}
func (*Send) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{1}
}
func (m *Send) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send.Unmarshal(m, b)
}
func (m *Send) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send.Marshal(b, m, deterministic)
}
func (dst *Send) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send.Merge(dst, src)
}
func (m *Send) XXX_Size() int {
	return xxx_messageInfo_Send.Size(m)
}
func (m *Send) XXX_DiscardUnknown() {
	xxx_messageInfo_Send.DiscardUnknown(m)
}

var xxx_messageInfo_Send proto.InternalMessageInfo

func (m *Send) GetCmd() Cmd {
	if m != nil {
		return m.Cmd
	}
	return Cmd_STREAM
}

func (m *Send) GetStreamSend() *StreamSend {
	if m != nil {
		return m.StreamSend
	}
	return nil
}

func (m *Send) GetStoreRequest() *StoreRequest {
	if m != nil {
		return m.StoreRequest
	}
	return nil
}

type StreamRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{2}
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

func (m *StreamRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *StreamRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *StreamRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type StreamSend struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamSend) Reset()         { *m = StreamSend{} }
func (m *StreamSend) String() string { return proto.CompactTextString(m) }
func (*StreamSend) ProtoMessage()    {}
func (*StreamSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{3}
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

func (m *StreamSend) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *StreamSend) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *StreamSend) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type StoreRequest struct {
	Cmd                  Store_Cmd       `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Store_Cmd" json:"cmd,omitempty"`
	Key                  string          `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Entity               *store.Entity   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	Filters              []*store.Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
	Updates              []*store.Update `protobuf:"bytes,5,rep,name=updates,proto3" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{4}
}
func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (dst *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(dst, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetCmd() Store_Cmd {
	if m != nil {
		return m.Cmd
	}
	return Store_Cmd_GET
}

func (m *StoreRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreRequest) GetEntity() *store.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *StoreRequest) GetFilters() []*store.Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *StoreRequest) GetUpdates() []*store.Update {
	if m != nil {
		return m.Updates
	}
	return nil
}

type StoreResponse struct {
	Cmd                  Store_Cmd     `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.Store_Cmd" json:"cmd,omitempty"`
	Success              bool          `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Entity               *store.Entity `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_babd969823b94968, []int{5}
}
func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (dst *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(dst, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetCmd() Store_Cmd {
	if m != nil {
		return m.Cmd
	}
	return Store_Cmd_GET
}

func (m *StoreResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *StoreResponse) GetEntity() *store.Entity {
	if m != nil {
		return m.Entity
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
	return fileDescriptor_service_babd969823b94968, []int{6}
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
	return fileDescriptor_service_babd969823b94968, []int{7}
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
	proto.RegisterType((*Receive)(nil), "runtime.Receive")
	proto.RegisterType((*Send)(nil), "runtime.Send")
	proto.RegisterType((*StreamRequest)(nil), "runtime.StreamRequest")
	proto.RegisterType((*StreamSend)(nil), "runtime.StreamSend")
	proto.RegisterType((*StoreRequest)(nil), "runtime.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "runtime.StoreResponse")
	proto.RegisterType((*HealthzRequest)(nil), "runtime.HealthzRequest")
	proto.RegisterType((*HealthzResponse)(nil), "runtime.HealthzResponse")
	proto.RegisterEnum("runtime.Cmd", Cmd_name, Cmd_value)
	proto.RegisterEnum("runtime.Store_Cmd", Store_Cmd_name, Store_Cmd_value)
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
	Stream(ctx context.Context, opts ...grpc.CallOption) (RuntimeService_StreamClient, error)
	Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error)
}

type runtimeServiceClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeServiceClient(cc *grpc.ClientConn) RuntimeServiceClient {
	return &runtimeServiceClient{cc}
}

func (c *runtimeServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (RuntimeService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeService_serviceDesc.Streams[0], "/runtime.RuntimeService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeServiceStreamClient{stream}
	return x, nil
}

type RuntimeService_StreamClient interface {
	Send(*Receive) error
	Recv() (*Send, error)
	grpc.ClientStream
}

type runtimeServiceStreamClient struct {
	grpc.ClientStream
}

func (x *runtimeServiceStreamClient) Send(m *Receive) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeServiceStreamClient) Recv() (*Send, error) {
	m := new(Send)
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
	Stream(RuntimeService_StreamServer) error
	Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error)
}

func RegisterRuntimeServiceServer(s *grpc.Server, srv RuntimeServiceServer) {
	s.RegisterService(&_RuntimeService_serviceDesc, srv)
}

func _RuntimeService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeServiceServer).Stream(&runtimeServiceStreamServer{stream})
}

type RuntimeService_StreamServer interface {
	Send(*Send) error
	Recv() (*Receive, error)
	grpc.ServerStream
}

type runtimeServiceStreamServer struct {
	grpc.ServerStream
}

func (x *runtimeServiceStreamServer) Send(m *Send) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeServiceStreamServer) Recv() (*Receive, error) {
	m := new(Receive)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
			ClientStreams: true,
		},
	},
	Metadata: "runtime/service.proto",
}

func init() { proto.RegisterFile("runtime/service.proto", fileDescriptor_service_babd969823b94968) }

var fileDescriptor_service_babd969823b94968 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0xd6, 0x4d, 0x4c, 0xa6, 0x49, 0x70, 0x17, 0x0a, 0x56, 0x84, 0x50, 0x64, 0xa8, 0x88,
	0x7a, 0xb0, 0x51, 0xe0, 0x84, 0x04, 0x12, 0xb4, 0x29, 0x20, 0x51, 0x15, 0x6d, 0x92, 0x03, 0x5c,
	0x22, 0xd7, 0x9e, 0x52, 0xab, 0xb1, 0x1d, 0xbc, 0xeb, 0x48, 0xe1, 0xc6, 0x5f, 0xf0, 0x07, 0x7c,
	0x06, 0xbf, 0x86, 0xbc, 0xbb, 0x76, 0x9d, 0xf6, 0x40, 0x0f, 0xdc, 0xf6, 0xbd, 0x37, 0xb3, 0x7e,
	0x6f, 0x32, 0x1b, 0xd8, 0xcb, 0xf2, 0x44, 0x44, 0x31, 0x7a, 0x1c, 0xb3, 0x55, 0x14, 0xa0, 0xbb,
	0xcc, 0x52, 0x91, 0x52, 0x53, 0xd3, 0xfd, 0x5d, 0x2e, 0xd2, 0x0c, 0xbd, 0x95, 0xbf, 0xc8, 0xb5,
	0xd6, 0xa7, 0x8a, 0x3a, 0x8f, 0x16, 0x02, 0xb3, 0x4d, 0x2e, 0x5f, 0x86, 0xbe, 0xd0, 0x75, 0xce,
	0x6f, 0x02, 0x26, 0xc3, 0x00, 0xa3, 0x15, 0xd2, 0xc7, 0x60, 0x04, 0x71, 0x68, 0x93, 0x01, 0x19,
	0xf6, 0x46, 0x1d, 0x57, 0xdf, 0xee, 0x1e, 0xc6, 0x21, 0x2b, 0x04, 0xfa, 0x1a, 0x7a, 0x5c, 0x64,
	0xe8, 0xc7, 0xf3, 0x0c, 0xbf, 0xe7, 0xc8, 0x85, 0xbd, 0x35, 0x20, 0xc3, 0x9d, 0xd1, 0x83, 0xaa,
	0x74, 0x22, 0x65, 0xa6, 0x54, 0xd6, 0xe5, 0x75, 0xa8, 0xda, 0xd3, 0x0c, 0xe7, 0x19, 0xf2, 0x65,
	0x9a, 0x70, 0xb4, 0x8d, 0x1b, 0xed, 0x69, 0x86, 0x4c, 0xab, 0x45, 0x7b, 0x0d, 0x3a, 0xbf, 0x08,
	0x6c, 0x4f, 0x30, 0x09, 0xff, 0x69, 0xf3, 0x25, 0xec, 0x68, 0x9b, 0x1c, 0x93, 0x50, 0x7b, 0xbc,
	0x77, 0xcd, 0x63, 0x71, 0x13, 0x03, 0x5e, 0x9d, 0xe9, 0x2b, 0xe8, 0x96, 0xee, 0x54, 0x36, 0x65,
	0x6e, 0xef, 0xba, 0x39, 0x15, 0xad, 0xc3, 0x6b, 0xc8, 0x99, 0x41, 0x77, 0x23, 0x39, 0xbd, 0x0f,
	0x4d, 0x91, 0x2e, 0xa3, 0x40, 0x9a, 0x6c, 0x33, 0x05, 0x0a, 0x16, 0x57, 0x98, 0xa8, 0xb1, 0xb5,
	0x99, 0x02, 0xd4, 0x06, 0x73, 0xe9, 0xaf, 0x17, 0xa9, 0x1f, 0xca, 0x4f, 0x76, 0x58, 0x09, 0x1d,
	0x06, 0x70, 0x65, 0xf6, 0x3f, 0xdd, 0xf9, 0x87, 0x40, 0xa7, 0x9e, 0x84, 0x3e, 0xad, 0x4f, 0x93,
	0x6e, 0xa6, 0x9d, 0x57, 0x33, 0xb5, 0xc0, 0xb8, 0xc4, 0xb5, 0xfe, 0x48, 0x71, 0xa4, 0xfb, 0xd0,
	0xc2, 0x44, 0x44, 0x62, 0xad, 0x07, 0xd5, 0x75, 0xe5, 0x48, 0xdc, 0xb1, 0x24, 0x99, 0x16, 0xe9,
	0x33, 0x30, 0xd5, 0x0e, 0x72, 0x7b, 0x7b, 0x60, 0xd4, 0xea, 0x8e, 0x25, 0xcb, 0x4a, 0xb5, 0x28,
	0x54, 0x8b, 0xc9, 0xed, 0xe6, 0x46, 0xe1, 0x4c, 0xb2, 0xac, 0x54, 0x9d, 0x55, 0x31, 0xec, 0xda,
	0x62, 0xdc, 0x32, 0x81, 0x0d, 0x26, 0xcf, 0x83, 0x00, 0x39, 0x97, 0x29, 0xee, 0xb0, 0x12, 0xde,
	0x32, 0x89, 0x63, 0x41, 0xef, 0x03, 0xfa, 0x0b, 0x71, 0xf1, 0xa3, 0xfc, 0xd9, 0x77, 0xe1, 0x6e,
	0xc5, 0x28, 0x2f, 0x07, 0x8f, 0xc0, 0x38, 0x8c, 0x43, 0x0a, 0xd0, 0x9a, 0x4c, 0xd9, 0xf8, 0xed,
	0x89, 0xd5, 0xa0, 0x6d, 0x68, 0x4e, 0xa6, 0xa7, 0x6c, 0x6c, 0x91, 0x83, 0x11, 0xb4, 0x2b, 0x57,
	0xd4, 0x04, 0xe3, 0xfd, 0x78, 0x6a, 0x35, 0x8a, 0xc3, 0xe7, 0xd9, 0xd4, 0x22, 0x45, 0xd7, 0xc9,
	0xe9, 0xd1, 0xc7, 0xe3, 0x2f, 0xd6, 0x56, 0x41, 0x1e, 0x8d, 0x3f, 0x59, 0xc6, 0xe8, 0x27, 0x81,
	0x1e, 0x53, 0x91, 0x26, 0xea, 0xf5, 0x53, 0x0f, 0x5a, 0x6a, 0x2f, 0xa8, 0x55, 0xa5, 0xd5, 0x6f,
	0xb8, 0xdf, 0xbd, 0xca, 0x8f, 0x49, 0xe8, 0x34, 0x86, 0xe4, 0x39, 0xa1, 0x6f, 0xc0, 0xd4, 0x46,
	0xe9, 0xc3, 0x4a, 0xdf, 0x0c, 0xd3, 0xb7, 0x6f, 0x0a, 0xfa, 0xe1, 0x35, 0xde, 0xed, 0x7f, 0x7d,
	0xf2, 0x2d, 0x12, 0x17, 0xf9, 0x99, 0x1b, 0xa4, 0xb1, 0xc7, 0xd3, 0xe0, 0x12, 0xc5, 0x79, 0x9e,
	0x04, 0x9e, 0xfc, 0x0f, 0xf1, 0x74, 0xe3, 0x59, 0x4b, 0xc2, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x23, 0x9b, 0x6d, 0xe4, 0xaf, 0x04, 0x00, 0x00,
}
