#!/bin/bash

# Aguarda o MySQL subir (mysql:3306)
echo "Waiting for MySQL..."
/usr/bin/wait-for-it mysql:3306 -t 30 -- echo "MySQL is up!"

# Agora podemos iniciar o wallet-core
echo "Starting wallet-core..."
exec /app/wallet-core
