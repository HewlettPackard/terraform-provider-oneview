All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html)

# [v1.7.1]
#### Bug fixes & Enhancements:
- [#285] (https://github.com/HewlettPackard/oneview-golang/issues/285) Fix Docker Setup section in README.md
- [#286] (https://github.com/HewlettPackard/oneview-golang/issues/286) GetIloIPAddress() does not work with GetServerHardwareList()

# [v1.7.0] (2020-11-05)
### Notes
- This release extends supports of the SDK to Oneview API Version 2200 and Image Streamer API Version 2000.

### Oneview Features supported
- Deployment Plan
- Ethernet Network
- Enclosure
- Enclosure Group
- FC Network
- FCOE Network
- Hypervisor Cluster Profiles
- Hypervisor Manager
- Interconnects
- Interconnect Types
- Logical Interconnect
- Logical Enclosure
- Logical Interconnect Group
- Network Set
- Scope
- Server Certificates
- Server Hardware
- Server Hardware Types
- Server Profile
- Server Profile Template
- Storage Pool
- Storage System
- Storage Template
- Storage Volume
- Storage Volume Attachments
- Tasks
- Uplink Set

# [v1.6.0] 
### Notes
- This release extends supports of the SDK to Oneview API2000.
- Minor Bug Fixes

### Oneview Features supported
- Enclosure
- Enclosure Group
- Ethernet Network
- FC Network
- FCOE Network
- Hypervisor Manager
- Hypervisor Cluster Profile
- Interconnects
- Interconnect Types
- Logical Enclosure
- Logical Interconnect
- Logical Interconnect Group
- Network Set
- Scope
- Server Certificates
- Server Hardware
- Server Hardware Types
- Server Profile
- Server Profile Template
- Storage Pool
- Storage Systems
- Storage Volume
- Storage Volume Attachments
- Storage Volume Template
- Tasks
- Uplink Set 

#### Bug fixes & Enhancements:
- [#239] (https://github.com/HewlettPackard/oneview-golang/issues/239) CreateProfileFromTemplate not working w OV 5.20
- [#257] (https://github.com/HewlettPackard/oneview-golang/issues/257) Add server name to ServerHardware structure in server_hardware.go

### Major changes:
- Refactored SDK to take default API version from Oneview appliance.
- Added support for publishing Docker Image whenever a new release is created in GitHub.

# [v1.5.0] (2020-07-17)
### Notes
- This release extends supports of the SDK to Oneview API1800.
- Added support for Id pools IPV4 Ranges.
- Added support for Email Notifications resource from API 1600.
- Minor Bug Fixes

### Oneview Features supported
- Enclosure
- Enclosure Group
- Ethernet Network
- FCOE Network
- FC Network
- Hypervisor Cluster Profile
- Hypervisor Manager
- Id pools IPV4 Ranges
- Interconnects
- Interconnect Types
- Logical Enclosure
- Logical Interconnect
- Logical Interconnect Types
- Network Set
- Scope
- Server Certificate
- Server Hardware
- Server Heardware Types
- Server Profile
- Server Profile Template
- Storage Pool
- Storage System
- Storage Volume
- Storage Volume Attachment
- Storage Volume Template
- Task
- Uplink Sets

#### Bug fixes & Enhancements:
- [#56] (https://github.com/HewlettPackard/oneview-golang/issues/56) Selecting available machines should be atomic
- [#141] (https://github.com/HewlettPackard/oneview-golang/issues/141) Query Parameter retains in the next endpoint call
- [#154] (https://github.com/HewlettPackard/oneview-golang/issues/154) Add json configuration file for creating NewOVClient to run examples

 
# [v1.4.0] (2020-05-27)
### Notes
- Added support for loginMsgAck attribute in login session POST request.
- This release supports API1600 minimally where we can test OneView API1600 with this SDK
### Oneview Features supported
- Deployment Plan
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
- Task
- Uplink Set


# [v1.3.0] (2020-04-17)
### Notes
- This release add supports for Hypervisor Manager, Hypervisor Cluster and Server Certificate for API V800,V1000 and V1200
### Oneview Features supported
- Hypervisor Cluster Profile
- Hypervisor Manager 
- Server Certificate

# [v1.2.0] (2019-12-13)
### Notes
- This release supports API1200 minimally where we can test OneView API1200 with this SDK. No new fields are added/deleted to support API1200
. Complete support will be done in next releases.
- Support for API V1020 for image streamer resource deployment plan.

### Oneview Features supported
- Deployment Plan
- Enclosure
- Enclosure Group
- FC Network
- FCoE Network
- Interconnect
- Interconnect type
- Logical Enclosure
- Logical interconnect
- Logical Interconnect Group
- Network Set
- Server hardware
- Server hardware type
- Server Profile
- Server Profile Template
- Storage Pool
- Storage System
- Storage Volume
- Storage Volume Attachment
- Storage Volume Template
- Task
- Uplink set

# [v1.1.0] (2019-12-03)
### Notes
 - Minor Bug Fixes
 - New Endpoint support for Server Hardware (firmware)
 - Server Profile support with Server Profile Template extended to 1000
 - Support for update compliance in Logical Interonnect.

### Major changes:
- Extended support of SDK to OneView API1000.

#### Bug fixes & Enhancements:
- [#153] (https://github.com/HewlettPackard/oneview-golang/issues/153) change in the enclosure.go file for editing  wrong fields

### Oneview Features supported
- Deployment Plan
- Enclosure
- Enclosure group
- Ethernet Network
- FC Network
- FCoE Network
- Interconnect
- Interconnect type
- Logical enclosure
- Logical interconnect
- Logical interconnect group
- Network Set
- Scope
- Server hardware type
- Server Profile
- Server Profile Template
- Storage pool
- Storage system
- Storage volume
- Storage volume attachment
- Storage volume template
- Task
- Uplink set

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
