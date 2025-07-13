#!/bin/bash

# Test script for denote-contacts with task creation
# This uses a separate test directory to avoid messing with your real contacts

echo "Testing denote-contacts with task creation..."

# Create test directory if it doesn't exist
TEST_DIR="./test-tasks-data"
if [ ! -d "$TEST_DIR" ]; then
    echo "Creating test directory: $TEST_DIR"
    mkdir -p "$TEST_DIR"
    
    # Copy a few sample contacts for testing
    if [ -d "./contacts-data" ]; then
        echo "Copying sample contacts..."
        cp ./contacts-data/*__contact.md "$TEST_DIR/" 2>/dev/null || echo "No sample contacts to copy"
    fi
fi

echo "Using test directory: $TEST_DIR"
echo ""
echo "To test task creation:"
echo "1. Press 'd' on any contact, choose a state like (f)ollowup"
echo "2. Press 'e' to edit a contact, change State to (p)ing" 
echo "3. Press 'c' to create new contact with State (s)cheduled"
echo ""
echo "Tasks will be created in: $TEST_DIR/*__task.md"
echo ""

# Set the contacts directory to our test directory
export DENOTE_CONTACTS_DIR="$TEST_DIR"

# Run the app
go run cmd/main.go

echo ""
echo "After testing, check for tasks:"
echo "ls -la $TEST_DIR/*__task.md"