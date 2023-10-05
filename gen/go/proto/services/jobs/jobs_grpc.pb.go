// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: services/jobs/jobs.proto

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
	JobsService_ColleaguesList_FullMethodName             = "/services.jobs.JobsService/ColleaguesList"
	JobsService_ConductListEntries_FullMethodName         = "/services.jobs.JobsService/ConductListEntries"
	JobsService_TimeclockStats_FullMethodName             = "/services.jobs.JobsService/TimeclockStats"
	JobsService_ConductCreateEntry_FullMethodName         = "/services.jobs.JobsService/ConductCreateEntry"
	JobsService_ConductUpdateEntry_FullMethodName         = "/services.jobs.JobsService/ConductUpdateEntry"
	JobsService_ConductDeleteEntry_FullMethodName         = "/services.jobs.JobsService/ConductDeleteEntry"
	JobsService_TimeclockListEntries_FullMethodName       = "/services.jobs.JobsService/TimeclockListEntries"
	JobsService_RequestsListEntries_FullMethodName        = "/services.jobs.JobsService/RequestsListEntries"
	JobsService_RequestsCreateEntry_FullMethodName        = "/services.jobs.JobsService/RequestsCreateEntry"
	JobsService_RequestsUpdateEntry_FullMethodName        = "/services.jobs.JobsService/RequestsUpdateEntry"
	JobsService_RequestsDeleteEntry_FullMethodName        = "/services.jobs.JobsService/RequestsDeleteEntry"
	JobsService_RequestsListTypes_FullMethodName          = "/services.jobs.JobsService/RequestsListTypes"
	JobsService_RequestsCreateOrUpdateType_FullMethodName = "/services.jobs.JobsService/RequestsCreateOrUpdateType"
	JobsService_RequestsDeleteType_FullMethodName         = "/services.jobs.JobsService/RequestsDeleteType"
	JobsService_RequestsPostComment_FullMethodName        = "/services.jobs.JobsService/RequestsPostComment"
	JobsService_RequestsDeleteComment_FullMethodName      = "/services.jobs.JobsService/RequestsDeleteComment"
)

// JobsServiceClient is the client API for JobsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobsServiceClient interface {
	// @perm
	ColleaguesList(ctx context.Context, in *ColleaguesListRequest, opts ...grpc.CallOption) (*ColleaguesListResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "All"}§[]string{"Own"}
	ConductListEntries(ctx context.Context, in *ConductListEntriesRequest, opts ...grpc.CallOption) (*ConductListEntriesResponse, error)
	// @perm: Name=ConductListEntries
	TimeclockStats(ctx context.Context, in *TimeclockStatsRequest, opts ...grpc.CallOption) (*TimeclockStatsResponse, error)
	// @perm
	ConductCreateEntry(ctx context.Context, in *ConductCreateEntryRequest, opts ...grpc.CallOption) (*ConductCreateEntryResponse, error)
	// @perm
	ConductUpdateEntry(ctx context.Context, in *ConductUpdateEntryRequest, opts ...grpc.CallOption) (*ConductUpdateEntryResponse, error)
	// @perm
	ConductDeleteEntry(ctx context.Context, in *ConductDeleteEntryRequest, opts ...grpc.CallOption) (*ConductDeleteEntryResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "All"}§[]string{"Own"}
	TimeclockListEntries(ctx context.Context, in *TimeclockListEntriesRequest, opts ...grpc.CallOption) (*TimeclockListEntriesResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "All"}§[]string{"Own"}
	RequestsListEntries(ctx context.Context, in *RequestsListEntriesRequest, opts ...grpc.CallOption) (*RequestsListEntriesResponse, error)
	// @perm
	RequestsCreateEntry(ctx context.Context, in *RequestsCreateEntryRequest, opts ...grpc.CallOption) (*RequestsCreateEntryResponse, error)
	// @perm
	RequestsUpdateEntry(ctx context.Context, in *RequestsUpdateEntryRequest, opts ...grpc.CallOption) (*RequestsUpdateEntryResponse, error)
	// @perm
	RequestsDeleteEntry(ctx context.Context, in *RequestsDeleteEntryRequest, opts ...grpc.CallOption) (*RequestsDeleteEntryResponse, error)
	// @perm: Name=RequestsListEntries
	RequestsListTypes(ctx context.Context, in *RequestsListTypesRequest, opts ...grpc.CallOption) (*RequestsListTypesResponse, error)
	// @perm
	RequestsCreateOrUpdateType(ctx context.Context, in *RequestsCreateOrUpdateTypeRequest, opts ...grpc.CallOption) (*RequestsCreateOrUpdateTypeResponse, error)
	// @perm
	RequestsDeleteType(ctx context.Context, in *RequestsDeleteTypeRequest, opts ...grpc.CallOption) (*RequestsDeleteTypeResponse, error)
	// @perm
	RequestsPostComment(ctx context.Context, in *RequestsPostCommentRequest, opts ...grpc.CallOption) (*RequestsPostCommentResponse, error)
	// @perm
	RequestsDeleteComment(ctx context.Context, in *RequestsDeleteCommentRequest, opts ...grpc.CallOption) (*RequestsDeleteCommentResponse, error)
}

type jobsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobsServiceClient(cc grpc.ClientConnInterface) JobsServiceClient {
	return &jobsServiceClient{cc}
}

func (c *jobsServiceClient) ColleaguesList(ctx context.Context, in *ColleaguesListRequest, opts ...grpc.CallOption) (*ColleaguesListResponse, error) {
	out := new(ColleaguesListResponse)
	err := c.cc.Invoke(ctx, JobsService_ColleaguesList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) ConductListEntries(ctx context.Context, in *ConductListEntriesRequest, opts ...grpc.CallOption) (*ConductListEntriesResponse, error) {
	out := new(ConductListEntriesResponse)
	err := c.cc.Invoke(ctx, JobsService_ConductListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) TimeclockStats(ctx context.Context, in *TimeclockStatsRequest, opts ...grpc.CallOption) (*TimeclockStatsResponse, error) {
	out := new(TimeclockStatsResponse)
	err := c.cc.Invoke(ctx, JobsService_TimeclockStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) ConductCreateEntry(ctx context.Context, in *ConductCreateEntryRequest, opts ...grpc.CallOption) (*ConductCreateEntryResponse, error) {
	out := new(ConductCreateEntryResponse)
	err := c.cc.Invoke(ctx, JobsService_ConductCreateEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) ConductUpdateEntry(ctx context.Context, in *ConductUpdateEntryRequest, opts ...grpc.CallOption) (*ConductUpdateEntryResponse, error) {
	out := new(ConductUpdateEntryResponse)
	err := c.cc.Invoke(ctx, JobsService_ConductUpdateEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) ConductDeleteEntry(ctx context.Context, in *ConductDeleteEntryRequest, opts ...grpc.CallOption) (*ConductDeleteEntryResponse, error) {
	out := new(ConductDeleteEntryResponse)
	err := c.cc.Invoke(ctx, JobsService_ConductDeleteEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) TimeclockListEntries(ctx context.Context, in *TimeclockListEntriesRequest, opts ...grpc.CallOption) (*TimeclockListEntriesResponse, error) {
	out := new(TimeclockListEntriesResponse)
	err := c.cc.Invoke(ctx, JobsService_TimeclockListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsListEntries(ctx context.Context, in *RequestsListEntriesRequest, opts ...grpc.CallOption) (*RequestsListEntriesResponse, error) {
	out := new(RequestsListEntriesResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsCreateEntry(ctx context.Context, in *RequestsCreateEntryRequest, opts ...grpc.CallOption) (*RequestsCreateEntryResponse, error) {
	out := new(RequestsCreateEntryResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsCreateEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsUpdateEntry(ctx context.Context, in *RequestsUpdateEntryRequest, opts ...grpc.CallOption) (*RequestsUpdateEntryResponse, error) {
	out := new(RequestsUpdateEntryResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsUpdateEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsDeleteEntry(ctx context.Context, in *RequestsDeleteEntryRequest, opts ...grpc.CallOption) (*RequestsDeleteEntryResponse, error) {
	out := new(RequestsDeleteEntryResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsDeleteEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsListTypes(ctx context.Context, in *RequestsListTypesRequest, opts ...grpc.CallOption) (*RequestsListTypesResponse, error) {
	out := new(RequestsListTypesResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsListTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsCreateOrUpdateType(ctx context.Context, in *RequestsCreateOrUpdateTypeRequest, opts ...grpc.CallOption) (*RequestsCreateOrUpdateTypeResponse, error) {
	out := new(RequestsCreateOrUpdateTypeResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsCreateOrUpdateType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsDeleteType(ctx context.Context, in *RequestsDeleteTypeRequest, opts ...grpc.CallOption) (*RequestsDeleteTypeResponse, error) {
	out := new(RequestsDeleteTypeResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsDeleteType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsPostComment(ctx context.Context, in *RequestsPostCommentRequest, opts ...grpc.CallOption) (*RequestsPostCommentResponse, error) {
	out := new(RequestsPostCommentResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsPostComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) RequestsDeleteComment(ctx context.Context, in *RequestsDeleteCommentRequest, opts ...grpc.CallOption) (*RequestsDeleteCommentResponse, error) {
	out := new(RequestsDeleteCommentResponse)
	err := c.cc.Invoke(ctx, JobsService_RequestsDeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobsServiceServer is the server API for JobsService service.
// All implementations must embed UnimplementedJobsServiceServer
// for forward compatibility
type JobsServiceServer interface {
	// @perm
	ColleaguesList(context.Context, *ColleaguesListRequest) (*ColleaguesListResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "All"}§[]string{"Own"}
	ConductListEntries(context.Context, *ConductListEntriesRequest) (*ConductListEntriesResponse, error)
	// @perm: Name=ConductListEntries
	TimeclockStats(context.Context, *TimeclockStatsRequest) (*TimeclockStatsResponse, error)
	// @perm
	ConductCreateEntry(context.Context, *ConductCreateEntryRequest) (*ConductCreateEntryResponse, error)
	// @perm
	ConductUpdateEntry(context.Context, *ConductUpdateEntryRequest) (*ConductUpdateEntryResponse, error)
	// @perm
	ConductDeleteEntry(context.Context, *ConductDeleteEntryRequest) (*ConductDeleteEntryResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "All"}§[]string{"Own"}
	TimeclockListEntries(context.Context, *TimeclockListEntriesRequest) (*TimeclockListEntriesResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "All"}§[]string{"Own"}
	RequestsListEntries(context.Context, *RequestsListEntriesRequest) (*RequestsListEntriesResponse, error)
	// @perm
	RequestsCreateEntry(context.Context, *RequestsCreateEntryRequest) (*RequestsCreateEntryResponse, error)
	// @perm
	RequestsUpdateEntry(context.Context, *RequestsUpdateEntryRequest) (*RequestsUpdateEntryResponse, error)
	// @perm
	RequestsDeleteEntry(context.Context, *RequestsDeleteEntryRequest) (*RequestsDeleteEntryResponse, error)
	// @perm: Name=RequestsListEntries
	RequestsListTypes(context.Context, *RequestsListTypesRequest) (*RequestsListTypesResponse, error)
	// @perm
	RequestsCreateOrUpdateType(context.Context, *RequestsCreateOrUpdateTypeRequest) (*RequestsCreateOrUpdateTypeResponse, error)
	// @perm
	RequestsDeleteType(context.Context, *RequestsDeleteTypeRequest) (*RequestsDeleteTypeResponse, error)
	// @perm
	RequestsPostComment(context.Context, *RequestsPostCommentRequest) (*RequestsPostCommentResponse, error)
	// @perm
	RequestsDeleteComment(context.Context, *RequestsDeleteCommentRequest) (*RequestsDeleteCommentResponse, error)
	mustEmbedUnimplementedJobsServiceServer()
}

// UnimplementedJobsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobsServiceServer struct {
}

func (UnimplementedJobsServiceServer) ColleaguesList(context.Context, *ColleaguesListRequest) (*ColleaguesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ColleaguesList not implemented")
}
func (UnimplementedJobsServiceServer) ConductListEntries(context.Context, *ConductListEntriesRequest) (*ConductListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConductListEntries not implemented")
}
func (UnimplementedJobsServiceServer) TimeclockStats(context.Context, *TimeclockStatsRequest) (*TimeclockStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeclockStats not implemented")
}
func (UnimplementedJobsServiceServer) ConductCreateEntry(context.Context, *ConductCreateEntryRequest) (*ConductCreateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConductCreateEntry not implemented")
}
func (UnimplementedJobsServiceServer) ConductUpdateEntry(context.Context, *ConductUpdateEntryRequest) (*ConductUpdateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConductUpdateEntry not implemented")
}
func (UnimplementedJobsServiceServer) ConductDeleteEntry(context.Context, *ConductDeleteEntryRequest) (*ConductDeleteEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConductDeleteEntry not implemented")
}
func (UnimplementedJobsServiceServer) TimeclockListEntries(context.Context, *TimeclockListEntriesRequest) (*TimeclockListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeclockListEntries not implemented")
}
func (UnimplementedJobsServiceServer) RequestsListEntries(context.Context, *RequestsListEntriesRequest) (*RequestsListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsListEntries not implemented")
}
func (UnimplementedJobsServiceServer) RequestsCreateEntry(context.Context, *RequestsCreateEntryRequest) (*RequestsCreateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsCreateEntry not implemented")
}
func (UnimplementedJobsServiceServer) RequestsUpdateEntry(context.Context, *RequestsUpdateEntryRequest) (*RequestsUpdateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsUpdateEntry not implemented")
}
func (UnimplementedJobsServiceServer) RequestsDeleteEntry(context.Context, *RequestsDeleteEntryRequest) (*RequestsDeleteEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsDeleteEntry not implemented")
}
func (UnimplementedJobsServiceServer) RequestsListTypes(context.Context, *RequestsListTypesRequest) (*RequestsListTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsListTypes not implemented")
}
func (UnimplementedJobsServiceServer) RequestsCreateOrUpdateType(context.Context, *RequestsCreateOrUpdateTypeRequest) (*RequestsCreateOrUpdateTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsCreateOrUpdateType not implemented")
}
func (UnimplementedJobsServiceServer) RequestsDeleteType(context.Context, *RequestsDeleteTypeRequest) (*RequestsDeleteTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsDeleteType not implemented")
}
func (UnimplementedJobsServiceServer) RequestsPostComment(context.Context, *RequestsPostCommentRequest) (*RequestsPostCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsPostComment not implemented")
}
func (UnimplementedJobsServiceServer) RequestsDeleteComment(context.Context, *RequestsDeleteCommentRequest) (*RequestsDeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsDeleteComment not implemented")
}
func (UnimplementedJobsServiceServer) mustEmbedUnimplementedJobsServiceServer() {}

// UnsafeJobsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobsServiceServer will
// result in compilation errors.
type UnsafeJobsServiceServer interface {
	mustEmbedUnimplementedJobsServiceServer()
}

func RegisterJobsServiceServer(s grpc.ServiceRegistrar, srv JobsServiceServer) {
	s.RegisterService(&JobsService_ServiceDesc, srv)
}

func _JobsService_ColleaguesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColleaguesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ColleaguesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_ColleaguesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ColleaguesList(ctx, req.(*ColleaguesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_ConductListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConductListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ConductListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_ConductListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ConductListEntries(ctx, req.(*ConductListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_TimeclockStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeclockStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).TimeclockStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_TimeclockStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).TimeclockStats(ctx, req.(*TimeclockStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_ConductCreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConductCreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ConductCreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_ConductCreateEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ConductCreateEntry(ctx, req.(*ConductCreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_ConductUpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConductUpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ConductUpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_ConductUpdateEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ConductUpdateEntry(ctx, req.(*ConductUpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_ConductDeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConductDeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ConductDeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_ConductDeleteEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ConductDeleteEntry(ctx, req.(*ConductDeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_TimeclockListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeclockListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).TimeclockListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_TimeclockListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).TimeclockListEntries(ctx, req.(*TimeclockListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsListEntries(ctx, req.(*RequestsListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsCreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsCreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsCreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsCreateEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsCreateEntry(ctx, req.(*RequestsCreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsUpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsUpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsUpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsUpdateEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsUpdateEntry(ctx, req.(*RequestsUpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsDeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsDeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsDeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsDeleteEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsDeleteEntry(ctx, req.(*RequestsDeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsListTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsListTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsListTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsListTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsListTypes(ctx, req.(*RequestsListTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsCreateOrUpdateType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsCreateOrUpdateTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsCreateOrUpdateType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsCreateOrUpdateType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsCreateOrUpdateType(ctx, req.(*RequestsCreateOrUpdateTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsDeleteType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsDeleteTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsDeleteType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsDeleteType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsDeleteType(ctx, req.(*RequestsDeleteTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsPostComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsPostCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsPostComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsPostComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsPostComment(ctx, req.(*RequestsPostCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_RequestsDeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsDeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).RequestsDeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_RequestsDeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).RequestsDeleteComment(ctx, req.(*RequestsDeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobsService_ServiceDesc is the grpc.ServiceDesc for JobsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.jobs.JobsService",
	HandlerType: (*JobsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ColleaguesList",
			Handler:    _JobsService_ColleaguesList_Handler,
		},
		{
			MethodName: "ConductListEntries",
			Handler:    _JobsService_ConductListEntries_Handler,
		},
		{
			MethodName: "TimeclockStats",
			Handler:    _JobsService_TimeclockStats_Handler,
		},
		{
			MethodName: "ConductCreateEntry",
			Handler:    _JobsService_ConductCreateEntry_Handler,
		},
		{
			MethodName: "ConductUpdateEntry",
			Handler:    _JobsService_ConductUpdateEntry_Handler,
		},
		{
			MethodName: "ConductDeleteEntry",
			Handler:    _JobsService_ConductDeleteEntry_Handler,
		},
		{
			MethodName: "TimeclockListEntries",
			Handler:    _JobsService_TimeclockListEntries_Handler,
		},
		{
			MethodName: "RequestsListEntries",
			Handler:    _JobsService_RequestsListEntries_Handler,
		},
		{
			MethodName: "RequestsCreateEntry",
			Handler:    _JobsService_RequestsCreateEntry_Handler,
		},
		{
			MethodName: "RequestsUpdateEntry",
			Handler:    _JobsService_RequestsUpdateEntry_Handler,
		},
		{
			MethodName: "RequestsDeleteEntry",
			Handler:    _JobsService_RequestsDeleteEntry_Handler,
		},
		{
			MethodName: "RequestsListTypes",
			Handler:    _JobsService_RequestsListTypes_Handler,
		},
		{
			MethodName: "RequestsCreateOrUpdateType",
			Handler:    _JobsService_RequestsCreateOrUpdateType_Handler,
		},
		{
			MethodName: "RequestsDeleteType",
			Handler:    _JobsService_RequestsDeleteType_Handler,
		},
		{
			MethodName: "RequestsPostComment",
			Handler:    _JobsService_RequestsPostComment_Handler,
		},
		{
			MethodName: "RequestsDeleteComment",
			Handler:    _JobsService_RequestsDeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/jobs/jobs.proto",
}
