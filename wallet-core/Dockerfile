# Etapa 1: Build
FROM golang:1.23.4 AS builder
WORKDIR /app
RUN apt-get update && apt-get install -y librdkafka-dev
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o wallet-core ./cmd/walletcore/main.go

# Etapa 2: Runtime
FROM golang:1.23.4
WORKDIR /app
COPY --from=builder /app/wallet-core /app/wallet-core
COPY --from=builder /app/migrations /app/migrations
COPY ./wait-for-it.sh /usr/bin/wait-for-it
RUN chmod +x /usr/bin/wait-for-it

EXPOSE 8080

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]