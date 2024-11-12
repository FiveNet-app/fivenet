package mailer

import (
	"context"
	"errors"
	"slices"

	"github.com/fivenet-app/fivenet/gen/go/proto/resources/mailer"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/rector"
	errorsmailer "github.com/fivenet-app/fivenet/gen/go/proto/services/mailer/errors"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	"github.com/fivenet-app/fivenet/pkg/grpc/errswrap"
	"github.com/fivenet-app/fivenet/query/fivenet/model"
	"github.com/fivenet-app/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

var tSettingsBlocks = table.FivenetMsgsSettingsBlocks

func (s *Server) SetThreadUserState(ctx context.Context, req *SetThreadUserStateRequest) (*SetThreadUserStateResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	check, err := s.access.CanUserAccessTarget(ctx, req.State.ThreadId, userInfo, mailer.AccessLevel_ACCESS_LEVEL_PARTICIPANT)
	if err != nil {
		return nil, errswrap.NewError(err, errorsmailer.ErrFailedQuery)
	}
	if !check && !userInfo.SuperUser {
		if !userInfo.SuperUser {
			return nil, errorsmailer.ErrFailedQuery
		}
	}

	tThreadsUserState := table.FivenetMsgsThreadsUserState
	stmt := tThreadsUserState.
		INSERT(
			tThreadsUserState.ThreadID,
			tThreadsUserState.UserID,
			tThreadsUserState.Unread,
			tThreadsUserState.LastRead,
			tThreadsUserState.Important,
			tThreadsUserState.Favorite,
			tThreadsUserState.Muted,
			tThreadsUserState.Archived,
		).
		VALUES(
			req.State.ThreadId,
			userInfo.UserId,
			req.State.Unread,
			req.State.LastRead,
			req.State.Important,
			req.State.Favorite,
			req.State.Muted,
			req.State.Archived,
		).
		ON_DUPLICATE_KEY_UPDATE(
			tThreadsUserState.LastRead.SET(jet.RawTimestamp("VALUES(`last_read`)")),
			tThreadsUserState.Important.SET(jet.Bool(req.State.Important)),
			tThreadsUserState.Favorite.SET(jet.Bool(req.State.Favorite)),
			tThreadsUserState.Muted.SET(jet.Bool(req.State.Muted)),
			tThreadsUserState.Archived.SET(jet.Bool(req.State.Archived)),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, errswrap.NewError(err, errorsmailer.ErrFailedQuery)
	}

	return &SetThreadUserStateResponse{}, nil
}

func (s *Server) setUnreadState(ctx context.Context, threadId uint64, userIds []int32) error {
	if len(userIds) == 0 {
		return nil
	}

	tThreadsUserState := table.FivenetMsgsThreadsUserState
	stmt := tThreadsUserState.
		INSERT(
			tThreadsUserState.ThreadID,
			tThreadsUserState.UserID,
			tThreadsUserState.Unread,
		)

	for _, userId := range userIds {
		stmt = stmt.VALUES(
			threadId,
			userId,
			true,
		)
	}

	stmt = stmt.ON_DUPLICATE_KEY_UPDATE(
		tThreadsUserState.Unread.SET(jet.RawBool("VALUES(`unread`)")),
	)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return err
	}

	return nil
}

func (s *Server) GetUserSettings(ctx context.Context, req *GetUserSettingsRequest) (*GetUserSettingsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	stmt := tSettingsBlocks.
		SELECT(
			tSettingsBlocks.SourceID,
			tSettingsBlocks.UserID,
			tUsers.Firstname,
			tUsers.Lastname,
			tUsers.Dateofbirth,
		).
		FROM(
			tSettingsBlocks.
				INNER_JOIN(tUsers,
					tUsers.ID.EQ(tSettingsBlocks.SourceID),
				),
		).
		WHERE(
			tSettingsBlocks.SourceID.EQ(jet.Int32(userInfo.UserId)),
		).
		LIMIT(25)

	resp := &GetUserSettingsResponse{}
	if err := stmt.QueryContext(ctx, s.db, &resp); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorsmailer.ErrFailedQuery)
		}
	}

	return resp, nil
}

type blockedUserHelper struct {
	SourceID int32
	UserID   int32
}

func (s *Server) SetUserSettings(ctx context.Context, req *SetUserSettingsRequest) (*SetUserSettingsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: MailerService_ServiceDesc.ServiceName,
		Method:  "SetUserSettings",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	settings, err := s.GetUserSettings(ctx, &GetUserSettingsRequest{})
	if err != nil {
		return nil, errswrap.NewError(err, errorsmailer.ErrFailedQuery)
	}

	// Handle blocked users changes
	if len(req.Settings.BlockedUsers) == 0 {
		if len(settings.Settings.BlockedUsers) > 0 {
			stmt := tSettingsBlocks.
				DELETE().
				WHERE(tSettingsBlocks.SourceID.EQ(jet.Int32(userInfo.UserId)))

			if _, err := stmt.ExecContext(ctx, s.db); err != nil {
				return nil, errswrap.NewError(err, errorsmailer.ErrFailedQuery)
			}
		}
	} else {
		toCreate := []*mailer.BlockedUser{}
		toUpdate := []*mailer.BlockedUser{}

		for _, bu := range req.Settings.BlockedUsers {
			if slices.ContainsFunc(settings.Settings.BlockedUsers, func(a *mailer.BlockedUser) bool {
				return a.UserId == bu.UserId
			}) {
				toUpdate = append(toUpdate, bu)
			} else {
				toCreate = append(toCreate, bu)
			}
		}

		// TODO
	}

	// TODO need a check logic for thread creation

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return nil, nil
}
