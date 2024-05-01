package modules

import (
	"context"
	"fmt"
	"net"

	grpcsvc "github.com/fivenet-app/fivenet/pkg/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	grpc_auth "github.com/fivenet-app/fivenet/pkg/grpc/interceptors/auth"
	grpc_permission "github.com/fivenet-app/fivenet/pkg/grpc/interceptors/permission"
)

type GRPCServerParams struct {
	fx.In

	LC fx.Lifecycle

	Logger *zap.Logger

	GRPCAuth *auth.GRPCAuth
	GRPCPerm *auth.GRPCPerm

	Services []grpcsvc.Service `group:"grpcservices"`
}

func TestGRPCServer(ctx context.Context) (*grpc.ClientConn, func(p GRPCServerParams) (*grpc.Server, error), error) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to test grpc server: %v", err)
	}

	return conn, func(p GRPCServerParams) (*grpc.Server, error) {
		srv := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				grpc_auth.UnaryServerInterceptor(p.GRPCAuth.GRPCAuthFunc),
				grpc_permission.UnaryServerInterceptor(p.GRPCPerm.GRPCPermissionUnaryFunc),
				recovery.UnaryServerInterceptor(),
			),
			grpc.ChainStreamInterceptor(
				grpc_auth.StreamServerInterceptor(p.GRPCAuth.GRPCAuthFunc),
				grpc_permission.StreamServerInterceptor(p.GRPCPerm.GRPCPermissionStreamFunc),
				recovery.StreamServerInterceptor(),
			),
		)

		for _, service := range p.Services {
			service.RegisterServer(srv)
		}

		p.LC.Append(fx.StartHook(func() error {
			go func() {
				if err := srv.Serve(lis); err != nil {
					p.Logger.Error("error serving test grpc server", zap.Error(err))
					return
				}
			}()

			return nil
		}))

		p.LC.Append(fx.StopHook(func() error {
			err := lis.Close()
			if err != nil {
				return err
			}
			srv.Stop()

			return nil
		}))

		return srv, nil
	}, nil
}
