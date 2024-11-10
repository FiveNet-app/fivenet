// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.3
// source: services/wiki/wiki.proto

package wiki

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	database "github.com/fivenet-app/fivenet/gen/go/proto/resources/common/database"
	wiki "github.com/fivenet-app/fivenet/gen/go/proto/resources/wiki"
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

type ListPagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Sort       *database.Sort              `protobuf:"bytes,2,opt,name=sort,proto3,oneof" json:"sort,omitempty"`
	// Search params
	Job      *string `protobuf:"bytes,3,opt,name=job,proto3,oneof" json:"job,omitempty"`
	RootOnly *bool   `protobuf:"varint,4,opt,name=root_only,json=rootOnly,proto3,oneof" json:"root_only,omitempty"`
	Search   *string `protobuf:"bytes,5,opt,name=search,proto3,oneof" json:"search,omitempty"`
}

func (x *ListPagesRequest) Reset() {
	*x = ListPagesRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPagesRequest) ProtoMessage() {}

func (x *ListPagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPagesRequest.ProtoReflect.Descriptor instead.
func (*ListPagesRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{0}
}

func (x *ListPagesRequest) GetPagination() *database.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPagesRequest) GetSort() *database.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListPagesRequest) GetJob() string {
	if x != nil && x.Job != nil {
		return *x.Job
	}
	return ""
}

func (x *ListPagesRequest) GetRootOnly() bool {
	if x != nil && x.RootOnly != nil {
		return *x.RootOnly
	}
	return false
}

func (x *ListPagesRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

type ListPagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Pages      []*wiki.PageShort            `protobuf:"bytes,2,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *ListPagesResponse) Reset() {
	*x = ListPagesResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPagesResponse) ProtoMessage() {}

func (x *ListPagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPagesResponse.ProtoReflect.Descriptor instead.
func (*ListPagesResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{1}
}

func (x *ListPagesResponse) GetPagination() *database.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPagesResponse) GetPages() []*wiki.PageShort {
	if x != nil {
		return x.Pages
	}
	return nil
}

type GetPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPageRequest) Reset() {
	*x = GetPageRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageRequest) ProtoMessage() {}

func (x *GetPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageRequest.ProtoReflect.Descriptor instead.
func (*GetPageRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{2}
}

func (x *GetPageRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *wiki.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetPageResponse) Reset() {
	*x = GetPageResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageResponse) ProtoMessage() {}

func (x *GetPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageResponse.ProtoReflect.Descriptor instead.
func (*GetPageResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{3}
}

func (x *GetPageResponse) GetPage() *wiki.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type CreatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *wiki.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *CreatePageRequest) Reset() {
	*x = CreatePageRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePageRequest) ProtoMessage() {}

func (x *CreatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePageRequest.ProtoReflect.Descriptor instead.
func (*CreatePageRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePageRequest) GetPage() *wiki.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type CreatePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *wiki.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *CreatePageResponse) Reset() {
	*x = CreatePageResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePageResponse) ProtoMessage() {}

func (x *CreatePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePageResponse.ProtoReflect.Descriptor instead.
func (*CreatePageResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePageResponse) GetPage() *wiki.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type UpdatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *wiki.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *UpdatePageRequest) Reset() {
	*x = UpdatePageRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageRequest) ProtoMessage() {}

func (x *UpdatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageRequest.ProtoReflect.Descriptor instead.
func (*UpdatePageRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePageRequest) GetPage() *wiki.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type UpdatePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *wiki.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *UpdatePageResponse) Reset() {
	*x = UpdatePageResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageResponse) ProtoMessage() {}

func (x *UpdatePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageResponse.ProtoReflect.Descriptor instead.
func (*UpdatePageResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePageResponse) GetPage() *wiki.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type DeletePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePageRequest) Reset() {
	*x = DeletePageRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageRequest) ProtoMessage() {}

func (x *DeletePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageRequest.ProtoReflect.Descriptor instead.
func (*DeletePageRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePageRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePageResponse) Reset() {
	*x = DeletePageResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageResponse) ProtoMessage() {}

func (x *DeletePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageResponse.ProtoReflect.Descriptor instead.
func (*DeletePageResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{9}
}

type ListPageActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	PageId     uint64                      `protobuf:"varint,2,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
}

func (x *ListPageActivityRequest) Reset() {
	*x = ListPageActivityRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPageActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPageActivityRequest) ProtoMessage() {}

func (x *ListPageActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPageActivityRequest.ProtoReflect.Descriptor instead.
func (*ListPageActivityRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{10}
}

func (x *ListPageActivityRequest) GetPagination() *database.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPageActivityRequest) GetPageId() uint64 {
	if x != nil {
		return x.PageId
	}
	return 0
}

type ListPageActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Activity   []*wiki.PageActivity         `protobuf:"bytes,2,rep,name=activity,proto3" json:"activity,omitempty"`
}

func (x *ListPageActivityResponse) Reset() {
	*x = ListPageActivityResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPageActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPageActivityResponse) ProtoMessage() {}

func (x *ListPageActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_wiki_wiki_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPageActivityResponse.ProtoReflect.Descriptor instead.
func (*ListPageActivityResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{11}
}

func (x *ListPageActivityResponse) GetPagination() *database.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPageActivityResponse) GetActivity() []*wiki.PageActivity {
	if x != nil {
		return x.Activity
	}
	return nil
}

var File_services_wiki_wiki_proto protoreflect.FileDescriptor

var file_services_wiki_wiki_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f,
	0x77, 0x69, 0x6b, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x1a, 0x28, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x77,
	0x69, 0x6b, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69,
	0x6b, 0x69, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x48, 0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x18, 0x32, 0x48, 0x01, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x02, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x03, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x6a, 0x6f, 0x62, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22,
	0x9d, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x56, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x57, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x32, 0x85, 0x04, 0x0a, 0x0b, 0x57, 0x69, 0x6b, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69,
	0x6b, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x51, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b,
	0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69,
	0x6b, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x77, 0x69, 0x6b, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x26, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77,
	0x69, 0x6b, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x3b, 0x77, 0x69, 0x6b, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_wiki_wiki_proto_rawDescOnce sync.Once
	file_services_wiki_wiki_proto_rawDescData = file_services_wiki_wiki_proto_rawDesc
)

func file_services_wiki_wiki_proto_rawDescGZIP() []byte {
	file_services_wiki_wiki_proto_rawDescOnce.Do(func() {
		file_services_wiki_wiki_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_wiki_wiki_proto_rawDescData)
	})
	return file_services_wiki_wiki_proto_rawDescData
}

var file_services_wiki_wiki_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_services_wiki_wiki_proto_goTypes = []any{
	(*ListPagesRequest)(nil),            // 0: services.wiki.ListPagesRequest
	(*ListPagesResponse)(nil),           // 1: services.wiki.ListPagesResponse
	(*GetPageRequest)(nil),              // 2: services.wiki.GetPageRequest
	(*GetPageResponse)(nil),             // 3: services.wiki.GetPageResponse
	(*CreatePageRequest)(nil),           // 4: services.wiki.CreatePageRequest
	(*CreatePageResponse)(nil),          // 5: services.wiki.CreatePageResponse
	(*UpdatePageRequest)(nil),           // 6: services.wiki.UpdatePageRequest
	(*UpdatePageResponse)(nil),          // 7: services.wiki.UpdatePageResponse
	(*DeletePageRequest)(nil),           // 8: services.wiki.DeletePageRequest
	(*DeletePageResponse)(nil),          // 9: services.wiki.DeletePageResponse
	(*ListPageActivityRequest)(nil),     // 10: services.wiki.ListPageActivityRequest
	(*ListPageActivityResponse)(nil),    // 11: services.wiki.ListPageActivityResponse
	(*database.PaginationRequest)(nil),  // 12: resources.common.database.PaginationRequest
	(*database.Sort)(nil),               // 13: resources.common.database.Sort
	(*database.PaginationResponse)(nil), // 14: resources.common.database.PaginationResponse
	(*wiki.PageShort)(nil),              // 15: resources.wiki.PageShort
	(*wiki.Page)(nil),                   // 16: resources.wiki.Page
	(*wiki.PageActivity)(nil),           // 17: resources.wiki.PageActivity
}
var file_services_wiki_wiki_proto_depIdxs = []int32{
	12, // 0: services.wiki.ListPagesRequest.pagination:type_name -> resources.common.database.PaginationRequest
	13, // 1: services.wiki.ListPagesRequest.sort:type_name -> resources.common.database.Sort
	14, // 2: services.wiki.ListPagesResponse.pagination:type_name -> resources.common.database.PaginationResponse
	15, // 3: services.wiki.ListPagesResponse.pages:type_name -> resources.wiki.PageShort
	16, // 4: services.wiki.GetPageResponse.page:type_name -> resources.wiki.Page
	16, // 5: services.wiki.CreatePageRequest.page:type_name -> resources.wiki.Page
	16, // 6: services.wiki.CreatePageResponse.page:type_name -> resources.wiki.Page
	16, // 7: services.wiki.UpdatePageRequest.page:type_name -> resources.wiki.Page
	16, // 8: services.wiki.UpdatePageResponse.page:type_name -> resources.wiki.Page
	12, // 9: services.wiki.ListPageActivityRequest.pagination:type_name -> resources.common.database.PaginationRequest
	14, // 10: services.wiki.ListPageActivityResponse.pagination:type_name -> resources.common.database.PaginationResponse
	17, // 11: services.wiki.ListPageActivityResponse.activity:type_name -> resources.wiki.PageActivity
	0,  // 12: services.wiki.WikiService.ListPages:input_type -> services.wiki.ListPagesRequest
	2,  // 13: services.wiki.WikiService.GetPage:input_type -> services.wiki.GetPageRequest
	4,  // 14: services.wiki.WikiService.CreatePage:input_type -> services.wiki.CreatePageRequest
	6,  // 15: services.wiki.WikiService.UpdatePage:input_type -> services.wiki.UpdatePageRequest
	8,  // 16: services.wiki.WikiService.DeletePage:input_type -> services.wiki.DeletePageRequest
	10, // 17: services.wiki.WikiService.ListPageActivity:input_type -> services.wiki.ListPageActivityRequest
	1,  // 18: services.wiki.WikiService.ListPages:output_type -> services.wiki.ListPagesResponse
	3,  // 19: services.wiki.WikiService.GetPage:output_type -> services.wiki.GetPageResponse
	5,  // 20: services.wiki.WikiService.CreatePage:output_type -> services.wiki.CreatePageResponse
	7,  // 21: services.wiki.WikiService.UpdatePage:output_type -> services.wiki.UpdatePageResponse
	9,  // 22: services.wiki.WikiService.DeletePage:output_type -> services.wiki.DeletePageResponse
	11, // 23: services.wiki.WikiService.ListPageActivity:output_type -> services.wiki.ListPageActivityResponse
	18, // [18:24] is the sub-list for method output_type
	12, // [12:18] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_services_wiki_wiki_proto_init() }
func file_services_wiki_wiki_proto_init() {
	if File_services_wiki_wiki_proto != nil {
		return
	}
	file_services_wiki_wiki_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_wiki_wiki_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_wiki_wiki_proto_goTypes,
		DependencyIndexes: file_services_wiki_wiki_proto_depIdxs,
		MessageInfos:      file_services_wiki_wiki_proto_msgTypes,
	}.Build()
	File_services_wiki_wiki_proto = out.File
	file_services_wiki_wiki_proto_rawDesc = nil
	file_services_wiki_wiki_proto_goTypes = nil
	file_services_wiki_wiki_proto_depIdxs = nil
}