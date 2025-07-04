// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.0
// source: interservice/event/event.proto

package event

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type EventType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventType) Reset() {
	*x = EventType{}
	mi := &file_interservice_event_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventType) ProtoMessage() {}

func (x *EventType) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventType.ProtoReflect.Descriptor instead.
func (*EventType) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Producer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	ProducerName  string                 `protobuf:"bytes,2,opt,name=producer_name,json=producerName,proto3" json:"producer_name,omitempty" toml:"producer_name,omitempty" mapstructure:"producer_name,omitempty"`
	ProducerType  string                 `protobuf:"bytes,3,opt,name=producer_type,json=producerType,proto3" json:"producer_type,omitempty" toml:"producer_type,omitempty" mapstructure:"producer_type,omitempty"`
	Tags          []string               `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" toml:"tags,omitempty" mapstructure:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Producer) Reset() {
	*x = Producer{}
	mi := &file_interservice_event_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Producer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Producer) ProtoMessage() {}

func (x *Producer) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Producer.ProtoReflect.Descriptor instead.
func (*Producer) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *Producer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Producer) GetProducerName() string {
	if x != nil {
		return x.ProducerName
	}
	return ""
}

func (x *Producer) GetProducerType() string {
	if x != nil {
		return x.ProducerType
	}
	return ""
}

func (x *Producer) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Actor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	ObjectType    string                 `protobuf:"bytes,2,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty" toml:"object_type,omitempty" mapstructure:"object_type,omitempty"`
	DisplayName   string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" toml:"display_name,omitempty" mapstructure:"display_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Actor) Reset() {
	*x = Actor{}
	mi := &file_interservice_event_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Actor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actor) ProtoMessage() {}

func (x *Actor) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actor.ProtoReflect.Descriptor instead.
func (*Actor) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *Actor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Actor) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *Actor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type Object struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	ObjectType    string                 `protobuf:"bytes,2,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty" toml:"object_type,omitempty" mapstructure:"object_type,omitempty"`
	DisplayName   string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" toml:"display_name,omitempty" mapstructure:"display_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Object) Reset() {
	*x = Object{}
	mi := &file_interservice_event_event_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{3}
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Object) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *Object) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type Target struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	ObjectType    string                 `protobuf:"bytes,2,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty" toml:"object_type,omitempty" mapstructure:"object_type,omitempty"`
	DisplayName   string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" toml:"display_name,omitempty" mapstructure:"display_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Target) Reset() {
	*x = Target{}
	mi := &file_interservice_event_event_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{4}
}

func (x *Target) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Target) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *Target) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type EventMsg struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventId       string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty" toml:"event_id,omitempty" mapstructure:"event_id,omitempty"`
	Type          *EventType             `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty" mapstructure:"type,omitempty"`
	Producer      *Producer              `protobuf:"bytes,3,opt,name=producer,proto3" json:"producer,omitempty" toml:"producer,omitempty" mapstructure:"producer,omitempty"`
	Tags          []string               `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" toml:"tags,omitempty" mapstructure:"tags,omitempty"`
	Published     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=published,proto3" json:"published,omitempty" toml:"published,omitempty" mapstructure:"published,omitempty"`
	Actor         *Actor                 `protobuf:"bytes,6,opt,name=actor,proto3" json:"actor,omitempty" toml:"actor,omitempty" mapstructure:"actor,omitempty"`
	Verb          string                 `protobuf:"bytes,7,opt,name=verb,proto3" json:"verb,omitempty" toml:"verb,omitempty" mapstructure:"verb,omitempty"`
	Object        *Object                `protobuf:"bytes,8,opt,name=object,proto3" json:"object,omitempty" toml:"object,omitempty" mapstructure:"object,omitempty"`
	Target        *Target                `protobuf:"bytes,9,opt,name=target,proto3" json:"target,omitempty" toml:"target,omitempty" mapstructure:"target,omitempty"`
	Data          *structpb.Struct       `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty" toml:"data,omitempty" mapstructure:"data,omitempty"`
	Projects      []string               `protobuf:"bytes,11,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventMsg) Reset() {
	*x = EventMsg{}
	mi := &file_interservice_event_event_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMsg) ProtoMessage() {}

func (x *EventMsg) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMsg.ProtoReflect.Descriptor instead.
func (*EventMsg) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{5}
}

func (x *EventMsg) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventMsg) GetType() *EventType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *EventMsg) GetProducer() *Producer {
	if x != nil {
		return x.Producer
	}
	return nil
}

func (x *EventMsg) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *EventMsg) GetPublished() *timestamppb.Timestamp {
	if x != nil {
		return x.Published
	}
	return nil
}

func (x *EventMsg) GetActor() *Actor {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *EventMsg) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *EventMsg) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *EventMsg) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *EventMsg) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EventMsg) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type EventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty" toml:"success,omitempty" mapstructure:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	mi := &file_interservice_event_event_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{6}
}

func (x *EventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PublishRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           *EventMsg              `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty" toml:"msg,omitempty" mapstructure:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	mi := &file_interservice_event_event_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{7}
}

func (x *PublishRequest) GetMsg() *EventMsg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type PublishResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty" toml:"success,omitempty" mapstructure:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_interservice_event_event_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{8}
}

func (x *PublishResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_interservice_event_event_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{9}
}

type SubscribeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	mi := &file_interservice_event_event_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{10}
}

type StartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	mi := &file_interservice_event_event_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{11}
}

type StartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	mi := &file_interservice_event_event_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{12}
}

type StopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	mi := &file_interservice_event_event_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{13}
}

type StopResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	mi := &file_interservice_event_event_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_event_event_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_interservice_event_event_proto_rawDescGZIP(), []int{14}
}

var File_interservice_event_event_proto protoreflect.FileDescriptor

const file_interservice_event_event_proto_rawDesc = "" +
	"\n" +
	"\x1einterservice/event/event.proto\x12\x1echef.automate.domain.event.api\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/protobuf/struct.proto\"\x1f\n" +
	"\tEventType\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"x\n" +
	"\bProducer\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12#\n" +
	"\rproducer_name\x18\x02 \x01(\tR\fproducerName\x12#\n" +
	"\rproducer_type\x18\x03 \x01(\tR\fproducerType\x12\x12\n" +
	"\x04tags\x18\x04 \x03(\tR\x04tags\"[\n" +
	"\x05Actor\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1f\n" +
	"\vobject_type\x18\x02 \x01(\tR\n" +
	"objectType\x12!\n" +
	"\fdisplay_name\x18\x03 \x01(\tR\vdisplayName\"\\\n" +
	"\x06Object\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1f\n" +
	"\vobject_type\x18\x02 \x01(\tR\n" +
	"objectType\x12!\n" +
	"\fdisplay_name\x18\x03 \x01(\tR\vdisplayName\"\\\n" +
	"\x06Target\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1f\n" +
	"\vobject_type\x18\x02 \x01(\tR\n" +
	"objectType\x12!\n" +
	"\fdisplay_name\x18\x03 \x01(\tR\vdisplayName\"\x92\x04\n" +
	"\bEventMsg\x12\x19\n" +
	"\bevent_id\x18\x01 \x01(\tR\aeventId\x12=\n" +
	"\x04type\x18\x02 \x01(\v2).chef.automate.domain.event.api.EventTypeR\x04type\x12D\n" +
	"\bproducer\x18\x03 \x01(\v2(.chef.automate.domain.event.api.ProducerR\bproducer\x12\x12\n" +
	"\x04tags\x18\x04 \x03(\tR\x04tags\x128\n" +
	"\tpublished\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tpublished\x12;\n" +
	"\x05actor\x18\x06 \x01(\v2%.chef.automate.domain.event.api.ActorR\x05actor\x12\x12\n" +
	"\x04verb\x18\a \x01(\tR\x04verb\x12>\n" +
	"\x06object\x18\b \x01(\v2&.chef.automate.domain.event.api.ObjectR\x06object\x12>\n" +
	"\x06target\x18\t \x01(\v2&.chef.automate.domain.event.api.TargetR\x06target\x12+\n" +
	"\x04data\x18\n" +
	" \x01(\v2\x17.google.protobuf.StructR\x04data\x12\x1a\n" +
	"\bprojects\x18\v \x03(\tR\bprojects\")\n" +
	"\rEventResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"L\n" +
	"\x0ePublishRequest\x12:\n" +
	"\x03msg\x18\x01 \x01(\v2(.chef.automate.domain.event.api.EventMsgR\x03msg\"+\n" +
	"\x0fPublishResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\x12\n" +
	"\x10SubscribeRequest\"\x13\n" +
	"\x11SubscribeResponse\"\x0e\n" +
	"\fStartRequest\"\x0f\n" +
	"\rStartResponse\"\r\n" +
	"\vStopRequest\"\x0e\n" +
	"\fStopResponse2\xb5\x03\n" +
	"\fEventService\x12j\n" +
	"\aPublish\x12..chef.automate.domain.event.api.PublishRequest\x1a/.chef.automate.domain.event.api.PublishResponse\x12p\n" +
	"\tSubscribe\x120.chef.automate.domain.event.api.SubscribeRequest\x1a1.chef.automate.domain.event.api.SubscribeResponse\x12d\n" +
	"\x05Start\x12,.chef.automate.domain.event.api.StartRequest\x1a-.chef.automate.domain.event.api.StartResponse\x12a\n" +
	"\x04Stop\x12+.chef.automate.domain.event.api.StopRequest\x1a,.chef.automate.domain.event.api.StopResponseB1Z/github.com/chef/automate/api/interservice/eventb\x06proto3"

var (
	file_interservice_event_event_proto_rawDescOnce sync.Once
	file_interservice_event_event_proto_rawDescData []byte
)

func file_interservice_event_event_proto_rawDescGZIP() []byte {
	file_interservice_event_event_proto_rawDescOnce.Do(func() {
		file_interservice_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_interservice_event_event_proto_rawDesc), len(file_interservice_event_event_proto_rawDesc)))
	})
	return file_interservice_event_event_proto_rawDescData
}

var file_interservice_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_interservice_event_event_proto_goTypes = []any{
	(*EventType)(nil),             // 0: chef.automate.domain.event.api.EventType
	(*Producer)(nil),              // 1: chef.automate.domain.event.api.Producer
	(*Actor)(nil),                 // 2: chef.automate.domain.event.api.Actor
	(*Object)(nil),                // 3: chef.automate.domain.event.api.Object
	(*Target)(nil),                // 4: chef.automate.domain.event.api.Target
	(*EventMsg)(nil),              // 5: chef.automate.domain.event.api.EventMsg
	(*EventResponse)(nil),         // 6: chef.automate.domain.event.api.EventResponse
	(*PublishRequest)(nil),        // 7: chef.automate.domain.event.api.PublishRequest
	(*PublishResponse)(nil),       // 8: chef.automate.domain.event.api.PublishResponse
	(*SubscribeRequest)(nil),      // 9: chef.automate.domain.event.api.SubscribeRequest
	(*SubscribeResponse)(nil),     // 10: chef.automate.domain.event.api.SubscribeResponse
	(*StartRequest)(nil),          // 11: chef.automate.domain.event.api.StartRequest
	(*StartResponse)(nil),         // 12: chef.automate.domain.event.api.StartResponse
	(*StopRequest)(nil),           // 13: chef.automate.domain.event.api.StopRequest
	(*StopResponse)(nil),          // 14: chef.automate.domain.event.api.StopResponse
	(*timestamppb.Timestamp)(nil), // 15: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 16: google.protobuf.Struct
}
var file_interservice_event_event_proto_depIdxs = []int32{
	0,  // 0: chef.automate.domain.event.api.EventMsg.type:type_name -> chef.automate.domain.event.api.EventType
	1,  // 1: chef.automate.domain.event.api.EventMsg.producer:type_name -> chef.automate.domain.event.api.Producer
	15, // 2: chef.automate.domain.event.api.EventMsg.published:type_name -> google.protobuf.Timestamp
	2,  // 3: chef.automate.domain.event.api.EventMsg.actor:type_name -> chef.automate.domain.event.api.Actor
	3,  // 4: chef.automate.domain.event.api.EventMsg.object:type_name -> chef.automate.domain.event.api.Object
	4,  // 5: chef.automate.domain.event.api.EventMsg.target:type_name -> chef.automate.domain.event.api.Target
	16, // 6: chef.automate.domain.event.api.EventMsg.data:type_name -> google.protobuf.Struct
	5,  // 7: chef.automate.domain.event.api.PublishRequest.msg:type_name -> chef.automate.domain.event.api.EventMsg
	7,  // 8: chef.automate.domain.event.api.EventService.Publish:input_type -> chef.automate.domain.event.api.PublishRequest
	9,  // 9: chef.automate.domain.event.api.EventService.Subscribe:input_type -> chef.automate.domain.event.api.SubscribeRequest
	11, // 10: chef.automate.domain.event.api.EventService.Start:input_type -> chef.automate.domain.event.api.StartRequest
	13, // 11: chef.automate.domain.event.api.EventService.Stop:input_type -> chef.automate.domain.event.api.StopRequest
	8,  // 12: chef.automate.domain.event.api.EventService.Publish:output_type -> chef.automate.domain.event.api.PublishResponse
	10, // 13: chef.automate.domain.event.api.EventService.Subscribe:output_type -> chef.automate.domain.event.api.SubscribeResponse
	12, // 14: chef.automate.domain.event.api.EventService.Start:output_type -> chef.automate.domain.event.api.StartResponse
	14, // 15: chef.automate.domain.event.api.EventService.Stop:output_type -> chef.automate.domain.event.api.StopResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_interservice_event_event_proto_init() }
func file_interservice_event_event_proto_init() {
	if File_interservice_event_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_interservice_event_event_proto_rawDesc), len(file_interservice_event_event_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_event_event_proto_goTypes,
		DependencyIndexes: file_interservice_event_event_proto_depIdxs,
		MessageInfos:      file_interservice_event_event_proto_msgTypes,
	}.Build()
	File_interservice_event_event_proto = out.File
	file_interservice_event_event_proto_goTypes = nil
	file_interservice_event_event_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.event.api.EventService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedEventServiceServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedEventServiceServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedEventServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.event.api.EventService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.event.api.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _EventService_Publish_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _EventService_Subscribe_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _EventService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _EventService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/event/event.proto",
}
