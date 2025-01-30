#!/bin/bash
set -e

echo "Waiting for MySQL to be ready..."
until mysqladmin ping -h "mysql" --silent; do
    echo "MySQL is unavailable - sleeping"
    sleep 2
done

echo "Running database migrations..."
migrate -path /app/migrations -database "mysql://root:root@tcp(mysql:3306)/database" up

echo "Starting application..."
exec /app/main
