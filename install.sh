#!/bin/bash
set -e

echo "Installing dependencies..."

# Initialize Go module if not already initialized
if [ ! -f "go.mod" ]; then
    go mod init NiftyGoGo
fi

# Download and tidy dependencies
go mod tidy

echo "Installation complete!"
