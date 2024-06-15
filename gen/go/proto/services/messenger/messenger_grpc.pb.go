// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: services/messenger/messenger.proto

package messenger

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
	MessengerService_ListThreads_FullMethodName          = "/services.messenger.MessengerService/ListThreads"
	MessengerService_GetThread_FullMethodName            = "/services.messenger.MessengerService/GetThread"
	MessengerService_CreateOrUpdateThread_FullMethodName = "/services.messenger.MessengerService/CreateOrUpdateThread"
	MessengerService_DeleteThread_FullMethodName         = "/services.messenger.MessengerService/DeleteThread"
	MessengerService_SetThreadUserState_FullMethodName   = "/services.messenger.MessengerService/SetThreadUserState"
	MessengerService_LeaveThread_FullMethodName          = "/services.messenger.MessengerService/LeaveThread"
	MessengerService_GetThreadMessages_FullMethodName    = "/services.messenger.MessengerService/GetThreadMessages"
	MessengerService_PostMessage_FullMethodName          = "/services.messenger.MessengerService/PostMessage"
	MessengerService_DeleteMessage_FullMethodName        = "/services.messenger.MessengerService/DeleteMessage"
)

// MessengerServiceClient is the client API for MessengerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerServiceClient interface {
	// @perm
	ListThreads(ctx context.Context, in *ListThreadsRequest, opts ...grpc.CallOption) (*ListThreadsResponse, error)
	// @perm: Name=ListThreads
	GetThread(ctx context.Context, in *GetThreadRequest, opts ...grpc.CallOption) (*GetThreadResponse, error)
	// @perm
	CreateOrUpdateThread(ctx context.Context, in *CreateOrUpdateThreadRequest, opts ...grpc.CallOption) (*CreateOrUpdateThreadResponse, error)
	// @perm
	DeleteThread(ctx context.Context, in *DeleteThreadRequest, opts ...grpc.CallOption) (*DeleteThreadResponse, error)
	// @perm: Name=ListThreads
	SetThreadUserState(ctx context.Context, in *SetThreadUserStateRequest, opts ...grpc.CallOption) (*SetThreadUserStateResponse, error)
	// @perm: Name=ListThreads
	LeaveThread(ctx context.Context, in *LeaveThreadRequest, opts ...grpc.CallOption) (*LeaveThreadResponse, error)
	// @perm: Name=ListThreads
	GetThreadMessages(ctx context.Context, in *GetThreadMessagesRequest, opts ...grpc.CallOption) (*GetThreadMessagesResponse, error)
	// @perm
	PostMessage(ctx context.Context, in *PostMessageRequest, opts ...grpc.CallOption) (*PostMessageResponse, error)
	// @perm: Name=SuperUser
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
}

type messengerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerServiceClient(cc grpc.ClientConnInterface) MessengerServiceClient {
	return &messengerServiceClient{cc}
}

func (c *messengerServiceClient) ListThreads(ctx context.Context, in *ListThreadsRequest, opts ...grpc.CallOption) (*ListThreadsResponse, error) {
	out := new(ListThreadsResponse)
	err := c.cc.Invoke(ctx, MessengerService_ListThreads_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) GetThread(ctx context.Context, in *GetThreadRequest, opts ...grpc.CallOption) (*GetThreadResponse, error) {
	out := new(GetThreadResponse)
	err := c.cc.Invoke(ctx, MessengerService_GetThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) CreateOrUpdateThread(ctx context.Context, in *CreateOrUpdateThreadRequest, opts ...grpc.CallOption) (*CreateOrUpdateThreadResponse, error) {
	out := new(CreateOrUpdateThreadResponse)
	err := c.cc.Invoke(ctx, MessengerService_CreateOrUpdateThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) DeleteThread(ctx context.Context, in *DeleteThreadRequest, opts ...grpc.CallOption) (*DeleteThreadResponse, error) {
	out := new(DeleteThreadResponse)
	err := c.cc.Invoke(ctx, MessengerService_DeleteThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) SetThreadUserState(ctx context.Context, in *SetThreadUserStateRequest, opts ...grpc.CallOption) (*SetThreadUserStateResponse, error) {
	out := new(SetThreadUserStateResponse)
	err := c.cc.Invoke(ctx, MessengerService_SetThreadUserState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) LeaveThread(ctx context.Context, in *LeaveThreadRequest, opts ...grpc.CallOption) (*LeaveThreadResponse, error) {
	out := new(LeaveThreadResponse)
	err := c.cc.Invoke(ctx, MessengerService_LeaveThread_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) GetThreadMessages(ctx context.Context, in *GetThreadMessagesRequest, opts ...grpc.CallOption) (*GetThreadMessagesResponse, error) {
	out := new(GetThreadMessagesResponse)
	err := c.cc.Invoke(ctx, MessengerService_GetThreadMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) PostMessage(ctx context.Context, in *PostMessageRequest, opts ...grpc.CallOption) (*PostMessageResponse, error) {
	out := new(PostMessageResponse)
	err := c.cc.Invoke(ctx, MessengerService_PostMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, MessengerService_DeleteMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServiceServer is the server API for MessengerService service.
// All implementations must embed UnimplementedMessengerServiceServer
// for forward compatibility
type MessengerServiceServer interface {
	// @perm
	ListThreads(context.Context, *ListThreadsRequest) (*ListThreadsResponse, error)
	// @perm: Name=ListThreads
	GetThread(context.Context, *GetThreadRequest) (*GetThreadResponse, error)
	// @perm
	CreateOrUpdateThread(context.Context, *CreateOrUpdateThreadRequest) (*CreateOrUpdateThreadResponse, error)
	// @perm
	DeleteThread(context.Context, *DeleteThreadRequest) (*DeleteThreadResponse, error)
	// @perm: Name=ListThreads
	SetThreadUserState(context.Context, *SetThreadUserStateRequest) (*SetThreadUserStateResponse, error)
	// @perm: Name=ListThreads
	LeaveThread(context.Context, *LeaveThreadRequest) (*LeaveThreadResponse, error)
	// @perm: Name=ListThreads
	GetThreadMessages(context.Context, *GetThreadMessagesRequest) (*GetThreadMessagesResponse, error)
	// @perm
	PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error)
	// @perm: Name=SuperUser
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	mustEmbedUnimplementedMessengerServiceServer()
}

// UnimplementedMessengerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServiceServer struct {
}

func (UnimplementedMessengerServiceServer) ListThreads(context.Context, *ListThreadsRequest) (*ListThreadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListThreads not implemented")
}
func (UnimplementedMessengerServiceServer) GetThread(context.Context, *GetThreadRequest) (*GetThreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThread not implemented")
}
func (UnimplementedMessengerServiceServer) CreateOrUpdateThread(context.Context, *CreateOrUpdateThreadRequest) (*CreateOrUpdateThreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateThread not implemented")
}
func (UnimplementedMessengerServiceServer) DeleteThread(context.Context, *DeleteThreadRequest) (*DeleteThreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThread not implemented")
}
func (UnimplementedMessengerServiceServer) SetThreadUserState(context.Context, *SetThreadUserStateRequest) (*SetThreadUserStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetThreadUserState not implemented")
}
func (UnimplementedMessengerServiceServer) LeaveThread(context.Context, *LeaveThreadRequest) (*LeaveThreadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveThread not implemented")
}
func (UnimplementedMessengerServiceServer) GetThreadMessages(context.Context, *GetThreadMessagesRequest) (*GetThreadMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThreadMessages not implemented")
}
func (UnimplementedMessengerServiceServer) PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMessage not implemented")
}
func (UnimplementedMessengerServiceServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedMessengerServiceServer) mustEmbedUnimplementedMessengerServiceServer() {}

// UnsafeMessengerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServiceServer will
// result in compilation errors.
type UnsafeMessengerServiceServer interface {
	mustEmbedUnimplementedMessengerServiceServer()
}

func RegisterMessengerServiceServer(s grpc.ServiceRegistrar, srv MessengerServiceServer) {
	s.RegisterService(&MessengerService_ServiceDesc, srv)
}

func _MessengerService_ListThreads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListThreadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).ListThreads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_ListThreads_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).ListThreads(ctx, req.(*ListThreadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_GetThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_GetThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetThread(ctx, req.(*GetThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_CreateOrUpdateThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).CreateOrUpdateThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_CreateOrUpdateThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).CreateOrUpdateThread(ctx, req.(*CreateOrUpdateThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_DeleteThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).DeleteThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_DeleteThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).DeleteThread(ctx, req.(*DeleteThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_SetThreadUserState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetThreadUserStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).SetThreadUserState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_SetThreadUserState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).SetThreadUserState(ctx, req.(*SetThreadUserStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_LeaveThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).LeaveThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_LeaveThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).LeaveThread(ctx, req.(*LeaveThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_GetThreadMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThreadMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).GetThreadMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_GetThreadMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).GetThreadMessages(ctx, req.(*GetThreadMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_PostMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).PostMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_PostMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).PostMessage(ctx, req.(*PostMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessengerService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessengerService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServiceServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessengerService_ServiceDesc is the grpc.ServiceDesc for MessengerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessengerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.messenger.MessengerService",
	HandlerType: (*MessengerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListThreads",
			Handler:    _MessengerService_ListThreads_Handler,
		},
		{
			MethodName: "GetThread",
			Handler:    _MessengerService_GetThread_Handler,
		},
		{
			MethodName: "CreateOrUpdateThread",
			Handler:    _MessengerService_CreateOrUpdateThread_Handler,
		},
		{
			MethodName: "DeleteThread",
			Handler:    _MessengerService_DeleteThread_Handler,
		},
		{
			MethodName: "SetThreadUserState",
			Handler:    _MessengerService_SetThreadUserState_Handler,
		},
		{
			MethodName: "LeaveThread",
			Handler:    _MessengerService_LeaveThread_Handler,
		},
		{
			MethodName: "GetThreadMessages",
			Handler:    _MessengerService_GetThreadMessages_Handler,
		},
		{
			MethodName: "PostMessage",
			Handler:    _MessengerService_PostMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _MessengerService_DeleteMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/messenger/messenger.proto",
}
