// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: protobuf/danmu.proto

package protobuf

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DanmuIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DanmuIdRequest) Reset() {
	*x = DanmuIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuIdRequest) ProtoMessage() {}

func (x *DanmuIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuIdRequest.ProtoReflect.Descriptor instead.
func (*DanmuIdRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{0}
}

func (x *DanmuIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DanmuUIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DanmuUIdRequest) Reset() {
	*x = DanmuUIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuUIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuUIdRequest) ProtoMessage() {}

func (x *DanmuUIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuUIdRequest.ProtoReflect.Descriptor instead.
func (*DanmuUIdRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{1}
}

func (x *DanmuUIdRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type DanmuRoomIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId uint32 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *DanmuRoomIdRequest) Reset() {
	*x = DanmuRoomIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuRoomIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuRoomIdRequest) ProtoMessage() {}

func (x *DanmuRoomIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuRoomIdRequest.ProtoReflect.Descriptor instead.
func (*DanmuRoomIdRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{2}
}

func (x *DanmuRoomIdRequest) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type DanmuUIdAndRoomIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomId uint32 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *DanmuUIdAndRoomIdRequest) Reset() {
	*x = DanmuUIdAndRoomIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuUIdAndRoomIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuUIdAndRoomIdRequest) ProtoMessage() {}

func (x *DanmuUIdAndRoomIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuUIdAndRoomIdRequest.ProtoReflect.Descriptor instead.
func (*DanmuUIdAndRoomIdRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{3}
}

func (x *DanmuUIdAndRoomIdRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DanmuUIdAndRoomIdRequest) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type DanmuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid     uint32 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomId  uint32 `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Visible bool   `protobuf:"varint,4,opt,name=visible,proto3" json:"visible,omitempty"`
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Color   string `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *DanmuRequest) Reset() {
	*x = DanmuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuRequest) ProtoMessage() {}

func (x *DanmuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuRequest.ProtoReflect.Descriptor instead.
func (*DanmuRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{4}
}

func (x *DanmuRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DanmuRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DanmuRequest) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *DanmuRequest) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *DanmuRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *DanmuRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type DanmuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    bool                    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	DanmuList []*DanmuResponse_Result `protobuf:"bytes,2,rep,name=danmu_list,json=danmuList,proto3" json:"danmu_list,omitempty"`
	TimeUsed  int64                   `protobuf:"varint,3,opt,name=time_used,json=timeUsed,proto3" json:"time_used,omitempty"`
}

func (x *DanmuResponse) Reset() {
	*x = DanmuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuResponse) ProtoMessage() {}

func (x *DanmuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuResponse.ProtoReflect.Descriptor instead.
func (*DanmuResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{5}
}

func (x *DanmuResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *DanmuResponse) GetDanmuList() []*DanmuResponse_Result {
	if x != nil {
		return x.DanmuList
	}
	return nil
}

func (x *DanmuResponse) GetTimeUsed() int64 {
	if x != nil {
		return x.TimeUsed
	}
	return 0
}

type DanmuChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DanmuChangeResponse) Reset() {
	*x = DanmuChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuChangeResponse) ProtoMessage() {}

func (x *DanmuChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuChangeResponse.ProtoReflect.Descriptor instead.
func (*DanmuChangeResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{6}
}

func (x *DanmuChangeResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type DanmuResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visible bool   `protobuf:"varint,1,opt,name=visible,proto3" json:"visible,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *DanmuResponse_Result) Reset() {
	*x = DanmuResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_danmu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmuResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmuResponse_Result) ProtoMessage() {}

func (x *DanmuResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_danmu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmuResponse_Result.ProtoReflect.Descriptor instead.
func (*DanmuResponse_Result) Descriptor() ([]byte, []int) {
	return file_protobuf_danmu_proto_rawDescGZIP(), []int{5, 0}
}

func (x *DanmuResponse_Result) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *DanmuResponse_Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_protobuf_danmu_proto protoreflect.FileDescriptor

var file_protobuf_danmu_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x61, 0x6e, 0x6d, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x22, 0x20, 0x0a, 0x0e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x23, 0x0a, 0x0f, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x55, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x18, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x55,
	0x49, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x93, 0x01,
	0x0a, 0x0c, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a,
	0x0a, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x09, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x64, 0x1a, 0x3c, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xfb, 0x02, 0x0a, 0x0c, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x42, 0x79, 0x55, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x55, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x42, 0x79, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x42, 0x79, 0x55, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x55, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x12, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_danmu_proto_rawDescOnce sync.Once
	file_protobuf_danmu_proto_rawDescData = file_protobuf_danmu_proto_rawDesc
)

func file_protobuf_danmu_proto_rawDescGZIP() []byte {
	file_protobuf_danmu_proto_rawDescOnce.Do(func() {
		file_protobuf_danmu_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_danmu_proto_rawDescData)
	})
	return file_protobuf_danmu_proto_rawDescData
}

var file_protobuf_danmu_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_danmu_proto_goTypes = []interface{}{
	(*DanmuIdRequest)(nil),           // 0: protobuf.DanmuIdRequest
	(*DanmuUIdRequest)(nil),          // 1: protobuf.DanmuUIdRequest
	(*DanmuRoomIdRequest)(nil),       // 2: protobuf.DanmuRoomIdRequest
	(*DanmuUIdAndRoomIdRequest)(nil), // 3: protobuf.DanmuUIdAndRoomIdRequest
	(*DanmuRequest)(nil),             // 4: protobuf.DanmuRequest
	(*DanmuResponse)(nil),            // 5: protobuf.DanmuResponse
	(*DanmuChangeResponse)(nil),      // 6: protobuf.DanmuChangeResponse
	(*DanmuResponse_Result)(nil),     // 7: protobuf.DanmuResponse.Result
}
var file_protobuf_danmu_proto_depIdxs = []int32{
	7, // 0: protobuf.DanmuResponse.danmu_list:type_name -> protobuf.DanmuResponse.Result
	0, // 1: protobuf.DanmuService.GetDanmuById:input_type -> protobuf.DanmuIdRequest
	1, // 2: protobuf.DanmuService.GetDanmuByUId:input_type -> protobuf.DanmuUIdRequest
	2, // 3: protobuf.DanmuService.GetDanmuByRoomId:input_type -> protobuf.DanmuRoomIdRequest
	3, // 4: protobuf.DanmuService.GetDanmuByUIdAndRoomId:input_type -> protobuf.DanmuUIdAndRoomIdRequest
	4, // 5: protobuf.DanmuService.AddDanmu:input_type -> protobuf.DanmuRequest
	5, // 6: protobuf.DanmuService.GetDanmuById:output_type -> protobuf.DanmuResponse
	5, // 7: protobuf.DanmuService.GetDanmuByUId:output_type -> protobuf.DanmuResponse
	5, // 8: protobuf.DanmuService.GetDanmuByRoomId:output_type -> protobuf.DanmuResponse
	5, // 9: protobuf.DanmuService.GetDanmuByUIdAndRoomId:output_type -> protobuf.DanmuResponse
	6, // 10: protobuf.DanmuService.AddDanmu:output_type -> protobuf.DanmuChangeResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_danmu_proto_init() }
func file_protobuf_danmu_proto_init() {
	if File_protobuf_danmu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_danmu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_danmu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuUIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_danmu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuRoomIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_danmu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuUIdAndRoomIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_danmu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_danmu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_danmu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuChangeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_danmu_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmuResponse_Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_danmu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_danmu_proto_goTypes,
		DependencyIndexes: file_protobuf_danmu_proto_depIdxs,
		MessageInfos:      file_protobuf_danmu_proto_msgTypes,
	}.Build()
	File_protobuf_danmu_proto = out.File
	file_protobuf_danmu_proto_rawDesc = nil
	file_protobuf_danmu_proto_goTypes = nil
	file_protobuf_danmu_proto_depIdxs = nil
}
