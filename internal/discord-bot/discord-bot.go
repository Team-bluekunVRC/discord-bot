// Package discordbot implements the core discord bot application logic.
package discordbot

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Team-bluekunVRC/discord-bot/ent"
	"github.com/Team-bluekunVRC/discord-bot/internal/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tritonmedia/pkg/service"

	// Used by ent
	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewBotService creates a new discord bot
func NewBotService(conf *Config) *BotService {
	return &BotService{conf: conf}
}

type BotService struct {
	service.Service
	conf *Config
}

func createDBConnection(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func (b *BotService) Run(ctx context.Context, log logrus.FieldLogger) error {
	log.Info("creating postgres connection")
	db, err := createDBConnection(b.conf.DatabaseURL)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()

	if err := db.Schema.Create(ctx); err != nil {
		return errors.Wrap(err, "failed to run migrations")
	}

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

	err = commands.Register(ctx, s, log, db)
	if err != nil {
		return errors.Wrap(err, "failed to register commands")
	}

	<-ctx.Done()

	log.WithError(ctx.Err()).Info("shutting down bot due to error")
	s.Close()

	return nil
}
