// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: connectors/source/perpetual/smart-contracts/L2/InsuranceFund/InsuranceFund.proto

package InsuranceFund

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

type OwnershipTransferred struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	PreviousOwner []byte                 `protobuf:"bytes,2,opt,name=PreviousOwner,proto3" json:"PreviousOwner,omitempty"` //	address
	NewOwner      []byte                 `protobuf:"bytes,3,opt,name=NewOwner,proto3" json:"NewOwner,omitempty"`           //	address
}

func (x *OwnershipTransferred) Reset() {
	*x = OwnershipTransferred{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnershipTransferred) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnershipTransferred) ProtoMessage() {}

func (x *OwnershipTransferred) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnershipTransferred.ProtoReflect.Descriptor instead.
func (*OwnershipTransferred) Descriptor() ([]byte, []int) {
	return file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescGZIP(), []int{0}
}

func (x *OwnershipTransferred) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *OwnershipTransferred) GetPreviousOwner() []byte {
	if x != nil {
		return x.PreviousOwner
	}
	return nil
}

func (x *OwnershipTransferred) GetNewOwner() []byte {
	if x != nil {
		return x.NewOwner
	}
	return nil
}

type ShutdownAllAmms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber []byte                 `protobuf:"bytes,2,opt,name=BlockNumber,proto3" json:"BlockNumber,omitempty"` //	uint256
}

func (x *ShutdownAllAmms) Reset() {
	*x = ShutdownAllAmms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownAllAmms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownAllAmms) ProtoMessage() {}

func (x *ShutdownAllAmms) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownAllAmms.ProtoReflect.Descriptor instead.
func (*ShutdownAllAmms) Descriptor() ([]byte, []int) {
	return file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescGZIP(), []int{1}
}

func (x *ShutdownAllAmms) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ShutdownAllAmms) GetBlockNumber() []byte {
	if x != nil {
		return x.BlockNumber
	}
	return nil
}

type TokenAdded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	TokenAddress []byte                 `protobuf:"bytes,2,opt,name=TokenAddress,proto3" json:"TokenAddress,omitempty"` //	address
}

func (x *TokenAdded) Reset() {
	*x = TokenAdded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenAdded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenAdded) ProtoMessage() {}

func (x *TokenAdded) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenAdded.ProtoReflect.Descriptor instead.
func (*TokenAdded) Descriptor() ([]byte, []int) {
	return file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescGZIP(), []int{2}
}

func (x *TokenAdded) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *TokenAdded) GetTokenAddress() []byte {
	if x != nil {
		return x.TokenAddress
	}
	return nil
}

type TokenRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	TokenAddress []byte                 `protobuf:"bytes,2,opt,name=TokenAddress,proto3" json:"TokenAddress,omitempty"` //	address
}

func (x *TokenRemoved) Reset() {
	*x = TokenRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRemoved) ProtoMessage() {}

func (x *TokenRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRemoved.ProtoReflect.Descriptor instead.
func (*TokenRemoved) Descriptor() ([]byte, []int) {
	return file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescGZIP(), []int{3}
}

func (x *TokenRemoved) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *TokenRemoved) GetTokenAddress() []byte {
	if x != nil {
		return x.TokenAddress
	}
	return nil
}

type Withdrawn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Withdrawer []byte                 `protobuf:"bytes,2,opt,name=Withdrawer,proto3" json:"Withdrawer,omitempty"` //	address
	Amount     []byte                 `protobuf:"bytes,3,opt,name=Amount,proto3" json:"Amount,omitempty"`         //	uint256
}

func (x *Withdrawn) Reset() {
	*x = Withdrawn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Withdrawn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdrawn) ProtoMessage() {}

func (x *Withdrawn) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdrawn.ProtoReflect.Descriptor instead.
func (*Withdrawn) Descriptor() ([]byte, []int) {
	return file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescGZIP(), []int{4}
}

func (x *Withdrawn) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Withdrawn) GetWithdrawer() []byte {
	if x != nil {
		return x.Withdrawer
	}
	return nil
}

func (x *Withdrawn) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

var File_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto protoreflect.FileDescriptor

var file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDesc = []byte{
	0x0a, 0x50, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x70, 0x65, 0x74, 0x75, 0x61, 0x6c, 0x2f, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x4c, 0x32,
	0x2f, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x2f, 0x49,
	0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6e,
	0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x4e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x0f, 0x53, 0x68, 0x75,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x41, 0x6c, 0x6c, 0x41, 0x6d, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x0a, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x02, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x02, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6f, 0x0a, 0x09, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x6e, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x62, 0x6c, 0x65,
	0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x70, 0x65, 0x74, 0x75, 0x61, 0x6c, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x4c,
	0x32, 0x2f, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescOnce sync.Once
	file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescData = file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDesc
)

func file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescGZIP() []byte {
	file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescOnce.Do(func() {
		file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescData)
	})
	return file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDescData
}

var file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_goTypes = []interface{}{
	(*OwnershipTransferred)(nil),  // 0: InsuranceFund.OwnershipTransferred
	(*ShutdownAllAmms)(nil),       // 1: InsuranceFund.ShutdownAllAmms
	(*TokenAdded)(nil),            // 2: InsuranceFund.TokenAdded
	(*TokenRemoved)(nil),          // 3: InsuranceFund.TokenRemoved
	(*Withdrawn)(nil),             // 4: InsuranceFund.Withdrawn
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_depIdxs = []int32{
	5, // 0: InsuranceFund.OwnershipTransferred.ts:type_name -> google.protobuf.Timestamp
	5, // 1: InsuranceFund.ShutdownAllAmms.ts:type_name -> google.protobuf.Timestamp
	5, // 2: InsuranceFund.TokenAdded.ts:type_name -> google.protobuf.Timestamp
	5, // 3: InsuranceFund.TokenRemoved.ts:type_name -> google.protobuf.Timestamp
	5, // 4: InsuranceFund.Withdrawn.ts:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_init()
}
func file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_init() {
	if File_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnershipTransferred); i {
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
		file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownAllAmms); i {
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
		file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenAdded); i {
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
		file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRemoved); i {
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
		file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Withdrawn); i {
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
			RawDescriptor: file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_goTypes,
		DependencyIndexes: file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_depIdxs,
		MessageInfos:      file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_msgTypes,
	}.Build()
	File_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto = out.File
	file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_rawDesc = nil
	file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_goTypes = nil
	file_connectors_source_perpetual_smart_contracts_L2_InsuranceFund_InsuranceFund_proto_depIdxs = nil
}
