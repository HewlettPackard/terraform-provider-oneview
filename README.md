# HPE OneView SDK for Terraform Provider

## Build Status 
|      OV Version|    8.30 |    8.20 |    8.10|    8.00|    7.20  |7.10  | 7.00 | 6.60 | 6.50 | 6.40 | 6.30 |6.20 |6.10 | 6.00 |
| --------------:| --------------:| --------------: | --------------: | --------------: |--------------: |--------------: |--------------: |--------------: |--------------: |--------------: |--------------: |--------------: |--------------: |--------------: |
| SDK Version/Tag |[v8.3.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v8.2.0-13)|[v8.3.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v8.2.0-13) |[v8.1.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v8.1.0-13) |[v8.0.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v8.0.0-13) |[v7.2.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v7.2.0-13) | [v7.1.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v7.1.0-13) |  [v7.0.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v7.0.0-13) | [v6.6.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.6.0-13) |  [v6.5.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.5.0-13) | [v6.4.0-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.4.0-13) |   [v6.3.1-13](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.3.1-13) |  [v6.2.0-13 ](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.2.0-13) |   [v6.1.0-13 ](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.1.0-13) |   [v6.0.0-13 ](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.0.0-13) | [v6.0.0-12 ](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v6.0.0-12)  | [v1.7.0-12 ](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v1.7.0-12) |
|    Build Status | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/4423603823)|[![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/4423603823) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/3702892830) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/3410062934) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/2982719597) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/2688385352) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/2311877982) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/2020207864) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/1686945274) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/1488877220) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/1216442369) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/1021599071) | [![Build status](https://github.com/HewlettPackard/terraform-provider-oneview/actions/workflows/test.yml/badge.svg)](https://github.com/HewlettPackard/terraform-provider-oneview/actions/runs/707591356) | [![Build status](https://api.travis-ci.com/HewlettPackard/terraform-provider-oneview.svg)](https://travis-ci.org/github/HewlettPackard/terraform-provider-oneview/builds) | [![Build status](https://api.travis-ci.com/HewlettPackard/terraform-provider-oneview.svg)](https://travis-ci.org/github/HewlettPackard/terraform-provider-oneview/builds) |
```
Note: v6.6.0-13 onwards supports Terraform version from 0.13.xx to 1.x.x
```
## Introduction

HPE OneView makes it simple to deploy and manage today’s complex hybrid cloud infrastructure. HPE OneView can help you transform your data center to software-defined, and it supports HPE’s broad portfolio of servers, storage, and networking solutions, ensuring the simple and automated management of your hybrid infrastructure. Software-defined intelligence enables a template-driven approach for deploying, provisioning, updating, and integrating compute, storage, and networking infrastructure.

The HPE OneView Terraform SDK enables developers to easily build integrations and scalable solutions with HPE OneView.

You can find the latest supported HPE OneView Terraform Provider SDK [here](https://github.com/HewlettPackard/terraform-provider-oneview/releases/tag/v8.3.0-13)

## What's New

HPE OneView Terraform SDK library extends support of the SDK to OneView REST API version 5200 (OneView v8.3)

Please refer to [notes](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/CHANGELOG.md) for more information on the changes , features supported and issues fixed in this version
### Migration to HPE OneView 7.0 from HPE OneView 6.x ( Non-Synergy appliance )
With HPE OneView 7.0 one can migrate his/her server hardware from HPE OneView 6.x using migrate functionality provided in the UI. 
With server hardware migration , corresponding server profile will also get migrated to HPE OneView 7.0 . 

#### Impact after migration
1. Idempotency issues post migration.
Post migration do run terraform refresh to sync with migrated resources.

    Steps: 
    -  Point your <TF_VAR_endpoint> to HPE OneView 8.3.
    -  Run `terraform refresh` to sync with  migrated resource 
  
2. Idempotency issue in server profile after migration.
In HPE OneView 7.0 the server hardware type field is not compatible with HPE OneView 6.x server hardware type. To resolve that, while creating server profile, make sure terraform configuration file is not having server hardware type field defined . 


  

## Getting Started 

## Installation and Configuration

## Installing `terraform-provider-oneview` with Go

HPE OneView SDK for Terraform can be installed from Source or Docker container installation methods. You can either use a docker container which will have the HPE OneView SDK for terraform installed or perform local installation manually.

### Docker Setup
The light weight containerized version of the HPE OneView SDK for Terraform is available in the [Docker Store](https://hub.docker.com/repository/docker/hewlettpackardenterprise/hpe-oneview-sdk-for-terraform/general). The Docker Store image tag consist of two sections: <sdk_version-OV_version>

```bash
# Download and store a local copy of oneview-sdk-for-terraform and use it as a Docker Image.
$ docker pull hewlettpackardenterprise/hpe-oneview-sdk-for-terraform:v8.3.0-13-OV8.3
# Run docker commands below given, which  will in turn create a sh session 
# where you can create files, issue commands and execute the examples.
$ docker run -it docker pull hewlettpackardenterprise/hpe-oneview-sdk-for-terraform:v8.3.0-13-OV8.3 /bin/sh
```

### Local Setup

Local installation requires
- Installing Go
```bash 
$ apt-get install build-essential git wget
$ wget https://golang.org/dl/go1.15.7.linux-amd64.tar.gz

#unzip and untar the file 
$ tar -zxvf go1.15.7.linux-amd64.tar.gz

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

	Install Terraform 1.1.x from [here](https://www.terraform.io/downloads.html) and save it into `/usr/local/bin/terraform` folder (create it if it doesn't exists).



- Install Oneview Terraform Provider SDK from Terraform Registry

     Terraform 0.13 added support for automatically downloading providers from
the terraform registry. Add the following to your terraform project

```hcl
terraform {
  required_version = ">= 0.13"
  required_providers {
    oneview = {
      source = "HewlettPackard/oneview"
    }
  }
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

### HPE Synergy Image Streamer

From Release 8.1, Image streamer is no longer supported.

### Testing the Terraform Resources: 

In the home directory of project(terraform-provider-oneview) user needs to create (or) copy  the file that is to be executed. 

Sample example manifests are available in [example](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/examples) directory. 

The following terraform commands has to be executed to test the example. 
```terraform
$ terraform init 
$ terraform plan 
$ terraform apply
```
Note: Only a single terraform file (example file) should exist in the home folder to execute the above mentioned three commands. Once the resource is tested move that file to examples folder. 

Note: Currently this SDK supports OneView API 5200 minimally, where we can test OneView API 5200 version with this SDK. If API version used is not supported then error will be thrown. If API version is not provided then appliance's maximum supported API version will be used. 

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

Note: Currently this SDK supports OneView API 5200 minimally where we can test OneView API 5200 version with this SDK. No new fields have been added/deleted to support API5200 version. Complete support will be done in next releases.If  API version is not provided then appliance's API version will be used. If API version used is not supported then error will be thrown.

### HPE OneView Community

[HPE OneView Community Forums](http://hpe.com/info/oneviewcommunity)

Learn more about HPE OneView at [hpe.com/info/oneview](https://hpe.com/info/oneview)
