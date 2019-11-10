package http

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
)

const ContextKeyRequestID string = "request_id"

func AddRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		id := uuid.New()
		ctx = context.WithValue(ctx, ContextKeyRequestID, id.String())
		request = request.WithContext(ctx)

		next.ServeHTTP(writer, request)
	})
}

func HandleRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Request received.", map[string]string{
			"request_id": request.Context().Value(ContextKeyRequestID).(string),
			"method":     request.Method,
			"path":       request.URL.Path,
		})

		next.ServeHTTP(writer, request)
	})
}
