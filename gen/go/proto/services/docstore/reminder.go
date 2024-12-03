package docstore

import (
	"context"
	"time"

	"github.com/fivenet-app/fivenet/gen/go/proto/resources/documents"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/rector"
	errorsdocstore "github.com/fivenet-app/fivenet/gen/go/proto/services/docstore/errors"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	"github.com/fivenet-app/fivenet/pkg/grpc/errswrap"
	"github.com/fivenet-app/fivenet/query/fivenet/model"
	"github.com/fivenet-app/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (s *Server) createOrUpdateWorkflowState(ctx context.Context, tx qrm.DB, documentId uint64, workflow *documents.Workflow) error {
	if workflow == nil || (!workflow.AutoClose && !workflow.Reminder) {
		return nil
	}

	now := time.Now()

	autoCloseTime := jet.TimestampExp(jet.NULL)
	if workflow.AutoClose && workflow.AutoCloseSettings != nil && workflow.AutoCloseSettings.Duration != nil {
		autoCloseTime = jet.TimestampT(now.Add(workflow.AutoCloseSettings.Duration.AsDuration()))
	}

	nextReminderTime := jet.TimestampExp(jet.NULL)
	if workflow.Reminder && workflow.ReminderSettings != nil && len(workflow.ReminderSettings.Reminders) > 0 {
		reminder := workflow.ReminderSettings.Reminders[0]

		nextReminderTime = jet.TimestampT(now.Add(reminder.Duration.AsDuration()))
	}

	tWorkflow := table.FivenetDocumentsWorkflowState
	stmt := tWorkflow.
		INSERT(
			tWorkflow.DocumentID,
			tWorkflow.AutoCloseTime,
			tWorkflow.NextReminderTime,
			tWorkflow.NextReminderCount,
		).
		VALUES(
			documentId,
			autoCloseTime,
			nextReminderTime,
			jet.NULL,
		).
		ON_DUPLICATE_KEY_UPDATE(
			tWorkflow.AutoCloseTime.SET(autoCloseTime),
			tWorkflow.NextReminderTime.SET(nextReminderTime),
			tWorkflow.NextReminderCount.SET(jet.IntExp(jet.NULL)),
		)

	if _, err := stmt.ExecContext(ctx, tx); err != nil {
		return err
	}

	return nil
}

func (s *Server) SetDocumentReminder(ctx context.Context, req *SetDocumentReminderRequest) (*SetDocumentReminderResponse, error) {
	trace.SpanFromContext(ctx).SetAttributes(attribute.Int64("fivenet.docstore.id", int64(req.DocumentId)))

	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "SetDocumentReminder",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	check, err := s.access.CanUserAccessTarget(ctx, req.DocumentId, userInfo, documents.AccessLevel_ACCESS_LEVEL_VIEW)
	if err != nil {
		return nil, errswrap.NewError(err, errorsdocstore.ErrFailedQuery)
	}
	if !check {
		return nil, errorsdocstore.ErrDocViewDenied
	}

	if req.ReminderTime == nil {
		if err := deleteWorkflowUserState(ctx, s.db, &documents.WorkflowUserState{
			DocumentId: req.DocumentId,
			UserId:     userInfo.UserId,
		}); err != nil {
			return nil, errswrap.NewError(err, errorsdocstore.ErrFailedQuery)
		}
	} else {
		if err := updateWorkflowUserState(ctx, s.db, &documents.WorkflowUserState{
			DocumentId:            req.DocumentId,
			UserId:                userInfo.UserId,
			ManualReminderTime:    req.ReminderTime,
			ManualReminderMessage: req.Message,
		}); err != nil {
			return nil, errswrap.NewError(err, errorsdocstore.ErrFailedQuery)
		}
	}

	return &SetDocumentReminderResponse{}, nil
}
