package manager

import (
	"context"

	"github.com/galexrt/fivenet/gen/go/proto/resources/centrum"
	errorscentrum "github.com/galexrt/fivenet/gen/go/proto/services/centrum/errors"
	eventscentrum "github.com/galexrt/fivenet/gen/go/proto/services/centrum/events"
	jet "github.com/go-jet/jet/v2/mysql"
	"google.golang.org/protobuf/proto"
)

func (s *Manager) UpdateSettingsInDB(ctx context.Context, job string, settings *centrum.Settings) (*centrum.Settings, error) {
	stmt := tCentrumSettings.
		INSERT(
			tCentrumSettings.Job,
			tCentrumSettings.Enabled,
			tCentrumSettings.Mode,
			tCentrumSettings.FallbackMode,
		).
		VALUES(
			job,
			settings.Enabled,
			settings.Mode,
			settings.FallbackMode,
		).
		ON_DUPLICATE_KEY_UPDATE(
			tCentrumSettings.Job.SET(jet.String(job)),
			tCentrumSettings.Enabled.SET(jet.Bool(settings.Enabled)),
			tCentrumSettings.Mode.SET(jet.Int32(int32(settings.Mode))),
			tCentrumSettings.FallbackMode.SET(jet.Int32(int32(settings.FallbackMode))),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, errorscentrum.ErrFailedQuery
	}

	// Load settings from database so they are updated in the "cache"
	if err := s.LoadSettings(ctx, job); err != nil {
		return nil, errorscentrum.ErrFailedQuery
	}

	set := s.GetSettings(job)

	data, err := proto.Marshal(set)
	if err != nil {
		return nil, errorscentrum.ErrFailedQuery
	}

	if _, err := s.js.Publish(eventscentrum.BuildSubject(eventscentrum.TopicGeneral, eventscentrum.TypeGeneralSettings, job), data); err != nil {
		return nil, err
	}

	return set, nil
}
