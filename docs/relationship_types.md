# Relationship Types

## Current Types (from test data)

The following relationship types are supported based on the test data in contacts-data:

1. **close** - Close friends (30 days default frequency)
2. **family** - Family members (30 days default frequency)  
3. **network** - Professional network (90 days default frequency)
4. **work** - Work colleagues (60 days default frequency)
5. **social** - Social acquaintances (no default frequency)
6. **providers** - Service providers (no default frequency)
7. **recruiters** - Recruiters (no default frequency)

## Hotkey Mappings

When selecting relationship types in the UI:

- **(f)** - family
- **(c)** - close
- **(n)** - network
- **(w)** - work
- **(r)** - recruiters
- **(p)** - providers
- **(s)** - social

## Changes Made

1. Fixed "recruiter" (singular) to "recruiters" (plural) to match test data
2. Removed "colleague" type that was incorrectly added - use "work" instead
3. Added "work" type to all UI selection menus
4. Updated column width in list view from 9 to 10 characters to accommodate "recruiters"

## Default Contact Frequencies

For contacts with `contact_style: periodic`, the default frequencies are:

- **close** - 30 days
- **family** - 30 days
- **work** - 60 days
- **network** - 90 days
- **social** - no default (must set custom)
- **providers** - no default (must set custom)
- **recruiters** - no default (must set custom)

These can be overridden with `custom_frequency_days` in the contact's frontmatter.