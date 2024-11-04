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
	Prefix string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
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

func (x *ListPagesRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
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

	Path *string `protobuf:"bytes,1,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Id   *uint64 `protobuf:"varint,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
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

func (x *GetPageRequest) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *GetPageRequest) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
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

type CreateOrUpdatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *wiki.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *CreateOrUpdatePageRequest) Reset() {
	*x = CreateOrUpdatePageRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrUpdatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdatePageRequest) ProtoMessage() {}

func (x *CreateOrUpdatePageRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateOrUpdatePageRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdatePageRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrUpdatePageRequest) GetPage() *wiki.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type CreateOrUpdatePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *wiki.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *CreateOrUpdatePageResponse) Reset() {
	*x = CreateOrUpdatePageResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrUpdatePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdatePageResponse) ProtoMessage() {}

func (x *CreateOrUpdatePageResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateOrUpdatePageResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdatePageResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{5}
}

func (x *CreateOrUpdatePageResponse) GetPage() *wiki.Page {
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
	mi := &file_services_wiki_wiki_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageRequest) ProtoMessage() {}

func (x *DeletePageRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeletePageRequest.ProtoReflect.Descriptor instead.
func (*DeletePageRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{6}
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
	mi := &file_services_wiki_wiki_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageResponse) ProtoMessage() {}

func (x *DeletePageResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeletePageResponse.ProtoReflect.Descriptor instead.
func (*DeletePageResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{7}
}

type GetPageHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPageHistoryRequest) Reset() {
	*x = GetPageHistoryRequest{}
	mi := &file_services_wiki_wiki_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPageHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageHistoryRequest) ProtoMessage() {}

func (x *GetPageHistoryRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPageHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetPageHistoryRequest) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{8}
}

func (x *GetPageHistoryRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPageHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *database.PaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Activity   []*wiki.PageActivity         `protobuf:"bytes,2,rep,name=activity,proto3" json:"activity,omitempty"`
}

func (x *GetPageHistoryResponse) Reset() {
	*x = GetPageHistoryResponse{}
	mi := &file_services_wiki_wiki_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPageHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageHistoryResponse) ProtoMessage() {}

func (x *GetPageHistoryResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPageHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetPageHistoryResponse) Descriptor() ([]byte, []int) {
	return file_services_wiki_wiki_proto_rawDescGZIP(), []int{9}
}

func (x *GetPageHistoryResponse) GetPagination() *database.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetPageHistoryResponse) GetActivity() []*wiki.PageActivity {
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
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x48, 0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x9d,
	0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x22, 0x52,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x05, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x45, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77,
	0x69, 0x6b, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x27,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b,
	0x69, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x32, 0xc4, 0x03, 0x0a, 0x0b, 0x57, 0x69, 0x6b,
	0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77,
	0x69, 0x6b, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69,
	0x6b, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x69, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69,
	0x6b, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x69,
	0x6b, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65,
	0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x3b, 0x77, 0x69, 0x6b,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_services_wiki_wiki_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_wiki_wiki_proto_goTypes = []any{
	(*ListPagesRequest)(nil),            // 0: services.wiki.ListPagesRequest
	(*ListPagesResponse)(nil),           // 1: services.wiki.ListPagesResponse
	(*GetPageRequest)(nil),              // 2: services.wiki.GetPageRequest
	(*GetPageResponse)(nil),             // 3: services.wiki.GetPageResponse
	(*CreateOrUpdatePageRequest)(nil),   // 4: services.wiki.CreateOrUpdatePageRequest
	(*CreateOrUpdatePageResponse)(nil),  // 5: services.wiki.CreateOrUpdatePageResponse
	(*DeletePageRequest)(nil),           // 6: services.wiki.DeletePageRequest
	(*DeletePageResponse)(nil),          // 7: services.wiki.DeletePageResponse
	(*GetPageHistoryRequest)(nil),       // 8: services.wiki.GetPageHistoryRequest
	(*GetPageHistoryResponse)(nil),      // 9: services.wiki.GetPageHistoryResponse
	(*database.PaginationRequest)(nil),  // 10: resources.common.database.PaginationRequest
	(*database.Sort)(nil),               // 11: resources.common.database.Sort
	(*database.PaginationResponse)(nil), // 12: resources.common.database.PaginationResponse
	(*wiki.PageShort)(nil),              // 13: resources.wiki.PageShort
	(*wiki.Page)(nil),                   // 14: resources.wiki.Page
	(*wiki.PageActivity)(nil),           // 15: resources.wiki.PageActivity
}
var file_services_wiki_wiki_proto_depIdxs = []int32{
	10, // 0: services.wiki.ListPagesRequest.pagination:type_name -> resources.common.database.PaginationRequest
	11, // 1: services.wiki.ListPagesRequest.sort:type_name -> resources.common.database.Sort
	12, // 2: services.wiki.ListPagesResponse.pagination:type_name -> resources.common.database.PaginationResponse
	13, // 3: services.wiki.ListPagesResponse.pages:type_name -> resources.wiki.PageShort
	14, // 4: services.wiki.GetPageResponse.page:type_name -> resources.wiki.Page
	14, // 5: services.wiki.CreateOrUpdatePageRequest.page:type_name -> resources.wiki.Page
	14, // 6: services.wiki.CreateOrUpdatePageResponse.page:type_name -> resources.wiki.Page
	12, // 7: services.wiki.GetPageHistoryResponse.pagination:type_name -> resources.common.database.PaginationResponse
	15, // 8: services.wiki.GetPageHistoryResponse.activity:type_name -> resources.wiki.PageActivity
	0,  // 9: services.wiki.WikiService.ListPages:input_type -> services.wiki.ListPagesRequest
	2,  // 10: services.wiki.WikiService.GetPage:input_type -> services.wiki.GetPageRequest
	4,  // 11: services.wiki.WikiService.CreateOrUpdatePage:input_type -> services.wiki.CreateOrUpdatePageRequest
	6,  // 12: services.wiki.WikiService.DeletePage:input_type -> services.wiki.DeletePageRequest
	8,  // 13: services.wiki.WikiService.GetPageHistory:input_type -> services.wiki.GetPageHistoryRequest
	1,  // 14: services.wiki.WikiService.ListPages:output_type -> services.wiki.ListPagesResponse
	3,  // 15: services.wiki.WikiService.GetPage:output_type -> services.wiki.GetPageResponse
	5,  // 16: services.wiki.WikiService.CreateOrUpdatePage:output_type -> services.wiki.CreateOrUpdatePageResponse
	7,  // 17: services.wiki.WikiService.DeletePage:output_type -> services.wiki.DeletePageResponse
	9,  // 18: services.wiki.WikiService.GetPageHistory:output_type -> services.wiki.GetPageHistoryResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_services_wiki_wiki_proto_init() }
func file_services_wiki_wiki_proto_init() {
	if File_services_wiki_wiki_proto != nil {
		return
	}
	file_services_wiki_wiki_proto_msgTypes[0].OneofWrappers = []any{}
	file_services_wiki_wiki_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_wiki_wiki_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
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
