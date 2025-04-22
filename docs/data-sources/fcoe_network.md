---
layout: "oneview"
page_title: "Oneview: fcoe_network"
sidebar_current: "docs-fcoe_network"
description: |-
 Gets information about an existing fcoe_network.
---

# oneview\_fcoe\_network

Use this data source to access the attributes of a fcoe network.

## Example Usage

```hcl
data "oneview_fcoe_network" "test" {
 name = "Test FCoE network"
}

output "oneview_fcoe_network_value" {
 value = "${data.oneview_fcoe_network.test.uri}"
}
```

## Argument Reference

* `name` - (Required) The name of the fc network.

## Attributes Reference

* `vlan_id` - (Required) The Virtual LAN (VLAN) identification number (integer) assigned to the network.
Changing this forces a new resource.

* `fabric_uri` - The URI of the fabric resource of which this resource is a member.

* `category` - Identifies the resource type.

* `description` - Brief description of the resource.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `connection_template_uri` - The connection template URI that is associated with this Fibre Channel network.

* `managed_san_uri` - The managed SAN URI that is associated with this Fibre Channel network.

* `scopesUri` - The URI for the resource scope assignments..

* `status` - Overall health status of the resource.

* `uri` - The URI of the resource.

* `type` - Type of the resource.

* `bandwidth` - The bandwidth assigned to the connection. 
  *  `maximum_bandwidth` - The maximum bandwidth of a network connection, expressed in Mbps.
  *  `typical_bandwidth` - The typical bandwidth of a network connection, expressed in Mbps.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.
