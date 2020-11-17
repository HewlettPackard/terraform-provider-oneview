# HPE OneView Terraform Provider bindings
| 5.50 Branch   |
| ------------- |
[![Build Status](https://travis-ci.org/HewlettPackard/terraform-provider-oneview.svg?branch=master)](https://travis-ci.org/HewlettPackard/terraform-provider-oneview)

## Introduction

HPE OneView makes it simple to deploy and manage today’s complex hybrid cloud infrastructure. HPE OneView can help you transform your data center to software-defined, and it supports HPE’s broad portfolio of servers, storage, and networking solutions, ensuring the simple and automated management of your hybrid infrastructure. Software-defined intelligence enables a template-driven approach for deploying, provisioning, updating, and integrating compute, storage, and networking infrastructure.

The HPE OneView Terraform Provider SDK dilivers library to easily interact with HPE OneView and HPE Image Streamer REST APIs. The HPE OneView Go SDK enables developers to easily build integrations and scalable solutions with HPE OneView and HPE Image Streamer.

You can find the latest supported HPE OneView Terraform Provider SDK [here](https://github.com/HewlettPackard/terraform-provider-oneview/releases/latest)

## Installing `terraform-provider-oneview` with Go

HPE OneView SDK for terraform can be installed from Source or Docker container installation methods. You can either use a docker container which will have the HPE OneView SDK for terraform installed or perform local installation manually.

### Docker Container Setup

We also provide a lightweight and easy way to test and run oneview-terraform. The `hewlettpackardenterprise/hpe-oneview-sdk-for-terraform:<tag>` docker image contains an installation of oneview-terraform installation you can use by just pulling down the Docker Image:

The Docker Store image `tag` consist of two sections: `<sdk_version-OV_version>`

Download and store a local copy of hpe-oneview-sdk-for-terraform and use it as a Docker image.
```bash
$ docker pull hewlettpackardenterprise/hpe-oneview-sdk-for-terraform:v1.6.0-OV5.5
```
Run docker commands and this will in turn create a sh session where you can create files, issue commands and execute the tests
```bash
$ docker run -it hewlettpackardenterprise/hpe-oneview-sdk-for-terraform:v1.6.0-OV5.5 /bin/sh
```

### Local Setup

Local installation requires 
- Installing Go
```bash 
$ apt-get install build-essential git wget
$ wget https://dl.google.com/go/go1.11.3.linux-amd64.tar.gz

#unzip and untar the file 
$ tar -zxvf "go1.11.linux-amd64.tar.gz", 

# move it to /usr/local/ and create directory for Go.
$ mv go1.11.3.linux-amd64.tar.gz /usr/local/ 
$ mkdir ~/go
```
```bash 
# Setting Environment Variable 
$ export GOROOT=/usr/local/go
$ export GOPATH=$HOME/go 
$ export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

- Installing Terraform
Install Terraform 0.11.x [from here](https://www.terraform.io/downloads.html) and save it into `/usr/local/bin/terraform` folder (create it if it doesn't exists).
```bash 
Note: This provider DOES NOT SUPPORT Terraform 0.12 or above.
```

- Install Oneview Terraform Provider SDK
```go
# Download the source code for terraform-provider-oneview
# Build the needed binary

go get github.com/HewlettPackard/terraform-provider-oneview
$ cd $GOPATH/src/github.com/HewlettPackard/terraform-provider-oneview    

go build -o terraform-provider-oneview  
$ mv $GOPATH/bin/terraform-provider-oneview ~/.terraform.d/plugins
```



## Configuration

- OneView Client Configuration
The OneView Client configuration options that can be passed during OneView Client object creation:
The variables are defined in [variable.tf](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/variables.tf) file.

```bash
# Following environment variables can be set for testing:

# Required
$ export TF_VAR_endpoint=<ov_endpoint>
$ export TF_VAR_username=<ov_username>
$ export TF_VAR_password=<ov_password>
$ export TF_VAR_ssl_enabled=false
$ export TF_VAR_ov_domain=<ov_domain>
```

```bash
#For authentication, you need to provide the provider information in examples:

provider "oneview" {
	ov_username  = "${var.username}"
	ov_password  = "${var.password}"
	ov_endpoint  = "${var.endpoint}"
	ov_sslverify = "${var.ssl_enabled}"
	ov_domain    = "${var.ov_domain}"
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

```

- Image Streamer Client Configuration: 


```bash
# The Image Streamer (I3S) client is very much similar to the OneView client. 
# Following extra environment variables should be set for testing:

# Required
$ export TF_VAR_i3s_endpoint=<i3s_endpoint>
```

```bash
# Here we create the Image Streamer(I3S) client.

provider "oneview" {
	ov_username      = "${var.username}"
	ov_password      = "${var.password}"
	ov_endpoint      = "${var.endpoint}"
	ov_i3s_endpoint  = "${var.i3s_endpoint}"
	ov_sslverify     = "${var.ssl_enabled}"
	ov_domain        = "${var.ov_domain}"
	ov_apiversion    = <i3s_apiversion>
	ov_ifmatch = "*"
}
```

### Testing the resources: 
In the home directory of project(terraform-provider-oneview) create the file which is to be executed. 
Find sample example manifests in [example](https://github.com/HewlettPackard/terraform-provider-oneview/tree/master/examples) directory. 

The following commands has to be executed to test the example. 
```terraform
$ terraform init 
$ terraform plan 
$ terraform apply
```
Note: Only a single terraform file (example file) should exist in the home folder to execute the above mentioned three commands. Once the resource is tested move that file to examples folder. 

Note: Currently this SDK supports OneView API 2200 minimally, where we can test OneView API 2200 version with this SDK. If API version used is not supported then error will be thrown. If API version is not provided then appliance's maximum supported API version will be used. 

## API Implementation

A detailed list of the HPE OneView REST interfaces that have been implemented in this SDK can be found in the [endpoints-support.md](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/endpoints-support.md) file.

## Getting Help 

Are you running into a road block? Have an issue with unexpected behavior? Feel free to open a new issue on the [issue tracker](https://github.com/HewlettPackard/terraform-provider-oneview/issues/new)

## License
This project is licensed under the Apache 2.0 license. Please see [LICENSE](LICENSE) for more info.

## Contributing and feature requests

We welcome your contributions to the HPE OneView for Terraform Provider SDK. 

**Contributing:** You know the drill. Fork it, branch it, change it, commit it, and pull-request it.
We are passionate about improving this project, and glad to accept help to make it better.
For more information refer [CONTRIBUTING.md](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/CONTRIBUTING.md) file.

NOTE: We reserve the right to reject changes that we feel do not fit the scope of this project, so for feature additions, please open an issue to discuss your ideas before doing the work.

**Feature Requests:** If you have a need that is not met by the current implementation, please let us know opening an new enhancement request/issue.
This feedback is important for us to deliver a useful product. 

## Additional Resources 

### HPE OneView Documentation

[HPE OneView Release Notes](http://hpe.com/info/OneView/docs)

[HPE OneView Support Matrix](http://hpe.com/info/OneView/docs)

[HPE OneView Installation Guide](http://hpe.com/info/OneView/docs)

[HPE OneView User Guide](http://hpe.com/info/OneView/docs)

[HPE OneView Online Help](http://hpe.com/info/OneView/docs)

[HPE OneView REST API Reference](http://hpe.com/info/OneView/docs)

[HPE OneView Firmware Management White Paper](http://hpe.com/info/OneView/docs)

[HPE OneView Deployment and Management White Paper](http://hpe.com/info/OneView/docs)

### HPE OneView Community

[HPE OneView Community Forums](http://hpe.com/info/oneviewcommunity)

Learn more about HPE OneView at [hpe.com/info/oneview](https://hpe.com/info/oneview)


## Hacking Guide

Use the [hacking guide](HACKING.md) to run local acceptance testing and debugging local contributions. Currently test cases are not supported portion of our CI/CD approval process but might be made available from this test suite in the future.  Contributions to the test suite is appreciated.

## License
This project is licensed under the Apache License, Version 2.0.  See LICENSE for full license text.
