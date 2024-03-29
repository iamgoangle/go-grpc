// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package welcome

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

type Hello struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Hello) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type HelloWithDeadlineRequest struct {
	Hello                *Hello   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWithDeadlineRequest) Reset()         { *m = HelloWithDeadlineRequest{} }
func (m *HelloWithDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWithDeadlineRequest) ProtoMessage()    {}
func (*HelloWithDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloWithDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWithDeadlineRequest.Unmarshal(m, b)
}
func (m *HelloWithDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWithDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *HelloWithDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWithDeadlineRequest.Merge(m, src)
}
func (m *HelloWithDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWithDeadlineRequest.Size(m)
}
func (m *HelloWithDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWithDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWithDeadlineRequest proto.InternalMessageInfo

func (m *HelloWithDeadlineRequest) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

type HelloWithDealLineResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWithDealLineResponse) Reset()         { *m = HelloWithDealLineResponse{} }
func (m *HelloWithDealLineResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWithDealLineResponse) ProtoMessage()    {}
func (*HelloWithDealLineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *HelloWithDealLineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWithDealLineResponse.Unmarshal(m, b)
}
func (m *HelloWithDealLineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWithDealLineResponse.Marshal(b, m, deterministic)
}
func (m *HelloWithDealLineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWithDealLineResponse.Merge(m, src)
}
func (m *HelloWithDealLineResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWithDealLineResponse.Size(m)
}
func (m *HelloWithDealLineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWithDealLineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWithDealLineResponse proto.InternalMessageInfo

func (m *HelloWithDealLineResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "welcome.Hello")
	proto.RegisterType((*HelloWithDeadlineRequest)(nil), "welcome.HelloWithDeadlineRequest")
	proto.RegisterType((*HelloWithDealLineResponse)(nil), "welcome.HelloWithDealLineResponse")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4f, 0xcd, 0x49, 0xce, 0xcf, 0x4d,
	0x55, 0x72, 0xe6, 0x62, 0xf5, 0x00, 0x89, 0x0b, 0xc9, 0x72, 0x71, 0xa5, 0x65, 0x16, 0x15, 0x97,
	0xc4, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x82, 0x45, 0xfc,
	0x12, 0x73, 0x53, 0x85, 0xa4, 0xb9, 0x38, 0x73, 0x12, 0x61, 0xb2, 0x4c, 0x60, 0x59, 0x0e, 0x90,
	0x00, 0x48, 0x52, 0xc9, 0x81, 0x4b, 0x02, 0x6c, 0x48, 0x78, 0x66, 0x49, 0x86, 0x4b, 0x6a, 0x62,
	0x4a, 0x4e, 0x66, 0x5e, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x0a, 0x17, 0x2b,
	0xd8, 0x62, 0xb0, 0x91, 0xdc, 0x46, 0x7c, 0x7a, 0x50, 0x9b, 0xf5, 0xc0, 0x3a, 0x82, 0x20, 0x92,
	0x4a, 0xc6, 0x5c, 0x92, 0xc8, 0x26, 0xe4, 0xf8, 0x80, 0x4d, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x15, 0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x81, 0x3a, 0x0b, 0xca, 0x33, 0xca,
	0xe1, 0xe2, 0x01, 0x6b, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x8a, 0xe1, 0x12, 0xc4,
	0x70, 0x86, 0x90, 0x22, 0xaa, 0x85, 0x58, 0x9c, 0x28, 0xa5, 0x84, 0x55, 0x09, 0x8a, 0x1b, 0x94,
	0x18, 0x92, 0xd8, 0xc0, 0x21, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x50, 0x15, 0xa9,
	0x48, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	HelloWithDeadline(ctx context.Context, in *HelloWithDeadlineRequest, opts ...grpc.CallOption) (*HelloWithDealLineResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) HelloWithDeadline(ctx context.Context, in *HelloWithDeadlineRequest, opts ...grpc.CallOption) (*HelloWithDealLineResponse, error) {
	out := new(HelloWithDealLineResponse)
	err := c.cc.Invoke(ctx, "/welcome.HelloService/HelloWithDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	HelloWithDeadline(context.Context, *HelloWithDeadlineRequest) (*HelloWithDealLineResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_HelloWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).HelloWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/welcome.HelloService/HelloWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).HelloWithDeadline(ctx, req.(*HelloWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "welcome.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWithDeadline",
			Handler:    _HelloService_HelloWithDeadline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
