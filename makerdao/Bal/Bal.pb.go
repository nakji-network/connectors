// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: Bal/Bal.proto

package Bal

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
		mi := &file_Bal_Bal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approval) ProtoMessage() {}

func (x *Approval) ProtoReflect() protoreflect.Message {
	mi := &file_Bal_Bal_proto_msgTypes[0]
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
	return file_Bal_Bal_proto_rawDescGZIP(), []int{0}
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

type RoleGranted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Role    []byte                 `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`       //	bytes32
	Account []byte                 `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"` //	address
	Sender  []byte                 `protobuf:"bytes,4,opt,name=Sender,proto3" json:"Sender,omitempty"`   //	address
}

func (x *RoleGranted) Reset() {
	*x = RoleGranted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Bal_Bal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGranted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGranted) ProtoMessage() {}

func (x *RoleGranted) ProtoReflect() protoreflect.Message {
	mi := &file_Bal_Bal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGranted.ProtoReflect.Descriptor instead.
func (*RoleGranted) Descriptor() ([]byte, []int) {
	return file_Bal_Bal_proto_rawDescGZIP(), []int{1}
}

func (x *RoleGranted) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *RoleGranted) GetRole() []byte {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *RoleGranted) GetAccount() []byte {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *RoleGranted) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

type RoleRevoked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Role    []byte                 `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`       //	bytes32
	Account []byte                 `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"` //	address
	Sender  []byte                 `protobuf:"bytes,4,opt,name=Sender,proto3" json:"Sender,omitempty"`   //	address
}

func (x *RoleRevoked) Reset() {
	*x = RoleRevoked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Bal_Bal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleRevoked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleRevoked) ProtoMessage() {}

func (x *RoleRevoked) ProtoReflect() protoreflect.Message {
	mi := &file_Bal_Bal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleRevoked.ProtoReflect.Descriptor instead.
func (*RoleRevoked) Descriptor() ([]byte, []int) {
	return file_Bal_Bal_proto_rawDescGZIP(), []int{2}
}

func (x *RoleRevoked) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *RoleRevoked) GetRole() []byte {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *RoleRevoked) GetAccount() []byte {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *RoleRevoked) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Id []byte                 `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"` //	uint256
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Bal_Bal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_Bal_Bal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_Bal_Bal_proto_rawDescGZIP(), []int{3}
}

func (x *Snapshot) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Snapshot) GetId() []byte {
	if x != nil {
		return x.Id
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
		mi := &file_Bal_Bal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Bal_Bal_proto_msgTypes[4]
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
	return file_Bal_Bal_proto_rawDescGZIP(), []int{4}
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

var File_Bal_Bal_proto protoreflect.FileDescriptor

var file_Bal_Bal_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x42, 0x61, 0x6c, 0x2f, 0x42, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x42, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x22, 0x7f, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x64, 0x22, 0x70, 0x0a,
	0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x36, 0x5a, 0x34, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72,
	0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x42, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Bal_Bal_proto_rawDescOnce sync.Once
	file_Bal_Bal_proto_rawDescData = file_Bal_Bal_proto_rawDesc
)

func file_Bal_Bal_proto_rawDescGZIP() []byte {
	file_Bal_Bal_proto_rawDescOnce.Do(func() {
		file_Bal_Bal_proto_rawDescData = protoimpl.X.CompressGZIP(file_Bal_Bal_proto_rawDescData)
	})
	return file_Bal_Bal_proto_rawDescData
}

var file_Bal_Bal_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Bal_Bal_proto_goTypes = []interface{}{
	(*Approval)(nil),              // 0: Bal.Approval
	(*RoleGranted)(nil),           // 1: Bal.RoleGranted
	(*RoleRevoked)(nil),           // 2: Bal.RoleRevoked
	(*Snapshot)(nil),              // 3: Bal.Snapshot
	(*Transfer)(nil),              // 4: Bal.Transfer
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_Bal_Bal_proto_depIdxs = []int32{
	5, // 0: Bal.Approval.ts:type_name -> google.protobuf.Timestamp
	5, // 1: Bal.RoleGranted.ts:type_name -> google.protobuf.Timestamp
	5, // 2: Bal.RoleRevoked.ts:type_name -> google.protobuf.Timestamp
	5, // 3: Bal.Snapshot.ts:type_name -> google.protobuf.Timestamp
	5, // 4: Bal.Transfer.ts:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Bal_Bal_proto_init() }
func file_Bal_Bal_proto_init() {
	if File_Bal_Bal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Bal_Bal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Bal_Bal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGranted); i {
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
		file_Bal_Bal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleRevoked); i {
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
		file_Bal_Bal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
		file_Bal_Bal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Bal_Bal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Bal_Bal_proto_goTypes,
		DependencyIndexes: file_Bal_Bal_proto_depIdxs,
		MessageInfos:      file_Bal_Bal_proto_msgTypes,
	}.Build()
	File_Bal_Bal_proto = out.File
	file_Bal_Bal_proto_rawDesc = nil
	file_Bal_Bal_proto_goTypes = nil
	file_Bal_Bal_proto_depIdxs = nil
}
