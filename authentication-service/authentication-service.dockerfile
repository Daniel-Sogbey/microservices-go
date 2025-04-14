FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o authApp ./cmd/api

RUN chmod +x authApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/authApp /app

CMD ["/app/authApp"]