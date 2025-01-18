package notificator

import (
	"database/sql"

	pbnotificator "github.com/fivenet-app/fivenet/gen/go/proto/services/notificator"
	"github.com/fivenet-app/fivenet/pkg/config/appconfig"
	"github.com/fivenet-app/fivenet/pkg/events"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth/userinfo"
	"github.com/fivenet-app/fivenet/pkg/mstlystcdata"
	"github.com/fivenet-app/fivenet/pkg/perms"
	"go.uber.org/fx"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

type Server struct {
	pbnotificator.NotificatorServiceServer

	logger   *zap.Logger
	db       *sql.DB
	p        perms.Permissions
	tm       *auth.TokenMgr
	ui       userinfo.UserInfoRetriever
	js       *events.JSWrapper
	enricher *mstlystcdata.Enricher
	appCfg   appconfig.IConfig
}

type Params struct {
	fx.In

	LC fx.Lifecycle

	Logger    *zap.Logger
	DB        *sql.DB
	Perms     perms.Permissions
	TM        *auth.TokenMgr
	UI        userinfo.UserInfoRetriever
	JS        *events.JSWrapper
	Enricher  *mstlystcdata.Enricher
	AppConfig appconfig.IConfig
}

func NewServer(p Params) *Server {
	s := &Server{
		logger:   p.Logger,
		db:       p.DB,
		p:        p.Perms,
		tm:       p.TM,
		ui:       p.UI,
		js:       p.JS,
		enricher: p.Enricher,
		appCfg:   p.AppConfig,
	}

	return s
}

func (s *Server) RegisterServer(srv *grpc.Server) {
	pbnotificator.RegisterNotificatorServiceServer(srv, s)
}

func (s *Server) GetPermsRemap() map[string]string {
	return pbnotificator.PermsRemap
}