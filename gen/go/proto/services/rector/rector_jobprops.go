package rector

import (
	"context"
	"errors"
	"strings"

	"github.com/galexrt/fivenet/gen/go/proto/resources/jobs"
	rector "github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	"github.com/galexrt/fivenet/pkg/auth"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

var (
	jobProps = table.FivenetJobProps
)

func (s *Server) GetJobProps(ctx context.Context, req *GetJobPropsRequest) (*GetJobPropsResponse, error) {
	_, job, _ := auth.GetUserInfoFromContext(ctx)

	jobProps := table.FivenetJobProps.AS("jobprops")
	stmt := jobProps.
		SELECT(
			jobProps.AllColumns,
		).
		FROM(jobProps).
		WHERE(
			jobProps.Job.EQ(jet.String(job)),
		).
		LIMIT(1)

	resp := &GetJobPropsResponse{
		JobProps: &jobs.JobProps{},
	}
	if err := stmt.QueryContext(ctx, s.db, resp.JobProps); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, err
		}
	}

	resp.JobProps.Default(job)

	return resp, nil
}
func (s *Server) SetJobProps(ctx context.Context, req *SetJobPropsRequest) (*SetJobPropsResponse, error) {
	userId, job, _ := auth.GetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorService_ServiceDesc.ServiceName,
		Method:  "SetJobProps",
		UserID:  userId,
		UserJob: job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	// Ensure that the job is the user's job
	req.JobProps.Job = job

	req.JobProps.LivemapMarkerColor = strings.ReplaceAll(req.JobProps.LivemapMarkerColor, "#", "")

	stmt := jobProps.
		INSERT(
			jobProps.Job,
			jobProps.Theme,
			jobProps.LivemapMarkerColor,
		).
		VALUES(
			req.JobProps.Job,
			req.JobProps.Theme,
			req.JobProps.LivemapMarkerColor,
		).
		ON_DUPLICATE_KEY_UPDATE(
			jobProps.Theme.SET(jet.String(req.JobProps.Theme)),
			jobProps.LivemapMarkerColor.SET(jet.String(req.JobProps.LivemapMarkerColor)),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	auditEntry.State = int16(rector.EVENT_TYPE_UPDATED)

	return &SetJobPropsResponse{}, nil
}
