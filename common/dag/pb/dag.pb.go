// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dag.proto

/*
Package dagpb is a generated protocol buffer package.

It is generated from these files:
	dag.proto

It has these top-level messages:
	Dag
	Node
*/
package dagpb

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Dag struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=Nodes" json:"Nodes,omitempty"`
}

func (m *Dag) Reset()                    { *m = Dag{} }
func (m *Dag) String() string            { return proto.CompactTextString(m) }
func (*Dag) ProtoMessage()               {}
func (*Dag) Descriptor() ([]byte, []int) { return fileDescriptorDag, []int{0} }

func (m *Dag) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Key      string  `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Index    int32   `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Children []int32 `protobuf:"varint,3,rep,packed,name=Children" json:"Children,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorDag, []int{1} }

func (m *Node) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Node) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Node) GetChildren() []int32 {
	if m != nil {
		return m.Children
	}
	return nil
}

func init() {
	proto.RegisterType((*Dag)(nil), "dagpb.Dag")
	proto.RegisterType((*Node)(nil), "dagpb.Node")
}

func init() { proto.RegisterFile("dag.proto", fileDescriptorDag) }

var fileDescriptorDag = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x49, 0x4c, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x49, 0x4c, 0x2f, 0x48, 0x52, 0xd2, 0xe0, 0x62,
	0x76, 0x49, 0x4c, 0x17, 0x52, 0xe4, 0x62, 0xf5, 0xcb, 0x4f, 0x49, 0x2d, 0x96, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x36, 0xe2, 0xd6, 0x03, 0xcb, 0xea, 0x81, 0xc4, 0x82, 0x20, 0x32, 0x4a, 0x5e, 0x5c,
	0x2c, 0x20, 0x86, 0x90, 0x00, 0x17, 0xb3, 0x77, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0xea, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x1a, 0x04, 0xe1, 0x08, 0x49, 0x71, 0x71, 0x38, 0x67, 0x64, 0xe6, 0xa4, 0x14, 0xa5, 0xe6,
	0x49, 0x30, 0x2b, 0x30, 0x6b, 0xb0, 0x06, 0xc1, 0xf9, 0x49, 0x6c, 0x60, 0x37, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xd9, 0x69, 0x2e, 0xcd, 0x90, 0x00, 0x00, 0x00,
}
