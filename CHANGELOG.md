All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html)

# [v1.1.0] (UnReleased)

### Notes
This version supports OneView appliances with version 4.2 using the OneView REST API version 1000.

- Support for Update Group in Logical Enclosure.
- Support for Update Compliance in Logical Interconnect.
- Usecase for synergy infrastructure provisioning with Network.
- Usecase for synergy infrastructure provisioning with i3s.

### Oneview Features supported
- Enclosure Group
- Ethernet network
- FC network
- FCoE network
- Interconnect
- Interconnect Type
- Logical enclosure
- Network set
- Server hardware
- Server hardware type
- Storage pool
- Storage system
- Storage volume
- Storage volume attachment
- Storage volume template

### Image Streamer Features supported
- Deployment Plan

# [v1.0.1]
### Notes
- Patch changes to sever profile to include boot order fixes.
- Usecase for infrasructure provisioning without storage and image streamer

# [v1.0.0]
### Notes
  This is the first release of the OneView modules in Terraform and it adds support to core features listed below.
  This version of the module supports OneView appliances with version 4.10, using the OneView REST API version 800.

### Major changes:
- Support for Go 1.11  
- Added example files for the resources for improved readability and usability.
- Added CHANGELOG to track versions, issues and improvements.
- Officially adopted Semantic Versioning for the SDK
- Added endpoints-support.md to track the supported and tested endpoints for the different HPE OneView REST APIs
- Added SNMP v3 configuration support to Logical Interconnect Group
- Added import support to all resources
- Updated contribution guidelines
- Support for datasources of resources.

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
- Server hardware
- Server hardware type
- Server profile
- Server profile template
- Storage system
- Uplink set
