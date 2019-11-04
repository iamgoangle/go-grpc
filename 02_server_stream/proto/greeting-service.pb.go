// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeting-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Greeting struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b4aa03c685ff65, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Greeting) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type CreateRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b4aa03c685ff65, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type CreateResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b4aa03c685ff65, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "v1.Greeting")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
}

func init() { proto.RegisterFile("greeting-service.proto", fileDescriptor_b9b4aa03c685ff65) }

var fileDescriptor_b9b4aa03c685ff65 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3d, 0x8f, 0x82, 0x40,
	0x10, 0x86, 0x0f, 0x0a, 0x02, 0x03, 0x77, 0xc9, 0x4d, 0x41, 0x0c, 0xb1, 0x30, 0x5b, 0xd1, 0x48,
	0x04, 0x1b, 0xad, 0x35, 0xd1, 0x1a, 0x7f, 0xc1, 0x6a, 0x46, 0x42, 0xc2, 0x97, 0xbb, 0x0b, 0xbf,
	0xdf, 0xb8, 0x2c, 0xa8, 0xe5, 0x3b, 0x4f, 0xe6, 0xfd, 0x80, 0xb0, 0x10, 0x44, 0xaa, 0x6c, 0x8a,
	0xb5, 0x24, 0x31, 0x94, 0x37, 0x4a, 0x3a, 0xd1, 0xaa, 0x16, 0xed, 0x21, 0x65, 0x47, 0x70, 0x4f,
	0x86, 0xe2, 0x12, 0xbc, 0x7b, 0x29, 0xa4, 0x6a, 0x78, 0x4d, 0x0b, 0x6b, 0x65, 0xc5, 0x5e, 0xfe,
	0x3e, 0x60, 0x04, 0x6e, 0xc5, 0x0d, 0xb4, 0x35, 0x9c, 0x35, 0xdb, 0xc3, 0xef, 0x41, 0x10, 0x57,
	0x94, 0xd3, 0xa3, 0x27, 0xa9, 0x30, 0x06, 0x77, 0x0a, 0xd5, 0x4e, 0x7e, 0x16, 0x24, 0x43, 0x9a,
	0x4c, 0x51, 0xf9, 0x4c, 0x59, 0x0c, 0x7f, 0xd3, 0xab, 0xec, 0xda, 0x46, 0x12, 0x86, 0xe0, 0x08,
	0x92, 0x7d, 0xa5, 0x4c, 0x07, 0xa3, 0xb2, 0x33, 0x04, 0xfa, 0xff, 0x32, 0x8e, 0xc0, 0x1d, 0xf8,
	0xa3, 0x56, 0x82, 0x78, 0x8d, 0xff, 0xaf, 0x80, 0xaf, 0x16, 0x11, 0x7e, 0x9e, 0x46, 0x77, 0xf6,
	0xb3, 0xb1, 0xae, 0x8e, 0xde, 0xbf, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x21, 0x19, 0x9d, 0xd8,
	0x19, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	GreetStream(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetStream(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/v1.GreetService/GreetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetStreamClient interface {
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type greetServiceGreetStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetStreamClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	GreetStream(*CreateRequest, GreetService_GreetStreamServer) error
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetStream(m, &greetServiceGreetStreamServer{stream})
}

type GreetService_GreetStreamServer interface {
	Send(*CreateResponse) error
	grpc.ServerStream
}

type greetServiceGreetStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetStreamServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetStream",
			Handler:       _GreetService_GreetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greeting-service.proto",
}
