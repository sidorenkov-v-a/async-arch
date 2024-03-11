package api

import (
	"encoding/json"
	"net/http"

	api_client "async-arch/task_tracker/api/generated"
)

type server struct {
}

func NewServer() *server {
	return &server{}
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	out := api_client.Error{
		Message: err.Error(),
	}

	_ = json.NewEncoder(w).Encode(out)
}

func internalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	out := api_client.Error{
		Message: err.Error(),
	}

	_ = json.NewEncoder(w).Encode(out)
}

func (s *server) CreateTask(w http.ResponseWriter, r *http.Request) {

}
