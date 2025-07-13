package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// updateCreate handles input in create view
func (m Model) updateCreate(msg tea.KeyMsg) (Model, tea.Cmd) {
	// If no field is being edited (field selection mode)
	if m.editField == -1 {
		switch msg.String() {
		case "esc":
			// Cancel without saving - return to entry view
			m.currentView = m.entryView
			m.editField = -1
			m.editValues = nil
			return m, nil

		case "q":
			// Save new contact and exit
			return m, m.saveNewContact()

		// Field selection hotkeys
		case "n":
			m.editField = fieldTitle
		case "e":
			m.editField = fieldEmail
		case "p":
			m.editField = fieldPhone
		case "c":
			m.editField = fieldCompany
		case "r":
			m.editField = fieldRole
		case "l":
			m.editField = fieldLocation
		case "t":
			m.editField = fieldRelationType
		case "s":
			m.editField = fieldContactStyle
		case "S":
			m.editField = fieldState
		case "T":
			m.editField = fieldTags
		}
	} else {
		// Field editing mode
		switch msg.String() {
		case "esc":
			// Cancel field edit, return to field selection
			m.editField = -1
			return m, nil

		case "enter":
			// Save field and return to field selection
			m.editField = -1
			return m, nil

		case "backspace":
			// Delete character from current field
			if len(m.editValues[m.editField]) > 0 {
				m.editValues[m.editField] = m.editValues[m.editField][:len(m.editValues[m.editField])-1]
			}

		default:
			// Handle special fields
			if m.editField == fieldRelationType {
				return m.handleCreateTypeSelection(msg), nil
			} else if m.editField == fieldContactStyle {
				return m.handleCreateStyleSelection(msg), nil
			} else if m.editField == fieldState {
				return m.handleCreateStateSelection(msg), nil
			}
			
			// Add character to current field
			if len(msg.String()) == 1 {
				m.editValues[m.editField] += msg.String()
			}
		}
	}

	return m, nil
}

// handleCreateTypeSelection handles relationship type selection for new contacts
func (m Model) handleCreateTypeSelection(msg tea.KeyMsg) Model {
	switch msg.String() {
	case "f":
		m.editValues[fieldRelationType] = "family"
		m.editField = -1 // Return to field selection
	case "c":
		m.editValues[fieldRelationType] = "close"
		m.editField = -1 // Return to field selection
	case "n":
		m.editValues[fieldRelationType] = "network"
		m.editField = -1 // Return to field selection
	case "w":
		m.editValues[fieldRelationType] = "work"
		m.editField = -1 // Return to field selection
	case "r":
		m.editValues[fieldRelationType] = "recruiters"
		m.editField = -1 // Return to field selection
	case "p":
		m.editValues[fieldRelationType] = "providers"
		m.editField = -1 // Return to field selection
	case "s":
		m.editValues[fieldRelationType] = "social"
		m.editField = -1 // Return to field selection
	}
	return m
}

// handleCreateStyleSelection handles contact style selection for new contacts
func (m Model) handleCreateStyleSelection(msg tea.KeyMsg) Model {
	switch msg.String() {
	case "p":
		m.editValues[fieldContactStyle] = "periodic"
		m.editField = -1 // Return to field selection
	case "a":
		m.editValues[fieldContactStyle] = "ambient"
		m.editField = -1 // Return to field selection
	case "t":
		m.editValues[fieldContactStyle] = "triggered"
		m.editField = -1 // Return to field selection
	}
	return m
}

// handleCreateStateSelection handles contact state selection for new contacts
func (m Model) handleCreateStateSelection(msg tea.KeyMsg) Model {
	switch msg.String() {
	case "o":
		m.editValues[fieldState] = "ok"
		m.editField = -1 // Return to field selection
	case "f":
		m.editValues[fieldState] = "followup"
		m.editField = -1 // Return to field selection
	case "p":
		m.editValues[fieldState] = "ping"
		m.editField = -1 // Return to field selection
	case "s":
		m.editValues[fieldState] = "scheduled"
		m.editField = -1 // Return to field selection
	case "t":
		m.editValues[fieldState] = "timeout"
		m.editField = -1 // Return to field selection
	}
	return m
}

// viewCreate renders the create view
func (m Model) viewCreate() string {
	var b strings.Builder

	// Title
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("214"))
	b.WriteString(titleStyle.Render("Create New Contact"))
	b.WriteString("\n\n")

	if m.editField == -1 {
		// Field selection mode - show all fields with hotkeys
		b.WriteString(m.renderCreateFieldSelectionView())
	} else {
		// Field editing mode - show the active field
		b.WriteString(m.renderCreateFieldEditingView())
	}

	// Pad to fill screen
	content := b.String()
	lines := strings.Split(content, "\n")
	for len(lines) < m.height-1 {
		lines = append(lines, "")
	}

	return strings.Join(lines, "\n")
}

// renderCreateFieldSelectionView shows all fields with hotkeys for selection
func (m Model) renderCreateFieldSelectionView() string {
	var b strings.Builder
	
	// Show all fields with their current values and hotkeys
	fieldHotkeys := []string{"n", "e", "p", "c", "r", "l", "t", "s", "S", "T"}
	
	for i := 0; i < fieldCount; i++ {
		label := fieldLabels[i]
		value := m.editValues[i]
		hotkey := fieldHotkeys[i]
		
		if value == "" {
			value = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true).Render("(empty)")
		}
		
		b.WriteString(fmt.Sprintf("  (%s) %s: %s\n", 
			lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Render(hotkey),
			lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Render(fmt.Sprintf("%-15s", label)),
			lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render(value)))
	}
	
	b.WriteString("\n")
	
	// Instructions
	instructionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	b.WriteString(instructionStyle.Render("Select field to edit • q: save & exit • Esc: cancel"))
	
	return b.String()
}

// renderCreateFieldEditingView shows the active field being edited
func (m Model) renderCreateFieldEditingView() string {
	var b strings.Builder
	
	label := fieldLabels[m.editField]
	value := m.editValues[m.editField]
	
	// Special rendering for certain fields
	if m.editField == fieldRelationType {
		b.WriteString(fmt.Sprintf("Editing %s:\n\n", label))
		b.WriteString(m.renderTypeOptions(value))
		b.WriteString("\n\n")
	} else if m.editField == fieldContactStyle {
		b.WriteString(fmt.Sprintf("Editing %s:\n\n", label))
		b.WriteString(m.renderStyleOptions(value))
		b.WriteString("\n\n")
	} else if m.editField == fieldState {
		b.WriteString(fmt.Sprintf("Editing %s:\n\n", label))
		b.WriteString(m.renderStateOptions(value))
		b.WriteString("\n\n")
	} else {
		b.WriteString(fmt.Sprintf("Editing %s:\n\n", label))
		// Show current value with cursor
		displayValue := value + "█"
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render(displayValue))
		b.WriteString("\n\n")
	}
	
	// Instructions
	instructionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	b.WriteString(instructionStyle.Render("Type to edit • Enter: save field • Esc: cancel field"))
	
	return b.String()
}

// initializeCreateValues initializes empty values for creating a new contact
func (m *Model) initializeCreateValues() {
	m.editValues = make([]string, fieldCount)
	
	// Set some sensible defaults
	m.editValues[fieldRelationType] = "network" // Default to network
	m.editValues[fieldContactStyle] = "periodic" // Default to periodic
	m.editValues[fieldState] = "ok" // Default to ok
}