---
layout: "oneview"
page_title: "Oneview: enclosure_group"
sidebar_current: "docs-oneview-enclosure-group"
description: |-
  Creates an enclosure-group.
---

# oneview\_enclosure\_group

Creates an enclosure group.

## Example Usage

```js
resource "oneview_enclosure_group" "default" {
  name = "default-enclosure-group"
  ip_addressing_mode = "External"
  enclosure_count = 3
  "interconnect_bay_mappings":
    [
        {
            "interconnect_bay": 1,
            "logical_interconnect_group_uri": "/rest/logical-interconnect-groups/aeef7314-527d-4053-868c-17b87df1b57c"
        },
	]
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

---

* `interconnect_bay_mappings` - (Optional) Defines which logical interconnect group is associated with each interconnect bay in which enclosure.

* `enclosure_count` - (Optional) The number of enclosures in the enclosure group.

* `ip_addressing_mode` - (Optional)  IPv4 address allocation for interconnects and device bay management processors. Used for Synergy only.

* `stacking_mode` - (Optional) Stacking mode of the enclosure group. Defaults to Enclosure.

* `interconnect_bay_mapping_count` - (Optional) The number of interconnect bay mappings. Defaults to 8.

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be initially assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.
