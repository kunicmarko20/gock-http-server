package api

import (
	"net/http"
)

type SetMockHandler struct {}

func (h *SetMockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
