package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DiscordBotToken string `envconfig:"DISCORD_BOT_TOKEN" required:"true"`
	DiscordGuildID  string `envconfig:"DISCORD_GUILD_ID" required:"true"`
	DiscordAppID    string `envconfig:"DISCORD_APP_ID" required:"true"`
}

func NewConfig() *Config {
	var config *Config

	envconfig.MustProcess("TARDY_TOAD", config)

	return config
}
