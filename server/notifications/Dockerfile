FROM golang:latest AS build

WORKDIR /app

COPY ./notifications/ ./notifications/
COPY ./common/ ./common/

WORKDIR /app/notifications

RUN GOOS=linux go build -o=./build ./cmd/api

EXPOSE 3000

CMD ["./build"]