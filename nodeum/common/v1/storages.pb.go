// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: nodeum/common/v1/storages.proto

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

type Bucket_Provider int32

const (
	Bucket_PROVIDER_UNSPECIFIED          Bucket_Provider = 0
	Bucket_PROVIDER_GENERIC_S3           Bucket_Provider = 1
	Bucket_PROVIDER_AMAZON_AWS_S3        Bucket_Provider = 2
	Bucket_PROVIDER_CLOUDIAN_HYPERSTORE  Bucket_Provider = 3
	Bucket_PROVIDER_SCALITY_RING         Bucket_Provider = 4
	Bucket_PROVIDER_DELL_EMC_ECS         Bucket_Provider = 5
	Bucket_PROVIDER_AZURE                Bucket_Provider = 6
	Bucket_PROVIDER_GOOGLE_CLOUD_STORAGE Bucket_Provider = 7
	Bucket_PROVIDER_OPENSTACK_SWIFT      Bucket_Provider = 8
	Bucket_PROVIDER_WASABI               Bucket_Provider = 9
	Bucket_PROVIDER_QUANTUM_ACTIVE_SCALE Bucket_Provider = 10
)

// Enum value maps for Bucket_Provider.
var (
	Bucket_Provider_name = map[int32]string{
		0:  "PROVIDER_UNSPECIFIED",
		1:  "PROVIDER_GENERIC_S3",
		2:  "PROVIDER_AMAZON_AWS_S3",
		3:  "PROVIDER_CLOUDIAN_HYPERSTORE",
		4:  "PROVIDER_SCALITY_RING",
		5:  "PROVIDER_DELL_EMC_ECS",
		6:  "PROVIDER_AZURE",
		7:  "PROVIDER_GOOGLE_CLOUD_STORAGE",
		8:  "PROVIDER_OPENSTACK_SWIFT",
		9:  "PROVIDER_WASABI",
		10: "PROVIDER_QUANTUM_ACTIVE_SCALE",
	}
	Bucket_Provider_value = map[string]int32{
		"PROVIDER_UNSPECIFIED":          0,
		"PROVIDER_GENERIC_S3":           1,
		"PROVIDER_AMAZON_AWS_S3":        2,
		"PROVIDER_CLOUDIAN_HYPERSTORE":  3,
		"PROVIDER_SCALITY_RING":         4,
		"PROVIDER_DELL_EMC_ECS":         5,
		"PROVIDER_AZURE":                6,
		"PROVIDER_GOOGLE_CLOUD_STORAGE": 7,
		"PROVIDER_OPENSTACK_SWIFT":      8,
		"PROVIDER_WASABI":               9,
		"PROVIDER_QUANTUM_ACTIVE_SCALE": 10,
	}
)

func (x Bucket_Provider) Enum() *Bucket_Provider {
	p := new(Bucket_Provider)
	*p = x
	return p
}

func (x Bucket_Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Bucket_Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeum_common_v1_storages_proto_enumTypes[0].Descriptor()
}

func (Bucket_Provider) Type() protoreflect.EnumType {
	return &file_nodeum_common_v1_storages_proto_enumTypes[0]
}

func (x Bucket_Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Bucket_Provider.Descriptor instead.
func (Bucket_Provider) EnumDescriptor() ([]byte, []int) {
	return file_nodeum_common_v1_storages_proto_rawDescGZIP(), []int{0, 0}
}

type NASShare_Protocol int32

const (
	NASShare_PROTOCOL_UNSPECIFIED   NASShare_Protocol = 0
	NASShare_PROTOCOL_SMB_V1        NASShare_Protocol = 1
	NASShare_PROTOCOL_NFS_V3        NASShare_Protocol = 2
	NASShare_PROTOCOL_STORE_NEXT_V5 NASShare_Protocol = 3
	NASShare_PROTOCOL_NFS_V4        NASShare_Protocol = 4
	NASShare_PROTOCOL_SMB_V2_1      NASShare_Protocol = 5
	NASShare_PROTOCOL_SMB_V3        NASShare_Protocol = 6
)

// Enum value maps for NASShare_Protocol.
var (
	NASShare_Protocol_name = map[int32]string{
		0: "PROTOCOL_UNSPECIFIED",
		1: "PROTOCOL_SMB_V1",
		2: "PROTOCOL_NFS_V3",
		3: "PROTOCOL_STORE_NEXT_V5",
		4: "PROTOCOL_NFS_V4",
		5: "PROTOCOL_SMB_V2_1",
		6: "PROTOCOL_SMB_V3",
	}
	NASShare_Protocol_value = map[string]int32{
		"PROTOCOL_UNSPECIFIED":   0,
		"PROTOCOL_SMB_V1":        1,
		"PROTOCOL_NFS_V3":        2,
		"PROTOCOL_STORE_NEXT_V5": 3,
		"PROTOCOL_NFS_V4":        4,
		"PROTOCOL_SMB_V2_1":      5,
		"PROTOCOL_SMB_V3":        6,
	}
)

func (x NASShare_Protocol) Enum() *NASShare_Protocol {
	p := new(NASShare_Protocol)
	*p = x
	return p
}

func (x NASShare_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NASShare_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeum_common_v1_storages_proto_enumTypes[1].Descriptor()
}

func (NASShare_Protocol) Type() protoreflect.EnumType {
	return &file_nodeum_common_v1_storages_proto_enumTypes[1]
}

func (x NASShare_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NASShare_Protocol.Descriptor instead.
func (NASShare_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_nodeum_common_v1_storages_proto_rawDescGZIP(), []int{1, 0}
}

type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ConnectorName string          `protobuf:"bytes,3,opt,name=connector_name,proto3" json:"connector_name,omitempty"`
	PrimaryName   string          `protobuf:"bytes,4,opt,name=primary_name,proto3" json:"primary_name,omitempty"`
	Url           string          `protobuf:"bytes,16,opt,name=url,proto3" json:"url,omitempty"`
	Provider      Bucket_Provider `protobuf:"varint,17,opt,name=provider,proto3,enum=nodeum.common.v1.Bucket_Provider" json:"provider,omitempty"`
	Region        string          `protobuf:"bytes,18,opt,name=region,proto3" json:"region,omitempty"`
	AccessKey     string          `protobuf:"bytes,19,opt,name=access_key,proto3" json:"access_key,omitempty"`
	SecretKey     string          `protobuf:"bytes,20,opt,name=secret_key,proto3" json:"secret_key,omitempty"`
	Options       string          `protobuf:"bytes,21,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_storages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_storages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_storages_proto_rawDescGZIP(), []int{0}
}

func (x *Bucket) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bucket) GetConnectorName() string {
	if x != nil {
		return x.ConnectorName
	}
	return ""
}

func (x *Bucket) GetPrimaryName() string {
	if x != nil {
		return x.PrimaryName
	}
	return ""
}

func (x *Bucket) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Bucket) GetProvider() Bucket_Provider {
	if x != nil {
		return x.Provider
	}
	return Bucket_PROVIDER_UNSPECIFIED
}

func (x *Bucket) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Bucket) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Bucket) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Bucket) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

type NASShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NasName     string            `protobuf:"bytes,3,opt,name=nas_name,proto3" json:"nas_name,omitempty"`
	PrimaryName string            `protobuf:"bytes,4,opt,name=primary_name,proto3" json:"primary_name,omitempty"`
	Protocol    NASShare_Protocol `protobuf:"varint,16,opt,name=protocol,proto3,enum=nodeum.common.v1.NASShare_Protocol" json:"protocol,omitempty"`
	Host        string            `protobuf:"bytes,17,opt,name=host,proto3" json:"host,omitempty"`
	Path        string            `protobuf:"bytes,18,opt,name=path,proto3" json:"path,omitempty"`
	Username    string            `protobuf:"bytes,19,opt,name=username,proto3" json:"username,omitempty"`
	Password    string            `protobuf:"bytes,20,opt,name=password,proto3" json:"password,omitempty"`
	Options     string            `protobuf:"bytes,21,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *NASShare) Reset() {
	*x = NASShare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_storages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NASShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NASShare) ProtoMessage() {}

func (x *NASShare) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_storages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NASShare.ProtoReflect.Descriptor instead.
func (*NASShare) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_storages_proto_rawDescGZIP(), []int{1}
}

func (x *NASShare) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NASShare) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NASShare) GetNasName() string {
	if x != nil {
		return x.NasName
	}
	return ""
}

func (x *NASShare) GetPrimaryName() string {
	if x != nil {
		return x.PrimaryName
	}
	return ""
}

func (x *NASShare) GetProtocol() NASShare_Protocol {
	if x != nil {
		return x.Protocol
	}
	return NASShare_PROTOCOL_UNSPECIFIED
}

func (x *NASShare) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *NASShare) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *NASShare) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NASShare) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NASShare) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_storages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_storages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_storages_proto_rawDescGZIP(), []int{2}
}

func (x *Container) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Storage:
	//	*Storage_Bucket
	//	*Storage_NasShare
	//	*Storage_Container
	Storage isStorage_Storage `protobuf_oneof:"storage"`
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeum_common_v1_storages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_nodeum_common_v1_storages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_nodeum_common_v1_storages_proto_rawDescGZIP(), []int{3}
}

func (m *Storage) GetStorage() isStorage_Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (x *Storage) GetBucket() *Bucket {
	if x, ok := x.GetStorage().(*Storage_Bucket); ok {
		return x.Bucket
	}
	return nil
}

func (x *Storage) GetNasShare() *NASShare {
	if x, ok := x.GetStorage().(*Storage_NasShare); ok {
		return x.NasShare
	}
	return nil
}

func (x *Storage) GetContainer() *Container {
	if x, ok := x.GetStorage().(*Storage_Container); ok {
		return x.Container
	}
	return nil
}

type isStorage_Storage interface {
	isStorage_Storage()
}

type Storage_Bucket struct {
	Bucket *Bucket `protobuf:"bytes,1,opt,name=bucket,proto3,oneof"`
}

type Storage_NasShare struct {
	NasShare *NASShare `protobuf:"bytes,2,opt,name=nas_share,proto3,oneof"`
}

type Storage_Container struct {
	Container *Container `protobuf:"bytes,3,opt,name=container,proto3,oneof"`
}

func (*Storage_Bucket) isStorage_Storage() {}

func (*Storage_NasShare) isStorage_Storage() {}

func (*Storage_Container) isStorage_Storage() {}

var File_nodeum_common_v1_storages_proto protoreflect.FileDescriptor

var file_nodeum_common_v1_storages_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0xfc, 0x04, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x3d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xbe, 0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x14, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x56,
	0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x53, 0x33, 0x10,
	0x01, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x4d,
	0x41, 0x5a, 0x4f, 0x4e, 0x5f, 0x41, 0x57, 0x53, 0x5f, 0x53, 0x33, 0x10, 0x02, 0x12, 0x20, 0x0a,
	0x1c, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x49,
	0x41, 0x4e, 0x5f, 0x48, 0x59, 0x50, 0x45, 0x52, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x03, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x43, 0x41, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52,
	0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x4c, 0x5f, 0x45, 0x4d, 0x43, 0x5f,
	0x45, 0x43, 0x53, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45,
	0x52, 0x5f, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x10, 0x06, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x52, 0x4f,
	0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x43, 0x4c, 0x4f,
	0x55, 0x44, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18,
	0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x53, 0x54, 0x41,
	0x43, 0x4b, 0x5f, 0x53, 0x57, 0x49, 0x46, 0x54, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52,
	0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x57, 0x41, 0x53, 0x41, 0x42, 0x49, 0x10, 0x09, 0x12,
	0x21, 0x0a, 0x1d, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x51, 0x55, 0x41, 0x4e,
	0x54, 0x55, 0x4d, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x45,
	0x10, 0x0a, 0x22, 0xd7, 0x03, 0x0a, 0x08, 0x4e, 0x41, 0x53, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x61, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x41, 0x53, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xab,
	0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x14, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f,
	0x4c, 0x5f, 0x53, 0x4d, 0x42, 0x5f, 0x56, 0x31, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x4e, 0x46, 0x53, 0x5f, 0x56, 0x33, 0x10, 0x02, 0x12,
	0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52,
	0x45, 0x5f, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x56, 0x35, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x4e, 0x46, 0x53, 0x5f, 0x56, 0x34, 0x10, 0x04,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x53, 0x4d, 0x42,
	0x5f, 0x56, 0x32, 0x5f, 0x31, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x54, 0x4f,
	0x43, 0x4f, 0x4c, 0x5f, 0x53, 0x4d, 0x42, 0x5f, 0x56, 0x33, 0x10, 0x06, 0x22, 0x2f, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc1, 0x01,
	0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3a, 0x0a,
	0x09, 0x6e, 0x61, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x41, 0x53, 0x53, 0x68, 0x61, 0x72, 0x65, 0x48, 0x00, 0x52, 0x09,
	0x6e, 0x61, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2d, 0x69, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d,
	0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x75, 0x6d, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodeum_common_v1_storages_proto_rawDescOnce sync.Once
	file_nodeum_common_v1_storages_proto_rawDescData = file_nodeum_common_v1_storages_proto_rawDesc
)

func file_nodeum_common_v1_storages_proto_rawDescGZIP() []byte {
	file_nodeum_common_v1_storages_proto_rawDescOnce.Do(func() {
		file_nodeum_common_v1_storages_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeum_common_v1_storages_proto_rawDescData)
	})
	return file_nodeum_common_v1_storages_proto_rawDescData
}

var file_nodeum_common_v1_storages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nodeum_common_v1_storages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nodeum_common_v1_storages_proto_goTypes = []interface{}{
	(Bucket_Provider)(0),   // 0: nodeum.common.v1.Bucket.Provider
	(NASShare_Protocol)(0), // 1: nodeum.common.v1.NASShare.Protocol
	(*Bucket)(nil),         // 2: nodeum.common.v1.Bucket
	(*NASShare)(nil),       // 3: nodeum.common.v1.NASShare
	(*Container)(nil),      // 4: nodeum.common.v1.Container
	(*Storage)(nil),        // 5: nodeum.common.v1.Storage
}
var file_nodeum_common_v1_storages_proto_depIdxs = []int32{
	0, // 0: nodeum.common.v1.Bucket.provider:type_name -> nodeum.common.v1.Bucket.Provider
	1, // 1: nodeum.common.v1.NASShare.protocol:type_name -> nodeum.common.v1.NASShare.Protocol
	2, // 2: nodeum.common.v1.Storage.bucket:type_name -> nodeum.common.v1.Bucket
	3, // 3: nodeum.common.v1.Storage.nas_share:type_name -> nodeum.common.v1.NASShare
	4, // 4: nodeum.common.v1.Storage.container:type_name -> nodeum.common.v1.Container
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nodeum_common_v1_storages_proto_init() }
func file_nodeum_common_v1_storages_proto_init() {
	if File_nodeum_common_v1_storages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodeum_common_v1_storages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucket); i {
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
		file_nodeum_common_v1_storages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NASShare); i {
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
		file_nodeum_common_v1_storages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_nodeum_common_v1_storages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
	file_nodeum_common_v1_storages_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Storage_Bucket)(nil),
		(*Storage_NasShare)(nil),
		(*Storage_Container)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nodeum_common_v1_storages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nodeum_common_v1_storages_proto_goTypes,
		DependencyIndexes: file_nodeum_common_v1_storages_proto_depIdxs,
		EnumInfos:         file_nodeum_common_v1_storages_proto_enumTypes,
		MessageInfos:      file_nodeum_common_v1_storages_proto_msgTypes,
	}.Build()
	File_nodeum_common_v1_storages_proto = out.File
	file_nodeum_common_v1_storages_proto_rawDesc = nil
	file_nodeum_common_v1_storages_proto_goTypes = nil
	file_nodeum_common_v1_storages_proto_depIdxs = nil
}