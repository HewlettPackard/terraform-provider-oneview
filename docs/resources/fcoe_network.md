---
layout: "oneview"
page_title: "Oneview: fcoe_network"
sidebar_current: "docs-oneview-fcoe-network"
description: |-
  Creates an fcoe network.
---

# oneview\_fcoe\_network

Creates an fcoe network.

## Example Usage

```js
resource "oneview_fcoe_network" "default" {
  name = "test-fcoe-network"
  vlanId = 71
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the resource.

* `vlanId` - (Required) The Virtual LAN (VLAN) identification number (integer) assigned to the network.
Changing this forces a new resource

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be initially assigned.

* `connection_template_uri` - The connection template URI that is associated with this Fibre Channel network.

* `managed_san_uri` - The managed SAN URI that is associated with this Fibre Channel network.

* `fabric_uri` - The URI of the fabric resource of which this resource is a member.

* `description` - Brief description of the resource.

* `bandwidth` - The bandwidth assigned to the connection. 
  *  `maximum_bandwidth` - The maximum bandwidth of a network connection, expressed in Mbps.
  *  `typical_bandwidth` - The typical bandwidth of a network connection, expressed in Mbps.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.
