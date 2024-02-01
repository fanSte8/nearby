FROM golang:latest AS build

WORKDIR /app

COPY ./gateway/ ./gateway/

WORKDIR /app/gateway

RUN GOOS=linux go build -o=./build ./cmd/api

EXPOSE 3000

CMD ["./build"]