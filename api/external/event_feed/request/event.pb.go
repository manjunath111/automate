// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/event_feed/request/event.proto

package request

import (
	query "github.com/chef/automate/api/external/common/query"
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

type GetEventTypeCountsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Filters to be applied to the request.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest events to return.
	Start int64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// Latest events to return.
	End           int64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventTypeCountsRequest) Reset() {
	*x = GetEventTypeCountsRequest{}
	mi := &file_external_event_feed_request_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventTypeCountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventTypeCountsRequest) ProtoMessage() {}

func (x *GetEventTypeCountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_request_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventTypeCountsRequest.ProtoReflect.Descriptor instead.
func (*GetEventTypeCountsRequest) Descriptor() ([]byte, []int) {
	return file_external_event_feed_request_event_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventTypeCountsRequest) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetEventTypeCountsRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetEventTypeCountsRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type GetEventFeedRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Filters to be applied to the request.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest events to return.
	Start int64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// Latest events to return.
	End int64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	// Count of events to return per page.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Used for pagination, to request results in ascending order.
	After int64 `protobuf:"varint,5,opt,name=after,proto3" json:"after,omitempty"`
	// Used for pagination, to request results in descending order.
	Before int64 `protobuf:"varint,6,opt,name=before,proto3" json:"before,omitempty"`
	// Used for pagination, to bookmark next set of results.
	Cursor string `protobuf:"bytes,7,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// Used to group similar events together.
	Collapse      bool `protobuf:"varint,8,opt,name=collapse,proto3" json:"collapse,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventFeedRequest) Reset() {
	*x = GetEventFeedRequest{}
	mi := &file_external_event_feed_request_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventFeedRequest) ProtoMessage() {}

func (x *GetEventFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_request_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventFeedRequest.ProtoReflect.Descriptor instead.
func (*GetEventFeedRequest) Descriptor() ([]byte, []int) {
	return file_external_event_feed_request_event_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventFeedRequest) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetEventFeedRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetEventFeedRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *GetEventFeedRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetEventFeedRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *GetEventFeedRequest) GetBefore() int64 {
	if x != nil {
		return x.Before
	}
	return 0
}

func (x *GetEventFeedRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *GetEventFeedRequest) GetCollapse() bool {
	if x != nil {
		return x.Collapse
	}
	return false
}

type GetEventTaskCountsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Filters to be applied to the request.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest events to return.
	Start int64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	// Latest events to return.
	End           int64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventTaskCountsRequest) Reset() {
	*x = GetEventTaskCountsRequest{}
	mi := &file_external_event_feed_request_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventTaskCountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventTaskCountsRequest) ProtoMessage() {}

func (x *GetEventTaskCountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_request_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventTaskCountsRequest.ProtoReflect.Descriptor instead.
func (*GetEventTaskCountsRequest) Descriptor() ([]byte, []int) {
	return file_external_event_feed_request_event_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventTaskCountsRequest) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetEventTaskCountsRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetEventTaskCountsRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type EventExportRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// JSON or CSV
	OutputType string `protobuf:"bytes,1,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	// Filters to be applied to the request.
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest events to return.
	Start int64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	// Latest events to return.
	End int64 `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	// Order the results should be returned in.
	Order         query.SortOrder `protobuf:"varint,5,opt,name=order,proto3,enum=chef.automate.api.common.query.SortOrder" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventExportRequest) Reset() {
	*x = EventExportRequest{}
	mi := &file_external_event_feed_request_event_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventExportRequest) ProtoMessage() {}

func (x *EventExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_event_feed_request_event_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventExportRequest.ProtoReflect.Descriptor instead.
func (*EventExportRequest) Descriptor() ([]byte, []int) {
	return file_external_event_feed_request_event_proto_rawDescGZIP(), []int{3}
}

func (x *EventExportRequest) GetOutputType() string {
	if x != nil {
		return x.OutputType
	}
	return ""
}

func (x *EventExportRequest) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *EventExportRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *EventExportRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *EventExportRequest) GetOrder() query.SortOrder {
	if x != nil {
		return x.Order
	}
	return query.SortOrder(0)
}

var File_external_event_feed_request_event_proto protoreflect.FileDescriptor

var file_external_event_feed_request_event_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x22, 0xd4, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xb6, 0x01, 0x0a, 0x12, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x3f, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_event_feed_request_event_proto_rawDescOnce sync.Once
	file_external_event_feed_request_event_proto_rawDescData = file_external_event_feed_request_event_proto_rawDesc
)

func file_external_event_feed_request_event_proto_rawDescGZIP() []byte {
	file_external_event_feed_request_event_proto_rawDescOnce.Do(func() {
		file_external_event_feed_request_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_event_feed_request_event_proto_rawDescData)
	})
	return file_external_event_feed_request_event_proto_rawDescData
}

var file_external_event_feed_request_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_external_event_feed_request_event_proto_goTypes = []any{
	(*GetEventTypeCountsRequest)(nil), // 0: chef.automate.api.event_feed.request.GetEventTypeCountsRequest
	(*GetEventFeedRequest)(nil),       // 1: chef.automate.api.event_feed.request.GetEventFeedRequest
	(*GetEventTaskCountsRequest)(nil), // 2: chef.automate.api.event_feed.request.GetEventTaskCountsRequest
	(*EventExportRequest)(nil),        // 3: chef.automate.api.event_feed.request.EventExportRequest
	(query.SortOrder)(0),              // 4: chef.automate.api.common.query.SortOrder
}
var file_external_event_feed_request_event_proto_depIdxs = []int32{
	4, // 0: chef.automate.api.event_feed.request.EventExportRequest.order:type_name -> chef.automate.api.common.query.SortOrder
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_external_event_feed_request_event_proto_init() }
func file_external_event_feed_request_event_proto_init() {
	if File_external_event_feed_request_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_event_feed_request_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_event_feed_request_event_proto_goTypes,
		DependencyIndexes: file_external_event_feed_request_event_proto_depIdxs,
		MessageInfos:      file_external_event_feed_request_event_proto_msgTypes,
	}.Build()
	File_external_event_feed_request_event_proto = out.File
	file_external_event_feed_request_event_proto_rawDesc = nil
	file_external_event_feed_request_event_proto_goTypes = nil
	file_external_event_feed_request_event_proto_depIdxs = nil
}
