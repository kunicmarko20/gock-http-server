FROM golang:1-alpine

ADD . /app/

WORKDIR /app

ENV GOPATH /app

RUN apk add --no-cache git && \
  go build main.go && \
  apk del git

ENTRYPOINT /app/app
