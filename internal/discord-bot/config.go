package discordbot

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// HTTPPort is the port for this application to listen on
	HTTPPort int `env:"HTTP_PORT" envDefault:"8000"`

	// HTTPAddress is the address for this application to listen on, minus the port
	// use HTTPPort for that instead.
	HTTPAddress string `env:"HTTP_ADDRESS" envDefault:"0.0.0.0"`

	// BotToken is the Discord token for this bot
	BotToken string `env:"BOT_TOKEN"`
}

// LoadConfig reads the configuration from env variables and returns it
func LoadConfig() (*Config, error) {
	cfg := Config{}
	return &cfg, env.Parse(&cfg)
}
