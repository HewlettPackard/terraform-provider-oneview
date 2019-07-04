# terraform-provider-oneview

[![Build Status](https://travis-ci.org/HewlettPackard/terraform-provider-oneview.svg?branch=master)](https://travis-ci.org/HewlettPackard/terraform-provider-oneview)

A Terraform provider for oneview

## Installing `terraform-provider-oneview` with Go

* Install Go 1.11. For previous versions, you may have to set your `$GOPATH` manually, if you haven't done it yet.
* Install Terraform 0.11.x or above [from here](https://www.terraform.io/downloads.html) and save it into `/usr/local/bin/terraform` folder (create it if it doesn't exists)
* Download the code by issuing a `go get` command.

```bash
# Download the source code for terraform-provider-oneview
# and build the needed binary, by saving it inside $GOPATH/bin
$ go get -u github.com/HewlettPackard/terraform-provider-oneview

# Copy the binary to have it along the terraform binary
$ mv $GOPATH/bin/terraform-provider-oneview /usr/local/bin/terraform
```

### Provider Information

For authentication, you need to provide the provider information in examples:


```bash
provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}
```

### License

This project is licensed under the Apache 2.0 license.

## Version and changes

To view history and notes for this version, view the [Changelog](CHANGELOG.md).
