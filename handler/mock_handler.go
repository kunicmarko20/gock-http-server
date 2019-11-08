package handler

import (
	"net/http"
)

type MockHandler struct {}

func (h *MockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
