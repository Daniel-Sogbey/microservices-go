FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=linux CGO_ENABLED=0 go build -o mailApp ./cmd/api

RUN chmod +x /app/mailApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app /app

CMD ["/app/mailApp"]
