---
layout: "oneview"
page_title: "Oneview: enclosure_group"
sidebar_current: "docs-enclosure_group"
description: |-
 Gets information about an existing enclosure_group.
---

# oneview\_enclosure\_group

Use this data source to access the attributes of a enclsoure group.

## Example Usage

```hcl
data "oneview_enclosure_group" "test" {
 name = "Test server hardware"
}

output "oneview_enclosure_group_value" {
 value = "${data.oneview_enclosure_group.test.uri}"
}
```

## Argument Reference

* `name` - (Required) The name of the enclosure group.

## Attributes Reference

* `ambient_temperature_mode` - The ambient temperature mode describing the environment in which the enclosure group should be optimized to operate.

* `associated_logical_interconnect_groups` - A sorted list of logical interconnect group URIs associated with the enclosure group.

* `category` - Identifies the resource type.

* `description` - Brief description of the resource.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `enclosure_count` - The number of enclosures in the enclosure group.

* `enclosure_type_uri` - The type of enclosures the group contains.

* `interconnect_bay_mapping_count` - The number of interconnect bay mappings.

* `interconnect_bay_mappings` - Defines which logical interconnect group is associated with each interconnect bay in which enclosure.

* `ip_addressing_mode` - nP addressing mode, Not used in the C7000 enclosure.

* `ip_range_uris` - Range of IP URIs, not used in the C7000 enclosure.

* `port_mapping_count` - The number of port mappings.

* `port_mappings` - Provides midplane port number to IO bay mapping.

* `power_mode` - Power mode of the enclosure group.

* `scopes_uri` - The URI for the resource scope assignments.

* `stacking_mode` -  Stacking mode of the enclosure group. Currently only the Enclosure mode is supported..

* `status` - Overall health status of the enclosure group.

* `uri` - The URI of the resource.

* `type` - Type of the resource.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned. 
