FROM golang:alpine AS build

WORKDIR /go/service

COPY . .

RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go build -o myapp

ENTRYPOINT ["/go/service/myapp"]