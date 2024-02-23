package api

import "net/http"

type Payload []byte

type HTTPHandler interface {
	GetPath() string
	Handle(request *http.Request) (Payload, error)
}
