// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goraph.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Graph struct {
	Edge                 []*Edge   `protobuf:"bytes,1,rep,name=edge,proto3" json:"edge,omitempty"`
	Latlon               []*LatLon `protobuf:"bytes,2,rep,name=latlon,proto3" json:"latlon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Graph) Reset()         { *m = Graph{} }
func (m *Graph) String() string { return proto.CompactTextString(m) }
func (*Graph) ProtoMessage()    {}
func (*Graph) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebeae2f27825ff85, []int{0}
}

func (m *Graph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Graph.Unmarshal(m, b)
}
func (m *Graph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Graph.Marshal(b, m, deterministic)
}
func (m *Graph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Graph.Merge(m, src)
}
func (m *Graph) XXX_Size() int {
	return xxx_messageInfo_Graph.Size(m)
}
func (m *Graph) XXX_DiscardUnknown() {
	xxx_messageInfo_Graph.DiscardUnknown(m)
}

var xxx_messageInfo_Graph proto.InternalMessageInfo

func (m *Graph) GetEdge() []*Edge {
	if m != nil {
		return m.Edge
	}
	return nil
}

func (m *Graph) GetLatlon() []*LatLon {
	if m != nil {
		return m.Latlon
	}
	return nil
}

// Our address book file is just one of these.
type Edge struct {
	EdgeId               int64    `protobuf:"varint,1,opt,name=edge_id,json=edgeId,proto3" json:"edge_id,omitempty"`
	From                 int64    `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	Cost                 float64  `protobuf:"fixed64,4,opt,name=cost,proto3" json:"cost,omitempty"`
	PointId              []int64  `protobuf:"varint,5,rep,packed,name=point_id,json=pointId,proto3" json:"point_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Edge) Reset()         { *m = Edge{} }
func (m *Edge) String() string { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()    {}
func (*Edge) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebeae2f27825ff85, []int{1}
}

func (m *Edge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Edge.Unmarshal(m, b)
}
func (m *Edge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Edge.Marshal(b, m, deterministic)
}
func (m *Edge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Edge.Merge(m, src)
}
func (m *Edge) XXX_Size() int {
	return xxx_messageInfo_Edge.Size(m)
}
func (m *Edge) XXX_DiscardUnknown() {
	xxx_messageInfo_Edge.DiscardUnknown(m)
}

var xxx_messageInfo_Edge proto.InternalMessageInfo

func (m *Edge) GetEdgeId() int64 {
	if m != nil {
		return m.EdgeId
	}
	return 0
}

func (m *Edge) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Edge) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Edge) GetCost() float64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *Edge) GetPointId() []int64 {
	if m != nil {
		return m.PointId
	}
	return nil
}

type LatLon struct {
	LatlonId             int64    `protobuf:"varint,1,opt,name=latlon_id,json=latlonId,proto3" json:"latlon_id,omitempty"`
	Lat                  float64  `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float64  `protobuf:"fixed64,3,opt,name=lon,proto3" json:"lon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LatLon) Reset()         { *m = LatLon{} }
func (m *LatLon) String() string { return proto.CompactTextString(m) }
func (*LatLon) ProtoMessage()    {}
func (*LatLon) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebeae2f27825ff85, []int{2}
}

func (m *LatLon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatLon.Unmarshal(m, b)
}
func (m *LatLon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatLon.Marshal(b, m, deterministic)
}
func (m *LatLon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatLon.Merge(m, src)
}
func (m *LatLon) XXX_Size() int {
	return xxx_messageInfo_LatLon.Size(m)
}
func (m *LatLon) XXX_DiscardUnknown() {
	xxx_messageInfo_LatLon.DiscardUnknown(m)
}

var xxx_messageInfo_LatLon proto.InternalMessageInfo

func (m *LatLon) GetLatlonId() int64 {
	if m != nil {
		return m.LatlonId
	}
	return 0
}

func (m *LatLon) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *LatLon) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func init() {
	proto.RegisterType((*Graph)(nil), "pb.Graph")
	proto.RegisterType((*Edge)(nil), "pb.Edge")
	proto.RegisterType((*LatLon)(nil), "pb.LatLon")
}

func init() { proto.RegisterFile("goraph.proto", fileDescriptor_ebeae2f27825ff85) }

var fileDescriptor_ebeae2f27825ff85 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4a, 0xc6, 0x30,
	0x10, 0x84, 0x49, 0xd2, 0xbf, 0x7f, 0x5d, 0x45, 0x64, 0x2f, 0x46, 0xf4, 0x50, 0x7a, 0xea, 0xa9,
	0x07, 0x7d, 0x06, 0x91, 0xc0, 0x7f, 0xca, 0x0b, 0x48, 0x6b, 0x6a, 0x2d, 0xd4, 0x6c, 0xa8, 0xfb,
	0xfe, 0xc8, 0x6e, 0x05, 0x6f, 0xb3, 0xdf, 0x64, 0x33, 0xc3, 0xc2, 0xcd, 0x42, 0xfb, 0x58, 0xbe,
	0x86, 0xb2, 0x13, 0x13, 0xda, 0x32, 0x75, 0x01, 0x4e, 0x6f, 0x82, 0xf0, 0x09, 0xaa, 0x39, 0x2d,
	0xb3, 0x37, 0xad, 0xeb, 0xaf, 0x9f, 0x9b, 0xa1, 0x4c, 0xc3, 0x6b, 0x5a, 0xe6, 0xa8, 0x14, 0x3b,
	0xa8, 0xb7, 0x91, 0x37, 0xca, 0xde, 0xaa, 0x0f, 0xe2, 0x5f, 0x46, 0xbe, 0x50, 0x8e, 0x7f, 0x4e,
	0xb7, 0x43, 0x25, 0x1b, 0x78, 0x0f, 0x67, 0xd9, 0x79, 0x5f, 0x93, 0x37, 0xad, 0xe9, 0x5d, 0xac,
	0x65, 0x0c, 0x09, 0x11, 0xaa, 0xcf, 0x9d, 0xbe, 0xbd, 0x55, 0xaa, 0x1a, 0x6f, 0xc1, 0x32, 0x79,
	0xa7, 0xc4, 0x32, 0xc9, 0x9b, 0x0f, 0xfa, 0x61, 0x5f, 0xb5, 0xa6, 0x37, 0x51, 0x35, 0x3e, 0x40,
	0x53, 0x68, 0xcd, 0x2c, 0x3f, 0x9e, 0x5a, 0xd7, 0xbb, 0x78, 0xd6, 0x39, 0xa4, 0x2e, 0x40, 0x7d,
	0xb4, 0xc0, 0x47, 0xb8, 0x3a, 0x7a, 0xfc, 0xe7, 0x36, 0x07, 0x08, 0x09, 0xef, 0xc0, 0x6d, 0x23,
	0x6b, 0xb0, 0x89, 0x22, 0x95, 0x50, 0xd6, 0x60, 0x21, 0x94, 0xa7, 0x5a, 0x8f, 0xf2, 0xf2, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x83, 0x89, 0x15, 0xc0, 0x24, 0x01, 0x00, 0x00,
}
