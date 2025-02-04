// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: event_student.proto

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
	EventStudentService_Create_FullMethodName         = "/education_management_service.EventStudentService/Create"
	EventStudentService_GetByID_FullMethodName        = "/education_management_service.EventStudentService/GetByID"
	EventStudentService_GetList_FullMethodName        = "/education_management_service.EventStudentService/GetList"
	EventStudentService_Update_FullMethodName         = "/education_management_service.EventStudentService/Update"
	EventStudentService_Delete_FullMethodName         = "/education_management_service.EventStudentService/Delete"
	EventStudentService_GetStudentByID_FullMethodName = "/education_management_service.EventStudentService/GetStudentByID"
)

// EventStudentServiceClient is the client API for EventStudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventStudentServiceClient interface {
	Create(ctx context.Context, in *CreateEventStudentRequest, opts ...grpc.CallOption) (*EventStudentResponse, error)
	GetByID(ctx context.Context, in *EventStudentID, opts ...grpc.CallOption) (*GetEventStudentResponse, error)
	GetList(ctx context.Context, in *GetListEventStudentRequest, opts ...grpc.CallOption) (*GetListEventStudentResponse, error)
	Update(ctx context.Context, in *UpdateEventStudentRequest, opts ...grpc.CallOption) (*GetEventStudentResponse, error)
	Delete(ctx context.Context, in *EventStudentID, opts ...grpc.CallOption) (*EventStudentEmpty, error)
	GetStudentByID(ctx context.Context, in *StudentID, opts ...grpc.CallOption) (*GetStudentWithEventsResponse, error)
}

type eventStudentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventStudentServiceClient(cc grpc.ClientConnInterface) EventStudentServiceClient {
	return &eventStudentServiceClient{cc}
}

func (c *eventStudentServiceClient) Create(ctx context.Context, in *CreateEventStudentRequest, opts ...grpc.CallOption) (*EventStudentResponse, error) {
	out := new(EventStudentResponse)
	err := c.cc.Invoke(ctx, EventStudentService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStudentServiceClient) GetByID(ctx context.Context, in *EventStudentID, opts ...grpc.CallOption) (*GetEventStudentResponse, error) {
	out := new(GetEventStudentResponse)
	err := c.cc.Invoke(ctx, EventStudentService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStudentServiceClient) GetList(ctx context.Context, in *GetListEventStudentRequest, opts ...grpc.CallOption) (*GetListEventStudentResponse, error) {
	out := new(GetListEventStudentResponse)
	err := c.cc.Invoke(ctx, EventStudentService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStudentServiceClient) Update(ctx context.Context, in *UpdateEventStudentRequest, opts ...grpc.CallOption) (*GetEventStudentResponse, error) {
	out := new(GetEventStudentResponse)
	err := c.cc.Invoke(ctx, EventStudentService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStudentServiceClient) Delete(ctx context.Context, in *EventStudentID, opts ...grpc.CallOption) (*EventStudentEmpty, error) {
	out := new(EventStudentEmpty)
	err := c.cc.Invoke(ctx, EventStudentService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStudentServiceClient) GetStudentByID(ctx context.Context, in *StudentID, opts ...grpc.CallOption) (*GetStudentWithEventsResponse, error) {
	out := new(GetStudentWithEventsResponse)
	err := c.cc.Invoke(ctx, EventStudentService_GetStudentByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventStudentServiceServer is the server API for EventStudentService service.
// All implementations must embed UnimplementedEventStudentServiceServer
// for forward compatibility
type EventStudentServiceServer interface {
	Create(context.Context, *CreateEventStudentRequest) (*EventStudentResponse, error)
	GetByID(context.Context, *EventStudentID) (*GetEventStudentResponse, error)
	GetList(context.Context, *GetListEventStudentRequest) (*GetListEventStudentResponse, error)
	Update(context.Context, *UpdateEventStudentRequest) (*GetEventStudentResponse, error)
	Delete(context.Context, *EventStudentID) (*EventStudentEmpty, error)
	GetStudentByID(context.Context, *StudentID) (*GetStudentWithEventsResponse, error)
	mustEmbedUnimplementedEventStudentServiceServer()
}

// UnimplementedEventStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventStudentServiceServer struct {
}

func (UnimplementedEventStudentServiceServer) Create(context.Context, *CreateEventStudentRequest) (*EventStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEventStudentServiceServer) GetByID(context.Context, *EventStudentID) (*GetEventStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedEventStudentServiceServer) GetList(context.Context, *GetListEventStudentRequest) (*GetListEventStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedEventStudentServiceServer) Update(context.Context, *UpdateEventStudentRequest) (*GetEventStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEventStudentServiceServer) Delete(context.Context, *EventStudentID) (*EventStudentEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEventStudentServiceServer) GetStudentByID(context.Context, *StudentID) (*GetStudentWithEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentByID not implemented")
}
func (UnimplementedEventStudentServiceServer) mustEmbedUnimplementedEventStudentServiceServer() {}

// UnsafeEventStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventStudentServiceServer will
// result in compilation errors.
type UnsafeEventStudentServiceServer interface {
	mustEmbedUnimplementedEventStudentServiceServer()
}

func RegisterEventStudentServiceServer(s grpc.ServiceRegistrar, srv EventStudentServiceServer) {
	s.RegisterService(&EventStudentService_ServiceDesc, srv)
}

func _EventStudentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStudentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventStudentService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStudentServiceServer).Create(ctx, req.(*CreateEventStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStudentService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventStudentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStudentServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventStudentService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStudentServiceServer).GetByID(ctx, req.(*EventStudentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStudentService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListEventStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStudentServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventStudentService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStudentServiceServer).GetList(ctx, req.(*GetListEventStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStudentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStudentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventStudentService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStudentServiceServer).Update(ctx, req.(*UpdateEventStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStudentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventStudentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStudentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventStudentService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStudentServiceServer).Delete(ctx, req.(*EventStudentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStudentService_GetStudentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStudentServiceServer).GetStudentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventStudentService_GetStudentByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStudentServiceServer).GetStudentByID(ctx, req.(*StudentID))
	}
	return interceptor(ctx, in, info, handler)
}

// EventStudentService_ServiceDesc is the grpc.ServiceDesc for EventStudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventStudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "education_management_service.EventStudentService",
	HandlerType: (*EventStudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EventStudentService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _EventStudentService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _EventStudentService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EventStudentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EventStudentService_Delete_Handler,
		},
		{
			MethodName: "GetStudentByID",
			Handler:    _EventStudentService_GetStudentByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event_student.proto",
}
