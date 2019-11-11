FROM golang:1-alpine

ADD . /app/

WORKDIR /app

ENV GOPATH /go

RUN apk add --no-cache git && \
  go build -o gock-http-server && \
  apk del git

ENTRYPOINT /app/gock-http-server
