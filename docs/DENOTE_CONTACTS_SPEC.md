# Denote Contacts Specification

## Overview

This document specifies the Denote format for contact records as part of a unified Denote-based personal information management system. Contacts are stored as Markdown files with YAML frontmatter, following Denote's file naming conventions.

## File Naming Convention

Contact files follow the standard Denote naming scheme:

```
YYYYMMDD--kebab-case-name__contact.md
```

Example: `20250712--sarah-chen__contact.md`

Components:
- **Date identifier**: `YYYYMMDD` format (contact creation date)
- **Double dash separator**: `--`
- **Kebab-case title**: Contact name in lowercase with hyphens
- **Double underscore separator**: `__`
- **Tag**: Always `contact` for contact files
- **Extension**: `.md` for Markdown

## YAML Frontmatter Schema

### Required Fields

```yaml
---
title: string          # Full contact name
date: YYYY-MM-DD      # Contact creation date
tags: [contact]       # Must include 'contact' tag
identifier: YYYYMMDD  # Denote identifier matching filename
---
```

### Optional Contact Fields

```yaml
email: string              # Email address
phone: string              # Phone number
company: string            # Company/organization
relationship_type: string  # Type: close, family, network, work
state: string              # Contact state (ok, needs_attention, etc.)
label: string              # Custom label/tag (e.g., @mentor)
contact_style: string      # periodic, ambient, triggered
custom_frequency_days: int # Override default contact frequency
last_contacted: YYYY-MM-DD # Last interaction date
last_bump_date: YYYY-MM-DD # Last review date
bump_count: int            # Number of times reviewed
follow_up_date: YYYY-MM-DD # Scheduled follow-up
deadline_date: YYYY-MM-DD  # Contact-related deadline
archived: boolean          # Archive status
archived_at: YYYY-MM-DD    # Archive date
basic_memory_url: string   # External reference URL
updated_at: YYYY-MM-DD     # Last modification date
```

### Relationship Types

- `close`: Close friends, inner circle (30-day default frequency)
- `family`: Family members (30-day default frequency)
- `network`: Professional network (90-day default frequency)
- `work`: Work colleagues (60-day default frequency)

### Contact Styles

- `periodic`: Regular check-ins based on relationship type
- `ambient`: Passive contact, no active outreach needed
- `triggered`: Contact only when specific events occur

## Body Content Structure

The Markdown body contains:

### Notes Section
General notes about the contact, background information, or context.

```markdown
## Notes

[Free-form text about the contact]
```

### Recent Interactions Section
Chronological log of interactions with the contact.

```markdown
## Recent Interactions

### YYYY-MM-DD HH:MM - [Type]

[Interaction notes]
```

Interaction types include:
- Email
- Call
- Text
- Meeting
- Social
- Bump (review without contact)
- Note (observation without contact)

## Integration with Denote Ecosystem

### Task Creation
Contacts can trigger task creation in denote-tasks through:
- Follow-up dates
- Deadline dates
- Overdue contact reminders

### Cross-References
Contacts can be referenced in:
- Notes (denote-tui) using `[[YYYYMMDD--contact-name]]` links
- Tasks (denote-tasks) for contact-related actions
- Projects for stakeholder tracking

### File Organization
All Denote files (notes, tasks, contacts) can coexist in a single directory:
- Filter by tags: `_note`, `_task`, `_contact`, `_project`
- Use Denote search to find specific content types
- Maintain unified timestamp-based ordering

## Example Contact File

```markdown
---
title: Sarah Chen
date: 2025-07-12
tags: [contact]
identifier: 20250712
email: sarah.chen@techstartup.com
phone: 555-0101
company: Tech Startup Inc
relationship_type: close
state: ok
label: @mentor
contact_style: periodic
last_contacted: 2025-07-10
updated_at: 2025-07-12
---

## Notes

Close friend from college, now CTO at Tech Startup Inc. Great mentor for technical decisions and career advice. Interested in AI/ML applications.

## Recent Interactions

### 2025-07-10 14:30 - Meeting

Coffee meeting to discuss new ML project. She offered to review our architecture proposal.

### 2025-06-15 10:00 - Email

Sent article about transformer models. She responded with helpful insights about implementation challenges.
```

## Implementation Notes

1. **File Creation**: Use contact creation date for filename, not current date
2. **Name Sanitization**: Convert to lowercase, replace spaces with hyphens, remove special characters
3. **Null Handling**: Omit optional fields if empty rather than including null values
4. **Date Format**: Use `YYYY-MM-DD` throughout for consistency
5. **Tag Standardization**: Always use lowercase `contact` tag

## Future Extensions

Potential additional fields for enhanced functionality:
- `social_media`: Object containing platform handles
- `address`: Physical address information
- `birthday`: For personal relationship management
- `timezone`: For scheduling across time zones
- `preferred_contact_method`: Email, phone, text, etc.
- `projects`: Array of related project identifiers
- `skills`: Array of expertise areas for network search