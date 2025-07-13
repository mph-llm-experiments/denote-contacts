package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Available interaction types
var interactionTypes = []struct {
	key   string
	value string
	label string
}{
	{"p", "phone", "Phone Call"},
	{"e", "email", "Email"},
	{"t", "text", "Text/SMS"},
	{"m", "meeting", "In-Person Meeting"},
	{"v", "video", "Video Call"},
	{"s", "social", "Social Media"},
	{"l", "mail", "Physical Mail"},
	{"o", "other", "Other"},
}

// Available contact states
var contactStates = []struct {
	key   string
	value string
	label string
	desc  string
}{
	{"o", "ok", "OK", "Contact is up to date"},
	{"f", "followup", "Follow Up", "Need to follow up"},
	{"p", "ping", "Ping", "Send a quick check-in"},
	{"s", "scheduled", "Scheduled", "Meeting/call is scheduled"},
	{"t", "timeout", "Timeout", "No response"},
}

// updateInteractionType handles input in the contact logging flow
func (m Model) updateInteractionType(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch m.contactLogStep {
	case 0: // Selecting interaction type
		switch msg.String() {
		case "esc", "q":
			// Cancel and go back
			m.resetContactLogging()
			return m, nil

		// Direct selection by hotkey
		case "p", "e", "t", "m", "v", "s", "l", "o":
			for _, it := range interactionTypes {
				if it.key == msg.String() {
					m.interactionType = it.value
					m.contactLogStep = 1 // Move to state selection
					return m, nil
				}
			}
		}

	case 1: // Selecting next state
		switch msg.String() {
		case "esc":
			// Go back to type selection
			m.contactLogStep = 0
			m.interactionType = ""
			return m, nil

		case "q":
			// Cancel entirely
			m.resetContactLogging()
			return m, nil

		// Direct selection by hotkey
		case "o", "f", "p", "s", "n", "t":
			for _, st := range contactStates {
				if st.key == msg.String() {
					m.interactionState = st.value
					m.contactLogStep = 2 // Move to note entry
					return m, nil
				}
			}
		}

	case 2: // Adding note
		switch msg.String() {
		case "esc":
			// Go back to state selection
			m.contactLogStep = 1
			m.interactionState = ""
			m.interactionNote = ""
			return m, nil

		case "ctrl+c", "ctrl+q":
			// Cancel entirely
			m.resetContactLogging()
			return m, nil

		case "enter":
			// Save without note
			if m.contactToMark != nil {
				return m, m.logContactInteraction(*m.contactToMark)
			}

		case "ctrl+s":
			// Save with note
			if m.contactToMark != nil && m.interactionNote != "" {
				return m, m.logContactInteraction(*m.contactToMark)
			}

		case "backspace":
			if len(m.interactionNote) > 0 {
				m.interactionNote = m.interactionNote[:len(m.interactionNote)-1]
			}

		default:
			// Add to note
			if len(msg.String()) == 1 && len(m.interactionNote) < 200 {
				m.interactionNote += msg.String()
			}
		}
	}

	return m, nil
}

// resetContactLogging clears the contact logging state and returns to previous view
func (m *Model) resetContactLogging() {
	m.currentView = m.entryView  // Return to where we came from
	m.contactToMark = nil
	m.interactionType = ""
	m.interactionState = ""
	m.interactionNote = ""
	m.contactLogStep = 0
}

// viewInteractionType renders the contact logging interface
func (m Model) viewInteractionType() string {
	if m.contactToMark == nil {
		return "No contact selected"
	}

	var b strings.Builder

	// Title
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("214"))
	
	// Check if this is a quick state change (starting at step 1)
	isQuickStateChange := m.contactLogStep == 1 && m.interactionType == "note"
	
	if isQuickStateChange {
		b.WriteString(titleStyle.Render("Change State"))
	} else {
		b.WriteString(titleStyle.Render("Log Contact"))
	}
	b.WriteString("\n\n")

	// Contact name
	if isQuickStateChange {
		b.WriteString(fmt.Sprintf("Contact: %s\n\n", 
			lipgloss.NewStyle().Bold(true).Render(m.contactToMark.Title)))
	} else {
		b.WriteString(fmt.Sprintf("Recording interaction with %s\n\n", 
			lipgloss.NewStyle().Bold(true).Render(m.contactToMark.Title)))
	}

	switch m.contactLogStep {
	case 0: // Interaction type selection
		// Progress indicator
		stepStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
		b.WriteString(stepStyle.Render("Step 1 of 3: Interaction Type"))
		b.WriteString("\n\n")

		// Prompt
		b.WriteString("How did you contact them?\n\n")

		// Options
		hotkeyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
		for _, it := range interactionTypes {
			b.WriteString(fmt.Sprintf("  %s  %s\n", 
				hotkeyStyle.Render("("+it.key+")"),
				it.label))
		}

		b.WriteString("\n")
		b.WriteString(hotkeyStyle.Render("Esc to cancel"))

	case 1: // State selection
		// Progress indicator
		stepStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
		if isQuickStateChange {
			b.WriteString(stepStyle.Render("Select new state"))
			b.WriteString("\n\n")
		} else {
			b.WriteString(stepStyle.Render("Step 2 of 3: Next State"))
			b.WriteString("\n")
			b.WriteString(fmt.Sprintf("Type: %s\n\n", filterActiveStyle.Render(m.interactionType)))
		}

		// Prompt
		b.WriteString("What's the next state for this contact?\n\n")

		// Options
		hotkeyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
		descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true)
		for _, st := range contactStates {
			b.WriteString(fmt.Sprintf("  %s  %-12s  %s\n", 
				hotkeyStyle.Render("("+st.key+")"),
				st.label,
				descStyle.Render(st.desc)))
		}

		b.WriteString("\n")
		b.WriteString(hotkeyStyle.Render("Esc to go back • q to cancel"))

	case 2: // Note entry
		// Progress indicator
		stepStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
		b.WriteString(stepStyle.Render("Step 3 of 3: Add Note (Optional)"))
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf("Type: %s • State: %s\n\n", 
			filterActiveStyle.Render(m.interactionType),
			filterActiveStyle.Render(m.interactionState)))

		// Note input
		b.WriteString("Add a note about this interaction:\n\n")
		
		// Text box
		boxStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("238")).
			Padding(1, 2).
			Width(60)
		
		noteText := m.interactionNote
		if noteText == "" {
			noteText = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true).Render("(optional)")
		}
		noteText += "█" // cursor
		
		b.WriteString(boxStyle.Render(noteText))
		b.WriteString("\n\n")

		hotkeyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
		b.WriteString(hotkeyStyle.Render("Enter to save • Ctrl+S to save with note • Esc to go back"))
	}

	// Pad to fill screen
	content := b.String()
	lines := strings.Split(content, "\n")
	for len(lines) < m.height-1 {
		lines = append(lines, "")
	}

	return strings.Join(lines, "\n")
}