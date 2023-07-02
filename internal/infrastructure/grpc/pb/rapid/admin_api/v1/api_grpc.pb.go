// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: rapid/admin_api/v1/api.proto

package admin_apiv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AdminV1Service_CreateAssetPresignedURL_FullMethodName = "/rapid.admin_api.v1.AdminV1Service/CreateAssetPresignedURL"
	AdminV1Service_GetTenant_FullMethodName               = "/rapid.admin_api.v1.AdminV1Service/GetTenant"
	AdminV1Service_ListTenants_FullMethodName             = "/rapid.admin_api.v1.AdminV1Service/ListTenants"
	AdminV1Service_CreateTenant_FullMethodName            = "/rapid.admin_api.v1.AdminV1Service/CreateTenant"
	AdminV1Service_UpdateTenant_FullMethodName            = "/rapid.admin_api.v1.AdminV1Service/UpdateTenant"
	AdminV1Service_DeleteTenant_FullMethodName            = "/rapid.admin_api.v1.AdminV1Service/DeleteTenant"
	AdminV1Service_CreateStaff_FullMethodName             = "/rapid.admin_api.v1.AdminV1Service/CreateStaff"
)

// AdminV1ServiceClient is the client API for AdminV1Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminV1ServiceClient interface {
	CreateAssetPresignedURL(ctx context.Context, in *CreateAssetPresignedURLRequest, opts ...grpc.CallOption) (*CreateAssetPresignedURLResponse, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error)
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error)
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*UpdateTenantResponse, error)
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error)
	CreateStaff(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error)
}

type adminV1ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminV1ServiceClient(cc grpc.ClientConnInterface) AdminV1ServiceClient {
	return &adminV1ServiceClient{cc}
}

func (c *adminV1ServiceClient) CreateAssetPresignedURL(ctx context.Context, in *CreateAssetPresignedURLRequest, opts ...grpc.CallOption) (*CreateAssetPresignedURLResponse, error) {
	out := new(CreateAssetPresignedURLResponse)
	err := c.cc.Invoke(ctx, AdminV1Service_CreateAssetPresignedURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminV1ServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, AdminV1Service_GetTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminV1ServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, AdminV1Service_ListTenants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminV1ServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error) {
	out := new(CreateTenantResponse)
	err := c.cc.Invoke(ctx, AdminV1Service_CreateTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminV1ServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*UpdateTenantResponse, error) {
	out := new(UpdateTenantResponse)
	err := c.cc.Invoke(ctx, AdminV1Service_UpdateTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminV1ServiceClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error) {
	out := new(DeleteTenantResponse)
	err := c.cc.Invoke(ctx, AdminV1Service_DeleteTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminV1ServiceClient) CreateStaff(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error) {
	out := new(CreateStaffResponse)
	err := c.cc.Invoke(ctx, AdminV1Service_CreateStaff_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminV1ServiceServer is the server API for AdminV1Service service.
// All implementations should embed UnimplementedAdminV1ServiceServer
// for forward compatibility
type AdminV1ServiceServer interface {
	CreateAssetPresignedURL(context.Context, *CreateAssetPresignedURLRequest) (*CreateAssetPresignedURLResponse, error)
	GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error)
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error)
	UpdateTenant(context.Context, *UpdateTenantRequest) (*UpdateTenantResponse, error)
	DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error)
	CreateStaff(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error)
}

// UnimplementedAdminV1ServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdminV1ServiceServer struct {
}

func (UnimplementedAdminV1ServiceServer) CreateAssetPresignedURL(context.Context, *CreateAssetPresignedURLRequest) (*CreateAssetPresignedURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssetPresignedURL not implemented")
}
func (UnimplementedAdminV1ServiceServer) GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedAdminV1ServiceServer) ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}
func (UnimplementedAdminV1ServiceServer) CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedAdminV1ServiceServer) UpdateTenant(context.Context, *UpdateTenantRequest) (*UpdateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenant not implemented")
}
func (UnimplementedAdminV1ServiceServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedAdminV1ServiceServer) CreateStaff(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaff not implemented")
}

// UnsafeAdminV1ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminV1ServiceServer will
// result in compilation errors.
type UnsafeAdminV1ServiceServer interface {
	mustEmbedUnimplementedAdminV1ServiceServer()
}

func RegisterAdminV1ServiceServer(s grpc.ServiceRegistrar, srv AdminV1ServiceServer) {
	s.RegisterService(&AdminV1Service_ServiceDesc, srv)
}

func _AdminV1Service_CreateAssetPresignedURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssetPresignedURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminV1ServiceServer).CreateAssetPresignedURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminV1Service_CreateAssetPresignedURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminV1ServiceServer).CreateAssetPresignedURL(ctx, req.(*CreateAssetPresignedURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminV1Service_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminV1ServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminV1Service_GetTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminV1ServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminV1Service_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminV1ServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminV1Service_ListTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminV1ServiceServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminV1Service_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminV1ServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminV1Service_CreateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminV1ServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminV1Service_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminV1ServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminV1Service_UpdateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminV1ServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminV1Service_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminV1ServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminV1Service_DeleteTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminV1ServiceServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminV1Service_CreateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminV1ServiceServer).CreateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminV1Service_CreateStaff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminV1ServiceServer).CreateStaff(ctx, req.(*CreateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminV1Service_ServiceDesc is the grpc.ServiceDesc for AdminV1Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminV1Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rapid.admin_api.v1.AdminV1Service",
	HandlerType: (*AdminV1ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAssetPresignedURL",
			Handler:    _AdminV1Service_CreateAssetPresignedURL_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _AdminV1Service_GetTenant_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _AdminV1Service_ListTenants_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _AdminV1Service_CreateTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _AdminV1Service_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _AdminV1Service_DeleteTenant_Handler,
		},
		{
			MethodName: "CreateStaff",
			Handler:    _AdminV1Service_CreateStaff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rapid/admin_api/v1/api.proto",
}
