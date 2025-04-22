---
layout: "oneview"
page_title: "Oneview: network_set"
sidebar_current: "docs-network_set"
description: |-
 Gets information about an existing network_set.
---

# oneview\_network\_set

Use this data source to access the attributes of a network set.

## Example Usage

```hcl
data "oneview_network_set" "test" {
 name = "Test network set"
}

output "oneview_network_set_value" {
 value = "${data.oneview_network_set.test.uri}"
}
```

## Argument Reference

* `name` - (Required) The name of the network set.

## Attributes Reference

* `network_uris` - (Optional) A set of Ethernet network URIs that will be members of this network set. 

* `native_network_uri` - (Optional) The URI of the network that will serve as the native network 

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.

* `connection_template_uri` - The URI of the existing connection template associated with this object.

* `scopesUri` - The URI for the resource scope assignments.

* `description` - Brief description of the resource.

* `network_set_type` - NetworkSet Type whose default value is "Regular".

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.
