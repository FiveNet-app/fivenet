package errorsqualifications

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFailedQuery = status.Error(codes.Internal, "errors.QualificationsService.ErrFailedQuery")
)
