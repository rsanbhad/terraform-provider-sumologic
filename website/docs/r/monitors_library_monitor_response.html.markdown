---
layout: "sumologic"
page_title: "SumoLogic: sumologic_monitors_library_monitor_response"
description: |-
  Provides a Sumologic Monitors LibraryMonitor Response
---

# sumologic_monitorsLibraryMonitorResponse
Provider to manage Sumologic Monitors LibraryMonitor Responses

## Example Usage
```hcl
resource "sumologic_monitors_library_monitor_response" "example_monitors_library_monitor_response" {
    modified_at = ""
    created_by = ""
    is_locked = ""
    monitor_type = ""
    is_system = ""
    name = ""
    parent_id = ""
    notifications = ""
    type = ""
    version = ""
    is_mutable = ""
    triggers = ""
    queries = ""
    description = ""
    modified_by = ""
    created_at = ""
    content_type = ""
}
```
## Argument reference

The following arguments are supported:

- `modified_at` - (Required) Last modification timestamp in UTC.
- `created_by` - (Required) Identifier of the user who created the resource.
- `is_locked` - (Required) Whether the object is locked.
- `monitor_type` - (Required) Monitor type of the monitor. Valid values: 1. `Logs`: A log based monitor. 2. `Metrics`: A metrics based monitor.
- `is_system` - (Required) System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can't be updated.
- `name` - (Required) Identifier of the monitor or folder.
- `parent_id` - (Required) Identifier of the parent folder.
- `notifications` - (Optional) The notifications that the monitor will send when the respective trigger condition is met.
- `type` - (Required) Type of the object model.
- `version` - (Required) Version of the monitor or folder.
- `is_mutable` - (Required) Immutable objects are 'READ-ONLY'.
- `triggers` - (Optional) Defines the conditions on which the monitor will trigger notifications.
- `queries` - (Optional) All queries relating to the monitor.
- `description` - (Required) Description of the monitor or folder.
- `modified_by` - (Required) Identifier of the user who last modified the resource.
- `created_at` - (Required) Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
- `content_type` - (Required) Type of the content. Valid values:
  1) Monitor
  2) Folder

The following attributes are exported:

- `id` - The internal ID of the monitors_library_monitor_response



[Back to Index][0]

[0]: ../README.md
