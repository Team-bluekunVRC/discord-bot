// Package discordbot implements the core discord bot application logic.
package discordbot

import (
	"context"

	"github.com/Team-bluekunVRC/discord-bot/internal/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tritonmedia/pkg/service"
)

// NewBotService creates a new discord bot
func NewBotService(conf *Config) *BotService {
	return &BotService{conf: conf}
}

type BotService struct {
	service.Service
	conf *Config
}

func (b *BotService) Run(ctx context.Context, log logrus.FieldLogger) error {
	log.Info("starting discord bot")
	s, err := discordgo.New("Bot " + b.conf.BotToken)
	if err != nil {
		return errors.Wrap(err, "failed to create bot")
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info("discord bot is ready")
	})

	err = s.Open()
	if err != nil {
		return errors.Wrap(err, "failed to open connection to discord")
	}

	err = commands.Register(ctx, s, log)
	if err != nil {
		return errors.Wrap(err, "failed to register commands")
	}

	<-ctx.Done()

	log.WithError(ctx.Err()).Info("shutting down bot due to error")
	s.Close()

	return nil
}
