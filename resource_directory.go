package main

import (
	"os"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDirectoryCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get(fsNAME).(string)
	fileMode := os.FileMode(0777)
	if mode1, ok := d.GetOk(fsMODE); ok {
		if mode2, err := strconv.ParseInt(mode1.(string), 16, 64); err == nil {
			fileMode = os.FileMode(mode2)
		}
	}
	if err := os.Mkdir(name, fileMode); err != nil {
		return err
	}
	d.SetId(name)
	return nil
}

func resourceDirectoryRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDirectoryUpdate(d *schema.ResourceData, m interface{}) error {
	oldName1, newName1 := d.GetChange(fsNAME)
	oldName := oldName1.(string)
	newName := newName1.(string)
	if oldName == "" {
		return resourceDirectoryCreate(d, m)
	}
	err := os.Rename(oldName, newName)
	return err
}

func resourceDirectoryDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get(fsNAME).(string)
	if err := os.Remove(name); err != nil {
		return err
	}
	return nil
}

func resourceDirectorySchema() map[string]*schema.Schema {
	result := map[string]*schema.Schema{
		fsNAME: &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		fsMODE: &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return result
}

func resourceDirectory() *schema.Resource {
	return &schema.Resource{
		Create: resourceDirectoryCreate,
		Read:   resourceDirectoryRead,
		Update: resourceDirectoryUpdate,
		Delete: resourceDirectoryDelete,

		Schema: resourceDirectorySchema(),
	}
}
