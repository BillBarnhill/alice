// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/getamis/alice/crypto/utils/message.proto

package utils

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Hash struct {
	Msgs                 []*any.Any `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_3146f5d2d896ba86, []int{0}
}

func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (m *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(m, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetMsgs() []*any.Any {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func init() {
	proto.RegisterType((*Hash)(nil), "utils.Hash")
}

func init() {
	proto.RegisterFile("github.com/getamis/alice/crypto/utils/message.proto", fileDescriptor_3146f5d2d896ba86)
}

var fileDescriptor_3146f5d2d896ba86 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x4f, 0x2d, 0x49, 0xcc, 0xcd, 0x2c, 0xd6, 0x4f,
	0xcc, 0xc9, 0x4c, 0x4e, 0xd5, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xd7, 0x2f, 0x2d, 0xc9, 0xcc,
	0x29, 0xd6, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x0b, 0x4a, 0x49, 0xa6, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x05, 0x93, 0x4a,
	0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0x2a, 0x94, 0x0c, 0xb8, 0x58, 0x3c, 0x12, 0x8b, 0x33, 0x84,
	0x34, 0xb8, 0x58, 0x72, 0x8b, 0xd3, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x44, 0xf4,
	0x20, 0x3a, 0xf4, 0x60, 0x3a, 0xf4, 0x1c, 0xf3, 0x2a, 0x83, 0xc0, 0x2a, 0x92, 0xd8, 0xc0, 0x62,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0xbc, 0xdd, 0xd8, 0x91, 0x00, 0x00, 0x00,
}
