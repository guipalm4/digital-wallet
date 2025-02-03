#!/usr/bin/env bash

echo "Creating topic 'transactions'..."
kafka-topics \
  --create \
  --topic transactions \
  --partitions 1 \
  --replication-factor 1 \
  --if-not-exists \
  --bootstrap-server kafka:29092

echo "Topic created successfully."

echo "Creating topic 'balances'..."
kafka-topics \
  --create \
  --topic balances \
  --partitions 1 \
  --replication-factor 1 \
  --if-not-exists \
  --bootstrap-server kafka:29092

echo "Topic created successfully."
