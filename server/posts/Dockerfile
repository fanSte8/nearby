FROM golang:latest AS build

WORKDIR /app

COPY ./posts/ ./posts/
COPY ./common/ ./common/

WORKDIR /app/posts

RUN GOOS=linux go build -o=./build ./cmd/api

EXPOSE 3000

CMD ["./build"]