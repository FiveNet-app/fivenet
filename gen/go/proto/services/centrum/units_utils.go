package centrum

import (
	"context"
	"strings"

	dispatch "github.com/galexrt/fivenet/gen/go/proto/resources/dispatch"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/pkg/utils/dbutils"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
)

func (s *Server) getUnit(job string, id uint64) (*dispatch.Unit, bool) {
	units, ok := s.state.Units.Load(job)
	if !ok {
		return nil, false
	}

	return units.Load(id)
}

func (s *Server) listUnits(job string) ([]*dispatch.Unit, bool) {
	us := []*dispatch.Unit{}

	units, ok := s.state.Units.Load(job)
	if !ok {
		return nil, false
	}

	units.Range(func(key uint64, unit *dispatch.Unit) bool {
		us = append(us, unit)
		return true
	})

	slices.SortFunc(us, func(a, b *dispatch.Unit) int {
		return strings.Compare(a.Name, b.Name)
	})

	return us, true
}

func (s *Server) getUnitStatusFromDB(ctx context.Context, job string, id uint64) (*dispatch.UnitStatus, error) {
	stmt := tUnitStatus.
		SELECT(
			tUnitStatus.ID,
			tUnitStatus.CreatedAt,
			tUnitStatus.UnitID,
			tUnitStatus.Status,
			tUnitStatus.Reason,
			tUnitStatus.Code,
			tUnitStatus.UserID,
			tUnitStatus.X,
			tUnitStatus.Y,
			tUnitStatus.Postal,
			tUnitStatus.CreatorID,
		).
		FROM(tUnitStatus).
		WHERE(
			tUnitStatus.ID.EQ(jet.Uint64(id)),
		).
		LIMIT(1)

	var dest dispatch.UnitStatus
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		return nil, err
	}

	if dest.UserId != nil {
		var err error
		dest.User, err = s.resolveUserShortById(ctx, *dest.UserId)
		if err != nil {
			return nil, err
		}
	}
	if dest.CreatorId != nil {
		var err error
		dest.Creator, err = s.resolveUserShortById(ctx, *dest.CreatorId)
		if err != nil {
			return nil, err
		}
	}

	if dest.UnitId > 0 {
		unit, ok := s.getUnit(job, dest.UnitId)
		if ok && unit != nil {
			newUnit := proto.Clone(unit)
			dest.Unit = newUnit.(*dispatch.Unit)
		}
	}

	return &dest, nil
}

func (s *Server) resolveUsersForUnit(ctx context.Context, u []*dispatch.UnitAssignment) ([]*dispatch.UnitAssignment, error) {
	userIds := make([]int32, len(u))
	for i := 0; i < len(u); i++ {
		userIds[i] = u[i].UserId
	}

	if len(userIds) == 0 {
		return nil, nil
	}

	us, err := s.resolveUserShortsByIds(ctx, userIds)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(u); i++ {
		u[i].User = us[i]
	}

	return u, nil
}

func (s *Server) getUnitIDForUserID(userId int32) (uint64, bool) {
	return s.state.UserIDToUnitID.Load(userId)
}

func (s *Server) updateUnitStatus(ctx context.Context, job string, unit *dispatch.Unit, in *dispatch.UnitStatus) error {
	// If the unit status is the same and is a status that shouldn't be duplicated, don't update the status again
	if unit.Status != nil &&
		unit.Status.Status == in.Status &&
		(in.Status == dispatch.StatusUnit_STATUS_UNIT_ON_BREAK ||
			in.Status == dispatch.StatusUnit_STATUS_UNIT_BUSY ||
			in.Status == dispatch.StatusUnit_STATUS_UNIT_UNAVAILABLE ||
			in.Status == dispatch.StatusUnit_STATUS_UNIT_AVAILABLE) {
		return nil
	}

	tUnitStatus := table.FivenetCentrumUnitsStatus
	stmt := tUnitStatus.
		INSERT(
			tUnitStatus.UnitID,
			tUnitStatus.Status,
			tUnitStatus.Reason,
			tUnitStatus.Code,
			tUnitStatus.UserID,
			tUnitStatus.X,
			tUnitStatus.Y,
			tUnitStatus.Postal,
			tUnitStatus.CreatorID,
		).
		VALUES(
			in.UnitId,
			in.Status,
			in.Reason,
			in.Code,
			in.UserId,
			in.X,
			in.Y,
			in.Postal,
			in.CreatorId,
		)

	res, err := stmt.ExecContext(ctx, s.db)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	status, err := s.getUnitStatusFromDB(ctx, job, uint64(lastId))
	if err != nil {
		return err
	}
	unit.Status = status

	data, err := proto.Marshal(unit)
	if err != nil {
		return err
	}
	s.events.JS.PublishAsync(buildSubject(TopicUnit, TypeUnitStatus, job, status.UnitId), data)

	return nil
}

func (s *Server) updateUnitAssignments(ctx context.Context, userInfo *userinfo.UserInfo, unit *dispatch.Unit, toAdd []int32, toRemove []int32) error {
	var x, y *float64
	var postal *string
	marker, ok := s.tracker.GetUserById(userInfo.UserId)
	if ok {
		x = &marker.Info.X
		y = &marker.Info.Y
		postal = marker.Info.Postal
	}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return ErrFailedQuery
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	tUnitUser := table.FivenetCentrumUnitsUsers
	if len(toRemove) > 0 {
		removeIds := make([]jet.Expression, len(toRemove))
		for i := 0; i < len(toRemove); i++ {
			removeIds[i] = jet.Int32(toRemove[i])
		}

		stmt := tUnitUser.
			DELETE().
			WHERE(jet.AND(
				tUnitUser.UnitID.EQ(jet.Uint64(unit.Id)),
				tUnitUser.UserID.IN(removeIds...),
			))

		if _, err := stmt.ExecContext(ctx, tx); err != nil {
			return err
		}

		for i := 0; i < len(unit.Users); i++ {
			for k := len(toRemove) - 1; k >= 0; k-- {
				if unit.Users[i].UserId == toRemove[k] {
					if err := s.updateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
						UnitId:    unit.Id,
						Status:    dispatch.StatusUnit_STATUS_UNIT_USER_REMOVED,
						UserId:    &toRemove[k],
						CreatorId: &userInfo.UserId,
						X:         x,
						Y:         y,
						Postal:    postal,
					}); err != nil {
						return err
					}

					unit.Users = utils.RemoveFromSlice(unit.Users, i)
					s.state.UserIDToUnitID.Delete(toRemove[k])
				}
			}
		}
	}

	if len(toAdd) > 0 {
		notFound := []int32{}
		addIds := []jet.IntegerExpression{}
		for i := 0; i < len(toAdd); i++ {
			_, ok := s.tracker.GetUserById(toAdd[i])
			if !ok {
				continue
			}

			// Skip already added units
			if utils.InSliceFunc(unit.Users, func(in *dispatch.UnitAssignment) bool {
				return in.UserId == toAdd[i]
			}) {
				continue
			}

			addIds = append(addIds, jet.Int32(toAdd[i]))
			notFound = append(notFound, toAdd[i])
		}

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
					return err
				}
			}
		}

		users, err := s.resolveUserShortsByIds(ctx, notFound)
		if err != nil {
			return err
		}

		for _, user := range users {
			unit.Users = append(unit.Users, &dispatch.UnitAssignment{
				UnitId: unit.Id,
				UserId: user.UserId,
				User:   user,
			})

			if err := s.updateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
				UnitId:    unit.Id,
				Status:    dispatch.StatusUnit_STATUS_UNIT_USER_ADDED,
				UserId:    &user.UserId,
				CreatorId: &userInfo.UserId,
				X:         x,
				Y:         y,
				Postal:    postal,
			}); err != nil {
				return err
			}

			s.state.UserIDToUnitID.Store(user.UserId, unit.Id)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return ErrFailedQuery
	}

	data, err := proto.Marshal(unit)
	if err != nil {
		return err
	}

	s.events.JS.PublishAsync(buildSubject(TopicUnit, TypeUnitUpdated, userInfo.Job, unit.Id), data)

	// Unit is empty, set unit status to be unavailable automatically
	if len(unit.Users) == 0 {
		if err := s.updateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
			UnitId:    unit.Id,
			Status:    dispatch.StatusUnit_STATUS_UNIT_UNAVAILABLE,
			UserId:    &userInfo.UserId,
			CreatorId: &userInfo.UserId,
			X:         x,
			Y:         y,
			Postal:    postal,
		}); err != nil {
			return err
		}
	}

	return nil
}
