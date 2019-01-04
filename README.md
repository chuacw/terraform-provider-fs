# Terraform Provider FileSystem

This is a simple Terraform Provider to create, rename and destroy directories on the local file system.

To use this on Windows:
<pre>
provider "fs" {
}

resource "fs_directory" {
    name = "c:\\newDirectoryName"
}

data "fs_file" "file1" {
  filename = "C:\\Go"
}

output "go-exists" {
  value = "filename: ${data.fs_file.file1.filename}, exists: ${data.fs_file.file1.exists}, isdir: ${data.fs_file.file1.isdir}"
}
</pre>
Note the double backslash requirement.

In the example above, the following 2 things happen:
 - "c:\newDirectoryName" is created.
 - if the directory "C:\Go" exists, then the output for "go-exists" is:
<pre>
go-exists = filename: C:\Go, exists: true, isdir: true
</pre>
