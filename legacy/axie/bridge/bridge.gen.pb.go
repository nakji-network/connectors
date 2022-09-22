// Code generated - DO NOT EDIT.
// This file is a generated protobuf definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: connectors/source/axie/bridge/bridge.gen.proto

package bridge

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TokenDeposited struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts               *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	DepositId        []byte                 `protobuf:"bytes,2,opt,name=depositId,proto3" json:"depositId,omitempty"` // uint256
	Owner            []byte                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	TokenAddress     []byte                 `protobuf:"bytes,4,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	SidechainAddress []byte                 `protobuf:"bytes,5,opt,name=sidechainAddress,proto3" json:"sidechainAddress,omitempty"`
	Standard         uint32                 `protobuf:"varint,6,opt,name=standard,proto3" json:"standard,omitempty"`
	TokenNumber      []byte                 `protobuf:"bytes,7,opt,name=tokenNumber,proto3" json:"tokenNumber,omitempty"` // uint256
}

func (x *TokenDeposited) Reset() {
	*x = TokenDeposited{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenDeposited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenDeposited) ProtoMessage() {}

func (x *TokenDeposited) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenDeposited.ProtoReflect.Descriptor instead.
func (*TokenDeposited) Descriptor() ([]byte, []int) {
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP(), []int{0}
}

func (x *TokenDeposited) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *TokenDeposited) GetDepositId() []byte {
	if x != nil {
		return x.DepositId
	}
	return nil
}

func (x *TokenDeposited) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *TokenDeposited) GetTokenAddress() []byte {
	if x != nil {
		return x.TokenAddress
	}
	return nil
}

func (x *TokenDeposited) GetSidechainAddress() []byte {
	if x != nil {
		return x.SidechainAddress
	}
	return nil
}

func (x *TokenDeposited) GetStandard() uint32 {
	if x != nil {
		return x.Standard
	}
	return 0
}

func (x *TokenDeposited) GetTokenNumber() []byte {
	if x != nil {
		return x.TokenNumber
	}
	return nil
}

type TokenWithdrew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	WithdrawId   []byte                 `protobuf:"bytes,2,opt,name=withdrawId,proto3" json:"withdrawId,omitempty"` // uint256
	Owner        []byte                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	TokenAddress []byte                 `protobuf:"bytes,4,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	TokenNumber  []byte                 `protobuf:"bytes,5,opt,name=tokenNumber,proto3" json:"tokenNumber,omitempty"` // uint256
}

func (x *TokenWithdrew) Reset() {
	*x = TokenWithdrew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenWithdrew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenWithdrew) ProtoMessage() {}

func (x *TokenWithdrew) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenWithdrew.ProtoReflect.Descriptor instead.
func (*TokenWithdrew) Descriptor() ([]byte, []int) {
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP(), []int{1}
}

func (x *TokenWithdrew) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *TokenWithdrew) GetWithdrawId() []byte {
	if x != nil {
		return x.WithdrawId
	}
	return nil
}

func (x *TokenWithdrew) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *TokenWithdrew) GetTokenAddress() []byte {
	if x != nil {
		return x.TokenAddress
	}
	return nil
}

func (x *TokenWithdrew) GetTokenNumber() []byte {
	if x != nil {
		return x.TokenNumber
	}
	return nil
}

type Unpaused struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Unpaused) Reset() {
	*x = Unpaused{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unpaused) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unpaused) ProtoMessage() {}

func (x *Unpaused) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unpaused.ProtoReflect.Descriptor instead.
func (*Unpaused) Descriptor() ([]byte, []int) {
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP(), []int{2}
}

func (x *Unpaused) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type AdminChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	OldAdmin []byte                 `protobuf:"bytes,2,opt,name=oldAdmin,proto3" json:"oldAdmin,omitempty"`
	NewAdmin []byte                 `protobuf:"bytes,3,opt,name=newAdmin,proto3" json:"newAdmin,omitempty"`
}

func (x *AdminChanged) Reset() {
	*x = AdminChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminChanged) ProtoMessage() {}

func (x *AdminChanged) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminChanged.ProtoReflect.Descriptor instead.
func (*AdminChanged) Descriptor() ([]byte, []int) {
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP(), []int{3}
}

func (x *AdminChanged) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *AdminChanged) GetOldAdmin() []byte {
	if x != nil {
		return x.OldAdmin
	}
	return nil
}

func (x *AdminChanged) GetNewAdmin() []byte {
	if x != nil {
		return x.NewAdmin
	}
	return nil
}

type AdminRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	OldAdmin []byte                 `protobuf:"bytes,2,opt,name=oldAdmin,proto3" json:"oldAdmin,omitempty"`
}

func (x *AdminRemoved) Reset() {
	*x = AdminRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRemoved) ProtoMessage() {}

func (x *AdminRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRemoved.ProtoReflect.Descriptor instead.
func (*AdminRemoved) Descriptor() ([]byte, []int) {
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP(), []int{4}
}

func (x *AdminRemoved) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *AdminRemoved) GetOldAdmin() []byte {
	if x != nil {
		return x.OldAdmin
	}
	return nil
}

type Paused struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Paused) Reset() {
	*x = Paused{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paused) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paused) ProtoMessage() {}

func (x *Paused) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paused.ProtoReflect.Descriptor instead.
func (*Paused) Descriptor() ([]byte, []int) {
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP(), []int{5}
}

func (x *Paused) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type ProxyUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	New []byte                 `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
	Old []byte                 `protobuf:"bytes,3,opt,name=old,proto3" json:"old,omitempty"`
}

func (x *ProxyUpdated) Reset() {
	*x = ProxyUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyUpdated) ProtoMessage() {}

func (x *ProxyUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyUpdated.ProtoReflect.Descriptor instead.
func (*ProxyUpdated) Descriptor() ([]byte, []int) {
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP(), []int{6}
}

func (x *ProxyUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ProxyUpdated) GetNew() []byte {
	if x != nil {
		return x.New
	}
	return nil
}

func (x *ProxyUpdated) GetOld() []byte {
	if x != nil {
		return x.Old
	}
	return nil
}

var File_connectors_source_axie_bridge_bridge_gen_proto protoreflect.FileDescriptor

var file_connectors_source_axie_bridge_bridge_gen_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x61, 0x78, 0x69, 0x65, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x69, 0x64, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x73, 0x69, 0x64, 0x65,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x65, 0x77, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x77, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x08, 0x55, 0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0x72, 0x0a, 0x0c,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x56, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x6c, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x6f, 0x6c, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x34, 0x0a, 0x06, 0x50, 0x61, 0x75, 0x73,
	0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0x5e,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x42, 0x20,
	0x5a, 0x1e, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectors_source_axie_bridge_bridge_gen_proto_rawDescOnce sync.Once
	file_connectors_source_axie_bridge_bridge_gen_proto_rawDescData = file_connectors_source_axie_bridge_bridge_gen_proto_rawDesc
)

func file_connectors_source_axie_bridge_bridge_gen_proto_rawDescGZIP() []byte {
	file_connectors_source_axie_bridge_bridge_gen_proto_rawDescOnce.Do(func() {
		file_connectors_source_axie_bridge_bridge_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectors_source_axie_bridge_bridge_gen_proto_rawDescData)
	})
	return file_connectors_source_axie_bridge_bridge_gen_proto_rawDescData
}

var file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_connectors_source_axie_bridge_bridge_gen_proto_goTypes = []interface{}{
	(*TokenDeposited)(nil),        // 0: bridge.TokenDeposited
	(*TokenWithdrew)(nil),         // 1: bridge.TokenWithdrew
	(*Unpaused)(nil),              // 2: bridge.Unpaused
	(*AdminChanged)(nil),          // 3: bridge.AdminChanged
	(*AdminRemoved)(nil),          // 4: bridge.AdminRemoved
	(*Paused)(nil),                // 5: bridge.Paused
	(*ProxyUpdated)(nil),          // 6: bridge.ProxyUpdated
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_connectors_source_axie_bridge_bridge_gen_proto_depIdxs = []int32{
	7, // 0: bridge.TokenDeposited.ts:type_name -> google.protobuf.Timestamp
	7, // 1: bridge.TokenWithdrew.ts:type_name -> google.protobuf.Timestamp
	7, // 2: bridge.Unpaused.ts:type_name -> google.protobuf.Timestamp
	7, // 3: bridge.AdminChanged.ts:type_name -> google.protobuf.Timestamp
	7, // 4: bridge.AdminRemoved.ts:type_name -> google.protobuf.Timestamp
	7, // 5: bridge.Paused.ts:type_name -> google.protobuf.Timestamp
	7, // 6: bridge.ProxyUpdated.ts:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_connectors_source_axie_bridge_bridge_gen_proto_init() }
func file_connectors_source_axie_bridge_bridge_gen_proto_init() {
	if File_connectors_source_axie_bridge_bridge_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenDeposited); i {
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
		file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenWithdrew); i {
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
		file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unpaused); i {
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
		file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminChanged); i {
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
		file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRemoved); i {
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
		file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paused); i {
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
		file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyUpdated); i {
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
			RawDescriptor: file_connectors_source_axie_bridge_bridge_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectors_source_axie_bridge_bridge_gen_proto_goTypes,
		DependencyIndexes: file_connectors_source_axie_bridge_bridge_gen_proto_depIdxs,
		MessageInfos:      file_connectors_source_axie_bridge_bridge_gen_proto_msgTypes,
	}.Build()
	File_connectors_source_axie_bridge_bridge_gen_proto = out.File
	file_connectors_source_axie_bridge_bridge_gen_proto_rawDesc = nil
	file_connectors_source_axie_bridge_bridge_gen_proto_goTypes = nil
	file_connectors_source_axie_bridge_bridge_gen_proto_depIdxs = nil
}
