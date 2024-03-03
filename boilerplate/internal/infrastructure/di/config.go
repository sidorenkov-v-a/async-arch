package di

import (
	"async-arch/boilerplate/pkg/config"
)

func NewConfig() (*config.Config, error) {
	return config.NewConfig()
}
