// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: eventsource.proto

package deposit

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EventParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId       string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventType     string `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	AggregateId   string `protobuf:"bytes,3,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateType string `protobuf:"bytes,4,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	EventData     string `protobuf:"bytes,5,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	Channel       string `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty"` // an optional field
}

func (x *EventParam) Reset() {
	*x = EventParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventsource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventParam) ProtoMessage() {}

func (x *EventParam) ProtoReflect() protoreflect.Message {
	mi := &file_eventsource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventParam.ProtoReflect.Descriptor instead.
func (*EventParam) Descriptor() ([]byte, []int) {
	return file_eventsource_proto_rawDescGZIP(), []int{0}
}

func (x *EventParam) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventParam) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *EventParam) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *EventParam) GetAggregateType() string {
	if x != nil {
		return x.AggregateType
	}
	return ""
}

func (x *EventParam) GetEventData() string {
	if x != nil {
		return x.EventData
	}
	return ""
}

func (x *EventParam) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type ResponseParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResponseParam) Reset() {
	*x = ResponseParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventsource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseParam) ProtoMessage() {}

func (x *ResponseParam) ProtoReflect() protoreflect.Message {
	mi := &file_eventsource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseParam.ProtoReflect.Descriptor instead.
func (*ResponseParam) Descriptor() ([]byte, []int) {
	return file_eventsource_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseParam) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *ResponseParam) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type EventFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId     string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AggregateId string `protobuf:"bytes,2,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
}

func (x *EventFilter) Reset() {
	*x = EventFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventsource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventFilter) ProtoMessage() {}

func (x *EventFilter) ProtoReflect() protoreflect.Message {
	mi := &file_eventsource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventFilter.ProtoReflect.Descriptor instead.
func (*EventFilter) Descriptor() ([]byte, []int) {
	return file_eventsource_proto_rawDescGZIP(), []int{2}
}

func (x *EventFilter) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventFilter) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventParam `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventsource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventsource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_eventsource_proto_rawDescGZIP(), []int{3}
}

func (x *EventResponse) GetEvents() []*EventParam {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_eventsource_proto protoreflect.FileDescriptor

var file_eventsource_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x22, 0xc9, 0x01, 0x0a,
	0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x44, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4b,
	0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x87, 0x01, 0x0a, 0x0a, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventsource_proto_rawDescOnce sync.Once
	file_eventsource_proto_rawDescData = file_eventsource_proto_rawDesc
)

func file_eventsource_proto_rawDescGZIP() []byte {
	file_eventsource_proto_rawDescOnce.Do(func() {
		file_eventsource_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventsource_proto_rawDescData)
	})
	return file_eventsource_proto_rawDescData
}

var file_eventsource_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eventsource_proto_goTypes = []interface{}{
	(*EventParam)(nil),    // 0: deposit.EventParam
	(*ResponseParam)(nil), // 1: deposit.ResponseParam
	(*EventFilter)(nil),   // 2: deposit.EventFilter
	(*EventResponse)(nil), // 3: deposit.EventResponse
}
var file_eventsource_proto_depIdxs = []int32{
	0, // 0: deposit.EventResponse.events:type_name -> deposit.EventParam
	2, // 1: deposit.EventStore.GetEvents:input_type -> deposit.EventFilter
	0, // 2: deposit.EventStore.CreateEvent:input_type -> deposit.EventParam
	3, // 3: deposit.EventStore.GetEvents:output_type -> deposit.EventResponse
	1, // 4: deposit.EventStore.CreateEvent:output_type -> deposit.ResponseParam
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eventsource_proto_init() }
func file_eventsource_proto_init() {
	if File_eventsource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventsource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventParam); i {
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
		file_eventsource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseParam); i {
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
		file_eventsource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventFilter); i {
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
		file_eventsource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
			RawDescriptor: file_eventsource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventsource_proto_goTypes,
		DependencyIndexes: file_eventsource_proto_depIdxs,
		MessageInfos:      file_eventsource_proto_msgTypes,
	}.Build()
	File_eventsource_proto = out.File
	file_eventsource_proto_rawDesc = nil
	file_eventsource_proto_goTypes = nil
	file_eventsource_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventStoreClient is the client API for EventStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventStoreClient interface {
	GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error)
	CreateEvent(ctx context.Context, in *EventParam, opts ...grpc.CallOption) (*ResponseParam, error)
}

type eventStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewEventStoreClient(cc grpc.ClientConnInterface) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/deposit.EventStore/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStoreClient) CreateEvent(ctx context.Context, in *EventParam, opts ...grpc.CallOption) (*ResponseParam, error) {
	out := new(ResponseParam)
	err := c.cc.Invoke(ctx, "/deposit.EventStore/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventStoreServer is the server API for EventStore service.
type EventStoreServer interface {
	GetEvents(context.Context, *EventFilter) (*EventResponse, error)
	CreateEvent(context.Context, *EventParam) (*ResponseParam, error)
}

// UnimplementedEventStoreServer can be embedded to have forward compatible implementations.
type UnimplementedEventStoreServer struct {
}

func (*UnimplementedEventStoreServer) GetEvents(context.Context, *EventFilter) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (*UnimplementedEventStoreServer) CreateEvent(context.Context, *EventParam) (*ResponseParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}

func RegisterEventStoreServer(s *grpc.Server, srv EventStoreServer) {
	s.RegisterService(&_EventStore_serviceDesc, srv)
}

func _EventStore_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deposit.EventStore/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).GetEvents(ctx, req.(*EventFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStore_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deposit.EventStore/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).CreateEvent(ctx, req.(*EventParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deposit.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _EventStore_GetEvents_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _EventStore_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventsource.proto",
}
