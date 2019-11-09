package handler

import (
	"../../mock/matching"
	"net/http"
)

type MockHandler struct {
	MockMatcher matching.MockMatcher
}

func (h *MockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	mock, err := h.MockMatcher.Match(request)

	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	for header, value := range mock.Response.Headers {
		writer.Header().Set(header, value)
	}

	writer.WriteHeader(mock.Response.StatusCode)
	_, _ = writer.Write([]byte(mock.Response.Body))
}
