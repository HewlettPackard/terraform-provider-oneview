---
layout: "oneview"
page_title: "Oneview: logical_interconnect_group"
sidebar_current: "docs-logical_interconnect_group"
description: |-
 Gets information about an existing logical_interconnect_group.
---

# oneview\_logical_interconnect_group

Use this data source to access the attributes of a Logical Interconnect Group.

## Example Usage

```hcl
data "oneview_logical_interconnect_group" "test" {
 name = "TestLIG"
}

output "oneview_logical_interconnect_group_value" {
 value = "${data.oneview_logical_interconnect_group.test.uri}"
}
```

## Argument Reference

* `name` - (Required) Name od teh Logical Interconnect Group.

## Attributes Reference

* `category` -  Identifies the resource type

* `description` - Brief description of the resource

* `eTag` - Entity tag/version ID of the resource.

* `enclosure_indexes` - The list of enclosure indices that are specified by this logical interconnect group. 
The value [-1] indicates that this is a single enclosure logical interconnect group for Virtual Connect SE FC Modules. 
The value [1] indicates that this is a single enclosure logical interconnect group for other supported interconnects. 
If you are building a logical interconnect group for use with a three enclosures interconnect link topology, 
the value needs to be [1,2,3].

* `ethernet_settings` - The Ethernet interconnect settings for the logical interconnect group.

* `fabric_uri` - The URI of the fabric resource of which this resource is a member.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.

* `interconnect_bay_set` - Interconnect bay associated with the logical interconnect group

* `interconnect_map_entry_template` - Interconnect map associated with the logical interconnect group

* `internal_network_uris` - A list of internal network URIs

* `quality_of_service` - The QOS configuration

* `redundancy_type` - The type of enclosure redundancy

* `snmp_configuration` -  The SNMP configuration for the logical interconnect group. Optional, if not supplied a default will be used

* `state` - Current state of the resource

* `telemetry_configuration` - The controls for collection of interconnect statistics. Optional, if not supplied a default will be used.  

* `type` - Uniquely identifies the type of the JSON object 

* `uplink_set` -  List of uplink sets in the logical interconnect group.

* `uri` - The URI of the created resource.
