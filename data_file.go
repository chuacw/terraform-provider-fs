package main

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFileRead,

		Schema: map[string]*schema.Schema{
			fsFILENAME: {
				Type:        schema.TypeString,
				Description: "Filename to check for existence",
				Required:    true,
				ForceNew:    false,
			},
			fsEXISTS: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			fsISDIR: {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceFileRead(d *schema.ResourceData, _ interface{}) error {
	path := d.Get(fsFILENAME).(string)
	d.SetId(path)
	if fi, err := os.Stat(path); !os.IsNotExist(err) {
		// path/to/whatever exists
		d.Set(fsEXISTS, true)
		switch mode := fi.Mode(); {
		case mode.IsDir():
			{
				d.Set(fsISDIR, true)
			}
		default:
			d.Set(fsISDIR, false)
		}
		return nil
	}
	d.Set(fsISDIR, false)
	return nil
}
