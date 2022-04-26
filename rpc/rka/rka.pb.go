// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: rpc/rka/rka.proto

package rka

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type RKA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NameRka   string                 `protobuf:"bytes,2,opt,name=name_rka,json=nameRka,proto3" json:"name_rka,omitempty"`
	TargetRka int32                  `protobuf:"varint,3,opt,name=target_rka,json=targetRka,proto3" json:"target_rka,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *RKA) Reset() {
	*x = RKA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RKA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RKA) ProtoMessage() {}

func (x *RKA) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RKA.ProtoReflect.Descriptor instead.
func (*RKA) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{1}
}

func (x *RKA) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RKA) GetNameRka() string {
	if x != nil {
		return x.NameRka
	}
	return ""
}

func (x *RKA) GetTargetRka() int32 {
	if x != nil {
		return x.TargetRka
	}
	return 0
}

func (x *RKA) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RKA) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *RKA) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type UserRKA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RkaId     int64                  `protobuf:"varint,2,opt,name=rka_id,json=rkaId,proto3" json:"rka_id,omitempty"`
	UserId    int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Visited   int64                  `protobuf:"varint,4,opt,name=visited,proto3" json:"visited,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *UserRKA) Reset() {
	*x = UserRKA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRKA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRKA) ProtoMessage() {}

func (x *UserRKA) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRKA.ProtoReflect.Descriptor instead.
func (*UserRKA) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{2}
}

func (x *UserRKA) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRKA) GetRkaId() int64 {
	if x != nil {
		return x.RkaId
	}
	return 0
}

func (x *UserRKA) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRKA) GetVisited() int64 {
	if x != nil {
		return x.Visited
	}
	return 0
}

func (x *UserRKA) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserRKA) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserRKA) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ResponseRKA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*RKA `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseRKA) Reset() {
	*x = ResponseRKA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRKA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRKA) ProtoMessage() {}

func (x *ResponseRKA) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRKA.ProtoReflect.Descriptor instead.
func (*ResponseRKA) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseRKA) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseRKA) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseRKA) GetData() []*RKA {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *User  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseUser) Reset() {
	*x = ResponseUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseUser) ProtoMessage() {}

func (x *ResponseUser) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseUser.ProtoReflect.Descriptor instead.
func (*ResponseUser) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseUser) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseUser) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseUser) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseUserRKA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *User  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseUserRKA) Reset() {
	*x = ResponseUserRKA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseUserRKA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseUserRKA) ProtoMessage() {}

func (x *ResponseUserRKA) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseUserRKA.ProtoReflect.Descriptor instead.
func (*ResponseUserRKA) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseUserRKA) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseUserRKA) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseUserRKA) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseAccumulationRKA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code            int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message         string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ProgressVisited []byte `protobuf:"bytes,3,opt,name=progress_visited,json=progressVisited,proto3" json:"progress_visited,omitempty"`
}

func (x *ResponseAccumulationRKA) Reset() {
	*x = ResponseAccumulationRKA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAccumulationRKA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAccumulationRKA) ProtoMessage() {}

func (x *ResponseAccumulationRKA) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAccumulationRKA.ProtoReflect.Descriptor instead.
func (*ResponseAccumulationRKA) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseAccumulationRKA) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseAccumulationRKA) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseAccumulationRKA) GetProgressVisited() []byte {
	if x != nil {
		return x.ProgressVisited
	}
	return nil
}

type RequestAccumulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Yaer   int32  `protobuf:"varint,2,opt,name=yaer,proto3" json:"yaer,omitempty"`
	Month  string `protobuf:"bytes,3,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *RequestAccumulation) Reset() {
	*x = RequestAccumulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rka_rka_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAccumulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAccumulation) ProtoMessage() {}

func (x *RequestAccumulation) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rka_rka_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAccumulation.ProtoReflect.Descriptor instead.
func (*RequestAccumulation) Descriptor() ([]byte, []int) {
	return file_rpc_rka_rka_proto_rawDescGZIP(), []int{7}
}

func (x *RequestAccumulation) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RequestAccumulation) GetYaer() int32 {
	if x != nil {
		return x.Yaer
	}
	return 0
}

func (x *RequestAccumulation) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

var File_rpc_rka_rka_proto protoreflect.FileDescriptor

var file_rpc_rka_rka_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x6b, 0x61, 0x2f, 0x72, 0x6b, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x6b, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x03, 0x52, 0x4b, 0x41, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x6b, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x6b, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6b, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6b, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x94, 0x02, 0x0a, 0x07, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x4b, 0x41, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x6b, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x6b, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x74, 0x65, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x59, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x4b, 0x41,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x72,
	0x6b, 0x61, 0x2e, 0x52, 0x4b, 0x41, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x6b, 0x61, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5e, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x4b, 0x41, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x6b, 0x61, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x72, 0x0a, 0x17, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x4b, 0x41, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x76,
	0x69, 0x73, 0x69, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x69, 0x73, 0x69, 0x74, 0x65, 0x64, 0x22, 0x58, 0x0a,
	0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x61, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x61, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x32, 0xe0, 0x01, 0x0a, 0x0a, 0x52, 0x6b, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x09, 0x2e, 0x72, 0x6b, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x72,
	0x6b, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x26, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x4b, 0x41, 0x12, 0x08, 0x2e, 0x72, 0x6b,
	0x61, 0x2e, 0x52, 0x4b, 0x41, 0x1a, 0x10, 0x2e, 0x72, 0x6b, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x4b, 0x41, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x4b, 0x41, 0x12, 0x0c, 0x2e, 0x72, 0x6b, 0x61, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x4b, 0x41, 0x1a, 0x14, 0x2e, 0x72, 0x6b, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x4b, 0x41, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x0f, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x4b, 0x41,
	0x12, 0x18, 0x2e, 0x72, 0x6b, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x72, 0x6b, 0x61,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x4b, 0x41, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x72, 0x70,
	0x63, 0x2f, 0x72, 0x6b, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_rka_rka_proto_rawDescOnce sync.Once
	file_rpc_rka_rka_proto_rawDescData = file_rpc_rka_rka_proto_rawDesc
)

func file_rpc_rka_rka_proto_rawDescGZIP() []byte {
	file_rpc_rka_rka_proto_rawDescOnce.Do(func() {
		file_rpc_rka_rka_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_rka_rka_proto_rawDescData)
	})
	return file_rpc_rka_rka_proto_rawDescData
}

var file_rpc_rka_rka_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_rka_rka_proto_goTypes = []interface{}{
	(*User)(nil),                    // 0: rka.User
	(*RKA)(nil),                     // 1: rka.RKA
	(*UserRKA)(nil),                 // 2: rka.UserRKA
	(*ResponseRKA)(nil),             // 3: rka.ResponseRKA
	(*ResponseUser)(nil),            // 4: rka.ResponseUser
	(*ResponseUserRKA)(nil),         // 5: rka.ResponseUserRKA
	(*ResponseAccumulationRKA)(nil), // 6: rka.ResponseAccumulationRKA
	(*RequestAccumulation)(nil),     // 7: rka.RequestAccumulation
	(*timestamppb.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_rpc_rka_rka_proto_depIdxs = []int32{
	8,  // 0: rka.User.created_at:type_name -> google.protobuf.Timestamp
	8,  // 1: rka.User.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 2: rka.User.deleted_at:type_name -> google.protobuf.Timestamp
	8,  // 3: rka.RKA.created_at:type_name -> google.protobuf.Timestamp
	8,  // 4: rka.RKA.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 5: rka.RKA.deleted_at:type_name -> google.protobuf.Timestamp
	8,  // 6: rka.UserRKA.created_at:type_name -> google.protobuf.Timestamp
	8,  // 7: rka.UserRKA.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 8: rka.UserRKA.deleted_at:type_name -> google.protobuf.Timestamp
	1,  // 9: rka.ResponseRKA.data:type_name -> rka.RKA
	0,  // 10: rka.ResponseUser.data:type_name -> rka.User
	0,  // 11: rka.ResponseUserRKA.data:type_name -> rka.User
	0,  // 12: rka.RkaService.Adduser:input_type -> rka.User
	1,  // 13: rka.RkaService.AddRKA:input_type -> rka.RKA
	2,  // 14: rka.RkaService.AddUserRKA:input_type -> rka.UserRKA
	7,  // 15: rka.RkaService.AccumulationRKA:input_type -> rka.RequestAccumulation
	4,  // 16: rka.RkaService.Adduser:output_type -> rka.ResponseUser
	3,  // 17: rka.RkaService.AddRKA:output_type -> rka.ResponseRKA
	5,  // 18: rka.RkaService.AddUserRKA:output_type -> rka.ResponseUserRKA
	6,  // 19: rka.RkaService.AccumulationRKA:output_type -> rka.ResponseAccumulationRKA
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_rpc_rka_rka_proto_init() }
func file_rpc_rka_rka_proto_init() {
	if File_rpc_rka_rka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_rka_rka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_rpc_rka_rka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RKA); i {
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
		file_rpc_rka_rka_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRKA); i {
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
		file_rpc_rka_rka_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRKA); i {
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
		file_rpc_rka_rka_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseUser); i {
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
		file_rpc_rka_rka_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseUserRKA); i {
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
		file_rpc_rka_rka_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAccumulationRKA); i {
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
		file_rpc_rka_rka_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAccumulation); i {
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
			RawDescriptor: file_rpc_rka_rka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_rka_rka_proto_goTypes,
		DependencyIndexes: file_rpc_rka_rka_proto_depIdxs,
		MessageInfos:      file_rpc_rka_rka_proto_msgTypes,
	}.Build()
	File_rpc_rka_rka_proto = out.File
	file_rpc_rka_rka_proto_rawDesc = nil
	file_rpc_rka_rka_proto_goTypes = nil
	file_rpc_rka_rka_proto_depIdxs = nil
}