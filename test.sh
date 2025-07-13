#!/bin/bash

# Test script for denote-contacts

echo "Testing denote-contacts with sample data..."
echo "Using contacts directory: ./contacts-data"
echo ""

# Set the contacts directory to our test data
export DENOTE_CONTACTS_DIR="$(pwd)/contacts-data"

# Run the app
go run cmd/main.go