// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.3
// source: resources/qualifications/access.proto

package qualifications

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

type AccessLevelUpdateMode int32

const (
	AccessLevelUpdateMode_ACCESS_LEVEL_UPDATE_MODE_UNSPECIFIED AccessLevelUpdateMode = 0
	AccessLevelUpdateMode_ACCESS_LEVEL_UPDATE_MODE_UPDATE      AccessLevelUpdateMode = 1
	AccessLevelUpdateMode_ACCESS_LEVEL_UPDATE_MODE_DELETE      AccessLevelUpdateMode = 2
	AccessLevelUpdateMode_ACCESS_LEVEL_UPDATE_MODE_CLEAR       AccessLevelUpdateMode = 3
)

// Enum value maps for AccessLevelUpdateMode.
var (
	AccessLevelUpdateMode_name = map[int32]string{
		0: "ACCESS_LEVEL_UPDATE_MODE_UNSPECIFIED",
		1: "ACCESS_LEVEL_UPDATE_MODE_UPDATE",
		2: "ACCESS_LEVEL_UPDATE_MODE_DELETE",
		3: "ACCESS_LEVEL_UPDATE_MODE_CLEAR",
	}
	AccessLevelUpdateMode_value = map[string]int32{
		"ACCESS_LEVEL_UPDATE_MODE_UNSPECIFIED": 0,
		"ACCESS_LEVEL_UPDATE_MODE_UPDATE":      1,
		"ACCESS_LEVEL_UPDATE_MODE_DELETE":      2,
		"ACCESS_LEVEL_UPDATE_MODE_CLEAR":       3,
	}
)

func (x AccessLevelUpdateMode) Enum() *AccessLevelUpdateMode {
	p := new(AccessLevelUpdateMode)
	*p = x
	return p
}

func (x AccessLevelUpdateMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessLevelUpdateMode) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_qualifications_access_proto_enumTypes[0].Descriptor()
}

func (AccessLevelUpdateMode) Type() protoreflect.EnumType {
	return &file_resources_qualifications_access_proto_enumTypes[0]
}

func (x AccessLevelUpdateMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessLevelUpdateMode.Descriptor instead.
func (AccessLevelUpdateMode) EnumDescriptor() ([]byte, []int) {
	return file_resources_qualifications_access_proto_rawDescGZIP(), []int{0}
}

type AccessLevel int32

const (
	AccessLevel_ACCESS_LEVEL_UNSPECIFIED AccessLevel = 0
	AccessLevel_ACCESS_LEVEL_BLOCKED     AccessLevel = 1
	AccessLevel_ACCESS_LEVEL_VIEW        AccessLevel = 2
	AccessLevel_ACCESS_LEVEL_REQUEST     AccessLevel = 3
	AccessLevel_ACCESS_LEVEL_TAKE        AccessLevel = 4
	AccessLevel_ACCESS_LEVEL_GRADE       AccessLevel = 5
	AccessLevel_ACCESS_LEVEL_MANAGE      AccessLevel = 6
	AccessLevel_ACCESS_LEVEL_EDIT        AccessLevel = 7
)

// Enum value maps for AccessLevel.
var (
	AccessLevel_name = map[int32]string{
		0: "ACCESS_LEVEL_UNSPECIFIED",
		1: "ACCESS_LEVEL_BLOCKED",
		2: "ACCESS_LEVEL_VIEW",
		3: "ACCESS_LEVEL_REQUEST",
		4: "ACCESS_LEVEL_TAKE",
		5: "ACCESS_LEVEL_GRADE",
		6: "ACCESS_LEVEL_MANAGE",
		7: "ACCESS_LEVEL_EDIT",
	}
	AccessLevel_value = map[string]int32{
		"ACCESS_LEVEL_UNSPECIFIED": 0,
		"ACCESS_LEVEL_BLOCKED":     1,
		"ACCESS_LEVEL_VIEW":        2,
		"ACCESS_LEVEL_REQUEST":     3,
		"ACCESS_LEVEL_TAKE":        4,
		"ACCESS_LEVEL_GRADE":       5,
		"ACCESS_LEVEL_MANAGE":      6,
		"ACCESS_LEVEL_EDIT":        7,
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
	return file_resources_qualifications_access_proto_enumTypes[1].Descriptor()
}

func (AccessLevel) Type() protoreflect.EnumType {
	return &file_resources_qualifications_access_proto_enumTypes[1]
}

func (x AccessLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessLevel.Descriptor instead.
func (AccessLevel) EnumDescriptor() ([]byte, []int) {
	return file_resources_qualifications_access_proto_rawDescGZIP(), []int{1}
}

type QualificationAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*QualificationJobAccess `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *QualificationAccess) Reset() {
	*x = QualificationAccess{}
	mi := &file_resources_qualifications_access_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QualificationAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QualificationAccess) ProtoMessage() {}

func (x *QualificationAccess) ProtoReflect() protoreflect.Message {
	mi := &file_resources_qualifications_access_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QualificationAccess.ProtoReflect.Descriptor instead.
func (*QualificationAccess) Descriptor() ([]byte, []int) {
	return file_resources_qualifications_access_proto_rawDescGZIP(), []int{0}
}

func (x *QualificationAccess) GetJobs() []*QualificationJobAccess {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type QualificationJobAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt       *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	QualificationId uint64               `protobuf:"varint,4,opt,name=qualification_id,json=qualificationId,proto3" json:"qualification_id,omitempty"`
	Job             string               `protobuf:"bytes,5,opt,name=job,proto3" json:"job,omitempty"`
	JobLabel        *string              `protobuf:"bytes,6,opt,name=job_label,json=jobLabel,proto3,oneof" json:"job_label,omitempty"`
	MinimumGrade    int32                `protobuf:"varint,7,opt,name=minimum_grade,json=minimumGrade,proto3" json:"minimum_grade,omitempty"`
	JobGradeLabel   *string              `protobuf:"bytes,8,opt,name=job_grade_label,json=jobGradeLabel,proto3,oneof" json:"job_grade_label,omitempty"`
	Access          AccessLevel          `protobuf:"varint,9,opt,name=access,proto3,enum=resources.qualifications.AccessLevel" json:"access,omitempty"`
}

func (x *QualificationJobAccess) Reset() {
	*x = QualificationJobAccess{}
	mi := &file_resources_qualifications_access_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QualificationJobAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QualificationJobAccess) ProtoMessage() {}

func (x *QualificationJobAccess) ProtoReflect() protoreflect.Message {
	mi := &file_resources_qualifications_access_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QualificationJobAccess.ProtoReflect.Descriptor instead.
func (*QualificationJobAccess) Descriptor() ([]byte, []int) {
	return file_resources_qualifications_access_proto_rawDescGZIP(), []int{1}
}

func (x *QualificationJobAccess) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QualificationJobAccess) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *QualificationJobAccess) GetQualificationId() uint64 {
	if x != nil {
		return x.QualificationId
	}
	return 0
}

func (x *QualificationJobAccess) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *QualificationJobAccess) GetJobLabel() string {
	if x != nil && x.JobLabel != nil {
		return *x.JobLabel
	}
	return ""
}

func (x *QualificationJobAccess) GetMinimumGrade() int32 {
	if x != nil {
		return x.MinimumGrade
	}
	return 0
}

func (x *QualificationJobAccess) GetJobGradeLabel() string {
	if x != nil && x.JobGradeLabel != nil {
		return *x.JobGradeLabel
	}
	return ""
}

func (x *QualificationJobAccess) GetAccess() AccessLevel {
	if x != nil {
		return x.Access
	}
	return AccessLevel_ACCESS_LEVEL_UNSPECIFIED
}

var File_resources_qualifications_access_proto protoreflect.FileDescriptor

var file_resources_qualifications_access_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5b, 0x0a, 0x13, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x44, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xc3, 0x03, 0x0a,
	0x16, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f,
	0x62, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x2d, 0x0a, 0x10, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0f, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x18, 0x14, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x29, 0x0a, 0x09, 0x6a, 0x6f, 0x62,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x18, 0x32, 0x48, 0x01, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x28, 0x00, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x18, 0x32, 0x48, 0x02, 0x52, 0x0d, 0x6a, 0x6f, 0x62, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2a, 0xaf, 0x01, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x24,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02,
	0x12, 0x22, 0x0a, 0x1e, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4c, 0x45,
	0x41, 0x52, 0x10, 0x03, 0x2a, 0xd5, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x41,
	0x4b, 0x45, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x41, 0x4e,
	0x41, 0x47, 0x45, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x10, 0x07, 0x42, 0x55, 0x5a, 0x53,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x74, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_qualifications_access_proto_rawDescOnce sync.Once
	file_resources_qualifications_access_proto_rawDescData = file_resources_qualifications_access_proto_rawDesc
)

func file_resources_qualifications_access_proto_rawDescGZIP() []byte {
	file_resources_qualifications_access_proto_rawDescOnce.Do(func() {
		file_resources_qualifications_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_qualifications_access_proto_rawDescData)
	})
	return file_resources_qualifications_access_proto_rawDescData
}

var file_resources_qualifications_access_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_resources_qualifications_access_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_qualifications_access_proto_goTypes = []any{
	(AccessLevelUpdateMode)(0),     // 0: resources.qualifications.AccessLevelUpdateMode
	(AccessLevel)(0),               // 1: resources.qualifications.AccessLevel
	(*QualificationAccess)(nil),    // 2: resources.qualifications.QualificationAccess
	(*QualificationJobAccess)(nil), // 3: resources.qualifications.QualificationJobAccess
	(*timestamp.Timestamp)(nil),    // 4: resources.timestamp.Timestamp
}
var file_resources_qualifications_access_proto_depIdxs = []int32{
	3, // 0: resources.qualifications.QualificationAccess.jobs:type_name -> resources.qualifications.QualificationJobAccess
	4, // 1: resources.qualifications.QualificationJobAccess.created_at:type_name -> resources.timestamp.Timestamp
	1, // 2: resources.qualifications.QualificationJobAccess.access:type_name -> resources.qualifications.AccessLevel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_resources_qualifications_access_proto_init() }
func file_resources_qualifications_access_proto_init() {
	if File_resources_qualifications_access_proto != nil {
		return
	}
	file_resources_qualifications_access_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_qualifications_access_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_qualifications_access_proto_goTypes,
		DependencyIndexes: file_resources_qualifications_access_proto_depIdxs,
		EnumInfos:         file_resources_qualifications_access_proto_enumTypes,
		MessageInfos:      file_resources_qualifications_access_proto_msgTypes,
	}.Build()
	File_resources_qualifications_access_proto = out.File
	file_resources_qualifications_access_proto_rawDesc = nil
	file_resources_qualifications_access_proto_goTypes = nil
	file_resources_qualifications_access_proto_depIdxs = nil
}
