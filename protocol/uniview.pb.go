// SPDX-FileCopyrightText: © Moritz Poldrack
// SPDX-License-Identifier: AGPL-3.0-or-later

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: protocol/uniview.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_EVENT_UNSPECIFIED       EventType = 0
	EventType_EVENT_JOIN              EventType = 1
	EventType_EVENT_PAUSE             EventType = 2
	EventType_EVENT_JUMP              EventType = 3
	EventType_EVENT_SERVER_CLOSE      EventType = 4
	EventType_EVENT_CLIENT_DISCONNECT EventType = 5
	EventType_EVENT_SERVER_PING       EventType = 6
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_JOIN",
		2: "EVENT_PAUSE",
		3: "EVENT_JUMP",
		4: "EVENT_SERVER_CLOSE",
		5: "EVENT_CLIENT_DISCONNECT",
		6: "EVENT_SERVER_PING",
	}
	EventType_value = map[string]int32{
		"EVENT_UNSPECIFIED":       0,
		"EVENT_JOIN":              1,
		"EVENT_PAUSE":             2,
		"EVENT_JUMP":              3,
		"EVENT_SERVER_CLOSE":      4,
		"EVENT_CLIENT_DISCONNECT": 5,
		"EVENT_SERVER_PING":       6,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_uniview_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_protocol_uniview_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_uniview_proto_rawDescGZIP(), []int{0}
}

type RoomEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.EventType" json:"type,omitempty"`
	// Types that are assignable to Event:
	//
	//	*RoomEvent_Join
	//	*RoomEvent_PauseEvent
	//	*RoomEvent_JumpEvent
	Event isRoomEvent_Event `protobuf_oneof:"event"`
}

func (x *RoomEvent) Reset() {
	*x = RoomEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_uniview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEvent) ProtoMessage() {}

func (x *RoomEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_uniview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomEvent.ProtoReflect.Descriptor instead.
func (*RoomEvent) Descriptor() ([]byte, []int) {
	return file_protocol_uniview_proto_rawDescGZIP(), []int{0}
}

func (x *RoomEvent) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_EVENT_UNSPECIFIED
}

func (m *RoomEvent) GetEvent() isRoomEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *RoomEvent) GetJoin() *RoomJoin {
	if x, ok := x.GetEvent().(*RoomEvent_Join); ok {
		return x.Join
	}
	return nil
}

func (x *RoomEvent) GetPauseEvent() *PlayPause {
	if x, ok := x.GetEvent().(*RoomEvent_PauseEvent); ok {
		return x.PauseEvent
	}
	return nil
}

func (x *RoomEvent) GetJumpEvent() *PlaybackJump {
	if x, ok := x.GetEvent().(*RoomEvent_JumpEvent); ok {
		return x.JumpEvent
	}
	return nil
}

type isRoomEvent_Event interface {
	isRoomEvent_Event()
}

type RoomEvent_Join struct {
	Join *RoomJoin `protobuf:"bytes,2,opt,name=join,proto3,oneof"`
}

type RoomEvent_PauseEvent struct {
	PauseEvent *PlayPause `protobuf:"bytes,3,opt,name=pauseEvent,proto3,oneof"`
}

type RoomEvent_JumpEvent struct {
	JumpEvent *PlaybackJump `protobuf:"bytes,4,opt,name=jumpEvent,proto3,oneof"`
}

func (*RoomEvent_Join) isRoomEvent_Event() {}

func (*RoomEvent_PauseEvent) isRoomEvent_Event() {}

func (*RoomEvent_JumpEvent) isRoomEvent_Event() {}

type RoomJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp *durationpb.Duration `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Url       string               `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RoomJoin) Reset() {
	*x = RoomJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_uniview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomJoin) ProtoMessage() {}

func (x *RoomJoin) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_uniview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomJoin.ProtoReflect.Descriptor instead.
func (*RoomJoin) Descriptor() ([]byte, []int) {
	return file_protocol_uniview_proto_rawDescGZIP(), []int{1}
}

func (x *RoomJoin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoomJoin) GetTimestamp() *durationpb.Duration {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *RoomJoin) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type PlayPause struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pause     bool                 `protobuf:"varint,1,opt,name=pause,proto3" json:"pause,omitempty"`
	Timestamp *durationpb.Duration `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PlayPause) Reset() {
	*x = PlayPause{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_uniview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayPause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayPause) ProtoMessage() {}

func (x *PlayPause) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_uniview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayPause.ProtoReflect.Descriptor instead.
func (*PlayPause) Descriptor() ([]byte, []int) {
	return file_protocol_uniview_proto_rawDescGZIP(), []int{2}
}

func (x *PlayPause) GetPause() bool {
	if x != nil {
		return x.Pause
	}
	return false
}

func (x *PlayPause) GetTimestamp() *durationpb.Duration {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type PlaybackJump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *durationpb.Duration `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PlaybackJump) Reset() {
	*x = PlaybackJump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_uniview_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaybackJump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaybackJump) ProtoMessage() {}

func (x *PlaybackJump) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_uniview_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaybackJump.ProtoReflect.Descriptor instead.
func (*PlaybackJump) Descriptor() ([]byte, []int) {
	return file_protocol_uniview_proto_rawDescGZIP(), []int{3}
}

func (x *PlaybackJump) GetTimestamp() *durationpb.Duration {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_protocol_uniview_proto protoreflect.FileDescriptor

var file_protocol_uniview_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6a, 0x6f, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6a,
	0x6f, 0x69, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x61, 0x75, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x61, 0x75, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0a,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x6a, 0x75,
	0x6d, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63,
	0x6b, 0x4a, 0x75, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x6a, 0x75, 0x6d, 0x70, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x08, 0x52,
	0x6f, 0x6f, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x5a, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x61,
	0x75, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x70, 0x61, 0x75, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x47, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x4a, 0x75,
	0x6d, 0x70, 0x12, 0x37, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x9f, 0x01, 0x0a, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4a, 0x55, 0x4d, 0x50, 0x10,
	0x03, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x32, 0x3f, 0x0a,
	0x07, 0x55, 0x6e, 0x69, 0x56, 0x69, 0x65, 0x77, 0x12, 0x34, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x23,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x2e, 0x73, 0x72, 0x2e, 0x68, 0x74, 0x2f, 0x7e, 0x6d, 0x70, 0x6c,
	0x64, 0x72, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_uniview_proto_rawDescOnce sync.Once
	file_protocol_uniview_proto_rawDescData = file_protocol_uniview_proto_rawDesc
)

func file_protocol_uniview_proto_rawDescGZIP() []byte {
	file_protocol_uniview_proto_rawDescOnce.Do(func() {
		file_protocol_uniview_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_uniview_proto_rawDescData)
	})
	return file_protocol_uniview_proto_rawDescData
}

var file_protocol_uniview_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_uniview_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_uniview_proto_goTypes = []interface{}{
	(EventType)(0),              // 0: protocol.EventType
	(*RoomEvent)(nil),           // 1: protocol.RoomEvent
	(*RoomJoin)(nil),            // 2: protocol.RoomJoin
	(*PlayPause)(nil),           // 3: protocol.PlayPause
	(*PlaybackJump)(nil),        // 4: protocol.PlaybackJump
	(*durationpb.Duration)(nil), // 5: google.protobuf.Duration
}
var file_protocol_uniview_proto_depIdxs = []int32{
	0, // 0: protocol.RoomEvent.type:type_name -> protocol.EventType
	2, // 1: protocol.RoomEvent.join:type_name -> protocol.RoomJoin
	3, // 2: protocol.RoomEvent.pauseEvent:type_name -> protocol.PlayPause
	4, // 3: protocol.RoomEvent.jumpEvent:type_name -> protocol.PlaybackJump
	5, // 4: protocol.RoomJoin.timestamp:type_name -> google.protobuf.Duration
	5, // 5: protocol.PlayPause.timestamp:type_name -> google.protobuf.Duration
	5, // 6: protocol.PlaybackJump.timestamp:type_name -> google.protobuf.Duration
	1, // 7: protocol.UniView.Room:input_type -> protocol.RoomEvent
	1, // 8: protocol.UniView.Room:output_type -> protocol.RoomEvent
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protocol_uniview_proto_init() }
func file_protocol_uniview_proto_init() {
	if File_protocol_uniview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_uniview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomEvent); i {
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
		file_protocol_uniview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomJoin); i {
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
		file_protocol_uniview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayPause); i {
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
		file_protocol_uniview_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaybackJump); i {
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
	file_protocol_uniview_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RoomEvent_Join)(nil),
		(*RoomEvent_PauseEvent)(nil),
		(*RoomEvent_JumpEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocol_uniview_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_uniview_proto_goTypes,
		DependencyIndexes: file_protocol_uniview_proto_depIdxs,
		EnumInfos:         file_protocol_uniview_proto_enumTypes,
		MessageInfos:      file_protocol_uniview_proto_msgTypes,
	}.Build()
	File_protocol_uniview_proto = out.File
	file_protocol_uniview_proto_rawDesc = nil
	file_protocol_uniview_proto_goTypes = nil
	file_protocol_uniview_proto_depIdxs = nil
}
