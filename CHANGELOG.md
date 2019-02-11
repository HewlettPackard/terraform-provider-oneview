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

#### Bug fixes & Enhancements:
- [#58] (https://github.com/HewlettPackard/terraform-provider-oneview/pull/58) Added SNMP v3 configuration support to Logical Interconnect Group
- [#56] (https://github.com/HewlettPackard/terraform-provider-oneview/pull/56) Added Synergy Support for Logical Interconnect Groups
- [#54] (https://github.com/HewlettPackard/terraform-provider-oneview/pull/54) Added import support to all resources
- [#51] (https://github.com/HewlettPackard/terraform-provider-oneview/pull/51) Updated contribution guidelines

### Oneview Features supported
- Enclosure    
- FC network       
