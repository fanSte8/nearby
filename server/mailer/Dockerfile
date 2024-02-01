FROM golang:latest AS build

WORKDIR /app

COPY ./mailer/ ./mailer/
COPY ./common/ ./common/

WORKDIR /app/mailer

RUN GOOS=linux go build -o=./build ./cmd/api

EXPOSE 3000

CMD ["./build"]