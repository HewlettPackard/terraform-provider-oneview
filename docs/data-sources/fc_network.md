---
layout: "oneview"
page_title: "Oneview: fc_network"
sidebar_current: "docs-fc_network"
description: |-
 Gets information about an existing fc_network.
---

# oneview\_fc\_network

Use this data source to access the attributes of a fc network.

## Example Usage

```hcl
data "oneview_fc_network" "test" {
 name = "Test fc network"
}

output "oneview_fc_network_value" {
 value = "${data.oneview_fc_network.test.uri}"
}
```

## Argument Reference

* `name` - (Required) The name of the fc network.

## Attributes Reference

* `fabric_type` - The supported Fibre Channel access method.

* `link_stability_time` - The time interval, expressed in seconds, to wait after a link that was previously offline becomes stable, before automatic redistribution occurs within the fabric.

* `category` - Identifies the resource type.

* `description` - Brief description of the resource.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `auto_login_redistribution` - Used for load balancing when logins are not evenly distributed over the Fibre Channel links, such as when an uplink that was previously down becomes available.

* `connection_template_uri` - The connection template URI that is associated with this Fibre Channel network.

* `managed_san_uri` - The managed SAN URI that is associated with this Fibre Channel network.

* `scopesUri` - The URI for the resource scope assignments..

* `status` - Overall health status of the resource.

* `uri` - The URI of the resource.

* `type` - Type of the resource.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.

* `bandwidth` - The bandwidth assigned to the connection. 
  *  `maximum_bandwidth` - The maximum bandwidth of a network connection, expressed in Mbps.
  *  `typical_bandwidth` - The typical bandwidth of a network connection, expressed in Mbps.
