// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: resources/dispatch/settings.proto

package dispatch

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CENTRUM_MODE int32

const (
	CENTRUM_MODE_MANUAL           CENTRUM_MODE = 0
	CENTRUM_MODE_CENTRAL_COMMAND  CENTRUM_MODE = 1
	CENTRUM_MODE_AUTO_ROUND_ROBIN CENTRUM_MODE = 2
)

// Enum value maps for CENTRUM_MODE.
var (
	CENTRUM_MODE_name = map[int32]string{
		0: "MANUAL",
		1: "CENTRAL_COMMAND",
		2: "AUTO_ROUND_ROBIN",
	}
	CENTRUM_MODE_value = map[string]int32{
		"MANUAL":           0,
		"CENTRAL_COMMAND":  1,
		"AUTO_ROUND_ROBIN": 2,
	}
)

func (x CENTRUM_MODE) Enum() *CENTRUM_MODE {
	p := new(CENTRUM_MODE)
	*p = x
	return p
}

func (x CENTRUM_MODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CENTRUM_MODE) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_dispatch_settings_proto_enumTypes[0].Descriptor()
}

func (CENTRUM_MODE) Type() protoreflect.EnumType {
	return &file_resources_dispatch_settings_proto_enumTypes[0]
}

func (x CENTRUM_MODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CENTRUM_MODE.Descriptor instead.
func (CENTRUM_MODE) EnumDescriptor() ([]byte, []int) {
	return file_resources_dispatch_settings_proto_rawDescGZIP(), []int{0}
}

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job          string       `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Enabled      bool         `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Mode         CENTRUM_MODE `protobuf:"varint,3,opt,name=mode,proto3,enum=resources.dispatch.CENTRUM_MODE" json:"mode,omitempty"`
	FallbackMode CENTRUM_MODE `protobuf:"varint,4,opt,name=fallback_mode,json=fallbackMode,proto3,enum=resources.dispatch.CENTRUM_MODE" json:"fallback_mode,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_dispatch_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dispatch_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_resources_dispatch_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Settings) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Settings) GetMode() CENTRUM_MODE {
	if x != nil {
		return x.Mode
	}
	return CENTRUM_MODE_MANUAL
}

func (x *Settings) GetFallbackMode() CENTRUM_MODE {
	if x != nil {
		return x.FallbackMode
	}
	return CENTRUM_MODE_MANUAL
}

var File_resources_dispatch_settings_proto protoreflect.FileDescriptor

var file_resources_dispatch_settings_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd0, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x19, 0x0a,
	0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x3e, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x55, 0x4d, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x43,
	0x45, 0x4e, 0x54, 0x52, 0x55, 0x4d, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4d,
	0x6f, 0x64, 0x65, 0x2a, 0x45, 0x0a, 0x0c, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x55, 0x4d, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x43, 0x45, 0x4e, 0x54, 0x52, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x4e, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x52, 0x4f, 0x55,
	0x4e, 0x44, 0x5f, 0x52, 0x4f, 0x42, 0x49, 0x4e, 0x10, 0x02, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74,
	0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x3b, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_dispatch_settings_proto_rawDescOnce sync.Once
	file_resources_dispatch_settings_proto_rawDescData = file_resources_dispatch_settings_proto_rawDesc
)

func file_resources_dispatch_settings_proto_rawDescGZIP() []byte {
	file_resources_dispatch_settings_proto_rawDescOnce.Do(func() {
		file_resources_dispatch_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_dispatch_settings_proto_rawDescData)
	})
	return file_resources_dispatch_settings_proto_rawDescData
}

var file_resources_dispatch_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_dispatch_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_dispatch_settings_proto_goTypes = []interface{}{
	(CENTRUM_MODE)(0), // 0: resources.dispatch.CENTRUM_MODE
	(*Settings)(nil),  // 1: resources.dispatch.Settings
}
var file_resources_dispatch_settings_proto_depIdxs = []int32{
	0, // 0: resources.dispatch.Settings.mode:type_name -> resources.dispatch.CENTRUM_MODE
	0, // 1: resources.dispatch.Settings.fallback_mode:type_name -> resources.dispatch.CENTRUM_MODE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resources_dispatch_settings_proto_init() }
func file_resources_dispatch_settings_proto_init() {
	if File_resources_dispatch_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_dispatch_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
			RawDescriptor: file_resources_dispatch_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_dispatch_settings_proto_goTypes,
		DependencyIndexes: file_resources_dispatch_settings_proto_depIdxs,
		EnumInfos:         file_resources_dispatch_settings_proto_enumTypes,
		MessageInfos:      file_resources_dispatch_settings_proto_msgTypes,
	}.Build()
	File_resources_dispatch_settings_proto = out.File
	file_resources_dispatch_settings_proto_rawDesc = nil
	file_resources_dispatch_settings_proto_goTypes = nil
	file_resources_dispatch_settings_proto_depIdxs = nil
}
