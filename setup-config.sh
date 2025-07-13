#!/bin/bash

# Setup script for denote-contacts configuration

CONFIG_DIR="$HOME/.config/denote-contacts"
CONFIG_FILE="$CONFIG_DIR/config.toml"

# Create config directory if it doesn't exist
mkdir -p "$CONFIG_DIR"

# Check if config already exists
if [ -f "$CONFIG_FILE" ]; then
    echo "Config file already exists at: $CONFIG_FILE"
    echo "Edit it to customize your settings."
else
    # Copy example config
    cp config.toml.example "$CONFIG_FILE"
    echo "Created config file at: $CONFIG_FILE"
    echo "Edit it to set your notes directory location."
fi

echo ""
echo "To customize your notes directory, edit:"
echo "  $CONFIG_FILE"
echo ""
echo "Or set the DENOTE_CONTACTS_DIR environment variable."