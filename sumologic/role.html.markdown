---
layout: "sumologic"
page_title: "SumoLogic: sumologic_role"
description: |-
  Provides a Sumologic Role
---

# sumologic_role
Provider to manage Sumologic Roles

## Example Usage
```hcl
resource "sumologic_role" "example_role" {
    name = "DataAdmin"
    description = "Manage data of the org."
    capabilities = "["manageContent","manageDataVolumeFeed","manageFieldExtractionRules","manageS3DataForwarding"]"
    filter_predicate = "!_sourceCategory=billing"
}
```
## Argument reference

The following arguments are supported:

- `name` - (Required) Name of the role.
- `description` - (Optional) Description of the role.
- `capabilities` - (Optional) List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are
### Data Management
  - viewCollectors
  - manageCollectors
  - manageBudgets
  - manageDataVolumeFeed
  - viewFieldExtraction
  - manageFieldExtractionRules
  - manageS3DataForwarding
  - manageContent
  - dataVolumeIndex
  - viewConnections
  - manageConnections
  - viewScheduledViews
  - manageScheduledViews
  - viewPartitions
  - managePartitions
  - viewFields
  - manageFields
  - viewAccountOverview
  - manageTokens
  - manageDataStreams

### Metrics
  - manageMonitors
  - metricsTransformation
  - metricsExtraction
  - metricsRules

### Security
  - managePasswordPolicy
  - ipWhitelisting
  - manageAccessKeys
  - manageSupportAccountAccess
  - manageAuditDataFeed
  - manageSaml
  - shareDashboardOutsideOrg
  - manageOrgSettings
  - changeDataAccessLevel

### Dashboards
  - shareDashboardWorld
  - shareDashboardWhitelist

### UserManagement
  - manageUsersAndRoles

### Observability
  - searchAuditIndex
  - auditEventIndex
- `filter_predicate` - (Optional) A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using '!_sourceCategory=billing' as a filter predicate will prevent users assigned to the role from viewing logs from the source category named 'billing'.

The following attributes are exported:

- `id` - The internal ID of the role



[Back to Index][0]

[0]: ../README.md
