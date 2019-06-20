// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverpb/server.proto

package serverpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9202d8f598083902, []int{0}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	// Full PMM version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Build timestamp.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// pmm-managed commit.
	PmmManagedCommit     string   `protobuf:"bytes,3,opt,name=pmm_managed_commit,json=pmmManagedCommit,proto3" json:"pmm_managed_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9202d8f598083902, []int{1}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *VersionResponse) GetPmmManagedCommit() string {
	if m != nil {
		return m.PmmManagedCommit
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "server.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "server.VersionResponse")
}

func init() { proto.RegisterFile("serverpb/server.proto", fileDescriptor_9202d8f598083902) }

var fileDescriptor_9202d8f598083902 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x4a, 0x03, 0x31,
	0x18, 0x64, 0x2b, 0xb4, 0x36, 0x45, 0x2d, 0x11, 0x75, 0x59, 0x04, 0x97, 0x16, 0xb1, 0x87, 0x76,
	0x83, 0xf5, 0xe2, 0xc5, 0x43, 0x15, 0x41, 0x04, 0x11, 0x56, 0x51, 0xf4, 0x52, 0xd2, 0xf6, 0xeb,
	0x12, 0x68, 0x7e, 0x4c, 0xd2, 0x7a, 0xf7, 0x11, 0xec, 0xa3, 0xf9, 0x0a, 0x3e, 0x88, 0x98, 0x6c,
	0x2a, 0xd2, 0xd3, 0xee, 0x37, 0xdf, 0x64, 0x32, 0x33, 0x41, 0x7b, 0x06, 0xf4, 0x02, 0xb4, 0x1a,
	0x11, 0xff, 0x93, 0x29, 0x2d, 0xad, 0xc4, 0x55, 0x3f, 0x25, 0x87, 0x85, 0x94, 0xc5, 0x0c, 0x08,
	0x55, 0x8c, 0x50, 0x21, 0xa4, 0xa5, 0x96, 0x49, 0x61, 0x3c, 0x2b, 0x39, 0x2a, 0xb7, 0x6e, 0x1a,
	0xcd, 0xa7, 0xc4, 0x32, 0x0e, 0xc6, 0x52, 0xae, 0x4a, 0x42, 0xd7, 0x7d, 0xc6, 0xbd, 0x02, 0x44,
	0xcf, 0xbc, 0xd3, 0xa2, 0x00, 0x4d, 0xa4, 0x72, 0x12, 0xeb, 0x72, 0xad, 0x26, 0xda, 0x7e, 0x02,
	0x6d, 0x98, 0x14, 0x39, 0xbc, 0xcd, 0xc1, 0xd8, 0xd6, 0x32, 0x42, 0x3b, 0x2b, 0xc8, 0x28, 0x29,
	0x0c, 0xe0, 0x18, 0xd5, 0x16, 0x1e, 0x8a, 0xa3, 0x34, 0xea, 0xd4, 0xf3, 0x30, 0xe2, 0x73, 0x54,
	0x5f, 0x19, 0x88, 0x2b, 0x69, 0xd4, 0x69, 0xf4, 0x93, 0xcc, 0x5b, 0xcc, 0x82, 0xc5, 0xec, 0x31,
	0x30, 0xf2, 0x3f, 0x32, 0xee, 0x22, 0xac, 0x38, 0x1f, 0x72, 0x2a, 0x68, 0x01, 0x93, 0xe1, 0x58,
	0x72, 0xce, 0x6c, 0xbc, 0xe1, 0xe4, 0x9b, 0x8a, 0xf3, 0x3b, 0xbf, 0xb8, 0x72, 0x78, 0xff, 0x05,
	0x55, 0x1f, 0x5c, 0x3d, 0xf8, 0x1e, 0xd5, 0x4a, 0x7b, 0x78, 0x3f, 0x2b, 0x0b, 0xfc, 0x1f, 0x21,
	0x39, 0x58, 0xc3, 0x7d, 0x8e, 0xd6, 0xee, 0xc7, 0xd7, 0xf7, 0xb2, 0xb2, 0x85, 0x1b, 0x64, 0x71,
	0x4a, 0xca, 0x08, 0x97, 0xcf, 0x9f, 0x83, 0xdb, 0xfc, 0x06, 0xd5, 0x26, 0x30, 0xa5, 0xf3, 0x99,
	0xc5, 0x17, 0x08, 0x0f, 0x44, 0x0a, 0x5a, 0x4b, 0x9d, 0xea, 0xf2, 0x64, 0x86, 0x4f, 0xd0, 0x71,
	0xd2, 0x6e, 0x93, 0x09, 0x4c, 0x99, 0x60, 0xbe, 0xcd, 0xf0, 0x8c, 0xd7, 0xbf, 0xd4, 0x70, 0xc7,
	0xeb, 0x66, 0x80, 0x47, 0x55, 0x57, 0xc0, 0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xda,
	0xce, 0x8a, 0xf0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	// Version returns PMM Server version.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/server.Server/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	// Version returns PMM Server version.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Server/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Server_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverpb/server.proto",
}