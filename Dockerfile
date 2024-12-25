FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/configs ./configs

EXPOSE 8080
ENTRYPOINT ["./server", "-env", "prod"] 