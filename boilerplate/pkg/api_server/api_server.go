package api_server

import (
	"net/http"

	"github.com/gorilla/mux"

	"async-arch/boilerplate/pkg/config"
)

type Server struct {
	config *config.APIServerConfig
}

func NewServer(config *config.APIServerConfig) *Server {
	return &Server{config: config}
}

func (s *Server) Run(router *mux.Router) error {
	return http.ListenAndServe(s.config.BindAddress, router)
}
