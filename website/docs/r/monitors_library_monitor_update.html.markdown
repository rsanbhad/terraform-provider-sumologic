---
layout: "sumologic"
page_title: "SumoLogic: sumologic_monitors_library_monitor_update"
description: |-
  Provides a Sumologic Monitors LibraryMonitor Update
---

# sumologic_monitorsLibraryMonitorUpdate
Provider to manage Sumologic Monitors LibraryMonitor Updates

## Example Usage
```hcl
resource "sumologic_monitors_library_monitor_update" "example_monitors_library_monitor_update" {
    description = ""
    monitor_type = ""
    name = ""
    notifications = ""
    triggers = ""
    queries = ""
    version = ""
    type = ""
}
```
## Argument reference

The following arguments are supported:

- `description` - (Required) The description of the monitor or folder.
- `monitor_type` - (Required) Monitor type of the monitor. Valid values: 1. `Logs`: A log based monitor. 2. `Metrics`: A metrics based monitor.
- `name` - (Required) The name of the monitor or folder.
- `notifications` - (Optional) The notifications that the monitor will send when the respective trigger condition is met.
- `triggers` - (Optional) Defines the conditions on which the monitor will trigger notifications.
- `queries` - (Optional) All queries relating to the monitor.
- `version` - (Required) The version of the monitor or folder.
- `type` - (Required) Type of the object model. Valid values:
  1) MonitorsLibraryMonitorWithVersion

The following attributes are exported:

- `id` - The internal ID of the monitors_library_monitor_update



[Back to Index][0]

[0]: ../README.md
