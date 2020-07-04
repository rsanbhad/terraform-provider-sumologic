// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------
package sumologic

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSumologicMonitorsLibraryMonitorResponse() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicMonitorsLibraryMonitorResponseCreate,
		Read:   resourceSumologicMonitorsLibraryMonitorResponseRead,
		Update: resourceSumologicMonitorsLibraryMonitorResponseUpdate,
		Delete: resourceSumologicMonitorsLibraryMonitorResponseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"post_request_map": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceToMonitorsLibraryMonitorResponse(d *schema.ResourceData) MonitorsLibraryMonitorResponse {
	rawNotifications := d.Get("notifications").([]interface{})
	notifications := make([]string, len(rawNotifications))
	for i, v := range rawNotifications {
		notifications[i] = v.(string)
	}
	rawTriggers := d.Get("triggers").([]interface{})
	triggers := make([]string, len(rawTriggers))
	for i, v := range rawTriggers {
		triggers[i] = v.(string)
	}
	rawQueries := d.Get("queries").([]interface{})
	queries := make([]string, len(rawQueries))
	for i, v := range rawQueries {
		queries[i] = v.(string)
	}

	return MonitorsLibraryMonitorResponse{
		CreatedBy:     d.Get("created_by").(string),
		Name:          d.Get("name").(string),
		ID:            d.Id(),
		CreatedAt:     d.Get("created_at").(string),
		MonitorType:   d.Get("monitor_type").(string),
		Description:   d.Get("description").(string),
		Queries:       queries,
		ModifiedBy:    d.Get("modified_by").(string),
		IsMutable:     d.Get("is_mutable").(bool),
		Version:       d.Get("version").(int),
		Notifications: notifications,
		Type:          d.Get("type").(string),
		ParentId:      d.Get("parent_id").(string),
		ModifiedAt:    d.Get("modified_at").(string),
		Triggers:      triggers,
		ContentType:   d.Get("content_type").(string),
		IsLocked:      d.Get("is_locked").(bool),
		IsSystem:      d.Get("is_system").(bool),
	}
}
