package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	NotesDirectory string `toml:"notes_directory"`
}

func Load() (*Config, error) {
	config := &Config{}
	
	// Default config file location
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	
	configPath := filepath.Join(homeDir, ".config", "denote-contacts", "config.toml")
	
	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Use defaults if no config file
		config.NotesDirectory = filepath.Join(homeDir, "Documents", "denote")
		return config, nil
	}
	
	// Load config from file
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		return nil, err
	}
	
	// Expand ~ in path if present
	if len(config.NotesDirectory) > 0 && config.NotesDirectory[0] == '~' {
		config.NotesDirectory = filepath.Join(homeDir, config.NotesDirectory[1:])
	}
	
	return config, nil
}