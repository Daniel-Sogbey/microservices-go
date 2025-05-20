FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOOS=linux CGO_ENABLED=0 go build -o listenerApp .

RUN chmod +x /app/listenerApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/listenerApp /app

CMD ["/app/listenerApp"]