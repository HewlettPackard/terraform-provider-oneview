---
layout: "oneview"
page_title: "Oneview: server_hardware"
sidebar_current: "docs-server_hardware"
description: |-
 Gets information about an existing server_hardware.
---

# oneview\_server_hardware

Use this data source to access the attributes of a Server Hardware.

## Example Usage

```hcl
data "oneview_server_hardware" "test" {
 name = "Test server hardware"
}

output "oneview_server_hardware_value" {
 value = "${data.oneview_server_hardware.test.uri}"
}
```

## Argument Reference

* `name` - (Required) For blade servers, it is the location based name of the server, which is formed by concatenating the enclosure name and the bay number. For rack servers, it is the serial number prefixed by word "ILO" e.g. ILOUSE31835LS

## Attributes Reference

* `power_state` - Current power state of the server hardware.

* `location_uri` - For blade servers, the enclosure in which this blade server resides. This URI can be used to retrieve information about the enclosure. This value is not set for rack mount servers.

* `mp_firmware_version` - The version of the firmware installed on the iLO.

* `uri` - The URI of the resource.

* `type` - Type of the resource.

* `server_group_uri` - For blade servers, this is the URI of the containing enclosure's enclosure group. This URI can be used to retrieve information about the enclosure group or to identify all the servers in the same group. This value is not set for rack mount servers.

* `server_hardware_type_uri` - URI of the server hardware type associated with the server hardware. This URI can be used to retrieve information about the server hardware type. 

* `server_profile_uri` - URI of a server profile assigned to this server hardware

* `uuid` - Universally Unique ID (UUID) of the server hardware.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.
