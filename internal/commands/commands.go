// Package commands handles implementing the command interface for this discord bot
package commands

import (
	"context"

	"github.com/Team-bluekunVRC/discord-bot/ent"
	helloworld "github.com/Team-bluekunVRC/discord-bot/internal/commands/hello-world"
	"github.com/Team-bluekunVRC/discord-bot/internal/commands/types"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// commands are the built-in commands
var commands = []types.Command{&helloworld.Command{}}

// Register registers the built-in commands onto the current discord
// bot
func Register(ctx context.Context, s *discordgo.Session, log logrus.FieldLogger, db *ent.Client) error {
	commandTable := make(map[string]types.HandlerFunc)
	for _, c := range commands {
		command := c.Info().Name
		log.WithField("command.name", command).Info("building command")

		// map the handler to the actual command table
		commandTable[command] = c.Handler(log.WithField("command.name", command), db)

		// register the command with discord
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", c.Info())
		if err != nil {
			log.WithError(err).Error("failed to register command")
			return err
		}
	}

	// register the global command handler (router)
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		command := i.ApplicationCommandData().Name
		if h, ok := commandTable[command]; ok {
			err := h(s, i)
			if err != nil {
				log.WithError(err).Error("command failed to execute")
			}

			return
		}

		log.WithField("command.name", command).Info("skipping unknown command")
	})

	return nil
}
