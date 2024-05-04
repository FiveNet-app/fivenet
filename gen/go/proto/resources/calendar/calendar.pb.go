// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resources/calendar/calendar.proto

package calendar

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

type RsvpResponses int32

const (
	RsvpResponses_RSVP_RESPONSES_UNSPECIFIED RsvpResponses = 0
	RsvpResponses_RSVP_RESPONSES_NO          RsvpResponses = 1
	RsvpResponses_RSVP_RESPONSES_MAYBE       RsvpResponses = 2
	RsvpResponses_RSVP_RESPONSES_YES         RsvpResponses = 3
)

// Enum value maps for RsvpResponses.
var (
	RsvpResponses_name = map[int32]string{
		0: "RSVP_RESPONSES_UNSPECIFIED",
		1: "RSVP_RESPONSES_NO",
		2: "RSVP_RESPONSES_MAYBE",
		3: "RSVP_RESPONSES_YES",
	}
	RsvpResponses_value = map[string]int32{
		"RSVP_RESPONSES_UNSPECIFIED": 0,
		"RSVP_RESPONSES_NO":          1,
		"RSVP_RESPONSES_MAYBE":       2,
		"RSVP_RESPONSES_YES":         3,
	}
)

func (x RsvpResponses) Enum() *RsvpResponses {
	p := new(RsvpResponses)
	*p = x
	return p
}

func (x RsvpResponses) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RsvpResponses) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_calendar_calendar_proto_enumTypes[0].Descriptor()
}

func (RsvpResponses) Type() protoreflect.EnumType {
	return &file_resources_calendar_calendar_proto_enumTypes[0]
}

func (x RsvpResponses) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RsvpResponses.Descriptor instead.
func (RsvpResponses) EnumDescriptor() ([]byte, []int) {
	return file_resources_calendar_calendar_proto_rawDescGZIP(), []int{0}
}

type Calendar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	DeletedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	Job       *string              `protobuf:"bytes,5,opt,name=job,proto3,oneof" json:"job,omitempty"`
	// @sanitize
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// @sanitize: method=StripTags
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Public      bool    `protobuf:"varint,8,opt,name=public,proto3" json:"public,omitempty"`
	Closed      bool    `protobuf:"varint,9,opt,name=closed,proto3" json:"closed,omitempty"`
	// @sanitize: method=StripTags
	Color      *string          `protobuf:"bytes,10,opt,name=color,proto3,oneof" json:"color,omitempty"`
	CreatorId  *int32           `protobuf:"varint,11,opt,name=creator_id,json=creatorId,proto3,oneof" json:"creator_id,omitempty"`
	Creator    *users.UserShort `protobuf:"bytes,12,opt,name=creator,proto3,oneof" json:"creator,omitempty" alias:"creator"` // @gotags: alias:"creator"
	CreatorJob string           `protobuf:"bytes,13,opt,name=creator_job,json=creatorJob,proto3" json:"creator_job,omitempty"`
	Access     *CalendarAccess  `protobuf:"bytes,14,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *Calendar) Reset() {
	*x = Calendar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_calendar_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calendar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calendar) ProtoMessage() {}

func (x *Calendar) ProtoReflect() protoreflect.Message {
	mi := &file_resources_calendar_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calendar.ProtoReflect.Descriptor instead.
func (*Calendar) Descriptor() ([]byte, []int) {
	return file_resources_calendar_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *Calendar) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Calendar) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Calendar) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Calendar) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Calendar) GetJob() string {
	if x != nil && x.Job != nil {
		return *x.Job
	}
	return ""
}

func (x *Calendar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Calendar) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Calendar) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *Calendar) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

func (x *Calendar) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *Calendar) GetCreatorId() int32 {
	if x != nil && x.CreatorId != nil {
		return *x.CreatorId
	}
	return 0
}

func (x *Calendar) GetCreator() *users.UserShort {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Calendar) GetCreatorJob() string {
	if x != nil {
		return x.CreatorJob
	}
	return ""
}

func (x *Calendar) GetAccess() *CalendarAccess {
	if x != nil {
		return x.Access
	}
	return nil
}

type CalendarEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	DeletedAt  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	CalendarId uint64               `protobuf:"varint,5,opt,name=calendar_id,json=calendarId,proto3" json:"calendar_id,omitempty"`
	Job        *string              `protobuf:"bytes,6,opt,name=job,proto3,oneof" json:"job,omitempty"`
	StartTime  *timestamp.Timestamp `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    *timestamp.Timestamp `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
	// @sanitize
	Title string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	// @sanitize
	Content    string           `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	Public     bool             `protobuf:"varint,11,opt,name=public,proto3" json:"public,omitempty"`
	RsvpOpen   *bool            `protobuf:"varint,12,opt,name=rsvp_open,json=rsvpOpen,proto3,oneof" json:"rsvp_open,omitempty"`
	CreatorId  *int32           `protobuf:"varint,13,opt,name=creator_id,json=creatorId,proto3,oneof" json:"creator_id,omitempty"`
	Creator    *users.UserShort `protobuf:"bytes,14,opt,name=creator,proto3,oneof" json:"creator,omitempty" alias:"creator"` // @gotags: alias:"creator"
	CreatorJob string           `protobuf:"bytes,15,opt,name=creator_job,json=creatorJob,proto3" json:"creator_job,omitempty"`
	Access     *CalendarAccess  `protobuf:"bytes,16,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *CalendarEntry) Reset() {
	*x = CalendarEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_calendar_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalendarEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarEntry) ProtoMessage() {}

func (x *CalendarEntry) ProtoReflect() protoreflect.Message {
	mi := &file_resources_calendar_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarEntry.ProtoReflect.Descriptor instead.
func (*CalendarEntry) Descriptor() ([]byte, []int) {
	return file_resources_calendar_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *CalendarEntry) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CalendarEntry) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CalendarEntry) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CalendarEntry) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *CalendarEntry) GetCalendarId() uint64 {
	if x != nil {
		return x.CalendarId
	}
	return 0
}

func (x *CalendarEntry) GetJob() string {
	if x != nil && x.Job != nil {
		return *x.Job
	}
	return ""
}

func (x *CalendarEntry) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *CalendarEntry) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *CalendarEntry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CalendarEntry) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CalendarEntry) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *CalendarEntry) GetRsvpOpen() bool {
	if x != nil && x.RsvpOpen != nil {
		return *x.RsvpOpen
	}
	return false
}

func (x *CalendarEntry) GetCreatorId() int32 {
	if x != nil && x.CreatorId != nil {
		return *x.CreatorId
	}
	return 0
}

func (x *CalendarEntry) GetCreator() *users.UserShort {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *CalendarEntry) GetCreatorJob() string {
	if x != nil {
		return x.CreatorJob
	}
	return ""
}

func (x *CalendarEntry) GetAccess() *CalendarAccess {
	if x != nil {
		return x.Access
	}
	return nil
}

type CalendarEntryRSVP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId   uint64               `protobuf:"varint,1,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UserId    int32                `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User      *users.UserShort     `protobuf:"bytes,4,opt,name=user,proto3,oneof" json:"user,omitempty"`
	Response  RsvpResponses        `protobuf:"varint,5,opt,name=response,proto3,enum=resources.calendar.RsvpResponses" json:"response,omitempty"`
}

func (x *CalendarEntryRSVP) Reset() {
	*x = CalendarEntryRSVP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_calendar_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalendarEntryRSVP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarEntryRSVP) ProtoMessage() {}

func (x *CalendarEntryRSVP) ProtoReflect() protoreflect.Message {
	mi := &file_resources_calendar_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarEntryRSVP.ProtoReflect.Descriptor instead.
func (*CalendarEntryRSVP) Descriptor() ([]byte, []int) {
	return file_resources_calendar_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *CalendarEntryRSVP) GetEntryId() uint64 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *CalendarEntryRSVP) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CalendarEntryRSVP) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CalendarEntryRSVP) GetUser() *users.UserShort {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CalendarEntryRSVP) GetResponse() RsvpResponses {
	if x != nil {
		return x.Response
	}
	return RsvpResponses_RSVP_RESPONSES_UNSPECIFIED
}

var File_resources_calendar_calendar_proto protoreflect.FileDescriptor

var file_resources_calendar_calendar_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x1a, 0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xde, 0x05, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x02, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x48, 0x03, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x03, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x04, 0x48,
	0x04, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x0c, 0x48, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x07, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x6a, 0x6f, 0x62, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x18, 0x14, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x12, 0x3a,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6a, 0x6f, 0x62, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x22, 0xe6, 0x06, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52,
	0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x03, 0x6a,
	0x6f, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18,
	0x14, 0x48, 0x03, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x03, 0x18, 0x80, 0x04, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa,
	0x42, 0x08, 0x72, 0x06, 0x10, 0x14, 0x28, 0xc0, 0x84, 0x3d, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x20, 0x0a, 0x09, 0x72,
	0x73, 0x76, 0x70, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05,
	0x52, 0x08, 0x72, 0x73, 0x76, 0x70, 0x4f, 0x70, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x06, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x07,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x12, 0x3a, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x6a, 0x6f, 0x62, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x73, 0x76, 0x70, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xae, 0x02,
	0x0a, 0x11, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x53, 0x56, 0x50, 0x12, 0x1d, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x48, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x47, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2e, 0x52, 0x73, 0x76, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2a, 0x78,
	0x0a, 0x0d, 0x52, 0x73, 0x76, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x1a, 0x52, 0x53, 0x56, 0x50, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x52, 0x53, 0x56, 0x50, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x53, 0x5f, 0x4e, 0x4f, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x53, 0x56, 0x50, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x53, 0x5f, 0x4d, 0x41, 0x59, 0x42, 0x45, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x52, 0x53, 0x56, 0x50, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x53, 0x5f, 0x59, 0x45, 0x53, 0x10, 0x03, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61,
	0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x3b, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_calendar_calendar_proto_rawDescOnce sync.Once
	file_resources_calendar_calendar_proto_rawDescData = file_resources_calendar_calendar_proto_rawDesc
)

func file_resources_calendar_calendar_proto_rawDescGZIP() []byte {
	file_resources_calendar_calendar_proto_rawDescOnce.Do(func() {
		file_resources_calendar_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_calendar_calendar_proto_rawDescData)
	})
	return file_resources_calendar_calendar_proto_rawDescData
}

var file_resources_calendar_calendar_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_calendar_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_calendar_calendar_proto_goTypes = []interface{}{
	(RsvpResponses)(0),          // 0: resources.calendar.RsvpResponses
	(*Calendar)(nil),            // 1: resources.calendar.Calendar
	(*CalendarEntry)(nil),       // 2: resources.calendar.CalendarEntry
	(*CalendarEntryRSVP)(nil),   // 3: resources.calendar.CalendarEntryRSVP
	(*timestamp.Timestamp)(nil), // 4: resources.timestamp.Timestamp
	(*users.UserShort)(nil),     // 5: resources.users.UserShort
	(*CalendarAccess)(nil),      // 6: resources.calendar.CalendarAccess
}
var file_resources_calendar_calendar_proto_depIdxs = []int32{
	4,  // 0: resources.calendar.Calendar.created_at:type_name -> resources.timestamp.Timestamp
	4,  // 1: resources.calendar.Calendar.updated_at:type_name -> resources.timestamp.Timestamp
	4,  // 2: resources.calendar.Calendar.deleted_at:type_name -> resources.timestamp.Timestamp
	5,  // 3: resources.calendar.Calendar.creator:type_name -> resources.users.UserShort
	6,  // 4: resources.calendar.Calendar.access:type_name -> resources.calendar.CalendarAccess
	4,  // 5: resources.calendar.CalendarEntry.created_at:type_name -> resources.timestamp.Timestamp
	4,  // 6: resources.calendar.CalendarEntry.updated_at:type_name -> resources.timestamp.Timestamp
	4,  // 7: resources.calendar.CalendarEntry.deleted_at:type_name -> resources.timestamp.Timestamp
	4,  // 8: resources.calendar.CalendarEntry.start_time:type_name -> resources.timestamp.Timestamp
	4,  // 9: resources.calendar.CalendarEntry.end_time:type_name -> resources.timestamp.Timestamp
	5,  // 10: resources.calendar.CalendarEntry.creator:type_name -> resources.users.UserShort
	6,  // 11: resources.calendar.CalendarEntry.access:type_name -> resources.calendar.CalendarAccess
	4,  // 12: resources.calendar.CalendarEntryRSVP.created_at:type_name -> resources.timestamp.Timestamp
	5,  // 13: resources.calendar.CalendarEntryRSVP.user:type_name -> resources.users.UserShort
	0,  // 14: resources.calendar.CalendarEntryRSVP.response:type_name -> resources.calendar.RsvpResponses
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_resources_calendar_calendar_proto_init() }
func file_resources_calendar_calendar_proto_init() {
	if File_resources_calendar_calendar_proto != nil {
		return
	}
	file_resources_calendar_access_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_resources_calendar_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calendar); i {
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
		file_resources_calendar_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalendarEntry); i {
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
		file_resources_calendar_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalendarEntryRSVP); i {
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
	file_resources_calendar_calendar_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_calendar_calendar_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_calendar_calendar_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_calendar_calendar_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_calendar_calendar_proto_goTypes,
		DependencyIndexes: file_resources_calendar_calendar_proto_depIdxs,
		EnumInfos:         file_resources_calendar_calendar_proto_enumTypes,
		MessageInfos:      file_resources_calendar_calendar_proto_msgTypes,
	}.Build()
	File_resources_calendar_calendar_proto = out.File
	file_resources_calendar_calendar_proto_rawDesc = nil
	file_resources_calendar_calendar_proto_goTypes = nil
	file_resources_calendar_calendar_proto_depIdxs = nil
}
