// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: rapid/public_api/v1/public.proto

package public_apiv1

import (
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rapid_public_api_v1_public_proto protoreflect.FileDescriptor

var file_rapid_public_api_v1_public_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x72,
	0x61, 0x70, 0x69, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x28, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbe, 0x03, 0x0a, 0x0f, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x56, 0x31, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01,
	0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x65, 0x65, 0x70, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x31, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x44, 0x65, 0x65, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x72, 0x61, 0x70,
	0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x65, 0x65, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x65, 0x70,
	0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x7b, 0x0a,
	0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x28, 0x2e,
	0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x12, 0x8d, 0x01, 0x0a, 0x0f, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x2b,
	0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x61,
	0x70, 0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0xf4, 0x01, 0x0a, 0x17, 0x63,
	0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x62, 0x79, 0x73, 0x73, 0x70, 0x61, 0x72, 0x61, 0x6e, 0x6f, 0x69, 0x61, 0x2f,
	0x72, 0x61, 0x70, 0x69, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x50, 0x58, 0xaa,
	0x02, 0x12, 0x52, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x52, 0x61, 0x70, 0x69, 0x64, 0x5c, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x52, 0x61, 0x70, 0x69,
	0x64, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x52, 0x61, 0x70,
	0x69, 0x64, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_rapid_public_api_v1_public_proto_goTypes = []interface{}{
	(*PublicDeepHealthCheckRequest)(nil),  // 0: rapid.public_api.v1.PublicDeepHealthCheckRequest
	(*PublicSignInRequest)(nil),           // 1: rapid.public_api.v1.PublicSignInRequest
	(*PublicGetTenantRequest)(nil),        // 2: rapid.public_api.v1.PublicGetTenantRequest
	(*PublicDeepHealthCheckResponse)(nil), // 3: rapid.public_api.v1.PublicDeepHealthCheckResponse
	(*PublicSignInResponse)(nil),          // 4: rapid.public_api.v1.PublicSignInResponse
	(*PublicGetTenantResponse)(nil),       // 5: rapid.public_api.v1.PublicGetTenantResponse
}
var file_rapid_public_api_v1_public_proto_depIdxs = []int32{
	0, // 0: rapid.public_api.v1.PublicV1Service.PublicDeepHealthCheck:input_type -> rapid.public_api.v1.PublicDeepHealthCheckRequest
	1, // 1: rapid.public_api.v1.PublicV1Service.PublicSignIn:input_type -> rapid.public_api.v1.PublicSignInRequest
	2, // 2: rapid.public_api.v1.PublicV1Service.PublicGetTenant:input_type -> rapid.public_api.v1.PublicGetTenantRequest
	3, // 3: rapid.public_api.v1.PublicV1Service.PublicDeepHealthCheck:output_type -> rapid.public_api.v1.PublicDeepHealthCheckResponse
	4, // 4: rapid.public_api.v1.PublicV1Service.PublicSignIn:output_type -> rapid.public_api.v1.PublicSignInResponse
	5, // 5: rapid.public_api.v1.PublicV1Service.PublicGetTenant:output_type -> rapid.public_api.v1.PublicGetTenantResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rapid_public_api_v1_public_proto_init() }
func file_rapid_public_api_v1_public_proto_init() {
	if File_rapid_public_api_v1_public_proto != nil {
		return
	}
	file_rapid_public_api_v1_health_check_proto_init()
	file_rapid_public_api_v1_tenant_proto_init()
	file_rapid_public_api_v1_authentication_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rapid_public_api_v1_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rapid_public_api_v1_public_proto_goTypes,
		DependencyIndexes: file_rapid_public_api_v1_public_proto_depIdxs,
	}.Build()
	File_rapid_public_api_v1_public_proto = out.File
	file_rapid_public_api_v1_public_proto_rawDesc = nil
	file_rapid_public_api_v1_public_proto_goTypes = nil
	file_rapid_public_api_v1_public_proto_depIdxs = nil
}
