---
layout: "oneview"
page_title: "Oneview: Uplink port relative value"
sidebar_current: "docs-relative_value"
description: |-
 Gets relative value of a given port name
---

# oneview\_relative value

Use this data source to get the relative value for a given port name

## Example Usage

```hcl
data "oneview_relative_value"  "rv"{
  port_name="Q2:1"
  interconnect_type_name="Virtual Connect SE 40Gb F8 Module for Synergy"
}

output "oneview_relative_value" {
 value = "${oneview_relative_value.rv.port_num}"
}
```

## Argument Reference

* `port_name` - (Required) The name of the enclsoure.
* `interconnect_type_name`- (Required)" The name of interconnect type

## Attributes Reference

* `port_name` - (Required) The name of the enclsoure.
* `interconnect_type_name`- (Required)" The name of interconnect type
* `port_num` - The relative value corresponding to the port name