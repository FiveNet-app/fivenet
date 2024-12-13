// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: resources/mailer/thread.proto

package mailer

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/fivenet-app/fivenet/gen/go/proto/resources/timestamp"
	users "github.com/fivenet-app/fivenet/gen/go/proto/resources/users"
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

type Thread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt      *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	DeletedAt      *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	CreatorEmailId uint64               `protobuf:"varint,5,opt,name=creator_email_id,json=creatorEmailId,proto3" json:"creator_email_id,omitempty"`
	CreatorEmail   *Email               `protobuf:"bytes,6,opt,name=creator_email,json=creatorEmail,proto3,oneof" json:"creator_email,omitempty"`
	CreatorId      *int32               `protobuf:"varint,7,opt,name=creator_id,json=creatorId,proto3,oneof" json:"creator_id,omitempty"`
	Creator        *users.UserShort     `protobuf:"bytes,8,opt,name=creator,proto3,oneof" json:"creator,omitempty" alias:"creator"` // @gotags: alias:"creator"
	// @sanitize: method=StripTags
	Title      string                  `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	Recipients []*ThreadRecipientEmail `protobuf:"bytes,10,rep,name=recipients,proto3" json:"recipients,omitempty"`
	State      *ThreadState            `protobuf:"bytes,11,opt,name=state,proto3,oneof" json:"state,omitempty" alias:"thread_state"` // @gotags: alias:"thread_state"
}

func (x *Thread) Reset() {
	*x = Thread{}
	mi := &file_resources_mailer_thread_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Thread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thread) ProtoMessage() {}

func (x *Thread) ProtoReflect() protoreflect.Message {
	mi := &file_resources_mailer_thread_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thread.ProtoReflect.Descriptor instead.
func (*Thread) Descriptor() ([]byte, []int) {
	return file_resources_mailer_thread_proto_rawDescGZIP(), []int{0}
}

func (x *Thread) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Thread) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Thread) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Thread) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Thread) GetCreatorEmailId() uint64 {
	if x != nil {
		return x.CreatorEmailId
	}
	return 0
}

func (x *Thread) GetCreatorEmail() *Email {
	if x != nil {
		return x.CreatorEmail
	}
	return nil
}

func (x *Thread) GetCreatorId() int32 {
	if x != nil && x.CreatorId != nil {
		return *x.CreatorId
	}
	return 0
}

func (x *Thread) GetCreator() *users.UserShort {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Thread) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Thread) GetRecipients() []*ThreadRecipientEmail {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *Thread) GetState() *ThreadState {
	if x != nil {
		return x.State
	}
	return nil
}

type ThreadRecipientEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	TargetId  uint64               `protobuf:"varint,4,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty" alias:"thread_id"` // @gotags: alias:"thread_id"
	EmailId   uint64               `protobuf:"varint,5,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	Email     *Email               `protobuf:"bytes,6,opt,name=email,proto3,oneof" json:"email,omitempty"`
}

func (x *ThreadRecipientEmail) Reset() {
	*x = ThreadRecipientEmail{}
	mi := &file_resources_mailer_thread_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThreadRecipientEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadRecipientEmail) ProtoMessage() {}

func (x *ThreadRecipientEmail) ProtoReflect() protoreflect.Message {
	mi := &file_resources_mailer_thread_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadRecipientEmail.ProtoReflect.Descriptor instead.
func (*ThreadRecipientEmail) Descriptor() ([]byte, []int) {
	return file_resources_mailer_thread_proto_rawDescGZIP(), []int{1}
}

func (x *ThreadRecipientEmail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ThreadRecipientEmail) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ThreadRecipientEmail) GetTargetId() uint64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *ThreadRecipientEmail) GetEmailId() uint64 {
	if x != nil {
		return x.EmailId
	}
	return 0
}

func (x *ThreadRecipientEmail) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type ThreadState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadId  uint64               `protobuf:"varint,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	EmailId   uint64               `protobuf:"varint,2,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	LastRead  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_read,json=lastRead,proto3,oneof" json:"last_read,omitempty"`
	Unread    bool                 `protobuf:"varint,4,opt,name=unread,proto3" json:"unread,omitempty"`
	Important bool                 `protobuf:"varint,5,opt,name=important,proto3" json:"important,omitempty"`
	Favorite  bool                 `protobuf:"varint,6,opt,name=favorite,proto3" json:"favorite,omitempty"`
	Muted     bool                 `protobuf:"varint,7,opt,name=muted,proto3" json:"muted,omitempty"`
	Archived  bool                 `protobuf:"varint,8,opt,name=archived,proto3" json:"archived,omitempty"`
}

func (x *ThreadState) Reset() {
	*x = ThreadState{}
	mi := &file_resources_mailer_thread_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ThreadState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadState) ProtoMessage() {}

func (x *ThreadState) ProtoReflect() protoreflect.Message {
	mi := &file_resources_mailer_thread_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadState.ProtoReflect.Descriptor instead.
func (*ThreadState) Descriptor() ([]byte, []int) {
	return file_resources_mailer_thread_proto_rawDescGZIP(), []int{2}
}

func (x *ThreadState) GetThreadId() uint64 {
	if x != nil {
		return x.ThreadId
	}
	return 0
}

func (x *ThreadState) GetEmailId() uint64 {
	if x != nil {
		return x.EmailId
	}
	return 0
}

func (x *ThreadState) GetLastRead() *timestamp.Timestamp {
	if x != nil {
		return x.LastRead
	}
	return nil
}

func (x *ThreadState) GetUnread() bool {
	if x != nil {
		return x.Unread
	}
	return false
}

func (x *ThreadState) GetImportant() bool {
	if x != nil {
		return x.Important
	}
	return false
}

func (x *ThreadState) GetFavorite() bool {
	if x != nil {
		return x.Favorite
	}
	return false
}

func (x *ThreadState) GetMuted() bool {
	if x != nil {
		return x.Muted
	}
	return false
}

func (x *ThreadState) GetArchived() bool {
	if x != nil {
		return x.Archived
	}
	return false
}

var File_resources_mailer_thread_proto protoreflect.FileDescriptor

var file_resources_mailer_thread_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x1a, 0x1c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x05, 0x0a, 0x06, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x01, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x2c, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x41,
	0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x48, 0x02,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x04, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10,
	0x03, 0x18, 0xff, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10,
	0x14, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x05, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xfb, 0x01, 0x0a,
	0x14, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x08, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x48, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x09, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x08, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e,
	0x72, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d,
	0x75, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x42, 0x45,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76,
	0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x3b, 0x6d,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_mailer_thread_proto_rawDescOnce sync.Once
	file_resources_mailer_thread_proto_rawDescData = file_resources_mailer_thread_proto_rawDesc
)

func file_resources_mailer_thread_proto_rawDescGZIP() []byte {
	file_resources_mailer_thread_proto_rawDescOnce.Do(func() {
		file_resources_mailer_thread_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_mailer_thread_proto_rawDescData)
	})
	return file_resources_mailer_thread_proto_rawDescData
}

var file_resources_mailer_thread_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_mailer_thread_proto_goTypes = []any{
	(*Thread)(nil),               // 0: resources.mailer.Thread
	(*ThreadRecipientEmail)(nil), // 1: resources.mailer.ThreadRecipientEmail
	(*ThreadState)(nil),          // 2: resources.mailer.ThreadState
	(*timestamp.Timestamp)(nil),  // 3: resources.timestamp.Timestamp
	(*Email)(nil),                // 4: resources.mailer.Email
	(*users.UserShort)(nil),      // 5: resources.users.UserShort
}
var file_resources_mailer_thread_proto_depIdxs = []int32{
	3,  // 0: resources.mailer.Thread.created_at:type_name -> resources.timestamp.Timestamp
	3,  // 1: resources.mailer.Thread.updated_at:type_name -> resources.timestamp.Timestamp
	3,  // 2: resources.mailer.Thread.deleted_at:type_name -> resources.timestamp.Timestamp
	4,  // 3: resources.mailer.Thread.creator_email:type_name -> resources.mailer.Email
	5,  // 4: resources.mailer.Thread.creator:type_name -> resources.users.UserShort
	1,  // 5: resources.mailer.Thread.recipients:type_name -> resources.mailer.ThreadRecipientEmail
	2,  // 6: resources.mailer.Thread.state:type_name -> resources.mailer.ThreadState
	3,  // 7: resources.mailer.ThreadRecipientEmail.created_at:type_name -> resources.timestamp.Timestamp
	4,  // 8: resources.mailer.ThreadRecipientEmail.email:type_name -> resources.mailer.Email
	3,  // 9: resources.mailer.ThreadState.last_read:type_name -> resources.timestamp.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_resources_mailer_thread_proto_init() }
func file_resources_mailer_thread_proto_init() {
	if File_resources_mailer_thread_proto != nil {
		return
	}
	file_resources_mailer_email_proto_init()
	file_resources_mailer_thread_proto_msgTypes[0].OneofWrappers = []any{}
	file_resources_mailer_thread_proto_msgTypes[1].OneofWrappers = []any{}
	file_resources_mailer_thread_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_mailer_thread_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_mailer_thread_proto_goTypes,
		DependencyIndexes: file_resources_mailer_thread_proto_depIdxs,
		MessageInfos:      file_resources_mailer_thread_proto_msgTypes,
	}.Build()
	File_resources_mailer_thread_proto = out.File
	file_resources_mailer_thread_proto_rawDesc = nil
	file_resources_mailer_thread_proto_goTypes = nil
	file_resources_mailer_thread_proto_depIdxs = nil
}
