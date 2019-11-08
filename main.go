package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

import "./handler"
import "./handler/api"

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/mock/").Handler(new(handler.MockHandler))
	r.Path("/api/reset").Methods("POST").Handler(new(api.ResetHandler))
	r.Path("/api/mock/{mock}").Methods("POST").Handler(new(api.SetMockHandler))

	_ = http.ListenAndServe(":8080", r)
}
