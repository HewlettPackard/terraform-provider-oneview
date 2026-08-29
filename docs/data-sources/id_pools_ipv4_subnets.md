---
layout: "oneview"
page_title: "Oneview: IPv4 Subnets"
sidebar_current: "docs-ipv4_subnets"
description: |-
 Gets information about an existing ipv4 subnet.
---

# oneview\_id_pools_ipv4_subnets

Use this data source to access the attributes of a Subnet.

## Example Usage

```hcl
data "oneview_id_pools_ipv4_subnets" "test" {
 subnet_id = "998gfjghf3254-jdfr844739"
}

output "oneview_subnet_value" {
 value = "${data.oneview_id_pools_ipv4_subnets.test.uri}"
}
```

## Argument Reference

* `subnet_id` - (Required) Id of the subnet whose details need to be retrieved.

## Attributes Reference

* `associated_resource` - Details of associated resource.

* `associated_task_uri` - URI of another task associated with it.

* `category` - Category of this resource.

* `collector_uri` - URI of the collector for the range.

* `dnsServers` - The list of DNS server IP addresses for IP range.

* `eTag` - Entity tag/version ID of the resource. 

* `gateway` - The gateway IP address for an IP range.

* `modified` - Date and time when the resource was last modified.

* `name` - Display name for the resource.

* `networkId` - The network ID for IP subnet.

* `rangeUris` - A list of range Uris.

* `subnetmask` - The subnet mask for an IP range.

* `type` - Type of the object .

* `uri` - URI of the resource.
