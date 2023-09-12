// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: services/livemapper/livemap.proto

package livemapper

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	livemap "github.com/galexrt/fivenet/gen/go/proto/resources/livemap"
	users "github.com/galexrt/fivenet/gen/go/proto/resources/users"
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

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_livemapper_livemap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_livemapper_livemap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_services_livemapper_livemap_proto_rawDescGZIP(), []int{0}
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobsUsers   []*users.Job          `protobuf:"bytes,1,rep,name=jobs_users,json=jobsUsers,proto3" json:"jobs_users,omitempty"`
	Users       []*livemap.UserMarker `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	JobsMarkers []*users.Job          `protobuf:"bytes,3,rep,name=jobs_markers,json=jobsMarkers,proto3" json:"jobs_markers,omitempty"`
	Markers     []*livemap.Marker     `protobuf:"bytes,4,rep,name=markers,proto3" json:"markers,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_livemapper_livemap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_livemapper_livemap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_services_livemapper_livemap_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResponse) GetJobsUsers() []*users.Job {
	if x != nil {
		return x.JobsUsers
	}
	return nil
}

func (x *StreamResponse) GetUsers() []*livemap.UserMarker {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *StreamResponse) GetJobsMarkers() []*users.Job {
	if x != nil {
		return x.JobsMarkers
	}
	return nil
}

func (x *StreamResponse) GetMarkers() []*livemap.Marker {
	if x != nil {
		return x.Markers
	}
	return nil
}

type CreateOrUpdateMarkerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Marker *livemap.Marker `protobuf:"bytes,1,opt,name=marker,proto3" json:"marker,omitempty"`
}

func (x *CreateOrUpdateMarkerRequest) Reset() {
	*x = CreateOrUpdateMarkerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_livemapper_livemap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateMarkerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateMarkerRequest) ProtoMessage() {}

func (x *CreateOrUpdateMarkerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_livemapper_livemap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateMarkerRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateMarkerRequest) Descriptor() ([]byte, []int) {
	return file_services_livemapper_livemap_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrUpdateMarkerRequest) GetMarker() *livemap.Marker {
	if x != nil {
		return x.Marker
	}
	return nil
}

type CreateOrUpdateMarkerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Marker *livemap.Marker `protobuf:"bytes,1,opt,name=marker,proto3" json:"marker,omitempty"`
}

func (x *CreateOrUpdateMarkerResponse) Reset() {
	*x = CreateOrUpdateMarkerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_livemapper_livemap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateMarkerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateMarkerResponse) ProtoMessage() {}

func (x *CreateOrUpdateMarkerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_livemapper_livemap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateMarkerResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateMarkerResponse) Descriptor() ([]byte, []int) {
	return file_services_livemapper_livemap_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrUpdateMarkerResponse) GetMarker() *livemap.Marker {
	if x != nil {
		return x.Marker
	}
	return nil
}

type DeleteMarkerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMarkerRequest) Reset() {
	*x = DeleteMarkerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_livemapper_livemap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMarkerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMarkerRequest) ProtoMessage() {}

func (x *DeleteMarkerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_livemapper_livemap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMarkerRequest.ProtoReflect.Descriptor instead.
func (*DeleteMarkerRequest) Descriptor() ([]byte, []int) {
	return file_services_livemapper_livemap_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMarkerRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteMarkerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMarkerResponse) Reset() {
	*x = DeleteMarkerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_livemapper_livemap_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMarkerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMarkerResponse) ProtoMessage() {}

func (x *DeleteMarkerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_livemapper_livemap_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMarkerResponse.ProtoReflect.Descriptor instead.
func (*DeleteMarkerResponse) Descriptor() ([]byte, []int) {
	return file_services_livemapper_livemap_proto_rawDescGZIP(), []int{5}
}

var File_services_livemapper_livemap_proto protoreflect.FileDescriptor

var file_services_livemapper_livemap_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6d,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x1a, 0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2f, 0x6c, 0x69, 0x76, 0x65,
	0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f,
	0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xe8, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x09, 0x6a, 0x6f,
	0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0c,
	0x6a, 0x6f, 0x62, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x73, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xca, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x76,
	0x65, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x7b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x63, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72,
	0x12, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x66, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x3b, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_livemapper_livemap_proto_rawDescOnce sync.Once
	file_services_livemapper_livemap_proto_rawDescData = file_services_livemapper_livemap_proto_rawDesc
)

func file_services_livemapper_livemap_proto_rawDescGZIP() []byte {
	file_services_livemapper_livemap_proto_rawDescOnce.Do(func() {
		file_services_livemapper_livemap_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_livemapper_livemap_proto_rawDescData)
	})
	return file_services_livemapper_livemap_proto_rawDescData
}

var file_services_livemapper_livemap_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_livemapper_livemap_proto_goTypes = []interface{}{
	(*StreamRequest)(nil),                // 0: services.livemapper.StreamRequest
	(*StreamResponse)(nil),               // 1: services.livemapper.StreamResponse
	(*CreateOrUpdateMarkerRequest)(nil),  // 2: services.livemapper.CreateOrUpdateMarkerRequest
	(*CreateOrUpdateMarkerResponse)(nil), // 3: services.livemapper.CreateOrUpdateMarkerResponse
	(*DeleteMarkerRequest)(nil),          // 4: services.livemapper.DeleteMarkerRequest
	(*DeleteMarkerResponse)(nil),         // 5: services.livemapper.DeleteMarkerResponse
	(*users.Job)(nil),                    // 6: resources.users.Job
	(*livemap.UserMarker)(nil),           // 7: resources.livemap.UserMarker
	(*livemap.Marker)(nil),               // 8: resources.livemap.Marker
}
var file_services_livemapper_livemap_proto_depIdxs = []int32{
	6, // 0: services.livemapper.StreamResponse.jobs_users:type_name -> resources.users.Job
	7, // 1: services.livemapper.StreamResponse.users:type_name -> resources.livemap.UserMarker
	6, // 2: services.livemapper.StreamResponse.jobs_markers:type_name -> resources.users.Job
	8, // 3: services.livemapper.StreamResponse.markers:type_name -> resources.livemap.Marker
	8, // 4: services.livemapper.CreateOrUpdateMarkerRequest.marker:type_name -> resources.livemap.Marker
	8, // 5: services.livemapper.CreateOrUpdateMarkerResponse.marker:type_name -> resources.livemap.Marker
	0, // 6: services.livemapper.LivemapperService.Stream:input_type -> services.livemapper.StreamRequest
	2, // 7: services.livemapper.LivemapperService.CreateOrUpdateMarker:input_type -> services.livemapper.CreateOrUpdateMarkerRequest
	4, // 8: services.livemapper.LivemapperService.DeleteMarker:input_type -> services.livemapper.DeleteMarkerRequest
	1, // 9: services.livemapper.LivemapperService.Stream:output_type -> services.livemapper.StreamResponse
	3, // 10: services.livemapper.LivemapperService.CreateOrUpdateMarker:output_type -> services.livemapper.CreateOrUpdateMarkerResponse
	5, // 11: services.livemapper.LivemapperService.DeleteMarker:output_type -> services.livemapper.DeleteMarkerResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_services_livemapper_livemap_proto_init() }
func file_services_livemapper_livemap_proto_init() {
	if File_services_livemapper_livemap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_livemapper_livemap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_services_livemapper_livemap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
		file_services_livemapper_livemap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateMarkerRequest); i {
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
		file_services_livemapper_livemap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateMarkerResponse); i {
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
		file_services_livemapper_livemap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMarkerRequest); i {
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
		file_services_livemapper_livemap_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMarkerResponse); i {
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
			RawDescriptor: file_services_livemapper_livemap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_livemapper_livemap_proto_goTypes,
		DependencyIndexes: file_services_livemapper_livemap_proto_depIdxs,
		MessageInfos:      file_services_livemapper_livemap_proto_msgTypes,
	}.Build()
	File_services_livemapper_livemap_proto = out.File
	file_services_livemapper_livemap_proto_rawDesc = nil
	file_services_livemapper_livemap_proto_goTypes = nil
	file_services_livemapper_livemap_proto_depIdxs = nil
}
