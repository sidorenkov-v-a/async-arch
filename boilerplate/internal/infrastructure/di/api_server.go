package di

import (
	"async-arch/boilerplate/internal/infrastructure/di/api_server"
	"async-arch/boilerplate/internal/infrastructure/di/config"
)

func NewAPIServer(config *config.APIServerConfig) *api_server.Server {
	return api_server.NewServer(config)
}
