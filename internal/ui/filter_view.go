package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	filterLabelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("250"))
	filterHotkeyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214")).
		Bold(true)
	filterValueStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252"))
	filterActiveStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Bold(true)
)

// updateFilter handles input in filter popup
func (m Model) updateFilter(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "q":
		// Close filter menu without changes
		m.showFilterPopup = false
		return m, nil
	
	// Clear all filters
	case "a":
		m.filterType = ""
		m.filterState = ""
		m.filterStatus = ""
		m.applyFilters()
		m.showFilterPopup = false
		m.message = "Cleared all filters"
		return m, clearMessageAfter(3 * time.Second)
	
	// Type filters
	case "f": // family
		return m.applyFilterAndReturn("type", "family", "Filtered to family")
	case "c": // close
		return m.applyFilterAndReturn("type", "close", "Filtered to close")
	case "n": // network
		return m.applyFilterAndReturn("type", "network", "Filtered to network")
	case "w": // work
		return m.applyFilterAndReturn("type", "work", "Filtered to work")
	case "r": // recruiters
		return m.applyFilterAndReturn("type", "recruiters", "Filtered to recruiters")
	case "p": // providers
		return m.applyFilterAndReturn("type", "providers", "Filtered to providers")
	case "s": // social
		return m.applyFilterAndReturn("type", "social", "Filtered to social")
	
	// State filters (using uppercase to avoid conflicts)
	case "F": // followup
		return m.applyFilterAndReturn("state", "followup", "Filtered to follow up")
	case "P": // ping
		return m.applyFilterAndReturn("state", "ping", "Filtered to ping")
	case "S": // scheduled
		return m.applyFilterAndReturn("state", "scheduled", "Filtered to scheduled")
	case "T": // timeout
		return m.applyFilterAndReturn("state", "timeout", "Filtered to timeout")
	
	// Status filters
	case "o": // overdue
		return m.applyFilterAndReturn("status", "overdue", "Filtered to overdue")
	case "d": // due soon (needs attention)
		return m.applyFilterAndReturn("status", "needsAttention", "Filtered to due soon")
	case "g": // good
		return m.applyFilterAndReturn("status", "ok", "Filtered to good timing")
	}
	
	return m, nil
}

// applyFilterAndReturn applies a single filter and returns to list
func (m Model) applyFilterAndReturn(filterType, value, message string) (Model, tea.Cmd) {
	// Clear all filters first (single filter at a time)
	m.filterType = ""
	m.filterState = ""
	m.filterStatus = ""
	
	// Apply the selected filter
	switch filterType {
	case "type":
		m.filterType = value
	case "state":
		m.filterState = value
	case "status":
		m.filterStatus = value
	}
	
	m.applyFilters()
	m.showFilterPopup = false
	m.message = message
	return m, clearMessageAfter(3 * time.Second)
}

// renderFilterPopup renders the filter selection popup
func (m Model) renderFilterPopup() string {
	var b strings.Builder
	
	// Title
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("214"))
	b.WriteString(titleStyle.Render("Filter Contacts"))
	b.WriteString("\n\n")
	
	// Currently active filter
	if m.filterType != "" || m.filterState != "" || m.filterStatus != "" {
		b.WriteString(filterLabelStyle.Render("Active: "))
		if m.filterType != "" {
			b.WriteString(filterActiveStyle.Render(fmt.Sprintf("Type: %s", m.filterType)))
		} else if m.filterState != "" {
			b.WriteString(filterActiveStyle.Render(fmt.Sprintf("State: %s", m.filterState)))
		} else if m.filterStatus != "" {
			statusDisplay := m.filterStatus
			switch m.filterStatus {
			case "overdue":
				statusDisplay = "overdue"
			case "needsAttention":
				statusDisplay = "due soon"
			case "ok":
				statusDisplay = "good timing"
			}
			b.WriteString(filterActiveStyle.Render(fmt.Sprintf("Status: %s", statusDisplay)))
		}
		b.WriteString("\n\n")
	}
	
	hotkeyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	
	// Clear option
	b.WriteString(fmt.Sprintf("  %s Clear all filters\n\n", hotkeyStyle.Render("(a)")))
	
	// Type section
	b.WriteString(filterLabelStyle.Render("By Type:"))
	b.WriteString("\n")
	typeOptions := []struct {
		key   string
		value string
		label string
	}{
		{"f", "family", "Family"},
		{"c", "close", "Close"},
		{"n", "network", "Network"},
		{"w", "work", "Work"},
		{"r", "recruiters", "Recruiters"},
		{"p", "providers", "Providers"},
		{"s", "social", "Social"},
	}
	
	for _, opt := range typeOptions {
		selected := m.filterType == opt.value
		if selected {
			b.WriteString(fmt.Sprintf("  %s %s %s\n", 
				hotkeyStyle.Render("("+opt.key+")"),
				filterActiveStyle.Render("●"),
				filterActiveStyle.Render(opt.label)))
		} else {
			b.WriteString(fmt.Sprintf("  %s   %s\n", 
				hotkeyStyle.Render("("+opt.key+")"),
				opt.label))
		}
	}
	
	b.WriteString("\n")
	
	// State section
	b.WriteString(filterLabelStyle.Render("By State:"))
	b.WriteString("\n")
	stateOptions := []struct {
		key   string
		value string
		label string
	}{
		{"F", "followup", "Follow Up"},
		{"P", "ping", "Ping"},
		{"S", "scheduled", "Scheduled"},
		{"T", "timeout", "Timeout"},
	}
	
	for _, opt := range stateOptions {
		selected := m.filterState == opt.value
		if selected {
			b.WriteString(fmt.Sprintf("  %s %s %s\n", 
				hotkeyStyle.Render("("+opt.key+")"),
				filterActiveStyle.Render("●"),
				filterActiveStyle.Render(opt.label)))
		} else {
			b.WriteString(fmt.Sprintf("  %s   %s\n", 
				hotkeyStyle.Render("("+opt.key+")"),
				opt.label))
		}
	}
	
	b.WriteString("\n")
	
	// Status section
	b.WriteString(filterLabelStyle.Render("By Status:"))
	b.WriteString("\n")
	statusOptions := []struct {
		key   string
		value string
		label string
	}{
		{"o", "overdue", "Overdue"},
		{"d", "needsAttention", "Due Soon"},
		{"g", "ok", "Good Timing"},
	}
	
	for _, opt := range statusOptions {
		selected := m.filterStatus == opt.value
		if selected {
			b.WriteString(fmt.Sprintf("  %s %s %s\n", 
				hotkeyStyle.Render("("+opt.key+")"),
				filterActiveStyle.Render("●"),
				filterActiveStyle.Render(opt.label)))
		} else {
			b.WriteString(fmt.Sprintf("  %s   %s\n", 
				hotkeyStyle.Render("("+opt.key+")"),
				opt.label))
		}
	}
	
	b.WriteString("\n")
	b.WriteString(hotkeyStyle.Render("Esc to cancel"))
	
	// Pad to fill screen
	content := b.String()
	lines := strings.Split(content, "\n")
	for len(lines) < m.height-1 {
		lines = append(lines, "")
	}
	
	return strings.Join(lines, "\n")
}

