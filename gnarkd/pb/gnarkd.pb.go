// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pb/gnarkd.proto

package pb

import (
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

type ProveJobResult_Status int32

const (
	ProveJobResult_WAITING_WITNESS ProveJobResult_Status = 0
	ProveJobResult_QUEUED          ProveJobResult_Status = 1
	ProveJobResult_RUNNING         ProveJobResult_Status = 2
	ProveJobResult_COMPLETED       ProveJobResult_Status = 3
	ProveJobResult_ERRORED         ProveJobResult_Status = 4
)

// Enum value maps for ProveJobResult_Status.
var (
	ProveJobResult_Status_name = map[int32]string{
		0: "WAITING_WITNESS",
		1: "QUEUED",
		2: "RUNNING",
		3: "COMPLETED",
		4: "ERRORED",
	}
	ProveJobResult_Status_value = map[string]int32{
		"WAITING_WITNESS": 0,
		"QUEUED":          1,
		"RUNNING":         2,
		"COMPLETED":       3,
		"ERRORED":         4,
	}
)

func (x ProveJobResult_Status) Enum() *ProveJobResult_Status {
	p := new(ProveJobResult_Status)
	*p = x
	return p
}

func (x ProveJobResult_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProveJobResult_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_gnarkd_proto_enumTypes[0].Descriptor()
}

func (ProveJobResult_Status) Type() protoreflect.EnumType {
	return &file_pb_gnarkd_proto_enumTypes[0]
}

func (x ProveJobResult_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProveJobResult_Status.Descriptor instead.
func (ProveJobResult_Status) EnumDescriptor() ([]byte, []int) {
	return file_pb_gnarkd_proto_rawDescGZIP(), []int{4, 0}
}

type ProveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CircuitID string `protobuf:"bytes,1,opt,name=circuitID,proto3" json:"circuitID,omitempty"`
	Witness   []byte `protobuf:"bytes,2,opt,name=witness,proto3" json:"witness,omitempty"`
}

func (x *ProveRequest) Reset() {
	*x = ProveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gnarkd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProveRequest) ProtoMessage() {}

func (x *ProveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gnarkd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProveRequest.ProtoReflect.Descriptor instead.
func (*ProveRequest) Descriptor() ([]byte, []int) {
	return file_pb_gnarkd_proto_rawDescGZIP(), []int{0}
}

func (x *ProveRequest) GetCircuitID() string {
	if x != nil {
		return x.CircuitID
	}
	return ""
}

func (x *ProveRequest) GetWitness() []byte {
	if x != nil {
		return x.Witness
	}
	return nil
}

type ProveResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *ProveResult) Reset() {
	*x = ProveResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gnarkd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProveResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProveResult) ProtoMessage() {}

func (x *ProveResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gnarkd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProveResult.ProtoReflect.Descriptor instead.
func (*ProveResult) Descriptor() ([]byte, []int) {
	return file_pb_gnarkd_proto_rawDescGZIP(), []int{1}
}

func (x *ProveResult) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type CreateProveJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CircuitID string `protobuf:"bytes,1,opt,name=circuitID,proto3" json:"circuitID,omitempty"`
}

func (x *CreateProveJobRequest) Reset() {
	*x = CreateProveJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gnarkd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProveJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProveJobRequest) ProtoMessage() {}

func (x *CreateProveJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gnarkd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProveJobRequest.ProtoReflect.Descriptor instead.
func (*CreateProveJobRequest) Descriptor() ([]byte, []int) {
	return file_pb_gnarkd_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProveJobRequest) GetCircuitID() string {
	if x != nil {
		return x.CircuitID
	}
	return ""
}

type CreateProveJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
}

func (x *CreateProveJobResponse) Reset() {
	*x = CreateProveJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gnarkd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProveJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProveJobResponse) ProtoMessage() {}

func (x *CreateProveJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gnarkd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProveJobResponse.ProtoReflect.Descriptor instead.
func (*CreateProveJobResponse) Descriptor() ([]byte, []int) {
	return file_pb_gnarkd_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProveJobResponse) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

type ProveJobResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID  string                `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Status ProveJobResult_Status `protobuf:"varint,2,opt,name=status,proto3,enum=gnarkd.ProveJobResult_Status" json:"status,omitempty"`
	Err    *string               `protobuf:"bytes,3,opt,name=err,proto3,oneof" json:"err,omitempty"`
	Proof  []byte                `protobuf:"bytes,4,opt,name=proof,proto3,oneof" json:"proof,omitempty"`
}

func (x *ProveJobResult) Reset() {
	*x = ProveJobResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gnarkd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProveJobResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProveJobResult) ProtoMessage() {}

func (x *ProveJobResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gnarkd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProveJobResult.ProtoReflect.Descriptor instead.
func (*ProveJobResult) Descriptor() ([]byte, []int) {
	return file_pb_gnarkd_proto_rawDescGZIP(), []int{4}
}

func (x *ProveJobResult) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *ProveJobResult) GetStatus() ProveJobResult_Status {
	if x != nil {
		return x.Status
	}
	return ProveJobResult_WAITING_WITNESS
}

func (x *ProveJobResult) GetErr() string {
	if x != nil && x.Err != nil {
		return *x.Err
	}
	return ""
}

func (x *ProveJobResult) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type SubscribeToProveJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
}

func (x *SubscribeToProveJobRequest) Reset() {
	*x = SubscribeToProveJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_gnarkd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToProveJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToProveJobRequest) ProtoMessage() {}

func (x *SubscribeToProveJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_gnarkd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToProveJobRequest.ProtoReflect.Descriptor instead.
func (*SubscribeToProveJobRequest) Descriptor() ([]byte, []int) {
	return file_pb_gnarkd_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeToProveJobRequest) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

var File_pb_gnarkd_proto protoreflect.FileDescriptor

var file_pb_gnarkd_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x64, 0x22, 0x46, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x69,
	0x72, 0x63, 0x75, 0x69, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x22, 0x23, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x35, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x49, 0x44, 0x22, 0x2e, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x22, 0xf5, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x64, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x88, 0x01, 0x01, 0x22,
	0x52, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x41, 0x49,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x57, 0x49, 0x54, 0x4e, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x45,
	0x44, 0x10, 0x04, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x65, 0x72, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x32, 0x0a, 0x1a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x32, 0xe3, 0x01, 0x0a, 0x07, 0x47, 0x72,
	0x6f, 0x74, 0x68, 0x31, 0x36, 0x12, 0x32, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x14,
	0x2e, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x64, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4f, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x67, 0x6e,
	0x61, 0x72, 0x6b, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6e, 0x61,
	0x72, 0x6b, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x13, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f,
	0x62, 0x12, 0x22, 0x2e, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x64, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x64, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x30, 0x01, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x79, 0x73, 0x2f, 0x67, 0x6e, 0x61, 0x72, 0x6b, 0x2f, 0x67, 0x6e,
	0x61, 0x72, 0x6b, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_gnarkd_proto_rawDescOnce sync.Once
	file_pb_gnarkd_proto_rawDescData = file_pb_gnarkd_proto_rawDesc
)

func file_pb_gnarkd_proto_rawDescGZIP() []byte {
	file_pb_gnarkd_proto_rawDescOnce.Do(func() {
		file_pb_gnarkd_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_gnarkd_proto_rawDescData)
	})
	return file_pb_gnarkd_proto_rawDescData
}

var file_pb_gnarkd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_gnarkd_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_gnarkd_proto_goTypes = []interface{}{
	(ProveJobResult_Status)(0),         // 0: gnarkd.ProveJobResult.Status
	(*ProveRequest)(nil),               // 1: gnarkd.ProveRequest
	(*ProveResult)(nil),                // 2: gnarkd.ProveResult
	(*CreateProveJobRequest)(nil),      // 3: gnarkd.CreateProveJobRequest
	(*CreateProveJobResponse)(nil),     // 4: gnarkd.CreateProveJobResponse
	(*ProveJobResult)(nil),             // 5: gnarkd.ProveJobResult
	(*SubscribeToProveJobRequest)(nil), // 6: gnarkd.SubscribeToProveJobRequest
}
var file_pb_gnarkd_proto_depIdxs = []int32{
	0, // 0: gnarkd.ProveJobResult.status:type_name -> gnarkd.ProveJobResult.Status
	1, // 1: gnarkd.Groth16.Prove:input_type -> gnarkd.ProveRequest
	3, // 2: gnarkd.Groth16.CreateProveJob:input_type -> gnarkd.CreateProveJobRequest
	6, // 3: gnarkd.Groth16.SubscribeToProveJob:input_type -> gnarkd.SubscribeToProveJobRequest
	2, // 4: gnarkd.Groth16.Prove:output_type -> gnarkd.ProveResult
	4, // 5: gnarkd.Groth16.CreateProveJob:output_type -> gnarkd.CreateProveJobResponse
	5, // 6: gnarkd.Groth16.SubscribeToProveJob:output_type -> gnarkd.ProveJobResult
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_gnarkd_proto_init() }
func file_pb_gnarkd_proto_init() {
	if File_pb_gnarkd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_gnarkd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProveRequest); i {
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
		file_pb_gnarkd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProveResult); i {
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
		file_pb_gnarkd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProveJobRequest); i {
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
		file_pb_gnarkd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProveJobResponse); i {
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
		file_pb_gnarkd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProveJobResult); i {
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
		file_pb_gnarkd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToProveJobRequest); i {
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
	file_pb_gnarkd_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_gnarkd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_gnarkd_proto_goTypes,
		DependencyIndexes: file_pb_gnarkd_proto_depIdxs,
		EnumInfos:         file_pb_gnarkd_proto_enumTypes,
		MessageInfos:      file_pb_gnarkd_proto_msgTypes,
	}.Build()
	File_pb_gnarkd_proto = out.File
	file_pb_gnarkd_proto_rawDesc = nil
	file_pb_gnarkd_proto_goTypes = nil
	file_pb_gnarkd_proto_depIdxs = nil
}