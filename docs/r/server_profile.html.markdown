---
layout: "oneview"
page_title: "Oneview: server_profile"
sidebar_current: "docs-oneview-server-profile"
description: |-
  Creates a server profile.
---

# oneview\_server\_profile

Creates a server profile.
et
## Example Usage

Create server profile using server profile template.
```js
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  template = "${oneview_server_profile_template.test.name}"
}
```

Create server profile without server profile template.
```js
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  hardware_name = "00AT092 bay 2"
  type = "ServerProfileV10"
}
```
Update server profile
```js
resource "oneview_server_profile" "default" {
        update_type = "patch"
        options = [
        {
          op = "replace"
          path = "/refreshState"
          value = "RefreshPending"
        }
        ]
        name = "TestSP_Renamed"
        type = "ServerProfileV10"
        server_hardware_type = "SY 480 Gen9 3"
        enclosure_group = "SYN03_EC"
        hardware_name = "SYN03_Frame3, bay 1"
}
```

Patch request for server profile
```js
resource "oneview_server_profile" "default" {
        update_type = "patch"
        options = [
        {
          op = "replace"
          path = "/refreshState"
          value = "RefreshPending"
        }
        ]
        name = "TestSP"
        type = "ServerProfileV10"
        server_hardware_type = "SY 480 Gen9 3"
        enclosure_group = "SYN03_EC"
        hardware_name = "SYN03_Frame3, bay 1"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

* `template` - (Optional) The name of the template you will use for the server profile. 

* `update_type` - (Required) Type of update of Server Profile.

	| NO |  Type of Update   |   Update String |
	|----|-------------------|-----------------|
	|  1 |`Update`           |'put'            |
	|----|-------------------|-----------------|
	|  1 |`Patch`            |'patch'          |
	|----|-------------------|-----------------|

- - -

- - -

* `public_connection` - (Optional) The name of the network that is going out to the public.

* `hardware_name` - (Optional) The name of the Server Hardware the server will be provisioned on.
  If this isn't used, a server hardware will be picked based on compatibility with the server profile template and any hw_filter(s) (see below).

* `hw_filter` - (Optional) List of filters to apply to the search for HW. See the OneView API docs pertaining to common filter query params, but the basic format is `[not] {attribute} {operator} '{value}'`. For example, `hw_filter = ["memoryMb >= 4096", "processorCoreCount = 4", "processorSpeedMhz >= 2400", "processorType regex '^Intel.*'"]`

* `power_state` - (Optional) Power state to enforce; `"on"` or `"off"`

* `type` - (Optional) The server profile version to be provisioned. Defaults to ServerProfileV5.
  Use ServerProfileV6 to use Image Streamer.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `serial_number` - A 10-byte value that is exposed to the Operating System as the server hardware's
  Serial Number. The value can be a virtual serial number, user defined serial number or physical serial
  number read from the server's ROM. It cannot be modified after the profile is created.

* `public_mac` - The MAC address of the NIC your public network is attached to.
  Need to specify public_connection to access this value. 
  
* `public_slot_id` - The slot id of the NIC your public network is attached to.
  Need to specify public_connection to access this value. 
  
* `ilo_ip` - The ILO ip address that is managing the server.

* `hardware_uri` - The URI of the hardware the server is provisioned on.

* `os_deployment_settings` - (Optional) OS Deployment settings applicable when deployment is invoked through a server profile.

* `boot_order` - Defines the order in which boot will be attempted on the available devices.

* `boot_mode` -  Boot mode settings to be configured on the server.

* `bios_option` - Server BIOS settings.

* `server_hardware_type` - Identifies the server hardware type for which the Server Profile was designed. 

* `enclosure_group` -  Identifies the enclosure group for which the Server Profile was designed.

* `affinity` - This identifies the behavior of the server profile when the server hardware is removed or replaced. 

* `hide_unused_flex_nics` - This setting controls the enumeration of physical functions that do not correspond to connections in a profile.

* `serial_number_type` -  Specifies the type of Serial Number and UUID to be programmed into the server ROM.

* `wwn_type` -  Specifies the type of WWN address to be programmed into the IO devices. 

* `mac_type` - Specifies the type of MAC address to be programmed into the IO Devices.

* `firmware` - Defines and enables firmware baseline management.

* `local_storage` -  Local storage settings to be configured on the server.

* `logical_drives` - The list of logical drives associated with the controller. 

* `san_storage` - The profile san storage configuration.

* `volume_attachments` - The list of storage volume attachments.