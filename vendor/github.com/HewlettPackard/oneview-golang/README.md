# HPE OneView golang bindings
| 5.50 Branch   |
| ------------- |
[![Build Status](https://travis-ci.org/HewlettPackard/oneview-golang.svg?branch=master)](https://travis-ci.org/HewlettPackard/oneview-golang)

## Introduction

HPE OneView makes it simple to deploy and manage today’s complex hybrid cloud infrastructure. HPE OneView can help you transform your data center to software-defined, and it supports HPE’s broad portfolio of servers, storage, and networking solutions, ensuring the simple and automated management of your hybrid infrastructure. Software-defined intelligence enables a template-driven approach for deploying, provisioning, updating, and integrating compute, storage, and networking infrastructure.

The HPE OneView Go SDK provides library to easily interact with HPE OneView and HPE Image Streamer REST APIs. The HPE OneView Go SDK enables developers to easily build integrations and scalable solutions with HPE OneView and HPE Image Streamer.

You can find the latest supported HPE OneView Go SDK [here](https://github.com/HewlettPackard/oneview-golang/releases/latest)

## What's New

HPE OneView Go library extends support of the SDK to OneView REST API version 2200 (OneView v5.50)

Please refer to [notes](https://github.com/HewlettPackard/oneview-golang/blob/master/CHANGELOG.md) for more information on the changes , features supported and issues fixed in this version

## Installation
HPE OneView SDK for Go can be installed from Source or Docker container installation methods. You can either use a docker container which will have the HPE OneView SDK for Go installed or perform local installation.

### Docker Container Setup

We also provide a lightweight and easy way to test and run oneview-golang. The `hewlettpackardenterprise/hpe-oneview-sdk-for-golang:<tag>` docker image contains an installation of oneview-golang installation you can use by just pulling down the Docker Image:

The Docker Store image `tag` consist of two sections: `<sdk_version-OV_version>`

Download and store a local copy of hpe-oneview-sdk-for-golang and use it as a Docker image.
```bash
$ docker pull hewlettpackardenterprise/hpe-oneview-sdk-for-golang:v1.7.0-OV5.5
```
Run docker commands and this will in turn create a sh session where you can create files, issue commands and execute the tests
```bash
$ docker run -it hewlettpackardenterprise/hpe-oneview-sdk-for-golang:v1.7.0-OV5.5 /bin/sh
```

### Local Setup
- Local installation requires 
Installing Go
```bash 
$ apt-get install build-essential git wget
$ wget https://dl.google.com/go/go1.11.3.linux-amd64.tar.gz
```
unzip and untar  with "tar -zxvf go1.11.linux-amd64.tar.gz", move it to /usr/local/ and create directory for Go.
```bash 
$ mv go1.11.3.linux-amd64.tar.gz /usr/local/ 
$ mkdir ~/go
```
Setting Environment Variable 
```bash 
$ export GOROOT=/usr/local/go
$ export GOPATH=$HOME/go 
$ export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
Install Oneview Go SDK
```go
go get -u github.com/HewlettPackard/oneview-golang 
```


## Configuration

### OneView Client Configuration
The OneView Client configuration options that can be passed during OneView Client object creation:

Following environment variables can be set for testing:

```bash
# Required
$ export ONEVIEW_OV_ENDPOINT=<ov_endpoint>
$ export ONEVIEW_OV_USER=<ov_username>
$ export ONEVIEW_OV_PASSWORD=<ov_password>
$ export ONEVIEW_OV_DOMAIN=LOCAL
$ export ONEVIEW_SSLVERIFY=false
$ export ONEVIEW_APIVERSION=<ov_apiversion>
```
Note: Currently this SDK supports OneView API 2200 minimally, where we can test OneView API 2200 version with this SDK. If API version is not provided then appliance's API version will be used. If API version used is not supported then error will be thrown.

```go
# Create a OneView client object:

import "github.com/HewlettPackard/oneview-golang/ov"

var ClientOV    *ov.OVClient

apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))

ovc := ClientOV.NewOVClient(
  os.Getenv("ONEVIEW_OV_USER"),       # This is to set the Oneview UserName
  os.Getenv("ONEVIEW_OV_PASSWORD"),   # This is to set the Oneview Password
  os.Getenv("ONEVIEW_OV_DOMAIN"),     # This is to set the Domain, default is LOCAL
  os.Getenv("ONEVIEW_OV_ENDPOINT"),   # This is to set the IP Address of the Oneview Appliance
  false,                              # This is to set the SSL, default is false
  apiversion,                         # This is to set the OneView REST API Version. Defaults to oneview appliance maximum supported REST API Version
  "*")
```
:lock: Tip: Check the file permissions because the password is stored in clear-text as Environment Variable.


### Image Streamer Client Configuration
The Image Streamer (I3S) client is very much similar to the OneView client. 
Here we create the Image Streamer(I3S) client through the Oneview client instance.
Following extra environment variables should be set for testing:

```bash
# Required
$ export ONEVIEW_I3S_ENDPOINT=<i3s_endpoint>
```

```go
# Create a OneView client object
import (
  "github.com/HewlettPackard/oneview-golang/ov"
  "github.com/HewlettPackard/oneview-golang/i3s" 
)

var (
  clientOV             *ov.OVClient
  i3sClient            *i3s.I3SClient
  endpoint             = os.Getenv("ONEVIEW_OV_ENDPOINT")
  i3sc_endpoint        = os.Getenv("ONEVIEW_I3S_ENDPOINT")
  username             = os.Getenv("ONEVIEW_OV_USER")
  password             = os.Getenv("ONEVIEW_OV_PASSWORD")
  domain               = os.Getenv("ONEVIEW_OV_DOMAIN")
)

apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
i3s_apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))

# Creates OV Client
ovc := clientOV.NewOVClient( username, password, domain, endpoint, false, api_version, "*")

# Gets Session ID
ovc.RefreshLogin()

# Create I3s Client using Session ID
i3sc := i3sClient.NewI3SClient(i3sc_endpoint, false, i3s_apiversion, ovc.APIKey)

```
:lock: Tip: Check the file permissions because the password is stored in clear-text as Environment Variable.

### Configuration Files

Configuration files can also be used to define client configuration (json or yaml formats). Here's an example json file:

```bash
Update the ``` oneview_config.json ``` file in example directory as below.
{
  "username":     "<ov_username>",  
  "password":     "<ov_password>",
  "endpoint":     "<ov_ip>",
  "domain":       "LOCAL",
  "apiversion":   <ov_apiversion>,
  "sslverify":    false,
  "ifmatch":      "*"
}
```

```go
# Create a OneView client object:
import 	"github.com/HewlettPackard/oneview-golang/ov"

var	ClientOV    *ov.OVClient
config, _ := ov.LoadConfigFile("oneview_config.json")

ovc := ClientOV.NewOVClient(
  config.UserName,        # This is to set the Oneview UserName
  config.Password,        # This is to set the Oneview Password
  config.Domain,          # This is to set the Domain, default is LOCAL
  config.Endpoint,        # This is to set the IP Address of the 
  config.SslVerify,       # This is to set the SSL, default is false
  config.ApiVersion,      # This is to set the OneView REST API Version. Defaults to oneview appliance maximum supported REST API Version
  config.IfMatch)
```
:lock: Tip: Check the file permissions if the password or token is stored in clear-text.

## API Implementation

A detailed list of the HPE OneView REST interfaces that have been implemented in this SDK can be found in the [endpoints-support.md](https://github.com/HewlettPackard/oneview-golang/blob/master/endpoints-support.md) file.


## Testing your changes

We use docker to build and test, run this project on a system that has docker. If you don't use docker, you will need to install and setup go-lang locally as well as any other make requirements.  
You can use `USE_CONTAINER=false` environment setting for make to avoid using docker. Otherwise make sure to have these tools:
- docker client and daemon
- gnu make tools

The Tests in GoLang are divided into two segments one is doing the acceptance test and another drives the Unit Tests.

### With Docker

```
make test
```

### Without docker

* Install golang 1.5 or higher(We recommend using Go 1.11)
* Install go packages listed in .travis.yml

The Test Data for these Tests are  supplied through JSON file stored at `test/data config_EGSL_tb200.json`

These Tests can be run locally, you must install the below dependencies as mention in .travis.yml and do export USE_CONTAINER=false

```go
go get github.com/mattn/goveralls
go get -u golang.org/x/lint/golint
```

These Tests are using make files to save the compile time. 
Below are the commands to run the tests.

```bash
$ sudo make vet
$ sudo make test-short
$ sudo make test-long
$ sudo make coverage-send
```

Note: ```make test``` runs all the above mentioned tests.

## Getting Help 

Are you running into a road block? Have an issue with unexpected behavior? Feel free to open a new issue on the [issue tracker](https://github.com/HewlettPackard/oneview-golang/issues/new)

## License
This project is licensed under the Apache 2.0 license. Please see [LICENSE](LICENSE) for more info.

## Contributing and feature requests

We welcome your contributions to the HPE OneView for Go SDK library. 

**Contributing:** You know the drill. Fork it, branch it, change it, commit it, and pull-request it.
We are passionate about improving this project, and glad to accept help to make it better.
For more information refer [CONTRIBUTING.md](https://github.com/HewlettPackard/oneview-golang/blob/master/CONTRIBUTING.md) file.

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
