package cmd

import (
	"github.com/fivenet-app/fivenet/pkg/croner"
	"github.com/fivenet-app/fivenet/pkg/discord"
	"go.uber.org/fx"
)

type DiscordCmd struct {
	ModuleCronAgent bool `help:"Run the cron agent." default:"false"`
}

func (c *DiscordCmd) Run(ctx *Context) error {
	fxOpts := getFxBaseOpts(Cli.StartTimeout, true)
	fxOpts = append(fxOpts, fx.Invoke(func(*discord.Bot) {}))

	if c.ModuleCronAgent {
		fxOpts = append(fxOpts, fx.Invoke(func(*croner.Agent) {}))
	}

	app := fx.New(fxOpts...)
	app.Run()

	return nil
}