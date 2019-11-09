package api

import (
	ghttp "../../../http"
	"../../../mock/repository"
	"net/http"
)

type SetMockHandler struct {
	mockRepository repository.MockRepository
}

func NewSetMockHandler(mockRepository repository.MockRepository) *SetMockHandler {
	return &SetMockHandler{mockRepository}
}

func (h *SetMockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	mock, err := new(ghttp.PayloadFactory).FromRequest(request)

	if err != nil {
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	h.mockRepository.Set(mock)
	writer.WriteHeader(204)
}
