// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: services/centrum/centrum.proto

package centrum

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

type CreateSquadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSquadRequest) Reset() {
	*x = CreateSquadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSquadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSquadRequest) ProtoMessage() {}

func (x *CreateSquadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSquadRequest.ProtoReflect.Descriptor instead.
func (*CreateSquadRequest) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{0}
}

type CreateSquadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSquadResponse) Reset() {
	*x = CreateSquadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSquadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSquadResponse) ProtoMessage() {}

func (x *CreateSquadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSquadResponse.ProtoReflect.Descriptor instead.
func (*CreateSquadResponse) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{1}
}

type UpdateSquadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSquadRequest) Reset() {
	*x = UpdateSquadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSquadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSquadRequest) ProtoMessage() {}

func (x *UpdateSquadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSquadRequest.ProtoReflect.Descriptor instead.
func (*UpdateSquadRequest) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{2}
}

type UpdateSquadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSquadResponse) Reset() {
	*x = UpdateSquadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSquadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSquadResponse) ProtoMessage() {}

func (x *UpdateSquadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSquadResponse.ProtoReflect.Descriptor instead.
func (*UpdateSquadResponse) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{3}
}

type DeleteSquadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSquadRequest) Reset() {
	*x = DeleteSquadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSquadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSquadRequest) ProtoMessage() {}

func (x *DeleteSquadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSquadRequest.ProtoReflect.Descriptor instead.
func (*DeleteSquadRequest) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{4}
}

type DeleteSquadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSquadResponse) Reset() {
	*x = DeleteSquadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSquadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSquadResponse) ProtoMessage() {}

func (x *DeleteSquadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSquadResponse.ProtoReflect.Descriptor instead.
func (*DeleteSquadResponse) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{5}
}

type AssignSquadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssignSquadRequest) Reset() {
	*x = AssignSquadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignSquadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignSquadRequest) ProtoMessage() {}

func (x *AssignSquadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignSquadRequest.ProtoReflect.Descriptor instead.
func (*AssignSquadRequest) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{6}
}

type AssignSquadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssignSquadResponse) Reset() {
	*x = AssignSquadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_centrum_centrum_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignSquadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignSquadResponse) ProtoMessage() {}

func (x *AssignSquadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_centrum_centrum_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignSquadResponse.ProtoReflect.Descriptor instead.
func (*AssignSquadResponse) Descriptor() ([]byte, []int) {
	return file_services_centrum_centrum_proto_rawDescGZIP(), []int{7}
}

var File_services_centrum_centrum_proto protoreflect.FileDescriptor

var file_services_centrum_centrum_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x75, 0x6d, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x75, 0x6d, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x15, 0x0a, 0x13, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x80, 0x03, 0x0a, 0x0e, 0x43, 0x65, 0x6e, 0x74, 0x72,
	0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75,
	0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x71, 0x75, 0x61, 0x64, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x71,
	0x75, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64,
	0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x75, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x0b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x71, 0x75, 0x61, 0x64, 0x12, 0x24, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x71, 0x75, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x71, 0x75, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f,
	0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x75, 0x6d, 0x3b, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_centrum_centrum_proto_rawDescOnce sync.Once
	file_services_centrum_centrum_proto_rawDescData = file_services_centrum_centrum_proto_rawDesc
)

func file_services_centrum_centrum_proto_rawDescGZIP() []byte {
	file_services_centrum_centrum_proto_rawDescOnce.Do(func() {
		file_services_centrum_centrum_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_centrum_centrum_proto_rawDescData)
	})
	return file_services_centrum_centrum_proto_rawDescData
}

var file_services_centrum_centrum_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_centrum_centrum_proto_goTypes = []interface{}{
	(*CreateSquadRequest)(nil),  // 0: services.centrum.CreateSquadRequest
	(*CreateSquadResponse)(nil), // 1: services.centrum.CreateSquadResponse
	(*UpdateSquadRequest)(nil),  // 2: services.centrum.UpdateSquadRequest
	(*UpdateSquadResponse)(nil), // 3: services.centrum.UpdateSquadResponse
	(*DeleteSquadRequest)(nil),  // 4: services.centrum.DeleteSquadRequest
	(*DeleteSquadResponse)(nil), // 5: services.centrum.DeleteSquadResponse
	(*AssignSquadRequest)(nil),  // 6: services.centrum.AssignSquadRequest
	(*AssignSquadResponse)(nil), // 7: services.centrum.AssignSquadResponse
}
var file_services_centrum_centrum_proto_depIdxs = []int32{
	0, // 0: services.centrum.CentrumService.CreateSquad:input_type -> services.centrum.CreateSquadRequest
	2, // 1: services.centrum.CentrumService.UpdateSquad:input_type -> services.centrum.UpdateSquadRequest
	4, // 2: services.centrum.CentrumService.DeleteSquad:input_type -> services.centrum.DeleteSquadRequest
	6, // 3: services.centrum.CentrumService.AssignSquad:input_type -> services.centrum.AssignSquadRequest
	1, // 4: services.centrum.CentrumService.CreateSquad:output_type -> services.centrum.CreateSquadResponse
	3, // 5: services.centrum.CentrumService.UpdateSquad:output_type -> services.centrum.UpdateSquadResponse
	5, // 6: services.centrum.CentrumService.DeleteSquad:output_type -> services.centrum.DeleteSquadResponse
	7, // 7: services.centrum.CentrumService.AssignSquad:output_type -> services.centrum.AssignSquadResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_centrum_centrum_proto_init() }
func file_services_centrum_centrum_proto_init() {
	if File_services_centrum_centrum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_centrum_centrum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSquadRequest); i {
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
		file_services_centrum_centrum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSquadResponse); i {
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
		file_services_centrum_centrum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSquadRequest); i {
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
		file_services_centrum_centrum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSquadResponse); i {
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
		file_services_centrum_centrum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSquadRequest); i {
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
		file_services_centrum_centrum_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSquadResponse); i {
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
		file_services_centrum_centrum_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignSquadRequest); i {
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
		file_services_centrum_centrum_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignSquadResponse); i {
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
			RawDescriptor: file_services_centrum_centrum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_centrum_centrum_proto_goTypes,
		DependencyIndexes: file_services_centrum_centrum_proto_depIdxs,
		MessageInfos:      file_services_centrum_centrum_proto_msgTypes,
	}.Build()
	File_services_centrum_centrum_proto = out.File
	file_services_centrum_centrum_proto_rawDesc = nil
	file_services_centrum_centrum_proto_goTypes = nil
	file_services_centrum_centrum_proto_depIdxs = nil
}
