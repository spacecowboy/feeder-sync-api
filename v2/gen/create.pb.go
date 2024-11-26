// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: feeder/v2/create.proto

package feeder_v2

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

type CreateChainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName string `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
}

func (x *CreateChainRequest) Reset() {
	*x = CreateChainRequest{}
	mi := &file_feeder_v2_create_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChainRequest) ProtoMessage() {}

func (x *CreateChainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeder_v2_create_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChainRequest.ProtoReflect.Descriptor instead.
func (*CreateChainRequest) Descriptor() ([]byte, []int) {
	return file_feeder_v2_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChainRequest) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

type CreateChainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Device *Device `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreateChainResponse) Reset() {
	*x = CreateChainResponse{}
	mi := &file_feeder_v2_create_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChainResponse) ProtoMessage() {}

func (x *CreateChainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeder_v2_create_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChainResponse.ProtoReflect.Descriptor instead.
func (*CreateChainResponse) Descriptor() ([]byte, []int) {
	return file_feeder_v2_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChainResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateChainResponse) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_feeder_v2_create_proto protoreflect.FileDescriptor

var file_feeder_v2_create_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x1a, 0x16, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x66, 0x65, 0x65,
	0x64, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x35, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x63, 0x6f, 0x77, 0x62, 0x6f, 0x79, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72,
	0x2d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x6e,
	0x3b, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_feeder_v2_create_proto_rawDescOnce sync.Once
	file_feeder_v2_create_proto_rawDescData = file_feeder_v2_create_proto_rawDesc
)

func file_feeder_v2_create_proto_rawDescGZIP() []byte {
	file_feeder_v2_create_proto_rawDescOnce.Do(func() {
		file_feeder_v2_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_feeder_v2_create_proto_rawDescData)
	})
	return file_feeder_v2_create_proto_rawDescData
}

var file_feeder_v2_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feeder_v2_create_proto_goTypes = []any{
	(*CreateChainRequest)(nil),  // 0: feeder.v2.CreateChainRequest
	(*CreateChainResponse)(nil), // 1: feeder.v2.CreateChainResponse
	(*User)(nil),                // 2: feeder.v2.User
	(*Device)(nil),              // 3: feeder.v2.Device
}
var file_feeder_v2_create_proto_depIdxs = []int32{
	2, // 0: feeder.v2.CreateChainResponse.user:type_name -> feeder.v2.User
	3, // 1: feeder.v2.CreateChainResponse.device:type_name -> feeder.v2.Device
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_feeder_v2_create_proto_init() }
func file_feeder_v2_create_proto_init() {
	if File_feeder_v2_create_proto != nil {
		return
	}
	file_feeder_v2_device_proto_init()
	file_feeder_v2_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feeder_v2_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feeder_v2_create_proto_goTypes,
		DependencyIndexes: file_feeder_v2_create_proto_depIdxs,
		MessageInfos:      file_feeder_v2_create_proto_msgTypes,
	}.Build()
	File_feeder_v2_create_proto = out.File
	file_feeder_v2_create_proto_rawDesc = nil
	file_feeder_v2_create_proto_goTypes = nil
	file_feeder_v2_create_proto_depIdxs = nil
}
