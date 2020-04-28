---
layout: "oneview"
page_title: "Oneview: hypervisor_manager"
sidebar_current: "docs-hypervisor_manager"
description: |-
 Gets information about an existing hypervisor_manager.
---

# oneview\_hypervisor\_manager

Use this data source to access the attributes of a hypervisor manager.

## Example Usage

```hcl
data "oneview_hypervisor_manager" "test" {
 name = "Test hypervisor manager"
}

output "oneview_hypervisor_manager_value" {
 value = "${data.oneview_hypervisor_manager.test.uri}"
}
```

## Argument Reference

* `uri` - (Required) The name of the hypervisor manager.

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
