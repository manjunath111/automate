// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/iam/v2/users.proto

package v2

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	request "github.com/chef/automate/api/external/iam/v2/request"
	response "github.com/chef/automate/api/external/iam/v2/response"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_external_iam_v2_users_proto protoreflect.FileDescriptor

var file_external_iam_v2_users_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x0b, 0x0a, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0xb3, 0x02, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0xd1, 0x01, 0x92, 0x41, 0x8f, 0x01, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x6a, 0x85, 0x01, 0x0a, 0x0e, 0x78, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2d,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x73, 0x32, 0x71, 0x0a, 0x6f, 0x2a, 0x6d, 0x0a,
	0x0e, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x06, 0x1a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x0a,
	0x5b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x51, 0x1a, 0x4f, 0x7b, 0x22, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x20, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x30, 0x30, 0x31, 0x72, 0x75, 0x6c, 0x65, 0x7a, 0x22, 0x2c, 0x20,
	0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x53, 0x61,
	0x66, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x7d, 0x8a, 0xb5, 0x18, 0x1d,
	0x0a, 0x09, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x10, 0x69, 0x61, 0x6d,
	0x3a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0xa1, 0x01, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x43, 0x92, 0x41, 0x07, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x8a, 0xb5, 0x18, 0x1b, 0x0a, 0x09, 0x69, 0x61, 0x6d, 0x3a, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a,
	0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0xa4,
	0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x4c, 0x92, 0x41, 0x07, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x8a, 0xb5, 0x18, 0x1f, 0x0a, 0x0e, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x3a, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x0d, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xb0, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x4f, 0x92, 0x41, 0x07, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x8a, 0xb5, 0x18, 0x22, 0x0a, 0x0e, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x3a, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x10, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x3a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9d, 0x02, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0xbb, 0x01, 0x92, 0x41, 0x70,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x6a, 0x67, 0x0a, 0x0e, 0x78, 0x2d, 0x63, 0x6f, 0x64,
	0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x55, 0x32, 0x53, 0x0a, 0x51, 0x2a,
	0x4f, 0x0a, 0x0e, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x06, 0x1a, 0x04, 0x4a, 0x53, 0x4f,
	0x4e, 0x0a, 0x3d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x33, 0x1a, 0x31, 0x7b,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4e, 0x65, 0x77, 0x20, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x20,
	0x22, 0x61, 0x53, 0x61, 0x66, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x7d,
	0x8a, 0xb5, 0x18, 0x22, 0x0a, 0x0e, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x10, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x1a,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xde, 0x02, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x66, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71,
	0x1a, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x22, 0xfc, 0x01, 0x92, 0x41, 0xa9,
	0x01, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x6a, 0x9f, 0x01, 0x0a, 0x0e, 0x78, 0x2d, 0x63,
	0x6f, 0x64, 0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x8c, 0x01, 0x32, 0x89,
	0x01, 0x0a, 0x86, 0x01, 0x2a, 0x83, 0x01, 0x0a, 0x0e, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x06, 0x1a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x0a, 0x71, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x67, 0x1a, 0x65, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d,
	0x79, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x4e, 0x65, 0x77, 0x53, 0x61,
	0x66, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x3a, 0x20, 0x22, 0x61, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x53, 0x61, 0x66, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x7d, 0x8a, 0xb5, 0x18, 0x2a, 0x0a, 0x12,
	0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x6c, 0x66, 0x3a, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x14, 0x69, 0x61, 0x6d, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x6c,
	0x66, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01,
	0x2a, 0x1a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f,
	0x73, 0x65, 0x6c, 0x66, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_external_iam_v2_users_proto_goTypes = []any{
	(*request.CreateUserReq)(nil),   // 0: chef.automate.api.iam.v2.CreateUserReq
	(*request.ListUsersReq)(nil),    // 1: chef.automate.api.iam.v2.ListUsersReq
	(*request.GetUserReq)(nil),      // 2: chef.automate.api.iam.v2.GetUserReq
	(*request.DeleteUserReq)(nil),   // 3: chef.automate.api.iam.v2.DeleteUserReq
	(*request.UpdateUserReq)(nil),   // 4: chef.automate.api.iam.v2.UpdateUserReq
	(*request.UpdateSelfReq)(nil),   // 5: chef.automate.api.iam.v2.UpdateSelfReq
	(*response.CreateUserResp)(nil), // 6: chef.automate.api.iam.v2.CreateUserResp
	(*response.ListUsersResp)(nil),  // 7: chef.automate.api.iam.v2.ListUsersResp
	(*response.GetUserResp)(nil),    // 8: chef.automate.api.iam.v2.GetUserResp
	(*response.DeleteUserResp)(nil), // 9: chef.automate.api.iam.v2.DeleteUserResp
	(*response.UpdateUserResp)(nil), // 10: chef.automate.api.iam.v2.UpdateUserResp
	(*response.UpdateSelfResp)(nil), // 11: chef.automate.api.iam.v2.UpdateSelfResp
}
var file_external_iam_v2_users_proto_depIdxs = []int32{
	0,  // 0: chef.automate.api.iam.v2.Users.CreateUser:input_type -> chef.automate.api.iam.v2.CreateUserReq
	1,  // 1: chef.automate.api.iam.v2.Users.ListUsers:input_type -> chef.automate.api.iam.v2.ListUsersReq
	2,  // 2: chef.automate.api.iam.v2.Users.GetUser:input_type -> chef.automate.api.iam.v2.GetUserReq
	3,  // 3: chef.automate.api.iam.v2.Users.DeleteUser:input_type -> chef.automate.api.iam.v2.DeleteUserReq
	4,  // 4: chef.automate.api.iam.v2.Users.UpdateUser:input_type -> chef.automate.api.iam.v2.UpdateUserReq
	5,  // 5: chef.automate.api.iam.v2.Users.UpdateSelf:input_type -> chef.automate.api.iam.v2.UpdateSelfReq
	6,  // 6: chef.automate.api.iam.v2.Users.CreateUser:output_type -> chef.automate.api.iam.v2.CreateUserResp
	7,  // 7: chef.automate.api.iam.v2.Users.ListUsers:output_type -> chef.automate.api.iam.v2.ListUsersResp
	8,  // 8: chef.automate.api.iam.v2.Users.GetUser:output_type -> chef.automate.api.iam.v2.GetUserResp
	9,  // 9: chef.automate.api.iam.v2.Users.DeleteUser:output_type -> chef.automate.api.iam.v2.DeleteUserResp
	10, // 10: chef.automate.api.iam.v2.Users.UpdateUser:output_type -> chef.automate.api.iam.v2.UpdateUserResp
	11, // 11: chef.automate.api.iam.v2.Users.UpdateSelf:output_type -> chef.automate.api.iam.v2.UpdateSelfResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_external_iam_v2_users_proto_init() }
func file_external_iam_v2_users_proto_init() {
	if File_external_iam_v2_users_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_iam_v2_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_iam_v2_users_proto_goTypes,
		DependencyIndexes: file_external_iam_v2_users_proto_depIdxs,
	}.Build()
	File_external_iam_v2_users_proto = out.File
	file_external_iam_v2_users_proto_rawDesc = nil
	file_external_iam_v2_users_proto_goTypes = nil
	file_external_iam_v2_users_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	// Create a user
	//
	// Creates a local user that can sign in to Automate and be a member of IAM teams
	// and IAM policies.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:create
	CreateUser(ctx context.Context, in *request.CreateUserReq, opts ...grpc.CallOption) (*response.CreateUserResp, error)
	// List all users
	//
	// Lists all local users.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:list
	ListUsers(ctx context.Context, in *request.ListUsersReq, opts ...grpc.CallOption) (*response.ListUsersResp, error)
	// Get a user
	//
	// Returns the details for a local user.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:get
	GetUser(ctx context.Context, in *request.GetUserReq, opts ...grpc.CallOption) (*response.GetUserResp, error)
	// Delete a user
	//
	// Deletes a local user.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:delete
	DeleteUser(ctx context.Context, in *request.DeleteUserReq, opts ...grpc.CallOption) (*response.DeleteUserResp, error)
	// Update a user
	//
	// Updates a local user's name or password.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:update
	UpdateUser(ctx context.Context, in *request.UpdateUserReq, opts ...grpc.CallOption) (*response.UpdateUserResp, error)
	// Update self (as user)
	//
	// Updates a local user's own name or password.
	// If changing the password, both "password" and "previous_password" are required.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:usersSelf:update
	UpdateSelf(ctx context.Context, in *request.UpdateSelfReq, opts ...grpc.CallOption) (*response.UpdateSelfResp, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUser(ctx context.Context, in *request.CreateUserReq, opts ...grpc.CallOption) (*response.CreateUserResp, error) {
	out := new(response.CreateUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListUsers(ctx context.Context, in *request.ListUsersReq, opts ...grpc.CallOption) (*response.ListUsersResp, error) {
	out := new(response.ListUsersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUser(ctx context.Context, in *request.GetUserReq, opts ...grpc.CallOption) (*response.GetUserResp, error) {
	out := new(response.GetUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUser(ctx context.Context, in *request.DeleteUserReq, opts ...grpc.CallOption) (*response.DeleteUserResp, error) {
	out := new(response.DeleteUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUser(ctx context.Context, in *request.UpdateUserReq, opts ...grpc.CallOption) (*response.UpdateUserResp, error) {
	out := new(response.UpdateUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateSelf(ctx context.Context, in *request.UpdateSelfReq, opts ...grpc.CallOption) (*response.UpdateSelfResp, error) {
	out := new(response.UpdateSelfResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/UpdateSelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	// Create a user
	//
	// Creates a local user that can sign in to Automate and be a member of IAM teams
	// and IAM policies.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:create
	CreateUser(context.Context, *request.CreateUserReq) (*response.CreateUserResp, error)
	// List all users
	//
	// Lists all local users.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:list
	ListUsers(context.Context, *request.ListUsersReq) (*response.ListUsersResp, error)
	// Get a user
	//
	// Returns the details for a local user.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:get
	GetUser(context.Context, *request.GetUserReq) (*response.GetUserResp, error)
	// Delete a user
	//
	// Deletes a local user.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:delete
	DeleteUser(context.Context, *request.DeleteUserReq) (*response.DeleteUserResp, error)
	// Update a user
	//
	// Updates a local user's name or password.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:users:update
	UpdateUser(context.Context, *request.UpdateUserReq) (*response.UpdateUserResp, error)
	// Update self (as user)
	//
	// Updates a local user's own name or password.
	// If changing the password, both "password" and "previous_password" are required.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//iam:usersSelf:update
	UpdateSelf(context.Context, *request.UpdateSelfReq) (*response.UpdateSelfResp, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) CreateUser(context.Context, *request.CreateUserReq) (*response.CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUsersServer) ListUsers(context.Context, *request.ListUsersReq) (*response.ListUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedUsersServer) GetUser(context.Context, *request.GetUserReq) (*response.GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUsersServer) DeleteUser(context.Context, *request.DeleteUserReq) (*response.DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedUsersServer) UpdateUser(context.Context, *request.UpdateUserReq) (*response.UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUsersServer) UpdateSelf(context.Context, *request.UpdateSelfReq) (*response.UpdateSelfResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSelf not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*request.CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListUsers(ctx, req.(*request.ListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*request.GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUser(ctx, req.(*request.DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*request.UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateSelfReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/UpdateSelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateSelf(ctx, req.(*request.UpdateSelfReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Users_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Users_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateSelf",
			Handler:    _Users_UpdateSelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/iam/v2/users.proto",
}
