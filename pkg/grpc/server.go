package grpc

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/galexrt/fivenet/pkg/config"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	grpc_auth "github.com/galexrt/fivenet/pkg/grpc/interceptors/auth"
	grpc_permission "github.com/galexrt/fivenet/pkg/grpc/interceptors/permission"
	grpc_sanitizer "github.com/galexrt/fivenet/pkg/grpc/interceptors/sanitizer"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/getsentry/sentry-go"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternalServer = status.Error(codes.Internal, "errors.general.internal_error.title;errors.general.internal_error.content")
)

func wrapLogger(log *zap.Logger) *zap.Logger {
	return log.Named("grpc_server")
}

var ServerModule = fx.Module("grpcserver",
	fx.Provide(
		NewServer,
	),
	fx.Decorate(wrapLogger),
)

type ServerParams struct {
	fx.In

	LC fx.Lifecycle

	Logger   *zap.Logger
	Config   *config.Config
	DB       *sql.DB
	TP       *tracesdk.TracerProvider
	Services []Service `group:"grpcservices"`
	TokenMgr *auth.TokenMgr
	UserInfo userinfo.UserInfoRetriever
	Perms    perms.Permissions
}

type ServerResult struct {
	fx.Out

	Server *grpc.Server
}

func NewServer(p ServerParams) (ServerResult, error) {
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}

	// Setup metrics
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	prometheus.MustRegister(srvMetrics)
	exemplarFromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}

	// Setup GRPC tracing
	otel.SetTracerProvider(p.TP)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Setup metric for panic recoveries
	panicsTotal := promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
		Name: "grpc_req_panics_recovered_total",
		Help: "Total number of gRPC requests recovered from internal panic.",
	})
	grpcPanicRecoveryHandler := func(pa any) (err error) {
		panicsTotal.Inc()

		p.Logger.Error("recovered from panic", zap.Any("err", pa), zap.Stack("stacktrace"))
		if e, ok := pa.(error); ok {
			sentry.CaptureException(e)
		}

		return ErrInternalServer
	}

	grpcAuth := auth.NewGRPCAuth(p.UserInfo, p.TokenMgr)
	grpcPerm := auth.NewGRPCPerms(p.Perms)

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.UnaryServerInterceptor(InterceptorLogger(p.Logger),
				logging.WithFieldsFromContext(logTraceID),
				logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
			),
			grpc_auth.UnaryServerInterceptor(grpcAuth.GRPCAuthFunc),
			validator.UnaryServerInterceptor(),
			grpc_permission.UnaryServerInterceptor(grpcPerm.GRPCPermissionUnaryFunc),
			grpc_sanitizer.UnaryServerInterceptor(),
			recovery.UnaryServerInterceptor(
				recovery.WithRecoveryHandler(grpcPanicRecoveryHandler),
			),
		),
		grpc.ChainStreamInterceptor(
			otelgrpc.StreamServerInterceptor(),
			srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			logging.StreamServerInterceptor(InterceptorLogger(p.Logger),
				logging.WithFieldsFromContext(logTraceID),
				logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
			),
			grpc_auth.StreamServerInterceptor(grpcAuth.GRPCAuthFunc),
			validator.StreamServerInterceptor(),
			grpc_permission.StreamServerInterceptor(grpcPerm.GRPCPermissionStreamFunc),
			recovery.StreamServerInterceptor(
				recovery.WithRecoveryHandler(grpcPanicRecoveryHandler),
			),
		),
	)

	for _, service := range p.Services {
		service.RegisterServer(srv)
	}

	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", p.Config.GRPC.Listen)
			if err != nil {
				return err
			}
			p.Logger.Info("grpc server listening", zap.String("address", p.Config.HTTP.Listen))
			go func() {
				if err := srv.Serve(ln); err != nil {
					p.Logger.Error("failed to serve grpc server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			go func() {
				srv.GracefulStop()
			}()
			// Wait 3 seconds before "forceful stop
			time.Sleep(3 * time.Second)

			srv.Stop()
			return nil
		},
	})

	return ServerResult{
		Server: srv,
	}, nil
}

// InterceptorLogger adapts zap logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)
		iter := logging.Fields(fields).Iterator()
		for iter.Next() {
			k, v := iter.At()
			f = append(f, zap.Any(k, v))
		}
		l := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg)
		case logging.LevelInfo:
			l.Info(msg)
		case logging.LevelWarn:
			l.Warn(msg)
		case logging.LevelError:
			l.Error(msg)
		default:
			l.Error(fmt.Sprintf("unknown log level '%v' for message", lvl), zap.String("msg", msg))
		}
	})
}