package manager

import (
	"context"

	"github.com/galexrt/fivenet/gen/go/proto/resources/dispatch"
	eventscentrum "github.com/galexrt/fivenet/gen/go/proto/services/centrum/events"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/pkg/utils/dbutils"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"google.golang.org/protobuf/proto"
)

func (s *Manager) GetUnitStatusFromDB(ctx context.Context, job string, id uint64) (*dispatch.UnitStatus, error) {
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
		unit, ok := s.GetUnit(job, dest.UnitId)
		if ok && unit != nil {
			newUnit := proto.Clone(unit)
			dest.Unit = newUnit.(*dispatch.Unit)
		}
	}

	return &dest, nil
}

func (s *Manager) UpdateUnitStatus(ctx context.Context, job string, unit *dispatch.Unit, in *dispatch.UnitStatus) error {
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

	status, err := s.GetUnitStatusFromDB(ctx, job, uint64(lastId))
	if err != nil {
		return err
	}
	unit.Status = status

	data, err := proto.Marshal(unit)
	if err != nil {
		return err
	}
	s.events.JS.PublishAsync(eventscentrum.BuildSubject(eventscentrum.TopicUnit, eventscentrum.TypeUnitStatus, job, status.UnitId), data)

	return nil
}

func (s *Manager) UpdateUnitAssignments(ctx context.Context, userInfo *userinfo.UserInfo, unit *dispatch.Unit, toAdd []int32, toRemove []int32) error {
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
		return err
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

		for i := len(unit.Users) - 1; i >= 0; i-- {
			if i > len(unit.Users)-1 {
				break
			}

			for k := 0; k < len(toRemove); k++ {
				if unit.Users[i].UserId == toRemove[k] {
					unit.Users = utils.RemoveFromSlice(unit.Users, i)

					if err := s.UpdateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
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

					s.UserIDToUnitID.Delete(toRemove[k])
					break
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

			if err := s.UpdateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
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

			s.UserIDToUnitID.Store(user.UserId, unit.Id)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	data, err := proto.Marshal(unit)
	if err != nil {
		return err
	}
	s.events.JS.PublishAsync(eventscentrum.BuildSubject(eventscentrum.TopicUnit, eventscentrum.TypeUnitUpdated, userInfo.Job, unit.Id), data)

	// Unit is empty, set unit status to be unavailable automatically
	if len(unit.Users) == 0 {
		if err := s.UpdateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
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