#!/bin/bash
set -e

echo "Installing dependencies..."
go mod tidy

echo "Running main.go..."
go run src/main.go
