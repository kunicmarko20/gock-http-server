FROM golang:1-alpine

ADD . /app/

WORKDIR /app

RUN apk add --no-cache git && \
  go get -u github.com/gorilla/mux github.com/google/uuid && \
  apk del git

ENV GOPATH /app
ENV BASE_PORT 3000
ENTRYPOINT /usr/local/go/bin/go run main.go
