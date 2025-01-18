// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.20.3
// source: resources/users/labels.proto

package users

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

type CitizenLabels struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*CitizenLabel        `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CitizenLabels) Reset() {
	*x = CitizenLabels{}
	mi := &file_resources_users_labels_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CitizenLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitizenLabels) ProtoMessage() {}

func (x *CitizenLabels) ProtoReflect() protoreflect.Message {
	mi := &file_resources_users_labels_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitizenLabels.ProtoReflect.Descriptor instead.
func (*CitizenLabels) Descriptor() ([]byte, []int) {
	return file_resources_users_labels_proto_rawDescGZIP(), []int{0}
}

func (x *CitizenLabels) GetList() []*CitizenLabel {
	if x != nil {
		return x.List
	}
	return nil
}

type CitizenLabel struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	Job   *string                `protobuf:"bytes,2,opt,name=job,proto3,oneof" json:"job,omitempty"`
	Name  string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// @sanitize: method=StripTags
	Color         string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CitizenLabel) Reset() {
	*x = CitizenLabel{}
	mi := &file_resources_users_labels_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CitizenLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitizenLabel) ProtoMessage() {}

func (x *CitizenLabel) ProtoReflect() protoreflect.Message {
	mi := &file_resources_users_labels_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitizenLabel.ProtoReflect.Descriptor instead.
func (*CitizenLabel) Descriptor() ([]byte, []int) {
	return file_resources_users_labels_proto_rawDescGZIP(), []int{1}
}

func (x *CitizenLabel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CitizenLabel) GetJob() string {
	if x != nil && x.Job != nil {
		return *x.Job
	}
	return ""
}

func (x *CitizenLabel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CitizenLabel) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

var File_resources_users_labels_proto protoreflect.FileDescriptor

var file_resources_users_labels_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0d, 0x43, 0x69, 0x74, 0x69,
	0x7a, 0x65, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65,
	0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x0a,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x43, 0x69, 0x74, 0x69, 0x7a,
	0x65, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x03, 0x6a,
	0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18,
	0x14, 0x48, 0x00, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x18, 0x30, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x32, 0x11,
	0x5e, 0x23, 0x5b, 0x41, 0x2d, 0x46, 0x61, 0x2d, 0x66, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36, 0x7d,
	0x24, 0x98, 0x01, 0x07, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x6a, 0x6f, 0x62, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_users_labels_proto_rawDescOnce sync.Once
	file_resources_users_labels_proto_rawDescData = file_resources_users_labels_proto_rawDesc
)

func file_resources_users_labels_proto_rawDescGZIP() []byte {
	file_resources_users_labels_proto_rawDescOnce.Do(func() {
		file_resources_users_labels_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_users_labels_proto_rawDescData)
	})
	return file_resources_users_labels_proto_rawDescData
}

var file_resources_users_labels_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_users_labels_proto_goTypes = []any{
	(*CitizenLabels)(nil), // 0: resources.users.CitizenLabels
	(*CitizenLabel)(nil),  // 1: resources.users.CitizenLabel
}
var file_resources_users_labels_proto_depIdxs = []int32{
	1, // 0: resources.users.CitizenLabels.list:type_name -> resources.users.CitizenLabel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resources_users_labels_proto_init() }
func file_resources_users_labels_proto_init() {
	if File_resources_users_labels_proto != nil {
		return
	}
	file_resources_users_labels_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_users_labels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_users_labels_proto_goTypes,
		DependencyIndexes: file_resources_users_labels_proto_depIdxs,
		MessageInfos:      file_resources_users_labels_proto_msgTypes,
	}.Build()
	File_resources_users_labels_proto = out.File
	file_resources_users_labels_proto_rawDesc = nil
	file_resources_users_labels_proto_goTypes = nil
	file_resources_users_labels_proto_depIdxs = nil
}