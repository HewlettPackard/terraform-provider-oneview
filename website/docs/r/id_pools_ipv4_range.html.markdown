---
layout: "oneview"
page_title: "Oneview: id_pools_ipv4_range"
sidebar_current: "docs-id_pools_ipv4_range"
description: |-
 Creates an ipv4 range.
---

# oneview\_id\_pools\_ipv4\_range

Creates the ipv4 range.

## Example Usage

```js
resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange"
  subnet_uri = "/rest/id-pools/ipv4/subnets/40f76df9-1e39-4e5a-81fc-14614efea5e8"
  start_stop_fragments {
    start_address = "<ip_address>"
    end_address = "<ip_address>"
  }
  start_stop_fragments {
    start_address = "<ip_address>"
    end_address = "<ip_address>"
  }
}
```

## Argument Reference

* `allocatedFragmentUri` - URI of allocated fragments for the range.

* `allocatedIdCount` - Count of IDs allocated from the range.

* `allocatorUri` - URI of the allocator for the range.

* `associatedResources` - A list of associated resources.

* `category` - Identifies the resource type.

* `collectorUri` - URI of the collector for the range.

* `created` - Date and time when the resource was created.

* `defaultRange` - Whether this is a default range.

* `eTag` - Entity tag/version ID of the resource.

* `enabled` - The status of the pool.

* `endAddress` - The end address of the range.

* `freeFragmentUri` - URI of free fragments for the range.

* `freeIdCount` - Count of IDs returned to the range.

* `modified` - Date and time when the resource was last modified.

* `name` - The name of the range.

* `prefix` - Prefix to be used in front of the generated IDs.

* `rangeCategory` - The category of the range.

* `reservedIdCount` - Count of IDs reserved in the range.

* `startAddress` - The start address of the range.

* `startStopFragments` - A list of start and end addresses of the range.

* `subnetUri` - The subnetUri of associated with an IP range.

* `totalCount` - Total count of IDs managed by the range.

* `type` - Uniquely identifies the type of the JSON object.

* `uri` - The canonical URI of the resource.
