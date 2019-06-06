All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html)

# [v1.0.0] (Unreleased)
### Notes
  Major release which extends support to OneView appliances with versions 4.10, using the OneView Rest API version 800.

### Major changes:
- Extended support of SDK to OneView API800.
- Support for Go 1.11  
- Added example files for the resources for improved readability and usability.
- Added CHANGELOG to track versions, issues and improvements.
- Officially adopted Semantic Versioning for the SDK
- Added endpoints-support.md to track the supported and tested endpoints for the different HPE OneView REST APIs
- Added SNMP v3 configuration support to Logical Interconnect Group
- Added import support to all resources
- Updated contribution guidelines

#### Bug fixes & Enhancements:
- [#47] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/47) Added Synergy Support for Logical Interconnect Groups

### Oneview Features supported
- Enclosure
- Enclosure group
- FC network
- Interconnect
- Interconnect type
- Logical enclosure
- Logical interconnect
- Logical interconnect group
- Scope
- Server hardware
- Server hardware type
- Server profile
- Uplink set