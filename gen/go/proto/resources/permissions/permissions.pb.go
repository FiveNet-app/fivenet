// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: resources/permissions/permissions.proto

package permissions

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/fivenet-app/fivenet/gen/go/proto/resources/timestamp"
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

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	Category  string               `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Name      string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	GuardName string               `protobuf:"bytes,5,opt,name=guard_name,json=guardName,proto3" json:"guard_name,omitempty"`
	Val       bool                 `protobuf:"varint,6,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_resources_permissions_permissions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_resources_permissions_permissions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_resources_permissions_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Permission) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Permission) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetGuardName() string {
	if x != nil {
		return x.GuardName
	}
	return ""
}

func (x *Permission) GetVal() bool {
	if x != nil {
		return x.Val
	}
	return false
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	Job           string               `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
	JobLabel      *string              `protobuf:"bytes,4,opt,name=job_label,json=jobLabel,proto3,oneof" json:"job_label,omitempty"`
	Grade         int32                `protobuf:"varint,5,opt,name=grade,proto3" json:"grade,omitempty"`
	JobGradeLabel *string              `protobuf:"bytes,6,opt,name=job_grade_label,json=jobGradeLabel,proto3,oneof" json:"job_grade_label,omitempty"`
	Permissions   []*Permission        `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Attributes    []*RoleAttribute     `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_resources_permissions_permissions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_resources_permissions_permissions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_resources_permissions_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Role) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Role) GetJobLabel() string {
	if x != nil && x.JobLabel != nil {
		return *x.JobLabel
	}
	return ""
}

func (x *Role) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Role) GetJobGradeLabel() string {
	if x != nil && x.JobGradeLabel != nil {
		return *x.JobGradeLabel
	}
	return ""
}

func (x *Role) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Role) GetAttributes() []*RoleAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type RawRoleAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId       uint64               `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	CreatedAt    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	AttrId       uint64               `protobuf:"varint,3,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
	PermissionId uint64               `protobuf:"varint,4,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	Category     string               `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Name         string               `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Key          string               `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	Type         string               `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	ValidValues  *AttributeValues     `protobuf:"bytes,9,opt,name=valid_values,json=validValues,proto3" json:"valid_values,omitempty"`
	Value        *AttributeValues     `protobuf:"bytes,10,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RawRoleAttribute) Reset() {
	*x = RawRoleAttribute{}
	mi := &file_resources_permissions_permissions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RawRoleAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawRoleAttribute) ProtoMessage() {}

func (x *RawRoleAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_resources_permissions_permissions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawRoleAttribute.ProtoReflect.Descriptor instead.
func (*RawRoleAttribute) Descriptor() ([]byte, []int) {
	return file_resources_permissions_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *RawRoleAttribute) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *RawRoleAttribute) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RawRoleAttribute) GetAttrId() uint64 {
	if x != nil {
		return x.AttrId
	}
	return 0
}

func (x *RawRoleAttribute) GetPermissionId() uint64 {
	if x != nil {
		return x.PermissionId
	}
	return 0
}

func (x *RawRoleAttribute) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *RawRoleAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RawRoleAttribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RawRoleAttribute) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RawRoleAttribute) GetValidValues() *AttributeValues {
	if x != nil {
		return x.ValidValues
	}
	return nil
}

func (x *RawRoleAttribute) GetValue() *AttributeValues {
	if x != nil {
		return x.Value
	}
	return nil
}

type RoleAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId       uint64               `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	CreatedAt    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	AttrId       uint64               `protobuf:"varint,3,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
	PermissionId uint64               `protobuf:"varint,4,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	Category     string               `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Name         string               `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Key          string               `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	Type         string               `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	ValidValues  *AttributeValues     `protobuf:"bytes,9,opt,name=valid_values,json=validValues,proto3" json:"valid_values,omitempty"`
	Value        *AttributeValues     `protobuf:"bytes,10,opt,name=value,proto3" json:"value,omitempty"`
	MaxValues    *AttributeValues     `protobuf:"bytes,11,opt,name=max_values,json=maxValues,proto3,oneof" json:"max_values,omitempty"`
}

func (x *RoleAttribute) Reset() {
	*x = RoleAttribute{}
	mi := &file_resources_permissions_permissions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAttribute) ProtoMessage() {}

func (x *RoleAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_resources_permissions_permissions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleAttribute.ProtoReflect.Descriptor instead.
func (*RoleAttribute) Descriptor() ([]byte, []int) {
	return file_resources_permissions_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *RoleAttribute) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *RoleAttribute) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RoleAttribute) GetAttrId() uint64 {
	if x != nil {
		return x.AttrId
	}
	return 0
}

func (x *RoleAttribute) GetPermissionId() uint64 {
	if x != nil {
		return x.PermissionId
	}
	return 0
}

func (x *RoleAttribute) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *RoleAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleAttribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RoleAttribute) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RoleAttribute) GetValidValues() *AttributeValues {
	if x != nil {
		return x.ValidValues
	}
	return nil
}

func (x *RoleAttribute) GetValue() *AttributeValues {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *RoleAttribute) GetMaxValues() *AttributeValues {
	if x != nil {
		return x.MaxValues
	}
	return nil
}

type AttributeValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ValidValues:
	//
	//	*AttributeValues_StringList
	//	*AttributeValues_JobList
	//	*AttributeValues_JobGradeList
	ValidValues isAttributeValues_ValidValues `protobuf_oneof:"valid_values"`
}

func (x *AttributeValues) Reset() {
	*x = AttributeValues{}
	mi := &file_resources_permissions_permissions_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttributeValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeValues) ProtoMessage() {}

func (x *AttributeValues) ProtoReflect() protoreflect.Message {
	mi := &file_resources_permissions_permissions_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeValues.ProtoReflect.Descriptor instead.
func (*AttributeValues) Descriptor() ([]byte, []int) {
	return file_resources_permissions_permissions_proto_rawDescGZIP(), []int{4}
}

func (m *AttributeValues) GetValidValues() isAttributeValues_ValidValues {
	if m != nil {
		return m.ValidValues
	}
	return nil
}

func (x *AttributeValues) GetStringList() *StringList {
	if x, ok := x.GetValidValues().(*AttributeValues_StringList); ok {
		return x.StringList
	}
	return nil
}

func (x *AttributeValues) GetJobList() *StringList {
	if x, ok := x.GetValidValues().(*AttributeValues_JobList); ok {
		return x.JobList
	}
	return nil
}

func (x *AttributeValues) GetJobGradeList() *JobGradeList {
	if x, ok := x.GetValidValues().(*AttributeValues_JobGradeList); ok {
		return x.JobGradeList
	}
	return nil
}

type isAttributeValues_ValidValues interface {
	isAttributeValues_ValidValues()
}

type AttributeValues_StringList struct {
	StringList *StringList `protobuf:"bytes,1,opt,name=string_list,json=stringList,proto3,oneof"`
}

type AttributeValues_JobList struct {
	JobList *StringList `protobuf:"bytes,2,opt,name=job_list,json=jobList,proto3,oneof"`
}

type AttributeValues_JobGradeList struct {
	JobGradeList *JobGradeList `protobuf:"bytes,3,opt,name=job_grade_list,json=jobGradeList,proto3,oneof"`
}

func (*AttributeValues_StringList) isAttributeValues_ValidValues() {}

func (*AttributeValues_JobList) isAttributeValues_ValidValues() {}

func (*AttributeValues_JobGradeList) isAttributeValues_ValidValues() {}

type StringList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings []string `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
}

func (x *StringList) Reset() {
	*x = StringList{}
	mi := &file_resources_permissions_permissions_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StringList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringList) ProtoMessage() {}

func (x *StringList) ProtoReflect() protoreflect.Message {
	mi := &file_resources_permissions_permissions_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringList.ProtoReflect.Descriptor instead.
func (*StringList) Descriptor() ([]byte, []int) {
	return file_resources_permissions_permissions_proto_rawDescGZIP(), []int{5}
}

func (x *StringList) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

type JobGradeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs map[string]int32 `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *JobGradeList) Reset() {
	*x = JobGradeList{}
	mi := &file_resources_permissions_permissions_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobGradeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobGradeList) ProtoMessage() {}

func (x *JobGradeList) ProtoReflect() protoreflect.Message {
	mi := &file_resources_permissions_permissions_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobGradeList.ProtoReflect.Descriptor instead.
func (*JobGradeList) Descriptor() ([]byte, []int) {
	return file_resources_permissions_permissions_proto_rawDescGZIP(), []int{6}
}

func (x *JobGradeList) GetJobs() map[string]int32 {
	if x != nil {
		return x.Jobs
	}
	return nil
}

var File_resources_permissions_permissions_proto protoreflect.FileDescriptor

var file_resources_permissions_permissions_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80,
	0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01, 0x52, 0x09, 0x67, 0x75, 0x61, 0x72, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x22, 0xb6, 0x03, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01, 0x52, 0x03, 0x6a, 0x6f, 0x62,
	0x12, 0x29, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x01, 0x52, 0x08,
	0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x05, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a,
	0x02, 0x28, 0x00, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x6a, 0x6f,
	0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x02, 0x52, 0x0d,
	0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x43, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6a,
	0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6a, 0x6f, 0x62,
	0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xcf, 0x03, 0x0a,
	0x10, 0x52, 0x61, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x42,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x61, 0x74, 0x74, 0x72, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x18, 0x80, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x18, 0xff, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0xa7,
	0x04, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x1b, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x42, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x61, 0x74, 0x74, 0x72, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x18, 0x80, 0x01, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18,
	0xff, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4a,
	0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x01, 0x52, 0x09, 0x6d, 0x61,
	0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x61,
	0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xf4, 0x01, 0x0a, 0x0f, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0e, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0c, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x0e, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0x26, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4a,
	0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x4a,
	0x6f, 0x62, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_permissions_permissions_proto_rawDescOnce sync.Once
	file_resources_permissions_permissions_proto_rawDescData = file_resources_permissions_permissions_proto_rawDesc
)

func file_resources_permissions_permissions_proto_rawDescGZIP() []byte {
	file_resources_permissions_permissions_proto_rawDescOnce.Do(func() {
		file_resources_permissions_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_permissions_permissions_proto_rawDescData)
	})
	return file_resources_permissions_permissions_proto_rawDescData
}

var file_resources_permissions_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_resources_permissions_permissions_proto_goTypes = []any{
	(*Permission)(nil),          // 0: resources.permissions.Permission
	(*Role)(nil),                // 1: resources.permissions.Role
	(*RawRoleAttribute)(nil),    // 2: resources.permissions.RawRoleAttribute
	(*RoleAttribute)(nil),       // 3: resources.permissions.RoleAttribute
	(*AttributeValues)(nil),     // 4: resources.permissions.AttributeValues
	(*StringList)(nil),          // 5: resources.permissions.StringList
	(*JobGradeList)(nil),        // 6: resources.permissions.JobGradeList
	nil,                         // 7: resources.permissions.JobGradeList.JobsEntry
	(*timestamp.Timestamp)(nil), // 8: resources.timestamp.Timestamp
}
var file_resources_permissions_permissions_proto_depIdxs = []int32{
	8,  // 0: resources.permissions.Permission.created_at:type_name -> resources.timestamp.Timestamp
	8,  // 1: resources.permissions.Role.created_at:type_name -> resources.timestamp.Timestamp
	0,  // 2: resources.permissions.Role.permissions:type_name -> resources.permissions.Permission
	3,  // 3: resources.permissions.Role.attributes:type_name -> resources.permissions.RoleAttribute
	8,  // 4: resources.permissions.RawRoleAttribute.created_at:type_name -> resources.timestamp.Timestamp
	4,  // 5: resources.permissions.RawRoleAttribute.valid_values:type_name -> resources.permissions.AttributeValues
	4,  // 6: resources.permissions.RawRoleAttribute.value:type_name -> resources.permissions.AttributeValues
	8,  // 7: resources.permissions.RoleAttribute.created_at:type_name -> resources.timestamp.Timestamp
	4,  // 8: resources.permissions.RoleAttribute.valid_values:type_name -> resources.permissions.AttributeValues
	4,  // 9: resources.permissions.RoleAttribute.value:type_name -> resources.permissions.AttributeValues
	4,  // 10: resources.permissions.RoleAttribute.max_values:type_name -> resources.permissions.AttributeValues
	5,  // 11: resources.permissions.AttributeValues.string_list:type_name -> resources.permissions.StringList
	5,  // 12: resources.permissions.AttributeValues.job_list:type_name -> resources.permissions.StringList
	6,  // 13: resources.permissions.AttributeValues.job_grade_list:type_name -> resources.permissions.JobGradeList
	7,  // 14: resources.permissions.JobGradeList.jobs:type_name -> resources.permissions.JobGradeList.JobsEntry
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_resources_permissions_permissions_proto_init() }
func file_resources_permissions_permissions_proto_init() {
	if File_resources_permissions_permissions_proto != nil {
		return
	}
	file_resources_permissions_permissions_proto_msgTypes[0].OneofWrappers = []any{}
	file_resources_permissions_permissions_proto_msgTypes[1].OneofWrappers = []any{}
	file_resources_permissions_permissions_proto_msgTypes[2].OneofWrappers = []any{}
	file_resources_permissions_permissions_proto_msgTypes[3].OneofWrappers = []any{}
	file_resources_permissions_permissions_proto_msgTypes[4].OneofWrappers = []any{
		(*AttributeValues_StringList)(nil),
		(*AttributeValues_JobList)(nil),
		(*AttributeValues_JobGradeList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_permissions_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_permissions_permissions_proto_goTypes,
		DependencyIndexes: file_resources_permissions_permissions_proto_depIdxs,
		MessageInfos:      file_resources_permissions_permissions_proto_msgTypes,
	}.Build()
	File_resources_permissions_permissions_proto = out.File
	file_resources_permissions_permissions_proto_rawDesc = nil
	file_resources_permissions_permissions_proto_goTypes = nil
	file_resources_permissions_permissions_proto_depIdxs = nil
}
