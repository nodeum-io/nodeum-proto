// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: nodeum/common/v1/result.proto

package commonv1

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

type FailedResult_ErrorType int32

const (
	FailedResult_ERROR_TYPE_UNSPECIFIED FailedResult_ErrorType = 0
	FailedResult_ERROR_TYPE_UNKNOWN     FailedResult_ErrorType = 1
	FailedResult_ERROR_TYPE_NOT_EXIST   FailedResult_ErrorType = 2
	FailedResult_ERROR_TYPE_NOT_DIR     FailedResult_ErrorType = 3
	FailedResult_ERROR_TYPE_PERMISSION  FailedResult_ErrorType = 4
	FailedResult_ERROR_TYPE_EXIST       FailedResult_ErrorType = 5
	FailedResult_ERROR_TYPE_NOT_EMPTY   FailedResult_ErrorType = 6
	FailedResult_ERROR_TYPE_TIMEOUT     FailedResult_ErrorType = 7
)

// Enum value maps for FailedResult_ErrorType.
var (
	FailedResult_ErrorType_name = map[int32]string{
		0: "ERROR_TYPE_UNSPECIFIED",
		1: "ERROR_TYPE_UNKNOWN",
		2: "ERROR_TYPE_NOT_EXIST",
		3: "ERROR_TYPE_NOT_DIR",
		4: "ERROR_TYPE_PERMISSION",
		5: "ERROR_TYPE_EXIST",
		6: "ERROR_TYPE_NOT_EMPTY",
		7: "ERROR_TYPE_TIMEOUT",
	}
	FailedResult_ErrorType_value = map[string]int32{
		"ERROR_TYPE_UNSPECIFIED": 0,
		"ERROR_TYPE_UNKNOWN":     1,
		"ERROR_TYPE_NOT_EXIST":   2,
		"ERROR_TYPE_NOT_DIR":     3,
		"ERROR_TYPE_PERMISSION":  4,
		"ERROR_TYPE_EXIST":       5,
		"ERROR_TYPE_NOT_EMPTY":   6,
		"ERROR_TYPE_TIMEOUT":     7,
	}
)

func (x FailedResult_ErrorType) Enum() *FailedResult_ErrorType {
	p := new(FailedResult_ErrorType)
	*p = x
	return p
}

func (x FailedResult_ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FailedResult_ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeum_common_v1_result_proto_enumTypes[0].Descriptor()
}

func (FailedResult_ErrorType) Type() protoreflect.EnumType {
	return &file_nodeum_common_v1_result_proto_enumTypes[0]
}

func (x FailedResult_ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FailedResult_ErrorType.Descriptor instead.
func (FailedResult_ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{6, 0}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Types that are assignable to Result:
	//	*Result_Progress
	//	*Result_Checkin
	//	*Result_Checkout
	//	*Result_Seen
	//	*Result_Over
	//	*Result_Failed
	Result isResult_Result `protobuf_oneof:"result"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (m *Result) GetResult() isResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *Result) GetProgress() *ProgressResult {
	if x, ok := x.GetResult().(*Result_Progress); ok {
		return x.Progress
	}
	return nil
}

func (x *Result) GetCheckin() *CheckInResult {
	if x, ok := x.GetResult().(*Result_Checkin); ok {
		return x.Checkin
	}
	return nil
}

func (x *Result) GetCheckout() *CheckOutResult {
	if x, ok := x.GetResult().(*Result_Checkout); ok {
		return x.Checkout
	}
	return nil
}

func (x *Result) GetSeen() *SeenResult {
	if x, ok := x.GetResult().(*Result_Seen); ok {
		return x.Seen
	}
	return nil
}

func (x *Result) GetOver() *OverResult {
	if x, ok := x.GetResult().(*Result_Over); ok {
		return x.Over
	}
	return nil
}

func (x *Result) GetFailed() *FailedResult {
	if x, ok := x.GetResult().(*Result_Failed); ok {
		return x.Failed
	}
	return nil
}

type isResult_Result interface {
	isResult_Result()
}

type Result_Progress struct {
	Progress *ProgressResult `protobuf:"bytes,2,opt,name=progress,proto3,oneof"`
}

type Result_Checkin struct {
	Checkin *CheckInResult `protobuf:"bytes,3,opt,name=checkin,proto3,oneof"`
}

type Result_Checkout struct {
	Checkout *CheckOutResult `protobuf:"bytes,4,opt,name=checkout,proto3,oneof"`
}

type Result_Seen struct {
	Seen *SeenResult `protobuf:"bytes,5,opt,name=seen,proto3,oneof"`
}

type Result_Over struct {
	Over *OverResult `protobuf:"bytes,6,opt,name=over,proto3,oneof"`
}

type Result_Failed struct {
	Failed *FailedResult `protobuf:"bytes,7,opt,name=failed,proto3,oneof"`
}

func (*Result_Progress) isResult_Result() {}

func (*Result_Checkin) isResult_Result() {}

func (*Result_Checkout) isResult_Result() {}

func (*Result_Seen) isResult_Result() {}

func (*Result_Over) isResult_Result() {}

func (*Result_Failed) isResult_Result() {}

type ProgressResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SizeDone int64   `protobuf:"varint,2,opt,name=size_done,proto3" json:"size_done,omitempty"`
	SizeTodo int64   `protobuf:"varint,3,opt,name=size_todo,proto3" json:"size_todo,omitempty"`
	Speed    float64 `protobuf:"fixed64,4,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *ProgressResult) Reset() {
	*x = ProgressResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressResult) ProtoMessage() {}

func (x *ProgressResult) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressResult.ProtoReflect.Descriptor instead.
func (*ProgressResult) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{1}
}

func (x *ProgressResult) GetSizeDone() int64 {
	if x != nil {
		return x.SizeDone
	}
	return 0
}

func (x *ProgressResult) GetSizeTodo() int64 {
	if x != nil {
		return x.SizeTodo
	}
	return 0
}

func (x *ProgressResult) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

type CheckInResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage  *Storage  `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Info     *NodeInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Indirect bool      `protobuf:"varint,4,opt,name=indirect,proto3" json:"indirect,omitempty"`
}

func (x *CheckInResult) Reset() {
	*x = CheckInResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInResult) ProtoMessage() {}

func (x *CheckInResult) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInResult.ProtoReflect.Descriptor instead.
func (*CheckInResult) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{2}
}

func (x *CheckInResult) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *CheckInResult) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *CheckInResult) GetIndirect() bool {
	if x != nil {
		return x.Indirect
	}
	return false
}

type CheckOutResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage *Storage `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Path    string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CheckOutResult) Reset() {
	*x = CheckOutResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOutResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOutResult) ProtoMessage() {}

func (x *CheckOutResult) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOutResult.ProtoReflect.Descriptor instead.
func (*CheckOutResult) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{3}
}

func (x *CheckOutResult) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *CheckOutResult) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SeenResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage *Storage  `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Info    *NodeInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SeenResult) Reset() {
	*x = SeenResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeenResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeenResult) ProtoMessage() {}

func (x *SeenResult) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeenResult.ProtoReflect.Descriptor instead.
func (*SeenResult) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{4}
}

func (x *SeenResult) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *SeenResult) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type OverResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OverResult) Reset() {
	*x = OverResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverResult) ProtoMessage() {}

func (x *OverResult) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverResult.ProtoReflect.Descriptor instead.
func (*OverResult) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{5}
}

type FailedResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage  *Storage               `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Path     string                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Error    FailedResult_ErrorType `protobuf:"varint,4,opt,name=error,proto3,enum=nodeum.common.v1.FailedResult_ErrorType" json:"error,omitempty"`
	ErrorStr string                 `protobuf:"bytes,5,opt,name=error_str,proto3" json:"error_str,omitempty"`
}

func (x *FailedResult) Reset() {
	*x = FailedResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_result_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailedResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailedResult) ProtoMessage() {}

func (x *FailedResult) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_result_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailedResult.ProtoReflect.Descriptor instead.
func (*FailedResult) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_result_proto_rawDescGZIP(), []int{6}
}

func (x *FailedResult) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *FailedResult) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FailedResult) GetError() FailedResult_ErrorType {
	if x != nil {
		return x.Error
	}
	return FailedResult_ERROR_TYPE_UNSPECIFIED
}

func (x *FailedResult) GetErrorStr() string {
	if x != nil {
		return x.ErrorStr
	}
	return ""
}

var File_nodeum_common_v1_result_proto protoreflect.FileDescriptor

var file_nodeum_common_v1_result_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x03, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x12, 0x3e, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x12, 0x32, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x04,
	0x73, 0x65, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x62, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x22, 0x90, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x22, 0x59, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x71,
	0x0a, 0x0a, 0x53, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x4f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x8c, 0x03, 0x0a, 0x0c, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x22, 0xd4, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45,
	0x4d, 0x50, 0x54, 0x59, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x07, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodeum_common_v1_result_proto_rawDescOnce sync.Once
	file_nodeum_common_v1_result_proto_rawDescData = file_nodeum_common_v1_result_proto_rawDesc
)

func file_nodeum_common_v1_result_proto_rawDescGZIP() []byte {
	file_nodeum_common_v1_result_proto_rawDescOnce.Do(func() {
		file_nodeum_common_v1_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeum_common_v1_result_proto_rawDescData)
	})
	return file_nodeum_common_v1_result_proto_rawDescData
}

var file_nodeum_common_v1_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nodeum_common_v1_result_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nodeum_common_v1_result_proto_goTypes = []interface{}{
	(FailedResult_ErrorType)(0), // 0: nodeum.common.v1.FailedResult.ErrorType
	(*Result)(nil),              // 1: nodeum.common.v1.Result
	(*ProgressResult)(nil),      // 2: nodeum.common.v1.ProgressResult
	(*CheckInResult)(nil),       // 3: nodeum.common.v1.CheckInResult
	(*CheckOutResult)(nil),      // 4: nodeum.common.v1.CheckOutResult
	(*SeenResult)(nil),          // 5: nodeum.common.v1.SeenResult
	(*OverResult)(nil),          // 6: nodeum.common.v1.OverResult
	(*FailedResult)(nil),        // 7: nodeum.common.v1.FailedResult
	(*Request)(nil),             // 8: nodeum.common.v1.Request
	(*Storage)(nil),             // 9: nodeum.common.v1.Storage
	(*NodeInfo)(nil),            // 10: nodeum.common.v1.NodeInfo
}
var file_nodeum_common_v1_result_proto_depIdxs = []int32{
	8,  // 0: nodeum.common.v1.Result.request:type_name -> nodeum.common.v1.Request
	2,  // 1: nodeum.common.v1.Result.progress:type_name -> nodeum.common.v1.ProgressResult
	3,  // 2: nodeum.common.v1.Result.checkin:type_name -> nodeum.common.v1.CheckInResult
	4,  // 3: nodeum.common.v1.Result.checkout:type_name -> nodeum.common.v1.CheckOutResult
	5,  // 4: nodeum.common.v1.Result.seen:type_name -> nodeum.common.v1.SeenResult
	6,  // 5: nodeum.common.v1.Result.over:type_name -> nodeum.common.v1.OverResult
	7,  // 6: nodeum.common.v1.Result.failed:type_name -> nodeum.common.v1.FailedResult
	9,  // 7: nodeum.common.v1.CheckInResult.storage:type_name -> nodeum.common.v1.Storage
	10, // 8: nodeum.common.v1.CheckInResult.info:type_name -> nodeum.common.v1.NodeInfo
	9,  // 9: nodeum.common.v1.CheckOutResult.storage:type_name -> nodeum.common.v1.Storage
	9,  // 10: nodeum.common.v1.SeenResult.storage:type_name -> nodeum.common.v1.Storage
	10, // 11: nodeum.common.v1.SeenResult.info:type_name -> nodeum.common.v1.NodeInfo
	9,  // 12: nodeum.common.v1.FailedResult.storage:type_name -> nodeum.common.v1.Storage
	0,  // 13: nodeum.common.v1.FailedResult.error:type_name -> nodeum.common.v1.FailedResult.ErrorType
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_nodeum_common_v1_result_proto_init() }
func file_nodeum_common_v1_result_proto_init() {
	if File_nodeum_common_v1_result_proto != nil {
		return
	}
	file_nodeum_common_v1_nodeinfo_proto_init()
	file_nodeum_common_v1_request_proto_init()
	file_nodeum_common_v1_storages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nodeum_common_v1_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_nodeum_common_v1_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressResult); i {
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
		file_nodeum_common_v1_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInResult); i {
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
		file_nodeum_common_v1_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOutResult); i {
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
		file_nodeum_common_v1_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeenResult); i {
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
		file_nodeum_common_v1_result_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverResult); i {
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
		file_nodeum_common_v1_result_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailedResult); i {
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
	file_nodeum_common_v1_result_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Result_Progress)(nil),
		(*Result_Checkin)(nil),
		(*Result_Checkout)(nil),
		(*Result_Seen)(nil),
		(*Result_Over)(nil),
		(*Result_Failed)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nodeum_common_v1_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nodeum_common_v1_result_proto_goTypes,
		DependencyIndexes: file_nodeum_common_v1_result_proto_depIdxs,
		EnumInfos:         file_nodeum_common_v1_result_proto_enumTypes,
		MessageInfos:      file_nodeum_common_v1_result_proto_msgTypes,
	}.Build()
	File_nodeum_common_v1_result_proto = out.File
	file_nodeum_common_v1_result_proto_rawDesc = nil
	file_nodeum_common_v1_result_proto_goTypes = nil
	file_nodeum_common_v1_result_proto_depIdxs = nil
}
