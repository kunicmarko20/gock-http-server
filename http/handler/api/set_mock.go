package api

import (
	ghttp "../../../http"
	"../../../mock/repository"
	"../../../propertyaccess"
	"log"
	"net/http"
)

type SetMockHandler struct {
	mockRepository   repository.MockRepository
	propertyAccessor *propertyaccess.PropertyAccessor
}

func NewSetMockHandler(
	mockRepository repository.MockRepository,
	propertyAccessor *propertyaccess.PropertyAccessor,
) *SetMockHandler {
	return &SetMockHandler{mockRepository, propertyAccessor}
}

func (h *SetMockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	mock, err := ghttp.NewPayloadToMockTransformer(h.propertyAccessor).FromRequest(request)

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
