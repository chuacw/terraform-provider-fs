# Terraform Provider FileSystem

This is a simple Terraform Provider to create, rename and destroy directories on the local file system.

To use this on Windows:
<pre>
provider "fs" {
}

resource "fs_directory" {
    name = "c:\\newDirectoryName"
}
</pre>
Note the double backslash requirement.
