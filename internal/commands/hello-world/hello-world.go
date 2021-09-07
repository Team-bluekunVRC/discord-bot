package helloworld

import (
	"github.com/Team-bluekunVRC/discord-bot/ent"
	"github.com/Team-bluekunVRC/discord-bot/internal/commands/types"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// Compile time validation that we match the interface
var _ types.Command = &Command{}

// Command is the struct containing the needed configuration
// for this bot command.
type Command struct{}

func (*Command) Info() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "hello-world",
		Description: "Returns hello world",
	}
}

func (*Command) Handler(log logrus.FieldLogger, _ *ent.Client) types.HandlerFunc {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
		log.Info("hello-world ran")

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello, world!",
			},
		})
		return nil
	}
}
