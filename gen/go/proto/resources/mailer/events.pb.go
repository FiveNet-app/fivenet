// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.20.3
// source: resources/mailer/events.proto

package mailer

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

type MailerEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*MailerEvent_EmailUpdate
	//	*MailerEvent_EmailDelete
	//	*MailerEvent_EmailSettingsUpdated
	//	*MailerEvent_ThreadUpdate
	//	*MailerEvent_ThreadDelete
	//	*MailerEvent_ThreadStateUpdate
	//	*MailerEvent_MessageUpdate
	//	*MailerEvent_MessageDelete
	Data          isMailerEvent_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MailerEvent) Reset() {
	*x = MailerEvent{}
	mi := &file_resources_mailer_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailerEvent) ProtoMessage() {}

func (x *MailerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_resources_mailer_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailerEvent.ProtoReflect.Descriptor instead.
func (*MailerEvent) Descriptor() ([]byte, []int) {
	return file_resources_mailer_events_proto_rawDescGZIP(), []int{0}
}

func (x *MailerEvent) GetData() isMailerEvent_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MailerEvent) GetEmailUpdate() *Email {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_EmailUpdate); ok {
			return x.EmailUpdate
		}
	}
	return nil
}

func (x *MailerEvent) GetEmailDelete() uint64 {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_EmailDelete); ok {
			return x.EmailDelete
		}
	}
	return 0
}

func (x *MailerEvent) GetEmailSettingsUpdated() *EmailSettings {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_EmailSettingsUpdated); ok {
			return x.EmailSettingsUpdated
		}
	}
	return nil
}

func (x *MailerEvent) GetThreadUpdate() *Thread {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_ThreadUpdate); ok {
			return x.ThreadUpdate
		}
	}
	return nil
}

func (x *MailerEvent) GetThreadDelete() uint64 {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_ThreadDelete); ok {
			return x.ThreadDelete
		}
	}
	return 0
}

func (x *MailerEvent) GetThreadStateUpdate() *ThreadState {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_ThreadStateUpdate); ok {
			return x.ThreadStateUpdate
		}
	}
	return nil
}

func (x *MailerEvent) GetMessageUpdate() *Message {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_MessageUpdate); ok {
			return x.MessageUpdate
		}
	}
	return nil
}

func (x *MailerEvent) GetMessageDelete() uint64 {
	if x != nil {
		if x, ok := x.Data.(*MailerEvent_MessageDelete); ok {
			return x.MessageDelete
		}
	}
	return 0
}

type isMailerEvent_Data interface {
	isMailerEvent_Data()
}

type MailerEvent_EmailUpdate struct {
	EmailUpdate *Email `protobuf:"bytes,1,opt,name=email_update,json=emailUpdate,proto3,oneof"`
}

type MailerEvent_EmailDelete struct {
	EmailDelete uint64 `protobuf:"varint,2,opt,name=email_delete,json=emailDelete,proto3,oneof"`
}

type MailerEvent_EmailSettingsUpdated struct {
	EmailSettingsUpdated *EmailSettings `protobuf:"bytes,3,opt,name=email_settings_updated,json=emailSettingsUpdated,proto3,oneof"`
}

type MailerEvent_ThreadUpdate struct {
	ThreadUpdate *Thread `protobuf:"bytes,4,opt,name=thread_update,json=threadUpdate,proto3,oneof"`
}

type MailerEvent_ThreadDelete struct {
	ThreadDelete uint64 `protobuf:"varint,5,opt,name=thread_delete,json=threadDelete,proto3,oneof"`
}

type MailerEvent_ThreadStateUpdate struct {
	ThreadStateUpdate *ThreadState `protobuf:"bytes,6,opt,name=thread_state_update,json=threadStateUpdate,proto3,oneof"`
}

type MailerEvent_MessageUpdate struct {
	MessageUpdate *Message `protobuf:"bytes,7,opt,name=message_update,json=messageUpdate,proto3,oneof"`
}

type MailerEvent_MessageDelete struct {
	MessageDelete uint64 `protobuf:"varint,8,opt,name=message_delete,json=messageDelete,proto3,oneof"`
}

func (*MailerEvent_EmailUpdate) isMailerEvent_Data() {}

func (*MailerEvent_EmailDelete) isMailerEvent_Data() {}

func (*MailerEvent_EmailSettingsUpdated) isMailerEvent_Data() {}

func (*MailerEvent_ThreadUpdate) isMailerEvent_Data() {}

func (*MailerEvent_ThreadDelete) isMailerEvent_Data() {}

func (*MailerEvent_ThreadStateUpdate) isMailerEvent_Data() {}

func (*MailerEvent_MessageUpdate) isMailerEvent_Data() {}

func (*MailerEvent_MessageDelete) isMailerEvent_Data() {}

var File_resources_mailer_events_proto protoreflect.FileDescriptor

var file_resources_mailer_events_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x1a, 0x1c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x04, 0x0a, 0x0b, 0x4d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x48, 0x00, 0x52, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x57, 0x0a, 0x16, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x48, 0x00, 0x52, 0x14, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x0d, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x13, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x11, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x0e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x3b, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resources_mailer_events_proto_rawDescOnce sync.Once
	file_resources_mailer_events_proto_rawDescData = file_resources_mailer_events_proto_rawDesc
)

func file_resources_mailer_events_proto_rawDescGZIP() []byte {
	file_resources_mailer_events_proto_rawDescOnce.Do(func() {
		file_resources_mailer_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_mailer_events_proto_rawDescData)
	})
	return file_resources_mailer_events_proto_rawDescData
}

var file_resources_mailer_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_mailer_events_proto_goTypes = []any{
	(*MailerEvent)(nil),   // 0: resources.mailer.MailerEvent
	(*Email)(nil),         // 1: resources.mailer.Email
	(*EmailSettings)(nil), // 2: resources.mailer.EmailSettings
	(*Thread)(nil),        // 3: resources.mailer.Thread
	(*ThreadState)(nil),   // 4: resources.mailer.ThreadState
	(*Message)(nil),       // 5: resources.mailer.Message
}
var file_resources_mailer_events_proto_depIdxs = []int32{
	1, // 0: resources.mailer.MailerEvent.email_update:type_name -> resources.mailer.Email
	2, // 1: resources.mailer.MailerEvent.email_settings_updated:type_name -> resources.mailer.EmailSettings
	3, // 2: resources.mailer.MailerEvent.thread_update:type_name -> resources.mailer.Thread
	4, // 3: resources.mailer.MailerEvent.thread_state_update:type_name -> resources.mailer.ThreadState
	5, // 4: resources.mailer.MailerEvent.message_update:type_name -> resources.mailer.Message
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resources_mailer_events_proto_init() }
func file_resources_mailer_events_proto_init() {
	if File_resources_mailer_events_proto != nil {
		return
	}
	file_resources_mailer_email_proto_init()
	file_resources_mailer_settings_proto_init()
	file_resources_mailer_message_proto_init()
	file_resources_mailer_thread_proto_init()
	file_resources_mailer_events_proto_msgTypes[0].OneofWrappers = []any{
		(*MailerEvent_EmailUpdate)(nil),
		(*MailerEvent_EmailDelete)(nil),
		(*MailerEvent_EmailSettingsUpdated)(nil),
		(*MailerEvent_ThreadUpdate)(nil),
		(*MailerEvent_ThreadDelete)(nil),
		(*MailerEvent_ThreadStateUpdate)(nil),
		(*MailerEvent_MessageUpdate)(nil),
		(*MailerEvent_MessageDelete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_mailer_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_mailer_events_proto_goTypes,
		DependencyIndexes: file_resources_mailer_events_proto_depIdxs,
		MessageInfos:      file_resources_mailer_events_proto_msgTypes,
	}.Build()
	File_resources_mailer_events_proto = out.File
	file_resources_mailer_events_proto_rawDesc = nil
	file_resources_mailer_events_proto_goTypes = nil
	file_resources_mailer_events_proto_depIdxs = nil
}
