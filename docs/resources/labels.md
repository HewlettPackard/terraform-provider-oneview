---
layout: "oneview"
page_title: "Oneview: label"
sidebar_current: "docs-oneview-label"
description: |-
  Creates labels for a resource.
---

# oneview\_label

Create labels for a resource.

## Example Usage

```hcl
resource "oneview_label" "default" {
  resource_uri = "/rest/server-profile-templates/xyz"
  labels {
    name = "NewLabel"
  }
}
```

## Argument Reference

The following arguments are supported: 

* `resource_uri` - (Required) The URI to the resource which will have these labels.

* `labels` - The block of label. Labels configuration is specified below.
  * `name` - The name of the label.
  * `uri` - A URI to the label.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `category` -  Identifies the resource type.

* `created` -  Date and time when the resource was created.

* `etag` -  Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resourcee.

* `modified` - Date and time when the resource was last modified.

* `type` - Uniquely identifies the type of the JSON object.

* `uri` - The canonical URI of the resource.
