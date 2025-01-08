// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.20.3
// source: resources/sync/activity.proto

package sync

import (
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

// Connect an identifier/license to the provider with the specified external id
// (e.g., auto discord social connect on server join)
type UserOAuth2Conn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProviderName  string                 `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	Identifier    string                 `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	ExternalId    string                 `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserOAuth2Conn) Reset() {
	*x = UserOAuth2Conn{}
	mi := &file_resources_sync_activity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserOAuth2Conn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOAuth2Conn) ProtoMessage() {}

func (x *UserOAuth2Conn) ProtoReflect() protoreflect.Message {
	mi := &file_resources_sync_activity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOAuth2Conn.ProtoReflect.Descriptor instead.
func (*UserOAuth2Conn) Descriptor() ([]byte, []int) {
	return file_resources_sync_activity_proto_rawDescGZIP(), []int{0}
}

func (x *UserOAuth2Conn) GetProviderName() string {
	if x != nil {
		return x.ProviderName
	}
	return ""
}

func (x *UserOAuth2Conn) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *UserOAuth2Conn) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

var File_resources_sync_activity_proto protoreflect.FileDescriptor

var file_resources_sync_activity_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x22,
	0x76, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70,
	0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x3b, 0x73, 0x79, 0x6e, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resources_sync_activity_proto_rawDescOnce sync.Once
	file_resources_sync_activity_proto_rawDescData = file_resources_sync_activity_proto_rawDesc
)

func file_resources_sync_activity_proto_rawDescGZIP() []byte {
	file_resources_sync_activity_proto_rawDescOnce.Do(func() {
		file_resources_sync_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_sync_activity_proto_rawDescData)
	})
	return file_resources_sync_activity_proto_rawDescData
}

var file_resources_sync_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_sync_activity_proto_goTypes = []any{
	(*UserOAuth2Conn)(nil), // 0: resources.sync.UserOAuth2Conn
}
var file_resources_sync_activity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resources_sync_activity_proto_init() }
func file_resources_sync_activity_proto_init() {
	if File_resources_sync_activity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_sync_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_sync_activity_proto_goTypes,
		DependencyIndexes: file_resources_sync_activity_proto_depIdxs,
		MessageInfos:      file_resources_sync_activity_proto_msgTypes,
	}.Build()
	File_resources_sync_activity_proto = out.File
	file_resources_sync_activity_proto_rawDesc = nil
	file_resources_sync_activity_proto_goTypes = nil
	file_resources_sync_activity_proto_depIdxs = nil
}
