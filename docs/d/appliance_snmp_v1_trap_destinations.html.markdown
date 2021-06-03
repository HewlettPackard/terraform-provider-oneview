---
layout: "oneview"
page_title: "Oneview: Appliance SNMPv1 Trap Destinations"
sidebar_current: "docs-appliance_snmpv1_trap_destinations"
description: |-
 Gets information about an existing appliance_snmpv1_trap_destinations.
---

# oneview\_appliance\_snmp\_v1\_trap\_destinations

Use this data source to access the attributes of a appliance_snmp_v1_trap_destinations.

## Example Usage

```hcl
data "oneview_appliance_snmp_v1_trap_destinations" "test" {
 destination_id = "4"
}
output "oneview_snmp_v1_trap_destination_value" {
 value = "${data.oneview_snmp_v1_trap_destinations.test}"
}
```

## Argument Reference

* `destination_id` - (Required) The trap destination Id.

## Attributes Reference

* `community_string` - The community string associated with this trap destination.

* `destination_address` - The IP address or host name of the trap destination.

* `port` - The trap destination port.

* `uri` - The URI of the resource.

