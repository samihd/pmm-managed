// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qanpb/profile.proto

package qanpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// ReportRequest defines filtering of metrics report for db server or other dimentions.
type ReportRequest struct {
	PeriodStartFrom      *timestamp.Timestamp   `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo        *timestamp.Timestamp   `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	GroupBy              string                 `protobuf:"bytes,3,opt,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Labels               []*ReportMapFieldEntry `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Columns              []string               `protobuf:"bytes,5,rep,name=columns,proto3" json:"columns,omitempty"`
	OrderBy              string                 `protobuf:"bytes,6,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Offset               uint32                 `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32                 `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	MainMetric           string                 `protobuf:"bytes,9,opt,name=main_metric,json=mainMetric,proto3" json:"main_metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41aef16fb960ebb7, []int{0}
}

func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (m *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(m, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetPeriodStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartFrom
	}
	return nil
}

func (m *ReportRequest) GetPeriodStartTo() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartTo
	}
	return nil
}

func (m *ReportRequest) GetGroupBy() string {
	if m != nil {
		return m.GroupBy
	}
	return ""
}

func (m *ReportRequest) GetLabels() []*ReportMapFieldEntry {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ReportRequest) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *ReportRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ReportRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReportRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReportRequest) GetMainMetric() string {
	if m != nil {
		return m.MainMetric
	}
	return ""
}

// ReportMapFieldEntry allows to pass labels/dimentions in form like {"server": ["db1", "db2"...]}.
type ReportMapFieldEntry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportMapFieldEntry) Reset()         { *m = ReportMapFieldEntry{} }
func (m *ReportMapFieldEntry) String() string { return proto.CompactTextString(m) }
func (*ReportMapFieldEntry) ProtoMessage()    {}
func (*ReportMapFieldEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_41aef16fb960ebb7, []int{1}
}

func (m *ReportMapFieldEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportMapFieldEntry.Unmarshal(m, b)
}
func (m *ReportMapFieldEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportMapFieldEntry.Marshal(b, m, deterministic)
}
func (m *ReportMapFieldEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportMapFieldEntry.Merge(m, src)
}
func (m *ReportMapFieldEntry) XXX_Size() int {
	return xxx_messageInfo_ReportMapFieldEntry.Size(m)
}
func (m *ReportMapFieldEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportMapFieldEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ReportMapFieldEntry proto.InternalMessageInfo

func (m *ReportMapFieldEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReportMapFieldEntry) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

// ReportReply is list of reports per quieryids, hosts etc.
type ReportReply struct {
	TotalRows            uint32   `protobuf:"varint,1,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Rows                 []*Row   `protobuf:"bytes,4,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportReply) Reset()         { *m = ReportReply{} }
func (m *ReportReply) String() string { return proto.CompactTextString(m) }
func (*ReportReply) ProtoMessage()    {}
func (*ReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_41aef16fb960ebb7, []int{2}
}

func (m *ReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportReply.Unmarshal(m, b)
}
func (m *ReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportReply.Marshal(b, m, deterministic)
}
func (m *ReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportReply.Merge(m, src)
}
func (m *ReportReply) XXX_Size() int {
	return xxx_messageInfo_ReportReply.Size(m)
}
func (m *ReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReportReply proto.InternalMessageInfo

func (m *ReportReply) GetTotalRows() uint32 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}

func (m *ReportReply) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReportReply) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReportReply) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

// Row define metrics for selected dimention.
type Row struct {
	Rank                 uint32             `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Dimension            string             `protobuf:"bytes,2,opt,name=dimension,proto3" json:"dimension,omitempty"`
	Metrics              map[string]*Metric `protobuf:"bytes,3,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Sparkline            []*Point           `protobuf:"bytes,4,rep,name=sparkline,proto3" json:"sparkline,omitempty"`
	Fingerprint          string             `protobuf:"bytes,5,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	NumQueries           uint32             `protobuf:"varint,6,opt,name=num_queries,json=numQueries,proto3" json:"num_queries,omitempty"`
	Qps                  float32            `protobuf:"fixed32,7,opt,name=qps,proto3" json:"qps,omitempty"`
	Load                 float32            `protobuf:"fixed32,8,opt,name=load,proto3" json:"load,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_41aef16fb960ebb7, []int{3}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Row) GetDimension() string {
	if m != nil {
		return m.Dimension
	}
	return ""
}

func (m *Row) GetMetrics() map[string]*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Row) GetSparkline() []*Point {
	if m != nil {
		return m.Sparkline
	}
	return nil
}

func (m *Row) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *Row) GetNumQueries() uint32 {
	if m != nil {
		return m.NumQueries
	}
	return 0
}

func (m *Row) GetQps() float32 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func (m *Row) GetLoad() float32 {
	if m != nil {
		return m.Load
	}
	return 0
}

// Metric cell.
type Metric struct {
	Stats                *Stat    `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_41aef16fb960ebb7, []int{4}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetStats() *Stat {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Stat is statistics of specific metric.
type Stat struct {
	Rate                 float32  `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	Cnt                  float32  `protobuf:"fixed32,2,opt,name=cnt,proto3" json:"cnt,omitempty"`
	Sum                  float32  `protobuf:"fixed32,3,opt,name=sum,proto3" json:"sum,omitempty"`
	Min                  float32  `protobuf:"fixed32,4,opt,name=min,proto3" json:"min,omitempty"`
	Max                  float32  `protobuf:"fixed32,5,opt,name=max,proto3" json:"max,omitempty"`
	P99                  float32  `protobuf:"fixed32,6,opt,name=p99,proto3" json:"p99,omitempty"`
	Avg                  float32  `protobuf:"fixed32,7,opt,name=avg,proto3" json:"avg,omitempty"`
	SumPerSec            float32  `protobuf:"fixed32,8,opt,name=sum_per_sec,json=sumPerSec,proto3" json:"sum_per_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_41aef16fb960ebb7, []int{5}
}

func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}
func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}
func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *Stat) GetCnt() float32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func (m *Stat) GetSum() float32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *Stat) GetMin() float32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Stat) GetMax() float32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *Stat) GetP99() float32 {
	if m != nil {
		return m.P99
	}
	return 0
}

func (m *Stat) GetAvg() float32 {
	if m != nil {
		return m.Avg
	}
	return 0
}

func (m *Stat) GetSumPerSec() float32 {
	if m != nil {
		return m.SumPerSec
	}
	return 0
}

func init() {
	proto.RegisterType((*ReportRequest)(nil), "qan.ReportRequest")
	proto.RegisterType((*ReportMapFieldEntry)(nil), "qan.ReportMapFieldEntry")
	proto.RegisterType((*ReportReply)(nil), "qan.ReportReply")
	proto.RegisterType((*Row)(nil), "qan.Row")
	proto.RegisterMapType((map[string]*Metric)(nil), "qan.Row.MetricsEntry")
	proto.RegisterType((*Metric)(nil), "qan.Metric")
	proto.RegisterType((*Stat)(nil), "qan.Stat")
}

func init() { proto.RegisterFile("qanpb/profile.proto", fileDescriptor_41aef16fb960ebb7) }

var fileDescriptor_41aef16fb960ebb7 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x86, 0xa8, 0x3f, 0x73, 0x54, 0xc1, 0xc9, 0xa6, 0x0d, 0x58, 0xc1, 0xad, 0x59, 0xb6, 0x07,
	0xb5, 0x68, 0xc4, 0xc0, 0xbd, 0x34, 0x06, 0x7a, 0x88, 0x01, 0xdb, 0xa7, 0xa0, 0xce, 0x3a, 0xa7,
	0x5c, 0x84, 0x95, 0xb4, 0x24, 0x16, 0x26, 0x77, 0x57, 0xbb, 0x4b, 0xcb, 0xba, 0xf6, 0x11, 0xda,
	0x77, 0xe8, 0x93, 0xf4, 0x0d, 0xfa, 0x0a, 0x3d, 0xf6, 0x21, 0x8a, 0x9d, 0xa5, 0x6a, 0x19, 0x10,
	0x90, 0x93, 0x66, 0x3e, 0xcd, 0x7c, 0xf3, 0x71, 0x3e, 0x72, 0xe0, 0xc5, 0x9a, 0x49, 0xbd, 0xc8,
	0xb5, 0x51, 0x85, 0xa8, 0xf8, 0x4c, 0x1b, 0xe5, 0x14, 0xe9, 0xae, 0x99, 0x9c, 0x9c, 0x94, 0x4a,
	0x95, 0x15, 0xcf, 0x99, 0x16, 0x39, 0x93, 0x52, 0x39, 0xe6, 0x84, 0x92, 0x36, 0x94, 0x4c, 0x4e,
	0xdb, 0x7f, 0x31, 0x5b, 0x34, 0x45, 0xee, 0x44, 0xcd, 0xad, 0x63, 0xb5, 0x6e, 0x0b, 0x7e, 0xc4,
	0x9f, 0xe5, 0xab, 0x92, 0xcb, 0x57, 0x76, 0xc3, 0xca, 0x92, 0x9b, 0x5c, 0x69, 0xa4, 0x38, 0x40,
	0x77, 0x1c, 0x64, 0xac, 0x99, 0x0c, 0x40, 0xf6, 0x6f, 0x04, 0x63, 0xca, 0xb5, 0x32, 0x8e, 0xf2,
	0x75, 0xc3, 0xad, 0x23, 0x57, 0xf0, 0x5c, 0x73, 0x23, 0xd4, 0x6a, 0x6e, 0x1d, 0x33, 0x6e, 0x5e,
	0x18, 0x55, 0x27, 0x9d, 0xb4, 0x33, 0x1d, 0x9d, 0x4d, 0x66, 0x41, 0xcd, 0x6c, 0xa7, 0x66, 0xf6,
	0x61, 0xa7, 0x86, 0x1e, 0x87, 0xa6, 0x5b, 0xdf, 0x73, 0x65, 0x54, 0x4d, 0x2e, 0xe0, 0xf8, 0x09,
	0x8f, 0x53, 0x49, 0xf4, 0x49, 0x96, 0xf1, 0x1e, 0xcb, 0x07, 0x45, 0xbe, 0x84, 0xa3, 0xd2, 0xa8,
	0x46, 0xcf, 0x17, 0xdb, 0xa4, 0x9b, 0x76, 0xa6, 0x31, 0x1d, 0x62, 0x7e, 0xb1, 0x25, 0xaf, 0x61,
	0x50, 0xb1, 0x05, 0xaf, 0x6c, 0xd2, 0x4b, 0xbb, 0xd3, 0xd1, 0x59, 0x32, 0xf3, 0x0f, 0x15, 0x1e,
	0xe5, 0x1d, 0xd3, 0x57, 0x82, 0x57, 0xab, 0x4b, 0xe9, 0xcc, 0x96, 0xb6, 0x75, 0x24, 0x81, 0xe1,
	0x52, 0x55, 0x4d, 0x2d, 0x6d, 0xd2, 0x4f, 0xbb, 0x9e, 0xab, 0x4d, 0xfd, 0x18, 0x65, 0x56, 0xdc,
	0xf8, 0x31, 0x83, 0x30, 0x06, 0xf3, 0x8b, 0x2d, 0x79, 0x09, 0x03, 0x55, 0x14, 0x96, 0xbb, 0x64,
	0x98, 0x76, 0xa6, 0x63, 0xda, 0x66, 0xe4, 0x73, 0xe8, 0x57, 0xa2, 0x16, 0x2e, 0x39, 0x42, 0x38,
	0x24, 0xe4, 0x14, 0x46, 0x35, 0x13, 0x72, 0x5e, 0x73, 0x67, 0xc4, 0x32, 0x89, 0x91, 0x0b, 0x3c,
	0xf4, 0x0e, 0x91, 0xec, 0x17, 0x78, 0x71, 0x40, 0x22, 0x79, 0x06, 0xdd, 0x3b, 0xbe, 0xc5, 0x2d,
	0xc7, 0xd4, 0x87, 0x9e, 0xff, 0x9e, 0x55, 0x0d, 0x4f, 0x22, 0x94, 0x1a, 0x92, 0xec, 0x01, 0x46,
	0x3b, 0xb3, 0x74, 0xb5, 0x25, 0x5f, 0x01, 0x38, 0xe5, 0x58, 0x35, 0x37, 0x6a, 0x63, 0xb1, 0x7b,
	0x4c, 0x63, 0x44, 0xa8, 0xda, 0xd8, 0x3d, 0xed, 0xd1, 0x61, 0xed, 0xdd, 0x7d, 0xed, 0x27, 0xd0,
	0x43, 0x9a, 0xb0, 0xce, 0xa3, 0xb0, 0x4e, 0xb5, 0xa1, 0x88, 0x66, 0x7f, 0x45, 0xd0, 0xa5, 0x6a,
	0x43, 0x08, 0xf4, 0x0c, 0x93, 0x77, 0xed, 0x30, 0x8c, 0xc9, 0x09, 0xc4, 0x2b, 0x51, 0x73, 0x69,
	0x85, 0x92, 0x38, 0x2a, 0xa6, 0x8f, 0x00, 0xc9, 0x61, 0x18, 0xd6, 0x61, 0x93, 0x2e, 0x52, 0x7f,
	0xb1, 0xa3, 0x9e, 0x85, 0xa5, 0xd8, 0x60, 0xd3, 0xae, 0x8a, 0x4c, 0x21, 0xb6, 0x9a, 0x99, 0xbb,
	0x4a, 0x48, 0xde, 0xaa, 0x01, 0x6c, 0xb9, 0x51, 0x42, 0x3a, 0xfa, 0xf8, 0x27, 0x49, 0x61, 0x54,
	0x08, 0x59, 0x72, 0xa3, 0x8d, 0x90, 0x2e, 0xe9, 0xe3, 0xe8, 0x7d, 0xc8, 0x1b, 0x22, 0x9b, 0x7a,
	0xbe, 0x6e, 0xb8, 0x11, 0xdc, 0xa2, 0xb9, 0x63, 0x0a, 0xb2, 0xa9, 0xdf, 0x07, 0xc4, 0x6f, 0x7e,
	0xad, 0x2d, 0x9a, 0x1b, 0x51, 0x1f, 0xfa, 0x27, 0xac, 0x14, 0x5b, 0xa1, 0xb1, 0x11, 0xc5, 0x78,
	0x72, 0x0d, 0x9f, 0xed, 0x6b, 0x3d, 0xe0, 0xd7, 0x37, 0x8f, 0x7e, 0xf9, 0x77, 0x7c, 0x84, 0x82,
	0x43, 0x4f, 0x6b, 0xde, 0x79, 0xf4, 0x73, 0x27, 0xfb, 0x1e, 0x06, 0x01, 0x24, 0xa7, 0xd0, 0xb7,
	0x8e, 0x39, 0xdb, 0x7e, 0x5a, 0x31, 0x36, 0xdc, 0x3a, 0xe6, 0x68, 0xc0, 0xb3, 0x3f, 0x3b, 0xd0,
	0xf3, 0x79, 0x58, 0xb9, 0xe3, 0x58, 0x18, 0x51, 0x8c, 0xbd, 0x80, 0xa5, 0x0c, 0xbe, 0x46, 0xd4,
	0x87, 0x1e, 0xb1, 0x4d, 0x8d, 0x96, 0x46, 0xd4, 0x87, 0x1e, 0xa9, 0x85, 0x4c, 0x7a, 0x01, 0xa9,
	0x85, 0x44, 0x84, 0x3d, 0xe0, 0x9e, 0x3c, 0xc2, 0x1e, 0x3c, 0xa2, 0xdf, 0xbc, 0xc1, 0xbd, 0x44,
	0xd4, 0x87, 0x1e, 0x61, 0xf7, 0xe5, 0x6e, 0x21, 0xec, 0xbe, 0x24, 0x5f, 0xc3, 0xc8, 0x36, 0xf5,
	0x5c, 0x73, 0x33, 0xb7, 0x7c, 0xd9, 0xee, 0x25, 0xb6, 0x4d, 0x7d, 0xc3, 0xcd, 0x2d, 0x5f, 0x9e,
	0x7d, 0x84, 0xe1, 0x4d, 0x38, 0x6b, 0xe4, 0x57, 0x88, 0xaf, 0xb9, 0x0b, 0xaf, 0x28, 0x21, 0x7b,
	0x5f, 0x64, 0x7b, 0x5c, 0x26, 0xcf, 0x9e, 0x60, 0xba, 0xda, 0x66, 0x27, 0xbf, 0xfd, 0xfd, 0xcf,
	0x1f, 0xd1, 0xcb, 0xec, 0x79, 0x7e, 0xff, 0xda, 0xdf, 0xa5, 0xfc, 0x7f, 0x82, 0xf3, 0xce, 0x0f,
	0x17, 0xef, 0x7f, 0x7f, 0x7b, 0x4d, 0x2f, 0x61, 0xb8, 0xe2, 0x05, 0x6b, 0x2a, 0x47, 0xce, 0x81,
	0xbc, 0x95, 0x29, 0x37, 0x46, 0x99, 0xd4, 0x70, 0xab, 0x95, 0xb4, 0x7c, 0x46, 0xbe, 0x83, 0x6c,
	0x92, 0x7e, 0x9b, 0xaf, 0x78, 0x21, 0xa4, 0x08, 0xa7, 0x0f, 0xef, 0xdc, 0xa5, 0xaf, 0xa3, 0x6d,
	0xd9, 0xc7, 0x3e, 0x62, 0x8b, 0x01, 0x9e, 0x9d, 0x9f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x59,
	0xa0, 0xa2, 0x25, 0x92, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileClient interface {
	// GetReport returns list of metrics group by queryid or other dimentions.
	GetReport(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetReport(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, "/qan.Profile/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
type ProfileServer interface {
	// GetReport returns list of metrics group by queryid or other dimentions.
	GetReport(context.Context, *ReportRequest) (*ReportReply, error)
}

// UnimplementedProfileServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (*UnimplementedProfileServer) GetReport(ctx context.Context, req *ReportRequest) (*ReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.Profile/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetReport(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qan.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReport",
			Handler:    _Profile_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/profile.proto",
}
