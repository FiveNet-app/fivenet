package croner

import (
	"context"
	"fmt"
	"time"

	"github.com/fivenet-app/fivenet/gen/go/proto/resources/common/cron"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/timestamp"
	"github.com/fivenet-app/fivenet/pkg/events"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
)

var AgentModule = fx.Module("cron",
	fx.Provide(
		NewAgent,
	),
)

type AgentParams struct {
	fx.In

	LC fx.Lifecycle

	Logger *zap.Logger
	JS     *events.JSWrapper

	Handlers *Handlers
}

type Agent struct {
	logger *zap.Logger
	ctx    context.Context
	js     *events.JSWrapper

	handlers *Handlers
}

func NewAgent(p AgentParams) (*Agent, error) {
	ctx, cancel := context.WithCancel(context.Background())

	ag := &Agent{
		logger: p.Logger.Named("cron_agent"),
		ctx:    ctx,
		js:     p.JS,

		handlers: p.Handlers,
	}

	p.LC.Append(fx.StartHook(func(c context.Context) error {
		if err := registerCronStreams(ctx, ag.js); err != nil {
			return err
		}

		return ag.registerSubscriptions(ctx, p.JS)
	}))

	p.LC.Append(fx.StopHook(func(c context.Context) error {
		cancel()

		return nil
	}))

	return ag, nil
}

func (ag *Agent) registerSubscriptions(ctx context.Context, js *events.JSWrapper) error {
	consumer, err := js.CreateConsumer(ctx, CronScheduleStreamName, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverNewPolicy,
		FilterSubject: fmt.Sprintf("%s.%s", CronScheduleSubject, CronScheduleTopic),
		MaxDeliver:    3,
	})
	if err != nil {
		return err
	}

	if _, err := consumer.Consume(ag.watchForEvents,
		js.ConsumeErrHandlerWithRestart(context.Background(), ag.logger,
			func(ctx context.Context, c context.Context) error {
				return ag.registerSubscriptions(c, js)
			})); err != nil {
		return err
	}

	return nil
}

func (ag *Agent) watchForEvents(msg jetstream.Msg) {
	job := &cron.CronjobSchedulerEvent{}
	if err := protojson.Unmarshal(msg.Data(), job); err != nil {
		ag.logger.Error("failed to unmarshal cron schedule msg", zap.String("subject", msg.Subject()), zap.Error(err))

		if err := msg.Nak(); err != nil {
			ag.logger.Error("failed to nack unmarshal cron schedule msg", zap.String("subject", msg.Subject()), zap.Error(err))
		}
		return
	}

	fn := ag.handlers.getCronjobHandler(job.Cronjob.Name)
	if fn == nil {
		if err := msg.Nak(); err != nil {
			ag.logger.Error("failed to nack unmarshal cron schedule msg", zap.String("subject", msg.Subject()), zap.Error(err))
		}
		return
	}

	if err := msg.InProgress(); err != nil {
		ag.logger.Error("failed to send in progress for cron schedule msg", zap.String("subject", msg.Subject()), zap.Error(err))
	}

	start := time.Now()
	err := fn(ag.ctx, job.Cronjob.Data)
	elapsed := time.Since(start)

	// Update timestamp
	now := timestamp.Now()
	job.Cronjob.Data.UpdatedAt = now

	if _, err := ag.js.PublishProto(ag.ctx, fmt.Sprintf("%s.%s", CronScheduleSubject, CronCompleteTopic), &cron.CronjobCompletedEvent{
		Name:    job.Cronjob.Name,
		Sucess:  err == nil,
		Elapsed: durationpb.New(elapsed),
		EndDate: now,

		Data: job.Cronjob.Data,
	}); err != nil {
		ag.logger.Error("failed to publish cron schedule completion msg", zap.String("subject", msg.Subject()), zap.Error(err))
		return
	}

	if err := msg.Ack(); err != nil {
		ag.logger.Error("failed to ack cron schedule msg", zap.String("subject", msg.Subject()), zap.Error(err))
		return
	}
}