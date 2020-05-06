FROM golang:latest

WORKDIR /go/src/github.com/root87x/go-root

COPY . /go/src/github.com/root87x/go-root

ENTRYPOINT go run main.go