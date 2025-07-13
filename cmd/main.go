package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mph-llm-experiments/denote-contacts/internal/ui"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("denote-contacts v0.1.0")
		os.Exit(0)
	}

	contactsDir := os.Getenv("DENOTE_CONTACTS_DIR")
	if contactsDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Could not determine home directory:", err)
		}
		contactsDir = homeDir + "/Documents/denote"
	}

	m := ui.NewModel(contactsDir)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}