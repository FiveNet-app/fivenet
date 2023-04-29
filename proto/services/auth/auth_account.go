package auth

import (
	"context"
	"errors"

	"github.com/galexrt/fivenet/pkg/auth"
	"github.com/galexrt/fivenet/pkg/config"
	accounts "github.com/galexrt/fivenet/proto/resources/accounts"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	oAuth2Accounts = table.FivenetOauth2Accounts

	GenericAccountErr = status.Error(codes.Internal, "Failed to get/update your account, please try again.")
	InvalidTokenErr   = status.Error(codes.Unauthenticated, "")
)

func (s *Server) GetAccountInfo(ctx context.Context, req *GetAccountInfoRequest) (*GetAccountInfoResponse, error) {
	token, err := auth.GetTokenFromGRPCContext(ctx)
	if err != nil {
		return nil, auth.InvalidTokenErr
	}

	claims, err := s.tm.ParseWithClaims(token)
	if err != nil {
		return nil, GenericAccountErr
	}

	// Load account
	acc, err := s.getAccountFromDB(ctx, account.ID.EQ(jet.Uint64(claims.AccountID)))
	if err != nil {
		return nil, GenericAccountErr
	}
	if acc.ID == 0 {
		return nil, GenericAccountErr
	}

	oauth2Providers := make([]*accounts.OAuth2Provider, len(config.C.OAuth2.Providers))
	for i := 0; i < len(oauth2Providers); i++ {
		p := config.C.OAuth2.Providers[i]
		oauth2Providers[i] = &accounts.OAuth2Provider{
			Name:     p.Name,
			Label:    p.Label,
			Homepage: p.Homepage,
		}
	}

	oAuth2Accounts := table.FivenetOauth2Accounts.AS("oauth2account")
	stmt := oAuth2Accounts.
		SELECT(
			oAuth2Accounts.AccountID,
			oAuth2Accounts.CreatedAt,
			oAuth2Accounts.Provider.AS("oauth2account.providername"),
			oAuth2Accounts.ExternalID,
			oAuth2Accounts.Username,
			oAuth2Accounts.Avatar,
		).
		FROM(
			oAuth2Accounts,
		).
		WHERE(
			oAuth2Accounts.AccountID.EQ(jet.Uint64(acc.ID)),
		).
		LIMIT(3)

	oauth2Conns := []*accounts.OAuth2Account{}
	if err := stmt.QueryContext(ctx, s.db, &oauth2Conns); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, GenericAccountErr
		}
	}

	// Set provider in the connections
	for i := 0; i < len(oauth2Conns); i++ {
		for _, p := range oauth2Providers {
			if p.Name == oauth2Conns[i].GetProviderName() {
				oauth2Conns[i].Provider = p
				break
			}
		}
	}

	return &GetAccountInfoResponse{
		Account:           accounts.ConvertFromAcc(acc),
		Oauth2Providers:   oauth2Providers,
		Oauth2Connections: oauth2Conns,
	}, nil
}
