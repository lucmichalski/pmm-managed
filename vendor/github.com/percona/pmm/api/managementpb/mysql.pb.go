// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/mysql.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
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

type AddMySQLRequest struct {
	// Node identifier on which a service is been running. Required.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Node and Service access address (DNS name or IP). Required.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port. Required.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,5,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,7,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,8,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// MySQL username for scraping metrics.
	Username string `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	// MySQL password for scraping metrics.
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	// FIXME remove
	QanUsername string `protobuf:"bytes,11,opt,name=qan_username,json=qanUsername,proto3" json:"qan_username,omitempty"`
	// FIXME remove
	QanPassword string `protobuf:"bytes,12,opt,name=qan_password,json=qanPassword,proto3" json:"qan_password,omitempty"`
	// If true, adds mysqld_exporter for provided service.
	MysqldExporter bool `protobuf:"varint,13,opt,name=mysqld_exporter,json=mysqldExporter,proto3" json:"mysqld_exporter,omitempty"`
	// If true, adds qan-mysql-perfschema-agent for provided service.
	QanMysqlPerfschema bool `protobuf:"varint,14,opt,name=qan_mysql_perfschema,json=qanMysqlPerfschema,proto3" json:"qan_mysql_perfschema,omitempty"`
	// If true, adds qan-mysql-slowlog-agent for provided service.
	QanMysqlSlowlog      bool     `protobuf:"varint,15,opt,name=qan_mysql_slowlog,json=qanMysqlSlowlog,proto3" json:"qan_mysql_slowlog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMySQLRequest) Reset()         { *m = AddMySQLRequest{} }
func (m *AddMySQLRequest) String() string { return proto.CompactTextString(m) }
func (*AddMySQLRequest) ProtoMessage()    {}
func (*AddMySQLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab81470951176953, []int{0}
}

func (m *AddMySQLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLRequest.Unmarshal(m, b)
}
func (m *AddMySQLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLRequest.Marshal(b, m, deterministic)
}
func (m *AddMySQLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLRequest.Merge(m, src)
}
func (m *AddMySQLRequest) XXX_Size() int {
	return xxx_messageInfo_AddMySQLRequest.Size(m)
}
func (m *AddMySQLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLRequest proto.InternalMessageInfo

func (m *AddMySQLRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddMySQLRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AddMySQLRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddMySQLRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AddMySQLRequest) GetPmmAgentId() string {
	if m != nil {
		return m.PmmAgentId
	}
	return ""
}

func (m *AddMySQLRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *AddMySQLRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *AddMySQLRequest) GetReplicationSet() string {
	if m != nil {
		return m.ReplicationSet
	}
	return ""
}

func (m *AddMySQLRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddMySQLRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddMySQLRequest) GetQanUsername() string {
	if m != nil {
		return m.QanUsername
	}
	return ""
}

func (m *AddMySQLRequest) GetQanPassword() string {
	if m != nil {
		return m.QanPassword
	}
	return ""
}

func (m *AddMySQLRequest) GetMysqldExporter() bool {
	if m != nil {
		return m.MysqldExporter
	}
	return false
}

func (m *AddMySQLRequest) GetQanMysqlPerfschema() bool {
	if m != nil {
		return m.QanMysqlPerfschema
	}
	return false
}

func (m *AddMySQLRequest) GetQanMysqlSlowlog() bool {
	if m != nil {
		return m.QanMysqlSlowlog
	}
	return false
}

type AddMySQLResponse struct {
	Service              *inventorypb.MySQLService            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MysqldExporter       *inventorypb.MySQLdExporter          `protobuf:"bytes,2,opt,name=mysqld_exporter,json=mysqldExporter,proto3" json:"mysqld_exporter,omitempty"`
	QanMysqlPerfschema   *inventorypb.QANMySQLPerfSchemaAgent `protobuf:"bytes,3,opt,name=qan_mysql_perfschema,json=qanMysqlPerfschema,proto3" json:"qan_mysql_perfschema,omitempty"`
	QanMysqlSlowlog      *inventorypb.QANMySQLSlowlogAgent    `protobuf:"bytes,4,opt,name=qan_mysql_slowlog,json=qanMysqlSlowlog,proto3" json:"qan_mysql_slowlog,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AddMySQLResponse) Reset()         { *m = AddMySQLResponse{} }
func (m *AddMySQLResponse) String() string { return proto.CompactTextString(m) }
func (*AddMySQLResponse) ProtoMessage()    {}
func (*AddMySQLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab81470951176953, []int{1}
}

func (m *AddMySQLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLResponse.Unmarshal(m, b)
}
func (m *AddMySQLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLResponse.Marshal(b, m, deterministic)
}
func (m *AddMySQLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLResponse.Merge(m, src)
}
func (m *AddMySQLResponse) XXX_Size() int {
	return xxx_messageInfo_AddMySQLResponse.Size(m)
}
func (m *AddMySQLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLResponse proto.InternalMessageInfo

func (m *AddMySQLResponse) GetService() *inventorypb.MySQLService {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *AddMySQLResponse) GetMysqldExporter() *inventorypb.MySQLdExporter {
	if m != nil {
		return m.MysqldExporter
	}
	return nil
}

func (m *AddMySQLResponse) GetQanMysqlPerfschema() *inventorypb.QANMySQLPerfSchemaAgent {
	if m != nil {
		return m.QanMysqlPerfschema
	}
	return nil
}

func (m *AddMySQLResponse) GetQanMysqlSlowlog() *inventorypb.QANMySQLSlowlogAgent {
	if m != nil {
		return m.QanMysqlSlowlog
	}
	return nil
}

func init() {
	proto.RegisterType((*AddMySQLRequest)(nil), "management.AddMySQLRequest")
	proto.RegisterType((*AddMySQLResponse)(nil), "management.AddMySQLResponse")
}

func init() { proto.RegisterFile("managementpb/mysql.proto", fileDescriptor_ab81470951176953) }

var fileDescriptor_ab81470951176953 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x4e, 0xdb, 0x3c,
	0x18, 0xc6, 0xbf, 0x96, 0xd2, 0x82, 0x5b, 0x28, 0x9f, 0xf5, 0x49, 0x9f, 0xd7, 0x31, 0x91, 0x95,
	0x03, 0x0a, 0x5b, 0x93, 0xc1, 0xa4, 0x1d, 0xec, 0x2c, 0x48, 0x1c, 0xa0, 0x01, 0x82, 0x74, 0x93,
	0xa6, 0x9d, 0x44, 0x6e, 0xfd, 0x12, 0xa2, 0x25, 0x76, 0x6a, 0xbb, 0xed, 0x38, 0xdd, 0x25, 0x6c,
	0xb7, 0x31, 0x69, 0x17, 0xb3, 0x0b, 0x98, 0x34, 0xed, 0x42, 0xa6, 0x38, 0x7f, 0x5a, 0xa0, 0x3b,
	0x6a, 0xfd, 0x3c, 0x3f, 0x3f, 0xce, 0xfb, 0xe6, 0x8d, 0x11, 0x89, 0x29, 0xa7, 0x01, 0xc4, 0xc0,
	0x75, 0x32, 0x74, 0xe2, 0x5b, 0x35, 0x8e, 0xec, 0x44, 0x0a, 0x2d, 0x30, 0x9a, 0x3b, 0x9d, 0x57,
	0x41, 0xa8, 0x6f, 0x26, 0x43, 0x7b, 0x24, 0x62, 0x27, 0x9e, 0x85, 0xfa, 0xa3, 0x98, 0x39, 0x81,
	0xe8, 0x1b, 0xb0, 0x3f, 0xa5, 0x51, 0xc8, 0xa8, 0x16, 0x52, 0x39, 0xe5, 0xdf, 0x2c, 0xa3, 0xb3,
	0x1d, 0x08, 0x11, 0x44, 0xe0, 0xd0, 0x24, 0x74, 0x28, 0xe7, 0x42, 0x53, 0x1d, 0x0a, 0xae, 0x72,
	0x97, 0x84, 0x7c, 0x0a, 0x5c, 0x0b, 0x79, 0x9b, 0x0c, 0x1d, 0x1a, 0x00, 0xd7, 0x85, 0xd3, 0x59,
	0x74, 0x14, 0xc8, 0x69, 0x38, 0x82, 0xc2, 0x7b, 0x6e, 0x7e, 0x46, 0xfd, 0x00, 0x78, 0x5f, 0xcd,
	0x68, 0x10, 0x80, 0x74, 0x44, 0x62, 0x72, 0x1f, 0x9e, 0xd1, 0xfd, 0x5e, 0x43, 0x6d, 0x97, 0xb1,
	0xf3, 0xdb, 0xc1, 0xd5, 0x99, 0x07, 0xe3, 0x09, 0x28, 0x8d, 0x77, 0x50, 0x83, 0x0b, 0x06, 0x7e,
	0xc8, 0x48, 0xc5, 0xaa, 0xf4, 0xd6, 0x8f, 0xeb, 0xbf, 0x7e, 0xee, 0x54, 0xdf, 0x57, 0xbc, 0x7a,
	0x2a, 0x9f, 0x32, 0xbc, 0x8f, 0x5a, 0xf9, 0xa1, 0x3e, 0xa7, 0x31, 0x90, 0xea, 0x1d, 0xaa, 0x99,
	0x7b, 0x17, 0x34, 0x06, 0x6c, 0xa1, 0x06, 0x65, 0x4c, 0x82, 0x52, 0x64, 0xe5, 0x0e, 0x55, 0xc8,
	0xb8, 0x83, 0x6a, 0x89, 0x90, 0x9a, 0xd4, 0xac, 0x4a, 0x6f, 0x23, 0xb3, 0xb7, 0xfe, 0xf1, 0x8c,
	0x86, 0x7b, 0xa8, 0x95, 0xc4, 0xb1, 0x6f, 0x6a, 0x4f, 0x1f, 0x67, 0xf5, 0x4e, 0x04, 0x4a, 0xe2,
	0xd8, 0x4d, 0xad, 0x53, 0x86, 0x2d, 0xd4, 0x04, 0x3e, 0x0d, 0xa5, 0xe0, 0xe9, 0x0b, 0x21, 0xf5,
	0x14, 0xf4, 0x16, 0x25, 0x4c, 0x50, 0x63, 0x14, 0x4d, 0x94, 0x06, 0x49, 0x1a, 0xc6, 0x2d, 0x96,
	0x78, 0x0f, 0xb5, 0x25, 0x24, 0x51, 0x38, 0x32, 0x9d, 0xf1, 0x15, 0x68, 0xb2, 0x66, 0x88, 0xcd,
	0x05, 0x79, 0x00, 0x1a, 0x77, 0xd0, 0xda, 0x44, 0x81, 0x34, 0x35, 0xaf, 0x1b, 0xa2, 0x5c, 0xa7,
	0x5e, 0x42, 0x95, 0x9a, 0x09, 0xc9, 0x08, 0xca, 0xbc, 0x62, 0x8d, 0x9f, 0xa2, 0xd6, 0x98, 0x72,
	0xbf, 0xdc, 0xdb, 0xcc, 0x9e, 0x6e, 0x4c, 0xf9, 0xbb, 0x62, 0x7b, 0x8e, 0x94, 0x11, 0xad, 0x12,
	0xb9, 0x2c, 0x52, 0xf6, 0x50, 0xdb, 0xcc, 0x1f, 0xf3, 0xe1, 0x53, 0xda, 0x1d, 0x90, 0x64, 0xc3,
	0xaa, 0xf4, 0xd6, 0xbc, 0xcd, 0x4c, 0x3e, 0xc9, 0x55, 0xfc, 0x02, 0xfd, 0x97, 0x66, 0x19, 0xd5,
	0x4f, 0x40, 0x5e, 0xab, 0xd1, 0x0d, 0xc4, 0x94, 0x6c, 0x1a, 0x1a, 0x8f, 0x29, 0x3f, 0x4f, 0xad,
	0xcb, 0xd2, 0xc1, 0x07, 0xe8, 0xdf, 0xf9, 0x0e, 0x15, 0x89, 0x59, 0x24, 0x02, 0xd2, 0x36, 0x78,
	0xbb, 0xc0, 0x07, 0x99, 0xdc, 0xfd, 0x56, 0x45, 0x5b, 0xf3, 0x89, 0x51, 0x89, 0xe0, 0x0a, 0xf0,
	0x21, 0x6a, 0xe4, 0x6f, 0xdd, 0x8c, 0x4c, 0xf3, 0xe8, 0x7f, 0xbb, 0x1c, 0x51, 0xdb, 0xa0, 0x83,
	0xcc, 0xf6, 0x0a, 0x0e, 0x1f, 0x3f, 0x2c, 0xa7, 0x6a, 0xb6, 0x3e, 0xba, 0xbf, 0xb5, 0xac, 0xec,
	0x41, 0xa5, 0x6f, 0xff, 0x52, 0xe9, 0x8a, 0x09, 0xea, 0x2e, 0x04, 0x5d, 0xb9, 0x17, 0x26, 0x2b,
	0x2d, 0x7a, 0x60, 0x20, 0x33, 0x39, 0x4b, 0xbb, 0xf1, 0x66, 0x59, 0x37, 0x6a, 0x26, 0x72, 0x67,
	0x49, 0x64, 0xde, 0x98, 0x2c, 0xef, 0x7e, 0xbb, 0x8e, 0x38, 0x5a, 0x35, 0x14, 0x06, 0xb4, 0xe2,
	0x32, 0x86, 0x1f, 0xdb, 0xf3, 0x7b, 0xc3, 0xbe, 0xf7, 0xe5, 0x75, 0xb6, 0x97, 0x9b, 0x59, 0x93,
	0xbb, 0xbb, 0x9f, 0x7f, 0xfc, 0xfe, 0x5a, 0x7d, 0xd2, 0x25, 0xce, 0xf4, 0xd0, 0x99, 0x83, 0x8e,
	0xa1, 0x1c, 0x97, 0xb1, 0xd7, 0x95, 0x83, 0x63, 0xff, 0x8b, 0x7b, 0xe1, 0x9d, 0xa1, 0x06, 0x83,
	0x6b, 0x3a, 0x89, 0x34, 0x76, 0x11, 0x76, 0xb9, 0x05, 0x52, 0x0a, 0x69, 0xc9, 0x3c, 0xc9, 0xc6,
	0xcf, 0xd0, 0x7e, 0x67, 0x6f, 0xd7, 0x61, 0x70, 0x1d, 0xf2, 0x30, 0xbb, 0x1b, 0x16, 0x2f, 0xba,
	0x93, 0x14, 0x2f, 0xce, 0xfd, 0xd0, 0x5a, 0xb4, 0x86, 0x75, 0x73, 0x71, 0xbc, 0xfc, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x03, 0xbd, 0xe8, 0x02, 0x1a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MySQLClient is the client API for MySQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MySQLClient interface {
	// Add MySQL adds MySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mysqld_exporter", and "qan_mysql_perfschema" agents
	// with provided "pmm_agent_id" and other parameters.
	Add(ctx context.Context, in *AddMySQLRequest, opts ...grpc.CallOption) (*AddMySQLResponse, error)
}

type mySQLClient struct {
	cc *grpc.ClientConn
}

func NewMySQLClient(cc *grpc.ClientConn) MySQLClient {
	return &mySQLClient{cc}
}

func (c *mySQLClient) Add(ctx context.Context, in *AddMySQLRequest, opts ...grpc.CallOption) (*AddMySQLResponse, error) {
	out := new(AddMySQLResponse)
	err := c.cc.Invoke(ctx, "/management.MySQL/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MySQLServer is the server API for MySQL service.
type MySQLServer interface {
	// Add MySQL adds MySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mysqld_exporter", and "qan_mysql_perfschema" agents
	// with provided "pmm_agent_id" and other parameters.
	Add(context.Context, *AddMySQLRequest) (*AddMySQLResponse, error)
}

func RegisterMySQLServer(s *grpc.Server, srv MySQLServer) {
	s.RegisterService(&_MySQL_serviceDesc, srv)
}

func _MySQL_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.MySQL/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServer).Add(ctx, req.(*AddMySQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MySQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.MySQL",
	HandlerType: (*MySQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MySQL_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/mysql.proto",
}