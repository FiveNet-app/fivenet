// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: resources/dispatch/dispatch.proto

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

type SQUAD_STATUS int32

const (
	SQUAD_STATUS_UNKNOWN     SQUAD_STATUS = 0
	SQUAD_STATUS_AVAILABLE   SQUAD_STATUS = 1
	SQUAD_STATUS_UNAVAILABLE SQUAD_STATUS = 2
	SQUAD_STATUS_PAUSE       SQUAD_STATUS = 3
)

// Enum value maps for SQUAD_STATUS.
var (
	SQUAD_STATUS_name = map[int32]string{
		0: "UNKNOWN",
		1: "AVAILABLE",
		2: "UNAVAILABLE",
		3: "PAUSE",
	}
	SQUAD_STATUS_value = map[string]int32{
		"UNKNOWN":     0,
		"AVAILABLE":   1,
		"UNAVAILABLE": 2,
		"PAUSE":       3,
	}
)

func (x SQUAD_STATUS) Enum() *SQUAD_STATUS {
	p := new(SQUAD_STATUS)
	*p = x
	return p
}

func (x SQUAD_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SQUAD_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_dispatch_dispatch_proto_enumTypes[0].Descriptor()
}

func (SQUAD_STATUS) Type() protoreflect.EnumType {
	return &file_resources_dispatch_dispatch_proto_enumTypes[0]
}

func (x SQUAD_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SQUAD_STATUS.Descriptor instead.
func (SQUAD_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_resources_dispatch_dispatch_proto_rawDescGZIP(), []int{0}
}

type DISPATCH_STATUS int32

const (
	DISPATCH_STATUS_UNASSIGNED  DISPATCH_STATUS = 0
	DISPATCH_STATUS_EN_ROUTE    DISPATCH_STATUS = 1
	DISPATCH_STATUS_AT_SCENE    DISPATCH_STATUS = 2
	DISPATCH_STATUS_NEED_BACKUP DISPATCH_STATUS = 3
)

// Enum value maps for DISPATCH_STATUS.
var (
	DISPATCH_STATUS_name = map[int32]string{
		0: "UNASSIGNED",
		1: "EN_ROUTE",
		2: "AT_SCENE",
		3: "NEED_BACKUP",
	}
	DISPATCH_STATUS_value = map[string]int32{
		"UNASSIGNED":  0,
		"EN_ROUTE":    1,
		"AT_SCENE":    2,
		"NEED_BACKUP": 3,
	}
)

func (x DISPATCH_STATUS) Enum() *DISPATCH_STATUS {
	p := new(DISPATCH_STATUS)
	*p = x
	return p
}

func (x DISPATCH_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DISPATCH_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_dispatch_dispatch_proto_enumTypes[1].Descriptor()
}

func (DISPATCH_STATUS) Type() protoreflect.EnumType {
	return &file_resources_dispatch_dispatch_proto_enumTypes[1]
}

func (x DISPATCH_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DISPATCH_STATUS.Descriptor instead.
func (DISPATCH_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_resources_dispatch_dispatch_proto_rawDescGZIP(), []int{1}
}

type Squad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	Name     string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Initials string             `protobuf:"bytes,3,opt,name=initials,proto3" json:"initials,omitempty"`
	Status   *SQUAD_STATUS      `protobuf:"varint,4,opt,name=status,proto3,enum=resources.dispatch.SQUAD_STATUS,oneof" json:"status,omitempty"`
	Assigned []*users.UserShort `protobuf:"bytes,5,rep,name=assigned,proto3" json:"assigned,omitempty"`
}

func (x *Squad) Reset() {
	*x = Squad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_dispatch_dispatch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Squad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Squad) ProtoMessage() {}

func (x *Squad) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dispatch_dispatch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Squad.ProtoReflect.Descriptor instead.
func (*Squad) Descriptor() ([]byte, []int) {
	return file_resources_dispatch_dispatch_proto_rawDescGZIP(), []int{0}
}

func (x *Squad) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Squad) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Squad) GetInitials() string {
	if x != nil {
		return x.Initials
	}
	return ""
}

func (x *Squad) GetStatus() SQUAD_STATUS {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return SQUAD_STATUS_UNKNOWN
}

func (x *Squad) GetAssigned() []*users.UserShort {
	if x != nil {
		return x.Assigned
	}
	return nil
}

type Dispatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"`                                     // @gotags: sql:"primary_key" alias:"id"
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty" alias:"created_at"` // @gotags: alias:"created_at"
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty" alias:"updated_at"` // @gotags: alias:"updated_at"
	Status      *DISPATCH_STATUS     `protobuf:"varint,4,opt,name=status,proto3,enum=resources.dispatch.DISPATCH_STATUS,oneof" json:"status,omitempty"`
	Title       string               `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description *string              `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Attributes  map[string]string    `protobuf:"bytes,7,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Dispatch) Reset() {
	*x = Dispatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_dispatch_dispatch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dispatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dispatch) ProtoMessage() {}

func (x *Dispatch) ProtoReflect() protoreflect.Message {
	mi := &file_resources_dispatch_dispatch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dispatch.ProtoReflect.Descriptor instead.
func (*Dispatch) Descriptor() ([]byte, []int) {
	return file_resources_dispatch_dispatch_proto_rawDescGZIP(), []int{1}
}

func (x *Dispatch) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dispatch) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Dispatch) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Dispatch) GetStatus() DISPATCH_STATUS {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return DISPATCH_STATUS_UNASSIGNED
}

func (x *Dispatch) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Dispatch) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Dispatch) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_resources_dispatch_dispatch_proto protoreflect.FileDescriptor

var file_resources_dispatch_dispatch_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x05, 0x53, 0x71, 0x75, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x18, 0x14, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x18, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x3d,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x53, 0x51, 0x55, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a,
	0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xe7, 0x03, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x42, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x44, 0x49, 0x53, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x4c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x46, 0x0a, 0x0c, 0x53, 0x51,
	0x55, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49,
	0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x55, 0x53, 0x45,
	0x10, 0x03, 0x2a, 0x4e, 0x0a, 0x0f, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x41, 0x53, 0x53, 0x49, 0x47,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x5f, 0x52, 0x4f, 0x55, 0x54,
	0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x54, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x45, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50,
	0x10, 0x03, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x3b, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_resources_dispatch_dispatch_proto_rawDescOnce sync.Once
	file_resources_dispatch_dispatch_proto_rawDescData = file_resources_dispatch_dispatch_proto_rawDesc
)

func file_resources_dispatch_dispatch_proto_rawDescGZIP() []byte {
	file_resources_dispatch_dispatch_proto_rawDescOnce.Do(func() {
		file_resources_dispatch_dispatch_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_dispatch_dispatch_proto_rawDescData)
	})
	return file_resources_dispatch_dispatch_proto_rawDescData
}

var file_resources_dispatch_dispatch_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_resources_dispatch_dispatch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_dispatch_dispatch_proto_goTypes = []interface{}{
	(SQUAD_STATUS)(0),           // 0: resources.dispatch.SQUAD_STATUS
	(DISPATCH_STATUS)(0),        // 1: resources.dispatch.DISPATCH_STATUS
	(*Squad)(nil),               // 2: resources.dispatch.Squad
	(*Dispatch)(nil),            // 3: resources.dispatch.Dispatch
	nil,                         // 4: resources.dispatch.Dispatch.AttributesEntry
	(*users.UserShort)(nil),     // 5: resources.users.UserShort
	(*timestamp.Timestamp)(nil), // 6: resources.timestamp.Timestamp
}
var file_resources_dispatch_dispatch_proto_depIdxs = []int32{
	0, // 0: resources.dispatch.Squad.status:type_name -> resources.dispatch.SQUAD_STATUS
	5, // 1: resources.dispatch.Squad.assigned:type_name -> resources.users.UserShort
	6, // 2: resources.dispatch.Dispatch.created_at:type_name -> resources.timestamp.Timestamp
	6, // 3: resources.dispatch.Dispatch.updated_at:type_name -> resources.timestamp.Timestamp
	1, // 4: resources.dispatch.Dispatch.status:type_name -> resources.dispatch.DISPATCH_STATUS
	4, // 5: resources.dispatch.Dispatch.attributes:type_name -> resources.dispatch.Dispatch.AttributesEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_resources_dispatch_dispatch_proto_init() }
func file_resources_dispatch_dispatch_proto_init() {
	if File_resources_dispatch_dispatch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_dispatch_dispatch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Squad); i {
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
		file_resources_dispatch_dispatch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dispatch); i {
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
	file_resources_dispatch_dispatch_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_dispatch_dispatch_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_dispatch_dispatch_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_dispatch_dispatch_proto_goTypes,
		DependencyIndexes: file_resources_dispatch_dispatch_proto_depIdxs,
		EnumInfos:         file_resources_dispatch_dispatch_proto_enumTypes,
		MessageInfos:      file_resources_dispatch_dispatch_proto_msgTypes,
	}.Build()
	File_resources_dispatch_dispatch_proto = out.File
	file_resources_dispatch_dispatch_proto_rawDesc = nil
	file_resources_dispatch_dispatch_proto_goTypes = nil
	file_resources_dispatch_dispatch_proto_depIdxs = nil
}
