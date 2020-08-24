// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/message.proto

package protobuf

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

type DanmuInternalMessage struct {
	Uid                  uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Rid                  uint32   `protobuf:"varint,2,opt,name=rid,proto3" json:"rid,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Color                string   `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DanmuInternalMessage) Reset()         { *m = DanmuInternalMessage{} }
func (m *DanmuInternalMessage) String() string { return proto.CompactTextString(m) }
func (*DanmuInternalMessage) ProtoMessage()    {}
func (*DanmuInternalMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{0}
}

func (m *DanmuInternalMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DanmuInternalMessage.Unmarshal(m, b)
}
func (m *DanmuInternalMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DanmuInternalMessage.Marshal(b, m, deterministic)
}
func (m *DanmuInternalMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DanmuInternalMessage.Merge(m, src)
}
func (m *DanmuInternalMessage) XXX_Size() int {
	return xxx_messageInfo_DanmuInternalMessage.Size(m)
}
func (m *DanmuInternalMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DanmuInternalMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DanmuInternalMessage proto.InternalMessageInfo

func (m *DanmuInternalMessage) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *DanmuInternalMessage) GetRid() uint32 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *DanmuInternalMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *DanmuInternalMessage) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func init() {
	proto.RegisterType((*DanmuInternalMessage)(nil), "protobuf.DanmuInternalMessage")
}

func init() { proto.RegisterFile("protobuf/message.proto", fileDescriptor_8368f5d77b0b9b7b) }

var fileDescriptor_8368f5d77b0b9b7b = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x03, 0x0b, 0x08,
	0x71, 0xc0, 0xc4, 0x95, 0xb2, 0xb8, 0x44, 0x5c, 0x12, 0xf3, 0x72, 0x4b, 0x3d, 0xf3, 0x4a, 0x52,
	0x8b, 0xf2, 0x12, 0x73, 0x7c, 0x21, 0xea, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x33, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x78, 0x83, 0x40, 0x4c, 0x90, 0x48, 0x51, 0x66, 0x8a, 0x04, 0x13, 0x44, 0xa4,
	0x28, 0x33, 0x45, 0x48, 0x82, 0x8b, 0x3d, 0x39, 0x3f, 0xaf, 0x24, 0x35, 0xaf, 0x44, 0x82, 0x59,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x12, 0xe1, 0x62, 0x4d, 0xce, 0xcf, 0xc9, 0x2f, 0x92,
	0x60, 0x01, 0x8b, 0x43, 0x38, 0x49, 0x6c, 0x60, 0x5b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xc8, 0x02, 0xaf, 0x1a, 0x96, 0x00, 0x00, 0x00,
}
