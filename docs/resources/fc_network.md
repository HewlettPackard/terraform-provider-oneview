---
layout: "oneview"
page_title: "Oneview: fc_network"
sidebar_current: "docs-oneview-fc-network"
description: |-
  Creates an fibre channel network.
---

# oneview\_fc\_network

Creates an fc network.

## Example Usage

```js
resource "oneview_fc_network" "default" {
  name = "test-fc-network"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

- - -

* `fabric_type` - (Optional) The supported Fibre Channel access method. 
  This defaults to FabricAttach.
  
* `link_stability_time` - (Optional) The time interval, expressed in seconds, to 
wait after a link that was previously offline becomes stable, before automatic redistribution occurs within the fabric. 
This value is not effective if autoLoginRedistribution is false.
This defaults to 30.

* `auto_login_redistribution` - (Optional) Used for load balancing when logins are not 
evenly distributed over the Fibre Channel links,such as when an uplink that was previously down becomes available. 
This defaults to false.

* `bandwidth` - The bandwidth assigned to the connection. 
  *  `maximum_bandwidth` - The maximum bandwidth of a network connection, expressed in Mbps.
  *  `typical_bandwidth` - The typical bandwidth of a network connection, expressed in Mbps.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.
