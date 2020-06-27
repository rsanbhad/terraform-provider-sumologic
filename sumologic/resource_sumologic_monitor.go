package sumologic

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSumologicMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicMonitorCreate,
		Read:   resourceSumologicMonitorRead,
		Update: resourceSumologicMonitorUpdate,
		Delete: resourceSumologicMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"version": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"createdAt": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"modifiedAt": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"createdBy": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"modifiedBy": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"parentId": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
				Default:  "0000000000000002",
			},
			"contentType": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
				Default:  "Monitor",
			},
			"isLocked": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"isSystem": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"isMutable": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"monitorType": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
				Default:  "Metrics",
			},
			"notifications": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"queries": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
			"triggers": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
		},
	}
}

func resourceSumologicMonitorRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	id := d.Id()
	monitor, err := c.GetMonitor(id)

	if err != nil {
		return err
	}

	if monitor == nil {
		log.Printf("[WARN] Monitor not found, removing from state: %v - %v", id, err)
		d.SetId("")
		return nil
	}

	d.Set("name", monitor.Name)
	d.Set("type", monitor.Type)
	d.Set("description", monitor.Description)
	d.Set("version", monitor.Version)
	d.Set("createdAt", monitor.CreatedAt)
	d.Set("createdBy", monitor.CreatedBy)
	d.Set("modifiedAt", monitor.ModifiedAt)
	d.Set("modifiedBy", monitor.ModifiedBy)
	d.Set("parentId", monitor.ParentID)
	d.Set("contentType", monitor.ContentType)
	d.Set("isLocked", monitor.IsLocked)
	d.Set("isSystem", monitor.IsSystem)
	d.Set("isMutable", monitor.IsMutable)
	d.Set("monitorType", monitor.MonitorType)
	d.Set("notifications", monitor.Notifications)
	d.Set("queries", monitor.Queries)
	d.Set("triggers", monitor.Triggers)

	return nil
}

func resourceSumologicMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)
	return c.DeleteMonitor(d.Id())
}

func resourceSumologicMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	if d.Id() == "" {
		monitor := resourceToMonitor(d)
		id, err := c.CreateMonitor(monitor)

		if err != nil {
			return err
		}

		d.SetId(id)
	}

	return resourceSumologicMonitorRead(d, meta)
}

func resourceSumologicMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	monitor := resourceToMonitor(d)

	err := c.UpdateMonitor(monitor)

	if err != nil {
		return err
	}

	return resourceSumologicMonitorRead(d, meta)
}

func resourceToMonitor(d *schema.ResourceData) Monitor {
	id, _ := strconv.Atoi(d.Id())

	return Monitor{
		ID:            d.Id(),
		Name:          d.Get("name").(string),
		Type:          d.Get("type").(string),
		Description:   d.Get("description").(string),
		Version:       d.Get("version").(string),
		CreatedAt:     d.Get("createdAt").(string),
		CreatedBy:     d.Get("createdBy").(string),
		ModifiedAt:    d.Get("modifiedAt").(string),
		ModifiedBy:    d.Get("modifiedBy").(string),
		ParentID:      d.Get("parentId").(string),
		ContentType:   d.Get("contentType").(string),
		IsLocked:      d.Get("isLocked").(string),
		IsSystem:      d.Get("isSystem").(string),
		IsMutable:     d.Get("isMutable").(string),
		MonitorType:   d.Get("monitorType").(string),
		Notifications: d.Get("notifications").(string),
		Queries:       d.Get("queries").(string),
		Triggers:      d.Get("triggers").(string),
	}
}
