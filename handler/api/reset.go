package api

import (
	"net/http"
)

type ResetHandler struct {}

func (h *ResetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
