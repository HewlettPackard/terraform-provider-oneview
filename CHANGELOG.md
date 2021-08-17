All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html)

# [v6.3.0-13] (Unreleased)

### Notes
- This release supports API3200 minimally where we can use OneView v6.30 with this SDK.

#### Major Changes
1. Extended support to Server Profile resource on DL Server.
2. Extended support to Server Profile Template resource on DL Server.
3. Extended support to Server Hardware resource for DL Server.
   - POST   /rest/server-hardware
   - PATCH  /rest/server-hardware/{id}
4. logical_interconnect_group: Changed the way we provide port_num, instead of list we have to provide integer value for port_num in each logical_port_config in uplink_set.

### Bug fixes & Enhancements:
- [#312] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/312) Unable to apply server_profile_template (Invalid character '.)

# [v6.2.0-13]

### Notes
- This release supports API3000 minimally where we can use OneView v6.20 with this SDK.

#### Major Changes
1. Added support to Appliance SNMPv3 Trap Destinations
   - GET /rest/appliance/snmpv3-trap-forwarding/destinations/{id}
   - POST /rest/appliance/snmpv3-trap-forwarding/destinations
   - PUT /rest/appliance/snmpv3-trap-forwarding/destinations/{id}
   - DELETE /rest/appliance/snmpv3-trap-forwarding/destinations/{id}
2. Added support to Appliance Time and Locale Configuration.
   - GET  /rest/appliance/configuration/time-locale
   - POST /rest/appliance/configuration/time-locale
3. Added support to Label Resource. 
   - GET /rest/labels/resources/**
   - PUT /rest/labels/resources/**
   - DELETE /rest/labels/resources/**
   - POST /rest/labels/resources
4. Added support to Appliance SNMPv1 Trap Destinations
   - GET     /rest/appliance/trap-destinations/{id}
   - POST    /rest/appliance/trap-destinations/{id}
   - PUT     /rest/appliance/trap-destinations/{id}
   - DELETE  /rest/appliance/trap-destinations/{id}
5. Added support to support to retreive the appliance supported version.
   - GET /rest/version
6. Added support to Appliance SSH Access Resource.
   - GET  /rest/appliance/ssh-access
   - PUT /rest/appliance/ssh-access
7. Added support to Firmware Drivers Resource.
   - GET    /rest/firmware-drivers/{id}
   - POST   /rest/firmware-drivers
   - DELETE /rest/firmware-drivers/{id}
8. Added support to Connection Templates Resource.
   - GET /rest/connection-templates/{id}
   - PUT /rest/connection-templates/{id}
9. Added support to Id Pools Ipv4 Range
   - GET     /rest/id-pools/ipv4/ranges/{id}
   - POST    /rest/id-pools/ipv4/ranges
   - PUT     /rest/id-pools/ipv4/ranges/{id}
   - DELETE  /rest/id-pools/ipv4/ranges/{id}
10. Added support to Id Pools IPv4 Subnets
   - GET      /rest/id-pools/ipv4/subnets/{id}
   - POST     /rest/id-pools/ipv4/subnets
   - PUT      /rest/id-pools/ipv4/subnets/{id}
   - PUT      /rest/id-pools/ipv4/subnets/{id}/allocator
   - PUT      /rest/id-pools/ipv4/subnets/{id}/collector
   - DELETE   /rest/id-pools/ipv4/subnets/{id}
11. Added support to SNMPv3 User
   - GET      /rest/appliance/snmpv3-trap-forwarding/users/{id}
   - POST     /rest/appliance/snmpv3-trap-forwarding/users
   - PUT      /rest/appliance/snmpv3-trap-forwarding/users/{id}   
   - DELETE   /rest/appliance/snmpv3-trap-forwarding/users/{id}


### Notes
- This release supports API3000 minimally where we can use OneView v6.20 with this SDK.

### Oneview Features supported
- Appliance SNMPv1 Trap Destinations
- Appliance SNMPv3 Trap Destinations
- Appliance SSH Access
- Appliance Time and Locale Configuration
- Connection Template
- Ethernet Network
- Enclosure
- Enclosure Group
- FC Network
- FCOE Network
- Firmware Drivers
- Hypervisor Manager
- Hypervisor Cluster Profile
- Id Pools Ipv4 Range
- Id Pools IPv4 Subnets
- Labels
- Interconnects
- Interconnect Types
- Logical Enclosure
- Logical Interconnects
- Logical Interconnect Groups 
- Network Set
- Scope
- Server Certificate
- Server Hardware
- Server Hardware Type
- Server Profile
- Server Profile Template
- Storage Pool
- Storage Volume
- Storage Volume Attachments
- Storage Volume Template
- Storage System
- Uplink Set

### Bug fixes & Enhancements:
- [#289] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/289) Terraform crashes for Logical Interconnect Group Read
- [#300] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/300) Can't create resource oneview_server_profile_template
- [#301] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/301) Can't create resource oneview_server_profile
- [#336] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/336) Terraform crash creating Server Profile Template 
- [#334] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/338) volume creation is not derived from volume template
- [#343] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/343) Creating Server Profile produces an Error 404

# [v6.1.0-13] 

### Notes
- This release supports API2800 minimally where we can use OneView v6.10 with this SDK. 
- Migration Support document is added to help the user migrate from older  sdk supporting Terraform 0.11.

### Bug fixes & Enhancements:
- [#272] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/272) more server profile/template issues not working as expected.
- [#280] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/280) local storage configuration for profiles - importConfiguration
- [#288] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/288) Terraform crashes on Logical Enclosure create
- [#295] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/295) server profile update in-place.

### Oneview Features supported
- Ethernet Network
- Enclosure
- Enclosure Group
- FC Network
- FCOE Network
- Hypervisor Manager
- Hypervisor Cluster Profile
- Interconnects
- Interconnect Types
- Logical Enclosure
- Logical Interconnects
- Logical Interconnect Groups 
- Network Set
- Scope
- Server Certificate
- Server Hardware
- Server Hardware Type
- Server Profile
- Server Profile Template
- Storage Pool
- Storage Volume
- Storage Volume Attachments
- Storage Volume Template
- Storage System
- Uplink Set

### Image Streamer resources supported:
- Deployment Plans

# [v6.0.0-13]
#### Major changes
1.This release is validated to support Terraform 0.13.  Terraform 0.13 support HCL2 and enables automated installation of the HPE OneView Terraform Provider.

### Bug fixes & Enhancements:
- [#277] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/277) SPT Import configuration issue
- [#285] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/285) Terraform crashes for Logical Interconnect Group Read. 

### Oneview Features supported
- Ethernet Network
- Enclosure
- Enclosure Group
- FC Network
- FCOE Network
- Hypervisor Manager
- Hypervisor Cluster Profile
- Interconnects
- Interconnect Types
- Logical Enclosure
- Logical Interconnects
- Logical Interconnect Groups 
- Network Set
- Scope
- Server Certificate
- Server Hardware
- Server Hardware Type
- Server Profile
- Server Profile Template
- Storage Pool
- Storage Volume
- Storage Volume Attachments
- Storage Volume Template
- Storage System
- Uplink Set

# [v6.0.0-12]
### Notes
- This release supports API2600 minimally where we can use OneView v6.00 with this SDK. 

### Oneview Features supported
- Ethernet Network
- Enclosure
- Enclosure Group
- FC Network
- FCOE Network
- Hypervisor Manager
- Hypervisor Cluster Profile
- Interconnects
- Interconnect Types
- Logical Enclosure
- Logical Interconnects
- Logical Interconnect Groups 
- Network Set
- Scope
- Server Certificate
- Server Hardware
- Server Hardware Type
- Server Profile
- Server Profile Template
- Storage Pool
- Storage Volume
- Storage Volume Attachments
- Storage Volume Template
- Storage System
- Uplink Set

### Image Streamer resources supported:
- Deployment Plans

### Bug fixes & Enhancements:
- [#238] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/238) server profile connection settings
- [#267] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/267) local_storage not working as expected.
- [#268] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/268) local storage issue ( defined in server profile template). 

# [v1.7.0-12]
### Notes
- This release supports API2400 minimally where we can use OneView v5.60 with this SDK. 

#### Major changes
1. Added support for new endpoints in logical interconnect resource, logical interconnect data source and logical interconnect group data source.
   - PUT /rest/logical-interconnects/{id}/portFlapSettings

### Oneview Features supported
- Ethernet Network
- Enclosure
- Enclosure Group
- FC Network
- FCOE Network
- Hypervisor Manager
- Hypervisor Cluster Profile
- Interconnects
- Interconnect Types
- Logical Enclosure
- Logical Interconnects
- Logical Interconnect Groups 
- Network Set
- Scope
- Server Certificate
- Server Hardware
- Server Hardware Type
- Server Profile
- Server Profile Template
- Storage Pool
- Storage Volume
- Storage Volume Attachments
- Storage Volume Template
- Storage System
- Uplink Set

### Bug fixes & Enhancements:
- [#233] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/233) terraform import oneview_server_profile_template missing settings.
- [#234] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/234) local storage configuration.
- [#242] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/242) serverprofile bootmode_secureboot.

# [v1.6.0]
### Notes
- This release supports API2200 minimally where we can use OneView API2200 with this SDK and ImageStreamer REST API version 2000 (I3S v5.40).  
- This release adds variable.tf file to provide easy access for adding Appliance Credentials through environment variables or during runtime. 
- No new fields are added/deleted to support API2200. 

### Oneview Features supported
- Enclosure
- Enclosure Group
- Ethernet Network
- FC Network
- FCOE Network
- Hypervisor Cluster Profile
- Hypervisor Manager
- Interconnects
- Interconnect Types
- Network Set
- Logical Enclosure
- Logical Interconnect
- Logical Interconnect Group
- Scope
- Server Certificate
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
 
### Image Streamer resources supported:
- Deployment Plans

### Bug fixes & Enhancements:
- [#118] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/118) mp_ip_address and mp_dns_name return NULL value.

# [v1.5.0]
### Notes
- This release supports API2000 minimally where we can use OneView API2000 with this SDK. No new fields are added/deleted to support API2000.

### Oneview Features supported
- Enclosure
- Enclosure Group
- Ethernet Network
- FC Network
- FCOE Network
- Hypervisor Cluster Profile
- Hypervisor Manager
- Interconnects
- Interrconnect Type
- Logical Encloure
- Logical Interconnect
- Logical Interconnect Group
- Network Set
- Scope
- Server Certificates
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

### Major changes:
- Refactored SDK to take default API version from Oneview appliance
- Added support for automatic publish of Docker Image when there is a new release in GitHub.

### Bug fixes & Enhancements:
- [#29] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/29) Add Description and Bios settings to server template
- [#63] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/63) Create util function to get scope by name while creating a resource instead of hardcoding it
- [#162] (https://github.com/HewlettPackard/terraform-provider-oneview/issues/162) Logical Interconnect Group dependentResourceUri error.

# [v1.4.0]
### Notes
- This release supports API1800 minimally where we can use OneView API1800 with this SDK. No new fields are added/deleted to support API1800.

### Oneview Features supported
- Enclosure
- Enclosure Group
- Ethernet Network
- FC Network
- FCOE Network
- Hypervisor Cluster Profile
- Hypervisor Manager
- Interconnect
- Interconnect Type
- Logical Enclosure
- Logical Interconnect
- Logical Interconnect Group
- Network Set
- Scopes
- Server Certificates
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

# [v1.3.0]
### Notes
- This release supports to API 800,1000,1200,1600 for Hypervisor Manager, Hypervisor Cluster Profile and Server Certificate resources.
- This release supports API1600 minimally where we can use OneView API1600 with this SDK. No new fields are added/deleted to support API1600.

### Oneview Features supported
- Deployment Plan
- Enclosure
- Enclosure Group
- Ethernet Network
- FC Network
- FCoE Network
- Hypervisor Cluster Profile
- Hypervisor Manager
- Interconnect
- Interconnect Type
- Logical Enclosure
- Logical Interconnect
- Logical Interconnect Group
- Network Set
- Scopes
- Server Certificate
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
