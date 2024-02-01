FROM golang:latest AS build

WORKDIR /app

COPY ./notifications/ ./notifications/
COPY ./common/ ./common/

WORKDIR /app/notifications

RUN GOOS=linux go build -o=./build ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/notifications/build ./build

EXPOSE 3000

CMD ["./build"]