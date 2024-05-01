package auth

import (
	"context"
	"fmt"
	"os"
	"testing"

	errorsauth "github.com/fivenet-app/fivenet/gen/go/proto/services/auth/errors"
	permsauth "github.com/fivenet-app/fivenet/gen/go/proto/services/auth/perms"
	"github.com/fivenet-app/fivenet/internal/modules"
	"github.com/fivenet-app/fivenet/internal/tests/proto"
	"github.com/fivenet-app/fivenet/internal/tests/servers"
	grpcserver "github.com/fivenet-app/fivenet/pkg/grpc"
	"github.com/fivenet-app/fivenet/pkg/perms"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestMain(m *testing.M) {
	if err := servers.TestDBServer.Setup(); err != nil {
		fmt.Println("failed to setup mysql test server: %w", err)
		return
	}
	defer servers.TestDBServer.Stop()

	if err := servers.TestNATSServer.Setup(); err != nil {
		fmt.Println("failed to setup nats test server: %w", err)
		return
	}
	defer servers.TestNATSServer.Stop()

	code := m.Run()

	os.Exit(code)
}

func TestFullAuthFlow(t *testing.T) {
	defer servers.TestDBServer.Reset()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientConn, grpcSrvModule, err := modules.TestGRPCServer(ctx)
	require.NoError(t, err)

	var srv *Server
	app := fxtest.New(t,
		modules.GetFxTestOpts(
			fx.Provide(grpcSrvModule),
			fx.Provide(grpcserver.AsService(func(p Params) *Server {
				srv = NewServer(p)
				return srv
			})),

			fx.Invoke(func(*grpc.Server) {}),
		)...,
	)
	assert.NotNil(t, app)

	app.RequireStart()
	assert.NotNil(t, srv)

	client := NewAuthServiceClient(clientConn)

	// First login without credentials
	loginReq := &LoginRequest{}
	loginReq.Username = ""
	loginReq.Password = ""
	res, err := client.Login(ctx, loginReq)
	assert.Error(t, err)
	assert.Nil(t, res)
	proto.CompareGRPCError(t, errorsauth.ErrInvalidLogin, err)

	// Login with invalid credentials
	loginReq.Username = "non-existant-username"
	loginReq.Password = "non-existant-password"
	res, err = client.Login(ctx, loginReq)
	assert.Error(t, err)
	assert.Nil(t, res)
	proto.CompareGRPCError(t, errorsauth.ErrInvalidLogin, err)

	// user-3: Login with valid account that has one char
	loginReq.Username = "user-3"
	loginReq.Password = "password"
	res, err = client.Login(ctx, loginReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	if res == nil {
		assert.FailNow(t, "user-3: Login with valid account failed, response is nil")
	}
	assert.NotEmpty(t, res.GetToken())

	// user-3: Create authenticated metadate and get characters (only has one char)
	md := metadata.New(map[string]string{"Authorization": "Bearer " + res.GetToken()})
	ctx = metadata.NewOutgoingContext(ctx, md)
	getCharsReq := &GetCharactersRequest{}
	getCharsRes, err := client.GetCharacters(ctx, getCharsReq)
	assert.NoError(t, err)
	assert.NotNil(t, getCharsRes)
	if getCharsRes == nil {
		assert.FailNow(t, "user-3: Empty char list returned for valid account that should have 2 chars")
	}
	assert.Len(t, getCharsRes.GetChars(), 1)

	// user-1: Login with valid account (2 chars)
	loginReq.Username = "user-1"
	loginReq.Password = "password"
	res, err = client.Login(ctx, loginReq)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	if res == nil {
		assert.FailNow(t, "user-1: Login with valid account failed, response is nil")
	}
	assert.NotEmpty(t, res.GetToken())

	// user-1: Create authenticated metadate and get characters
	md = metadata.New(map[string]string{"Authorization": "Bearer " + res.GetToken()})
	ctx = metadata.NewOutgoingContext(ctx, md)
	getCharsReq = &GetCharactersRequest{}
	getCharsRes, err = client.GetCharacters(ctx, getCharsReq)
	assert.NoError(t, err)
	assert.NotNil(t, getCharsRes)
	if getCharsRes == nil {
		assert.FailNow(t, "user-1: Empty char list returned for valid account that should have 2 chars")
	}
	assert.Len(t, getCharsRes.GetChars(), 2)

	// user-1: Choose an invalid character
	chooseCharReq := &ChooseCharacterRequest{}
	chooseCharReq.CharId = 2 // Char id 2 is `user-2`'s char
	chooseCharRes, err := client.ChooseCharacter(ctx, chooseCharReq)
	assert.Error(t, err)
	assert.Nil(t, chooseCharRes)
	proto.CompareGRPCError(t, errorsauth.ErrUnableToChooseChar, err)

	role, err := srv.ps.GetRoleByJobAndGrade(ctx, "ambulance", 1)
	assert.NoError(t, err)
	assert.NotNil(t, role)

	perm, err := srv.ps.GetPermission(ctx, permsauth.AuthServicePerm, permsauth.AuthServiceChooseCharacterPerm)
	assert.NoError(t, err)
	assert.NotNil(t, perm)

	// user-1: Choose valid character but we don't have permissions on the role
	err = srv.ps.RemovePermissionsFromRole(ctx, role.ID, perm.Id)
	assert.NoError(t, err)
	// Disable choose char perm from ambulance rank 1 role, `user-1` is a medic rank 1+
	err = srv.ps.UpdateRolePermissions(ctx, role.ID, perms.AddPerm{
		Id:  perm.Id,
		Val: false,
	})
	assert.NoError(t, err)
	chooseCharReq.CharId = 1
	chooseCharRes, err = client.ChooseCharacter(ctx, chooseCharReq)
	assert.Error(t, err)
	assert.Nil(t, chooseCharRes)
	proto.CompareGRPCError(t, errorsauth.ErrUnableToChooseChar, err)

	// user-1: Choose valid character, now we add a permssion
	err = srv.ps.UpdateRolePermissions(ctx, role.ID, perms.AddPerm{
		Id:  perm.Id,
		Val: true,
	})
	assert.NoError(t, err)
	chooseCharReq.CharId = 1
	chooseCharRes, err = client.ChooseCharacter(ctx, chooseCharReq)
	assert.NoError(t, err)
	assert.NotNil(t, chooseCharRes)

	app.RequireStop()
}
