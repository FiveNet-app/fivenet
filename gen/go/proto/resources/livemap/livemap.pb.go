// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.20.3
// source: resources/livemap/livemap.proto

package livemap

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	centrum "github.com/fivenet-app/fivenet/gen/go/proto/resources/centrum"
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

type MarkerType int32

const (
	MarkerType_MARKER_TYPE_UNSPECIFIED MarkerType = 0
	MarkerType_MARKER_TYPE_DOT         MarkerType = 1
	MarkerType_MARKER_TYPE_CIRCLE      MarkerType = 2
	MarkerType_MARKER_TYPE_ICON        MarkerType = 3
)

// Enum value maps for MarkerType.
var (
	MarkerType_name = map[int32]string{
		0: "MARKER_TYPE_UNSPECIFIED",
		1: "MARKER_TYPE_DOT",
		2: "MARKER_TYPE_CIRCLE",
		3: "MARKER_TYPE_ICON",
	}
	MarkerType_value = map[string]int32{
		"MARKER_TYPE_UNSPECIFIED": 0,
		"MARKER_TYPE_DOT":         1,
		"MARKER_TYPE_CIRCLE":      2,
		"MARKER_TYPE_ICON":        3,
	}
)

func (x MarkerType) Enum() *MarkerType {
	p := new(MarkerType)
	*p = x
	return p
}

func (x MarkerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MarkerType) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_livemap_livemap_proto_enumTypes[0].Descriptor()
}

func (MarkerType) Type() protoreflect.EnumType {
	return &file_resources_livemap_livemap_proto_enumTypes[0]
}

func (x MarkerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MarkerType.Descriptor instead.
func (MarkerType) EnumDescriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{0}
}

type MarkerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	Job       string               `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	JobLabel  string               `protobuf:"bytes,12,opt,name=job_label,json=jobLabel,proto3" json:"job_label,omitempty"`
	// @sanitize
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// @sanitize
	Description *string `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	X           float64 `protobuf:"fixed64,7,opt,name=x,proto3" json:"x,omitempty"`
	Y           float64 `protobuf:"fixed64,8,opt,name=y,proto3" json:"y,omitempty"`
	// @sanitize
	Postal *string `protobuf:"bytes,9,opt,name=postal,proto3,oneof" json:"postal,omitempty"`
	Color  *string `protobuf:"bytes,10,opt,name=color,proto3,oneof" json:"color,omitempty"`
	Icon   *string `protobuf:"bytes,11,opt,name=icon,proto3,oneof" json:"icon,omitempty"`
}

func (x *MarkerInfo) Reset() {
	*x = MarkerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkerInfo) ProtoMessage() {}

func (x *MarkerInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MarkerInfo.ProtoReflect.Descriptor instead.
func (*MarkerInfo) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{0}
}

func (x *MarkerInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MarkerInfo) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MarkerInfo) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MarkerInfo) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *MarkerInfo) GetJobLabel() string {
	if x != nil {
		return x.JobLabel
	}
	return ""
}

func (x *MarkerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MarkerInfo) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *MarkerInfo) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MarkerInfo) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *MarkerInfo) GetPostal() string {
	if x != nil && x.Postal != nil {
		return *x.Postal
	}
	return ""
}

func (x *MarkerInfo) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *MarkerInfo) GetIcon() string {
	if x != nil && x.Icon != nil {
		return *x.Icon
	}
	return ""
}

type UserMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info   *MarkerInfo      `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	UserId int32            `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User   *users.UserShort `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty" alias:"user"` // @gotags: alias:"user"
	UnitId *uint64          `protobuf:"varint,4,opt,name=unit_id,json=unitId,proto3,oneof" json:"unit_id,omitempty"`
	Unit   *centrum.Unit    `protobuf:"bytes,5,opt,name=unit,proto3,oneof" json:"unit,omitempty"`
}

func (x *UserMarker) Reset() {
	*x = UserMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMarker) ProtoMessage() {}

func (x *UserMarker) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserMarker.ProtoReflect.Descriptor instead.
func (*UserMarker) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{1}
}

func (x *UserMarker) GetInfo() *MarkerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *UserMarker) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
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

func (x *UserMarker) GetUnit() *centrum.Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

type MarkerMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *MarkerInfo          `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Type      MarkerType           `protobuf:"varint,2,opt,name=type,proto3,enum=resources.livemap.MarkerType" json:"type,omitempty" alias:"markerType"` // @gotags: alias:"markerType"
	ExpiresAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3,oneof" json:"expires_at,omitempty"`
	Data      *MarkerData          `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty" alias:"markerData"` // @gotags: alias:"markerData"
	CreatorId *int32               `protobuf:"varint,5,opt,name=creator_id,json=creatorId,proto3,oneof" json:"creator_id,omitempty"`
	Creator   *users.UserShort     `protobuf:"bytes,6,opt,name=creator,proto3,oneof" json:"creator,omitempty"`
}

func (x *MarkerMarker) Reset() {
	*x = MarkerMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkerMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkerMarker) ProtoMessage() {}

func (x *MarkerMarker) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MarkerMarker.ProtoReflect.Descriptor instead.
func (*MarkerMarker) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{2}
}

func (x *MarkerMarker) GetInfo() *MarkerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *MarkerMarker) GetType() MarkerType {
	if x != nil {
		return x.Type
	}
	return MarkerType_MARKER_TYPE_UNSPECIFIED
}

func (x *MarkerMarker) GetExpiresAt() *timestamp.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *MarkerMarker) GetData() *MarkerData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MarkerMarker) GetCreatorId() int32 {
	if x != nil && x.CreatorId != nil {
		return *x.CreatorId
	}
	return 0
}

func (x *MarkerMarker) GetCreator() *users.UserShort {
	if x != nil {
		return x.Creator
	}
	return nil
}

type MarkerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*MarkerData_Circle
	//	*MarkerData_Icon
	Data isMarkerData_Data `protobuf_oneof:"data"`
}

func (x *MarkerData) Reset() {
	*x = MarkerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkerData) ProtoMessage() {}

func (x *MarkerData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_livemap_livemap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkerData.ProtoReflect.Descriptor instead.
func (*MarkerData) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{3}
}

func (m *MarkerData) GetData() isMarkerData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *MarkerData) GetCircle() *CircleMarker {
	if x, ok := x.GetData().(*MarkerData_Circle); ok {
		return x.Circle
	}
	return nil
}

func (x *MarkerData) GetIcon() *IconMarker {
	if x, ok := x.GetData().(*MarkerData_Icon); ok {
		return x.Icon
	}
	return nil
}

type isMarkerData_Data interface {
	isMarkerData_Data()
}

type MarkerData_Circle struct {
	Circle *CircleMarker `protobuf:"bytes,3,opt,name=circle,proto3,oneof"`
}

type MarkerData_Icon struct {
	Icon *IconMarker `protobuf:"bytes,4,opt,name=icon,proto3,oneof"`
}

func (*MarkerData_Circle) isMarkerData_Data() {}

func (*MarkerData_Icon) isMarkerData_Data() {}

type CircleMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Radius  int32    `protobuf:"varint,1,opt,name=radius,proto3" json:"radius,omitempty"`
	Opacity *float32 `protobuf:"fixed32,2,opt,name=opacity,proto3,oneof" json:"opacity,omitempty"`
}

func (x *CircleMarker) Reset() {
	*x = CircleMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CircleMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CircleMarker) ProtoMessage() {}

func (x *CircleMarker) ProtoReflect() protoreflect.Message {
	mi := &file_resources_livemap_livemap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CircleMarker.ProtoReflect.Descriptor instead.
func (*CircleMarker) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{4}
}

func (x *CircleMarker) GetRadius() int32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *CircleMarker) GetOpacity() float32 {
	if x != nil && x.Opacity != nil {
		return *x.Opacity
	}
	return 0
}

type Coords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coords) Reset() {
	*x = Coords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coords) ProtoMessage() {}

func (x *Coords) ProtoReflect() protoreflect.Message {
	mi := &file_resources_livemap_livemap_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coords.ProtoReflect.Descriptor instead.
func (*Coords) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{5}
}

func (x *Coords) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coords) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type IconMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *IconMarker) Reset() {
	*x = IconMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_livemap_livemap_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IconMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IconMarker) ProtoMessage() {}

func (x *IconMarker) ProtoReflect() protoreflect.Message {
	mi := &file_resources_livemap_livemap_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IconMarker.ProtoReflect.Descriptor instead.
func (*IconMarker) Descriptor() ([]byte, []int) {
	return file_resources_livemap_livemap_proto_rawDescGZIP(), []int{6}
}

func (x *IconMarker) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

var File_resources_livemap_livemap_proto protoreflect.FileDescriptor

var file_resources_livemap_livemap_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65,
	0x6d, 0x61, 0x70, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6d, 0x61, 0x70, 0x1a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1,
	0x03, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x6a,
	0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12,
	0x24, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x30, 0x48, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x32, 0x11, 0x5e, 0x23, 0x5b,
	0x41, 0x2d, 0x46, 0x61, 0x2d, 0x66, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36, 0x7d, 0x24, 0x98, 0x01,
	0x07, 0x48, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x63,
	0x6f, 0x6e, 0x22, 0xfa, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x06, 0x75,
	0x6e, 0x69, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x75, 0x6d, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x48,
	0x01, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x22,
	0xf4, 0x02, 0x0a, 0x0c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d,
	0x61, 0x70, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x48, 0x02,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x12, 0x33, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6d,
	0x61, 0x70, 0x2e, 0x49, 0x63, 0x6f, 0x6e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a,
	0x0c, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6f, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x22, 0x24, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x20, 0x0a, 0x0a, 0x49, 0x63, 0x6f, 0x6e, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x2a, 0x6c, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x4f, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x52, 0x4b,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x49, 0x52, 0x43, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x43, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70,
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
var file_resources_livemap_livemap_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_resources_livemap_livemap_proto_goTypes = []interface{}{
	(MarkerType)(0),             // 0: resources.livemap.MarkerType
	(*MarkerInfo)(nil),          // 1: resources.livemap.MarkerInfo
	(*UserMarker)(nil),          // 2: resources.livemap.UserMarker
	(*MarkerMarker)(nil),        // 3: resources.livemap.MarkerMarker
	(*MarkerData)(nil),          // 4: resources.livemap.MarkerData
	(*CircleMarker)(nil),        // 5: resources.livemap.CircleMarker
	(*Coords)(nil),              // 6: resources.livemap.Coords
	(*IconMarker)(nil),          // 7: resources.livemap.IconMarker
	(*timestamp.Timestamp)(nil), // 8: resources.timestamp.Timestamp
	(*users.UserShort)(nil),     // 9: resources.users.UserShort
	(*centrum.Unit)(nil),        // 10: resources.centrum.Unit
}
var file_resources_livemap_livemap_proto_depIdxs = []int32{
	8,  // 0: resources.livemap.MarkerInfo.created_at:type_name -> resources.timestamp.Timestamp
	8,  // 1: resources.livemap.MarkerInfo.updated_at:type_name -> resources.timestamp.Timestamp
	1,  // 2: resources.livemap.UserMarker.info:type_name -> resources.livemap.MarkerInfo
	9,  // 3: resources.livemap.UserMarker.user:type_name -> resources.users.UserShort
	10, // 4: resources.livemap.UserMarker.unit:type_name -> resources.centrum.Unit
	1,  // 5: resources.livemap.MarkerMarker.info:type_name -> resources.livemap.MarkerInfo
	0,  // 6: resources.livemap.MarkerMarker.type:type_name -> resources.livemap.MarkerType
	8,  // 7: resources.livemap.MarkerMarker.expires_at:type_name -> resources.timestamp.Timestamp
	4,  // 8: resources.livemap.MarkerMarker.data:type_name -> resources.livemap.MarkerData
	9,  // 9: resources.livemap.MarkerMarker.creator:type_name -> resources.users.UserShort
	5,  // 10: resources.livemap.MarkerData.circle:type_name -> resources.livemap.CircleMarker
	7,  // 11: resources.livemap.MarkerData.icon:type_name -> resources.livemap.IconMarker
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_resources_livemap_livemap_proto_init() }
func file_resources_livemap_livemap_proto_init() {
	if File_resources_livemap_livemap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_livemap_livemap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkerInfo); i {
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
		file_resources_livemap_livemap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkerMarker); i {
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
		file_resources_livemap_livemap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkerData); i {
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
		file_resources_livemap_livemap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CircleMarker); i {
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
		file_resources_livemap_livemap_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coords); i {
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
		file_resources_livemap_livemap_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IconMarker); i {
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
	file_resources_livemap_livemap_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_resources_livemap_livemap_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_resources_livemap_livemap_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_resources_livemap_livemap_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MarkerData_Circle)(nil),
		(*MarkerData_Icon)(nil),
	}
	file_resources_livemap_livemap_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_livemap_livemap_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
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
