---
layout: "sumologic"
page_title: "SumoLogic: sumologic_partition"
description: |-
  Provides a Sumologic Partition
---

# sumologic_partition
Provider to manage Sumologic Partitions

## Example Usage
```hcl
resource "sumologic_partition" "example_partition" {
    routing_expression = "_sourcecategory=*/Apache"
    is_compliant = "false"
    analytics_tier = "enhanced"
    name = "apache"
    retention_period = "365"
}
```
## Argument reference

The following arguments are supported:

- `routing_expression` - (Required) The query that defines the data to be included in the partition.
- `is_compliant` - (Optional) Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later.
- `analytics_tier` - (Optional) The Cloud Flex analytics tier for your data; only relevant if your account has basic analytics enabled. Possible values are:
  1. `enhanced`
  2. `basic`
  3. `cold`
- `name` - (Optional) The name of the partition.
- `retention_period` - (Optional) The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled.

The following attributes are exported:

- `id` - The internal ID of the partition



[Back to Index][0]

[0]: ../README.md
