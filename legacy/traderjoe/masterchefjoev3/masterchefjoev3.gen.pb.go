// Code generated - DO NOT EDIT.
// This file is a generated protobuf definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: connectors/source/traderjoe/masterchefjoev3/masterchefjoev3.gen.proto

package masterchefjoev3

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

type Harvest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	User   []byte                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pid    []byte                 `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`       // uint256
	Amount []byte                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // uint256
}

func (x *Harvest) Reset() {
	*x = Harvest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Harvest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Harvest) ProtoMessage() {}

func (x *Harvest) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Harvest.ProtoReflect.Descriptor instead.
func (*Harvest) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{0}
}

func (x *Harvest) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Harvest) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Harvest) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *Harvest) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

type Withdraw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	User   []byte                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pid    []byte                 `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`       // uint256
	Amount []byte                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // uint256
}

func (x *Withdraw) Reset() {
	*x = Withdraw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Withdraw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdraw) ProtoMessage() {}

func (x *Withdraw) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdraw.ProtoReflect.Descriptor instead.
func (*Withdraw) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{1}
}

func (x *Withdraw) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Withdraw) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Withdraw) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *Withdraw) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

type OwnershipTransferred struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	PreviousOwner []byte                 `protobuf:"bytes,2,opt,name=previousOwner,proto3" json:"previousOwner,omitempty"`
	NewOwner      []byte                 `protobuf:"bytes,3,opt,name=newOwner,proto3" json:"newOwner,omitempty"`
}

func (x *OwnershipTransferred) Reset() {
	*x = OwnershipTransferred{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnershipTransferred) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnershipTransferred) ProtoMessage() {}

func (x *OwnershipTransferred) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[2]
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
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{2}
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

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Pid        []byte                 `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`               // uint256
	AllocPoint []byte                 `protobuf:"bytes,3,opt,name=allocPoint,proto3" json:"allocPoint,omitempty"` // uint256
	Rewarder   []byte                 `protobuf:"bytes,4,opt,name=rewarder,proto3" json:"rewarder,omitempty"`
	Overwrite  bool                   `protobuf:"varint,5,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{3}
}

func (x *Set) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Set) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *Set) GetAllocPoint() []byte {
	if x != nil {
		return x.AllocPoint
	}
	return nil
}

func (x *Set) GetRewarder() []byte {
	if x != nil {
		return x.Rewarder
	}
	return nil
}

func (x *Set) GetOverwrite() bool {
	if x != nil {
		return x.Overwrite
	}
	return false
}

type Deposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	User   []byte                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pid    []byte                 `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`       // uint256
	Amount []byte                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // uint256
}

func (x *Deposit) Reset() {
	*x = Deposit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deposit) ProtoMessage() {}

func (x *Deposit) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deposit.ProtoReflect.Descriptor instead.
func (*Deposit) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{4}
}

func (x *Deposit) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Deposit) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Deposit) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *Deposit) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

type EmergencyWithdraw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	User   []byte                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pid    []byte                 `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`       // uint256
	Amount []byte                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // uint256
}

func (x *EmergencyWithdraw) Reset() {
	*x = EmergencyWithdraw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmergencyWithdraw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmergencyWithdraw) ProtoMessage() {}

func (x *EmergencyWithdraw) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmergencyWithdraw.ProtoReflect.Descriptor instead.
func (*EmergencyWithdraw) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{5}
}

func (x *EmergencyWithdraw) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *EmergencyWithdraw) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *EmergencyWithdraw) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *EmergencyWithdraw) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

type Init struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Init) Reset() {
	*x = Init{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Init) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Init) ProtoMessage() {}

func (x *Init) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Init.ProtoReflect.Descriptor instead.
func (*Init) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{6}
}

func (x *Init) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type UpdatePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts                  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Pid                 []byte                 `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`                                 // uint256
	LastRewardTimestamp []byte                 `protobuf:"bytes,3,opt,name=lastRewardTimestamp,proto3" json:"lastRewardTimestamp,omitempty"` // uint256
	LpSupply            []byte                 `protobuf:"bytes,4,opt,name=lpSupply,proto3" json:"lpSupply,omitempty"`                       // uint256
	AccJoePerShare      []byte                 `protobuf:"bytes,5,opt,name=accJoePerShare,proto3" json:"accJoePerShare,omitempty"`           // uint256
}

func (x *UpdatePool) Reset() {
	*x = UpdatePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePool) ProtoMessage() {}

func (x *UpdatePool) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePool.ProtoReflect.Descriptor instead.
func (*UpdatePool) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePool) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *UpdatePool) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *UpdatePool) GetLastRewardTimestamp() []byte {
	if x != nil {
		return x.LastRewardTimestamp
	}
	return nil
}

func (x *UpdatePool) GetLpSupply() []byte {
	if x != nil {
		return x.LpSupply
	}
	return nil
}

func (x *UpdatePool) GetAccJoePerShare() []byte {
	if x != nil {
		return x.AccJoePerShare
	}
	return nil
}

type Add struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Pid        []byte                 `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`               // uint256
	AllocPoint []byte                 `protobuf:"bytes,3,opt,name=allocPoint,proto3" json:"allocPoint,omitempty"` // uint256
	LpToken    []byte                 `protobuf:"bytes,4,opt,name=lpToken,proto3" json:"lpToken,omitempty"`
	Rewarder   []byte                 `protobuf:"bytes,5,opt,name=rewarder,proto3" json:"rewarder,omitempty"`
}

func (x *Add) Reset() {
	*x = Add{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Add) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Add) ProtoMessage() {}

func (x *Add) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Add.ProtoReflect.Descriptor instead.
func (*Add) Descriptor() ([]byte, []int) {
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP(), []int{8}
}

func (x *Add) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Add) GetPid() []byte {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *Add) GetAllocPoint() []byte {
	if x != nil {
		return x.AllocPoint
	}
	return nil
}

func (x *Add) GetLpToken() []byte {
	if x != nil {
		return x.LpToken
	}
	return nil
}

func (x *Add) GetRewarder() []byte {
	if x != nil {
		return x.Rewarder
	}
	return nil
}

var File_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto protoreflect.FileDescriptor

var file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x6a, 0x6f, 0x65, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x63, 0x68, 0x65, 0x66, 0x6a, 0x6f, 0x65, 0x76, 0x33, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x63, 0x68, 0x65, 0x66, 0x6a, 0x6f, 0x65, 0x76, 0x33, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x63,
	0x68, 0x65, 0x66, 0x6a, 0x6f, 0x65, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x07, 0x48, 0x61, 0x72,
	0x76, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x74,
	0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x9d, 0x01, 0x0a, 0x03,
	0x53, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x73, 0x0a, 0x07, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02,
	0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x7d, 0x0a, 0x11, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x32, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x02, 0x74, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x69, 0x64,
	0x12, 0x30, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x70, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x70, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x26,
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x4a, 0x6f, 0x65, 0x50, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x4a, 0x6f, 0x65, 0x50, 0x65,
	0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6c,
	0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x29, 0x5a, 0x27, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x63, 0x68, 0x65, 0x66, 0x6a, 0x6f, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescOnce sync.Once
	file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescData = file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDesc
)

func file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescGZIP() []byte {
	file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescOnce.Do(func() {
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescData)
	})
	return file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDescData
}

var file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_goTypes = []interface{}{
	(*Harvest)(nil),               // 0: masterchefjoev3.Harvest
	(*Withdraw)(nil),              // 1: masterchefjoev3.Withdraw
	(*OwnershipTransferred)(nil),  // 2: masterchefjoev3.OwnershipTransferred
	(*Set)(nil),                   // 3: masterchefjoev3.Set
	(*Deposit)(nil),               // 4: masterchefjoev3.Deposit
	(*EmergencyWithdraw)(nil),     // 5: masterchefjoev3.EmergencyWithdraw
	(*Init)(nil),                  // 6: masterchefjoev3.Init
	(*UpdatePool)(nil),            // 7: masterchefjoev3.UpdatePool
	(*Add)(nil),                   // 8: masterchefjoev3.Add
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_depIdxs = []int32{
	9, // 0: masterchefjoev3.Harvest.ts:type_name -> google.protobuf.Timestamp
	9, // 1: masterchefjoev3.Withdraw.ts:type_name -> google.protobuf.Timestamp
	9, // 2: masterchefjoev3.OwnershipTransferred.ts:type_name -> google.protobuf.Timestamp
	9, // 3: masterchefjoev3.Set.ts:type_name -> google.protobuf.Timestamp
	9, // 4: masterchefjoev3.Deposit.ts:type_name -> google.protobuf.Timestamp
	9, // 5: masterchefjoev3.EmergencyWithdraw.ts:type_name -> google.protobuf.Timestamp
	9, // 6: masterchefjoev3.Init.ts:type_name -> google.protobuf.Timestamp
	9, // 7: masterchefjoev3.UpdatePool.ts:type_name -> google.protobuf.Timestamp
	9, // 8: masterchefjoev3.Add.ts:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_init() }
func file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_init() {
	if File_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Harvest); i {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Withdraw); i {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deposit); i {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmergencyWithdraw); i {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Init); i {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePool); i {
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
		file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Add); i {
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
			RawDescriptor: file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_goTypes,
		DependencyIndexes: file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_depIdxs,
		MessageInfos:      file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_msgTypes,
	}.Build()
	File_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto = out.File
	file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_rawDesc = nil
	file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_goTypes = nil
	file_connectors_source_traderjoe_masterchefjoev3_masterchefjoev3_gen_proto_depIdxs = nil
}
