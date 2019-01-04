package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func Test_dataSourceFileRead(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	directory := usr.HomeDir
	if strings.Contains(directory, `\`) {
		directory = strings.Replace(directory, `\`, `\\`, -1)
	}
	executable := os.Args[0]
	if strings.Contains(executable, `\`) {
		executable = strings.Replace(executable, `\`, `\\`, -1)
	}
	var cases = []struct {
		path                string
		config              string
		key                 string
		wantExists, wantDir bool
	}{
		{
			directory,
			fmt.Sprintf(`data "fs_file" "dir1" {
               filename = "%s"
			}`, directory),
			"data.fs_file.dir1",
			true, true,
		},
		{
			directory + "2",
			fmt.Sprintf(`data "fs_file" "dir2" {
               filename = "%s"
			}`, directory+"2"),
			"data.fs_file.dir2",
			false, false,
		},
		{
			executable,
			fmt.Sprintf(`data "fs_file" "dir3" {
               filename = "%s"
			}`, executable),
			"data.fs_file.dir3",
			true, false,
		},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				Providers: testProviders,
				Steps: []resource.TestStep{
					{
						Config: test.config,
						Check: func(s *terraform.State) error {
							m := s.RootModule()
							i := m.Resources[test.key].Primary
							filename := i.Attributes[fsFILENAME]
							existsStr := i.Attributes[fsEXISTS]
							exists, _ := strconv.ParseBool(existsStr)

							isDirStr := i.Attributes[fsISDIR]
							isDir, _ := strconv.ParseBool(isDirStr)
							fmt.Printf("File: %s, exists: %v, isDir: %v\n", filename, exists, isDir)
							if isDir != test.wantDir || exists != test.wantExists {
								t.Fatalf("Expected: %v, %v, got: %v, %v", test.wantDir, test.wantExists, isDir, exists)
							}
							return nil
						},
					},
				},
			})
		})
	}
}
