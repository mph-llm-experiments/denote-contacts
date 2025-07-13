# Filter Behavior

## Overview

The filter system in denote-contacts uses a single-filter approach for simplicity. When you press 'f' to open the filter menu, you can select one filter option which will immediately apply and return you to the filtered list.

## How It Works

1. Press 'f' from the list view to open the filter menu
2. Select any filter option with its hotkey
3. The filter is immediately applied and you return to the list
4. A message flashes showing what filter was applied
5. The header shows the active filter

## Filter Options

### Clear Filter
- **(a)** - Clear all filters and show all contacts

### Type Filters
- **(f)** - Family
- **(c)** - Close
- **(n)** - Network
- **(w)** - Work
- **(r)** - Recruiters
- **(p)** - Providers
- **(s)** - Social

### State Filters (uppercase to avoid conflicts)
- **(F)** - Follow Up
- **(P)** - Ping
- **(S)** - Scheduled
- **(T)** - Timeout

### Status Filters
- **(o)** - Overdue (contacts past their frequency)
- **(d)** - Due Soon (contacts within 7 days of frequency)
- **(g)** - Good Timing (contacts within half their frequency)

## Behavior Notes

- Only one filter can be active at a time
- Selecting a new filter replaces the previous one
- Search (/) works independently and in addition to filters
- The active filter shows in the header (e.g., "[3/10] 10 of 45 (type: family)")
- Press 'f' then 'a' to clear filters quickly

## Examples

1. View only family contacts: `f` then `f`
2. View overdue contacts: `f` then `o`
3. View contacts needing follow up: `f` then `F`
4. Clear filter and see all: `f` then `a`