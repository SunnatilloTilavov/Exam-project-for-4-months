// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: schedule.proto

package education_management_service

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
	ScheduleService_GetByID_FullMethodName      = "/education_management_service.ScheduleService/GetByID"
	ScheduleService_GetList_FullMethodName      = "/education_management_service.ScheduleService/GetList"
	ScheduleService_Create_FullMethodName       = "/education_management_service.ScheduleService/Create"
	ScheduleService_Update_FullMethodName       = "/education_management_service.ScheduleService/Update"
	ScheduleService_Delete_FullMethodName       = "/education_management_service.ScheduleService/Delete"
	ScheduleService_GetListMonth_FullMethodName = "/education_management_service.ScheduleService/GetListMonth"
)

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	GetByID(ctx context.Context, in *ScheduleID, opts ...grpc.CallOption) (*GetScheduleResponse, error)
	GetList(ctx context.Context, in *GetListScheduleRequest, opts ...grpc.CallOption) (*GetListScheduleResponse, error)
	Create(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	Update(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*GetScheduleResponse, error)
	Delete(ctx context.Context, in *ScheduleID, opts ...grpc.CallOption) (*ScheduleEmpty, error)
	GetListMonth(ctx context.Context, in *GetListScheduleMonthRequest, opts ...grpc.CallOption) (*GetListScheduleResponse, error)
}

type scheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleServiceClient(cc grpc.ClientConnInterface) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) GetByID(ctx context.Context, in *ScheduleID, opts ...grpc.CallOption) (*GetScheduleResponse, error) {
	out := new(GetScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetList(ctx context.Context, in *GetListScheduleRequest, opts ...grpc.CallOption) (*GetListScheduleResponse, error) {
	out := new(GetListScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) Create(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) Update(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*GetScheduleResponse, error) {
	out := new(GetScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) Delete(ctx context.Context, in *ScheduleID, opts ...grpc.CallOption) (*ScheduleEmpty, error) {
	out := new(ScheduleEmpty)
	err := c.cc.Invoke(ctx, ScheduleService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetListMonth(ctx context.Context, in *GetListScheduleMonthRequest, opts ...grpc.CallOption) (*GetListScheduleResponse, error) {
	out := new(GetListScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetListMonth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
// All implementations must embed UnimplementedScheduleServiceServer
// for forward compatibility
type ScheduleServiceServer interface {
	GetByID(context.Context, *ScheduleID) (*GetScheduleResponse, error)
	GetList(context.Context, *GetListScheduleRequest) (*GetListScheduleResponse, error)
	Create(context.Context, *CreateScheduleRequest) (*ScheduleResponse, error)
	Update(context.Context, *UpdateScheduleRequest) (*GetScheduleResponse, error)
	Delete(context.Context, *ScheduleID) (*ScheduleEmpty, error)
	GetListMonth(context.Context, *GetListScheduleMonthRequest) (*GetListScheduleResponse, error)
	mustEmbedUnimplementedScheduleServiceServer()
}

// UnimplementedScheduleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServiceServer struct {
}

func (UnimplementedScheduleServiceServer) GetByID(context.Context, *ScheduleID) (*GetScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedScheduleServiceServer) GetList(context.Context, *GetListScheduleRequest) (*GetListScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedScheduleServiceServer) Create(context.Context, *CreateScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedScheduleServiceServer) Update(context.Context, *UpdateScheduleRequest) (*GetScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedScheduleServiceServer) Delete(context.Context, *ScheduleID) (*ScheduleEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedScheduleServiceServer) GetListMonth(context.Context, *GetListScheduleMonthRequest) (*GetListScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListMonth not implemented")
}
func (UnimplementedScheduleServiceServer) mustEmbedUnimplementedScheduleServiceServer() {}

// UnsafeScheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServiceServer will
// result in compilation errors.
type UnsafeScheduleServiceServer interface {
	mustEmbedUnimplementedScheduleServiceServer()
}

func RegisterScheduleServiceServer(s grpc.ServiceRegistrar, srv ScheduleServiceServer) {
	s.RegisterService(&ScheduleService_ServiceDesc, srv)
}

func _ScheduleService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetByID(ctx, req.(*ScheduleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetList(ctx, req.(*GetListScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).Create(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).Update(ctx, req.(*UpdateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).Delete(ctx, req.(*ScheduleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetListMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListScheduleMonthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetListMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetListMonth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetListMonth(ctx, req.(*GetListScheduleMonthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduleService_ServiceDesc is the grpc.ServiceDesc for ScheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "education_management_service.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _ScheduleService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ScheduleService_GetList_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ScheduleService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ScheduleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ScheduleService_Delete_Handler,
		},
		{
			MethodName: "GetListMonth",
			Handler:    _ScheduleService_GetListMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule.proto",
}
