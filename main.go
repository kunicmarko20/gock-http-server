package main

import (
	"./http/handler"
	"./http/handler/api"
	"./mock"
	"./mock/matching"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	repository := new(mock.Repository)
	matcher := matching.MockMatcher{repository}

	router := mux.NewRouter()
	router.PathPrefix("/mock/").Handler(&handler.MockHandler{matcher})
	router.Path("/api/reset").Methods("POST").Handler(&api.ResetHandler{repository})
	router.Path("/api/mock/{mock}").Methods("POST").Handler(&api.SetMockHandler{repository})

	_ = http.ListenAndServe(":8080", router)
}
