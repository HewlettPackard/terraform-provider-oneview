---
layout: "oneview"
page_title: "Oneview: network_set"
sidebar_current: "docs-oneview-network-set"
description: |-
  Creates a network set.
---

# oneview\_network\_set

Creates a network set.

## Example Usage
## Empty Network Set
```js
resource "oneview_network_set" "default" {
  name = "test-network-set"
}
```
## With networks 
```js
resource "oneview_network_set" "default" {
  name = "test-network-set"
  network_uris = ["${oneview_ethernet_network.default.*.uri}"]
  native_network_uri = "${oneview_ethernet_network.default.1.uri}"
}
```
## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

- - -

* `network_uris` - (Optional) A set of Ethernet network URIs that will be members of this network set. 
  NOTE: all Ethernet networks in a network set must have unique VLAN IDs.
  
* `native_network_uri` - (Optional) The URI of the network that will serve as the native network 
  in the network set. It must be in the set of network_uris.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.

* `connection_template_uri` - The URI of the existing connection template associated with this object.

* `scopesUri` - The URI for the resource scope assignments.

* `description` - Brief description of the resource.

* `network_set_type` - NetworkSet Type whose default value is "Regular".

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.
