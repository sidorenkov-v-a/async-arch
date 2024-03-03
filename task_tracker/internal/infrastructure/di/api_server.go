package di

import (
	"async-arch/task_tracker/pkg/api_server"
	"async-arch/task_tracker/pkg/env"
)

func NewAPIServer(config *env.Server) *api_server.Server {
	return api_server.NewServer(config)
}
