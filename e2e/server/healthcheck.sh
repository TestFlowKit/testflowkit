#!/bin/sh

# Health check script for the e2e server
# This script checks if the server is responding on port 3000

set -e

# Wait for the server to be ready
echo "Checking server health..."

# Try to connect to the server
if curl -f http://localhost:3000 > /dev/null 2>&1; then
    echo "Server is healthy and responding"
    exit 0
else
    echo "Server is not responding"
    exit 1
fi 