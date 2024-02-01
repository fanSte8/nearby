FROM golang:latest AS build

WORKDIR /app

COPY ./gateway/ ./gateway/

WORKDIR /app/gateway

RUN GOOS=linux go build -o=./build ./cmd/api

FROM ubuntu:latest

WORKDIR /app

COPY --from=build /app/gateway/build ./build

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 3000

CMD ["./build"]