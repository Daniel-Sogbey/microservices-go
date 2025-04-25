#base go image
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=linux CGO_ENABLED=0 go build -o loggerApp ./cmd/api

RUN chmod +x /app/loggerApp

#build tiny docker image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/loggerApp /app

CMD ["/app/loggerApp"]