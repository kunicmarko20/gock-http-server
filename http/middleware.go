package http

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type RequestID string

const ContextKeyRequestID RequestID = "request_id"

func AddRequestID(next http.Handler) http.Handler {
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
