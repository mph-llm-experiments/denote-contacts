# denote-contacts

A focused terminal-based contacts management system that uses the [Denote](https://protesilaos.com/emacs/denote) file naming convention. This tool is specifically designed for managing contacts as markdown files with a consistent, searchable format.

**Note**: This project is not affiliated with the Denote project. It simply adopts Denote's excellent file naming convention for consistent contact identification.

## Important consideration before using this code or interacting with this codebase

This application is an experiment in using Claude Code as the primary driver the development of a small, focused app that concerns itself with the owner's particular point of view on the task it is accomplishing.

As such, this is not meant to be what people think of as "an open source project," because I don't have a commitment to building a community around it and don't have the bandwidth to maintain it beyond "fix bugs I find in the process of pushing it in a direction that works for me."

It's important to understand this for a few reasons:

1. If you use this code, you'll be using something largely written by an LLM with all the things we know this entails in 2025: Potential inefficiency, security risks, and the risk of data loss.

2. If you use this code, you'll be using something that works for me the way I would like it to work. If it doesn't do what you want it to do, or if it fails in some way particular to your preferred environment, tools, or use cases, your best option is to take advantage of its very liberal license and fork it.

3. I'll make a best effort to only tag the codebase when it is in a working state with no bugs that functional testing has revealed.

While I appreciate and applaud assorted efforts to certify code and projects AI-free, I think it's also helpful to post commentary like this up front: Yes, this was largely written by an LLM so treat it accordingly. Don't think of it like code you can engage with, think of it like someone's take on how to do a task or solve a problem.

## Overview

denote-contacts provides a streamlined TUI (Terminal User Interface) for managing personal and professional contacts. Each contact is stored as a markdown file with YAML frontmatter, using Denote's timestamp-based naming convention for unique identification.

### Key Features

- **Focused Purpose**: Exclusively for contact management (not a general notes system)
- **Smart Reminders**: Set contact frequencies based on relationship type
- **Visual Status**: See at a glance who's overdue, due soon, or on track
- **Quick Actions**: Log interactions, bump reviews, edit details with single keystrokes
- **Flexible Organization**: Tag and categorize contacts by type, style, and custom tags
- **Task Integration**: Automatically creates tasks in [denote-tasks](https://github.com/pdxmph/denote-tasks) when contacts need attention

## Installation

```bash
go install github.com/mph-llm-experiments/denote-contacts@latest
```

Or clone and build:

```bash
git clone https://github.com/mph-llm-experiments/denote-contacts.git
cd denote-contacts
go build
```

## Usage

```bash
# Run with default contacts directory (~/contacts)
denote-contacts

# Specify a custom directory
denote-contacts ~/my-contacts

# Use the DENOTE_CONTACTS_DIR environment variable
export DENOTE_CONTACTS_DIR=~/my-contacts
denote-contacts
```

## Contact File Format

Contacts are stored as markdown files with YAML frontmatter:

```yaml
---
title: Jane Smith
identifier: 20240715T093045
date: 2024-07-15
tags: [contact]
email: jane@example.com
phone: 555-0123
company: Acme Corp
role: Senior Engineer
location: Portland, OR
relationship_type: work
contact_style: periodic
custom_frequency_days: 30
state: ok
last_contacted: 2024-07-01T10:30:00Z
---
## Notes

Met at tech conference...
```

### File Naming

Files follow the Denote convention:

```
YYYYMMDDTHHMMSS--kebab-case-name__contact.md
```

Example: `20240715T093045--jane-smith__contact.md`

## Keyboard Controls

### List View

- **Navigation**
  - `j/↓` - Move down
  - `k/↑` - Move up
  - `g/Home` - Go to top
  - `G/End` - Go to bottom
  - `Ctrl+d` - Page down
  - `Ctrl+u` - Page up

- **Actions**
  - `Enter` - View contact details
  - `d` - Log interaction (contacted)
  - `s` - Quick state change
  - `T` - Quick type change
  - `b` - Bump (mark as reviewed)
  - `e` - Edit contact
  - `c` - Create new contact
  - `/` - Search
  - `f` - Filter
  - `q` - Quit

### Detail View

- `e` - Edit contact
- `d` - Log interaction
- `b` - Bump contact
- `q/Esc` - Back to list

### Filter Options

Press `f` from the list view to filter. Select one option to immediately apply:

- **By Type**: (f)amily, (c)lose, (n)etwork, (w)ork, (r)ecruiters, (p)roviders, (s)ocial
- **By State**: (F)ollow up, (P)ing, (S)cheduled, (T)imeout
- **By Status**: (o)verdue, (d)ue soon, (g)ood timing
- **Clear**: (a) - Show all contacts

## Contact Types & Default Frequencies

When using `contact_style: periodic`, these defaults apply:

- **close** - 30 days
- **family** - 30 days
- **work** - 60 days
- **network** - 90 days
- **social** - No default
- **providers** - No default
- **recruiters** - No default

Override with `custom_frequency_days` in the frontmatter.

## Contact Styles

- **periodic** - Regular check-ins based on frequency
- **ambient** - Passive monitoring, no reminders
- **triggered** - Event-based contact

## Contact States

- **ok** - Up to date
- **followup** - Need to follow up
- **ping** - Send a quick check-in
- **scheduled** - Meeting/call is scheduled
- **timeout** - No response, needs attention

## Task Integration

When a contact's state changes to one requiring action (followup, ping, scheduled, timeout), denote-contacts automatically creates a task in [denote-tasks](https://github.com/pdxmph/denote-tasks) format. Tasks are created in `~/notes` by default and include:

- Link to the contact via `contact_id` field
- Appropriate action verb (Follow up with, Ping, Meeting with, etc.)
- Same label as the contact (if set)
- Tagged with `task` and `contact-{state}`

## Status Indicators

- **●** (red) - Overdue
- **!** (yellow) - Due soon (within 7 days)
- **●** (green) - Good timing (recently contacted)
- **○** (gray) - OK / No frequency set

## Tips

1. Use tags to group contacts (e.g., `#portland`, `#conference`, `#client`)
2. Set `contact_style: ambient` for contacts you only reach out to when needed
3. Use the bump feature (`b`) to acknowledge you've thought about a contact without logging an interaction
4. Quick filters are your friend - learn the hotkeys for fast navigation
5. The search (`/`) is fuzzy and searches names, companies, emails, tags, and roles

## Contributing

Pull requests welcome! Please keep in mind this tool's focused purpose - it's specifically for contact management, not a general notes system.

## License

MIT
