All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html)

# [v1.3.0](unreleased)
### Notes
- This release supports for API 800,1000,1200  for resource listed below.

### Oneview Features supported
- Hypervisor Manager
- Server Certificate

# [v1.2.0]
### Notes
- This release supports API1200 minimally where we can use OneView API1200 with this SDK. No new fields are added/deleted to support API1200. Complete support will be done in next releases.
- Support for API V1020 for image streamer resource deployment plan.

### Oneview Features supported
- Deployment Plan
- Enclosure
- Enclosure Group
- FC Network
- FCoE Network
- Interconnect
- Interconnect Type
- Logical Enclosure
- Logical Interconnect
- Logical Interconnect Group
- Network Set
- Server Hardware
- Server Hardware Type
- Server Profile
- Server Profile Template
- Storage Pool
- Storage System
- Storage Volume
- Storage Volume Attachment
- Storage Volume Template
- Uplink Set

# [v1.1.0]

### Notes
This version supports OneView appliances with version 4.2 using the OneView REST API version 1000.

- Support for Update Group in Logical Enclosure.
- Support for Update Compliance in Logical Interconnect.
- Usecase for synergy infrastructure provisioning with Network.
- Usecase for synergy infrastructure provisioning with i3s.

### Oneview Features supported
- Enclosure
- Enclosure Group
- Ethernet network
- FC network
- FCoE network
- Interconnect
- Interconnect Type
- Logical enclosure
- Logical interconnect
- Logical interconnect group
- Network set
- Scope
- Server hardware
- Server hardware type
- Server profile
- Server profile template
- Storage pool
- Storage system
- Storage volume
- Storage volume attachment
- Storage volume template
- Uplink set

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
