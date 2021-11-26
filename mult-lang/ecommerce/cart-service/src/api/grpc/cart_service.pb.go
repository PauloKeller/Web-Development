// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart_service.proto

package cart_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Cart struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string               `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Itens                []string             `protobuf:"bytes,3,rep,name=itens,proto3" json:"itens,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{0}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cart) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Cart) GetItens() []string {
	if m != nil {
		return m.Itens
	}
	return nil
}

func (m *Cart) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Cart) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateCartRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCartRequest) Reset()         { *m = CreateCartRequest{} }
func (m *CreateCartRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCartRequest) ProtoMessage()    {}
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{1}
}

func (m *CreateCartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCartRequest.Unmarshal(m, b)
}
func (m *CreateCartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCartRequest.Marshal(b, m, deterministic)
}
func (m *CreateCartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCartRequest.Merge(m, src)
}
func (m *CreateCartRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCartRequest.Size(m)
}
func (m *CreateCartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCartRequest proto.InternalMessageInfo

func (m *CreateCartRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*Cart)(nil), "cart_grpc.Cart")
	proto.RegisterType((*CreateCartRequest)(nil), "cart_grpc.CreateCartRequest")
}

func init() {
	proto.RegisterFile("cart_service.proto", fileDescriptor_c9a99120c5507bc1)
}

var fileDescriptor_c9a99120c5507bc1 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0xdd, 0xdc, 0x1f, 0xc9, 0x1c, 0x28, 0x2e, 0x82, 0x21, 0x08, 0x86, 0x54, 0x29, 0x64,
	0x0f, 0xce, 0xca, 0xc2, 0xe2, 0xb8, 0x4a, 0xb0, 0x8a, 0xf6, 0x61, 0x2f, 0x3b, 0x86, 0x05, 0xcf,
	0xc4, 0xd9, 0x89, 0x5f, 0xce, 0x2f, 0x27, 0xd9, 0xbd, 0x68, 0xc0, 0xc2, 0x72, 0x66, 0xde, 0x6f,
	0xde, 0x9b, 0x01, 0x59, 0x6b, 0xe2, 0xca, 0x21, 0x7d, 0xda, 0x1a, 0x55, 0x47, 0x2d, 0xb7, 0x32,
	0xf6, 0xbd, 0x86, 0xba, 0x3a, 0xbd, 0x69, 0xda, 0xb6, 0x79, 0xc3, 0xb5, 0x1f, 0xec, 0xfb, 0xd7,
	0x35, 0xdb, 0x03, 0x3a, 0xd6, 0x87, 0x2e, 0x68, 0xf3, 0x2f, 0x01, 0xf3, 0x9d, 0x26, 0x96, 0x67,
	0x10, 0x59, 0x93, 0x88, 0x4c, 0x14, 0x71, 0x19, 0x59, 0x23, 0xaf, 0xe0, 0xb4, 0x77, 0x48, 0x95,
	0x35, 0x49, 0xe4, 0x9b, 0xcb, 0xa1, 0x7c, 0x34, 0xf2, 0x12, 0x16, 0x96, 0xf1, 0xdd, 0x25, 0xb3,
	0x6c, 0x56, 0xc4, 0x65, 0x28, 0xe4, 0x3d, 0x40, 0x4d, 0xa8, 0x19, 0x4d, 0xa5, 0x39, 0x99, 0x67,
	0xa2, 0x58, 0x6d, 0x52, 0x15, 0xdc, 0xd5, 0xe8, 0xae, 0x5e, 0x46, 0xf7, 0x32, 0x3e, 0xaa, 0xb7,
	0x3c, 0xa0, 0x7d, 0x67, 0x46, 0x74, 0xf1, 0x3f, 0x7a, 0x54, 0x6f, 0x39, 0xbf, 0x85, 0x8b, 0x9d,
	0xdf, 0x33, 0x9c, 0x50, 0xe2, 0x47, 0x8f, 0x8e, 0xa7, 0xc9, 0xc5, 0x34, 0xf9, 0xe6, 0x09, 0x56,
	0x83, 0xee, 0x39, 0x3c, 0x4b, 0x3e, 0x00, 0xfc, 0xc2, 0xf2, 0x5a, 0xfd, 0x7c, 0x4d, 0xfd, 0xd9,
	0x99, 0x9e, 0x4f, 0xa7, 0x9a, 0x38, 0x3f, 0xd9, 0x2f, 0x7d, 0xb4, 0xbb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0x15, 0x2f, 0xda, 0x82, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*Cart, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*Cart, error) {
	out := new(Cart)
	err := c.cc.Invoke(ctx, "/cart_grpc.CartService/CreateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	CreateCart(context.Context, *CreateCartRequest) (*Cart, error)
}

// UnimplementedCartServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (*UnimplementedCartServiceServer) CreateCart(ctx context.Context, req *CreateCartRequest) (*Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart_grpc.CartService/CreateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateCart(ctx, req.(*CreateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cart_grpc.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _CartService_CreateCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_service.proto",
}