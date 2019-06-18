---
layout: "oneview"
page_title: "Oneview: server hardware type"
sidebar_current: "docs-server-hardware-type"
description: |-
  Gets information about an existing server hardware type.
---

# oneview\_server\_hardware\_type

Use this data source to access the attributes of an Server Hardware Type.

## Example Usage

```hcl
data "oneview_server_hardware_type" "test" {
  name = "Test Server Hardware Type"
}

output "oneview_server_hardware_type_value" {
  value = "${data.oneview_server_hardware_type.test.uri}"
}
```

## Argument Reference

* `name` - (Required) A unique name for the resource.

## Attributes Reference


* `category` - The category of the interconnect.

* `uri` - The URI of the created resource.

* `etag` - Entity tag/version ID of the resource.
