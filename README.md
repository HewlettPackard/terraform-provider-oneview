# HPE OneView SDK for Terraform Provider 

## Build Status 

OV Version | 5.60 | 5.50 | 5.40 | 5.30 |
| ------------- |:-------------:| -------------:| -------------:| -------------:|
SDK Version/Tag | [v1.7.0-0.11](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v1.7.0) |[v1.6.0](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v1.6.0) | [v1.5.0](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v1.5.0) | [v1.4.0](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v1.4.0) |
Build Status | ![Build status](https://ci.appveyor.com/api/projects/status/u84505l6syp70013?svg=true) | ![Build status](https://ci.appveyor.com/api/projects/status/u84505l6syp70013?svg=true) | ![Build status](https://ci.appveyor.com/api/projects/status/u84505l6syp70013?svg=true) | ![Build status](https://ci.appveyor.com/api/projects/status/u84505l6syp70013?svg=true)|

## Introduction

HPE OneView makes it simple to deploy and manage today’s complex hybrid cloud infrastructure. HPE OneView can help you transform your data center to software-defined, and it supports HPE’s broad portfolio of servers, storage, and networking solutions, ensuring the simple and automated management of your hybrid infrastructure. Software-defined intelligence enables a template-driven approach for deploying, provisioning, updating, and integrating compute, storage, and networking infrastructure.

The HPE OneView Terraform SDK enables developers to easily build integrations and scalable solutions with HPE OneView and HPE Image Streamer.

You can find the latest supported HPE OneView Terraform Provider SDK [here](https://github.com/HewlettPackard/terraform-provider-oneview/releases/latest)

## What's New

HPE OneView Terraform SDK library extends support of the SDK to OneView REST API version 2400 (OneView v5.60)

Please refer to [notes](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/CHANGELOG.md) for more information on the changes , features supported and issues fixed in this version

## Getting Started 

## Installation and Configuration

## Installing `terraform-provider-oneview` with Go

HPE OneView SDK for Terraform can be installed from Source or Docker container installation methods. You can either use a docker container which will have the HPE OneView SDK for terraform installed or perform local installation manually.

### Docker Setup
The light weight containerized version of the HPE OneView SDK for Terraform is available in the [Docker Store](https://store.docker.com/community/images/hewlettpackardenterprise/hewlettpackardenterprise/hpe-oneview-sdk-for-terraform). The Docker Store image tag consist of two sections: <sdk_version-OV_version>

```bash
# Download and store a local copy of oneview-sdk-for-terraform and use it as a Docker Image.
$ docker pull hewlettpackardenterprise/hpe-oneview-sdk-for-terraform:v1.7.0-0.11-OV5.6
# Run docker commands below given, which  will in turn create a sh session 
# where you can create files, issue commands and execute the examples.
$ docker run -it docker pull hewlettpackardenterprise/hpe-oneview-sdk-for-terraform:v1.7.0-0.11-OV5.6 /bin/sh
```

### Local Setup

Local installation requires
- Installing Go
```bash 
$ apt-get install build-essential git wget
$ wget https://dl.google.com/go/go1.11.3.linux-amd64.tar.gz

# unzip and untar the file 
$ tar -zxvf go1.11.linux-amd64.tar.gz

# move it to /usr/local/ and create directory for Go.
$ mv go/ /usr/local/ 
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
Note: This provider SUPPORT Terraform 0.12 or above on Terraform-0.12 branch.
```

- Install Oneview Terraform Provider SDK

```go
# Download the source code for terraform-provider-oneview
# Build the needed binary

$ go get github.com/HewlettPackard/terraform-provider-oneview
$ cd $GOPATH/src/github.com/HewlettPackard/terraform-provider-oneview    
$ go build -o terraform-provider-oneview
```

## Configuration

### Environment Variables

Following environment variables can be set for testing:

```bash
# Required
$ export TF_VAR_endpoint=<ov_endpoint>
$ export TF_VAR_username=<ov_username>
$ export TF_VAR_password=<ov_password>
$ export TF_VAR_ssl_enabled=false
$ export TF_VAR_ov_domain=<ov_domain>
```

### OneView Client Configuration

The OneView Client configuration options that can be passed during OneView Client object creation:
The variables are defined in [variable.tf](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/variables.tf) file.

For OneView authentication, you need to provide the provider information in examples:

```terraform
# Create terraform OneView client
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

:lock: Tip: Check the file permissions because the password is stored in clear-text as Environment Variable.

### Image Streamer Client Configuration: 

The Image Streamer (I3S) client is very much similar to the OneView client. 
Following extra environment variables should be set for testing:

```bash
# Required
$ export TF_VAR_i3s_endpoint=<i3s_endpoint>
```
Here we create the Image Streamer(I3S) client.

```terraform
# Create I3s Client
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

:lock: Tip: Check the file permissions because the password is stored in clear-text as Environment Variable.

### Testing the Terraform Resources: 

In the home directory of project(terraform-provider-oneview) user needs to create (or) copy  the file that is to be executed. 

Sample example manifests are available in [example](https://github.com/HewlettPackard/terraform-provider-oneview/tree/master/examples) directory. 

The following terraform commands has to be executed to test the example. 
```terraform
$ terraform init 
$ terraform plan 
$ terraform apply
```
Note: Only a single terraform file (example file) should exist in the home folder to execute the above mentioned three commands. Once the resource is tested move that file to examples folder. 

Note: Currently this SDK supports OneView API 2400 minimally, where we can test OneView API 2400 version with this SDK. If API version used is not supported then error will be thrown. If API version is not provided then appliance's maximum supported API version will be used. 

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


### HPE OneView Community

[HPE OneView Community Forums](http://hpe.com/info/oneviewcommunity)

Learn more about HPE OneView at [hpe.com/info/oneview](https://hpe.com/info/oneview)

