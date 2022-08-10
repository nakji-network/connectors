// Code generated - DO NOT EDIT.
// This file is a generated protobuf definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: WOOPP.proto

package WOOPP

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

type Paused struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Account []byte                 `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *Paused) Reset() {
	*x = Paused{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paused) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paused) ProtoMessage() {}

func (x *Paused) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[0]
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
	return file_WOOPP_proto_rawDescGZIP(), []int{0}
}

func (x *Paused) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Paused) GetAccount() []byte {
	if x != nil {
		return x.Account
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
		mi := &file_WOOPP_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnershipTransferred) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnershipTransferred) ProtoMessage() {}

func (x *OwnershipTransferred) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[1]
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
	return file_WOOPP_proto_rawDescGZIP(), []int{1}
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

type ParametersUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BaseToken    []byte                 `protobuf:"bytes,2,opt,name=baseToken,proto3" json:"baseToken,omitempty"`
	NewThreshold []byte                 `protobuf:"bytes,3,opt,name=newThreshold,proto3" json:"newThreshold,omitempty"` // uint256
	NewR         []byte                 `protobuf:"bytes,4,opt,name=newR,proto3" json:"newR,omitempty"`                 // uint256
}

func (x *ParametersUpdated) Reset() {
	*x = ParametersUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParametersUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParametersUpdated) ProtoMessage() {}

func (x *ParametersUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParametersUpdated.ProtoReflect.Descriptor instead.
func (*ParametersUpdated) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{2}
}

func (x *ParametersUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ParametersUpdated) GetBaseToken() []byte {
	if x != nil {
		return x.BaseToken
	}
	return nil
}

func (x *ParametersUpdated) GetNewThreshold() []byte {
	if x != nil {
		return x.NewThreshold
	}
	return nil
}

func (x *ParametersUpdated) GetNewR() []byte {
	if x != nil {
		return x.NewR
	}
	return nil
}

type FeeManagerUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	NewFeeManager []byte                 `protobuf:"bytes,2,opt,name=newFeeManager,proto3" json:"newFeeManager,omitempty"`
}

func (x *FeeManagerUpdated) Reset() {
	*x = FeeManagerUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeManagerUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeManagerUpdated) ProtoMessage() {}

func (x *FeeManagerUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeManagerUpdated.ProtoReflect.Descriptor instead.
func (*FeeManagerUpdated) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{3}
}

func (x *FeeManagerUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *FeeManagerUpdated) GetNewFeeManager() []byte {
	if x != nil {
		return x.NewFeeManager
	}
	return nil
}

type RewardManagerUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts               *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	NewRewardManager []byte                 `protobuf:"bytes,2,opt,name=newRewardManager,proto3" json:"newRewardManager,omitempty"`
}

func (x *RewardManagerUpdated) Reset() {
	*x = RewardManagerUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardManagerUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardManagerUpdated) ProtoMessage() {}

func (x *RewardManagerUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardManagerUpdated.ProtoReflect.Descriptor instead.
func (*RewardManagerUpdated) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{4}
}

func (x *RewardManagerUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *RewardManagerUpdated) GetNewRewardManager() []byte {
	if x != nil {
		return x.NewRewardManager
	}
	return nil
}

type Unpaused struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Account []byte                 `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *Unpaused) Reset() {
	*x = Unpaused{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unpaused) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unpaused) ProtoMessage() {}

func (x *Unpaused) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[5]
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
	return file_WOOPP_proto_rawDescGZIP(), []int{5}
}

func (x *Unpaused) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Unpaused) GetAccount() []byte {
	if x != nil {
		return x.Account
	}
	return nil
}

type WooGuardianUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts             *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	NewWooGuardian []byte                 `protobuf:"bytes,2,opt,name=newWooGuardian,proto3" json:"newWooGuardian,omitempty"`
}

func (x *WooGuardianUpdated) Reset() {
	*x = WooGuardianUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WooGuardianUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WooGuardianUpdated) ProtoMessage() {}

func (x *WooGuardianUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WooGuardianUpdated.ProtoReflect.Descriptor instead.
func (*WooGuardianUpdated) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{6}
}

func (x *WooGuardianUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *WooGuardianUpdated) GetNewWooGuardian() []byte {
	if x != nil {
		return x.NewWooGuardian
	}
	return nil
}

type WooSwap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	FromToken  []byte                 `protobuf:"bytes,2,opt,name=fromToken,proto3" json:"fromToken,omitempty"`
	ToToken    []byte                 `protobuf:"bytes,3,opt,name=toToken,proto3" json:"toToken,omitempty"`
	FromAmount []byte                 `protobuf:"bytes,4,opt,name=fromAmount,proto3" json:"fromAmount,omitempty"` // uint256
	ToAmount   []byte                 `protobuf:"bytes,5,opt,name=toAmount,proto3" json:"toAmount,omitempty"`     // uint256
	From       []byte                 `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	To         []byte                 `protobuf:"bytes,7,opt,name=to,proto3" json:"to,omitempty"`
	RebateTo   []byte                 `protobuf:"bytes,8,opt,name=rebateTo,proto3" json:"rebateTo,omitempty"`
}

func (x *WooSwap) Reset() {
	*x = WooSwap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WooSwap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WooSwap) ProtoMessage() {}

func (x *WooSwap) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WooSwap.ProtoReflect.Descriptor instead.
func (*WooSwap) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{7}
}

func (x *WooSwap) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *WooSwap) GetFromToken() []byte {
	if x != nil {
		return x.FromToken
	}
	return nil
}

func (x *WooSwap) GetToToken() []byte {
	if x != nil {
		return x.ToToken
	}
	return nil
}

func (x *WooSwap) GetFromAmount() []byte {
	if x != nil {
		return x.FromAmount
	}
	return nil
}

func (x *WooSwap) GetToAmount() []byte {
	if x != nil {
		return x.ToAmount
	}
	return nil
}

func (x *WooSwap) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *WooSwap) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *WooSwap) GetRebateTo() []byte {
	if x != nil {
		return x.RebateTo
	}
	return nil
}

type OwnershipTransferPrepared struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	PreviousOwner []byte                 `protobuf:"bytes,2,opt,name=previousOwner,proto3" json:"previousOwner,omitempty"`
	NewOwner      []byte                 `protobuf:"bytes,3,opt,name=newOwner,proto3" json:"newOwner,omitempty"`
}

func (x *OwnershipTransferPrepared) Reset() {
	*x = OwnershipTransferPrepared{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnershipTransferPrepared) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnershipTransferPrepared) ProtoMessage() {}

func (x *OwnershipTransferPrepared) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnershipTransferPrepared.ProtoReflect.Descriptor instead.
func (*OwnershipTransferPrepared) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{8}
}

func (x *OwnershipTransferPrepared) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *OwnershipTransferPrepared) GetPreviousOwner() []byte {
	if x != nil {
		return x.PreviousOwner
	}
	return nil
}

func (x *OwnershipTransferPrepared) GetNewOwner() []byte {
	if x != nil {
		return x.NewOwner
	}
	return nil
}

type WooracleUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	NewWooracle []byte                 `protobuf:"bytes,2,opt,name=newWooracle,proto3" json:"newWooracle,omitempty"`
}

func (x *WooracleUpdated) Reset() {
	*x = WooracleUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WooracleUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WooracleUpdated) ProtoMessage() {}

func (x *WooracleUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WooracleUpdated.ProtoReflect.Descriptor instead.
func (*WooracleUpdated) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{9}
}

func (x *WooracleUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *WooracleUpdated) GetNewWooracle() []byte {
	if x != nil {
		return x.NewWooracle
	}
	return nil
}

type StrategistUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Strategist []byte                 `protobuf:"bytes,2,opt,name=strategist,proto3" json:"strategist,omitempty"`
	Flag       bool                   `protobuf:"varint,3,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *StrategistUpdated) Reset() {
	*x = StrategistUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategistUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategistUpdated) ProtoMessage() {}

func (x *StrategistUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategistUpdated.ProtoReflect.Descriptor instead.
func (*StrategistUpdated) Descriptor() ([]byte, []int) {
	return file_WOOPP_proto_rawDescGZIP(), []int{10}
}

func (x *StrategistUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *StrategistUpdated) GetStrategist() []byte {
	if x != nil {
		return x.Strategist
	}
	return nil
}

func (x *StrategistUpdated) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

type Withdraw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Token  []byte                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	To     []byte                 `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Amount []byte                 `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // uint256
}

func (x *Withdraw) Reset() {
	*x = Withdraw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WOOPP_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Withdraw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdraw) ProtoMessage() {}

func (x *Withdraw) ProtoReflect() protoreflect.Message {
	mi := &file_WOOPP_proto_msgTypes[11]
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
	return file_WOOPP_proto_rawDescGZIP(), []int{11}
}

func (x *Withdraw) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Withdraw) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *Withdraw) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Withdraw) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

var File_WOOPP_proto protoreflect.FileDescriptor

var file_WOOPP_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x57, 0x4f, 0x4f, 0x50, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x57,
	0x4f, 0x4f, 0x50, 0x50, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x06, 0x50, 0x61, 0x75, 0x73, 0x65, 0x64, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x95, 0x01, 0x0a,
	0x11, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x6e, 0x65, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x52, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x6e, 0x65, 0x77, 0x52, 0x22, 0x65, 0x0a, 0x11, 0x46, 0x65, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x46, 0x65, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6e, 0x65,
	0x77, 0x46, 0x65, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x14, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x6e, 0x65, 0x77, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x08, 0x55,
	0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x02, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x68, 0x0a,
	0x12, 0x57, 0x6f, 0x6f, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x57, 0x6f, 0x6f, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x57, 0x6f, 0x6f, 0x47,
	0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x22, 0xe9, 0x01, 0x0a, 0x07, 0x57, 0x6f, 0x6f, 0x53,
	0x77, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x6f, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x62, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x62, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x19, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x22,
	0x5f, 0x0a, 0x0f, 0x57, 0x6f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x57, 0x6f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x57, 0x6f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x22, 0x73, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x66, 0x6c, 0x61, 0x67, 0x22, 0x74, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x6b, 0x6a, 0x69, 0x2d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x77, 0x6f, 0x6f, 0x66, 0x69, 0x2f, 0x57, 0x4f, 0x4f, 0x50, 0x50, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WOOPP_proto_rawDescOnce sync.Once
	file_WOOPP_proto_rawDescData = file_WOOPP_proto_rawDesc
)

func file_WOOPP_proto_rawDescGZIP() []byte {
	file_WOOPP_proto_rawDescOnce.Do(func() {
		file_WOOPP_proto_rawDescData = protoimpl.X.CompressGZIP(file_WOOPP_proto_rawDescData)
	})
	return file_WOOPP_proto_rawDescData
}

var file_WOOPP_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_WOOPP_proto_goTypes = []interface{}{
	(*Paused)(nil),                    // 0: WOOPP.Paused
	(*OwnershipTransferred)(nil),      // 1: WOOPP.OwnershipTransferred
	(*ParametersUpdated)(nil),         // 2: WOOPP.ParametersUpdated
	(*FeeManagerUpdated)(nil),         // 3: WOOPP.FeeManagerUpdated
	(*RewardManagerUpdated)(nil),      // 4: WOOPP.RewardManagerUpdated
	(*Unpaused)(nil),                  // 5: WOOPP.Unpaused
	(*WooGuardianUpdated)(nil),        // 6: WOOPP.WooGuardianUpdated
	(*WooSwap)(nil),                   // 7: WOOPP.WooSwap
	(*OwnershipTransferPrepared)(nil), // 8: WOOPP.OwnershipTransferPrepared
	(*WooracleUpdated)(nil),           // 9: WOOPP.WooracleUpdated
	(*StrategistUpdated)(nil),         // 10: WOOPP.StrategistUpdated
	(*Withdraw)(nil),                  // 11: WOOPP.Withdraw
	(*timestamppb.Timestamp)(nil),     // 12: google.protobuf.Timestamp
}
var file_WOOPP_proto_depIdxs = []int32{
	12, // 0: WOOPP.Paused.ts:type_name -> google.protobuf.Timestamp
	12, // 1: WOOPP.OwnershipTransferred.ts:type_name -> google.protobuf.Timestamp
	12, // 2: WOOPP.ParametersUpdated.ts:type_name -> google.protobuf.Timestamp
	12, // 3: WOOPP.FeeManagerUpdated.ts:type_name -> google.protobuf.Timestamp
	12, // 4: WOOPP.RewardManagerUpdated.ts:type_name -> google.protobuf.Timestamp
	12, // 5: WOOPP.Unpaused.ts:type_name -> google.protobuf.Timestamp
	12, // 6: WOOPP.WooGuardianUpdated.ts:type_name -> google.protobuf.Timestamp
	12, // 7: WOOPP.WooSwap.ts:type_name -> google.protobuf.Timestamp
	12, // 8: WOOPP.OwnershipTransferPrepared.ts:type_name -> google.protobuf.Timestamp
	12, // 9: WOOPP.WooracleUpdated.ts:type_name -> google.protobuf.Timestamp
	12, // 10: WOOPP.StrategistUpdated.ts:type_name -> google.protobuf.Timestamp
	12, // 11: WOOPP.Withdraw.ts:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_WOOPP_proto_init() }
func file_WOOPP_proto_init() {
	if File_WOOPP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WOOPP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_WOOPP_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_WOOPP_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParametersUpdated); i {
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
		file_WOOPP_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeManagerUpdated); i {
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
		file_WOOPP_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardManagerUpdated); i {
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
		file_WOOPP_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_WOOPP_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WooGuardianUpdated); i {
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
		file_WOOPP_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WooSwap); i {
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
		file_WOOPP_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnershipTransferPrepared); i {
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
		file_WOOPP_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WooracleUpdated); i {
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
		file_WOOPP_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategistUpdated); i {
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
		file_WOOPP_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_WOOPP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WOOPP_proto_goTypes,
		DependencyIndexes: file_WOOPP_proto_depIdxs,
		MessageInfos:      file_WOOPP_proto_msgTypes,
	}.Build()
	File_WOOPP_proto = out.File
	file_WOOPP_proto_rawDesc = nil
	file_WOOPP_proto_goTypes = nil
	file_WOOPP_proto_depIdxs = nil
}
