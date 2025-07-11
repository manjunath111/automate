// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.0
// source: external/iam/v2/request/introspect.proto

package request

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

type IntrospectAllReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntrospectAllReq) Reset() {
	*x = IntrospectAllReq{}
	mi := &file_external_iam_v2_request_introspect_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectAllReq) ProtoMessage() {}

func (x *IntrospectAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_introspect_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectAllReq.ProtoReflect.Descriptor instead.
func (*IntrospectAllReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_introspect_proto_rawDescGZIP(), []int{0}
}

type IntrospectSomeReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Paths         []string               `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntrospectSomeReq) Reset() {
	*x = IntrospectSomeReq{}
	mi := &file_external_iam_v2_request_introspect_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectSomeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectSomeReq) ProtoMessage() {}

func (x *IntrospectSomeReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_introspect_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectSomeReq.ProtoReflect.Descriptor instead.
func (*IntrospectSomeReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_introspect_proto_rawDescGZIP(), []int{1}
}

func (x *IntrospectSomeReq) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

type IntrospectReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Path          string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Parameters    []string               `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntrospectReq) Reset() {
	*x = IntrospectReq{}
	mi := &file_external_iam_v2_request_introspect_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectReq) ProtoMessage() {}

func (x *IntrospectReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_introspect_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectReq.ProtoReflect.Descriptor instead.
func (*IntrospectReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_introspect_proto_rawDescGZIP(), []int{2}
}

func (x *IntrospectReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *IntrospectReq) GetParameters() []string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_external_iam_v2_request_introspect_proto protoreflect.FileDescriptor

const file_external_iam_v2_request_introspect_proto_rawDesc = "" +
	"\n" +
	"(external/iam/v2/request/introspect.proto\x12\x18chef.automate.api.iam.v2\"\x12\n" +
	"\x10IntrospectAllReq\")\n" +
	"\x11IntrospectSomeReq\x12\x14\n" +
	"\x05paths\x18\x01 \x03(\tR\x05paths\"C\n" +
	"\rIntrospectReq\x12\x12\n" +
	"\x04path\x18\x01 \x01(\tR\x04path\x12\x1e\n" +
	"\n" +
	"parameters\x18\x02 \x03(\tR\n" +
	"parametersB6Z4github.com/chef/automate/api/external/iam/v2/requestb\x06proto3"

var (
	file_external_iam_v2_request_introspect_proto_rawDescOnce sync.Once
	file_external_iam_v2_request_introspect_proto_rawDescData []byte
)

func file_external_iam_v2_request_introspect_proto_rawDescGZIP() []byte {
	file_external_iam_v2_request_introspect_proto_rawDescOnce.Do(func() {
		file_external_iam_v2_request_introspect_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_external_iam_v2_request_introspect_proto_rawDesc), len(file_external_iam_v2_request_introspect_proto_rawDesc)))
	})
	return file_external_iam_v2_request_introspect_proto_rawDescData
}

var file_external_iam_v2_request_introspect_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_external_iam_v2_request_introspect_proto_goTypes = []any{
	(*IntrospectAllReq)(nil),  // 0: chef.automate.api.iam.v2.IntrospectAllReq
	(*IntrospectSomeReq)(nil), // 1: chef.automate.api.iam.v2.IntrospectSomeReq
	(*IntrospectReq)(nil),     // 2: chef.automate.api.iam.v2.IntrospectReq
}
var file_external_iam_v2_request_introspect_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_iam_v2_request_introspect_proto_init() }
func file_external_iam_v2_request_introspect_proto_init() {
	if File_external_iam_v2_request_introspect_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_external_iam_v2_request_introspect_proto_rawDesc), len(file_external_iam_v2_request_introspect_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_iam_v2_request_introspect_proto_goTypes,
		DependencyIndexes: file_external_iam_v2_request_introspect_proto_depIdxs,
		MessageInfos:      file_external_iam_v2_request_introspect_proto_msgTypes,
	}.Build()
	File_external_iam_v2_request_introspect_proto = out.File
	file_external_iam_v2_request_introspect_proto_goTypes = nil
	file_external_iam_v2_request_introspect_proto_depIdxs = nil
}
