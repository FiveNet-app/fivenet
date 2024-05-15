package oauth2

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/fivenet-app/fivenet/pkg/config"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	"github.com/fivenet-app/fivenet/pkg/server/oauth2/providers"
	"github.com/fivenet-app/fivenet/pkg/utils"
	"github.com/fivenet-app/fivenet/pkg/utils/dbutils"
	"github.com/fivenet-app/fivenet/query/fivenet/model"
	"github.com/fivenet-app/fivenet/query/fivenet/table"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var (
	tAccs      = table.FivenetAccounts
	tOAuthAccs = table.FivenetOauth2Accounts
)

type Params struct {
	fx.In

	Logger *zap.Logger
	DB     *sql.DB
	TM     *auth.TokenMgr
	Config *config.Config
}

type OAuth2 struct {
	logger *zap.Logger
	db     *sql.DB
	tm     *auth.TokenMgr

	domain       string
	oauthConfigs map[string]providers.IProvider
}

func New(p Params) *OAuth2 {
	if len(p.Config.OAuth2.Providers) == 0 {
		return nil
	}

	o := &OAuth2{
		logger:       p.Logger,
		db:           p.DB,
		tm:           p.TM,
		domain:       p.Config.HTTP.Sessions.Domain,
		oauthConfigs: make(map[string]providers.IProvider, len(p.Config.OAuth2.Providers)),
	}

	for _, p := range p.Config.OAuth2.Providers {
		cfg := &oauth2.Config{
			RedirectURL:  p.RedirectURL,
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			Scopes:       p.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:   p.Endpoints.AuthURL,
				TokenURL:  p.Endpoints.TokenURL,
				AuthStyle: oauth2.AuthStyleInParams,
			},
		}
		var provider providers.IProvider
		switch p.Type {
		case config.OAuth2ProviderDiscord:
			provider = &providers.Discord{
				BaseProvider: providers.BaseProvider{
					Name: p.Name,
				},
			}
		default:
			provider = &providers.Generic{
				BaseProvider: providers.BaseProvider{
					Name:          p.Name,
					UserInfoURL:   p.Endpoints.UserInfoURL,
					DefaultAvatar: p.DefaultAvatar,
				},
			}
		}

		provider.SetOauthConfig(cfg)
		provider.SetMapping(p.Mapping)

		o.oauthConfigs[p.Name] = provider
	}

	return o
}

func (o *OAuth2) RegisterHTTP(e *gin.Engine) {
	oauth := e.Group("/api/oauth2")
	{
		oauth.GET("/login/:provider", o.Login)
		oauth.POST("/login/:provider", o.Login)
		oauth.GET("/callback/:provider", o.Callback)
		oauth.POST("/callback/:provider", o.Callback)
	}
}

func (o *OAuth2) GetProvider(c *gin.Context) (providers.IProvider, error) {
	param, ok := c.Params.Get("provider")
	if !ok {
		return nil, fmt.Errorf("no provider found")
	}

	provider, ok := o.oauthConfigs[param]
	if !ok {
		return nil, fmt.Errorf("no provider found")
	}

	return provider, nil
}

const (
	AccountInfoRedirBase string = "/auth/account-info"
	LoginRedirBase       string = "/auth/login"

	ReasonInternalError string = "internal_error"
)

func (o *OAuth2) handleRedirect(c *gin.Context, connectOnly bool, success bool, reason string) {
	if !success {
		redirURL := ""
		if connectOnly {
			redirURL = AccountInfoRedirBase + "?oauth2Connect=failed"
		} else {
			redirURL = LoginRedirBase + "?oauth2Login=failed"
		}

		if reason != "" {
			redirURL = redirURL + "&reason=" + url.QueryEscape(reason)
		}

		c.Redirect(http.StatusTemporaryRedirect, redirURL)
		return
	}

	redirURL := ""
	if connectOnly {
		redirURL = AccountInfoRedirBase + "?oauth2Connect=success"
	} else {
		redirURL = LoginRedirBase + "?oauth2Login=success"
	}

	c.Redirect(http.StatusTemporaryRedirect, redirURL)
}

func (o *OAuth2) Login(c *gin.Context) {
	sess := sessions.DefaultMany(c, "fivenet_oauth2_state")
	connectOnly := false
	connectOnlyVal := c.Query("connect-only")
	if connectOnlyVal != "" {
		var err error
		connectOnly, err = strconv.ParseBool(connectOnlyVal)
		if err != nil {
			o.logger.Error("failed to parse connect only var", zap.Error(err))
			o.handleRedirect(c, false, false, "invalid_request_connect_only")
			return
		}
	}

	tokenVal, err := c.Cookie("fivenet_token")
	if err != nil {
		o.logger.Error("failed to parse token cookie", zap.Error(err))
		o.handleRedirect(c, false, false, "invalid_request_token")
		return
	}
	if tokenVal != "" {
		sess.Set("token", tokenVal)
	}

	state, err := utils.GenerateRandomString(64)
	if err != nil {
		o.logger.Error("failed to generate random string", zap.Error(err))
		o.handleRedirect(c, connectOnly, false, ReasonInternalError)
		return
	}

	sess.Set("connect-only", connectOnly)
	sess.Set("state", state)
	sess.Save()

	provider, err := o.GetProvider(c)
	if err != nil {
		o.logger.Error("failed to get provider", zap.Error(err))
		o.handleRedirect(c, connectOnly, false, "invalid_provider")
		return
	}

	// Redirect to provider for login
	c.Redirect(http.StatusTemporaryRedirect, provider.GetRedirect(state))
}

func (o *OAuth2) Callback(c *gin.Context) {
	sess := sessions.DefaultMany(c, "fivenet_oauth2_state")
	sessState := sess.Get("state")
	if sessState == nil {
		o.handleRedirect(c, false, false, "invalid_state")
		return
	}

	state := sessState.(string)
	connectOnly := false
	sessConnectOnly := sess.Get("connect-only")
	if sessConnectOnly != nil {
		connectOnly = sessConnectOnly.(bool)
	}

	var token string
	sessToken := sess.Get("token")
	if sessToken != nil {
		token = sessToken.(string)
	}

	// Remove vars from session
	sess.Delete("connect-only")
	sess.Delete("state")
	sess.Delete("token")
	sess.Save()

	if c.Request.FormValue("state") != state {
		o.handleRedirect(c, connectOnly, false, "invalid_state_404")
		return
	}

	provider, err := o.GetProvider(c)
	if err != nil {
		o.logger.Error("failed to get provider", zap.Error(err))
		o.handleRedirect(c, connectOnly, false, "invalid_provider")
		return
	}

	userInfo, err := provider.GetUserInfo(c, c.Request.FormValue("code"))
	if err != nil {
		o.logger.Error("failed to get userinfo from provider", zap.Error(err))
		o.handleRedirect(c, connectOnly, false, "provider_failed")
		return
	}

	if connectOnly {
		claims, err := o.tm.ParseWithClaims(token)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, LoginRedirBase+"?oauth2Connect=failed&reason=token_invalid")
			return
		}

		if err := o.storeUserInfo(c, claims.AccID, provider.GetName(), userInfo); err != nil {
			o.logger.Error("failed to store user info", zap.Error(err))
			if !dbutils.IsDuplicateError(err) {
				o.handleRedirect(c, connectOnly, false, ReasonInternalError)
			} else {
				o.handleRedirect(c, connectOnly, false, "already_in_use")
			}
			return
		}

		o.handleRedirect(c, connectOnly, true, "")
		return
	}

	// Take care of logging user in
	account, err := o.getUserInfo(c, provider.GetName(), userInfo)
	if err != nil {
		o.logger.Error("failed to store userinfo in database", zap.Error(err))
		o.handleRedirect(c, connectOnly, false, ReasonInternalError)
		return
	}

	if account == nil {
		c.Redirect(http.StatusTemporaryRedirect, LoginRedirBase+"?oauth2Login=failed&reason=unconnected")
		return
	} else if account.ID == 0 {
		o.logger.Error("invalid account id from userinfo", zap.Error(err))
		o.handleRedirect(c, connectOnly, true, ReasonInternalError)
		return
	}

	claims := auth.BuildTokenClaimsFromAccount(account, nil)
	newToken, err := o.tm.NewWithClaims(claims)
	if err != nil {
		o.logger.Error("failed to token from account", zap.Error(err))
		o.handleRedirect(c, connectOnly, true, ReasonInternalError)
		return
	}

	c.SetCookie(auth.CookieName, newToken, 6*24*60*60, "/", o.domain, false, true)

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf(LoginRedirBase+"?oauth2Login=success&u=%s&exp=%d",
		url.QueryEscape(*account.Username),
		claims.ExpiresAt.Time.UTC().UnixNano()/1e6,
	))
}

func (o *OAuth2) getUserInfo(ctx context.Context, provider string, userInfo *providers.UserInfo) (*model.FivenetAccounts, error) {
	stmt := tOAuthAccs.
		SELECT(
			tAccs.ID,
			tAccs.Username,
			tAccs.License,
		).
		FROM(tOAuthAccs.
			INNER_JOIN(tAccs,
				tAccs.ID.EQ(tOAuthAccs.AccountID),
			),
		).
		WHERE(jet.AND(
			tOAuthAccs.Provider.EQ(jet.String(provider)),
			tOAuthAccs.ExternalID.EQ(jet.String(userInfo.ID)),
			tAccs.Enabled.IS_TRUE(),
		)).
		LIMIT(1)

	var dest model.FivenetAccounts
	if err := stmt.QueryContext(ctx, o.db, &dest); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
		return nil, nil
	}

	return &dest, nil
}

func (o *OAuth2) storeUserInfo(ctx context.Context, accountId uint64, provider string, userInfo *providers.UserInfo) error {
	stmt := tOAuthAccs.
		INSERT(
			tOAuthAccs.AccountID,
			tOAuthAccs.Provider,
			tOAuthAccs.ExternalID,
			tOAuthAccs.Username,
			tOAuthAccs.Avatar,
		).
		VALUES(
			accountId,
			provider,
			userInfo.ID,
			userInfo.Username,
			userInfo.Avatar,
		)

	if _, err := stmt.ExecContext(ctx, o.db); err != nil {
		return err
	}

	return nil
}
