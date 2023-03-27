// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: MCDCropperImp/MCDCropperImp.proto

package MCDCropperImp

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

type Hope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	From []byte                 `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"` //	address
	To   []byte                 `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`     //	address
}

func (x *Hope) Reset() {
	*x = Hope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCDCropperImp_MCDCropperImp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hope) ProtoMessage() {}

func (x *Hope) ProtoReflect() protoreflect.Message {
	mi := &file_MCDCropperImp_MCDCropperImp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hope.ProtoReflect.Descriptor instead.
func (*Hope) Descriptor() ([]byte, []int) {
	return file_MCDCropperImp_MCDCropperImp_proto_rawDescGZIP(), []int{0}
}

func (x *Hope) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Hope) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Hope) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

type NewProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Usr []byte                 `protobuf:"bytes,2,opt,name=Usr,proto3" json:"Usr,omitempty"` //	address
	Urp []byte                 `protobuf:"bytes,3,opt,name=Urp,proto3" json:"Urp,omitempty"` //	address
}

func (x *NewProxy) Reset() {
	*x = NewProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCDCropperImp_MCDCropperImp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProxy) ProtoMessage() {}

func (x *NewProxy) ProtoReflect() protoreflect.Message {
	mi := &file_MCDCropperImp_MCDCropperImp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProxy.ProtoReflect.Descriptor instead.
func (*NewProxy) Descriptor() ([]byte, []int) {
	return file_MCDCropperImp_MCDCropperImp_proto_rawDescGZIP(), []int{1}
}

func (x *NewProxy) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *NewProxy) GetUsr() []byte {
	if x != nil {
		return x.Usr
	}
	return nil
}

func (x *NewProxy) GetUrp() []byte {
	if x != nil {
		return x.Urp
	}
	return nil
}

type Nope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	From []byte                 `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"` //	address
	To   []byte                 `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`     //	address
}

func (x *Nope) Reset() {
	*x = Nope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCDCropperImp_MCDCropperImp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nope) ProtoMessage() {}

func (x *Nope) ProtoReflect() protoreflect.Message {
	mi := &file_MCDCropperImp_MCDCropperImp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nope.ProtoReflect.Descriptor instead.
func (*Nope) Descriptor() ([]byte, []int) {
	return file_MCDCropperImp_MCDCropperImp_proto_rawDescGZIP(), []int{2}
}

func (x *Nope) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Nope) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Nope) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

var File_MCDCropperImp_MCDCropperImp_proto protoreflect.FileDescriptor

var file_MCDCropperImp_MCDCropperImp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x4d, 0x43, 0x44, 0x43, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x49, 0x6d, 0x70, 0x2f,
	0x4d, 0x43, 0x44, 0x43, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x49, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x4d, 0x43, 0x44, 0x43, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x49,
	0x6d, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x04, 0x48, 0x6f, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x6f, 0x22, 0x5a, 0x0a, 0x08, 0x4e,
	0x65, 0x77, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x55, 0x73, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x55, 0x72, 0x70, 0x22, 0x56, 0x0a, 0x04, 0x4e, 0x6f, 0x70, 0x65, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x6f, 0x42,
	0x40, 0x5a, 0x3e, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72,
	0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x4d, 0x43, 0x44, 0x43, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x49, 0x6d,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MCDCropperImp_MCDCropperImp_proto_rawDescOnce sync.Once
	file_MCDCropperImp_MCDCropperImp_proto_rawDescData = file_MCDCropperImp_MCDCropperImp_proto_rawDesc
)

func file_MCDCropperImp_MCDCropperImp_proto_rawDescGZIP() []byte {
	file_MCDCropperImp_MCDCropperImp_proto_rawDescOnce.Do(func() {
		file_MCDCropperImp_MCDCropperImp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MCDCropperImp_MCDCropperImp_proto_rawDescData)
	})
	return file_MCDCropperImp_MCDCropperImp_proto_rawDescData
}

var file_MCDCropperImp_MCDCropperImp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_MCDCropperImp_MCDCropperImp_proto_goTypes = []interface{}{
	(*Hope)(nil),                  // 0: MCDCropperImp.Hope
	(*NewProxy)(nil),              // 1: MCDCropperImp.NewProxy
	(*Nope)(nil),                  // 2: MCDCropperImp.Nope
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_MCDCropperImp_MCDCropperImp_proto_depIdxs = []int32{
	3, // 0: MCDCropperImp.Hope.ts:type_name -> google.protobuf.Timestamp
	3, // 1: MCDCropperImp.NewProxy.ts:type_name -> google.protobuf.Timestamp
	3, // 2: MCDCropperImp.Nope.ts:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_MCDCropperImp_MCDCropperImp_proto_init() }
func file_MCDCropperImp_MCDCropperImp_proto_init() {
	if File_MCDCropperImp_MCDCropperImp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MCDCropperImp_MCDCropperImp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hope); i {
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
		file_MCDCropperImp_MCDCropperImp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProxy); i {
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
		file_MCDCropperImp_MCDCropperImp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nope); i {
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
			RawDescriptor: file_MCDCropperImp_MCDCropperImp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MCDCropperImp_MCDCropperImp_proto_goTypes,
		DependencyIndexes: file_MCDCropperImp_MCDCropperImp_proto_depIdxs,
		MessageInfos:      file_MCDCropperImp_MCDCropperImp_proto_msgTypes,
	}.Build()
	File_MCDCropperImp_MCDCropperImp_proto = out.File
	file_MCDCropperImp_MCDCropperImp_proto_rawDesc = nil
	file_MCDCropperImp_MCDCropperImp_proto_goTypes = nil
	file_MCDCropperImp_MCDCropperImp_proto_depIdxs = nil
}