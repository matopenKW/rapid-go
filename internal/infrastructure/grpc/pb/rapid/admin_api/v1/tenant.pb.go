// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: rapid/admin_api/v1/tenant.proto

package admin_apiv1

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/abyssparanoia/rapid-go/internal/infrastructure/grpc/pb/rapid/model/v1"
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

type AdminGetTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *AdminGetTenantRequest) Reset() {
	*x = AdminGetTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetTenantRequest) ProtoMessage() {}

func (x *AdminGetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetTenantRequest.ProtoReflect.Descriptor instead.
func (*AdminGetTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *AdminGetTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type AdminGetTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *v1.Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *AdminGetTenantResponse) Reset() {
	*x = AdminGetTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetTenantResponse) ProtoMessage() {}

func (x *AdminGetTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetTenantResponse.ProtoReflect.Descriptor instead.
func (*AdminGetTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *AdminGetTenantResponse) GetTenant() *v1.Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type AdminListTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *AdminListTenantsRequest) Reset() {
	*x = AdminListTenantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminListTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminListTenantsRequest) ProtoMessage() {}

func (x *AdminListTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminListTenantsRequest.ProtoReflect.Descriptor instead.
func (*AdminListTenantsRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *AdminListTenantsRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AdminListTenantsRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminListTenantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants    []*v1.Tenant   `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	Pagination *v1.Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *AdminListTenantsResponse) Reset() {
	*x = AdminListTenantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminListTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminListTenantsResponse) ProtoMessage() {}

func (x *AdminListTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminListTenantsResponse.ProtoReflect.Descriptor instead.
func (*AdminListTenantsResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *AdminListTenantsResponse) GetTenants() []*v1.Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *AdminListTenantsResponse) GetPagination() *v1.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type AdminCreateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AdminCreateTenantRequest) Reset() {
	*x = AdminCreateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateTenantRequest) ProtoMessage() {}

func (x *AdminCreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateTenantRequest.ProtoReflect.Descriptor instead.
func (*AdminCreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *AdminCreateTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AdminCreateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *v1.Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *AdminCreateTenantResponse) Reset() {
	*x = AdminCreateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateTenantResponse) ProtoMessage() {}

func (x *AdminCreateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateTenantResponse.ProtoReflect.Descriptor instead.
func (*AdminCreateTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *AdminCreateTenantResponse) GetTenant() *v1.Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type AdminUpdateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string  `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Name     *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *AdminUpdateTenantRequest) Reset() {
	*x = AdminUpdateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateTenantRequest) ProtoMessage() {}

func (x *AdminUpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*AdminUpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{6}
}

func (x *AdminUpdateTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *AdminUpdateTenantRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type AdminUpdateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *v1.Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *AdminUpdateTenantResponse) Reset() {
	*x = AdminUpdateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateTenantResponse) ProtoMessage() {}

func (x *AdminUpdateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateTenantResponse.ProtoReflect.Descriptor instead.
func (*AdminUpdateTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{7}
}

func (x *AdminUpdateTenantResponse) GetTenant() *v1.Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type AdminDeleteTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *AdminDeleteTenantRequest) Reset() {
	*x = AdminDeleteTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteTenantRequest) ProtoMessage() {}

func (x *AdminDeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*AdminDeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{8}
}

func (x *AdminDeleteTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type AdminDeleteTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminDeleteTenantResponse) Reset() {
	*x = AdminDeleteTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteTenantResponse) ProtoMessage() {}

func (x *AdminDeleteTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_tenant_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteTenantResponse.ProtoReflect.Descriptor instead.
func (*AdminDeleteTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_tenant_proto_rawDescGZIP(), []int{9}
}

var File_rapid_admin_api_v1_tenant_proto protoreflect.FileDescriptor

var file_rapid_admin_api_v1_tenant_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x3a, 0x11, 0x92, 0x41, 0x0e, 0x0a, 0x0c,
	0xd2, 0x01, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x16,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a, 0x0e, 0x92, 0x41, 0x0b, 0x0a, 0x09, 0xd2, 0x01, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xa6, 0x01, 0x0a, 0x18,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x70, 0x69,
	0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x1c, 0x92, 0x41, 0x19, 0x0a, 0x17, 0xd2, 0x01, 0x07,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0xd2, 0x01, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x18, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x0c, 0x92, 0x41, 0x09, 0x0a, 0x07, 0xd2, 0x01, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x19, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a,
	0x0e, 0x92, 0x41, 0x0b, 0x0a, 0x09, 0xd2, 0x01, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22,
	0x6c, 0x0a, 0x18, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x3a, 0x11, 0x92, 0x41, 0x0e, 0x0a, 0x0c, 0xd2, 0x01, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a,
	0x19, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x70,
	0x69, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a, 0x0e, 0x92, 0x41, 0x0b, 0x0a,
	0x09, 0xd2, 0x01, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x18, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x3a, 0x11, 0x92, 0x41, 0x0e, 0x0a, 0x0c, 0xd2, 0x01, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0xed, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x70, 0x69,
	0x64, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x79, 0x73, 0x73, 0x70,
	0x61, 0x72, 0x61, 0x6e, 0x6f, 0x69, 0x61, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2d, 0x67, 0x6f,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x41, 0x58, 0xaa, 0x02, 0x11, 0x52, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x52, 0x61, 0x70, 0x69,
	0x64, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d,
	0x52, 0x61, 0x70, 0x69, 0x64, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x52, 0x61, 0x70, 0x69, 0x64, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rapid_admin_api_v1_tenant_proto_rawDescOnce sync.Once
	file_rapid_admin_api_v1_tenant_proto_rawDescData = file_rapid_admin_api_v1_tenant_proto_rawDesc
)

func file_rapid_admin_api_v1_tenant_proto_rawDescGZIP() []byte {
	file_rapid_admin_api_v1_tenant_proto_rawDescOnce.Do(func() {
		file_rapid_admin_api_v1_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_rapid_admin_api_v1_tenant_proto_rawDescData)
	})
	return file_rapid_admin_api_v1_tenant_proto_rawDescData
}

var file_rapid_admin_api_v1_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rapid_admin_api_v1_tenant_proto_goTypes = []interface{}{
	(*AdminGetTenantRequest)(nil),     // 0: rapid.admin_api.v1.AdminGetTenantRequest
	(*AdminGetTenantResponse)(nil),    // 1: rapid.admin_api.v1.AdminGetTenantResponse
	(*AdminListTenantsRequest)(nil),   // 2: rapid.admin_api.v1.AdminListTenantsRequest
	(*AdminListTenantsResponse)(nil),  // 3: rapid.admin_api.v1.AdminListTenantsResponse
	(*AdminCreateTenantRequest)(nil),  // 4: rapid.admin_api.v1.AdminCreateTenantRequest
	(*AdminCreateTenantResponse)(nil), // 5: rapid.admin_api.v1.AdminCreateTenantResponse
	(*AdminUpdateTenantRequest)(nil),  // 6: rapid.admin_api.v1.AdminUpdateTenantRequest
	(*AdminUpdateTenantResponse)(nil), // 7: rapid.admin_api.v1.AdminUpdateTenantResponse
	(*AdminDeleteTenantRequest)(nil),  // 8: rapid.admin_api.v1.AdminDeleteTenantRequest
	(*AdminDeleteTenantResponse)(nil), // 9: rapid.admin_api.v1.AdminDeleteTenantResponse
	(*v1.Tenant)(nil),                 // 10: rapid.model.v1.Tenant
	(*v1.Pagination)(nil),             // 11: rapid.model.v1.Pagination
}
var file_rapid_admin_api_v1_tenant_proto_depIdxs = []int32{
	10, // 0: rapid.admin_api.v1.AdminGetTenantResponse.tenant:type_name -> rapid.model.v1.Tenant
	10, // 1: rapid.admin_api.v1.AdminListTenantsResponse.tenants:type_name -> rapid.model.v1.Tenant
	11, // 2: rapid.admin_api.v1.AdminListTenantsResponse.pagination:type_name -> rapid.model.v1.Pagination
	10, // 3: rapid.admin_api.v1.AdminCreateTenantResponse.tenant:type_name -> rapid.model.v1.Tenant
	10, // 4: rapid.admin_api.v1.AdminUpdateTenantResponse.tenant:type_name -> rapid.model.v1.Tenant
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rapid_admin_api_v1_tenant_proto_init() }
func file_rapid_admin_api_v1_tenant_proto_init() {
	if File_rapid_admin_api_v1_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rapid_admin_api_v1_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetTenantRequest); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetTenantResponse); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminListTenantsRequest); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminListTenantsResponse); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateTenantRequest); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateTenantResponse); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateTenantRequest); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateTenantResponse); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteTenantRequest); i {
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
		file_rapid_admin_api_v1_tenant_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteTenantResponse); i {
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
	file_rapid_admin_api_v1_tenant_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rapid_admin_api_v1_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rapid_admin_api_v1_tenant_proto_goTypes,
		DependencyIndexes: file_rapid_admin_api_v1_tenant_proto_depIdxs,
		MessageInfos:      file_rapid_admin_api_v1_tenant_proto_msgTypes,
	}.Build()
	File_rapid_admin_api_v1_tenant_proto = out.File
	file_rapid_admin_api_v1_tenant_proto_rawDesc = nil
	file_rapid_admin_api_v1_tenant_proto_goTypes = nil
	file_rapid_admin_api_v1_tenant_proto_depIdxs = nil
}