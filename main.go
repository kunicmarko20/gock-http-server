package main

import (
	"./http/handler"
	"./http/handler/api"
	"./mock/matching"
	"./mock/repository"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	mockRepository := new(repository.InMemoryMockRepository)
	matcher := matching.MockMatcher{mockRepository}

	router := mux.NewRouter()
	router.PathPrefix("/mock/").Handler(handler.NewMockHandler(matcher))
	router.Path("/api/reset").Methods("POST").Handler(api.NewResetHandler(mockRepository))
	router.Path("/api/mock/{mock}").Methods("POST").Handler((api.NewSetMockHandler(mockRepository)))

	_ = http.ListenAndServe(":8080", router)
}
