package audit

import (
	"context"
	"database/sql"
	"time"

	"github.com/galexrt/fivenet/pkg/config"
	jet "github.com/go-jet/jet/v2/mysql"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var RetentionModule = fx.Module("audit_retention",
	fx.Provide(
		NewRetention,
	),
)

type Retention struct {
	ctx    context.Context
	tracer trace.Tracer
	logger *zap.Logger
	db     *sql.DB

	auditRetentionDays *int
}

type RetentionParams struct {
	fx.In

	LC fx.Lifecycle

	Logger *zap.Logger
	TP     *tracesdk.TracerProvider
	DB     *sql.DB
	Cfg    *config.Config
}

func NewRetention(p RetentionParams) *Retention {
	ctx, cancel := context.WithCancel(context.Background())

	r := &Retention{
		ctx:    ctx,
		logger: p.Logger.Named("audit_retention"),
		tracer: p.TP.Tracer("audit-retention"),
		db:     p.DB,

		auditRetentionDays: p.Cfg.Game.AuditRetentionDays,
	}

	p.LC.Append(fx.StartHook(func(_ context.Context) error {
		go func() {
			for {
				func() {
					ctx, span := r.tracer.Start(r.ctx, "audit-retention")
					defer span.End()

					r.run(ctx)

					select {
					case <-ctx.Done():
						return
					case <-time.After(30 * time.Minute):
					}
				}()
			}
		}()
		return nil
	}))

	p.LC.Append(fx.StopHook(func(ctx context.Context) error {
		cancel()
		return nil
	}))

	return r
}

func (r *Retention) run(ctx context.Context) error {
	if r.auditRetentionDays != nil {
		// Now minus retention days
		t := time.Now().AddDate(0, 0, -*r.auditRetentionDays)
		if err := r.Cleanup(t); err != nil {
			return err
		}
	}

	return nil
}

func (r *Retention) Cleanup(before time.Time) error {
	r.logger.Debug("starting audit store cleanup", zap.Time("before_time", before))

	stmt := tAudit.
		DELETE().
		WHERE(tAudit.CreatedAt.LT_EQ(
			jet.TimestampT(before),
		))

	res, err := stmt.ExecContext(r.ctx, r.db)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	r.logger.Info("audit store cleanup completed", zap.Int64("affected_rows", affected))

	return nil
}
