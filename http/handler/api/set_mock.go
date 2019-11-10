package api

import (
	ghttp "../../../http"
	"../../../mock/repository"
	"log"
	"net/http"
)

type SetMockHandler struct {
	mockRepository repository.MockRepository
}

func NewSetMockHandler(mockRepository repository.MockRepository) *SetMockHandler {
	return &SetMockHandler{mockRepository}
}

func (h *SetMockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	mock, err := new(ghttp.PayloadToMockTransformer).FromRequest(request)

	if err != nil {
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	h.mockRepository.Set(mock)

	log.Println("Updated mock.", map[string]string{
		"mock": mock.Name(),
	})

	writer.WriteHeader(204)
}
