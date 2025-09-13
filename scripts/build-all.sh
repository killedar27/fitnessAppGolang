#!/bin/bash

# Build all services script
# Usage: ./scripts/build-all.sh [clean]

set -e

echo "Building all services..."

# Clean build if requested
if [ "$1" = "clean" ]; then
    echo "Cleaning previous builds..."
    go clean -cache
    go mod download
fi

# Build user-service
echo "Building user-service..."
cd services/user-service
go build -o ../../bin/user-service ./cmd/server
cd ../..

# Build order-service
echo "Building order-service..."
cd services/order-service
go build -o ../../bin/order-service ./cmd/server
cd ../..

echo "All services built successfully!"
echo "Binaries are available in the bin/ directory"
