package rector

import (
	"context"

	rector "github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
)

var (
	tConfig = table.FivenetConfig
)

func (s *Server) GetAppConfig(ctx context.Context, req *GetAppConfigRequest) (*GetAppConfigResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorConfigService_ServiceDesc.ServiceName,
		Method:  "GetAppConfig",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	config, err := s.appCfg.LoadFromDB(ctx)
	if err != nil {
		return nil, err
	}

	return &GetAppConfigResponse{
		Config: config,
	}, nil
}

func (s *Server) UpdateAppConfig(ctx context.Context, req *UpdateAppConfigRequest) (*UpdateAppConfigResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorConfigService_ServiceDesc.ServiceName,
		Method:  "UpdateAppConfig",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	req.Config.Default()

	stmt := tConfig.
		INSERT(
			tConfig.Key,
			tConfig.AppConfig,
		).
		VALUES(
			1,
			req.Config,
		).
		ON_DUPLICATE_KEY_UPDATE(
			tConfig.AppConfig.SET(jet.RawString("VALUES(`app_config`)")),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	// Update default perms
	defaultPerms := make([]string, len(req.Config.Perms.Default))
	for i := 0; i < len(req.Config.Perms.Default); i++ {
		defaultPerms[i] = perms.BuildGuard(perms.Category(req.Config.Perms.Default[i].Category), perms.Name(req.Config.Perms.Default[i].Name))
	}
	if err := s.ps.SetDefaultRolePerms(ctx, defaultPerms); err != nil {
		return nil, err
	}

	// Update config state
	if err := s.appCfg.Update(req.Config); err != nil {
		return nil, err
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	config, err := s.appCfg.LoadFromDB(ctx)
	if err != nil {
		return nil, err
	}

	return &UpdateAppConfigResponse{
		Config: config,
	}, nil
}
