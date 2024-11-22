// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.20.3
// source: resources/calendar/access.proto

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

type AccessLevel int32

const (
	AccessLevel_ACCESS_LEVEL_UNSPECIFIED AccessLevel = 0
	AccessLevel_ACCESS_LEVEL_BLOCKED     AccessLevel = 1
	AccessLevel_ACCESS_LEVEL_VIEW        AccessLevel = 2
	AccessLevel_ACCESS_LEVEL_SHARE       AccessLevel = 3
	AccessLevel_ACCESS_LEVEL_EDIT        AccessLevel = 4
	AccessLevel_ACCESS_LEVEL_MANAGE      AccessLevel = 5
)

// Enum value maps for AccessLevel.
var (
	AccessLevel_name = map[int32]string{
		0: "ACCESS_LEVEL_UNSPECIFIED",
		1: "ACCESS_LEVEL_BLOCKED",
		2: "ACCESS_LEVEL_VIEW",
		3: "ACCESS_LEVEL_SHARE",
		4: "ACCESS_LEVEL_EDIT",
		5: "ACCESS_LEVEL_MANAGE",
	}
	AccessLevel_value = map[string]int32{
		"ACCESS_LEVEL_UNSPECIFIED": 0,
		"ACCESS_LEVEL_BLOCKED":     1,
		"ACCESS_LEVEL_VIEW":        2,
		"ACCESS_LEVEL_SHARE":       3,
		"ACCESS_LEVEL_EDIT":        4,
		"ACCESS_LEVEL_MANAGE":      5,
	}
)

func (x AccessLevel) Enum() *AccessLevel {
	p := new(AccessLevel)
	*p = x
	return p
}

func (x AccessLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_calendar_access_proto_enumTypes[0].Descriptor()
}

func (AccessLevel) Type() protoreflect.EnumType {
	return &file_resources_calendar_access_proto_enumTypes[0]
}

func (x AccessLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessLevel.Descriptor instead.
func (AccessLevel) EnumDescriptor() ([]byte, []int) {
	return file_resources_calendar_access_proto_rawDescGZIP(), []int{0}
}

type CalendarAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs  []*CalendarJobAccess  `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty" alias:"job_access"`   // @gotags: alias:"job_access"
	Users []*CalendarUserAccess `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty" alias:"user_access"` // @gotags: alias:"user_access"
}

func (x *CalendarAccess) Reset() {
	*x = CalendarAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_calendar_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalendarAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarAccess) ProtoMessage() {}

func (x *CalendarAccess) ProtoReflect() protoreflect.Message {
	mi := &file_resources_calendar_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarAccess.ProtoReflect.Descriptor instead.
func (*CalendarAccess) Descriptor() ([]byte, []int) {
	return file_resources_calendar_access_proto_rawDescGZIP(), []int{0}
}

func (x *CalendarAccess) GetJobs() []*CalendarJobAccess {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *CalendarAccess) GetUsers() []*CalendarUserAccess {
	if x != nil {
		return x.Users
	}
	return nil
}

type CalendarJobAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	TargetId      uint64               `protobuf:"varint,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty" alias:"calendar_id"` // @gotags: alias:"calendar_id"
	Job           string               `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	JobLabel      *string              `protobuf:"bytes,5,opt,name=job_label,json=jobLabel,proto3,oneof" json:"job_label,omitempty"`
	MinimumGrade  int32                `protobuf:"varint,6,opt,name=minimum_grade,json=minimumGrade,proto3" json:"minimum_grade,omitempty"`
	JobGradeLabel *string              `protobuf:"bytes,7,opt,name=job_grade_label,json=jobGradeLabel,proto3,oneof" json:"job_grade_label,omitempty"`
	Access        AccessLevel          `protobuf:"varint,8,opt,name=access,proto3,enum=resources.calendar.AccessLevel" json:"access,omitempty"`
}

func (x *CalendarJobAccess) Reset() {
	*x = CalendarJobAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_calendar_access_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalendarJobAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarJobAccess) ProtoMessage() {}

func (x *CalendarJobAccess) ProtoReflect() protoreflect.Message {
	mi := &file_resources_calendar_access_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarJobAccess.ProtoReflect.Descriptor instead.
func (*CalendarJobAccess) Descriptor() ([]byte, []int) {
	return file_resources_calendar_access_proto_rawDescGZIP(), []int{1}
}

func (x *CalendarJobAccess) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CalendarJobAccess) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CalendarJobAccess) GetTargetId() uint64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *CalendarJobAccess) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *CalendarJobAccess) GetJobLabel() string {
	if x != nil && x.JobLabel != nil {
		return *x.JobLabel
	}
	return ""
}

func (x *CalendarJobAccess) GetMinimumGrade() int32 {
	if x != nil {
		return x.MinimumGrade
	}
	return 0
}

func (x *CalendarJobAccess) GetJobGradeLabel() string {
	if x != nil && x.JobGradeLabel != nil {
		return *x.JobGradeLabel
	}
	return ""
}

func (x *CalendarJobAccess) GetAccess() AccessLevel {
	if x != nil {
		return x.Access
	}
	return AccessLevel_ACCESS_LEVEL_UNSPECIFIED
}

type CalendarUserAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	TargetId  uint64               `protobuf:"varint,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty" alias:"calendar_id"` // @gotags: alias:"calendar_id"
	UserId    int32                `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User      *users.UserShort     `protobuf:"bytes,5,opt,name=user,proto3,oneof" json:"user,omitempty"`
	Access    AccessLevel          `protobuf:"varint,6,opt,name=access,proto3,enum=resources.calendar.AccessLevel" json:"access,omitempty"`
}

func (x *CalendarUserAccess) Reset() {
	*x = CalendarUserAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_calendar_access_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalendarUserAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarUserAccess) ProtoMessage() {}

func (x *CalendarUserAccess) ProtoReflect() protoreflect.Message {
	mi := &file_resources_calendar_access_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarUserAccess.ProtoReflect.Descriptor instead.
func (*CalendarUserAccess) Descriptor() ([]byte, []int) {
	return file_resources_calendar_access_proto_rawDescGZIP(), []int{2}
}

func (x *CalendarUserAccess) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CalendarUserAccess) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CalendarUserAccess) GetTargetId() uint64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *CalendarUserAccess) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CalendarUserAccess) GetUser() *users.UserShort {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CalendarUserAccess) GetAccess() AccessLevel {
	if x != nil {
		return x.Access
	}
	return AccessLevel_ACCESS_LEVEL_UNSPECIFIED
}

var File_resources_calendar_access_proto protoreflect.FileDescriptor

var file_resources_calendar_access_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9d, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x4a,
	0x6f, 0x62, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x10, 0x14, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x46, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x14, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0xaa, 0x03, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x4a, 0x6f, 0x62,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x29, 0x0a, 0x09, 0x6a, 0x6f,
	0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x01, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d,
	0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x02, 0x52, 0x0d, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xbf, 0x02,
	0x0a, 0x12, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x09,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x33, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2a,
	0xa4, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1c, 0x0a, 0x18, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x14, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53,
	0x48, 0x41, 0x52, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x10, 0x04, 0x12, 0x17, 0x0a,
	0x13, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x41,
	0x4e, 0x41, 0x47, 0x45, 0x10, 0x05, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70,
	0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x3b, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_calendar_access_proto_rawDescOnce sync.Once
	file_resources_calendar_access_proto_rawDescData = file_resources_calendar_access_proto_rawDesc
)

func file_resources_calendar_access_proto_rawDescGZIP() []byte {
	file_resources_calendar_access_proto_rawDescOnce.Do(func() {
		file_resources_calendar_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_calendar_access_proto_rawDescData)
	})
	return file_resources_calendar_access_proto_rawDescData
}

var file_resources_calendar_access_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_calendar_access_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_calendar_access_proto_goTypes = []interface{}{
	(AccessLevel)(0),            // 0: resources.calendar.AccessLevel
	(*CalendarAccess)(nil),      // 1: resources.calendar.CalendarAccess
	(*CalendarJobAccess)(nil),   // 2: resources.calendar.CalendarJobAccess
	(*CalendarUserAccess)(nil),  // 3: resources.calendar.CalendarUserAccess
	(*timestamp.Timestamp)(nil), // 4: resources.timestamp.Timestamp
	(*users.UserShort)(nil),     // 5: resources.users.UserShort
}
var file_resources_calendar_access_proto_depIdxs = []int32{
	2, // 0: resources.calendar.CalendarAccess.jobs:type_name -> resources.calendar.CalendarJobAccess
	3, // 1: resources.calendar.CalendarAccess.users:type_name -> resources.calendar.CalendarUserAccess
	4, // 2: resources.calendar.CalendarJobAccess.created_at:type_name -> resources.timestamp.Timestamp
	0, // 3: resources.calendar.CalendarJobAccess.access:type_name -> resources.calendar.AccessLevel
	4, // 4: resources.calendar.CalendarUserAccess.created_at:type_name -> resources.timestamp.Timestamp
	5, // 5: resources.calendar.CalendarUserAccess.user:type_name -> resources.users.UserShort
	0, // 6: resources.calendar.CalendarUserAccess.access:type_name -> resources.calendar.AccessLevel
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_resources_calendar_access_proto_init() }
func file_resources_calendar_access_proto_init() {
	if File_resources_calendar_access_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_calendar_access_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalendarAccess); i {
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
		file_resources_calendar_access_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalendarJobAccess); i {
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
		file_resources_calendar_access_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalendarUserAccess); i {
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
	file_resources_calendar_access_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_calendar_access_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_calendar_access_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_calendar_access_proto_goTypes,
		DependencyIndexes: file_resources_calendar_access_proto_depIdxs,
		EnumInfos:         file_resources_calendar_access_proto_enumTypes,
		MessageInfos:      file_resources_calendar_access_proto_msgTypes,
	}.Build()
	File_resources_calendar_access_proto = out.File
	file_resources_calendar_access_proto_rawDesc = nil
	file_resources_calendar_access_proto_goTypes = nil
	file_resources_calendar_access_proto_depIdxs = nil
}
