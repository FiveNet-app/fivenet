// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.20.3
// source: resources/jobs/activity.proto

package jobs

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

type JobsUserActivityType int32

const (
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED  JobsUserActivityType = 0
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_HIRED        JobsUserActivityType = 1
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_FIRED        JobsUserActivityType = 2
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_PROMOTED     JobsUserActivityType = 3
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_DEMOTED      JobsUserActivityType = 4
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE JobsUserActivityType = 5
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_NOTE         JobsUserActivityType = 6
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_LABELS       JobsUserActivityType = 7
	JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_NAME         JobsUserActivityType = 8
)

// Enum value maps for JobsUserActivityType.
var (
	JobsUserActivityType_name = map[int32]string{
		0: "JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED",
		1: "JOBS_USER_ACTIVITY_TYPE_HIRED",
		2: "JOBS_USER_ACTIVITY_TYPE_FIRED",
		3: "JOBS_USER_ACTIVITY_TYPE_PROMOTED",
		4: "JOBS_USER_ACTIVITY_TYPE_DEMOTED",
		5: "JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE",
		6: "JOBS_USER_ACTIVITY_TYPE_NOTE",
		7: "JOBS_USER_ACTIVITY_TYPE_LABELS",
		8: "JOBS_USER_ACTIVITY_TYPE_NAME",
	}
	JobsUserActivityType_value = map[string]int32{
		"JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED":  0,
		"JOBS_USER_ACTIVITY_TYPE_HIRED":        1,
		"JOBS_USER_ACTIVITY_TYPE_FIRED":        2,
		"JOBS_USER_ACTIVITY_TYPE_PROMOTED":     3,
		"JOBS_USER_ACTIVITY_TYPE_DEMOTED":      4,
		"JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE": 5,
		"JOBS_USER_ACTIVITY_TYPE_NOTE":         6,
		"JOBS_USER_ACTIVITY_TYPE_LABELS":       7,
		"JOBS_USER_ACTIVITY_TYPE_NAME":         8,
	}
)

func (x JobsUserActivityType) Enum() *JobsUserActivityType {
	p := new(JobsUserActivityType)
	*p = x
	return p
}

func (x JobsUserActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobsUserActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_jobs_activity_proto_enumTypes[0].Descriptor()
}

func (JobsUserActivityType) Type() protoreflect.EnumType {
	return &file_resources_jobs_activity_proto_enumTypes[0]
}

func (x JobsUserActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobsUserActivityType.Descriptor instead.
func (JobsUserActivityType) EnumDescriptor() ([]byte, []int) {
	return file_resources_jobs_activity_proto_rawDescGZIP(), []int{0}
}

type JobsUserActivity struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	Id           uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" sql:"primary_key" alias:"id"` // @gotags: sql:"primary_key" alias:"id"
	CreatedAt    *timestamp.Timestamp   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	Job          string                 `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	SourceUserId *int32                 `protobuf:"varint,5,opt,name=source_user_id,json=sourceUserId,proto3,oneof" json:"source_user_id,omitempty"`
	SourceUser   *Colleague             `protobuf:"bytes,6,opt,name=source_user,json=sourceUser,proto3,oneof" json:"source_user,omitempty" alias:"source_user"` // @gotags: alias:"source_user"
	TargetUserId int32                  `protobuf:"varint,7,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	TargetUser   *Colleague             `protobuf:"bytes,8,opt,name=target_user,json=targetUser,proto3" json:"target_user,omitempty" alias:"target_user"` // @gotags: alias:"target_user"
	ActivityType JobsUserActivityType   `protobuf:"varint,9,opt,name=activity_type,json=activityType,proto3,enum=resources.jobs.JobsUserActivityType" json:"activity_type,omitempty"`
	// @sanitize
	Reason        string                `protobuf:"bytes,10,opt,name=reason,proto3" json:"reason,omitempty"`
	Data          *JobsUserActivityData `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobsUserActivity) Reset() {
	*x = JobsUserActivity{}
	mi := &file_resources_jobs_activity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobsUserActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsUserActivity) ProtoMessage() {}

func (x *JobsUserActivity) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_activity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsUserActivity.ProtoReflect.Descriptor instead.
func (*JobsUserActivity) Descriptor() ([]byte, []int) {
	return file_resources_jobs_activity_proto_rawDescGZIP(), []int{0}
}

func (x *JobsUserActivity) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JobsUserActivity) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *JobsUserActivity) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *JobsUserActivity) GetSourceUserId() int32 {
	if x != nil && x.SourceUserId != nil {
		return *x.SourceUserId
	}
	return 0
}

func (x *JobsUserActivity) GetSourceUser() *Colleague {
	if x != nil {
		return x.SourceUser
	}
	return nil
}

func (x *JobsUserActivity) GetTargetUserId() int32 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

func (x *JobsUserActivity) GetTargetUser() *Colleague {
	if x != nil {
		return x.TargetUser
	}
	return nil
}

func (x *JobsUserActivity) GetActivityType() JobsUserActivityType {
	if x != nil {
		return x.ActivityType
	}
	return JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_UNSPECIFIED
}

func (x *JobsUserActivity) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *JobsUserActivity) GetData() *JobsUserActivityData {
	if x != nil {
		return x.Data
	}
	return nil
}

type JobsUserActivityData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*JobsUserActivityData_AbsenceDate
	//	*JobsUserActivityData_GradeChange
	//	*JobsUserActivityData_LabelsChange
	//	*JobsUserActivityData_NameChange
	Data          isJobsUserActivityData_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobsUserActivityData) Reset() {
	*x = JobsUserActivityData{}
	mi := &file_resources_jobs_activity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobsUserActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsUserActivityData) ProtoMessage() {}

func (x *JobsUserActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_activity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsUserActivityData.ProtoReflect.Descriptor instead.
func (*JobsUserActivityData) Descriptor() ([]byte, []int) {
	return file_resources_jobs_activity_proto_rawDescGZIP(), []int{1}
}

func (x *JobsUserActivityData) GetData() isJobsUserActivityData_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *JobsUserActivityData) GetAbsenceDate() *ColleagueAbsenceDate {
	if x != nil {
		if x, ok := x.Data.(*JobsUserActivityData_AbsenceDate); ok {
			return x.AbsenceDate
		}
	}
	return nil
}

func (x *JobsUserActivityData) GetGradeChange() *ColleagueGradeChange {
	if x != nil {
		if x, ok := x.Data.(*JobsUserActivityData_GradeChange); ok {
			return x.GradeChange
		}
	}
	return nil
}

func (x *JobsUserActivityData) GetLabelsChange() *ColleagueLabelsChange {
	if x != nil {
		if x, ok := x.Data.(*JobsUserActivityData_LabelsChange); ok {
			return x.LabelsChange
		}
	}
	return nil
}

func (x *JobsUserActivityData) GetNameChange() *ColleagueNameChange {
	if x != nil {
		if x, ok := x.Data.(*JobsUserActivityData_NameChange); ok {
			return x.NameChange
		}
	}
	return nil
}

type isJobsUserActivityData_Data interface {
	isJobsUserActivityData_Data()
}

type JobsUserActivityData_AbsenceDate struct {
	AbsenceDate *ColleagueAbsenceDate `protobuf:"bytes,1,opt,name=absence_date,json=absenceDate,proto3,oneof"`
}

type JobsUserActivityData_GradeChange struct {
	GradeChange *ColleagueGradeChange `protobuf:"bytes,2,opt,name=grade_change,json=gradeChange,proto3,oneof"`
}

type JobsUserActivityData_LabelsChange struct {
	LabelsChange *ColleagueLabelsChange `protobuf:"bytes,3,opt,name=labels_change,json=labelsChange,proto3,oneof"`
}

type JobsUserActivityData_NameChange struct {
	NameChange *ColleagueNameChange `protobuf:"bytes,4,opt,name=name_change,json=nameChange,proto3,oneof"`
}

func (*JobsUserActivityData_AbsenceDate) isJobsUserActivityData_Data() {}

func (*JobsUserActivityData_GradeChange) isJobsUserActivityData_Data() {}

func (*JobsUserActivityData_LabelsChange) isJobsUserActivityData_Data() {}

func (*JobsUserActivityData_NameChange) isJobsUserActivityData_Data() {}

type ColleagueAbsenceDate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AbsenceBegin  *timestamp.Timestamp   `protobuf:"bytes,1,opt,name=absence_begin,json=absenceBegin,proto3" json:"absence_begin,omitempty"`
	AbsenceEnd    *timestamp.Timestamp   `protobuf:"bytes,2,opt,name=absence_end,json=absenceEnd,proto3" json:"absence_end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColleagueAbsenceDate) Reset() {
	*x = ColleagueAbsenceDate{}
	mi := &file_resources_jobs_activity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColleagueAbsenceDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleagueAbsenceDate) ProtoMessage() {}

func (x *ColleagueAbsenceDate) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_activity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleagueAbsenceDate.ProtoReflect.Descriptor instead.
func (*ColleagueAbsenceDate) Descriptor() ([]byte, []int) {
	return file_resources_jobs_activity_proto_rawDescGZIP(), []int{2}
}

func (x *ColleagueAbsenceDate) GetAbsenceBegin() *timestamp.Timestamp {
	if x != nil {
		return x.AbsenceBegin
	}
	return nil
}

func (x *ColleagueAbsenceDate) GetAbsenceEnd() *timestamp.Timestamp {
	if x != nil {
		return x.AbsenceEnd
	}
	return nil
}

type ColleagueGradeChange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Grade         int32                  `protobuf:"varint,1,opt,name=grade,proto3" json:"grade,omitempty"`
	GradeLabel    string                 `protobuf:"bytes,2,opt,name=grade_label,json=gradeLabel,proto3" json:"grade_label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColleagueGradeChange) Reset() {
	*x = ColleagueGradeChange{}
	mi := &file_resources_jobs_activity_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColleagueGradeChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleagueGradeChange) ProtoMessage() {}

func (x *ColleagueGradeChange) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_activity_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleagueGradeChange.ProtoReflect.Descriptor instead.
func (*ColleagueGradeChange) Descriptor() ([]byte, []int) {
	return file_resources_jobs_activity_proto_rawDescGZIP(), []int{3}
}

func (x *ColleagueGradeChange) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *ColleagueGradeChange) GetGradeLabel() string {
	if x != nil {
		return x.GradeLabel
	}
	return ""
}

type ColleagueLabelsChange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Added         []*Label               `protobuf:"bytes,1,rep,name=added,proto3" json:"added,omitempty"`
	Removed       []*Label               `protobuf:"bytes,2,rep,name=removed,proto3" json:"removed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColleagueLabelsChange) Reset() {
	*x = ColleagueLabelsChange{}
	mi := &file_resources_jobs_activity_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColleagueLabelsChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleagueLabelsChange) ProtoMessage() {}

func (x *ColleagueLabelsChange) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_activity_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleagueLabelsChange.ProtoReflect.Descriptor instead.
func (*ColleagueLabelsChange) Descriptor() ([]byte, []int) {
	return file_resources_jobs_activity_proto_rawDescGZIP(), []int{4}
}

func (x *ColleagueLabelsChange) GetAdded() []*Label {
	if x != nil {
		return x.Added
	}
	return nil
}

func (x *ColleagueLabelsChange) GetRemoved() []*Label {
	if x != nil {
		return x.Removed
	}
	return nil
}

type ColleagueNameChange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Prefix        *string                `protobuf:"bytes,1,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
	Suffix        *string                `protobuf:"bytes,2,opt,name=suffix,proto3,oneof" json:"suffix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColleagueNameChange) Reset() {
	*x = ColleagueNameChange{}
	mi := &file_resources_jobs_activity_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColleagueNameChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColleagueNameChange) ProtoMessage() {}

func (x *ColleagueNameChange) ProtoReflect() protoreflect.Message {
	mi := &file_resources_jobs_activity_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColleagueNameChange.ProtoReflect.Descriptor instead.
func (*ColleagueNameChange) Descriptor() ([]byte, []int) {
	return file_resources_jobs_activity_proto_rawDescGZIP(), []int{5}
}

func (x *ColleagueNameChange) GetPrefix() string {
	if x != nil && x.Prefix != nil {
		return *x.Prefix
	}
	return ""
}

func (x *ColleagueNameChange) GetSuffix() string {
	if x != nil && x.Suffix != nil {
		return *x.Suffix
	}
	return ""
}

var File_resources_jobs_activity_proto protoreflect.FileDescriptor

var file_resources_jobs_activity_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x1a,
	0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x04, 0x0a, 0x10,
	0x4a, 0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x03,
	0x6a, 0x6f, 0x62, 0x12, 0x32, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x20, 0x00, 0x48, 0x01, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x48, 0x02, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e,
	0x4a, 0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x22, 0xcf, 0x02, 0x0a,
	0x14, 0x4a, 0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x0c, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x49, 0x0a, 0x0c, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x42, 0x0b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x9c,
	0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x41, 0x62, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x61, 0x62, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c,
	0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x3f, 0x0a, 0x0b,
	0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x61, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x22, 0x4d, 0x0a,
	0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x75, 0x0a, 0x15,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x22, 0x65, 0x0a, 0x13, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x2a, 0xe2, 0x02, 0x0a, 0x14, 0x4a,
	0x6f, 0x62, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d,
	0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49,
	0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x21, 0x0a, 0x1d, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x4a, 0x4f, 0x42, 0x53,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x28, 0x0a,
	0x24, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x42, 0x53, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x4a, 0x4f, 0x42, 0x53, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x45, 0x10, 0x06, 0x12, 0x22, 0x0a, 0x1e, 0x4a, 0x4f, 0x42,
	0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x53, 0x10, 0x07, 0x12, 0x20, 0x0a,
	0x1c, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x08, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x69, 0x76, 0x65, 0x6e, 0x65,
	0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x3b, 0x6a, 0x6f,
	0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_jobs_activity_proto_rawDescOnce sync.Once
	file_resources_jobs_activity_proto_rawDescData = file_resources_jobs_activity_proto_rawDesc
)

func file_resources_jobs_activity_proto_rawDescGZIP() []byte {
	file_resources_jobs_activity_proto_rawDescOnce.Do(func() {
		file_resources_jobs_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_jobs_activity_proto_rawDescData)
	})
	return file_resources_jobs_activity_proto_rawDescData
}

var file_resources_jobs_activity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_jobs_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_resources_jobs_activity_proto_goTypes = []any{
	(JobsUserActivityType)(0),     // 0: resources.jobs.JobsUserActivityType
	(*JobsUserActivity)(nil),      // 1: resources.jobs.JobsUserActivity
	(*JobsUserActivityData)(nil),  // 2: resources.jobs.JobsUserActivityData
	(*ColleagueAbsenceDate)(nil),  // 3: resources.jobs.ColleagueAbsenceDate
	(*ColleagueGradeChange)(nil),  // 4: resources.jobs.ColleagueGradeChange
	(*ColleagueLabelsChange)(nil), // 5: resources.jobs.ColleagueLabelsChange
	(*ColleagueNameChange)(nil),   // 6: resources.jobs.ColleagueNameChange
	(*timestamp.Timestamp)(nil),   // 7: resources.timestamp.Timestamp
	(*Colleague)(nil),             // 8: resources.jobs.Colleague
	(*Label)(nil),                 // 9: resources.jobs.Label
}
var file_resources_jobs_activity_proto_depIdxs = []int32{
	7,  // 0: resources.jobs.JobsUserActivity.created_at:type_name -> resources.timestamp.Timestamp
	8,  // 1: resources.jobs.JobsUserActivity.source_user:type_name -> resources.jobs.Colleague
	8,  // 2: resources.jobs.JobsUserActivity.target_user:type_name -> resources.jobs.Colleague
	0,  // 3: resources.jobs.JobsUserActivity.activity_type:type_name -> resources.jobs.JobsUserActivityType
	2,  // 4: resources.jobs.JobsUserActivity.data:type_name -> resources.jobs.JobsUserActivityData
	3,  // 5: resources.jobs.JobsUserActivityData.absence_date:type_name -> resources.jobs.ColleagueAbsenceDate
	4,  // 6: resources.jobs.JobsUserActivityData.grade_change:type_name -> resources.jobs.ColleagueGradeChange
	5,  // 7: resources.jobs.JobsUserActivityData.labels_change:type_name -> resources.jobs.ColleagueLabelsChange
	6,  // 8: resources.jobs.JobsUserActivityData.name_change:type_name -> resources.jobs.ColleagueNameChange
	7,  // 9: resources.jobs.ColleagueAbsenceDate.absence_begin:type_name -> resources.timestamp.Timestamp
	7,  // 10: resources.jobs.ColleagueAbsenceDate.absence_end:type_name -> resources.timestamp.Timestamp
	9,  // 11: resources.jobs.ColleagueLabelsChange.added:type_name -> resources.jobs.Label
	9,  // 12: resources.jobs.ColleagueLabelsChange.removed:type_name -> resources.jobs.Label
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_resources_jobs_activity_proto_init() }
func file_resources_jobs_activity_proto_init() {
	if File_resources_jobs_activity_proto != nil {
		return
	}
	file_resources_jobs_colleagues_proto_init()
	file_resources_jobs_labels_proto_init()
	file_resources_jobs_activity_proto_msgTypes[0].OneofWrappers = []any{}
	file_resources_jobs_activity_proto_msgTypes[1].OneofWrappers = []any{
		(*JobsUserActivityData_AbsenceDate)(nil),
		(*JobsUserActivityData_GradeChange)(nil),
		(*JobsUserActivityData_LabelsChange)(nil),
		(*JobsUserActivityData_NameChange)(nil),
	}
	file_resources_jobs_activity_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_jobs_activity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_jobs_activity_proto_goTypes,
		DependencyIndexes: file_resources_jobs_activity_proto_depIdxs,
		EnumInfos:         file_resources_jobs_activity_proto_enumTypes,
		MessageInfos:      file_resources_jobs_activity_proto_msgTypes,
	}.Build()
	File_resources_jobs_activity_proto = out.File
	file_resources_jobs_activity_proto_rawDesc = nil
	file_resources_jobs_activity_proto_goTypes = nil
	file_resources_jobs_activity_proto_depIdxs = nil
}
