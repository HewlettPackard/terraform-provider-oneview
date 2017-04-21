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

```js
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  template = "${oneview_server_profile_template.test.name}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

* `template` - (Required) The name of the template you will use for the server profile. 

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
