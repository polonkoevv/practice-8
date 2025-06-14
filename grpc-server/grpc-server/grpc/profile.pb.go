// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: profile.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	mi := &file_profile_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	mi := &file_profile_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	mi := &file_profile_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProfileRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ProfileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileResponse) Reset() {
	*x = ProfileResponse{}
	mi := &file_profile_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResponse) ProtoMessage() {}

func (x *ProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResponse.ProtoReflect.Descriptor instead.
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DeleteProfileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProfileResponse) Reset() {
	*x = DeleteProfileResponse{}
	mi := &file_profile_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileResponse) ProtoMessage() {}

func (x *DeleteProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileResponse.ProtoReflect.Descriptor instead.
func (*DeleteProfileResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteProfileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteProfileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EmptyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_profile_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{5}
}

type ProfileListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Profiles      []*ProfileResponse     `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileListResponse) Reset() {
	*x = ProfileListResponse{}
	mi := &file_profile_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileListResponse) ProtoMessage() {}

func (x *ProfileListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileListResponse.ProtoReflect.Descriptor instead.
func (*ProfileListResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{6}
}

func (x *ProfileListResponse) GetProfiles() []*ProfileResponse {
	if x != nil {
		return x.Profiles
	}
	return nil
}

var File_profile_proto protoreflect.FileDescriptor

const file_profile_proto_rawDesc = "" +
	"\n" +
	"\rprofile.proto\x12\aprofile\" \n" +
	"\x0eProfileRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"@\n" +
	"\x14CreateProfileRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\"P\n" +
	"\x14UpdateProfileRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\"K\n" +
	"\x0fProfileResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\"K\n" +
	"\x15DeleteProfileResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\x0e\n" +
	"\fEmptyRequest\"K\n" +
	"\x13ProfileListResponse\x124\n" +
	"\bprofiles\x18\x01 \x03(\v2\x18.profile.ProfileResponseR\bprofiles2\xf6\x02\n" +
	"\x0eProfileService\x12?\n" +
	"\n" +
	"GetProfile\x12\x17.profile.ProfileRequest\x1a\x18.profile.ProfileResponse\x12H\n" +
	"\rCreateProfile\x12\x1d.profile.CreateProfileRequest\x1a\x18.profile.ProfileResponse\x12H\n" +
	"\rUpdateProfile\x12\x1d.profile.UpdateProfileRequest\x1a\x18.profile.ProfileResponse\x12H\n" +
	"\rDeleteProfile\x12\x17.profile.ProfileRequest\x1a\x1e.profile.DeleteProfileResponse\x12E\n" +
	"\x0eGetAllProfiles\x12\x15.profile.EmptyRequest\x1a\x1c.profile.ProfileListResponseB\x12Z\x10grpc-server/grpcb\x06proto3"

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData []byte
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_profile_proto_rawDesc), len(file_profile_proto_rawDesc)))
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_profile_proto_goTypes = []any{
	(*ProfileRequest)(nil),        // 0: profile.ProfileRequest
	(*CreateProfileRequest)(nil),  // 1: profile.CreateProfileRequest
	(*UpdateProfileRequest)(nil),  // 2: profile.UpdateProfileRequest
	(*ProfileResponse)(nil),       // 3: profile.ProfileResponse
	(*DeleteProfileResponse)(nil), // 4: profile.DeleteProfileResponse
	(*EmptyRequest)(nil),          // 5: profile.EmptyRequest
	(*ProfileListResponse)(nil),   // 6: profile.ProfileListResponse
}
var file_profile_proto_depIdxs = []int32{
	3, // 0: profile.ProfileListResponse.profiles:type_name -> profile.ProfileResponse
	0, // 1: profile.ProfileService.GetProfile:input_type -> profile.ProfileRequest
	1, // 2: profile.ProfileService.CreateProfile:input_type -> profile.CreateProfileRequest
	2, // 3: profile.ProfileService.UpdateProfile:input_type -> profile.UpdateProfileRequest
	0, // 4: profile.ProfileService.DeleteProfile:input_type -> profile.ProfileRequest
	5, // 5: profile.ProfileService.GetAllProfiles:input_type -> profile.EmptyRequest
	3, // 6: profile.ProfileService.GetProfile:output_type -> profile.ProfileResponse
	3, // 7: profile.ProfileService.CreateProfile:output_type -> profile.ProfileResponse
	3, // 8: profile.ProfileService.UpdateProfile:output_type -> profile.ProfileResponse
	4, // 9: profile.ProfileService.DeleteProfile:output_type -> profile.DeleteProfileResponse
	6, // 10: profile.ProfileService.GetAllProfiles:output_type -> profile.ProfileListResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_profile_proto_rawDesc), len(file_profile_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}
