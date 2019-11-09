package api

import (
	"../../../mock/repository"
	"net/http"
)

type ResetHandler struct {
	mockRepository repository.MockRepository
}

func NewResetHandler(mockRepository repository.MockRepository) *ResetHandler {
	return &ResetHandler{mockRepository}
}

func (h *ResetHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.mockRepository.Reset()

	writer.WriteHeader(204)
}
