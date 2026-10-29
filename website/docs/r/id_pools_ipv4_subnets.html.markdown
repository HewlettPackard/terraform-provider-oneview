---
layout: "oneview"
page_title: "Oneview: IPv4 Subnets"
sidebar_current: "docs-ipv4_subnets"
description: |-
 Creates a ipv4 subnet.
---

# oneview\_id_pools_ipv4_subnets

Creates a new Subnet.

## Example Usage

```hcl
resource "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  name="SubnetTF"
  network_id="<networkId>"
  subnet_mask="<subnetMask>"
  gateway="<gateway>"
  domain= "Terraform.com"
}

```

## Argument Reference

* `gateway` - (Required) The gateway IP address for an IP range.

* `network_id` - (Required) Id of the subnet whose details need to be retrieved.

* `subnetmask` - (Required) The subnet mask for an IP range.

* `name` - Display name for the resource.

* `domain` - The domain for an IP range.

## Attributes Reference

* `associated_resource` - Details of associated resource.

* `associated_task_uri` - URI of another task associated with it.

* `category` - Category of this resource.

* `collector_uri` - URI of the collector for the range.

* `dnsServers` - The list of DNS server IP addresses for IP range.

* `eTag` - Entity tag/version ID of the resource. 

* `modified` - Date and time when the resource was last modified.

* `rangeUris` - A list of range Uris.

* `type` - Type of the object.

* `uri` - URI of the resource.
