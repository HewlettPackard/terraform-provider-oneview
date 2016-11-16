# terraform-provider-oneview
A Terraform provider for oneview

## Installation 

* Install golang 1.6 or better
* Install Terraform 0.7.0 or better
* Create a file structure similar to ~/workspace/go/src/github.com/HewlettPackard/
* Set GOPATH=~/workspace/go

```
  cd ~/workspace/go/src/github.com/HewlettPackard/
  git clone https://github.com/HewlettPackard/terraform-provider-oneview.git
  cd terraform-provider-oneview
  go get
  go build -o terraform-provider-oneview
```
Move terraform-provider-oneview into the directory that Terraform is installed in

## Example terraform file 
```
provider "oneview" {
  ov_username   = "Administrator"
  ov_password   = "thisisapassword"
  ov_endpoint   = "https://oneview_instance.com"
}

resource "oneview_server_profile" "test" {
  name             = "test"
  template         = "Web Server Template"
}
```
### License

This project is licensed under the Apache 2.0 license.
