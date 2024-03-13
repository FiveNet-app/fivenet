package jobs

import (
	"context"
	"errors"

	database "github.com/galexrt/fivenet/gen/go/proto/resources/common/database"
	jobs "github.com/galexrt/fivenet/gen/go/proto/resources/jobs"
	"github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	errorsjobs "github.com/galexrt/fivenet/gen/go/proto/services/jobs/errors"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/grpc/errswrap"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

var (
	tQualiResults = table.FivenetJobsQualificationsResults.AS("qualificationresult")
)

func (s *Server) ListQualificationsResults(ctx context.Context, req *ListQualificationsResultsRequest) (*ListQualificationsResultsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	condition := tQualiResults.UserID.EQ(jet.Int32(userInfo.UserId))

	if req.QualificationId != nil {
		check, err := s.checkIfUserHasAccessToQuali(ctx, *req.QualificationId, userInfo, jobs.AccessLevel_ACCESS_LEVEL_GRADE)
		if err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
		if !check {
			return nil, errorsjobs.ErrFailedQuery
		}

		condition = condition.AND(tQualiResults.QualificationID.EQ(jet.Uint64(*req.QualificationId)))
	} else {
		condition = condition.AND(jet.AND(
			tQuali.DeletedAt.IS_NULL(),
			jet.OR(
				tQuali.CreatorID.EQ(jet.Int32(userInfo.UserId)),
				jet.AND(
					tQJobAccess.Access.IS_NOT_NULL(),
					tQJobAccess.Access.NOT_EQ(jet.Int32(int32(jobs.AccessLevel_ACCESS_LEVEL_BLOCKED))),
				),
			),
		))
	}

	if len(req.Status) > 0 {
		statuses := []jet.Expression{}
		for i := 0; i < len(req.Status); i++ {
			statuses = append(statuses, jet.Int16(int16(req.Status[i])))
		}

		condition = condition.AND(tQualiResults.Status.IN(statuses...))
	}

	countStmt := tQualiResults.
		SELECT(
			jet.COUNT(tQualiResults.ID).AS("datacount.totalcount"),
		).
		FROM(
			tQualiResults.
				INNER_JOIN(tQuali,
					tQuali.ID.EQ(tQualiResults.QualificationID),
				).
				LEFT_JOIN(tCreator,
					tQuali.CreatorID.EQ(tCreator.ID),
				).
				LEFT_JOIN(tQJobAccess,
					tQJobAccess.QualificationID.EQ(tQuali.ID).
						AND(tQJobAccess.Job.EQ(jet.String(userInfo.Job))).
						AND(tQJobAccess.MinimumGrade.LT_EQ(jet.Int32(userInfo.JobGrade))),
				),
		).
		WHERE(condition)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		return nil, errswrap.NewError(errorsjobs.ErrFailedQuery, err)
	}

	pag, limit := req.Pagination.GetResponseWithPageSize(count.TotalCount, QualificationsPageSize)
	resp := &ListQualificationsResultsResponse{
		Pagination: pag,
		Results:    []*jobs.Qualification{},
	}
	if count.TotalCount <= 0 {
		return resp, nil
	}

	stmt := tQualiResults.
		SELECT(
			tQualiResults.ID,
			tQualiResults.CreatedAt,
			tQualiResults.QualificationID,
			tQualiResults.UserID,
			tQualiResults.Status,
			tQualiResults.Score,
			tQualiResults.Summary,
			tQuali.ID,
			tQuali.CreatedAt,
			tQuali.UpdatedAt,
			tQuali.Job,
			tQuali.Closed,
			tQuali.Abbreviation,
			tQuali.Title,
			tQuali.Description,
			tQuali.Content,
			tQuali.CreatorJob,
			tQuali.CreatorID,
			tCreator.ID,
			tCreator.Identifier,
			tCreator.Job,
			tCreator.JobGrade,
			tCreator.Firstname,
			tCreator.Lastname,
			tCreator.Dateofbirth,
		).
		FROM(
			tQualiResults.
				INNER_JOIN(tQuali,
					tQuali.ID.EQ(tQualiResults.QualificationID),
				).
				LEFT_JOIN(tCreator,
					tQuali.CreatorID.EQ(tCreator.ID),
				).
				LEFT_JOIN(tQJobAccess,
					tQJobAccess.QualificationID.EQ(tQuali.ID).
						AND(tQJobAccess.Job.EQ(jet.String(userInfo.Job))).
						AND(tQJobAccess.MinimumGrade.LT_EQ(jet.Int32(userInfo.JobGrade))),
				),
		).
		WHERE(condition).
		OFFSET(req.Pagination.Offset).
		LIMIT(limit)

	if err := stmt.QueryContext(ctx, s.db, &resp.Results); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	return resp, nil
}

func (s *Server) CreateOrUpdateQualificationResult(ctx context.Context, req *CreateOrUpdateQualificationResultRequest) (*CreateOrUpdateQualificationResultResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: JobsQualificationsService_ServiceDesc.ServiceName,
		Method:  "CreateOrUpdateQualificationResult",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	ok, err := s.checkIfUserHasAccessToQuali(ctx, req.Result.QualificationId, userInfo, jobs.AccessLevel_ACCESS_LEVEL_GRADE)
	if err != nil {
		return nil, errswrap.NewError(errorsjobs.ErrFailedQuery, err)
	}
	if !ok {
		return nil, errorsjobs.ErrFailedQuery
	}

	if req.Result.Id <= 0 {
		stmt := tQualiResults.
			INSERT(
				tQualiResults.QualificationID,
				tQualiResults.UserID,
				tQualiResults.Status,
				tQualiResults.Score,
				tQualiResults.Summary,
				tQualiResults.CreatorID,
				tQualiResults.CreatorJob,
			).
			VALUES(
				req.Result.QualificationId,
				req.Result.UserId,
				req.Result.Status,
				req.Result.Score,
				req.Result.Summary,
				userInfo.UserId,
				userInfo.Job,
			)

		res, err := stmt.ExecContext(ctx, s.db)
		if err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}

		lastId, err := res.LastInsertId()
		if err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}

		req.Result.Id = uint64(lastId)

		auditEntry.State = int16(rector.EventType_EVENT_TYPE_CREATED)
	} else {
		stmt := tQualiResults.
			UPDATE(
				tQualiResults.QualificationID,
				tQualiResults.UserID,
				tQualiResults.Status,
				tQualiResults.Score,
				tQualiResults.Summary,
			).
			SET(
				tQualiResults.QualificationID.SET(jet.Uint64(req.Result.QualificationId)),
				tQualiResults.UserID.SET(jet.Int32(req.Result.UserId)),
				tQualiResults.Status.SET(jet.Int16(int16(req.Result.Status))),
				tQualiResults.Score.SET(jet.Uint32(*req.Result.Score)),
				tQualiResults.Summary.SET(jet.String(req.Result.Summary)),
			).
			WHERE(tQualiResults.ID.EQ(jet.Uint64(req.Result.Id)))

		if _, err := stmt.ExecContext(ctx, s.db); err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}

		auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)
	}

	result, err := s.getQualificationResult(ctx, tQualiResults.ID.EQ(jet.Uint64(req.Result.Id)), userInfo)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}

	return &CreateOrUpdateQualificationResultResponse{
		Result: result,
	}, nil
}

func (s *Server) getQualificationResult(ctx context.Context, condition jet.BoolExpression, userInfo *userinfo.UserInfo) (*jobs.QualificationResult, error) {
	var quali jobs.QualificationResult

	stmt := s.getQualificationQuery(condition, nil, userInfo).
		LIMIT(1)

	if err := stmt.QueryContext(ctx, s.db, &quali); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(errorsjobs.ErrFailedQuery, err)
		}
	}

	if quali.Creator != nil {
		s.enricher.EnrichJobInfo(quali.Creator)
	}

	return &quali, nil
}
