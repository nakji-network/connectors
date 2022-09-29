// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: flow.proto

package flow

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Id       []byte                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ParentID []byte                 `protobuf:"bytes,3,opt,name=parentID,proto3" json:"parentID,omitempty"`
	Height   uint64                 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_flow_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Block) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Block) GetParentID() []byte {
	if x != nil {
		return x.ParentID
	}
	return nil
}

func (x *Block) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts                 *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Id                 []byte                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	BlockID            []byte                  `protobuf:"bytes,3,opt,name=blockID,proto3" json:"blockID,omitempty"`
	CollectionID       []byte                  `protobuf:"bytes,4,opt,name=collectionID,proto3" json:"collectionID,omitempty"`
	Script             []byte                  `protobuf:"bytes,5,opt,name=script,proto3" json:"script,omitempty"`
	Arguments          [][]byte                `protobuf:"bytes,6,rep,name=arguments,proto3" json:"arguments,omitempty"`
	ReferenceBlockID   []byte                  `protobuf:"bytes,7,opt,name=referenceBlockID,proto3" json:"referenceBlockID,omitempty"`
	GasLimit           uint64                  `protobuf:"varint,8,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	ProposalKey        *ProposalKey            `protobuf:"bytes,9,opt,name=proposalKey,proto3" json:"proposalKey,omitempty"`
	Payer              []byte                  `protobuf:"bytes,10,opt,name=payer,proto3" json:"payer,omitempty"`
	Authorizers        [][]byte                `protobuf:"bytes,11,rep,name=authorizers,proto3" json:"authorizers,omitempty"`
	PayloadSignatures  []*TransactionSignature `protobuf:"bytes,12,rep,name=payloadSignatures,proto3" json:"payloadSignatures,omitempty"`
	EnvelopeSignatures []*TransactionSignature `protobuf:"bytes,13,rep,name=envelopeSignatures,proto3" json:"envelopeSignatures,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_flow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_flow_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Transaction) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Transaction) GetBlockID() []byte {
	if x != nil {
		return x.BlockID
	}
	return nil
}

func (x *Transaction) GetCollectionID() []byte {
	if x != nil {
		return x.CollectionID
	}
	return nil
}

func (x *Transaction) GetScript() []byte {
	if x != nil {
		return x.Script
	}
	return nil
}

func (x *Transaction) GetArguments() [][]byte {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Transaction) GetReferenceBlockID() []byte {
	if x != nil {
		return x.ReferenceBlockID
	}
	return nil
}

func (x *Transaction) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Transaction) GetProposalKey() *ProposalKey {
	if x != nil {
		return x.ProposalKey
	}
	return nil
}

func (x *Transaction) GetPayer() []byte {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *Transaction) GetAuthorizers() [][]byte {
	if x != nil {
		return x.Authorizers
	}
	return nil
}

func (x *Transaction) GetPayloadSignatures() []*TransactionSignature {
	if x != nil {
		return x.PayloadSignatures
	}
	return nil
}

func (x *Transaction) GetEnvelopeSignatures() []*TransactionSignature {
	if x != nil {
		return x.EnvelopeSignatures
	}
	return nil
}

type ProposalKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	KeyIndex       int64  `protobuf:"varint,2,opt,name=keyIndex,proto3" json:"keyIndex,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,3,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
}

func (x *ProposalKey) Reset() {
	*x = ProposalKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalKey) ProtoMessage() {}

func (x *ProposalKey) ProtoReflect() protoreflect.Message {
	mi := &file_flow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalKey.ProtoReflect.Descriptor instead.
func (*ProposalKey) Descriptor() ([]byte, []int) {
	return file_flow_proto_rawDescGZIP(), []int{2}
}

func (x *ProposalKey) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ProposalKey) GetKeyIndex() int64 {
	if x != nil {
		return x.KeyIndex
	}
	return 0
}

func (x *ProposalKey) GetSequenceNumber() uint64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

type TransactionSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	SignerIndex int64  `protobuf:"varint,2,opt,name=signerIndex,proto3" json:"signerIndex,omitempty"`
	KeyIndex    int64  `protobuf:"varint,3,opt,name=keyIndex,proto3" json:"keyIndex,omitempty"`
	Signature   []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *TransactionSignature) Reset() {
	*x = TransactionSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionSignature) ProtoMessage() {}

func (x *TransactionSignature) ProtoReflect() protoreflect.Message {
	mi := &file_flow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionSignature.ProtoReflect.Descriptor instead.
func (*TransactionSignature) Descriptor() ([]byte, []int) {
	return file_flow_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionSignature) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TransactionSignature) GetSignerIndex() int64 {
	if x != nil {
		return x.SignerIndex
	}
	return 0
}

func (x *TransactionSignature) GetKeyIndex() int64 {
	if x != nil {
		return x.KeyIndex
	}
	return 0
}

func (x *TransactionSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_flow_proto protoreflect.FileDescriptor

var file_flow_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c,
	0x6f, 0x77, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x88, 0x04, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x10,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72,
	0x73, 0x12, 0x48, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x11, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x12, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x12, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x6b, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x61, 0x6b, 0x6a, 0x69, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_proto_rawDescOnce sync.Once
	file_flow_proto_rawDescData = file_flow_proto_rawDesc
)

func file_flow_proto_rawDescGZIP() []byte {
	file_flow_proto_rawDescOnce.Do(func() {
		file_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_proto_rawDescData)
	})
	return file_flow_proto_rawDescData
}

var file_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flow_proto_goTypes = []interface{}{
	(*Block)(nil),                 // 0: flow.Block
	(*Transaction)(nil),           // 1: flow.Transaction
	(*ProposalKey)(nil),           // 2: flow.ProposalKey
	(*TransactionSignature)(nil),  // 3: flow.TransactionSignature
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_flow_proto_depIdxs = []int32{
	4, // 0: flow.Block.ts:type_name -> google.protobuf.Timestamp
	4, // 1: flow.Transaction.ts:type_name -> google.protobuf.Timestamp
	2, // 2: flow.Transaction.proposalKey:type_name -> flow.ProposalKey
	3, // 3: flow.Transaction.payloadSignatures:type_name -> flow.TransactionSignature
	3, // 4: flow.Transaction.envelopeSignatures:type_name -> flow.TransactionSignature
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_flow_proto_init() }
func file_flow_proto_init() {
	if File_flow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_flow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_flow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalKey); i {
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
		file_flow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionSignature); i {
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
			RawDescriptor: file_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_proto_goTypes,
		DependencyIndexes: file_flow_proto_depIdxs,
		MessageInfos:      file_flow_proto_msgTypes,
	}.Build()
	File_flow_proto = out.File
	file_flow_proto_rawDesc = nil
	file_flow_proto_goTypes = nil
	file_flow_proto_depIdxs = nil
}
