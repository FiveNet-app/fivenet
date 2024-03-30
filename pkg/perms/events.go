package perms

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/galexrt/fivenet/pkg/events"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

const (
	BaseSubject events.Subject = "perms"

	RoleCreatedSubject    events.Type = "roleperm.create"
	RolePermUpdateSubject events.Type = "roleperm.update"
	RoleDeletedSubject    events.Type = "roleperm.delete"
	RoleAttrUpdateSubject events.Type = "roleattr.update"
	JobAttrUpdateSubject  events.Type = "jobattr.update"
)

type RoleIDEvent struct {
	RoleID uint64
}

type JobAttrUpdateEvent struct {
	Job string
}

func (p *Perms) registerSubscriptions(ctx context.Context, c context.Context) error {
	cfg := jetstream.StreamConfig{
		Name:        "PERMS",
		Description: "Perms system events",
		Retention:   jetstream.InterestPolicy,
		Subjects:    []string{fmt.Sprintf("%s.>", BaseSubject)},
		Discard:     jetstream.DiscardOld,
		MaxAge:      15 * time.Second,
		Storage:     jetstream.MemoryStorage,
	}

	if _, err := p.js.CreateOrUpdateStream(ctx, cfg); err != nil {
		return err
	}

	consumer, err := p.js.CreateConsumer(ctx, cfg.Name, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverNewPolicy,
		FilterSubject: fmt.Sprintf("%s.>", BaseSubject),
	})
	if err != nil {
		return err
	}

	if p.jsCons != nil {
		p.jsCons.Stop()
		p.jsCons = nil
	}

	p.jsCons, err = consumer.Consume(p.handleMessageFunc(c), p.js.ConsumeErrHandlerWithRestart(c, p.logger, p.registerSubscriptions))
	if err != nil {
		return err
	}

	return nil
}

func (p *Perms) handleMessageFunc(ctx context.Context) jetstream.MessageHandler {
	return func(msg jetstream.Msg) {
		if err := msg.Ack(); err != nil {
			p.logger.Error("failed to ack message", zap.Error(err))
		}

		p.logger.Debug("received event message", zap.String("subject", msg.Subject()))

		switch events.Type(strings.TrimPrefix(msg.Subject(), string(BaseSubject)+".")) {
		case RoleCreatedSubject:
			fallthrough
		case RolePermUpdateSubject:
			event := &RoleIDEvent{}
			if err := json.Unmarshal(msg.Data(), event); err != nil {
				p.logger.Error("failed to unmarshal message event data", zap.Error(err))
				return
			}

			if err := p.loadRoles(ctx, event.RoleID); err != nil {
				p.logger.Error("failed to load role for role data load", zap.Error(err))
				return
			}

			if err := p.loadRolePermissions(ctx, event.RoleID); err != nil {
				p.logger.Error("failed to load updated role permissions", zap.Error(err))
				return
			}

		case RoleAttrUpdateSubject:
			event := &RoleIDEvent{}
			if err := json.Unmarshal(msg.Data(), event); err != nil {
				p.logger.Error("failed to unmarshal message event data", zap.Error(err))
				return
			}

			if err := p.loadRoles(ctx, event.RoleID); err != nil {
				p.logger.Error("failed to load role for role data load", zap.Error(err))
				return
			}

			if err := p.loadRoleAttributes(ctx, event.RoleID); err != nil {
				p.logger.Error("failed to load updated role permissions", zap.Error(err))
				return
			}

		case RoleDeletedSubject:
			event := &RoleIDEvent{}
			if err := json.Unmarshal(msg.Data(), event); err != nil {
				p.logger.Error("failed to unmarshal message event data", zap.Error(err))
				return
			}

			role, err := p.GetRole(ctx, event.RoleID)
			if err != nil {
				p.logger.Error("failed to deleted role message event data", zap.Error(err))
				return
			}
			// Remove role from local state
			p.deleteRole(role)

		case JobAttrUpdateSubject:
			event := &JobAttrUpdateEvent{}
			if err := json.Unmarshal(msg.Data(), event); err != nil {
				p.logger.Error("failed to unmarshal message event data", zap.Error(err))
				return
			}

			if err := p.loadJobAttrs(ctx, event.Job); err != nil {
				p.logger.Error("failed to update job attributes", zap.Error(err))
				return
			}

		default:
			p.logger.Error("unknown type of perms events message received")
		}
	}
}

func (p *Perms) publishMessage(ctx context.Context, subj events.Type, data any) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err := p.js.Publish(ctx, fmt.Sprintf("%s.%s", BaseSubject, subj), out); err != nil {
		return err
	}

	return nil
}
