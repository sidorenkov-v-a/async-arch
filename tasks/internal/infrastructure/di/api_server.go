package di

import (
	"async-arch/tasks/pkg/api_server"
	"async-arch/tasks/pkg/env"
)

func NewAPIServer(config *env.Server) *api_server.Server {
	return api_server.NewServer(config)
}
