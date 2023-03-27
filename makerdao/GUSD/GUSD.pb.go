// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: GUSD/GUSD.proto

package GUSD

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

type CustodianChangeConfirmed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	LockId       []byte                 `protobuf:"bytes,2,opt,name=LockId,proto3" json:"LockId,omitempty"`             //	bytes32
	NewCustodian []byte                 `protobuf:"bytes,3,opt,name=NewCustodian,proto3" json:"NewCustodian,omitempty"` //	address
}

func (x *CustodianChangeConfirmed) Reset() {
	*x = CustodianChangeConfirmed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUSD_GUSD_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustodianChangeConfirmed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustodianChangeConfirmed) ProtoMessage() {}

func (x *CustodianChangeConfirmed) ProtoReflect() protoreflect.Message {
	mi := &file_GUSD_GUSD_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustodianChangeConfirmed.ProtoReflect.Descriptor instead.
func (*CustodianChangeConfirmed) Descriptor() ([]byte, []int) {
	return file_GUSD_GUSD_proto_rawDescGZIP(), []int{0}
}

func (x *CustodianChangeConfirmed) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *CustodianChangeConfirmed) GetLockId() []byte {
	if x != nil {
		return x.LockId
	}
	return nil
}

func (x *CustodianChangeConfirmed) GetNewCustodian() []byte {
	if x != nil {
		return x.NewCustodian
	}
	return nil
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	From  []byte                 `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`   //	address
	To    []byte                 `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`       //	address
	Value []byte                 `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"` //	uint256
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUSD_GUSD_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_GUSD_GUSD_proto_msgTypes[1]
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
	return file_GUSD_GUSD_proto_rawDescGZIP(), []int{1}
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

func (x *Transfer) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type Approval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Owner   []byte                 `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`     //	address
	Spender []byte                 `protobuf:"bytes,3,opt,name=Spender,proto3" json:"Spender,omitempty"` //	address
	Value   []byte                 `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`     //	uint256
}

func (x *Approval) Reset() {
	*x = Approval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUSD_GUSD_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approval) ProtoMessage() {}

func (x *Approval) ProtoReflect() protoreflect.Message {
	mi := &file_GUSD_GUSD_proto_msgTypes[2]
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
	return file_GUSD_GUSD_proto_rawDescGZIP(), []int{2}
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

func (x *Approval) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ImplChangeRequested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	LockId       []byte                 `protobuf:"bytes,2,opt,name=LockId,proto3" json:"LockId,omitempty"`             //	bytes32
	MsgSender    []byte                 `protobuf:"bytes,3,opt,name=MsgSender,proto3" json:"MsgSender,omitempty"`       //	address
	ProposedImpl []byte                 `protobuf:"bytes,4,opt,name=ProposedImpl,proto3" json:"ProposedImpl,omitempty"` //	address
}

func (x *ImplChangeRequested) Reset() {
	*x = ImplChangeRequested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUSD_GUSD_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImplChangeRequested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImplChangeRequested) ProtoMessage() {}

func (x *ImplChangeRequested) ProtoReflect() protoreflect.Message {
	mi := &file_GUSD_GUSD_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImplChangeRequested.ProtoReflect.Descriptor instead.
func (*ImplChangeRequested) Descriptor() ([]byte, []int) {
	return file_GUSD_GUSD_proto_rawDescGZIP(), []int{3}
}

func (x *ImplChangeRequested) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ImplChangeRequested) GetLockId() []byte {
	if x != nil {
		return x.LockId
	}
	return nil
}

func (x *ImplChangeRequested) GetMsgSender() []byte {
	if x != nil {
		return x.MsgSender
	}
	return nil
}

func (x *ImplChangeRequested) GetProposedImpl() []byte {
	if x != nil {
		return x.ProposedImpl
	}
	return nil
}

type ImplChangeConfirmed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	LockId  []byte                 `protobuf:"bytes,2,opt,name=LockId,proto3" json:"LockId,omitempty"`   //	bytes32
	NewImpl []byte                 `protobuf:"bytes,3,opt,name=NewImpl,proto3" json:"NewImpl,omitempty"` //	address
}

func (x *ImplChangeConfirmed) Reset() {
	*x = ImplChangeConfirmed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUSD_GUSD_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImplChangeConfirmed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImplChangeConfirmed) ProtoMessage() {}

func (x *ImplChangeConfirmed) ProtoReflect() protoreflect.Message {
	mi := &file_GUSD_GUSD_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImplChangeConfirmed.ProtoReflect.Descriptor instead.
func (*ImplChangeConfirmed) Descriptor() ([]byte, []int) {
	return file_GUSD_GUSD_proto_rawDescGZIP(), []int{4}
}

func (x *ImplChangeConfirmed) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ImplChangeConfirmed) GetLockId() []byte {
	if x != nil {
		return x.LockId
	}
	return nil
}

func (x *ImplChangeConfirmed) GetNewImpl() []byte {
	if x != nil {
		return x.NewImpl
	}
	return nil
}

type CustodianChangeRequested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts                *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	LockId            []byte                 `protobuf:"bytes,2,opt,name=LockId,proto3" json:"LockId,omitempty"`                       //	bytes32
	MsgSender         []byte                 `protobuf:"bytes,3,opt,name=MsgSender,proto3" json:"MsgSender,omitempty"`                 //	address
	ProposedCustodian []byte                 `protobuf:"bytes,4,opt,name=ProposedCustodian,proto3" json:"ProposedCustodian,omitempty"` //	address
}

func (x *CustodianChangeRequested) Reset() {
	*x = CustodianChangeRequested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUSD_GUSD_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustodianChangeRequested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustodianChangeRequested) ProtoMessage() {}

func (x *CustodianChangeRequested) ProtoReflect() protoreflect.Message {
	mi := &file_GUSD_GUSD_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustodianChangeRequested.ProtoReflect.Descriptor instead.
func (*CustodianChangeRequested) Descriptor() ([]byte, []int) {
	return file_GUSD_GUSD_proto_rawDescGZIP(), []int{5}
}

func (x *CustodianChangeRequested) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *CustodianChangeRequested) GetLockId() []byte {
	if x != nil {
		return x.LockId
	}
	return nil
}

func (x *CustodianChangeRequested) GetMsgSender() []byte {
	if x != nil {
		return x.MsgSender
	}
	return nil
}

func (x *CustodianChangeRequested) GetProposedCustodian() []byte {
	if x != nil {
		return x.ProposedCustodian
	}
	return nil
}

var File_GUSD_GUSD_proto protoreflect.FileDescriptor

var file_GUSD_GUSD_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x55, 0x53, 0x44, 0x2f, 0x47, 0x55, 0x53, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x47, 0x55, 0x53, 0x44, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x18, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x65, 0x77,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x4e, 0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e, 0x22, 0x70, 0x0a,
	0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x7c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9b, 0x01,
	0x0a, 0x13, 0x49, 0x6d, 0x70, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x73, 0x67,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x4d, 0x73,
	0x67, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x6c, 0x22, 0x73, 0x0a, 0x13, 0x49,
	0x6d, 0x70, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x49, 0x6d, 0x70,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4e, 0x65, 0x77, 0x49, 0x6d, 0x70, 0x6c,
	0x22, 0xaa, 0x01, 0x0a, 0x18, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x64, 0x69, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e, 0x42, 0x37, 0x5a,
	0x35, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61,
	0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x47, 0x55, 0x53, 0x44, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GUSD_GUSD_proto_rawDescOnce sync.Once
	file_GUSD_GUSD_proto_rawDescData = file_GUSD_GUSD_proto_rawDesc
)

func file_GUSD_GUSD_proto_rawDescGZIP() []byte {
	file_GUSD_GUSD_proto_rawDescOnce.Do(func() {
		file_GUSD_GUSD_proto_rawDescData = protoimpl.X.CompressGZIP(file_GUSD_GUSD_proto_rawDescData)
	})
	return file_GUSD_GUSD_proto_rawDescData
}

var file_GUSD_GUSD_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_GUSD_GUSD_proto_goTypes = []interface{}{
	(*CustodianChangeConfirmed)(nil), // 0: GUSD.CustodianChangeConfirmed
	(*Transfer)(nil),                 // 1: GUSD.Transfer
	(*Approval)(nil),                 // 2: GUSD.Approval
	(*ImplChangeRequested)(nil),      // 3: GUSD.ImplChangeRequested
	(*ImplChangeConfirmed)(nil),      // 4: GUSD.ImplChangeConfirmed
	(*CustodianChangeRequested)(nil), // 5: GUSD.CustodianChangeRequested
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_GUSD_GUSD_proto_depIdxs = []int32{
	6, // 0: GUSD.CustodianChangeConfirmed.ts:type_name -> google.protobuf.Timestamp
	6, // 1: GUSD.Transfer.ts:type_name -> google.protobuf.Timestamp
	6, // 2: GUSD.Approval.ts:type_name -> google.protobuf.Timestamp
	6, // 3: GUSD.ImplChangeRequested.ts:type_name -> google.protobuf.Timestamp
	6, // 4: GUSD.ImplChangeConfirmed.ts:type_name -> google.protobuf.Timestamp
	6, // 5: GUSD.CustodianChangeRequested.ts:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_GUSD_GUSD_proto_init() }
func file_GUSD_GUSD_proto_init() {
	if File_GUSD_GUSD_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GUSD_GUSD_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustodianChangeConfirmed); i {
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
		file_GUSD_GUSD_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_GUSD_GUSD_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_GUSD_GUSD_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImplChangeRequested); i {
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
		file_GUSD_GUSD_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImplChangeConfirmed); i {
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
		file_GUSD_GUSD_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustodianChangeRequested); i {
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
			RawDescriptor: file_GUSD_GUSD_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GUSD_GUSD_proto_goTypes,
		DependencyIndexes: file_GUSD_GUSD_proto_depIdxs,
		MessageInfos:      file_GUSD_GUSD_proto_msgTypes,
	}.Build()
	File_GUSD_GUSD_proto = out.File
	file_GUSD_GUSD_proto_rawDesc = nil
	file_GUSD_GUSD_proto_goTypes = nil
	file_GUSD_GUSD_proto_depIdxs = nil
}
