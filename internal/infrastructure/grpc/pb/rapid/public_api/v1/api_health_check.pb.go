// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: rapid/public_api/v1/api_health_check.proto

package public_apiv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeepHealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeepHealthCheckRequest) Reset() {
	*x = DeepHealthCheckRequest{}
	mi := &file_rapid_public_api_v1_api_health_check_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeepHealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeepHealthCheckRequest) ProtoMessage() {}

func (x *DeepHealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_public_api_v1_api_health_check_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeepHealthCheckRequest.ProtoReflect.Descriptor instead.
func (*DeepHealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_rapid_public_api_v1_api_health_check_proto_rawDescGZIP(), []int{0}
}

type DeepHealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseStatus string `protobuf:"bytes,1,opt,name=database_status,json=databaseStatus,proto3" json:"database_status,omitempty"`
}

func (x *DeepHealthCheckResponse) Reset() {
	*x = DeepHealthCheckResponse{}
	mi := &file_rapid_public_api_v1_api_health_check_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeepHealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeepHealthCheckResponse) ProtoMessage() {}

func (x *DeepHealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_public_api_v1_api_health_check_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeepHealthCheckResponse.ProtoReflect.Descriptor instead.
func (*DeepHealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_rapid_public_api_v1_api_health_check_proto_rawDescGZIP(), []int{1}
}

func (x *DeepHealthCheckResponse) GetDatabaseStatus() string {
	if x != nil {
		return x.DatabaseStatus
	}
	return ""
}

var File_rapid_public_api_v1_api_health_check_proto protoreflect.FileDescriptor

var file_rapid_public_api_v1_api_health_check_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x72, 0x61,
	0x70, 0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x65, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x17, 0x44,
	0x65, 0x65, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a,
	0x17, 0x92, 0x41, 0x14, 0x0a, 0x12, 0xd2, 0x01, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0xfc, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x41, 0x70, 0x69, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x79, 0x73, 0x73, 0x70, 0x61, 0x72,
	0x61, 0x6e, 0x6f, 0x69, 0x61, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x61, 0x70, 0x69, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x50, 0x58, 0xaa, 0x02, 0x12, 0x52, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x52, 0x61, 0x70,
	0x69, 0x64, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1e, 0x52, 0x61, 0x70, 0x69, 0x64, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x14, 0x52, 0x61, 0x70, 0x69, 0x64, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rapid_public_api_v1_api_health_check_proto_rawDescOnce sync.Once
	file_rapid_public_api_v1_api_health_check_proto_rawDescData = file_rapid_public_api_v1_api_health_check_proto_rawDesc
)

func file_rapid_public_api_v1_api_health_check_proto_rawDescGZIP() []byte {
	file_rapid_public_api_v1_api_health_check_proto_rawDescOnce.Do(func() {
		file_rapid_public_api_v1_api_health_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_rapid_public_api_v1_api_health_check_proto_rawDescData)
	})
	return file_rapid_public_api_v1_api_health_check_proto_rawDescData
}

var file_rapid_public_api_v1_api_health_check_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rapid_public_api_v1_api_health_check_proto_goTypes = []any{
	(*DeepHealthCheckRequest)(nil),  // 0: rapid.public_api.v1.DeepHealthCheckRequest
	(*DeepHealthCheckResponse)(nil), // 1: rapid.public_api.v1.DeepHealthCheckResponse
}
var file_rapid_public_api_v1_api_health_check_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rapid_public_api_v1_api_health_check_proto_init() }
func file_rapid_public_api_v1_api_health_check_proto_init() {
	if File_rapid_public_api_v1_api_health_check_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rapid_public_api_v1_api_health_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rapid_public_api_v1_api_health_check_proto_goTypes,
		DependencyIndexes: file_rapid_public_api_v1_api_health_check_proto_depIdxs,
		MessageInfos:      file_rapid_public_api_v1_api_health_check_proto_msgTypes,
	}.Build()
	File_rapid_public_api_v1_api_health_check_proto = out.File
	file_rapid_public_api_v1_api_health_check_proto_rawDesc = nil
	file_rapid_public_api_v1_api_health_check_proto_goTypes = nil
	file_rapid_public_api_v1_api_health_check_proto_depIdxs = nil
}
