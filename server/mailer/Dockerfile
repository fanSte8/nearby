FROM golang:latest AS build

WORKDIR /app

COPY ./mailer/ ./mailer/
COPY ./common/ ./common/

WORKDIR /app/mailer

RUN GOOS=linux go build -o=./build ./cmd/api

FROM ubuntu:latest

WORKDIR /app

COPY --from=build /app/mailer/build ./build

EXPOSE 3000

CMD ["./build"]