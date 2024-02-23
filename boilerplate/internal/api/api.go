package api

import (
	"encoding/json"
	"net/http"

	"async-arch/boilerplate/internal/infrastructure/contract"
)

type ResponseError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type wrapper struct {
	log contract.Log
}

func NewWrapper(log contract.Log) *wrapper {
	return &wrapper{log: log}
}

func (w *wrapper) Handle(inner HTTPHandler) (string, http.HandlerFunc) {
	path := inner.GetPath()

	return path, func(response http.ResponseWriter, request *http.Request) {
		payload, err := inner.Handle(request)

		w.process(payload, err, path, response)
	}
}

func (w *wrapper) process(payload Payload, err error, path string, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.processError(err, path, response)
	}

	response.WriteHeader(http.StatusOK)

	_, _ = response.Write(payload)
}

func (w *wrapper) processError(err error, path string, response http.ResponseWriter) {
	w.log.WithError(err).Error(path)

	response.WriteHeader(http.StatusInternalServerError)

	respErr := ResponseError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}

	payload, _ := json.Marshal(respErr)
	_, _ = response.Write(payload)
}
