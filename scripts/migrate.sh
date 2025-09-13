#!/bin/bash

# Database migration script for fitness app
# Usage: ./scripts/migrate.sh [up|down|create] [service_name]

set -e

SERVICE_NAME=${2:-"user-service"}
MIGRATION_DIR="services/$SERVICE_NAME/migrations"

case $1 in
  "up")
    echo "Running migrations up for $SERVICE_NAME..."
    migrate -path $MIGRATION_DIR -database "postgres://postgres:password@localhost:5432/fitness_app?sslmode=disable" up
    ;;
  "down")
    echo "Running migrations down for $SERVICE_NAME..."
    migrate -path $MIGRATION_DIR -database "postgres://postgres:password@localhost:5432/fitness_app?sslmode=disable" down
    ;;
  "create")
    if [ -z "$3" ]; then
      echo "Usage: ./scripts/migrate.sh create [service_name] [migration_name]"
      exit 1
    fi
    echo "Creating migration $3 for $SERVICE_NAME..."
    migrate create -ext sql -dir $MIGRATION_DIR -seq $3
    ;;
  *)
    echo "Usage: ./scripts/migrate.sh [up|down|create] [service_name] [migration_name]"
    echo "Examples:"
    echo "  ./scripts/migrate.sh up user-service"
    echo "  ./scripts/migrate.sh down order-service"
    echo "  ./scripts/migrate.sh create user-service add_user_table"
    exit 1
    ;;
esac
