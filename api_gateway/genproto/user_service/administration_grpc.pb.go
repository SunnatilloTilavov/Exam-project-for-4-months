// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: administration.proto

package user_service

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
	AdministrationService_Create_FullMethodName        = "/user_service.AdministrationService/Create"
	AdministrationService_GetByID_FullMethodName       = "/user_service.AdministrationService/GetByID"
	AdministrationService_GetList_FullMethodName       = "/user_service.AdministrationService/GetList"
	AdministrationService_Update_FullMethodName        = "/user_service.AdministrationService/Update"
	AdministrationService_Delete_FullMethodName        = "/user_service.AdministrationService/Delete"
	AdministrationService_GetReportList_FullMethodName = "/user_service.AdministrationService/GetReportList"
)

// AdministrationServiceClient is the client API for AdministrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdministrationServiceClient interface {
	Create(ctx context.Context, in *CreateAdministrationRequest, opts ...grpc.CallOption) (*AdministrationResponse, error)
	GetByID(ctx context.Context, in *AdministrationID, opts ...grpc.CallOption) (*GetAdministrationResponse, error)
	GetList(ctx context.Context, in *GetListAdministrationRequest, opts ...grpc.CallOption) (*GetListAdministrationResponse, error)
	Update(ctx context.Context, in *UpdateAdministrationRequest, opts ...grpc.CallOption) (*GetAdministrationResponse, error)
	Delete(ctx context.Context, in *AdministrationID, opts ...grpc.CallOption) (*AdministrationEmpty, error)
	GetReportList(ctx context.Context, in *GetReportListAdministrationRequest, opts ...grpc.CallOption) (*GetReportListAdministrationResponse, error)
}

type administrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdministrationServiceClient(cc grpc.ClientConnInterface) AdministrationServiceClient {
	return &administrationServiceClient{cc}
}

func (c *administrationServiceClient) Create(ctx context.Context, in *CreateAdministrationRequest, opts ...grpc.CallOption) (*AdministrationResponse, error) {
	out := new(AdministrationResponse)
	err := c.cc.Invoke(ctx, AdministrationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) GetByID(ctx context.Context, in *AdministrationID, opts ...grpc.CallOption) (*GetAdministrationResponse, error) {
	out := new(GetAdministrationResponse)
	err := c.cc.Invoke(ctx, AdministrationService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) GetList(ctx context.Context, in *GetListAdministrationRequest, opts ...grpc.CallOption) (*GetListAdministrationResponse, error) {
	out := new(GetListAdministrationResponse)
	err := c.cc.Invoke(ctx, AdministrationService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) Update(ctx context.Context, in *UpdateAdministrationRequest, opts ...grpc.CallOption) (*GetAdministrationResponse, error) {
	out := new(GetAdministrationResponse)
	err := c.cc.Invoke(ctx, AdministrationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) Delete(ctx context.Context, in *AdministrationID, opts ...grpc.CallOption) (*AdministrationEmpty, error) {
	out := new(AdministrationEmpty)
	err := c.cc.Invoke(ctx, AdministrationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) GetReportList(ctx context.Context, in *GetReportListAdministrationRequest, opts ...grpc.CallOption) (*GetReportListAdministrationResponse, error) {
	out := new(GetReportListAdministrationResponse)
	err := c.cc.Invoke(ctx, AdministrationService_GetReportList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdministrationServiceServer is the server API for AdministrationService service.
// All implementations must embed UnimplementedAdministrationServiceServer
// for forward compatibility
type AdministrationServiceServer interface {
	Create(context.Context, *CreateAdministrationRequest) (*AdministrationResponse, error)
	GetByID(context.Context, *AdministrationID) (*GetAdministrationResponse, error)
	GetList(context.Context, *GetListAdministrationRequest) (*GetListAdministrationResponse, error)
	Update(context.Context, *UpdateAdministrationRequest) (*GetAdministrationResponse, error)
	Delete(context.Context, *AdministrationID) (*AdministrationEmpty, error)
	GetReportList(context.Context, *GetReportListAdministrationRequest) (*GetReportListAdministrationResponse, error)
	mustEmbedUnimplementedAdministrationServiceServer()
}

// UnimplementedAdministrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdministrationServiceServer struct {
}

func (UnimplementedAdministrationServiceServer) Create(context.Context, *CreateAdministrationRequest) (*AdministrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdministrationServiceServer) GetByID(context.Context, *AdministrationID) (*GetAdministrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedAdministrationServiceServer) GetList(context.Context, *GetListAdministrationRequest) (*GetListAdministrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdministrationServiceServer) Update(context.Context, *UpdateAdministrationRequest) (*GetAdministrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdministrationServiceServer) Delete(context.Context, *AdministrationID) (*AdministrationEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdministrationServiceServer) GetReportList(context.Context, *GetReportListAdministrationRequest) (*GetReportListAdministrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportList not implemented")
}
func (UnimplementedAdministrationServiceServer) mustEmbedUnimplementedAdministrationServiceServer() {}

// UnsafeAdministrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdministrationServiceServer will
// result in compilation errors.
type UnsafeAdministrationServiceServer interface {
	mustEmbedUnimplementedAdministrationServiceServer()
}

func RegisterAdministrationServiceServer(s grpc.ServiceRegistrar, srv AdministrationServiceServer) {
	s.RegisterService(&AdministrationService_ServiceDesc, srv)
}

func _AdministrationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdministrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).Create(ctx, req.(*CreateAdministrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministrationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).GetByID(ctx, req.(*AdministrationID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAdministrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).GetList(ctx, req.(*GetListAdministrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdministrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).Update(ctx, req.(*UpdateAdministrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministrationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).Delete(ctx, req.(*AdministrationID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_GetReportList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportListAdministrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).GetReportList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_GetReportList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).GetReportList(ctx, req.(*GetReportListAdministrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdministrationService_ServiceDesc is the grpc.ServiceDesc for AdministrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdministrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.AdministrationService",
	HandlerType: (*AdministrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AdministrationService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _AdministrationService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _AdministrationService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdministrationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdministrationService_Delete_Handler,
		},
		{
			MethodName: "GetReportList",
			Handler:    _AdministrationService_GetReportList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "administration.proto",
}
