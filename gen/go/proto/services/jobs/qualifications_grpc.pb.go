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
	JobsQualificationsService_ListQualifications_FullMethodName  = "/services.jobs.JobsQualificationsService/ListQualifications"
	JobsQualificationsService_CreateQualification_FullMethodName = "/services.jobs.JobsQualificationsService/CreateQualification"
	JobsQualificationsService_UpdateQualification_FullMethodName = "/services.jobs.JobsQualificationsService/UpdateQualification"
	JobsQualificationsService_DeleteQualification_FullMethodName = "/services.jobs.JobsQualificationsService/DeleteQualification"
)

// JobsQualificationsServiceClient is the client API for JobsQualificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobsQualificationsServiceClient interface {
	// @perm
	ListQualifications(ctx context.Context, in *ListQualificationsRequest, opts ...grpc.CallOption) (*ListQualificationsResponse, error)
	// @perm
	CreateQualification(ctx context.Context, in *CreateQualificationRequest, opts ...grpc.CallOption) (*CreateQualificationResponse, error)
	// @perm
	UpdateQualification(ctx context.Context, in *UpdateQualificationRequest, opts ...grpc.CallOption) (*UpdateQualificationResponse, error)
	// @perm
	DeleteQualification(ctx context.Context, in *DeleteQualificationRequest, opts ...grpc.CallOption) (*DeleteQualificationResponse, error)
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

// JobsQualificationsServiceServer is the server API for JobsQualificationsService service.
// All implementations must embed UnimplementedJobsQualificationsServiceServer
// for forward compatibility
type JobsQualificationsServiceServer interface {
	// @perm
	ListQualifications(context.Context, *ListQualificationsRequest) (*ListQualificationsResponse, error)
	// @perm
	CreateQualification(context.Context, *CreateQualificationRequest) (*CreateQualificationResponse, error)
	// @perm
	UpdateQualification(context.Context, *UpdateQualificationRequest) (*UpdateQualificationResponse, error)
	// @perm
	DeleteQualification(context.Context, *DeleteQualificationRequest) (*DeleteQualificationResponse, error)
	mustEmbedUnimplementedJobsQualificationsServiceServer()
}

// UnimplementedJobsQualificationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobsQualificationsServiceServer struct {
}

func (UnimplementedJobsQualificationsServiceServer) ListQualifications(context.Context, *ListQualificationsRequest) (*ListQualificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQualifications not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/jobs/qualifications.proto",
}
