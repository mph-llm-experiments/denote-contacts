# CLAUDE.md - Project Context for denote-contacts

This file contains important context about the denote-contacts project to help AI assistants understand the codebase, architecture decisions, and current state.

## Project Overview

**denote-contacts** is a focused contacts management tool built on the Denote file naming convention. It uses Denote format for consistent file identification and backward compatibility while providing powerful contact management features.

## Important Documents

**ALWAYS READ THESE FIRST:**

- `/docs/DENOTE_CONTACTS_SPEC.md` - Contact file format specification
- `/docs/TUI_SPECIFICATION.md` - UI patterns from denote-tasks
- `/docs/CONTACTS_TUI_ARCHITECTURE.md` - Technical architecture
- `/docs/DENOTE_ECOSYSTEM_VISION.md` - Overall vision and integration
- `/docs/DENOTE_TASKS_INTEGRATION.md` - Task system integration

## Denote File Format

Contact files MUST follow this exact naming pattern:
- Pattern: `YYYYMMDD--kebab-case-name__contact.md`
- Example: `20250607--sarah-chen__contact.md`
- The date is 8 digits ONLY (YYYYMMDD), no time component
- The identifier in frontmatter matches this 8-digit date
- Double underscore (`__`) appears before the file type tag
- Filename uses single tag `contact` to identify file type
- Additional tags go in the YAML frontmatter, not the filename

## Tag System

### Filename Tags
- Contact files MUST use `__contact` in the filename
- This identifies the file type in the Denote system
- Only ONE tag in the filename (the file type identifier)

### Frontmatter Tags
- The `tags` array in YAML MUST include `contact`
- Additional tags can be added for organization:
  ```yaml
  tags: [contact, personal, tech, portland]
  ```
- These additional tags are for categorization and search

## Architecture Principles

1. **Denote Format** - Use Denote naming for consistent IDs
2. **Contacts Focus** - Only contacts files, no general notes
3. **No External Dependencies** - No TaskWarrior, Things, dstask, or SQLite
4. **Simplicity** - Focused functionality over feature creep

## Testing Guidelines

### CRITICAL RULE: NEVER MARK FEATURES AS COMPLETE WITHOUT HUMAN TESTING

**STOP AND READ:** Any feature implementation MUST be marked as "IMPLEMENTED BUT NOT TESTED" until the human has confirmed it works. This includes:

- Never marking issues as "✅ Completed" without human confirmation
- Never updating todo lists to "completed" for untested features
- Always use phrases like "implemented but needs testing" or "code complete, awaiting manual testing"
- NEVER assume code that compiles successfully actually works

### For TUI Development

**IMPORTANT FOR AI ASSISTANTS:** It is IMPOSSIBLE to test TUI applications in this environment. NEVER attempt to run or test the TUI. Instead, always ask the user to test the features and provide feedback. TUI applications require an interactive terminal which is not available in this context.

Since TUI applications can't be tested in this environment:

1. **Implement features completely** before declaring done
2. **Document what needs manual testing** in PROGRESS.md
3. **Create test configurations** (never modify user configs)
4. **List specific test cases** for human testing
5. **Ask the user to test** rather than attempting to test yourself

### Manual Testing Checklist

When implementing TUI features, provide:

- Step-by-step testing instructions
- Expected behavior for each feature
- Edge cases to verify
- Sample test data if needed

## TUI Implementation Guidelines

Follow the exact patterns from pdxmph/denote-tasks:
- Bubble Tea Model-View-Update architecture
- List view with specific column widths (matching denote-tasks)
- Consistent hotkey system (j/k navigation, vim-style movements)
- Color scheme: Lipgloss colors (252 base, 214 selected, etc.)
- No modal dialogs - use inline editing patterns
- Status line showing counts and current state

## Contact Management Features

### Relationship Types & Frequencies
- close: 30 days
- family: 30 days  
- network: 90 days
- work: 60 days
- social: No default frequency
- Custom frequencies override defaults

### Contact Styles
- periodic: Regular check-ins based on frequency
- ambient: Passive monitoring, no active reminders
- triggered: Event-based contact only

### Interaction Types
email, call, text, meeting, social, bump, note

### The "Bump" Concept
A bump is reviewing a contact without actually contacting them. It updates last_bump_date but NOT last_contacted.

## Contact List Display Format

Following denote-tasks column layout patterns:
- Status indicator (●/○/!) for overdue/ok/needs attention
- Relationship badge [close]/[family]/[network]/[work]
- Days since contact (right-aligned, 3 chars)
- Name (left-aligned, truncated to fit)
- Company/role (if available)
- Tags (excluding 'contact' tag)
- Last interaction indicator

## Expected Hotkeys (denote-tasks patterns)

### List View
- j/k, ↑/↓: Navigate
- g/G: Go to top/bottom
- Ctrl+d/u: Page down/up
- Enter: View contact details
- c: Create new contact
- d: Mark contacted (select interaction type)
- b: Bump contact
- /: Search/filter
- t: Create task in denote-tasks
- E: Open in external editor
- x: Delete (with confirmation)
- q/Esc: Quit

### Detail View
- Tab: Next field
- Shift+Tab: Previous field
- Enter: Edit current field
- Esc: Back to list

## Common Pitfalls to Avoid

1. **Don't modify user configs** - Always use test configs
2. **Don't assume TUI works** - It needs terminal testing
3. **Don't add non-contacts features** - This is a contacts management tool only
4. **Don't add general notes support** - Contacts files only
5. **Don't forget PROGRESS.md** - Update it regularly
6. **Don't add caching** - We're working with small text files. Always question any caching you find and ask if we can remove it. Caching causes staleness bugs without meaningful performance benefits.
7. **Don't stray from focus** - If it's not about contacts, it doesn't belong

## Questions/Decisions

- **Why Bubble Tea?** - Modern, well-maintained, good docs
- **Why not fork contacts-tui?** - Too much legacy, want clean start
- **Why contacts-only?** - Focus and simplicity beat feature creep
- **Why keep Denote format?** - Consistent IDs, backward compatibility
- **Why remove notes?** - Clear purpose, simpler codebase
- **Why unified architecture?** - Easier maintenance, consistent behavior

## Performance Philosophy

We prioritize simplicity and correctness over premature optimization:

- Always read files fresh from disk - no caching
- These are small markdown files (typically < 200 lines)
- File I/O is negligible compared to user interaction time
- Eliminating cache eliminates an entire class of staleness bugs

The Denote format is used purely for its excellent file naming convention and ID system, not because we're trying to be a general Denote file manager.

## Ecosystem Integration

denote-contacts is designed to work with:
- **denote-tasks**: Generate follow-up tasks, reference contacts
- **notes-tui**: Link contacts in meeting notes
- All use same directory structure and Denote naming

Integration is one-way: contacts can create tasks, but doesn't depend on task systems.

## Format Mistakes to Prevent

NEVER use these incorrect formats:
- ❌ `20250607--sarah-chen__contact__personal.md` (tags go in frontmatter, not filename)
- ❌ `20250607--sarah-chen_contact.md` (must use double underscore)
- ❌ `20250607-sarah-chen__contact.md` (must use double dash after date)
- ❌ `YYYYMMDDTHHMMSS--name__contact.md` (no time in filename)

ALWAYS use:
- ✅ `YYYYMMDD--kebab-case-name__contact.md` (single type tag in filename)
- ✅ Additional tags in frontmatter: `tags: [contact, work, engineering]`

## Implementation Order

1. Core file parsing (reuse Denote parser)
2. List view matching denote-tasks layout
3. Basic navigation and display
4. Contact detail view
5. Mark contacted functionality
6. Bump feature
7. Search/filter
8. Create/edit contacts
9. Task integration hooks
10. Advanced features (tags, relationships)

## Explicit Non-Goals

- NO task backend integrations (TaskWarrior, Things, dstask)
- NO SQLite or database layer
- NO general notes functionality
- NO caching of any kind
- NO modal dialogs (use inline editing)
- NO custom configuration formats (use standard Denote)

## Development Workflow

1. Study denote-tasks implementation for UI patterns
2. Implement features following the exact column widths and styling
3. Document testing requirements in PROGRESS.md
4. Never modify the existing Denote file parser - it's battle-tested
5. Focus on contact-specific features only

## Denote Parser Requirements

The existing Denote parser expects:
1. Exactly 8 digits for the date (YYYYMMDD)
2. Double dash (`--`) after the date
3. Kebab-case name (lowercase, hyphens for spaces)
4. Double underscore (`__`) before the tag
5. Single tag only in filename (multiple tags in frontmatter)
6. `.md` extension

Do NOT modify the parser - it's battle-tested and correct.

---
