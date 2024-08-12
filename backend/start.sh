#!/bin/bash
set -e

echo "Starting wait-for-it script"
/usr/local/bin/wait-for-it.sh postgres:5432
echo "PostgreSQL is up"

echo "Running Goose migrations"
goose -dir /app/sql/schema postgres "$GOOSE_URL" up
echo "Migrations complete"

echo "Starting main application"
./main
echo "Main application exited"
