---
layout: "sumologic"
page_title: "SumoLogic: sumologic_monitors_library_monitor"
description: |-
  Provides a Sumologic Monitors LibraryMonitor
---

# sumologic_monitorsLibraryMonitor
Provider to manage Sumologic Monitors LibraryMonitors

## Example Usage
```hcl
resource "sumologic_monitors_library_monitor" "example_monitors_library_monitor" {
    monitor_type = ""
    name = ""
    notifications = ""
    triggers = ""
    type = ""
    queries = ""
    description = ""
}
```
## Argument reference

The following arguments are supported:

- `monitor_type` - (Required) Monitor type of the monitor. Valid values: 1. `Logs`: A log based monitor. 2. `Metrics`: A metrics based monitor.
- `name` - (Required) Name of the monitor or folder.
- `notifications` - (Optional) The notifications that the monitor will send when the respective trigger condition is met.
- `triggers` - (Optional) Defines the conditions on which the monitor will trigger notifications.
- `type` - (Required) Type of the object model. Valid values:
  1) MonitorsLibraryMonitor
- `queries` - (Optional) All queries relating to the monitor.
- `description` - (Required) Description of the monitor or folder.

The following attributes are exported:

- `id` - The internal ID of the monitors_library_monitor



[Back to Index][0]

[0]: ../README.md
