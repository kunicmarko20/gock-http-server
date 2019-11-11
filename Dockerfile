FROM golang:1-alpine

ADD . /app/

WORKDIR /app

RUN apk add --no-cache git && \
  go get -u github.com/gorilla/mux github.com/google/uuid github.com/pkg/errors && \
  go build main.go && \
  apk del git

ENV GOPATH /app
ENTRYPOINT /app/app
