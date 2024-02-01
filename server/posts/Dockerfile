FROM golang:latest AS build

WORKDIR /app

COPY ./posts/ ./posts/
COPY ./common/ ./common/

WORKDIR /app/posts

RUN GOOS=linux go build -o=./build ./cmd/api

FROM ubuntu:latest

WORKDIR /app

COPY --from=build /app/posts/build ./build

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 3000

CMD ["./build"]