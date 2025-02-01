#!/usr/bin/env bash

echo "Creating topic 'transactions'..."
docker-compose exec kafka \
  kafka-topics \
    --create \
    --topic transactions \
    --partitions 1 \
    --replication-factor 1 \
    --if-not-exists \
    --bootstrap-server kafka:29092

echo "Topic created successfully."