package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mph-llm-experiments/denote-contacts/internal/config"
	"github.com/mph-llm-experiments/denote-contacts/internal/ui"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("denote-contacts v0.1.0")
		os.Exit(0)
	}

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	
	// Allow environment variable to override config
	contactsDir := os.Getenv("DENOTE_CONTACTS_DIR")
	if contactsDir == "" {
		contactsDir = cfg.NotesDirectory
	}

	m := ui.NewModel(contactsDir)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}