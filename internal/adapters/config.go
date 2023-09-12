package adapters

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DiscordClientID string	`toml:"discord-client-id" env-required`
	DiscordClientSecret string `toml:"discord-client-secret" env-required`
	DiscordToken string `toml:"discord-token" env-required`
}

func NewConfig() (*Config, error) {
	var config Config

	err := cleanenv.ReadConfig("dev-config.toml", &config)
	if err != nil {
		err = cleanenv.ReadEnv(&config)
		if err != nil {
			return nil, err
		}
	}

	return &config, nil
}