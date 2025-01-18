package qualifications

import (
	"context"

	"github.com/fivenet-app/fivenet/gen/go/proto/resources/qualifications"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/rector"
	pbqualifications "github.com/fivenet-app/fivenet/gen/go/proto/services/qualifications"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	"github.com/fivenet-app/fivenet/pkg/grpc/errswrap"
	"github.com/fivenet-app/fivenet/query/fivenet/model"
	errorsqualifications "github.com/fivenet-app/fivenet/services/qualifications/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (s *Server) GetQualificationAccess(ctx context.Context, req *pbqualifications.GetQualificationAccessRequest) (*pbqualifications.GetQualificationAccessResponse, error) {
	trace.SpanFromContext(ctx).SetAttributes(attribute.Int64("fivenet.qualifications.id", int64(req.QualificationId)))

	userInfo := auth.MustGetUserInfoFromContext(ctx)
	check, err := s.access.CanUserAccessTarget(ctx, req.QualificationId, userInfo, qualifications.AccessLevel_ACCESS_LEVEL_VIEW)
	if err != nil {
		return nil, errswrap.NewError(err, errorsqualifications.ErrFailedQuery)
	}
	if !check {
		return nil, errorsqualifications.ErrFailedQuery
	}

	access, err := s.access.Jobs.List(ctx, s.db, req.QualificationId)
	if err != nil {
		return nil, errswrap.NewError(err, errorsqualifications.ErrFailedQuery)
	}

	for i := 0; i < len(access); i++ {
		s.enricher.EnrichJobInfo(access[i])
	}

	resp := &pbqualifications.GetQualificationAccessResponse{
		Access: &qualifications.QualificationAccess{
			Jobs: access,
		},
	}

	return resp, nil
}

func (s *Server) SetQualificationAccess(ctx context.Context, req *pbqualifications.SetQualificationAccessRequest) (*pbqualifications.SetQualificationAccessResponse, error) {
	trace.SpanFromContext(ctx).SetAttributes(attribute.Int64("fivenet.qualifications.id", int64(req.QualificationId)))

	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: pbqualifications.QualificationsService_ServiceDesc.ServiceName,
		Method:  "SetQualificationAccess",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	check, err := s.access.CanUserAccessTarget(ctx, req.QualificationId, userInfo, qualifications.AccessLevel_ACCESS_LEVEL_EDIT)
	if err != nil {
		return nil, errswrap.NewError(err, errorsqualifications.ErrFailedQuery)
	}
	if !check {
		return nil, errorsqualifications.ErrFailedQuery
	}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errswrap.NewError(err, errorsqualifications.ErrFailedQuery)
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	if req.Access != nil {
		if _, err := s.access.HandleAccessChanges(ctx, tx, req.QualificationId, req.Access.Jobs, nil, nil); err != nil {
			return nil, errswrap.NewError(err, errorsqualifications.ErrFailedQuery)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, errswrap.NewError(err, errorsqualifications.ErrFailedQuery)
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &pbqualifications.SetQualificationAccessResponse{}, nil
}