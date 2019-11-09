package api

import (
	"../../../mock/repository"
	"net/http"
)

type SetMockHandler struct {
	MockRepository repository.MockRepository
}

func (h *SetMockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}
