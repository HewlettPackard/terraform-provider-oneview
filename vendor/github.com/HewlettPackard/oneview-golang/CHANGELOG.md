All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html)

# [v1.1.0] (UNRELEASED)
### Notes
 - Minor Bug Fixes
 - New Endpoint support for Server Hardware (firmware)

#### Bug fixes & Enhancements:
- [#153] (https://github.com/HewlettPackard/oneview-golang/issues/153) change in the enclosure.go file for editing  wrong fields

# [v1.0.0] (2019-02-07)
### Notes
  Major release which extends support to OneView appliances with versions 4.10, using the OneView Rest API version 800.

### Major changes:
- Extended support of SDK to OneView API800.
- Support for Go 1.11    
- Added example files for the resources for improved readability and usability.
- Added CHANGELOG to track versions, issues and improvements.
- Officially adopted Semantic Versioning for the SDK
- Added endpoints-support.md to track the supported and tested endpoints for the different HPE OneView REST APIs

#### Bug fixes & Enhancements:
- [#139] (https://github.com/HewlettPackard/oneview-golang/issues/139) fix examples due to change in the ov.go file for creating a NewOVClient
- [#119] (https://github.com/HewlettPackard/oneview-golang/issues/119) Server Profile Template support for API 500+
- [#111] (https://github.com/HewlettPackard/oneview-golang/issues/111) Travis build failed related to golint dependency
- [#108] (https://github.com/HewlettPackard/oneview-golang/issues/108) Add change log file to keep a track of each version
- [#106] (https://github.com/HewlettPackard/oneview-golang/issues/106) Building DockerImage fails with the current Go version

### Oneview Features supported
- Enclosure    
- Enclosure group
- Ethernet network
- FC network       
- Interconnect
- Interconnect type
- Logical enclosure
- Logical interconnect
- Logical interconnect group
- Scope
- Server profile
- Server profile template
- Storage pool
- Storage system
- Storage volume attachment
- Server hardware
- Server hardware type
- Task
- Uplink set
- Volume

# [v0.8.2] (2017-03-09)
### Notes
  This version of the SDK supports Image streamer with versions 2.00.00 or higher, using the OneView Rest API version 300.
### Major changes:
  Improvements and refactoring.

### Oneview Features supported
- Artifacts Bundle
- Deployment plan
- Enclosure Group
- Golden image
- OS Build Plan
- OS Deployment plan
- OS volume
- Plan script
- Server profile

# [v0.8.1] (2016-09-14)
### Notes
  This version supports some more resources for Oneview appliance for Rest API 120/200.

### Oneview Features supported
- FC network    
- Server hardware Type    
- Storage Volume

# [v0.5.5] (2016-07-15)
### Notes
  This is the first release of the OneView SDK in Go and it adds support to core features.
  This version of the SDK supports OneView appliances with versions 2.00.00 or higher, using the OneView Rest API version 120 or 200.
  This version is compatible with docker-machine-oneview.

#### Bug fixes & Enhancements:
  - [#61] (https://github.com/HewlettPackard/oneview-golang/issues/61) Issues when running make test target locally
  - [#40] (https://github.com/HewlettPackard/oneview-golang/issues/40) refresh login should check if the current session id is good

### Oneview Features supported
- Ethernet network    
- FCoE network
- Interconnect type   
- Logical interconnect group
- Logical switch
- Logical switch group
- Network Set
- Server hardware
- Server profile
- Server profile template
- Storage
- Switch Type
- Task
- Version
