---
layout: "sumologic"
page_title: "SumoLogic: sumologic_scheduled_view"
description: |-
  Provides a Sumologic Scheduled View
---

# sumologic_scheduledView
Provider to manage Sumologic Scheduled Views

## Example Usage
```hcl
resource "sumologic_scheduled_view" "example_scheduled_view" {
    parsing_mode = "AutoParse"
    query = "_sourcecategory=*/Apache"
    retention_period = "365"
    index_name = "TestScheduledView"
    start_time = ""
}
```
## Argument reference

The following arguments are supported:

- `parsing_mode` - (Optional) Define the parsing mode to scan the JSON format log messages. Possible values are:
  1. `AutoParse`
  2. `Manual`
In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
- `query` - (Required) The query that defines the data to be included in the scheduled view.
- `retention_period` - (Optional) The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled.
- `index_name` - (Required) Name of the index for the scheduled view.
- `start_time` - (Required) Start timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.

The following attributes are exported:

- `id` - The internal ID of the scheduled_view



[Back to Index][0]

[0]: ../README.md
