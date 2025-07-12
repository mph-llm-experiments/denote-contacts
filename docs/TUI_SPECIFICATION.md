# Denote Tasks TUI Specification

This document provides a comprehensive specification of the Terminal User Interface (TUI) implementation for denote-tasks. It's designed to serve as a blueprint for implementing similar task management interfaces.

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [UI Layout and Components](#ui-layout-and-components)
3. [List View Specification](#list-view-specification)
4. [Detail View Specification](#detail-view-specification)
5. [Editing Mechanisms](#editing-mechanisms)
6. [Hotkey System](#hotkey-system)
7. [Color Scheme and Styling](#color-scheme-and-styling)
8. [Navigation Patterns](#navigation-patterns)
9. [State Management](#state-management)
10. [Code Architecture Patterns](#code-architecture-patterns)

## Architecture Overview

### Core Components

The TUI is built using the **Bubble Tea** framework with the following key files:

- `model.go` - Central state management and data structures
- `views.go` - Rendering logic for different UI views
- `keys.go` - Keyboard input handling
- `task_view.go` - Task detail view implementation
- `project_view.go` - Project detail view implementation
- `constants.go` - UI constants and configuration
- `field_renderer.go` - Consistent field display logic
- `navigation.go` - Common navigation patterns

### Model-View Pattern

The application follows a Model-View pattern where:
- **Model** (`Model` struct) contains all application state
- **View** methods render the current state
- **Update** methods handle state transitions based on user input

## UI Layout and Components

### Main List View Layout

```
┌─────────────────────────────────────────────────────┐
│ Denote Tasks                                        │ ← Title
│ 42 tasks | Area: work | Priority: p1 | Sort: due ↑ │ ← Status line
│                                                     │
│ > ○ [p1]      [2024-01-15]  Task title...   [tag]  │ ← Task line
│   ✓ [p2] [3]  [2024-01-20]  Another task... [tag]  │
│   ⏸      [5]               Paused task...          │
│                                                     │
│ j/k:nav • /:search • enter:view • c:create • q:quit│ ← Footer
└─────────────────────────────────────────────────────┘
```

### Components

1. **Header** (140 lines height)
   - Title line
   - Status/filter line
   - Empty separator line

2. **List Area** (dynamic height)
   - Scrollable task/project list
   - Visual cursor (">")
   - Status symbols, priority, dates, etc.

3. **Footer** (1 line)
   - Context-sensitive hotkey hints

## List View Specification

### Task Line Format

Each task line follows this exact format:

```
[selector] [status] [priority] [estimate] [due_date]  [title] [tags] [area] [project]
```

Field specifications:

| Field | Width | Alignment | Example |
|-------|-------|-----------|---------|
| Selector | 1 | Left | ">" or " " |
| Status | 1 | Left | "○", "✓", "⏸", "→", "⨯" |
| Priority | 4 | Left | "[p1]", "[p2]", "[p3]", "    " |
| Estimate | 5 | Right | "[  3]", "     " |
| Due Date | 12 | Left | "[2024-01-15]" |
| Title | 40 | Left | "Task title with notes indicator ≡" |
| Tags | 20 | Left | "[work, urgent]" |
| Area | 10 | Left | "(personal)" |
| Project | 15 | Left | "→ Project Name" |

### Project Line Format

Projects use the same format as tasks with these differences:
- Status symbol is "▶" for active projects
- No estimate field (but space is preserved for alignment)
- No project field (since it IS a project)
- Entire line uses cyan color when project is active

### Special Features

1. **Notes Indicator**: Tasks/projects with notes show "≡" prefix on title
2. **Due Date Divider**: When sorted by due date, shows "────→ due today" divider
3. **Project Links**: Tasks show their project in cyan if project is active
4. **Truncation**: Long text truncated with "..." suffix

## Detail View Specification

### Task Detail View

```
┌─────────────────────────────────────────────────────┐
│ Task Details                                        │
│                                                     │
│   (T)itle      : Complete the TUI documentation    │
│   (s)tatus     : ○ open                           │
│   (p)riority   : p1                                │
│   (d)ue Date   : 2024-01-15                       │
│   (a)rea       : work                              │
│   es(t)imate   : 5                                 │
│   ta(g)s       : documentation ui                  │
│   pro(j)ect    : Denote Tasks                     │
│                                                     │
│   File         : /path/to/task.md                  │
│   ID           : 20240101T120000                   │
│ ─────────────────────────────────────────────────── │
│ This is the task body content. It can contain      │
│ multiple lines of notes and documentation.         │
│                                                     │
│ [2024-01-10 Wed]: Added initial notes              │
│ [2024-01-12 Fri]: Updated progress                 │
│                                                     │
│ q:back • E:edit • p:priority • s:status • l:log    │
└─────────────────────────────────────────────────────┘
```

### Project Detail View (with tabs)

```
┌─────────────────────────────────────────────────────┐
│ Project: Denote Tasks                              │
│ ┌──────────┐ ┌───────┐                            │
│ │ Tasks (5) │ │ Notes │                            │ ← Active tab highlighted
│ └──────────┘ └───────┘                            │
│   (T)itle      : Denote Tasks                     │
│   (s)tatus     : ● active                          │
│   (p)riority   : p1                                │
│   (d)ue Date   : 2024-02-01                       │
│   (a)rea       : development                       │
│   ta(g)s       : software opensource               │
│ ─────────────────────────────────────────────────── │
│ > ○ [p1] [2024-01-15]  Implement TUI...           │ ← Task list
│   ✓ [p2] [2024-01-10]  Design architecture...     │
│                                                     │
│ tab:switch • n:new task • enter:view • x:delete    │
└─────────────────────────────────────────────────────┘
```

## Editing Mechanisms

### Inline Editing

For simple fields (priority, status, due date), editing happens inline:

1. **Popup Editors** (for due date, tags, estimate):
   ```
   ┌─────────────────────────────────────┐
   │ Edit Due Date                       │
   │                                     │
   │ Examples: today, tomorrow, 7d, fri  │
   │ Format: YYYY-MM-DD or natural      │
   │                                     │
   │ Input: 2024-01-20█                 │
   │ → 2024-01-20                       │
   │                                     │
   │ Enter to save, Esc to cancel        │
   └─────────────────────────────────────┘
   ```

2. **In-View Editing** (task view fields):
   - Field label changes to show it's editable
   - Cursor appears in the field value
   - Background color changes to indicate edit mode

### Multi-Field Forms

For task creation:

```
┌─────────────────────────────────────────────────────┐
│ Create New Task                                     │
│                                                     │
│ → Title: Write documentation█                       │
│   Priority: p1 (p1, p2, p3)                        │
│   Due Date: 2024-01-20 (YYYY-MM-DD or natural)    │
│   Area: work (life context)                        │
│   Project: Denote Tasks (press Enter to select)    │
│   Estimate: 5 (numeric value)                      │
│   Tags: doc ui (space-separated)                   │
│                                                     │
│ ↑/↓ to navigate, Enter to save, Esc to cancel      │
└─────────────────────────────────────────────────────┘
```

### Project Selection

```
┌─────────────────────────────────────────────────────┐
│ Select Project                                      │
│                                                     │
│   0. ✗ (None - unassign from project)             │
│ > 1. ● Denote Tasks (development) [2024-02-01]    │
│   2. ● Another Project (personal)                  │
│   3. ⏸ Paused Project (work)                      │
│                                                     │
│ j/k:nav • 1-9:quick • Enter:select • Esc:cancel    │
└─────────────────────────────────────────────────────┘
```

## Hotkey System

### Global Navigation Keys

| Key | Action | Context |
|-----|--------|---------|
| `j`, `↓` | Move down | Lists, menus |
| `k`, `↑` | Move up | Lists, menus |
| `g` | Go to top | Lists |
| `G` | Go to bottom | Lists |
| `gg` | Go to top (vim style) | Lists |
| `Ctrl+d` | Page down | Lists |
| `Ctrl+u` | Page up | Lists |
| `Tab` | Next field/tab | Forms, views |
| `Shift+Tab` | Previous field/tab | Forms, views |

### Mode-Specific Keys

#### Normal Mode (List View)
| Key | Action | Notes |
|-----|--------|-------|
| `/` | Search | Fuzzy search, `#tag` for tags |
| `Enter` | View details | Opens task/project view |
| `c` | Create | Task or project based on view |
| `0` | Clear priority | |
| `1-3` | Set priority | p1, p2, p3 |
| `d` | Edit due date | Opens popup editor |
| `e` | Edit estimate | Tasks only |
| `t` | Edit tags | Opens popup editor |
| `s` | Change state | Tasks only |
| `x` | Delete | With confirmation |
| `E` | External editor | |
| `l` | Add log entry | Tasks only |
| `f` | Filter menu | |
| `S` | Sort menu | |
| `P` | Projects view | Toggle |
| `T` | Tasks view | Toggle |
| `?` | Help | |
| `q` | Quit | |

#### Task/Project View
| Key | Action | Notes |
|-----|--------|-------|
| `q`, `Esc` | Back to list | |
| `T` | Edit title | Direct edit |
| `p` | Edit priority | |
| `s` | Edit status | |
| `d` | Edit due date | |
| `a` | Edit area | |
| `t` | Edit tags | |
| `e` | Edit estimate | Tasks only |
| `j` | Edit project | Tasks only |
| `l` | Add log | Tasks only |
| `E` | External editor | |
| `Tab` | Switch tabs | Projects only |

### Quick Actions

- **Priority**: Press `1`, `2`, `3`, or `0` (clear)
- **State Menu**: Press letter for state (`o`pen, `p`aused, `d`one, etc.)
- **Project Selection**: Press `1-9` for quick select

## Color Scheme and Styling

### Color Palette

| Color | Lipgloss Color | Usage |
|-------|----------------|-------|
| Base Text | 252 | Normal text (light gray) |
| Selected | 214 | Orange for selection |
| Title | 99 | Purple for titles |
| Done | 70 | Green for completed |
| Overdue | 196 | Red for overdue |
| Priority High | 196 | Red for p1 |
| Priority Medium | 214 | Orange for p2 |
| Priority Low | 248 | Light gray for p3 |
| Project/Active | 135/51 | Purple/Cyan |
| Paused | 243 | Dim gray |
| Delegated | 33 | Blue |
| Dropped | 240 | Dark gray |
| Help Text | 248 | Bright gray |
| Status | 245 | Medium gray |

### Styling Rules

1. **Hierarchy**: Selected items override all other styling
2. **State Colors**: Task state determines line color
3. **Priority Badges**: Always colored individually
4. **Active Projects**: Use cyan for emphasis
5. **Overdue Items**: Red text, bold
6. **Empty Values**: Italic, darker gray

## Navigation Patterns

### List Navigation

```go
// Handled by NavigationHandler
- Single step: j/k, up/down arrows
- Jump to ends: g (top), G (bottom)  
- Page movement: Ctrl+d (down), Ctrl+u (up)
- Vim double-tap: gg (top)
```

### Scrolling Behavior

1. **Viewport Management**:
   - Cursor stays in middle third when possible
   - Smooth scrolling (one line at a time)
   - Page size = 10 lines

2. **Edge Cases**:
   - At top: cursor at first item
   - At bottom: cursor at last item
   - Empty list: show help message

### Tab Navigation

Projects use horizontal tabs:
- `Tab` key cycles through tabs
- Active tab highlighted with orange border
- Tab shows count (e.g., "Tasks (5)")

## State Management

### Core State Structure

```go
type Model struct {
    // Core data
    files      []denote.File    // All files
    filtered   []denote.File    // After filters
    cursor     int              // List position
    
    // UI State
    mode       Mode             // Current mode
    width      int             // Terminal width
    height     int             // Terminal height
    
    // Filters
    searchQuery    string
    areaFilter     string
    priorityFilter string
    stateFilter    string
    soonFilter     bool
    projectFilter  bool
    
    // Viewing state
    viewingTask    *denote.Task
    viewingProject *denote.Project
    editingField   string
    editBuffer     string
    editCursor     int
}
```

### Mode Transitions

```
Normal ─┬─> Search (/)
        ├─> TaskView (Enter)
        ├─> ProjectView (Enter on project)
        ├─> Create (c)
        ├─> Help (?)
        ├─> FilterMenu (f)
        ├─> SortMenu (S)
        └─> Various edit modes (d, t, e)
```

### Data Flow

1. **File Scanning**: Always reads fresh from disk (no caching)
2. **Filtering**: Applied in order: type → search → area → priority → state → soon
3. **Sorting**: Applied after filtering
4. **Rendering**: Only visible items processed

## Code Architecture Patterns

### 1. Separation of Concerns

```
├── State Management (model.go)
│   ├── Data structures
│   ├── State transitions
│   └── Business logic
│
├── Rendering (views.go, task_view.go, project_view.go)
│   ├── Layout computation
│   ├── Style application
│   └── Component composition
│
├── Input Handling (keys.go, task_view_keys.go, project_view_keys.go)
│   ├── Mode-specific handlers
│   ├── Navigation logic
│   └── Command execution
│
└── Utilities
    ├── Field rendering (field_renderer.go)
    ├── Navigation (navigation.go)
    └── Constants (constants.go)
```

### 2. Consistent Patterns

#### Field Rendering Pattern
```go
// All fields follow same pattern
func (m Model) renderFieldWithHotkey(label, value, emptyText, hotkey string) string {
    // 1. Determine display label with hotkey
    // 2. Check if editing
    // 3. Apply appropriate styling
    // 4. Return formatted string
}
```

#### Mode Handler Pattern
```go
func (m Model) handleXModeKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "esc":
        // Return to previous mode
    case "enter":
        // Confirm action
    default:
        // Mode-specific handling
    }
}
```

#### List Rendering Pattern
```go
// 1. Calculate visible range
// 2. Iterate visible items
// 3. Render each line
// 4. Apply selection highlighting
// 5. Join with newlines
```

### 3. Key Design Decisions

1. **No Caching**: Always read fresh from disk
   - Simple, correct, no staleness bugs
   - Performance is fine for text files

2. **Mode-Based UI**: Each mode has dedicated handler
   - Clear state transitions
   - Easy to add new modes
   - Predictable behavior

3. **Consistent Hotkeys**: Same key does same thing
   - `d` always edits due date
   - `t` always edits tags
   - Numbers always set priority

4. **Visual Feedback**: Every action has feedback
   - Status messages
   - Color changes
   - Cursor indicators

5. **Flexible Rendering**: Components are composable
   - Headers, lists, footers
   - Reusable field renderers
   - Consistent styling

### 4. Extension Points

To adapt for other domains (e.g., contacts):

1. **Replace Data Structures**:
   - Change `Task`/`Project` to domain entities
   - Update metadata fields
   - Adjust file parsing

2. **Modify List Columns**:
   - Update line format in `renderXLine()`
   - Adjust column widths in constants
   - Change sort fields

3. **Customize Detail Views**:
   - Replace task fields with domain fields
   - Update edit handlers
   - Adjust validation rules

4. **Adapt Hotkeys**:
   - Keep navigation keys
   - Replace action keys
   - Update help text

5. **Adjust Styling**:
   - Keep color system
   - Update status symbols
   - Modify priority levels

## Implementation Checklist

When implementing a similar TUI:

- [ ] Set up Bubble Tea framework
- [ ] Define data model and state structure
- [ ] Implement file reading/parsing
- [ ] Create main list view rendering
- [ ] Add keyboard navigation
- [ ] Implement filtering system
- [ ] Add sorting capabilities
- [ ] Create detail view
- [ ] Implement inline editing
- [ ] Add popup editors
- [ ] Create form inputs
- [ ] Implement search
- [ ] Add status messages
- [ ] Create help system
- [ ] Test all hotkeys
- [ ] Verify color accessibility
- [ ] Handle edge cases
- [ ] Add data validation
- [ ] Implement confirmations
- [ ] Polish transitions