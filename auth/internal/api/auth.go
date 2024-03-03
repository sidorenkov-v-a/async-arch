package api

import (
	"encoding/json"
	"net/http"

	api_client "async-arch/auth/api/generated"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := api_client.User{
		Email:     "test@test.com",
		FirstName: "first_name",
		LastName:  "last_name",
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
