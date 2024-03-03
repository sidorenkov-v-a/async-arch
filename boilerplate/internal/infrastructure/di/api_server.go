package di

import (
	"async-arch/boilerplate/pkg/api_server"
	"async-arch/boilerplate/pkg/config"
)

func NewAPIServer(config *config.APIServerConfig) *api_server.Server {
	return api_server.NewServer(config)
}
