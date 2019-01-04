package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Test_resourceDirectoryCreateUpdateDelete(t *testing.T) {
	directory := `d:\\terraform`
	var cases = []struct {
		path   string
		config string
	}{
		{
			directory,
			fmt.Sprintf(`resource "fs_directory" "dir1" {
               name = "%s"
            }`, directory),
		},
	}

	for _, tt := range cases {
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				{
					Config: tt.config,
					Check: func(s *terraform.State) error {
						if _, err := os.Stat(tt.path); os.IsNotExist(err) {
							return fmt.Errorf("config:\n%s\n,got: %s", tt.config, err)
						}
						return nil
					},
				},
				{
					Config: tt.config,
					Check: func(s *terraform.State) error {
						r := &schema.Resource{
							Schema: map[string]*schema.Schema{
								fsNAME: &schema.Schema{
									Type: schema.TypeString,
								},
							},
							Update: resourceDirectoryUpdate,
						}
						oldname := directory
						newname := oldname + "2"
						s1 := &terraform.InstanceState{
							ID: "foo",
							Attributes: map[string]string{
								fsNAME: oldname,
							},
						}
						d1 := &terraform.InstanceDiff{
							Attributes: map[string]*terraform.ResourceAttrDiff{
								fsNAME: &terraform.ResourceAttrDiff{
									New: newname,
								},
							},
						}
						if newState, err := r.Apply(s1, d1, nil); err != nil {
							t.Fatalf("err: %s", err)
						} else {
							d2 := &terraform.InstanceDiff{
								Attributes: map[string]*terraform.ResourceAttrDiff{
									fsNAME: &terraform.ResourceAttrDiff{
										New: oldname,
									},
								},
							}
							if _, err := r.Apply(newState, d2, nil); err != nil {
								t.Fatalf("err: %s", err)
							}
						}

						return nil
					},
				},
			},
			CheckDestroy: func(*terraform.State) error {
				if _, err := os.Stat(tt.path); os.IsNotExist(err) {
					return nil
				}
				return errors.New("directory did not get destroyed")
			},
		})
	}
}

func Test_resourceDirectoryCreate(t *testing.T) {
	directory := `d:\\terraform`
	var cases = []struct {
		path   string
		config string
	}{
		{
			directory,
			fmt.Sprintf(`resource "fs_directory" "dir1" {
               name = "%s"
            }`, directory),
		},
	}

	for _, tt := range cases {
		r.UnitTest(t, r.TestCase{
			Providers: testProviders,
			Steps: []r.TestStep{
				{
					Config: tt.config,
					Check: func(s *terraform.State) error {
						if _, err := os.Stat(tt.path); os.IsNotExist(err) {
							return fmt.Errorf("config:\n%s\n,got: %s", tt.config, err)
						}
						return nil
					},
				},
			},
			CheckDestroy: func(*terraform.State) error {
				if _, err := os.Stat(tt.path); os.IsNotExist(err) {
					return nil
				}
				return errors.New("directory did not get destroyed")
			},
		})
	}
}

var testProviders = map[string]terraform.ResourceProvider{
	"fs": FSProvider(),
}

func Test_resourceDirectoryUpdate(t *testing.T) {
	// Covered in Test_resourceDirectoryCreateUpdateDelete
}
