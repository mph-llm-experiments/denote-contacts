package parser

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mph-llm-experiments/denote-contacts/internal/model"
	"gopkg.in/yaml.v3"
)

// ParseContactFile parses a Denote-format contact file
func ParseContactFile(path string) (model.Contact, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return model.Contact{}, fmt.Errorf("error reading file: %w", err)
	}

	// Split frontmatter and content
	parts := bytes.SplitN(content, []byte("---\n"), 3)
	if len(parts) < 3 {
		return model.Contact{}, fmt.Errorf("invalid file format: no frontmatter found")
	}

	// Parse YAML frontmatter
	var contact model.Contact
	if err := yaml.Unmarshal(parts[1], &contact); err != nil {
		return model.Contact{}, fmt.Errorf("error parsing frontmatter: %w", err)
	}

	// Validate required fields
	if !containsTag(contact.Tags, "contact") {
		return model.Contact{}, fmt.Errorf("not a contact file: missing 'contact' tag")
	}

	// Set runtime fields
	contact.FilePath = path
	contact.Content = string(parts[2])

	// Parse filename to extract identifier if not set
	if contact.Identifier == "" {
		basename := strings.TrimSuffix(filepath.Base(path), ".md")
		if idx := strings.Index(basename, "--"); idx >= 0 {
			contact.Identifier = basename[:idx]
		}
	}

	return contact, nil
}

// SaveContactFile saves a contact to a Denote-format file
func SaveContactFile(contact model.Contact) error {
	// Generate filename if needed
	if contact.FilePath == "" {
		contact.FilePath = GenerateFilename(contact)
	}

	// Ensure updated_at is set
	contact.UpdatedAt = time.Now()

	// Marshal frontmatter
	frontmatter, err := yaml.Marshal(contact)
	if err != nil {
		return fmt.Errorf("error marshaling frontmatter: %w", err)
	}

	// Build file content
	var content bytes.Buffer
	content.WriteString("---\n")
	content.Write(frontmatter)
	content.WriteString("---\n")
	content.WriteString(contact.Content)

	// Write file
	return os.WriteFile(contact.FilePath, content.Bytes(), 0644)
}

// GenerateFilename generates a Denote-compliant filename for a contact
func GenerateFilename(contact model.Contact) string {
	// Use creation date or current date
	date := contact.Date
	if date.IsZero() {
		date = time.Now()
	}

	// Format: YYYYMMDD--kebab-case-name__contact.md
	identifier := date.Format("20060102")
	name := strings.ToLower(contact.Title)
	name = strings.ReplaceAll(name, " ", "-")
	name = sanitizeName(name)

	return fmt.Sprintf("%s--%s__contact.md", identifier, name)
}

// sanitizeName removes special characters and ensures valid filename
func sanitizeName(name string) string {
	// Remove special characters, keep only alphanumeric and hyphens
	var result strings.Builder
	for _, ch := range name {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '-' {
			result.WriteRune(ch)
		}
	}
	return result.String()
}

// containsTag checks if a tag exists in the tags slice
func containsTag(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}