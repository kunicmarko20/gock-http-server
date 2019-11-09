package api

import (
	"../../../mock/repository"
	"net/http"
)

type ResetHandler struct {
	MockRepository repository.MockRepository
}
func (h *ResetHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.MockRepository.Reset()
}
