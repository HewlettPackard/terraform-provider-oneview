---
layout: "oneview"
page_title: "Oneview: labels"
sidebar_current: "docs-labels"
description: |-
 Gets information about an existing label.
---

# oneview\_labels

Use this data source to access the attributes of a label.

## Example Usage

```hcl
data "oneview_label" "test" {
 resource_uri = "/rest/resource/uri"
}

output "oneview_label_value" {
 value = "${data.oneview_label.test.labels}"
}
```

## Argument Reference

* `resource_uri` - (Required) The URI to the resource that has these labels.

## Argument Reference

In addition to the argument listed above, the following computed attributes are exported:

* `labels` - The block of label. Labels configuration is specified below.
  * `name` - The name of the label.
  * `uri` - A URI to the label.

* `category` -  Identifies the resource type.

* `created` -  Date and time when the resource was created.

* `etag` -  Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resourcee.

* `modified` - Date and time when the resource was last modified.

* `type` - Uniquely identifies the type of the JSON object.

* `uri` - The canonical URI of the resource.
