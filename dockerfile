# Etapa de build
FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o app cmd/app/main.go

# Etapa final (leve)
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]