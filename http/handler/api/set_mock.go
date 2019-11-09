package api

import (
	"../../../mock"
	"net/http"
)

type SetMockHandler struct {
	MockRepository *mock.Repository
}

func (h *SetMockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
}
