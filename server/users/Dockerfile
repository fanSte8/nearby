FROM golang:latest AS build

WORKDIR /app

COPY ./users/ ./users/
COPY ./common/ ./common/

WORKDIR /app/users

RUN GOOS=linux go build -o=./build ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/users/build ./build

EXPOSE 3000

CMD ["./build"]