// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addition.proto

package bkp

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

type AddRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	AnotherNumber        int32    `protobuf:"varint,2,opt,name=anotherNumber,proto3" json:"anotherNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9e7dbd1dd38dde4, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *AddRequest) GetAnotherNumber() int32 {
	if m != nil {
		return m.AnotherNumber
	}
	return 0
}

type AddResponse struct {
	Sum                  int64    `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9e7dbd1dd38dde4, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "addition.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "addition.AddResponse")
}

func init() { proto.RegisterFile("addition.proto", fileDescriptor_a9e7dbd1dd38dde4) }

var fileDescriptor_a9e7dbd1dd38dde4 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4c, 0x49, 0xc9,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xbc,
	0xb8, 0xb8, 0x1c, 0x53, 0x52, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8,
	0xf2, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x21,
	0x15, 0x2e, 0xde, 0xc4, 0xbc, 0xfc, 0x92, 0x8c, 0xd4, 0x22, 0x3f, 0x88, 0x34, 0x13, 0x58, 0x1a,
	0x55, 0x50, 0x49, 0x9e, 0x8b, 0x1b, 0x6c, 0x56, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x00,
	0x17, 0x73, 0x71, 0x69, 0x2e, 0xd8, 0x24, 0xe6, 0x20, 0x10, 0xd3, 0xc8, 0x81, 0x8b, 0xc3, 0x11,
	0x6a, 0xb1, 0x90, 0x09, 0x17, 0xb3, 0x63, 0x4a, 0x8a, 0x90, 0x88, 0x1e, 0xdc, 0x69, 0x08, 0x77,
	0x48, 0x89, 0xa2, 0x89, 0x42, 0x4c, 0x54, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xdf, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x10, 0xc8, 0x37, 0xe9, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdditionClient is the client API for Addition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdditionClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type additionClient struct {
	cc grpc.ClientConnInterface
}

func NewAdditionClient(cc grpc.ClientConnInterface) AdditionClient {
	return &additionClient{cc}
}

func (c *additionClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/addition.Addition/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdditionServer is the server API for Addition service.
type AdditionServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
}

// UnimplementedAdditionServer can be embedded to have forward compatible implementations.
type UnimplementedAdditionServer struct {
}

func (*UnimplementedAdditionServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterAdditionServer(s *grpc.Server, srv AdditionServer) {
	s.RegisterService(&_Addition_serviceDesc, srv)
}

func _Addition_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addition.Addition/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Addition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addition.Addition",
	HandlerType: (*AdditionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Addition_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addition.proto",
}
