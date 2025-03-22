FROM golang:1.23.3-alpine3.19

WORKDIR /src/app

RUN go install github.com/cosmtrek/air@v1.45.0

COPY . .

RUN go mod tidy

