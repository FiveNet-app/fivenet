// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: services/qualifications/qualifications.proto

package qualifications

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
	QualificationsService_ListQualifications_FullMethodName                 = "/services.qualifications.QualificationsService/ListQualifications"
	QualificationsService_GetQualification_FullMethodName                   = "/services.qualifications.QualificationsService/GetQualification"
	QualificationsService_CreateQualification_FullMethodName                = "/services.qualifications.QualificationsService/CreateQualification"
	QualificationsService_UpdateQualification_FullMethodName                = "/services.qualifications.QualificationsService/UpdateQualification"
	QualificationsService_DeleteQualification_FullMethodName                = "/services.qualifications.QualificationsService/DeleteQualification"
	QualificationsService_ListQualificationRequests_FullMethodName          = "/services.qualifications.QualificationsService/ListQualificationRequests"
	QualificationsService_CreateOrUpdateQualificationRequest_FullMethodName = "/services.qualifications.QualificationsService/CreateOrUpdateQualificationRequest"
	QualificationsService_DeleteQualificationReq_FullMethodName             = "/services.qualifications.QualificationsService/DeleteQualificationReq"
	QualificationsService_ListQualificationsResults_FullMethodName          = "/services.qualifications.QualificationsService/ListQualificationsResults"
	QualificationsService_CreateOrUpdateQualificationResult_FullMethodName  = "/services.qualifications.QualificationsService/CreateOrUpdateQualificationResult"
	QualificationsService_DeleteQualificationResult_FullMethodName          = "/services.qualifications.QualificationsService/DeleteQualificationResult"
	QualificationsService_GetExamInfo_FullMethodName                        = "/services.qualifications.QualificationsService/GetExamInfo"
	QualificationsService_TakeExam_FullMethodName                           = "/services.qualifications.QualificationsService/TakeExam"
	QualificationsService_SubmitExam_FullMethodName                         = "/services.qualifications.QualificationsService/SubmitExam"
	QualificationsService_GetUserExam_FullMethodName                        = "/services.qualifications.QualificationsService/GetUserExam"
)

// QualificationsServiceClient is the client API for QualificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QualificationsServiceClient interface {
	// @perm
	ListQualifications(ctx context.Context, in *ListQualificationsRequest, opts ...grpc.CallOption) (*ListQualificationsResponse, error)
	// @perm: Name=ListQualifications
	GetQualification(ctx context.Context, in *GetQualificationRequest, opts ...grpc.CallOption) (*GetQualificationResponse, error)
	// @perm
	CreateQualification(ctx context.Context, in *CreateQualificationRequest, opts ...grpc.CallOption) (*CreateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	UpdateQualification(ctx context.Context, in *UpdateQualificationRequest, opts ...grpc.CallOption) (*UpdateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	DeleteQualification(ctx context.Context, in *DeleteQualificationRequest, opts ...grpc.CallOption) (*DeleteQualificationResponse, error)
	// @perm: Name=ListQualifications
	ListQualificationRequests(ctx context.Context, in *ListQualificationRequestsRequest, opts ...grpc.CallOption) (*ListQualificationRequestsResponse, error)
	// @perm: Name=ListQualifications
	CreateOrUpdateQualificationRequest(ctx context.Context, in *CreateOrUpdateQualificationRequestRequest, opts ...grpc.CallOption) (*CreateOrUpdateQualificationRequestResponse, error)
	// @perm
	DeleteQualificationReq(ctx context.Context, in *DeleteQualificationReqRequest, opts ...grpc.CallOption) (*DeleteQualificationReqResponse, error)
	// @perm: Name=ListQualifications
	ListQualificationsResults(ctx context.Context, in *ListQualificationsResultsRequest, opts ...grpc.CallOption) (*ListQualificationsResultsResponse, error)
	// @perm
	CreateOrUpdateQualificationResult(ctx context.Context, in *CreateOrUpdateQualificationResultRequest, opts ...grpc.CallOption) (*CreateOrUpdateQualificationResultResponse, error)
	// @perm
	DeleteQualificationResult(ctx context.Context, in *DeleteQualificationResultRequest, opts ...grpc.CallOption) (*DeleteQualificationResultResponse, error)
	// @perm: Name=ListQualifications
	GetExamInfo(ctx context.Context, in *GetExamInfoRequest, opts ...grpc.CallOption) (*GetExamInfoResponse, error)
	// @perm: Name=ListQualifications
	TakeExam(ctx context.Context, in *TakeExamRequest, opts ...grpc.CallOption) (*TakeExamResponse, error)
	// @perm: Name=ListQualifications
	SubmitExam(ctx context.Context, in *SubmitExamRequest, opts ...grpc.CallOption) (*SubmitExamResponse, error)
	// @perm: Name=CreateOrUpdateQualificationResult
	GetUserExam(ctx context.Context, in *GetUserExamRequest, opts ...grpc.CallOption) (*GetUserExamResponse, error)
}

type qualificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQualificationsServiceClient(cc grpc.ClientConnInterface) QualificationsServiceClient {
	return &qualificationsServiceClient{cc}
}

func (c *qualificationsServiceClient) ListQualifications(ctx context.Context, in *ListQualificationsRequest, opts ...grpc.CallOption) (*ListQualificationsResponse, error) {
	out := new(ListQualificationsResponse)
	err := c.cc.Invoke(ctx, QualificationsService_ListQualifications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) GetQualification(ctx context.Context, in *GetQualificationRequest, opts ...grpc.CallOption) (*GetQualificationResponse, error) {
	out := new(GetQualificationResponse)
	err := c.cc.Invoke(ctx, QualificationsService_GetQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) CreateQualification(ctx context.Context, in *CreateQualificationRequest, opts ...grpc.CallOption) (*CreateQualificationResponse, error) {
	out := new(CreateQualificationResponse)
	err := c.cc.Invoke(ctx, QualificationsService_CreateQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) UpdateQualification(ctx context.Context, in *UpdateQualificationRequest, opts ...grpc.CallOption) (*UpdateQualificationResponse, error) {
	out := new(UpdateQualificationResponse)
	err := c.cc.Invoke(ctx, QualificationsService_UpdateQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) DeleteQualification(ctx context.Context, in *DeleteQualificationRequest, opts ...grpc.CallOption) (*DeleteQualificationResponse, error) {
	out := new(DeleteQualificationResponse)
	err := c.cc.Invoke(ctx, QualificationsService_DeleteQualification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) ListQualificationRequests(ctx context.Context, in *ListQualificationRequestsRequest, opts ...grpc.CallOption) (*ListQualificationRequestsResponse, error) {
	out := new(ListQualificationRequestsResponse)
	err := c.cc.Invoke(ctx, QualificationsService_ListQualificationRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) CreateOrUpdateQualificationRequest(ctx context.Context, in *CreateOrUpdateQualificationRequestRequest, opts ...grpc.CallOption) (*CreateOrUpdateQualificationRequestResponse, error) {
	out := new(CreateOrUpdateQualificationRequestResponse)
	err := c.cc.Invoke(ctx, QualificationsService_CreateOrUpdateQualificationRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) DeleteQualificationReq(ctx context.Context, in *DeleteQualificationReqRequest, opts ...grpc.CallOption) (*DeleteQualificationReqResponse, error) {
	out := new(DeleteQualificationReqResponse)
	err := c.cc.Invoke(ctx, QualificationsService_DeleteQualificationReq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) ListQualificationsResults(ctx context.Context, in *ListQualificationsResultsRequest, opts ...grpc.CallOption) (*ListQualificationsResultsResponse, error) {
	out := new(ListQualificationsResultsResponse)
	err := c.cc.Invoke(ctx, QualificationsService_ListQualificationsResults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) CreateOrUpdateQualificationResult(ctx context.Context, in *CreateOrUpdateQualificationResultRequest, opts ...grpc.CallOption) (*CreateOrUpdateQualificationResultResponse, error) {
	out := new(CreateOrUpdateQualificationResultResponse)
	err := c.cc.Invoke(ctx, QualificationsService_CreateOrUpdateQualificationResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) DeleteQualificationResult(ctx context.Context, in *DeleteQualificationResultRequest, opts ...grpc.CallOption) (*DeleteQualificationResultResponse, error) {
	out := new(DeleteQualificationResultResponse)
	err := c.cc.Invoke(ctx, QualificationsService_DeleteQualificationResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) GetExamInfo(ctx context.Context, in *GetExamInfoRequest, opts ...grpc.CallOption) (*GetExamInfoResponse, error) {
	out := new(GetExamInfoResponse)
	err := c.cc.Invoke(ctx, QualificationsService_GetExamInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) TakeExam(ctx context.Context, in *TakeExamRequest, opts ...grpc.CallOption) (*TakeExamResponse, error) {
	out := new(TakeExamResponse)
	err := c.cc.Invoke(ctx, QualificationsService_TakeExam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) SubmitExam(ctx context.Context, in *SubmitExamRequest, opts ...grpc.CallOption) (*SubmitExamResponse, error) {
	out := new(SubmitExamResponse)
	err := c.cc.Invoke(ctx, QualificationsService_SubmitExam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qualificationsServiceClient) GetUserExam(ctx context.Context, in *GetUserExamRequest, opts ...grpc.CallOption) (*GetUserExamResponse, error) {
	out := new(GetUserExamResponse)
	err := c.cc.Invoke(ctx, QualificationsService_GetUserExam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QualificationsServiceServer is the server API for QualificationsService service.
// All implementations must embed UnimplementedQualificationsServiceServer
// for forward compatibility
type QualificationsServiceServer interface {
	// @perm
	ListQualifications(context.Context, *ListQualificationsRequest) (*ListQualificationsResponse, error)
	// @perm: Name=ListQualifications
	GetQualification(context.Context, *GetQualificationRequest) (*GetQualificationResponse, error)
	// @perm
	CreateQualification(context.Context, *CreateQualificationRequest) (*CreateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	UpdateQualification(context.Context, *UpdateQualificationRequest) (*UpdateQualificationResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	DeleteQualification(context.Context, *DeleteQualificationRequest) (*DeleteQualificationResponse, error)
	// @perm: Name=ListQualifications
	ListQualificationRequests(context.Context, *ListQualificationRequestsRequest) (*ListQualificationRequestsResponse, error)
	// @perm: Name=ListQualifications
	CreateOrUpdateQualificationRequest(context.Context, *CreateOrUpdateQualificationRequestRequest) (*CreateOrUpdateQualificationRequestResponse, error)
	// @perm
	DeleteQualificationReq(context.Context, *DeleteQualificationReqRequest) (*DeleteQualificationReqResponse, error)
	// @perm: Name=ListQualifications
	ListQualificationsResults(context.Context, *ListQualificationsResultsRequest) (*ListQualificationsResultsResponse, error)
	// @perm
	CreateOrUpdateQualificationResult(context.Context, *CreateOrUpdateQualificationResultRequest) (*CreateOrUpdateQualificationResultResponse, error)
	// @perm
	DeleteQualificationResult(context.Context, *DeleteQualificationResultRequest) (*DeleteQualificationResultResponse, error)
	// @perm: Name=ListQualifications
	GetExamInfo(context.Context, *GetExamInfoRequest) (*GetExamInfoResponse, error)
	// @perm: Name=ListQualifications
	TakeExam(context.Context, *TakeExamRequest) (*TakeExamResponse, error)
	// @perm: Name=ListQualifications
	SubmitExam(context.Context, *SubmitExamRequest) (*SubmitExamResponse, error)
	// @perm: Name=CreateOrUpdateQualificationResult
	GetUserExam(context.Context, *GetUserExamRequest) (*GetUserExamResponse, error)
	mustEmbedUnimplementedQualificationsServiceServer()
}

// UnimplementedQualificationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQualificationsServiceServer struct {
}

func (UnimplementedQualificationsServiceServer) ListQualifications(context.Context, *ListQualificationsRequest) (*ListQualificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQualifications not implemented")
}
func (UnimplementedQualificationsServiceServer) GetQualification(context.Context, *GetQualificationRequest) (*GetQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQualification not implemented")
}
func (UnimplementedQualificationsServiceServer) CreateQualification(context.Context, *CreateQualificationRequest) (*CreateQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQualification not implemented")
}
func (UnimplementedQualificationsServiceServer) UpdateQualification(context.Context, *UpdateQualificationRequest) (*UpdateQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQualification not implemented")
}
func (UnimplementedQualificationsServiceServer) DeleteQualification(context.Context, *DeleteQualificationRequest) (*DeleteQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQualification not implemented")
}
func (UnimplementedQualificationsServiceServer) ListQualificationRequests(context.Context, *ListQualificationRequestsRequest) (*ListQualificationRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQualificationRequests not implemented")
}
func (UnimplementedQualificationsServiceServer) CreateOrUpdateQualificationRequest(context.Context, *CreateOrUpdateQualificationRequestRequest) (*CreateOrUpdateQualificationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateQualificationRequest not implemented")
}
func (UnimplementedQualificationsServiceServer) DeleteQualificationReq(context.Context, *DeleteQualificationReqRequest) (*DeleteQualificationReqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQualificationReq not implemented")
}
func (UnimplementedQualificationsServiceServer) ListQualificationsResults(context.Context, *ListQualificationsResultsRequest) (*ListQualificationsResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQualificationsResults not implemented")
}
func (UnimplementedQualificationsServiceServer) CreateOrUpdateQualificationResult(context.Context, *CreateOrUpdateQualificationResultRequest) (*CreateOrUpdateQualificationResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateQualificationResult not implemented")
}
func (UnimplementedQualificationsServiceServer) DeleteQualificationResult(context.Context, *DeleteQualificationResultRequest) (*DeleteQualificationResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQualificationResult not implemented")
}
func (UnimplementedQualificationsServiceServer) GetExamInfo(context.Context, *GetExamInfoRequest) (*GetExamInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExamInfo not implemented")
}
func (UnimplementedQualificationsServiceServer) TakeExam(context.Context, *TakeExamRequest) (*TakeExamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeExam not implemented")
}
func (UnimplementedQualificationsServiceServer) SubmitExam(context.Context, *SubmitExamRequest) (*SubmitExamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitExam not implemented")
}
func (UnimplementedQualificationsServiceServer) GetUserExam(context.Context, *GetUserExamRequest) (*GetUserExamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserExam not implemented")
}
func (UnimplementedQualificationsServiceServer) mustEmbedUnimplementedQualificationsServiceServer() {}

// UnsafeQualificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QualificationsServiceServer will
// result in compilation errors.
type UnsafeQualificationsServiceServer interface {
	mustEmbedUnimplementedQualificationsServiceServer()
}

func RegisterQualificationsServiceServer(s grpc.ServiceRegistrar, srv QualificationsServiceServer) {
	s.RegisterService(&QualificationsService_ServiceDesc, srv)
}

func _QualificationsService_ListQualifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQualificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).ListQualifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_ListQualifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).ListQualifications(ctx, req.(*ListQualificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_GetQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).GetQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_GetQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).GetQualification(ctx, req.(*GetQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_CreateQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).CreateQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_CreateQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).CreateQualification(ctx, req.(*CreateQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_UpdateQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).UpdateQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_UpdateQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).UpdateQualification(ctx, req.(*UpdateQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_DeleteQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).DeleteQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_DeleteQualification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).DeleteQualification(ctx, req.(*DeleteQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_ListQualificationRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQualificationRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).ListQualificationRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_ListQualificationRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).ListQualificationRequests(ctx, req.(*ListQualificationRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_CreateOrUpdateQualificationRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateQualificationRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).CreateOrUpdateQualificationRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_CreateOrUpdateQualificationRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).CreateOrUpdateQualificationRequest(ctx, req.(*CreateOrUpdateQualificationRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_DeleteQualificationReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQualificationReqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).DeleteQualificationReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_DeleteQualificationReq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).DeleteQualificationReq(ctx, req.(*DeleteQualificationReqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_ListQualificationsResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQualificationsResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).ListQualificationsResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_ListQualificationsResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).ListQualificationsResults(ctx, req.(*ListQualificationsResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_CreateOrUpdateQualificationResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateQualificationResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).CreateOrUpdateQualificationResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_CreateOrUpdateQualificationResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).CreateOrUpdateQualificationResult(ctx, req.(*CreateOrUpdateQualificationResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_DeleteQualificationResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQualificationResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).DeleteQualificationResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_DeleteQualificationResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).DeleteQualificationResult(ctx, req.(*DeleteQualificationResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_GetExamInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExamInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).GetExamInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_GetExamInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).GetExamInfo(ctx, req.(*GetExamInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_TakeExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeExamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).TakeExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_TakeExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).TakeExam(ctx, req.(*TakeExamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_SubmitExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitExamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).SubmitExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_SubmitExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).SubmitExam(ctx, req.(*SubmitExamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QualificationsService_GetUserExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserExamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QualificationsServiceServer).GetUserExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QualificationsService_GetUserExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QualificationsServiceServer).GetUserExam(ctx, req.(*GetUserExamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QualificationsService_ServiceDesc is the grpc.ServiceDesc for QualificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QualificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.qualifications.QualificationsService",
	HandlerType: (*QualificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListQualifications",
			Handler:    _QualificationsService_ListQualifications_Handler,
		},
		{
			MethodName: "GetQualification",
			Handler:    _QualificationsService_GetQualification_Handler,
		},
		{
			MethodName: "CreateQualification",
			Handler:    _QualificationsService_CreateQualification_Handler,
		},
		{
			MethodName: "UpdateQualification",
			Handler:    _QualificationsService_UpdateQualification_Handler,
		},
		{
			MethodName: "DeleteQualification",
			Handler:    _QualificationsService_DeleteQualification_Handler,
		},
		{
			MethodName: "ListQualificationRequests",
			Handler:    _QualificationsService_ListQualificationRequests_Handler,
		},
		{
			MethodName: "CreateOrUpdateQualificationRequest",
			Handler:    _QualificationsService_CreateOrUpdateQualificationRequest_Handler,
		},
		{
			MethodName: "DeleteQualificationReq",
			Handler:    _QualificationsService_DeleteQualificationReq_Handler,
		},
		{
			MethodName: "ListQualificationsResults",
			Handler:    _QualificationsService_ListQualificationsResults_Handler,
		},
		{
			MethodName: "CreateOrUpdateQualificationResult",
			Handler:    _QualificationsService_CreateOrUpdateQualificationResult_Handler,
		},
		{
			MethodName: "DeleteQualificationResult",
			Handler:    _QualificationsService_DeleteQualificationResult_Handler,
		},
		{
			MethodName: "GetExamInfo",
			Handler:    _QualificationsService_GetExamInfo_Handler,
		},
		{
			MethodName: "TakeExam",
			Handler:    _QualificationsService_TakeExam_Handler,
		},
		{
			MethodName: "SubmitExam",
			Handler:    _QualificationsService_SubmitExam_Handler,
		},
		{
			MethodName: "GetUserExam",
			Handler:    _QualificationsService_GetUserExam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/qualifications/qualifications.proto",
}
