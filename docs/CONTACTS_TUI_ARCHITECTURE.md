# contacts-tui Architecture Documentation

## Overview

contacts-tui is a terminal-based contact management application built with Go and the Bubble Tea TUI framework. It provides a keyboard-driven interface for managing personal and professional contacts with SQLite storage and task system integration.

## Core Architecture

### Technology Stack
- **Language**: Go 1.21+
- **TUI Framework**: Bubble Tea (Elm-inspired Model-Update-View pattern)
- **Database**: SQLite with migration support
- **Configuration**: TOML format
- **Task Integration**: Pluggable backend system (Things, Taskwarrior, dstask)

### Project Structure
```
contacts-tui/
├── main.go                    # Entry point, CLI handling
├── internal/
│   ├── config/               # Configuration management
│   │   └── config.go        # TOML config load/save
│   ├── db/                  # Database layer
│   │   ├── db.go           # Database operations
│   │   ├── models.go       # Data structures
│   │   ├── migrations.go   # Schema migrations
│   │   └── fixtures.go     # Test data generation
│   ├── export/             # Export functionality
│   │   ├── exporter.go     # Export interface
│   │   └── denote.go       # Denote format exporter
│   ├── tasks/              # Task backend system
│   │   ├── backend.go      # Backend interface
│   │   ├── things.go       # Things.app integration
│   │   ├── taskwarrior.go  # Taskwarrior integration
│   │   └── dstask.go       # dstask integration
│   └── tui/                # Terminal UI
│       ├── app.go          # Main TUI model
│       ├── views.go        # UI components
│       └── keys.go         # Keyboard shortcuts
└── migrations/             # SQL migration files
```

## Key Components

### 1. Database Layer (`internal/db/`)

#### Models
```go
type Contact struct {
    ID                   int
    Name                 string
    Email                sql.NullString
    Phone                sql.NullString
    Company              sql.NullString
    RelationshipType     string        // close, family, network, work
    State                sql.NullString // ok, needs_attention, etc.
    Notes                sql.NullString
    Label                sql.NullString // Custom tags like @mentor
    ContactStyle         string        // periodic, ambient, triggered
    CustomFrequencyDays  sql.NullInt64
    ContactedAt          sql.NullTime
    LastBumpDate         sql.NullTime
    BumpCount            int
    // ... additional fields
}

type Log struct {
    ID              int
    ContactID       int
    InteractionDate time.Time
    InteractionType string // email, call, text, meeting, etc.
    Notes           sql.NullString
    CreatedAt       time.Time
}
```

#### Key Database Functions
- `Open(dbPath string) (*DB, error)` - Opens database connection
- `ListContacts() ([]Contact, error)` - Retrieves all contacts
- `GetContact(id int) (*Contact, error)` - Get single contact
- `AddContact(contact Contact) (int64, error)` - Create new contact
- `UpdateContact(contact Contact) error` - Update existing contact
- `DeleteContact(contactID int) error` - Delete contact and logs
- `MarkContacted(contactID int, type, notes string) error` - Log interaction
- `BumpContact(contactID int) error` - Mark as reviewed
- `GetContactInteractions(contactID int, limit int) ([]Log, error)` - Get logs

### 2. TUI Layer (`internal/tui/`)

#### Model Structure
The TUI follows Bubble Tea's Model-Update-View pattern:

```go
type Model struct {
    db              *db.DB
    config          *config.Config
    contacts        []db.Contact
    selectedIdx     int
    viewMode        ViewMode
    searchQuery     string
    // ... additional state
}
```

#### View Modes
- **ListView**: Main contact list with filtering/sorting
- **DetailView**: Single contact details with interaction history
- **EditView**: Form for editing contact information
- **AddView**: Form for creating new contacts
- **SearchView**: Search interface

#### Key Bindings
- `j/k` or `↓/↑`: Navigate list
- `Enter`: View contact details
- `n`: New contact
- `e`: Edit contact
- `d`: Mark contacted (with note prompt)
- `b`: Bump contact (mark reviewed)
- `D`: Delete contact
- `/`: Search
- `q`: Quit
- `?`: Help

### 3. Configuration (`internal/config/`)

```toml
[database]
path = "~/path/to/contacts.db"

[tasks]
backend = "things"  # or "taskwarrior", "dstask", "" (auto-detect)
project = "Contacts"  # Project name for task creation

[ui]
theme = "default"
```

### 4. Export System (`internal/export/`)

#### Exporter Interface
```go
type Exporter interface {
    Export(contacts []db.Contact, getInteractions func(int) ([]db.Log, error)) error
}
```

#### Denote Exporter
Converts contacts to Denote Markdown format with:
- Standardized filenames: `YYYYMMDD--name__contact.md`
- YAML frontmatter with all contact metadata
- Interaction logs in the body

### 5. Task Backend System (`internal/tasks/`)

#### Backend Interface
```go
type Backend interface {
    Name() string
    IsAvailable() bool
    CreateTask(title, note string, tags []string) error
    UpdateTask(id, title, note string) error
    DeleteTask(id string) error
    GetTask(id string) (*Task, error)
    SyncTasks(contacts []db.Contact) error
}
```

## Key Algorithms

### Contact Overdue Calculation
```go
func (c Contact) IsOverdue() bool {
    // Archived/ambient/triggered contacts are never overdue
    // Use custom frequency if set
    // Otherwise use relationship type defaults:
    // - close/family: 30 days
    // - network: 90 days
    // - work: 60 days
}
```

### Name Sanitization for Filenames
```go
func SanitizeForFilename(s string) string {
    // Convert to lowercase
    // Replace spaces with hyphens
    // Remove special characters
    // Handle empty results
}
```

## Integration Points

### 1. Task System Integration
- Create tasks for follow-ups
- Link contacts to existing tasks
- Sync contact state with task completion

### 2. Notes Integration (Future)
- Cross-reference contacts in notes
- Link meeting notes to contacts
- Extract contact mentions

### 3. Denote Ecosystem
- Unified file format
- Cross-linking between content types
- Shared tag taxonomy

## Data Flow

1. **Contact Creation**
   - User fills form in TUI
   - Validate required fields
   - Save to SQLite database
   - Optional: Create associated task

2. **Interaction Logging**
   - User selects contact and interaction type
   - Prompt for notes
   - Update ContactedAt timestamp
   - Create interaction log entry

3. **Export Process**
   - Load all contacts from database
   - For each contact:
     - Generate Denote-compliant filename
     - Create YAML frontmatter
     - Fetch interaction logs
     - Write Markdown file

## Error Handling

- Database errors: Wrapped with context
- File I/O errors: User-friendly messages
- Task backend errors: Graceful fallback
- UI errors: Non-fatal, show message

## Performance Considerations

- Lazy loading for large contact lists
- Indexed database queries
- Batch operations for export
- Minimal UI redraws

## Testing Strategy

- Fixtures database for development
- Migration testing
- Export format validation
- Task backend mocking

## Future Enhancements

1. **Sync Capabilities**
   - Multi-device synchronization
   - Conflict resolution
   - Cloud backup

2. **Enhanced Search**
   - Full-text search
   - Tag-based filtering
   - Relationship mapping

3. **Automation**
   - Birthday reminders
   - Follow-up scheduling
   - Bulk operations

4. **Analytics**
   - Contact frequency reports
   - Relationship health metrics
   - Network visualization