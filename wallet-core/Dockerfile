# Etapa 1: Build
FROM golang:1.23.4 AS builder
WORKDIR /app
RUN apt-get update && apt-get install -y librdkafka-dev
#COPY go.mod go.sum ./
#RUN go mod download

COPY . .
RUN go mod tidy
RUN go build -o wallet-core ./cmd/walletcore/main.go

# Etapa 2: Runtime
FROM golang:1.23.4
WORKDIR /app
COPY --from=builder /app/wallet-core /app/wallet-core
COPY --from=builder /app/migrations /app/migrations

EXPOSE 8080
CMD ["/app/wallet-core"]
