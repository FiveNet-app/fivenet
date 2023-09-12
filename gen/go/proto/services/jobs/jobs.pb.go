// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: services/jobs/jobs.proto

package jobs

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	database "github.com/galexrt/fivenet/gen/go/proto/resources/common/database"
	jobs "github.com/galexrt/fivenet/gen/go/proto/resources/jobs"
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

type ColleaguesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Search params
	SearchName string `protobuf:"bytes,2,opt,name=search_name,json=searchName,proto3" json:"search_name,omitempty"`
}

func (x *ColleaguesListRequest) Reset() {
	*x = ColleaguesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColleaguesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleaguesListRequest) ProtoMessage() {}

func (x *ColleaguesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleaguesListRequest.ProtoReflect.Descriptor instead.
func (*ColleaguesListRequest) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{0}
}

func (x *ColleaguesListRequest) GetPagination() *database.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ColleaguesListRequest) GetSearchName() string {
	if x != nil {
		return x.SearchName
	}
	return ""
}

type ColleaguesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Users      []*users.User                `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ColleaguesListResponse) Reset() {
	*x = ColleaguesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColleaguesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleaguesListResponse) ProtoMessage() {}

func (x *ColleaguesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleaguesListResponse.ProtoReflect.Descriptor instead.
func (*ColleaguesListResponse) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{1}
}

func (x *ColleaguesListResponse) GetPagination() *database.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ColleaguesListResponse) GetUsers() []*users.User {
	if x != nil {
		return x.Users
	}
	return nil
}

type ConductListEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Search params
	Types       []jobs.CONDUCT_TYPE `protobuf:"varint,2,rep,packed,name=types,proto3,enum=resources.jobs.CONDUCT_TYPE" json:"types,omitempty"`
	ShowExpired *bool               `protobuf:"varint,3,opt,name=show_expired,json=showExpired,proto3,oneof" json:"show_expired,omitempty"`
	UserIds     []int32             `protobuf:"varint,4,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *ConductListEntriesRequest) Reset() {
	*x = ConductListEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductListEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductListEntriesRequest) ProtoMessage() {}

func (x *ConductListEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductListEntriesRequest.ProtoReflect.Descriptor instead.
func (*ConductListEntriesRequest) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{2}
}

func (x *ConductListEntriesRequest) GetPagination() *database.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ConductListEntriesRequest) GetTypes() []jobs.CONDUCT_TYPE {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *ConductListEntriesRequest) GetShowExpired() bool {
	if x != nil && x.ShowExpired != nil {
		return *x.ShowExpired
	}
	return false
}

func (x *ConductListEntriesRequest) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type ConductListEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Entries    []*jobs.ConductEntry         `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ConductListEntriesResponse) Reset() {
	*x = ConductListEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductListEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductListEntriesResponse) ProtoMessage() {}

func (x *ConductListEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductListEntriesResponse.ProtoReflect.Descriptor instead.
func (*ConductListEntriesResponse) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{3}
}

func (x *ConductListEntriesResponse) GetPagination() *database.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ConductListEntriesResponse) GetEntries() []*jobs.ConductEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type ConductCreateEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *jobs.ConductEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *ConductCreateEntryRequest) Reset() {
	*x = ConductCreateEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductCreateEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductCreateEntryRequest) ProtoMessage() {}

func (x *ConductCreateEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductCreateEntryRequest.ProtoReflect.Descriptor instead.
func (*ConductCreateEntryRequest) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{4}
}

func (x *ConductCreateEntryRequest) GetEntry() *jobs.ConductEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type ConductCreateEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *jobs.ConductEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *ConductCreateEntryResponse) Reset() {
	*x = ConductCreateEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductCreateEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductCreateEntryResponse) ProtoMessage() {}

func (x *ConductCreateEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductCreateEntryResponse.ProtoReflect.Descriptor instead.
func (*ConductCreateEntryResponse) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{5}
}

func (x *ConductCreateEntryResponse) GetEntry() *jobs.ConductEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type ConductUpdateEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *jobs.ConductEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *ConductUpdateEntryRequest) Reset() {
	*x = ConductUpdateEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductUpdateEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductUpdateEntryRequest) ProtoMessage() {}

func (x *ConductUpdateEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductUpdateEntryRequest.ProtoReflect.Descriptor instead.
func (*ConductUpdateEntryRequest) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{6}
}

func (x *ConductUpdateEntryRequest) GetEntry() *jobs.ConductEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type ConductUpdateEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *jobs.ConductEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *ConductUpdateEntryResponse) Reset() {
	*x = ConductUpdateEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductUpdateEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductUpdateEntryResponse) ProtoMessage() {}

func (x *ConductUpdateEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductUpdateEntryResponse.ProtoReflect.Descriptor instead.
func (*ConductUpdateEntryResponse) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{7}
}

func (x *ConductUpdateEntryResponse) GetEntry() *jobs.ConductEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type ConductDeleteEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConductDeleteEntryRequest) Reset() {
	*x = ConductDeleteEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductDeleteEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductDeleteEntryRequest) ProtoMessage() {}

func (x *ConductDeleteEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductDeleteEntryRequest.ProtoReflect.Descriptor instead.
func (*ConductDeleteEntryRequest) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{8}
}

func (x *ConductDeleteEntryRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConductDeleteEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConductDeleteEntryResponse) Reset() {
	*x = ConductDeleteEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_jobs_jobs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConductDeleteEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConductDeleteEntryResponse) ProtoMessage() {}

func (x *ConductDeleteEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_jobs_jobs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConductDeleteEntryResponse.ProtoReflect.Descriptor instead.
func (*ConductDeleteEntryResponse) Descriptor() ([]byte, []int) {
	return file_services_jobs_jobs_proto_rawDescGZIP(), []int{9}
}

var File_services_jobs_jobs_proto protoreflect.FileDescriptor

var file_services_jobs_jobs_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x1a, 0x28, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x56, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xfb, 0x01, 0x0a, 0x19, 0x43,
	0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x32, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73,
	0x2e, 0x43, 0x4f, 0x4e, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x68,
	0x6f, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x68, 0x6f, 0x77,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6e,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x59,
	0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x75, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x1a, 0x43, 0x6f, 0x6e,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x59, 0x0a, 0x19, 0x43,
	0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x5a, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x22, 0x2b, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1c, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x98, 0x04,
	0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a,
	0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x12,
	0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f,
	0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x69, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f,
	0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a,
	0x12, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x66,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62,
	0x73, 0x3b, 0x6a, 0x6f, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_jobs_jobs_proto_rawDescOnce sync.Once
	file_services_jobs_jobs_proto_rawDescData = file_services_jobs_jobs_proto_rawDesc
)

func file_services_jobs_jobs_proto_rawDescGZIP() []byte {
	file_services_jobs_jobs_proto_rawDescOnce.Do(func() {
		file_services_jobs_jobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_jobs_jobs_proto_rawDescData)
	})
	return file_services_jobs_jobs_proto_rawDescData
}

var file_services_jobs_jobs_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_jobs_jobs_proto_goTypes = []interface{}{
	(*ColleaguesListRequest)(nil),       // 0: services.jobs.ColleaguesListRequest
	(*ColleaguesListResponse)(nil),      // 1: services.jobs.ColleaguesListResponse
	(*ConductListEntriesRequest)(nil),   // 2: services.jobs.ConductListEntriesRequest
	(*ConductListEntriesResponse)(nil),  // 3: services.jobs.ConductListEntriesResponse
	(*ConductCreateEntryRequest)(nil),   // 4: services.jobs.ConductCreateEntryRequest
	(*ConductCreateEntryResponse)(nil),  // 5: services.jobs.ConductCreateEntryResponse
	(*ConductUpdateEntryRequest)(nil),   // 6: services.jobs.ConductUpdateEntryRequest
	(*ConductUpdateEntryResponse)(nil),  // 7: services.jobs.ConductUpdateEntryResponse
	(*ConductDeleteEntryRequest)(nil),   // 8: services.jobs.ConductDeleteEntryRequest
	(*ConductDeleteEntryResponse)(nil),  // 9: services.jobs.ConductDeleteEntryResponse
	(*database.PaginationRequest)(nil),  // 10: resources.common.database.PaginationRequest
	(*database.PaginationResponse)(nil), // 11: resources.common.database.PaginationResponse
	(*users.User)(nil),                  // 12: resources.users.User
	(jobs.CONDUCT_TYPE)(0),              // 13: resources.jobs.CONDUCT_TYPE
	(*jobs.ConductEntry)(nil),           // 14: resources.jobs.ConductEntry
}
var file_services_jobs_jobs_proto_depIdxs = []int32{
	10, // 0: services.jobs.ColleaguesListRequest.pagination:type_name -> resources.common.database.PaginationRequest
	11, // 1: services.jobs.ColleaguesListResponse.pagination:type_name -> resources.common.database.PaginationResponse
	12, // 2: services.jobs.ColleaguesListResponse.users:type_name -> resources.users.User
	10, // 3: services.jobs.ConductListEntriesRequest.pagination:type_name -> resources.common.database.PaginationRequest
	13, // 4: services.jobs.ConductListEntriesRequest.types:type_name -> resources.jobs.CONDUCT_TYPE
	11, // 5: services.jobs.ConductListEntriesResponse.pagination:type_name -> resources.common.database.PaginationResponse
	14, // 6: services.jobs.ConductListEntriesResponse.entries:type_name -> resources.jobs.ConductEntry
	14, // 7: services.jobs.ConductCreateEntryRequest.entry:type_name -> resources.jobs.ConductEntry
	14, // 8: services.jobs.ConductCreateEntryResponse.entry:type_name -> resources.jobs.ConductEntry
	14, // 9: services.jobs.ConductUpdateEntryRequest.entry:type_name -> resources.jobs.ConductEntry
	14, // 10: services.jobs.ConductUpdateEntryResponse.entry:type_name -> resources.jobs.ConductEntry
	0,  // 11: services.jobs.JobsService.ColleaguesList:input_type -> services.jobs.ColleaguesListRequest
	2,  // 12: services.jobs.JobsService.ConductListEntries:input_type -> services.jobs.ConductListEntriesRequest
	4,  // 13: services.jobs.JobsService.ConductCreateEntry:input_type -> services.jobs.ConductCreateEntryRequest
	6,  // 14: services.jobs.JobsService.ConductUpdateEntry:input_type -> services.jobs.ConductUpdateEntryRequest
	8,  // 15: services.jobs.JobsService.ConductDeleteEntry:input_type -> services.jobs.ConductDeleteEntryRequest
	1,  // 16: services.jobs.JobsService.ColleaguesList:output_type -> services.jobs.ColleaguesListResponse
	3,  // 17: services.jobs.JobsService.ConductListEntries:output_type -> services.jobs.ConductListEntriesResponse
	5,  // 18: services.jobs.JobsService.ConductCreateEntry:output_type -> services.jobs.ConductCreateEntryResponse
	7,  // 19: services.jobs.JobsService.ConductUpdateEntry:output_type -> services.jobs.ConductUpdateEntryResponse
	9,  // 20: services.jobs.JobsService.ConductDeleteEntry:output_type -> services.jobs.ConductDeleteEntryResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_services_jobs_jobs_proto_init() }
func file_services_jobs_jobs_proto_init() {
	if File_services_jobs_jobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_jobs_jobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColleaguesListRequest); i {
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
		file_services_jobs_jobs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColleaguesListResponse); i {
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
		file_services_jobs_jobs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductListEntriesRequest); i {
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
		file_services_jobs_jobs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductListEntriesResponse); i {
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
		file_services_jobs_jobs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductCreateEntryRequest); i {
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
		file_services_jobs_jobs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductCreateEntryResponse); i {
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
		file_services_jobs_jobs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductUpdateEntryRequest); i {
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
		file_services_jobs_jobs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductUpdateEntryResponse); i {
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
		file_services_jobs_jobs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductDeleteEntryRequest); i {
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
		file_services_jobs_jobs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConductDeleteEntryResponse); i {
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
	file_services_jobs_jobs_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_jobs_jobs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_jobs_jobs_proto_goTypes,
		DependencyIndexes: file_services_jobs_jobs_proto_depIdxs,
		MessageInfos:      file_services_jobs_jobs_proto_msgTypes,
	}.Build()
	File_services_jobs_jobs_proto = out.File
	file_services_jobs_jobs_proto_rawDesc = nil
	file_services_jobs_jobs_proto_goTypes = nil
	file_services_jobs_jobs_proto_depIdxs = nil
}
