package di

import (
	"async-arch/boilerplate/pkg/api_server"
	"async-arch/boilerplate/pkg/env"
)

func NewAPIServer(config *env.Server) *api_server.Server {
	return api_server.NewServer(config)
}
