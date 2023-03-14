// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: nodeum/micro/v1/monitoring.proto

package microv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Service_Node_Status int32

const (
	Service_Node_STATUS_UNSPECIFIED Service_Node_Status = 0
	Service_Node_STATUS_ALIVE       Service_Node_Status = 1
	Service_Node_STATUS_DEAD        Service_Node_Status = 2
)

// Enum value maps for Service_Node_Status.
var (
	Service_Node_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_ALIVE",
		2: "STATUS_DEAD",
	}
	Service_Node_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_ALIVE":       1,
		"STATUS_DEAD":        2,
	}
)

func (x Service_Node_Status) Enum() *Service_Node_Status {
	p := new(Service_Node_Status)
	*p = x
	return p
}

func (x Service_Node_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service_Node_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeum_micro_v1_monitoring_proto_enumTypes[0].Descriptor()
}

func (Service_Node_Status) Type() protoreflect.EnumType {
	return &file_nodeum_micro_v1_monitoring_proto_enumTypes[0]
}

func (x Service_Node_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service_Node_Status.Descriptor instead.
func (Service_Node_Status) EnumDescriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{4, 0, 0}
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{0}
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Hostname string                 `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Ip       string                 `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Memory   int64                  `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Cpus     int32                  `protobuf:"varint,5,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Time     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	Timezone string                 `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Uptime   int64                  `protobuf:"varint,8,opt,name=uptime,proto3" json:"uptime,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *InfoResponse) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *InfoResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *InfoResponse) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *InfoResponse) GetCpus() int32 {
	if x != nil {
		return x.Cpus
	}
	return 0
}

func (x *InfoResponse) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *InfoResponse) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *InfoResponse) GetUptime() int64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional service name
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{2}
}

func (x *StatsRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp of recording
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// unix timestamp
	Started int64 `protobuf:"varint,2,opt,name=started,proto3" json:"started,omitempty"`
	// in seconds
	Uptime int64 `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// in bytes
	Memory int64 `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	// num threads
	Threads int64 `protobuf:"varint,5,opt,name=threads,proto3" json:"threads,omitempty"`
	// total gc in nanoseconds
	Gc int64 `protobuf:"varint,6,opt,name=gc,proto3" json:"gc,omitempty"`
	// total number of requests
	Requests int64 `protobuf:"varint,7,opt,name=requests,proto3" json:"requests,omitempty"`
	// total number of errors
	Errors int64 `protobuf:"varint,8,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{3}
}

func (x *StatsResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *StatsResponse) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *StatsResponse) GetUptime() int64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *StatsResponse) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *StatsResponse) GetThreads() int64 {
	if x != nil {
		return x.Threads
	}
	return 0
}

func (x *StatsResponse) GetGc() int64 {
	if x != nil {
		return x.Gc
	}
	return 0
}

func (x *StatsResponse) GetRequests() int64 {
	if x != nil {
		return x.Requests
	}
	return 0
}

func (x *StatsResponse) GetErrors() int64 {
	if x != nil {
		return x.Errors
	}
	return 0
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string          `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Nodes   []*Service_Node `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{4}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service) GetNodes() []*Service_Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type ListServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListServicesRequest) Reset() {
	*x = ListServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesRequest) ProtoMessage() {}

func (x *ListServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesRequest.ProtoReflect.Descriptor instead.
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{5}
}

type ListServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListServicesResponse) Reset() {
	*x = ListServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesResponse) ProtoMessage() {}

func (x *ListServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesResponse.ProtoReflect.Descriptor instead.
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{6}
}

func (x *ListServicesResponse) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type ListHostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHostsRequest) Reset() {
	*x = ListHostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostsRequest) ProtoMessage() {}

func (x *ListHostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostsRequest.ProtoReflect.Descriptor instead.
func (*ListHostsRequest) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{7}
}

type ListHostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts []*InfoResponse `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
}

func (x *ListHostsResponse) Reset() {
	*x = ListHostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostsResponse) ProtoMessage() {}

func (x *ListHostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostsResponse.ProtoReflect.Descriptor instead.
func (*ListHostsResponse) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{8}
}

func (x *ListHostsResponse) GetHosts() []*InfoResponse {
	if x != nil {
		return x.Hosts
	}
	return nil
}

type Service_Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string              `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Stats   *StatsResponse      `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	Status  Service_Node_Status `protobuf:"varint,4,opt,name=status,proto3,enum=nodeum.micro.v1.Service_Node_Status" json:"status,omitempty"`
}

func (x *Service_Node) Reset() {
	*x = Service_Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service_Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service_Node) ProtoMessage() {}

func (x *Service_Node) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_micro_v1_monitoring_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service_Node.ProtoReflect.Descriptor instead.
func (*Service_Node) Descriptor() ([]byte, []int) {
	return file_nodeum_micro_v1_monitoring_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Service_Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Service_Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Service_Node) GetStats() *StatsResponse {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Service_Node) GetStatus() Service_Node_Status {
	if x != nil {
		return x.Status
	}
	return Service_Node_STATUS_UNSPECIFIED
}

var File_nodeum_micro_v1_monitoring_proto protoreflect.FileDescriptor

var file_nodeum_micro_v1_monitoring_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xde, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x70,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x70, 0x75, 0x73, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xd5, 0x01, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x67, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x67, 0x63,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0xd8, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x1a, 0xe9, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x02, 0x22,
	0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x68, 0x6f, 0x73,
	0x74, 0x73, 0x32, 0x9f, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0xe7, 0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12,
	0x09, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_nodeum_micro_v1_monitoring_proto_rawDescOnce sync.Once
	file_nodeum_micro_v1_monitoring_proto_rawDescData = file_nodeum_micro_v1_monitoring_proto_rawDesc
)

func file_nodeum_micro_v1_monitoring_proto_rawDescGZIP() []byte {
	file_nodeum_micro_v1_monitoring_proto_rawDescOnce.Do(func() {
		file_nodeum_micro_v1_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeum_micro_v1_monitoring_proto_rawDescData)
	})
	return file_nodeum_micro_v1_monitoring_proto_rawDescData
}

var file_nodeum_micro_v1_monitoring_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nodeum_micro_v1_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_nodeum_micro_v1_monitoring_proto_goTypes = []interface{}{
	(Service_Node_Status)(0),      // 0: nodeum.micro.v1.Service.Node.Status
	(*InfoRequest)(nil),           // 1: nodeum.micro.v1.InfoRequest
	(*InfoResponse)(nil),          // 2: nodeum.micro.v1.InfoResponse
	(*StatsRequest)(nil),          // 3: nodeum.micro.v1.StatsRequest
	(*StatsResponse)(nil),         // 4: nodeum.micro.v1.StatsResponse
	(*Service)(nil),               // 5: nodeum.micro.v1.Service
	(*ListServicesRequest)(nil),   // 6: nodeum.micro.v1.ListServicesRequest
	(*ListServicesResponse)(nil),  // 7: nodeum.micro.v1.ListServicesResponse
	(*ListHostsRequest)(nil),      // 8: nodeum.micro.v1.ListHostsRequest
	(*ListHostsResponse)(nil),     // 9: nodeum.micro.v1.ListHostsResponse
	(*Service_Node)(nil),          // 10: nodeum.micro.v1.Service.Node
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_nodeum_micro_v1_monitoring_proto_depIdxs = []int32{
	11, // 0: nodeum.micro.v1.InfoResponse.time:type_name -> google.protobuf.Timestamp
	10, // 1: nodeum.micro.v1.Service.nodes:type_name -> nodeum.micro.v1.Service.Node
	5,  // 2: nodeum.micro.v1.ListServicesResponse.services:type_name -> nodeum.micro.v1.Service
	2,  // 3: nodeum.micro.v1.ListHostsResponse.hosts:type_name -> nodeum.micro.v1.InfoResponse
	4,  // 4: nodeum.micro.v1.Service.Node.stats:type_name -> nodeum.micro.v1.StatsResponse
	0,  // 5: nodeum.micro.v1.Service.Node.status:type_name -> nodeum.micro.v1.Service.Node.Status
	1,  // 6: nodeum.micro.v1.DebugService.Info:input_type -> nodeum.micro.v1.InfoRequest
	3,  // 7: nodeum.micro.v1.DebugService.Stats:input_type -> nodeum.micro.v1.StatsRequest
	6,  // 8: nodeum.micro.v1.MonitoringService.ListServices:input_type -> nodeum.micro.v1.ListServicesRequest
	8,  // 9: nodeum.micro.v1.MonitoringService.ListHosts:input_type -> nodeum.micro.v1.ListHostsRequest
	2,  // 10: nodeum.micro.v1.DebugService.Info:output_type -> nodeum.micro.v1.InfoResponse
	4,  // 11: nodeum.micro.v1.DebugService.Stats:output_type -> nodeum.micro.v1.StatsResponse
	7,  // 12: nodeum.micro.v1.MonitoringService.ListServices:output_type -> nodeum.micro.v1.ListServicesResponse
	9,  // 13: nodeum.micro.v1.MonitoringService.ListHosts:output_type -> nodeum.micro.v1.ListHostsResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_nodeum_micro_v1_monitoring_proto_init() }
func file_nodeum_micro_v1_monitoring_proto_init() {
	if File_nodeum_micro_v1_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodeum_micro_v1_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsResponse); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesRequest); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicesResponse); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostsRequest); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostsResponse); i {
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
		file_nodeum_micro_v1_monitoring_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service_Node); i {
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
			RawDescriptor: file_nodeum_micro_v1_monitoring_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_nodeum_micro_v1_monitoring_proto_goTypes,
		DependencyIndexes: file_nodeum_micro_v1_monitoring_proto_depIdxs,
		EnumInfos:         file_nodeum_micro_v1_monitoring_proto_enumTypes,
		MessageInfos:      file_nodeum_micro_v1_monitoring_proto_msgTypes,
	}.Build()
	File_nodeum_micro_v1_monitoring_proto = out.File
	file_nodeum_micro_v1_monitoring_proto_rawDesc = nil
	file_nodeum_micro_v1_monitoring_proto_goTypes = nil
	file_nodeum_micro_v1_monitoring_proto_depIdxs = nil
}
