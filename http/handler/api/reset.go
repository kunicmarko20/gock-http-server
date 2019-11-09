package api

import (
	"../../../mock"
	"net/http"
)

type ResetHandler struct {
	MockRepository *mock.Repository
}
func (h *ResetHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.MockRepository.Reset()
}
