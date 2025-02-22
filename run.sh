#!/bin/bash
set -e

echo "Initializing project..."

# Initialize Go modules
go mod init NiftyGoGo
go mod tidy

echo "Running main.go..."
go run src/main.go
