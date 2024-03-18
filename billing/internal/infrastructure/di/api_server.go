package di

import (
	"async-arch/billing/pkg/api_server"
	"async-arch/billing/pkg/env"
)

func NewAPIServer(config *env.Server) *api_server.Server {
	return api_server.NewServer(config)
}
