package main

import (
	"fmt"
	"github.com/gorilla/mux"
	ghttp "github.com/kunicmarko20/gock-http-server/http"
	"github.com/kunicmarko20/gock-http-server/http/handler"
	"github.com/kunicmarko20/gock-http-server/http/handler/api"
	"github.com/kunicmarko20/gock-http-server/mock/matching"
	"github.com/kunicmarko20/gock-http-server/mock/repository"
	"github.com/kunicmarko20/gock-http-server/propertyaccess"
	"net/http"
	"os"
)

func main() {
	port, present := os.LookupEnv("BASE_PORT")

	if !present {
		fmt.Println("The environment variable \"BASE_PORT\" must be provided.")
		return
	}
	_ = http.ListenAndServe(":"+port, newRouter())
}

func newRouter() *mux.Router {
	mockRepository := repository.NewInMemoryMockRepository()
	matcher := matching.NewRepositorySourcedMockMatcher(mockRepository)
	propertyAccessor := propertyaccess.NewPropertyAccessor()

	router := mux.NewRouter()
	router.Use(ghttp.AddRequestID, ghttp.HandleRequest)
	router.PathPrefix("/mock/").Handler(handler.NewMockHandler(matcher))
	router.Path("/api/reset").Methods("POST").Handler(api.NewResetHandler(mockRepository))
	router.Path("/api/mock/{mock}").Methods("POST").Handler(api.NewSetMockHandler(mockRepository, propertyAccessor))

	return router
}
