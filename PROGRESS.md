# Progress Tracking

## Implemented (Not Tested)

### Basic Structure (2025-01-12)
- ✓ Go project structure with go.mod
- ✓ Basic Bubble Tea app skeleton
- ✓ Contact model/types matching spec
- ✓ Denote file parser for contacts
- ✓ Basic list view with denote-tasks layout
- ✓ Navigation commands (j/k, g/G, Ctrl+d/u)
- ✓ Stub implementations for other views

### Testing Completed ✓
1. **Basic App Launch** ✓
   - App launches successfully with `./test.sh`
   - Shows TUI with contact list from test data
   - Displays 188 contacts

2. **Navigation Testing** ✓
   - j/k moves cursor up/down properly
   - q quits the application
   - Other navigation keys ready for testing

3. **Contact Display** ✓
   - Status indicators working (● for overdue periodic, ○ for others)
   - Relationship badges display correctly [close], [family], etc.
   - Days since contact shows properly (- for never contacted)
   - Contact names display correctly

### Still Needs Testing
- Enter to open detail view (IMPLEMENTED - needs testing)
- / for live fuzzy search (IMPLEMENTED - needs testing)
- f for filter popup (IMPLEMENTED - needs testing)
- g/G for top/bottom navigation
- Ctrl+d/u for page navigation
- d for mark contacted
- b for bump

## TODO

### High Priority
- [x] Implement contact detail view (IMPLEMENTED - needs testing)
- [x] Implement search/filter functionality (IMPLEMENTED - needs testing)
  - Live fuzzy search with / key
  - Filter popup with f key (type, state, status)
- [x] Implement contact logging functionality (d key) (IMPLEMENTED - needs testing)
  - Multi-step process:
    1. Select interaction type (phone, email, text, meeting, video, social, mail, other)
    2. Choose next state (ok, followup, ping, scheduled, notes, timeout)
    3. Optional: Add a note about the interaction
  - Updates last_contacted date and interaction type
  - Updates contact state based on selection
  - Adds timestamped note to contact content if provided
  - Shows contextual success message
  - Works in both list and detail views
- [x] Implement bump feature (b key) (IMPLEMENTED - needs testing)
  - Updates last_bump_date without changing last_contacted
  - Increments bump_count
  - Shows message with bump count
  - Works in both list and detail views
- [x] Add proper error handling for missing directory (IMPLEMENTED - needs testing)
  - Checks if contacts directory exists before attempting operations
  - Clear error messages for missing directory: "contacts directory '/path' does not exist. Please create it or check your configuration"
  - Validates directory is actually a directory, not a file
  - Better error messages for all save/load operations with contact names and specific failure reasons
  - Prevents crashes and provides actionable error information

### Medium Priority  
- [x] Add create new contact (c key) (IMPLEMENTED - needs testing)
  - Press 'c' in list view to create new contact
  - Same interface as edit view with field selection
  - Required field validation (name must not be empty)
  - Generates proper Denote filename format: YYYYMMDDTHHMMSS--name-slug__contact.md
  - Sets sensible defaults: network type, periodic style, ok state
  - Returns to list view with success message
  - Automatically reloads contact list to include new contact
- [x] Implement edit contact functionality (IMPLEMENTED - needs testing)
  - Press 'e' in list view or detail view to enter edit mode
  - Shows field selection view with hotkeys: (n)ame, (e)mail, (p)hone, (c)ompany, (r)ole, (l)ocation, (t)ype, (s)tyle, (S)tate, (T)ags
  - Press hotkey to select field for editing
  - Type to edit field value
  - Enter to save field and return to field selection
  - Type selection for relationship type: (f)amily, (c)lose, (C)olleague, (n)etwork, (r)ecruiter, (p)roviders, (s)ocial
  - Style selection for contact style: (p)eriodic, (a)mbient, (t)riggered
  - q to save all changes and exit (when in field selection mode)
  - Esc to cancel field edit or cancel entire edit session
  - Updates the contact's updated_at timestamp
  - Returns to appropriate view (list or detail) with success message after saving

### Low Priority
- [x] Task integration with denote-tasks (IMPLEMENTED - needs testing)
  - Automatically creates tasks when contacts change to action-requiring states
  - States that trigger tasks: followup, ping, scheduled, notes, timeout
  - Task includes contact_id field linking to contact's identifier
  - Task includes label field if contact has a label
  - Tasks are created in the same directory as contacts
  - Task format matches denote-tasks specification
  - Works when:
    - Logging interactions that change state (d key)
    - Editing contacts that change state (e key)
    - Creating new contacts with action states (c key)
- [ ] External editor support (E key)
- [ ] Delete contact with confirmation (x key)
- [ ] Export functionality

## Known Issues
- Parser doesn't handle all edge cases yet
- No config file support
- No interactive prompts for mark contacted

## Fixes Applied (2025-01-12)

### Initial Implementation
- Fixed window size initialization
- Fixed relationship badge width for consistent columns
- Added contact sorting (overdue first, then by days since contact)
- Added proper loading status
- Fixed tea.WindowSize compilation error
- Fixed status indicator logic to respect contact_style
- Simplified line formatting to match expected output
- Only show overdue/attention for periodic contacts
- Added visible cursor indicator (>) 
- Added position indicator [n/total] in header
- Changed sort to alphabetical by name for clarity
- Added company/role and tags to display
- Shows more metadata in list view
- Redesigned as columnar display with:
  - Contact style icons (↻ periodic, ◦ ambient, ! triggered)
  - Fixed emoji width issues
  - Contact state column (excluding "ok" state)
  - Removed last contacted date (redundant with days)
  - Fixed-width columns for better readability
  - Column headers
  - Two-line footer with icon legend
  - Reordered columns: Name first, then days, type, state, company, tags

### Latest Fixes
- Added more spacing between name and days columns
- Fixed handling of future dates (shows negative days)
- State column should now display properly
- Improved column alignment
- Added space between status and style icons
- Increased type column width to 8 chars for "providers"
- Added "providers" as a valid relationship type

## Notes
- Using test data from contacts-data directory
- Following denote-tasks UI patterns closely
- No caching - always reads fresh from disk