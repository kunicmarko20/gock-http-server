package handler

import (
	ghttp "../../http"
	"../../mock/matching"
	"encoding/json"
	"log"
	"net/http"
)

type MockHandler struct {
	mockMatcher matching.MockMatcher
}

func NewMockHandler(mockMatcher matching.MockMatcher) *MockHandler {
	return &MockHandler{mockMatcher}
}

func (h *MockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	requestHeaders, _ := json.Marshal(request.Header)

	log.Println("Trying to match request to mock.", map[string]string{
		"headers":    string(requestHeaders),
		"request_id": request.Context().Value(ghttp.ContextKeyRequestID).(string),
		"method":     request.Method,
		"path":       request.URL.Path,
	})

	mock, err := h.mockMatcher.Match(request)

	if err != nil {
		log.Println("Request did not match a mock.", map[string]string{
			"request_id": request.Context().Value(ghttp.ContextKeyRequestID).(string),
		})

		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}

	log.Println("Request matched a mock.", map[string]string{
		"request_id": request.Context().Value(ghttp.ContextKeyRequestID).(string),
		"mock":       mock.Name(),
	})

	for header, value := range mock.Response().Headers {
		writer.Header().Set(header, value)
	}

	writer.WriteHeader(mock.Response().StatusCode)
	_, _ = writer.Write([]byte(mock.Response().Content))
}
