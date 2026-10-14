---
layout: "oneview"
page_title: "Oneview: SNMPv3 Trap Destinations"
sidebar_current: "docs-snmpv3_trap_destinations"
description: |-
  Creates a snmpv3_trap_destinations.
---

# oneview\_snmpv3\_trap\_destinations

Creates a SNMPv3 Trap Destinations.

## Example Usage

```js
resource "oneview_appliance_snmpv3_trap_destinations" "snmptrap" {
    destination_address = "1.1.1.1"
    port = 162
    user_id = "41b96bbb-8f31-44e1-a3aa-8681e3d7c56c"
}
```

## Argument Reference

The following arguments are supported: 

* `destination_address` - The IP address or host name of the trap destination.

* `port` - The trap destination port.

* `user_id` - The user settings Id.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `category` - Identifies the resource type.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `user_uri` - The unique user setting uri.

* `id_field` - The trap destination Id.

* `trap_type` -  Trap notification type.

* `uri` - The URI of the resource.

* `type` - Type of the resource.