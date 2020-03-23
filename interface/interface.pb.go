// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interface.proto

package gke_grpc_service

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

type HealthCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

type HealthCheckResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

type HelloWorldRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{2}
}

func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (m *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(m, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloWorldResponse struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{3}
}

func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (m *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(m, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "GKEgRPCService.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "GKEgRPCService.HealthCheckResponse")
	proto.RegisterType((*HelloWorldRequest)(nil), "GKEgRPCService.HelloWorldRequest")
	proto.RegisterType((*HelloWorldResponse)(nil), "GKEgRPCService.HelloWorldResponse")
}

func init() { proto.RegisterFile("interface.proto", fileDescriptor_3ef53c9e620778f1) }

var fileDescriptor_3ef53c9e620778f1 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0x88, 0xe2, 0x08, 0x6a, 0xc7, 0x2a, 0x25, 0xf4, 0xa0, 0xeb, 0x41, 0xf1, 0x90,
	0x80, 0xde, 0x3c, 0x1a, 0xc4, 0x82, 0x20, 0x12, 0x0f, 0x82, 0x97, 0xb0, 0xc6, 0x31, 0x09, 0x5d,
	0x77, 0xd2, 0xdd, 0x55, 0xef, 0xbe, 0x82, 0x8f, 0xe6, 0x2b, 0xf8, 0x0e, 0x5e, 0xa5, 0x9b, 0x94,
	0xda, 0x16, 0x7a, 0xdb, 0x3f, 0xbf, 0xfd, 0xbe, 0xdf, 0xb0, 0xb0, 0x5d, 0x69, 0x47, 0xe6, 0x45,
	0xe6, 0x14, 0xd5, 0x86, 0x1d, 0xe3, 0xd6, 0xf5, 0xcd, 0x55, 0x91, 0xde, 0x25, 0xf7, 0x64, 0xde,
	0xab, 0x9c, 0xc2, 0x7e, 0xc1, 0x5c, 0x28, 0x8a, 0x65, 0x5d, 0xc5, 0x52, 0x6b, 0x76, 0xd2, 0x55,
	0xac, 0x6d, 0x43, 0x8b, 0x2e, 0xe0, 0x80, 0xa4, 0x72, 0x65, 0x52, 0x52, 0x3e, 0x4c, 0x69, 0xf4,
	0x46, 0xd6, 0x89, 0x3d, 0xd8, 0x9d, 0x39, 0xb5, 0x35, 0x6b, 0x4b, 0xe2, 0x18, 0x3a, 0x03, 0x52,
	0x8a, 0x1f, 0xd8, 0xa8, 0xe7, 0x96, 0x45, 0x84, 0xd5, 0x5b, 0xf9, 0x4a, 0xbd, 0xe0, 0x20, 0x38,
	0xd9, 0x48, 0xfd, 0x5a, 0x44, 0xe3, 0xd4, 0x29, 0xd8, 0x3c, 0xc7, 0x1e, 0xac, 0x27, 0xac, 0x1d,
	0x69, 0xd7, 0xc2, 0x93, 0xed, 0xd9, 0x6f, 0x00, 0x73, 0xda, 0x68, 0x60, 0xf3, 0x9f, 0x02, 0x8a,
	0x68, 0xf6, 0x3e, 0x5a, 0xb4, 0x0e, 0x8f, 0x96, 0x32, 0xed, 0x0c, 0xe1, 0xe7, 0xf7, 0xcf, 0xd7,
	0x4a, 0x17, 0x31, 0x36, 0x64, 0x5d, 0x5c, 0x7a, 0x24, 0xcb, 0x7d, 0xc9, 0x08, 0x60, 0xaa, 0x8d,
	0x87, 0x8b, 0x71, 0x73, 0xb3, 0x87, 0x62, 0x19, 0xd2, 0x16, 0xf6, 0x7d, 0xe1, 0xbe, 0xe8, 0x4c,
	0x0a, 0x95, 0xe2, 0xec, 0x63, 0x8c, 0x5c, 0x04, 0xa7, 0x97, 0xf8, 0xb8, 0x53, 0x0c, 0x29, 0x2b,
	0x4c, 0x9d, 0x67, 0xb6, 0x09, 0x79, 0x5a, 0xf3, 0x5f, 0x73, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0xfe, 0x23, 0xa6, 0xdb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GKEgRPCServiceClient is the client API for GKEgRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GKEgRPCServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type gKEgRPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewGKEgRPCServiceClient(cc *grpc.ClientConn) GKEgRPCServiceClient {
	return &gKEgRPCServiceClient{cc}
}

func (c *gKEgRPCServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/GKEgRPCService.GKEgRPCService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gKEgRPCServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/GKEgRPCService.GKEgRPCService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GKEgRPCServiceServer is the server API for GKEgRPCService service.
type GKEgRPCServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
}

// UnimplementedGKEgRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGKEgRPCServiceServer struct {
}

func (*UnimplementedGKEgRPCServiceServer) HealthCheck(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (*UnimplementedGKEgRPCServiceServer) HelloWorld(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}

func RegisterGKEgRPCServiceServer(s *grpc.Server, srv GKEgRPCServiceServer) {
	s.RegisterService(&_GKEgRPCService_serviceDesc, srv)
}

func _GKEgRPCService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GKEgRPCServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GKEgRPCService.GKEgRPCService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GKEgRPCServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GKEgRPCService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GKEgRPCServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GKEgRPCService.GKEgRPCService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GKEgRPCServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GKEgRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GKEgRPCService.GKEgRPCService",
	HandlerType: (*GKEgRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _GKEgRPCService_HealthCheck_Handler,
		},
		{
			MethodName: "HelloWorld",
			Handler:    _GKEgRPCService_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interface.proto",
}
