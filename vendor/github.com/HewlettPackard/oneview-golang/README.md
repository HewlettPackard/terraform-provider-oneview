# HPE OneView golang bindings

[![Build Status](https://travis-ci.org/HewlettPackard/oneview-golang.svg?branch=master)](https://travis-ci.org/HewlettPackard/oneview-golang)

HPE OneView allows you to treat your physical infrastructure as code, and now
you can integrate your favorite tools based in golang with HPE OneView.

## Build requirements
We use docker to build and test, run this project on a system that has docker.
If you don't use docker, you will need to install and setup go-lang locally as
well as any other make requirements.  You can use `USE_CONTAINER=false` environment
setting for make to avoid using docker. Otherwise make sure to have these tools:
- docker client and daemon
- gnu make tools

### Docker Container Setup

We also provide a lightweight and easy way to test and run `oneview-golang`. The `hewlettpackardenterprise/hpe-oneview-sdk-for-golang:<tag>` docker image contains an installation of oneview-golang installation you can use by just pulling down the Docker Image:

The Docker Store image `tag` consist of two sections: `<sdk_version-OV_version>`

```bash
# Download and store a local copy of hpe-oneview-sdk-for-golang and
# use it as a Docker image.
$ docker pull hewlettpackardenterprise/hpe-oneview-sdk-for-golang:v1.5.0-OV5.3
#Run docker commands and this will in turn create
# a sh session where you can create files, issue commands and execute the tests
$ docker run -it hewlettpackardenterprise/hpe-oneview-sdk-for-golang:v1.5.0-OV5.3 /bin/sh
```


### Environment Variables

Following environment variables can be set for testing:

```bash
# Required
export ONEVIEW_OV_ENDPOINT=<ov_endpoint>

export ONEVIEW_OV_USER=<ov_username>
export ONEVIEW_OV_PASSWORD=<ov_password>
export ONEVIEW_OV_DOMAIN=LOCAL
export ONEVIEW_SSLVERIFY=false
export ONEVIEW_APIVERSION=<ov_apiversion>
```
Note: Currently this SDK supports OneView API 1800 minimally where we can test OneView API 1800 version with this SDK. No new fields have been added/deleted to support API 1800 version. Complete support will be done in next releases.

## Testing your changes

### From a container
```
make test
```

### Without docker
* Install golang 1.5 or higher(We recommend using Go 1.11)
* Install go packages listed in .travis.yml
```
USE_CONTAINER=false make test
```

## Contributing

We welcome your contributions to the HPE OneView golang SDK. See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

## Hacking Guide

Use the [hacking guide](HACKING.md) to run local acceptance testing and debugging local contributions.
Currently test cases are not supported portion of our CI/CD approval process but might be made available from this test suite in the future.   Contributions to the test suite is appreciated.

## License
This project is licensed under the Apache License, Version 2.0.  See LICENSE for full license text.
