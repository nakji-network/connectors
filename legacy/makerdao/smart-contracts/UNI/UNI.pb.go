// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: connectors/source/makerdao/smart-contracts/UNI/UNI.proto

package UNI

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

type Approval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Owner   []byte                 `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`     //	address
	Spender []byte                 `protobuf:"bytes,3,opt,name=Spender,proto3" json:"Spender,omitempty"` //	address
	Amount  []byte                 `protobuf:"bytes,4,opt,name=Amount,proto3" json:"Amount,omitempty"`   //	uint256
}

func (x *Approval) Reset() {
	*x = Approval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approval) ProtoMessage() {}

func (x *Approval) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Approval.ProtoReflect.Descriptor instead.
func (*Approval) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescGZIP(), []int{0}
}

func (x *Approval) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Approval) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Approval) GetSpender() []byte {
	if x != nil {
		return x.Spender
	}
	return nil
}

func (x *Approval) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

type DelegateChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Delegator    []byte                 `protobuf:"bytes,2,opt,name=Delegator,proto3" json:"Delegator,omitempty"`       //	address
	FromDelegate []byte                 `protobuf:"bytes,3,opt,name=FromDelegate,proto3" json:"FromDelegate,omitempty"` //	address
	ToDelegate   []byte                 `protobuf:"bytes,4,opt,name=ToDelegate,proto3" json:"ToDelegate,omitempty"`     //	address
}

func (x *DelegateChanged) Reset() {
	*x = DelegateChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegateChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegateChanged) ProtoMessage() {}

func (x *DelegateChanged) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegateChanged.ProtoReflect.Descriptor instead.
func (*DelegateChanged) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescGZIP(), []int{1}
}

func (x *DelegateChanged) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *DelegateChanged) GetDelegator() []byte {
	if x != nil {
		return x.Delegator
	}
	return nil
}

func (x *DelegateChanged) GetFromDelegate() []byte {
	if x != nil {
		return x.FromDelegate
	}
	return nil
}

func (x *DelegateChanged) GetToDelegate() []byte {
	if x != nil {
		return x.ToDelegate
	}
	return nil
}

type DelegateVotesChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Delegate        []byte                 `protobuf:"bytes,2,opt,name=Delegate,proto3" json:"Delegate,omitempty"`               //	address
	PreviousBalance []byte                 `protobuf:"bytes,3,opt,name=PreviousBalance,proto3" json:"PreviousBalance,omitempty"` //	uint256
	NewBalance      []byte                 `protobuf:"bytes,4,opt,name=NewBalance,proto3" json:"NewBalance,omitempty"`           //	uint256
}

func (x *DelegateVotesChanged) Reset() {
	*x = DelegateVotesChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegateVotesChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegateVotesChanged) ProtoMessage() {}

func (x *DelegateVotesChanged) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegateVotesChanged.ProtoReflect.Descriptor instead.
func (*DelegateVotesChanged) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescGZIP(), []int{2}
}

func (x *DelegateVotesChanged) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *DelegateVotesChanged) GetDelegate() []byte {
	if x != nil {
		return x.Delegate
	}
	return nil
}

func (x *DelegateVotesChanged) GetPreviousBalance() []byte {
	if x != nil {
		return x.PreviousBalance
	}
	return nil
}

func (x *DelegateVotesChanged) GetNewBalance() []byte {
	if x != nil {
		return x.NewBalance
	}
	return nil
}

type MinterChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Minter    []byte                 `protobuf:"bytes,2,opt,name=Minter,proto3" json:"Minter,omitempty"`       //	address
	NewMinter []byte                 `protobuf:"bytes,3,opt,name=NewMinter,proto3" json:"NewMinter,omitempty"` //	address
}

func (x *MinterChanged) Reset() {
	*x = MinterChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinterChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinterChanged) ProtoMessage() {}

func (x *MinterChanged) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinterChanged.ProtoReflect.Descriptor instead.
func (*MinterChanged) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescGZIP(), []int{3}
}

func (x *MinterChanged) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *MinterChanged) GetMinter() []byte {
	if x != nil {
		return x.Minter
	}
	return nil
}

func (x *MinterChanged) GetNewMinter() []byte {
	if x != nil {
		return x.NewMinter
	}
	return nil
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	From   []byte                 `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`     //	address
	To     []byte                 `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`         //	address
	Amount []byte                 `protobuf:"bytes,4,opt,name=Amount,proto3" json:"Amount,omitempty"` //	uint256
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescGZIP(), []int{4}
}

func (x *Transfer) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Transfer) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Transfer) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Transfer) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

var File_connectors_source_makerdao_smart_contracts_UNI_UNI_proto protoreflect.FileDescriptor

var file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x55, 0x4e, 0x49,
	0x2f, 0x55, 0x4e, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x55, 0x4e, 0x49, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7e, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x9f, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x54, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x4e, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x4e, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x71, 0x0a,
	0x0d, 0x4d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x4d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x4d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x4e, 0x65, 0x77, 0x4d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x22, 0x72, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x36, 0x5a, 0x34, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x55, 0x4e, 0x49, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescOnce sync.Once
	file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescData = file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDesc
)

func file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescGZIP() []byte {
	file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescOnce.Do(func() {
		file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescData)
	})
	return file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDescData
}

var file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_goTypes = []interface{}{
	(*Approval)(nil),              // 0: UNI.Approval
	(*DelegateChanged)(nil),       // 1: UNI.DelegateChanged
	(*DelegateVotesChanged)(nil),  // 2: UNI.DelegateVotesChanged
	(*MinterChanged)(nil),         // 3: UNI.MinterChanged
	(*Transfer)(nil),              // 4: UNI.Transfer
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_depIdxs = []int32{
	5, // 0: UNI.Approval.ts:type_name -> google.protobuf.Timestamp
	5, // 1: UNI.DelegateChanged.ts:type_name -> google.protobuf.Timestamp
	5, // 2: UNI.DelegateVotesChanged.ts:type_name -> google.protobuf.Timestamp
	5, // 3: UNI.MinterChanged.ts:type_name -> google.protobuf.Timestamp
	5, // 4: UNI.Transfer.ts:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_init() }
func file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_init() {
	if File_connectors_source_makerdao_smart_contracts_UNI_UNI_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Approval); i {
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
		file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegateChanged); i {
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
		file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegateVotesChanged); i {
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
		file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinterChanged); i {
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
		file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
			RawDescriptor: file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_goTypes,
		DependencyIndexes: file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_depIdxs,
		MessageInfos:      file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_msgTypes,
	}.Build()
	File_connectors_source_makerdao_smart_contracts_UNI_UNI_proto = out.File
	file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_rawDesc = nil
	file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_goTypes = nil
	file_connectors_source_makerdao_smart_contracts_UNI_UNI_proto_depIdxs = nil
}
