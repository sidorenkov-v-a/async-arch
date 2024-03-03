package api_server

import (
	"net/http"

	"github.com/gorilla/mux"

	"async-arch/task_tracker/pkg/env"
)

type Server struct {
	config *env.Server
}

func NewServer(config *env.Server) *Server {
	return &Server{config: config}
}

func (s *Server) Run(router *mux.Router) error {
	return http.ListenAndServe(s.config.BindAddr, router)
}
