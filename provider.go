package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// FSProvider to return a Terraform provider
func FSProvider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"fs_directory": resourceDirectory(),
			"fs_file":      resourceFile(),
		},
	}
}
