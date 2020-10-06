// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.13.0
// source: auth/auth_api.proto

package authv1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User email.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// User password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{1}
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User email.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// User password.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{2}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session ID.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Session expiration time.
	ExpireTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SignInResponse) GetExpireTime() *timestamp.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

type SignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutRequest) Reset() {
	*x = SignOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRequest) ProtoMessage() {}

func (x *SignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRequest.ProtoReflect.Descriptor instead.
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{4}
}

type SignOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutResponse) Reset() {
	*x = SignOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutResponse) ProtoMessage() {}

func (x *SignOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutResponse.ProtoReflect.Descriptor instead.
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{5}
}

type RefreshSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshSessionRequest) Reset() {
	*x = RefreshSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshSessionRequest) ProtoMessage() {}

func (x *RefreshSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshSessionRequest.ProtoReflect.Descriptor instead.
func (*RefreshSessionRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{6}
}

type RefreshSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session expiration time.
	ExpireTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *RefreshSessionResponse) Reset() {
	*x = RefreshSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshSessionResponse) ProtoMessage() {}

func (x *RefreshSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshSessionResponse.ProtoReflect.Descriptor instead.
func (*RefreshSessionResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{7}
}

func (x *RefreshSessionResponse) GetExpireTime() *timestamp.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User email.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{8}
}

func (x *ResetPasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResetPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetPasswordResponse) Reset() {
	*x = ResetPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordResponse) ProtoMessage() {}

func (x *ResetPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordResponse.ProtoReflect.Descriptor instead.
func (*ResetPasswordResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_api_proto_rawDescGZIP(), []int{9}
}

var File_auth_auth_api_proto protoreflect.FileDescriptor

var file_auth_auth_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a,
	0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x6c, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x10,
	0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x11, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x16,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x8a, 0x04, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x41, 0x50, 0x49, 0x12, 0x5b,
	0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x27, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x27, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e,
	0x4f, 0x75, 0x74, 0x12, 0x28, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x70, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a,
	0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2e,
	0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0d, 0x5a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_api_proto_rawDescOnce sync.Once
	file_auth_auth_api_proto_rawDescData = file_auth_auth_api_proto_rawDesc
)

func file_auth_auth_api_proto_rawDescGZIP() []byte {
	file_auth_auth_api_proto_rawDescOnce.Do(func() {
		file_auth_auth_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_api_proto_rawDescData)
	})
	return file_auth_auth_api_proto_rawDescData
}

var (
	file_auth_auth_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
	file_auth_auth_api_proto_goTypes  = []interface{}{
		(*SignUpRequest)(nil),          // 0: percona.platform.auth.v1.SignUpRequest
		(*SignUpResponse)(nil),         // 1: percona.platform.auth.v1.SignUpResponse
		(*SignInRequest)(nil),          // 2: percona.platform.auth.v1.SignInRequest
		(*SignInResponse)(nil),         // 3: percona.platform.auth.v1.SignInResponse
		(*SignOutRequest)(nil),         // 4: percona.platform.auth.v1.SignOutRequest
		(*SignOutResponse)(nil),        // 5: percona.platform.auth.v1.SignOutResponse
		(*RefreshSessionRequest)(nil),  // 6: percona.platform.auth.v1.RefreshSessionRequest
		(*RefreshSessionResponse)(nil), // 7: percona.platform.auth.v1.RefreshSessionResponse
		(*ResetPasswordRequest)(nil),   // 8: percona.platform.auth.v1.ResetPasswordRequest
		(*ResetPasswordResponse)(nil),  // 9: percona.platform.auth.v1.ResetPasswordResponse
		(*timestamp.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	}
)

var file_auth_auth_api_proto_depIdxs = []int32{
	10, // 0: percona.platform.auth.v1.SignInResponse.expire_time:type_name -> google.protobuf.Timestamp
	10, // 1: percona.platform.auth.v1.RefreshSessionResponse.expire_time:type_name -> google.protobuf.Timestamp
	0,  // 2: percona.platform.auth.v1.AuthAPI.SignUp:input_type -> percona.platform.auth.v1.SignUpRequest
	2,  // 3: percona.platform.auth.v1.AuthAPI.SignIn:input_type -> percona.platform.auth.v1.SignInRequest
	4,  // 4: percona.platform.auth.v1.AuthAPI.SignOut:input_type -> percona.platform.auth.v1.SignOutRequest
	6,  // 5: percona.platform.auth.v1.AuthAPI.RefreshSession:input_type -> percona.platform.auth.v1.RefreshSessionRequest
	8,  // 6: percona.platform.auth.v1.AuthAPI.ResetPassword:input_type -> percona.platform.auth.v1.ResetPasswordRequest
	1,  // 7: percona.platform.auth.v1.AuthAPI.SignUp:output_type -> percona.platform.auth.v1.SignUpResponse
	3,  // 8: percona.platform.auth.v1.AuthAPI.SignIn:output_type -> percona.platform.auth.v1.SignInResponse
	5,  // 9: percona.platform.auth.v1.AuthAPI.SignOut:output_type -> percona.platform.auth.v1.SignOutResponse
	7,  // 10: percona.platform.auth.v1.AuthAPI.RefreshSession:output_type -> percona.platform.auth.v1.RefreshSessionResponse
	9,  // 11: percona.platform.auth.v1.AuthAPI.ResetPassword:output_type -> percona.platform.auth.v1.ResetPasswordResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_auth_auth_api_proto_init() }
func file_auth_auth_api_proto_init() {
	if File_auth_auth_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_auth_auth_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_auth_auth_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_auth_auth_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_auth_auth_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutRequest); i {
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
		file_auth_auth_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutResponse); i {
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
		file_auth_auth_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshSessionRequest); i {
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
		file_auth_auth_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshSessionResponse); i {
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
		file_auth_auth_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_auth_auth_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordResponse); i {
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
			RawDescriptor: file_auth_auth_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_api_proto_goTypes,
		DependencyIndexes: file_auth_auth_api_proto_depIdxs,
		MessageInfos:      file_auth_auth_api_proto_msgTypes,
	}.Build()
	File_auth_auth_api_proto = out.File
	file_auth_auth_api_proto_rawDesc = nil
	file_auth_auth_api_proto_goTypes = nil
	file_auth_auth_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConnInterface
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthAPIClient is the client API for AuthAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthAPIClient interface {
	// SignUp creates new user with given email and password.
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	// SignIn checks user authentication, creates session and returns session ID.
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// SignOut terminates user session.
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	// RefreshSession refreshes session timeout.
	RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error)
	// ResetPassword initiates user's password reset procedure.
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
}

type authAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthAPIClient(cc grpc.ClientConnInterface) AuthAPIClient {
	return &authAPIClient{cc}
}

func (c *authAPIClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error) {
	out := new(RefreshSessionResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/RefreshSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthAPIServer is the server API for AuthAPI service.
type AuthAPIServer interface {
	// SignUp creates new user with given email and password.
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	// SignIn checks user authentication, creates session and returns session ID.
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// SignOut terminates user session.
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	// RefreshSession refreshes session timeout.
	RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error)
	// ResetPassword initiates user's password reset procedure.
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
}

// UnimplementedAuthAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAuthAPIServer struct {
}

func (*UnimplementedAuthAPIServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func (*UnimplementedAuthAPIServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}

func (*UnimplementedAuthAPIServer) SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func (*UnimplementedAuthAPIServer) RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}

func (*UnimplementedAuthAPIServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}

func RegisterAuthAPIServer(s *grpc.Server, srv AuthAPIServer) {
	s.RegisterService(&_AuthAPI_serviceDesc, srv)
}

func _AuthAPI_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/RefreshSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).RefreshSession(ctx, req.(*RefreshSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "percona.platform.auth.v1.AuthAPI",
	HandlerType: (*AuthAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthAPI_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthAPI_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthAPI_SignOut_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _AuthAPI_RefreshSession_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthAPI_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth_api.proto",
}