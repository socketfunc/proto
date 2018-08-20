// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway/service.proto

package gateway // import "github.com/socketfunc/proto/gateway"

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

type PublishRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                string   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{0}
}
func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (dst *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(dst, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *PublishRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PublishResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{1}
}
func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (dst *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(dst, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

func (m *PublishResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SubscribeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{2}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (dst *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(dst, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type SubscribeSend struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeSend) Reset()         { *m = SubscribeSend{} }
func (m *SubscribeSend) String() string { return proto.CompactTextString(m) }
func (*SubscribeSend) ProtoMessage()    {}
func (*SubscribeSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{3}
}
func (m *SubscribeSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeSend.Unmarshal(m, b)
}
func (m *SubscribeSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeSend.Marshal(b, m, deterministic)
}
func (dst *SubscribeSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeSend.Merge(dst, src)
}
func (m *SubscribeSend) XXX_Size() int {
	return xxx_messageInfo_SubscribeSend.Size(m)
}
func (m *SubscribeSend) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeSend.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeSend proto.InternalMessageInfo

func (m *SubscribeSend) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UnsubscribeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsubscribeRequest) Reset()         { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()    {}
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{4}
}
func (m *UnsubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeRequest.Unmarshal(m, b)
}
func (m *UnsubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeRequest.Marshal(b, m, deterministic)
}
func (dst *UnsubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeRequest.Merge(dst, src)
}
func (m *UnsubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeRequest.Size(m)
}
func (m *UnsubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeRequest proto.InternalMessageInfo

func (m *UnsubscribeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UnsubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type UnsubscribeResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsubscribeResponse) Reset()         { *m = UnsubscribeResponse{} }
func (m *UnsubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeResponse) ProtoMessage()    {}
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{5}
}
func (m *UnsubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeResponse.Unmarshal(m, b)
}
func (m *UnsubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeResponse.Marshal(b, m, deterministic)
}
func (dst *UnsubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeResponse.Merge(dst, src)
}
func (m *UnsubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeResponse.Size(m)
}
func (m *UnsubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeResponse proto.InternalMessageInfo

func (m *UnsubscribeResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type DisconnectRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{6}
}
func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (dst *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(dst, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DisconnectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectResponse) Reset()         { *m = DisconnectResponse{} }
func (m *DisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*DisconnectResponse) ProtoMessage()    {}
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{7}
}
func (m *DisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectResponse.Unmarshal(m, b)
}
func (m *DisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectResponse.Marshal(b, m, deterministic)
}
func (dst *DisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectResponse.Merge(dst, src)
}
func (m *DisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_DisconnectResponse.Size(m)
}
func (m *DisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectResponse proto.InternalMessageInfo

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{8}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_95560bd8c6e6b1ae, []int{9}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PublishRequest)(nil), "gateway.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "gateway.PublishResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "gateway.SubscribeRequest")
	proto.RegisterType((*SubscribeSend)(nil), "gateway.SubscribeSend")
	proto.RegisterType((*UnsubscribeRequest)(nil), "gateway.UnsubscribeRequest")
	proto.RegisterType((*UnsubscribeResponse)(nil), "gateway.UnsubscribeResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "gateway.DisconnectRequest")
	proto.RegisterType((*DisconnectResponse)(nil), "gateway.DisconnectResponse")
	proto.RegisterType((*PingRequest)(nil), "gateway.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "gateway.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (GatewayService_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	Ping(ctx context.Context, opts ...grpc.CallOption) (GatewayService_PingClient, error)
}

type gatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayServiceClient(cc *grpc.ClientConn) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/gateway.GatewayService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (GatewayService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GatewayService_serviceDesc.Streams[0], "/gateway.GatewayService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GatewayService_SubscribeClient interface {
	Recv() (*SubscribeSend, error)
	grpc.ClientStream
}

type gatewayServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *gatewayServiceSubscribeClient) Recv() (*SubscribeSend, error) {
	m := new(SubscribeSend)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/gateway.GatewayService/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/gateway.GatewayService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Ping(ctx context.Context, opts ...grpc.CallOption) (GatewayService_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GatewayService_serviceDesc.Streams[1], "/gateway.GatewayService/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayServicePingClient{stream}
	return x, nil
}

type GatewayService_PingClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type gatewayServicePingClient struct {
	grpc.ClientStream
}

func (x *gatewayServicePingClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gatewayServicePingClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Subscribe(*SubscribeRequest, GatewayService_SubscribeServer) error
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	Ping(GatewayService_PingServer) error
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServiceServer).Subscribe(m, &gatewayServiceSubscribeServer{stream})
}

type GatewayService_SubscribeServer interface {
	Send(*SubscribeSend) error
	grpc.ServerStream
}

type gatewayServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *gatewayServiceSubscribeServer) Send(m *SubscribeSend) error {
	return x.ServerStream.SendMsg(m)
}

func _GatewayService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.GatewayService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GatewayServiceServer).Ping(&gatewayServicePingServer{stream})
}

type GatewayService_PingServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type gatewayServicePingServer struct {
	grpc.ServerStream
}

func (x *gatewayServicePingServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gatewayServicePingServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _GatewayService_Publish_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _GatewayService_Unsubscribe_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _GatewayService_Disconnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _GatewayService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Ping",
			Handler:       _GatewayService_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gateway/service.proto",
}

func init() { proto.RegisterFile("gateway/service.proto", fileDescriptor_service_95560bd8c6e6b1ae) }

var fileDescriptor_service_95560bd8c6e6b1ae = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x6f, 0xda, 0x30,
	0x18, 0x25, 0x19, 0x1b, 0xe3, 0x03, 0xb2, 0xcd, 0x83, 0x2d, 0x0b, 0x3b, 0x20, 0xa3, 0x49, 0x4c,
	0x93, 0x12, 0xb4, 0x5d, 0xb6, 0x1d, 0x7a, 0x40, 0x95, 0x90, 0x7a, 0x42, 0x41, 0xbd, 0xf4, 0x96,
	0x38, 0x06, 0xac, 0x52, 0x3b, 0xc5, 0x0e, 0x15, 0xff, 0x6f, 0xff, 0x90, 0x2a, 0x3f, 0x30, 0x3f,
	0x42, 0x5b, 0xb5, 0xc7, 0xf7, 0x7d, 0xef, 0xf9, 0xd9, 0xef, 0xc9, 0xd0, 0x99, 0x07, 0x8a, 0xde,
	0x05, 0x1b, 0x4f, 0xd2, 0xd5, 0x9a, 0x11, 0xea, 0xc6, 0x2b, 0xa1, 0x04, 0xaa, 0x15, 0x63, 0x3c,
	0x03, 0x6b, 0x92, 0x84, 0x4b, 0x26, 0x17, 0x3e, 0xbd, 0x4d, 0xa8, 0x54, 0xc8, 0x02, 0x93, 0x45,
	0xb6, 0xd1, 0x33, 0x06, 0x75, 0xdf, 0x64, 0x11, 0x6a, 0xc3, 0x5b, 0x25, 0x62, 0x46, 0x6c, 0x33,
	0x1b, 0xe5, 0x20, 0x9d, 0xd2, 0x35, 0xe5, 0xca, 0x7e, 0x93, 0x4f, 0x33, 0x80, 0x6c, 0xa8, 0xc5,
	0xc1, 0x66, 0x29, 0x82, 0xc8, 0xae, 0xf6, 0x8c, 0x41, 0xd3, 0xdf, 0x42, 0xfc, 0x0b, 0x3e, 0x68,
	0x1f, 0x19, 0x0b, 0x2e, 0x69, 0x4a, 0x96, 0x09, 0x21, 0x54, 0xca, 0xcc, 0xed, 0xbd, 0xbf, 0x85,
	0xf8, 0x2f, 0x7c, 0x9c, 0x26, 0xa1, 0x24, 0x2b, 0x16, 0xd2, 0x17, 0x5d, 0x0b, 0xff, 0x84, 0x96,
	0x56, 0x4e, 0x29, 0x8f, 0x9e, 0x30, 0xf9, 0x0f, 0xe8, 0x92, 0xcb, 0xd7, 0xd9, 0x78, 0xf0, 0xf9,
	0x40, 0xfb, 0xec, 0x8b, 0xfa, 0xf0, 0xe9, 0x9c, 0x49, 0x22, 0x38, 0xa7, 0x44, 0x3d, 0xe2, 0x85,
	0xdb, 0x80, 0xf6, 0x49, 0xf9, 0xa1, 0xb8, 0x05, 0x8d, 0x09, 0xe3, 0xf3, 0x42, 0x84, 0x2d, 0x68,
	0xe6, 0x30, 0x5f, 0xff, 0xbe, 0x37, 0xc1, 0x1a, 0xe7, 0x65, 0x4e, 0xf3, 0x8a, 0xd1, 0x19, 0xd4,
	0x8a, 0xac, 0xd1, 0x57, 0xb7, 0x28, 0xda, 0x3d, 0x6c, 0xd9, 0xb1, 0xcb, 0x8b, 0xc2, 0xaf, 0x82,
	0x46, 0x50, 0xd7, 0x21, 0xa2, 0x6f, 0x9a, 0x78, 0x5c, 0x89, 0xf3, 0xa5, 0xbc, 0x4a, 0x33, 0xc7,
	0x95, 0xa1, 0x81, 0x2e, 0xa0, 0xb1, 0x97, 0x10, 0xea, 0x6a, 0x6a, 0x39, 0x73, 0xe7, 0xfb, 0xe9,
	0xa5, 0xbe, 0xcf, 0x18, 0x60, 0x97, 0x0b, 0x72, 0x34, 0xbb, 0x94, 0xa8, 0xd3, 0x3d, 0xb9, 0xd3,
	0x07, 0xfd, 0x83, 0x6a, 0x9a, 0x1d, 0x6a, 0xef, 0x1e, 0xbf, 0x4b, 0xd6, 0xe9, 0x1c, 0x4d, 0xb7,
	0xb2, 0x81, 0x31, 0x34, 0x46, 0x3f, 0xae, 0xfa, 0x73, 0xa6, 0x16, 0x49, 0xe8, 0x12, 0x71, 0xe3,
	0x49, 0x41, 0xae, 0xa9, 0x9a, 0x25, 0x9c, 0x78, 0xd9, 0x7f, 0xf2, 0x0a, 0x65, 0xf8, 0x2e, 0x83,
	0x7f, 0x1e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x80, 0x90, 0x27, 0xf5, 0x77, 0x03, 0x00, 0x00,
}
