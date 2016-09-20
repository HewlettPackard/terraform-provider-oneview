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
  server_template = "${oneview_server_profile_template.test.name}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

* `server_template` - (Required) The name of the template you will use for the server profile. 

- - -

* `public_connection` - (Optional) The name of the network that is going out to the public.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `serial_number` - A 10-byte value that is exposed to the Operating System as the server hardware's
Serial Number. The value can be a virtual serial number, user defined serial number or physical serial
number read from the server's ROM. It cannot be modified after the profile is created.

* `public_mac` - The MAC address of the NIC your public network is attached to.

* `public_slot_id` - The slot id of the NIC your public network is attached to.
