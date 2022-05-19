// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: connectors/source/makerdao/smart-contracts/CLIPPER_MOM/CLIPPER_MOM.proto

package CLIPPER_MOM

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

type SetOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	OldOwner []byte                 `protobuf:"bytes,2,opt,name=OldOwner,proto3" json:"OldOwner,omitempty"` //	address
	NewOwner []byte                 `protobuf:"bytes,3,opt,name=NewOwner,proto3" json:"NewOwner,omitempty"` //	address
}

func (x *SetOwner) Reset() {
	*x = SetOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOwner) ProtoMessage() {}

func (x *SetOwner) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOwner.ProtoReflect.Descriptor instead.
func (*SetOwner) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescGZIP(), []int{0}
}

func (x *SetOwner) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *SetOwner) GetOldOwner() []byte {
	if x != nil {
		return x.OldOwner
	}
	return nil
}

func (x *SetOwner) GetNewOwner() []byte {
	if x != nil {
		return x.NewOwner
	}
	return nil
}

type SetAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	OldAuthority []byte                 `protobuf:"bytes,2,opt,name=OldAuthority,proto3" json:"OldAuthority,omitempty"` //	address
	NewAuthority []byte                 `protobuf:"bytes,3,opt,name=NewAuthority,proto3" json:"NewAuthority,omitempty"` //	address
}

func (x *SetAuthority) Reset() {
	*x = SetAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthority) ProtoMessage() {}

func (x *SetAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAuthority.ProtoReflect.Descriptor instead.
func (*SetAuthority) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescGZIP(), []int{1}
}

func (x *SetAuthority) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *SetAuthority) GetOldAuthority() []byte {
	if x != nil {
		return x.OldAuthority
	}
	return nil
}

func (x *SetAuthority) GetNewAuthority() []byte {
	if x != nil {
		return x.NewAuthority
	}
	return nil
}

type SetBreaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Clip  []byte                 `protobuf:"bytes,2,opt,name=Clip,proto3" json:"Clip,omitempty"`   //	address
	Level []byte                 `protobuf:"bytes,3,opt,name=Level,proto3" json:"Level,omitempty"` //	uint256
}

func (x *SetBreaker) Reset() {
	*x = SetBreaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBreaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBreaker) ProtoMessage() {}

func (x *SetBreaker) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBreaker.ProtoReflect.Descriptor instead.
func (*SetBreaker) Descriptor() ([]byte, []int) {
	return file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescGZIP(), []int{2}
}

func (x *SetBreaker) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *SetBreaker) GetClip() []byte {
	if x != nil {
		return x.Clip
	}
	return nil
}

func (x *SetBreaker) GetLevel() []byte {
	if x != nil {
		return x.Level
	}
	return nil
}

var File_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto protoreflect.FileDescriptor

var file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDesc = []byte{
	0x0a, 0x48, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x43, 0x4c, 0x49,
	0x50, 0x50, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x4d, 0x2f, 0x43, 0x4c, 0x49, 0x50, 0x50, 0x45, 0x52,
	0x5f, 0x4d, 0x4f, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x43, 0x4c, 0x49, 0x50,
	0x50, 0x45, 0x52, 0x5f, 0x4d, 0x4f, 0x4d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x4f, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x4e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x4e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x6c, 0x64, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x4f, 0x6c, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x65, 0x77,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x62, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6c, 0x69, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x43, 0x6c, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x42, 0x3e, 0x5a, 0x3c, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b,
	0x65, 0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x43, 0x4c, 0x49, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x4d, 0x4f,
	0x4d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescOnce sync.Once
	file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescData = file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDesc
)

func file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescGZIP() []byte {
	file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescOnce.Do(func() {
		file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescData)
	})
	return file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDescData
}

var file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_goTypes = []interface{}{
	(*SetOwner)(nil),              // 0: CLIPPER_MOM.SetOwner
	(*SetAuthority)(nil),          // 1: CLIPPER_MOM.SetAuthority
	(*SetBreaker)(nil),            // 2: CLIPPER_MOM.SetBreaker
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_depIdxs = []int32{
	3, // 0: CLIPPER_MOM.SetOwner.ts:type_name -> google.protobuf.Timestamp
	3, // 1: CLIPPER_MOM.SetAuthority.ts:type_name -> google.protobuf.Timestamp
	3, // 2: CLIPPER_MOM.SetBreaker.ts:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_init() }
func file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_init() {
	if File_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOwner); i {
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
		file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAuthority); i {
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
		file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBreaker); i {
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
			RawDescriptor: file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_goTypes,
		DependencyIndexes: file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_depIdxs,
		MessageInfos:      file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_msgTypes,
	}.Build()
	File_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto = out.File
	file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_rawDesc = nil
	file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_goTypes = nil
	file_connectors_source_makerdao_smart_contracts_CLIPPER_MOM_CLIPPER_MOM_proto_depIdxs = nil
}
