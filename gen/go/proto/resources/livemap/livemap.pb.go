// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: resources/livemap/livemap.proto

package livemap

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	dispatch "github.com/galexrt/fivenet/gen/go/proto/resources/dispatch"
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

type MARKER_TYPE int32

const (
	MARKER_TYPE_GENERIC  MARKER_TYPE = 0
	MARKER_TYPE_USER     MARKER_TYPE = 1
	MARKER_TYPE_DISPATCH MARKER_TYPE = 2
	MARKER_TYPE_CIRCLE   MARKER_TYPE = 3
)

// Enum value maps for MARKER_TYPE.
var (
	MARKER_TYPE_name = map[int32]string{
		0: "GENERIC",
		1: "USER",
		2: "DISPATCH",
		3: "CIRCLE",
	}
	MARKER_TYPE_value = map[string]int32{
		"GENERIC":  0,
		"USER":     1,
		"DISPATCH": 2,
		"CIRCLE":   3,
	}
)

func (x MARKER_TYPE) Enum() *MARKER_TYPE {
	p := new(MARKER_TYPE)
	*p = x
	return p
}

func (x MARKER_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MARKER_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_livemap_livemap_proto_enumTypes[0].Descriptor()
}

func (MARKER_TYPE) Type() protoreflect.EnumType {
	return &file_resources_livemap_livemap_proto_enumTypes[0]
}

func (x MARKER_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MARKER_TYPE.Descriptor instead.
func (MARKER_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{0}
}

type GenericMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	X         float64              `protobuf:"fixed64,3,opt,name=x,proto3" json:"x,omitempty"`
	Y         float64              `protobuf:"fixed64,4,opt,name=y,proto3" json:"y,omitempty"`
	Name      string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string               `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	IconColor string               `protobuf:"bytes,7,opt,name=icon_color,json=iconColor,proto3" json:"icon_color,omitempty"`
	Popup     string               `protobuf:"bytes,8,opt,name=popup,proto3" json:"popup,omitempty"`
}

func (x *GenericMarker) Reset() {
	*x = GenericMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericMarker) ProtoMessage() {}

func (x *GenericMarker) ProtoReflect() protoreflect.Message {
	mi := &file_resources_livemap_livemap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericMarker.ProtoReflect.Descriptor instead.
func (*GenericMarker) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{0}
}

func (x *GenericMarker) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GenericMarker) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *GenericMarker) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *GenericMarker) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *GenericMarker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenericMarker) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *GenericMarker) GetIconColor() string {
	if x != nil {
		return x.IconColor
	}
	return ""
}

func (x *GenericMarker) GetPopup() string {
	if x != nil {
		return x.Popup
	}
	return ""
}

type DispatchMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Marker   *GenericMarker `protobuf:"bytes,1,opt,name=marker,proto3" json:"marker,omitempty"`
	Job      string         `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty" alias:"job"`                                 // @gotags: alias:"job"
	JobLabel *string        `protobuf:"bytes,3,opt,name=job_label,json=jobLabel,proto3,oneof" json:"job_label,omitempty" alias:"job_label"` // @gotags: alias:"job_label"
	Active   bool           `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *DispatchMarker) Reset() {
	*x = DispatchMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchMarker) ProtoMessage() {}

func (x *DispatchMarker) ProtoReflect() protoreflect.Message {
	mi := &file_resources_livemap_livemap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchMarker.ProtoReflect.Descriptor instead.
func (*DispatchMarker) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{1}
}

func (x *DispatchMarker) GetMarker() *GenericMarker {
	if x != nil {
		return x.Marker
	}
	return nil
}

func (x *DispatchMarker) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *DispatchMarker) GetJobLabel() string {
	if x != nil && x.JobLabel != nil {
		return *x.JobLabel
	}
	return ""
}

func (x *DispatchMarker) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type UserMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Marker *GenericMarker   `protobuf:"bytes,1,opt,name=marker,proto3" json:"marker,omitempty"`
	User   *users.UserShort `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" alias:"user"` // @gotags: alias:"user"
	UnitId *uint64          `protobuf:"varint,3,opt,name=unit_id,json=unitId,proto3,oneof" json:"unit_id,omitempty"`
	Unit   *dispatch.Unit   `protobuf:"bytes,4,opt,name=unit,proto3,oneof" json:"unit,omitempty"`
}

func (x *UserMarker) Reset() {
	*x = UserMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMarker) ProtoMessage() {}

func (x *UserMarker) ProtoReflect() protoreflect.Message {
	mi := &file_resources_livemap_livemap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMarker.ProtoReflect.Descriptor instead.
func (*UserMarker) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{2}
}

func (x *UserMarker) GetMarker() *GenericMarker {
	if x != nil {
		return x.Marker
	}
	return nil
}

func (x *UserMarker) GetUser() *users.UserShort {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserMarker) GetUnitId() uint64 {
	if x != nil && x.UnitId != nil {
		return *x.UnitId
	}
	return 0
}

func (x *UserMarker) GetUnit() *dispatch.Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

var File_resources_livemap_livemap_proto protoreflect.FileDescriptor

var file_resources_livemap_livemap_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65,
	0x6d, 0x61, 0x70, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6d, 0x61, 0x70, 0x1a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe0, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x6f, 0x70, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x70,
	0x75, 0x70, 0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x29, 0x0a, 0x09, 0x6a, 0x6f,
	0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xdc, 0x01, 0x0a, 0x0a,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x48, 0x01, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x2a, 0x3e, 0x0a, 0x0b, 0x4d, 0x41,
	0x52, 0x4b, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x45, 0x4e,
	0x45, 0x52, 0x49, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x49, 0x52, 0x43, 0x4c, 0x45, 0x10, 0x03, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x74,
	0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x3b, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_livemap_livemap_proto_rawDescOnce sync.Once
	file_resources_livemap_livemap_proto_rawDescData = file_resources_livemap_livemap_proto_rawDesc
)

func file_resources_livemap_livemap_proto_rawDescGZIP() []byte {
	file_resources_livemap_livemap_proto_rawDescOnce.Do(func() {
		file_resources_livemap_livemap_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_livemap_livemap_proto_rawDescData)
	})
	return file_resources_livemap_livemap_proto_rawDescData
}

var file_resources_livemap_livemap_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_livemap_livemap_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resources_livemap_livemap_proto_goTypes = []interface{}{
	(MARKER_TYPE)(0),            // 0: resources.livemap.MARKER_TYPE
	(*GenericMarker)(nil),       // 1: resources.livemap.GenericMarker
	(*DispatchMarker)(nil),      // 2: resources.livemap.DispatchMarker
	(*UserMarker)(nil),          // 3: resources.livemap.UserMarker
	(*timestamp.Timestamp)(nil), // 4: resources.timestamp.Timestamp
	(*users.UserShort)(nil),     // 5: resources.users.UserShort
	(*dispatch.Unit)(nil),       // 6: resources.dispatch.Unit
}
var file_resources_livemap_livemap_proto_depIdxs = []int32{
	4, // 0: resources.livemap.GenericMarker.updated_at:type_name -> resources.timestamp.Timestamp
	1, // 1: resources.livemap.DispatchMarker.marker:type_name -> resources.livemap.GenericMarker
	1, // 2: resources.livemap.UserMarker.marker:type_name -> resources.livemap.GenericMarker
	5, // 3: resources.livemap.UserMarker.user:type_name -> resources.users.UserShort
	6, // 4: resources.livemap.UserMarker.unit:type_name -> resources.dispatch.Unit
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resources_livemap_livemap_proto_init() }
func file_resources_livemap_livemap_proto_init() {
	if File_resources_livemap_livemap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_livemap_livemap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericMarker); i {
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
		file_resources_livemap_livemap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchMarker); i {
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
		file_resources_livemap_livemap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMarker); i {
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
	file_resources_livemap_livemap_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_livemap_livemap_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_livemap_livemap_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_livemap_livemap_proto_goTypes,
		DependencyIndexes: file_resources_livemap_livemap_proto_depIdxs,
		EnumInfos:         file_resources_livemap_livemap_proto_enumTypes,
		MessageInfos:      file_resources_livemap_livemap_proto_msgTypes,
	}.Build()
	File_resources_livemap_livemap_proto = out.File
	file_resources_livemap_livemap_proto_rawDesc = nil
	file_resources_livemap_livemap_proto_goTypes = nil
	file_resources_livemap_livemap_proto_depIdxs = nil
}
