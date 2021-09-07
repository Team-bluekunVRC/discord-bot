// Package types contains definitions for the command interfaces
package types

import (
	"github.com/Team-bluekunVRC/discord-bot/ent"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// HandlerFunc is a function called when a command is ran
type HandlerFunc func(*discordgo.Session, *discordgo.InteractionCreate) error

// Command is a Discord slash command
type Command interface {
	// Info returns information about the provided command
	Info() *discordgo.ApplicationCommand

	// Handler is the actual handler logic for this command
	Handler(logrus.FieldLogger, *ent.Client) HandlerFunc
}
