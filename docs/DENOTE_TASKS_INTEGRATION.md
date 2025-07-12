# Denote Tasks Integration Guide

## Overview

This guide describes how denote-contacts integrates with denote-tasks to create a unified personal information management system using Denote as the underlying data structure.

## Integration Architecture

### Shared Concepts

1. **Denote File Format**
   - Both apps use Markdown files with YAML frontmatter
   - Consistent naming: `YYYYMMDD--title__tag.md`
   - Tags distinguish content types: `_contact`, `_task`

2. **Cross-References**
   - Tasks can reference contacts: `contact_id: 20250712`
   - Contacts can have associated tasks: `task_ids: [20250713, 20250714]`
   - Use Denote links: `[[20250712--sarah-chen]]`

3. **Shared Directory**
   - All Denote files coexist in one directory
   - Filter by tags to show specific content types
   - Unified timestamp-based ordering

## Task Creation from Contacts

### Automatic Task Generation

Contacts can trigger task creation based on:

1. **Overdue Contacts**
   ```yaml
   title: "Follow up with Sarah Chen"
   tags: [task, contact-followup]
   contact_id: 20250712
   due_date: 2025-07-15
   ```

2. **Follow-up Dates**
   ```yaml
   title: "Schedule meeting with Sarah Chen"
   tags: [task, contact-meeting]
   contact_id: 20250712
   due_date: 2025-07-20
   ```

3. **Contact Deadlines**
   ```yaml
   title: "Complete project review for Sarah Chen"
   tags: [task, contact-deadline]
   contact_id: 20250712
   due_date: 2025-07-25
   ```

### Manual Task Creation

From the contacts-tui interface:
- Press `t` on a contact to create a related task
- Task inherits contact context
- Automatically links back to contact

## Data Synchronization

### Contact State Updates

When a task is completed in denote-tasks:
1. Check if task has `contact_id`
2. Update contact's `last_contacted` if task type is `contact-followup`
3. Add interaction log entry to contact
4. Clear follow-up date if applicable

### Task State Updates

When a contact is marked as contacted:
1. Find related open tasks
2. Mark follow-up tasks as completed
3. Update task notes with interaction details

## Implementation Interface

### denote-tasks Should Expose

```go
// TaskService interface for contacts integration
type TaskService interface {
    // Create a new task linked to a contact
    CreateContactTask(contactID string, taskType string, title string, dueDate time.Time) (string, error)
    
    // Get all tasks for a contact
    GetContactTasks(contactID string) ([]Task, error)
    
    // Complete a contact-related task
    CompleteContactTask(taskID string, notes string) error
    
    // Update contact interaction from task
    NotifyContactInteraction(contactID string, interactionType string, notes string) error
}
```

### denote-contacts Should Expose

```go
// ContactService interface for task integration
type ContactService interface {
    // Get contact by Denote ID
    GetContactByDenoteID(denoteID string) (*Contact, error)
    
    // Update contact from task completion
    UpdateContactFromTask(denoteID string, taskType string, completionNotes string) error
    
    // Get overdue contacts for task generation
    GetOverdueContacts() ([]Contact, error)
    
    // Check if contact needs follow-up
    NeedsFollowUp(denoteID string) (bool, *time.Time, error)
}
```

## File Format Extensions

### Task File with Contact Reference

```markdown
---
title: Follow up with Sarah Chen
date: 2025-07-13
tags: [task, contact-followup]
identifier: 20250713
status: pending
priority: medium
due_date: 2025-07-15
contact_id: 20250712
contact_name: Sarah Chen
---

## Description

Quarterly check-in with Sarah to discuss:
- Project progress
- Team updates
- Future collaboration opportunities

## Context

Last contacted: 2025-04-10 (95 days ago)
Relationship type: close (30-day frequency)
```

### Contact File with Task References

```markdown
---
title: Sarah Chen
date: 2025-07-12
tags: [contact]
identifier: 20250712
# ... other contact fields ...
active_tasks: [20250713, 20250714]
completed_tasks: [20250601, 20250515]
---
```

## UI Integration Points

### In denote-contacts

1. **Task Indicator**
   - Show task count next to contact name
   - Highlight contacts with pending tasks
   - Quick key to view related tasks

2. **Task Creation**
   - `t`: Create general task
   - `f`: Create follow-up task
   - `m`: Create meeting task
   - `r`: Create reminder task

### In denote-tasks

1. **Contact Context**
   - Show contact name in task list
   - Display relationship type and frequency
   - Show last interaction date

2. **Quick Actions**
   - `c`: View related contact
   - `i`: Log interaction on completion
   - `l`: Link to contact

## Workflow Examples

### Example 1: Overdue Contact Workflow

1. denote-contacts detects Sarah Chen is overdue (90 days)
2. Automatically creates task: "Follow up with Sarah Chen"
3. User completes task in denote-tasks
4. denote-contacts updates last_contacted date
5. Interaction log added to contact file

### Example 2: Project Collaboration

1. User creates task: "Review ML proposal with Sarah"
2. Links task to Sarah Chen contact
3. Task appears in Sarah's contact view
4. Completing task adds project note to contact

### Example 3: Birthday Reminder

1. Contact has birthday field set
2. Annual task auto-created 1 week before
3. Task completion logs "birthday greeting" interaction
4. Next year's task scheduled automatically

## Configuration

### Shared Configuration

```toml
[denote]
directory = "~/Documents/denote"
file_extensions = ["md", "org", "txt"]

[integration]
enable_task_sync = true
enable_contact_sync = true
auto_create_followup_tasks = true
followup_task_lead_time_days = 3

[task_templates]
followup = "Follow up with {contact_name}"
meeting = "Schedule meeting with {contact_name}"
deadline = "Complete {subject} for {contact_name}"
```

## Best Practices

1. **Consistent Tagging**
   - Use hierarchical tags: `task/contact-followup`
   - Maintain tag taxonomy across apps
   - Document custom tags

2. **ID Management**
   - Use Denote identifier as primary key
   - Store as YYYYMMDD format
   - Handle ID conflicts gracefully

3. **Sync Frequency**
   - Real-time for user actions
   - Batch sync for automated tasks
   - Handle offline scenarios

4. **Data Integrity**
   - Validate cross-references
   - Handle deleted contacts/tasks
   - Maintain audit trail

## Future Enhancements

1. **Bi-directional Sync**
   - Full CRDT-based conflict resolution
   - Multi-device synchronization
   - Real-time collaboration

2. **Advanced Automation**
   - Rule-based task generation
   - Smart follow-up scheduling
   - Relationship health monitoring

3. **Analytics Integration**
   - Contact/task correlation metrics
   - Productivity insights
   - Network analysis