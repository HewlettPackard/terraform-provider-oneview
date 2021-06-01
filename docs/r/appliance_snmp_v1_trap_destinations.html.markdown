---
layout: "oneview"
page_title: "Oneview: Appliance SNMPv1 Trap Destinations"
sidebar_current: "docs-appliance_snmpv1_trap_destinations"
description: |-
 Creates a  appliance_snmp_v1_trap_destinations.
---

# oneview\_appliance\_snmp\_v1\_trap\_destinations

Creates a snmp_v1_trap_destinations.

## Example Usage

```hcl
resource "oneview_appliance_snmp_v1_trap_destinations" "test" {
 destination_id = "4"
 community_string = "Test1"
 destination_address = "1.1.1.1"
 port = 170
}
```

## Argument Reference

* `destination_id` - (Required) The trap destination Id.

* `community_string` - The community string associated with this trap destination.

* `destination_address` - The IP address or host name of the trap destination.

* `port` - The trap destination port.

## Attributes Reference

* `uri` - The URI of the resource.

