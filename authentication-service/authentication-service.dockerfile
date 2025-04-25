FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=linux CGO_ENABLED=0 go build -o authApp ./cmd/api

RUN chmod +x /app/authApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/authApp /app

CMD ["/app/authApp"]