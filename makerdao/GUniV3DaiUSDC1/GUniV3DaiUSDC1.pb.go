// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: GUniV3DaiUSDC1/GUniV3DaiUSDC1.proto

package GUniV3DaiUSDC1

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

type ProxyAdminTransferred struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	PreviousAdmin []byte                 `protobuf:"bytes,2,opt,name=PreviousAdmin,proto3" json:"PreviousAdmin,omitempty"` //	address
	NewAdmin      []byte                 `protobuf:"bytes,3,opt,name=NewAdmin,proto3" json:"NewAdmin,omitempty"`           //	address
}

func (x *ProxyAdminTransferred) Reset() {
	*x = ProxyAdminTransferred{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyAdminTransferred) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyAdminTransferred) ProtoMessage() {}

func (x *ProxyAdminTransferred) ProtoReflect() protoreflect.Message {
	mi := &file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyAdminTransferred.ProtoReflect.Descriptor instead.
func (*ProxyAdminTransferred) Descriptor() ([]byte, []int) {
	return file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyAdminTransferred) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ProxyAdminTransferred) GetPreviousAdmin() []byte {
	if x != nil {
		return x.PreviousAdmin
	}
	return nil
}

func (x *ProxyAdminTransferred) GetNewAdmin() []byte {
	if x != nil {
		return x.NewAdmin
	}
	return nil
}

type ProxyImplementationUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts                     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	PreviousImplementation []byte                 `protobuf:"bytes,2,opt,name=PreviousImplementation,proto3" json:"PreviousImplementation,omitempty"` //	address
	NewImplementation      []byte                 `protobuf:"bytes,3,opt,name=NewImplementation,proto3" json:"NewImplementation,omitempty"`           //	address
}

func (x *ProxyImplementationUpdated) Reset() {
	*x = ProxyImplementationUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyImplementationUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyImplementationUpdated) ProtoMessage() {}

func (x *ProxyImplementationUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyImplementationUpdated.ProtoReflect.Descriptor instead.
func (*ProxyImplementationUpdated) Descriptor() ([]byte, []int) {
	return file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyImplementationUpdated) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ProxyImplementationUpdated) GetPreviousImplementation() []byte {
	if x != nil {
		return x.PreviousImplementation
	}
	return nil
}

func (x *ProxyImplementationUpdated) GetNewImplementation() []byte {
	if x != nil {
		return x.NewImplementation
	}
	return nil
}

var File_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto protoreflect.FileDescriptor

var file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDesc = []byte{
	0x0a, 0x23, 0x47, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x44, 0x61, 0x69, 0x55, 0x53, 0x44, 0x43, 0x31,
	0x2f, 0x47, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x44, 0x61, 0x69, 0x55, 0x53, 0x44, 0x43, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x47, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x44, 0x61, 0x69,
	0x55, 0x53, 0x44, 0x43, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xae,
	0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x4e, 0x65,
	0x77, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x41, 0x5a, 0x3f, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72,
	0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x47, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x44, 0x61, 0x69, 0x55, 0x53, 0x44,
	0x43, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescOnce sync.Once
	file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescData = file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDesc
)

func file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescGZIP() []byte {
	file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescOnce.Do(func() {
		file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescData = protoimpl.X.CompressGZIP(file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescData)
	})
	return file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDescData
}

var file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_goTypes = []interface{}{
	(*ProxyAdminTransferred)(nil),      // 0: GUniV3DaiUSDC1.ProxyAdminTransferred
	(*ProxyImplementationUpdated)(nil), // 1: GUniV3DaiUSDC1.ProxyImplementationUpdated
	(*timestamppb.Timestamp)(nil),      // 2: google.protobuf.Timestamp
}
var file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_depIdxs = []int32{
	2, // 0: GUniV3DaiUSDC1.ProxyAdminTransferred.ts:type_name -> google.protobuf.Timestamp
	2, // 1: GUniV3DaiUSDC1.ProxyImplementationUpdated.ts:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_init() }
func file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_init() {
	if File_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyAdminTransferred); i {
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
		file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyImplementationUpdated); i {
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
			RawDescriptor: file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_goTypes,
		DependencyIndexes: file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_depIdxs,
		MessageInfos:      file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_msgTypes,
	}.Build()
	File_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto = out.File
	file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_rawDesc = nil
	file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_goTypes = nil
	file_GUniV3DaiUSDC1_GUniV3DaiUSDC1_proto_depIdxs = nil
}
