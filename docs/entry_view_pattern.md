# Entry View Pattern

## Overview

The entry view pattern is a state management approach used throughout the denote-contacts codebase to handle navigation after completing operations that temporarily change the view (like editing, creating, or quick actions).

## How It Works

1. **Capture Entry Point**: When transitioning to an operation view (edit, create, quick type, etc.), the current view is stored in `m.entryView`:
   ```go
   m.entryView = m.currentView  // Capture where we came from
   m.currentView = ViewEdit
   ```

2. **Return to Entry**: After completing or canceling the operation, return to the captured view:
   ```go
   m.currentView = m.entryView  // Return to where we came from
   ```

## Benefits

- **Simplifies Navigation Logic**: No need for complex conditionals to determine where to return
- **Consistent Behavior**: All operations follow the same pattern
- **Flexible**: Works whether coming from list view, detail view, or any future view
- **Maintainable**: Adding new views or operations doesn't require updating return logic

## Implementation

The pattern is implemented in:

- **model.go**: 
  - `entryView ViewMode` field in Model struct
  - Initialized to ViewList as default
  - Used in contactUpdatedMsg handling for all operation completions

- **list_view.go**: Captures entry view for operations initiated from list:
  - Edit contact ('e')
  - Create contact ('c')
  - Contact interaction ('d')
  - Quick state change ('s')
  - Quick type change ('T')

- **detail_view.go**: Captures entry view for operations initiated from detail:
  - Edit contact ('e')
  - Contact interaction ('d')

- **Operation views**: Use entryView when canceling or completing:
  - edit_view.go
  - create_view.go
  - interaction_view.go
  - quicktype_view.go

## Example Usage

```go
// Starting an edit from list view
case "e":
    if m.cursor < len(m.filtered) {
        contact := m.filtered[m.cursor]
        m.editingContact = &contact
        m.initializeEditValues(contact)
        m.entryView = m.currentView  // Capture we came from list
        m.currentView = ViewEdit
        m.editField = -1
    }

// Canceling edit returns to entry view
case "esc":
    m.currentView = m.entryView  // Returns to list
    m.editingContact = nil
    m.editField = -1
    m.editValues = nil
    return m, nil

// Completing edit also returns to entry view (in contactUpdatedMsg handler)
} else if m.currentView == ViewEdit {
    m.currentView = m.entryView  // Returns to list
    m.editingContact = nil
    m.editField = -1
    m.editValues = nil
}
```

## Adding New Operations

When adding a new operation that temporarily changes the view:

1. Capture entry view before transitioning:
   ```go
   m.entryView = m.currentView
   m.currentView = ViewNewOperation
   ```

2. Return to entry view on cancel:
   ```go
   case "esc":
       m.currentView = m.entryView
       // Clean up operation state
   ```

3. Add handling in contactUpdatedMsg (if operation saves data):
   ```go
   } else if m.currentView == ViewNewOperation {
       m.currentView = m.entryView
       // Clean up operation state
   }
   ```

This pattern ensures consistent navigation behavior across all operations.