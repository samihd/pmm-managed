// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/mongodb.proto

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

type AddMongoDBRequest struct {
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
	// MongoDB username for exporter and QAN agent access.
	Username string `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	// MongoDB password for exporter and QAN agent access.
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	// If true, adds mongodb_exporter for provided service.
	MongodbExporter bool `protobuf:"varint,11,opt,name=mongodb_exporter,json=mongodbExporter,proto3" json:"mongodb_exporter,omitempty"`
	// If true, adds qan-mongodb-profiler-agent for provided service.
	QanMongodbProfiler   bool     `protobuf:"varint,12,opt,name=qan_mongodb_profiler,json=qanMongodbProfiler,proto3" json:"qan_mongodb_profiler,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMongoDBRequest) Reset()         { *m = AddMongoDBRequest{} }
func (m *AddMongoDBRequest) String() string { return proto.CompactTextString(m) }
func (*AddMongoDBRequest) ProtoMessage()    {}
func (*AddMongoDBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_593aa4f9c0b43a5e, []int{0}
}

func (m *AddMongoDBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMongoDBRequest.Unmarshal(m, b)
}
func (m *AddMongoDBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMongoDBRequest.Marshal(b, m, deterministic)
}
func (m *AddMongoDBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMongoDBRequest.Merge(m, src)
}
func (m *AddMongoDBRequest) XXX_Size() int {
	return xxx_messageInfo_AddMongoDBRequest.Size(m)
}
func (m *AddMongoDBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMongoDBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMongoDBRequest proto.InternalMessageInfo

func (m *AddMongoDBRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddMongoDBRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AddMongoDBRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddMongoDBRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AddMongoDBRequest) GetPmmAgentId() string {
	if m != nil {
		return m.PmmAgentId
	}
	return ""
}

func (m *AddMongoDBRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *AddMongoDBRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *AddMongoDBRequest) GetReplicationSet() string {
	if m != nil {
		return m.ReplicationSet
	}
	return ""
}

func (m *AddMongoDBRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddMongoDBRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddMongoDBRequest) GetMongodbExporter() bool {
	if m != nil {
		return m.MongodbExporter
	}
	return false
}

func (m *AddMongoDBRequest) GetQanMongodbProfiler() bool {
	if m != nil {
		return m.QanMongodbProfiler
	}
	return false
}

type AddMongoDBResponse struct {
	Service              *inventorypb.MongoDBService          `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MongodbExporter      *inventorypb.MongoDBExporter         `protobuf:"bytes,2,opt,name=mongodb_exporter,json=mongodbExporter,proto3" json:"mongodb_exporter,omitempty"`
	QanMongodbProfiler   *inventorypb.QANMongoDBProfilerAgent `protobuf:"bytes,3,opt,name=qan_mongodb_profiler,json=qanMongodbProfiler,proto3" json:"qan_mongodb_profiler,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AddMongoDBResponse) Reset()         { *m = AddMongoDBResponse{} }
func (m *AddMongoDBResponse) String() string { return proto.CompactTextString(m) }
func (*AddMongoDBResponse) ProtoMessage()    {}
func (*AddMongoDBResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_593aa4f9c0b43a5e, []int{1}
}

func (m *AddMongoDBResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMongoDBResponse.Unmarshal(m, b)
}
func (m *AddMongoDBResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMongoDBResponse.Marshal(b, m, deterministic)
}
func (m *AddMongoDBResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMongoDBResponse.Merge(m, src)
}
func (m *AddMongoDBResponse) XXX_Size() int {
	return xxx_messageInfo_AddMongoDBResponse.Size(m)
}
func (m *AddMongoDBResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMongoDBResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMongoDBResponse proto.InternalMessageInfo

func (m *AddMongoDBResponse) GetService() *inventorypb.MongoDBService {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *AddMongoDBResponse) GetMongodbExporter() *inventorypb.MongoDBExporter {
	if m != nil {
		return m.MongodbExporter
	}
	return nil
}

func (m *AddMongoDBResponse) GetQanMongodbProfiler() *inventorypb.QANMongoDBProfilerAgent {
	if m != nil {
		return m.QanMongodbProfiler
	}
	return nil
}

func init() {
	proto.RegisterType((*AddMongoDBRequest)(nil), "management.AddMongoDBRequest")
	proto.RegisterType((*AddMongoDBResponse)(nil), "management.AddMongoDBResponse")
}

func init() { proto.RegisterFile("managementpb/mongodb.proto", fileDescriptor_593aa4f9c0b43a5e) }

var fileDescriptor_593aa4f9c0b43a5e = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xd1, 0x6e, 0xd3, 0x3a,
	0x18, 0xc7, 0x4f, 0xba, 0x9d, 0x66, 0x73, 0x7b, 0xce, 0x76, 0xac, 0x73, 0x11, 0x22, 0x60, 0x51,
	0x11, 0x5a, 0x07, 0x34, 0x86, 0x4d, 0xe2, 0x82, 0xbb, 0x4c, 0xec, 0x62, 0x12, 0x9b, 0x20, 0xe3,
	0x02, 0x71, 0x13, 0xb9, 0xf5, 0xb7, 0x60, 0xd1, 0xd8, 0x99, 0xed, 0xb6, 0x70, 0xc1, 0x0d, 0x8f,
	0x00, 0x4f, 0xc4, 0x33, 0xf0, 0x00, 0x48, 0x08, 0xf1, 0x1c, 0x28, 0x8e, 0xb3, 0x66, 0xea, 0x76,
	0xd5, 0xfa, 0xfb, 0xff, 0xfc, 0xb7, 0xff, 0xfe, 0xec, 0xa0, 0xb0, 0xa0, 0x82, 0xe6, 0x50, 0x80,
	0x30, 0xe5, 0x98, 0x14, 0x52, 0xe4, 0x92, 0x8d, 0xe3, 0x52, 0x49, 0x23, 0x31, 0x5a, 0x6a, 0xe1,
	0xd3, 0x9c, 0x9b, 0x77, 0xb3, 0x71, 0x3c, 0x91, 0x05, 0x29, 0x16, 0xdc, 0xbc, 0x97, 0x0b, 0x92,
	0xcb, 0x91, 0x05, 0x47, 0x73, 0x3a, 0xe5, 0x8c, 0x1a, 0xa9, 0x34, 0xb9, 0xfc, 0x5b, 0x7b, 0x84,
	0xb7, 0x73, 0x29, 0xf3, 0x29, 0x10, 0x5a, 0x72, 0x42, 0x85, 0x90, 0x86, 0x1a, 0x2e, 0x85, 0x76,
	0x6a, 0xc0, 0xc5, 0x1c, 0x84, 0x91, 0xea, 0x63, 0x39, 0x26, 0x34, 0x07, 0x61, 0x1a, 0x25, 0x6c,
	0x2b, 0x1a, 0xd4, 0x9c, 0x4f, 0xa0, 0xd1, 0x1e, 0xd9, 0x9f, 0xc9, 0x28, 0x07, 0x31, 0xd2, 0x0b,
	0x9a, 0xe7, 0xa0, 0x88, 0x2c, 0xad, 0xef, 0xea, 0x1a, 0x83, 0x6f, 0x6b, 0xe8, 0xbf, 0x84, 0xb1,
	0x93, 0x2a, 0xda, 0xf3, 0xc3, 0x14, 0x2e, 0x66, 0xa0, 0x0d, 0xde, 0x41, 0xbe, 0x90, 0x0c, 0x32,
	0xce, 0x02, 0x2f, 0xf2, 0x86, 0x9b, 0x87, 0xdd, 0x9f, 0x3f, 0x76, 0x3a, 0x6f, 0xbc, 0xb4, 0x5b,
	0x95, 0x8f, 0x19, 0xde, 0x43, 0x7d, 0xb7, 0x6c, 0x26, 0x68, 0x01, 0x41, 0xe7, 0x0a, 0xd5, 0x73,
	0xda, 0x29, 0x2d, 0x00, 0x47, 0xc8, 0xa7, 0x8c, 0x29, 0xd0, 0x3a, 0x58, 0xbb, 0x42, 0x35, 0x65,
	0x1c, 0xa2, 0xf5, 0x52, 0x2a, 0x13, 0xac, 0x47, 0xde, 0xf0, 0x9f, 0x5a, 0xde, 0xfe, 0x2b, 0xb5,
	0x35, 0x3c, 0x44, 0xfd, 0xb2, 0x28, 0x32, 0x9b, 0xbe, 0xda, 0xce, 0xdf, 0x57, 0x2c, 0x50, 0x59,
	0x14, 0x49, 0x25, 0x1d, 0x33, 0x1c, 0xa1, 0x1e, 0x88, 0x39, 0x57, 0x52, 0x54, 0x2d, 0x09, 0xba,
	0x15, 0x98, 0xb6, 0x4b, 0x38, 0x40, 0xfe, 0x64, 0x3a, 0xd3, 0x06, 0x54, 0xe0, 0x5b, 0xb5, 0x19,
	0xe2, 0x5d, 0xb4, 0xa5, 0xa0, 0x9c, 0xf2, 0x89, 0x3d, 0x9b, 0x4c, 0x83, 0x09, 0x36, 0x2c, 0xf1,
	0x6f, 0xab, 0x7c, 0x06, 0x06, 0x87, 0x68, 0x63, 0xa6, 0x41, 0xd9, 0xcc, 0x9b, 0x96, 0xb8, 0x1c,
	0x57, 0x5a, 0x49, 0xb5, 0x5e, 0x48, 0xc5, 0x02, 0x54, 0x6b, 0xcd, 0x18, 0xef, 0xa1, 0x6d, 0x77,
	0x7b, 0x32, 0xf8, 0x50, 0x05, 0x03, 0x15, 0xf4, 0x22, 0x6f, 0xb8, 0x91, 0x6e, 0xb9, 0xfa, 0x91,
	0x2b, 0xe3, 0xc7, 0xe8, 0xff, 0x0b, 0x2a, 0xb2, 0x06, 0x2f, 0x95, 0x3c, 0xe7, 0x53, 0x50, 0x41,
	0xdf, 0xe2, 0xf8, 0x82, 0x8a, 0x93, 0x5a, 0x7a, 0xe9, 0x94, 0xc1, 0x6f, 0x0f, 0xe1, 0x76, 0x0f,
	0x75, 0x29, 0x85, 0x06, 0x7c, 0x80, 0x7c, 0xd7, 0x07, 0xdb, 0xc4, 0xde, 0xfe, 0xad, 0xf8, 0xf2,
	0xda, 0xc4, 0x0e, 0x3e, 0xab, 0x81, 0xb4, 0x21, 0xf1, 0xd1, 0x35, 0x1b, 0xed, 0xd8, 0xd9, 0xe1,
	0xea, 0xec, 0x66, 0xcf, 0xab, 0x21, 0x5e, 0xdf, 0x10, 0x62, 0xcd, 0x5a, 0x0d, 0x5a, 0x56, 0xaf,
	0x92, 0x53, 0xe7, 0xd6, 0xe4, 0xb1, 0x0d, 0xbd, 0x2e, 0xe8, 0xfe, 0x27, 0xe4, 0x3b, 0x16, 0x2b,
	0x84, 0x96, 0x91, 0xf1, 0x9d, 0x78, 0xf9, 0x18, 0xe3, 0x95, 0xeb, 0x1c, 0xde, 0xbd, 0x49, 0xae,
	0x4f, 0x6a, 0x70, 0xff, 0xf3, 0xf7, 0x5f, 0x5f, 0x3b, 0x3b, 0x83, 0x90, 0xcc, 0x9f, 0x90, 0x25,
	0x4a, 0x1c, 0x47, 0x12, 0xc6, 0x9e, 0x79, 0x0f, 0x0e, 0xb3, 0x2f, 0xc9, 0x69, 0xfa, 0x02, 0xf9,
	0x0c, 0xce, 0xe9, 0x6c, 0x6a, 0x70, 0x82, 0x70, 0x22, 0x22, 0x50, 0x4a, 0xaa, 0x48, 0x39, 0xaf,
	0x18, 0x3f, 0x44, 0x7b, 0xe1, 0xee, 0x3d, 0xc2, 0xe0, 0x9c, 0x0b, 0x5e, 0x3f, 0xbb, 0xf6, 0x57,
	0xe4, 0xa8, 0xc2, 0x9b, 0x95, 0xdf, 0xf6, 0xdb, 0xd2, 0xb8, 0x6b, 0xdf, 0xe4, 0xc1, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0x3b, 0x53, 0xe8, 0x77, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MongoDBClient is the client API for MongoDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MongoDBClient interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error)
}

type mongoDBClient struct {
	cc *grpc.ClientConn
}

func NewMongoDBClient(cc *grpc.ClientConn) MongoDBClient {
	return &mongoDBClient{cc}
}

func (c *mongoDBClient) AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error) {
	out := new(AddMongoDBResponse)
	err := c.cc.Invoke(ctx, "/management.MongoDB/AddMongoDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoDBServer is the server API for MongoDB service.
type MongoDBServer interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(context.Context, *AddMongoDBRequest) (*AddMongoDBResponse, error)
}

func RegisterMongoDBServer(s *grpc.Server, srv MongoDBServer) {
	s.RegisterService(&_MongoDB_serviceDesc, srv)
}

func _MongoDB_AddMongoDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMongoDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).AddMongoDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.MongoDB/AddMongoDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).AddMongoDB(ctx, req.(*AddMongoDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MongoDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.MongoDB",
	HandlerType: (*MongoDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMongoDB",
			Handler:    _MongoDB_AddMongoDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/mongodb.proto",
}