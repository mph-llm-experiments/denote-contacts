# Quick Type Change Test Instructions

## Test Steps

1. Start the application
2. Navigate to any contact in the list
3. Press 'T' to trigger quick type change
4. Expected: You should see a "Change Type" screen with the current contact's name and type
5. Press any type hotkey (f, c, C, n, r, p, s)
6. Expected: You should immediately return to the list view with a success message
7. The contact's type should be updated in the list

## What to verify:
- Pressing 'T' shows the quick type selection interface
- Selecting a type immediately returns to the list (not the edit screen)
- The success message shows the change
- The contact's type is updated in the list view

## Common issues to check:
- If you end up in the edit screen, the fix didn't work
- If the type doesn't change, there's a save issue
- If you get stuck in the quick type view, the state transition failed