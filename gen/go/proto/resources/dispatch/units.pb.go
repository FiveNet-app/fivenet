// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: resources/dispatch/units.proto

package dispatch

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/galexrt/fivenet/gen/go/proto/resources/timestamp"
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

type UNIT_STATUS int32

const (
	UNIT_STATUS_UNKNOWN      UNIT_STATUS = 0
	UNIT_STATUS_USER_ADDED   UNIT_STATUS = 1
	UNIT_STATUS_USER_REMOVED UNIT_STATUS = 2
	UNIT_STATUS_UNAVAILABLE  UNIT_STATUS = 3
	UNIT_STATUS_AVAILABLE    UNIT_STATUS = 4
	UNIT_STATUS_ON_BREAK     UNIT_STATUS = 5
	UNIT_STATUS_BUSY         UNIT_STATUS = 6
)

// Enum value maps for UNIT_STATUS.
var (
	UNIT_STATUS_name = map[int32]string{
		0: "UNKNOWN",
		1: "USER_ADDED",
		2: "USER_REMOVED",
		3: "UNAVAILABLE",
		4: "AVAILABLE",
		5: "ON_BREAK",
		6: "BUSY",
	}
	UNIT_STATUS_value = map[string]int32{
		"UNKNOWN":      0,
		"USER_ADDED":   1,
		"USER_REMOVED": 2,
		"UNAVAILABLE":  3,
		"AVAILABLE":    4,
		"ON_BREAK":     5,
		"BUSY":         6,
	}
)

func (x UNIT_STATUS) Enum() *UNIT_STATUS {
	p := new(UNIT_STATUS)
	*p = x
	return p
}

func (x UNIT_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UNIT_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_dispatch_units_proto_enumTypes[0].Descriptor()
}

func (UNIT_STATUS) Type() protoreflect.EnumType {
	return &file_resources_dispatch_units_proto_enumTypes[0]
}

func (x UNIT_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UNIT_STATUS.Descriptor instead.
func (UNIT_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_resources_dispatch_units_proto_rawDescGZIP(), []int{0}
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	Job         string               `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	Name        string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Initials    string               `protobuf:"bytes,6,opt,name=initials,proto3" json:"initials,omitempty"`
	Color       *string              `protobuf:"bytes,7,opt,name=color,proto3,oneof" json:"color,omitempty"`
	Description *string              `protobuf:"bytes,8,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Status      *UnitStatus          `protobuf:"bytes,9,opt,name=status,proto3,oneof" json:"status,omitempty"`
	// repeated UnitStatus statuses = 10;
	Users []*UnitAssignment `protobuf:"bytes,11,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_dispatch_units_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dispatch_units_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_resources_dispatch_units_proto_rawDescGZIP(), []int{0}
}

func (x *Unit) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Unit) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Unit) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Unit) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Unit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Unit) GetInitials() string {
	if x != nil {
		return x.Initials
	}
	return ""
}

func (x *Unit) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *Unit) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Unit) GetStatus() *UnitStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Unit) GetUsers() []*UnitAssignment {
	if x != nil {
		return x.Users
	}
	return nil
}

type UnitAssignments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitId uint64            `protobuf:"varint,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	Job    string            `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Users  []*UnitAssignment `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UnitAssignments) Reset() {
	*x = UnitAssignments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_dispatch_units_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitAssignments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitAssignments) ProtoMessage() {}

func (x *UnitAssignments) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dispatch_units_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitAssignments.ProtoReflect.Descriptor instead.
func (*UnitAssignments) Descriptor() ([]byte, []int) {
	return file_resources_dispatch_units_proto_rawDescGZIP(), []int{1}
}

func (x *UnitAssignments) GetUnitId() uint64 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *UnitAssignments) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *UnitAssignments) GetUsers() []*UnitAssignment {
	if x != nil {
		return x.Users
	}
	return nil
}

type UnitAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitId uint64           `protobuf:"varint,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty" sql:"primary_key" alias:"unit_id"` // @gotags: sql:"primary_key" alias:"unit_id"
	UserId int32            `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" sql:"primary_key" alias:"user_id"` // @gotags: sql:"primary_key" alias:"user_id"
	User   *users.UserShort `protobuf:"bytes,3,opt,name=user,proto3,oneof" json:"user,omitempty"`
}

func (x *UnitAssignment) Reset() {
	*x = UnitAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_dispatch_units_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitAssignment) ProtoMessage() {}

func (x *UnitAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dispatch_units_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitAssignment.ProtoReflect.Descriptor instead.
func (*UnitAssignment) Descriptor() ([]byte, []int) {
	return file_resources_dispatch_units_proto_rawDescGZIP(), []int{2}
}

func (x *UnitAssignment) GetUnitId() uint64 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *UnitAssignment) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UnitAssignment) GetUser() *users.UserShort {
	if x != nil {
		return x.User
	}
	return nil
}

type UnitStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UnitId    uint64               `protobuf:"varint,3,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	Status    UNIT_STATUS          `protobuf:"varint,4,opt,name=status,proto3,enum=resources.dispatch.UNIT_STATUS" json:"status,omitempty"`
	Reason    *string              `protobuf:"bytes,5,opt,name=reason,proto3,oneof" json:"reason,omitempty"`
	Code      *string              `protobuf:"bytes,6,opt,name=code,proto3,oneof" json:"code,omitempty"`
	UserId    *int32               `protobuf:"varint,7,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	User      *users.UserShort     `protobuf:"bytes,8,opt,name=user,proto3,oneof" json:"user,omitempty"`
	X         *float64             `protobuf:"fixed64,9,opt,name=x,proto3,oneof" json:"x,omitempty"`
	Y         *float64             `protobuf:"fixed64,10,opt,name=y,proto3,oneof" json:"y,omitempty"`
	CreatorId *int32               `protobuf:"varint,11,opt,name=creator_id,json=creatorId,proto3,oneof" json:"creator_id,omitempty"`
	Creator   *users.UserShort     `protobuf:"bytes,12,opt,name=creator,proto3,oneof" json:"creator,omitempty"`
}

func (x *UnitStatus) Reset() {
	*x = UnitStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_dispatch_units_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitStatus) ProtoMessage() {}

func (x *UnitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dispatch_units_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitStatus.ProtoReflect.Descriptor instead.
func (*UnitStatus) Descriptor() ([]byte, []int) {
	return file_resources_dispatch_units_proto_rawDescGZIP(), []int{3}
}

func (x *UnitStatus) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UnitStatus) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UnitStatus) GetUnitId() uint64 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *UnitStatus) GetStatus() UNIT_STATUS {
	if x != nil {
		return x.Status
	}
	return UNIT_STATUS_UNKNOWN
}

func (x *UnitStatus) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *UnitStatus) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *UnitStatus) GetUserId() int32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *UnitStatus) GetUser() *users.UserShort {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UnitStatus) GetX() float64 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *UnitStatus) GetY() float64 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *UnitStatus) GetCreatorId() int32 {
	if x != nil && x.CreatorId != nil {
		return *x.CreatorId
	}
	return 0
}

func (x *UnitStatus) GetCreator() *users.UserShort {
	if x != nil {
		return x.Creator
	}
	return nil
}

var File_resources_dispatch_units_proto protoreflect.FileDescriptor

var file_resources_dispatch_units_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8e, 0x04, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x1d, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x03, 0x18, 0x18, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x06, 0x48, 0x02, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x18, 0xff, 0x01, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x55, 0x6e, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x7f, 0x0a, 0x0f, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03,
	0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x38, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x55, 0x6e, 0x69, 0x74,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x22, 0xce, 0x04, 0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x41, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01, 0x48, 0x01, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18,
	0x14, 0x48, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x04, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x78, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x05, 0x52, 0x01, 0x78, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x79,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x06, 0x52, 0x01, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2b,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x07, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x42, 0x04,
	0x0a, 0x02, 0x5f, 0x78, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x2a, 0x74, 0x0a, 0x0b, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x4e, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x59, 0x10, 0x06, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72,
	0x74, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x3b, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_dispatch_units_proto_rawDescOnce sync.Once
	file_resources_dispatch_units_proto_rawDescData = file_resources_dispatch_units_proto_rawDesc
)

func file_resources_dispatch_units_proto_rawDescGZIP() []byte {
	file_resources_dispatch_units_proto_rawDescOnce.Do(func() {
		file_resources_dispatch_units_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_dispatch_units_proto_rawDescData)
	})
	return file_resources_dispatch_units_proto_rawDescData
}

var file_resources_dispatch_units_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_dispatch_units_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_resources_dispatch_units_proto_goTypes = []interface{}{
	(UNIT_STATUS)(0),            // 0: resources.dispatch.UNIT_STATUS
	(*Unit)(nil),                // 1: resources.dispatch.Unit
	(*UnitAssignments)(nil),     // 2: resources.dispatch.UnitAssignments
	(*UnitAssignment)(nil),      // 3: resources.dispatch.UnitAssignment
	(*UnitStatus)(nil),          // 4: resources.dispatch.UnitStatus
	(*timestamp.Timestamp)(nil), // 5: resources.timestamp.Timestamp
	(*users.UserShort)(nil),     // 6: resources.users.UserShort
}
var file_resources_dispatch_units_proto_depIdxs = []int32{
	5,  // 0: resources.dispatch.Unit.created_at:type_name -> resources.timestamp.Timestamp
	5,  // 1: resources.dispatch.Unit.updated_at:type_name -> resources.timestamp.Timestamp
	4,  // 2: resources.dispatch.Unit.status:type_name -> resources.dispatch.UnitStatus
	3,  // 3: resources.dispatch.Unit.users:type_name -> resources.dispatch.UnitAssignment
	3,  // 4: resources.dispatch.UnitAssignments.users:type_name -> resources.dispatch.UnitAssignment
	6,  // 5: resources.dispatch.UnitAssignment.user:type_name -> resources.users.UserShort
	5,  // 6: resources.dispatch.UnitStatus.created_at:type_name -> resources.timestamp.Timestamp
	0,  // 7: resources.dispatch.UnitStatus.status:type_name -> resources.dispatch.UNIT_STATUS
	6,  // 8: resources.dispatch.UnitStatus.user:type_name -> resources.users.UserShort
	6,  // 9: resources.dispatch.UnitStatus.creator:type_name -> resources.users.UserShort
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_resources_dispatch_units_proto_init() }
func file_resources_dispatch_units_proto_init() {
	if File_resources_dispatch_units_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_dispatch_units_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
		file_resources_dispatch_units_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitAssignments); i {
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
		file_resources_dispatch_units_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitAssignment); i {
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
		file_resources_dispatch_units_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitStatus); i {
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
	file_resources_dispatch_units_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_dispatch_units_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_resources_dispatch_units_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_dispatch_units_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_dispatch_units_proto_goTypes,
		DependencyIndexes: file_resources_dispatch_units_proto_depIdxs,
		EnumInfos:         file_resources_dispatch_units_proto_enumTypes,
		MessageInfos:      file_resources_dispatch_units_proto_msgTypes,
	}.Build()
	File_resources_dispatch_units_proto = out.File
	file_resources_dispatch_units_proto_rawDesc = nil
	file_resources_dispatch_units_proto_goTypes = nil
	file_resources_dispatch_units_proto_depIdxs = nil
}
