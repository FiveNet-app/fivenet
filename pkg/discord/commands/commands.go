package commands

import (
	"fmt"
	"slices"

	"github.com/bwmarrin/discordgo"
	"github.com/galexrt/fivenet/pkg/config"
	"go.uber.org/zap"
)

type CommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate)

type CommandFactory = func(cfg *config.Config) (*discordgo.ApplicationCommand, CommandHandler, error)

var (
	CommandsFactories = map[string]CommandFactory{}
	Commands          = map[string]*discordgo.ApplicationCommand{}
	CommandHandlers   = map[string]CommandHandler{}
)

type Cmds struct {
	logger *zap.Logger

	discord *discordgo.Session
}

func New(logger *zap.Logger, s *discordgo.Session, cfg *config.Config) (*Cmds, error) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	c := &Cmds{
		logger:  logger,
		discord: s,
	}

	for _, factory := range CommandsFactories {
		command, handler, err := factory(cfg)
		if err != nil {
			return nil, err
		}

		Commands[command.Name] = command
		CommandHandlers[command.Name] = handler
	}

	return c, nil
}

func (c *Cmds) Register(guild *discordgo.Guild) error {
	cmds, err := c.discord.ApplicationCommands(c.discord.State.User.ID, guild.ID)
	if err != nil {
		return err
	}

	c.logger.Info("registering commands", zap.Int("count", len(Commands)))
	for _, command := range Commands {
		if slices.ContainsFunc(cmds, func(cmd *discordgo.ApplicationCommand) bool {
			return cmd.Name == command.Name
		}) {
			c.logger.Debug(fmt.Sprintf("command '%v' already registered", command.Name), zap.String("discord_guild_id", guild.ID))
			continue
		}

		if _, err := c.discord.ApplicationCommandCreate(c.discord.State.User.ID, guild.ID, command); err != nil {
			return fmt.Errorf("cannot create '%v' command for guild '%s'. %w", command.Name, guild.ID, err)
		}
	}

	if err := c.removeDuplicateDispatches(guild); err != nil {
		return err
	}

	return nil
}

func (c *Cmds) removeDuplicateDispatches(guild *discordgo.Guild) error {
	cmds, err := c.discord.ApplicationCommands(c.discord.State.User.ID, guild.ID)
	if err != nil {
		return err
	}

	// Remove duplicate commands
	duplicates := GetDuplicateCommands(cmds)
	c.logger.Info("removing duplicate commands", zap.Int("duplicates", len(duplicates)))
	for _, command := range duplicates {
		if err := c.discord.ApplicationCommandDelete(c.discord.State.User.ID, guild.ID, command.ID); err != nil {
			return fmt.Errorf("cannot delete '%v' duplicate command for guild '%s'. %w", command.Name, guild.ID, err)
		}
	}

	return nil
}

func GetDuplicateCommands(in []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	allKeys := make(map[string]bool)
	list := []*discordgo.ApplicationCommand{}

	for _, item := range in {
		if _, value := allKeys[item.Name]; !value {
			allKeys[item.Name] = true
		} else {
			list = append(list, item)
		}
	}

	return list
}
