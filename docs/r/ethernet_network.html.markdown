---
layout: "oneview"
page_title: "Oneview: ethernet_network"
sidebar_current: "docs-oneview-ethernet-network"
description: |-
  Creates an ethernet network.
---

# oneview\_ethernet\_network

Creates an ethernet network.

## Example Usage

```js
resource "oneview_ethernet_network" "default" {
  name = "test-ethernet-network"
  vlan_id = 71
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the resource.

* `vlan_id` - (Required) The Virtual LAN (VLAN) identification number (integer) assigned to the network.
Changing this forces a new resource.

- - -

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned. 
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.

* `purpose` - (Optional) A description of the network's role within the logical interconnect.
  This defaults to General.

* `private_network` - (Optional) When enabled, the network is configured so that all downlink (server) ports
  connected to the network are prevented from communicating with each other within the logical interconnect.
  This defaults to false.

* `smart_link` - (Optional) When enabled, the network is configured so that, within a logical interconnect,
  all uplinks that carry the network are monitored. This defaults to false.

* `ethernet_network_type` - (Optional) The type of Ethernet network. This defaults to Tagged.

* `bandwidth` - The bandwidth assigned to the connection. 
  *  `maximum_bandwidth` - The maximum bandwidth of a network connection, expressed in Mbps.
  *  `typical_bandwidth` - The typical bandwidth of a network connection, expressed in Mbps.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.
