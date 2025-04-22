---
layout: "oneview"
page_title: "Oneview: SNMPv3 Trap Destinations"
sidebar_current: "docs-snmpv3_trap_destinations"
description: |-
 Gets information about an existing snmpv3_trap_destinations.
---

# oneview\_appliance\_snmpv3\_trap\_destinations

Use this data source to access the attributes of a snmpv3_trap_destinations.

## Example Usage

```hcl
data "oneview_appliance_snmpv3_trap_destinations" "test" {
 id_field = "67003649-af34-4a92-a46a-137855ddc8f7"
}

output "oneview_snmpv3_trap_destinations" {
 value = "${data.oneview_snmpv3_trap_destinations.test}"
}
```

## Argument Reference

* `id_field` - (Required) The trap destination Id.

## Attributes Reference

* `destination_address` - The IP address or host name of the trap destination.

* `engine_id` - Trap destination engine ID.

* `port` - The trap destination port.

* `category` - Identifies the resource type.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `user_id` - The user settings Id

* `user_uri` - The unique user setting uri

* `trap_type` -  Trap notification type.

* `uri` - The URI of the resource.

* `type` - Type of the resource.