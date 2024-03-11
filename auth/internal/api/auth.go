package api

import (
	"encoding/json"
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"

	api_client "async-arch/auth/api/generated"
	"async-arch/auth/internal/pkg/usecase/login_user"
	"async-arch/auth/internal/pkg/usecase/register_user"
)

type server struct {
	registerUserUsecase register_user.Usecase
	loginUserUsecase    login_user.Usecase
}

func NewServer(registerUserUsecase register_user.Usecase, loginUserUsecase login_user.Usecase) *server {
	return &server{registerUserUsecase: registerUserUsecase, loginUserUsecase: loginUserUsecase}
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

func (s *server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	in := api_client.UserRegister{}

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		internalError(w, err)
		return
	}

	user, err := s.registerUserUsecase.Run(r.Context(), register_user.In{
		Email:     string(in.Email),
		Role:      in.Role,
		Password:  in.Password,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	})
	if err != nil {
		badRequest(w, err)
		return
	}

	out := api_client.User{
		Email:     openapi_types.Email(user.Email),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}

	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		internalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) LoginUser(w http.ResponseWriter, r *http.Request) {
	in := api_client.UserLogin{}

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		internalError(w, err)
		return
	}

	token, err := s.loginUserUsecase.Run(r.Context(), login_user.In{
		Email:    string(in.Email),
		Password: in.Password,
	})
	if err != nil {
		badRequest(w, err)
		return
	}

	out := api_client.Token{Token: token}

	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		internalError(w, err)
		return
	}

	w.Header().Add("X-Bearer", token)
}
