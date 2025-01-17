package jobs

import (
	"context"
	"errors"
	"slices"
	"strings"

	database "github.com/fivenet-app/fivenet/gen/go/proto/resources/common/database"
	jobs "github.com/fivenet-app/fivenet/gen/go/proto/resources/jobs"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/rector"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/users"
	errorsjobs "github.com/fivenet-app/fivenet/gen/go/proto/services/jobs/errors"
	permsjobs "github.com/fivenet-app/fivenet/gen/go/proto/services/jobs/perms"
	"github.com/fivenet-app/fivenet/pkg/access"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth/userinfo"
	"github.com/fivenet-app/fivenet/pkg/grpc/errswrap"
	"github.com/fivenet-app/fivenet/pkg/perms"
	"github.com/fivenet-app/fivenet/pkg/utils"
	"github.com/fivenet-app/fivenet/query/fivenet/model"
	"github.com/fivenet-app/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

const (
	ColleaguesDefaultPageSize = 20
)

var (
	tJobsUserProps    = table.FivenetJobsUserProps.AS("jobs_user_props")
	tJobsUserActivity = table.FivenetJobsUserActivity
)

func (s *Server) ListColleagues(ctx context.Context, req *ListColleaguesRequest) (*ListColleaguesResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	// Access Permission Check
	typesAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceGetColleaguePerm, permsjobs.JobsServiceGetColleagueTypesPermField)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	var types perms.StringList
	if typesAttr != nil {
		types = typesAttr.([]string)
	}

	tUser := tUser.AS("colleague")
	condition := tUser.Job.EQ(jet.String(userInfo.Job)).
		AND(s.customDB.Conditions.User.GetFilter(tUser.Alias()))

	if req.UserId != nil && *req.UserId > 0 {
		condition = condition.AND(tUser.ID.EQ(jet.Int32(*req.UserId)))
	} else {
		req.Search = strings.TrimSpace(req.Search)
		req.Search = strings.ReplaceAll(req.Search, "%", "")
		req.Search = strings.ReplaceAll(req.Search, " ", "%")
		if req.Search != "" {
			req.Search = "%" + req.Search + "%"
			condition = condition.AND(
				jet.CONCAT(tUser.Firstname, jet.String(" "), tUser.Lastname).
					LIKE(jet.String(req.Search)),
			)
		}
	}

	if req.Absent != nil && *req.Absent {
		condition = condition.AND(
			jet.AND(
				tJobsUserProps.AbsenceBegin.IS_NOT_NULL(),
				tJobsUserProps.AbsenceEnd.IS_NOT_NULL(),
				tJobsUserProps.AbsenceBegin.LT_EQ(jet.CURRENT_DATE()),
				tJobsUserProps.AbsenceEnd.GT_EQ(jet.CURRENT_DATE()),
			))
	}

	if (slices.Contains(types, "Labels") || userInfo.SuperUser) && len(req.LabelIds) > 0 {
		labelIds := []jet.Expression{}
		for _, labelId := range req.LabelIds {
			labelIds = append(labelIds, jet.Uint64(labelId))
		}

		condition = condition.AND(
			tUserLabels.LabelID.IN(labelIds...),
		)
	}

	if req.NamePrefix != nil {
		*req.NamePrefix = strings.TrimSpace(*req.NamePrefix)
		*req.NamePrefix = strings.ReplaceAll(*req.NamePrefix, "%", "")
		*req.NamePrefix = strings.ReplaceAll(*req.NamePrefix, " ", "%")
		if *req.NamePrefix != "" {
			*req.NamePrefix = "%" + *req.NamePrefix + "%"
			condition = condition.AND(jet.AND(
				tJobsUserProps.NamePrefix.IS_NOT_NULL(),
				tJobsUserProps.NamePrefix.LIKE(jet.String(*req.NamePrefix)),
			))
		}
	}

	if req.NameSuffix != nil {
		*req.NameSuffix = strings.TrimSpace(*req.NameSuffix)
		*req.NameSuffix = strings.ReplaceAll(*req.NameSuffix, "%", "")
		*req.NameSuffix = strings.ReplaceAll(*req.NameSuffix, " ", "%")
		if *req.NameSuffix != "" {
			*req.NameSuffix = "%" + *req.NameSuffix + "%"
			condition = condition.AND(jet.AND(
				tJobsUserProps.NameSuffix.IS_NOT_NULL(),
				tJobsUserProps.NameSuffix.LIKE(jet.String(*req.NameSuffix)),
			))
		}
	}

	// Get total count of values
	countStmt := tUser.
		SELECT(
			jet.COUNT(tUser.ID).AS("datacount.totalcount"),
		).
		OPTIMIZER_HINTS(jet.OptimizerHint("idx_users_firstname_lastname_fulltext"))

	if (slices.Contains(types, "Labels") || userInfo.SuperUser) && len(req.LabelIds) > 0 {
		countStmt = countStmt.
			FROM(
				tUser.
					LEFT_JOIN(tJobsUserProps,
						tJobsUserProps.UserID.EQ(tUser.ID).
							AND(tUser.Job.EQ(jet.String(userInfo.Job))),
					).
					INNER_JOIN(tUserLabels,
						tUserLabels.UserID.EQ(tUser.ID).
							AND(tUserLabels.Job.EQ(jet.String(userInfo.Job))),
					).
					LEFT_JOIN(tJobLabels,
						tJobLabels.ID.EQ(tUserLabels.LabelID),
					),
			)
	} else {
		countStmt = countStmt.
			FROM(
				tUser.
					LEFT_JOIN(tJobsUserProps,
						tJobsUserProps.UserID.EQ(tUser.ID).
							AND(tUser.Job.EQ(jet.String(userInfo.Job))),
					),
			)
	}

	countStmt = countStmt.
		WHERE(condition)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	pag, limit := req.Pagination.GetResponseWithPageSize(count.TotalCount, ColleaguesDefaultPageSize)
	resp := &ListColleaguesResponse{
		Pagination: pag,
	}
	if count.TotalCount <= 0 {
		return resp, nil
	}

	// Convert proto sort to db sorting
	orderBys := []jet.OrderByClause{}
	if req.Sort != nil {
		var columns []jet.Column
		switch req.Sort.Column {
		case "name":
			columns = append(columns, tUser.Firstname, tUser.Lastname)
		case "rank":
			fallthrough
		default:
			columns = append(columns, tUser.JobGrade)
		}

		for _, column := range columns {
			if req.Sort.Direction == database.AscSortDirection {
				orderBys = append(orderBys, column.ASC())
			} else {
				orderBys = append(orderBys, column.DESC())
			}
		}
	} else {
		orderBys = append(orderBys,
			tUser.JobGrade.ASC(),
			tUser.Firstname.ASC(),
			tUser.Lastname.ASC(),
		)
	}

	stmt := tUser.
		SELECT(
			tUser.ID,
			tUser.Job,
			tUser.JobGrade,
			tUser.Firstname,
			tUser.Lastname,
			tUser.Dateofbirth,
			tUser.PhoneNumber,
			tUserProps.Avatar.AS("colleague.avatar"),
			tUserProps.Email.AS("colleague.email"),
			tJobsUserProps.UserID,
			tJobsUserProps.Job,
			tJobsUserProps.AbsenceBegin,
			tJobsUserProps.AbsenceEnd,
			tJobsUserProps.NamePrefix,
			tJobsUserProps.NameSuffix,
		).
		OPTIMIZER_HINTS(jet.OptimizerHint("idx_users_firstname_lastname_fulltext"))

	if (slices.Contains(types, "Labels") || userInfo.SuperUser) && len(req.LabelIds) > 0 {
		stmt = stmt.
			FROM(
				tUser.
					LEFT_JOIN(tUserProps,
						tUserProps.UserID.EQ(tUser.ID),
					).
					LEFT_JOIN(tJobsUserProps,
						tJobsUserProps.UserID.EQ(tUser.ID).
							AND(tJobsUserProps.Job.EQ(jet.String(userInfo.Job))),
					).
					INNER_JOIN(tUserLabels,
						tUserLabels.UserID.EQ(tUser.ID).
							AND(tUserLabels.Job.EQ(jet.String(userInfo.Job))),
					).
					LEFT_JOIN(tJobLabels,
						tJobLabels.ID.EQ(tUserLabels.LabelID),
					),
			)
	} else {
		stmt = stmt.
			FROM(
				tUser.
					LEFT_JOIN(tUserProps,
						tUserProps.UserID.EQ(tUser.ID),
					).
					LEFT_JOIN(tJobsUserProps,
						tJobsUserProps.UserID.EQ(tUser.ID).
							AND(tJobsUserProps.Job.EQ(jet.String(userInfo.Job))),
					),
			)
	}

	stmt = stmt.
		WHERE(condition).
		OFFSET(req.Pagination.Offset).
		GROUP_BY(tUser.ID).
		ORDER_BY(orderBys...).
		LIMIT(limit)

	if err := stmt.QueryContext(ctx, s.db, &resp.Colleagues); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	if slices.Contains(types, "Labels") || userInfo.SuperUser {
		userIds := []jet.Expression{}
		for _, colleague := range resp.Colleagues {
			userIds = append(userIds, jet.Int32(colleague.UserId))
		}

		labelsStmt := tUserLabels.
			SELECT(
				tUserLabels.UserID.AS("user_id"),
				tJobLabels.ID,
				tJobLabels.Job,
				tJobLabels.Name,
				tJobLabels.Color,
			).
			FROM(
				tUserLabels.
					LEFT_JOIN(tJobLabels,
						tJobLabels.ID.EQ(tUserLabels.LabelID),
					),
			).
			WHERE(jet.AND(
				tJobLabels.Job.EQ(jet.String(userInfo.Job)),
				tUserLabels.UserID.IN(userIds...),
			)).
			ORDER_BY(
				tJobLabels.Order.ASC(),
			)

		labels := []*struct {
			UserId int32 `sql:"primary_key" alias:"userId"`
			Labels *jobs.Labels
		}{}
		if err := labelsStmt.QueryContext(ctx, s.db, &labels); err != nil {
			if !errors.Is(err, qrm.ErrNoRows) {
				return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
			}
		}

		for _, props := range labels {
			idx := slices.IndexFunc(resp.Colleagues, func(c *jobs.Colleague) bool {
				return c.UserId == props.UserId
			})
			if idx == -1 {
				continue
			}

			colleague := resp.Colleagues[idx]
			if colleague.Props == nil {
				colleague.Props = &jobs.JobsUserProps{
					UserId: colleague.UserId,
					Job:    userInfo.Job,
				}
			}

			colleague.Props.Labels = props.Labels
		}
	}

	resp.Pagination.Update(len(resp.Colleagues))

	for i := 0; i < len(resp.Colleagues); i++ {
		s.enricher.EnrichJobInfo(resp.Colleagues[i])
	}

	return resp, nil
}

func (s *Server) getColleague(ctx context.Context, userInfo *userinfo.UserInfo, job string, userId int32, withColumns []jet.Projection) (*jobs.Colleague, error) {
	tUser := tUser.AS("colleague")
	columns := []jet.Projection{
		tUser.Firstname,
		tUser.Lastname,
		tUser.Job,
		tUser.JobGrade,
		tUser.Dateofbirth,
		tUser.PhoneNumber,
		tUserProps.Avatar.AS("colleague.avatar"),
		tUserProps.Email.AS("colleague.email"),
		tJobsUserProps.UserID,
		tJobsUserProps.Job,
		tJobsUserProps.AbsenceBegin,
		tJobsUserProps.AbsenceEnd,
		tJobsUserProps.NamePrefix,
		tJobsUserProps.NameSuffix,
	}
	columns = append(columns, withColumns...)

	stmt := tUser.
		SELECT(
			tUser.ID,
			columns...,
		).
		FROM(
			tUser.
				LEFT_JOIN(tUserProps,
					tUserProps.UserID.EQ(tUser.ID),
				).
				LEFT_JOIN(tJobsUserProps,
					tJobsUserProps.UserID.EQ(tUser.ID).
						AND(tJobsUserProps.Job.EQ(jet.String(job))),
				),
		).
		WHERE(
			tUser.ID.EQ(jet.Int32(userId)),
		).
		LIMIT(1)

	dest := &jobs.Colleague{}
	if err := stmt.QueryContext(ctx, s.db, dest); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
	}

	if dest.UserId == 0 {
		return nil, nil
	}

	if dest.Props == nil {
		dest.Props = &jobs.JobsUserProps{
			UserId: dest.UserId,
			Job:    userInfo.Job,
		}
	}

	labels, err := s.getUserLabels(ctx, userInfo, userId)
	if err != nil {
		return nil, err
	}
	dest.Props.Labels = labels

	s.enricher.EnrichJobInfo(dest)

	return dest, nil
}

func (s *Server) GetColleague(ctx context.Context, req *GetColleagueRequest) (*GetColleagueResponse, error) {
	trace.SpanFromContext(ctx).SetAttributes(attribute.Int64("fivenet.jobs.colleague.id", int64(req.UserId)))

	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: JobsService_ServiceDesc.ServiceName,
		Method:  "GetColleague",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	// Access Permission Check
	accessAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceGetColleaguePerm, permsjobs.JobsServiceGetColleagueAccessPermField)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	var colleagueAccess perms.StringList
	if accessAttr != nil {
		colleagueAccess = accessAttr.([]string)
	}

	targetUser, err := s.getColleague(ctx, userInfo, userInfo.Job, req.UserId, nil)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	if targetUser == nil {
		return nil, errorsjobs.ErrNotFoundOrNoPerms
	}

	infoOnly := req.InfoOnly != nil && *req.InfoOnly

	check := access.CheckIfHasAccess(colleagueAccess, userInfo, targetUser.Job, &users.UserShort{
		UserId:   targetUser.UserId,
		Job:      targetUser.Job,
		JobGrade: targetUser.JobGrade,
	})
	if !check {
		if userInfo.Job != targetUser.Job || !infoOnly {
			return nil, errorsjobs.ErrFailedQuery
		}
	}

	// Field Permission Check
	var types perms.StringList
	if !infoOnly {
		typesAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceGetColleaguePerm, permsjobs.JobsServiceGetColleagueTypesPermField)
		if err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
		if typesAttr != nil {
			types = typesAttr.([]string)
		}
	}
	if userInfo.SuperUser {
		types = []string{"Note"}
	}

	withColumns := []jet.Projection{}
	for _, fType := range types {
		switch fType {
		case "Note":
			withColumns = append(withColumns, tJobsUserProps.Note)
		}
	}

	colleague, err := s.getColleague(ctx, userInfo, userInfo.Job, targetUser.UserId, withColumns)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_VIEWED)

	return &GetColleagueResponse{
		Colleague: colleague,
	}, nil
}

func (s *Server) GetSelf(ctx context.Context, req *GetSelfRequest) (*GetSelfResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	// Field Permission Check
	typesAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceGetColleaguePerm, permsjobs.JobsServiceGetColleagueTypesPermField)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	var types perms.StringList
	if typesAttr != nil {
		types = typesAttr.([]string)
	}
	if userInfo.SuperUser {
		types = append(types, "AbsenceDate", "Note")
	}

	withColumns := []jet.Projection{}
	for _, fType := range types {
		switch fType {
		case "Note":
			withColumns = append(withColumns, tJobsUserProps.Note)
		}
	}

	colleague, err := s.getColleague(ctx, userInfo, userInfo.Job, userInfo.UserId, withColumns)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}

	return &GetSelfResponse{
		Colleague: colleague,
	}, nil
}

func (s *Server) getJobsUserProps(ctx context.Context, userInfo *userinfo.UserInfo, userId int32, job string, fields []string) (*jobs.JobsUserProps, error) {
	tJobsUserProps := table.FivenetJobsUserProps.AS("jobsuserprops")
	columns := []jet.Projection{
		tJobsUserProps.Job,
		tJobsUserProps.AbsenceBegin,
		tJobsUserProps.AbsenceEnd,
		tJobsUserProps.NamePrefix,
		tJobsUserProps.NameSuffix,
	}

	for _, field := range fields {
		switch field {
		case "Note":
			columns = append(columns, tJobsUserProps.Note)
		}
	}

	stmt := tJobsUserProps.
		SELECT(
			tJobsUserProps.UserID,
			columns...,
		).
		FROM(tJobsUserProps).
		WHERE(jet.AND(
			tJobsUserProps.UserID.EQ(jet.Int32(userId)),
			tJobsUserProps.Job.EQ(jet.String(job)),
		)).
		LIMIT(1)

	dest := &jobs.JobsUserProps{
		UserId: userId,
	}
	if err := stmt.QueryContext(ctx, s.db, dest); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	labels, err := s.getUserLabels(ctx, userInfo, userId)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	dest.Labels = labels

	return dest, nil
}

func (s *Server) SetJobsUserProps(ctx context.Context, req *SetJobsUserPropsRequest) (*SetJobsUserPropsResponse, error) {
	trace.SpanFromContext(ctx).SetAttributes(attribute.Int64("fivenet.jobs.colleague.id", int64(req.Props.UserId)))

	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: JobsService_ServiceDesc.ServiceName,
		Method:  "SetJobsUserProps",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	if req.Reason == "" {
		return nil, errorsjobs.ErrReasonRequired
	}

	// Access Permission Check
	accessAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceSetJobsUserPropsPerm, permsjobs.JobsServiceSetJobsUserPropsAccessPermField)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	var colleagueAccess perms.StringList
	if accessAttr != nil {
		colleagueAccess = accessAttr.([]string)
	}

	targetUser, err := s.getColleague(ctx, userInfo, userInfo.Job, req.Props.UserId, nil)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	if targetUser == nil {
		return nil, errorsjobs.ErrNotFoundOrNoPerms
	}

	if !access.CheckIfHasAccess(colleagueAccess, userInfo, targetUser.Job, &users.UserShort{
		UserId:   targetUser.UserId,
		Job:      targetUser.Job,
		JobGrade: targetUser.JobGrade,
	}) {
		return nil, errorsjobs.ErrFailedQuery
	}

	// Types Permission Check
	typesAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceSetJobsUserPropsPerm, permsjobs.JobsServiceSetJobsUserPropsTypesPermField)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	var types perms.StringList
	if typesAttr != nil {
		types = typesAttr.([]string)
	}
	if userInfo.SuperUser {
		types = []string{"AbsenceDate", "Note", "Labels", "Name"}
	}

	props, err := s.getJobsUserProps(ctx, userInfo, targetUser.UserId, targetUser.Job, types)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}

	absenceBegin := jet.DateExp(jet.NULL)
	absenceEnd := jet.DateExp(jet.NULL)
	if req.Props.AbsenceBegin != nil && req.Props.AbsenceEnd != nil {
		// Allow users to set their own absence date regardless of types perms check
		if userInfo.UserId != targetUser.UserId && !slices.Contains(types, "AbsenceDate") && !userInfo.SuperUser {
			return nil, errorsjobs.ErrPropsAbsenceDenied
		}

		if req.Props.AbsenceBegin.Timestamp == nil {
			req.Props.AbsenceBegin = nil
		} else {
			absenceBegin = jet.DateT(req.Props.AbsenceBegin.AsTime())
		}

		if req.Props.AbsenceEnd.Timestamp == nil {
			req.Props.AbsenceEnd = nil
		} else {
			absenceEnd = jet.DateT(req.Props.AbsenceEnd.AsTime())
		}
	} else {
		req.Props.AbsenceBegin = props.AbsenceBegin
		req.Props.AbsenceEnd = props.AbsenceEnd
	}

	tJobsUserProps := table.FivenetJobsUserProps
	updateSets := []jet.ColumnAssigment{
		tJobsUserProps.AbsenceBegin.SET(jet.DateExp(jet.Raw("VALUES(`absence_begin`)"))),
		tJobsUserProps.AbsenceEnd.SET(jet.DateExp(jet.Raw("VALUES(`absence_end`)"))),
	}

	// Generate the update sets
	if req.Props.Note != nil {
		if !slices.Contains(types, "Note") && !userInfo.SuperUser {
			return nil, errorsjobs.ErrPropsNoteDenied
		}

		updateSets = append(updateSets, tJobsUserProps.Note.SET(jet.String(*req.Props.Note)))
	} else {
		req.Props.Note = props.Note
	}

	if req.Props.Labels != nil {
		// Check if user is allowed to update labels
		if !slices.Contains(types, "Labels") && !userInfo.SuperUser {
			return nil, errorsjobs.ErrPropsAbsenceDenied
		}

		added, _ := utils.SlicesDifferenceFunc(props.Labels.List, req.Props.Labels.List,
			func(in *jobs.Label) uint64 {
				return in.Id
			})

		valid, err := s.validateLabels(ctx, userInfo, added)
		if err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
		if !valid {
			return nil, errorsjobs.ErrPropsLabelsDenied
		}
	} else {
		req.Props.Labels = props.Labels
	}

	if req.Props.NamePrefix != nil || req.Props.NameSuffix != nil {
		if !slices.Contains(types, "Name") && !userInfo.SuperUser {
			return nil, errorsjobs.ErrPropsNameDenied
		}

		if req.Props.NamePrefix != nil {
			updateSets = append(updateSets, tJobsUserProps.NamePrefix.SET(jet.String(*req.Props.NamePrefix)))
		} else {
			req.Props.NamePrefix = props.NamePrefix
		}
		if req.Props.NameSuffix != nil {
			updateSets = append(updateSets, tJobsUserProps.NameSuffix.SET(jet.String(*req.Props.NameSuffix)))
		} else {
			req.Props.NameSuffix = props.NameSuffix
		}
	} else {
		req.Props.NamePrefix = props.NamePrefix
		req.Props.NameSuffix = props.NameSuffix
	}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	stmt := tJobsUserProps.
		INSERT(
			tJobsUserProps.UserID,
			tJobsUserProps.Job,
			tJobsUserProps.AbsenceBegin,
			tJobsUserProps.AbsenceEnd,
			tJobsUserProps.Note,
			tJobsUserProps.NamePrefix,
			tJobsUserProps.NameSuffix,
		).
		VALUES(
			targetUser.UserId,
			targetUser.Job,
			absenceBegin,
			absenceEnd,
			req.Props.Note,
			req.Props.NamePrefix,
			req.Props.NameSuffix,
		).
		ON_DUPLICATE_KEY_UPDATE(
			updateSets...,
		)

	if _, err := stmt.ExecContext(ctx, tx); err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}

	if req.Props.Labels != nil && !proto.Equal(req.Props.Labels, props.Labels) {
		added, removed := utils.SlicesDifferenceFunc(props.Labels.List, req.Props.Labels.List,
			func(in *jobs.Label) uint64 {
				return in.Id
			})

		valid, err := s.validateLabels(ctx, userInfo, added)
		if err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
		if !valid {
			return nil, errorsjobs.ErrPropsLabelsDenied
		}

		if err := s.updateLabels(ctx, tx, targetUser.UserId, targetUser.Job, added, removed); err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}

		if err := s.addJobsUserActivity(ctx, tx, &jobs.JobsUserActivity{
			Job:          userInfo.Job,
			SourceUserId: userInfo.UserId,
			TargetUserId: targetUser.UserId,
			ActivityType: jobs.JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_LABELS,
			Reason:       req.Reason,
			Data: &jobs.JobsUserActivityData{
				Data: &jobs.JobsUserActivityData_LabelsChange{
					LabelsChange: &jobs.ColleagueLabelsChange{
						Added:   added,
						Removed: removed,
					},
				},
			},
		}); err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	// Compare absence dates if any were set
	if req.Props.AbsenceBegin != nil && (props.AbsenceBegin == nil || req.Props.AbsenceBegin.AsTime().Compare(props.AbsenceBegin.AsTime()) != 0) ||
		req.Props.AbsenceEnd != nil && (props.AbsenceEnd == nil || req.Props.AbsenceEnd.AsTime().Compare(props.AbsenceEnd.AsTime()) != 0) {
		if err := s.addJobsUserActivity(ctx, tx, &jobs.JobsUserActivity{
			Job:          userInfo.Job,
			SourceUserId: userInfo.UserId,
			TargetUserId: targetUser.UserId,
			ActivityType: jobs.JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_ABSENCE_DATE,
			Reason:       req.Reason,
			Data: &jobs.JobsUserActivityData{
				Data: &jobs.JobsUserActivityData_AbsenceDate{
					AbsenceDate: &jobs.ColleagueAbsenceDate{
						AbsenceBegin: req.Props.AbsenceBegin,
						AbsenceEnd:   req.Props.AbsenceEnd,
					},
				},
			},
		}); err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	if (req.Props.Note == nil && props.Note != nil) || (req.Props.Note != nil && props.Note == nil) {
		if err := s.addJobsUserActivity(ctx, tx, &jobs.JobsUserActivity{
			Job:          userInfo.Job,
			SourceUserId: userInfo.UserId,
			TargetUserId: targetUser.UserId,
			ActivityType: jobs.JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_NOTE,
			Reason:       req.Reason,
		}); err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	if req.Props.NamePrefix != nil && (props.NamePrefix == nil || req.Props.NamePrefix != props.NamePrefix) ||
		req.Props.NameSuffix != nil && (props.NameSuffix == nil || req.Props.NameSuffix != props.NameSuffix) {
		if err := s.addJobsUserActivity(ctx, tx, &jobs.JobsUserActivity{
			Job:          userInfo.Job,
			SourceUserId: userInfo.UserId,
			TargetUserId: targetUser.UserId,
			ActivityType: jobs.JobsUserActivityType_JOBS_USER_ACTIVITY_TYPE_NAME,
			Reason:       req.Reason,
			Data: &jobs.JobsUserActivityData{
				Data: &jobs.JobsUserActivityData_NameChange{
					NameChange: &jobs.ColleagueNameChange{
						Prefix: req.Props.NamePrefix,
						Suffix: req.Props.NameSuffix,
					},
				},
			},
		}); err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	props, err = s.getJobsUserProps(ctx, userInfo, targetUser.UserId, targetUser.Job, types)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}

	return &SetJobsUserPropsResponse{
		Props: props,
	}, nil
}

func (s *Server) addJobsUserActivity(ctx context.Context, tx qrm.DB, activity *jobs.JobsUserActivity) error {
	stmt := tJobsUserActivity.
		INSERT(
			tJobsUserActivity.Job,
			tJobsUserActivity.SourceUserID,
			tJobsUserActivity.TargetUserID,
			tJobsUserActivity.ActivityType,
			tJobsUserActivity.Reason,
			tJobsUserActivity.Data,
		).
		VALUES(
			activity.Job,
			activity.SourceUserId,
			activity.TargetUserId,
			activity.ActivityType,
			activity.Reason,
			activity.Data,
		)

	if _, err := stmt.ExecContext(ctx, tx); err != nil {
		return err
	}

	return nil
}

func (s *Server) getConditionForColleagueAccess(actTable *table.FivenetJobsUserActivityTable, usersTable *table.UsersTable, levels []string, userInfo *userinfo.UserInfo) jet.BoolExpression {
	condition := jet.Bool(true)
	if userInfo.SuperUser {
		return condition
	}

	// If no levels set, assume "Own" as a safe default
	if len(levels) == 0 {
		return actTable.TargetUserID.EQ(jet.Int32(userInfo.UserId))
	}

	if slices.Contains(levels, "Any") {
		return condition
	}
	if slices.Contains(levels, "Lower_Rank") {
		return usersTable.ID.LT(jet.Int32(userInfo.JobGrade))
	}
	if slices.Contains(levels, "Same_Rank") {
		return usersTable.ID.LT_EQ(jet.Int32(userInfo.JobGrade))
	}
	if slices.Contains(levels, "Own") {
		return usersTable.ID.EQ(jet.Int32(userInfo.UserId))
	}

	return jet.Bool(false)
}

func (s *Server) ListColleagueActivity(ctx context.Context, req *ListColleagueActivityRequest) (*ListColleagueActivityResponse, error) {
	trace.SpanFromContext(ctx).SetAttributes(attribute.IntSlice("fivenet.jobs.colleagues.user_ids", utils.SliceInt32ToInt(req.UserIds)))

	userInfo := auth.MustGetUserInfoFromContext(ctx)

	// Access Field Permission Check
	accessAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceGetColleaguePerm, permsjobs.JobsServiceGetColleagueAccessPermField)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	var colleagueAccess perms.StringList
	if accessAttr != nil {
		colleagueAccess = accessAttr.([]string)
	}

	tJobsUserActivity := tJobsUserActivity.AS("jobsuseractivity")
	tTargetUser := tUser.AS("target_user")
	tSourceUser := tUser.AS("source_user")

	condition := tJobsUserActivity.Job.EQ(jet.String(userInfo.Job))

	resp := &ListColleagueActivityResponse{
		Pagination: &database.PaginationResponse{
			PageSize: ColleaguesDefaultPageSize,
		},
	}

	if len(req.ActivityTypes) == 0 {
		return resp, nil
	}

	// If no user IDs given or more than 2, show all the user has access to
	if len(req.UserIds) == 0 || len(req.UserIds) >= 2 {
		condition = condition.AND(s.getConditionForColleagueAccess(tJobsUserActivity, tTargetUser, colleagueAccess, userInfo))

		if len(req.UserIds) >= 2 {
			// More than 2 user ids
			userIds := make([]jet.Expression, len(req.UserIds))
			for i := 0; i < len(req.UserIds); i++ {
				userIds[i] = jet.Int32(req.UserIds[i])
			}

			condition = condition.AND(tTargetUser.ID.IN(userIds...))
		}
	} else {
		userId := req.UserIds[0]

		targetUser, err := s.getColleague(ctx, userInfo, userInfo.Job, userId, nil)
		if err != nil {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
		if targetUser == nil {
			return nil, errorsjobs.ErrNotFoundOrNoPerms
		}

		if !access.CheckIfHasAccess(colleagueAccess, userInfo, targetUser.Job, &users.UserShort{
			UserId:   targetUser.UserId,
			Job:      targetUser.Job,
			JobGrade: targetUser.JobGrade,
		}) {
			return nil, errorsjobs.ErrFailedQuery
		}

		condition = condition.AND(tJobsUserActivity.TargetUserID.EQ(jet.Int32(userId)))
	}

	// Types Field Permission Check
	typesAttr, err := s.ps.Attr(userInfo, permsjobs.JobsServicePerm, permsjobs.JobsServiceListColleagueActivityPerm, permsjobs.JobsServiceListColleagueActivityTypesPermField)
	if err != nil {
		return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
	}
	var types perms.StringList
	if typesAttr != nil {
		types = typesAttr.([]string)
	}
	if len(types) == 0 {
		if !userInfo.SuperUser {
			return resp, nil
		} else {
			types = append(types, "HIRED", "FIRED", "PROMOTED", "DEMOTED", "ABSENCE_DATE", "NOTE", "LABELS", "NAME")
		}
	}

	if len(req.ActivityTypes) > 0 {
		req.ActivityTypes = slices.DeleteFunc(req.ActivityTypes, func(t jobs.JobsUserActivityType) bool {
			return !slices.ContainsFunc(types, func(s string) bool {
				return strings.Contains(t.String(), "JOBS_USER_ACTIVITY_TYPE_"+s)
			})
		})
	}

	condTypes := []jet.Expression{}
	for _, aType := range req.ActivityTypes {
		condTypes = append(condTypes, jet.Int16(int16(aType)))
	}

	if len(condTypes) == 0 {
		return resp, nil
	}

	condition = condition.AND(tJobsUserActivity.ActivityType.IN(condTypes...))

	// Get total count of values
	countStmt := tJobsUserActivity.
		SELECT(
			jet.COUNT(tJobsUserActivity.ID).AS("datacount.totalcount"),
		).
		FROM(
			tJobsUserActivity.
				INNER_JOIN(tTargetUser,
					tTargetUser.ID.EQ(tJobsUserActivity.TargetUserID),
				),
		).
		WHERE(condition)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorsjobs.ErrFailedQuery)
		}
	}

	pag, limit := req.Pagination.GetResponseWithPageSize(count.TotalCount, ColleaguesDefaultPageSize)
	resp.Pagination = pag
	if count.TotalCount <= 0 {
		return resp, nil
	}

	// Convert proto sort to db sorting
	orderBys := []jet.OrderByClause{}
	if req.Sort != nil {
		var column jet.Column
		switch req.Sort.Column {
		case "createdAt":
			fallthrough
		default:
			column = tJobsUserActivity.CreatedAt
		}

		if req.Sort.Direction == database.AscSortDirection {
			orderBys = append(orderBys, column.ASC())
		} else {
			orderBys = append(orderBys, column.DESC())
		}
	} else {
		orderBys = append(orderBys, tJobsUserActivity.CreatedAt.DESC())
	}

	tTargetUserProps := tUserProps.AS("target_user_props")
	tTargetJobsUserProps := tJobsUserProps.AS("fivenet_jobs_user_props")
	tSourceUserProps := tUserProps.AS("source_user_props")

	stmt := tJobsUserActivity.
		SELECT(
			tJobsUserActivity.ID,
			tJobsUserActivity.CreatedAt,
			tJobsUserActivity.Job,
			tJobsUserActivity.SourceUserID,
			tJobsUserActivity.TargetUserID,
			tJobsUserActivity.ActivityType,
			tJobsUserActivity.Reason,
			tJobsUserActivity.Data,
			tTargetUser.ID,
			tTargetUser.Job,
			tTargetUser.JobGrade,
			tTargetUser.Firstname,
			tTargetUser.Lastname,
			tTargetUser.Dateofbirth,
			tTargetUser.PhoneNumber,
			tTargetUserProps.Avatar.AS("target_user.avatar"),
			tTargetJobsUserProps.UserID,
			tTargetJobsUserProps.Job,
			tTargetJobsUserProps.AbsenceBegin,
			tTargetJobsUserProps.AbsenceEnd,
			tTargetJobsUserProps.NamePrefix,
			tTargetJobsUserProps.NameSuffix,
			tSourceUser.ID,
			tSourceUser.Job,
			tSourceUser.JobGrade,
			tSourceUser.Firstname,
			tSourceUser.Lastname,
			tSourceUser.Dateofbirth,
			tSourceUser.PhoneNumber,
			tSourceUserProps.Avatar.AS("source_user.avatar"),
		).
		FROM(
			tJobsUserActivity.
				INNER_JOIN(tTargetUser,
					tTargetUser.ID.EQ(tJobsUserActivity.TargetUserID),
				).
				LEFT_JOIN(tTargetUserProps,
					tTargetUserProps.UserID.EQ(tTargetUser.ID),
				).
				LEFT_JOIN(tTargetJobsUserProps,
					tTargetJobsUserProps.UserID.EQ(tTargetUser.ID).
						AND(tTargetUser.Job.EQ(jet.String(userInfo.Job))),
				).
				LEFT_JOIN(tSourceUser,
					tSourceUser.ID.EQ(tJobsUserActivity.SourceUserID),
				).
				LEFT_JOIN(tSourceUserProps,
					tSourceUserProps.UserID.EQ(tSourceUser.ID),
				),
		).
		WHERE(condition).
		OFFSET(pag.Offset).
		ORDER_BY(orderBys...).
		LIMIT(limit)

	if err := stmt.QueryContext(ctx, s.db, &resp.Activity); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
	}

	pag.Update(len(resp.Activity))

	jobInfoFn := s.enricher.EnrichJobInfoSafeFunc(userInfo)
	for i := 0; i < len(resp.Activity); i++ {
		if resp.Activity[i].SourceUser != nil {
			jobInfoFn(resp.Activity[i].SourceUser)
		}

		if resp.Activity[i].TargetUser != nil {
			jobInfoFn(resp.Activity[i].TargetUser)
		}
	}

	return resp, nil
}
