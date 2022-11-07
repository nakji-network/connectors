// Code generated - DO NOT EDIT.
// This file is a generated protobuf definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: WooCrossChainRouterV1.proto

package WooCrossChainRouterV1

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

type WooCrossSwapOnSrcChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber     uint64                 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Index           uint64                 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	TxHash          []byte                 `protobuf:"bytes,4,opt,name=txHash,proto3" json:"txHash,omitempty"`
	RefId           []byte                 `protobuf:"bytes,5,opt,name=refId,proto3" json:"refId,omitempty"` // uint256
	Sender          []byte                 `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	To              []byte                 `protobuf:"bytes,7,opt,name=to,proto3" json:"to,omitempty"`
	FromToken       []byte                 `protobuf:"bytes,8,opt,name=fromToken,proto3" json:"fromToken,omitempty"`
	FromAmount      []byte                 `protobuf:"bytes,9,opt,name=fromAmount,proto3" json:"fromAmount,omitempty"`            // uint256
	MinQuoteAmount  []byte                 `protobuf:"bytes,10,opt,name=minQuoteAmount,proto3" json:"minQuoteAmount,omitempty"`   // uint256
	RealQuoteAmount []byte                 `protobuf:"bytes,11,opt,name=realQuoteAmount,proto3" json:"realQuoteAmount,omitempty"` // uint256
	ContractAddress string                 `protobuf:"bytes,12,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (x *WooCrossSwapOnSrcChain) Reset() {
	*x = WooCrossSwapOnSrcChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WooCrossChainRouterV1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WooCrossSwapOnSrcChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WooCrossSwapOnSrcChain) ProtoMessage() {}

func (x *WooCrossSwapOnSrcChain) ProtoReflect() protoreflect.Message {
	mi := &file_WooCrossChainRouterV1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WooCrossSwapOnSrcChain.ProtoReflect.Descriptor instead.
func (*WooCrossSwapOnSrcChain) Descriptor() ([]byte, []int) {
	return file_WooCrossChainRouterV1_proto_rawDescGZIP(), []int{0}
}

func (x *WooCrossSwapOnSrcChain) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *WooCrossSwapOnSrcChain) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *WooCrossSwapOnSrcChain) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetRefId() []byte {
	if x != nil {
		return x.RefId
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetFromToken() []byte {
	if x != nil {
		return x.FromToken
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetFromAmount() []byte {
	if x != nil {
		return x.FromAmount
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetMinQuoteAmount() []byte {
	if x != nil {
		return x.MinQuoteAmount
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetRealQuoteAmount() []byte {
	if x != nil {
		return x.RealQuoteAmount
	}
	return nil
}

func (x *WooCrossSwapOnSrcChain) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

type WooCrossSwapOnDstChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber     uint64                 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Index           uint64                 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	TxHash          []byte                 `protobuf:"bytes,4,opt,name=txHash,proto3" json:"txHash,omitempty"`
	RefId           []byte                 `protobuf:"bytes,5,opt,name=refId,proto3" json:"refId,omitempty"` // uint256
	Sender          []byte                 `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	To              []byte                 `protobuf:"bytes,7,opt,name=to,proto3" json:"to,omitempty"`
	BridgedToken    []byte                 `protobuf:"bytes,8,opt,name=bridgedToken,proto3" json:"bridgedToken,omitempty"`
	BridgedAmount   []byte                 `protobuf:"bytes,9,opt,name=bridgedAmount,proto3" json:"bridgedAmount,omitempty"` // uint256
	ToToken         []byte                 `protobuf:"bytes,10,opt,name=toToken,proto3" json:"toToken,omitempty"`
	RealToToken     []byte                 `protobuf:"bytes,11,opt,name=realToToken,proto3" json:"realToToken,omitempty"`
	MinToAmount     []byte                 `protobuf:"bytes,12,opt,name=minToAmount,proto3" json:"minToAmount,omitempty"`   // uint256
	RealToAmount    []byte                 `protobuf:"bytes,13,opt,name=realToAmount,proto3" json:"realToAmount,omitempty"` // uint256
	ContractAddress string                 `protobuf:"bytes,14,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (x *WooCrossSwapOnDstChain) Reset() {
	*x = WooCrossSwapOnDstChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WooCrossChainRouterV1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WooCrossSwapOnDstChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WooCrossSwapOnDstChain) ProtoMessage() {}

func (x *WooCrossSwapOnDstChain) ProtoReflect() protoreflect.Message {
	mi := &file_WooCrossChainRouterV1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WooCrossSwapOnDstChain.ProtoReflect.Descriptor instead.
func (*WooCrossSwapOnDstChain) Descriptor() ([]byte, []int) {
	return file_WooCrossChainRouterV1_proto_rawDescGZIP(), []int{1}
}

func (x *WooCrossSwapOnDstChain) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *WooCrossSwapOnDstChain) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *WooCrossSwapOnDstChain) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetRefId() []byte {
	if x != nil {
		return x.RefId
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetBridgedToken() []byte {
	if x != nil {
		return x.BridgedToken
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetBridgedAmount() []byte {
	if x != nil {
		return x.BridgedAmount
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetToToken() []byte {
	if x != nil {
		return x.ToToken
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetRealToToken() []byte {
	if x != nil {
		return x.RealToToken
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetMinToAmount() []byte {
	if x != nil {
		return x.MinToAmount
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetRealToAmount() []byte {
	if x != nil {
		return x.RealToAmount
	}
	return nil
}

func (x *WooCrossSwapOnDstChain) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

type OwnershipTransferred struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber     uint64                 `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Index           uint64                 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	TxHash          []byte                 `protobuf:"bytes,4,opt,name=txHash,proto3" json:"txHash,omitempty"`
	PreviousOwner   []byte                 `protobuf:"bytes,5,opt,name=previousOwner,proto3" json:"previousOwner,omitempty"`
	NewOwner        []byte                 `protobuf:"bytes,6,opt,name=newOwner,proto3" json:"newOwner,omitempty"`
	ContractAddress string                 `protobuf:"bytes,7,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (x *OwnershipTransferred) Reset() {
	*x = OwnershipTransferred{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WooCrossChainRouterV1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnershipTransferred) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnershipTransferred) ProtoMessage() {}

func (x *OwnershipTransferred) ProtoReflect() protoreflect.Message {
	mi := &file_WooCrossChainRouterV1_proto_msgTypes[2]
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
	return file_WooCrossChainRouterV1_proto_rawDescGZIP(), []int{2}
}

func (x *OwnershipTransferred) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *OwnershipTransferred) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *OwnershipTransferred) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *OwnershipTransferred) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
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

func (x *OwnershipTransferred) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

var File_WooCrossChainRouterV1_proto protoreflect.FileDescriptor

var file_WooCrossChainRouterV1_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x57, 0x6f, 0x6f, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x56, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x57,
	0x6f, 0x6f, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x56, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x16, 0x57, 0x6f, 0x6f, 0x43, 0x72, 0x6f,
	0x73, 0x73, 0x53, 0x77, 0x61, 0x70, 0x4f, 0x6e, 0x53, 0x72, 0x63, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x66, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x65, 0x66,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x66, 0x72,
	0x6f, 0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x6c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x65, 0x61, 0x6c, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0xc8, 0x03, 0x0a, 0x16, 0x57, 0x6f, 0x6f, 0x43, 0x72, 0x6f, 0x73,
	0x73, 0x53, 0x77, 0x61, 0x70, 0x4f, 0x6e, 0x44, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x66, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a,
	0x0d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xfe, 0x01, 0x0a, 0x14, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x02, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65,
	0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x65,
	0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x61, 0x6b, 0x6a, 0x69, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x77, 0x6f, 0x6f, 0x66, 0x69, 0x5f, 0x70, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2f, 0x57, 0x6f, 0x6f, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_WooCrossChainRouterV1_proto_rawDescOnce sync.Once
	file_WooCrossChainRouterV1_proto_rawDescData = file_WooCrossChainRouterV1_proto_rawDesc
)

func file_WooCrossChainRouterV1_proto_rawDescGZIP() []byte {
	file_WooCrossChainRouterV1_proto_rawDescOnce.Do(func() {
		file_WooCrossChainRouterV1_proto_rawDescData = protoimpl.X.CompressGZIP(file_WooCrossChainRouterV1_proto_rawDescData)
	})
	return file_WooCrossChainRouterV1_proto_rawDescData
}

var file_WooCrossChainRouterV1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_WooCrossChainRouterV1_proto_goTypes = []interface{}{
	(*WooCrossSwapOnSrcChain)(nil), // 0: WooCrossChainRouterV1.WooCrossSwapOnSrcChain
	(*WooCrossSwapOnDstChain)(nil), // 1: WooCrossChainRouterV1.WooCrossSwapOnDstChain
	(*OwnershipTransferred)(nil),   // 2: WooCrossChainRouterV1.OwnershipTransferred
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_WooCrossChainRouterV1_proto_depIdxs = []int32{
	3, // 0: WooCrossChainRouterV1.WooCrossSwapOnSrcChain.ts:type_name -> google.protobuf.Timestamp
	3, // 1: WooCrossChainRouterV1.WooCrossSwapOnDstChain.ts:type_name -> google.protobuf.Timestamp
	3, // 2: WooCrossChainRouterV1.OwnershipTransferred.ts:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_WooCrossChainRouterV1_proto_init() }
func file_WooCrossChainRouterV1_proto_init() {
	if File_WooCrossChainRouterV1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WooCrossChainRouterV1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WooCrossSwapOnSrcChain); i {
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
		file_WooCrossChainRouterV1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WooCrossSwapOnDstChain); i {
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
		file_WooCrossChainRouterV1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_WooCrossChainRouterV1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WooCrossChainRouterV1_proto_goTypes,
		DependencyIndexes: file_WooCrossChainRouterV1_proto_depIdxs,
		MessageInfos:      file_WooCrossChainRouterV1_proto_msgTypes,
	}.Build()
	File_WooCrossChainRouterV1_proto = out.File
	file_WooCrossChainRouterV1_proto_rawDesc = nil
	file_WooCrossChainRouterV1_proto_goTypes = nil
	file_WooCrossChainRouterV1_proto_depIdxs = nil
}
