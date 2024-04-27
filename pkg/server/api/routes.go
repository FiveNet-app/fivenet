package api

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/galexrt/fivenet/pkg/config"
	"github.com/galexrt/fivenet/pkg/config/appconfig"
	"github.com/galexrt/fivenet/pkg/version"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Routes struct {
	logger *zap.Logger

	clientCfg *atomic.Pointer[ClientConfig]
}

type Params struct {
	fx.In

	LC fx.Lifecycle

	Logger    *zap.Logger
	Config    *config.Config
	AppConfig appconfig.IConfig
}

func New(p Params) *Routes {
	r := &Routes{
		logger:    p.Logger,
		clientCfg: &atomic.Pointer[ClientConfig]{},
	}

	providers := make([]*ProviderConfig, len(p.Config.OAuth2.Providers))

	for i, p := range p.Config.OAuth2.Providers {
		providers[i] = &ProviderConfig{
			Name:  p.Name,
			Label: p.Label,
			Icon:  p.Icon,
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	p.LC.Append(fx.StartHook(func(_ context.Context) error {
		r.handleAppConfigUpdate(providers, p.AppConfig.Get())

		// Handle app config updates
		go func() {
			configUpdateCh := p.AppConfig.Subscribe()
			for {
				select {
				case <-ctx.Done():
					p.AppConfig.Unsubscribe(configUpdateCh)
					return

				case cfg := <-configUpdateCh:
					r.handleAppConfigUpdate(providers, cfg)
				}
			}
		}()

		return nil
	}))

	p.LC.Append(fx.StopHook(func() error {
		cancel()

		return nil
	}))

	return r
}

func (r *Routes) handleAppConfigUpdate(providers []*ProviderConfig, appCfg *appconfig.Cfg) {
	clientCfg := r.buildClientConfig(providers, appCfg)
	r.clientCfg.Store(clientCfg)
}

func (r *Routes) buildClientConfig(providers []*ProviderConfig, appCfg *appconfig.Cfg) *ClientConfig {
	clientCfg := &ClientConfig{
		Version: version.Version,
		Login: LoginConfig{
			SignupEnabled: appCfg.Auth.SignupEnabled,
			LastCharLock:  appCfg.Auth.LastCharLock,
			Providers:     providers,
		},
		Discord: Discord{},
		Links:   Links{},
	}

	clientCfg.Discord.BotInviteURL = appCfg.Discord.InviteUrl

	if appCfg.Website.Links.Imprint != nil {
		clientCfg.Links.Imprint = appCfg.Website.Links.Imprint
	}
	if appCfg.Website.Links.PrivacyPolicy != nil {
		clientCfg.Links.PrivacyPolicy = appCfg.Website.Links.PrivacyPolicy
	}

	return clientCfg
}

var versionInfo = &Version{
	Version: version.Version,
}

func (r *Routes) RegisterHTTP(e *gin.Engine) {
	// API Base
	g := e.Group("/api")
	{
		g.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, PingResponse)
		})

		g.POST("/config", func(c *gin.Context) {
			c.JSON(http.StatusOK, r.clientCfg.Load())
		})

		g.GET("/clear-site-data", func(c *gin.Context) {
			c.Header("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\"")
			c.String(http.StatusOK, "Your local site data should be cleared now, please go back to the FiveNet homepage yourself.")
		})

		g.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, versionInfo)
		})
	}
}
