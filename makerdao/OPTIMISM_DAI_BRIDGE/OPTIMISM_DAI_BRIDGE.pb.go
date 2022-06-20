// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: OPTIMISM_DAI_BRIDGE.proto

package OPTIMISM_DAI_BRIDGE

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

type ERC20WithdrawalFinalized struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	L1Token []byte                 `protobuf:"bytes,2,opt,name=L1Token,proto3" json:"L1Token,omitempty"` //	address
	L2Token []byte                 `protobuf:"bytes,3,opt,name=L2Token,proto3" json:"L2Token,omitempty"` //	address
	From    []byte                 `protobuf:"bytes,4,opt,name=From,proto3" json:"From,omitempty"`       //	address
	To      []byte                 `protobuf:"bytes,5,opt,name=To,proto3" json:"To,omitempty"`           //	address
	Amount  []byte                 `protobuf:"bytes,6,opt,name=Amount,proto3" json:"Amount,omitempty"`   //	uint256
	Data    []byte                 `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`       //	bytes
}

func (x *ERC20WithdrawalFinalized) Reset() {
	*x = ERC20WithdrawalFinalized{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ERC20WithdrawalFinalized) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ERC20WithdrawalFinalized) ProtoMessage() {}

func (x *ERC20WithdrawalFinalized) ProtoReflect() protoreflect.Message {
	mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ERC20WithdrawalFinalized.ProtoReflect.Descriptor instead.
func (*ERC20WithdrawalFinalized) Descriptor() ([]byte, []int) {
	return file_OPTIMISM_DAI_BRIDGE_proto_rawDescGZIP(), []int{0}
}

func (x *ERC20WithdrawalFinalized) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ERC20WithdrawalFinalized) GetL1Token() []byte {
	if x != nil {
		return x.L1Token
	}
	return nil
}

func (x *ERC20WithdrawalFinalized) GetL2Token() []byte {
	if x != nil {
		return x.L2Token
	}
	return nil
}

func (x *ERC20WithdrawalFinalized) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ERC20WithdrawalFinalized) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *ERC20WithdrawalFinalized) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *ERC20WithdrawalFinalized) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

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
		mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rely) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rely) ProtoMessage() {}

func (x *Rely) ProtoReflect() protoreflect.Message {
	mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[1]
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
	return file_OPTIMISM_DAI_BRIDGE_proto_rawDescGZIP(), []int{1}
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

type Closed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Closed) Reset() {
	*x = Closed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Closed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Closed) ProtoMessage() {}

func (x *Closed) ProtoReflect() protoreflect.Message {
	mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Closed.ProtoReflect.Descriptor instead.
func (*Closed) Descriptor() ([]byte, []int) {
	return file_OPTIMISM_DAI_BRIDGE_proto_rawDescGZIP(), []int{2}
}

func (x *Closed) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
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
		mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deny) ProtoMessage() {}

func (x *Deny) ProtoReflect() protoreflect.Message {
	mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[3]
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
	return file_OPTIMISM_DAI_BRIDGE_proto_rawDescGZIP(), []int{3}
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

type ERC20DepositInitiated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	L1Token []byte                 `protobuf:"bytes,2,opt,name=L1Token,proto3" json:"L1Token,omitempty"` //	address
	L2Token []byte                 `protobuf:"bytes,3,opt,name=L2Token,proto3" json:"L2Token,omitempty"` //	address
	From    []byte                 `protobuf:"bytes,4,opt,name=From,proto3" json:"From,omitempty"`       //	address
	To      []byte                 `protobuf:"bytes,5,opt,name=To,proto3" json:"To,omitempty"`           //	address
	Amount  []byte                 `protobuf:"bytes,6,opt,name=Amount,proto3" json:"Amount,omitempty"`   //	uint256
	Data    []byte                 `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`       //	bytes
}

func (x *ERC20DepositInitiated) Reset() {
	*x = ERC20DepositInitiated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ERC20DepositInitiated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ERC20DepositInitiated) ProtoMessage() {}

func (x *ERC20DepositInitiated) ProtoReflect() protoreflect.Message {
	mi := &file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ERC20DepositInitiated.ProtoReflect.Descriptor instead.
func (*ERC20DepositInitiated) Descriptor() ([]byte, []int) {
	return file_OPTIMISM_DAI_BRIDGE_proto_rawDescGZIP(), []int{4}
}

func (x *ERC20DepositInitiated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ERC20DepositInitiated) GetL1Token() []byte {
	if x != nil {
		return x.L1Token
	}
	return nil
}

func (x *ERC20DepositInitiated) GetL2Token() []byte {
	if x != nil {
		return x.L2Token
	}
	return nil
}

func (x *ERC20DepositInitiated) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ERC20DepositInitiated) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *ERC20DepositInitiated) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *ERC20DepositInitiated) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_OPTIMISM_DAI_BRIDGE_proto protoreflect.FileDescriptor

var file_OPTIMISM_DAI_BRIDGE_proto_rawDesc = []byte{
	0x0a, 0x19, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x4d, 0x5f, 0x44, 0x41, 0x49, 0x5f, 0x42,
	0x52, 0x49, 0x44, 0x47, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x4f, 0x50, 0x54,
	0x49, 0x4d, 0x49, 0x53, 0x4d, 0x5f, 0x44, 0x41, 0x49, 0x5f, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xca, 0x01, 0x0a, 0x18, 0x45, 0x52, 0x43, 0x32, 0x30, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x31,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4c, 0x31, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4c, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x44,
	0x0a, 0x04, 0x52, 0x65, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02,
	0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x55, 0x73, 0x72, 0x22, 0x34, 0x0a, 0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0x44, 0x0a, 0x04, 0x44, 0x65,
	0x6e, 0x79, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x55, 0x73, 0x72,
	0x22, 0xc7, 0x01, 0x0a, 0x15, 0x45, 0x52, 0x43, 0x32, 0x30, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x31, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4c, 0x31, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x4c, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x4c, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72,
	0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x46, 0x5a, 0x44, 0x62, 0x6c,
	0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x4f,
	0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x4d, 0x5f, 0x44, 0x41, 0x49, 0x5f, 0x42, 0x52, 0x49, 0x44,
	0x47, 0x45, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OPTIMISM_DAI_BRIDGE_proto_rawDescOnce sync.Once
	file_OPTIMISM_DAI_BRIDGE_proto_rawDescData = file_OPTIMISM_DAI_BRIDGE_proto_rawDesc
)

func file_OPTIMISM_DAI_BRIDGE_proto_rawDescGZIP() []byte {
	file_OPTIMISM_DAI_BRIDGE_proto_rawDescOnce.Do(func() {
		file_OPTIMISM_DAI_BRIDGE_proto_rawDescData = protoimpl.X.CompressGZIP(file_OPTIMISM_DAI_BRIDGE_proto_rawDescData)
	})
	return file_OPTIMISM_DAI_BRIDGE_proto_rawDescData
}

var file_OPTIMISM_DAI_BRIDGE_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_OPTIMISM_DAI_BRIDGE_proto_goTypes = []interface{}{
	(*ERC20WithdrawalFinalized)(nil), // 0: OPTIMISM_DAI_BRIDGE.ERC20WithdrawalFinalized
	(*Rely)(nil),                     // 1: OPTIMISM_DAI_BRIDGE.Rely
	(*Closed)(nil),                   // 2: OPTIMISM_DAI_BRIDGE.Closed
	(*Deny)(nil),                     // 3: OPTIMISM_DAI_BRIDGE.Deny
	(*ERC20DepositInitiated)(nil),    // 4: OPTIMISM_DAI_BRIDGE.ERC20DepositInitiated
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
}
var file_OPTIMISM_DAI_BRIDGE_proto_depIdxs = []int32{
	5, // 0: OPTIMISM_DAI_BRIDGE.ERC20WithdrawalFinalized.ts:type_name -> google.protobuf.Timestamp
	5, // 1: OPTIMISM_DAI_BRIDGE.Rely.ts:type_name -> google.protobuf.Timestamp
	5, // 2: OPTIMISM_DAI_BRIDGE.Closed.ts:type_name -> google.protobuf.Timestamp
	5, // 3: OPTIMISM_DAI_BRIDGE.Deny.ts:type_name -> google.protobuf.Timestamp
	5, // 4: OPTIMISM_DAI_BRIDGE.ERC20DepositInitiated.ts:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_OPTIMISM_DAI_BRIDGE_proto_init() }
func file_OPTIMISM_DAI_BRIDGE_proto_init() {
	if File_OPTIMISM_DAI_BRIDGE_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ERC20WithdrawalFinalized); i {
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
		file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Closed); i {
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
		file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_OPTIMISM_DAI_BRIDGE_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ERC20DepositInitiated); i {
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
			RawDescriptor: file_OPTIMISM_DAI_BRIDGE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OPTIMISM_DAI_BRIDGE_proto_goTypes,
		DependencyIndexes: file_OPTIMISM_DAI_BRIDGE_proto_depIdxs,
		MessageInfos:      file_OPTIMISM_DAI_BRIDGE_proto_msgTypes,
	}.Build()
	File_OPTIMISM_DAI_BRIDGE_proto = out.File
	file_OPTIMISM_DAI_BRIDGE_proto_rawDesc = nil
	file_OPTIMISM_DAI_BRIDGE_proto_goTypes = nil
	file_OPTIMISM_DAI_BRIDGE_proto_depIdxs = nil
}