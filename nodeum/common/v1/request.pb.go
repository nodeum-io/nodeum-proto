// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: nodeum/common/v1/request.proto

package commonv1

import (
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

type CopyRequest_Type int32

const (
	CopyRequest_TYPE_UNSPECIFIED CopyRequest_Type = 0
	CopyRequest_TYPE_COPY        CopyRequest_Type = 1
	CopyRequest_TYPE_MOVE        CopyRequest_Type = 2
)

// Enum value maps for CopyRequest_Type.
var (
	CopyRequest_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_COPY",
		2: "TYPE_MOVE",
	}
	CopyRequest_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_COPY":        1,
		"TYPE_MOVE":        2,
	}
)

func (x CopyRequest_Type) Enum() *CopyRequest_Type {
	p := new(CopyRequest_Type)
	*p = x
	return p
}

func (x CopyRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CopyRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeum_common_v1_request_proto_enumTypes[0].Descriptor()
}

func (CopyRequest_Type) Type() protoreflect.EnumType {
	return &file_nodeum_common_v1_request_proto_enumTypes[0]
}

func (x CopyRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CopyRequest_Type.Descriptor instead.
func (CopyRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_nodeum_common_v1_request_proto_rawDescGZIP(), []int{1, 0}
}

type ScanRequest_Type int32

const (
	ScanRequest_TYPE_UNSPECIFIED ScanRequest_Type = 0
	ScanRequest_TYPE_SCAN        ScanRequest_Type = 1
	ScanRequest_TYPE_READDIR     ScanRequest_Type = 2
)

// Enum value maps for ScanRequest_Type.
var (
	ScanRequest_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_SCAN",
		2: "TYPE_READDIR",
	}
	ScanRequest_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_SCAN":        1,
		"TYPE_READDIR":     2,
	}
)

func (x ScanRequest_Type) Enum() *ScanRequest_Type {
	p := new(ScanRequest_Type)
	*p = x
	return p
}

func (x ScanRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScanRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeum_common_v1_request_proto_enumTypes[1].Descriptor()
}

func (ScanRequest_Type) Type() protoreflect.EnumType {
	return &file_nodeum_common_v1_request_proto_enumTypes[1]
}

func (x ScanRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScanRequest_Type.Descriptor instead.
func (ScanRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_nodeum_common_v1_request_proto_rawDescGZIP(), []int{3, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Request:
	//	*Request_Copy
	//	*Request_Remove
	//	*Request_Scan
	Request isRequest_Request `protobuf_oneof:"request"`
	Options *structpb.Struct  `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Request) GetRequest() isRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Request) GetCopy() *CopyRequest {
	if x, ok := x.GetRequest().(*Request_Copy); ok {
		return x.Copy
	}
	return nil
}

func (x *Request) GetRemove() *RemoveRequest {
	if x, ok := x.GetRequest().(*Request_Remove); ok {
		return x.Remove
	}
	return nil
}

func (x *Request) GetScan() *ScanRequest {
	if x, ok := x.GetRequest().(*Request_Scan); ok {
		return x.Scan
	}
	return nil
}

func (x *Request) GetOptions() *structpb.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type isRequest_Request interface {
	isRequest_Request()
}

type Request_Copy struct {
	Copy *CopyRequest `protobuf:"bytes,2,opt,name=copy,proto3,oneof"`
}

type Request_Remove struct {
	Remove *RemoveRequest `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

type Request_Scan struct {
	Scan *ScanRequest `protobuf:"bytes,4,opt,name=scan,proto3,oneof"`
}

func (*Request_Copy) isRequest_Request() {}

func (*Request_Remove) isRequest_Request() {}

func (*Request_Scan) isRequest_Request() {}

type CopyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            CopyRequest_Type `protobuf:"varint,2,opt,name=type,proto3,enum=nodeum.common.v1.CopyRequest_Type" json:"type,omitempty"`
	Source          *Storage         `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	SourcePath      string           `protobuf:"bytes,4,opt,name=source_path,proto3" json:"source_path,omitempty"`
	Destination     *Storage         `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
	DestinationPath string           `protobuf:"bytes,6,opt,name=destination_path,proto3" json:"destination_path,omitempty"`
	Info            *NodeInfo        `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CopyRequest) Reset() {
	*x = CopyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyRequest) ProtoMessage() {}

func (x *CopyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyRequest.ProtoReflect.Descriptor instead.
func (*CopyRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_request_proto_rawDescGZIP(), []int{1}
}

func (x *CopyRequest) GetType() CopyRequest_Type {
	if x != nil {
		return x.Type
	}
	return CopyRequest_TYPE_UNSPECIFIED
}

func (x *CopyRequest) GetSource() *Storage {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CopyRequest) GetSourcePath() string {
	if x != nil {
		return x.SourcePath
	}
	return ""
}

func (x *CopyRequest) GetDestination() *Storage {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *CopyRequest) GetDestinationPath() string {
	if x != nil {
		return x.DestinationPath
	}
	return ""
}

func (x *CopyRequest) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage *Storage  `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Path    string    `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Info    *NodeInfo `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_request_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveRequest) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *RemoveRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RemoveRequest) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    ScanRequest_Type `protobuf:"varint,2,opt,name=type,proto3,enum=nodeum.common.v1.ScanRequest_Type" json:"type,omitempty"`
	Storage *Storage         `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Path    string           `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_request_proto_rawDescGZIP(), []int{3}
}

func (x *ScanRequest) GetType() ScanRequest_Type {
	if x != nil {
		return x.Type
	}
	return ScanRequest_TYPE_UNSPECIFIED
}

func (x *ScanRequest) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *ScanRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_nodeum_common_v1_request_proto protoreflect.FileDescriptor

var file_nodeum_common_v1_request_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33,
	0x0a, 0x04, 0x63, 0x6f, 0x70, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x63,
	0x6f, 0x70, 0x79, 0x12, 0x39, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x73, 0x63, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x73,
	0x63, 0x61, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xef, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x3b,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x3a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f,
	0x50, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x56,
	0x45, 0x10, 0x02, 0x22, 0x88, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2e,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xcd,
	0x01, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x3d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x44, 0x49, 0x52, 0x10, 0x02, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodeum_common_v1_request_proto_rawDescOnce sync.Once
	file_nodeum_common_v1_request_proto_rawDescData = file_nodeum_common_v1_request_proto_rawDesc
)

func file_nodeum_common_v1_request_proto_rawDescGZIP() []byte {
	file_nodeum_common_v1_request_proto_rawDescOnce.Do(func() {
		file_nodeum_common_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeum_common_v1_request_proto_rawDescData)
	})
	return file_nodeum_common_v1_request_proto_rawDescData
}

var file_nodeum_common_v1_request_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nodeum_common_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nodeum_common_v1_request_proto_goTypes = []interface{}{
	(CopyRequest_Type)(0),   // 0: nodeum.common.v1.CopyRequest.Type
	(ScanRequest_Type)(0),   // 1: nodeum.common.v1.ScanRequest.Type
	(*Request)(nil),         // 2: nodeum.common.v1.Request
	(*CopyRequest)(nil),     // 3: nodeum.common.v1.CopyRequest
	(*RemoveRequest)(nil),   // 4: nodeum.common.v1.RemoveRequest
	(*ScanRequest)(nil),     // 5: nodeum.common.v1.ScanRequest
	(*structpb.Struct)(nil), // 6: google.protobuf.Struct
	(*Storage)(nil),         // 7: nodeum.common.v1.Storage
	(*NodeInfo)(nil),        // 8: nodeum.common.v1.NodeInfo
}
var file_nodeum_common_v1_request_proto_depIdxs = []int32{
	3,  // 0: nodeum.common.v1.Request.copy:type_name -> nodeum.common.v1.CopyRequest
	4,  // 1: nodeum.common.v1.Request.remove:type_name -> nodeum.common.v1.RemoveRequest
	5,  // 2: nodeum.common.v1.Request.scan:type_name -> nodeum.common.v1.ScanRequest
	6,  // 3: nodeum.common.v1.Request.options:type_name -> google.protobuf.Struct
	0,  // 4: nodeum.common.v1.CopyRequest.type:type_name -> nodeum.common.v1.CopyRequest.Type
	7,  // 5: nodeum.common.v1.CopyRequest.source:type_name -> nodeum.common.v1.Storage
	7,  // 6: nodeum.common.v1.CopyRequest.destination:type_name -> nodeum.common.v1.Storage
	8,  // 7: nodeum.common.v1.CopyRequest.info:type_name -> nodeum.common.v1.NodeInfo
	7,  // 8: nodeum.common.v1.RemoveRequest.storage:type_name -> nodeum.common.v1.Storage
	8,  // 9: nodeum.common.v1.RemoveRequest.info:type_name -> nodeum.common.v1.NodeInfo
	1,  // 10: nodeum.common.v1.ScanRequest.type:type_name -> nodeum.common.v1.ScanRequest.Type
	7,  // 11: nodeum.common.v1.ScanRequest.storage:type_name -> nodeum.common.v1.Storage
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_nodeum_common_v1_request_proto_init() }
func file_nodeum_common_v1_request_proto_init() {
	if File_nodeum_common_v1_request_proto != nil {
		return
	}
	file_nodeum_common_v1_nodeinfo_proto_init()
	file_nodeum_common_v1_storages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nodeum_common_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_nodeum_common_v1_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyRequest); i {
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
		file_nodeum_common_v1_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_nodeum_common_v1_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRequest); i {
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
	file_nodeum_common_v1_request_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_Copy)(nil),
		(*Request_Remove)(nil),
		(*Request_Scan)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nodeum_common_v1_request_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nodeum_common_v1_request_proto_goTypes,
		DependencyIndexes: file_nodeum_common_v1_request_proto_depIdxs,
		EnumInfos:         file_nodeum_common_v1_request_proto_enumTypes,
		MessageInfos:      file_nodeum_common_v1_request_proto_msgTypes,
	}.Build()
	File_nodeum_common_v1_request_proto = out.File
	file_nodeum_common_v1_request_proto_rawDesc = nil
	file_nodeum_common_v1_request_proto_goTypes = nil
	file_nodeum_common_v1_request_proto_depIdxs = nil
}
