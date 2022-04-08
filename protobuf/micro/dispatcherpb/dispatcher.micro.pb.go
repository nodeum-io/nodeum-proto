// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: protobuf/dispatcher.micro.proto

package dispatcherpb

import (
	taskpb "github.com/nodeum-io/nodeum-plugins/protobuf/types/taskpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//	*StartRequest_Task
	//	*StartRequest_TaskId
	Source   isStartRequest_Source `protobuf_oneof:"source"`
	Metadata *structpb.Struct      `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatcher_micro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatcher_micro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_protobuf_dispatcher_micro_proto_rawDescGZIP(), []int{0}
}

func (m *StartRequest) GetSource() isStartRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *StartRequest) GetTask() *taskpb.Task {
	if x, ok := x.GetSource().(*StartRequest_Task); ok {
		return x.Task
	}
	return nil
}

func (x *StartRequest) GetTaskId() string {
	if x, ok := x.GetSource().(*StartRequest_TaskId); ok {
		return x.TaskId
	}
	return ""
}

func (x *StartRequest) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type isStartRequest_Source interface {
	isStartRequest_Source()
}

type StartRequest_Task struct {
	Task *taskpb.Task `protobuf:"bytes,1,opt,name=task,proto3,oneof"`
}

type StartRequest_TaskId struct {
	TaskId string `protobuf:"bytes,2,opt,name=task_id,proto3,oneof"`
}

func (*StartRequest_Task) isStartRequest_Source() {}

func (*StartRequest_TaskId) isStartRequest_Source() {}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg       string            `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Execution *taskpb.Execution `protobuf:"bytes,2,opt,name=execution,proto3" json:"execution,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatcher_micro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatcher_micro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_protobuf_dispatcher_micro_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *StartResponse) GetExecution() *taskpb.Execution {
	if x != nil {
		return x.Execution
	}
	return nil
}

type TriggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TriggerRequest) Reset() {
	*x = TriggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatcher_micro_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerRequest) ProtoMessage() {}

func (x *TriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatcher_micro_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerRequest.ProtoReflect.Descriptor instead.
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_dispatcher_micro_proto_rawDescGZIP(), []int{2}
}

func (x *TriggerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TriggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *TriggerResponse) Reset() {
	*x = TriggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatcher_micro_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerResponse) ProtoMessage() {}

func (x *TriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatcher_micro_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerResponse.ProtoReflect.Descriptor instead.
func (*TriggerResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_dispatcher_micro_proto_rawDescGZIP(), []int{3}
}

func (x *TriggerResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_protobuf_dispatcher_micro_proto protoreflect.FileDescriptor

var file_protobuf_dispatcher_micro_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x00, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x5b, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a,
	0x0e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x23, 0x0a, 0x0f, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0xf6, 0x03, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x74, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x28, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x05, 0x50, 0x61, 0x75,
	0x73, 0x65, 0x12, 0x2a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x22, 0x10, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x61, 0x75, 0x73, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12,
	0x2a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x11, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x78, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x2a, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_dispatcher_micro_proto_rawDescOnce sync.Once
	file_protobuf_dispatcher_micro_proto_rawDescData = file_protobuf_dispatcher_micro_proto_rawDesc
)

func file_protobuf_dispatcher_micro_proto_rawDescGZIP() []byte {
	file_protobuf_dispatcher_micro_proto_rawDescOnce.Do(func() {
		file_protobuf_dispatcher_micro_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_dispatcher_micro_proto_rawDescData)
	})
	return file_protobuf_dispatcher_micro_proto_rawDescData
}

var file_protobuf_dispatcher_micro_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_dispatcher_micro_proto_goTypes = []interface{}{
	(*StartRequest)(nil),     // 0: nodeum.protobuf.dispatcher.StartRequest
	(*StartResponse)(nil),    // 1: nodeum.protobuf.dispatcher.StartResponse
	(*TriggerRequest)(nil),   // 2: nodeum.protobuf.dispatcher.TriggerRequest
	(*TriggerResponse)(nil),  // 3: nodeum.protobuf.dispatcher.TriggerResponse
	(*taskpb.Task)(nil),      // 4: nodeum.protobuf.Task
	(*structpb.Struct)(nil),  // 5: google.protobuf.Struct
	(*taskpb.Execution)(nil), // 6: nodeum.protobuf.Execution
}
var file_protobuf_dispatcher_micro_proto_depIdxs = []int32{
	4, // 0: nodeum.protobuf.dispatcher.StartRequest.task:type_name -> nodeum.protobuf.Task
	5, // 1: nodeum.protobuf.dispatcher.StartRequest.metadata:type_name -> google.protobuf.Struct
	6, // 2: nodeum.protobuf.dispatcher.StartResponse.execution:type_name -> nodeum.protobuf.Execution
	0, // 3: nodeum.protobuf.dispatcher.Dispatcher.Start:input_type -> nodeum.protobuf.dispatcher.StartRequest
	2, // 4: nodeum.protobuf.dispatcher.Dispatcher.Pause:input_type -> nodeum.protobuf.dispatcher.TriggerRequest
	2, // 5: nodeum.protobuf.dispatcher.Dispatcher.Resume:input_type -> nodeum.protobuf.dispatcher.TriggerRequest
	2, // 6: nodeum.protobuf.dispatcher.Dispatcher.Stop:input_type -> nodeum.protobuf.dispatcher.TriggerRequest
	1, // 7: nodeum.protobuf.dispatcher.Dispatcher.Start:output_type -> nodeum.protobuf.dispatcher.StartResponse
	3, // 8: nodeum.protobuf.dispatcher.Dispatcher.Pause:output_type -> nodeum.protobuf.dispatcher.TriggerResponse
	3, // 9: nodeum.protobuf.dispatcher.Dispatcher.Resume:output_type -> nodeum.protobuf.dispatcher.TriggerResponse
	3, // 10: nodeum.protobuf.dispatcher.Dispatcher.Stop:output_type -> nodeum.protobuf.dispatcher.TriggerResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_dispatcher_micro_proto_init() }
func file_protobuf_dispatcher_micro_proto_init() {
	if File_protobuf_dispatcher_micro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_dispatcher_micro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_protobuf_dispatcher_micro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_protobuf_dispatcher_micro_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerRequest); i {
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
		file_protobuf_dispatcher_micro_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerResponse); i {
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
	file_protobuf_dispatcher_micro_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StartRequest_Task)(nil),
		(*StartRequest_TaskId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_dispatcher_micro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_dispatcher_micro_proto_goTypes,
		DependencyIndexes: file_protobuf_dispatcher_micro_proto_depIdxs,
		MessageInfos:      file_protobuf_dispatcher_micro_proto_msgTypes,
	}.Build()
	File_protobuf_dispatcher_micro_proto = out.File
	file_protobuf_dispatcher_micro_proto_rawDesc = nil
	file_protobuf_dispatcher_micro_proto_goTypes = nil
	file_protobuf_dispatcher_micro_proto_depIdxs = nil
}
