// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: auth/auth.proto

package auth

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	common "github.com/galexrt/arpanet/proto/common"
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string              `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Chars []*common.Character `protobuf:"bytes,2,rep,name=chars,proto3" json:"chars,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetChars() []*common.Character {
	if x != nil {
		return x.Chars
	}
	return nil
}

type ChooseCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *ChooseCharacterRequest) Reset() {
	*x = ChooseCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChooseCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChooseCharacterRequest) ProtoMessage() {}

func (x *ChooseCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChooseCharacterRequest.ProtoReflect.Descriptor instead.
func (*ChooseCharacterRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ChooseCharacterRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ChooseCharacterRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type ChooseCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ChooseCharacterResponse) Reset() {
	*x = ChooseCharacterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChooseCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChooseCharacterResponse) ProtoMessage() {}

func (x *ChooseCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChooseCharacterResponse.ProtoReflect.Descriptor instead.
func (*ChooseCharacterResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ChooseCharacterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{5}
}

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x16, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x52, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x63,
	0x68, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x52, 0x05, 0x63, 0x68, 0x61, 0x72, 0x73, 0x22, 0x59, 0x0a, 0x16, 0x43, 0x68, 0x6f, 0x6f,
	0x73, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x2e, 0x18, 0x40, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x17, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdf, 0x01,
	0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x43, 0x68,
	0x6f, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x6f, 0x6f, 0x73,
	0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x67,
	0x65, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61,
	0x6c, 0x65, 0x78, 0x72, 0x74, 0x2f, 0x61, 0x72, 0x70, 0x61, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData = file_auth_auth_proto_rawDesc
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_proto_rawDescData)
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_auth_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),            // 0: gen.auth.LoginRequest
	(*LoginResponse)(nil),           // 1: gen.auth.LoginResponse
	(*ChooseCharacterRequest)(nil),  // 2: gen.auth.ChooseCharacterRequest
	(*ChooseCharacterResponse)(nil), // 3: gen.auth.ChooseCharacterResponse
	(*LogoutRequest)(nil),           // 4: gen.auth.LogoutRequest
	(*LogoutResponse)(nil),          // 5: gen.auth.LogoutResponse
	(*common.Character)(nil),        // 6: gen.common.Character
}
var file_auth_auth_proto_depIdxs = []int32{
	6, // 0: gen.auth.LoginResponse.chars:type_name -> gen.common.Character
	0, // 1: gen.auth.AccountService.Login:input_type -> gen.auth.LoginRequest
	2, // 2: gen.auth.AccountService.ChooseCharacter:input_type -> gen.auth.ChooseCharacterRequest
	4, // 3: gen.auth.AccountService.Logout:input_type -> gen.auth.LogoutRequest
	1, // 4: gen.auth.AccountService.Login:output_type -> gen.auth.LoginResponse
	3, // 5: gen.auth.AccountService.ChooseCharacter:output_type -> gen.auth.ChooseCharacterResponse
	5, // 6: gen.auth.AccountService.Logout:output_type -> gen.auth.LogoutResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChooseCharacterRequest); i {
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
		file_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChooseCharacterResponse); i {
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
		file_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}
