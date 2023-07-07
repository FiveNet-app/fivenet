package centrum

import (
	"context"

	database "github.com/galexrt/fivenet/gen/go/proto/resources/common/database"
	"github.com/galexrt/fivenet/gen/go/proto/resources/dispatch"
	"github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/utils/dbutils"
	"github.com/galexrt/fivenet/pkg/utils/syncx"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
)

var (
	tUnits      = table.FivenetCentrumUnits.AS("unit")
	tUnitStatus = table.FivenetCentrumUnitsStatus.AS("unitstatus")
	tUnitUser   = table.FivenetCentrumUnitsUsers.AS("unitassignment")
	tUser       = table.Users.AS("usershort")
)

func (s *Server) loadUnits(ctx context.Context) error {
	condition := tUnitStatus.ID.IS_NULL().OR(
		tUnitStatus.ID.EQ(
			jet.RawInt(`SELECT MAX(unitstatus.id) FROM fivenet_centrum_units_status AS unitstatus WHERE unitstatus.unit_id = unit.id`),
		),
	)

	stmt := tUnits.
		SELECT(
			tUnits.ID,
			tUnits.Job,
			tUnits.Name,
			tUnits.Initials,
			tUnits.Color,
			tUnits.Description,
			tUnitStatus.ID,
			tUnitStatus.CreatedAt,
			tUnitStatus.UnitID,
			tUnitStatus.Status,
			tUnitStatus.Reason,
			tUnitStatus.Code,
			tUnitStatus.UserID,
			tUnitStatus.InSquad,
			tUnitStatus.X,
			tUnitStatus.Y,
			tUnitUser.UnitID,
			tUnitUser.UserID,
			tUnitUser.Identifier,
		).
		FROM(
			tUnits.
				LEFT_JOIN(tUnitStatus,
					tUnitStatus.UnitID.EQ(tUnits.ID),
				).
				LEFT_JOIN(tUnitUser,
					tUnitUser.UnitID.EQ(tUnits.ID),
				),
		).
		WHERE(condition).
		ORDER_BY(
			tUnits.Job.ASC(),
			tUnits.Name.ASC(),
			tUnitStatus.Status.ASC(),
		)

	units := []*dispatch.Unit{}
	if err := stmt.QueryContext(ctx, s.db, &units); err != nil {
		return err
	}

	for i := 0; i < len(units); i++ {
		var err error
		units[i].Users, err = s.resolveUsersForUnit(ctx, units[i].Users)
		if err != nil {
			return err
		}

		jobUnits, _ := s.units.LoadOrStore(units[i].Job, &syncx.Map[uint64, *dispatch.Unit]{})
		jobUnits.Store(units[i].Id, units[i])
	}

	return nil
}

func (s *Server) ListUnits(ctx context.Context, req *ListUnitsRequest) (*ListUnitsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "ListUnits",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	resp := &ListUnitsResponse{
		Units: []*dispatch.Unit{},
	}

	val, ok := s.units.Load(userInfo.Job)
	if val == nil || !ok {
		return resp, nil
	}

	units := []*dispatch.Unit{}
	val.Range(func(key uint64, value *dispatch.Unit) bool {
		units = append(units, value)
		return true
	})

	resp.Units = units

	auditEntry.State = int16(rector.EVENT_TYPE_VIEWED)

	return resp, nil
}

func (s *Server) CreateOrUpdateUnit(ctx context.Context, req *CreateOrUpdateUnitRequest) (*CreateOrUpdateUnitResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "CreateOrUpdateUnit",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	resp := &CreateOrUpdateUnitResponse{}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, ErrFailedQuery
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	// No unit id set
	if req.Unit.Id <= 0 {
		tUnits := table.FivenetCentrumUnits
		stmt := tUnits.
			INSERT(
				tUnits.Job,
				tUnits.Name,
				tUnits.Initials,
				tUnits.Color,
				tUnits.Description,
			).
			VALUES(
				userInfo.Job,
				req.Unit.Name,
				req.Unit.Initials,
				req.Unit.Color,
				req.Unit.Description,
			)

		result, err := stmt.ExecContext(ctx, tx)
		if err != nil {
			return nil, err
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		unit, err := s.getUnit(ctx, userInfo, uint64(lastId))
		if err != nil {
			return nil, err
		}

		resp.Unit = unit

		stmt = tUnitStatus.
			INSERT(
				tUnitStatus.UnitID,
				tUnitStatus.Status,
				tUnitStatus.UserID,
			).
			VALUES(
				resp.Unit.Id,
				dispatch.UNIT_STATUS_UNKNOWN,
				userInfo.UserId,
			)

		if _, err := stmt.ExecContext(ctx, tx); err != nil {
			return nil, err
		}

		auditEntry.State = int16(rector.EVENT_TYPE_CREATED)
	} else {
		stmt := tUnits.
			UPDATE(
				tUnits.Job,
				tUnits.Name,
				tUnits.Initials,
				tUnits.Color,
				tUnits.Description,
			).
			SET(
				userInfo.Job,
				req.Unit.Name,
				req.Unit.Initials,
				req.Unit.Color,
				req.Unit.Description,
			).
			WHERE(jet.AND(
				tUnits.Job.EQ(jet.String(userInfo.Job)),
				tUnits.Job.EQ(jet.String(userInfo.Job)),
			))

		if _, err := stmt.ExecContext(ctx, tx); err != nil {
			return nil, err
		}

		unit, err := s.getUnit(ctx, userInfo, req.Unit.Id)
		if err != nil {
			return nil, err
		}

		resp.Unit = unit

		auditEntry.State = int16(rector.EVENT_TYPE_UPDATED)
	}

	return resp, nil
}

func (s *Server) DeleteUnit(ctx context.Context, req *DeleteUnitRequest) (*DeleteUnitResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "DeleteUnit",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EVENT_TYPE_DELETED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	stmt := tUnits.
		DELETE().
		WHERE(jet.AND(
			tUnits.Job.EQ(jet.String(userInfo.Job)),
			tUnits.ID.EQ(jet.Uint64(req.UnitId)),
		)).
		LIMIT(1)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	units, ok := s.units.Load(userInfo.Job)
	if ok {
		units.Delete(req.UnitId)
	}

	auditEntry.State = int16(rector.EVENT_TYPE_DELETED)

	return &DeleteUnitResponse{}, nil
}

func (s *Server) UpdateUnitStatus(ctx context.Context, req *UpdateUnitStatusRequest) (*UpdateUnitStatusResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	unit, err := s.getUnit(ctx, userInfo, req.UnitId)
	if err != nil {
		return nil, err
	}

	found := false
	for i := 0; i < len(unit.Users); i++ {
		if unit.Users[i].UserId == userInfo.UserId {
			found = true
			break
		}
	}
	if !found {
		return nil, ErrFailedQuery
	}

	stmt := tUnitStatus.
		INSERT(
			tUnitStatus.UnitID,
			tUnitStatus.Status,
			tUnitStatus.Reason,
			tUnitStatus.Code,
			tUnitStatus.UserID,
			tUnitStatus.InSquad,
		).
		VALUES(
			req.UnitId,
			req.Status,
			req.Reason,
			req.Code,
			userInfo.UserId,
			false,
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	return &UpdateUnitStatusResponse{}, nil
}

func (s *Server) AssignUnit(ctx context.Context, req *AssignUnitRequest) (*AssignUnitResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	unit, err := s.getUnit(ctx, userInfo, req.UnitId)
	if err != nil {
		return nil, err
	}
	if unit.Job != userInfo.Job {
		return nil, nil
	}

	addIds := make([]jet.Expression, len(req.ToAdd))
	for i := 0; i < len(req.ToAdd); i++ {
		addIds[i] = jet.Int32(req.ToAdd[i])
	}
	removeIds := make([]jet.Expression, len(req.ToRemove))
	for i := 0; i < len(req.ToRemove); i++ {
		removeIds[i] = jet.Int32(req.ToRemove[i])
	}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, ErrFailedQuery
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	if len(removeIds) > 0 {
		stmt := tUnitUser.
			DELETE().
			WHERE(jet.AND(
				tUnitUser.UnitID.EQ(jet.Uint64(unit.Id)),
				tUnitUser.UserID.IN(removeIds...),
			))

		if _, err := stmt.ExecContext(ctx, tx); err != nil {
			return nil, err
		}

		for i := 0; i < len(unit.Users); i++ {
			for k := 0; k < len(req.ToRemove); k++ {
				if unit.Users[i].UserId == req.ToRemove[k] {
					break
				}
			}
		}
	}

	if len(addIds) > 0 {
		for _, id := range addIds {
			stmt := tUnitUser.
				INSERT(
					tUnitUser.UnitID,
					tUnitUser.UserID,
				).
				VALUES(
					unit.Id,
					id,
				)

			if _, err := stmt.ExecContext(ctx, tx); err != nil {
				if !dbutils.IsDuplicateError(err) {
					return nil, err
				}
			}

		}

		found := []int32{}
		for k := 0; k < len(req.ToAdd); k++ {
			for i := 0; i < len(unit.Users); i++ {
				if unit.Users[i].UserId == req.ToAdd[k] {
					found = append(found, req.ToAdd[k])
				}
			}
		}

		users, err := s.resolveUsersById(ctx, found)
		if err != nil {
			return nil, err
		}
		assignments := []*dispatch.UnitAssignment{}
		for _, v := range users {
			assignments = append(assignments, &dispatch.UnitAssignment{
				UnitId: unit.Id,
				UserId: v.UserId,
				User:   v,
			})
		}
		unit.Users = assignments
	}

	return &AssignUnitResponse{}, nil
}

func (s *Server) ListUnitActivity(ctx context.Context, req *ListActivityRequest) (*ListUnitActivityResponse, error) {
	countStmt := tUnitStatus.
		SELECT(
			jet.COUNT(jet.DISTINCT(tUnitStatus.ID)).AS("datacount.totalcount"),
		).
		FROM(tUnitStatus).
		WHERE(
			tUnitStatus.UnitID.EQ(jet.Uint64(req.Id)),
		)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		return nil, ErrFailedQuery
	}

	pag, limit := req.Pagination.GetResponseWithPageSize(10)
	resp := &ListUnitActivityResponse{
		Pagination: pag,
	}
	if count.TotalCount <= 0 {
		return resp, nil
	}

	stmt := tUnitStatus.
		SELECT(
			tUnitStatus.ID,
			tUnitStatus.CreatedAt,
			tUnitStatus.UnitID,
			tUnitStatus.Status,
			tUnitStatus.Reason,
			tUnitStatus.Code,
			tUnitStatus.UserID,
			tUnitStatus.InSquad,
			tUnitStatus.X,
			tUnitStatus.Y,
			tUser.Firstname,
			tUser.Lastname,
			tUser.Dateofbirth,
			tUser.Job,
		).
		FROM(
			tUnitStatus.
				LEFT_JOIN(tUser,
					tUser.ID.EQ(tUnitStatus.UserID),
				),
		).
		WHERE(
			tUnitStatus.UnitID.EQ(jet.Uint64(req.Id)),
		).
		ORDER_BY(tUnitStatus.ID.DESC()).
		OFFSET(req.Pagination.Offset).
		LIMIT(limit)

	if err := stmt.QueryContext(ctx, s.db, &resp.Activity); err != nil {
		return nil, err
	}

	resp.Pagination.Update(count.TotalCount, len(resp.Activity))

	return resp, nil
}
