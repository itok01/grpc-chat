// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcchat.proto

package grpcchat

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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_67007a5b5af5d17f, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Message struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Register             bool     `protobuf:"varint,3,opt,name=register,proto3" json:"register,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_67007a5b5af5d17f, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetRegister() bool {
	if m != nil {
		return m.Register
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "grpcchat.User")
	proto.RegisterType((*Message)(nil), "grpcchat.Message")
}

func init() { proto.RegisterFile("grpcchat.proto", fileDescriptor_67007a5b5af5d17f) }

var fileDescriptor_67007a5b5af5d17f = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2f, 0x2a, 0x48,
	0x4e, 0xce, 0x48, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xa4,
	0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa5, 0x58, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4,
	0xf4, 0x54, 0x21, 0x25, 0x2e, 0x96, 0xd2, 0xe2, 0xd4, 0x22, 0xb0, 0x34, 0xb7, 0x11, 0x9f, 0x1e,
	0xdc, 0x3c, 0x90, 0xe6, 0x20, 0xb0, 0x1c, 0xc8, 0x88, 0x92, 0xd4, 0x8a, 0x12, 0x09, 0x26, 0x88,
	0x11, 0x20, 0xb6, 0x90, 0x14, 0x17, 0x47, 0x51, 0x6a, 0x7a, 0x66, 0x71, 0x49, 0x6a, 0x91, 0x04,
	0xb3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x9c, 0x6f, 0xe4, 0xc9, 0xc5, 0xe1, 0x5e, 0x54, 0x90, 0xec,
	0x9c, 0x91, 0x58, 0x22, 0x64, 0xcb, 0xc5, 0x1f, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x96, 0x5a, 0x04,
	0xb3, 0x52, 0x10, 0x61, 0x09, 0x54, 0x48, 0x0a, 0x53, 0x48, 0x89, 0x41, 0x83, 0xd1, 0x80, 0xd1,
	0x49, 0x8d, 0x4b, 0x38, 0x39, 0x3f, 0x57, 0x2f, 0xb3, 0x24, 0x3f, 0xdb, 0xc0, 0x10, 0xae, 0xc8,
	0x89, 0x17, 0x66, 0x7e, 0x00, 0xc8, 0xd7, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0xef, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x27, 0x2b, 0x82, 0x1e, 0x10, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcChatClient is the client API for GrpcChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcChatClient interface {
	TransferMessage(ctx context.Context, opts ...grpc.CallOption) (GrpcChat_TransferMessageClient, error)
}

type grpcChatClient struct {
	cc *grpc.ClientConn
}

func NewGrpcChatClient(cc *grpc.ClientConn) GrpcChatClient {
	return &grpcChatClient{cc}
}

func (c *grpcChatClient) TransferMessage(ctx context.Context, opts ...grpc.CallOption) (GrpcChat_TransferMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcChat_serviceDesc.Streams[0], "/grpcchat.GrpcChat/TransferMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcChatTransferMessageClient{stream}
	return x, nil
}

type GrpcChat_TransferMessageClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type grpcChatTransferMessageClient struct {
	grpc.ClientStream
}

func (x *grpcChatTransferMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *grpcChatTransferMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GrpcChatServer is the server API for GrpcChat service.
type GrpcChatServer interface {
	TransferMessage(GrpcChat_TransferMessageServer) error
}

// UnimplementedGrpcChatServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcChatServer struct {
}

func (*UnimplementedGrpcChatServer) TransferMessage(srv GrpcChat_TransferMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferMessage not implemented")
}

func RegisterGrpcChatServer(s *grpc.Server, srv GrpcChatServer) {
	s.RegisterService(&_GrpcChat_serviceDesc, srv)
}

func _GrpcChat_TransferMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GrpcChatServer).TransferMessage(&grpcChatTransferMessageServer{stream})
}

type GrpcChat_TransferMessageServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type grpcChatTransferMessageServer struct {
	grpc.ServerStream
}

func (x *grpcChatTransferMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *grpcChatTransferMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GrpcChat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcchat.GrpcChat",
	HandlerType: (*GrpcChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferMessage",
			Handler:       _GrpcChat_TransferMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcchat.proto",
}
