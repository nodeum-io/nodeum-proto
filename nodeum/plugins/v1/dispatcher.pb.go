// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: nodeum/plugins/v1/dispatcher.proto

package pluginsv1

import (
	v1 "github.com/nodeum-io/nodeum-proto/nodeum/common/v1"
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

type BeforeTaskExecutedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     *v1.Task         `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Metadata *structpb.Struct `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *BeforeTaskExecutedRequest) Reset() {
	*x = BeforeTaskExecutedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeforeTaskExecutedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeforeTaskExecutedRequest) ProtoMessage() {}

func (x *BeforeTaskExecutedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeforeTaskExecutedRequest.ProtoReflect.Descriptor instead.
func (*BeforeTaskExecutedRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_plugins_v1_dispatcher_proto_rawDescGZIP(), []int{0}
}

func (x *BeforeTaskExecutedRequest) GetTask() *v1.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *BeforeTaskExecutedRequest) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type BeforeTaskExecutedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     *v1.Task         `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Metadata *structpb.Struct `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *BeforeTaskExecutedResponse) Reset() {
	*x = BeforeTaskExecutedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeforeTaskExecutedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeforeTaskExecutedResponse) ProtoMessage() {}

func (x *BeforeTaskExecutedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeforeTaskExecutedResponse.ProtoReflect.Descriptor instead.
func (*BeforeTaskExecutedResponse) Descriptor() ([]byte, []int) {
	return file_nodeum_plugins_v1_dispatcher_proto_rawDescGZIP(), []int{1}
}

func (x *BeforeTaskExecutedResponse) GetTask() *v1.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *BeforeTaskExecutedResponse) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type BeforeRequestDispatchedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Execution *v1.Execution `protobuf:"bytes,1,opt,name=execution,proto3" json:"execution,omitempty"`
	Request   *v1.Request   `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *BeforeRequestDispatchedRequest) Reset() {
	*x = BeforeRequestDispatchedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeforeRequestDispatchedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeforeRequestDispatchedRequest) ProtoMessage() {}

func (x *BeforeRequestDispatchedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeforeRequestDispatchedRequest.ProtoReflect.Descriptor instead.
func (*BeforeRequestDispatchedRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_plugins_v1_dispatcher_proto_rawDescGZIP(), []int{2}
}

func (x *BeforeRequestDispatchedRequest) GetExecution() *v1.Execution {
	if x != nil {
		return x.Execution
	}
	return nil
}

func (x *BeforeRequestDispatchedRequest) GetRequest() *v1.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type BeforeRequestDispatchedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *v1.Request `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Discard bool        `protobuf:"varint,3,opt,name=discard,proto3" json:"discard,omitempty"`
}

func (x *BeforeRequestDispatchedResponse) Reset() {
	*x = BeforeRequestDispatchedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeforeRequestDispatchedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeforeRequestDispatchedResponse) ProtoMessage() {}

func (x *BeforeRequestDispatchedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_plugins_v1_dispatcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeforeRequestDispatchedResponse.ProtoReflect.Descriptor instead.
func (*BeforeRequestDispatchedResponse) Descriptor() ([]byte, []int) {
	return file_nodeum_plugins_v1_dispatcher_proto_rawDescGZIP(), []int{3}
}

func (x *BeforeRequestDispatchedResponse) GetRequest() *v1.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *BeforeRequestDispatchedResponse) GetDiscard() bool {
	if x != nil {
		return x.Discard
	}
	return false
}

var File_nodeum_plugins_v1_dispatcher_proto protoreflect.FileDescriptor

var file_nodeum_plugins_v1_dispatcher_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x19, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x33, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x7d, 0x0a, 0x1a, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x90, 0x01, 0x0a, 0x1e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x70, 0x0a, 0x1f, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69,
	0x73, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73,
	0x63, 0x61, 0x72, 0x64, 0x32, 0x8f, 0x02, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x71, 0x0a, 0x12, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x17, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x12,
	0x31, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodeum_plugins_v1_dispatcher_proto_rawDescOnce sync.Once
	file_nodeum_plugins_v1_dispatcher_proto_rawDescData = file_nodeum_plugins_v1_dispatcher_proto_rawDesc
)

func file_nodeum_plugins_v1_dispatcher_proto_rawDescGZIP() []byte {
	file_nodeum_plugins_v1_dispatcher_proto_rawDescOnce.Do(func() {
		file_nodeum_plugins_v1_dispatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeum_plugins_v1_dispatcher_proto_rawDescData)
	})
	return file_nodeum_plugins_v1_dispatcher_proto_rawDescData
}

var file_nodeum_plugins_v1_dispatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nodeum_plugins_v1_dispatcher_proto_goTypes = []interface{}{
	(*BeforeTaskExecutedRequest)(nil),       // 0: nodeum.plugins.v1.BeforeTaskExecutedRequest
	(*BeforeTaskExecutedResponse)(nil),      // 1: nodeum.plugins.v1.BeforeTaskExecutedResponse
	(*BeforeRequestDispatchedRequest)(nil),  // 2: nodeum.plugins.v1.BeforeRequestDispatchedRequest
	(*BeforeRequestDispatchedResponse)(nil), // 3: nodeum.plugins.v1.BeforeRequestDispatchedResponse
	(*v1.Task)(nil),                         // 4: nodeum.common.v1.Task
	(*structpb.Struct)(nil),                 // 5: google.protobuf.Struct
	(*v1.Execution)(nil),                    // 6: nodeum.common.v1.Execution
	(*v1.Request)(nil),                      // 7: nodeum.common.v1.Request
}
var file_nodeum_plugins_v1_dispatcher_proto_depIdxs = []int32{
	4, // 0: nodeum.plugins.v1.BeforeTaskExecutedRequest.task:type_name -> nodeum.common.v1.Task
	5, // 1: nodeum.plugins.v1.BeforeTaskExecutedRequest.metadata:type_name -> google.protobuf.Struct
	4, // 2: nodeum.plugins.v1.BeforeTaskExecutedResponse.task:type_name -> nodeum.common.v1.Task
	5, // 3: nodeum.plugins.v1.BeforeTaskExecutedResponse.metadata:type_name -> google.protobuf.Struct
	6, // 4: nodeum.plugins.v1.BeforeRequestDispatchedRequest.execution:type_name -> nodeum.common.v1.Execution
	7, // 5: nodeum.plugins.v1.BeforeRequestDispatchedRequest.request:type_name -> nodeum.common.v1.Request
	7, // 6: nodeum.plugins.v1.BeforeRequestDispatchedResponse.request:type_name -> nodeum.common.v1.Request
	0, // 7: nodeum.plugins.v1.DispatcherPluginService.BeforeTaskExecuted:input_type -> nodeum.plugins.v1.BeforeTaskExecutedRequest
	2, // 8: nodeum.plugins.v1.DispatcherPluginService.BeforeRequestDispatched:input_type -> nodeum.plugins.v1.BeforeRequestDispatchedRequest
	1, // 9: nodeum.plugins.v1.DispatcherPluginService.BeforeTaskExecuted:output_type -> nodeum.plugins.v1.BeforeTaskExecutedResponse
	3, // 10: nodeum.plugins.v1.DispatcherPluginService.BeforeRequestDispatched:output_type -> nodeum.plugins.v1.BeforeRequestDispatchedResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_nodeum_plugins_v1_dispatcher_proto_init() }
func file_nodeum_plugins_v1_dispatcher_proto_init() {
	if File_nodeum_plugins_v1_dispatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodeum_plugins_v1_dispatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeforeTaskExecutedRequest); i {
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
		file_nodeum_plugins_v1_dispatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeforeTaskExecutedResponse); i {
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
		file_nodeum_plugins_v1_dispatcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeforeRequestDispatchedRequest); i {
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
		file_nodeum_plugins_v1_dispatcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeforeRequestDispatchedResponse); i {
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
			RawDescriptor: file_nodeum_plugins_v1_dispatcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodeum_plugins_v1_dispatcher_proto_goTypes,
		DependencyIndexes: file_nodeum_plugins_v1_dispatcher_proto_depIdxs,
		MessageInfos:      file_nodeum_plugins_v1_dispatcher_proto_msgTypes,
	}.Build()
	File_nodeum_plugins_v1_dispatcher_proto = out.File
	file_nodeum_plugins_v1_dispatcher_proto_rawDesc = nil
	file_nodeum_plugins_v1_dispatcher_proto_goTypes = nil
	file_nodeum_plugins_v1_dispatcher_proto_depIdxs = nil
}
