// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users_service.proto

package candyland_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2683e726c16c873, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateUserRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2683e726c16c873, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateUserReply struct {
	WasCreated           bool     `protobuf:"varint,1,opt,name=wasCreated,proto3" json:"wasCreated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReply) Reset()         { *m = CreateUserReply{} }
func (m *CreateUserReply) String() string { return proto.CompactTextString(m) }
func (*CreateUserReply) ProtoMessage()    {}
func (*CreateUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2683e726c16c873, []int{2}
}

func (m *CreateUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReply.Unmarshal(m, b)
}
func (m *CreateUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReply.Marshal(b, m, deterministic)
}
func (m *CreateUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReply.Merge(m, src)
}
func (m *CreateUserReply) XXX_Size() int {
	return xxx_messageInfo_CreateUserReply.Size(m)
}
func (m *CreateUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReply proto.InternalMessageInfo

func (m *CreateUserReply) GetWasCreated() bool {
	if m != nil {
		return m.WasCreated
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "candyland.grpc.User")
	proto.RegisterType((*CreateUserRequest)(nil), "candyland.grpc.CreateUserRequest")
	proto.RegisterType((*CreateUserReply)(nil), "candyland.grpc.CreateUserReply")
}

func init() {
	proto.RegisterFile("users_service.proto", fileDescriptor_e2683e726c16c873)
}

var fileDescriptor_e2683e726c16c873 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0xb7, 0x2b, 0xed, 0x20, 0x2b, 0x46, 0x0f, 0xa5, 0xa2, 0xab, 0x7b, 0xf2, 0xd4,
	0x45, 0x7d, 0x04, 0xef, 0x22, 0x15, 0xcf, 0x6b, 0xb6, 0x9d, 0x2d, 0x81, 0xa4, 0x8d, 0x99, 0x54,
	0xe9, 0x51, 0xf0, 0xc1, 0xa5, 0x09, 0x8b, 0x55, 0xb1, 0xc7, 0xf9, 0xbf, 0x49, 0xf8, 0x92, 0x1f,
	0x4e, 0x3a, 0x42, 0x43, 0x1b, 0x42, 0xf3, 0x26, 0x4a, 0xcc, 0xb5, 0x69, 0x6d, 0xcb, 0x16, 0x25,
	0x6f, 0xaa, 0x5e, 0xf2, 0xa6, 0xca, 0x6b, 0xa3, 0xcb, 0x6c, 0x59, 0xb7, 0x6d, 0x2d, 0x71, 0xed,
	0xe8, 0xb6, 0xdb, 0xad, 0xad, 0x50, 0x48, 0x96, 0x2b, 0xed, 0x0f, 0xac, 0x3e, 0x03, 0x88, 0x9e,
	0x09, 0x0d, 0x5b, 0x40, 0x28, 0xaa, 0x34, 0xb8, 0x0c, 0xae, 0x93, 0x22, 0x14, 0x15, 0x3b, 0x07,
	0xd8, 0x09, 0x43, 0x76, 0xd3, 0x70, 0x85, 0x69, 0xe8, 0xf2, 0xc4, 0x25, 0x0f, 0x5c, 0x21, 0x3b,
	0x83, 0x44, 0xf2, 0x3d, 0x9d, 0x39, 0x1a, 0x0f, 0x81, 0x83, 0x19, 0xc4, 0x83, 0x9c, 0x63, 0x91,
	0x67, 0xfb, 0x99, 0x9d, 0xc2, 0x1c, 0x15, 0x17, 0x32, 0x9d, 0x3b, 0xe0, 0x87, 0xd5, 0x47, 0x00,
	0xc7, 0xf7, 0x06, 0xb9, 0xc5, 0x41, 0xa6, 0xc0, 0xd7, 0x0e, 0xc9, 0xfe, 0x72, 0x08, 0x26, 0x1d,
	0xc2, 0x09, 0x87, 0xd9, 0x7f, 0x0e, 0xd1, 0xd8, 0xe1, 0x06, 0x8e, 0xc6, 0x0a, 0x5a, 0xf6, 0xec,
	0x02, 0xe0, 0x9d, 0x93, 0x4f, 0xfd, 0xe7, 0xc4, 0xc5, 0x28, 0xb9, 0x7d, 0x81, 0xc3, 0x61, 0x99,
	0x9e, 0x7c, 0x09, 0xec, 0x11, 0xe0, 0xfb, 0x0a, 0x76, 0x95, 0xff, 0x6c, 0x23, 0xff, 0xf3, 0xc2,
	0x6c, 0x39, 0xb5, 0xa2, 0x65, 0xbf, 0x3d, 0x70, 0x35, 0xdd, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0xf6, 0xce, 0xab, 0xee, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, "/candyland.grpc.UsersService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
type UsersServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
}

// UnimplementedUsersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (*UnimplementedUsersServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUsersServiceServer(s *grpc.Server, srv UsersServiceServer) {
	s.RegisterService(&_UsersService_serviceDesc, srv)
}

func _UsersService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candyland.grpc.UsersService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "candyland.grpc.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UsersService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users_service.proto",
}
