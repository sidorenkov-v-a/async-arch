package config

import (
	"github.com/BurntSushi/toml"
)

const configPath = "app.toml"

type APIServerConfig struct {
	BindAddress string `toml:"bind_addr"`
}

type Config struct {
	APIServer APIServerConfig `toml:"api_server"`
}

func NewConfig() (*Config, error) {
	var config Config

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
