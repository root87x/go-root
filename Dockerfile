FROM golang:latest

WORKDIR /go/src/github.com/root87x/go-root

COPY . /go/src/github.com/root87x/go-root

RUN go run ./main.go

EXPOSE 8080