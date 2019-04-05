---
layout: "oneview"
page_title: "Oneview: logical-interconnects"
sidebar_current: "docs-logical-interconnects"
description: |-
  Gets information about an existing logical interconnects.
---

# oneview\_logical\_interconnects

Use this data source to access the attributes of an Logical Interconnect.

## Example Usage

```hcl
data "oneview_logical_interconnect" "test" {
  name = "Test Interconnect"
}

output "oneview_logical_interconnect_value" {
  value = "${data.oneview_logical_interconnect.test.uri}"
}
```

## Argument Reference

* `name` - (Required) A unique name for the resource.

## Attributes Reference


* `category` - The category of the logical interconnect.

* `uri` - The URI of the created resource.

* `type` - Type of the resource.
