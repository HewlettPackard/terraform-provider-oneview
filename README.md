# terraform-provider-oneview
A Terraform provider for oneview

## Installation 
```
  go build -o terraform-provider-oneview
```
Move terraform-provider-oneview into the directory that Terraform is installed in

## Example terraform file 
```
provider "oneview" {
  username   = "Administrator"
  password   = "thisisapassword"
  endpoint   = "https://oneview_instance.com"
  sslverify  = false
  apiversion = 200
}

resource "oneview_server_profile" "test" {
  name             = "test"
  server_template  = "Web Server Template"
}
```
### License

This project is licensed under the Apache 2.0 license. Please see LICENSE for more info.
