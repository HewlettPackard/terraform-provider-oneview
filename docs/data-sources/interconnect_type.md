---
layout: "oneview"
page_title: "Oneview: interconnect_type"
sidebar_current: "docs-interconnect_type"
description: |-
 Gets information about an existing interconnect type.
---

# oneview\_interconnect_type

Use this data source to access the attributes of a Server Hardware.

## Example Usage

```hcl
data "oneview_interconnect_type" "test" {
 name = "Test interconnect type"
}

output "oneview_interconnect_type_value" {
 value = "${data.oneview_interconnect_type.test.uri}"
}
```

## Argument Reference

* `name` - (Required) A unique name for the resource

## Attributes Reference

* `category` - The category of the interconnect.

* `uri` - The URI of the created resource.

* `etag` - Entity tag/version ID of the resource.

* `port_info` - port capabilities for this interconnect type

* `downlink_port_capability` - The downlink port capabilities for this interconnect type 

* `state` - Current state of the resource
