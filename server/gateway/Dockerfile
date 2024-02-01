FROM golang:latest AS build

WORKDIR /app

COPY ./gateway/ ./gateway/

WORKDIR /app/gateway

RUN GOOS=linux go build -o=./build ./cmd/api

FROM centos:latest

WORKDIR /app

COPY --from=build /app/gateway/build ./build

EXPOSE 3000

CMD ["./build"]