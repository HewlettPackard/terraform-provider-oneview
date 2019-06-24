---
layout: "oneview"
page_title: "Oneview: interconnects"
sidebar_current: "docs-interconnects"
description: |-
  Gets information about an existing interconnects.
---

# oneview\_interconnects

Use this data source to access the attributes of an Interconnect.

## Example Usage

```hcl
data "oneview_interconnect" "test" {
  name = "Test Interconnect"
}

output "oneview_interconnect_value" {
  value = "${data.oneview_interconnect.test.uri}"
}
```

## Argument Reference

* `name` - (Required) A unique name for the resource.

## Attributes Reference


* `category` - The category of the interconnect.

* `uri` - The URI of the created resource.

* `type` - Type of the resource.
