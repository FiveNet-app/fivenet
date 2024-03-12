// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: services/jobs/qualifications.proto

package jobs

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
	JobsQualificationsService_ListQualifications_FullMethodName        = "/services.jobs.JobsQualificationsService/ListQualifications"
	JobsQualificationsService_GetQualification_FullMethodName          = "/services.jobs.JobsQualificationsService/GetQualification"
	JobsQualificationsService_ListQualificationsResults_FullMethodName = "/services.jobs.JobsQualificationsService/ListQualificationsResults"
	JobsQualificationsService_CreateQualification_FullMethodName       = "/services.jobs.JobsQualificationsService/CreateQualification"
	JobsQualificationsService_UpdateQualification_FullMethodName       = "/services.jobs.JobsQualificationsService/UpdateQualification"
	JobsQualificationsService_DeleteQualification_FullMethodName       = "/services.jobs.JobsQualificationsService/DeleteQualification"
	JobsQualificationsService_ListQualificationRequests_FullMethodName = "/services.jobs.JobsQualificationsService/ListQualificationRequests"
)

// JobsQualificationsServiceClient is the client API for JobsQualificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobsQualificationsServiceClient interface {
	// @perm
	ListQualifications(ctx context.Context, in *ListQualificationsRequest, opts ...grpc.CallOption) (*ListQualificationsResponse, error)
	// @perm
	GetQualification(ctx context.Context, in *GetQualificationRequest, opts ...grpc.CallOption) (*GetQualificationResponse, error)
	// @perm
	ListQualificationsResults(ctx context.Context, in *ListQualificationsResultsRequest, opts ...grpc.CallOption) (*ListQualificationsResultsResponse, error)
	// @perm
	CreateQualification(ctx context.Context, in *CreateQualificationRequest, opts ...grpc.CallOption) (*CreateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	UpdateQualification(ctx context.Context, in *UpdateQualificationRequest, opts ...grpc.CallOption) (*UpdateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	DeleteQualification(ctx context.Context, in *DeleteQualificationRequest, opts ...grpc.CallOption) (*DeleteQualificationResponse, error)
	// @perm
	ListQualificationRequests(ctx context.Context, in *ListQualificationRequestsRequest, opts ...grpc.CallOption) (*ListQualificationRequestsResponse, error)
}

type jobsQualificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobsQualificationsServiceClient(cc grpc.ClientConnInterface) JobsQualificationsServiceClient {
	return &jobsQualificationsServiceClient{cc}
}

func (c *jobsQualificationsServiceClient) ListQualifications(ctx context.Context, in *ListQualificationsRequest, opts ...grpc.CallOption) (*ListQualificationsResponse, error) {
	out := new(ListQualificationsResponse)
	err := c.cc.Invoke(ctx, JobsQualificationsService_ListQualifications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsQualificationsServiceClient) GetQualification(ctx context.Context, in *GetQualificationRequest, opts ...grpc.CallOption) (*GetQualificationResponse, error) {
	out := new(GetQualificationResponse)
	err := c.cc.Invoke(ctx, JobsQualificationsService_GetQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsQualificationsServiceClient) ListQualificationsResults(ctx context.Context, in *ListQualificationsResultsRequest, opts ...grpc.CallOption) (*ListQualificationsResultsResponse, error) {
	out := new(ListQualificationsResultsResponse)
	err := c.cc.Invoke(ctx, JobsQualificationsService_ListQualificationsResults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsQualificationsServiceClient) CreateQualification(ctx context.Context, in *CreateQualificationRequest, opts ...grpc.CallOption) (*CreateQualificationResponse, error) {
	out := new(CreateQualificationResponse)
	err := c.cc.Invoke(ctx, JobsQualificationsService_CreateQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsQualificationsServiceClient) UpdateQualification(ctx context.Context, in *UpdateQualificationRequest, opts ...grpc.CallOption) (*UpdateQualificationResponse, error) {
	out := new(UpdateQualificationResponse)
	err := c.cc.Invoke(ctx, JobsQualificationsService_UpdateQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsQualificationsServiceClient) DeleteQualification(ctx context.Context, in *DeleteQualificationRequest, opts ...grpc.CallOption) (*DeleteQualificationResponse, error) {
	out := new(DeleteQualificationResponse)
	err := c.cc.Invoke(ctx, JobsQualificationsService_DeleteQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsQualificationsServiceClient) ListQualificationRequests(ctx context.Context, in *ListQualificationRequestsRequest, opts ...grpc.CallOption) (*ListQualificationRequestsResponse, error) {
	out := new(ListQualificationRequestsResponse)
	err := c.cc.Invoke(ctx, JobsQualificationsService_ListQualificationRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobsQualificationsServiceServer is the server API for JobsQualificationsService service.
// All implementations must embed UnimplementedJobsQualificationsServiceServer
// for forward compatibility
type JobsQualificationsServiceServer interface {
	// @perm
	ListQualifications(context.Context, *ListQualificationsRequest) (*ListQualificationsResponse, error)
	// @perm
	GetQualification(context.Context, *GetQualificationRequest) (*GetQualificationResponse, error)
	// @perm
	ListQualificationsResults(context.Context, *ListQualificationsResultsRequest) (*ListQualificationsResultsResponse, error)
	// @perm
	CreateQualification(context.Context, *CreateQualificationRequest) (*CreateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	UpdateQualification(context.Context, *UpdateQualificationRequest) (*UpdateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	DeleteQualification(context.Context, *DeleteQualificationRequest) (*DeleteQualificationResponse, error)
	// @perm
	ListQualificationRequests(context.Context, *ListQualificationRequestsRequest) (*ListQualificationRequestsResponse, error)
	mustEmbedUnimplementedJobsQualificationsServiceServer()
}

// UnimplementedJobsQualificationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobsQualificationsServiceServer struct {
}

func (UnimplementedJobsQualificationsServiceServer) ListQualifications(context.Context, *ListQualificationsRequest) (*ListQualificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQualifications not implemented")
}
func (UnimplementedJobsQualificationsServiceServer) GetQualification(context.Context, *GetQualificationRequest) (*GetQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQualification not implemented")
}
func (UnimplementedJobsQualificationsServiceServer) ListQualificationsResults(context.Context, *ListQualificationsResultsRequest) (*ListQualificationsResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQualificationsResults not implemented")
}
func (UnimplementedJobsQualificationsServiceServer) CreateQualification(context.Context, *CreateQualificationRequest) (*CreateQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQualification not implemented")
}
func (UnimplementedJobsQualificationsServiceServer) UpdateQualification(context.Context, *UpdateQualificationRequest) (*UpdateQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQualification not implemented")
}
func (UnimplementedJobsQualificationsServiceServer) DeleteQualification(context.Context, *DeleteQualificationRequest) (*DeleteQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQualification not implemented")
}
func (UnimplementedJobsQualificationsServiceServer) ListQualificationRequests(context.Context, *ListQualificationRequestsRequest) (*ListQualificationRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQualificationRequests not implemented")
}
func (UnimplementedJobsQualificationsServiceServer) mustEmbedUnimplementedJobsQualificationsServiceServer() {
}

// UnsafeJobsQualificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobsQualificationsServiceServer will
// result in compilation errors.
type UnsafeJobsQualificationsServiceServer interface {
	mustEmbedUnimplementedJobsQualificationsServiceServer()
}

func RegisterJobsQualificationsServiceServer(s grpc.ServiceRegistrar, srv JobsQualificationsServiceServer) {
	s.RegisterService(&JobsQualificationsService_ServiceDesc, srv)
}

func _JobsQualificationsService_ListQualifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQualificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsQualificationsServiceServer).ListQualifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsQualificationsService_ListQualifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsQualificationsServiceServer).ListQualifications(ctx, req.(*ListQualificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsQualificationsService_GetQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsQualificationsServiceServer).GetQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsQualificationsService_GetQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsQualificationsServiceServer).GetQualification(ctx, req.(*GetQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsQualificationsService_ListQualificationsResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQualificationsResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsQualificationsServiceServer).ListQualificationsResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsQualificationsService_ListQualificationsResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsQualificationsServiceServer).ListQualificationsResults(ctx, req.(*ListQualificationsResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsQualificationsService_CreateQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsQualificationsServiceServer).CreateQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsQualificationsService_CreateQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsQualificationsServiceServer).CreateQualification(ctx, req.(*CreateQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsQualificationsService_UpdateQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsQualificationsServiceServer).UpdateQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsQualificationsService_UpdateQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsQualificationsServiceServer).UpdateQualification(ctx, req.(*UpdateQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsQualificationsService_DeleteQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsQualificationsServiceServer).DeleteQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsQualificationsService_DeleteQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsQualificationsServiceServer).DeleteQualification(ctx, req.(*DeleteQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsQualificationsService_ListQualificationRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQualificationRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsQualificationsServiceServer).ListQualificationRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsQualificationsService_ListQualificationRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsQualificationsServiceServer).ListQualificationRequests(ctx, req.(*ListQualificationRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobsQualificationsService_ServiceDesc is the grpc.ServiceDesc for JobsQualificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobsQualificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.jobs.JobsQualificationsService",
	HandlerType: (*JobsQualificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListQualifications",
			Handler:    _JobsQualificationsService_ListQualifications_Handler,
		},
		{
			MethodName: "GetQualification",
			Handler:    _JobsQualificationsService_GetQualification_Handler,
		},
		{
			MethodName: "ListQualificationsResults",
			Handler:    _JobsQualificationsService_ListQualificationsResults_Handler,
		},
		{
			MethodName: "CreateQualification",
			Handler:    _JobsQualificationsService_CreateQualification_Handler,
		},
		{
			MethodName: "UpdateQualification",
			Handler:    _JobsQualificationsService_UpdateQualification_Handler,
		},
		{
			MethodName: "DeleteQualification",
			Handler:    _JobsQualificationsService_DeleteQualification_Handler,
		},
		{
			MethodName: "ListQualificationRequests",
			Handler:    _JobsQualificationsService_ListQualificationRequests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/jobs/qualifications.proto",
}
