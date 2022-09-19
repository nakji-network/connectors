// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: flowfees.proto

package flowfees

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

type TokensDeposited struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber uint64                 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TxID        []byte                 `protobuf:"bytes,3,opt,name=txID,proto3" json:"txID,omitempty"`
	LogIndex    uint64                 `protobuf:"varint,4,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Amount      float64                `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TokensDeposited) Reset() {
	*x = TokensDeposited{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flowfees_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensDeposited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensDeposited) ProtoMessage() {}

func (x *TokensDeposited) ProtoReflect() protoreflect.Message {
	mi := &file_flowfees_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensDeposited.ProtoReflect.Descriptor instead.
func (*TokensDeposited) Descriptor() ([]byte, []int) {
	return file_flowfees_proto_rawDescGZIP(), []int{0}
}

func (x *TokensDeposited) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *TokensDeposited) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *TokensDeposited) GetTxID() []byte {
	if x != nil {
		return x.TxID
	}
	return nil
}

func (x *TokensDeposited) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *TokensDeposited) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TokensWithdrawn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber uint64                 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TxID        []byte                 `protobuf:"bytes,3,opt,name=txID,proto3" json:"txID,omitempty"`
	LogIndex    uint64                 `protobuf:"varint,4,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Amount      float64                `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TokensWithdrawn) Reset() {
	*x = TokensWithdrawn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flowfees_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensWithdrawn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensWithdrawn) ProtoMessage() {}

func (x *TokensWithdrawn) ProtoReflect() protoreflect.Message {
	mi := &file_flowfees_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensWithdrawn.ProtoReflect.Descriptor instead.
func (*TokensWithdrawn) Descriptor() ([]byte, []int) {
	return file_flowfees_proto_rawDescGZIP(), []int{1}
}

func (x *TokensWithdrawn) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *TokensWithdrawn) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *TokensWithdrawn) GetTxID() []byte {
	if x != nil {
		return x.TxID
	}
	return nil
}

func (x *TokensWithdrawn) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *TokensWithdrawn) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type FeesDeducted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber     uint64                 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TxID            []byte                 `protobuf:"bytes,3,opt,name=txID,proto3" json:"txID,omitempty"`
	LogIndex        uint64                 `protobuf:"varint,4,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Amount          float64                `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	InclusionEffort float64                `protobuf:"fixed64,6,opt,name=inclusionEffort,proto3" json:"inclusionEffort,omitempty"`
	ExecutionEffort float64                `protobuf:"fixed64,7,opt,name=executionEffort,proto3" json:"executionEffort,omitempty"`
}

func (x *FeesDeducted) Reset() {
	*x = FeesDeducted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flowfees_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeesDeducted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeesDeducted) ProtoMessage() {}

func (x *FeesDeducted) ProtoReflect() protoreflect.Message {
	mi := &file_flowfees_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeesDeducted.ProtoReflect.Descriptor instead.
func (*FeesDeducted) Descriptor() ([]byte, []int) {
	return file_flowfees_proto_rawDescGZIP(), []int{2}
}

func (x *FeesDeducted) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *FeesDeducted) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *FeesDeducted) GetTxID() []byte {
	if x != nil {
		return x.TxID
	}
	return nil
}

func (x *FeesDeducted) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *FeesDeducted) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FeesDeducted) GetInclusionEffort() float64 {
	if x != nil {
		return x.InclusionEffort
	}
	return 0
}

func (x *FeesDeducted) GetExecutionEffort() float64 {
	if x != nil {
		return x.ExecutionEffort
	}
	return 0
}

type FeeParametersChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts                  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber         uint64                 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TxID                []byte                 `protobuf:"bytes,3,opt,name=txID,proto3" json:"txID,omitempty"`
	LogIndex            uint64                 `protobuf:"varint,4,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	SurgeFactor         float64                `protobuf:"fixed64,5,opt,name=surgeFactor,proto3" json:"surgeFactor,omitempty"`
	InclusionEffortCost float64                `protobuf:"fixed64,6,opt,name=inclusionEffortCost,proto3" json:"inclusionEffortCost,omitempty"`
	ExecutionEffortCost float64                `protobuf:"fixed64,7,opt,name=executionEffortCost,proto3" json:"executionEffortCost,omitempty"`
}

func (x *FeeParametersChanged) Reset() {
	*x = FeeParametersChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flowfees_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeParametersChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeParametersChanged) ProtoMessage() {}

func (x *FeeParametersChanged) ProtoReflect() protoreflect.Message {
	mi := &file_flowfees_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeParametersChanged.ProtoReflect.Descriptor instead.
func (*FeeParametersChanged) Descriptor() ([]byte, []int) {
	return file_flowfees_proto_rawDescGZIP(), []int{3}
}

func (x *FeeParametersChanged) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *FeeParametersChanged) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *FeeParametersChanged) GetTxID() []byte {
	if x != nil {
		return x.TxID
	}
	return nil
}

func (x *FeeParametersChanged) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *FeeParametersChanged) GetSurgeFactor() float64 {
	if x != nil {
		return x.SurgeFactor
	}
	return 0
}

func (x *FeeParametersChanged) GetInclusionEffortCost() float64 {
	if x != nil {
		return x.InclusionEffortCost
	}
	return 0
}

func (x *FeeParametersChanged) GetExecutionEffortCost() float64 {
	if x != nil {
		return x.ExecutionEffortCost
	}
	return 0
}

var File_flowfees_proto protoreflect.FileDescriptor

var file_flowfees_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x6c, 0x6f, 0x77, 0x66, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x6c, 0x6f, 0x77, 0x66, 0x65, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0f,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x64, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x78, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x6e, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c,
	0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xf8, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x73, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x78, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x66, 0x66, 0x6f, 0x72, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x66, 0x66,
	0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x22, 0x9a, 0x02, 0x0a, 0x14, 0x46,
	0x65, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x74, 0x78, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x72, 0x67, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x73, 0x75, 0x72, 0x67, 0x65, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x66, 0x66, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x66, 0x66, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x13, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x66, 0x66,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x6b, 0x6a, 0x69, 0x2d, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x66, 0x65, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flowfees_proto_rawDescOnce sync.Once
	file_flowfees_proto_rawDescData = file_flowfees_proto_rawDesc
)

func file_flowfees_proto_rawDescGZIP() []byte {
	file_flowfees_proto_rawDescOnce.Do(func() {
		file_flowfees_proto_rawDescData = protoimpl.X.CompressGZIP(file_flowfees_proto_rawDescData)
	})
	return file_flowfees_proto_rawDescData
}

var file_flowfees_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flowfees_proto_goTypes = []interface{}{
	(*TokensDeposited)(nil),       // 0: flowfees.TokensDeposited
	(*TokensWithdrawn)(nil),       // 1: flowfees.TokensWithdrawn
	(*FeesDeducted)(nil),          // 2: flowfees.FeesDeducted
	(*FeeParametersChanged)(nil),  // 3: flowfees.FeeParametersChanged
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_flowfees_proto_depIdxs = []int32{
	4, // 0: flowfees.TokensDeposited.ts:type_name -> google.protobuf.Timestamp
	4, // 1: flowfees.TokensWithdrawn.ts:type_name -> google.protobuf.Timestamp
	4, // 2: flowfees.FeesDeducted.ts:type_name -> google.protobuf.Timestamp
	4, // 3: flowfees.FeeParametersChanged.ts:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_flowfees_proto_init() }
func file_flowfees_proto_init() {
	if File_flowfees_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flowfees_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokensDeposited); i {
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
		file_flowfees_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokensWithdrawn); i {
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
		file_flowfees_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeesDeducted); i {
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
		file_flowfees_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeParametersChanged); i {
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
			RawDescriptor: file_flowfees_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flowfees_proto_goTypes,
		DependencyIndexes: file_flowfees_proto_depIdxs,
		MessageInfos:      file_flowfees_proto_msgTypes,
	}.Build()
	File_flowfees_proto = out.File
	file_flowfees_proto_rawDesc = nil
	file_flowfees_proto_goTypes = nil
	file_flowfees_proto_depIdxs = nil
}
