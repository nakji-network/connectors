// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: Changelog/Changelog.proto

package Changelog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rely struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Usr []byte                 `protobuf:"bytes,2,opt,name=Usr,proto3" json:"Usr,omitempty"` //	address
}

func (x *Rely) Reset() {
	*x = Rely{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Changelog_Changelog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rely) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rely) ProtoMessage() {}

func (x *Rely) ProtoReflect() protoreflect.Message {
	mi := &file_Changelog_Changelog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rely.ProtoReflect.Descriptor instead.
func (*Rely) Descriptor() ([]byte, []int) {
	return file_Changelog_Changelog_proto_rawDescGZIP(), []int{0}
}

func (x *Rely) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Rely) GetUsr() []byte {
	if x != nil {
		return x.Usr
	}
	return nil
}

type RemoveAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Key []byte                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"` //	bytes32
}

func (x *RemoveAddress) Reset() {
	*x = RemoveAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Changelog_Changelog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAddress) ProtoMessage() {}

func (x *RemoveAddress) ProtoReflect() protoreflect.Message {
	mi := &file_Changelog_Changelog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAddress.ProtoReflect.Descriptor instead.
func (*RemoveAddress) Descriptor() ([]byte, []int) {
	return file_Changelog_Changelog_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveAddress) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *RemoveAddress) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type UpdateAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Key  []byte                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`   //	bytes32
	Addr []byte                 `protobuf:"bytes,3,opt,name=Addr,proto3" json:"Addr,omitempty"` //	address
}

func (x *UpdateAddress) Reset() {
	*x = UpdateAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Changelog_Changelog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAddress) ProtoMessage() {}

func (x *UpdateAddress) ProtoReflect() protoreflect.Message {
	mi := &file_Changelog_Changelog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAddress.ProtoReflect.Descriptor instead.
func (*UpdateAddress) Descriptor() ([]byte, []int) {
	return file_Changelog_Changelog_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAddress) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *UpdateAddress) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *UpdateAddress) GetAddr() []byte {
	if x != nil {
		return x.Addr
	}
	return nil
}

type UpdateIPFS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Ipfs string                 `protobuf:"bytes,2,opt,name=Ipfs,proto3" json:"Ipfs,omitempty"` //	string
}

func (x *UpdateIPFS) Reset() {
	*x = UpdateIPFS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Changelog_Changelog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIPFS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIPFS) ProtoMessage() {}

func (x *UpdateIPFS) ProtoReflect() protoreflect.Message {
	mi := &file_Changelog_Changelog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIPFS.ProtoReflect.Descriptor instead.
func (*UpdateIPFS) Descriptor() ([]byte, []int) {
	return file_Changelog_Changelog_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateIPFS) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *UpdateIPFS) GetIpfs() string {
	if x != nil {
		return x.Ipfs
	}
	return ""
}

type UpdateSha256Sum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Sha256Sum string                 `protobuf:"bytes,2,opt,name=Sha256sum,proto3" json:"Sha256sum,omitempty"` //	string
}

func (x *UpdateSha256Sum) Reset() {
	*x = UpdateSha256Sum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Changelog_Changelog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSha256Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSha256Sum) ProtoMessage() {}

func (x *UpdateSha256Sum) ProtoReflect() protoreflect.Message {
	mi := &file_Changelog_Changelog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSha256Sum.ProtoReflect.Descriptor instead.
func (*UpdateSha256Sum) Descriptor() ([]byte, []int) {
	return file_Changelog_Changelog_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSha256Sum) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *UpdateSha256Sum) GetSha256Sum() string {
	if x != nil {
		return x.Sha256Sum
	}
	return ""
}

type UpdateVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Version string                 `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"` //	string
}

func (x *UpdateVersion) Reset() {
	*x = UpdateVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Changelog_Changelog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVersion) ProtoMessage() {}

func (x *UpdateVersion) ProtoReflect() protoreflect.Message {
	mi := &file_Changelog_Changelog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVersion.ProtoReflect.Descriptor instead.
func (*UpdateVersion) Descriptor() ([]byte, []int) {
	return file_Changelog_Changelog_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateVersion) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *UpdateVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Deny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Usr []byte                 `protobuf:"bytes,2,opt,name=Usr,proto3" json:"Usr,omitempty"` //	address
}

func (x *Deny) Reset() {
	*x = Deny{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Changelog_Changelog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deny) ProtoMessage() {}

func (x *Deny) ProtoReflect() protoreflect.Message {
	mi := &file_Changelog_Changelog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deny.ProtoReflect.Descriptor instead.
func (*Deny) Descriptor() ([]byte, []int) {
	return file_Changelog_Changelog_proto_rawDescGZIP(), []int{6}
}

func (x *Deny) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Deny) GetUsr() []byte {
	if x != nil {
		return x.Usr
	}
	return nil
}

var File_Changelog_Changelog_proto protoreflect.FileDescriptor

var file_Changelog_Changelog_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x04, 0x52, 0x65, 0x6c, 0x79, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x55, 0x73, 0x72, 0x22, 0x4d, 0x0a,
	0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x61, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x22,
	0x4c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x50, 0x46, 0x53, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x70, 0x66,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x70, 0x66, 0x73, 0x22, 0x5b, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x53, 0x75, 0x6d,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x73, 0x75, 0x6d, 0x22, 0x55, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x44, 0x0a, 0x04, 0x44, 0x65, 0x6e, 0x79, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x55, 0x73, 0x72, 0x42, 0x3c, 0x5a, 0x3a, 0x62, 0x6c, 0x65, 0x70, 0x2e,
	0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Changelog_Changelog_proto_rawDescOnce sync.Once
	file_Changelog_Changelog_proto_rawDescData = file_Changelog_Changelog_proto_rawDesc
)

func file_Changelog_Changelog_proto_rawDescGZIP() []byte {
	file_Changelog_Changelog_proto_rawDescOnce.Do(func() {
		file_Changelog_Changelog_proto_rawDescData = protoimpl.X.CompressGZIP(file_Changelog_Changelog_proto_rawDescData)
	})
	return file_Changelog_Changelog_proto_rawDescData
}

var file_Changelog_Changelog_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Changelog_Changelog_proto_goTypes = []interface{}{
	(*Rely)(nil),                  // 0: Changelog.Rely
	(*RemoveAddress)(nil),         // 1: Changelog.RemoveAddress
	(*UpdateAddress)(nil),         // 2: Changelog.UpdateAddress
	(*UpdateIPFS)(nil),            // 3: Changelog.UpdateIPFS
	(*UpdateSha256Sum)(nil),       // 4: Changelog.UpdateSha256Sum
	(*UpdateVersion)(nil),         // 5: Changelog.UpdateVersion
	(*Deny)(nil),                  // 6: Changelog.Deny
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_Changelog_Changelog_proto_depIdxs = []int32{
	7, // 0: Changelog.Rely.ts:type_name -> google.protobuf.Timestamp
	7, // 1: Changelog.RemoveAddress.ts:type_name -> google.protobuf.Timestamp
	7, // 2: Changelog.UpdateAddress.ts:type_name -> google.protobuf.Timestamp
	7, // 3: Changelog.UpdateIPFS.ts:type_name -> google.protobuf.Timestamp
	7, // 4: Changelog.UpdateSha256Sum.ts:type_name -> google.protobuf.Timestamp
	7, // 5: Changelog.UpdateVersion.ts:type_name -> google.protobuf.Timestamp
	7, // 6: Changelog.Deny.ts:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Changelog_Changelog_proto_init() }
func file_Changelog_Changelog_proto_init() {
	if File_Changelog_Changelog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Changelog_Changelog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rely); i {
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
		file_Changelog_Changelog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAddress); i {
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
		file_Changelog_Changelog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAddress); i {
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
		file_Changelog_Changelog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIPFS); i {
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
		file_Changelog_Changelog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSha256Sum); i {
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
		file_Changelog_Changelog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVersion); i {
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
		file_Changelog_Changelog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deny); i {
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
			RawDescriptor: file_Changelog_Changelog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Changelog_Changelog_proto_goTypes,
		DependencyIndexes: file_Changelog_Changelog_proto_depIdxs,
		MessageInfos:      file_Changelog_Changelog_proto_msgTypes,
	}.Build()
	File_Changelog_Changelog_proto = out.File
	file_Changelog_Changelog_proto_rawDesc = nil
	file_Changelog_Changelog_proto_goTypes = nil
	file_Changelog_Changelog_proto_depIdxs = nil
}
