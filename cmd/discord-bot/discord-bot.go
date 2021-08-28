package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	discordbot "github.com/Team-bluekunVRC/discord-bot/internal/discord-bot"
	"github.com/getoutreach/gobox/pkg/app"
	"github.com/sirupsen/logrus"
	"github.com/tritonmedia/pkg/service"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	log := logrus.New().WithContext(ctx)

	//nolint:gocritic importShadow
	app := cli.App{
		Name:    "discord-bot",
		Version: app.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log-format",
				Usage:   "set the formatter for logs",
				EnvVars: []string{"LOG_FORMAT"},
				Value:   "JSON",
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		logFormat := c.String("log-format")
		switch strings.ToLower(logFormat) {
		case "json":
			logrus.SetFormatter(&logrus.JSONFormatter{})
		case "text":
			logrus.SetFormatter(&logrus.TextFormatter{})
		default:
			return fmt.Errorf("unknown log format: %s", logFormat)
		}

		// update our logger's formatter, since we created before
		// we set it
		log.Logger.Formatter = logrus.StandardLogger().Formatter

		// load the configuration
		conf, err := discordbot.LoadConfig()
		if err != nil {
			log.WithError(err).Fatal("failed to load configuration")
		}

		// start the service runner, which handles context cancellation
		// and threading
		r := service.NewServiceRunner(ctx, []service.Service{discordbot.NewBotService(conf)})
		sigC := make(chan os.Signal)

		// listen for signals that we want to cancel on, and cancel
		// the context if one is passed
		signal.Notify(sigC, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-sigC
			cancel()
		}()

		return r.Run(ctx, log)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
