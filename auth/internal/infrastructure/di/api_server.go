package di

import (
	"async-arch/auth/pkg/api_server"
	"async-arch/auth/pkg/env"
)

func NewAPIServer(config *env.Server) *api_server.Server {
	return api_server.NewServer(config)
}
