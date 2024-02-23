package di

import (
	"async-arch/boilerplate/internal/infrastructure/di/config"
)

func NewConfig() (*config.Config, error) {
	return config.NewConfig()
}
