package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	// name := d.Get(fsNAME).(string)
	// d.SetId(name)
	return resourceDirectoryRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFileUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFileRead(d, m)
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Update: resourceFileUpdate,
		Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{
			fsNAME: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
