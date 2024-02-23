package index

import (
	"net/http"

	"async-arch/boilerplate/internal/api"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) GetPath() string {
	return "/"
}

func (h *handler) Handle(request *http.Request) (api.Payload, error) {

	return nil, nil
}
