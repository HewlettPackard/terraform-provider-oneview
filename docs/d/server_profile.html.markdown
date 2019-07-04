---
layout: "oneview"
page_title: "Oneview: server_profile"
sidebar_current: "docs-oneview-server-profile"
description: |-
	Gets information about an existing server_profile.
---

# oneview\_server\_profile

Use this data source to access the attributes of a Server Profile.

## Example Usage

```js
data "oneview_server_profile" "sp" {
        name = "TestAll"
}

output "oneview_server_profile_value" {
        value = "${data.oneview_server_profile.sp.uri}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

- - -

* `public_connection` -  The name of the network that is going out to the public.

* `hardware_name` -  The name of the Server Hardware the server will be provisioned on.
  If this isn't used, a server hardware will be picked based on compatibility with the server profile template and any hw_filter(s) (see below).

* `type` -  The server profile version to be provisioned. Defaults to ServerProfileV5.
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
