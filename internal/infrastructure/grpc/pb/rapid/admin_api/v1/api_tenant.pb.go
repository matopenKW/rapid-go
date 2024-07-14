// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: rapid/admin_api/v1/api_tenant.proto

package admin_apiv1

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

type GetTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *GetTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type GetTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *GetTenantResponse) Reset() {
	*x = GetTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantResponse) ProtoMessage() {}

func (x *GetTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantResponse.ProtoReflect.Descriptor instead.
func (*GetTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *GetTenantResponse) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type ListTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListTenantsRequest) Reset() {
	*x = ListTenantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsRequest) ProtoMessage() {}

func (x *ListTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsRequest.ProtoReflect.Descriptor instead.
func (*ListTenantsRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *ListTenantsRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTenantsRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListTenantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants    []*Tenant   `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListTenantsResponse) Reset() {
	*x = ListTenantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsResponse) ProtoMessage() {}

func (x *ListTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsResponse.ProtoReflect.Descriptor instead.
func (*ListTenantsResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *ListTenantsResponse) GetTenants() []*Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *ListTenantsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *CreateTenantResponse) Reset() {
	*x = CreateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantResponse) ProtoMessage() {}

func (x *CreateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantResponse.ProtoReflect.Descriptor instead.
func (*CreateTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTenantResponse) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type UpdateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string  `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Name     *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *UpdateTenantRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type UpdateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *UpdateTenantResponse) Reset() {
	*x = UpdateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantResponse) ProtoMessage() {}

func (x *UpdateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantResponse.ProtoReflect.Descriptor instead.
func (*UpdateTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTenantResponse) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type DeleteTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *DeleteTenantRequest) Reset() {
	*x = DeleteTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantRequest) ProtoMessage() {}

func (x *DeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type DeleteTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTenantResponse) Reset() {
	*x = DeleteTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantResponse) ProtoMessage() {}

func (x *DeleteTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rapid_admin_api_v1_api_tenant_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantResponse.ProtoReflect.Descriptor instead.
func (*DeleteTenantResponse) Descriptor() ([]byte, []int) {
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP(), []int{9}
}

var File_rapid_admin_api_v1_api_tenant_proto protoreflect.FileDescriptor

var file_rapid_admin_api_v1_api_tenant_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x72, 0x61, 0x70, 0x69, 0x64,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x3a, 0x11, 0x92, 0x41,
	0x0e, 0x0a, 0x0c, 0xd2, 0x01, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x57, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a, 0x0e, 0x92, 0x41, 0x0b, 0x0a, 0x09, 0xd2,
	0x01, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x61, 0x70,
	0x69, 0x64, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x1c, 0x92, 0x41, 0x19, 0x0a, 0x17, 0xd2, 0x01, 0x07,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0xd2, 0x01, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a,
	0x0c, 0x92, 0x41, 0x09, 0x0a, 0x07, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a, 0x0e, 0x92, 0x41, 0x0b, 0x0a, 0x09,
	0xd2, 0x01, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x67, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x11, 0x92, 0x41, 0x0e, 0x0a, 0x0c, 0xd2, 0x01, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x70,
	0x69, 0x64, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a, 0x0e,
	0x92, 0x41, 0x0b, 0x0a, 0x09, 0xd2, 0x01, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x45,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x3a, 0x11, 0x92, 0x41, 0x0e, 0x0a, 0x0c, 0xd2, 0x01, 0x09, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xf0, 0x01,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x41, 0x70, 0x69, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x79, 0x73, 0x73, 0x70, 0x61, 0x72, 0x61,
	0x6e, 0x6f, 0x69, 0x61, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x61,
	0x70, 0x69, 0x64, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52,
	0x41, 0x58, 0xaa, 0x02, 0x11, 0x52, 0x61, 0x70, 0x69, 0x64, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x52, 0x61, 0x70, 0x69, 0x64, 0x5c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x52, 0x61, 0x70,
	0x69, 0x64, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x52, 0x61, 0x70,
	0x69, 0x64, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rapid_admin_api_v1_api_tenant_proto_rawDescOnce sync.Once
	file_rapid_admin_api_v1_api_tenant_proto_rawDescData = file_rapid_admin_api_v1_api_tenant_proto_rawDesc
)

func file_rapid_admin_api_v1_api_tenant_proto_rawDescGZIP() []byte {
	file_rapid_admin_api_v1_api_tenant_proto_rawDescOnce.Do(func() {
		file_rapid_admin_api_v1_api_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_rapid_admin_api_v1_api_tenant_proto_rawDescData)
	})
	return file_rapid_admin_api_v1_api_tenant_proto_rawDescData
}

var file_rapid_admin_api_v1_api_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rapid_admin_api_v1_api_tenant_proto_goTypes = []any{
	(*GetTenantRequest)(nil),     // 0: rapid.admin_api.v1.GetTenantRequest
	(*GetTenantResponse)(nil),    // 1: rapid.admin_api.v1.GetTenantResponse
	(*ListTenantsRequest)(nil),   // 2: rapid.admin_api.v1.ListTenantsRequest
	(*ListTenantsResponse)(nil),  // 3: rapid.admin_api.v1.ListTenantsResponse
	(*CreateTenantRequest)(nil),  // 4: rapid.admin_api.v1.CreateTenantRequest
	(*CreateTenantResponse)(nil), // 5: rapid.admin_api.v1.CreateTenantResponse
	(*UpdateTenantRequest)(nil),  // 6: rapid.admin_api.v1.UpdateTenantRequest
	(*UpdateTenantResponse)(nil), // 7: rapid.admin_api.v1.UpdateTenantResponse
	(*DeleteTenantRequest)(nil),  // 8: rapid.admin_api.v1.DeleteTenantRequest
	(*DeleteTenantResponse)(nil), // 9: rapid.admin_api.v1.DeleteTenantResponse
	(*Tenant)(nil),               // 10: rapid.admin_api.v1.Tenant
	(*Pagination)(nil),           // 11: rapid.admin_api.v1.Pagination
}
var file_rapid_admin_api_v1_api_tenant_proto_depIdxs = []int32{
	10, // 0: rapid.admin_api.v1.GetTenantResponse.tenant:type_name -> rapid.admin_api.v1.Tenant
	10, // 1: rapid.admin_api.v1.ListTenantsResponse.tenants:type_name -> rapid.admin_api.v1.Tenant
	11, // 2: rapid.admin_api.v1.ListTenantsResponse.pagination:type_name -> rapid.admin_api.v1.Pagination
	10, // 3: rapid.admin_api.v1.CreateTenantResponse.tenant:type_name -> rapid.admin_api.v1.Tenant
	10, // 4: rapid.admin_api.v1.UpdateTenantResponse.tenant:type_name -> rapid.admin_api.v1.Tenant
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rapid_admin_api_v1_api_tenant_proto_init() }
func file_rapid_admin_api_v1_api_tenant_proto_init() {
	if File_rapid_admin_api_v1_api_tenant_proto != nil {
		return
	}
	file_rapid_admin_api_v1_model_pagination_proto_init()
	file_rapid_admin_api_v1_model_tenant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTenantRequest); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetTenantResponse); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListTenantsRequest); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListTenantsResponse); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTenantRequest); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTenantResponse); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTenantRequest); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTenantResponse); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTenantRequest); i {
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
		file_rapid_admin_api_v1_api_tenant_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTenantResponse); i {
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
	file_rapid_admin_api_v1_api_tenant_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rapid_admin_api_v1_api_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rapid_admin_api_v1_api_tenant_proto_goTypes,
		DependencyIndexes: file_rapid_admin_api_v1_api_tenant_proto_depIdxs,
		MessageInfos:      file_rapid_admin_api_v1_api_tenant_proto_msgTypes,
	}.Build()
	File_rapid_admin_api_v1_api_tenant_proto = out.File
	file_rapid_admin_api_v1_api_tenant_proto_rawDesc = nil
	file_rapid_admin_api_v1_api_tenant_proto_goTypes = nil
	file_rapid_admin_api_v1_api_tenant_proto_depIdxs = nil
}
