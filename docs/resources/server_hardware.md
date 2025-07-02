---
layout: "oneview"
page_title: "Oneview: server_hardware"
sidebar_current: "docs-server_hardware"
description: |-
 Creates/Updates server_hardware.
---

# oneview\_server_hardware

Use this source to create/update Server Hardware.

## Example Usage

```hcl
resource "oneview_server_hardware" "sh" {
  configuration_state = "Managed"
  hostname   = "<serverIp>"
  username = "dcs"
  password = "dcs"
  licensing_intent = "OneView"
}

```

## Argument Reference

* `hostname` - (Required)IP address or host name of a rack mount server's iLO management processor.
* `username` -  (Required) iLO user account login name with Administrator User Accounts privilege.
* `password` - (Required) The password for the specified username.
* `licensing_intent` - (Required) The type of product license to assign to the server hardware.

## Attributes Reference

* `configuration_state` - Specifies the desired server state.
  
* `force` - Specify 'true' to force the addition of the server and take ownership away from any other manager. Use this optional flag with caution. The default is 'false'.

* `maintenance_mode` - Maintenance mode of the server hardware, true or false.

* `uid_state` - Uid state of the the server hardware.

* `one_time_boot` - Boot options of the server hardware.

* `power_state` - Current power state of the server hardware.

* `location_uri` - For blade servers, the enclosure in which this blade server resides. This URI can be used to retrieve information about the enclosure. This value is not set for rack mount servers.

* `mp_firmware_version` - The version of the firmware installed on the iLO.

* `uri` - The URI of the resource.

* `type` - Type of the resource.

* `server_group_uri` - For blade servers, this is the URI of the containing enclosure's enclosure group. This URI can be used to retrieve information about the enclosure group or to identify all the servers in the same group. This value is not set for rack mount servers.

* `server_hardware_type_uri` - URI of the server hardware type associated with the server hardware. This URI can be used to retrieve information about the server hardware type. 

* `server_profile_uri` - URI of a server profile assigned to this server hardware

* `uuid` - Universally Unique ID (UUID) of the server hardware.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.
